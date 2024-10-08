/*
Retool API

Go to Settings > API to get started. Once you generate an API token, use bearer token authentication to make requests.

API version: 2.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ResourcesAPIService ResourcesAPI service
type ResourcesAPIService service

type ApiResourcesGetRequest struct {
	ctx context.Context
	ApiService *ResourcesAPIService
	nextToken *string
	resourceType *string
}

func (r ApiResourcesGetRequest) NextToken(nextToken string) ApiResourcesGetRequest {
	r.nextToken = &nextToken
	return r
}

func (r ApiResourcesGetRequest) ResourceType(resourceType string) ApiResourcesGetRequest {
	r.resourceType = &resourceType
	return r
}

func (r ApiResourcesGetRequest) Execute() (*ResourcesGet200Response, *http.Response, error) {
	return r.ApiService.ResourcesGetExecute(r)
}

/*
ResourcesGet Get resources

Available from API version 2.4.0+ and onprem version 3.28.0+. Gets a paginated list of resources. The API token must have the "Resources > Read" scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiResourcesGetRequest
*/
func (a *ResourcesAPIService) ResourcesGet(ctx context.Context) ApiResourcesGetRequest {
	return ApiResourcesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResourcesGet200Response
func (a *ResourcesAPIService) ResourcesGetExecute(r ApiResourcesGetRequest) (*ResourcesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResourcesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesAPIService.ResourcesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nextToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "next_token", r.nextToken, "")
	}
	if r.resourceType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resource_type", r.resourceType, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiResourcesPostRequest struct {
	ctx context.Context
	ApiService *ResourcesAPIService
	resourcesPostRequest *ResourcesPostRequest
}

func (r ApiResourcesPostRequest) ResourcesPostRequest(resourcesPostRequest ResourcesPostRequest) ApiResourcesPostRequest {
	r.resourcesPostRequest = &resourcesPostRequest
	return r
}

func (r ApiResourcesPostRequest) Execute() (*ResourcesPost200Response, *http.Response, error) {
	return r.ApiService.ResourcesPostExecute(r)
}

/*
ResourcesPost Create a resource

Available from onprem version 3.72.0+. Creates a resource. The API token must have the "Resources > Write" scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiResourcesPostRequest
*/
func (a *ResourcesAPIService) ResourcesPost(ctx context.Context) ApiResourcesPostRequest {
	return ApiResourcesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResourcesPost200Response
func (a *ResourcesAPIService) ResourcesPostExecute(r ApiResourcesPostRequest) (*ResourcesPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResourcesPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesAPIService.ResourcesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.resourcesPostRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiResourcesResourceIdDeleteRequest struct {
	ctx context.Context
	ApiService *ResourcesAPIService
	resourceId ResourcesGet200ResponseDataInnerId
}

func (r ApiResourcesResourceIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ResourcesResourceIdDeleteExecute(r)
}

/*
ResourcesResourceIdDelete Delete resource

Available from API version 2.4.0+ and onprem version 3.28.0+. Deletes a resource with the given id. The API token must have the "Resources > Write" scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceId
 @return ApiResourcesResourceIdDeleteRequest
*/
func (a *ResourcesAPIService) ResourcesResourceIdDelete(ctx context.Context, resourceId ResourcesGet200ResponseDataInnerId) ApiResourcesResourceIdDeleteRequest {
	return ApiResourcesResourceIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		resourceId: resourceId,
	}
}

// Execute executes the request
func (a *ResourcesAPIService) ResourcesResourceIdDeleteExecute(r ApiResourcesResourceIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesAPIService.ResourcesResourceIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/{resourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", url.PathEscape(parameterValueToString(r.resourceId, "resourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiResourcesResourceIdGetRequest struct {
	ctx context.Context
	ApiService *ResourcesAPIService
	resourceId ResourcesGet200ResponseDataInnerId
}

func (r ApiResourcesResourceIdGetRequest) Execute() (*ResourcesResourceIdGet200Response, *http.Response, error) {
	return r.ApiService.ResourcesResourceIdGetExecute(r)
}

/*
ResourcesResourceIdGet Get resource by id

Available from API version 2.4.0+ and onprem version 3.28.0+. Returns a specific resource. The API token must have the "Resources > Read" scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resourceId
 @return ApiResourcesResourceIdGetRequest
*/
func (a *ResourcesAPIService) ResourcesResourceIdGet(ctx context.Context, resourceId ResourcesGet200ResponseDataInnerId) ApiResourcesResourceIdGetRequest {
	return ApiResourcesResourceIdGetRequest{
		ApiService: a,
		ctx: ctx,
		resourceId: resourceId,
	}
}

// Execute executes the request
//  @return ResourcesResourceIdGet200Response
func (a *ResourcesAPIService) ResourcesResourceIdGetExecute(r ApiResourcesResourceIdGetRequest) (*ResourcesResourceIdGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResourcesResourceIdGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcesAPIService.ResourcesResourceIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/resources/{resourceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"resourceId"+"}", url.PathEscape(parameterValueToString(r.resourceId, "resourceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
