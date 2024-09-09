// Package space provides implementation of the Space resource.
package space

import (
	"context"

	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type spaceResource struct {
	client *api.APIClient
}

// Ensure SpaceResource implements the tfsdk.Resource interface.
var (
	_ resource.Resource                = &spaceResource{}
	_ resource.ResourceWithConfigure   = &spaceResource{}
	_ resource.ResourceWithImportState = &spaceResource{}
)

type spaceResourceModel struct {
	ID            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	Domain        types.String `tfsdk:"domain"`
	CreateOptions types.Object `tfsdk:"create_options"`
}

type spaceCreateOptionsModel struct {
	CopySSOSettings               types.Bool `tfsdk:"copy_sso_settings"`
	CopyBrandingAndThemesSettings types.Bool `tfsdk:"copy_branding_and_themes_settings"`
	UsersToCopyAsAdmins           types.List `tfsdk:"users_to_copy_as_admins"`
	CreateAdminUser               types.Bool `tfsdk:"create_admin_user"`
}

func getDefaultCreateOptions(diags *diag.Diagnostics) *basetypes.ObjectValue {
	emptyList, localDiags := types.ListValue(types.StringType, []attr.Value{})
	diags.Append(localDiags...)
	if diags.HasError() {
		return nil
	}

	defaultCreateOptionsTypes := map[string]attr.Type{
		"copy_sso_settings":                 types.BoolType,
		"copy_branding_and_themes_settings": types.BoolType,
		"users_to_copy_as_admins":           types.ListType{types.StringType},
		"create_admin_user":                 types.BoolType,
	}
	defaultCreateOptionsValues := map[string]attr.Value{
		"copy_sso_settings":                 types.BoolValue(false),
		"copy_branding_and_themes_settings": types.BoolValue(false),
		"users_to_copy_as_admins":           emptyList,
		"create_admin_user":                 types.BoolValue(true),
	}
	defaultCreateOptions, localDiags := types.ObjectValue(defaultCreateOptionsTypes, defaultCreateOptionsValues)
	diags.Append(localDiags...)
	if diags.HasError() {
		return nil
	}
	return &defaultCreateOptions
}

// Create new Space resource.
func NewResource() resource.Resource {
	return &spaceResource{}
}

func (r *spaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_space"
}

func (r *spaceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	emptyList, diags := types.ListValue(types.StringType, []attr.Value{})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	defaultCreateOptions := getDefaultCreateOptions(&resp.Diagnostics)
	if resp.Diagnostics.HasError() || defaultCreateOptions == nil {
		return
	}

	resp.Schema = schema.Schema{
		Description: "Space resource allows you to create and manage Spaces in Retool. The provider must be configured using the hostname of the admin Space.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the space.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the space.",
			},
			"domain": schema.StringAttribute{
				Required:    true,
				Description: "The domain of the space. Use full domain for self-hosted Retool, subdomain for Retool Cloud.",
			},
			"create_options": schema.SingleNestedAttribute{
				Description: "Options for creating the space.",
				Computed:    true,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"copy_sso_settings": schema.BoolAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Copy SSO settings from the admin Space.",
						Default:     booldefault.StaticBool(false),
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"copy_branding_and_themes_settings": schema.BoolAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Copy branding and themes settings from the admin Space.",
						Default:     booldefault.StaticBool(false),
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"users_to_copy_as_admins": schema.ListAttribute{
						Optional:    true,
						Computed:    true,
						Description: "List of emails of users from the admin space that need to be added to the new space as admins.",
						ElementType: types.StringType,
						Default:     listdefault.StaticValue(emptyList),
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"create_admin_user": schema.BoolAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Create an admin user in the new space for the creator instead of just sending out an invite.",
						Default:     booldefault.StaticBool(true),
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
				Default: objectdefault.StaticValue(*defaultCreateOptions),
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
					objectplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *spaceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*utils.ProviderData)
	if !ok {
		resp.Diagnostics.AddError("Unexpected Resource Configure Type", "Expected *utils.ProviderData, got: %T. Please report this issue to the provider developers.")
		return
	}
	r.client = providerData.Client
}

func (r *spaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan spaceResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var createOptions spaceCreateOptionsModel
	diags = plan.CreateOptions.As(ctx, &createOptions, basetypes.ObjectAsOptions{})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Creating Space", map[string]interface{}{"name": plan.Name, "domain": plan.Domain})

	usersToCopyStr := utils.GetStringListValue(ctx, createOptions.UsersToCopyAsAdmins, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	request := api.SpacesPostRequest{
		Name:   plan.Name.ValueString(),
		Domain: plan.Domain.ValueString(),
		Options: &api.SpacesPostRequestOptions{
			CopySsoSettings:               createOptions.CopySSOSettings.ValueBoolPointer(),
			CopyBrandingAndThemesSettings: createOptions.CopyBrandingAndThemesSettings.ValueBoolPointer(),
			UsersToCopyAsAdmins:           usersToCopyStr,
			CreateAdminUser:               createOptions.CreateAdminUser.ValueBoolPointer(),
		},
	}
	response, httpResponse, err := r.client.SpacesAPI.SpacesPost(ctx).SpacesPostRequest(request).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Space",
			"Could not create Space, unexpected error: "+err.Error(),
		)
		tflog.Error(ctx, "Error creating Space", utils.AddHTTPStatusCode(map[string]interface{}{"error": err.Error()}, httpResponse))
		return
	}

	plan.ID = types.StringValue(response.Data.Id)
	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Space created", map[string]interface{}{"id": plan.ID})
}

func (r *spaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state spaceResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	spaceID := state.ID.ValueString()
	tflog.Info(ctx, "Reading Space", map[string]interface{}{"id": spaceID})
	response, httpResponse, err := r.client.SpacesAPI.SpacesSpaceIdGet(ctx, spaceID).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			tflog.Info(ctx, "Space not found", map[string]any{"id": spaceID})
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Error reading Space",
			"Could not read Space, unexpected error: "+err.Error(),
		)
		tflog.Error(ctx, "Error reading Space", utils.AddHTTPStatusCode(map[string]interface{}{"error": err.Error(), "id": spaceID}, httpResponse))
		return
	}
	state.Name = types.StringValue(response.Data.Name)
	state.Domain = types.StringValue(response.Data.Domain)
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Space read", map[string]interface{}{"id": spaceID})
}

