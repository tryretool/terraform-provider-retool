---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "retool_folder Resource - retool"
subcategory: ""
description: |-
  
---

# retool_folder (Resource)



## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `folder_type` (String) The type of the folder: (app|file|resource|workflow).
- `name` (String) The name of the folder.

### Optional

- `parent_folder_id` (String) The id of the parent folder.

### Read-Only

- `id` (String) The id of the folder. Currently this is the same as legacy_id but will be different in the future.
- `is_system_folder` (Boolean) Whether the folder is a system folder.
- `legacy_id` (String) The legacy id of the folder.

## Import

Import is supported using the following syntax:

```shell
# Folder can be imported by specifying the id
terraform import retool_folder.example app_123
```