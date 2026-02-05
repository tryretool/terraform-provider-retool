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

// TestAccPermissions_Screen tests screen permissions functionality.
// Note: Screen permissions currently work with group subjects but not user subjects.
func TestAccPermissions_Screen(t *testing.T) {
	testConfig := `
resource "retool_group" "test_group_screen" {
	name = "tf-acc-test-group-screen"
}

resource "retool_permissions" "test_screen_permissions" {
	subject = {
		type = "group"
		id = retool_group.test_group_screen.id
	}
	permissions = [
		{
			object = {
				type = "screen"
				id = "4d0f13b1-b919-4d52-ad19-d0b51f6dbcf7"
				app_id = "2cc01bea-0222-11f1-a079-afa417c1dd75"
			}
			access_level = "use"
		},
	]
}
`

	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Create screen permissions
			{
				Config: testConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("retool_permissions.test_screen_permissions", "subject.type", "group"),
					resource.TestCheckResourceAttrPair("retool_permissions.test_screen_permissions", "subject.id", "retool_group.test_group_screen", "id"),
					resource.TestCheckResourceAttr("retool_permissions.test_screen_permissions", "permissions.#", "1"),
					resource.TestCheckResourceAttr("retool_permissions.test_screen_permissions", "permissions.0.object.type", "screen"),
					resource.TestCheckResourceAttr("retool_permissions.test_screen_permissions", "permissions.0.object.id", "4d0f13b1-b919-4d52-ad19-d0b51f6dbcf7"),
					resource.TestCheckResourceAttr("retool_permissions.test_screen_permissions", "permissions.0.object.app_id", "2cc01bea-0222-11f1-a079-afa417c1dd75"),
					resource.TestCheckResourceAttr("retool_permissions.test_screen_permissions", "permissions.0.access_level", "use"),
				),
			},
		},
	})
}
