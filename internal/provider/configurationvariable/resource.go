// Package configurationvariable provides the implementation of the Configuration Variable resource.
package configurationvariable

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// Ensure ConfigurationVariableResource implements the tfsdk.Resource interface.
var (
	_ resource.Resource                = &configurationVariableResource{}
	_ resource.ResourceWithConfigure   = &configurationVariableResource{}
	_ resource.ResourceWithImportState = &configurationVariableResource{}
)

// ConfigurationVariable schema structure.
type configurationVariableResource struct {
	client *api.APIClient
}

// configurationVariableValueModel maps the values attribute of the configuration variable resource.
type configurationVariableValueModel struct {
	EnvironmentID types.String `tfsdk:"environment_id"`
	Value         types.String `tfsdk:"value"`
}

// configurationVariableResourceModel defines the data model for the configuration variable resource.
type configurationVariableResourceModel struct {
	ID          types.String                      `tfsdk:"id"`
	Name        types.String                      `tfsdk:"name"`
	Description types.String                      `tfsdk:"description"`
	Secret      types.Bool                        `tfsdk:"secret"`
	Values      []configurationVariableValueModel `tfsdk:"values"`
}

// Create a new Configuration Variable resource.
func NewResource() resource.Resource {
	return &configurationVariableResource{}
}

// Configure add the provider configured client to the resource.
func (r *configurationVariableResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*utils.ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = providerData.Client
}

// Metadata associated with the configuration variable resource.
func (r *configurationVariableResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_configuration_variable"
}

// Schema returnrs the schema for the Configuration Variable resource.
func (r *configurationVariableResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages a Retool configuration variable. Configuration variables are used to store environment-specific values that can be referenced throughout your Retool applications.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "The unique identifier for the configuration variable.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the configuration variable. This is how you'll reference it in your Retool applications.",
				Validators:  []validator.String{stringvalidator.LengthBetween(1, 255)},
			},
			"description": schema.StringAttribute{
				Optional:      true,
				Description:   "A brief description of the configuration variable's purpose.",
				Validators:    []validator.String{stringvalidator.LengthBetween(0, 1024)},
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"secret": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the configuration variable is a secret. Secrets are encrypted and not exposed in the Retool UI.",
				Default:     booldefault.StaticBool(false),
			},
			"values": schema.ListNestedAttribute{
				Description: "A list of environment-specific values for the configuration variable.",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"environment_id": schema.StringAttribute{
							Description: "The ID of the environment this value is associated with.",
							Required:    true,
							Validators:  []validator.String{stringvalidator.LengthBetween(1, 255)},
						},
						"value": schema.StringAttribute{
							Description: "The value of the configuration variable for the specified environment.",
							Required:    true,
							Sensitive:   true,
							Validators:  []validator.String{stringvalidator.LengthBetween(1, 4096)},
						},
					},
				},
			},
		},
	}
}

// Create creates a configuration variable.
func (r *configurationVariableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan configurationVariableResourceModel

	// Read Terraform plan data into the model.
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var configuratonVariable api.ConfigurationVariablesPostRequest
	configuratonVariable.Name = plan.Name.ValueString()
	if !plan.Description.IsNull() {
		configuratonVariable.Description = plan.Description.ValueStringPointer()
	}
	configuratonVariable.Secret = plan.Secret.ValueBool()

	var values []api.ConfigurationVariablesGet200ResponseDataInnerValuesInner
	for _, v := range plan.Values {
		values = append(values, api.ConfigurationVariablesGet200ResponseDataInnerValuesInner{
			EnvironmentId: v.EnvironmentID.ValueString(),
			Value:         v.Value.ValueString(),
		})
	}
	configuratonVariable.Values = values

	response, httpResponse, err := r.client.ConfigurationVariablesAPI.ConfigurationVariablesPost(ctx).ConfigurationVariablesPostRequest(configuratonVariable).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating configuration variable",
			fmt.Sprintf("Could not create configuration variable, unexpected error: %s", err.Error()),
		)
		tflog.Error(ctx, "Error creating configuration variable", utils.AddHTTPStatusCode(map[string]interface{}{"error": err.Error()}, httpResponse))
		return
	}

	if response.Data.Id == "" {
		resp.Diagnostics.AddError(
			"Error creating configuration variable",
			"Could not create configuration variable, no ID was returned.",
		)
		tflog.Error(ctx, "Error creating configuration variable, no ID was returned", utils.AddHTTPStatusCode(map[string]interface{}{}, httpResponse))
		return
	}

	// Map response body to schema and populate Computed attribute values.
	plan.ID = types.StringValue(response.Data.Id)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error setting state after creating configuration variable", map[string]interface{}{"id": plan.ID.ValueString()})
		return
	}

	tflog.Info(ctx, "Created configuration variable", map[string]interface{}{"id": plan.ID.ValueString()})
}

