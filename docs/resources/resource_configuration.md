---
page_title: "retool_resource_configuration Resource - terraform-provider-retool"
subcategory: ""
description: |-
  Manages a resource configuration in Retool.
---

# retool_resource_configuration (Resource)

A resource configuration defines environment-specific configuration for a Retool resource. This allows you to configure different connection settings, credentials, and options for a resource across different environments (e.g., production, staging, development).

For example, you might have a PostgreSQL resource that connects to different database instances in production vs. staging environments, each with different credentials and connection strings.

## Example Usage

```terraform
# First, create a resource
resource "retool_resource" "postgres" {
  display_name = "PostgreSQL Database"
  type         = "postgresql"
  options = jsonencode({
    host     = "localhost"
    port     = 5432
    database = "mydb"
    user     = "default_user"
  })
}

# Get environments
data "retool_environments" "all" {}

# Create a production-specific configuration
resource "retool_resource_configuration" "postgres_production" {
  resource_id    = retool_resource.postgres.id
  environment_id = [for env in data.retool_environments.all.environments : env.id if env.name == "production"][0]

  options = jsonencode({
    host     = "prod-db.example.com"
    port     = 5432
    database = "production_db"
    user     = "prod_user"
    password = "prod_password"
    ssl      = true
  })
}

# Create a staging-specific configuration
resource "retool_resource_configuration" "postgres_staging" {
  resource_id    = retool_resource.postgres.id
  environment_id = [for env in data.retool_environments.all.environments : env.id if env.name == "staging"][0]

  options = jsonencode({
    host     = "staging-db.example.com"
    port     = 5432
    database = "staging_db"
    user     = "staging_user"
    password = "staging_password"
    ssl      = true
  })
}
```

### REST API Example

```terraform
# Create a REST API resource
resource "retool_resource" "api" {
  display_name = "External API"
  type         = "restapi"
  options = jsonencode({
    base_url = "https://api.example.com"
  })
}

# Production configuration with different base URL and authentication
resource "retool_resource_configuration" "api_production" {
  resource_id    = retool_resource.api.id
  environment_id = data.retool_environments.all.environments[0].id

  options = jsonencode({
    base_url = "https://api.production.example.com"
    headers = {
      "Authorization" = "Bearer ${var.prod_api_token}"
      "X-Environment" = "production"
    }
  })
}
```

## Schema

### Required

- `resource_id` (String) The UUID or name of the resource to configure. Cannot be changed after creation.
- `environment_id` (String) The UUID of the environment for this configuration. Cannot be changed after creation.
- `options` (String, Sensitive) JSON string containing the environment-specific resource configuration options. The structure varies by resource type and mirrors the options structure used when creating a resource.

### Read-Only

- `id` (String) The UUID of the resource configuration.
- `created_at` (String) The timestamp when the resource configuration was created.
- `updated_at` (String) The timestamp when the resource configuration was last updated.

## Import

Resource configurations can be imported using their ID:

```shell
terraform import retool_resource_configuration.example 01234567-89ab-cdef-0123-456789abcdef
```

**Note:** After importing, you will need to manually set the `options` field in your Terraform configuration to match the current configuration.

