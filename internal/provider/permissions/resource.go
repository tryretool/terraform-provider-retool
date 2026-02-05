// Package permissions provides the implementation of the Permissions resource.
package permissions

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// Ensure GroupResource implements the tfsdk.Resource interface.
var (
	_ resource.Resource                = &permissionResource{}
	_ resource.ResourceWithConfigure   = &permissionResource{}
	_ resource.ResourceWithImportState = &permissionResource{}
)

type permissionResource struct {
	client *api.APIClient
}

// The key is a combination of subject type and id.
type permissionsResourceModel struct {
	Subject     types.Object      `tfsdk:"subject"`
	Permissions []permissionModel `tfsdk:"permissions"`
}

type permissionSubjectModel struct {
	ID   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

type permissionObjectModel struct {
	ID    types.String `tfsdk:"id"`
	Type  types.String `tfsdk:"type"`
	AppID types.String `tfsdk:"app_id"`
}

func (m permissionObjectModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":     types.StringType,
		"type":   types.StringType,
		"app_id": types.StringType,
	}
}

// A model for a single permission for a given subject.
type permissionModel struct {
	Object      types.Object `tfsdk:"object"`
	AccessLevel types.String `tfsdk:"access_level"`
}

// Create new Permissions resource.
func NewResource() resource.Resource {
	return &permissionResource{}
}

func (r *permissionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *permissionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_permissions"
}

func (r *permissionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Permissions resource can be used to grant a user or a permission group access to an object: app, folder, resource, resource configuration, or screen.",
		Attributes: map[string]schema.Attribute{
			"subject": schema.SingleNestedAttribute{
				Required:    true,
				Description: "Permission subject (who can perform the action).",
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required:    true,
						Description: "The ID of the subject.",
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(), // Changing the subject of permission requires replacing it.
						},
					},
					"type": schema.StringAttribute{
						Required:    true,
						Description: "The type of the subject - user or group.",
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(), // Changing the subject of permission requires replacing it.
						},
						Validators: []validator.String{
							stringvalidator.OneOf("user", "group"),
						},
					},
				},
			},
			"permissions": schema.SetNestedAttribute{
				Required:    true,
				Description: "A set of permissions.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"object": schema.SingleNestedAttribute{
							Required:    true,
							Description: "Permission object (that the action is performed on).",
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Required:    true,
									Description: "The ID of the object.",
								},
								"type": schema.StringAttribute{
									Required:    true,
									Description: "The type of the object - app, folder, resource, resource_configuration, or screen.",
									Validators: []validator.String{
										stringvalidator.OneOf("app", "folder", "resource", "resource_configuration", "screen"),
									},
								},
								"app_id": schema.StringAttribute{
									Optional:    true,
									Description: "The app ID (required when type is 'screen').",
								},
							},
						},
						"access_level": schema.StringAttribute{
							Required:    true,
							Description: "The access level of the permission. Accepted values are 'own', 'edit', 'use'.",
							Validators: []validator.String{
								stringvalidator.OneOf("own", "edit", "use"),
							},
						},
					},
				},
			},
		},
	}
}

func createNewAPIPermissionsSubject(subjectModel permissionSubjectModel) api.PermissionsListObjectsPostRequestSubject {
	subject := api.PermissionsListObjectsPostRequestSubject{}
	// ...OneOf represents a group subject, while OneOf1 represents a user subject.
	if subjectModel.Type.ValueString() == "group" {
		groupID, err := strconv.Atoi(subjectModel.ID.ValueString())
		if err != nil {
			return api.PermissionsListObjectsPostRequestSubject{}
		}
		floatGroupID := float32(groupID) // Our client uses float32 to represent "number" ids.
		groupSubject := api.NewGroup("group", *api.NewNullableFloat32(&floatGroupID))
		subject = api.GroupAsPermissionsListObjectsPostRequestSubject(groupSubject)
	} else if subjectModel.Type.ValueString() == "user" {
		userSubject := api.NewUser("user", subjectModel.ID.ValueString())
		subject = api.UserAsPermissionsListObjectsPostRequestSubject(userSubject)
	}
	return subject
}

