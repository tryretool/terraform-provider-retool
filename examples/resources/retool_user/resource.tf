resource "retool_user" "example" {
  email      = "user@example.com"
  first_name = "John"
  last_name  = "Doe"
  active     = true
  metadata   = "{\"department\":\"engineering\",\"team\":\"platform\"}"
  user_type  = "standard"
}

