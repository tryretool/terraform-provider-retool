resource "retool_group" "example" {
  name                       = "Test Group"
  universal_app_access       = "use"
  universal_resource_access  = "own"
  universal_workflow_access  = "none"
  user_list_access           = false
  audit_log_access           = true
  unpublished_release_access = false
  usage_analytics_access     = true
  account_details_access     = true
  landing_page_app_id        = "c37676ba-116f-11ea-b17d-d7734e1526f2"
}
