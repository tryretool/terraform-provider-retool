package group

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// Ensure GroupResource implements the tfsdk.Resource interface
var (
	_ resource.Resource                = &groupResource{}
	_ resource.ResourceWithConfigure   = &groupResource{}
	_ resource.ResourceWithImportState = &groupResource{}
)

// groupResource schema structure
type groupResource struct {
	client *api.APIClient
}

// groupResourceModel defines the data model for the Group resource
type groupResourceModel struct {
	Id                          types.String `tfsdk:"id"`
	LegacyId                    types.String `tfsdk:"legacy_id"`
	Name                        types.String `tfsdk:"name"`
	UniversalAppAccess          types.String `tfsdk:"universal_app_access"`
	UniversalResourceAccess     types.String `tfsdk:"universal_resource_access"`
	UniversalWorkflowAccess     types.String `tfsdk:"universal_workflow_access"`
	UniversalQueryLibraryAccess types.String `tfsdk:"universal_query_library_access"`
	UserListAccess              types.Bool   `tfsdk:"user_list_access"`
	AuditLogAccess              types.Bool   `tfsdk:"audit_log_access"`
	UnpublishedReleaseAccess    types.Bool   `tfsdk:"unpublished_release_access"`
	UsageAnalyticsAccess        types.Bool   `tfsdk:"usage_analytics_access"`
	AccountDetailsAccess        types.Bool   `tfsdk:"account_details_access"`
	LandingPageAppId            types.String `tfsdk:"landing_page_app_id"`
}

func NewResource() resource.Resource {
	return &groupResource{}
}

// Configure adds the provider configured client to the resource.
func (r *groupResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
}

// Metadata associated with the Group resource
func (r *groupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

// Schema returns the schema for the Group resource
func (r *groupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the group. Currently this is the same legacy_id but will change in the future.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"legacy_id": schema.StringAttribute{
				Computed:    true,
				Description: "The legacy ID of the group.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the group.",
			},
			"universal_app_access": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("none"),
				Description: "The universal app access level for the group. This denotes the access level that this group has for all apps.",
				Validators: []validator.String{
					stringvalidator.OneOf("none", "use", "edit", "own"),
				},
			},
			"universal_resource_access": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("none"),
				Description: "The universal resource access level for the group. This denotes the access level that this group has for all resources.",
				Validators: []validator.String{
					stringvalidator.OneOf("none", "use", "edit", "own"),
				},
			},
			"universal_workflow_access": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("none"),
				Description: "The universal workflow access level for the group. This denotes the access level that this group has for all workflows.",
				Validators: []validator.String{
					stringvalidator.OneOf("none", "use", "edit", "own"),
				},
			},
			"universal_query_library_access": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("none"),
				Description: "Level of access that the group has to the Query Library.",
				Validators: []validator.String{
					stringvalidator.OneOf("none", "use", "edit"),
				},
			},
			"user_list_access": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Whether the group has access to the user list.",
			},
			"audit_log_access": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Whether the group has access to the audit log.",
			},
			"unpublished_release_access": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Whether the group has access to unpublished releases.",
			},
			"usage_analytics_access": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Whether the group has access to usage analytics.",
			},
			"account_details_access": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Whether the group has access to account details.",
			},
			"landing_page_app_id": schema.StringAttribute{
				Optional:    true,
				Description: "The app ID of the landing page.",
			},
		},
	}
}

// Create creates the group resource and sets the initial Terraform state.
func (r *groupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan groupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	var group api.GroupsPostRequest
	group.Name = plan.Name.ValueString()
	group.UniversalAppAccess = plan.UniversalAppAccess.ValueStringPointer()
	group.UniversalResourceAccess = plan.UniversalResourceAccess.ValueStringPointer()
	group.UniversalWorkflowAccess = plan.UniversalWorkflowAccess.ValueStringPointer()
	group.UniversalQueryLibraryAccess = plan.UniversalQueryLibraryAccess.ValueStringPointer()
	group.UserListAccess = plan.UserListAccess.ValueBoolPointer()
	group.AuditLogAccess = plan.AuditLogAccess.ValueBoolPointer()
	group.UnpublishedReleaseAccess = plan.UnpublishedReleaseAccess.ValueBoolPointer()
	group.UsageAnalyticsAccess = plan.UsageAnalyticsAccess.ValueBoolPointer()
	group.AccountDetailsAccess = plan.AccountDetailsAccess.ValueBoolPointer()
	if !plan.LandingPageAppId.IsNull() && !plan.LandingPageAppId.IsUnknown() {
		group.LandingPageAppId.Set(plan.LandingPageAppId.ValueStringPointer())
	}

	tflog.Info(ctx, "Creating a group", map[string]interface{}{"name": plan.Name.ValueString()})

	// Create new group
	response, httpResponse, err := r.client.GroupsAPI.GroupsPost(ctx).GroupsPostRequest(group).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating group",
			"Could not create group, unexpected error: "+err.Error(),
		)
		tflog.Error(ctx, "Error creating group", utils.AddHttpStatusCode(map[string]interface{}{"error": err.Error()}, httpResponse))
		return
	}

	if response.Data.Id.Get() == nil || response.Data.LegacyId.Get() == nil {
		resp.Diagnostics.AddError(
			"Error creating group",
			"Could not create group, unexpected error: missing ID or LegacyID",
		)
		tflog.Error(ctx, "Error creating group", map[string]interface{}{"error": "missing ID or LegacyID"})
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.Id = types.StringValue(utils.Float32PtrToIntString(response.Data.Id.Get()))
	plan.LegacyId = types.StringValue(utils.Float32PtrToIntString(response.Data.LegacyId.Get()))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error creating group", map[string]interface{}{"error": "Could not set state"})
		return
	}

	tflog.Info(ctx, "Group created", map[string]interface{}{"id": response.Data.Id.Get(), "legacyId": response.Data.LegacyId.Get()})
}

