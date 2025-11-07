// Package sourcecontrol provides implementation of the Source Control resource.
package sourcecontrol

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// Ensure sourceControlResource implements the tfsdk.Resource interface.
var (
	_ resource.Resource                = &sourceControlResource{}
	_ resource.ResourceWithConfigure   = &sourceControlResource{}
	_ resource.ResourceWithImportState = &sourceControlResource{}
)

type sourceControlResource struct {
	client *api.APIClient
}

type appAuthConfigModel struct {
	AppID          types.String `tfsdk:"app_id"`
	InstallationID types.String `tfsdk:"installation_id"`
	PrivateKey     types.String `tfsdk:"private_key"`
}

func (m appAuthConfigModel) attributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"app_id":          types.StringType,
		"installation_id": types.StringType,
		"private_key":     types.StringType,
	}
}

type githubConfigModel struct {
	PersonalAccessToken types.String `tfsdk:"personal_access_token"`
	AppAuthentication   types.Object `tfsdk:"app_authentication"`
	URL                 types.String `tfsdk:"url"`
	EnterpriseAPIURL    types.String `tfsdk:"enterprise_api_url"`
}

func (m githubConfigModel) attributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"personal_access_token": types.StringType,
		"app_authentication":    types.ObjectType{AttrTypes: appAuthConfigModel{}.attributeTypes()},
		"url":                   types.StringType,
		"enterprise_api_url":    types.StringType,
	}
}

type gitlabConfigModel struct {
	ProjectID          types.String `tfsdk:"project_id"`
	URL                types.String `tfsdk:"url"`
	ProjectAccessToken types.String `tfsdk:"project_access_token"`
}

func (m gitlabConfigModel) attributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"project_id":           types.StringType,
		"url":                  types.StringType,
		"project_access_token": types.StringType,
	}
}

type awsCodeCommitConfigModel struct {
	URL             types.String `tfsdk:"url"`
	Region          types.String `tfsdk:"region"`
	AccessKeyID     types.String `tfsdk:"access_key_id"`
	SecretAccessKey types.String `tfsdk:"secret_access_key"`
	HTTPSUsername   types.String `tfsdk:"https_username"`
	HTTPSPassword   types.String `tfsdk:"https_password"`
}

func (m awsCodeCommitConfigModel) attributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"url":               types.StringType,
		"region":            types.StringType,
		"access_key_id":     types.StringType,
		"secret_access_key": types.StringType,
		"https_username":    types.StringType,
		"https_password":    types.StringType,
	}
}

type bitbucketConfigModel struct {
	Username         types.String `tfsdk:"username"`
	URL              types.String `tfsdk:"url"`
	EnterpriseAPIURL types.String `tfsdk:"enterprise_api_url"`
	AppPassword      types.String `tfsdk:"app_password"`
}

func (m bitbucketConfigModel) attributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"username":           types.StringType,
		"url":                types.StringType,
		"enterprise_api_url": types.StringType,
		"app_password":       types.StringType,
	}
}

type azureReposConfigModel struct {
	URL                 types.String `tfsdk:"url"`
	Project             types.String `tfsdk:"project"`
	User                types.String `tfsdk:"user"`
	PersonalAccessToken types.String `tfsdk:"personal_access_token"`
	UseBasicAuth        types.Bool   `tfsdk:"use_basic_auth"`
}

func (m azureReposConfigModel) attributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"url":                   types.StringType,
		"project":               types.StringType,
		"user":                  types.StringType,
		"personal_access_token": types.StringType,
		"use_basic_auth":        types.BoolType,
	}
}

type sourceControlModel struct {
	Org           types.String `tfsdk:"org"`
	Repo          types.String `tfsdk:"repo"`
	DefaultBranch types.String `tfsdk:"default_branch"`
	RepoVersion   types.String `tfsdk:"repo_version"`
	GitHub        types.Object `tfsdk:"github"`
	GitLab        types.Object `tfsdk:"gitlab"`
	AWSCodeCommit types.Object `tfsdk:"aws_codecommit"`
	Bitbucket     types.Object `tfsdk:"bitbucket"`
	AzureRepos    types.Object `tfsdk:"azure_repos"`
}

