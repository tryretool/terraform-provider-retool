package sso_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/tryretool/terraform-provider-retool/internal/acctest"
)

const testGoogleConfig = `
	resource "retool_sso" "sso" {
		disable_email_password_login = true
		google = {
			client_id     = "client_id"
			client_secret = "client_secret"
		}
	}	
		`
const testOidcConfig = `
	resource "retool_sso" "sso" {
		disable_email_password_login = true
		oidc = {
			client_id     = "client_id"
			client_secret = "client_secret"
			scopes = "openid email profile"
			auth_url = "https://example.com/auth"
			token_url = "https://example.com/token"
			jwt_email_key = "idToken.email"
			jwt_roles_key = "idToken.roles"
			jwt_first_name_key = "idToken.given_name"
			jwt_last_name_key = "idToken.family_name"
			roles_mapping = {
				"admin" = "admin_group"
				"user" = "user_group"
			}
			jit_enabled = true
			restricted_domains = ["example.com", "example.org"]
			trigger_login_automatically = true
		}
	}
`

const testGoogleOidcConfig = `
	resource "retool_sso" "sso" {
		disable_email_password_login = true
		google = {
			client_id     = "client_id"
			client_secret = "client_secret"
		}
		oidc = {
			client_id     = "client_id"
			client_secret = "client_secret"
			scopes = "openid email profile"
			auth_url = "https://example.com/auth"
			token_url = "https://example.com/token"
			jwt_email_key = "idToken.email"
			jwt_roles_key = "idToken.roles"
			jwt_first_name_key = "idToken.given_name"
			jwt_last_name_key = "idToken.family_name"
			roles_mapping = {
				"admin" = "admin_group"
				"user" = "user_group"
			}
			jit_enabled = true
			restricted_domains = ["example.com", "example.org"]
			trigger_login_automatically = true
		}
	}
`

const testSamlConfig = `
	resource "retool_sso" "sso" {
		disable_email_password_login = true
		saml = {
			idp_metadata_xml = "metadata_xml"
			first_name_attribute = "first_name"
			last_name_attribute = "last_name"
			groups_attribute = "groups"
			sync_group_claims = true
			roles_mapping = {
			"admin" = ["admin_group"]
			"user" = ["user_group","another_user_group"]
			}
			jit_enabled = true
			restricted_domains = ["example.com", "example.org"]
			trigger_login_automatically = true

			ldap_sync_group_claims = true
			ldap_config = {
			server_url = "ldap://example.com"
			base_domain_components = "dc=example,dc=com"
			server_name = "ldap.example.com"
			server_key = "server_key"
			server_certificate = "server_cert"
			}
		}
	}
`

const testGoogleSamlConfig = `
	resource "retool_sso" "sso" {
		disable_email_password_login = true
		google = {
			client_id     = "client_id"
			client_secret = "client_secret"
		}
		saml = {
			idp_metadata_xml = "metadata_xml"
			first_name_attribute = "first_name"
			last_name_attribute = "last_name"
			groups_attribute = "groups"
			sync_group_claims = true
			roles_mapping = {
				"admin" = ["admin_group"]
				"user" = ["user_group","another_user_group"]
			}
			jit_enabled = true
			restricted_domains = ["example.com", "example.org"]
			trigger_login_automatically = true

			ldap_sync_group_claims = true
			ldap_config = {
				server_url = "ldap://example.com"
				base_domain_components = "dc=example,dc=com"
				server_name = "ldap.example.com"
				server_key = "server_key"
				server_certificate = "server_cert"
			}
		}
	}
`

