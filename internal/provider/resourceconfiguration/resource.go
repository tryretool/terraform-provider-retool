// Package resourceconfiguration provides the implementation of the Retool Resource Configuration resource.
package resourceconfiguration

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// Ensure resourceConfigurationResource implements the tfsdk.Resource interface.
var (
	_ resource.Resource                = &resourceConfigurationResource{}
	_ resource.ResourceWithConfigure   = &resourceConfigurationResource{}
	_ resource.ResourceWithImportState = &resourceConfigurationResource{}
)

// resourceConfigurationResource schema structure.
type resourceConfigurationResource struct {
	client *api.APIClient
}

// resourceConfigurationModel defines the data model for the Resource Configuration resource.
type resourceConfigurationModel struct {
	ID            types.String `tfsdk:"id"`
	ResourceID    types.String `tfsdk:"resource_id"`
	EnvironmentID types.String `tfsdk:"environment_id"`
	Options       types.String `tfsdk:"options"`
	CreatedAt     types.String `tfsdk:"created_at"`
	UpdatedAt     types.String `tfsdk:"updated_at"`
}

// NewResource creates a new Resource Configuration resource.
func NewResource() resource.Resource {
	return &resourceConfigurationResource{}
}

// Configure adds the provider configured client to the resource.
func (r *resourceConfigurationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*utils.ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *utils.ProviderData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.client = providerData.Client
}

// Metadata associated with the Resource Configuration resource.
func (r *resourceConfigurationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource_configuration"
}

// Schema returns the schema for the Resource Configuration resource.
func (r *resourceConfigurationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `A resource configuration defines environment-specific configuration for a Retool resource. This allows you to configure different connection settings, credentials, and options for a resource across different environments (e.g., production, staging, development).

For example, you might have a PostgreSQL resource that connects to different database instances in production vs. staging environments, each with different credentials and connection strings.`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The UUID of the resource configuration.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"resource_id": schema.StringAttribute{
				Required:    true,
				Description: "The UUID or name of the resource to configure. Cannot be changed after creation.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"environment_id": schema.StringAttribute{
				Required:    true,
				Description: "The UUID of the environment for this configuration. Cannot be changed after creation.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"options": schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: "JSON string containing the environment-specific resource configuration options. The structure varies by resource type and mirrors the options structure used when creating a resource.",
				PlanModifiers: []planmodifier.String{
					JSONSemanticEquals(),
				},
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The timestamp when the resource configuration was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: "The timestamp when the resource configuration was last updated.",
			},
		},
	}
}

// Create creates the Resource Configuration and sets the initial Terraform state.
func (r *resourceConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan resourceConfigurationModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Parse options JSON string to map.
	var optionsMap map[string]interface{}
	if err := json.Unmarshal([]byte(plan.Options.ValueString()), &optionsMap); err != nil {
		resp.Diagnostics.AddError(
			"Invalid options JSON",
			fmt.Sprintf("Could not parse options as JSON: %s", err.Error()),
		)
		return
	}

	// Create the request body manually to avoid sending empty fields.
	requestBody := map[string]interface{}{
		"resource_id":    plan.ResourceID.ValueString(),
		"environment_id": plan.EnvironmentID.ValueString(),
		"options":        optionsMap,
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating request",
			fmt.Sprintf("Could not marshal request: %s", err.Error()),
		)
		return
	}

	var configRequest api.ResourceConfigurationsPostRequest
	if err := json.Unmarshal(requestJSON, &configRequest); err != nil {
		resp.Diagnostics.AddError(
			"Error creating request",
			fmt.Sprintf("Could not create request: %s", err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Creating a resource configuration", map[string]interface{}{
		"resource_id":    plan.ResourceID.ValueString(),
		"environment_id": plan.EnvironmentID.ValueString(),
	})

	// Create new resource configuration.
	response, httpResponse, err := r.client.ResourceConfigurationsAPI.ResourceConfigurationsPost(ctx).ResourceConfigurationsPostRequest(configRequest).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating resource configuration",
			fmt.Sprintf("Could not create resource configuration: %s", err.Error()),
		)
		tflog.Error(ctx, "Error creating resource configuration", utils.AddHTTPStatusCode(map[string]interface{}{"error": err.Error()}, httpResponse))
		return
	}

	// Map response body to schema and populate Computed attribute values.
	plan.ID = types.StringValue(response.Data.Id)
	plan.ResourceID = types.StringValue(response.Data.Resource.Id)
	plan.EnvironmentID = types.StringValue(response.Data.Environment.Id)
	plan.CreatedAt = types.StringValue(response.Data.CreatedAt)
	plan.UpdatedAt = types.StringValue(response.Data.UpdatedAt)
	// Keep the options in state after creation.

	// Set state to fully populated data.
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Resource configuration created", map[string]interface{}{
		"id":             response.Data.Id,
		"resource_id":    response.Data.Resource.Id,
		"environment_id": response.Data.Environment.Id,
	})
}

