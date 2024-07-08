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

To run unit tests, run the following command from the root folder:
```
make test-unit
```

You can run specific tests by setting `FILTER` env var:
```
FILTER=TestNewProvider make test-unit
``` 

To run acceptance tests, set `RETOOL_HOST`, `RETOOL_SCHEME` and `RETOOL_ACCESS_TOKEN` env vars, then run
```
make test-acc
```
Note that the acceptance tests will create, modify and delete resources in the Retool org.
You can run specific tests by setting `FILTER` env var as well.

## Documentation
Run `go generate` in the root of this repository to generate provider docs. See https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-documentation-generation for more details on how to add new examples.
