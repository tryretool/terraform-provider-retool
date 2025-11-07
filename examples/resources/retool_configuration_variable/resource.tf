resource "retool_configuration_variable" "test_config_var" {
  name = "Test Config Var"
  values = [
    {
      environment_id = "a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d" # prod
      value          = "value1"
    },
    {
      environment_id = "b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e" # staging
      value          = "value2"
    }
  ]
}

resource "retool_configuration_variable" "test_config_var_with_description" {
  name        = "Test Config Var with description"
  description = "This is a test configuration variable with a description"
  values = [
    {
      environment_id = "a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d" # prod
      value          = "value1"
    },
    {
      environment_id = "b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e" # staging
      value          = "value2"
    }
  ]
}

resource "retool_configuration_variable" "test_config_var_as_secret" {
  name        = "Test Secret Config Var"
  description = "This is a secret configuration variable"
  secret      = true
  values = [
    {
      environment_id = "a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d" # prod
      value          = "secret_value1"
    },
    {
      environment_id = "b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e" # staging
      value          = "secret_value2"
    }
  ]
}
