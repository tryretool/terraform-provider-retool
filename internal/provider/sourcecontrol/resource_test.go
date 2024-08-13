package sourcecontrol_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/tryretool/terraform-provider-retool/internal/acctest"
)

const testGithubAppConfig = `
	resource "retool_source_control" "scm" {
		org = "my-org"
		repo = "my-repo"
		default_branch = "main"
		repo_version = "1.0.0"
		github = {
			app_authentication = {
				app_id = "123456"
				installation_id = "123457"
				private_key = "private-key"
			}
			url = "https://github.com"
			enterprise_api_url = "https://github.mycompany.com/api/v3"
		}
	}	
`

const testGithubPersonalConfig = `
	resource "retool_source_control" "scm" {
		org = "my-org"
		repo = "my-repo"
		default_branch = "main"
		repo_version = "1.0.0"
		github = {
			personal_access_token = "personal-access-token"
			url = "https://github.com"
			enterprise_api_url = "https://github.mycompany.com/api/v3"
		}
	}	
`

const testGitLabConfig = `
	resource "retool_source_control" "scm" {
		org = "my-org"
		repo = "my-repo"
		default_branch = "main"
		repo_version = "1.0.0"
		gitlab = {
			url = "https://gitlab.com"
			project_access_token = "project-access-token"
			project_id = "12345"
		}
	}
`

const testAWSCodeCommitConfig = `
	resource "retool_source_control" "scm" {
		org = "my-org"
		repo = "my-repo"
		default_branch = "main"
		repo_version = "1.0.0"
		aws_codecommit = {
			url = "https://git-codecommit.us-west-2.amazonaws.com/v1/repos/my-repo"
			region = "us-west-2"
			access_key_id = "access-key-id"
			secret_access_key = "secret-access-key"
			https_username = "https-username"
			https_password = "https-password"
		}
	}
`

const testBitbucketConfig = `
	resource "retool_source_control" "scm" {
		org = "my-org"
		repo = "my-repo"
		default_branch = "main"
		repo_version = "1.0.0"
		bitbucket = {
			username = "username"
			app_password = "app_password"
			url = "https://bitbucket.org"
			enterprise_api_url = "https://bitbucket.mycompany.com/api/v3"
		}
	}
`

const testAzureReposConfig = `
	resource "retool_source_control" "scm" {
		org = "my-org"
		repo = "my-repo"
		default_branch = "main"
		repo_version = "1.0.0"
		azure_repos = {
			url = "https://dev.azure.com/organization"
			project = "project"
			user = "user-id"
			personal_access_token = "personal-access-token"
			use_basic_auth = true
		}
	}
`

const testAzureReposUpdatedConfig = `
	resource "retool_source_control" "scm" {
		org = "my-org-updated"
		repo = "my-repo-updated"
		default_branch = "main-updated"
		repo_version = "2.0.0"
		azure_repos = {
			url = "https://dev.azure.com/organization-updated"
			project = "project-updated"
			user = "user-id-updated"
			personal_access_token = "personal-access-token-updated"
			use_basic_auth = false
		}
	}
`

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestAccSourceControl(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Read and Create.
			{
				Config: testGithubAppConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_source_control.scm", "org", "my-org"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "repo", "my-repo"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "default_branch", "main"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "repo_version", "1.0.0"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "github.app_authentication.app_id", "123456"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "github.app_authentication.installation_id", "123457"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "github.app_authentication.private_key", "private-key"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "github.url", "https://github.com"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "github.enterprise_api_url", "https://github.mycompany.com/api/v3"),
				),
				ExpectNonEmptyPlan: true, // Because it'd want to refresh secret strings.
			},
			{
				Config: testGithubPersonalConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_source_control.scm", "org", "my-org"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "repo", "my-repo"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "default_branch", "main"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "repo_version", "1.0.0"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "github.personal_access_token", "personal-access-token"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "github.url", "https://github.com"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "github.enterprise_api_url", "https://github.mycompany.com/api/v3"),
				),
				ExpectNonEmptyPlan: true, // Because it'd want to refresh secret strings.
			},
			{
				Config: testGitLabConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_source_control.scm", "org", "my-org"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "repo", "my-repo"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "default_branch", "main"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "repo_version", "1.0.0"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "gitlab.url", "https://gitlab.com"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "gitlab.project_access_token", "project-access-token"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "gitlab.project_id", "12345"),
				),
				ExpectNonEmptyPlan: true, // Because it'd want to refresh secret strings.
			},
			{
				Config: testAWSCodeCommitConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_source_control.scm", "org", "my-org"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "repo", "my-repo"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "default_branch", "main"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "repo_version", "1.0.0"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "aws_codecommit.url", "https://git-codecommit.us-west-2.amazonaws.com/v1/repos/my-repo"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "aws_codecommit.region", "us-west-2"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "aws_codecommit.access_key_id", "access-key-id"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "aws_codecommit.secret_access_key", "secret-access-key"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "aws_codecommit.https_username", "https-username"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "aws_codecommit.https_password", "https-password"),
				),
				ExpectNonEmptyPlan: true, // Because it'd want to refresh secret strings.
			},
			{
				Config: testBitbucketConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_source_control.scm", "org", "my-org"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "repo", "my-repo"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "default_branch", "main"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "repo_version", "1.0.0"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "bitbucket.username", "username"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "bitbucket.app_password", "app_password"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "bitbucket.url", "https://bitbucket.org"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "bitbucket.enterprise_api_url", "https://bitbucket.mycompany.com/api/v3"),
				),
				ExpectNonEmptyPlan: true, // Because it'd want to refresh secret strings.
			},
			{
				Config: testAzureReposConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_source_control.scm", "org", "my-org"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "repo", "my-repo"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "default_branch", "main"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "repo_version", "1.0.0"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "azure_repos.url", "https://dev.azure.com/organization"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "azure_repos.project", "project"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "azure_repos.user", "user-id"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "azure_repos.personal_access_token", "personal-access-token"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "azure_repos.use_basic_auth", "true"),
				),
				ExpectNonEmptyPlan: true, // Because it'd want to refresh secret strings.
			},
			// Update.
			{
				Config: testAzureReposUpdatedConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_source_control.scm", "org", "my-org-updated"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "repo", "my-repo-updated"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "default_branch", "main-updated"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "repo_version", "2.0.0"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "azure_repos.url", "https://dev.azure.com/organization-updated"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "azure_repos.project", "project-updated"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "azure_repos.user", "user-id-updated"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "azure_repos.personal_access_token", "personal-access-token-updated"),
					resource.TestCheckResourceAttr("retool_source_control.scm", "azure_repos.use_basic_auth", "false"),
				),
				ExpectNonEmptyPlan: true, // Because it'd want to refresh secret strings.
			},
		},
	})
}

func sweepSourceControl(region string) error {
	log.Printf("Sweeping Source Control config in region %s", region)
	client, err := acctest.SweeperClient()
	if err != nil {
		return err
	}

	_, err = client.SourceControlAPI.SourceControlConfigDelete(context.Background()).Execute()
	if err != nil {
		return fmt.Errorf("Error deleting Source Control config: %s", err.Error())
	}
	return nil
}

func init() {
	resource.AddTestSweepers("retool_source_control", &resource.Sweeper{
		Name: "retool_source_control",
		F:    sweepSourceControl,
	})
}
