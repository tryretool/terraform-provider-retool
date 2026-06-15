package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/stretchr/testify/assert"
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
