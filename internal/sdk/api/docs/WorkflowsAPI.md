# \WorkflowsAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowsGet**](WorkflowsAPI.md#WorkflowsGet) | **Get** /workflows | Get all workflows
[**WorkflowsWorkflowIdGet**](WorkflowsAPI.md#WorkflowsWorkflowIdGet) | **Get** /workflows/{workflowId} | Get a workflow



## WorkflowsGet

> WorkflowsGet200Response WorkflowsGet(ctx).Execute()

Get all workflows



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
	resp, r, err := apiClient.WorkflowsAPI.WorkflowsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.WorkflowsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkflowsGet`: WorkflowsGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.WorkflowsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowsGetRequest struct via the builder pattern


### Return type

[**WorkflowsGet200Response**](WorkflowsGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowsWorkflowIdGet

> WorkflowsWorkflowIdGet200Response WorkflowsWorkflowIdGet(ctx, workflowId).Execute()

Get a workflow



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
	workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Workflow ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.WorkflowsWorkflowIdGet(context.Background(), workflowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.WorkflowsWorkflowIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkflowsWorkflowIdGet`: WorkflowsWorkflowIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.WorkflowsWorkflowIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The Workflow ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowsWorkflowIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowsWorkflowIdGet200Response**](WorkflowsWorkflowIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

