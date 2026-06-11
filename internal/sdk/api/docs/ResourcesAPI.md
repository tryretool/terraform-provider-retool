# \ResourcesAPI

All URIs are relative to *https://stable-4-0.retool.dev/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResourcesGet**](ResourcesAPI.md#ResourcesGet) | **Get** /resources | List resources
[**ResourcesPost**](ResourcesAPI.md#ResourcesPost) | **Post** /resources | Create a resource
[**ResourcesResourceIdDelete**](ResourcesAPI.md#ResourcesResourceIdDelete) | **Delete** /resources/{resourceId} | Delete a resource
[**ResourcesResourceIdGet**](ResourcesAPI.md#ResourcesResourceIdGet) | **Get** /resources/{resourceId} | Get a resource
[**ResourcesResourceIdPatch**](ResourcesAPI.md#ResourcesResourceIdPatch) | **Patch** /resources/{resourceId} | Update a resource



## ResourcesGet

> ResourcesGet200Response ResourcesGet(ctx).ResourceType(resourceType).NameContains(nameContains).Limit(limit).NextToken(nextToken).Execute()

List resources



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
	resourceType := "resourceType_example" // string |  (optional)
	nameContains := "sales" // string |  (optional)
	limit := int32(50) // int32 |  (optional)
	nextToken := "eyJsYXN0SWQiOjEyM30..." // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.ResourcesGet(context.Background()).ResourceType(resourceType).NameContains(nameContains).Limit(limit).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.ResourcesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourcesGet`: ResourcesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.ResourcesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResourcesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceType** | **string** |  | 
 **nameContains** | **string** |  | 
 **limit** | **int32** |  | 
 **nextToken** | **string** |  | 

### Return type

[**ResourcesGet200Response**](ResourcesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourcesPost

> ResourcesPost200Response ResourcesPost(ctx).ResourcesPostRequest(resourcesPostRequest).Execute()

Create a resource



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
	resourcesPostRequest := *openapiclient.NewResourcesPostRequest("Type_example", "DisplayName_example", map[string]interface{}{"key": interface{}(123)}) // ResourcesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.ResourcesPost(context.Background()).ResourcesPostRequest(resourcesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.ResourcesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourcesPost`: ResourcesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.ResourcesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResourcesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourcesPostRequest** | [**ResourcesPostRequest**](ResourcesPostRequest.md) |  | 

### Return type

[**ResourcesPost200Response**](ResourcesPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourcesResourceIdDelete

> ResourcesResourceIdDelete(ctx, resourceId).DeleteUnderlyingDB(deleteUnderlyingDB).Execute()

Delete a resource



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
	resourceId := "resourceId_example" // string | 
	deleteUnderlyingDB := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcesAPI.ResourcesResourceIdDelete(context.Background(), resourceId).DeleteUnderlyingDB(deleteUnderlyingDB).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.ResourcesResourceIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourcesResourceIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteUnderlyingDB** | **bool** |  | 

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


## ResourcesResourceIdGet

> ResourcesResourceIdGet200Response ResourcesResourceIdGet(ctx, resourceId).Execute()

Get a resource



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
	resourceId := "resourceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.ResourcesResourceIdGet(context.Background(), resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.ResourcesResourceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourcesResourceIdGet`: ResourcesResourceIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.ResourcesResourceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourcesResourceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourcesResourceIdGet200Response**](ResourcesResourceIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourcesResourceIdPatch

> ResourcesResourceIdPatch200Response ResourcesResourceIdPatch(ctx, resourceId).ResourcesResourceIdPatchRequest(resourcesResourceIdPatchRequest).Execute()

Update a resource



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
	resourceId := "resourceId_example" // string | 
	resourcesResourceIdPatchRequest := *openapiclient.NewResourcesResourceIdPatchRequest([]openapiclient.ReplaceOperation{*openapiclient.NewReplaceOperation("Op_example", "Path_example")}) // ResourcesResourceIdPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.ResourcesResourceIdPatch(context.Background(), resourceId).ResourcesResourceIdPatchRequest(resourcesResourceIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPI.ResourcesResourceIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourcesResourceIdPatch`: ResourcesResourceIdPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourcesAPI.ResourcesResourceIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourcesResourceIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourcesResourceIdPatchRequest** | [**ResourcesResourceIdPatchRequest**](ResourcesResourceIdPatchRequest.md) |  | 

### Return type

[**ResourcesResourceIdPatch200Response**](ResourcesResourceIdPatch200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

