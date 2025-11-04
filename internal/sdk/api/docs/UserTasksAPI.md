# \UserTasksAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserTasksGet**](UserTasksAPI.md#UserTasksGet) | **Get** /user_tasks | List User Tasks
[**UserTasksTaskIdCancelPatch**](UserTasksAPI.md#UserTasksTaskIdCancelPatch) | **Patch** /user_tasks/{taskId}/cancel | Cancels a user task.
[**UserTasksTaskIdReassignPatch**](UserTasksAPI.md#UserTasksTaskIdReassignPatch) | **Patch** /user_tasks/{taskId}/reassign | Reassign a user task to a group or a user.
[**UserTasksTaskIdSubmitPatch**](UserTasksAPI.md#UserTasksTaskIdSubmitPatch) | **Patch** /user_tasks/{taskId}/submit | Submits a user task.



## UserTasksGet

> UserTasksGet200Response UserTasksGet(ctx).AssignedToUsers(assignedToUsers).Statuses(statuses).WorkflowIds(workflowIds).Limit(limit).NextToken(nextToken).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()

List User Tasks



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
	assignedToUsers := *openapiclient.NewUserTasksGetAssignedToUsersParameter() // UserTasksGetAssignedToUsersParameter | Filter tasks by assignment. Provide one or more user ids. If not provided, returns all accessible tasks. (optional)
	statuses := "statuses_example" // string | A comma-separated list of statuses. (optional)
	workflowIds := "workflowIds_example" // string | A comma-separated list of workflow IDs. (optional)
	limit := "limit_example" // string | The maximum number of tasks to include in the response. (optional)
	nextToken := "nextToken_example" // string | The next token from the previous API response. This is used for paginating through an arbitrary number of tasks. (optional)
	startingAfter := "startingAfter_example" // string | The ID of the task following which the response will include subsequent tasks. (optional)
	endingBefore := "endingBefore_example" // string | The ID of the task before which the response will include preceding tasks. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTasksAPI.UserTasksGet(context.Background()).AssignedToUsers(assignedToUsers).Statuses(statuses).WorkflowIds(workflowIds).Limit(limit).NextToken(nextToken).StartingAfter(startingAfter).EndingBefore(endingBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTasksAPI.UserTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTasksGet`: UserTasksGet200Response
	fmt.Fprintf(os.Stdout, "Response from `UserTasksAPI.UserTasksGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUserTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assignedToUsers** | [**UserTasksGetAssignedToUsersParameter**](UserTasksGetAssignedToUsersParameter.md) | Filter tasks by assignment. Provide one or more user ids. If not provided, returns all accessible tasks. | 
 **statuses** | **string** | A comma-separated list of statuses. | 
 **workflowIds** | **string** | A comma-separated list of workflow IDs. | 
 **limit** | **string** | The maximum number of tasks to include in the response. | 
 **nextToken** | **string** | The next token from the previous API response. This is used for paginating through an arbitrary number of tasks. | 
 **startingAfter** | **string** | The ID of the task following which the response will include subsequent tasks. | 
 **endingBefore** | **string** | The ID of the task before which the response will include preceding tasks. | 

### Return type

[**UserTasksGet200Response**](UserTasksGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTasksTaskIdCancelPatch

> UserTasksTaskIdSubmitPatch200Response UserTasksTaskIdCancelPatch(ctx, taskId).UserTasksTaskIdCancelPatchRequest(userTasksTaskIdCancelPatchRequest).Execute()

Cancels a user task.



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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the user task.
	userTasksTaskIdCancelPatchRequest := *openapiclient.NewUserTasksTaskIdCancelPatchRequest("Email_example") // UserTasksTaskIdCancelPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTasksAPI.UserTasksTaskIdCancelPatch(context.Background(), taskId).UserTasksTaskIdCancelPatchRequest(userTasksTaskIdCancelPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTasksAPI.UserTasksTaskIdCancelPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTasksTaskIdCancelPatch`: UserTasksTaskIdSubmitPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `UserTasksAPI.UserTasksTaskIdCancelPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The ID of the user task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserTasksTaskIdCancelPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userTasksTaskIdCancelPatchRequest** | [**UserTasksTaskIdCancelPatchRequest**](UserTasksTaskIdCancelPatchRequest.md) |  | 

### Return type

[**UserTasksTaskIdSubmitPatch200Response**](UserTasksTaskIdSubmitPatch200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTasksTaskIdReassignPatch

> UserTasksTaskIdSubmitPatch200Response UserTasksTaskIdReassignPatch(ctx, taskId).UserTasksTaskIdReassignPatchRequest(userTasksTaskIdReassignPatchRequest).Execute()

Reassign a user task to a group or a user.



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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the user task.
	userTasksTaskIdReassignPatchRequest := *openapiclient.NewUserTasksTaskIdReassignPatchRequest("Email_example", []openapiclient.UserTasksTaskIdReassignPatchRequestAssigneesInner{*openapiclient.NewUserTasksTaskIdReassignPatchRequestAssigneesInner("Type_example", "Id_example")}) // UserTasksTaskIdReassignPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTasksAPI.UserTasksTaskIdReassignPatch(context.Background(), taskId).UserTasksTaskIdReassignPatchRequest(userTasksTaskIdReassignPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTasksAPI.UserTasksTaskIdReassignPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTasksTaskIdReassignPatch`: UserTasksTaskIdSubmitPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `UserTasksAPI.UserTasksTaskIdReassignPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The ID of the user task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserTasksTaskIdReassignPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userTasksTaskIdReassignPatchRequest** | [**UserTasksTaskIdReassignPatchRequest**](UserTasksTaskIdReassignPatchRequest.md) |  | 

### Return type

[**UserTasksTaskIdSubmitPatch200Response**](UserTasksTaskIdSubmitPatch200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTasksTaskIdSubmitPatch

> UserTasksTaskIdSubmitPatch200Response UserTasksTaskIdSubmitPatch(ctx, taskId).UserTasksTaskIdSubmitPatchRequest(userTasksTaskIdSubmitPatchRequest).Execute()

Submits a user task.



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
	taskId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the user task.
	userTasksTaskIdSubmitPatchRequest := *openapiclient.NewUserTasksTaskIdSubmitPatchRequest("Email_example", map[string]interface{}{"key": interface{}(123)}) // UserTasksTaskIdSubmitPatchRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTasksAPI.UserTasksTaskIdSubmitPatch(context.Background(), taskId).UserTasksTaskIdSubmitPatchRequest(userTasksTaskIdSubmitPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTasksAPI.UserTasksTaskIdSubmitPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UserTasksTaskIdSubmitPatch`: UserTasksTaskIdSubmitPatch200Response
	fmt.Fprintf(os.Stdout, "Response from `UserTasksAPI.UserTasksTaskIdSubmitPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The ID of the user task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserTasksTaskIdSubmitPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userTasksTaskIdSubmitPatchRequest** | [**UserTasksTaskIdSubmitPatchRequest**](UserTasksTaskIdSubmitPatchRequest.md) |  | 

### Return type

[**UserTasksTaskIdSubmitPatch200Response**](UserTasksTaskIdSubmitPatch200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