// Read a Resource Configuration.
func (r *resourceConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state resourceConfigurationModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	configID := state.ID.ValueString()

	// Note: Due to a bug in the generated SDK, the GET response model doesn't have a .Data wrapper
	// even though the API returns one. We need to make a raw HTTP call and parse it manually.
	// The actual API response format is identical for GET, POST, and PATCH operations.

	// Get the API client configuration to make a raw HTTP request.
	cfg := r.client.GetConfig()

	// Construct the full URL using the host and scheme from the configuration.
	url := fmt.Sprintf("%s://%s/api/v2/resource_configurations/%s", cfg.Scheme, cfg.Host, configID)
	httpReq, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating request",
			fmt.Sprintf("Could not create HTTP request: %s", err.Error()),
		)
		return
	}

	// Add authentication and other headers.
	for key, value := range cfg.DefaultHeader {
		httpReq.Header.Set(key, value)
	}
	httpReq.Header.Set("Accept", "application/json")

	// Execute the request using the client's HTTP client.
	httpResponse, err := cfg.HTTPClient.Do(httpReq)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading resource configuration",
			fmt.Sprintf("Could not read resource configuration with ID %s: %s", configID, err.Error()),
		)
		return
	}
	defer func() { _ = httpResponse.Body.Close() }()

	// Handle 404.
	if httpResponse.StatusCode == 404 {
		tflog.Info(ctx, "Resource configuration not found", map[string]any{"configurationID": configID})
		resp.State.RemoveResource(ctx)
		return
	}

	// Read the response body.
	bodyBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading response",
			fmt.Sprintf("Could not read response body: %s", err.Error()),
		)
		return
	}

	// Check for error status codes.
	if httpResponse.StatusCode >= 300 {
		resp.Diagnostics.AddError(
			"Error reading resource configuration",
			fmt.Sprintf("API returned status %d: %s", httpResponse.StatusCode, string(bodyBytes)),
		)
		return
	}

	// Parse the response using the correct Post200Response type (which has .Data wrapper).
	var configuration api.ResourceConfigurationsPost200Response
	if err := json.Unmarshal(bodyBytes, &configuration); err != nil {
		resp.Diagnostics.AddError(
			"Error parsing response",
			fmt.Sprintf("Could not parse response: %s", err.Error()),
		)
		return
	}

	// Map the API response to the Terraform state.
	state.ID = types.StringValue(configuration.Data.Id)
	state.ResourceID = types.StringValue(configuration.Data.Resource.Id)
	state.EnvironmentID = types.StringValue(configuration.Data.Environment.Id)
	state.CreatedAt = types.StringValue(configuration.Data.CreatedAt)
	state.UpdatedAt = types.StringValue(configuration.Data.UpdatedAt)

	// The API returns options, so we can update from the response.
	// However, options may contain sensitive data, so we'll marshal it back to JSON.
	optionsJSON, err := json.Marshal(configuration.Data.Options)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error processing options",
			fmt.Sprintf("Could not marshal options to JSON: %s", err.Error()),
		)
		return
	}
	state.Options = types.StringValue(string(optionsJSON))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update a Resource Configuration.
