# \SSOAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SsoConfigDelete**](SSOAPI.md#SsoConfigDelete) | **Delete** /sso/config | Remove SSO configuration
[**SsoConfigGet**](SSOAPI.md#SsoConfigGet) | **Get** /sso/config | Get SSO configuration
[**SsoConfigPost**](SSOAPI.md#SsoConfigPost) | **Post** /sso/config | Set SSO configuration



## SsoConfigDelete

> SsoConfigDelete(ctx).Execute()

Remove SSO configuration



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
	r, err := apiClient.SSOAPI.SsoConfigDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOAPI.SsoConfigDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSsoConfigDeleteRequest struct via the builder pattern


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


## SsoConfigGet

> SsoConfigGet200Response SsoConfigGet(ctx).Execute()

Get SSO configuration



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
	resp, r, err := apiClient.SSOAPI.SsoConfigGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SSOAPI.SsoConfigGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SsoConfigGet`: SsoConfigGet200Response
	fmt.Fprintf(os.Stdout, "Response from `SSOAPI.SsoConfigGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSsoConfigGetRequest struct via the builder pattern


### Return type

[**SsoConfigGet200Response**](SsoConfigGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	ssoConfigPostRequest := *openapiclient.NewSsoConfigPostRequest(openapiclient._sso_config_post_request_data{Google: openapiclient.NewGoogle("ConfigType_example", "GoogleClientId_example", "GoogleClientSecret_example", false)}) // SsoConfigPostRequest |  (optional)

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

