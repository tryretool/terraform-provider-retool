# Go client library for Retool API
The library code is generated from OpenAPI spec using [OpenAPI Generator](https://openapi-generator.tech/).

## Install
You can install OpenAPI Generator using Brew:
```
brew install openapi-generator
```

NOTE: You might have to install newer version of JDK that we have by default on our laptops. I downloaded one from https://adoptium.net/ and set my JAVA_HOME env var to `/Library/Java/JavaVirtualMachines/temurin-21.jdk/Contents/Home`

## Prepare the OpenAPI spec
If you want to update the client code to cover recent changes in the Retool API, you'd need to update `openAPISpec.json` file first:
- download new spec from e.g. https://api.retool.com/api/v2/spec
- format it so that it's human-readable ("Format Document" command in VSCode would do)
- manually make sure that every `tags` element has only one value. 
  
Why do we have to manually edit `tags`? Otherwise OpenAPI Generator will generate the same request/response structure for every tag value and you'll get compile errors. 
Leave the tag that corresponds to the top-level element in the endpoint path. So, if the tags are `["Users", "User Attributes"]` for `/users/{userId}/user_attributes/{attributeName}`, you should only leave `["Users"]`.
See https://github.com/OpenAPITools/openapi-generator/issues/741 for more details and potential alternatives that we may explore in the future.

## Generate client library code
Just run `go generate` in this folder. All generated code lives in the `api` folder.
Command line parameters for `openapi-generator` invocation are controlled via `generate.go` file.

## Test client library
There's a simple executable in `client` folder that makes an API request to localhost:3000 and prints out the response.
You can run it as follows:
```
RETOOL_ACCESS_TOKEN=<your API token generated on localhost:3000> go run client/main.go
```
