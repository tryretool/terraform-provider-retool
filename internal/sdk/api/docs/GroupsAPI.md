# \GroupsAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsGet**](GroupsAPI.md#GroupsGet) | **Get** /groups | List groups
[**GroupsGroupIdDelete**](GroupsAPI.md#GroupsGroupIdDelete) | **Delete** /groups/{groupId} | Delete group
[**GroupsGroupIdGet**](GroupsAPI.md#GroupsGroupIdGet) | **Get** /groups/{groupId} | Get group
[**GroupsGroupIdMembersPost**](GroupsAPI.md#GroupsGroupIdMembersPost) | **Post** /groups/{groupId}/members | Add users to group
[**GroupsGroupIdMembersUserSidDelete**](GroupsAPI.md#GroupsGroupIdMembersUserSidDelete) | **Delete** /groups/{groupId}/members/{userSid} | Remove user from group
[**GroupsGroupIdPatch**](GroupsAPI.md#GroupsGroupIdPatch) | **Patch** /groups/{groupId} | Update group
[**GroupsGroupIdPut**](GroupsAPI.md#GroupsGroupIdPut) | **Put** /groups/{groupId} | Update group
[**GroupsGroupIdUserInvitesPost**](GroupsAPI.md#GroupsGroupIdUserInvitesPost) | **Post** /groups/{groupId}/user_invites | Add user invites to group
[**GroupsGroupIdUserInvitesUserInviteIdDelete**](GroupsAPI.md#GroupsGroupIdUserInvitesUserInviteIdDelete) | **Delete** /groups/{groupId}/user_invites/{userInviteId} | Remove user invite from group
[**GroupsPost**](GroupsAPI.md#GroupsPost) | **Post** /groups | Create group



## GroupsGet

> GroupsGet200Response GroupsGet(ctx).Execute()

List groups



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
	resp, r, err := apiClient.GroupsAPI.GroupsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGet`: GroupsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGetRequest struct via the builder pattern


### Return type

[**GroupsGet200Response**](GroupsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGroupIdDelete

> GroupsGroupIdDelete(ctx, groupId).Execute()

Delete group



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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.GroupsGroupIdDelete(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsGroupIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdDeleteRequest struct via the builder pattern


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


## GroupsGroupIdGet

> GroupsGroupIdGet200Response GroupsGroupIdGet(ctx, groupId).Execute()

Get group



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
	groupId := "groupId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsGroupIdGet(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsGroupIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGroupIdGet`: GroupsGroupIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsGroupIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupsGroupIdGet200Response**](GroupsGroupIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGroupIdMembersPost

> GroupsGroupIdPut200Response GroupsGroupIdMembersPost(ctx, groupId).GroupsGroupIdMembersPostRequest(groupsGroupIdMembersPostRequest).Execute()

Add users to group



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
	groupId := "groupId_example" // string | 
	groupsGroupIdMembersPostRequest := *openapiclient.NewGroupsGroupIdMembersPostRequest([]openapiclient.GroupsGroupIdPutRequestMembersInner{*openapiclient.NewGroupsGroupIdPutRequestMembersInner("Id_example", false)}) // GroupsGroupIdMembersPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsGroupIdMembersPost(context.Background(), groupId).GroupsGroupIdMembersPostRequest(groupsGroupIdMembersPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsGroupIdMembersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGroupIdMembersPost`: GroupsGroupIdPut200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsGroupIdMembersPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdMembersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupsGroupIdMembersPostRequest** | [**GroupsGroupIdMembersPostRequest**](GroupsGroupIdMembersPostRequest.md) |  | 

### Return type

[**GroupsGroupIdPut200Response**](GroupsGroupIdPut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGroupIdMembersUserSidDelete

> GroupsGroupIdPut200Response GroupsGroupIdMembersUserSidDelete(ctx, groupId, userSid).Execute()

Remove user from group



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
	groupId := "groupId_example" // string | 
	userSid := "userSid_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsGroupIdMembersUserSidDelete(context.Background(), groupId, userSid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsGroupIdMembersUserSidDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGroupIdMembersUserSidDelete`: GroupsGroupIdPut200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsGroupIdMembersUserSidDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 
**userSid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdMembersUserSidDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupsGroupIdPut200Response**](GroupsGroupIdPut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGroupIdPatch

> GroupsGroupIdPut200Response GroupsGroupIdPatch(ctx, groupId).GroupsGroupIdPatchRequest(groupsGroupIdPatchRequest).Execute()

Update group



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
	groupId := "groupId_example" // string | 
	groupsGroupIdPatchRequest := *openapiclient.NewGroupsGroupIdPatchRequest([]openapiclient.UsersUserIdPatchRequestOperationsInner{*openapiclient.NewUsersUserIdPatchRequestOperationsInner("Op_example", "Path_example")}) // GroupsGroupIdPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsGroupIdPatch(context.Background(), groupId).GroupsGroupIdPatchRequest(groupsGroupIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsGroupIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGroupIdPatch`: GroupsGroupIdPut200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsGroupIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupsGroupIdPatchRequest** | [**GroupsGroupIdPatchRequest**](GroupsGroupIdPatchRequest.md) |  | 

### Return type

[**GroupsGroupIdPut200Response**](GroupsGroupIdPut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGroupIdPut

> GroupsGroupIdPut200Response GroupsGroupIdPut(ctx, groupId).GroupsGroupIdPutRequest(groupsGroupIdPutRequest).Execute()

Update group



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
	groupId := "groupId_example" // string | 
	groupsGroupIdPutRequest := *openapiclient.NewGroupsGroupIdPutRequest() // GroupsGroupIdPutRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsGroupIdPut(context.Background(), groupId).GroupsGroupIdPutRequest(groupsGroupIdPutRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsGroupIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGroupIdPut`: GroupsGroupIdPut200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsGroupIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupsGroupIdPutRequest** | [**GroupsGroupIdPutRequest**](GroupsGroupIdPutRequest.md) |  | 

### Return type

[**GroupsGroupIdPut200Response**](GroupsGroupIdPut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGroupIdUserInvitesPost

> GroupsGroupIdPut200Response GroupsGroupIdUserInvitesPost(ctx, groupId).GroupsGroupIdUserInvitesPostRequest(groupsGroupIdUserInvitesPostRequest).Execute()

Add user invites to group



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
	groupId := "groupId_example" // string | 
	groupsGroupIdUserInvitesPostRequest := *openapiclient.NewGroupsGroupIdUserInvitesPostRequest([]float32{float32(123)}) // GroupsGroupIdUserInvitesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsGroupIdUserInvitesPost(context.Background(), groupId).GroupsGroupIdUserInvitesPostRequest(groupsGroupIdUserInvitesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsGroupIdUserInvitesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGroupIdUserInvitesPost`: GroupsGroupIdPut200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsGroupIdUserInvitesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdUserInvitesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupsGroupIdUserInvitesPostRequest** | [**GroupsGroupIdUserInvitesPostRequest**](GroupsGroupIdUserInvitesPostRequest.md) |  | 

### Return type

[**GroupsGroupIdPut200Response**](GroupsGroupIdPut200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGroupIdUserInvitesUserInviteIdDelete

> GroupsGroupIdUserInvitesUserInviteIdDelete(ctx, groupId, userInviteId).Execute()

Remove user invite from group



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
	groupId := "groupId_example" // string | 
	userInviteId := "userInviteId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.GroupsGroupIdUserInvitesUserInviteIdDelete(context.Background(), groupId, userInviteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsGroupIdUserInvitesUserInviteIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** |  | 
**userInviteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdUserInvitesUserInviteIdDeleteRequest struct via the builder pattern


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


## GroupsPost

> GroupsPost200Response GroupsPost(ctx).GroupsPostRequest(groupsPostRequest).Execute()

Create group



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
	groupsPostRequest := *openapiclient.NewGroupsPostRequest("Name_example") // GroupsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsPost(context.Background()).GroupsPostRequest(groupsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsPost`: GroupsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupsPostRequest** | [**GroupsPostRequest**](GroupsPostRequest.md) |  | 

### Return type

[**GroupsPost200Response**](GroupsPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

