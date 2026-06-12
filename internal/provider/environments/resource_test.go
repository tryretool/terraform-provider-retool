package environments_test

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestAccEnvironment(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			{
				Config: testAccEnvironmentConfig("tf-acc-test-environment", "Test environment for acceptance tests", "#FF5733"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("retool_environment.test", "id"),
					resource.TestCheckResourceAttr("retool_environment.test", "name", "tf-acc-test-environment"),
					resource.TestCheckResourceAttr("retool_environment.test", "description", "Test environment for acceptance tests"),
					resource.TestCheckResourceAttr("retool_environment.test", "color", "#FF5733"),
					resource.TestCheckResourceAttrSet("retool_environment.test", "default"),
					resource.TestCheckResourceAttrSet("retool_environment.test", "created_at"),
					resource.TestCheckResourceAttrSet("retool_environment.test", "updated_at"),
				),
			},
			{
				Config: testAccEnvironmentConfig("tf-acc-test-environment-updated", "Updated test environment", "#00FF00"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("retool_environment.test", "id"),
					resource.TestCheckResourceAttr("retool_environment.test", "name", "tf-acc-test-environment-updated"),
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
				Config: testAccEnvironmentConfigNoDescription("tf-acc-test-environment-no-desc", "#0000FF"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("retool_environment.test", "id"),
					resource.TestCheckResourceAttr("retool_environment.test", "name", "tf-acc-test-environment-no-desc"),
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

func sweepEnvironments(region string) error {
	log.Printf("Sweeping environments in region %s", region)
	client, err := acctest.SweeperClient()
	if err != nil {
		return fmt.Errorf("could not create Retool API client: %w", err)
	}

	environments, _, err := client.EnvironmentsAPI.EnvironmentsGet(context.Background()).Execute()
	if err != nil {
		return fmt.Errorf("error reading environments: %s", err.Error())
	}

	for _, environment := range environments.Data {
		// Never delete the default environment.
		if environment.Default {
			continue
		}
		if strings.HasPrefix(environment.Name, "tf-acc") {
			log.Printf("Deleting environment %s (%s)\n", environment.Name, environment.Id)
			_, err := client.EnvironmentsAPI.EnvironmentsEnvironmentIdDelete(context.Background(), environment.Id).Execute()
			if err != nil {
				return fmt.Errorf("error deleting environment %s: %s", environment.Id, err.Error())
			}
			log.Printf("Deleted environment %s\n", environment.Id)
		}
	}
	return nil
}

func init() {
	resource.AddTestSweepers("retool_environment", &resource.Sweeper{
		Name: "retool_environment",
		F:    sweepEnvironments,
	})
}
