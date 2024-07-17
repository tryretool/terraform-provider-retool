# \SpacesAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SpacesCopyElementsPost**](SpacesAPI.md#SpacesCopyElementsPost) | **Post** /spaces/copyElements | Copies apps, queries, resources, and workflows from one space to another
[**SpacesGet**](SpacesAPI.md#SpacesGet) | **Get** /spaces | List spaces
[**SpacesPost**](SpacesAPI.md#SpacesPost) | **Post** /spaces | Create a space
[**SpacesSpaceIdDelete**](SpacesAPI.md#SpacesSpaceIdDelete) | **Delete** /spaces/{spaceId} | Delete a space
[**SpacesSpaceIdGet**](SpacesAPI.md#SpacesSpaceIdGet) | **Get** /spaces/{spaceId} | Get space
[**SpacesSpaceIdPut**](SpacesAPI.md#SpacesSpaceIdPut) | **Put** /spaces/{spaceId} | Update space



## SpacesCopyElementsPost

> SpacesCopyElementsPost201Response SpacesCopyElementsPost(ctx).SpacesCopyElementsPostRequest(spacesCopyElementsPostRequest).Execute()

Copies apps, queries, resources, and workflows from one space to another



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
	spacesCopyElementsPostRequest := *openapiclient.NewSpacesCopyElementsPostRequest([]openapiclient.SpacesCopyElementsPostRequestResourceIdsInner{*openapiclient.NewSpacesCopyElementsPostRequestResourceIdsInner()}, []string{"QueryLibraryQueryIds_example"}, []string{"AppIds_example"}, []string{"WorkflowIds_example"}, "DestinationSpaceId_example") // SpacesCopyElementsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.SpacesCopyElementsPost(context.Background()).SpacesCopyElementsPostRequest(spacesCopyElementsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.SpacesCopyElementsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpacesCopyElementsPost`: SpacesCopyElementsPost201Response
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.SpacesCopyElementsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpacesCopyElementsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spacesCopyElementsPostRequest** | [**SpacesCopyElementsPostRequest**](SpacesCopyElementsPostRequest.md) |  | 

### Return type

[**SpacesCopyElementsPost201Response**](SpacesCopyElementsPost201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpacesGet

> SpacesGet200Response SpacesGet(ctx).Execute()

List spaces



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
	resp, r, err := apiClient.SpacesAPI.SpacesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.SpacesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpacesGet`: SpacesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.SpacesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSpacesGetRequest struct via the builder pattern


### Return type

[**SpacesGet200Response**](SpacesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpacesPost

> SpacesPost200Response SpacesPost(ctx).SpacesPostRequest(spacesPostRequest).Execute()

Create a space



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
	spacesPostRequest := *openapiclient.NewSpacesPostRequest("Name_example", "Domain_example") // SpacesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.SpacesPost(context.Background()).SpacesPostRequest(spacesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.SpacesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpacesPost`: SpacesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.SpacesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSpacesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **spacesPostRequest** | [**SpacesPostRequest**](SpacesPostRequest.md) |  | 

### Return type

[**SpacesPost200Response**](SpacesPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpacesSpaceIdDelete

> SpacesSpaceIdDelete(ctx, spaceId).Execute()

Delete a space



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
	spaceId := "spaceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SpacesAPI.SpacesSpaceIdDelete(context.Background(), spaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.SpacesSpaceIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpacesSpaceIdDeleteRequest struct via the builder pattern


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


## SpacesSpaceIdGet

> SpacesSpaceIdGet200Response SpacesSpaceIdGet(ctx, spaceId).Execute()

Get space



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
	spaceId := "spaceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.SpacesSpaceIdGet(context.Background(), spaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.SpacesSpaceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpacesSpaceIdGet`: SpacesSpaceIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.SpacesSpaceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpacesSpaceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpacesSpaceIdGet200Response**](SpacesSpaceIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SpacesSpaceIdPut

> SpacesSpaceIdPut200Response SpacesSpaceIdPut(ctx, spaceId).Execute()

Update space



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
	spaceId := "spaceId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SpacesAPI.SpacesSpaceIdPut(context.Background(), spaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SpacesAPI.SpacesSpaceIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SpacesSpaceIdPut`: SpacesSpaceIdPut200Response
	fmt.Fprintf(os.Stdout, "Response from `SpacesAPI.SpacesSpaceIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSpacesSpaceIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SpacesSpaceIdPut200Response**](SpacesSpaceIdPut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

