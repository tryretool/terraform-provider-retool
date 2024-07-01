# ResourceConfigurationsGet200ResponseDataInnerOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseOptions** | [**SnowflakeOptionsDatabaseOptions**](SnowflakeOptionsDatabaseOptions.md) |  | 
**AccountIdentifier** | **string** |  | 
**UserRole** | Pointer to **string** |  | [optional] 
**AuthenticationOptions** | [**RestAPIOptionsAuthenticationOptions**](RestAPIOptionsAuthenticationOptions.md) |  | 
**BaseUrl** | **string** | Use the absolute URL (e.g https://example.com). | 
**UrlParameters** | Pointer to **[][]string** |  | [optional] 
**Headers** | Pointer to **[][]string** |  | [optional] 
**ExtraBodyValues** | Pointer to **[][]string** | Extra body values are not passed for GET or HEAD requests. | [optional] 
**CookiesToForward** | Pointer to **[]string** | You can use the pattern COOKIE_your_cookie_name in the Headers section in order to implement the double-cookie submit pattern. | [optional] 
**ForwardAllCookies** | Pointer to **bool** | This is useful if you have dynamic cookie names. | [optional] 

## Methods

### NewResourceConfigurationsGet200ResponseDataInnerOptions

`func NewResourceConfigurationsGet200ResponseDataInnerOptions(databaseOptions SnowflakeOptionsDatabaseOptions, accountIdentifier string, authenticationOptions RestAPIOptionsAuthenticationOptions, baseUrl string, ) *ResourceConfigurationsGet200ResponseDataInnerOptions`

NewResourceConfigurationsGet200ResponseDataInnerOptions instantiates a new ResourceConfigurationsGet200ResponseDataInnerOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConfigurationsGet200ResponseDataInnerOptionsWithDefaults

`func NewResourceConfigurationsGet200ResponseDataInnerOptionsWithDefaults() *ResourceConfigurationsGet200ResponseDataInnerOptions`

NewResourceConfigurationsGet200ResponseDataInnerOptionsWithDefaults instantiates a new ResourceConfigurationsGet200ResponseDataInnerOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseOptions

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetDatabaseOptions() SnowflakeOptionsDatabaseOptions`

GetDatabaseOptions returns the DatabaseOptions field if non-nil, zero value otherwise.

### GetDatabaseOptionsOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetDatabaseOptionsOk() (*SnowflakeOptionsDatabaseOptions, bool)`

GetDatabaseOptionsOk returns a tuple with the DatabaseOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseOptions

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetDatabaseOptions(v SnowflakeOptionsDatabaseOptions)`

SetDatabaseOptions sets DatabaseOptions field to given value.


### GetAccountIdentifier

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetAccountIdentifier() string`

GetAccountIdentifier returns the AccountIdentifier field if non-nil, zero value otherwise.

### GetAccountIdentifierOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetAccountIdentifierOk() (*string, bool)`

GetAccountIdentifierOk returns a tuple with the AccountIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIdentifier

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetAccountIdentifier(v string)`

SetAccountIdentifier sets AccountIdentifier field to given value.


### GetUserRole

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### GetAuthenticationOptions

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetAuthenticationOptions() RestAPIOptionsAuthenticationOptions`

GetAuthenticationOptions returns the AuthenticationOptions field if non-nil, zero value otherwise.

### GetAuthenticationOptionsOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetAuthenticationOptionsOk() (*RestAPIOptionsAuthenticationOptions, bool)`

GetAuthenticationOptionsOk returns a tuple with the AuthenticationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOptions

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetAuthenticationOptions(v RestAPIOptionsAuthenticationOptions)`

SetAuthenticationOptions sets AuthenticationOptions field to given value.


### GetBaseUrl

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetUrlParameters

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetUrlParameters() [][]string`

GetUrlParameters returns the UrlParameters field if non-nil, zero value otherwise.

### GetUrlParametersOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetUrlParametersOk() (*[][]string, bool)`

GetUrlParametersOk returns a tuple with the UrlParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlParameters

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetUrlParameters(v [][]string)`

SetUrlParameters sets UrlParameters field to given value.

### HasUrlParameters

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasUrlParameters() bool`

HasUrlParameters returns a boolean if a field has been set.

### GetHeaders

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetHeaders() [][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetHeadersOk() (*[][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetHeaders(v [][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetExtraBodyValues

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetExtraBodyValues() [][]string`

GetExtraBodyValues returns the ExtraBodyValues field if non-nil, zero value otherwise.

### GetExtraBodyValuesOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetExtraBodyValuesOk() (*[][]string, bool)`

GetExtraBodyValuesOk returns a tuple with the ExtraBodyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraBodyValues

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetExtraBodyValues(v [][]string)`

SetExtraBodyValues sets ExtraBodyValues field to given value.

### HasExtraBodyValues

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasExtraBodyValues() bool`

HasExtraBodyValues returns a boolean if a field has been set.

### GetCookiesToForward

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetCookiesToForward() []string`

GetCookiesToForward returns the CookiesToForward field if non-nil, zero value otherwise.

### GetCookiesToForwardOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetCookiesToForwardOk() (*[]string, bool)`

GetCookiesToForwardOk returns a tuple with the CookiesToForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookiesToForward

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetCookiesToForward(v []string)`

SetCookiesToForward sets CookiesToForward field to given value.

### HasCookiesToForward

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasCookiesToForward() bool`

HasCookiesToForward returns a boolean if a field has been set.

### GetForwardAllCookies

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetForwardAllCookies() bool`

GetForwardAllCookies returns the ForwardAllCookies field if non-nil, zero value otherwise.

### GetForwardAllCookiesOk

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) GetForwardAllCookiesOk() (*bool, bool)`

GetForwardAllCookiesOk returns a tuple with the ForwardAllCookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardAllCookies

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) SetForwardAllCookies(v bool)`

SetForwardAllCookies sets ForwardAllCookies field to given value.

### HasForwardAllCookies

`func (o *ResourceConfigurationsGet200ResponseDataInnerOptions) HasForwardAllCookies() bool`

HasForwardAllCookies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