const testUpdatedGoogleSamlConfig = `
	resource "retool_sso" "sso" {
		disable_email_password_login = false
		google = {
			client_id     = "new_client_id"
			client_secret = "new_client_secret"
		}
		saml = {
			idp_metadata_xml = "new_metadata_xml"
			first_name_attribute = "new_first_name"
			last_name_attribute = "new_last_name"
			groups_attribute = "new_groups"
			sync_group_claims = false
			roles_mapping = {
				"admin" = ["new_admin_group"]
				"user" = ["new_user_group","another_user_group"]
				"guest" = ["guest_group"]
			}
			jit_enabled = false
			restricted_domains = ["example.com", "newexample.org"]
			trigger_login_automatically = false

			ldap_sync_group_claims = true
			ldap_config = {
				server_url = "ldap://newexample.com"
				base_domain_components = "dc=newexample,dc=com"
				server_name = "ldap.newexample.com"
				server_key = "new_server_key"
				server_certificate = "new_server_cert"
			}
		}
	}
`

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestAccSSO(t *testing.T) {
	acctest.Test(t, resource.TestCase{
		Steps: []resource.TestStep{
			// Read and Create
			{
				Config: testGoogleConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_sso.sso", "google.client_id", "client_id"),
					resource.TestCheckResourceAttr("retool_sso.sso", "google.client_secret", "client_secret"),
					resource.TestCheckResourceAttrSet("retool_sso.sso", "google.encrypted_client_secret"),
					resource.TestCheckResourceAttr("retool_sso.sso", "disable_email_password_login", "true"),
				),
			},
			{
				Config: testOidcConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.client_id", "client_id"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.client_secret", "client_secret"),
					resource.TestCheckResourceAttrSet("retool_sso.sso", "oidc.encrypted_client_secret"),
					resource.TestCheckResourceAttr("retool_sso.sso", "disable_email_password_login", "true"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.scopes", "openid email profile"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.auth_url", "https://example.com/auth"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.token_url", "https://example.com/token"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.jwt_email_key", "idToken.email"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.jwt_roles_key", "idToken.roles"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.jwt_first_name_key", "idToken.given_name"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.jwt_last_name_key", "idToken.family_name"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.roles_mapping.admin", "admin_group"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.roles_mapping.user", "user_group"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.jit_enabled", "true"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.restricted_domains.#", "2"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.restricted_domains.0", "example.com"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.restricted_domains.1", "example.org"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.trigger_login_automatically", "true"),
				),
			},
			{
				Config: testGoogleOidcConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_sso.sso", "google.client_id", "client_id"),
					resource.TestCheckResourceAttr("retool_sso.sso", "google.client_secret", "client_secret"),
					resource.TestCheckResourceAttrSet("retool_sso.sso", "google.encrypted_client_secret"),
					resource.TestCheckResourceAttr("retool_sso.sso", "disable_email_password_login", "true"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.client_id", "client_id"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.client_secret", "client_secret"),
					resource.TestCheckResourceAttrSet("retool_sso.sso", "oidc.encrypted_client_secret"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.scopes", "openid email profile"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.auth_url", "https://example.com/auth"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.token_url", "https://example.com/token"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.jwt_email_key", "idToken.email"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.jwt_roles_key", "idToken.roles"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.jwt_first_name_key", "idToken.given_name"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.jwt_last_name_key", "idToken.family_name"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.roles_mapping.admin", "admin_group"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.roles_mapping.user", "user_group"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.jit_enabled", "true"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.restricted_domains.#", "2"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.restricted_domains.0", "example.com"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.restricted_domains.1", "example.org"),
					resource.TestCheckResourceAttr("retool_sso.sso", "oidc.trigger_login_automatically", "true"),
				),
			},
			{
				Config: testSamlConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_sso.sso", "disable_email_password_login", "true"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.idp_metadata_xml", "metadata_xml"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.first_name_attribute", "first_name"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.last_name_attribute", "last_name"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.groups_attribute", "groups"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.sync_group_claims", "true"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.admin.#", "1"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.admin.0", "admin_group"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.user.#", "2"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.user.0", "user_group"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.user.1", "another_user_group"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.jit_enabled", "true"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.restricted_domains.#", "2"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.restricted_domains.0", "example.com"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.restricted_domains.1", "example.org"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.trigger_login_automatically", "true"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_sync_group_claims", "true"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_config.server_url", "ldap://example.com"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_config.base_domain_components", "dc=example,dc=com"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_config.server_name", "ldap.example.com"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_config.server_key", "server_key"),
					resource.TestCheckResourceAttrSet("retool_sso.sso", "saml.ldap_config.encrypted_server_key"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_config.server_certificate", "server_cert"),
					resource.TestCheckResourceAttrSet("retool_sso.sso", "saml.ldap_config.encrypted_server_certificate"),
				),
			},
			{
				Config: testGoogleSamlConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_sso.sso", "google.client_id", "client_id"),
					resource.TestCheckResourceAttr("retool_sso.sso", "google.client_secret", "client_secret"),
					resource.TestCheckResourceAttrSet("retool_sso.sso", "google.encrypted_client_secret"),
					resource.TestCheckResourceAttr("retool_sso.sso", "disable_email_password_login", "true"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.idp_metadata_xml", "metadata_xml"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.first_name_attribute", "first_name"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.last_name_attribute", "last_name"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.groups_attribute", "groups"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.sync_group_claims", "true"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.admin.#", "1"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.admin.0", "admin_group"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.user.#", "2"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.user.0", "user_group"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.user.1", "another_user_group"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.jit_enabled", "true"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.restricted_domains.#", "2"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.restricted_domains.0", "example.com"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.restricted_domains.1", "example.org"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.trigger_login_automatically", "true"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_sync_group_claims", "true"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_config.server_url", "ldap://example.com"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_config.base_domain_components", "dc=example,dc=com"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_config.server_name", "ldap.example.com"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_config.server_key", "server_key"),
					resource.TestCheckResourceAttrSet("retool_sso.sso", "saml.ldap_config.encrypted_server_key"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_config.server_certificate", "server_cert"),
					resource.TestCheckResourceAttrSet("retool_sso.sso", "saml.ldap_config.encrypted_server_certificate"),
				),
			},
			// Update and Read
			{
				Config: testUpdatedGoogleSamlConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("retool_sso.sso", "google.client_id", "new_client_id"),
					resource.TestCheckResourceAttr("retool_sso.sso", "google.client_secret", "new_client_secret"),
					resource.TestCheckResourceAttrSet("retool_sso.sso", "google.encrypted_client_secret"),
					resource.TestCheckResourceAttr("retool_sso.sso", "disable_email_password_login", "false"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.idp_metadata_xml", "new_metadata_xml"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.first_name_attribute", "new_first_name"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.last_name_attribute", "new_last_name"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.groups_attribute", "new_groups"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.sync_group_claims", "false"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.admin.#", "1"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.admin.0", "new_admin_group"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.user.#", "2"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.user.0", "new_user_group"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.user.1", "another_user_group"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.guest.#", "1"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.roles_mapping.guest.0", "guest_group"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.jit_enabled", "false"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.restricted_domains.#", "2"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.restricted_domains.0", "example.com"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.restricted_domains.1", "newexample.org"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.trigger_login_automatically", "false"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_sync_group_claims", "true"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_config.server_url", "ldap://newexample.com"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_config.base_domain_components", "dc=newexample,dc=com"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_config.server_name", "ldap.newexample.com"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_config.server_key", "new_server_key"),
					resource.TestCheckResourceAttrSet("retool_sso.sso", "saml.ldap_config.encrypted_server_key"),
					resource.TestCheckResourceAttr("retool_sso.sso", "saml.ldap_config.server_certificate", "new_server_cert"),
					resource.TestCheckResourceAttrSet("retool_sso.sso", "saml.ldap_config.encrypted_server_certificate"),
				),
			},
		},
	})
}

func sweepSso(region string) error {
	log.Printf("Sweeping SSO config in region %s", region)
	client, err := acctest.SweeperClient()
	if err != nil {
		return err
	}

	_, err = client.SSOAPI.SsoConfigDelete(context.Background()).Execute()
	if err != nil {
		return fmt.Errorf("Error deleting SSO config: %s", err.Error())
	}
	return nil
}

func init() {
	resource.AddTestSweepers("retool_sso", &resource.Sweeper{
		Name: "retool_sso",
		F:    sweepSso,
	})
}