func (r *spaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state spaceResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plan spaceResourceModel
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	spaceID := state.ID.ValueString()
	tflog.Info(ctx, "Updating Space", map[string]interface{}{"id": spaceID, "name": plan.Name.ValueString(), "domain": plan.Domain.ValueString()})

	request := api.SpacesSpaceIdPutRequest{
		Name:   plan.Name.ValueString(),
		Domain: plan.Domain.ValueString(),
	}
	_, httpResponse, err := r.client.SpacesAPI.SpacesSpaceIdPut(ctx, spaceID).SpacesSpaceIdPutRequest(request).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Space",
			"Could not update Space with id "+spaceID+", unexpected error: "+err.Error(),
		)
		tflog.Error(ctx, "Error updating Space", utils.AddHTTPStatusCode(map[string]interface{}{"error": err.Error(), "id": spaceID}, httpResponse))
		return
	}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Space updated", map[string]interface{}{"id": spaceID})
}

func (r *spaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state spaceResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	spaceID := state.ID.ValueString()
	tflog.Info(ctx, "Deleting Space", map[string]interface{}{"id": spaceID})
	httpResponse, err := r.client.SpacesAPI.SpacesSpaceIdDelete(ctx, spaceID).Execute()
	if err != nil && !(httpResponse != nil && httpResponse.StatusCode == 404) {
		resp.Diagnostics.AddError(
			"Error deleting Space",
			"Could not delete Space with id "+spaceID+", unexpected error: "+err.Error(),
		)
		tflog.Error(ctx, "Error deleting Space", utils.AddHTTPStatusCode(map[string]interface{}{"error": err.Error(), "id": spaceID}, httpResponse))
		return
	}
}

// ImportState allows importing of a Space resource.
func (r *spaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute.
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	// Need to set default create options so that TF doesn't attempt to re-create the space.
	defaultCreateOptions := getDefaultCreateOptions(&resp.Diagnostics)
	if resp.Diagnostics.HasError() || defaultCreateOptions == nil {
		return
	}
	resp.State.SetAttribute(ctx, path.Root("create_options"), *defaultCreateOptions)
}
