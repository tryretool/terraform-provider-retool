# \UsersAPI

All URIs are relative to *https://stable-4-0.retool.dev/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGet**](UsersAPI.md#UsersGet) | **Get** /users | List users
[**UsersPost**](UsersAPI.md#UsersPost) | **Post** /users | Create a user
[**UsersReset2faUserIdPut**](UsersAPI.md#UsersReset2faUserIdPut) | **Put** /users/reset2fa/{userId} | Reset a user&#39;s two-factor authentication
[**UsersUserIdDelete**](UsersAPI.md#UsersUserIdDelete) | **Delete** /users/{userId} | Delete a user
[**UsersUserIdGet**](UsersAPI.md#UsersUserIdGet) | **Get** /users/{userId} | Get a user
[**UsersUserIdLogoutPost**](UsersAPI.md#UsersUserIdLogoutPost) | **Post** /users/{userId}/logout | Log out a user
[**UsersUserIdPatch**](UsersAPI.md#UsersUserIdPatch) | **Patch** /users/{userId} | Update a user



## UsersGet

> UsersGet200Response UsersGet(ctx).Email(email).FirstName(firstName).LastName(lastName).Limit(limit).NextToken(nextToken).Execute()

List users



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
	email := "email_example" // string |  (optional)
	firstName := "firstName_example" // string |  (optional)
	lastName := "lastName_example" // string |  (optional)
	limit := int32(50) // int32 |  (optional)
	nextToken := "eyJsYXN0SWQiOjEyM30..." // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersGet(context.Background()).Email(email).FirstName(firstName).LastName(lastName).Limit(limit).NextToken(nextToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersGet`: UsersGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** |  | 
 **firstName** | **string** |  | 
 **lastName** | **string** |  | 
 **limit** | **int32** |  | 
 **nextToken** | **string** |  | 

### Return type

[**UsersGet200Response**](UsersGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersPost

> UsersPost200Response UsersPost(ctx).UsersPostRequest(usersPostRequest).Execute()

Create a user



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
	usersPostRequest := *openapiclient.NewUsersPostRequest("Email_example", "FirstName_example", "LastName_example") // UsersPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersPost(context.Background()).UsersPostRequest(usersPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersPost`: UsersPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usersPostRequest** | [**UsersPostRequest**](UsersPostRequest.md) |  | 

### Return type

[**UsersPost200Response**](UsersPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersReset2faUserIdPut

> UsersReset2faUserIdPut(ctx, userId).Execute()

Reset a user's two-factor authentication



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersReset2faUserIdPut(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersReset2faUserIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersReset2faUserIdPutRequest struct via the builder pattern


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


## UsersUserIdDelete

> UsersUserIdDelete(ctx, userId).Execute()

Delete a user



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersUserIdDelete(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUserIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserIdDeleteRequest struct via the builder pattern


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


## UsersUserIdGet

> UsersUserIdGet200Response UsersUserIdGet(ctx, userId).IncludeGroups(includeGroups).Execute()

Get a user



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
	includeGroups := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUserIdGet(context.Background(), userId).IncludeGroups(includeGroups).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUserIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUserIdGet`: UsersUserIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeGroups** | **bool** |  | [default to false]

### Return type

[**UsersUserIdGet200Response**](UsersUserIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserIdLogoutPost

> UsersUserIdLogoutPost200Response UsersUserIdLogoutPost(ctx, userId).UsersUserIdLogoutPostRequest(usersUserIdLogoutPostRequest).Execute()

Log out a user



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
	usersUserIdLogoutPostRequest := *openapiclient.NewUsersUserIdLogoutPostRequest() // UsersUserIdLogoutPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUserIdLogoutPost(context.Background(), userId).UsersUserIdLogoutPostRequest(usersUserIdLogoutPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUserIdLogoutPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUserIdLogoutPost`: UsersUserIdLogoutPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUserIdLogoutPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserIdLogoutPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usersUserIdLogoutPostRequest** | [**UsersUserIdLogoutPostRequest**](UsersUserIdLogoutPostRequest.md) |  | 

### Return type

[**UsersUserIdLogoutPost200Response**](UsersUserIdLogoutPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserIdPatch

> UsersUserIdPatch200Response UsersUserIdPatch(ctx, userId).UsersUserIdPatchRequest(usersUserIdPatchRequest).Execute()

Update a user



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
	usersUserIdPatchRequest := *openapiclient.NewUsersUserIdPatchRequest([]openapiclient.UsersUserIdPatchRequestOperationsInner{*openapiclient.NewUsersUserIdPatchRequestOperationsInner("Op_example", "Path_example")}) // UsersUserIdPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUserIdPatch(context.Background(), userId).UsersUserIdPatchRequest(usersUserIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUserIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUserIdPatch`: UsersUserIdPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUserIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usersUserIdPatchRequest** | [**UsersUserIdPatchRequest**](UsersUserIdPatchRequest.md) |  | 

### Return type

[**UsersUserIdPatch200Response**](UsersUserIdPatch200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

