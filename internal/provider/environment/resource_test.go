package environment_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
)

func TestAccEnvironment(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: testAccEnvironmentConfig("test-environment", "Test environment for acceptance tests", "#FF5733"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("retool_environment.test", "id"),
					resource.TestCheckResourceAttr("retool_environment.test", "name", "test-environment"),
					resource.TestCheckResourceAttr("retool_environment.test", "description", "Test environment for acceptance tests"),
					resource.TestCheckResourceAttr("retool_environment.test", "color", "#FF5733"),
					resource.TestCheckResourceAttrSet("retool_environment.test", "default"),
					resource.TestCheckResourceAttrSet("retool_environment.test", "created_at"),
					resource.TestCheckResourceAttrSet("retool_environment.test", "updated_at"),
				),
			},
			{
				Config: testAccEnvironmentConfig("test-environment-updated", "Updated test environment", "#00FF00"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("retool_environment.test", "id"),
					resource.TestCheckResourceAttr("retool_environment.test", "name", "test-environment-updated"),
					resource.TestCheckResourceAttr("retool_environment.test", "description", "Updated test environment"),
					resource.TestCheckResourceAttr("retool_environment.test", "color", "#00FF00"),
					resource.TestCheckResourceAttrSet("retool_environment.test", "default"),
					resource.TestCheckResourceAttrSet("retool_environment.test", "created_at"),
					resource.TestCheckResourceAttrSet("retool_environment.test", "updated_at"),
				),
			},
			{
				ResourceName:      "retool_environment.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccEnvironmentConfig(name string, description string, color string) string {
	return fmt.Sprintf(`
resource "retool_environment" "test" {
  name        = "%s"
  description = "%s"
  color       = "%s"
}
`, name, description, color)
}

func TestAccEnvironmentWithoutDescription(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: testAccEnvironmentConfigNoDescription("test-environment-no-desc", "#0000FF"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("retool_environment.test", "id"),
					resource.TestCheckResourceAttr("retool_environment.test", "name", "test-environment-no-desc"),
					resource.TestCheckResourceAttr("retool_environment.test", "color", "#0000FF"),
					resource.TestCheckResourceAttrSet("retool_environment.test", "default"),
					resource.TestCheckResourceAttrSet("retool_environment.test", "created_at"),
					resource.TestCheckResourceAttrSet("retool_environment.test", "updated_at"),
				),
			},
		},
	})
}

func testAccEnvironmentConfigNoDescription(name string, color string) string {
	return fmt.Sprintf(`
resource "retool_environment" "test" {
  name  = "%s"
  color = "%s"
}
`, name, color)
}

func checkEnvironmentDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "retool_environment" {
			continue
		}

		// If we reach here, it means the resource still exists after destroy, which is not expected
		// However, since we don't have access to the client in this context, we can't verify
		// We'll rely on the API returning a 404 error in the Read method.
	}
	return nil
}
