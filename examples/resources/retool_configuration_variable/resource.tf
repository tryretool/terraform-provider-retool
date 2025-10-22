resource "retool_configuration_variable" "test_config_var" {
  name = "Test Space"
  value = [
    {
      environment_id = "production"
      value          = "value1"
    },
    {
      environment_id = "staging"
      value          = "value2"
    }
  ]
}

resource "retool_configuration_variable" "test_config_var_with_descritpion" {
  name        = "Test Space with creation options"
  description = "This is a test configuration variable with a description"
  value = [
    {
      environment_id = "production"
      value          = "value1"
    },
    {
      environment_id = "staging"
      value          = "value2"
    }
  ]
}

resource "retool_configuration_variable" "test_config_var_as_secret" {
  name        = "Test Space with creation options"
  description = "This is a test configuration variable with a description"
  secret      = true
  value = [
    {
      environment_id = "production"
      value          = "value1"
    },
    {
      environment_id = "staging"
      value          = "value2"
    }
  ]
}