func createNewAPIPermissionsObject(objectModel permissionObjectModel) api.PermissionsGrantPostRequestObject {
	objType := objectModel.Type.ValueString()
	objID := objectModel.ID.ValueString()

	// Use the appropriate constructor and wrapper function based on object type.
	switch objType {
	case "app":
		app := api.NewApp(objType, objID)
		return api.AppAsPermissionsGrantPostRequestObject(app)
	case "folder":
		folder := api.NewFolder(objType, objID)
		return api.FolderAsPermissionsGrantPostRequestObject(folder)
	case "resource":
		resource := api.NewResource(objType, objID)
		return api.ResourceAsPermissionsGrantPostRequestObject(resource)
	case "resource_configuration":
		resourceConfig := api.NewResourceConfiguration(objType, objID)
		return api.ResourceConfigurationAsPermissionsGrantPostRequestObject(resourceConfig)
	case "screen":
		appID := objectModel.AppID.ValueString()
		screen := api.NewScreen(objType, objID, appID)
		return api.ScreenAsPermissionsGrantPostRequestObject(screen)
	default:
		// Fallback: return empty object if type is unknown.
		return api.PermissionsGrantPostRequestObject{}
	}
}

func getPermissionID(subject permissionSubjectModel, object permissionObjectModel) string {
	return subject.Type.ValueString() + "|" + subject.ID.ValueString() + "|" + object.Type.ValueString() + "|" + object.ID.ValueString()
}

// Make an API call to revoke a permission.
func (r *permissionResource) revokePermission(ctx context.Context, subject permissionSubjectModel, permission permissionModel) diag.Diagnostics {
	var object permissionObjectModel

	diags := permission.Object.As(ctx, &object, basetypes.ObjectAsOptions{})
	if diags.HasError() {
		return diags
	}

	permissionID := getPermissionID(subject, object)
	tflog.Info(ctx, "Deleting permission", map[string]any{"id": permissionID})

	request := api.NewPermissionsRevokePostRequest(createNewAPIPermissionsSubject(subject), createNewAPIPermissionsObject(object))
	_, httpResponse, err := r.client.PermissionsAPI.PermissionsRevokePost(ctx).PermissionsRevokePostRequest(*request).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			// If the permission does not exist, we can ignore the error.
			tflog.Info(ctx, "Permission not found", map[string]any{"id": permissionID})
			return diags
		}
		diags.AddError(
			"Error deleting permission",
			"Could not delete permission with ID "+permissionID+": "+err.Error(),
		)
		tflog.Error(ctx, "Error Deleting Permission", utils.AddHTTPStatusCode(map[string]any{"error": err.Error(), "id": permissionID}, httpResponse))
		return diags
	}
	tflog.Info(ctx, "Permission deleted", map[string]any{"id": permissionID})
	return diags
}

// Make an API call to grant a permission.
func (r *permissionResource) grantPermission(ctx context.Context, subject permissionSubjectModel, permission permissionModel) diag.Diagnostics {
	var object permissionObjectModel

	diags := permission.Object.As(ctx, &object, basetypes.ObjectAsOptions{})
	if diags.HasError() {
		return diags
	}

	// Only used for logging.
	permissionID := getPermissionID(subject, object)

	// Generate API request body from plan.
	apiSubject := createNewAPIPermissionsSubject(subject)
	apiObject := createNewAPIPermissionsObject(object)

	grantRequest := api.NewPermissionsGrantPostRequest(apiSubject, apiObject, permission.AccessLevel.ValueString())

	tflog.Info(ctx, "Creating a permission", map[string]interface{}{"id": permissionID, "access level": permission.AccessLevel.ValueString()})

	// Grant the permission.
	_, httpResponse, err := r.client.PermissionsAPI.PermissionsGrantPost(ctx).PermissionsGrantPostRequest(*grantRequest).Execute()
	if err != nil {
		diags.AddError(
			"Error creating permission",
			"Could not create permission, unexpected error: "+err.Error(),
		)
		tflog.Error(ctx, "Error creating permission", utils.AddHTTPStatusCode(map[string]interface{}{"error": err.Error()}, httpResponse))
		return diags
	}
	tflog.Info(ctx, "Permission created", map[string]interface{}{"id": permissionID})
	return diags
}

