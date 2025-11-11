// Package user provides the implementation of the User resource.
package user

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// Ensure userResource implements the tfsdk.Resource interface.
var (
	_ resource.Resource                = &userResource{}
	_ resource.ResourceWithConfigure   = &userResource{}
	_ resource.ResourceWithImportState = &userResource{}
)

// userResource schema structure.
type userResource struct {
	client *api.APIClient
}

// userResourceModel defines the data model for the User resource.
type userResourceModel struct {
	ID        types.String `tfsdk:"id"`
	LegacyID  types.String `tfsdk:"legacy_id"`
	Email     types.String `tfsdk:"email"`
	FirstName types.String `tfsdk:"first_name"`
	LastName  types.String `tfsdk:"last_name"`
	Active    types.Bool   `tfsdk:"active"`
	Metadata  types.String `tfsdk:"metadata"`
	UserType  types.String `tfsdk:"user_type"`
	// Read-only computed fields.
	CreatedAt            types.String `tfsdk:"created_at"`
	LastActive           types.String `tfsdk:"last_active"`
	IsAdmin              types.Bool   `tfsdk:"is_admin"`
	TwoFactorAuthEnabled types.Bool   `tfsdk:"two_factor_auth_enabled"`
}

// Create new User resource.
func NewResource() resource.Resource {
	return &userResource{}
}

// Configure adds the provider configured client to the resource.
func (r *userResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// Metadata associated with the User resource.
func (r *userResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

// Schema returns the schema for the User resource.
func (r *userResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `A user represents an individual with access to the Retool instance. Users can be granted access to apps, resources, and workflows through group memberships and permissions.

Note: Deleting a user resource sets the user to inactive (active: false) rather than permanently removing them from Retool.`,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the user. Currently this is the same as legacy_id but will change in the future.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"legacy_id": schema.StringAttribute{
				Computed:    true,
				Description: "The legacy ID of the user.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"email": schema.StringAttribute{
				Required:    true,
				Description: "The email address of the user. Cannot be changed after creation.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"first_name": schema.StringAttribute{
				Required:    true,
				Description: "The first name of the user.",
			},
			"last_name": schema.StringAttribute{
				Required:    true,
				Description: "The last name of the user.",
			},
			"active": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
				Description: "Whether the user is active or not. Default is true.",
			},
			"metadata": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "JSON string containing custom metadata for the user.",
			},
			"user_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The user type. Accepted values vary by Retool instance configuration.",
			},
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The timestamp when the user was created.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_active": schema.StringAttribute{
				Computed:    true,
				Description: "The timestamp when the user was last active.",
			},
			"is_admin": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether the user is an admin or not.",
			},
			"two_factor_auth_enabled": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether two-factor authentication is enabled for this user.",
			},
		},
	}
}

// Create creates the user resource and sets the initial Terraform state.
func (r *userResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan.
	var plan userResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan.
	user := api.NewUsersPostRequest(
		plan.Email.ValueString(),
		plan.FirstName.ValueString(),
		plan.LastName.ValueString(),
	)

	if !plan.Active.IsNull() && !plan.Active.IsUnknown() {
		user.SetActive(plan.Active.ValueBool())
	}

	if !plan.Metadata.IsNull() && !plan.Metadata.IsUnknown() {
		// Parse metadata JSON string to map[string]interface{}.
		var metadata map[string]interface{}
		metadataStr := plan.Metadata.ValueString()
		if metadataStr != "" {
			if err := utils.JSONStringToMap(metadataStr, &metadata); err != nil {
				resp.Diagnostics.AddError(
					"Invalid metadata JSON",
					"Could not parse metadata as JSON: "+err.Error(),
				)
				return
			}
			user.SetMetadata(metadata)
		}
	}

	if !plan.UserType.IsNull() && !plan.UserType.IsUnknown() {
		user.SetUserType(plan.UserType.ValueString())
	}

	tflog.Info(ctx, "Creating a user", map[string]interface{}{"email": plan.Email.ValueString()})

	// Create new user.
	response, httpResponse, err := r.client.UsersAPI.UsersPost(ctx).UsersPostRequest(*user).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating user",
			"Could not create user, unexpected error: "+err.Error(),
		)
		tflog.Error(ctx, "Error creating user", utils.AddHTTPStatusCode(map[string]interface{}{"error": err.Error()}, httpResponse))
		return
	}

	if response.Data.Id == "" {
		resp.Diagnostics.AddError(
			"Error creating user",
			"Could not create user, unexpected error: missing ID",
		)
		tflog.Error(ctx, "Error creating user", map[string]interface{}{"error": "missing ID"})
		return
	}

	// Map response body to schema and populate Computed attribute values.
	plan.ID = types.StringValue(response.Data.Id)
	plan.LegacyID = types.StringValue(fmt.Sprintf("%.0f", response.Data.LegacyId))
	plan.Email = types.StringValue(response.Data.Email)
	plan.Active = types.BoolValue(response.Data.Active)

	// First_name and last_name are required, so they should always be in the response.
	if response.Data.FirstName.Get() != nil {
		plan.FirstName = types.StringValue(*response.Data.FirstName.Get())
	}
	if response.Data.LastName.Get() != nil {
		plan.LastName = types.StringValue(*response.Data.LastName.Get())
	}

	plan.CreatedAt = types.StringValue(response.Data.CreatedAt.String())
	
	if response.Data.LastActive.Get() != nil {
		plan.LastActive = types.StringValue(response.Data.LastActive.Get().String())
	} else {
		plan.LastActive = types.StringNull()
	}
	
	plan.IsAdmin = types.BoolValue(response.Data.IsAdmin)
	plan.UserType = types.StringValue(response.Data.UserType)
	plan.TwoFactorAuthEnabled = types.BoolValue(response.Data.TwoFactorAuthEnabled)

	// Handle metadata.
	switch {
	case len(response.Data.Metadata) > 0:
		metadataStr, err := utils.MapToJSONString(response.Data.Metadata)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error serializing metadata",
				"Could not serialize metadata: "+err.Error(),
			)
			return
		}
		plan.Metadata = types.StringValue(metadataStr)
	case plan.Metadata.IsNull():
		// Only set to null if it was null in the plan.
		plan.Metadata = types.StringNull()
	default:
		// If plan had metadata but response doesn't, set to empty JSON object.
		plan.Metadata = types.StringValue("{}")
	}

	// Set state to fully populated data.
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error creating user", map[string]interface{}{"error": "Could not set state"})
		return
	}

	tflog.Info(ctx, "User created", map[string]interface{}{"id": response.Data.Id, "email": response.Data.Email})
}

