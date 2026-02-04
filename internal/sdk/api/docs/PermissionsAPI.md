# \PermissionsAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermissionsAccessListObjectTypeObjectIdGet**](PermissionsAPI.md#PermissionsAccessListObjectTypeObjectIdGet) | **Get** /permissions/accessList/{objectType}/{objectId} | Get the access list of an app or folder
[**PermissionsGrantPost**](PermissionsAPI.md#PermissionsGrantPost) | **Post** /permissions/grant | Grant permissions
[**PermissionsListObjectsPost**](PermissionsAPI.md#PermissionsListObjectsPost) | **Post** /permissions/listObjects | List objects a group can access
[**PermissionsRevokePost**](PermissionsAPI.md#PermissionsRevokePost) | **Post** /permissions/revoke | Revoke permissions
[**RolePermissionsRemoveRoleGrantsForSubjectPost**](PermissionsAPI.md#RolePermissionsRemoveRoleGrantsForSubjectPost) | **Post** /role_permissions/remove_role_grants_for_subject/ | Delete a role grant for a subject
[**RolePermissionsRoleGrantsGet**](PermissionsAPI.md#RolePermissionsRoleGrantsGet) | **Get** /role_permissions/role_grants | Get role grants
[**RolePermissionsRoleGrantsPost**](PermissionsAPI.md#RolePermissionsRoleGrantsPost) | **Post** /role_permissions/role_grants | Create a role grant
[**RolePermissionsRoleGrantsRoleGrantIdDelete**](PermissionsAPI.md#RolePermissionsRoleGrantsRoleGrantIdDelete) | **Delete** /role_permissions/role_grants/{roleGrantId} | Delete a role grant
[**RolePermissionsRoleGrantsSubjectTypeSubjectIdGet**](PermissionsAPI.md#RolePermissionsRoleGrantsSubjectTypeSubjectIdGet) | **Get** /role_permissions/role_grants/{subjectType}/{subjectId} | Get role grants for a subject
[**RolePermissionsRolesGet**](PermissionsAPI.md#RolePermissionsRolesGet) | **Get** /role_permissions/roles | List roles
[**RolePermissionsRolesPost**](PermissionsAPI.md#RolePermissionsRolesPost) | **Post** /role_permissions/roles | Create a role
[**RolePermissionsRolesRoleIdDelete**](PermissionsAPI.md#RolePermissionsRolesRoleIdDelete) | **Delete** /role_permissions/roles/{roleId} | Delete a role
[**RolePermissionsRolesRoleIdPatch**](PermissionsAPI.md#RolePermissionsRolesRoleIdPatch) | **Patch** /role_permissions/roles/{roleId} | Update a role



## PermissionsAccessListObjectTypeObjectIdGet

> PermissionsAccessListObjectTypeObjectIdGet200Response PermissionsAccessListObjectTypeObjectIdGet(ctx, objectId, objectType).Execute()

Get the access list of an app or folder



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
	objectId := "objectId_example" // string | Object ID. Apps can be accessed by numeric id or UUID. Folders can be accessed by numeric id.
	objectType := "objectType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.PermissionsAccessListObjectTypeObjectIdGet(context.Background(), objectId, objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.PermissionsAccessListObjectTypeObjectIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionsAccessListObjectTypeObjectIdGet`: PermissionsAccessListObjectTypeObjectIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.PermissionsAccessListObjectTypeObjectIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | Object ID. Apps can be accessed by numeric id or UUID. Folders can be accessed by numeric id. | 
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsAccessListObjectTypeObjectIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PermissionsAccessListObjectTypeObjectIdGet200Response**](PermissionsAccessListObjectTypeObjectIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGrantPost

> PermissionsGrantPost200Response PermissionsGrantPost(ctx).PermissionsGrantPostRequest(permissionsGrantPostRequest).Execute()

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
	permissionsGrantPostRequest := *openapiclient.NewPermissionsGrantPostRequest(openapiclient._permissions_listObjects_post_request_subject{Group: openapiclient.NewGroup("Type_example", NullableFloat32(123))}, openapiclient._permissions_grant_post_request_object{App: openapiclient.NewApp("Type_example", "Id_example")}, "AccessLevel_example") // PermissionsGrantPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.PermissionsGrantPost(context.Background()).PermissionsGrantPostRequest(permissionsGrantPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.PermissionsGrantPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionsGrantPost`: PermissionsGrantPost200Response
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

[**PermissionsGrantPost200Response**](PermissionsGrantPost200Response.md)

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

List objects a group can access



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
	permissionsListObjectsPostRequest := *openapiclient.NewPermissionsListObjectsPostRequest(openapiclient._permissions_listObjects_post_request_subject{Group: openapiclient.NewGroup("Type_example", NullableFloat32(123))}, "ObjectType_example") // PermissionsListObjectsPostRequest |  (optional)

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

> PermissionsGrantPost200Response PermissionsRevokePost(ctx).PermissionsRevokePostRequest(permissionsRevokePostRequest).Execute()

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
	permissionsRevokePostRequest := *openapiclient.NewPermissionsRevokePostRequest(openapiclient._permissions_listObjects_post_request_subject{Group: openapiclient.NewGroup("Type_example", NullableFloat32(123))}, openapiclient._permissions_grant_post_request_object{App: openapiclient.NewApp("Type_example", "Id_example")}) // PermissionsRevokePostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.PermissionsRevokePost(context.Background()).PermissionsRevokePostRequest(permissionsRevokePostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.PermissionsRevokePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionsRevokePost`: PermissionsGrantPost200Response
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

[**PermissionsGrantPost200Response**](PermissionsGrantPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePermissionsRemoveRoleGrantsForSubjectPost

> RolePermissionsRemoveRoleGrantsForSubjectPost(ctx).RolePermissionsRoleGrantsPostRequest(rolePermissionsRoleGrantsPostRequest).Execute()

Delete a role grant for a subject



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
	rolePermissionsRoleGrantsPostRequest := *openapiclient.NewRolePermissionsRoleGrantsPostRequest("RoleId_example", "SubjectType_example", "SubjectId_example") // RolePermissionsRoleGrantsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionsAPI.RolePermissionsRemoveRoleGrantsForSubjectPost(context.Background()).RolePermissionsRoleGrantsPostRequest(rolePermissionsRoleGrantsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.RolePermissionsRemoveRoleGrantsForSubjectPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRolePermissionsRemoveRoleGrantsForSubjectPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rolePermissionsRoleGrantsPostRequest** | [**RolePermissionsRoleGrantsPostRequest**](RolePermissionsRoleGrantsPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePermissionsRoleGrantsGet

> RolePermissionsRoleGrantsGet(ctx).Execute()

Get role grants



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
	r, err := apiClient.PermissionsAPI.RolePermissionsRoleGrantsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.RolePermissionsRoleGrantsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRolePermissionsRoleGrantsGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePermissionsRoleGrantsPost

> RolePermissionsRoleGrantsPost(ctx).RolePermissionsRoleGrantsPostRequest(rolePermissionsRoleGrantsPostRequest).Execute()

Create a role grant



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
	rolePermissionsRoleGrantsPostRequest := *openapiclient.NewRolePermissionsRoleGrantsPostRequest("RoleId_example", "SubjectType_example", "SubjectId_example") // RolePermissionsRoleGrantsPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionsAPI.RolePermissionsRoleGrantsPost(context.Background()).RolePermissionsRoleGrantsPostRequest(rolePermissionsRoleGrantsPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.RolePermissionsRoleGrantsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRolePermissionsRoleGrantsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rolePermissionsRoleGrantsPostRequest** | [**RolePermissionsRoleGrantsPostRequest**](RolePermissionsRoleGrantsPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePermissionsRoleGrantsRoleGrantIdDelete

> RolePermissionsRoleGrantsRoleGrantIdDelete(ctx, roleGrantId).Execute()

Delete a role grant



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
	roleGrantId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the role grant

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionsAPI.RolePermissionsRoleGrantsRoleGrantIdDelete(context.Background(), roleGrantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.RolePermissionsRoleGrantsRoleGrantIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleGrantId** | **string** | The id of the role grant | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolePermissionsRoleGrantsRoleGrantIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePermissionsRoleGrantsSubjectTypeSubjectIdGet

> RolePermissionsRoleGrantsSubjectTypeSubjectIdGet(ctx, subjectType, subjectId).Execute()

Get role grants for a subject



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
	subjectType := "subjectType_example" // string | The type of the subject
	subjectId := "subjectId_example" // string | The id of the subject

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionsAPI.RolePermissionsRoleGrantsSubjectTypeSubjectIdGet(context.Background(), subjectType, subjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.RolePermissionsRoleGrantsSubjectTypeSubjectIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectType** | **string** | The type of the subject | 
**subjectId** | **string** | The id of the subject | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolePermissionsRoleGrantsSubjectTypeSubjectIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePermissionsRolesGet

> RolePermissionsRolesGet200Response RolePermissionsRolesGet(ctx).Execute()

List roles



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
	resp, r, err := apiClient.PermissionsAPI.RolePermissionsRolesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.RolePermissionsRolesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RolePermissionsRolesGet`: RolePermissionsRolesGet200Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.RolePermissionsRolesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRolePermissionsRolesGetRequest struct via the builder pattern


### Return type

[**RolePermissionsRolesGet200Response**](RolePermissionsRolesGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePermissionsRolesPost

> RolePermissionsRolesPost201Response RolePermissionsRolesPost(ctx).RolePermissionsRolesPostRequest(rolePermissionsRolesPostRequest).Execute()

Create a role



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
	rolePermissionsRolesPostRequest := *openapiclient.NewRolePermissionsRolesPostRequest("Name_example", []string{"ObjectScopes_example"}, []string{"OrganizationScopes_example"}) // RolePermissionsRolesPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.RolePermissionsRolesPost(context.Background()).RolePermissionsRolesPostRequest(rolePermissionsRolesPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.RolePermissionsRolesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RolePermissionsRolesPost`: RolePermissionsRolesPost201Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.RolePermissionsRolesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRolePermissionsRolesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rolePermissionsRolesPostRequest** | [**RolePermissionsRolesPostRequest**](RolePermissionsRolesPostRequest.md) |  | 

### Return type

[**RolePermissionsRolesPost201Response**](RolePermissionsRolesPost201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePermissionsRolesRoleIdDelete

> RolePermissionsRolesRoleIdDelete(ctx, roleId).Execute()

Delete a role



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
	roleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the role

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionsAPI.RolePermissionsRolesRoleIdDelete(context.Background(), roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.RolePermissionsRolesRoleIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The id of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolePermissionsRolesRoleIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolePermissionsRolesRoleIdPatch

> RolePermissionsRolesPost201Response RolePermissionsRolesRoleIdPatch(ctx, roleId).RolePermissionsRolesRoleIdPatchRequest(rolePermissionsRolesRoleIdPatchRequest).Execute()

Update a role



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
	roleId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The id of the role
	rolePermissionsRolesRoleIdPatchRequest := *openapiclient.NewRolePermissionsRolesRoleIdPatchRequest() // RolePermissionsRolesRoleIdPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.RolePermissionsRolesRoleIdPatch(context.Background(), roleId).RolePermissionsRolesRoleIdPatchRequest(rolePermissionsRolesRoleIdPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.RolePermissionsRolesRoleIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RolePermissionsRolesRoleIdPatch`: RolePermissionsRolesPost201Response
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.RolePermissionsRolesRoleIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | The id of the role | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolePermissionsRolesRoleIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rolePermissionsRolesRoleIdPatchRequest** | [**RolePermissionsRolesRoleIdPatchRequest**](RolePermissionsRolesRoleIdPatchRequest.md) |  | 

### Return type

[**RolePermissionsRolesPost201Response**](RolePermissionsRolesPost201Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

