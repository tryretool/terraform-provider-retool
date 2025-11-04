# \EnvironmentsAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentsEnvironmentIdDelete**](EnvironmentsAPI.md#EnvironmentsEnvironmentIdDelete) | **Delete** /environments/{environmentId} | Delete environment
[**EnvironmentsEnvironmentIdGet**](EnvironmentsAPI.md#EnvironmentsEnvironmentIdGet) | **Get** /environments/{environmentId} | Get an environment
[**EnvironmentsEnvironmentIdPatch**](EnvironmentsAPI.md#EnvironmentsEnvironmentIdPatch) | **Patch** /environments/{environmentId} | Update an environment
[**EnvironmentsGet**](EnvironmentsAPI.md#EnvironmentsGet) | **Get** /environments | Get environments
[**EnvironmentsPost**](EnvironmentsAPI.md#EnvironmentsPost) | **Post** /environments | Create environment



## EnvironmentsEnvironmentIdDelete

> EnvironmentsEnvironmentIdDelete(ctx, environmentId).Execute()

Delete environment



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
	r, err := apiClient.EnvironmentsAPI.EnvironmentsEnvironmentIdDelete(context.Background(), environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.EnvironmentsEnvironmentIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIdDeleteRequest struct via the builder pattern


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


## EnvironmentsEnvironmentIdGet

> EnvironmentsEnvironmentIdGet200Response EnvironmentsEnvironmentIdGet(ctx, environmentId).Execute()

Get an environment



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
	resp, r, err := apiClient.EnvironmentsAPI.EnvironmentsEnvironmentIdGet(context.Background(), environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.EnvironmentsEnvironmentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsEnvironmentIdGet`: EnvironmentsEnvironmentIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.EnvironmentsEnvironmentIdGet`: %v\n", resp)
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


## EnvironmentsEnvironmentIdPatch

> EnvironmentsEnvironmentIdPatch200Response EnvironmentsEnvironmentIdPatch(ctx, environmentId).EnvironmentsEnvironmentIdPatchRequest(environmentsEnvironmentIdPatchRequest).Execute()

Update an environment



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
	environmentsEnvironmentIdPatchRequest := *openapiclient.NewEnvironmentsEnvironmentIdPatchRequest([]openapiclient.ReplaceOperation{*openapiclient.NewReplaceOperation("Op_example", "Path_example")}) // EnvironmentsEnvironmentIdPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsAPI.EnvironmentsEnvironmentIdPatch(context.Background(), environmentId).EnvironmentsEnvironmentIdPatchRequest(environmentsEnvironmentIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.EnvironmentsEnvironmentIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsEnvironmentIdPatch`: EnvironmentsEnvironmentIdPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.EnvironmentsEnvironmentIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsEnvironmentIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentsEnvironmentIdPatchRequest** | [**EnvironmentsEnvironmentIdPatchRequest**](EnvironmentsEnvironmentIdPatchRequest.md) |  | 

### Return type

[**EnvironmentsEnvironmentIdPatch200Response**](EnvironmentsEnvironmentIdPatch200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsGet

> EnvironmentsGet200Response EnvironmentsGet(ctx).Execute()

Get environments



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
	resp, r, err := apiClient.EnvironmentsAPI.EnvironmentsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.EnvironmentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsGet`: EnvironmentsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.EnvironmentsGet`: %v\n", resp)
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


## EnvironmentsPost

> EnvironmentsPost200Response EnvironmentsPost(ctx).EnvironmentsPostRequest(environmentsPostRequest).Execute()

Create environment



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
	environmentsPostRequest := *openapiclient.NewEnvironmentsPostRequest("Name_example", "#FFFFFF") // EnvironmentsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsAPI.EnvironmentsPost(context.Background()).EnvironmentsPostRequest(environmentsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsAPI.EnvironmentsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnvironmentsPost`: EnvironmentsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsAPI.EnvironmentsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentsPostRequest** | [**EnvironmentsPostRequest**](EnvironmentsPostRequest.md) |  | 

### Return type

[**EnvironmentsPost200Response**](EnvironmentsPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

