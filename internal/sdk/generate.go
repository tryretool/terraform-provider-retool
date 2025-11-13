// Package sdk : contains the generated code from the openAPI spec
//
//go:generate sh -c "OPENAPI_GENERATOR_VERSION=7.6.0 openapi-generator-cli generate -i openAPISpec.json -g go -o ./api -c generatorConfig.yaml --minimal-update && ./fix_generated_code.sh"
package sdk
