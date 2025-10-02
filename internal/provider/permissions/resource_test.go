package permissions_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

const testPermissionsConfig = `
resource "retool_group" "test_group" {
	name = "tf-acc-test-group"
}

resource "retool_folder" "test_folder1" {
	name = "tf-acc-test-folder1"
	folder_type = "app"
}

resource "retool_folder" "test_folder2" {
	name = "tf-acc-test-folder2"
	parent_folder_id = retool_folder.test_folder1.id
	folder_type = "app"
}

resource "retool_permissions" "test_permissions" {
	subject = {
		type = "group"
		id = retool_group.test_group.id
	}
	permissions = [
		{
			object = {
				type = "folder"
				id = retool_folder.test_folder1.id
			}
			access_level = "use"
		},
	]
}
`

const testUpdatedPermissionsConfig = `
resource "retool_group" "test_group" {
	name = "tf-acc-test-group"
}

resource "retool_folder" "test_folder1" {
	name = "tf-acc-test-folder1"
	folder_type = "app"
}

resource "retool_folder" "test_folder2" {
	name = "tf-acc-test-folder2"
	parent_folder_id = retool_folder.test_folder1.id
	folder_type = "app"
}

resource "retool_permissions" "test_permissions" {
	subject = {
		type = "group"
		id = retool_group.test_group.id
	}
	permissions = [
		{
			object = {
				type = "folder"
				id = retool_folder.test_folder2.id
			}
			access_level = "own"
		},
	]
}
`

// Even though we don't define sweepers for permissions, we need to define this function so that the tests are not run when other resources are sweeped.
func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func importStateIDFunc(state *terraform.State) (string, error) {
	permissions, ok := state.RootModule().Resources["retool_permissions.test_permissions"]
	if !ok {
		return "", fmt.Errorf("Resource not found")
	}
	return "group|" + permissions.Primary.Attributes["subject.id"], nil
}

func TestAccPermissions(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Read and Create.
			{
				Config: testPermissionsConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("retool_permissions.test_permissions", "subject.type", "group"),
					resource.TestCheckResourceAttrPair("retool_permissions.test_permissions", "subject.id", "retool_group.test_group", "id"),
					resource.TestCheckResourceAttr("retool_permissions.test_permissions", "permissions.#", "1"),
					resource.TestCheckResourceAttr("retool_permissions.test_permissions", "permissions.0.object.type", "folder"),
					resource.TestCheckResourceAttrPair("retool_permissions.test_permissions", "permissions.0.object.id", "retool_folder.test_folder1", "id"),
					resource.TestCheckResourceAttr("retool_permissions.test_permissions", "permissions.0.access_level", "use"),
				),
			},
			// Import state.
			{
				ResourceName:                         "retool_permissions.test_permissions",
				ImportStateIdFunc:                    importStateIDFunc,
				ImportState:                          true,
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "subject.id",
			},
			// Update and Read.
			{
				Config: testUpdatedPermissionsConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("retool_permissions.test_permissions", "subject.type", "group"),
					resource.TestCheckResourceAttrPair("retool_permissions.test_permissions", "subject.id", "retool_group.test_group", "id"),
					resource.TestCheckResourceAttr("retool_permissions.test_permissions", "permissions.#", "1"),
					resource.TestCheckResourceAttr("retool_permissions.test_permissions", "permissions.0.object.type", "folder"),
					resource.TestCheckResourceAttrPair("retool_permissions.test_permissions", "permissions.0.object.id", "retool_folder.test_folder2", "id"),
					resource.TestCheckResourceAttr("retool_permissions.test_permissions", "permissions.0.access_level", "own"),
				),
			},
		},
	})
}


// Helper function to get an API client for direct API calls in tests
func getTestAPIClient(t *testing.T) *api.APIClient {
	client, err := acctest.SweeperClient()
	if err != nil {
		t.Fatalf("Failed to create API client: %s", err)
	}
	return client
}

// Helper function to create an unmanaged permission via direct API call
func createUnmanagedPermission(t *testing.T, client *api.APIClient, subjectType string, subjectID string, objectType string, objectID string, accessLevel string) {
	var apiSubject api.PermissionsListObjectsPostRequestSubject

	if subjectType == "group" {
		groupID, err := strconv.Atoi(subjectID)
		if err != nil {
			t.Fatalf("Failed to convert group ID to int: %s", err)
		}
		floatGroupID := float32(groupID)
		apiSubject.PermissionsListObjectsPostRequestSubjectOneOf = api.NewPermissionsListObjectsPostRequestSubjectOneOf("group", *api.NewNullableFloat32(&floatGroupID))
	} else if subjectType == "user" {
		apiSubject.PermissionsListObjectsPostRequestSubjectOneOf1 = api.NewPermissionsListObjectsPostRequestSubjectOneOf1("user", subjectID)
	}

	apiObject := api.PermissionsGrantPostRequestObject{
		PermissionsGrantPostRequestObjectOneOf: api.NewPermissionsGrantPostRequestObjectOneOf(objectType, objectID),
	}

	grantRequest := api.NewPermissionsGrantPostRequest(apiSubject, apiObject, accessLevel)

	_, httpResponse, err := client.PermissionsAPI.PermissionsGrantPost(context.Background()).PermissionsGrantPostRequest(*grantRequest).Execute()
	if err != nil {
		t.Fatalf("Failed to create unmanaged permission: %s, HTTP Status: %d", err, httpResponse.StatusCode)
	}
}

