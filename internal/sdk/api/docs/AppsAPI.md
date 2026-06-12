# \AppsAPI

All URIs are relative to *https://stable-4-0.retool.dev/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsAppIdDelete**](AppsAPI.md#AppsAppIdDelete) | **Delete** /apps/{appId} | Delete an app
[**AppsAppIdGet**](AppsAPI.md#AppsAppIdGet) | **Get** /apps/{appId} | Get an app
[**AppsCloneAppPost**](AppsAPI.md#AppsCloneAppPost) | **Post** /apps/cloneApp | Clone an app
[**AppsGet**](AppsAPI.md#AppsGet) | **Get** /apps | List apps
[**AppsToolscriptValidatePost**](AppsAPI.md#AppsToolscriptValidatePost) | **Post** /apps/toolscript/validate | Validate Toolscript



## AppsAppIdDelete

> AppsAppIdDelete(ctx, appId).Execute()

Delete an app



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
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppsAPI.AppsAppIdDelete(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AppsAppIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAppIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsAppIdGet

> AppsAppIdGet200Response AppsAppIdGet(ctx, appId).Execute()

Get an app



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
	appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.AppsAppIdGet(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AppsAppIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsAppIdGet`: AppsAppIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.AppsAppIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppsAppIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppsAppIdGet200Response**](AppsAppIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsCloneAppPost

> AppsAppIdGet200Response AppsCloneAppPost(ctx).AppsCloneAppPostRequest(appsCloneAppPostRequest).Execute()

Clone an app



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
	appsCloneAppPostRequest := *openapiclient.NewAppsCloneAppPostRequest("AppId_example", "NewAppName_example") // AppsCloneAppPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.AppsCloneAppPost(context.Background()).AppsCloneAppPostRequest(appsCloneAppPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AppsCloneAppPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsCloneAppPost`: AppsAppIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.AppsCloneAppPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsCloneAppPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appsCloneAppPostRequest** | [**AppsCloneAppPostRequest**](AppsCloneAppPostRequest.md) |  | 

### Return type

[**AppsAppIdGet200Response**](AppsAppIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsGet

> AppsGet200Response AppsGet(ctx).UsingResource(usingResource).NameContains(nameContains).Limit(limit).NextToken(nextToken).Execute()

List apps



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
	usingResource := *openapiclient.NewAppsGetUsingResourceParameter() // AppsGetUsingResourceParameter |  (optional)
	nameContains := "dashboard" // string |  (optional)
	limit := int32(50) // int32 |  (optional)
	nextToken := "eyJsYXN0SWQiOjEyM30..." // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.AppsGet(context.Background()).UsingResource(usingResource).NameContains(nameContains).Limit(limit).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AppsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsGet`: AppsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.AppsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usingResource** | [**AppsGetUsingResourceParameter**](AppsGetUsingResourceParameter.md) |  | 
 **nameContains** | **string** |  | 
 **limit** | **int32** |  | 
 **nextToken** | **string** |  | 

### Return type

[**AppsGet200Response**](AppsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppsToolscriptValidatePost

> AppsToolscriptValidatePost200Response AppsToolscriptValidatePost(ctx).AppsToolscriptValidatePostRequest(appsToolscriptValidatePostRequest).Execute()

Validate Toolscript



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
	appsToolscriptValidatePostRequest := *openapiclient.NewAppsToolscriptValidatePostRequest([]openapiclient.AppsToolscriptValidatePostRequestToolscriptInner{*openapiclient.NewAppsToolscriptValidatePostRequestToolscriptInner("Path_example", "Contents_example")}) // AppsToolscriptValidatePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppsAPI.AppsToolscriptValidatePost(context.Background()).AppsToolscriptValidatePostRequest(appsToolscriptValidatePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppsAPI.AppsToolscriptValidatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppsToolscriptValidatePost`: AppsToolscriptValidatePost200Response
	fmt.Fprintf(os.Stdout, "Response from `AppsAPI.AppsToolscriptValidatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppsToolscriptValidatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appsToolscriptValidatePostRequest** | [**AppsToolscriptValidatePostRequest**](AppsToolscriptValidatePostRequest.md) |  | 

### Return type

[**AppsToolscriptValidatePost200Response**](AppsToolscriptValidatePost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

