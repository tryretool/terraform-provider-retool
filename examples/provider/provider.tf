provider "retool" {
  host         = "your-retool-instance.com"
  scheme       = "https"
  access_token = "your-access-token"
}


# Create an app folder
resource "retool_folder" "production_apps" {
  name        = "Production Apps"
  folder_type = "app"
}

# Create a permission group
resource "retool_group" "production_apps_users" {
  name = "Production Apps Users"
}

# Grant "use" permissions on the folder to the permission group
resource "retool_permissions" "production_apps_users" {
  subject = {
    type = "group"
    id   = retool_group.production_apps_users.id
  }
  permissions = [
    {
      object = {
        type = "folder"
        id   = retool_folder.production_apps.id
      }
      access_level = "use"
    }
  ]
}
