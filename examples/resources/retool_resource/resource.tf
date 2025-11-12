resource "retool_resource" "api" {
  display_name = "My REST API"
  type         = "restapi"

  options = jsonencode({
    base_url = "https://api.example.com"
  })

  lifecycle {
    # Ignore changes to options since they can't be read back from the API
    ignore_changes = [options]
  }
}

resource "retool_resource" "api" {
  display_name = "My REST API"
  type         = "restapi"

  options = jsonencode({
    base_url = "https://api.example.com"
    authentication_options = {
      authentication_type = "basic"
      basic_username      = "admin"
      basic_password      = "password"
    }
  })

  lifecycle {
    # Ignore changes to options since they can't be read back from the API
    ignore_changes = [options]
  }
}

resource "retool_resource" "database" {
  display_name = "Production Database"
  type         = "postgresql"

  options = jsonencode({
    host              = "db.example.com"
    port              = 5432
    database_name     = "myapp"
    database_username = var.db_username
    database_password = var.db_password
    ssl_settings = {
      ssl_enabled = true
    }
  })

  lifecycle {
    ignore_changes = [options]
  }
}