// Create new Source Control resource.
func NewResource() resource.Resource {
	return &sourceControlResource{}
}

// Configure adds the provider configured client to the resource.
func (r *sourceControlResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*utils.ProviderData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			"Expected *utils.ProviderData, got: %T. Please report this issue to the provider developers.",
		)
	}
	r.client = providerData.Client
}

// Metadata associated with the Source Control resource.
func (r *sourceControlResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_control"
}

// Schema returns the schema for the Source Control resource.
func (r *sourceControlResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Source Control resource allows you to connect Retool to a source control provider.",
		Attributes: map[string]schema.Attribute{
			"org": schema.StringAttribute{
				Required:    true,
				Description: "The user or organization to which the repository belongs to.",
			},
			"repo": schema.StringAttribute{
				Required:    true,
				Description: "The name of the repository you created to use with Retool.",
			},
			"default_branch": schema.StringAttribute{
				Required:    true,
				Description: "The default branch, e.g., main.",
			},
			"repo_version": schema.StringAttribute{
				Optional:    true,
				Description: "Repositories using Toolscript are 2.0.0. Repositories using legacy YAML are 1.0.0.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"github": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "GitHub Source Control settings for App authentication",
				Attributes: map[string]schema.Attribute{
					"personal_access_token": schema.StringAttribute{
						Optional:    true,
						Description: "The GitHub project access token to authenticate to the GitHub API. Required if app_authentication is not set.",
						Sensitive:   true,
					},
					"app_authentication": schema.SingleNestedAttribute{
						Optional:    true,
						Description: "GitHub App authentication settings. Required if personal_access_token is not set.",
						Attributes: map[string]schema.Attribute{
							"app_id": schema.StringAttribute{
								Required:    true,
								Description: "The GitHub App ID.",
							},
							"installation_id": schema.StringAttribute{
								Required:    true,
								Description: "The GitHub installation ID. This can be found at the end of the installation URL.",
							},
							"private_key": schema.StringAttribute{
								Required:    true,
								Description: "The base64-encoded private key.",
								Sensitive:   true,
							},
						},
					},
					"url": schema.StringAttribute{
						Optional:    true,
						Description: "The domain used to access your self-hosted GitHub instance.",
					},
					"enterprise_api_url": schema.StringAttribute{
						Optional:    true,
						Description: "The REST API route for your self-hosted GitHub instance. Defaults to https://[hostname]/api/v3.",
					},
				},
			},
			"gitlab": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "GitLab Source Control settings",
				Attributes: map[string]schema.Attribute{
					"project_id": schema.StringAttribute{
						Required:    true,
						Description: "The numerical project ID for your GitLab project. Find this ID listed below the project's name on the project's homepage.",
					},
					"url": schema.StringAttribute{
						Required:    true,
						Description: "Your base GitLab URL. On GitLab Cloud, this is always https://gitlab.com. On GitLab self-managed, this is the URL where your instance is hosted.",
					},
					"project_access_token": schema.StringAttribute{
						Required:    true,
						Description: "The GitLab project access token to authenticate to the GitLab API.",
						Sensitive:   true,
					},
				},
			},
			"aws_codecommit": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "AWS CodeCommit Source Control settings",
				Attributes: map[string]schema.Attribute{
					"url": schema.StringAttribute{
						Required:    true,
						Description: "The domain used to access your self-hosted AWS CodeCommit instance.",
					},
					"region": schema.StringAttribute{
						Required:    true,
						Description: "The region of the CodeCommit repository.",
					},
					"access_key_id": schema.StringAttribute{
						Required:    true,
						Description: "The Access key ID from your AWSCodeCommitFullAccess policy.",
						Sensitive:   true,
					},
					"secret_access_key": schema.StringAttribute{
						Required:    true,
						Description: "The Secret Access Key from your AWSCodeCommitFullAccess policy",
						Sensitive:   true,
					},
					"https_username": schema.StringAttribute{
						Required:    true,
						Description: "The HTTPS username from your security credentials.",
					},
					"https_password": schema.StringAttribute{
						Required:    true,
						Description: "The HTTPS password from your security credentials.",
						Sensitive:   true,
					},
				},
			},
			"bitbucket": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Bitbucket Source Control settings",
				Attributes: map[string]schema.Attribute{
					"username": schema.StringAttribute{
						Required:    true,
						Description: "Your Bitbucket username.",
					},
					"url": schema.StringAttribute{
						Required:    true,
						Description: "The domain used to access your self-hosted Bitbucket instance. Defaults to https://bitbucket.org/.",
					},
					"enterprise_api_url": schema.StringAttribute{
						Required:    true,
						Description: "The REST API route for your self-hosted Bitbucket instance. Defaults to https://api.bitbucket.org/2.0.",
					},
					"app_password": schema.StringAttribute{
						Required:    true,
						Description: "Your Bitbucket app password.",
						Sensitive:   true,
					},
				},
			},
			"azure_repos": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Azure Repos Source Control settings",
				Attributes: map[string]schema.Attribute{
					"url": schema.StringAttribute{
						Required:    true,
						Description: "Your base Azure URL. For Azure Cloud, this is always http://dev.azure.com. For Azure self-managed, this is the URL where your instance is hosted.",
					},
					"project": schema.StringAttribute{
						Required:    true,
						Description: "Your new or existing Azure DevOps project.",
					},
					"user": schema.StringAttribute{
						Required:    true,
						Description: "The Azure Repos username.",
					},
					"personal_access_token": schema.StringAttribute{
						Required:    true,
						Description: "The Azure project access tokens to authenticate to the Azure API.",
						Sensitive:   true,
					},
					"use_basic_auth": schema.BoolAttribute{
						Required:    true,
						Description: "Set this to true if you are using self-hosted Azure Repos.",
					},
				},
			},
		},
	}
}

