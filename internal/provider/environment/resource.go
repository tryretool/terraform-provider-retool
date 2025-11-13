// Package environment provides implementation of the Environment resource.
package environment

import (
	"context"

	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type environmentResource struct {
	client *api.APIClient
}

// Ensure EnvironmentResource implements the tfsdk.Resource interface.
var (
	_ resource.Resource                = &environmentResource{}
	_ resource.ResourceWithConfigure   = &environmentResource{}
	_ resource.ResourceWithImportState = &environmentResource{}
)

type environmentResourceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Color       types.String `tfsdk:"color"`
	Default     types.Bool   `tfsdk:"default"`
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
}

// Create new Environment resource.
func NewResource() resource.Resource {
	return &environmentResource{}
}

func (r *environmentResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_environment"
}

func (r *environmentResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Environment resource allows you to create and manage Environments in Retool. Environments are used to manage different settings for your Retool apps in different contexts, such as development, staging, and production.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the environment.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the environment.",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "A brief description of the environment.",
			},
			"color": schema.StringAttribute{
				Required:    true,
				Description: "The hexadecimal color code for the environment (e.g., #FF5733).",
			},
			"default": schema.BoolAttribute{
				Computed:    true,
				Description: "Indicates if this is the default environment.",
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The timestamp when the environment was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				Computed:    true,
				Description: "The timestamp when the environment was last updated.",
			},
		},
	}
}

func (r *environmentResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *environmentResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan environmentResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Creating Environment", map[string]interface{}{"name": plan.Name})

	request := api.EnvironmentsPostRequest{
		Name:  plan.Name.ValueString(),
		Color: plan.Color.ValueString(),
	}
	if !plan.Description.IsNull() && !plan.Description.IsUnknown() {
		description := plan.Description.ValueString()
		request.Description = &description
	}

	response, httpResponse, err := r.client.EnvironmentsAPI.EnvironmentsPost(ctx).EnvironmentsPostRequest(request).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Environment",
			"Could not create Environment, unexpected error: "+err.Error(),
		)
		tflog.Error(ctx, "Error creating Environment", utils.AddHTTPStatusCode(map[string]interface{}{"error": err.Error()}, httpResponse))
		return
	}

	plan.ID = types.StringValue(response.Data.Id)
	plan.Name = types.StringValue(response.Data.Name)
	plan.Description = types.StringPointerValue(response.Data.Description.Get())
	plan.Color = types.StringValue(response.Data.Color)
	plan.Default = types.BoolValue(response.Data.Default)
	plan.CreatedAt = types.StringValue(response.Data.CreatedAt)
	plan.UpdatedAt = types.StringValue(response.Data.UpdatedAt)

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Environment created", map[string]interface{}{"id": plan.ID})
}

func (r *environmentResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state environmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentID := state.ID.ValueString()
	tflog.Info(ctx, "Reading Environment", map[string]interface{}{"id": environmentID})
	response, httpResponse, err := r.client.EnvironmentsAPI.EnvironmentsEnvironmentIdGet(ctx, environmentID).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			tflog.Info(ctx, "Environment not found", map[string]any{"id": environmentID})
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Error reading Environment",
			"Could not read Environment, unexpected error: "+err.Error(),
		)
		tflog.Error(ctx, "Error reading Environment", utils.AddHTTPStatusCode(map[string]interface{}{"error": err.Error(), "id": environmentID}, httpResponse))
		return
	}
	state.Name = types.StringValue(response.Data.Name)
	state.Description = types.StringPointerValue(response.Data.Description.Get())
	state.Color = types.StringValue(response.Data.Color)
	state.Default = types.BoolValue(response.Data.Default)
	state.CreatedAt = types.StringValue(response.Data.CreatedAt)
	state.UpdatedAt = types.StringValue(response.Data.UpdatedAt)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Environment read", map[string]interface{}{"id": environmentID})
}

func (r *environmentResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state environmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plan environmentResourceModel
	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentID := state.ID.ValueString()
	tflog.Info(ctx, "Updating Environment", map[string]interface{}{"id": environmentID, "name": plan.Name.ValueString()})

	var operations []api.ReplaceOperation

	// Check if name changed
	if !plan.Name.Equal(state.Name) {
		nameOp := api.NewReplaceOperation("replace", "/name")
		nameOp.Value = plan.Name.ValueStringPointer()
		operations = append(operations, *nameOp)
	}

	// Check if description changed
	if !plan.Description.Equal(state.Description) {
		descOp := api.NewReplaceOperation("replace", "/description")
		if plan.Description.IsNull() {
			descOp.SetValue(nil)
		} else {
			descOp.Value = plan.Description.ValueStringPointer()
		}
		operations = append(operations, *descOp)
	}

	// Check if color changed
	if !plan.Color.Equal(state.Color) {
		colorOp := api.NewReplaceOperation("replace", "/color")
		colorOp.Value = plan.Color.ValueStringPointer()
		operations = append(operations, *colorOp)
	}

	// Only make the PATCH request if there are operations to perform
	if len(operations) > 0 {
		patchRequest := api.NewEnvironmentsEnvironmentIdPatchRequest(operations)
		response, httpResponse, err := r.client.EnvironmentsAPI.EnvironmentsEnvironmentIdPatch(ctx, environmentID).EnvironmentsEnvironmentIdPatchRequest(*patchRequest).Execute()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error updating Environment",
				"Could not update Environment with id "+environmentID+", unexpected error: "+err.Error(),
			)
			tflog.Error(ctx, "Error updating Environment", utils.AddHTTPStatusCode(map[string]interface{}{"error": err.Error(), "id": environmentID}, httpResponse))
			return
		}

		plan.Name = types.StringValue(response.Data.Name)
		plan.Description = types.StringPointerValue(response.Data.Description.Get())
		plan.Color = types.StringValue(response.Data.Color)
		plan.Default = types.BoolValue(response.Data.Default)
		plan.CreatedAt = types.StringValue(response.Data.CreatedAt)
		plan.UpdatedAt = types.StringValue(response.Data.UpdatedAt)
	}

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Info(ctx, "Environment updated", map[string]interface{}{"id": environmentID})
}

func (r *environmentResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state environmentResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentID := state.ID.ValueString()
	tflog.Info(ctx, "Deleting Environment", map[string]interface{}{"id": environmentID})
	httpResponse, err := r.client.EnvironmentsAPI.EnvironmentsEnvironmentIdDelete(ctx, environmentID).Execute()
	if err != nil && !(httpResponse != nil && httpResponse.StatusCode == 404) {
		resp.Diagnostics.AddError(
			"Error deleting Environment",
			"Could not delete Environment with id "+environmentID+", unexpected error: "+err.Error(),
		)
		tflog.Error(ctx, "Error deleting Environment", utils.AddHTTPStatusCode(map[string]interface{}{"error": err.Error(), "id": environmentID}, httpResponse))
		return
	}
}

// ImportState allows importing of an Environment resource.
func (r *environmentResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute.
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
