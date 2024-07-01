# \AppThemesAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppThemesIdDelete**](AppThemesAPI.md#AppThemesIdDelete) | **Delete** /app_themes/{id} | Delete app theme
[**AppThemesIdGet**](AppThemesAPI.md#AppThemesIdGet) | **Get** /app_themes/{id} | Get app theme
[**AppThemesPost**](AppThemesAPI.md#AppThemesPost) | **Post** /app_themes | Create app theme
[**AppThemesPut**](AppThemesAPI.md#AppThemesPut) | **Put** /app_themes | Update app theme



## AppThemesIdDelete

> AppThemesIdDelete(ctx, id).Execute()

Delete app theme



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppThemesAPI.AppThemesIdDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppThemesAPI.AppThemesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppThemesIdDeleteRequest struct via the builder pattern


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


## AppThemesIdGet

> AppThemesIdGet200Response AppThemesIdGet(ctx, id).Execute()

Get app theme



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
	id := "id_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppThemesAPI.AppThemesIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppThemesAPI.AppThemesIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppThemesIdGet`: AppThemesIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AppThemesAPI.AppThemesIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppThemesIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppThemesIdGet200Response**](AppThemesIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppThemesPost

> AppThemesPut200Response AppThemesPost(ctx).AppThemesPutRequest(appThemesPutRequest).Execute()

Create app theme



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
	appThemesPutRequest := *openapiclient.NewAppThemesPutRequest(float32(123), float32(123), "Name_example", map[string]interface{}{"key": interface{}(123)}) // AppThemesPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppThemesAPI.AppThemesPost(context.Background()).AppThemesPutRequest(appThemesPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppThemesAPI.AppThemesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppThemesPost`: AppThemesPut200Response
	fmt.Fprintf(os.Stdout, "Response from `AppThemesAPI.AppThemesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppThemesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appThemesPutRequest** | [**AppThemesPutRequest**](AppThemesPutRequest.md) |  | 

### Return type

[**AppThemesPut200Response**](AppThemesPut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppThemesPut

> AppThemesPut200Response AppThemesPut(ctx).AppThemesPutRequest(appThemesPutRequest).Execute()

Update app theme



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
	appThemesPutRequest := *openapiclient.NewAppThemesPutRequest(float32(123), float32(123), "Name_example", map[string]interface{}{"key": interface{}(123)}) // AppThemesPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppThemesAPI.AppThemesPut(context.Background()).AppThemesPutRequest(appThemesPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppThemesAPI.AppThemesPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppThemesPut`: AppThemesPut200Response
	fmt.Fprintf(os.Stdout, "Response from `AppThemesAPI.AppThemesPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppThemesPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appThemesPutRequest** | [**AppThemesPutRequest**](AppThemesPutRequest.md) |  | 

### Return type

[**AppThemesPut200Response**](AppThemesPut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

