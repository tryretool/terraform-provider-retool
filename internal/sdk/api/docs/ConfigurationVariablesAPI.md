# \ConfigurationVariablesAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigurationVariablesGet**](ConfigurationVariablesAPI.md#ConfigurationVariablesGet) | **Get** /configuration_variables | List configuration variables and their values
[**ConfigurationVariablesIdDelete**](ConfigurationVariablesAPI.md#ConfigurationVariablesIdDelete) | **Delete** /configuration_variables/{id} | Delete configuration variable
[**ConfigurationVariablesIdGet**](ConfigurationVariablesAPI.md#ConfigurationVariablesIdGet) | **Get** /configuration_variables/{id} | Retreive a single configuration variable and its values
[**ConfigurationVariablesIdPut**](ConfigurationVariablesAPI.md#ConfigurationVariablesIdPut) | **Put** /configuration_variables/{id} | Update a configuration variable
[**ConfigurationVariablesPost**](ConfigurationVariablesAPI.md#ConfigurationVariablesPost) | **Post** /configuration_variables | Create a configuration variable



## ConfigurationVariablesGet

> ConfigurationVariablesGet200Response ConfigurationVariablesGet(ctx).Execute()

List configuration variables and their values



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
	resp, r, err := apiClient.ConfigurationVariablesAPI.ConfigurationVariablesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVariablesAPI.ConfigurationVariablesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigurationVariablesGet`: ConfigurationVariablesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationVariablesAPI.ConfigurationVariablesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiConfigurationVariablesGetRequest struct via the builder pattern


### Return type

[**ConfigurationVariablesGet200Response**](ConfigurationVariablesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurationVariablesIdDelete

> ConfigurationVariablesIdDelete(ctx, id).Execute()

Delete configuration variable



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the configuration variable

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationVariablesAPI.ConfigurationVariablesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVariablesAPI.ConfigurationVariablesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the configuration variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigurationVariablesIdDeleteRequest struct via the builder pattern


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


## ConfigurationVariablesIdGet

> ConfigurationVariablesPost200Response ConfigurationVariablesIdGet(ctx, id).Execute()

Retreive a single configuration variable and its values



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the configuration variable

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationVariablesAPI.ConfigurationVariablesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVariablesAPI.ConfigurationVariablesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigurationVariablesIdGet`: ConfigurationVariablesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationVariablesAPI.ConfigurationVariablesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the configuration variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigurationVariablesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationVariablesPost200Response**](ConfigurationVariablesPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurationVariablesIdPut

> ConfigurationVariablesPost200Response ConfigurationVariablesIdPut(ctx, id).ConfigurationVariablesPostRequest(configurationVariablesPostRequest).Execute()

Update a configuration variable



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the configuration variable
	configurationVariablesPostRequest := *openapiclient.NewConfigurationVariablesPostRequest("Name_example", false, []openapiclient.ConfigurationVariablesGet200ResponseDataInnerValuesInner{*openapiclient.NewConfigurationVariablesGet200ResponseDataInnerValuesInner("EnvironmentId_example", "Value_example")}) // ConfigurationVariablesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationVariablesAPI.ConfigurationVariablesIdPut(context.Background(), id).ConfigurationVariablesPostRequest(configurationVariablesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVariablesAPI.ConfigurationVariablesIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigurationVariablesIdPut`: ConfigurationVariablesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationVariablesAPI.ConfigurationVariablesIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the configuration variable | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigurationVariablesIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configurationVariablesPostRequest** | [**ConfigurationVariablesPostRequest**](ConfigurationVariablesPostRequest.md) |  | 

### Return type

[**ConfigurationVariablesPost200Response**](ConfigurationVariablesPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfigurationVariablesPost

> ConfigurationVariablesPost200Response ConfigurationVariablesPost(ctx).ConfigurationVariablesPostRequest(configurationVariablesPostRequest).Execute()

Create a configuration variable



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
	configurationVariablesPostRequest := *openapiclient.NewConfigurationVariablesPostRequest("Name_example", false, []openapiclient.ConfigurationVariablesGet200ResponseDataInnerValuesInner{*openapiclient.NewConfigurationVariablesGet200ResponseDataInnerValuesInner("EnvironmentId_example", "Value_example")}) // ConfigurationVariablesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationVariablesAPI.ConfigurationVariablesPost(context.Background()).ConfigurationVariablesPostRequest(configurationVariablesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationVariablesAPI.ConfigurationVariablesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigurationVariablesPost`: ConfigurationVariablesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationVariablesAPI.ConfigurationVariablesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigurationVariablesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationVariablesPostRequest** | [**ConfigurationVariablesPostRequest**](ConfigurationVariablesPostRequest.md) |  | 

### Return type

[**ConfigurationVariablesPost200Response**](ConfigurationVariablesPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

