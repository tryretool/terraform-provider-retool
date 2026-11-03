package provider

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/tryretool/terraform-provider-retool/internal/provider/utils"
)

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestNewProvider(t *testing.T) {
	version := "1.0.0"
	p := New(version)()

	assert.NotNil(t, p)
	assert.Equal(t, version, p.(*retoolProvider).version)
}

func TestProviderMetadata(t *testing.T) {
	p := &retoolProvider{version: "1.0.0"}

	req := provider.MetadataRequest{}
	resp := provider.MetadataResponse{}
	p.Metadata(context.Background(), req, &resp)

	assert.Equal(t, "retool", resp.TypeName)
}

func TestIsCompatibleVersion(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    bool
	}{
		{"clean version above minimum", "4.8.0", true},
		{"exact minimum", "4.0.0", true},
		{"below minimum", "3.334.0", false},
		// Retool Cloud 4.x reports a build-hash suffix that x/mod/semver
		// rejects as an invalid (leading-zero) pre-release identifier.
		{"build-hash suffix above minimum", "4.8.0-0838649", true},
		{"build-hash suffix below minimum", "3.99.0-0838649", false},
		{"build-metadata suffix above minimum", "4.8.0+0838649", true},
		{"empty version", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isCompatibleVersion(tt.version))
		})
	}
}

func configureProvider(t *testing.T, p provider.Provider, host string, requestsPerMinute *int64) *http.Client {
	t.Helper()
	ctx := context.Background()

	schemaResp := provider.SchemaResponse{}
	p.Schema(ctx, provider.SchemaRequest{}, &schemaResp)

	rpm := tftypes.NewValue(tftypes.Number, nil)
	if requestsPerMinute != nil {
		rpm = tftypes.NewValue(tftypes.Number, *requestsPerMinute)
	}
	raw := tftypes.NewValue(schemaResp.Schema.Type().TerraformType(ctx), map[string]tftypes.Value{
		"host":                tftypes.NewValue(tftypes.String, host),
		"scheme":              tftypes.NewValue(tftypes.String, "http"),
		"access_token":        tftypes.NewValue(tftypes.String, "test-token"),
		"requests_per_minute": rpm,
		"unix_socket_path":    tftypes.NewValue(tftypes.String, nil),
	})

	resp := provider.ConfigureResponse{}
	p.Configure(ctx, provider.ConfigureRequest{Config: tfsdk.Config{Schema: schemaResp.Schema, Raw: raw}}, &resp)
	require.False(t, resp.Diagnostics.HasError(), "%v", resp.Diagnostics)

	providerData, ok := resp.ResourceData.(*utils.ProviderData)
	require.True(t, ok)
	return providerData.Client.GetConfig().HTTPClient
}

func TestConfigureDoesNotModifySharedHTTPClients(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"version":"4.8.0"}`))
	}))
	defer server.Close()
	host := strings.TrimPrefix(server.URL, "http://")
	t.Setenv("RETOOL_HOST", "")
	t.Setenv("RETOOL_SCHEME", "")
	t.Setenv("RETOOL_ACCESS_TOKEN", "")

	t.Run("http.DefaultClient is left alone", func(t *testing.T) {
		originalTransport := http.DefaultClient.Transport
		defer func() { http.DefaultClient.Transport = originalTransport }()

		apiClient := configureProvider(t, New("test")(), host, nil)
		configureProvider(t, New("test")(), host, nil)

		assert.Equal(t, originalTransport, http.DefaultClient.Transport)
		assert.IsType(t, &utils.ThrottledTransport{}, apiClient.Transport)
	})

	t.Run("client passed to NewWithHTTPClient is left alone", func(t *testing.T) {
		client := &http.Client{}
		p := NewWithHTTPClient("test", client)()

		apiClient := configureProvider(t, p, host, nil)
		configureProvider(t, p, host, nil)

		assert.Nil(t, client.Transport)
		assert.IsType(t, &utils.ThrottledTransport{}, apiClient.Transport)
	})

	t.Run("disabling rate limiting after an earlier Configure", func(t *testing.T) {
		originalTransport := http.DefaultClient.Transport
		defer func() { http.DefaultClient.Transport = originalTransport }()

		configureProvider(t, New("test")(), host, nil)
		disabled := int64(-1)
		apiClient := configureProvider(t, New("test")(), host, &disabled)

		_, throttled := apiClient.Transport.(*utils.ThrottledTransport)
		assert.False(t, throttled, "requests_per_minute = -1 should not be rate limited")
	})
}
