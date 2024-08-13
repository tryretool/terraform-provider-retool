package sourcecontrol

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

func testAppGitHubConfig(t *testing.T, client *api.APIClient) {
	apiRequestConfigPart := api.SourceControlConfigGet200ResponseDataOneOfConfigOneOf{
		Type:           "App",
		AppId:          "app_id",
		InstallationId: "installation_id",
		PrivateKey:     "private_key",
	}
	apiRequest := api.SourceControlConfigPutRequest{
		Config: api.SourceControlConfigPutRequestConfig{
			SourceControlConfigGet200ResponseDataOneOf: &api.SourceControlConfigGet200ResponseDataOneOf{
				Provider:      "GitHub",
				Org:           "org",
				Repo:          "repo",
				DefaultBranch: "default_branch",
				Config: api.SourceControlConfigGet200ResponseDataOneOfConfig{
					SourceControlConfigGet200ResponseDataOneOfConfigOneOf: &apiRequestConfigPart,
				},
			},
		},
	}

	response, httpResponse, err := client.SourceControlAPI.SourceControlConfigPut(context.Background()).SourceControlConfigPutRequest(apiRequest).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Data.SourceControlConfigGet200ResponseDataOneOf.Provider, "GitHub")
	assert.Equal(t, response.Data.SourceControlConfigGet200ResponseDataOneOf.Config.SourceControlConfigGet200ResponseDataOneOfConfigOneOf.Type, "App")
	assert.NotNil(t, httpResponse)
}

func testPersonalGitHubConfig(t *testing.T, client *api.APIClient) {
	apiRequestConfigPart := api.SourceControlConfigGet200ResponseDataOneOfConfigOneOf1{
		Type:                "Personal",
		PersonalAccessToken: "personal_access_token",
	}
	apiRequest := api.SourceControlConfigPutRequest{
		Config: api.SourceControlConfigPutRequestConfig{
			SourceControlConfigGet200ResponseDataOneOf: &api.SourceControlConfigGet200ResponseDataOneOf{
				Provider:      "GitHub",
				Org:           "org",
				Repo:          "repo",
				DefaultBranch: "default_branch",
				Config: api.SourceControlConfigGet200ResponseDataOneOfConfig{
					SourceControlConfigGet200ResponseDataOneOfConfigOneOf1: &apiRequestConfigPart,
				},
			},
		},
	}

	response, httpResponse, err := client.SourceControlAPI.SourceControlConfigPut(context.Background()).SourceControlConfigPutRequest(apiRequest).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Data.SourceControlConfigGet200ResponseDataOneOf.Provider, "GitHub")
	assert.Equal(t, response.Data.SourceControlConfigGet200ResponseDataOneOf.Config.SourceControlConfigGet200ResponseDataOneOfConfigOneOf1.Type, "Personal")
	assert.NotNil(t, httpResponse)
}

func testGitLabConfig(t *testing.T, client *api.APIClient) {
	apiRequest := api.SourceControlConfigPutRequest{
		Config: api.SourceControlConfigPutRequestConfig{
			SourceControlConfigGet200ResponseDataOneOf1: &api.SourceControlConfigGet200ResponseDataOneOf1{
				Provider:      "GitLab",
				Org:           "org",
				Repo:          "repo",
				DefaultBranch: "default_branch",
				Config: api.SourceControlConfigGet200ResponseDataOneOf1Config{
					ProjectId:          1234,
					Url:                "https://gitlab.com",
					ProjectAccessToken: "project_access_token",
				},
			},
		},
	}

	response, httpResponse, err := client.SourceControlAPI.SourceControlConfigPut(context.Background()).SourceControlConfigPutRequest(apiRequest).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Data.SourceControlConfigGet200ResponseDataOneOf1.Provider, "GitLab")
	assert.Equal(t, response.Data.SourceControlConfigGet200ResponseDataOneOf1.Config.ProjectId, float32(1234))
	assert.NotNil(t, httpResponse)
}

