package permissions_test

import (
	// "context".
	"fmt"
	// "strconv".
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
	// "github.com/tryretool/terraform-provider-retool/internal/sdk/api".
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

const testUpdatedPermissionsConfig = `resource "retool_group" "test_group" {
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

<<<<<<< HEAD
=======
// TestAccPermissions_ManagedDeletion verifies that the permissions resource properly
// manages only the permissions it creates by testing updates and removals.
func TestAccPermissions_ManagedDeletion(t *testing.T) {
	// Step 1: Create permissions for two folders.
	configWithTwoPerms := `
resource "retool_group" "test_group" {
	name = "tf-acc-test-group-managed-del"
}

resource "retool_folder" "test_folder1" {
	name = "tf-acc-test-folder-managed-1"
	folder_type = "app"
}

resource "retool_folder" "test_folder2" {
	name = "tf-acc-test-folder-managed-2"
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
		{
			object = {
				type = "folder"
				id = retool_folder.test_folder2.id
			}
			access_level = "edit"
		},
	]
}
`

	// Step 2: Remove one permission from config.
	configWithOnePerm := `
resource "retool_group" "test_group" {
	name = "tf-acc-test-group-managed-del"
}

resource "retool_folder" "test_folder1" {
	name = "tf-acc-test-folder-managed-1"
	folder_type = "app"
}

resource "retool_folder" "test_folder2" {
	name = "tf-acc-test-folder-managed-2"
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
			access_level = "edit"
		},
	]
}
`

	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Create permissions for two folders.
			{
				Config: configWithTwoPerms,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("retool_permissions.test_permissions", "permissions.#", "2"),
					// Don't check specific indices as order may vary.
				),
			},
			// Remove one permission from config.
			{
				Config: configWithOnePerm,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify only one permission remains.
					resource.TestCheckResourceAttr("retool_permissions.test_permissions", "permissions.#", "1"),
					resource.TestCheckResourceAttr("retool_permissions.test_permissions", "permissions.0.access_level", "edit"),
				),
			},
		},
	})
}

>>>>>>> main
// TestAccPermissions_ImportAndRead verifies that importing permissions and subsequent
// reads correctly maintain the imported permissions in state.
func TestAccPermissions_ImportAndRead(t *testing.T) {
	config := `
resource "retool_group" "test_group_import" {
	name = "tf-acc-test-group-import"
}

resource "retool_folder" "test_folder_import" {
	name = "tf-acc-test-folder-import"
	folder_type = "app"
}

resource "retool_permissions" "test_permissions_import" {
	subject = {
		type = "group"
		id = retool_group.test_group_import.id
	}
	permissions = [
		{
			object = {
				type = "folder"
				id = retool_folder.test_folder_import.id
			}
			access_level = "own"
		},
	]
}
`

	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Create the permissions.
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("retool_permissions.test_permissions_import", "permissions.#", "1"),
					resource.TestCheckResourceAttr("retool_permissions.test_permissions_import", "permissions.0.access_level", "own"),
				),
			},
			// Import the permissions.
			{
				ResourceName: "retool_permissions.test_permissions_import",
				ImportState:  true,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					permissions, ok := state.RootModule().Resources["retool_permissions.test_permissions_import"]
					if !ok {
						return "", fmt.Errorf("Resource not found")
					}
					return "group|" + permissions.Primary.Attributes["subject.id"], nil
				},
				ImportStateVerify:                    true,
				ImportStateVerifyIdentifierAttribute: "subject.id",
			},
			// Re-apply config to trigger Read() and verify state is maintained.
			{
				Config: config,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("retool_permissions.test_permissions_import", "permissions.#", "1"),
					resource.TestCheckResourceAttr("retool_permissions.test_permissions_import", "permissions.0.access_level", "own"),
				),
			},
		},
	})
}
