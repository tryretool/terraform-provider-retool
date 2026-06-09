# \UserAttributesAPI

All URIs are relative to *https://stable-4-0.retool.dev/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserAttributesGet**](UserAttributesAPI.md#UserAttributesGet) | **Get** /user_attributes | List organization user attributes
[**UserAttributesIdDelete**](UserAttributesAPI.md#UserAttributesIdDelete) | **Delete** /user_attributes/{id} | Delete an organization user attribute
[**UserAttributesIdPatch**](UserAttributesAPI.md#UserAttributesIdPatch) | **Patch** /user_attributes/{id} | Update an organization user attribute
[**UserAttributesPost**](UserAttributesAPI.md#UserAttributesPost) | **Post** /user_attributes | Create an organization user attribute
[**UsersUserIdUserAttributesAttributeNameDelete**](UserAttributesAPI.md#UsersUserIdUserAttributesAttributeNameDelete) | **Delete** /users/{userId}/user_attributes/{attributeName} | Delete a user attribute
[**UsersUserIdUserAttributesPost**](UserAttributesAPI.md#UsersUserIdUserAttributesPost) | **Post** /users/{userId}/user_attributes | Create or update a user attribute



## UserAttributesGet

> UserAttributesGet200Response UserAttributesGet(ctx).Execute()

List organization user attributes



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

Delete an organization user attribute



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	updateExistingUsers := "updateExistingUsers_example" // string |  (optional)

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
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserAttributesIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateExistingUsers** | **string** |  | 

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

Update an organization user attribute



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
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
**id** | **string** |  | 

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

Create an organization user attribute



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


## UsersUserIdUserAttributesAttributeNameDelete

> UsersUserIdUserAttributesAttributeNameDelete200Response UsersUserIdUserAttributesAttributeNameDelete(ctx, userId, attributeName).Execute()

Delete a user attribute



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
	userId := "user_1234" // string | 
	attributeName := "attributeName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAttributesAPI.UsersUserIdUserAttributesAttributeNameDelete(context.Background(), userId, attributeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAttributesAPI.UsersUserIdUserAttributesAttributeNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUserIdUserAttributesAttributeNameDelete`: UsersUserIdUserAttributesAttributeNameDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAttributesAPI.UsersUserIdUserAttributesAttributeNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**attributeName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserIdUserAttributesAttributeNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UsersUserIdUserAttributesAttributeNameDelete200Response**](UsersUserIdUserAttributesAttributeNameDelete200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserIdUserAttributesPost

> UsersUserIdUserAttributesPost200Response UsersUserIdUserAttributesPost(ctx, userId).UsersUserIdUserAttributesPostRequest(usersUserIdUserAttributesPostRequest).Execute()

Create or update a user attribute



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
	userId := "user_1234" // string | 
	usersUserIdUserAttributesPostRequest := *openapiclient.NewUsersUserIdUserAttributesPostRequest("Name_example", "Value_example") // UsersUserIdUserAttributesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAttributesAPI.UsersUserIdUserAttributesPost(context.Background(), userId).UsersUserIdUserAttributesPostRequest(usersUserIdUserAttributesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAttributesAPI.UsersUserIdUserAttributesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUserIdUserAttributesPost`: UsersUserIdUserAttributesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAttributesAPI.UsersUserIdUserAttributesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserIdUserAttributesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usersUserIdUserAttributesPostRequest** | [**UsersUserIdUserAttributesPostRequest**](UsersUserIdUserAttributesPostRequest.md) |  | 

### Return type

[**UsersUserIdUserAttributesPost200Response**](UsersUserIdUserAttributesPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