func (r *resourceConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan resourceConfigurationModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state resourceConfigurationModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	configID := state.ID.ValueString()

	// Build patch operations based on changes.
	operations := []api.ReplaceOperation{}

	// Check if options have changed.
	if !plan.Options.Equal(state.Options) {
		// Parse the new options JSON.
		var optionsMap map[string]interface{}
		if err := json.Unmarshal([]byte(plan.Options.ValueString()), &optionsMap); err != nil {
			resp.Diagnostics.AddError(
				"Invalid options JSON",
				fmt.Sprintf("Could not parse options as JSON: %s", err.Error()),
			)
			return
		}

		replaceOp := api.NewReplaceOperation("replace", "/options")
		replaceOp.SetValue(optionsMap)
		operations = append(operations, *replaceOp)
	}

	if len(operations) == 0 {
		tflog.Info(ctx, "No changes detected for resource configuration", map[string]any{"configurationID": configID})
		// No changes, just return current state.
		diags = resp.State.Set(ctx, plan)
		resp.Diagnostics.Append(diags...)
		return
	}

	// Prepare the update payload.
	patchRequest := api.NewResourceConfigurationsConfigurationIdPatchRequest(operations)

	// Perform the update operation.
	updateResponse, httpResponse, err := r.client.ResourceConfigurationsAPI.ResourceConfigurationsConfigurationIdPatch(ctx, configID).ResourceConfigurationsConfigurationIdPatchRequest(*patchRequest).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating resource configuration",
			fmt.Sprintf("Could not update resource configuration with ID %s: %s", configID, err.Error()),
		)
		tflog.Error(ctx, "Error updating resource configuration", utils.AddHTTPStatusCode(map[string]any{"configurationID": configID, "error": err.Error()}, httpResponse))
		return
	}

	// Update the state with response data.
	plan.ID = types.StringValue(updateResponse.Data.Id)
	plan.ResourceID = types.StringValue(updateResponse.Data.Resource.Id)
	plan.EnvironmentID = types.StringValue(updateResponse.Data.Environment.Id)
	plan.CreatedAt = types.StringValue(updateResponse.Data.CreatedAt)
	plan.UpdatedAt = types.StringValue(updateResponse.Data.UpdatedAt)

	// Note: We don't update options from the API response because:.
	// 1. The API adds default values that weren't in the plan.
	// 2. This would cause Terraform to see inconsistent values for sensitive attributes.
	// Instead, we keep the planned value which is what the user specified.

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Resource configuration updated", map[string]interface{}{"id": configID})
}

// Delete a Resource Configuration.
func (r *resourceConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state resourceConfigurationModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	configID := state.ID.ValueString()

	// Delete the resource configuration.
	httpResponse, err := r.client.ResourceConfigurationsAPI.ResourceConfigurationsConfigurationIdDelete(ctx, configID).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			// Resource configuration already doesn't exist, nothing to do.
			tflog.Info(ctx, "Resource configuration not found during delete, treating as already deleted", map[string]any{"configurationID": configID})
			return
		}
		resp.Diagnostics.AddError(
			"Error Deleting Resource Configuration",
			fmt.Sprintf("Could not delete resource configuration with ID %s: %s", configID, err.Error()),
		)
		tflog.Error(ctx, "Error deleting resource configuration", utils.AddHTTPStatusCode(map[string]any{"error": err.Error(), "configurationID": configID}, httpResponse))
		return
	}

	tflog.Info(ctx, "Resource configuration deleted", map[string]interface{}{"id": configID})
}

// ImportState allows importing of a Resource Configuration.
func (r *resourceConfigurationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute.
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
