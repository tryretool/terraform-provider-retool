resource "retool_environment" "development" {
  name        = "Development"
  description = "Development environment"
  color       = "#FFA500"
}

resource "retool_environment" "staging" {
  name        = "Staging"
  description = "Staging environment for testing"
  color       = "#FFFF00"
}

resource "retool_environment" "production" {
  name        = "Production"
  description = "Production environment"
  color       = "#FF0000"
}

resource "retool_environment" "minimal" {
  name  = "Testing"
  color = "#00FF00"
}

