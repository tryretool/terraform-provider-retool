# \EnvironmentAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsEnvironmentIdGet**](EnvironmentAPI.md#EnvironmentsEnvironmentIdGet) | **Get** /environments/{environmentId} | Get an environment by id
[**EnvironmentsGet**](EnvironmentAPI.md#EnvironmentsGet) | **Get** /environments | Get organization environments



## EnvironmentsEnvironmentIdGet

> EnvironmentsEnvironmentIdGet200Response EnvironmentsEnvironmentIdGet(ctx, environmentId).Execute()

Get an environment by id



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
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentAPI.EnvironmentsEnvironmentIdGet(context.Background(), environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentsEnvironmentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsEnvironmentIdGet`: EnvironmentsEnvironmentIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentsEnvironmentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentsEnvironmentIdGet200Response**](EnvironmentsEnvironmentIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsGet

> EnvironmentsGet200Response EnvironmentsGet(ctx).Execute()

Get organization environments



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
	resp, r, err := apiClient.EnvironmentAPI.EnvironmentsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentAPI.EnvironmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsGet`: EnvironmentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentAPI.EnvironmentsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsGetRequest struct via the builder pattern


### Return type

[**EnvironmentsGet200Response**](EnvironmentsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

