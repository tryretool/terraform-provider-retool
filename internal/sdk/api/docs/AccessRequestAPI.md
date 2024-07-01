# \AccessRequestAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessRequestsAccessRequestIdGet**](AccessRequestAPI.md#AccessRequestsAccessRequestIdGet) | **Get** /access_requests/{accessRequestId} | Get access request
[**AccessRequestsAccessRequestIdPatch**](AccessRequestAPI.md#AccessRequestsAccessRequestIdPatch) | **Patch** /access_requests/{accessRequestId} | Approve or deny an access request
[**AccessRequestsGet**](AccessRequestAPI.md#AccessRequestsGet) | **Get** /access_requests | Get organization access requests



## AccessRequestsAccessRequestIdGet

> AccessRequestsAccessRequestIdGet200Response AccessRequestsAccessRequestIdGet(ctx, accessRequestId).Execute()

Get access request



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
	accessRequestId := "accessRequestId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestAPI.AccessRequestsAccessRequestIdGet(context.Background(), accessRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestAPI.AccessRequestsAccessRequestIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessRequestsAccessRequestIdGet`: AccessRequestsAccessRequestIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestAPI.AccessRequestsAccessRequestIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessRequestsAccessRequestIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessRequestsAccessRequestIdGet200Response**](AccessRequestsAccessRequestIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessRequestsAccessRequestIdPatch

> AccessRequestsAccessRequestIdPatch200Response AccessRequestsAccessRequestIdPatch(ctx, accessRequestId).AccessRequestsAccessRequestIdPatchRequest(accessRequestsAccessRequestIdPatchRequest).Execute()

Approve or deny an access request



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
	accessRequestId := "accessRequestId_example" // string | 
	accessRequestsAccessRequestIdPatchRequest := *openapiclient.NewAccessRequestsAccessRequestIdPatchRequest([]openapiclient.UsersUserIdPatchRequestOperationsInnerAnyOf2{*openapiclient.NewUsersUserIdPatchRequestOperationsInnerAnyOf2("Op_example", "Path_example")}) // AccessRequestsAccessRequestIdPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestAPI.AccessRequestsAccessRequestIdPatch(context.Background(), accessRequestId).AccessRequestsAccessRequestIdPatchRequest(accessRequestsAccessRequestIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestAPI.AccessRequestsAccessRequestIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessRequestsAccessRequestIdPatch`: AccessRequestsAccessRequestIdPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestAPI.AccessRequestsAccessRequestIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccessRequestsAccessRequestIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessRequestsAccessRequestIdPatchRequest** | [**AccessRequestsAccessRequestIdPatchRequest**](AccessRequestsAccessRequestIdPatchRequest.md) |  | 

### Return type

[**AccessRequestsAccessRequestIdPatch200Response**](AccessRequestsAccessRequestIdPatch200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessRequestsGet

> AccessRequestsGet200Response AccessRequestsGet(ctx).Status(status).Execute()

Get organization access requests



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
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRequestAPI.AccessRequestsGet(context.Background()).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRequestAPI.AccessRequestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessRequestsGet`: AccessRequestsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AccessRequestAPI.AccessRequestsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessRequestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | 

### Return type

[**AccessRequestsGet200Response**](AccessRequestsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

