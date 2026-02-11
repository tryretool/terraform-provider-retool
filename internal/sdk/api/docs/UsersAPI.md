# \UsersAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersGet**](UsersAPI.md#UsersGet) | **Get** /users | List users
[**UsersPost**](UsersAPI.md#UsersPost) | **Post** /users | Create user
[**UsersReset2faUserIdPut**](UsersAPI.md#UsersReset2faUserIdPut) | **Put** /users/reset2fa/{userId} | Resets a user&#39;s existing two factor authentication setting
[**UsersUserIdDelete**](UsersAPI.md#UsersUserIdDelete) | **Delete** /users/{userId} | Delete a user
[**UsersUserIdGet**](UsersAPI.md#UsersUserIdGet) | **Get** /users/{userId} | Get a user
[**UsersUserIdPatch**](UsersAPI.md#UsersUserIdPatch) | **Patch** /users/{userId} | Update a user
[**UsersUserIdUserAttributesAttributeNameDelete**](UsersAPI.md#UsersUserIdUserAttributesAttributeNameDelete) | **Delete** /users/{userId}/user_attributes/{attributeName} | Delete a user attribute
[**UsersUserIdUserAttributesPost**](UsersAPI.md#UsersUserIdUserAttributesPost) | **Post** /users/{userId}/user_attributes | Add or update a user attribute



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
	email := "email_example" // string | Email address of user (optional)
	firstName := "firstName_example" // string | First name of user (optional)
	lastName := "lastName_example" // string | Last name of user (optional)
	limit := int32(50) // int32 | Maximum number of items to return per page. If not provided, all items are returned. When provided, enables pagination and the response will include next_token for retrieving subsequent pages. Valid range: 1-100. (optional)
	nextToken := "eyJsYXN0SWQiOjEyM30..." // string | Cursor token for retrieving the next page of results. Obtained from the next_token field of a previous paginated response. Only valid when the limit parameter is also provided. (optional)

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
 **email** | **string** | Email address of user | 
 **firstName** | **string** | First name of user | 
 **lastName** | **string** | Last name of user | 
 **limit** | **int32** | Maximum number of items to return per page. If not provided, all items are returned. When provided, enables pagination and the response will include next_token for retrieving subsequent pages. Valid range: 1-100. | 
 **nextToken** | **string** | Cursor token for retrieving the next page of results. Obtained from the next_token field of a previous paginated response. Only valid when the limit parameter is also provided. | 

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

Create user



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

Resets a user's existing two factor authentication setting



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
	userId := "userId_example" // string | 

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
	userId := "userId_example" // string | 

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
	userId := "userId_example" // string | 
	includeGroups := true // bool | Whether to include the groups array in the response. (optional) (default to false)

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

 **includeGroups** | **bool** | Whether to include the groups array in the response. | [default to false]

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
	userId := "userId_example" // string | 
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
	userId := "userId_example" // string | 
	attributeName := "attributeName_example" // string | The name of the user attribute to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUserIdUserAttributesAttributeNameDelete(context.Background(), userId, attributeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUserIdUserAttributesAttributeNameDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUserIdUserAttributesAttributeNameDelete`: UsersUserIdUserAttributesAttributeNameDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUserIdUserAttributesAttributeNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**attributeName** | **string** | The name of the user attribute to delete | 

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

Add or update a user attribute



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
	userId := "userId_example" // string | 
	usersUserIdUserAttributesPostRequest := *openapiclient.NewUsersUserIdUserAttributesPostRequest("Name_example", "Value_example") // UsersUserIdUserAttributesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUserIdUserAttributesPost(context.Background(), userId).UsersUserIdUserAttributesPostRequest(usersUserIdUserAttributesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUserIdUserAttributesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUserIdUserAttributesPost`: UsersUserIdUserAttributesPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUserIdUserAttributesPost`: %v\n", resp)
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

