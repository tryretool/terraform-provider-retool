data "retool_groups" "all_groups" {}

output "all_groups" {
  value = data.retool_groups.all_groups.groups
}
