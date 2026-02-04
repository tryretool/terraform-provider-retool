# ResourcesPostRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseOptions** | [**SnowflakeOptionsDatabaseOptions**](SnowflakeOptionsDatabaseOptions.md) |  | 
**AccountIdentifier** | **string** |  | 
**UserRole** | Pointer to **string** |  | [optional] 
**AuthenticationOptions** | [**GRPCOptionsAuthenticationOptions**](GRPCOptionsAuthenticationOptions.md) |  | 
**BaseUrl** | **string** | Use the absolute URL (e.g. https://example.com). | 
**UrlParameters** | Pointer to **[][]string** |  | [optional] 
**Headers** | Pointer to **[][]string** |  | [optional] 
**ExtraBodyValues** | Pointer to **[][]string** | Extra body values are not passed for GET or HEAD requests. | [optional] 
**CookiesToForward** | Pointer to **[]string** | You can use the pattern COOKIE_your_cookie_name in the Headers section in order to implement the double-cookie submit pattern. | [optional] 
**ForwardAllCookies** | Pointer to **bool** | This is useful if you have dynamic cookie names. | [optional] 
**MaxReceiveMessageLength** | Pointer to **string** |  | [optional] 
**MaxSendMessageLength** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **[][]string** |  | [optional] 
**Options** | Pointer to **interface{}** |  | [optional] 
**ProtoFileUrl** | Pointer to **string** |  | [optional] 
**ProtoFileUrlHeaders** | Pointer to **[][]string** |  | [optional] 
**ProtoSource** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**ServiceNames** | Pointer to **[]string** |  | [optional] 
**VerifySessionActionEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewResourcesPostRequestOptions

`func NewResourcesPostRequestOptions(databaseOptions SnowflakeOptionsDatabaseOptions, accountIdentifier string, authenticationOptions GRPCOptionsAuthenticationOptions, baseUrl string, ) *ResourcesPostRequestOptions`

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

`func (o *ResourcesPostRequestOptions) GetAuthenticationOptions() GRPCOptionsAuthenticationOptions`

GetAuthenticationOptions returns the AuthenticationOptions field if non-nil, zero value otherwise.

### GetAuthenticationOptionsOk

`func (o *ResourcesPostRequestOptions) GetAuthenticationOptionsOk() (*GRPCOptionsAuthenticationOptions, bool)`

GetAuthenticationOptionsOk returns a tuple with the AuthenticationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOptions

`func (o *ResourcesPostRequestOptions) SetAuthenticationOptions(v GRPCOptionsAuthenticationOptions)`

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

### GetMaxReceiveMessageLength

`func (o *ResourcesPostRequestOptions) GetMaxReceiveMessageLength() string`

GetMaxReceiveMessageLength returns the MaxReceiveMessageLength field if non-nil, zero value otherwise.

### GetMaxReceiveMessageLengthOk

`func (o *ResourcesPostRequestOptions) GetMaxReceiveMessageLengthOk() (*string, bool)`

GetMaxReceiveMessageLengthOk returns a tuple with the MaxReceiveMessageLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReceiveMessageLength

`func (o *ResourcesPostRequestOptions) SetMaxReceiveMessageLength(v string)`

SetMaxReceiveMessageLength sets MaxReceiveMessageLength field to given value.

### HasMaxReceiveMessageLength

`func (o *ResourcesPostRequestOptions) HasMaxReceiveMessageLength() bool`

HasMaxReceiveMessageLength returns a boolean if a field has been set.

### GetMaxSendMessageLength

`func (o *ResourcesPostRequestOptions) GetMaxSendMessageLength() string`

GetMaxSendMessageLength returns the MaxSendMessageLength field if non-nil, zero value otherwise.

### GetMaxSendMessageLengthOk

`func (o *ResourcesPostRequestOptions) GetMaxSendMessageLengthOk() (*string, bool)`

GetMaxSendMessageLengthOk returns a tuple with the MaxSendMessageLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSendMessageLength

`func (o *ResourcesPostRequestOptions) SetMaxSendMessageLength(v string)`

SetMaxSendMessageLength sets MaxSendMessageLength field to given value.

### HasMaxSendMessageLength

`func (o *ResourcesPostRequestOptions) HasMaxSendMessageLength() bool`

HasMaxSendMessageLength returns a boolean if a field has been set.

### GetMetadata

`func (o *ResourcesPostRequestOptions) GetMetadata() [][]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourcesPostRequestOptions) GetMetadataOk() (*[][]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourcesPostRequestOptions) SetMetadata(v [][]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResourcesPostRequestOptions) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOptions

`func (o *ResourcesPostRequestOptions) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ResourcesPostRequestOptions) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ResourcesPostRequestOptions) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ResourcesPostRequestOptions) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *ResourcesPostRequestOptions) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *ResourcesPostRequestOptions) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetProtoFileUrl

