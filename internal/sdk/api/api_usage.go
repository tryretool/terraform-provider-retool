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
	"reflect"
)


// UsageAPIService UsageAPI service
type UsageAPIService service

type ApiUsageAppDetailsGetRequest struct {
	ctx context.Context
	ApiService *UsageAPIService
	startDate *string
	appName *string
	orgIds *[]string
	endDate *string
}

func (r ApiUsageAppDetailsGetRequest) StartDate(startDate string) ApiUsageAppDetailsGetRequest {
	r.startDate = &startDate
	return r
}

func (r ApiUsageAppDetailsGetRequest) AppName(appName string) ApiUsageAppDetailsGetRequest {
	r.appName = &appName
	return r
}

func (r ApiUsageAppDetailsGetRequest) OrgIds(orgIds []string) ApiUsageAppDetailsGetRequest {
	r.orgIds = &orgIds
	return r
}

func (r ApiUsageAppDetailsGetRequest) EndDate(endDate string) ApiUsageAppDetailsGetRequest {
	r.endDate = &endDate
	return r
}

func (r ApiUsageAppDetailsGetRequest) Execute() (*UsageAppDetailsGet200Response, *http.Response, error) {
	return r.ApiService.UsageAppDetailsGetExecute(r)
}

/*
UsageAppDetailsGet The app details for the selected app and organizations

The detailed app usage for the selected organizations. Contains information about the app, including the breakdown of saves and views in the specified time range. The API token must have the 'usage' scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUsageAppDetailsGetRequest
*/
func (a *UsageAPIService) UsageAppDetailsGet(ctx context.Context) ApiUsageAppDetailsGetRequest {
	return ApiUsageAppDetailsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UsageAppDetailsGet200Response
func (a *UsageAPIService) UsageAppDetailsGetExecute(r ApiUsageAppDetailsGetRequest) (*UsageAppDetailsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsageAppDetailsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageAPIService.UsageAppDetailsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage/app_details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.startDate == nil {
		return localVarReturnValue, nil, reportError("startDate is required and must be specified")
	}
	if r.appName == nil {
		return localVarReturnValue, nil, reportError("appName is required and must be specified")
	}

	if r.orgIds != nil {
		t := *r.orgIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "org_ids", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "org_ids", t, "multi")
		}
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "")
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "app_name", r.appName, "")
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

type ApiUsageAppSummaryGetRequest struct {
	ctx context.Context
	ApiService *UsageAPIService
	startDate *string
	orgIds *[]string
	endDate *string
}

func (r ApiUsageAppSummaryGetRequest) StartDate(startDate string) ApiUsageAppSummaryGetRequest {
	r.startDate = &startDate
	return r
}

func (r ApiUsageAppSummaryGetRequest) OrgIds(orgIds []string) ApiUsageAppSummaryGetRequest {
	r.orgIds = &orgIds
	return r
}

func (r ApiUsageAppSummaryGetRequest) EndDate(endDate string) ApiUsageAppSummaryGetRequest {
	r.endDate = &endDate
	return r
}

func (r ApiUsageAppSummaryGetRequest) Execute() (*UsageAppSummaryGet200Response, *http.Response, error) {
	return r.ApiService.UsageAppSummaryGetExecute(r)
}

/*
UsageAppSummaryGet The app summaries for the selected organizations

The app summaries for the selected organizations. Contains information about the app like saves, views, unique editors and viewers in the specified time range. The API token must have the 'usage' scope. 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUsageAppSummaryGetRequest
*/
func (a *UsageAPIService) UsageAppSummaryGet(ctx context.Context) ApiUsageAppSummaryGetRequest {
	return ApiUsageAppSummaryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UsageAppSummaryGet200Response
func (a *UsageAPIService) UsageAppSummaryGetExecute(r ApiUsageAppSummaryGetRequest) (*UsageAppSummaryGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsageAppSummaryGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageAPIService.UsageAppSummaryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage/app_summary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.startDate == nil {
		return localVarReturnValue, nil, reportError("startDate is required and must be specified")
	}

	if r.orgIds != nil {
		t := *r.orgIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "org_ids", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "org_ids", t, "multi")
		}
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "")
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "")
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

type ApiUsageGetRequest struct {
	ctx context.Context
	ApiService *UsageAPIService
	startDate *string
	orgIds *[]string
	endDate *string
}

func (r ApiUsageGetRequest) StartDate(startDate string) ApiUsageGetRequest {
	r.startDate = &startDate
	return r
}

func (r ApiUsageGetRequest) OrgIds(orgIds []string) ApiUsageGetRequest {
	r.orgIds = &orgIds
	return r
}

func (r ApiUsageGetRequest) EndDate(endDate string) ApiUsageGetRequest {
	r.endDate = &endDate
	return r
}

func (r ApiUsageGetRequest) Execute() (*UsageGet200Response, *http.Response, error) {
	return r.ApiService.UsageGetExecute(r)
}

/*
UsageGet The usage summary for the selected organizations

The usage summary for the selected organizations. Contains information about usage like page saves, page views, active users, and the growth in those fields in the specified time range. The API token must have the 'usage' scope. 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUsageGetRequest
*/
func (a *UsageAPIService) UsageGet(ctx context.Context) ApiUsageGetRequest {
	return ApiUsageGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UsageGet200Response
func (a *UsageAPIService) UsageGetExecute(r ApiUsageGetRequest) (*UsageGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsageGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageAPIService.UsageGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.startDate == nil {
		return localVarReturnValue, nil, reportError("startDate is required and must be specified")
	}

	if r.orgIds != nil {
		t := *r.orgIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "org_ids", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "org_ids", t, "multi")
		}
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "")
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "")
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