// Read a Group resource
func (r *groupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state groupResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Use the ID from the state to read the group
	groupID := state.Id.ValueString()
	group, httpResponse, err := r.client.GroupsAPI.GroupsGroupIdGet(ctx, groupID).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			tflog.Info(ctx, "Group not found", map[string]any{"groupID": groupID})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error reading group",
			fmt.Sprintf("Could not read group with ID %s: %s", groupID, err.Error()),
		)
		tflog.Error(ctx, "Error reading group", utils.AddHttpStatusCode(map[string]any{"groupID": groupID, "error": err.Error()}, httpResponse))
		return
	}

	// Map the API response to the Terraform state
	state.Id = types.StringValue(utils.Float32PtrToIntString(group.Data.Id.Get()))
	state.LegacyId = types.StringValue(utils.Float32PtrToIntString(group.Data.LegacyId.Get()))
	state.Name = types.StringValue(group.Data.Name)
	state.UniversalAppAccess = types.StringValue(group.Data.UniversalAppAccess)
	state.UniversalResourceAccess = types.StringValue(group.Data.UniversalResourceAccess)
	state.UniversalWorkflowAccess = types.StringValue(group.Data.UniversalWorkflowAccess)
	state.UniversalQueryLibraryAccess = types.StringValue(group.Data.UniversalQueryLibraryAccess)
	state.UserListAccess = types.BoolValue(group.Data.UserListAccess)
	state.AuditLogAccess = types.BoolValue(group.Data.AuditLogAccess)
	state.UnpublishedReleaseAccess = types.BoolValue(group.Data.UnpublishedReleaseAccess)
	state.UsageAnalyticsAccess = types.BoolValue(group.Data.UsageAnalyticsAccess)
	state.AccountDetailsAccess = types.BoolValue(group.Data.AccountDetailsAccess)
	state.LandingPageAppId = types.StringPointerValue(group.Data.LandingPageAppId.Get())

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update a Group resource
func (r *groupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan groupResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state groupResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// First we need to get the group by ID to get the current values of the members and user_invites fields, since we don't have those in the plan.
	groupID := state.Id.ValueString()
	group, httpResponse, err := r.client.GroupsAPI.GroupsGroupIdGet(ctx, groupID).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading group",
			fmt.Sprintf("Could not read group with ID %s: %s", groupID, err.Error()),
		)
		tflog.Error(ctx, "Error reading group", utils.AddHttpStatusCode(map[string]any{"groupID": groupID, "error": err.Error()}, httpResponse))
		return
	}

	// Prepare the update payload
	updatePayload := api.GroupsGroupIdPutRequest{
		Name:                        plan.Name.ValueStringPointer(),
		UniversalAppAccess:          plan.UniversalAppAccess.ValueStringPointer(),
		UniversalResourceAccess:     plan.UniversalResourceAccess.ValueStringPointer(),
		UniversalWorkflowAccess:     plan.UniversalWorkflowAccess.ValueStringPointer(),
		UniversalQueryLibraryAccess: plan.UniversalQueryLibraryAccess.ValueStringPointer(),
		UserListAccess:              plan.UserListAccess.ValueBoolPointer(),
		AuditLogAccess:              plan.AuditLogAccess.ValueBoolPointer(),
		UnpublishedReleaseAccess:    plan.UnpublishedReleaseAccess.ValueBoolPointer(),
		UsageAnalyticsAccess:        plan.UsageAnalyticsAccess.ValueBoolPointer(),
		AccountDetailsAccess:        plan.AccountDetailsAccess.ValueBoolPointer(),
		Members:                     convertGroupMembersToPutRequestType(group.Data.Members),
		UserInvites:                 group.Data.UserInvites,
	}
	updatePayload.LandingPageAppId.Set(plan.LandingPageAppId.ValueStringPointer())

	// Perform the update operation
	_, httpResponse, err = r.client.GroupsAPI.GroupsGroupIdPut(ctx, groupID).GroupsGroupIdPutRequest(updatePayload).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating group",
			fmt.Sprintf("Could not update group with ID %s: %s", groupID, err.Error()),
		)
		tflog.Error(ctx, "Error updating group", utils.AddHttpStatusCode(map[string]any{"groupID": groupID, "error": err.Error()}, httpResponse))
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func convertGroupMembersToPutRequestType(responseMembers []api.GroupsGroupIdGet200ResponseDataMembersInner) []api.GroupsGroupIdPutRequestMembersInner {
	members := make([]api.GroupsGroupIdPutRequestMembersInner, len(responseMembers))
	for i, m := range responseMembers {
		members[i] = api.GroupsGroupIdPutRequestMembersInner{
			Id:           m.Id,
			IsGroupAdmin: m.IsGroupAdmin,
		}
	}
	return members
}

// Delete a Group resource
func (r *groupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state groupResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	groupId := state.Id.ValueString()
	httpResponse, err := r.client.GroupsAPI.GroupsGroupIdDelete(ctx, groupId).Execute()
	if err != nil && !(httpResponse != nil && httpResponse.StatusCode == 404) { // it's ok to not find the group when deleting
		resp.Diagnostics.AddError(
			"Error Deleting Group",
			"Could not delete group with ID "+groupId+": "+err.Error(),
		)
		tflog.Error(ctx, "Error Deleting Group", utils.AddHttpStatusCode(map[string]any{"error": err.Error(), "groupId": groupId}, httpResponse))
		return
	}
}

// ImportState allows importing of a Group resource
func (r *groupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