// Custom validation implementation to prevent common errors in the Source Control config.
func (r *sourceControlResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var model sourceControlModel
	diags := req.Config.Get(ctx, &model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var sourceControlProviders = []bool{
		!utils.IsEmptyObject(model.GitHub),
		!utils.IsEmptyObject(model.GitLab),
		!utils.IsEmptyObject(model.AWSCodeCommit),
		!utils.IsEmptyObject(model.Bitbucket),
		!utils.IsEmptyObject(model.AzureRepos),
	}
	var countSetProviders int
	for _, providerSet := range sourceControlProviders {
		if providerSet {
			countSetProviders++
		}
	}
	if countSetProviders == 0 {
		resp.Diagnostics.AddError("No Source Control provider set", "Exactly one Source Control provider config must be provided")
		return
	}
	if countSetProviders > 1 {
		resp.Diagnostics.AddError("Multiple Source Control providers set", "Only one Source Control provider config can be provided")
		return
	}

	if !utils.IsEmptyObject(model.GitHub) {
		var githubConfig githubConfigModel
		diags := model.GitHub.As(ctx, &githubConfig, basetypes.ObjectAsOptions{})
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		if githubConfig.PersonalAccessToken.ValueStringPointer() == nil && utils.IsEmptyObject(githubConfig.AppAuthentication) {
			resp.Diagnostics.AddError("GitHub authentication config is missing", "Either a personal access token or app authentication must be configured")
			return
		}
	}
}

func updateSourceControlConfig(ctx context.Context, client *api.APIClient, model sourceControlModel, globalDiags *diag.Diagnostics) {
	var config api.SourceControlConfigPutRequestConfig
	switch {
	case !utils.IsEmptyObject(model.GitHub):
		var githubConfig githubConfigModel
		diags := model.GitHub.As(ctx, &githubConfig, basetypes.ObjectAsOptions{})
		globalDiags.Append(diags...)
		if globalDiags.HasError() {
			return
		}
		innerConfig := api.GitHubConfig{}

		if !utils.IsEmptyObject(githubConfig.AppAuthentication) {
			var appAuthConfig appAuthConfigModel
			diags := githubConfig.AppAuthentication.As(ctx, &appAuthConfig, basetypes.ObjectAsOptions{})
			globalDiags.Append(diags...)
			if globalDiags.HasError() {
				return
			}
			innerConfig.GitHubConfigAnyOf = &api.GitHubConfigAnyOf{
				Type:             "App",
				AppId:            appAuthConfig.AppID.ValueString(),
				InstallationId:   appAuthConfig.InstallationID.ValueString(),
				PrivateKey:       appAuthConfig.PrivateKey.ValueString(),
				Url:              githubConfig.URL.ValueStringPointer(),
				EnterpriseApiUrl: githubConfig.EnterpriseAPIURL.ValueStringPointer(),
			}
		} else {
			// Assuming here that the personal access token is set.
			innerConfig.GitHubConfigAnyOf1 = &api.GitHubConfigAnyOf1{
				Type:                "Personal",
				PersonalAccessToken: githubConfig.PersonalAccessToken.ValueString(),
				Url:                 githubConfig.URL.ValueStringPointer(),
				EnterpriseApiUrl:    githubConfig.EnterpriseAPIURL.ValueStringPointer(),
			}
		}

		github := api.NewGitHub(innerConfig, "GitHub", model.Org.ValueString(), model.Repo.ValueString(), model.DefaultBranch.ValueString())
		github.RepoVersion = model.RepoVersion.ValueStringPointer()
		config = api.SourceControlConfigPutRequestConfig{
			GitHub: github,
		}
	case !utils.IsEmptyObject(model.GitLab):
		var gitlabConfig gitlabConfigModel
		diags := model.GitLab.As(ctx, &gitlabConfig, basetypes.ObjectAsOptions{})
		globalDiags.Append(diags...)
		if globalDiags.HasError() {
			return
		}

		projectID, err := strconv.ParseFloat(gitlabConfig.ProjectID.ValueString(), 32)
		if err != nil {
			globalDiags.AddError("Invalid project ID", "The project ID must be a number")
			return
		}
		gitlabConfigObj := api.NewGitLabConfig(float32(projectID), gitlabConfig.URL.ValueString(), gitlabConfig.ProjectAccessToken.ValueString())
		gitlab := api.NewGitLab(*gitlabConfigObj, "GitLab", model.Org.ValueString(), model.Repo.ValueString(), model.DefaultBranch.ValueString())
		gitlab.RepoVersion = model.RepoVersion.ValueStringPointer()
		config = api.SourceControlConfigPutRequestConfig{
			GitLab: gitlab,
		}
	case !utils.IsEmptyObject(model.AWSCodeCommit):
		var awsCodeCommitConfig awsCodeCommitConfigModel
		diags := model.AWSCodeCommit.As(ctx, &awsCodeCommitConfig, basetypes.ObjectAsOptions{})
		globalDiags.Append(diags...)
		if globalDiags.HasError() {
			return
		}

		awsConfigObj := api.NewAWSCodeCommitConfig(awsCodeCommitConfig.URL.ValueString(), awsCodeCommitConfig.Region.ValueString(), awsCodeCommitConfig.AccessKeyID.ValueString(), awsCodeCommitConfig.SecretAccessKey.ValueString(), awsCodeCommitConfig.HTTPSUsername.ValueString(), awsCodeCommitConfig.HTTPSPassword.ValueString())
		awsCodeCommit := api.NewAWSCodeCommit(*awsConfigObj, "AWS CodeCommit", model.Org.ValueString(), model.Repo.ValueString(), model.DefaultBranch.ValueString())
		awsCodeCommit.RepoVersion = model.RepoVersion.ValueStringPointer()
		config = api.SourceControlConfigPutRequestConfig{
			AWSCodeCommit: awsCodeCommit,
		}
	case !utils.IsEmptyObject(model.Bitbucket):
		var bitbucketConfig bitbucketConfigModel
		diags := model.Bitbucket.As(ctx, &bitbucketConfig, basetypes.ObjectAsOptions{})
		globalDiags.Append(diags...)
		if globalDiags.HasError() {
			return
		}
		bitbucketConfigInner := api.NewBitbucketConfigAnyOf("AppPassword", bitbucketConfig.Username.ValueString(), bitbucketConfig.AppPassword.ValueString())
		bitbucketConfigInner.Url = bitbucketConfig.URL.ValueStringPointer()
		bitbucketConfigInner.EnterpriseApiUrl = bitbucketConfig.EnterpriseAPIURL.ValueStringPointer()
		bitbucketConfigWrapped := api.BitbucketConfig{
			BitbucketConfigAnyOf: bitbucketConfigInner,
		}
		bitbucket := api.NewBitbucket(bitbucketConfigWrapped, "Bitbucket", model.Org.ValueString(), model.Repo.ValueString(), model.DefaultBranch.ValueString())
		bitbucket.RepoVersion = model.RepoVersion.ValueStringPointer()
		config = api.SourceControlConfigPutRequestConfig{
			Bitbucket: bitbucket,
		}
	case !utils.IsEmptyObject(model.AzureRepos):
		var azureReposConfig azureReposConfigModel
		diags := model.AzureRepos.As(ctx, &azureReposConfig, basetypes.ObjectAsOptions{})
		globalDiags.Append(diags...)
		if globalDiags.HasError() {
			return
		}
		azureConfigObj := api.NewAzureReposConfig(azureReposConfig.URL.ValueString(), azureReposConfig.Project.ValueString(), azureReposConfig.User.ValueString(), azureReposConfig.PersonalAccessToken.ValueString(), azureReposConfig.UseBasicAuth.ValueBool())
		azureRepos := api.NewAzureRepos(*azureConfigObj, "Azure Repos", model.Org.ValueString(), model.Repo.ValueString(), model.DefaultBranch.ValueString())
		azureRepos.RepoVersion = model.RepoVersion.ValueStringPointer()
		config = api.SourceControlConfigPutRequestConfig{
			AzureRepos: azureRepos,
		}
	}

	apiRequest := api.SourceControlConfigPutRequest{
		Config: config,
	}

	_, httpResponse, err := client.SourceControlAPI.SourceControlConfigPut(context.Background()).SourceControlConfigPutRequest(apiRequest).Execute()
	if err != nil {
		globalDiags.AddError("Failed to update Source Control config", err.Error())
		tflog.Error(ctx, "Error updating Source Control config", utils.AddHTTPStatusCode(map[string]interface{}{"error": err.Error()}, httpResponse))
		return
	}
}

// Updates Source Control config and sets state.
func (r *sourceControlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan sourceControlModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateSourceControlConfig(ctx, r.client, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state to fully populated data.
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error creating Source Control conifg", map[string]interface{}{"error": "Could not set state"})
		return
	}

	tflog.Info(ctx, "Source Control config created")
}

// Read Source Control config.
func (r *sourceControlResource) Read(ctx context.Context, _ resource.ReadRequest, resp *resource.ReadResponse) {
	response, httpResponse, err := r.client.SourceControlAPI.SourceControlConfigGet(context.Background()).Execute()
	if err != nil {
		if httpResponse != nil && httpResponse.StatusCode == 404 {
			tflog.Info(ctx, "Source Control is not configured")
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError(
			"Error reading Source Control config",
			fmt.Sprintf("Could not read Source Control config: %s", err.Error()),
		)
		tflog.Error(ctx, "Error reading Source Control config", utils.AddHTTPStatusCode(map[string]any{"error": err.Error()}, httpResponse))
		return
	}

	var state sourceControlModel
	// First init all object fields to empty objects.
	state.GitHub = types.ObjectNull(githubConfigModel{}.attributeTypes())
	state.GitLab = types.ObjectNull(gitlabConfigModel{}.attributeTypes())
	state.AWSCodeCommit = types.ObjectNull(awsCodeCommitConfigModel{}.attributeTypes())
	state.Bitbucket = types.ObjectNull(bitbucketConfigModel{}.attributeTypes())
	state.AzureRepos = types.ObjectNull(azureReposConfigModel{}.attributeTypes())

	switch {
	case response.Data.GitHub != nil:
		state.Org = types.StringValue(response.Data.GitHub.Org)
		state.Repo = types.StringValue(response.Data.GitHub.Repo)
		state.DefaultBranch = types.StringValue(response.Data.GitHub.DefaultBranch)
		state.RepoVersion = types.StringPointerValue(response.Data.GitHub.RepoVersion)

		githubConfigModel := githubConfigModel{}
		if response.Data.GitHub.Config.GitHubConfigAnyOf != nil {
			appAuthConfig := appAuthConfigModel{
				AppID:          types.StringValue(response.Data.GitHub.Config.GitHubConfigAnyOf.AppId),
				InstallationID: types.StringValue(response.Data.GitHub.Config.GitHubConfigAnyOf.InstallationId),
				PrivateKey:     types.StringNull(), // API sends placeholder value that we don't need.
			}
			appAuthConfigModelObj, diags := types.ObjectValueFrom(ctx, appAuthConfig.attributeTypes(), appAuthConfig)
			resp.Diagnostics.Append(diags...)
			if resp.Diagnostics.HasError() {
				return
			}
			githubConfigModel.PersonalAccessToken = types.StringNull()
			githubConfigModel.AppAuthentication = appAuthConfigModelObj
			githubConfigModel.URL = types.StringPointerValue(response.Data.GitHub.Config.GitHubConfigAnyOf.Url)
			githubConfigModel.EnterpriseAPIURL = types.StringPointerValue(response.Data.GitHub.Config.GitHubConfigAnyOf.EnterpriseApiUrl)
		} else {
			githubConfigModel.PersonalAccessToken = types.StringNull() // API sends placeholder value that we don't need.
			githubConfigModel.AppAuthentication = types.ObjectNull(appAuthConfigModel{}.attributeTypes())
			githubConfigModel.URL = types.StringPointerValue(response.Data.GitHub.Config.GitHubConfigAnyOf1.Url)
			githubConfigModel.EnterpriseAPIURL = types.StringPointerValue(response.Data.GitHub.Config.GitHubConfigAnyOf1.EnterpriseApiUrl)
		}
		githubConfigModelObj, diags := types.ObjectValueFrom(ctx, githubConfigModel.attributeTypes(), githubConfigModel)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		state.GitHub = githubConfigModelObj

	case response.Data.GitLab != nil:
		state.Org = types.StringValue(response.Data.GitLab.Org)
		state.Repo = types.StringValue(response.Data.GitLab.Repo)
		state.DefaultBranch = types.StringValue(response.Data.GitLab.DefaultBranch)
		state.RepoVersion = types.StringPointerValue(response.Data.GitLab.RepoVersion)

		gitlabConfigModel := gitlabConfigModel{
			ProjectID:          types.StringValue(utils.Float32PtrToIntString(&response.Data.GitLab.Config.ProjectId)),
			URL:                types.StringValue(response.Data.GitLab.Config.Url),
			ProjectAccessToken: types.StringNull(), // API sends placeholder value that we don't need.
		}
		gitlabConfigModelObj, diags := types.ObjectValueFrom(ctx, gitlabConfigModel.attributeTypes(), gitlabConfigModel)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		state.GitLab = gitlabConfigModelObj

	case response.Data.AWSCodeCommit != nil:
		state.Org = types.StringValue(response.Data.AWSCodeCommit.Org)
		state.Repo = types.StringValue(response.Data.AWSCodeCommit.Repo)
		state.DefaultBranch = types.StringValue(response.Data.AWSCodeCommit.DefaultBranch)
		state.RepoVersion = types.StringPointerValue(response.Data.AWSCodeCommit.RepoVersion)

		awsCodeCommitConfigModel := awsCodeCommitConfigModel{
			URL:             types.StringValue(response.Data.AWSCodeCommit.Config.Url),
			Region:          types.StringValue(response.Data.AWSCodeCommit.Config.Region),
			AccessKeyID:     types.StringNull(), // API sends placeholder value that we don't need.
			SecretAccessKey: types.StringNull(), // API sends placeholder value that we don't need,.
			HTTPSUsername:   types.StringValue(response.Data.AWSCodeCommit.Config.HttpsUsername),
			HTTPSPassword:   types.StringNull(), // API sends placeholder value that we don't need.
		}
		awsCodeCommitConfigModelObj, diags := types.ObjectValueFrom(ctx, awsCodeCommitConfigModel.attributeTypes(), awsCodeCommitConfigModel)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		state.AWSCodeCommit = awsCodeCommitConfigModelObj

	case response.Data.Bitbucket != nil:
		state.Org = types.StringValue(response.Data.Bitbucket.Org)
		state.Repo = types.StringValue(response.Data.Bitbucket.Repo)
		state.DefaultBranch = types.StringValue(response.Data.Bitbucket.DefaultBranch)
		state.RepoVersion = types.StringPointerValue(response.Data.Bitbucket.RepoVersion)

		bitbucketConfigModel := bitbucketConfigModel{
			Username:         types.StringValue(response.Data.Bitbucket.Config.BitbucketConfigAnyOf.Username),
			URL:              types.StringPointerValue(response.Data.Bitbucket.Config.BitbucketConfigAnyOf.Url),
			EnterpriseAPIURL: types.StringPointerValue(response.Data.Bitbucket.Config.BitbucketConfigAnyOf.EnterpriseApiUrl),
			AppPassword:      types.StringNull(), // API sends placeholder value that we don't need.
		}
		bitbucketConfigModelObj, diags := types.ObjectValueFrom(ctx, bitbucketConfigModel.attributeTypes(), bitbucketConfigModel)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		state.Bitbucket = bitbucketConfigModelObj

	case response.Data.AzureRepos != nil:
		state.Org = types.StringValue(response.Data.AzureRepos.Org)
		state.Repo = types.StringValue(response.Data.AzureRepos.Repo)
		state.DefaultBranch = types.StringValue(response.Data.AzureRepos.DefaultBranch)
		state.RepoVersion = types.StringPointerValue(response.Data.AzureRepos.RepoVersion)

		azureReposConfigModel := azureReposConfigModel{
			URL:                 types.StringValue(response.Data.AzureRepos.Config.Url),
			Project:             types.StringValue(response.Data.AzureRepos.Config.Project),
			User:                types.StringValue(response.Data.AzureRepos.Config.User),
			PersonalAccessToken: types.StringNull(), // API sends placeholder value that we don't need.
			UseBasicAuth:        types.BoolValue(response.Data.AzureRepos.Config.UseBasicAuth),
		}
		azureReposConfigModelObj, diags := types.ObjectValueFrom(ctx, azureReposConfigModel.attributeTypes(), azureReposConfigModel)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		state.AzureRepos = azureReposConfigModelObj
	}
	diags := resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error reading Source Control conifg", map[string]interface{}{"error": "Could not set state"})
		return
	}
}

// Update Source Control config.
func (r *sourceControlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan sourceControlModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateSourceControlConfig(ctx, r.client, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state to fully populated data.
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		tflog.Error(ctx, "Error updating Source Control conifg", map[string]interface{}{"error": "Could not set state"})
		return
	}

	tflog.Info(ctx, "Source Control config updated")
}

// Delete Source Control config.
func (r *sourceControlResource) Delete(ctx context.Context, _ resource.DeleteRequest, resp *resource.DeleteResponse) {
	httpResponse, err := r.client.SourceControlAPI.SourceControlConfigDelete(context.Background()).Execute()
	if err != nil && !(httpResponse != nil && httpResponse.StatusCode == 404) { // It's ok to not find the source control config when deleting.
		resp.Diagnostics.AddError(
			"Error Deleting Source Control Config",
			"Could not Source Control config: "+err.Error(),
		)
		tflog.Error(ctx, "Error Deleting Source Control Config", utils.AddHTTPStatusCode(map[string]any{"error": err.Error()}, httpResponse))
		return
	}
}

// Import Source Control config.
func (r *sourceControlResource) ImportState(ctx context.Context, _ resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	emptyModel := sourceControlModel{
		GitHub:        types.ObjectNull(githubConfigModel{}.attributeTypes()),
		GitLab:        types.ObjectNull(gitlabConfigModel{}.attributeTypes()),
		AWSCodeCommit: types.ObjectNull(awsCodeCommitConfigModel{}.attributeTypes()),
		Bitbucket:     types.ObjectNull(bitbucketConfigModel{}.attributeTypes()),
		AzureRepos:    types.ObjectNull(azureReposConfigModel{}.attributeTypes()),
	}
	diags := resp.State.Set(ctx, emptyModel)
	resp.Diagnostics.Append(diags...)
}