func (r *permissionResource) fetchPermissionsForSubject(ctx context.Context, subject permissionSubjectModel) ([]permissionModel, diag.Diagnostics) {
	var permissions []permissionModel
	var allDiags diag.Diagnostics

	subjectID := subject.ID.ValueString() + "|" + subject.Type.ValueString()

	// Note: "screen" is omitted from this list because while the grant API supports screens,
	// the listObjects API does not yet support them (as of API v3.323.0).
	// Screen permissions are tracked in state but not refreshed from the API.
	for _, objectType := range []string{"app", "folder", "resource", "resource_configuration"} {
		request := api.NewPermissionsListObjectsPostRequest(createNewAPIPermissionsSubject(subject), objectType)

		tflog.Info(ctx, "Fetching permissions", map[string]interface{}{"subjectId": subjectID, "objectType": objectType})

		permissionsResponse, httpResponse, err := r.client.PermissionsAPI.PermissionsListObjectsPost(ctx).PermissionsListObjectsPostRequest(*request).Execute()
		if err != nil {
			allDiags.AddError(
				"Error reading permission",
				fmt.Sprintf("Could not read permissions for id: %s, object type: %s, error: %s", subjectID, objectType, err.Error()),
			)
			tflog.Error(ctx, "Error reading group", utils.AddHTTPStatusCode(map[string]any{"permissionId": subjectID, "objectType": objectType, "error": err.Error()}, httpResponse))
			return nil, allDiags
		}

		// Now let's populate the state with permissions based on our API response.
		for _, obj := range permissionsResponse.Data {
			var objID string
			var accessLevel string
			var appID string

			// Check which variant is populated (folder, app, screen, resource, or resource_configuration).
			switch {
			case obj.PermissionsListObjectsPost200ResponseDataInnerAnyOf != nil:
				// Folder.
				objID = obj.PermissionsListObjectsPost200ResponseDataInnerAnyOf.Id
				accessLevel = obj.PermissionsListObjectsPost200ResponseDataInnerAnyOf.AccessLevel
			case obj.PermissionsListObjectsPost200ResponseDataInnerAnyOf1 != nil:
				// App.
				objID = obj.PermissionsListObjectsPost200ResponseDataInnerAnyOf1.Id
				accessLevel = obj.PermissionsListObjectsPost200ResponseDataInnerAnyOf1.AccessLevel
			case obj.PermissionsListObjectsPost200ResponseDataInnerAnyOf2 != nil:
				// Screen.
				objID = obj.PermissionsListObjectsPost200ResponseDataInnerAnyOf2.Id
				accessLevel = obj.PermissionsListObjectsPost200ResponseDataInnerAnyOf2.AccessLevel
				// Note: Screen responses don't include appId, so we can't populate it here
			case obj.PermissionsListObjectsPost200ResponseDataInnerAnyOf3 != nil:
				// Resource.
				objID = obj.PermissionsListObjectsPost200ResponseDataInnerAnyOf3.Id
				accessLevel = obj.PermissionsListObjectsPost200ResponseDataInnerAnyOf3.AccessLevel
			case obj.PermissionsListObjectsPost200ResponseDataInnerAnyOf4 != nil:
				// Resource Configuration.
				objID = obj.PermissionsListObjectsPost200ResponseDataInnerAnyOf4.Id
				accessLevel = obj.PermissionsListObjectsPost200ResponseDataInnerAnyOf4.AccessLevel
			default:
				// None of the variants matched - log and skip this permission.
				tflog.Warn(ctx, "Permission object did not match any known type variant, skipping", map[string]interface{}{"objectType": objectType})
				continue
			}

			// Skip if objID or accessLevel are empty.
			if objID == "" || accessLevel == "" {
				tflog.Warn(ctx, "Permission object has empty ID or access level, skipping", map[string]interface{}{"objectType": objectType, "objID": objID, "accessLevel": accessLevel})
				continue
			}

			objValue := permissionObjectModel{
				ID:   types.StringValue(objID),
				Type: types.StringValue(objectType),
			}
			// Set appID if available (only for screens in request, but not returned in response)
			if appID != "" {
				objValue.AppID = types.StringValue(appID)
			} else {
				objValue.AppID = types.StringNull()
			}
			object, diags := types.ObjectValueFrom(ctx, objValue.AttributeTypes(), objValue)
			allDiags.Append(diags...)
			if allDiags.HasError() {
				return nil, allDiags
			}
			permissions = append(permissions, permissionModel{
				Object:      object,
				AccessLevel: types.StringValue(accessLevel),
			})
		}
	}
	return permissions, allDiags
}

func (r *permissionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan.
	var plan permissionsResourceModel
	var planSubject permissionSubjectModel

	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = plan.Subject.As(ctx, &planSubject, basetypes.ObjectAsOptions{})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, planPermission := range plan.Permissions {
		diags = r.grantPermission(ctx, planSubject, planPermission)

		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	// Set state to fully populated data.
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error creating permissions", map[string]interface{}{"error": "Could not set state"})
		return
	}
}

