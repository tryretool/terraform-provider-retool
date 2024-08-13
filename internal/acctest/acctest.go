// Package acctest provides a harness for our acceptance tests and test sweepers.
package acctest

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/tryretool/terraform-provider-retool/internal/provider"
	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// This file borrows heavily from Auth0's Terraform provider: https://github.com/auth0/terraform-provider-auth0/blob/main/internal/acctest/acctest.go
var (
	providerTestFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"retool": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
)

// Wrapper around resource.Test that allows for HTTP recordings.
func Test(t *testing.T, testCase resource.TestCase) {
	if httpRecordingsAreEnabled() {
		httpRecorder := newHTTPRecorder(t)
		testCase.ProtoV6ProviderFactories = testFactoriesWithHTTPRecordings(httpRecorder)
		resource.ParallelTest(t, testCase)

		return
	}

	testCase.ProtoV6ProviderFactories = providerTestFactories
	resource.Test(t, testCase)
}

// Init API client to be used by the test sweepers.
func SweeperClient() (*api.APIClient, error) {
	host := os.Getenv("RETOOL_HOST")
	if host == "" {
		return nil, fmt.Errorf("RETOOL_HOST must be set")
	}
	scheme := os.Getenv("RETOOL_SCHEME")
	if scheme == "" {
		scheme = "https"
	}
	accessToken := os.Getenv("RETOOL_ACCESS_TOKEN")
	if accessToken == "" {
		return nil, fmt.Errorf("RETOOL_ACCESS_TOKEN must be set")
	}

	clientConfig := api.NewConfiguration()
	clientConfig.Host = host
	clientConfig.Scheme = scheme
	clientConfig.Servers = api.ServerConfigurations{
		api.ServerConfiguration{
			URL: "/api/v2",
		},
	}

	clientConfig.AddDefaultHeader("Authorization", "Bearer "+accessToken)
	return api.NewAPIClient(clientConfig), nil
}

func httpRecordingsAreEnabled() bool {
	httpRecordings := os.Getenv("RETOOL_HTTP_RECORDINGS")
	return httpRecordings == "true" || httpRecordings == "1" || httpRecordings == "on"
}

func testFactoriesWithHTTPRecordings(httpRecorder *recorder.Recorder) map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"retool": func() (tfprotov6.ProviderServer, error) {
			retoolProvider := provider.NewWithHTTPClient("test", httpRecorder.GetDefaultClient())()
			return providerserver.NewProtocol6WithError(retoolProvider)()
		},
	}
}
