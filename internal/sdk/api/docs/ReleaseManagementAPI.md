# \ReleaseManagementAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SourceControlManifestsGet**](ReleaseManagementAPI.md#SourceControlManifestsGet) | **Get** /source_control/manifests | Lists all release manifests
[**SourceControlManifestsManifestNameAppsAppUuidDeletePost**](ReleaseManagementAPI.md#SourceControlManifestsManifestNameAppsAppUuidDeletePost) | **Post** /source_control/manifests/{manifestName}/apps/{appUuid}/delete | Delete the entry for an app from a release manifest
[**SourceControlManifestsManifestNameAppsAppUuidPut**](ReleaseManagementAPI.md#SourceControlManifestsManifestNameAppsAppUuidPut) | **Put** /source_control/manifests/{manifestName}/apps/{appUuid} | Set release configuration for app in manifest
[**SourceControlManifestsManifestNameDeletePost**](ReleaseManagementAPI.md#SourceControlManifestsManifestNameDeletePost) | **Post** /source_control/manifests/{manifestName}/delete | Delete a release manifest
[**SourceControlManifestsManifestNameElementUuidDeletePost**](ReleaseManagementAPI.md#SourceControlManifestsManifestNameElementUuidDeletePost) | **Post** /source_control/manifests/{manifestName}/{elementUuid}/delete | Delete element from release manifest
[**SourceControlManifestsManifestNameElementUuidPut**](ReleaseManagementAPI.md#SourceControlManifestsManifestNameElementUuidPut) | **Put** /source_control/manifests/{manifestName}/{elementUuid} | Set release version for element in manifest
[**SourceControlManifestsManifestNameGet**](ReleaseManagementAPI.md#SourceControlManifestsManifestNameGet) | **Get** /source_control/manifests/{manifestName} | Get a specific release manifest
[**SourceControlManifestsManifestNamePut**](ReleaseManagementAPI.md#SourceControlManifestsManifestNamePut) | **Put** /source_control/manifests/{manifestName} | Set release manifest
[**SourceControlReleasesAppsAppUuidGet**](ReleaseManagementAPI.md#SourceControlReleasesAppsAppUuidGet) | **Get** /source_control/releases/apps/{appUuid} | Lists all available releases for the given app.
[**SourceControlReleasesAppsAppUuidPost**](ReleaseManagementAPI.md#SourceControlReleasesAppsAppUuidPost) | **Post** /source_control/releases/apps/{appUuid} | Create a release artifact
[**SourceControlReleasesElementUuidGet**](ReleaseManagementAPI.md#SourceControlReleasesElementUuidGet) | **Get** /source_control/releases/{elementUuid} | Lists all available releases for the given element.
[**SourceControlReleasesElementUuidPost**](ReleaseManagementAPI.md#SourceControlReleasesElementUuidPost) | **Post** /source_control/releases/{elementUuid} | Create a release artifact



## SourceControlManifestsGet

> SourceControlManifestsGet200Response SourceControlManifestsGet(ctx).Execute()

Lists all release manifests



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
	resp, r, err := apiClient.ReleaseManagementAPI.SourceControlManifestsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.SourceControlManifestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsGet`: SourceControlManifestsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.SourceControlManifestsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlManifestsGetRequest struct via the builder pattern


### Return type

[**SourceControlManifestsGet200Response**](SourceControlManifestsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlManifestsManifestNameAppsAppUuidDeletePost

> SourceControlManifestsManifestNamePut200Response SourceControlManifestsManifestNameAppsAppUuidDeletePost(ctx, appUuid, manifestName).SourceControlManifestsManifestNameDeletePostRequest(sourceControlManifestsManifestNameDeletePostRequest).Execute()

Delete the entry for an app from a release manifest



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
	appUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The uuid to specify an app of interest within the release manifest. This should be the uuid found in the source control repository, which may differ from the organization specific uuid.
	manifestName := "manifestName_example" // string | Identifier for the manifest of interest
	sourceControlManifestsManifestNameDeletePostRequest := *openapiclient.NewSourceControlManifestsManifestNameDeletePostRequest() // SourceControlManifestsManifestNameDeletePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.SourceControlManifestsManifestNameAppsAppUuidDeletePost(context.Background(), appUuid, manifestName).SourceControlManifestsManifestNameDeletePostRequest(sourceControlManifestsManifestNameDeletePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.SourceControlManifestsManifestNameAppsAppUuidDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsManifestNameAppsAppUuidDeletePost`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.SourceControlManifestsManifestNameAppsAppUuidDeletePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appUuid** | **string** | The uuid to specify an app of interest within the release manifest. This should be the uuid found in the source control repository, which may differ from the organization specific uuid. | 
**manifestName** | **string** | Identifier for the manifest of interest | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlManifestsManifestNameAppsAppUuidDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sourceControlManifestsManifestNameDeletePostRequest** | [**SourceControlManifestsManifestNameDeletePostRequest**](SourceControlManifestsManifestNameDeletePostRequest.md) |  | 

### Return type

[**SourceControlManifestsManifestNamePut200Response**](SourceControlManifestsManifestNamePut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlManifestsManifestNameAppsAppUuidPut

> SourceControlManifestsManifestNamePut200Response SourceControlManifestsManifestNameAppsAppUuidPut(ctx, appUuid, manifestName).SourceControlManifestsManifestNameAppsAppUuidPutRequest(sourceControlManifestsManifestNameAppsAppUuidPutRequest).Execute()

Set release configuration for app in manifest



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
	appUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The uuid to specify an app of interest within the release manifest. This should be the uuid found in the source control repository, which may differ from the organization specific uuid.
	manifestName := "manifestName_example" // string | Identifier for the manifest of interest
	sourceControlManifestsManifestNameAppsAppUuidPutRequest := *openapiclient.NewSourceControlManifestsManifestNameAppsAppUuidPutRequest(*openapiclient.NewSourceControlManifestsManifestNameAppsAppUuidPutRequestRelease()) // SourceControlManifestsManifestNameAppsAppUuidPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.SourceControlManifestsManifestNameAppsAppUuidPut(context.Background(), appUuid, manifestName).SourceControlManifestsManifestNameAppsAppUuidPutRequest(sourceControlManifestsManifestNameAppsAppUuidPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.SourceControlManifestsManifestNameAppsAppUuidPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsManifestNameAppsAppUuidPut`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.SourceControlManifestsManifestNameAppsAppUuidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appUuid** | **string** | The uuid to specify an app of interest within the release manifest. This should be the uuid found in the source control repository, which may differ from the organization specific uuid. | 
**manifestName** | **string** | Identifier for the manifest of interest | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlManifestsManifestNameAppsAppUuidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **sourceControlManifestsManifestNameAppsAppUuidPutRequest** | [**SourceControlManifestsManifestNameAppsAppUuidPutRequest**](SourceControlManifestsManifestNameAppsAppUuidPutRequest.md) |  | 

### Return type

[**SourceControlManifestsManifestNamePut200Response**](SourceControlManifestsManifestNamePut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlManifestsManifestNameDeletePost

> SourceControlManifestsManifestNamePut200Response SourceControlManifestsManifestNameDeletePost(ctx, manifestName).SourceControlManifestsManifestNameDeletePostRequest(sourceControlManifestsManifestNameDeletePostRequest).Execute()

Delete a release manifest



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
	manifestName := "manifestName_example" // string | Identifier for the manifest of interest
	sourceControlManifestsManifestNameDeletePostRequest := *openapiclient.NewSourceControlManifestsManifestNameDeletePostRequest() // SourceControlManifestsManifestNameDeletePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.SourceControlManifestsManifestNameDeletePost(context.Background(), manifestName).SourceControlManifestsManifestNameDeletePostRequest(sourceControlManifestsManifestNameDeletePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.SourceControlManifestsManifestNameDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsManifestNameDeletePost`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.SourceControlManifestsManifestNameDeletePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manifestName** | **string** | Identifier for the manifest of interest | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlManifestsManifestNameDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceControlManifestsManifestNameDeletePostRequest** | [**SourceControlManifestsManifestNameDeletePostRequest**](SourceControlManifestsManifestNameDeletePostRequest.md) |  | 

### Return type

[**SourceControlManifestsManifestNamePut200Response**](SourceControlManifestsManifestNamePut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlManifestsManifestNameElementUuidDeletePost

> SourceControlManifestsManifestNamePut200Response SourceControlManifestsManifestNameElementUuidDeletePost(ctx, elementUuid, manifestName).ElementType(elementType).SourceControlManifestsManifestNameDeletePostRequest(sourceControlManifestsManifestNameDeletePostRequest).Execute()

Delete element from release manifest



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
	elementUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The uuid to specify an element of interest within the release manifest.
	manifestName := "manifestName_example" // string | Identifier for the manifest of interest
	elementType := "elementType_example" // string | The type of the element.
	sourceControlManifestsManifestNameDeletePostRequest := *openapiclient.NewSourceControlManifestsManifestNameDeletePostRequest() // SourceControlManifestsManifestNameDeletePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.SourceControlManifestsManifestNameElementUuidDeletePost(context.Background(), elementUuid, manifestName).ElementType(elementType).SourceControlManifestsManifestNameDeletePostRequest(sourceControlManifestsManifestNameDeletePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.SourceControlManifestsManifestNameElementUuidDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsManifestNameElementUuidDeletePost`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.SourceControlManifestsManifestNameElementUuidDeletePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**elementUuid** | **string** | The uuid to specify an element of interest within the release manifest. | 
**manifestName** | **string** | Identifier for the manifest of interest | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlManifestsManifestNameElementUuidDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **elementType** | **string** | The type of the element. | 
 **sourceControlManifestsManifestNameDeletePostRequest** | [**SourceControlManifestsManifestNameDeletePostRequest**](SourceControlManifestsManifestNameDeletePostRequest.md) |  | 

### Return type

[**SourceControlManifestsManifestNamePut200Response**](SourceControlManifestsManifestNamePut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlManifestsManifestNameElementUuidPut

> SourceControlManifestsManifestNamePut200Response SourceControlManifestsManifestNameElementUuidPut(ctx, elementUuid, manifestName).ElementType(elementType).SourceControlManifestsManifestNameAppsAppUuidPutRequest(sourceControlManifestsManifestNameAppsAppUuidPutRequest).Execute()

Set release version for element in manifest



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
	elementUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The uuid to specify an element of interest within the release manifest.
	manifestName := "manifestName_example" // string | Identifier for the manifest of interest
	elementType := "elementType_example" // string | The type of the element.
	sourceControlManifestsManifestNameAppsAppUuidPutRequest := *openapiclient.NewSourceControlManifestsManifestNameAppsAppUuidPutRequest(*openapiclient.NewSourceControlManifestsManifestNameAppsAppUuidPutRequestRelease()) // SourceControlManifestsManifestNameAppsAppUuidPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.SourceControlManifestsManifestNameElementUuidPut(context.Background(), elementUuid, manifestName).ElementType(elementType).SourceControlManifestsManifestNameAppsAppUuidPutRequest(sourceControlManifestsManifestNameAppsAppUuidPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.SourceControlManifestsManifestNameElementUuidPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsManifestNameElementUuidPut`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.SourceControlManifestsManifestNameElementUuidPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**elementUuid** | **string** | The uuid to specify an element of interest within the release manifest. | 
**manifestName** | **string** | Identifier for the manifest of interest | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlManifestsManifestNameElementUuidPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **elementType** | **string** | The type of the element. | 
 **sourceControlManifestsManifestNameAppsAppUuidPutRequest** | [**SourceControlManifestsManifestNameAppsAppUuidPutRequest**](SourceControlManifestsManifestNameAppsAppUuidPutRequest.md) |  | 

### Return type

[**SourceControlManifestsManifestNamePut200Response**](SourceControlManifestsManifestNamePut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlManifestsManifestNameGet

> SourceControlManifestsManifestNameGet200Response SourceControlManifestsManifestNameGet(ctx, manifestName).Execute()

Get a specific release manifest



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
	manifestName := "manifestName_example" // string | Identifier for the manifest of interest

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.SourceControlManifestsManifestNameGet(context.Background(), manifestName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.SourceControlManifestsManifestNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsManifestNameGet`: SourceControlManifestsManifestNameGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.SourceControlManifestsManifestNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manifestName** | **string** | Identifier for the manifest of interest | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlManifestsManifestNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceControlManifestsManifestNameGet200Response**](SourceControlManifestsManifestNameGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlManifestsManifestNamePut

> SourceControlManifestsManifestNamePut200Response SourceControlManifestsManifestNamePut(ctx, manifestName).SourceControlManifestsManifestNamePutRequest(sourceControlManifestsManifestNamePutRequest).Execute()

Set release manifest



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
	manifestName := "manifestName_example" // string | Identifier for the manifest of interest
	sourceControlManifestsManifestNamePutRequest := *openapiclient.NewSourceControlManifestsManifestNamePutRequest() // SourceControlManifestsManifestNamePutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.SourceControlManifestsManifestNamePut(context.Background(), manifestName).SourceControlManifestsManifestNamePutRequest(sourceControlManifestsManifestNamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.SourceControlManifestsManifestNamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsManifestNamePut`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.SourceControlManifestsManifestNamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**manifestName** | **string** | Identifier for the manifest of interest | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlManifestsManifestNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceControlManifestsManifestNamePutRequest** | [**SourceControlManifestsManifestNamePutRequest**](SourceControlManifestsManifestNamePutRequest.md) |  | 

### Return type

[**SourceControlManifestsManifestNamePut200Response**](SourceControlManifestsManifestNamePut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlReleasesAppsAppUuidGet

> SourceControlReleasesAppsAppUuidGet200Response SourceControlReleasesAppsAppUuidGet(ctx, appUuid).Execute()

Lists all available releases for the given app.



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
	appUuid := "appUuid_example" // string | The uuid of the app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.SourceControlReleasesAppsAppUuidGet(context.Background(), appUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.SourceControlReleasesAppsAppUuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlReleasesAppsAppUuidGet`: SourceControlReleasesAppsAppUuidGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.SourceControlReleasesAppsAppUuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appUuid** | **string** | The uuid of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlReleasesAppsAppUuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceControlReleasesAppsAppUuidGet200Response**](SourceControlReleasesAppsAppUuidGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlReleasesAppsAppUuidPost

> SourceControlManifestsManifestNamePut200Response SourceControlReleasesAppsAppUuidPost(ctx, appUuid).SourceControlReleasesAppsAppUuidPostRequest(sourceControlReleasesAppsAppUuidPostRequest).Execute()

Create a release artifact



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
	appUuid := "appUuid_example" // string | The uuid of the app.
	sourceControlReleasesAppsAppUuidPostRequest := *openapiclient.NewSourceControlReleasesAppsAppUuidPostRequest("1.0.0") // SourceControlReleasesAppsAppUuidPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.SourceControlReleasesAppsAppUuidPost(context.Background(), appUuid).SourceControlReleasesAppsAppUuidPostRequest(sourceControlReleasesAppsAppUuidPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.SourceControlReleasesAppsAppUuidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlReleasesAppsAppUuidPost`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.SourceControlReleasesAppsAppUuidPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appUuid** | **string** | The uuid of the app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlReleasesAppsAppUuidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceControlReleasesAppsAppUuidPostRequest** | [**SourceControlReleasesAppsAppUuidPostRequest**](SourceControlReleasesAppsAppUuidPostRequest.md) |  | 

### Return type

[**SourceControlManifestsManifestNamePut200Response**](SourceControlManifestsManifestNamePut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlReleasesElementUuidGet

> SourceControlReleasesAppsAppUuidGet200Response SourceControlReleasesElementUuidGet(ctx, elementUuid).ElementType(elementType).Execute()

Lists all available releases for the given element.



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
	elementUuid := "elementUuid_example" // string | The uuid of the element.
	elementType := "elementType_example" // string | The type of the element.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.SourceControlReleasesElementUuidGet(context.Background(), elementUuid).ElementType(elementType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.SourceControlReleasesElementUuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlReleasesElementUuidGet`: SourceControlReleasesAppsAppUuidGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.SourceControlReleasesElementUuidGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**elementUuid** | **string** | The uuid of the element. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlReleasesElementUuidGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **elementType** | **string** | The type of the element. | 

### Return type

[**SourceControlReleasesAppsAppUuidGet200Response**](SourceControlReleasesAppsAppUuidGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlReleasesElementUuidPost

> SourceControlManifestsManifestNamePut200Response SourceControlReleasesElementUuidPost(ctx, elementUuid).ElementType(elementType).SourceControlReleasesAppsAppUuidPostRequest(sourceControlReleasesAppsAppUuidPostRequest).Execute()

Create a release artifact



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
	elementUuid := "elementUuid_example" // string | The uuid of the element.
	elementType := "elementType_example" // string | The type of the element.
	sourceControlReleasesAppsAppUuidPostRequest := *openapiclient.NewSourceControlReleasesAppsAppUuidPostRequest("1.0.0") // SourceControlReleasesAppsAppUuidPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReleaseManagementAPI.SourceControlReleasesElementUuidPost(context.Background(), elementUuid).ElementType(elementType).SourceControlReleasesAppsAppUuidPostRequest(sourceControlReleasesAppsAppUuidPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReleaseManagementAPI.SourceControlReleasesElementUuidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlReleasesElementUuidPost`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `ReleaseManagementAPI.SourceControlReleasesElementUuidPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**elementUuid** | **string** | The uuid of the element. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlReleasesElementUuidPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **elementType** | **string** | The type of the element. | 
 **sourceControlReleasesAppsAppUuidPostRequest** | [**SourceControlReleasesAppsAppUuidPostRequest**](SourceControlReleasesAppsAppUuidPostRequest.md) |  | 

### Return type

[**SourceControlManifestsManifestNamePut200Response**](SourceControlManifestsManifestNamePut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

