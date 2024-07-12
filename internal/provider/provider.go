package provider

import (
	"context"
	"net/http"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/tryretool/terraform-provider-retool/internal/provider/folder"
	"github.com/tryretool/terraform-provider-retool/internal/provider/group"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &retoolProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &retoolProvider{
			version: version,
		}
	}
}

func NewWithHttpClient(version string, httpClient *http.Client) func() provider.Provider {
	return func() provider.Provider {
		return &retoolProvider{
			version:    version,
			httpClient: httpClient,
		}
	}
}

// retoolProvider is the provider implementation.
type retoolProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version    string
	httpClient *http.Client
}

type retoolProviderModel struct {
	Host        types.String `tfsdk:"host"`
	Scheme      types.String `tfsdk:"scheme"`
	AccessToken types.String `tfsdk:"access_token"`
}

// Metadata returns the provider type name.
func (p *retoolProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "retool"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *retoolProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Description: "The host of the Retool instance, organization or Space, e.g. 'example.retool.com'",
				Optional:    true,
			},
			"scheme": schema.StringAttribute{
				Description: "The scheme of the Retool instance, e.g. 'https'",
				Optional:    true,
			},
			"access_token": schema.StringAttribute{
				Description: "The access token for the Retool API",
				Optional:    true,
				Sensitive:   true,
			},
		},
	}
}

// Configure prepares a Retool API client for data sources and resources.
func (p *retoolProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Retool API client")
	// Retrieve provider data from configuration
	var config retoolProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Retool API Host",
			"The provider cannot create the Retool API client as there is an unknown configuration value for the Retool API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the RETOOL_HOST environment variable.",
		)
	}

	if config.AccessToken.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("access_token"),
			"Unknown Retool API Access Token",
			"The provider cannot create the Retool API client as there is an unknown configuration value for the Retool API Access Token. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the RETOOL_ACCESS_TOKEN environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("RETOOL_HOST")
	scheme := os.Getenv("RETOOL_SCHEME")
	if scheme == "" {
		scheme = "https"
	}
	accessToken := os.Getenv("RETOOL_ACCESS_TOKEN")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.Scheme.IsNull() && config.Scheme.ValueString() != "" {
		scheme = config.Scheme.ValueString()
	}

	if !config.AccessToken.IsNull() {
		accessToken = config.AccessToken.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Retool API Host",
			"The provider cannot create the Retool API client as there is a missing or empty value for the Retool API host. "+
				"Set the host value in the configuration or use the RETOOL_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if accessToken == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("access_token"),
			"Missing Retool API Access Token",
			"The provider cannot create the Retool API client as there is a missing or empty value for the Retool API Access Token. "+
				"Access token can be created in the Retool instance under the 'Settings > Retool API' section. "+
				"Set the access token value in the configuration or use the RETOOL_ACCESS_TOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	clientConfig := api.NewConfiguration()
	clientConfig.Host = host
	clientConfig.Scheme = scheme
	clientConfig.Servers = api.ServerConfigurations{
		api.ServerConfiguration{
			URL: "/api/v2",
		},
	}
	// We need this to be able to record and replay HTTP interactions in the acceptance tests
	if p.httpClient != nil {
		clientConfig.HTTPClient = p.httpClient
	}

	clientConfig.AddDefaultHeader("Authorization", "Bearer "+accessToken)
	client := api.NewAPIClient(clientConfig)

	// Make the Retool client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
	tflog.Info(ctx, "Retool API client configured")
}

// DataSources defines the data sources implemented in the provider.
func (p *retoolProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		folder.NewDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *retoolProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		folder.NewResource,
		group.NewResource,
	}
}
