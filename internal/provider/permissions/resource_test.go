package permissions_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
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
