# \AccessTokensAPI

All URIs are relative to *http://.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessTokensAccessTokenIdGet**](AccessTokensAPI.md#AccessTokensAccessTokenIdGet) | **Get** /access_tokens/:accessTokenId | Get a single access token
[**AccessTokensGet**](AccessTokensAPI.md#AccessTokensGet) | **Get** /access_tokens | List access tokens
[**AccessTokensPost**](AccessTokensAPI.md#AccessTokensPost) | **Post** /access_tokens | Create an access token



## AccessTokensAccessTokenIdGet

> AccessTokensAccessTokenIdGet200Response AccessTokensAccessTokenIdGet(ctx).Execute()

Get a single access token



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
	resp, r, err := apiClient.AccessTokensAPI.AccessTokensAccessTokenIdGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPI.AccessTokensAccessTokenIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessTokensAccessTokenIdGet`: AccessTokensAccessTokenIdGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPI.AccessTokensAccessTokenIdGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccessTokensAccessTokenIdGetRequest struct via the builder pattern


### Return type

[**AccessTokensAccessTokenIdGet200Response**](AccessTokensAccessTokenIdGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessTokensGet

> AccessTokensGet200Response AccessTokensGet(ctx).Execute()

List access tokens



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
	resp, r, err := apiClient.AccessTokensAPI.AccessTokensGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPI.AccessTokensGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessTokensGet`: AccessTokensGet200Response
	fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPI.AccessTokensGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAccessTokensGetRequest struct via the builder pattern


### Return type

[**AccessTokensGet200Response**](AccessTokensGet200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessTokensPost

> AccessTokensPost200Response AccessTokensPost(ctx).AccessTokensPostRequest(accessTokensPostRequest).Execute()

Create an access token



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
	accessTokensPostRequest := *openapiclient.NewAccessTokensPostRequest("Label_example", []string{"Scopes_example"}) // AccessTokensPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessTokensAPI.AccessTokensPost(context.Background()).AccessTokensPostRequest(accessTokensPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensAPI.AccessTokensPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessTokensPost`: AccessTokensPost200Response
	fmt.Fprintf(os.Stdout, "Response from `AccessTokensAPI.AccessTokensPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessTokensPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessTokensPostRequest** | [**AccessTokensPostRequest**](AccessTokensPostRequest.md) |  | 

### Return type

[**AccessTokensPost200Response**](AccessTokensPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

