resource "retool_source_control" "scm" {
  org            = "my-org"
  repo           = "my-repo"
  default_branch = "main"
  repo_version   = "1.0.0"
  github = {
    app_authentication = {
      app_id          = "123456"
      installation_id = "123457"
      private_key     = "private-key"
    }
    url                = "https://github.com"
    enterprise_api_url = "https://github.mycompany.com/api/v3"
  }
}
