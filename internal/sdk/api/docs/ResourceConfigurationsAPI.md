# \ResourceConfigurationsAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResourceConfigurationsConfigurationIdDelete**](ResourceConfigurationsAPI.md#ResourceConfigurationsConfigurationIdDelete) | **Delete** /resource_configurations/{configurationId} | Delete resource configuration
[**ResourceConfigurationsConfigurationIdGet**](ResourceConfigurationsAPI.md#ResourceConfigurationsConfigurationIdGet) | **Get** /resource_configurations/{configurationId} | Get resource configuration by id
[**ResourceConfigurationsConfigurationIdPatch**](ResourceConfigurationsAPI.md#ResourceConfigurationsConfigurationIdPatch) | **Patch** /resource_configurations/{configurationId} | Update a resource configuration
[**ResourceConfigurationsGet**](ResourceConfigurationsAPI.md#ResourceConfigurationsGet) | **Get** /resource_configurations | Get resource configurations
[**ResourceConfigurationsPost**](ResourceConfigurationsAPI.md#ResourceConfigurationsPost) | **Post** /resource_configurations | Create resource configuration



## ResourceConfigurationsConfigurationIdDelete

> ResourceConfigurationsConfigurationIdDelete(ctx, configurationId).Execute()

Delete resource configuration



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
	configurationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The resource configuration id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceConfigurationsAPI.ResourceConfigurationsConfigurationIdDelete(context.Background(), configurationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceConfigurationsAPI.ResourceConfigurationsConfigurationIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | The resource configuration id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceConfigurationsConfigurationIdDeleteRequest struct via the builder pattern


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


## ResourceConfigurationsConfigurationIdGet

> ResourceConfigurationsConfigurationIdGet200Response ResourceConfigurationsConfigurationIdGet(ctx, configurationId).Execute()

Get resource configuration by id



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
	configurationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The resource configuration id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceConfigurationsAPI.ResourceConfigurationsConfigurationIdGet(context.Background(), configurationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceConfigurationsAPI.ResourceConfigurationsConfigurationIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceConfigurationsConfigurationIdGet`: ResourceConfigurationsConfigurationIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourceConfigurationsAPI.ResourceConfigurationsConfigurationIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | The resource configuration id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceConfigurationsConfigurationIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourceConfigurationsConfigurationIdGet200Response**](ResourceConfigurationsConfigurationIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceConfigurationsConfigurationIdPatch

> ResourceConfigurationsConfigurationIdPatch200Response ResourceConfigurationsConfigurationIdPatch(ctx, configurationId).ResourceConfigurationsConfigurationIdPatchRequest(resourceConfigurationsConfigurationIdPatchRequest).Execute()

Update a resource configuration



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
	configurationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The resource configuration id.
	resourceConfigurationsConfigurationIdPatchRequest := *openapiclient.NewResourceConfigurationsConfigurationIdPatchRequest([]openapiclient.ReplaceOperation{*openapiclient.NewReplaceOperation("Op_example", "Path_example")}) // ResourceConfigurationsConfigurationIdPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceConfigurationsAPI.ResourceConfigurationsConfigurationIdPatch(context.Background(), configurationId).ResourceConfigurationsConfigurationIdPatchRequest(resourceConfigurationsConfigurationIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceConfigurationsAPI.ResourceConfigurationsConfigurationIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceConfigurationsConfigurationIdPatch`: ResourceConfigurationsConfigurationIdPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourceConfigurationsAPI.ResourceConfigurationsConfigurationIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configurationId** | **string** | The resource configuration id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourceConfigurationsConfigurationIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceConfigurationsConfigurationIdPatchRequest** | [**ResourceConfigurationsConfigurationIdPatchRequest**](ResourceConfigurationsConfigurationIdPatchRequest.md) |  | 

### Return type

[**ResourceConfigurationsConfigurationIdPatch200Response**](ResourceConfigurationsConfigurationIdPatch200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceConfigurationsGet

> ResourceConfigurationsGet200Response ResourceConfigurationsGet(ctx).NextToken(nextToken).ResourceType(resourceType).ResourceId(resourceId).EnvironmentId(environmentId).Execute()

Get resource configurations



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
	nextToken := "nextToken_example" // string | The token of the current page (optional)
	resourceType := "resourceType_example" // string | The type of resource. (optional)
	resourceId := "resourceId_example" // string | The uuid or name for the resource. (optional)
	environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceConfigurationsAPI.ResourceConfigurationsGet(context.Background()).NextToken(nextToken).ResourceType(resourceType).ResourceId(resourceId).EnvironmentId(environmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceConfigurationsAPI.ResourceConfigurationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceConfigurationsGet`: ResourceConfigurationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourceConfigurationsAPI.ResourceConfigurationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResourceConfigurationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextToken** | **string** | The token of the current page | 
 **resourceType** | **string** | The type of resource. | 
 **resourceId** | **string** | The uuid or name for the resource. | 
 **environmentId** | **string** |  | 

### Return type

[**ResourceConfigurationsGet200Response**](ResourceConfigurationsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResourceConfigurationsPost

> ResourceConfigurationsPost200Response ResourceConfigurationsPost(ctx).ResourceConfigurationsPostRequest(resourceConfigurationsPostRequest).Execute()

Create resource configuration



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
	resourceConfigurationsPostRequest := *openapiclient.NewResourceConfigurationsPostRequest("ResourceId_example", "EnvironmentId_example", *openapiclient.NewResourcesPostRequestOptions(*openapiclient.NewSnowflakeOptionsDatabaseOptions("Name_example"), "AccountIdentifier_example", *openapiclient.NewRestAPIOptionsAuthenticationOptions("AuthenticationType_example", "Auth0Domain_example", "Auth0ClientId_example", "Auth0ClientSecret_example", "Auth0CustomAudience_example", "AmazonAwsRegion_example", "AmazonServiceName_example", "AmazonAccessKeyId_example", "AmazonSecretAccessKey_example", "BasicUsername_example", "DigestUsername_example", "DigestPassword_example", *openapiclient.NewRestAPIOptionsAuthenticationOptionsAnyOf4Oauth1SignatureMethod(), "Oauth1ConsumerKey_example", "Oauth1ConsumerSecret_example", "Oauth1TokenKey_example", "Oauth1TokenSecret_example", "Oauth1RealmParameter_example", "Oauth2ClientId_example", "Oauth2ClientSecret_example", "Oauth2AuthUrl_example", "Oauth2AccessTokenUrl_example"), "BaseUrl_example")) // ResourceConfigurationsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceConfigurationsAPI.ResourceConfigurationsPost(context.Background()).ResourceConfigurationsPostRequest(resourceConfigurationsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceConfigurationsAPI.ResourceConfigurationsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResourceConfigurationsPost`: ResourceConfigurationsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ResourceConfigurationsAPI.ResourceConfigurationsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResourceConfigurationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceConfigurationsPostRequest** | [**ResourceConfigurationsPostRequest**](ResourceConfigurationsPostRequest.md) |  | 

### Return type

[**ResourceConfigurationsPost200Response**](ResourceConfigurationsPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

