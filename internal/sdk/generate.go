//go:generate openapi-generator generate -i openAPISpec.json -g go -o ./api -c generatorConfig.yaml --minimal-update
package sdk
