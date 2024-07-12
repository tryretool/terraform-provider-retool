# \ResourceConfigurationsAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResourceConfigurationsConfigurationIdDelete**](ResourceConfigurationsAPI.md#ResourceConfigurationsConfigurationIdDelete) | **Delete** /resource_configurations/{configurationId} | Delete resource configuration
[**ResourceConfigurationsConfigurationIdGet**](ResourceConfigurationsAPI.md#ResourceConfigurationsConfigurationIdGet) | **Get** /resource_configurations/{configurationId} | Get resource configuration by id
[**ResourceConfigurationsGet**](ResourceConfigurationsAPI.md#ResourceConfigurationsGet) | **Get** /resource_configurations | Get resource configurations



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
	configurationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

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
**configurationId** | **string** |  | 

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
	configurationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

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
**configurationId** | **string** |  | 

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
	nextToken := "nextToken_example" // string |  (optional)
	resourceType := "resourceType_example" // string |  (optional)
	resourceId := *openapiclient.NewResourcesGet200ResponseDataInnerId() // ResourcesGet200ResponseDataInnerId |  (optional)
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
 **nextToken** | **string** |  | 
 **resourceType** | **string** |  | 
 **resourceId** | [**ResourcesGet200ResponseDataInnerId**](ResourcesGet200ResponseDataInnerId.md) |  | 
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

