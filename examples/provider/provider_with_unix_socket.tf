provider "retool" {
  host             = "localhost"
  scheme           = "http"
  access_token     = "your-access-token"
  unix_socket_path = "/var/run/retool.sock"
}

# When using unix_socket_path, the provider will connect to the Retool API
# via the specified Unix socket instead of using TCP connections.
# This is useful for scenarios where:
# - Retool is running in a container and exposing a Unix socket
# - You want to reduce network overhead by using local socket communication
# - You need to connect through a Unix socket proxy

# Example resource using the unix socket connection
resource "retool_folder" "example_apps" {
  name        = "Example Apps"
  folder_type = "app"
}
