# \ResourcesAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResourcesGet**](ResourcesAPI.md#ResourcesGet) | **Get** /resources | Get resources
[**ResourcesPost**](ResourcesAPI.md#ResourcesPost) | **Post** /resources | Create a resource
[**ResourcesResourceIdDelete**](ResourcesAPI.md#ResourcesResourceIdDelete) | **Delete** /resources/{resourceId} | Delete resource
[**ResourcesResourceIdGet**](ResourcesAPI.md#ResourcesResourceIdGet) | **Get** /resources/{resourceId} | Get resource by id
[**ResourcesResourceIdPatch**](ResourcesAPI.md#ResourcesResourceIdPatch) | **Patch** /resources/{resourceId} | Update a resource



## ResourcesGet

> ResourcesGet200Response ResourcesGet(ctx).ResourceType(resourceType).Limit(limit).NextToken(nextToken).Execute()

Get resources



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
	resourceType := "resourceType_example" // string | The type of resource. (optional)
	limit := int32(50) // int32 | Maximum number of items to return per page. If not provided, all items are returned. When provided, enables pagination and the response will include next_token for retrieving subsequent pages. Valid range: 1-100. (optional)
	nextToken := "eyJsYXN0SWQiOjEyM30..." // string | Cursor token for retrieving the next page of results. Obtained from the next_token field of a previous paginated response. Only valid when the limit parameter is also provided. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.ResourcesGet(context.Background()).ResourceType(resourceType).Limit(limit).NextToken(nextToken).Execute()
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
 **resourceType** | **string** | The type of resource. | 
 **limit** | **int32** | Maximum number of items to return per page. If not provided, all items are returned. When provided, enables pagination and the response will include next_token for retrieving subsequent pages. Valid range: 1-100. | 
 **nextToken** | **string** | Cursor token for retrieving the next page of results. Obtained from the next_token field of a previous paginated response. Only valid when the limit parameter is also provided. | 

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
	resourcesPostRequest := *openapiclient.NewResourcesPostRequest("Type_example", "DisplayName_example", *openapiclient.NewResourcesPostRequestOptions(*openapiclient.NewSnowflakeOptionsDatabaseOptions("Name_example"), "AccountIdentifier_example", *openapiclient.NewGRPCOptionsAuthenticationOptions("AuthenticationType_example", "Auth0Domain_example", "Auth0ClientId_example", "Auth0ClientSecret_example", "Auth0CustomAudience_example", "Oauth2AccessTokenUrl_example", "Oauth2AuthUrl_example", "Oauth2ClientId_example", "Oauth2ClientSecret_example"), "BaseUrl_example")) // ResourcesPostRequest |  (optional)

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

Delete resource



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
	resourceId := "resourceId_example" // string | The uuid or name for the resource.
	deleteUnderlyingDB := true // bool | Whether to delete the underlying retool database when deleting the resource. (optional)

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
**resourceId** | **string** | The uuid or name for the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourcesResourceIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteUnderlyingDB** | **bool** | Whether to delete the underlying retool database when deleting the resource. | 

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

Get resource by id



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
	resourceId := "resourceId_example" // string | The uuid or name for the resource.

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
**resourceId** | **string** | The uuid or name for the resource. | 

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
	resourceId := "resourceId_example" // string | The uuid or name for the resource.
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
**resourceId** | **string** | The uuid or name for the resource. | 

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