// Read a User resource.
func (r *userResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state userResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Use the ID from the state to read the user.
	userID := state.ID.ValueString()
	user, httpResponse, err := r.client.UsersAPI.UsersUserIdGet(ctx, userID).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			tflog.Info(ctx, "User not found", map[string]any{"userID": userID})
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error reading user",
			fmt.Sprintf("Could not read user with ID %s: %s", userID, err.Error()),
		)
		tflog.Error(ctx, "Error reading user", utils.AddHTTPStatusCode(map[string]any{"userID": userID, "error": err.Error()}, httpResponse))
		return
	}

	// Map the API response to the Terraform state.
	state.ID = types.StringValue(user.Data.Id)
	state.LegacyID = types.StringValue(fmt.Sprintf("%.0f", user.Data.LegacyId))
	state.Email = types.StringValue(user.Data.Email)
	state.Active = types.BoolValue(user.Data.Active)
	
	if user.Data.FirstName.Get() != nil {
		state.FirstName = types.StringValue(*user.Data.FirstName.Get())
	} else {
		state.FirstName = types.StringNull()
	}
	
	if user.Data.LastName.Get() != nil {
		state.LastName = types.StringValue(*user.Data.LastName.Get())
	} else {
		state.LastName = types.StringNull()
	}
	
	state.CreatedAt = types.StringValue(user.Data.CreatedAt.String())
	
	if user.Data.LastActive.Get() != nil {
		state.LastActive = types.StringValue(user.Data.LastActive.Get().String())
	} else {
		state.LastActive = types.StringNull()
	}
	
	state.IsAdmin = types.BoolValue(user.Data.IsAdmin)
	state.UserType = types.StringValue(user.Data.UserType)
	state.TwoFactorAuthEnabled = types.BoolValue(user.Data.TwoFactorAuthEnabled)

	// Handle metadata.
	if len(user.Data.Metadata) > 0 {
		metadataStr, err := utils.MapToJSONString(user.Data.Metadata)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error serializing metadata",
				"Could not serialize metadata: "+err.Error(),
			)
			return
		}
		state.Metadata = types.StringValue(metadataStr)
	} else {
		state.Metadata = types.StringNull()
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update a User resource.
func (r *userResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan userResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var state userResourceModel
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	userID := state.ID.ValueString()

	// Build patch operations based on changes.
	operations := []api.UsersUserIdPatchRequestOperationsInner{}

	// Check for changes and build patch operations
	// Note: Email cannot be changed via PATCH according to API restrictions
	
	if !plan.FirstName.Equal(state.FirstName) {
		replaceOp := api.NewReplaceOperation("replace", "/first_name")
		replaceOp.SetValue(plan.FirstName.ValueString())
		op := api.UsersUserIdPatchRequestOperationsInner{
			ReplaceOperation: replaceOp,
		}
		operations = append(operations, op)
	}

	if !plan.LastName.Equal(state.LastName) {
		replaceOp := api.NewReplaceOperation("replace", "/last_name")
		replaceOp.SetValue(plan.LastName.ValueString())
		op := api.UsersUserIdPatchRequestOperationsInner{
			ReplaceOperation: replaceOp,
		}
		operations = append(operations, op)
	}

	if !plan.Active.Equal(state.Active) {
		replaceOp := api.NewReplaceOperation("replace", "/active")
		replaceOp.SetValue(plan.Active.ValueBool())
		op := api.UsersUserIdPatchRequestOperationsInner{
			ReplaceOperation: replaceOp,
		}
		operations = append(operations, op)
	}

	if !plan.Metadata.Equal(state.Metadata) {
		var metadata map[string]interface{}
		if !plan.Metadata.IsNull() && plan.Metadata.ValueString() != "" {
			if err := utils.JSONStringToMap(plan.Metadata.ValueString(), &metadata); err != nil {
				resp.Diagnostics.AddError(
					"Invalid metadata JSON",
					"Could not parse metadata as JSON: "+err.Error(),
				)
				return
			}
		} else {
			metadata = map[string]interface{}{}
		}
		replaceOp := api.NewReplaceOperation("replace", "/metadata")
		replaceOp.SetValue(metadata)
		op := api.UsersUserIdPatchRequestOperationsInner{
			ReplaceOperation: replaceOp,
		}
		operations = append(operations, op)
	}

	if !plan.UserType.Equal(state.UserType) && !plan.UserType.IsNull() && plan.UserType.ValueString() != "" {
		replaceOp := api.NewReplaceOperation("replace", "/user_type")
		replaceOp.SetValue(plan.UserType.ValueString())
		op := api.UsersUserIdPatchRequestOperationsInner{
			ReplaceOperation: replaceOp,
		}
		operations = append(operations, op)
	}

	if len(operations) == 0 {
		tflog.Info(ctx, "No changes detected for user", map[string]any{"userID": userID})
		// No changes, just return current state
		diags = resp.State.Set(ctx, plan)
		resp.Diagnostics.Append(diags...)
		return
	}

	// Prepare the update payload
	patchRequest := api.NewUsersUserIdPatchRequest(operations)

	// Perform the update operation
	updateResponse, httpResponse, err := r.client.UsersAPI.UsersUserIdPatch(ctx, userID).UsersUserIdPatchRequest(*patchRequest).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating user",
			fmt.Sprintf("Could not update user with ID %s: %s", userID, err.Error()),
		)
		tflog.Error(ctx, "Error updating user", utils.AddHTTPStatusCode(map[string]any{"userID": userID, "error": err.Error()}, httpResponse))
		return
	}

	// Update the state with response data
	plan.ID = types.StringValue(updateResponse.Data.Id)
	plan.LegacyID = types.StringValue(fmt.Sprintf("%.0f", updateResponse.Data.LegacyId))
	plan.Email = types.StringValue(updateResponse.Data.Email)
	plan.Active = types.BoolValue(updateResponse.Data.Active)
	
	if updateResponse.Data.FirstName.Get() != nil {
		plan.FirstName = types.StringValue(*updateResponse.Data.FirstName.Get())
	}
	if updateResponse.Data.LastName.Get() != nil {
		plan.LastName = types.StringValue(*updateResponse.Data.LastName.Get())
	}
	
	plan.CreatedAt = types.StringValue(updateResponse.Data.CreatedAt.String())
	
	if updateResponse.Data.LastActive.Get() != nil {
		plan.LastActive = types.StringValue(updateResponse.Data.LastActive.Get().String())
	} else {
		plan.LastActive = types.StringNull()
	}
	
	plan.IsAdmin = types.BoolValue(updateResponse.Data.IsAdmin)
	plan.UserType = types.StringValue(updateResponse.Data.UserType)
	plan.TwoFactorAuthEnabled = types.BoolValue(updateResponse.Data.TwoFactorAuthEnabled)
	
	// Handle metadata
	if len(updateResponse.Data.Metadata) > 0 {
		metadataStr, err := utils.MapToJSONString(updateResponse.Data.Metadata)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error serializing metadata",
				"Could not serialize metadata: "+err.Error(),
			)
			return
		}
		plan.Metadata = types.StringValue(metadataStr)
	} else {
		plan.Metadata = types.StringNull()
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete a User resource.
// Note: Deleting a user in Retool disables them (active: false) rather than removing them entirely.
func (r *userResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state userResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	userID := state.ID.ValueString()

	// Delete (disable) the user
	httpResponse, err := r.client.UsersAPI.UsersUserIdDelete(ctx, userID).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			// User already doesn't exist, nothing to do
			tflog.Info(ctx, "User not found during delete, treating as already deleted", map[string]any{"userID": userID})
			return
		}
		resp.Diagnostics.AddError(
			"Error Deleting User",
			fmt.Sprintf("Could not delete user with ID %s: %s", userID, err.Error()),
		)
		tflog.Error(ctx, "Error deleting user", utils.AddHTTPStatusCode(map[string]any{"error": err.Error(), "userId": userID}, httpResponse))
		return
	}

	tflog.Info(ctx, "User deleted (disabled)", map[string]interface{}{"id": userID})
}

// ImportState allows importing of a User resource.
func (r *userResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute.
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

