package folder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource                = &folderResource{}
	_ resource.ResourceWithConfigure   = &folderResource{}
	_ resource.ResourceWithImportState = &folderResource{}
)

// NewResource is a helper function to simplify the provider implementation.
func NewResource() resource.Resource {
	return &folderResource{}
}

// folderResource is the resource implementation.
type folderResource struct {
	client            *api.APIClient
	rootFolderIdCache *map[string]string
}

// Configure adds the provider configured client to the resource.
func (r *folderResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*utils.ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			"Expected *utils.ProviderData, got: %T. Please report this issue to the provider developers.",
		)
	}
	r.client = providerData.Client
	r.rootFolderIdCache = providerData.RootFolderIdCache
}

// Metadata returns the resource type name.
func (r *folderResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_folder"
}

// Schema defines the schema for the resource.
func (r *folderResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The id of the folder.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"legacy_id": schema.StringAttribute{
				Computed:    true,
				Description: "The legacy id of the folder.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the folder.",
			},
			"parent_folder_id": schema.StringAttribute{
				Computed:    true,
				Optional:    true,
				Description: "The id of the parent folder.",
				Default:     stringdefault.StaticString(ROOT_FOLDER_ID),
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"is_system_folder": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether the folder is a system folder.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"folder_type": schema.StringAttribute{
				Required:    true,
				Description: "The type of the folder: (app|file|resource|workflow).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(), // Changing the folder type requires replacing the resource
				},
				Validators: []validator.String{
					stringvalidator.OneOf("app", "file", "resource", "workflow"),
				},
			},
		},
	}
}

func (r *folderResource) getTrueParentFolderId(ctx context.Context, folderType, parentFolderId string, diags *diag.Diagnostics) string {
	if parentFolderId == ROOT_FOLDER_ID {
		// Get root folder ID
		rootFolderId, err := getRootFolderId(ctx, folderType, r.client, r.rootFolderIdCache)
		if err != nil {
			diags.AddError(
				"Error converting root folder id to the actual id",
				"Could not find root folder for type "+folderType+": "+err.Error(),
			)
			return parentFolderId
		}
		return rootFolderId
	}
	return parentFolderId
}

// Create creates the resource and sets the initial Terraform state.
func (r *folderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan folderModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	var folder api.FoldersPostRequest
	parentFolderId := r.getTrueParentFolderId(ctx, plan.FolderType.ValueString(), plan.ParentFolderId.ValueString(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	folder.Name = plan.Name.ValueString()
	folder.ParentFolderId.Set(&parentFolderId)
	folder.FolderType = plan.FolderType.ValueString()

	tflog.Info(ctx, "Creating a folder", map[string]any{"name": plan.Name.ValueString(), "folderType": plan.FolderType.ValueString(), "parentFolderId": parentFolderId})
	// Create new folder
	response, httpResponse, err := r.client.FoldersAPI.FoldersPost(ctx).FoldersPostRequest(folder).Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating folder",
			"Could not create folder, unexpected error: "+err.Error(),
		)
		tflog.Error(ctx, "Error creating folder", utils.AddHttpStatusCode(map[string]any{"error": err.Error()}, httpResponse))
		return
	}
	tflog.Info(ctx, "Folder created", map[string]any{"id": response.Data.Id, "legacyId": response.Data.LegacyId, "isSystemFolder": response.Data.IsSystemFolder})

	// Map response body to schema and populate Computed attribute values
	plan.Id = types.StringValue(response.Data.Id)
	plan.LegacyId = types.StringValue(response.Data.LegacyId)
	plan.IsSystemFolder = types.BoolValue(response.Data.IsSystemFolder)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *folderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state folderModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed folder value from Retool API
	tflog.Info(ctx, "Reading folder", map[string]any{"id": state.Id})
	response, httpResponse, err := r.client.FoldersAPI.FoldersFolderIdGet(ctx, state.Id.ValueString()).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			tflog.Info(ctx, "Folder not found", map[string]any{"id": state.Id})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error Reading Folder",
			"Could not read folder ID "+state.Id.ValueString()+": "+err.Error(),
		)
		tflog.Error(ctx, "Error Reading Folder", utils.AddHttpStatusCode(map[string]any{"error": err.Error()}, httpResponse))
		return
	}

	// Update the state
	state.LegacyId = types.StringValue(response.Data.LegacyId)
	state.Name = types.StringValue(response.Data.Name)
	// To keep the state consistent, we need to convert the root folder ID to the string constant ROOT_FOLDER_ID
	state.ParentFolderId = types.StringPointerValue(maybeReplaceRootFolderIdWithConstant(ctx, response.Data.FolderType, response.Data.ParentFolderId.Get(), r.client, r.rootFolderIdCache, &resp.Diagnostics))
	state.IsSystemFolder = types.BoolValue(response.Data.IsSystemFolder)
	state.FolderType = types.StringValue(response.Data.FolderType)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *folderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state folderModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plan folderModel
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if plan.Name.Equal(state.Name) && plan.ParentFolderId.Equal(state.ParentFolderId) {
		// No changes
		tflog.Info(ctx, "No changes detected for folder", map[string]any{"id": state.Id})
		return
	}

	// Generate API request body
	patchReq := api.FoldersFolderIdPatchRequest{}
	if !plan.Name.Equal(state.Name) {
		op := api.NewUsersUserIdPatchRequestOperationsInnerAnyOf("replace", "/name")
		op.Value = plan.Name.ValueString()
		patchReq.Operations = append(patchReq.Operations, api.FoldersFolderIdPatchRequestOperationsInner{UsersUserIdPatchRequestOperationsInnerAnyOf: op})
	}
	if !plan.ParentFolderId.Equal(state.ParentFolderId) {
		parentFolderId := r.getTrueParentFolderId(ctx, plan.FolderType.ValueString(), plan.ParentFolderId.ValueString(), &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
		op := api.NewUsersUserIdPatchRequestOperationsInnerAnyOf("replace", "/parent_folder_id")
		op.Value = parentFolderId
		patchReq.Operations = append(patchReq.Operations, api.FoldersFolderIdPatchRequestOperationsInner{UsersUserIdPatchRequestOperationsInnerAnyOf: op})
	}
	_, httpResponse, err := r.client.FoldersAPI.FoldersFolderIdPatch(ctx, state.Id.ValueString()).FoldersFolderIdPatchRequest(patchReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Folder",
			"Could not update folder "+state.Id.ValueString()+", unexpected error: "+err.Error(),
		)
		tflog.Error(ctx, "Error Updating Folder", utils.AddHttpStatusCode(map[string]any{"error": err.Error()}, httpResponse))
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *folderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state folderModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing order
	recursive := true
	deleteRequest := api.FoldersFolderIdDeleteRequest{}
	deleteRequest.Recursive = &recursive

	httpResponse, err := r.client.FoldersAPI.FoldersFolderIdDelete(ctx, state.Id.ValueString()).FoldersFolderIdDeleteRequest(deleteRequest).Execute()
	if err != nil && !(httpResponse != nil && httpResponse.StatusCode == 404) { // it's ok to not find the resource being deleted
		resp.Diagnostics.AddError(
			"Error Deleting Folder",
			"Could not delete folder"+state.Id.ValueString()+", unexpected error: "+err.Error(),
		)
		tflog.Error(ctx, "Error Deleting Folder", utils.AddHttpStatusCode(map[string]any{"error": err.Error()}, httpResponse))
		return
	}
}

func (r *folderResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
