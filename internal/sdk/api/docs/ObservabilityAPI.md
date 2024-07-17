# \ObservabilityAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ObservabilityConfigConfigIdDelete**](ObservabilityAPI.md#ObservabilityConfigConfigIdDelete) | **Delete** /observability/config/{configId} | Delete an observability provider configuration
[**ObservabilityConfigConfigIdPut**](ObservabilityAPI.md#ObservabilityConfigConfigIdPut) | **Put** /observability/config/{configId} | Update an observability provider configuration
[**ObservabilityConfigGet**](ObservabilityAPI.md#ObservabilityConfigGet) | **Get** /observability/config | Get observability provider configurations
[**ObservabilityConfigPost**](ObservabilityAPI.md#ObservabilityConfigPost) | **Post** /observability/config | Create a new observability provider configuration
[**ObservabilityProviderProviderTestPost**](ObservabilityAPI.md#ObservabilityProviderProviderTestPost) | **Post** /observability/provider/{provider}/test | Send a test error event to the observability provider



## ObservabilityConfigConfigIdDelete

> ObservabilityConfigConfigIdDelete(ctx, configId).Execute()

Delete an observability provider configuration



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
	configId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObservabilityAPI.ObservabilityConfigConfigIdDelete(context.Background(), configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityAPI.ObservabilityConfigConfigIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiObservabilityConfigConfigIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ObservabilityConfigConfigIdPut

> ObservabilityConfigConfigIdPut200Response ObservabilityConfigConfigIdPut(ctx, configId).ObservabilityConfigConfigIdPutRequest(observabilityConfigConfigIdPutRequest).Execute()

Update an observability provider configuration



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
	configId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	observabilityConfigConfigIdPutRequest := *openapiclient.NewObservabilityConfigConfigIdPutRequest() // ObservabilityConfigConfigIdPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservabilityAPI.ObservabilityConfigConfigIdPut(context.Background(), configId).ObservabilityConfigConfigIdPutRequest(observabilityConfigConfigIdPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityAPI.ObservabilityConfigConfigIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObservabilityConfigConfigIdPut`: ObservabilityConfigConfigIdPut200Response
	fmt.Fprintf(os.Stdout, "Response from `ObservabilityAPI.ObservabilityConfigConfigIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiObservabilityConfigConfigIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **observabilityConfigConfigIdPutRequest** | [**ObservabilityConfigConfigIdPutRequest**](ObservabilityConfigConfigIdPutRequest.md) |  | 

### Return type

[**ObservabilityConfigConfigIdPut200Response**](ObservabilityConfigConfigIdPut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObservabilityConfigGet

> ObservabilityConfigGet200Response ObservabilityConfigGet(ctx).Execute()

Get observability provider configurations



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
	resp, r, err := apiClient.ObservabilityAPI.ObservabilityConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityAPI.ObservabilityConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObservabilityConfigGet`: ObservabilityConfigGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ObservabilityAPI.ObservabilityConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiObservabilityConfigGetRequest struct via the builder pattern


### Return type

[**ObservabilityConfigGet200Response**](ObservabilityConfigGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObservabilityConfigPost

> ObservabilityConfigPost200Response ObservabilityConfigPost(ctx).ObservabilityConfigPostRequest(observabilityConfigPostRequest).Execute()

Create a new observability provider configuration



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
	observabilityConfigPostRequest := *openapiclient.NewObservabilityConfigPostRequest(openapiclient._observability_config_get_200_response_data_inner_config{ObservabilityConfigGet200ResponseDataInnerConfigOneOf: openapiclient.NewObservabilityConfigGet200ResponseDataInnerConfigOneOf("Provider_example", "Dsn_example")}, false) // ObservabilityConfigPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservabilityAPI.ObservabilityConfigPost(context.Background()).ObservabilityConfigPostRequest(observabilityConfigPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityAPI.ObservabilityConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObservabilityConfigPost`: ObservabilityConfigPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ObservabilityAPI.ObservabilityConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiObservabilityConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **observabilityConfigPostRequest** | [**ObservabilityConfigPostRequest**](ObservabilityConfigPostRequest.md) |  | 

### Return type

[**ObservabilityConfigPost200Response**](ObservabilityConfigPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObservabilityProviderProviderTestPost

> ObservabilityProviderProviderTestPost200Response ObservabilityProviderProviderTestPost(ctx, provider).Execute()

Send a test error event to the observability provider



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
	provider := *openapiclient.NewObservabilityProviderProviderTestPostProviderParameter() // ObservabilityProviderProviderTestPostProviderParameter | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObservabilityAPI.ObservabilityProviderProviderTestPost(context.Background(), provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObservabilityAPI.ObservabilityProviderProviderTestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ObservabilityProviderProviderTestPost`: ObservabilityProviderProviderTestPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ObservabilityAPI.ObservabilityProviderProviderTestPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | [**ObservabilityProviderProviderTestPostProviderParameter**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiObservabilityProviderProviderTestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObservabilityProviderProviderTestPost200Response**](ObservabilityProviderProviderTestPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