`func (o *ResourcesPostRequestOptions) GetProtoFileUrl() string`

GetProtoFileUrl returns the ProtoFileUrl field if non-nil, zero value otherwise.

### GetProtoFileUrlOk

`func (o *ResourcesPostRequestOptions) GetProtoFileUrlOk() (*string, bool)`

GetProtoFileUrlOk returns a tuple with the ProtoFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoFileUrl

`func (o *ResourcesPostRequestOptions) SetProtoFileUrl(v string)`

SetProtoFileUrl sets ProtoFileUrl field to given value.

### HasProtoFileUrl

`func (o *ResourcesPostRequestOptions) HasProtoFileUrl() bool`

HasProtoFileUrl returns a boolean if a field has been set.

### GetProtoFileUrlHeaders

`func (o *ResourcesPostRequestOptions) GetProtoFileUrlHeaders() [][]string`

GetProtoFileUrlHeaders returns the ProtoFileUrlHeaders field if non-nil, zero value otherwise.

### GetProtoFileUrlHeadersOk

`func (o *ResourcesPostRequestOptions) GetProtoFileUrlHeadersOk() (*[][]string, bool)`

GetProtoFileUrlHeadersOk returns a tuple with the ProtoFileUrlHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoFileUrlHeaders

`func (o *ResourcesPostRequestOptions) SetProtoFileUrlHeaders(v [][]string)`

SetProtoFileUrlHeaders sets ProtoFileUrlHeaders field to given value.

### HasProtoFileUrlHeaders

`func (o *ResourcesPostRequestOptions) HasProtoFileUrlHeaders() bool`

HasProtoFileUrlHeaders returns a boolean if a field has been set.

### SetProtoFileUrlHeadersNil

`func (o *ResourcesPostRequestOptions) SetProtoFileUrlHeadersNil(b bool)`

 SetProtoFileUrlHeadersNil sets the value for ProtoFileUrlHeaders to be an explicit nil

### UnsetProtoFileUrlHeaders
`func (o *ResourcesPostRequestOptions) UnsetProtoFileUrlHeaders()`

UnsetProtoFileUrlHeaders ensures that no value is present for ProtoFileUrlHeaders, not even an explicit nil
### GetProtoSource

`func (o *ResourcesPostRequestOptions) GetProtoSource() string`

GetProtoSource returns the ProtoSource field if non-nil, zero value otherwise.

### GetProtoSourceOk

`func (o *ResourcesPostRequestOptions) GetProtoSourceOk() (*string, bool)`

GetProtoSourceOk returns a tuple with the ProtoSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoSource

`func (o *ResourcesPostRequestOptions) SetProtoSource(v string)`

SetProtoSource sets ProtoSource field to given value.

### HasProtoSource

`func (o *ResourcesPostRequestOptions) HasProtoSource() bool`

HasProtoSource returns a boolean if a field has been set.

### GetServiceName

`func (o *ResourcesPostRequestOptions) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ResourcesPostRequestOptions) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ResourcesPostRequestOptions) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ResourcesPostRequestOptions) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceNames

`func (o *ResourcesPostRequestOptions) GetServiceNames() []string`

GetServiceNames returns the ServiceNames field if non-nil, zero value otherwise.

### GetServiceNamesOk

`func (o *ResourcesPostRequestOptions) GetServiceNamesOk() (*[]string, bool)`

GetServiceNamesOk returns a tuple with the ServiceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNames

`func (o *ResourcesPostRequestOptions) SetServiceNames(v []string)`

SetServiceNames sets ServiceNames field to given value.

### HasServiceNames

`func (o *ResourcesPostRequestOptions) HasServiceNames() bool`

HasServiceNames returns a boolean if a field has been set.

### GetVerifySessionActionEnabled

`func (o *ResourcesPostRequestOptions) GetVerifySessionActionEnabled() bool`

GetVerifySessionActionEnabled returns the VerifySessionActionEnabled field if non-nil, zero value otherwise.

### GetVerifySessionActionEnabledOk

`func (o *ResourcesPostRequestOptions) GetVerifySessionActionEnabledOk() (*bool, bool)`

GetVerifySessionActionEnabledOk returns a tuple with the VerifySessionActionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySessionActionEnabled

`func (o *ResourcesPostRequestOptions) SetVerifySessionActionEnabled(v bool)`

SetVerifySessionActionEnabled sets VerifySessionActionEnabled field to given value.

### HasVerifySessionActionEnabled

`func (o *ResourcesPostRequestOptions) HasVerifySessionActionEnabled() bool`

HasVerifySessionActionEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


