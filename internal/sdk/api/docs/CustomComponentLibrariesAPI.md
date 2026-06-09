# \CustomComponentLibrariesAPI

All URIs are relative to *https://stable-4-0.retool.dev/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomComponentLibrariesGet**](CustomComponentLibrariesAPI.md#CustomComponentLibrariesGet) | **Get** /custom_component_libraries | List custom component libraries
[**CustomComponentLibrariesLibraryIdGet**](CustomComponentLibrariesAPI.md#CustomComponentLibrariesLibraryIdGet) | **Get** /custom_component_libraries/{libraryId} | Get a custom component library
[**CustomComponentLibrariesLibraryIdRevisionsGet**](CustomComponentLibrariesAPI.md#CustomComponentLibrariesLibraryIdRevisionsGet) | **Get** /custom_component_libraries/{libraryId}/revisions | List revisions of a custom component library
[**CustomComponentLibrariesLibraryIdRevisionsPost**](CustomComponentLibrariesAPI.md#CustomComponentLibrariesLibraryIdRevisionsPost) | **Post** /custom_component_libraries/{libraryId}/revisions | Create a custom component library revision
[**CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet**](CustomComponentLibrariesAPI.md#CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet) | **Get** /custom_component_libraries/{libraryId}/revisions/{revisionId}/files | List files for a custom component library revision
[**CustomComponentLibrariesPost**](CustomComponentLibrariesAPI.md#CustomComponentLibrariesPost) | **Post** /custom_component_libraries | Create a custom component library



## CustomComponentLibrariesGet

> CustomComponentLibrariesGet200Response CustomComponentLibrariesGet(ctx).LibraryId(libraryId).Execute()

List custom component libraries



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
	libraryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomComponentLibrariesAPI.CustomComponentLibrariesGet(context.Background()).LibraryId(libraryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomComponentLibrariesAPI.CustomComponentLibrariesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomComponentLibrariesGet`: CustomComponentLibrariesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomComponentLibrariesAPI.CustomComponentLibrariesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomComponentLibrariesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **libraryId** | **string** |  | 

### Return type

[**CustomComponentLibrariesGet200Response**](CustomComponentLibrariesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomComponentLibrariesLibraryIdGet

> CustomComponentLibrariesLibraryIdGet200Response CustomComponentLibrariesLibraryIdGet(ctx, libraryId).Execute()

Get a custom component library



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
	libraryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomComponentLibrariesAPI.CustomComponentLibrariesLibraryIdGet(context.Background(), libraryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomComponentLibrariesAPI.CustomComponentLibrariesLibraryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomComponentLibrariesLibraryIdGet`: CustomComponentLibrariesLibraryIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomComponentLibrariesAPI.CustomComponentLibrariesLibraryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomComponentLibrariesLibraryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomComponentLibrariesLibraryIdGet200Response**](CustomComponentLibrariesLibraryIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomComponentLibrariesLibraryIdRevisionsGet

> CustomComponentLibrariesLibraryIdRevisionsGet200Response CustomComponentLibrariesLibraryIdRevisionsGet(ctx, libraryId).Execute()

List revisions of a custom component library



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
	libraryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomComponentLibrariesAPI.CustomComponentLibrariesLibraryIdRevisionsGet(context.Background(), libraryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomComponentLibrariesAPI.CustomComponentLibrariesLibraryIdRevisionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomComponentLibrariesLibraryIdRevisionsGet`: CustomComponentLibrariesLibraryIdRevisionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomComponentLibrariesAPI.CustomComponentLibrariesLibraryIdRevisionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomComponentLibrariesLibraryIdRevisionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomComponentLibrariesLibraryIdRevisionsGet200Response**](CustomComponentLibrariesLibraryIdRevisionsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomComponentLibrariesLibraryIdRevisionsPost

> CustomComponentLibrariesLibraryIdRevisionsPost200Response CustomComponentLibrariesLibraryIdRevisionsPost(ctx, libraryId).VersionBump(versionBump).Files(files).Id(id).Version(version).Execute()

Create a custom component library revision



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
	libraryId := "libraryId_example" // string | 
	versionBump := "versionBump_example" // string | 
	files := os.NewFile(1234, "some_file") // *os.File | 
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Specifies a specific id to use for the library. Used for syncronizing libraries across Retool Instances. (optional)
	version := "version_example" // string | A specific version tag to use. Also specify version_bump as 'specify_version'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomComponentLibrariesAPI.CustomComponentLibrariesLibraryIdRevisionsPost(context.Background(), libraryId).VersionBump(versionBump).Files(files).Id(id).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomComponentLibrariesAPI.CustomComponentLibrariesLibraryIdRevisionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomComponentLibrariesLibraryIdRevisionsPost`: CustomComponentLibrariesLibraryIdRevisionsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomComponentLibrariesAPI.CustomComponentLibrariesLibraryIdRevisionsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomComponentLibrariesLibraryIdRevisionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versionBump** | **string** |  | 
 **files** | ***os.File** |  | 
 **id** | **string** | Specifies a specific id to use for the library. Used for syncronizing libraries across Retool Instances. | 
 **version** | **string** | A specific version tag to use. Also specify version_bump as &#39;specify_version&#39;. | 

### Return type

[**CustomComponentLibrariesLibraryIdRevisionsPost200Response**](CustomComponentLibrariesLibraryIdRevisionsPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet

> CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet(ctx, libraryId, revisionId).Execute()

List files for a custom component library revision



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
	libraryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	revisionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomComponentLibrariesAPI.CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet(context.Background(), libraryId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomComponentLibrariesAPI.CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet`: CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomComponentLibrariesAPI.CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**libraryId** | **string** |  | 
**revisionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response**](CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CustomComponentLibrariesPost

> CustomComponentLibrariesPost200Response CustomComponentLibrariesPost(ctx).CustomComponentLibrariesPostRequest(customComponentLibrariesPostRequest).Execute()

Create a custom component library



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
	customComponentLibrariesPostRequest := *openapiclient.NewCustomComponentLibrariesPostRequest("Name_example", "Description_example", "Label_example") // CustomComponentLibrariesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomComponentLibrariesAPI.CustomComponentLibrariesPost(context.Background()).CustomComponentLibrariesPostRequest(customComponentLibrariesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomComponentLibrariesAPI.CustomComponentLibrariesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomComponentLibrariesPost`: CustomComponentLibrariesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomComponentLibrariesAPI.CustomComponentLibrariesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomComponentLibrariesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customComponentLibrariesPostRequest** | [**CustomComponentLibrariesPostRequest**](CustomComponentLibrariesPostRequest.md) |  | 

### Return type

[**CustomComponentLibrariesPost200Response**](CustomComponentLibrariesPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

