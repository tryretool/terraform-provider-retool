# Go client library for Retool API
The library code is generated from OpenAPI spec using [OpenAPI Generator](https://openapi-generator.tech/).

## Install
Install OpenAPI Generator [launcher script](https://openapi-generator.tech/docs/installation#bash-launcher-script)
```
brew install mvn

mkdir -p ~/bin/openapitools
curl https://raw.githubusercontent.com/OpenAPITools/openapi-generator/master/bin/utils/openapi-generator-cli.sh > ~/bin/openapitools/openapi-generator-cli
chmod u+x ~/bin/openapitools/openapi-generator-cli
export PATH=$PATH:~/bin/openapitools/
```

You'd want to update your shell file (.bashrc or equivalent) so that it adds `~/bin/openapitools/` to your `PATH` .

NOTE: You might have to install newer version of JDK that we have by default on our laptops. I downloaded one from https://adoptium.net/ and set my JAVA_HOME env var to `/Library/Java/JavaVirtualMachines/temurin-21.jdk/Contents/Home`

## Prepare the OpenAPI spec
If you want to update the client code to cover recent changes in the Retool API, you'd need to update `openAPISpec.json` file first:
- download new spec from e.g. https://api.retool.com/api/v2/spec
- format it so that it's human-readable ("Format Document" command in VSCode would do)
- manually make sure that every `tags` element has only one value. 
  
Why do we have to manually edit `tags`? Otherwise OpenAPI Generator will generate the same request/response structure for every tag value and you'll get compile errors. 
Leave the tag that corresponds to the top-level element in the endpoint path. So, if the tags are `["Users", "User Attributes"]` for `/users/{userId}/user_attributes/{attributeName}`, you should only leave `["Users"]`.
See https://github.com/OpenAPITools/openapi-generator/issues/741 for more details and potential alternatives that we may explore in the future.

NOTE: To make the client work with our API, I had to manually update `openAPISpec.json` in quite a few places. So, I'd recommend manually copying the fragment of the spec that you need into the file instead of trying to replace the file completely.

One of these is for the imprecise conversion of zod discriminated unions (used internally) to openapi - instead of `oneOf` schemas, `anyOf` is more accurate for nested object types.

## Generate client library code
Just run `go generate` in this folder. All generated code lives in the `api` folder.
Command line parameters for `openapi-generator` invocation are controlled via `generate.go` file.

## Test client library
There's a simple executable in `client` folder that makes an API request to localhost:3000 and prints out the response.
You can run it as follows:
```
RETOOL_ACCESS_TOKEN=<your API token generated on localhost:3000> go run client/main.go
```


## Notes from updating to 2.9.0

LLMs have sped this up a fair bit. With Sonnet 4.5 in cursor I was able to add a copy of the original spec this was based on (2.4.0), point it at the modified spec in this repo, and add the spec we were targeting (2.9.0). The following propmt yielded really good results:
```
@openAPISpecv2.4.0.json was updated to @openAPISpec.json to support terraform client autogeneration as described in @README.md 

Understand those changes, and attempt to update autogeneration to v2.9.0 as described in @openApiSpecv2.9.json 

You can understand what commands to run to build the library, and get more context on how the original file was modified in the README file
```

Following that a cursor agent was able to make remarkable progress on updating the provider code to respect this, including tests.