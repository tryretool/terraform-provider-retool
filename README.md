# Retool Terraform Provider
Work in progress, more docs to come later

The repository contains code-generated API client library for Retool API - see more details in [internal/sdk/README.md](/internal/sdk/README.md)

## Development
Install Go, make sure `GOPATH` variable is set and `$GOPATH/bin` is in the `PATH`.

Follow the [instructions](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider#prepare-terraform-for-local-provider-install) to use local version of Retool provider.

## Testing
You can test the provider, resources and data sources against your local dev:
```
cd examples/provider-initialization
RETOOL_HOST=localhost:3000 RETOOL_SCHEME=http RETOOL_ACCESS_TOKEN=<the API token you created> terraform plan
```

## Documentation
Run `go generate` in the root of this repository to generate provider docs. See https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-documentation-generation for more details on how to add new examples.