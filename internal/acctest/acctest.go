package acctest

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"gopkg.in/dnaeon/go-vcr.v3/recorder"

	"github.com/tryretool/terraform-provider-retool/internal/provider"
)

const (
	ProviderTestConfig = `
	provider "retool" {
	}
	`
)

// This file borrows heavily from Auth0's Terraform provider: https://github.com/auth0/terraform-provider-auth0/blob/main/internal/acctest/acctest.go

var (
	providerTestFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"retool": providerserver.NewProtocol6WithError(provider.New("test")()),
	}
)

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

func httpRecordingsAreEnabled() bool {
	httpRecordings := os.Getenv("RETOOL_HTTP_RECORDINGS")
	return httpRecordings == "true" || httpRecordings == "1" || httpRecordings == "on"
}

func testFactoriesWithHTTPRecordings(httpRecorder *recorder.Recorder) map[string]func() (tfprotov6.ProviderServer, error) {
	return map[string]func() (tfprotov6.ProviderServer, error){
		"retool": func() (tfprotov6.ProviderServer, error) {
			retoolProvider := provider.NewWithHttpClient("test", httpRecorder.GetDefaultClient())()
			return providerserver.NewProtocol6WithError(retoolProvider)()
		},
	}
}
