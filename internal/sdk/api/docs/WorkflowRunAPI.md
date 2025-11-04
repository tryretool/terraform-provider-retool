# \WorkflowRunAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkflowRunIdGet**](WorkflowRunAPI.md#WorkflowRunIdGet) | **Get** /workflow_run/{id} | Get Workflow Run Details



## WorkflowRunIdGet

> WorkflowRunIdGet200Response WorkflowRunIdGet(ctx, id).Execute()

Get Workflow Run Details



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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the workflow run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowRunAPI.WorkflowRunIdGet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowRunAPI.WorkflowRunIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkflowRunIdGet`: WorkflowRunIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkflowRunAPI.WorkflowRunIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the workflow run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowRunIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowRunIdGet200Response**](WorkflowRunIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

