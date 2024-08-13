// Package provider provides the implementation of the Retool Terraform provider.
package provider

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"golang.org/x/mod/semver"

	"github.com/tryretool/terraform-provider-retool/internal/provider/folder"
	"github.com/tryretool/terraform-provider-retool/internal/provider/group"
	"github.com/tryretool/terraform-provider-retool/internal/provider/permissions"
	"github.com/tryretool/terraform-provider-retool/internal/provider/sourcecontrol"
	"github.com/tryretool/terraform-provider-retool/internal/provider/sourcecontrolsettings"
	"github.com/tryretool/terraform-provider-retool/internal/provider/space"
	"github.com/tryretool/terraform-provider-retool/internal/provider/sso"
	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &retoolProvider{}
)

const MINIMUM_RETOOL_VERSION = "v3.75.0"

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
	// Version is set to the provider version on release, "dev" when the
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

type HealthCheckResponse struct {
	Version string `json:"version"`
}

func checkMinimalVersion(ctx context.Context, host string, scheme string) bool {
	// Create HTTP client, make GET /api/checkHealth request, parse the version field out of the JSON response.
	httpResponse, err := http.Get(scheme + "://" + host + "/api/checkHealth")
	if err != nil {
		tflog.Error(ctx, "Failed to check Retool version", map[string]any{"error": err})
		return false
	}
	defer httpResponse.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		tflog.Error(ctx, "Failed to read response body", map[string]any{"error": err})
		return false
	}

	// Parse the JSON response.
	var healthCheck HealthCheckResponse
	err = json.Unmarshal(body, &healthCheck)
	if err != nil {
		tflog.Error(ctx, "Failed to parse JSON response from healthcheck", map[string]any{"error": err, "body": string(body)})
		return false
	}
	tflog.Info(ctx, "Retool version", map[string]any{"version": healthCheck.Version})

	return semver.Compare("v"+healthCheck.Version, MINIMUM_RETOOL_VERSION) >= 0
}

// Configure prepares a Retool API client for data sources and resources.
func (p *retoolProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Retool API client")
	// Retrieve provider data from configuration.
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

	// We only check the minimum version if there's no HTTP client override
	// This is a hacky way to avoid doing the check when running acceptance tests in "record" or "replay" mode.
	if p.httpClient == nil && !checkMinimalVersion(ctx, host, scheme) {
		resp.Diagnostics.AddError("Incompatible Retool version", "The Retool instance version is not supported. Minimum version required is "+MINIMUM_RETOOL_VERSION)
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
	// We need this to be able to record and replay HTTP interactions in the acceptance tests.
	if p.httpClient != nil {
		clientConfig.HTTPClient = p.httpClient
	}

	clientConfig.AddDefaultHeader("Authorization", "Bearer "+accessToken)
	client := api.NewAPIClient(clientConfig)

	// Make the Retool client available during DataSource and Resource type Configure methods.
	// Also, init the cache for the root folder ids.
	rootFolderIdCache := make(map[string]string)
	providerData := utils.ProviderData{
		Client:            client,
		RootFolderIdCache: &rootFolderIdCache,
	}
	resp.DataSourceData = &providerData
	resp.ResourceData = &providerData
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
		permissions.NewResource,
		space.NewResource,
		sso.NewResource,
		sourcecontrol.NewResource,
		sourcecontrolsettings.NewResource,
	}
}