func (r *permissionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state permissionsResourceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var stateSubject permissionSubjectModel
	var managedPermissionKeys = make(map[string]bool)

	for _, permission := range state.Permissions {
		var obj permissionObjectModel
		diags := permission.Object.As(ctx, &obj, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			return
		}
		key := obj.ID.ValueString() + "|" + obj.Type.ValueString()
		managedPermissionKeys[key] = true
	}

	diags = state.Subject.As(ctx, &stateSubject, basetypes.ObjectAsOptions{})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	allPermissions, diags := r.fetchPermissionsForSubject(ctx, stateSubject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var filteredPermissions []permissionModel
	for _, perm := range allPermissions {
		var obj permissionObjectModel
		diags := perm.Object.As(ctx, &obj, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
		key := obj.ID.ValueString() + "|" + obj.Type.ValueString()
		if managedPermissionKeys[key] {
			filteredPermissions = append(filteredPermissions, perm)
		}
	}

	// Add screen permissions from state since the listObjects API doesn't support screens yet.
	// Screen permissions can be granted but not listed, so we preserve them from state.
	for _, statePerm := range state.Permissions {
		var obj permissionObjectModel
		diags := statePerm.Object.As(ctx, &obj, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			resp.Diagnostics.Append(diags...)
			return
		}
		if obj.Type.ValueString() == "screen" {
			filteredPermissions = append(filteredPermissions, statePerm)
		}
	}

	state.Permissions = filteredPermissions

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

type objectKey struct {
	ID   string
	Type string
}

// Converts a list of permissions to a map of permissions by object, for quick lookup.
func mapPermissionsByOjbect(ctx context.Context, permissions []permissionModel) (map[objectKey]permissionModel, diag.Diagnostics) {
	permissionsByObject := make(map[objectKey]permissionModel)
	var allDiags diag.Diagnostics
	for _, perm := range permissions {
		var obj permissionObjectModel
		diags := perm.Object.As(ctx, &obj, basetypes.ObjectAsOptions{})
		allDiags.Append(diags...)
		if diags.HasError() {
			return nil, allDiags
		}
		permissionsByObject[objectKey{ID: obj.ID.ValueString(), Type: obj.Type.ValueString()}] = perm
	}
	return permissionsByObject, allDiags
}

type permissionDiff struct {
	Revoke []permissionModel
	Grant  []permissionModel
}

func computePermissionDiff(ctx context.Context, statePermissions []permissionModel, planPermissions []permissionModel) (permissionDiff, diag.Diagnostics) {
	var diff permissionDiff

	// Maps to hold the sets of permissions for quick lookup.
	stateMap, diags := mapPermissionsByOjbect(ctx, statePermissions)
	if diags.HasError() {
		return diff, diags
	}
	planMap, diags := mapPermissionsByOjbect(ctx, planPermissions)
	if diags.HasError() {
		return diff, diags
	}

	// Find permissions to revoke (in state but not in plan).
	for key, statePerm := range stateMap {
		if _, found := planMap[key]; !found {
			diff.Revoke = append(diff.Revoke, statePerm)
		}
	}

	// Find permissions to grant (in plan but not in state).
	for key, planPerm := range planMap {
		if statePerm, found := stateMap[key]; !found || statePerm.AccessLevel.ValueString() != planPerm.AccessLevel.ValueString() {
			diff.Grant = append(diff.Grant, planPerm)
		}
	}

	return diff, nil
}

func (r *permissionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var state permissionsResourceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var stateSubject permissionSubjectModel

	diags = state.Subject.As(ctx, &stateSubject, basetypes.ObjectAsOptions{})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var plan permissionsResourceModel

	diags = req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diff, diags := computePermissionDiff(ctx, state.Permissions, plan.Permissions)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	subjectID := stateSubject.ID.ValueString() + "|" + stateSubject.Type.ValueString()
	tflog.Info(ctx, "Revoking old permissions", map[string]any{"subject_id": subjectID})
	// Revoke old permissions.
	for _, perm := range diff.Revoke {
		diags = r.revokePermission(ctx, stateSubject, perm)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	tflog.Info(ctx, "Granting new permissions", map[string]any{"subject_id": subjectID})
	// Grant new permissions.
	for _, perm := range diff.Grant {
		diags = r.grantPermission(ctx, stateSubject, perm)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	tflog.Info(ctx, "Permissions updated", map[string]any{"subject_id": subjectID})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *permissionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state permissionsResourceModel

	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var stateSubject permissionSubjectModel

	diags = state.Subject.As(ctx, &stateSubject, basetypes.ObjectAsOptions{})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	for _, statePermission := range state.Permissions {
		diags = r.revokePermission(ctx, stateSubject, statePermission)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
}

func (r *permissionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// We expect id to be in the format 'subjectType|subjectId'.
	parts := strings.Split(req.ID, "|")
	if len(parts) != 2 {
		resp.Diagnostics.AddError("Invalid import ID", "Import ID must be in the format 'subjectType|id'.")
		return
	}
	subjType := parts[0]
	id := parts[1]
	subject := permissionSubjectModel{
		ID:   types.StringValue(id),
		Type: types.StringValue(subjType),
	}
	resp.State.SetAttribute(ctx, path.Root("subject"), subject)

	allPermissions, diags := r.fetchPermissionsForSubject(ctx, subject)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.State.SetAttribute(ctx, path.Root("permissions"), allPermissions)
}
