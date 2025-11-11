// Package retoolresource provides the implementation of the Retool Resource resource.
// Note: "retoolresource" package name is used to avoid collision with Go's resource keyword.
package retoolresource

import (
	"context"
	"encoding/json"
	"fmt"

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

// Ensure retoolResourceResource implements the tfsdk.Resource interface.
var (
	_ resource.Resource                = &retoolResourceResource{}
	_ resource.ResourceWithConfigure   = &retoolResourceResource{}
	_ resource.ResourceWithImportState = &retoolResourceResource{}
)

// retoolResourceResource schema structure.
type retoolResourceResource struct {
	client *api.APIClient
}

// retoolResourceModel defines the data model for the Retool Resource resource.
type retoolResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Type        types.String `tfsdk:"type"`
	DisplayName types.String `tfsdk:"display_name"`
	Options     types.String `tfsdk:"options"`
	Protected   types.Bool   `tfsdk:"protected"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
}

// NewResource creates a new Retool Resource resource.
func NewResource() resource.Resource {
	return &retoolResourceResource{}
}

// Configure adds the provider configured client to the resource.
func (r *retoolResourceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Metadata associated with the Retool Resource resource.
func (r *retoolResourceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_resource"
}

// Schema returns the schema for the Retool Resource resource.
func (r *retoolResourceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `A Retool resource represents a connection to an external service such as a database, API, or other data source.

**Important Note:** Due to API security design, the options field (which contains connection credentials and configuration) is write-only and cannot be read back after creation. This means:
- Changes to options will not be detected by Terraform after the resource is created
- The options value will not be stored in Terraform state after initial creation
- To update resource options, you must use the Retool UI or recreate the resource

This resource is best used for managing the resource metadata (display_name, type) while the actual connection configuration is managed through other means.`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The UUID of the resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "The type of resource (e.g., 'restapi', 'postgresql', 'mysql', 'snowflake'). Cannot be changed after creation.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: "The display name of the resource.",
			},
			"options": schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: "JSON string containing the resource configuration options. The structure varies by resource type. This field is write-only and cannot be read back after creation.",
			},
			"protected": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether the resource is protected in source control.",
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The timestamp when the resource was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: "The timestamp when the resource was last updated.",
			},
		},
	}
}

// Create creates the Retool Resource and sets the initial Terraform state.
func (r *retoolResourceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan retoolResourceModel
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
		"type":         plan.Type.ValueString(),
		"display_name": plan.DisplayName.ValueString(),
		"options":      optionsMap,
	}

	requestJSON, err := json.Marshal(requestBody)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating request",
			fmt.Sprintf("Could not marshal request: %s", err.Error()),
		)
		return
	}

	var resourceRequest api.ResourcesPostRequest
	if err := json.Unmarshal(requestJSON, &resourceRequest); err != nil {
		resp.Diagnostics.AddError(
			"Error creating request",
			fmt.Sprintf("Could not create request: %s", err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Creating a Retool resource", map[string]interface{}{
		"type":         plan.Type.ValueString(),
		"display_name": plan.DisplayName.ValueString(),
	})

	// Create new resource.
	response, httpResponse, err := r.client.ResourcesAPI.ResourcesPost(ctx).ResourcesPostRequest(resourceRequest).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Retool resource",
			fmt.Sprintf("Could not create resource: %s", err.Error()),
		)
		tflog.Error(ctx, "Error creating Retool resource", utils.AddHTTPStatusCode(map[string]interface{}{"error": err.Error()}, httpResponse))
		return
	}

	// Map response body to schema and populate Computed attribute values.
	plan.ID = types.StringValue(response.Data.Id)
	plan.Type = types.StringValue(response.Data.Type)
	plan.DisplayName = types.StringValue(response.Data.DisplayName)
	plan.Protected = types.BoolValue(response.Data.Protected)
	plan.CreatedAt = types.StringValue(response.Data.CreatedAt)
	plan.UpdatedAt = types.StringValue(response.Data.UpdatedAt)
	// Note: We keep the options in state after creation, but it won't be refreshed on read.

	// Set state to fully populated data.
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Retool resource created", map[string]interface{}{
		"id":           response.Data.Id,
		"display_name": response.Data.DisplayName,
	})
}

// Read a Retool Resource.
func (r *retoolResourceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state retoolResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceID := state.ID.ValueString()
	resource, httpResponse, err := r.client.ResourcesAPI.ResourcesResourceIdGet(ctx, resourceID).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			tflog.Info(ctx, "Retool resource not found", map[string]any{"resourceID": resourceID})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error reading Retool resource",
			fmt.Sprintf("Could not read resource with ID %s: %s", resourceID, err.Error()),
		)
		tflog.Error(ctx, "Error reading Retool resource", utils.AddHTTPStatusCode(map[string]any{"resourceID": resourceID, "error": err.Error()}, httpResponse))
		return
	}

	// Map the API response to the Terraform state.
	state.ID = types.StringValue(resource.Data.Id)
	state.Type = types.StringValue(resource.Data.Type)
	state.DisplayName = types.StringValue(resource.Data.DisplayName)
	state.Protected = types.BoolValue(resource.Data.Protected)
	state.CreatedAt = types.StringValue(resource.Data.CreatedAt)
	state.UpdatedAt = types.StringValue(resource.Data.UpdatedAt)
	// Note: options are not returned by the API, so we keep the existing state value.

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update a Retool Resource.
func (r *retoolResourceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan retoolResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state retoolResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceID := state.ID.ValueString()

	// Build patch operations based on changes.
	operations := []api.ReplaceOperation{}

	// Only display_name can be updated (type requires replacement, options cannot be updated via API).
	if !plan.DisplayName.Equal(state.DisplayName) {
		replaceOp := api.NewReplaceOperation("replace", "/display_name")
		replaceOp.SetValue(plan.DisplayName.ValueString())
		operations = append(operations, *replaceOp)
	}

	if len(operations) == 0 {
		tflog.Info(ctx, "No changes detected for Retool resource", map[string]any{"resourceID": resourceID})
		// No changes, just return current state.
		diags = resp.State.Set(ctx, plan)
		resp.Diagnostics.Append(diags...)
		return
	}

	// Prepare the update payload.
	patchRequest := api.NewResourcesResourceIdPatchRequest(operations)

	// Perform the update operation.
	updateResponse, httpResponse, err := r.client.ResourcesAPI.ResourcesResourceIdPatch(ctx, resourceID).ResourcesResourceIdPatchRequest(*patchRequest).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating Retool resource",
			fmt.Sprintf("Could not update resource with ID %s: %s", resourceID, err.Error()),
		)
		tflog.Error(ctx, "Error updating Retool resource", utils.AddHTTPStatusCode(map[string]any{"resourceID": resourceID, "error": err.Error()}, httpResponse))
		return
	}

	// Update the state with response data.
	plan.ID = types.StringValue(updateResponse.Data.Id)
	plan.Type = types.StringValue(updateResponse.Data.Type)
	plan.DisplayName = types.StringValue(updateResponse.Data.DisplayName)
	plan.Protected = types.BoolValue(updateResponse.Data.Protected)
	plan.CreatedAt = types.StringValue(updateResponse.Data.CreatedAt)
	plan.UpdatedAt = types.StringValue(updateResponse.Data.UpdatedAt)
	// Note: options are not returned by the API, so we keep the existing state value.

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Retool resource updated", map[string]interface{}{"id": resourceID})
}

// Delete a Retool Resource.
func (r *retoolResourceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state retoolResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceID := state.ID.ValueString()

	// Delete the resource.
	httpResponse, err := r.client.ResourcesAPI.ResourcesResourceIdDelete(ctx, resourceID).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			// Resource already doesn't exist, nothing to do.
			tflog.Info(ctx, "Retool resource not found during delete, treating as already deleted", map[string]any{"resourceID": resourceID})
			return
		}
		resp.Diagnostics.AddError(
			"Error Deleting Retool Resource",
			fmt.Sprintf("Could not delete resource with ID %s: %s", resourceID, err.Error()),
		)
		tflog.Error(ctx, "Error deleting Retool resource", utils.AddHTTPStatusCode(map[string]any{"error": err.Error(), "resourceID": resourceID}, httpResponse))
		return
	}

	tflog.Info(ctx, "Retool resource deleted", map[string]interface{}{"id": resourceID})
}

// ImportState allows importing of a Retool Resource.
func (r *retoolResourceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute.
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	// Add a warning about options not being imported.
	resp.Diagnostics.AddWarning(
		"Options Not Imported",
		"The options field cannot be imported as it is not returned by the API for security reasons. "+
			"You will need to manually set the options field in your Terraform configuration to match the resource's current configuration. "+
			"Consider using lifecycle { ignore_changes = [options] } if you don't need Terraform to manage the connection details.",
	)
}
