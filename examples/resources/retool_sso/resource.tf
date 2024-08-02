resource "retool_sso" "sso" {
  disable_email_password_login = true
  google = {
    client_id     = "client_id"
    client_secret = "client_secret"
  }
  saml = {
    idp_metadata_xml     = "metadata_xml"
    first_name_attribute = "first_name"
    last_name_attribute  = "last_name"
    groups_attribute     = "groups"
    sync_group_claims    = true
    roles_mapping = {
      "admin" = ["admin_group"]
      "user"  = ["user_group", "another_user_group"]
    }
    jit_enabled                 = true
    restricted_domains          = ["example.com", "example.org"]
    trigger_login_automatically = true

    ldap_sync_group_claims = true
    ldap_config = {
      server_url             = "ldap://example.com"
      base_domain_components = "dc=example,dc=com"
      server_name            = "ldap.example.com"
      server_key             = "server_key"
      server_certificate     = "server_cert"
    }
  }
}
