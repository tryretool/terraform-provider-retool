data "retool_environments" "all_environments" {}

output "all_environments" {
  value = data.retool_environments.all_environments.environments
}
