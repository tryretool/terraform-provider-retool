// Package sourcecontrolsettings provides the implementation of the Source Control Settings resource.
package sourcecontrolsettings

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// Ensure SCM settings implements the tfsdk.Resource interface.
var (
	_ resource.Resource              = &scmSettingsResource{}
	_ resource.ResourceWithConfigure = &scmSettingsResource{}
)

type scmSettingsResource struct {
	client *api.APIClient
}

type scmSettingsModel struct {
	AutoBranchNamingEnabled          types.Bool   `tfsdk:"auto_branch_naming_enabled"`
	CustomPullRequestTemplateEnabled types.Bool   `tfsdk:"custom_pull_request_template_enabled"`
	CustomPullRequestTemplate        types.String `tfsdk:"custom_pull_request_template"`
	VersionControlLocked             types.Bool   `tfsdk:"version_control_locked"`
}

func NewResource() resource.Resource {
	return &scmSettingsResource{}
}

// Configure adds the provider configured client to the resource.
func (r *scmSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Metadata associated with the Source Control Settings resource.
func (r *scmSettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_control_settings"
}

// Schema returns the schema for the Source Control resource.
func (r *scmSettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"auto_branch_naming_enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When enabled, Retool automatically suggests a branch name on branch creation. Defaults to true.",
				Default:     booldefault.StaticBool(true),
			},
			"custom_pull_request_template_enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When enabled, Retool will use the template specified to create pull requests. Defaults to false.",
				Default:     booldefault.StaticBool(false),
			},
			"custom_pull_request_template": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Pull requests created from Retool will use the template specified.",
				Default:     stringdefault.StaticString(""),
			},
			"version_control_locked": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When set to true, creates a read-only instance of Retool, where app editing is disabled. Defaults to false.",
				Default:     booldefault.StaticBool(false),
			},
		},
	}
}

func updateSourceControlSettings(ctx context.Context, client *api.APIClient, model scmSettingsModel, globalDiags *diag.Diagnostics) {
	apiRequest := api.SourceControlSettingsPutRequest{
		Settings: api.SourceControlSettingsPutRequestSettings{
			AutoBranchNamingEnabled:          model.AutoBranchNamingEnabled.ValueBoolPointer(),
			CustomPullRequestTemplateEnabled: model.CustomPullRequestTemplateEnabled.ValueBoolPointer(),
			CustomPullRequestTemplate:        model.CustomPullRequestTemplate.ValueStringPointer(),
			VersionControlLocked:             model.VersionControlLocked.ValueBoolPointer(),
		},
	}

	_, httpResponse, err := client.SourceControlAPI.SourceControlSettingsPut(context.Background()).SourceControlSettingsPutRequest(apiRequest).Execute()
	if err != nil {
		globalDiags.AddError("Failed to update Source Control settings", err.Error())
		tflog.Error(ctx, "Error updating Source Control settings", utils.AddHttpStatusCode(map[string]interface{}{"error": err.Error()}, httpResponse))
		return
	}
}

// Updates Source Control settings and sets state.
func (r *scmSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan scmSettingsModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateSourceControlSettings(ctx, r.client, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state to fully populated data.
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error updating Source Control settings", map[string]interface{}{"error": "Could not set state"})
		return
	}

	tflog.Info(ctx, "Source Control settings updated")
}

// Read Source Control settings.
func (r *scmSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	response, httpResponse, err := r.client.SourceControlAPI.SourceControlSettingsGet(context.Background()).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			tflog.Info(ctx, "Source Control settings are not configured")
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error reading Source Control settings",
			fmt.Sprintf("Could not read Source Control config: %s", err.Error()),
		)
		tflog.Error(ctx, "Error reading Source Control config", utils.AddHttpStatusCode(map[string]any{"error": err.Error()}, httpResponse))
		return
	}

	var state scmSettingsModel

	state.AutoBranchNamingEnabled = types.BoolValue(response.Data.AutoBranchNamingEnabled)
	state.CustomPullRequestTemplateEnabled = types.BoolValue(response.Data.CustomPullRequestTemplateEnabled)
	state.CustomPullRequestTemplate = types.StringValue(response.Data.CustomPullRequestTemplate)
	state.VersionControlLocked = types.BoolValue(response.Data.VersionControlLocked)

	diags := resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error reading Source Control settings", map[string]interface{}{"error": "Could not set state"})
		return
	}
}

// Update Source Control settings.
func (r *scmSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan scmSettingsModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateSourceControlSettings(ctx, r.client, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state to fully populated data.
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error updating Source Control settings", map[string]interface{}{"error": "Could not set state"})
		return
	}

	tflog.Info(ctx, "Source Control settings updated")
}

// Delete Source Control settings.
func (r *scmSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// There's no DELETE endpoint, so we'll just set everything to default values
	// and call the update function.
	model := scmSettingsModel{
		AutoBranchNamingEnabled:          types.BoolValue(true),
		CustomPullRequestTemplateEnabled: types.BoolValue(false),
		CustomPullRequestTemplate:        types.StringValue(""),
		VersionControlLocked:             types.BoolValue(false),
	}
	updateSourceControlSettings(ctx, r.client, model, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
}
