package permissions

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// TestScreenPermissionObjectCreation tests that we can create screen permission objects correctly.
func TestScreenPermissionObjectCreation(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		model    permissionObjectModel
		expected string // Expected type.
	}{
		{
			name: "folder object",
			model: permissionObjectModel{
				ID:    types.StringValue("folder_123"),
				Type:  types.StringValue("folder"),
				AppID: types.StringNull(),
			},
			expected: "folder",
		},
		{
			name: "app object",
			model: permissionObjectModel{
				ID:    types.StringValue("app-uuid"),
				Type:  types.StringValue("app"),
				AppID: types.StringNull(),
			},
			expected: "app",
		},
		{
			name: "resource object",
			model: permissionObjectModel{
				ID:    types.StringValue("resource-uuid"),
				Type:  types.StringValue("resource"),
				AppID: types.StringNull(),
			},
			expected: "resource",
		},
		{
			name: "resource_configuration object",
			model: permissionObjectModel{
				ID:    types.StringValue("config-uuid"),
				Type:  types.StringValue("resource_configuration"),
				AppID: types.StringNull(),
			},
			expected: "resource_configuration",
		},
		{
			name: "screen object with app_id",
			model: permissionObjectModel{
				ID:    types.StringValue("screen-uuid"),
				Type:  types.StringValue("screen"),
				AppID: types.StringValue("app-uuid"),
			},
			expected: "screen",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
		obj := createNewAPIPermissionsObject(tc.model)

		// Verify the object was created (not empty).
		// We can't easily inspect the oneOf structure, but we can verify it's not nil
		// by checking if it has any populated fields
			hasValue := false

			if obj.Folder != nil {
				hasValue = true
				if tc.expected != "folder" {
					t.Errorf("Expected type %s but got folder", tc.expected)
				}
			}
			if obj.App != nil {
				hasValue = true
				if tc.expected != "app" {
					t.Errorf("Expected type %s but got app", tc.expected)
				}
			}
			if obj.Resource != nil {
				hasValue = true
				if tc.expected != "resource" {
					t.Errorf("Expected type %s but got resource", tc.expected)
				}
			}
			if obj.ResourceConfiguration != nil {
				hasValue = true
				if tc.expected != "resource_configuration" {
					t.Errorf("Expected type %s but got resource_configuration", tc.expected)
				}
			}
			if obj.Screen != nil {
				hasValue = true
				if tc.expected != "screen" {
					t.Errorf("Expected type %s but got screen", tc.expected)
				}
				// Verify screen has the app_id set
				if obj.Screen.GetAppId() != tc.model.AppID.ValueString() {
					t.Errorf("Expected app_id %s but got %s", tc.model.AppID.ValueString(), obj.Screen.GetAppId())
				}
			}

			if !hasValue {
				t.Error("Created object has no populated fields")
			}
		})
	}
}

// TestScreenSubjectCreation tests that we can create permission subjects correctly.
func TestScreenSubjectCreation(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		model    permissionSubjectModel
		expected string // Expected type.
	}{
		{
			name: "user subject",
			model: permissionSubjectModel{
				ID:   types.StringValue("user-uuid"),
				Type: types.StringValue("user"),
			},
			expected: "user",
		},
		{
			name: "group subject",
			model: permissionSubjectModel{
				ID:   types.StringValue("123"),
				Type: types.StringValue("group"),
			},
			expected: "group",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			subject := createNewAPIPermissionsSubject(tc.model)

			// Verify the subject was created
			hasValue := false

			if subject.User != nil {
				hasValue = true
				if tc.expected != "user" {
					t.Errorf("Expected type %s but got user", tc.expected)
				}
			}
			if subject.Group != nil {
				hasValue = true
				if tc.expected != "group" {
					t.Errorf("Expected type %s but got group", tc.expected)
				}
			}

			if !hasValue {
				t.Error("Created subject has no populated fields")
			}
		})
	}
}

// TestPermissionObjectModelAttributeTypes verifies the attribute types include app_id.
func TestPermissionObjectModelAttributeTypes(t *testing.T) {
	t.Parallel()

	model := permissionObjectModel{}
	attrs := model.AttributeTypes()

	expectedAttrs := []string{"id", "type", "app_id"}
	for _, attr := range expectedAttrs {
		if _, ok := attrs[attr]; !ok {
			t.Errorf("Expected attribute %s not found in AttributeTypes", attr)
		}
	}

	if len(attrs) != len(expectedAttrs) {
		t.Errorf("Expected %d attributes but got %d", len(expectedAttrs), len(attrs))
	}
}

