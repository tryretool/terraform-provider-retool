# \OrganizationAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationAiSettingsGet**](OrganizationAPI.md#OrganizationAiSettingsGet) | **Get** /organization/ai_settings | Get organization AI settings
[**OrganizationAiSettingsPut**](OrganizationAPI.md#OrganizationAiSettingsPut) | **Put** /organization/ai_settings | Update organization AI settings
[**OrganizationAnalyticsIntegrationsGet**](OrganizationAPI.md#OrganizationAnalyticsIntegrationsGet) | **Get** /organization/analytics_integrations | Get organization third-party Analytics Integrations settings
[**OrganizationGet**](OrganizationAPI.md#OrganizationGet) | **Get** /organization/ | Get organization
[**OrganizationInvalidateSessionsPost**](OrganizationAPI.md#OrganizationInvalidateSessionsPost) | **Post** /organization/invalidate_sessions/ | Invalidates all sessions for your organization
[**OrganizationPatch**](OrganizationAPI.md#OrganizationPatch) | **Patch** /organization/ | Update advanced settings on organization



## OrganizationAiSettingsGet

> OrganizationAiSettingsGet200Response OrganizationAiSettingsGet(ctx).Execute()

Get organization AI settings



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
	resp, r, err := apiClient.OrganizationAPI.OrganizationAiSettingsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.OrganizationAiSettingsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationAiSettingsGet`: OrganizationAiSettingsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.OrganizationAiSettingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationAiSettingsGetRequest struct via the builder pattern


### Return type

[**OrganizationAiSettingsGet200Response**](OrganizationAiSettingsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationAiSettingsPut

> OrganizationAiSettingsPut200Response OrganizationAiSettingsPut(ctx).OrganizationAiSettingsGet200ResponseData(organizationAiSettingsGet200ResponseData).Execute()

Update organization AI settings



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
	organizationAiSettingsGet200ResponseData := *openapiclient.NewOrganizationAiSettingsGet200ResponseData(false, false, false, false) // OrganizationAiSettingsGet200ResponseData |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.OrganizationAiSettingsPut(context.Background()).OrganizationAiSettingsGet200ResponseData(organizationAiSettingsGet200ResponseData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.OrganizationAiSettingsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationAiSettingsPut`: OrganizationAiSettingsPut200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.OrganizationAiSettingsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationAiSettingsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationAiSettingsGet200ResponseData** | [**OrganizationAiSettingsGet200ResponseData**](OrganizationAiSettingsGet200ResponseData.md) |  | 

### Return type

[**OrganizationAiSettingsPut200Response**](OrganizationAiSettingsPut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationAnalyticsIntegrationsGet

> OrganizationAnalyticsIntegrationsGet200Response OrganizationAnalyticsIntegrationsGet(ctx).Execute()

Get organization third-party Analytics Integrations settings



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
	resp, r, err := apiClient.OrganizationAPI.OrganizationAnalyticsIntegrationsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.OrganizationAnalyticsIntegrationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationAnalyticsIntegrationsGet`: OrganizationAnalyticsIntegrationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.OrganizationAnalyticsIntegrationsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationAnalyticsIntegrationsGetRequest struct via the builder pattern


### Return type

[**OrganizationAnalyticsIntegrationsGet200Response**](OrganizationAnalyticsIntegrationsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationGet

> OrganizationGet200Response OrganizationGet(ctx).Execute()

Get organization



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
	resp, r, err := apiClient.OrganizationAPI.OrganizationGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.OrganizationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationGet`: OrganizationGet200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.OrganizationGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationGetRequest struct via the builder pattern


### Return type

[**OrganizationGet200Response**](OrganizationGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrganizationInvalidateSessionsPost

> OrganizationInvalidateSessionsPost(ctx).Execute()

Invalidates all sessions for your organization



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
	r, err := apiClient.OrganizationAPI.OrganizationInvalidateSessionsPost(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.OrganizationInvalidateSessionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationInvalidateSessionsPostRequest struct via the builder pattern


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


## OrganizationPatch

> OrganizationPatch200Response OrganizationPatch(ctx).OrganizationPatchRequest(organizationPatchRequest).Execute()

Update advanced settings on organization



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
	organizationPatchRequest := *openapiclient.NewOrganizationPatchRequest([]openapiclient.ReplaceOperation{*openapiclient.NewReplaceOperation("Op_example", "Path_example")}) // OrganizationPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationAPI.OrganizationPatch(context.Background()).OrganizationPatchRequest(organizationPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationAPI.OrganizationPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationPatch`: OrganizationPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `OrganizationAPI.OrganizationPatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationPatchRequest** | [**OrganizationPatchRequest**](OrganizationPatchRequest.md) |  | 

### Return type

[**OrganizationPatch200Response**](OrganizationPatch200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

