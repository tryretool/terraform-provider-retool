# \ResourcesAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResourcesGet**](ResourcesAPI.md#ResourcesGet) | **Get** /resources | Get resources
[**ResourcesResourceIdDelete**](ResourcesAPI.md#ResourcesResourceIdDelete) | **Delete** /resources/{resourceId} | Delete resource
[**ResourcesResourceIdGet**](ResourcesAPI.md#ResourcesResourceIdGet) | **Get** /resources/{resourceId} | Get resource by id



## ResourcesGet

> ResourcesGet200Response ResourcesGet(ctx).NextToken(nextToken).ResourceType(resourceType).Execute()

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
	nextToken := "nextToken_example" // string |  (optional)
	resourceType := "resourceType_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourcesAPI.ResourcesGet(context.Background()).NextToken(nextToken).ResourceType(resourceType).Execute()
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
 **nextToken** | **string** |  | 
 **resourceType** | **string** |  | 

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


## ResourcesResourceIdDelete

> ResourcesResourceIdDelete(ctx, resourceId).Execute()

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
	resourceId := *openapiclient.NewResourcesGet200ResponseDataInnerId() // ResourcesGet200ResponseDataInnerId | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcesAPI.ResourcesResourceIdDelete(context.Background(), resourceId).Execute()
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
**resourceId** | [**ResourcesGet200ResponseDataInnerId**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourcesResourceIdDeleteRequest struct via the builder pattern


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
	resourceId := *openapiclient.NewResourcesGet200ResponseDataInnerId() // ResourcesGet200ResponseDataInnerId | 

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
**resourceId** | [**ResourcesGet200ResponseDataInnerId**](.md) |  | 

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

