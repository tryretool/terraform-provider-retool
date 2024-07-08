package acctest

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/tryretool/terraform-provider-retool/internal/provider"
)

const (
	ProviderTestConfig = `
	provider "retool" {
	}
	`
)

var (
	providerTestFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"retool": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
)

func Test(t *testing.T, testCase resource.TestCase) {
	testCase.ProtoV6ProviderFactories = providerTestFactories
	resource.Test(t, testCase)
}
