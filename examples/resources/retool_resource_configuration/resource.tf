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

  lifecycle {
    # Ignore changes to options since the API adds default values
    ignore_changes = [options]
  }
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

  lifecycle {
    ignore_changes = [options]
  }
}

