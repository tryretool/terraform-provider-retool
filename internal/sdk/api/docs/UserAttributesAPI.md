# \UserAttributesAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserAttributesGet**](UserAttributesAPI.md#UserAttributesGet) | **Get** /user_attributes | Get organization user attributes.



## UserAttributesGet

> UserAttributesGet200Response UserAttributesGet(ctx).Execute()

Get organization user attributes.



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