// Read a configuration variable.
func (r *configurationVariableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state configurationVariableResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	configurationVariableID := state.ID.ValueString()
	tflog.Info(ctx, "Reading configuration variable", map[string]interface{}{"id": configurationVariableID})

	response, httpResponse, err := r.client.ConfigurationVariablesAPI.ConfigurationVariablesIdGet(ctx, configurationVariableID).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			tflog.Info(ctx, "Configuration variable not found, removing from state", map[string]interface{}{"id": configurationVariableID})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error reading configuration variable",
			fmt.Sprintf("Could not read configuration variable with ID %s: %s", configurationVariableID, err.Error()),
		)
		tflog.Error(ctx, "Error reading configuration variable", utils.AddHTTPStatusCode(map[string]interface{}{"id": configurationVariableID, "error": err.Error()}, httpResponse))
		return
	}

	state.ID = types.StringValue(response.Data.Id)
	state.Name = types.StringValue(response.Data.Name)
	if response.Data.Description != nil {
		state.Description = types.StringValue(*response.Data.Description)
	} else {
		state.Description = types.StringNull()
	}

	state.Secret = types.BoolValue(response.Data.Secret)

	// For secrets, the API returns encrypted values that we cannot use.
	// We preserve values from the existing state (which come from the config).
	// During import, state values will be empty/null, which signals to the user they must provide them.
	if response.Data.Secret {
		// Preserve existing state values for secrets since API returns encrypted placeholders.
		// We need to update the environment IDs from the API response, but keep the values from state.
		existingValues := make(map[string]types.String)
		for _, v := range state.Values {
			existingValues[v.EnvironmentID.ValueString()] = v.Value
		}

		state.Values = nil
		for _, v := range response.Data.Values {
			envID := v.EnvironmentId
			existingValue, hasExisting := existingValues[envID]

			// Use existing value if present (normal refresh), otherwise null (import scenario).
			if hasExisting && !existingValue.IsNull() {
				state.Values = append(state.Values, configurationVariableValueModel{
					EnvironmentID: types.StringValue(envID),
					Value:         existingValue,
				})
			} else {
				state.Values = append(state.Values, configurationVariableValueModel{
					EnvironmentID: types.StringValue(envID),
					Value:         types.StringNull(), // API sends placeholder, set to null for import.
				})
				tflog.Warn(ctx, "Configuration variable is a secret with no value in state. Value must be provided in your Terraform configuration.", map[string]interface{}{"id": configurationVariableID, "environment_id": envID})
			}
		}
	} else {
		// Clear current values and repopulate from API response to handle deletions correctly.
		state.Values = nil
		for _, v := range response.Data.Values {
			state.Values = append(state.Values, configurationVariableValueModel{
				EnvironmentID: types.StringValue(v.EnvironmentId),
				Value:         types.StringValue(v.Value),
			})
		}
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error setting state after reading configuration variable", map[string]interface{}{"id": state.ID.ValueString()})
		return
	}

	tflog.Info(ctx, "Read configuration variable", map[string]interface{}{"id": state.ID.ValueString()})
}

// Update a configuration variable.
func (r *configurationVariableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan configurationVariableResourceModel

	// Read Terraform plan data into the model.
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	configurationVariableID := plan.ID.ValueString()
	tflog.Info(ctx, "Updating configuration variable", map[string]interface{}{"id": configurationVariableID})

	var values []api.ConfigurationVariablesGet200ResponseDataInnerValuesInner
	for _, v := range plan.Values {
		values = append(values, api.ConfigurationVariablesGet200ResponseDataInnerValuesInner{
			EnvironmentId: v.EnvironmentID.ValueString(),
			Value:         v.Value.ValueString(),
		})
	}

	updatePayload := api.ConfigurationVariablesPostRequest{
		Name:        plan.Name.ValueString(),
		Description: plan.Description.ValueStringPointer(),
		Secret:      plan.Secret.ValueBool(),
		Values:      values,
	}

	_, httpResponse, err := r.client.ConfigurationVariablesAPI.ConfigurationVariablesIdPut(ctx, configurationVariableID).ConfigurationVariablesPostRequest(updatePayload).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating configuration variable",
			fmt.Sprintf("Could not update configuration variable with ID %s: %s", configurationVariableID, err.Error()),
		)
		tflog.Error(ctx, "Error updating configuration variable", utils.AddHTTPStatusCode(map[string]interface{}{"id": configurationVariableID, "error": err.Error()}, httpResponse))
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error setting state after updating configuration variable", map[string]interface{}{"id": plan.ID.ValueString()})
		return
	}

	tflog.Info(ctx, "Updated configuration variable", map[string]interface{}{"id": plan.ID.ValueString()})
}

// Delete deletes a configuration variable.
func (r *configurationVariableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state configurationVariableResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	configurationVariableID := state.ID.ValueString()
	tflog.Info(ctx, "Deleting configuration variable", map[string]interface{}{"id": configurationVariableID})

	httpResponse, err := r.client.ConfigurationVariablesAPI.ConfigurationVariablesIdDelete(ctx, configurationVariableID).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			tflog.Info(ctx, "Configuration variable not found, removing from state", map[string]interface{}{"id": configurationVariableID})
			return
		}
		resp.Diagnostics.AddError(
			"Error deleting configuration variable",
			fmt.Sprintf("Could not delete configuration variable with ID %s: %s", configurationVariableID, err.Error()),
		)
		tflog.Error(ctx, "Error deleting configuration variable", utils.AddHTTPStatusCode(map[string]interface{}{"id": configurationVariableID, "error": err.Error()}, httpResponse))
		return
	}

	tflog.Info(ctx, "Deleted configuration variable", map[string]interface{}{"id": configurationVariableID})
}

// ImportState imports a configuration variable resource.
func (r *configurationVariableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
