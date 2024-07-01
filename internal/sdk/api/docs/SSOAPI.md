# \SSOAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SsoConfigPost**](SSOAPI.md#SsoConfigPost) | **Post** /sso/config | Set SSO configuration



## SsoConfigPost

> SsoConfigPost200Response SsoConfigPost(ctx).SsoConfigPostRequest(ssoConfigPostRequest).Execute()

Set SSO configuration



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
	ssoConfigPostRequest := *openapiclient.NewSsoConfigPostRequest(openapiclient._sso_config_post_request_data{SsoConfigPostRequestDataOneOf: openapiclient.NewSsoConfigPostRequestDataOneOf("ConfigType_example", "GoogleClientId_example", "GoogleClientSecret_example", false)}) // SsoConfigPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SSOAPI.SsoConfigPost(context.Background()).SsoConfigPostRequest(ssoConfigPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOAPI.SsoConfigPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SsoConfigPost`: SsoConfigPost200Response
	fmt.Fprintf(os.Stdout, "Response from `SSOAPI.SsoConfigPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSsoConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssoConfigPostRequest** | [**SsoConfigPostRequest**](SsoConfigPostRequest.md) |  | 

### Return type

[**SsoConfigPost200Response**](SsoConfigPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

