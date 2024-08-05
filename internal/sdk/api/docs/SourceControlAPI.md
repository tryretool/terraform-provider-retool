# \SourceControlAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SourceControlConfigDelete**](SourceControlAPI.md#SourceControlConfigDelete) | **Delete** /source_control/config | Delete source control provider configuration
[**SourceControlConfigGet**](SourceControlAPI.md#SourceControlConfigGet) | **Get** /source_control/config | Get source control configuration
[**SourceControlConfigPost**](SourceControlAPI.md#SourceControlConfigPost) | **Post** /source_control/config | Create source control configuration
[**SourceControlConfigPut**](SourceControlAPI.md#SourceControlConfigPut) | **Put** /source_control/config | Set source control configuration
[**SourceControlDeployPost**](SourceControlAPI.md#SourceControlDeployPost) | **Post** /source_control/deploy | Trigger deployment of latest changes
[**SourceControlDeploymentIdGet**](SourceControlAPI.md#SourceControlDeploymentIdGet) | **Get** /source_control/deployment/{id} | Get deployment by id
[**SourceControlSettingsGet**](SourceControlAPI.md#SourceControlSettingsGet) | **Get** /source_control/settings | Get source control settings
[**SourceControlSettingsPut**](SourceControlAPI.md#SourceControlSettingsPut) | **Put** /source_control/settings | Set source control settings
[**SourceControlTestConnectionGet**](SourceControlAPI.md#SourceControlTestConnectionGet) | **Get** /source_control/test_connection | Tests source control connection
[**SourceControlTestDeployPost**](SourceControlAPI.md#SourceControlTestDeployPost) | **Post** /source_control/test_deploy | Test source control changes



## SourceControlConfigDelete

> SourceControlConfigDelete(ctx).Execute()

Delete source control provider configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SourceControlAPI.SourceControlConfigDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlConfigDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlConfigDeleteRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlConfigGet

> SourceControlConfigGet200Response SourceControlConfigGet(ctx).Execute()

Get source control configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlConfigGet`: SourceControlConfigGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlConfigGetRequest struct via the builder pattern


### Return type

[**SourceControlConfigGet200Response**](SourceControlConfigGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlConfigPost

> SourceControlConfigPost200Response SourceControlConfigPost(ctx).SourceControlConfigPutRequest(sourceControlConfigPutRequest).Execute()

Create source control configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sourceControlConfigPutRequest := *openapiclient.NewSourceControlConfigPutRequest(openapiclient._source_control_config_put_request_config{SourceControlConfigGet200ResponseDataOneOf: openapiclient.NewSourceControlConfigGet200ResponseDataOneOf(openapiclient._source_control_config_get_200_response_data_oneOf_config{SourceControlConfigGet200ResponseDataOneOfConfigOneOf: openapiclient.NewSourceControlConfigGet200ResponseDataOneOfConfigOneOf("Type_example", "AppId_example", "InstallationId_example", "PrivateKey_example")}, "Provider_example", "Org_example", "Repo_example", "DefaultBranch_example")}) // SourceControlConfigPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlConfigPost(context.Background()).SourceControlConfigPutRequest(sourceControlConfigPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlConfigPost`: SourceControlConfigPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceControlConfigPutRequest** | [**SourceControlConfigPutRequest**](SourceControlConfigPutRequest.md) |  | 

### Return type

[**SourceControlConfigPost200Response**](SourceControlConfigPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlConfigPut

> SourceControlConfigPut200Response SourceControlConfigPut(ctx).SourceControlConfigPutRequest(sourceControlConfigPutRequest).Execute()

Set source control configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sourceControlConfigPutRequest := *openapiclient.NewSourceControlConfigPutRequest(openapiclient._source_control_config_put_request_config{SourceControlConfigGet200ResponseDataOneOf: openapiclient.NewSourceControlConfigGet200ResponseDataOneOf(openapiclient._source_control_config_get_200_response_data_oneOf_config{SourceControlConfigGet200ResponseDataOneOfConfigOneOf: openapiclient.NewSourceControlConfigGet200ResponseDataOneOfConfigOneOf("Type_example", "AppId_example", "InstallationId_example", "PrivateKey_example")}, "Provider_example", "Org_example", "Repo_example", "DefaultBranch_example")}) // SourceControlConfigPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlConfigPut(context.Background()).SourceControlConfigPutRequest(sourceControlConfigPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlConfigPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlConfigPut`: SourceControlConfigPut200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlConfigPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlConfigPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceControlConfigPutRequest** | [**SourceControlConfigPutRequest**](SourceControlConfigPutRequest.md) |  | 

### Return type

[**SourceControlConfigPut200Response**](SourceControlConfigPut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlDeployPost

> SourceControlDeployPost200Response SourceControlDeployPost(ctx).Execute()

Trigger deployment of latest changes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlDeployPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlDeployPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlDeployPost`: SourceControlDeployPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlDeployPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlDeployPostRequest struct via the builder pattern


### Return type

[**SourceControlDeployPost200Response**](SourceControlDeployPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlDeploymentIdGet

> SourceControlDeployPost200Response SourceControlDeploymentIdGet(ctx, id).Execute()

Get deployment by id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlDeploymentIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlDeploymentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlDeploymentIdGet`: SourceControlDeployPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlDeploymentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlDeploymentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceControlDeployPost200Response**](SourceControlDeployPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlSettingsGet

> SourceControlSettingsGet200Response SourceControlSettingsGet(ctx).Execute()

Get source control settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlSettingsGet`: SourceControlSettingsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlSettingsGetRequest struct via the builder pattern


### Return type

[**SourceControlSettingsGet200Response**](SourceControlSettingsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlSettingsPut

> SourceControlSettingsPut200Response SourceControlSettingsPut(ctx).SourceControlSettingsPutRequest(sourceControlSettingsPutRequest).Execute()

Set source control settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sourceControlSettingsPutRequest := *openapiclient.NewSourceControlSettingsPutRequest(*openapiclient.NewSourceControlSettingsPutRequestSettings()) // SourceControlSettingsPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlSettingsPut(context.Background()).SourceControlSettingsPutRequest(sourceControlSettingsPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlSettingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlSettingsPut`: SourceControlSettingsPut200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlSettingsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceControlSettingsPutRequest** | [**SourceControlSettingsPutRequest**](SourceControlSettingsPutRequest.md) |  | 

### Return type

[**SourceControlSettingsPut200Response**](SourceControlSettingsPut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlTestConnectionGet

> SourceControlTestConnectionGet200Response SourceControlTestConnectionGet(ctx).Execute()

Tests source control connection



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlTestConnectionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlTestConnectionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlTestConnectionGet`: SourceControlTestConnectionGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlTestConnectionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlTestConnectionGetRequest struct via the builder pattern


### Return type

[**SourceControlTestConnectionGet200Response**](SourceControlTestConnectionGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlTestDeployPost

> SourceControlTestDeployPost200Response SourceControlTestDeployPost(ctx).SourceControlTestDeployPostRequest(sourceControlTestDeployPostRequest).Execute()

Test source control changes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sourceControlTestDeployPostRequest := *openapiclient.NewSourceControlTestDeployPostRequest(*openapiclient.NewSourceControlTestDeployPostRequestDeployParams("CommitSha_example")) // SourceControlTestDeployPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlTestDeployPost(context.Background()).SourceControlTestDeployPostRequest(sourceControlTestDeployPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlTestDeployPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlTestDeployPost`: SourceControlTestDeployPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlTestDeployPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlTestDeployPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceControlTestDeployPostRequest** | [**SourceControlTestDeployPostRequest**](SourceControlTestDeployPostRequest.md) |  | 

### Return type

[**SourceControlTestDeployPost200Response**](SourceControlTestDeployPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

