resource "retool_source_control_settings" "scm_settings" {
  auto_branch_naming_enabled           = false
  custom_pull_request_template_enabled = true
  custom_pull_request_template         = "custom-pull-request-template"
  version_control_locked               = true
  force_uuid_mapping                   = false
}
