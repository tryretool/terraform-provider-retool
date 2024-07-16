resource "retool_permissions" "test_permissions" {
  subject = {
    type = "group"
    id   = retool_group.test_group.id
  }
  permissions = [
    {
      object = {
        type = "folder"
        id   = retool_folder.test_folder2.id
      }
      access_level = "own"
    },
  ]
}
