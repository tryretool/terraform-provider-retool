package sdk

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/tryretool/terraform-provider-retool/internal/sdk/api"
)

// A provider build outlives the spec it was generated from: Retool keeps adding
// response fields, and every instance on a newer release sends them. Decoding has
// to ignore what it doesn't recognize instead of failing the read.
func TestUnmarshal_IgnoresUnknownResponseFields(t *testing.T) {
	t.Parallel()

	body := `{
		"success": true,
		"data": {
			"id": "my-api",
			"type": "restapi",
			"display_name": "My API",
			"folder_id": null,
			"protected": false,
			"some_field_from_a_future_release": {"nested": ["anything"]},
			"created_at": "2019-02-08T11:45:48.899Z",
			"updated_at": "2019-02-24T18:28:18.790Z"
		}
	}`

	var response api.ResourcesResourceIdGet200Response
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		t.Fatalf("unexpected error decoding response with an unknown field: %s", err)
	}

	if response.Data.Id != "my-api" {
		t.Errorf("expected id my-api, got %s", response.Data.Id)
	}
	if response.Data.DisplayName != "My API" {
		t.Errorf("expected display_name My API, got %s", response.Data.DisplayName)
	}
}

// Dropping DisallowUnknownFields must not also drop the required-property check.
func TestUnmarshal_StillRejectsMissingRequiredFields(t *testing.T) {
	t.Parallel()

	body := `{
		"success": true,
		"data": {
			"type": "restapi",
			"display_name": "My API",
			"protected": false,
			"created_at": "2019-02-08T11:45:48.899Z",
			"updated_at": "2019-02-24T18:28:18.790Z"
		}
	}`

	var response api.ResourcesResourceIdGet200Response
	err := json.Unmarshal([]byte(body), &response)
	if err == nil {
		t.Fatal("expected an error for a response missing the required id field")
	}
	if !strings.Contains(err.Error(), "id") {
		t.Errorf("expected the error to name the missing property, got: %s", err)
	}
}

// oneOf/anyOf variants keep DisallowUnknownFields — it is what tells the variants
// apart. GoogleSAML is the case that proves it: its required properties are a
// superset of both Google's and SAML's, so the required-property check alone
// cannot rule those out. Without strictness all three match and the union fails
// with "data matches more than one schema in oneOf".
func TestUnmarshal_UnionVariantsStayDiscriminated(t *testing.T) {
	t.Parallel()

	body := `{
		"config_type": "google-saml",
		"google_client_id": "client-id",
		"google_client_secret": "client-secret",
		"disable_email_password_login": false,
		"idp_metadata_xml": "<EntityDescriptor/>",
		"saml_first_name_attribute": "firstName",
		"saml_last_name_attribute": "lastName",
		"saml_sync_group_claims": false,
		"jit_enabled": false,
		"trigger_login_automatically": false
	}`

	var ssoConfig api.SsoConfigPost200ResponseData
	if err := json.Unmarshal([]byte(body), &ssoConfig); err != nil {
		t.Fatalf("unexpected error decoding SSO config: %s", err)
	}

	if ssoConfig.GoogleSAML == nil {
		t.Fatal("expected the payload to resolve to the GoogleSAML variant")
	}
	if ssoConfig.Google != nil || ssoConfig.SAML != nil || ssoConfig.OIDC != nil || ssoConfig.GoogleOIDC != nil {
		t.Error("expected exactly one variant to match")
	}
}

// anyOf unions return on the first variant that decodes, so a lenient variant
// earlier in the list swallows payloads meant for a later one. Snowflake's first
// variant requires only authentication_type; without strictness a username /
// password payload matches it and both credentials are silently dropped.
func TestUnmarshal_AnyOfPicksTheVariantThatFits(t *testing.T) {
	t.Parallel()

	body := `{
		"authentication_type": "basic",
		"username": "svc_account",
		"password": "hunter2"
	}`

	var authOptions api.SnowflakeOptionsAuthenticationOptions
	if err := json.Unmarshal([]byte(body), &authOptions); err != nil {
		t.Fatalf("unexpected error decoding Snowflake auth options: %s", err)
	}

	if authOptions.SnowflakeOptionsAuthenticationOptionsAnyOf1 == nil {
		t.Fatal("expected the payload to resolve to the username/password variant")
	}
	if authOptions.SnowflakeOptionsAuthenticationOptionsAnyOf != nil {
		t.Error("expected the bearer-token variant not to match")
	}

	credentials := authOptions.SnowflakeOptionsAuthenticationOptionsAnyOf1
	if credentials.Username != "svc_account" || credentials.Password != "hunter2" {
		t.Errorf("credentials were dropped: username=%q password=%q", credentials.Username, credentials.Password)
	}
}
