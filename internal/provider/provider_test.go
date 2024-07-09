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
