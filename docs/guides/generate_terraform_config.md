---
page_title: Auto-generating Terraform config from existing Retool organization
description: |-
  How to auto-generate Terraform configuration files from your Retool organization
---

# Auto-generating Terraform config files

You can use [Retool CLI](https://www.npmjs.com/package/retool-cli) to automatically generate Terraform configuration based on existing Retool organization. 
This can be useful for controlling that organization via Terraform or copying its configuration to some other organization.

**NOTE** Auto-generation functionality is _experimental_ and is primarily intended to generate a best-effort approximation of Terraform config, which in many cases will still need to be manually tweaked in order to work properly. Use with caution.

## 1. Install Retool CLI tool
Follow the instructions [here](https://www.npmjs.com/package/retool-cli). You'll need version 1.0.29 or newer.

## 2. Create API access token in your org and set environment variables

Go to **Settings** > **Retool API** in your Retool organization and create access token with the following scopes: 
- source_control:read
- groups:read
- spaces:read
- folders:read
- permissions:all:read

Add `write` scopes if you plan to reuse the token in Retool Terraform provider later.

Set the following environment variables in your terminal:
```sh
export RETOOL_HOST=retool.example.com # domain name of the Retool org
export RETOOL_ACCESS_TOKEN=retool_123sdkfjsld # API access token
export RETOOL_SCHEME=http # Only if your Retool instance is using HTTP instead of HTTPS
```

## 3. Generate Terraform configuration
You can now run Retool CLI in your terminal:

```sh
retool terraform -c generated.tf -i import.tf
```

The file `generated.tf` will contain Terraform resource definitions. For example:
```hcl
resource "retool_folder" "production_apps" {
  name = "Production Apps"
  folder_type = "app"
}

resource "retool_group" "production_apps_user" {
  name = "Production Apps User"
  universal_app_access = "none"
  universal_resource_access = "none"
  universal_workflow_access = "none"
  universal_query_library_access = "none"
  user_list_access = false
  audit_log_access = false
  unpublished_release_access = false
  usage_analytics_access = false
  account_details_access = false
}
```

The file `import.tf` will map Terraform resources to ids of the objects in your Retool organization, so that Terraform knows not to create new ones:
```hcl
import {
  to = retool_folder.production_apps
  id = "app_74"
}

import {
  to = retool_group.production_apps_user
  id = "55"
}
```

## 4. Configure Retool Terraform provider and apply your configuration
Create a file `main.tf` and add the following code to it:
```hcl
terraform {
    required_providers {
        retool = {
            source = "tryretool/retool"
        }
    }
}

# Configure the Retool provider with the access token.
provider "retool" {
}
```

Note that in this example, the provider will use the environment variables you defined earlier to get the hostname and API access token. 
Depending on your scenario, you might want to tweak this configuration to use non-default Terraform backend, read the API access token from a secret manager etc.

## 5. Apply the configuration your Retool organization
You can now run:
```sh
terraform init
terraform plan
```

Carefully inspect Terraform output to make sure it will not inadvertently modify or destroy objects in your Retool organization.
Once you're happy with the plan, you can apply it:
```sh
terraform apply
```

# Other scenarios and common pitfalls

## Copying your configuration to another Retool organization
You can follow the process described above with the following modifications:
- You don't need to generate `imports.tf` file - just drop the `-i` option when running Retool CLI in step 3.
- You'll need to provide the hostname of your target organization when configuring Retool provider in step 4.
  
## Sensitive values
Retool CLI can't read unecrypted secrets from Retool org, so any sensitive values will be set to `null` in the generated Terraform config. 

## Hardcoded object ids
Retool CLI makes the best effort to generate "portable" Terraform configuration that can be applied to other Retool organizations.
However, in some scenarios the generated configuration will contain object IDs (app and workflow UUIDs, for example) from the source Retool organization.
Since these ids won't exist in the target Retool organization, you'd need to remove or update them in the config before trying to apply it to another Retool organization.

For example, the generated config may look like this:
```hcl
resource "retool_folder" "production_apps" {
  name = "Production Apps"
  folder_type = "app"
}

resource "retool_group" "production_apps_user" {
  name = "Production Apps User"
  universal_app_access = "none"
  universal_resource_access = "none"
  universal_workflow_access = "none"
  universal_query_library_access = "none"
  user_list_access = false
  audit_log_access = false
  unpublished_release_access = false
  usage_analytics_access = false
  account_details_access = false
}

resource "retool_permissions" "production_apps_user_permissions" {
  subject = {
    type = "group"
    id = retool_group.production_apps_user.id
  }
  permissions = [
    {
      object = {
        type = "folder"
        id = retool_folder.production_apps.id
      }
      access_level = "use"
    },
    {
      object = {
        type = "app"
        id = "123-29384-abcdef"
      }
      access_level = "use"
    },
  ]
}
```

This fragment is not portable, since it refers to UUID of an app that only exists in the source Retool organization:
```hcl
...
    {
      object = {
        type = "app"
        id = "123-29384-abcdef"
      }
      access_level = "use"
    },
...
```
