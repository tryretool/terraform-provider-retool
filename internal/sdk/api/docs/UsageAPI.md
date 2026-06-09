# \UsageAPI

All URIs are relative to *https://stable-4-0.retool.dev/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsageAppDetailsGet**](UsageAPI.md#UsageAppDetailsGet) | **Get** /usage/app_details | Get app usage details
[**UsageAppSummaryGet**](UsageAPI.md#UsageAppSummaryGet) | **Get** /usage/app_summary | Get app usage summaries
[**UsageGet**](UsageAPI.md#UsageGet) | **Get** /usage | Get usage summary
[**UsageOrganizationsGet**](UsageAPI.md#UsageOrganizationsGet) | **Get** /usage/organizations | List organizations
[**UsageUserDetailsGet**](UsageAPI.md#UsageUserDetailsGet) | **Get** /usage/user_details | Get user usage details
[**UsageUserSummaryGet**](UsageAPI.md#UsageUserSummaryGet) | **Get** /usage/user_summary | Get user usage summaries



## UsageAppDetailsGet

> UsageAppDetailsGet200Response UsageAppDetailsGet(ctx).StartDate(startDate).AppName(appName).OrgIds(orgIds).EndDate(endDate).Limit(limit).NextToken(nextToken).Execute()

Get app usage details



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
	startDate := "2024-01-15" // string | 
	appName := "appName_example" // string | 
	orgIds := "org_id1,org_id2" // string |  (optional)
	endDate := "2024-01-30" // string |  (optional)
	limit := int32(50) // int32 |  (optional)
	nextToken := "eyJsYXN0SWQiOjEyM30..." // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.UsageAppDetailsGet(context.Background()).StartDate(startDate).AppName(appName).OrgIds(orgIds).EndDate(endDate).Limit(limit).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.UsageAppDetailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageAppDetailsGet`: UsageAppDetailsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.UsageAppDetailsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsageAppDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **appName** | **string** |  | 
 **orgIds** | **string** |  | 
 **endDate** | **string** |  | 
 **limit** | **int32** |  | 
 **nextToken** | **string** |  | 

### Return type

[**UsageAppDetailsGet200Response**](UsageAppDetailsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsageAppSummaryGet

> UsageAppSummaryGet200Response UsageAppSummaryGet(ctx).StartDate(startDate).OrgIds(orgIds).EndDate(endDate).Execute()

Get app usage summaries



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
	startDate := "2024-01-15" // string | 
	orgIds := "org_id1,org_id2" // string |  (optional)
	endDate := "2024-01-30" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.UsageAppSummaryGet(context.Background()).StartDate(startDate).OrgIds(orgIds).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.UsageAppSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageAppSummaryGet`: UsageAppSummaryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.UsageAppSummaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsageAppSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **orgIds** | **string** |  | 
 **endDate** | **string** |  | 

### Return type

[**UsageAppSummaryGet200Response**](UsageAppSummaryGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsageGet

> UsageGet200Response UsageGet(ctx).StartDate(startDate).OrgIds(orgIds).EndDate(endDate).Execute()

Get usage summary



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
	startDate := "2024-01-15" // string | 
	orgIds := "org_id1,org_id2" // string |  (optional)
	endDate := "2024-01-30" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.UsageGet(context.Background()).StartDate(startDate).OrgIds(orgIds).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.UsageGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageGet`: UsageGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.UsageGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsageGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **orgIds** | **string** |  | 
 **endDate** | **string** |  | 

### Return type

[**UsageGet200Response**](UsageGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsageOrganizationsGet

> UsageOrganizationsGet200Response UsageOrganizationsGet(ctx).Execute()

List organizations



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
	resp, r, err := apiClient.UsageAPI.UsageOrganizationsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.UsageOrganizationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageOrganizationsGet`: UsageOrganizationsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.UsageOrganizationsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUsageOrganizationsGetRequest struct via the builder pattern


### Return type

[**UsageOrganizationsGet200Response**](UsageOrganizationsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsageUserDetailsGet

> UsageUserDetailsGet200Response UsageUserDetailsGet(ctx).StartDate(startDate).Email(email).OrgIds(orgIds).EndDate(endDate).Execute()

Get user usage details



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
	startDate := "2024-01-15" // string | 
	email := "email_example" // string | 
	orgIds := "org_id1,org_id2" // string |  (optional)
	endDate := "2024-01-30" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.UsageUserDetailsGet(context.Background()).StartDate(startDate).Email(email).OrgIds(orgIds).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.UsageUserDetailsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageUserDetailsGet`: UsageUserDetailsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.UsageUserDetailsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsageUserDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **email** | **string** |  | 
 **orgIds** | **string** |  | 
 **endDate** | **string** |  | 

### Return type

[**UsageUserDetailsGet200Response**](UsageUserDetailsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsageUserSummaryGet

> UsageUserSummaryGet200Response UsageUserSummaryGet(ctx).StartDate(startDate).OrgIds(orgIds).EndDate(endDate).Execute()

Get user usage summaries



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
	startDate := "2024-01-15" // string | 
	orgIds := "org_id1,org_id2" // string |  (optional)
	endDate := "2024-01-30" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsageAPI.UsageUserSummaryGet(context.Background()).StartDate(startDate).OrgIds(orgIds).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageAPI.UsageUserSummaryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsageUserSummaryGet`: UsageUserSummaryGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsageAPI.UsageUserSummaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsageUserSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startDate** | **string** |  | 
 **orgIds** | **string** |  | 
 **endDate** | **string** |  | 

### Return type

[**UsageUserSummaryGet200Response**](UsageUserSummaryGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

