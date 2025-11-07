package sourcecontrol

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

func testAppGitHubConfig(t *testing.T, client *api.APIClient) {
	githubConfig := api.GitHubConfig{
		GitHubConfigAnyOf: &api.GitHubConfigAnyOf{
		Type:           "App",
		AppId:          "app_id",
		InstallationId: "installation_id",
		PrivateKey:     "private_key",
		},
	}
	github := api.NewGitHub(githubConfig, "GitHub", "org", "repo", "default_branch")
	apiRequest := api.SourceControlConfigPutRequest{
		Config: api.SourceControlConfigPutRequestConfig{
			GitHub: github,
		},
	}

	response, httpResponse, err := client.SourceControlAPI.SourceControlConfigPut(context.Background()).SourceControlConfigPutRequest(apiRequest).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Data.GitHub.Provider, "GitHub")
	assert.Equal(t, response.Data.GitHub.Config.GitHubConfigAnyOf.Type, "App")
	assert.NotNil(t, httpResponse)
}

func testPersonalGitHubConfig(t *testing.T, client *api.APIClient) {
	githubConfig := api.GitHubConfig{
		GitHubConfigAnyOf1: &api.GitHubConfigAnyOf1{
		Type:                "Personal",
		PersonalAccessToken: "personal_access_token",
		},
	}
	github := api.NewGitHub(githubConfig, "GitHub", "org", "repo", "default_branch")
	apiRequest := api.SourceControlConfigPutRequest{
		Config: api.SourceControlConfigPutRequestConfig{
			GitHub: github,
		},
	}

	response, httpResponse, err := client.SourceControlAPI.SourceControlConfigPut(context.Background()).SourceControlConfigPutRequest(apiRequest).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Data.GitHub.Provider, "GitHub")
	assert.Equal(t, response.Data.GitHub.Config.GitHubConfigAnyOf1.Type, "Personal")
	assert.NotNil(t, httpResponse)
}

func testGitLabConfig(t *testing.T, client *api.APIClient) {
	gitlabConfig := api.NewGitLabConfig(1234, "https://gitlab.com", "project_access_token")
	gitlab := api.NewGitLab(*gitlabConfig, "GitLab", "org", "repo", "default_branch")
	apiRequest := api.SourceControlConfigPutRequest{
		Config: api.SourceControlConfigPutRequestConfig{
			GitLab: gitlab,
		},
	}

	response, httpResponse, err := client.SourceControlAPI.SourceControlConfigPut(context.Background()).SourceControlConfigPutRequest(apiRequest).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Data.GitLab.Provider, "GitLab")
	assert.Equal(t, response.Data.GitLab.Config.ProjectId, float32(1234))
	assert.NotNil(t, httpResponse)
}

func testCodeCommitConfig(t *testing.T, client *api.APIClient) {
	awsConfig := api.NewAWSCodeCommitConfig("https://git-codecommit.us-west-2.amazonaws.com", "us-west-2", "access_key_id", "secret_access_key", "https_username", "https_password")
	awsCodeCommit := api.NewAWSCodeCommit(*awsConfig, "AWS CodeCommit", "org", "repo", "default_branch")
	apiRequest := api.SourceControlConfigPutRequest{
		Config: api.SourceControlConfigPutRequestConfig{
			AWSCodeCommit: awsCodeCommit,
		},
	}

	response, httpResponse, err := client.SourceControlAPI.SourceControlConfigPut(context.Background()).SourceControlConfigPutRequest(apiRequest).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Data.AWSCodeCommit.Provider, "AWS CodeCommit")
	assert.Equal(t, response.Data.AWSCodeCommit.Config.Region, "us-west-2")
	assert.NotNil(t, httpResponse)
}

func testBitbucketConfig(t *testing.T, client *api.APIClient) {
	bitbucketConfigInner := api.NewBitbucketConfigAnyOf("bitbucket", "username", "app_password")
	bitbucketConfig := api.BitbucketConfig{
		BitbucketConfigAnyOf: bitbucketConfigInner,
	}
	bitbucket := api.NewBitbucket(bitbucketConfig, "Bitbucket", "org", "repo", "default_branch")
	apiRequest := api.SourceControlConfigPutRequest{
		Config: api.SourceControlConfigPutRequestConfig{
			Bitbucket: bitbucket,
		},
	}

	response, httpResponse, err := client.SourceControlAPI.SourceControlConfigPut(context.Background()).SourceControlConfigPutRequest(apiRequest).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Data.Bitbucket.Provider, "Bitbucket")
	assert.Equal(t, response.Data.Bitbucket.Config.BitbucketConfigAnyOf.Username, "username")
	assert.NotNil(t, httpResponse)
}

func testAzureReposConfig(t *testing.T, client *api.APIClient) {
	azureConfig := api.NewAzureReposConfig("https://dev.azure.com/organization", "project", "user", "personal_access_token", true)
	azureRepos := api.NewAzureRepos(*azureConfig, "Azure Repos", "org", "repo", "default_branch")
	apiRequest := api.SourceControlConfigPutRequest{
		Config: api.SourceControlConfigPutRequestConfig{
			AzureRepos: azureRepos,
		},
	}

	response, httpResponse, err := client.SourceControlAPI.SourceControlConfigPut(context.Background()).SourceControlConfigPutRequest(apiRequest).Execute()
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, response.Data.AzureRepos.Provider, "Azure Repos")
	assert.Equal(t, response.Data.AzureRepos.Config.Url, "https://dev.azure.com/organization")
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
