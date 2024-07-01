# \CustomComponentLibraryAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomComponentLibrariesGet**](CustomComponentLibraryAPI.md#CustomComponentLibrariesGet) | **Get** /custom_component_libraries | Get a list of all custom component libraries
[**CustomComponentLibrariesLibraryIdGet**](CustomComponentLibraryAPI.md#CustomComponentLibrariesLibraryIdGet) | **Get** /custom_component_libraries/{libraryId} | Get a single custom component libraries
[**CustomComponentLibrariesLibraryIdRevisionsGet**](CustomComponentLibraryAPI.md#CustomComponentLibrariesLibraryIdRevisionsGet) | **Get** /custom_component_libraries/{libraryId}/revisions | Gets a list of all the revisions of a custom component library
[**CustomComponentLibrariesLibraryIdRevisionsPost**](CustomComponentLibraryAPI.md#CustomComponentLibrariesLibraryIdRevisionsPost) | **Post** /custom_component_libraries/{libraryId}/revisions | Create a new custom component library revision
[**CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet**](CustomComponentLibraryAPI.md#CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet) | **Get** /custom_component_libraries/{libraryId}/revisions/{revisionId}/files | Gets all files associated with a custom component library revision.
[**CustomComponentLibrariesPost**](CustomComponentLibraryAPI.md#CustomComponentLibrariesPost) | **Post** /custom_component_libraries | Create a new custom component library



## CustomComponentLibrariesGet

> CustomComponentLibrariesGet200Response CustomComponentLibrariesGet(ctx).NextToken(nextToken).Execute()

Get a list of all custom component libraries



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CustomComponentLibraryAPI.CustomComponentLibrariesGet(context.Background()).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomComponentLibraryAPI.CustomComponentLibrariesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomComponentLibrariesGet`: CustomComponentLibrariesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomComponentLibraryAPI.CustomComponentLibrariesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCustomComponentLibrariesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** |  | 

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

Get a single custom component libraries



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
	resp, r, err := apiClient.CustomComponentLibraryAPI.CustomComponentLibrariesLibraryIdGet(context.Background(), libraryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomComponentLibraryAPI.CustomComponentLibrariesLibraryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomComponentLibrariesLibraryIdGet`: CustomComponentLibrariesLibraryIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomComponentLibraryAPI.CustomComponentLibrariesLibraryIdGet`: %v\n", resp)
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

Gets a list of all the revisions of a custom component library



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
	resp, r, err := apiClient.CustomComponentLibraryAPI.CustomComponentLibrariesLibraryIdRevisionsGet(context.Background(), libraryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomComponentLibraryAPI.CustomComponentLibrariesLibraryIdRevisionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomComponentLibrariesLibraryIdRevisionsGet`: CustomComponentLibrariesLibraryIdRevisionsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomComponentLibraryAPI.CustomComponentLibrariesLibraryIdRevisionsGet`: %v\n", resp)
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

Create a new custom component library revision



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
	resp, r, err := apiClient.CustomComponentLibraryAPI.CustomComponentLibrariesLibraryIdRevisionsPost(context.Background(), libraryId).VersionBump(versionBump).Files(files).Id(id).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomComponentLibraryAPI.CustomComponentLibrariesLibraryIdRevisionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomComponentLibrariesLibraryIdRevisionsPost`: CustomComponentLibrariesLibraryIdRevisionsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomComponentLibraryAPI.CustomComponentLibrariesLibraryIdRevisionsPost`: %v\n", resp)
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

Gets all files associated with a custom component library revision.



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
	resp, r, err := apiClient.CustomComponentLibraryAPI.CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet(context.Background(), libraryId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomComponentLibraryAPI.CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet`: CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomComponentLibraryAPI.CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet`: %v\n", resp)
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

Create a new custom component library



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
	resp, r, err := apiClient.CustomComponentLibraryAPI.CustomComponentLibrariesPost(context.Background()).CustomComponentLibrariesPostRequest(customComponentLibrariesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CustomComponentLibraryAPI.CustomComponentLibrariesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CustomComponentLibrariesPost`: CustomComponentLibrariesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `CustomComponentLibraryAPI.CustomComponentLibrariesPost`: %v\n", resp)
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

