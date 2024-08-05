package source_control

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// Ensure SSOResource implements the tfsdk.Resource interface
var (
	_ resource.Resource              = &sourceControlResource{}
	_ resource.ResourceWithConfigure = &sourceControlResource{}
)

// ssoResource schema structure
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
	AccessKeyID     types.String `tfsdk:"access_key_id`
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

// Metadata associated with the SSO resource
func (r *sourceControlResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_source_control"
}

// Schema returns the schema for the SSO resource
func (r *sourceControlResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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

// Custom validation implementation to prevent common errors in the Source Control config
func (r *sourceControlResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var model sourceControlModel
	diags := req.Config.Get(ctx, &model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: check that exactly one of the source control providers is set
	// TODO: check that GitHub has either a personal access token or app authentication
}

// Updates Source Control config and sets state
func (r *sourceControlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// TODO: implement
}

// Read SSO config
func (r *sourceControlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}


// Update Source Control config
func (r *sourceControlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// TODO: implement
}

// Delete Source Control config
func (r *sourceControlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// TODO: implement
}
