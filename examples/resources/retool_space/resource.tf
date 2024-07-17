resource "retool_space" "test_space" {
  name   = "Test Space"
  domain = "space1.example.com"
}

resource "retool_space" "test_space_with_create_options" {
  name   = "Test Space with creation options"
  domain = "space2.example.com"
  create_options = { # Note that changing any of these values will result in the old space being deleted and a new space being created
    copy_sso_settings                 = true
    copy_branding_and_themes_settings = true
    create_admin_user                 = false
    users_to_copy_as_admins           = ["admin@example.com"]
  }
}
