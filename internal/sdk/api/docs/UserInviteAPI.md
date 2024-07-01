# \UserInviteAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserInvitesGet**](UserInviteAPI.md#UserInvitesGet) | **Get** /user_invites | Get organization user invites
[**UserInvitesPost**](UserInviteAPI.md#UserInvitesPost) | **Post** /user_invites | Create a new user invite
[**UserInvitesUserInviteIdDelete**](UserInviteAPI.md#UserInvitesUserInviteIdDelete) | **Delete** /user_invites/{userInviteId} | Delete user invite
[**UserInvitesUserInviteIdGet**](UserInviteAPI.md#UserInvitesUserInviteIdGet) | **Get** /user_invites/{userInviteId} | Get user invite
[**UserInvitesUserInviteIdUserAttributesAttributeNameDelete**](UserInviteAPI.md#UserInvitesUserInviteIdUserAttributesAttributeNameDelete) | **Delete** /user_invites/{userInviteId}/user_attributes/{attributeName} | Delete a user attribute on a user invite
[**UserInvitesUserInviteIdUserAttributesPost**](UserInviteAPI.md#UserInvitesUserInviteIdUserAttributesPost) | **Post** /user_invites/{userInviteId}/user_attributes | Add or update user attributes on a user invite



## UserInvitesGet

> UserInvitesGet200Response UserInvitesGet(ctx).Execute()

Get organization user invites



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
	resp, r, err := apiClient.UserInviteAPI.UserInvitesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserInviteAPI.UserInvitesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserInvitesGet`: UserInvitesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserInviteAPI.UserInvitesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserInvitesGetRequest struct via the builder pattern


### Return type

[**UserInvitesGet200Response**](UserInvitesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInvitesPost

> UserInvitesPost200Response UserInvitesPost(ctx).UserInvitesPostRequest(userInvitesPostRequest).Execute()

Create a new user invite



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
	userInvitesPostRequest := *openapiclient.NewUserInvitesPostRequest("Email_example") // UserInvitesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserInviteAPI.UserInvitesPost(context.Background()).UserInvitesPostRequest(userInvitesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserInviteAPI.UserInvitesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserInvitesPost`: UserInvitesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserInviteAPI.UserInvitesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserInvitesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userInvitesPostRequest** | [**UserInvitesPostRequest**](UserInvitesPostRequest.md) |  | 

### Return type

[**UserInvitesPost200Response**](UserInvitesPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInvitesUserInviteIdDelete

> UserInvitesUserInviteIdDelete(ctx, userInviteId).Execute()

Delete user invite



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
	userInviteId := "userInviteId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserInviteAPI.UserInvitesUserInviteIdDelete(context.Background(), userInviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserInviteAPI.UserInvitesUserInviteIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userInviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserInvitesUserInviteIdDeleteRequest struct via the builder pattern


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


## UserInvitesUserInviteIdGet

> UserInvitesUserInviteIdGet200Response UserInvitesUserInviteIdGet(ctx, userInviteId).Execute()

Get user invite



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
	userInviteId := "userInviteId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserInviteAPI.UserInvitesUserInviteIdGet(context.Background(), userInviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserInviteAPI.UserInvitesUserInviteIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserInvitesUserInviteIdGet`: UserInvitesUserInviteIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserInviteAPI.UserInvitesUserInviteIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userInviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserInvitesUserInviteIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserInvitesUserInviteIdGet200Response**](UserInvitesUserInviteIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInvitesUserInviteIdUserAttributesAttributeNameDelete

> UserInvitesUserInviteIdUserAttributesAttributeNameDelete200Response UserInvitesUserInviteIdUserAttributesAttributeNameDelete(ctx, userInviteId, attributeName).Execute()

Delete a user attribute on a user invite



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
	userInviteId := "userInviteId_example" // string | 
	attributeName := "attributeName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserInviteAPI.UserInvitesUserInviteIdUserAttributesAttributeNameDelete(context.Background(), userInviteId, attributeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserInviteAPI.UserInvitesUserInviteIdUserAttributesAttributeNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserInvitesUserInviteIdUserAttributesAttributeNameDelete`: UserInvitesUserInviteIdUserAttributesAttributeNameDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `UserInviteAPI.UserInvitesUserInviteIdUserAttributesAttributeNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userInviteId** | **string** |  | 
**attributeName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserInvitesUserInviteIdUserAttributesAttributeNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserInvitesUserInviteIdUserAttributesAttributeNameDelete200Response**](UserInvitesUserInviteIdUserAttributesAttributeNameDelete200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserInvitesUserInviteIdUserAttributesPost

> UserInvitesUserInviteIdUserAttributesPost200Response UserInvitesUserInviteIdUserAttributesPost(ctx, userInviteId).UserInvitesUserInviteIdUserAttributesPostRequest(userInvitesUserInviteIdUserAttributesPostRequest).Execute()

Add or update user attributes on a user invite



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
	userInviteId := "userInviteId_example" // string | 
	userInvitesUserInviteIdUserAttributesPostRequest := *openapiclient.NewUserInvitesUserInviteIdUserAttributesPostRequest("Name_example", "Value_example") // UserInvitesUserInviteIdUserAttributesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserInviteAPI.UserInvitesUserInviteIdUserAttributesPost(context.Background(), userInviteId).UserInvitesUserInviteIdUserAttributesPostRequest(userInvitesUserInviteIdUserAttributesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserInviteAPI.UserInvitesUserInviteIdUserAttributesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserInvitesUserInviteIdUserAttributesPost`: UserInvitesUserInviteIdUserAttributesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UserInviteAPI.UserInvitesUserInviteIdUserAttributesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userInviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserInvitesUserInviteIdUserAttributesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userInvitesUserInviteIdUserAttributesPostRequest** | [**UserInvitesUserInviteIdUserAttributesPostRequest**](UserInvitesUserInviteIdUserAttributesPostRequest.md) |  | 

### Return type

[**UserInvitesUserInviteIdUserAttributesPost200Response**](UserInvitesUserInviteIdUserAttributesPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

