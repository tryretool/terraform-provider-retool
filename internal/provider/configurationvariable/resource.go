package configurationvariable

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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
	_ resource.Resource                = &ConfigurationVariableResource{}
	_ resource.ResourceWithConfigure   = &ConfigurationVariableResource{}
	_ resource.ResourceWithImportState = &ConfigurationVariableResource{}
)

type ConfigurationVariableResource struct {
	client *api.APIClient
}

type configurationVariableValueModel struct {
	EnvironmentID types.String `tfsdk:"environment_id"`
	Value         types.String `tfsdk:"value"`
}

func (m configurationVariableValueModel) attributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"environment_id": types.StringType,
		"value":          types.StringType,
	}
}

type configurationVariableResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Secret      types.Bool   `tfsdk:"secret"`
	Values      types.List   `tfsdk:"values"`
}

func NewResource() resource.Resource {
	return &ConfigurationVariableResource{}
}

// Configure add the provider configured client to the resource.
func (r *ConfigurationVariableResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*api.APIClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *api.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = providerData
}

// Metadata associated with the configuration variable resource.
func (r *ConfigurationVariableResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_configuration_variable"
}

// Schema returnrs the schema for the Configuration Variable resource.
func (r *ConfigurationVariableResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
							Validators:  []validator.String{stringvalidator.LengthBetween(1, 4096)},
						},
					},
				},
			},
		},
	}
}

func (r *ConfigurationVariableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan configurationVariableResourceModel

	// Read Terraform plan data into the model
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

	var valuesObjectsList = plan.Values
	var values []api.ConfigurationVariablesGet200ResponseDataInnerValuesInner
	diags = valuesObjectsList.ElementsAs(ctx, &values, false)
	resp.Diagnostics.Append(diags...)

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

// Read a configuration variable
func (r *ConfigurationVariableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
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

	values := make([]configurationVariableValueModel, 0, len(response.Data.Values))
	for _, v := range response.Data.Values {
		values = append(values, configurationVariableValueModel{
			EnvironmentID: types.StringValue(v.EnvironmentId),
			Value:         types.StringValue(v.Value),
		})
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error setting state after reading configuration variable", map[string]interface{}{"id": state.ID.ValueString()})
		return
	}

	tflog.Info(ctx, "Read configuration variable", map[string]interface{}{"id": state.ID.ValueString()})
}

func (r *ConfigurationVariableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan configurationVariableResourceModel

	// Read Terraform plan data into the model
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	configurationVariableID := plan.ID.ValueString()

	valuesObjectsList := plan.Values

	var values []api.ConfigurationVariablesGet200ResponseDataInnerValuesInner
	diags = valuesObjectsList.ElementsAs(ctx, &values, false)
	resp.Diagnostics.Append(diags...)

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

func (r *ConfigurationVariableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
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
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error deleting configuration variable",
			fmt.Sprintf("Could not delete configuration variable with ID %s: %s", configurationVariableID, err.Error()),
		)
		tflog.Error(ctx, "Error deleting configuration variable", utils.AddHTTPStatusCode(map[string]interface{}{"id": configurationVariableID, "error": err.Error()}, httpResponse))
		return
	}

	resp.State.RemoveResource(ctx)

	tflog.Info(ctx, "Deleted configuration variable", map[string]interface{}{"id": configurationVariableID})
}

func (r *ConfigurationVariableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
