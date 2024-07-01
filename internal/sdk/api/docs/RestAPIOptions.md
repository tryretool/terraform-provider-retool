# RestAPIOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | **string** | Use the absolute URL (e.g https://example.com). | 
**UrlParameters** | Pointer to **[][]string** |  | [optional] 
**Headers** | Pointer to **[][]string** |  | [optional] 
**ExtraBodyValues** | Pointer to **[][]string** | Extra body values are not passed for GET or HEAD requests. | [optional] 
**CookiesToForward** | Pointer to **[]string** | You can use the pattern COOKIE_your_cookie_name in the Headers section in order to implement the double-cookie submit pattern. | [optional] 
**ForwardAllCookies** | Pointer to **bool** | This is useful if you have dynamic cookie names. | [optional] 
**AuthenticationOptions** | Pointer to [**RestAPIOptionsAuthenticationOptions**](RestAPIOptionsAuthenticationOptions.md) |  | [optional] 

## Methods

### NewRestAPIOptions

`func NewRestAPIOptions(baseUrl string, ) *RestAPIOptions`

NewRestAPIOptions instantiates a new RestAPIOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestAPIOptionsWithDefaults

`func NewRestAPIOptionsWithDefaults() *RestAPIOptions`

NewRestAPIOptionsWithDefaults instantiates a new RestAPIOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *RestAPIOptions) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *RestAPIOptions) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *RestAPIOptions) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetUrlParameters

`func (o *RestAPIOptions) GetUrlParameters() [][]string`

GetUrlParameters returns the UrlParameters field if non-nil, zero value otherwise.

### GetUrlParametersOk

`func (o *RestAPIOptions) GetUrlParametersOk() (*[][]string, bool)`

GetUrlParametersOk returns a tuple with the UrlParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlParameters

`func (o *RestAPIOptions) SetUrlParameters(v [][]string)`

SetUrlParameters sets UrlParameters field to given value.

### HasUrlParameters

`func (o *RestAPIOptions) HasUrlParameters() bool`

HasUrlParameters returns a boolean if a field has been set.

### GetHeaders

`func (o *RestAPIOptions) GetHeaders() [][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *RestAPIOptions) GetHeadersOk() (*[][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *RestAPIOptions) SetHeaders(v [][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *RestAPIOptions) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetExtraBodyValues

`func (o *RestAPIOptions) GetExtraBodyValues() [][]string`

GetExtraBodyValues returns the ExtraBodyValues field if non-nil, zero value otherwise.

### GetExtraBodyValuesOk

`func (o *RestAPIOptions) GetExtraBodyValuesOk() (*[][]string, bool)`

GetExtraBodyValuesOk returns a tuple with the ExtraBodyValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraBodyValues

`func (o *RestAPIOptions) SetExtraBodyValues(v [][]string)`

SetExtraBodyValues sets ExtraBodyValues field to given value.

### HasExtraBodyValues

`func (o *RestAPIOptions) HasExtraBodyValues() bool`

HasExtraBodyValues returns a boolean if a field has been set.

### GetCookiesToForward

`func (o *RestAPIOptions) GetCookiesToForward() []string`

GetCookiesToForward returns the CookiesToForward field if non-nil, zero value otherwise.

### GetCookiesToForwardOk

`func (o *RestAPIOptions) GetCookiesToForwardOk() (*[]string, bool)`

GetCookiesToForwardOk returns a tuple with the CookiesToForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookiesToForward

`func (o *RestAPIOptions) SetCookiesToForward(v []string)`

SetCookiesToForward sets CookiesToForward field to given value.

### HasCookiesToForward

`func (o *RestAPIOptions) HasCookiesToForward() bool`

HasCookiesToForward returns a boolean if a field has been set.

### GetForwardAllCookies

`func (o *RestAPIOptions) GetForwardAllCookies() bool`

GetForwardAllCookies returns the ForwardAllCookies field if non-nil, zero value otherwise.

### GetForwardAllCookiesOk

`func (o *RestAPIOptions) GetForwardAllCookiesOk() (*bool, bool)`

GetForwardAllCookiesOk returns a tuple with the ForwardAllCookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardAllCookies

`func (o *RestAPIOptions) SetForwardAllCookies(v bool)`

SetForwardAllCookies sets ForwardAllCookies field to given value.

### HasForwardAllCookies

`func (o *RestAPIOptions) HasForwardAllCookies() bool`

HasForwardAllCookies returns a boolean if a field has been set.

### GetAuthenticationOptions

`func (o *RestAPIOptions) GetAuthenticationOptions() RestAPIOptionsAuthenticationOptions`

GetAuthenticationOptions returns the AuthenticationOptions field if non-nil, zero value otherwise.

### GetAuthenticationOptionsOk

`func (o *RestAPIOptions) GetAuthenticationOptionsOk() (*RestAPIOptionsAuthenticationOptions, bool)`

GetAuthenticationOptionsOk returns a tuple with the AuthenticationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOptions

`func (o *RestAPIOptions) SetAuthenticationOptions(v RestAPIOptionsAuthenticationOptions)`

SetAuthenticationOptions sets AuthenticationOptions field to given value.

### HasAuthenticationOptions

`func (o *RestAPIOptions) HasAuthenticationOptions() bool`

HasAuthenticationOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


