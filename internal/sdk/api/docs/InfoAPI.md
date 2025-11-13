# \InfoAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InfoIpAllowlistGet**](InfoAPI.md#InfoIpAllowlistGet) | **Get** /info/ip_allowlist | Get IP allowlist by region



## InfoIpAllowlistGet

> InfoIpAllowlistGet200Response InfoIpAllowlistGet(ctx).Execute()

Get IP allowlist by region



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
	resp, r, err := apiClient.InfoAPI.InfoIpAllowlistGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoAPI.InfoIpAllowlistGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InfoIpAllowlistGet`: InfoIpAllowlistGet200Response
	fmt.Fprintf(os.Stdout, "Response from `InfoAPI.InfoIpAllowlistGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInfoIpAllowlistGetRequest struct via the builder pattern


### Return type

[**InfoIpAllowlistGet200Response**](InfoIpAllowlistGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

