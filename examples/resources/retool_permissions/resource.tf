# Folder permissions example
resource "retool_permissions" "folder_permissions" {
  subject = {
    type = "group"
    id   = retool_group.test_group.id
  }
  permissions = [
    {
      object = {
        type = "folder"
        id   = retool_folder.test_folder.id
      }
      access_level = "own"
    },
  ]
}

# Screen permissions example
resource "retool_permissions" "screen_permissions" {
  subject = {
    type = "user"
    id   = retool_user.test_user.id
  }
  permissions = [
    {
      object = {
        type   = "screen"
        id     = "12345678-1234-1234-1234-123456789012"
        app_id = "87654321-4321-4321-4321-210987654321"
      }
      access_level = "use"
    },
  ]
}
