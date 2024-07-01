resource "retool_folder" "example" {
  name        = "Terraform Example Folder"
  folder_type = "app"
}

resource "retool_folder" "example_subfolder" {
  name             = "Example Subfolder"
  folder_type      = "app"
  parent_folder_id = retool_folder.example.id
}

resource "retool_folder" "resource_example" {
  name        = "Terraform Example Folder"
  folder_type = "resource"
}