// TestScreenAPIObjectConstruction verifies we can construct a Screen API object.
func TestScreenAPIObjectConstruction(t *testing.T) {
	t.Parallel()

	screenID := "screen-uuid-123"
	appID := "app-uuid-456"

	screen := api.NewScreen("screen", screenID, appID)

	if screen == nil {
		t.Fatal("Expected non-nil screen object")
	}

	if screen.GetType() != "screen" {
		t.Errorf("Expected type 'screen' but got '%s'", screen.GetType())
	}

	if screen.GetId() != screenID {
		t.Errorf("Expected id '%s' but got '%s'", screenID, screen.GetId())
	}

	if screen.GetAppId() != appID {
		t.Errorf("Expected appId '%s' but got '%s'", appID, screen.GetAppId())
	}
}

// TestScreenPermissionObjectConversion verifies screen objects convert correctly.
func TestScreenPermissionObjectConversion(t *testing.T) {
	t.Parallel()

	screenID := "screen-uuid-789"
	appID := "app-uuid-012"

	screen := api.NewScreen("screen", screenID, appID)
	permObj := api.ScreenAsPermissionsGrantPostRequestObject(screen)

	if permObj.Screen == nil {
		t.Fatal("Expected Screen field to be populated")
	}

	if permObj.Screen.GetId() != screenID {
		t.Errorf("Expected screen id '%s' but got '%s'", screenID, permObj.Screen.GetId())
	}

	if permObj.Screen.GetAppId() != appID {
		t.Errorf("Expected screen appId '%s' but got '%s'", appID, permObj.Screen.GetAppId())
	}
}

// TestGetPermissionID verifies the permission ID generation includes all object types.
func TestGetPermissionID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		subject  permissionSubjectModel
		object   permissionObjectModel
		expected string
	}{
		{
			name: "user with folder",
			subject: permissionSubjectModel{
				ID:   types.StringValue("user-123"),
				Type: types.StringValue("user"),
			},
			object: permissionObjectModel{
				ID:   types.StringValue("folder-456"),
				Type: types.StringValue("folder"),
			},
			expected: "user|user-123|folder|folder-456",
		},
		{
			name: "group with screen",
			subject: permissionSubjectModel{
				ID:   types.StringValue("789"),
				Type: types.StringValue("group"),
			},
			object: permissionObjectModel{
				ID:    types.StringValue("screen-abc"),
				Type:  types.StringValue("screen"),
				AppID: types.StringValue("app-def"),
			},
			expected: "group|789|screen|screen-abc",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			id := getPermissionID(tc.subject, tc.object)
			if id != tc.expected {
				t.Errorf("Expected permission ID '%s' but got '%s'", tc.expected, id)
			}
		})
	}
}

// TestPermissionObjectModelWithContext verifies we can convert models with context.
func TestPermissionObjectModelWithContext(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	// Test screen object with app_id
	screenObj := permissionObjectModel{
		ID:    types.StringValue("screen-123"),
		Type:  types.StringValue("screen"),
		AppID: types.StringValue("app-456"),
	}

	objValue, diags := types.ObjectValueFrom(ctx, screenObj.AttributeTypes(), screenObj)
	if diags.HasError() {
		t.Fatalf("Failed to create object value: %v", diags)
	}

	// Convert back
	var retrieved permissionObjectModel
	diags = objValue.As(ctx, &retrieved, basetypes.ObjectAsOptions{})
	if diags.HasError() {
		t.Fatalf("Failed to retrieve object value: %v", diags)
	}

	if retrieved.ID.ValueString() != "screen-123" {
		t.Errorf("Expected ID 'screen-123' but got '%s'", retrieved.ID.ValueString())
	}

	if retrieved.Type.ValueString() != "screen" {
		t.Errorf("Expected Type 'screen' but got '%s'", retrieved.Type.ValueString())
	}

	if retrieved.AppID.ValueString() != "app-456" {
		t.Errorf("Expected AppID 'app-456' but got '%s'", retrieved.AppID.ValueString())
	}
}
