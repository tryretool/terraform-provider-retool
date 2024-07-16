package folder

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
	client *api.APIClient
}

// Configure adds the provider configured client to the resource.
func (r *folderResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*api.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			"Expected *api.APIClient, got: %T. Please report this issue to the provider developers.",
		)
	}
	r.client = client
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

func (r *folderResource) getRootFolderId(ctx context.Context, folderType string) (string, error) {
	tflog.Info(ctx, "Getting root folder ID", map[string]any{"folderType": folderType})
	response, httpResponse, err := r.client.FoldersAPI.FoldersGet(ctx).Execute()
	if err != nil {
		tflog.Error(ctx, "Error getting root folder ID", utils.AddHttpStatusCode(map[string]any{"error": err.Error()}, httpResponse))
		return "", err
	}

	for _, folder := range response.Data {
		if folder.FolderType == folderType && folder.ParentFolderId.Get() == nil {
			return folder.Id, nil
		}
	}
	tflog.Error(ctx, "Root folder not found", map[string]any{"folderType": folderType})

	return "", fmt.Errorf("root folder not found for type %s", folderType)
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
	var parentFolderId string // We can't really create a folder without a parent folder ID, so we need to get the root folder id if it's not set

	if plan.ParentFolderId.IsNull() || plan.ParentFolderId.IsUnknown() {
		// Get root folder ID
		rootFolderId, err := r.getRootFolderId(ctx, plan.FolderType.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error creating folder",
				"Could not find root folder for type "+plan.FolderType.ValueString()+": "+err.Error(),
			)
			return
		}
		parentFolderId = rootFolderId
	} else {
		parentFolderId = plan.ParentFolderId.ValueString()
	}

	folder.Name = plan.Name.ValueString()
	folder.ParentFolderId.Set(&parentFolderId)
	folder.FolderType = plan.FolderType.ValueString()

	tflog.Info(ctx, "Creating a folder", map[string]any{"name": plan.Name, "folderType": plan.FolderType, "parentFolderId": plan.ParentFolderId})
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
	plan.ParentFolderId = types.StringPointerValue(response.Data.ParentFolderId.Get())

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
	state.ParentFolderId = types.StringPointerValue(response.Data.ParentFolderId.Get())
	state.IsSystemFolder = types.BoolValue(response.Data.IsSystemFolder)
	state.FolderType = types.StringValue(response.Data.FolderType)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *folderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// TODO(dzmitry.kishylau): Implement the Update method when we have correspnding API endpoint.
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
	httpResponse, err := r.client.FoldersAPI.FoldersFolderIdDelete(ctx, state.Id.ValueString()).Execute()
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
