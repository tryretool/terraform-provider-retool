# ResourcesPostRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseOptions** | [**SnowflakeOptionsDatabaseOptions**](SnowflakeOptionsDatabaseOptions.md) |  | 
**AccountIdentifier** | **string** |  | 
**UserRole** | Pointer to **string** |  | [optional] 
**AuthenticationOptions** | [**RestAPIOptionsAuthenticationOptions**](RestAPIOptionsAuthenticationOptions.md) |  | 
**BaseUrl** | **string** | Use the absolute URL (e.g. https://example.com). | 
**UrlParameters** | Pointer to **[][]string** |  | [optional] 
**Headers** | Pointer to **[][]string** |  | [optional] 
**ExtraBodyValues** | Pointer to **[][]string** | Extra body values are not passed for GET or HEAD requests. | [optional] 
**CookiesToForward** | Pointer to **[]string** | You can use the pattern COOKIE_your_cookie_name in the Headers section in order to implement the double-cookie submit pattern. | [optional] 
**ForwardAllCookies** | Pointer to **bool** | This is useful if you have dynamic cookie names. | [optional] 

## Methods

### NewResourcesPostRequestOptions

`func NewResourcesPostRequestOptions(databaseOptions SnowflakeOptionsDatabaseOptions, accountIdentifier string, authenticationOptions RestAPIOptionsAuthenticationOptions, baseUrl string, ) *ResourcesPostRequestOptions`

NewResourcesPostRequestOptions instantiates a new ResourcesPostRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesPostRequestOptionsWithDefaults

`func NewResourcesPostRequestOptionsWithDefaults() *ResourcesPostRequestOptions`

NewResourcesPostRequestOptionsWithDefaults instantiates a new ResourcesPostRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseOptions

`func (o *ResourcesPostRequestOptions) GetDatabaseOptions() SnowflakeOptionsDatabaseOptions`

GetDatabaseOptions returns the DatabaseOptions field if non-nil, zero value otherwise.

### GetDatabaseOptionsOk

`func (o *ResourcesPostRequestOptions) GetDatabaseOptionsOk() (*SnowflakeOptionsDatabaseOptions, bool)`

GetDatabaseOptionsOk returns a tuple with the DatabaseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseOptions

`func (o *ResourcesPostRequestOptions) SetDatabaseOptions(v SnowflakeOptionsDatabaseOptions)`

SetDatabaseOptions sets DatabaseOptions field to given value.


### GetAccountIdentifier

`func (o *ResourcesPostRequestOptions) GetAccountIdentifier() string`

GetAccountIdentifier returns the AccountIdentifier field if non-nil, zero value otherwise.

### GetAccountIdentifierOk

`func (o *ResourcesPostRequestOptions) GetAccountIdentifierOk() (*string, bool)`

GetAccountIdentifierOk returns a tuple with the AccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifier

`func (o *ResourcesPostRequestOptions) SetAccountIdentifier(v string)`

SetAccountIdentifier sets AccountIdentifier field to given value.


### GetUserRole

`func (o *ResourcesPostRequestOptions) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *ResourcesPostRequestOptions) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *ResourcesPostRequestOptions) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *ResourcesPostRequestOptions) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### GetAuthenticationOptions

`func (o *ResourcesPostRequestOptions) GetAuthenticationOptions() RestAPIOptionsAuthenticationOptions`

GetAuthenticationOptions returns the AuthenticationOptions field if non-nil, zero value otherwise.

### GetAuthenticationOptionsOk

`func (o *ResourcesPostRequestOptions) GetAuthenticationOptionsOk() (*RestAPIOptionsAuthenticationOptions, bool)`

GetAuthenticationOptionsOk returns a tuple with the AuthenticationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOptions

`func (o *ResourcesPostRequestOptions) SetAuthenticationOptions(v RestAPIOptionsAuthenticationOptions)`

SetAuthenticationOptions sets AuthenticationOptions field to given value.


### GetBaseUrl

`func (o *ResourcesPostRequestOptions) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ResourcesPostRequestOptions) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ResourcesPostRequestOptions) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetUrlParameters

`func (o *ResourcesPostRequestOptions) GetUrlParameters() [][]string`

GetUrlParameters returns the UrlParameters field if non-nil, zero value otherwise.

### GetUrlParametersOk

`func (o *ResourcesPostRequestOptions) GetUrlParametersOk() (*[][]string, bool)`

GetUrlParametersOk returns a tuple with the UrlParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlParameters

`func (o *ResourcesPostRequestOptions) SetUrlParameters(v [][]string)`

SetUrlParameters sets UrlParameters field to given value.

### HasUrlParameters

`func (o *ResourcesPostRequestOptions) HasUrlParameters() bool`

HasUrlParameters returns a boolean if a field has been set.

### GetHeaders

`func (o *ResourcesPostRequestOptions) GetHeaders() [][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ResourcesPostRequestOptions) GetHeadersOk() (*[][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ResourcesPostRequestOptions) SetHeaders(v [][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ResourcesPostRequestOptions) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetExtraBodyValues

`func (o *ResourcesPostRequestOptions) GetExtraBodyValues() [][]string`

GetExtraBodyValues returns the ExtraBodyValues field if non-nil, zero value otherwise.

### GetExtraBodyValuesOk

`func (o *ResourcesPostRequestOptions) GetExtraBodyValuesOk() (*[][]string, bool)`

GetExtraBodyValuesOk returns a tuple with the ExtraBodyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraBodyValues

`func (o *ResourcesPostRequestOptions) SetExtraBodyValues(v [][]string)`

SetExtraBodyValues sets ExtraBodyValues field to given value.

### HasExtraBodyValues

`func (o *ResourcesPostRequestOptions) HasExtraBodyValues() bool`

HasExtraBodyValues returns a boolean if a field has been set.

### GetCookiesToForward

`func (o *ResourcesPostRequestOptions) GetCookiesToForward() []string`

GetCookiesToForward returns the CookiesToForward field if non-nil, zero value otherwise.

### GetCookiesToForwardOk

`func (o *ResourcesPostRequestOptions) GetCookiesToForwardOk() (*[]string, bool)`

GetCookiesToForwardOk returns a tuple with the CookiesToForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookiesToForward

`func (o *ResourcesPostRequestOptions) SetCookiesToForward(v []string)`

SetCookiesToForward sets CookiesToForward field to given value.

### HasCookiesToForward

`func (o *ResourcesPostRequestOptions) HasCookiesToForward() bool`

HasCookiesToForward returns a boolean if a field has been set.

### GetForwardAllCookies

`func (o *ResourcesPostRequestOptions) GetForwardAllCookies() bool`

GetForwardAllCookies returns the ForwardAllCookies field if non-nil, zero value otherwise.

### GetForwardAllCookiesOk

`func (o *ResourcesPostRequestOptions) GetForwardAllCookiesOk() (*bool, bool)`

GetForwardAllCookiesOk returns a tuple with the ForwardAllCookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardAllCookies

`func (o *ResourcesPostRequestOptions) SetForwardAllCookies(v bool)`

SetForwardAllCookies sets ForwardAllCookies field to given value.

### HasForwardAllCookies

`func (o *ResourcesPostRequestOptions) HasForwardAllCookies() bool`

HasForwardAllCookies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


