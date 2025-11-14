data "retool_users" "all_users" {}

output "all_users" {
  value = data.retool_users.all_users.users
}

# Filter active users
output "active_users" {
  value = [for user in data.retool_users.all_users.users : user if user.active]
}

# Filter admin users
output "admin_users" {
  value = [for user in data.retool_users.all_users.users : user if user.is_admin]
}