type ApiUsageOrganizationsGetRequest struct {
	ctx context.Context
	ApiService *UsageAPIService
}

func (r ApiUsageOrganizationsGetRequest) Execute() (*UsageOrganizationsGet200Response, *http.Response, error) {
	return r.ApiService.UsageOrganizationsGetExecute(r)
}

/*
UsageOrganizationsGet List organizations

Returns a list of organizations that the token has scope to access. The API token must have the 'usage' scope. 'usage:organization' scope returns just the organization, 'usage:spaces' returns all the spaces under the admin organization, 'usage:license' returns all the organizations that use that license key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUsageOrganizationsGetRequest
*/
func (a *UsageAPIService) UsageOrganizationsGet(ctx context.Context) ApiUsageOrganizationsGetRequest {
	return ApiUsageOrganizationsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UsageOrganizationsGet200Response
func (a *UsageAPIService) UsageOrganizationsGetExecute(r ApiUsageOrganizationsGetRequest) (*UsageOrganizationsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsageOrganizationsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageAPIService.UsageOrganizationsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage/organizations"

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

type ApiUsageUserDetailsGetRequest struct {
	ctx context.Context
	ApiService *UsageAPIService
	startDate *string
	email *string
	orgIds *[]string
	endDate *string
}

func (r ApiUsageUserDetailsGetRequest) StartDate(startDate string) ApiUsageUserDetailsGetRequest {
	r.startDate = &startDate
	return r
}

func (r ApiUsageUserDetailsGetRequest) Email(email string) ApiUsageUserDetailsGetRequest {
	r.email = &email
	return r
}

func (r ApiUsageUserDetailsGetRequest) OrgIds(orgIds []string) ApiUsageUserDetailsGetRequest {
	r.orgIds = &orgIds
	return r
}

func (r ApiUsageUserDetailsGetRequest) EndDate(endDate string) ApiUsageUserDetailsGetRequest {
	r.endDate = &endDate
	return r
}

func (r ApiUsageUserDetailsGetRequest) Execute() (*UsageUserDetailsGet200Response, *http.Response, error) {
	return r.ApiService.UsageUserDetailsGetExecute(r)
}

/*
UsageUserDetailsGet The user details for the selected user and organizations

Detailed usage by user for the selected organizations. Contains information about the user, including the breakdown of saves and views for apps in the specified time range. The API token must have the 'usage' scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUsageUserDetailsGetRequest
*/
func (a *UsageAPIService) UsageUserDetailsGet(ctx context.Context) ApiUsageUserDetailsGetRequest {
	return ApiUsageUserDetailsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UsageUserDetailsGet200Response
func (a *UsageAPIService) UsageUserDetailsGetExecute(r ApiUsageUserDetailsGetRequest) (*UsageUserDetailsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsageUserDetailsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageAPIService.UsageUserDetailsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage/user_details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.startDate == nil {
		return localVarReturnValue, nil, reportError("startDate is required and must be specified")
	}
	if r.email == nil {
		return localVarReturnValue, nil, reportError("email is required and must be specified")
	}

	if r.orgIds != nil {
		t := *r.orgIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "org_ids", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "org_ids", t, "multi")
		}
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "")
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "email", r.email, "")
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

type ApiUsageUserSummaryGetRequest struct {
	ctx context.Context
	ApiService *UsageAPIService
	startDate *string
	orgIds *[]string
	endDate *string
}

func (r ApiUsageUserSummaryGetRequest) StartDate(startDate string) ApiUsageUserSummaryGetRequest {
	r.startDate = &startDate
	return r
}

func (r ApiUsageUserSummaryGetRequest) OrgIds(orgIds []string) ApiUsageUserSummaryGetRequest {
	r.orgIds = &orgIds
	return r
}

func (r ApiUsageUserSummaryGetRequest) EndDate(endDate string) ApiUsageUserSummaryGetRequest {
	r.endDate = &endDate
	return r
}

func (r ApiUsageUserSummaryGetRequest) Execute() (*UsageUserSummaryGet200Response, *http.Response, error) {
	return r.ApiService.UsageUserSummaryGetExecute(r)
}

/*
UsageUserSummaryGet The summaries of user usage for the selected organizations

The summaries of the user including email, last active time, the number of apps viewed and edited, for the selected organizations in the specified time range. The API token must have the 'usage' scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUsageUserSummaryGetRequest
*/
func (a *UsageAPIService) UsageUserSummaryGet(ctx context.Context) ApiUsageUserSummaryGetRequest {
	return ApiUsageUserSummaryGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UsageUserSummaryGet200Response
func (a *UsageAPIService) UsageUserSummaryGetExecute(r ApiUsageUserSummaryGetRequest) (*UsageUserSummaryGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UsageUserSummaryGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsageAPIService.UsageUserSummaryGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/usage/user_summary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.startDate == nil {
		return localVarReturnValue, nil, reportError("startDate is required and must be specified")
	}

	if r.orgIds != nil {
		t := *r.orgIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "org_ids", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "org_ids", t, "multi")
		}
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "start_date", r.startDate, "")
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end_date", r.endDate, "")
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