# \FoldersAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FoldersFolderIdDelete**](FoldersAPI.md#FoldersFolderIdDelete) | **Delete** /folders/{folderId} | Delete folder
[**FoldersFolderIdGet**](FoldersAPI.md#FoldersFolderIdGet) | **Get** /folders/{folderId} | Get a folder
[**FoldersGet**](FoldersAPI.md#FoldersGet) | **Get** /folders | List folders
[**FoldersPost**](FoldersAPI.md#FoldersPost) | **Post** /folders | Create folder



## FoldersFolderIdDelete

> FoldersFolderIdDelete(ctx, folderId).Execute()

Delete folder



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
	folderId := "folderId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FoldersAPI.FoldersFolderIdDelete(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.FoldersFolderIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFoldersFolderIdDeleteRequest struct via the builder pattern


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


## FoldersFolderIdGet

> FoldersFolderIdGet200Response FoldersFolderIdGet(ctx, folderId).Execute()

Get a folder



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
	folderId := "folderId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.FoldersFolderIdGet(context.Background(), folderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.FoldersFolderIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FoldersFolderIdGet`: FoldersFolderIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.FoldersFolderIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFoldersFolderIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FoldersFolderIdGet200Response**](FoldersFolderIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FoldersGet

> FoldersGet200Response FoldersGet(ctx).Execute()

List folders



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
	resp, r, err := apiClient.FoldersAPI.FoldersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.FoldersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FoldersGet`: FoldersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.FoldersGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFoldersGetRequest struct via the builder pattern


### Return type

[**FoldersGet200Response**](FoldersGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FoldersPost

> FoldersPost200Response FoldersPost(ctx).FoldersPostRequest(foldersPostRequest).Execute()

Create folder



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
	foldersPostRequest := *openapiclient.NewFoldersPostRequest("Name_example", "FolderType_example") // FoldersPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FoldersAPI.FoldersPost(context.Background()).FoldersPostRequest(foldersPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FoldersAPI.FoldersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FoldersPost`: FoldersPost200Response
	fmt.Fprintf(os.Stdout, "Response from `FoldersAPI.FoldersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFoldersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **foldersPostRequest** | [**FoldersPostRequest**](FoldersPostRequest.md) |  | 

### Return type

[**FoldersPost200Response**](FoldersPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