func testCodeCommitConfig(t *testing.T, client *api.APIClient) {
	apiRequest := api.SourceControlConfigPutRequest{
		Config: api.SourceControlConfigPutRequestConfig{
			SourceControlConfigGet200ResponseDataOneOf2: &api.SourceControlConfigGet200ResponseDataOneOf2{
				Provider:      "AWS CodeCommit",
				Org:           "org",
				Repo:          "repo",
				DefaultBranch: "default_branch",
				Config: api.SourceControlConfigGet200ResponseDataOneOf2Config{
					Url:             "https://git-codecommit.us-west-2.amazonaws.com",
					Region:          "us-west-2",
					AccessKeyId:     "access_key_id",
					SecretAccessKey: "secret_access_key",
					HttpsUsername:   "https_username",
					HttpsPassword:   "https_password",
				},
			},
		},
	}

	response, httpResponse, err := client.SourceControlAPI.SourceControlConfigPut(context.Background()).SourceControlConfigPutRequest(apiRequest).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Data.SourceControlConfigGet200ResponseDataOneOf2.Provider, "AWS CodeCommit")
	assert.Equal(t, response.Data.SourceControlConfigGet200ResponseDataOneOf2.Config.Region, "us-west-2")
	assert.NotNil(t, httpResponse)
}

func testBitbucketConfig(t *testing.T, client *api.APIClient) {
	apiRequest := api.SourceControlConfigPutRequest{
		Config: api.SourceControlConfigPutRequestConfig{
			SourceControlConfigGet200ResponseDataOneOf3: &api.SourceControlConfigGet200ResponseDataOneOf3{
				Provider:      "Bitbucket",
				Org:           "org",
				Repo:          "repo",
				DefaultBranch: "default_branch",
				Config: api.SourceControlConfigGet200ResponseDataOneOf3Config{
					Username:    "username",
					AppPassword: "app_password",
				},
			},
		},
	}

	response, httpResponse, err := client.SourceControlAPI.SourceControlConfigPut(context.Background()).SourceControlConfigPutRequest(apiRequest).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Data.SourceControlConfigGet200ResponseDataOneOf3.Provider, "Bitbucket")
	assert.Equal(t, response.Data.SourceControlConfigGet200ResponseDataOneOf3.Config.Username, "username")
	assert.NotNil(t, httpResponse)
}

func testAzureReposConfig(t *testing.T, client *api.APIClient) {
	apiRequest := api.SourceControlConfigPutRequest{
		Config: api.SourceControlConfigPutRequestConfig{
			SourceControlConfigGet200ResponseDataOneOf4: &api.SourceControlConfigGet200ResponseDataOneOf4{
				Provider:      "Azure Repos",
				Org:           "org",
				Repo:          "repo",
				DefaultBranch: "default_branch",
				Config: api.SourceControlConfigGet200ResponseDataOneOf4Config{
					Url:                 "https://dev.azure.com/organization",
					Project:             "project",
					User:                "user",
					PersonalAccessToken: "personal_access_token",
					UseBasicAuth:        true,
				},
			},
		},
	}

	response, httpResponse, err := client.SourceControlAPI.SourceControlConfigPut(context.Background()).SourceControlConfigPutRequest(apiRequest).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Data.SourceControlConfigGet200ResponseDataOneOf4.Provider, "Azure Repos")
	assert.Equal(t, response.Data.SourceControlConfigGet200ResponseDataOneOf4.Config.Url, "https://dev.azure.com/organization")
	assert.NotNil(t, httpResponse)
}

func TestSourceControlAPI(t *testing.T) {
	t.Skip("Skipping optional integration test")
	host := os.Getenv("RETOOL_HOST")
	scheme := os.Getenv("RETOOL_SCHEME")
	if scheme == "" {
		scheme = "https"
	}
	accessToken := os.Getenv("RETOOL_ACCESS_TOKEN")

	clientConfig := api.NewConfiguration()
	clientConfig.Host = host
	clientConfig.Scheme = scheme
	clientConfig.Servers = api.ServerConfigurations{
		api.ServerConfiguration{
			URL: "/api/v2",
		},
	}
	clientConfig.AddDefaultHeader("Authorization", "Bearer "+accessToken)
	client := api.NewAPIClient(clientConfig)

	testAppGitHubConfig(t, client)
	testPersonalGitHubConfig(t, client)
	testGitLabConfig(t, client)
	testCodeCommitConfig(t, client)
	testBitbucketConfig(t, client)
	testAzureReposConfig(t, client)
}
