# \PermissionsAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermissionsGrantPost**](PermissionsAPI.md#PermissionsGrantPost) | **Post** /permissions/grant | Grant permissions
[**PermissionsListObjectsPost**](PermissionsAPI.md#PermissionsListObjectsPost) | **Post** /permissions/listObjects | List folders a group can access
[**PermissionsRevokePost**](PermissionsAPI.md#PermissionsRevokePost) | **Post** /permissions/revoke | Revoke permissions



## PermissionsGrantPost

> PermissionsListObjectsPost200Response PermissionsGrantPost(ctx).PermissionsGrantPostRequest(permissionsGrantPostRequest).Execute()

Grant permissions



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
	permissionsGrantPostRequest := *openapiclient.NewPermissionsGrantPostRequest(openapiclient._permissions_listObjects_post_request_subject{PermissionsListObjectsPostRequestSubjectOneOf: openapiclient.NewPermissionsListObjectsPostRequestSubjectOneOf("Type_example", NullableFloat32(123))}, openapiclient._permissions_grant_post_request_object{PermissionsGrantPostRequestObjectOneOf: openapiclient.NewPermissionsGrantPostRequestObjectOneOf("Type_example", "Id_example")}, "AccessLevel_example") // PermissionsGrantPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.PermissionsGrantPost(context.Background()).PermissionsGrantPostRequest(permissionsGrantPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.PermissionsGrantPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionsGrantPost`: PermissionsListObjectsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.PermissionsGrantPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGrantPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissionsGrantPostRequest** | [**PermissionsGrantPostRequest**](PermissionsGrantPostRequest.md) |  | 

### Return type

[**PermissionsListObjectsPost200Response**](PermissionsListObjectsPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsListObjectsPost

> PermissionsListObjectsPost200Response PermissionsListObjectsPost(ctx).PermissionsListObjectsPostRequest(permissionsListObjectsPostRequest).Execute()

List folders a group can access



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
	permissionsListObjectsPostRequest := *openapiclient.NewPermissionsListObjectsPostRequest(openapiclient._permissions_listObjects_post_request_subject{PermissionsListObjectsPostRequestSubjectOneOf: openapiclient.NewPermissionsListObjectsPostRequestSubjectOneOf("Type_example", NullableFloat32(123))}, "ObjectType_example") // PermissionsListObjectsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.PermissionsListObjectsPost(context.Background()).PermissionsListObjectsPostRequest(permissionsListObjectsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.PermissionsListObjectsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionsListObjectsPost`: PermissionsListObjectsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.PermissionsListObjectsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsListObjectsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissionsListObjectsPostRequest** | [**PermissionsListObjectsPostRequest**](PermissionsListObjectsPostRequest.md) |  | 

### Return type

[**PermissionsListObjectsPost200Response**](PermissionsListObjectsPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsRevokePost

> PermissionsListObjectsPost200Response PermissionsRevokePost(ctx).PermissionsRevokePostRequest(permissionsRevokePostRequest).Execute()

Revoke permissions



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
	permissionsRevokePostRequest := *openapiclient.NewPermissionsRevokePostRequest(openapiclient._permissions_listObjects_post_request_subject{PermissionsListObjectsPostRequestSubjectOneOf: openapiclient.NewPermissionsListObjectsPostRequestSubjectOneOf("Type_example", NullableFloat32(123))}, openapiclient._permissions_grant_post_request_object{PermissionsGrantPostRequestObjectOneOf: openapiclient.NewPermissionsGrantPostRequestObjectOneOf("Type_example", "Id_example")}) // PermissionsRevokePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.PermissionsRevokePost(context.Background()).PermissionsRevokePostRequest(permissionsRevokePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.PermissionsRevokePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionsRevokePost`: PermissionsListObjectsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.PermissionsRevokePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsRevokePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissionsRevokePostRequest** | [**PermissionsRevokePostRequest**](PermissionsRevokePostRequest.md) |  | 

### Return type

[**PermissionsListObjectsPost200Response**](PermissionsListObjectsPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

