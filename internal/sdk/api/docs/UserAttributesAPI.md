# \UserAttributesAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserAttributesGet**](UserAttributesAPI.md#UserAttributesGet) | **Get** /user_attributes | Get organization user attributes
[**UserAttributesIdDelete**](UserAttributesAPI.md#UserAttributesIdDelete) | **Delete** /user_attributes/{id} | Delete an user attribute from the organization
[**UserAttributesIdPatch**](UserAttributesAPI.md#UserAttributesIdPatch) | **Patch** /user_attributes/{id} | Update organization user attribute
[**UserAttributesPost**](UserAttributesAPI.md#UserAttributesPost) | **Post** /user_attributes | Create a new user attribute for the organization



## UserAttributesGet

> UserAttributesGet200Response UserAttributesGet(ctx).Execute()

Get organization user attributes



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
	resp, r, err := apiClient.UserAttributesAPI.UserAttributesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAttributesAPI.UserAttributesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserAttributesGet`: UserAttributesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAttributesAPI.UserAttributesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserAttributesGetRequest struct via the builder pattern


### Return type

[**UserAttributesGet200Response**](UserAttributesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserAttributesIdDelete

> UserAttributesIdDelete(ctx, id).UpdateExistingUsers(updateExistingUsers).Execute()

Delete an user attribute from the organization



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the user attribute
	updateExistingUsers := "updateExistingUsers_example" // string | Whether to update existing users with the deleted attribute (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAttributesAPI.UserAttributesIdDelete(context.Background(), id).UpdateExistingUsers(updateExistingUsers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAttributesAPI.UserAttributesIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the user attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserAttributesIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExistingUsers** | **string** | Whether to update existing users with the deleted attribute | 

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


## UserAttributesIdPatch

> UserAttributesIdPatch200Response UserAttributesIdPatch(ctx, id).UserAttributesIdPatchRequest(userAttributesIdPatchRequest).Execute()

Update organization user attribute



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the user attribute
	userAttributesIdPatchRequest := *openapiclient.NewUserAttributesIdPatchRequest([]openapiclient.ReplaceOperation{*openapiclient.NewReplaceOperation("Op_example", "Path_example")}) // UserAttributesIdPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAttributesAPI.UserAttributesIdPatch(context.Background(), id).UserAttributesIdPatchRequest(userAttributesIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAttributesAPI.UserAttributesIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserAttributesIdPatch`: UserAttributesIdPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAttributesAPI.UserAttributesIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the user attribute | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserAttributesIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userAttributesIdPatchRequest** | [**UserAttributesIdPatchRequest**](UserAttributesIdPatchRequest.md) |  | 

### Return type

[**UserAttributesIdPatch200Response**](UserAttributesIdPatch200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserAttributesPost

> UserAttributesPost200Response UserAttributesPost(ctx).UserAttributesPostRequest(userAttributesPostRequest).Execute()

Create a new user attribute for the organization



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
	userAttributesPostRequest := *openapiclient.NewUserAttributesPostRequest("Name_example", "Label_example", "DataType_example") // UserAttributesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAttributesAPI.UserAttributesPost(context.Background()).UserAttributesPostRequest(userAttributesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAttributesAPI.UserAttributesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserAttributesPost`: UserAttributesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAttributesAPI.UserAttributesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserAttributesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAttributesPostRequest** | [**UserAttributesPostRequest**](UserAttributesPostRequest.md) |  | 

### Return type

[**UserAttributesPost200Response**](UserAttributesPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

