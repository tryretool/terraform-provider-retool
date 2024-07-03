terraform {
  required_providers {
    retool = {
      source = "registry.terraform.io/retool/retool"
    }
  }
}

provider "retool" {}

data "retool_folders" "all_folders" {}

output "all_folders" {
  value = data.retool_folders.all_folders.folders
}