// Helper function to verify a permission exists via API
func checkPermissionExists(t *testing.T, client *api.APIClient, subjectType string, subjectID string, objectType string, objectID string) bool {
	var apiSubject api.PermissionsListObjectsPostRequestSubject

	if subjectType == "group" {
		groupID, err := strconv.Atoi(subjectID)
		if err != nil {
			t.Fatalf("Failed to convert group ID to int: %s", err)
		}
		floatGroupID := float32(groupID)
		apiSubject.PermissionsListObjectsPostRequestSubjectOneOf = api.NewPermissionsListObjectsPostRequestSubjectOneOf("group", *api.NewNullableFloat32(&floatGroupID))
	} else if subjectType == "user" {
		apiSubject.PermissionsListObjectsPostRequestSubjectOneOf1 = api.NewPermissionsListObjectsPostRequestSubjectOneOf1("user", subjectID)
	}

	request := api.NewPermissionsListObjectsPostRequest(apiSubject, objectType)

	permissionsResponse, _, err := client.PermissionsAPI.PermissionsListObjectsPost(context.Background()).PermissionsListObjectsPostRequest(*request).Execute()
	if err != nil {
		t.Fatalf("Failed to list permissions: %s", err)
	}

	// Check if the object exists in the response
	for _, obj := range permissionsResponse.Data {
		if obj.PermissionsListObjectsPost200ResponseDataInnerOneOf != nil {
			if obj.PermissionsListObjectsPost200ResponseDataInnerOneOf.Id == objectID {
				return true
			}
		}
	}

	return false
}

// Verify that removing a managed permission doesn't delete unmanaged permissions
func TestAccPermissions_ManagedDeletion(t *testing.T) {
	client := getTestAPIClient(t)
	var groupID, folder1ID, folder3ID string

	
	configWithManagedPerm := `
resource "retool_group" "test_group" {
	name = "tf-acc-test-group-managed-del"
}

resource "retool_folder" "test_folder1" {
	name = "tf-acc-test-folder-managed-1"
	folder_type = "app"
}

resource "retool_folder" "test_folder3" {
	name = "tf-acc-test-folder-unmanaged-3"
	folder_type = "app"
}

resource "retool_permissions" "test_permissions" {
	subject = {
		type = "group"
		id = retool_group.test_group.id
	}
	permissions = [
		{
			object = {
				type = "folder"
				id = retool_folder.test_folder1.id
			}
			access_level = "use"
		},
	]
}
`

	
	emptyConfig := `
resource "retool_group" "test_group" {
	name = "tf-acc-test-group-managed-del"
}

resource "retool_folder" "test_folder1" {
	name = "tf-acc-test-folder-managed-1"
	folder_type = "app"
}

resource "retool_folder" "test_folder3" {
	name = "tf-acc-test-folder-unmanaged-3"
	folder_type = "app"
}
`

	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: configWithManagedPerm,
				Check: resource.ComposeAggregateTestCheckFunc(

					func(s *terraform.State) error {
						groupRes := s.RootModule().Resources["retool_group.test_group"]
						groupID = groupRes.Primary.Attributes["id"]

						folder1Res := s.RootModule().Resources["retool_folder.test_folder1"]
						folder1ID = folder1Res.Primary.Attributes["id"]

						folder3Res := s.RootModule().Resources["retool_folder.test_folder3"]
						folder3ID = folder3Res.Primary.Attributes["id"]

						// Create unmanaged permission via direct API call
						createUnmanagedPermission(t, client, "group", groupID, "folder", folder3ID, "edit")

						return nil
					},
					
					resource.TestCheckResourceAttr("retool_permissions.test_permissions", "permissions.#", "1"),
					resource.TestCheckResourceAttr("retool_permissions.test_permissions", "permissions.0.access_level", "use"),
				),
			},
			// Remove the permissions resource and verify unmanaged permission still exists
			{
				Config: emptyConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify unmanaged permission (folder3) still exists via API
					func(s *terraform.State) error {
						exists := checkPermissionExists(t, client, "group", groupID, "folder", folder3ID)
						if !exists {
							return fmt.Errorf("Unmanaged permission for folder3 was incorrectly deleted")
						}
						return nil
					},
					// Verify managed permission (folder1) was deleted via API
					func(s *terraform.State) error {
						exists := checkPermissionExists(t, client, "group", groupID, "folder", folder1ID)
						if exists {
							return fmt.Errorf("Managed permission for folder1 was not deleted")
						}
						return nil
					},
				),
			},
		},
	})

	// Cleanup
	if groupID != "" && folder3ID != "" {
		var apiSubject api.PermissionsListObjectsPostRequestSubject
		groupIDInt, _ := strconv.Atoi(groupID)
		floatGroupID := float32(groupIDInt)
		apiSubject.PermissionsListObjectsPostRequestSubjectOneOf = api.NewPermissionsListObjectsPostRequestSubjectOneOf("group", *api.NewNullableFloat32(&floatGroupID))

		apiObject := api.PermissionsGrantPostRequestObject{
			PermissionsGrantPostRequestObjectOneOf: api.NewPermissionsGrantPostRequestObjectOneOf("folder", folder3ID),
		}

		revokeRequest := api.NewPermissionsRevokePostRequest(apiSubject, apiObject)
		client.PermissionsAPI.PermissionsRevokePost(context.Background()).PermissionsRevokePostRequest(*revokeRequest).Execute()
	}
}

