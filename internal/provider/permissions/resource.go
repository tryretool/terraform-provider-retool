// Package permissions provides the implementation of the Permissions resource.
package permissions

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	_ resource.Resource              = &permissionResource{}
	_ resource.ResourceWithConfigure = &permissionResource{}
	// Note that unlike other resources, we don't implement ResourceWithImportState here, because there's not much to import for permission.
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
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

type permissionObjectModel struct {
	Id   types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
}

func (m permissionObjectModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":   types.StringType,
		"type": types.StringType,
	}
}

// A model for a single permission for a given subject.
type permissionModel struct {
	Object      types.Object `tfsdk:"object"`
	AccessLevel types.String `tfsdk:"access_level"`
}

func NewResource() resource.Resource {
	return &permissionResource{}
}

func (r *permissionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *permissionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_permissions"
}

func (r *permissionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
									Description: "The type of the object - app, folder, resource, or resource_configuration.",
									Validators: []validator.String{
										stringvalidator.OneOf("app", "folder", "resource", "resource_configuration"),
									},
								},
							},
						},
						"access_level": schema.StringAttribute{
							Required:    true,
							Description: "The access level of the permission.",
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

func createNewApiPermissionsSubject(subjectModel permissionSubjectModel) api.PermissionsListObjectsPostRequestSubject {
	subject := api.PermissionsListObjectsPostRequestSubject{}
	// ...OneOf represents a group subject, while OneOf1 represents a user subject.
	if subjectModel.Type.ValueString() == "group" {
		groupId, err := strconv.Atoi(subjectModel.Id.ValueString())
		if err != nil {
			return api.PermissionsListObjectsPostRequestSubject{}
		}
		floatGroupId := float32(groupId) // Our client uses float32 to represent "number" ids.
		subject.PermissionsListObjectsPostRequestSubjectOneOf = api.NewPermissionsListObjectsPostRequestSubjectOneOf("group", *api.NewNullableFloat32(&floatGroupId))
	} else if subjectModel.Type.ValueString() == "user" {
		subject.PermissionsListObjectsPostRequestSubjectOneOf1 = api.NewPermissionsListObjectsPostRequestSubjectOneOf1("user", subjectModel.Id.ValueString())
	}
	return subject
}

func createNewApiPermissionsObject(objectModel permissionObjectModel) api.PermissionsGrantPostRequestObject {
	object := api.PermissionsGrantPostRequestObject{
		PermissionsGrantPostRequestObjectOneOf: api.NewPermissionsGrantPostRequestObjectOneOf(objectModel.Type.ValueString(), objectModel.Id.ValueString()),
	}
	return object
}

func getPermissionId(subject permissionSubjectModel, object permissionObjectModel) string {
	return subject.Type.ValueString() + "|" + subject.Id.ValueString() + "|" + object.Type.ValueString() + "|" + object.Id.ValueString()
}

// Make an API call to revoke a permission.
func (r *permissionResource) revokePermission(ctx context.Context, subject permissionSubjectModel, permission permissionModel) diag.Diagnostics {
	var object permissionObjectModel

	diags := permission.Object.As(ctx, &object, basetypes.ObjectAsOptions{})
	if diags.HasError() {
		return diags
	}

	permissionId := getPermissionId(subject, object)
	tflog.Info(ctx, "Deleting permission", map[string]any{"id": permissionId})

	request := api.NewPermissionsRevokePostRequest(createNewApiPermissionsSubject(subject), createNewApiPermissionsObject(object))
	_, httpResponse, err := r.client.PermissionsAPI.PermissionsRevokePost(ctx).PermissionsRevokePostRequest(*request).Execute()
	if err != nil {
		diags.AddError(
			"Error deleting permission",
			"Could not delete permission with ID "+permissionId+": "+err.Error(),
		)
		tflog.Error(ctx, "Error Deleting Permission", utils.AddHttpStatusCode(map[string]any{"error": err.Error(), "id": permissionId}, httpResponse))
		return diags
	}
	tflog.Info(ctx, "Permission deleted", map[string]any{"id": permissionId})
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
	permissionId := getPermissionId(subject, object)

	// Generate API request body from plan.
	apiSubject := createNewApiPermissionsSubject(subject)
	apiObject := createNewApiPermissionsObject(object)

	grantRequest := api.NewPermissionsGrantPostRequest(apiSubject, apiObject, permission.AccessLevel.ValueString())

	tflog.Info(ctx, "Creating a permission", map[string]interface{}{"id": permissionId, "access level": permission.AccessLevel.ValueString()})

	// Grant the permission.
	_, httpResponse, err := r.client.PermissionsAPI.PermissionsGrantPost(ctx).PermissionsGrantPostRequest(*grantRequest).Execute()
	if err != nil {
		diags.AddError(
			"Error creating permission",
			"Could not create permission, unexpected error: "+err.Error(),
		)
		tflog.Error(ctx, "Error creating permission", utils.AddHttpStatusCode(map[string]interface{}{"error": err.Error()}, httpResponse))
		return diags
	}
	tflog.Info(ctx, "Permission created", map[string]interface{}{"id": permissionId})
	return diags
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

	diags = state.Subject.As(ctx, &stateSubject, basetypes.ObjectAsOptions{})
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var permissions []permissionModel

	subjectId := stateSubject.Id.ValueString() + "|" + stateSubject.Type.ValueString()

	// We'll need to get all the permissions for the given subject.
	for _, objectType := range []string{"app", "folder", "resource", "resource_configuration"} {
		request := api.NewPermissionsListObjectsPostRequest(createNewApiPermissionsSubject(stateSubject), objectType)

		tflog.Info(ctx, "Reading permission", map[string]interface{}{"subjectId": subjectId})

		permissionsResponse, httpResponse, err := r.client.PermissionsAPI.PermissionsListObjectsPost(ctx).PermissionsListObjectsPostRequest(*request).Execute()
		if err != nil {
			resp.Diagnostics.AddError(
				"Error reading permission",
				fmt.Sprintf("Could not read permissions for id: %s, object type: %s, error: %s", subjectId, objectType, err.Error()),
			)
			tflog.Error(ctx, "Error reading group", utils.AddHttpStatusCode(map[string]any{"permissionId": subjectId, "objectType": objectType, "error": err.Error()}, httpResponse))
			return
		}

		// Now let's populate the state with permissions based on our API response.
		for _, obj := range permissionsResponse.Data {
			var objId string
			var accessLevel string
			if obj.PermissionsListObjectsPost200ResponseDataInnerOneOf != nil {
				objId = obj.PermissionsListObjectsPost200ResponseDataInnerOneOf.Id
				accessLevel = obj.PermissionsListObjectsPost200ResponseDataInnerOneOf.AccessLevel
			}
			objValue := permissionObjectModel{
				Id:   types.StringValue(objId),
				Type: types.StringValue(objectType),
			}
			object, diags := types.ObjectValueFrom(ctx, objValue.AttributeTypes(), objValue)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
			permissions = append(permissions, permissionModel{
				Object:      object,
				AccessLevel: types.StringValue(accessLevel),
			})
		}
	}

	state.Permissions = permissions

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

type objectKey struct {
	Id   string
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
		permissionsByObject[objectKey{Id: obj.Id.ValueString(), Type: obj.Type.ValueString()}] = perm
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

	subjectId := stateSubject.Id.ValueString() + "|" + stateSubject.Type.ValueString()
	tflog.Info(ctx, "Revoking old permissions", map[string]any{"subject_id": subjectId})
	// Revoke old permissions.
	for _, perm := range diff.Revoke {
		diags = r.revokePermission(ctx, stateSubject, perm)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	tflog.Info(ctx, "Granting new permissions", map[string]any{"subject_id": subjectId})
	// Grant new permissions.
	for _, perm := range diff.Grant {
		diags = r.grantPermission(ctx, stateSubject, perm)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	tflog.Info(ctx, "Permissions updated", map[string]any{"subject_id": subjectId})

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
