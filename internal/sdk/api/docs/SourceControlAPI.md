# \SourceControlAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SourceControlConfigDelete**](SourceControlAPI.md#SourceControlConfigDelete) | **Delete** /source_control/config | Delete source control provider configuration
[**SourceControlConfigGet**](SourceControlAPI.md#SourceControlConfigGet) | **Get** /source_control/config | Get source control configuration
[**SourceControlConfigPost**](SourceControlAPI.md#SourceControlConfigPost) | **Post** /source_control/config | Create source control configuration
[**SourceControlConfigPut**](SourceControlAPI.md#SourceControlConfigPut) | **Put** /source_control/config | Set source control configuration
[**SourceControlDeployPost**](SourceControlAPI.md#SourceControlDeployPost) | **Post** /source_control/deploy | Trigger deployment of latest changes
[**SourceControlDeploymentIdGet**](SourceControlAPI.md#SourceControlDeploymentIdGet) | **Get** /source_control/deployment/{id} | Get a deployment
[**SourceControlManifestsGet**](SourceControlAPI.md#SourceControlManifestsGet) | **Get** /source_control/manifests | Lists all release manifests
[**SourceControlManifestsManifestNameAppsAppUuidDeletePost**](SourceControlAPI.md#SourceControlManifestsManifestNameAppsAppUuidDeletePost) | **Post** /source_control/manifests/{manifestName}/apps/{appUuid}/delete | Delete the entry for an app from a release manifest
[**SourceControlManifestsManifestNameAppsAppUuidPut**](SourceControlAPI.md#SourceControlManifestsManifestNameAppsAppUuidPut) | **Put** /source_control/manifests/{manifestName}/apps/{appUuid} | Set release manifest
[**SourceControlManifestsManifestNameDeletePost**](SourceControlAPI.md#SourceControlManifestsManifestNameDeletePost) | **Post** /source_control/manifests/{manifestName}/delete | Delete a release manifest
[**SourceControlManifestsManifestNameElementUuidDeletePost**](SourceControlAPI.md#SourceControlManifestsManifestNameElementUuidDeletePost) | **Post** /source_control/manifests/{manifestName}/{elementUuid}/delete | Delete the entry for an app from a release manifest
[**SourceControlManifestsManifestNameElementUuidPut**](SourceControlAPI.md#SourceControlManifestsManifestNameElementUuidPut) | **Put** /source_control/manifests/{manifestName}/{elementUuid} | Set release manifest
[**SourceControlManifestsManifestNameGet**](SourceControlAPI.md#SourceControlManifestsManifestNameGet) | **Get** /source_control/manifests/{manifestName} | Get a specific release manifest
[**SourceControlManifestsManifestNamePut**](SourceControlAPI.md#SourceControlManifestsManifestNamePut) | **Put** /source_control/manifests/{manifestName} | Set release manifest
[**SourceControlReleasesAppsAppUuidGet**](SourceControlAPI.md#SourceControlReleasesAppsAppUuidGet) | **Get** /source_control/releases/apps/{appUuid} | Lists all available releases for the given app.
[**SourceControlReleasesAppsAppUuidPost**](SourceControlAPI.md#SourceControlReleasesAppsAppUuidPost) | **Post** /source_control/releases/apps/{appUuid} | Create a release artifact
[**SourceControlReleasesElementUuidGet**](SourceControlAPI.md#SourceControlReleasesElementUuidGet) | **Get** /source_control/releases/{elementUuid} | Lists all available releases for the given element.
[**SourceControlReleasesElementUuidPost**](SourceControlAPI.md#SourceControlReleasesElementUuidPost) | **Post** /source_control/releases/{elementUuid} | Create a release artifact
[**SourceControlSettingsGet**](SourceControlAPI.md#SourceControlSettingsGet) | **Get** /source_control/settings | Get source control settings
[**SourceControlSettingsPut**](SourceControlAPI.md#SourceControlSettingsPut) | **Put** /source_control/settings | Set source control settings
[**SourceControlTestConnectionGet**](SourceControlAPI.md#SourceControlTestConnectionGet) | **Get** /source_control/test_connection | Tests source control connection
[**SourceControlTestDeployPost**](SourceControlAPI.md#SourceControlTestDeployPost) | **Post** /source_control/test_deploy | Test source control changes



## SourceControlConfigDelete

> SourceControlConfigDelete(ctx).Execute()

Delete source control provider configuration



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
	r, err := apiClient.SourceControlAPI.SourceControlConfigDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlConfigDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlConfigDeleteRequest struct via the builder pattern


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


## SourceControlConfigGet

> SourceControlConfigGet200Response SourceControlConfigGet(ctx).Execute()

Get source control configuration



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
	resp, r, err := apiClient.SourceControlAPI.SourceControlConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlConfigGet`: SourceControlConfigGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlConfigGetRequest struct via the builder pattern


### Return type

[**SourceControlConfigGet200Response**](SourceControlConfigGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlConfigPost

> SourceControlConfigPost200Response SourceControlConfigPost(ctx).SourceControlConfigPutRequest(sourceControlConfigPutRequest).Execute()

Create source control configuration



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
	sourceControlConfigPutRequest := *openapiclient.NewSourceControlConfigPutRequest(*openapiclient.NewSourceControlConfigPutRequestConfig(*openapiclient.NewAzureReposConfig("Url_example", "Project_example", "User_example", "PersonalAccessToken_example", false), "Provider_example", "Org_example", "Repo_example", "DefaultBranch_example")) // SourceControlConfigPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlConfigPost(context.Background()).SourceControlConfigPutRequest(sourceControlConfigPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlConfigPost`: SourceControlConfigPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceControlConfigPutRequest** | [**SourceControlConfigPutRequest**](SourceControlConfigPutRequest.md) |  | 

### Return type

[**SourceControlConfigPost200Response**](SourceControlConfigPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlConfigPut

> SourceControlConfigPut200Response SourceControlConfigPut(ctx).SourceControlConfigPutRequest(sourceControlConfigPutRequest).Execute()

Set source control configuration



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
	sourceControlConfigPutRequest := *openapiclient.NewSourceControlConfigPutRequest(*openapiclient.NewSourceControlConfigPutRequestConfig(*openapiclient.NewAzureReposConfig("Url_example", "Project_example", "User_example", "PersonalAccessToken_example", false), "Provider_example", "Org_example", "Repo_example", "DefaultBranch_example")) // SourceControlConfigPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlConfigPut(context.Background()).SourceControlConfigPutRequest(sourceControlConfigPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlConfigPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlConfigPut`: SourceControlConfigPut200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlConfigPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlConfigPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceControlConfigPutRequest** | [**SourceControlConfigPutRequest**](SourceControlConfigPutRequest.md) |  | 

### Return type

[**SourceControlConfigPut200Response**](SourceControlConfigPut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlDeployPost

> SourceControlDeployPost200Response SourceControlDeployPost(ctx).Execute()

Trigger deployment of latest changes



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
	resp, r, err := apiClient.SourceControlAPI.SourceControlDeployPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlDeployPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlDeployPost`: SourceControlDeployPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlDeployPost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlDeployPostRequest struct via the builder pattern


### Return type

[**SourceControlDeployPost200Response**](SourceControlDeployPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlDeploymentIdGet

> SourceControlDeployPost200Response SourceControlDeploymentIdGet(ctx, id).Execute()

Get a deployment



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The deployment ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlDeploymentIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlDeploymentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlDeploymentIdGet`: SourceControlDeployPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlDeploymentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The deployment ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlDeploymentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SourceControlDeployPost200Response**](SourceControlDeployPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	resp, r, err := apiClient.SourceControlAPI.SourceControlManifestsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlManifestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsGet`: SourceControlManifestsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlManifestsGet`: %v\n", resp)
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
	resp, r, err := apiClient.SourceControlAPI.SourceControlManifestsManifestNameAppsAppUuidDeletePost(context.Background(), appUuid, manifestName).SourceControlManifestsManifestNameDeletePostRequest(sourceControlManifestsManifestNameDeletePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlManifestsManifestNameAppsAppUuidDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsManifestNameAppsAppUuidDeletePost`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlManifestsManifestNameAppsAppUuidDeletePost`: %v\n", resp)
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
	appUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The uuid to specify an app of interest within the release manifest. This should be the uuid found in the source control repository, which may differ from the organization specific uuid.
	manifestName := "manifestName_example" // string | Identifier for the manifest of interest
	sourceControlManifestsManifestNameAppsAppUuidPutRequest := *openapiclient.NewSourceControlManifestsManifestNameAppsAppUuidPutRequest(*openapiclient.NewSourceControlManifestsManifestNameAppsAppUuidPutRequestRelease()) // SourceControlManifestsManifestNameAppsAppUuidPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlManifestsManifestNameAppsAppUuidPut(context.Background(), appUuid, manifestName).SourceControlManifestsManifestNameAppsAppUuidPutRequest(sourceControlManifestsManifestNameAppsAppUuidPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlManifestsManifestNameAppsAppUuidPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsManifestNameAppsAppUuidPut`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlManifestsManifestNameAppsAppUuidPut`: %v\n", resp)
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
	resp, r, err := apiClient.SourceControlAPI.SourceControlManifestsManifestNameDeletePost(context.Background(), manifestName).SourceControlManifestsManifestNameDeletePostRequest(sourceControlManifestsManifestNameDeletePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlManifestsManifestNameDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsManifestNameDeletePost`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlManifestsManifestNameDeletePost`: %v\n", resp)
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
	elementUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The uuid to specify an element of interest within the release manifest.
	manifestName := "manifestName_example" // string | Identifier for the manifest of interest
	elementType := "elementType_example" // string | The type of the element.
	sourceControlManifestsManifestNameDeletePostRequest := *openapiclient.NewSourceControlManifestsManifestNameDeletePostRequest() // SourceControlManifestsManifestNameDeletePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlManifestsManifestNameElementUuidDeletePost(context.Background(), elementUuid, manifestName).ElementType(elementType).SourceControlManifestsManifestNameDeletePostRequest(sourceControlManifestsManifestNameDeletePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlManifestsManifestNameElementUuidDeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsManifestNameElementUuidDeletePost`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlManifestsManifestNameElementUuidDeletePost`: %v\n", resp)
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
	elementUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The uuid to specify an element of interest within the release manifest.
	manifestName := "manifestName_example" // string | Identifier for the manifest of interest
	elementType := "elementType_example" // string | The type of the element.
	sourceControlManifestsManifestNameAppsAppUuidPutRequest := *openapiclient.NewSourceControlManifestsManifestNameAppsAppUuidPutRequest(*openapiclient.NewSourceControlManifestsManifestNameAppsAppUuidPutRequestRelease()) // SourceControlManifestsManifestNameAppsAppUuidPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlManifestsManifestNameElementUuidPut(context.Background(), elementUuid, manifestName).ElementType(elementType).SourceControlManifestsManifestNameAppsAppUuidPutRequest(sourceControlManifestsManifestNameAppsAppUuidPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlManifestsManifestNameElementUuidPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsManifestNameElementUuidPut`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlManifestsManifestNameElementUuidPut`: %v\n", resp)
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
	resp, r, err := apiClient.SourceControlAPI.SourceControlManifestsManifestNameGet(context.Background(), manifestName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlManifestsManifestNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsManifestNameGet`: SourceControlManifestsManifestNameGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlManifestsManifestNameGet`: %v\n", resp)
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
	resp, r, err := apiClient.SourceControlAPI.SourceControlManifestsManifestNamePut(context.Background(), manifestName).SourceControlManifestsManifestNamePutRequest(sourceControlManifestsManifestNamePutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlManifestsManifestNamePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlManifestsManifestNamePut`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlManifestsManifestNamePut`: %v\n", resp)
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
	resp, r, err := apiClient.SourceControlAPI.SourceControlReleasesAppsAppUuidGet(context.Background(), appUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlReleasesAppsAppUuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlReleasesAppsAppUuidGet`: SourceControlReleasesAppsAppUuidGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlReleasesAppsAppUuidGet`: %v\n", resp)
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
	resp, r, err := apiClient.SourceControlAPI.SourceControlReleasesAppsAppUuidPost(context.Background(), appUuid).SourceControlReleasesAppsAppUuidPostRequest(sourceControlReleasesAppsAppUuidPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlReleasesAppsAppUuidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlReleasesAppsAppUuidPost`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlReleasesAppsAppUuidPost`: %v\n", resp)
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
	resp, r, err := apiClient.SourceControlAPI.SourceControlReleasesElementUuidGet(context.Background(), elementUuid).ElementType(elementType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlReleasesElementUuidGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlReleasesElementUuidGet`: SourceControlReleasesAppsAppUuidGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlReleasesElementUuidGet`: %v\n", resp)
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
	resp, r, err := apiClient.SourceControlAPI.SourceControlReleasesElementUuidPost(context.Background(), elementUuid).ElementType(elementType).SourceControlReleasesAppsAppUuidPostRequest(sourceControlReleasesAppsAppUuidPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlReleasesElementUuidPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlReleasesElementUuidPost`: SourceControlManifestsManifestNamePut200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlReleasesElementUuidPost`: %v\n", resp)
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


## SourceControlSettingsGet

> SourceControlSettingsGet200Response SourceControlSettingsGet(ctx).Execute()

Get source control settings



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
	resp, r, err := apiClient.SourceControlAPI.SourceControlSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlSettingsGet`: SourceControlSettingsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlSettingsGetRequest struct via the builder pattern


### Return type

[**SourceControlSettingsGet200Response**](SourceControlSettingsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlSettingsPut

> SourceControlSettingsPut200Response SourceControlSettingsPut(ctx).SourceControlSettingsPutRequest(sourceControlSettingsPutRequest).Execute()

Set source control settings



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
	sourceControlSettingsPutRequest := *openapiclient.NewSourceControlSettingsPutRequest(*openapiclient.NewSourceControlSettingsPutRequestSettings()) // SourceControlSettingsPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlSettingsPut(context.Background()).SourceControlSettingsPutRequest(sourceControlSettingsPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlSettingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlSettingsPut`: SourceControlSettingsPut200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlSettingsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceControlSettingsPutRequest** | [**SourceControlSettingsPutRequest**](SourceControlSettingsPutRequest.md) |  | 

### Return type

[**SourceControlSettingsPut200Response**](SourceControlSettingsPut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlTestConnectionGet

> SourceControlTestConnectionGet200Response SourceControlTestConnectionGet(ctx).Execute()

Tests source control connection



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
	resp, r, err := apiClient.SourceControlAPI.SourceControlTestConnectionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlTestConnectionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlTestConnectionGet`: SourceControlTestConnectionGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlTestConnectionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlTestConnectionGetRequest struct via the builder pattern


### Return type

[**SourceControlTestConnectionGet200Response**](SourceControlTestConnectionGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourceControlTestDeployPost

> SourceControlTestDeployPost200Response SourceControlTestDeployPost(ctx).SourceControlTestDeployPostRequest(sourceControlTestDeployPostRequest).Execute()

Test source control changes



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
	sourceControlTestDeployPostRequest := *openapiclient.NewSourceControlTestDeployPostRequest(*openapiclient.NewSourceControlTestDeployPostRequestDeployParams("CommitSha_example")) // SourceControlTestDeployPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.SourceControlTestDeployPost(context.Background()).SourceControlTestDeployPostRequest(sourceControlTestDeployPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.SourceControlTestDeployPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourceControlTestDeployPost`: SourceControlTestDeployPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.SourceControlTestDeployPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSourceControlTestDeployPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceControlTestDeployPostRequest** | [**SourceControlTestDeployPostRequest**](SourceControlTestDeployPostRequest.md) |  | 

### Return type

[**SourceControlTestDeployPost200Response**](SourceControlTestDeployPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