// Verify that Read() only returns managed permissions
func TestAccPermissions_ReadOnlyManaged(t *testing.T) {
	client := getTestAPIClient(t)
	var groupID, folder2ID, folder4ID string

	configWithManagedPerm := `
resource "retool_group" "test_group_read" {
	name = "tf-acc-test-group-read"
}

resource "retool_folder" "test_folder2" {
	name = "tf-acc-test-folder-managed-2"
	folder_type = "app"
}

resource "retool_folder" "test_folder4" {
	name = "tf-acc-test-folder-unmanaged-4"
	folder_type = "app"
}

resource "retool_permissions" "test_permissions_read" {
	subject = {
		type = "group"
		id = retool_group.test_group_read.id
	}
	permissions = [
		{
			object = {
				type = "folder"
				id = retool_folder.test_folder2.id
			}
			access_level = "own"
		},
	]
}
`

	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Create managed permission + unmanaged permission
			{
				Config: configWithManagedPerm,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Capture IDs and create unmanaged permission
					func(s *terraform.State) error {
						groupRes := s.RootModule().Resources["retool_group.test_group_read"]
						groupID = groupRes.Primary.Attributes["id"]

						folder2Res := s.RootModule().Resources["retool_folder.test_folder2"]
						folder2ID = folder2Res.Primary.Attributes["id"]

						folder4Res := s.RootModule().Resources["retool_folder.test_folder4"]
						folder4ID = folder4Res.Primary.Attributes["id"]

						// Create unmanaged permission via direct API call
						createUnmanagedPermission(t, client, "group", groupID, "folder", folder4ID, "use")

						return nil
					},
					// Verify Terraform state only contains 1 managed permission
					resource.TestCheckResourceAttr("retool_permissions.test_permissions_read", "permissions.#", "1"),
					resource.TestCheckResourceAttr("retool_permissions.test_permissions_read", "permissions.0.access_level", "own"),
				),
			},
			// Re-apply same config to trigger Read() and verify state still has only 1 permission
			{
				Config: configWithManagedPerm,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify state still only has 1 managed permission (not 2)
					resource.TestCheckResourceAttr("retool_permissions.test_permissions_read", "permissions.#", "1"),
					// Verify via API that both permissions exist
					func(s *terraform.State) error {
						// Check managed permission exists
						managedExists := checkPermissionExists(t, client, "group", groupID, "folder", folder2ID)
						if !managedExists {
							return fmt.Errorf("Managed permission for folder2 not found via API")
						}

						// Check unmanaged permission exists
						unmanagedExists := checkPermissionExists(t, client, "group", groupID, "folder", folder4ID)
						if !unmanagedExists {
							return fmt.Errorf("Unmanaged permission for folder4 not found via API")
						}

						return nil
					},
				),
			},
		},
	})

	// Cleanup
	if groupID != "" && folder4ID != "" {
		var apiSubject api.PermissionsListObjectsPostRequestSubject
		groupIDInt, _ := strconv.Atoi(groupID)
		floatGroupID := float32(groupIDInt)
		apiSubject.PermissionsListObjectsPostRequestSubjectOneOf = api.NewPermissionsListObjectsPostRequestSubjectOneOf("group", *api.NewNullableFloat32(&floatGroupID))

		apiObject := api.PermissionsGrantPostRequestObject{
			PermissionsGrantPostRequestObjectOneOf: api.NewPermissionsGrantPostRequestObjectOneOf("folder", folder4ID),
		}

		revokeRequest := api.NewPermissionsRevokePostRequest(apiSubject, apiObject)
		client.PermissionsAPI.PermissionsRevokePost(context.Background()).PermissionsRevokePostRequest(*revokeRequest).Execute()
	}
}