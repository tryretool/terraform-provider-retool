# GRPCOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationOptions** | Pointer to [**GRPCOptionsAuthenticationOptions**](GRPCOptionsAuthenticationOptions.md) |  | [optional] 
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

### NewGRPCOptions

`func NewGRPCOptions() *GRPCOptions`

NewGRPCOptions instantiates a new GRPCOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGRPCOptionsWithDefaults

`func NewGRPCOptionsWithDefaults() *GRPCOptions`

NewGRPCOptionsWithDefaults instantiates a new GRPCOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationOptions

`func (o *GRPCOptions) GetAuthenticationOptions() GRPCOptionsAuthenticationOptions`

GetAuthenticationOptions returns the AuthenticationOptions field if non-nil, zero value otherwise.

### GetAuthenticationOptionsOk

`func (o *GRPCOptions) GetAuthenticationOptionsOk() (*GRPCOptionsAuthenticationOptions, bool)`

GetAuthenticationOptionsOk returns a tuple with the AuthenticationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationOptions

`func (o *GRPCOptions) SetAuthenticationOptions(v GRPCOptionsAuthenticationOptions)`

SetAuthenticationOptions sets AuthenticationOptions field to given value.

### HasAuthenticationOptions

`func (o *GRPCOptions) HasAuthenticationOptions() bool`

HasAuthenticationOptions returns a boolean if a field has been set.

### GetMaxReceiveMessageLength

`func (o *GRPCOptions) GetMaxReceiveMessageLength() string`

GetMaxReceiveMessageLength returns the MaxReceiveMessageLength field if non-nil, zero value otherwise.

### GetMaxReceiveMessageLengthOk

`func (o *GRPCOptions) GetMaxReceiveMessageLengthOk() (*string, bool)`

GetMaxReceiveMessageLengthOk returns a tuple with the MaxReceiveMessageLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxReceiveMessageLength

`func (o *GRPCOptions) SetMaxReceiveMessageLength(v string)`

SetMaxReceiveMessageLength sets MaxReceiveMessageLength field to given value.

### HasMaxReceiveMessageLength

`func (o *GRPCOptions) HasMaxReceiveMessageLength() bool`

HasMaxReceiveMessageLength returns a boolean if a field has been set.

### GetMaxSendMessageLength

`func (o *GRPCOptions) GetMaxSendMessageLength() string`

GetMaxSendMessageLength returns the MaxSendMessageLength field if non-nil, zero value otherwise.

### GetMaxSendMessageLengthOk

`func (o *GRPCOptions) GetMaxSendMessageLengthOk() (*string, bool)`

GetMaxSendMessageLengthOk returns a tuple with the MaxSendMessageLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSendMessageLength

`func (o *GRPCOptions) SetMaxSendMessageLength(v string)`

SetMaxSendMessageLength sets MaxSendMessageLength field to given value.

### HasMaxSendMessageLength

`func (o *GRPCOptions) HasMaxSendMessageLength() bool`

HasMaxSendMessageLength returns a boolean if a field has been set.

### GetMetadata

`func (o *GRPCOptions) GetMetadata() [][]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GRPCOptions) GetMetadataOk() (*[][]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GRPCOptions) SetMetadata(v [][]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GRPCOptions) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOptions

`func (o *GRPCOptions) GetOptions() interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GRPCOptions) GetOptionsOk() (*interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GRPCOptions) SetOptions(v interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GRPCOptions) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *GRPCOptions) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *GRPCOptions) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil
### GetProtoFileUrl

`func (o *GRPCOptions) GetProtoFileUrl() string`

GetProtoFileUrl returns the ProtoFileUrl field if non-nil, zero value otherwise.

### GetProtoFileUrlOk

`func (o *GRPCOptions) GetProtoFileUrlOk() (*string, bool)`

GetProtoFileUrlOk returns a tuple with the ProtoFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoFileUrl

`func (o *GRPCOptions) SetProtoFileUrl(v string)`

SetProtoFileUrl sets ProtoFileUrl field to given value.

### HasProtoFileUrl

`func (o *GRPCOptions) HasProtoFileUrl() bool`

HasProtoFileUrl returns a boolean if a field has been set.

### GetProtoFileUrlHeaders

`func (o *GRPCOptions) GetProtoFileUrlHeaders() [][]string`

GetProtoFileUrlHeaders returns the ProtoFileUrlHeaders field if non-nil, zero value otherwise.

### GetProtoFileUrlHeadersOk

`func (o *GRPCOptions) GetProtoFileUrlHeadersOk() (*[][]string, bool)`

GetProtoFileUrlHeadersOk returns a tuple with the ProtoFileUrlHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoFileUrlHeaders

`func (o *GRPCOptions) SetProtoFileUrlHeaders(v [][]string)`

SetProtoFileUrlHeaders sets ProtoFileUrlHeaders field to given value.

### HasProtoFileUrlHeaders

`func (o *GRPCOptions) HasProtoFileUrlHeaders() bool`

HasProtoFileUrlHeaders returns a boolean if a field has been set.

### SetProtoFileUrlHeadersNil

`func (o *GRPCOptions) SetProtoFileUrlHeadersNil(b bool)`

 SetProtoFileUrlHeadersNil sets the value for ProtoFileUrlHeaders to be an explicit nil

### UnsetProtoFileUrlHeaders
`func (o *GRPCOptions) UnsetProtoFileUrlHeaders()`

UnsetProtoFileUrlHeaders ensures that no value is present for ProtoFileUrlHeaders, not even an explicit nil
### GetProtoSource

`func (o *GRPCOptions) GetProtoSource() string`

GetProtoSource returns the ProtoSource field if non-nil, zero value otherwise.

### GetProtoSourceOk

`func (o *GRPCOptions) GetProtoSourceOk() (*string, bool)`

GetProtoSourceOk returns a tuple with the ProtoSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoSource

`func (o *GRPCOptions) SetProtoSource(v string)`

SetProtoSource sets ProtoSource field to given value.

### HasProtoSource

`func (o *GRPCOptions) HasProtoSource() bool`

HasProtoSource returns a boolean if a field has been set.

### GetServiceName

`func (o *GRPCOptions) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *GRPCOptions) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *GRPCOptions) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *GRPCOptions) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetServiceNames

`func (o *GRPCOptions) GetServiceNames() []string`

GetServiceNames returns the ServiceNames field if non-nil, zero value otherwise.

### GetServiceNamesOk

`func (o *GRPCOptions) GetServiceNamesOk() (*[]string, bool)`

GetServiceNamesOk returns a tuple with the ServiceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceNames

`func (o *GRPCOptions) SetServiceNames(v []string)`

SetServiceNames sets ServiceNames field to given value.

### HasServiceNames

`func (o *GRPCOptions) HasServiceNames() bool`

HasServiceNames returns a boolean if a field has been set.

### GetVerifySessionActionEnabled

`func (o *GRPCOptions) GetVerifySessionActionEnabled() bool`

GetVerifySessionActionEnabled returns the VerifySessionActionEnabled field if non-nil, zero value otherwise.

### GetVerifySessionActionEnabledOk

`func (o *GRPCOptions) GetVerifySessionActionEnabledOk() (*bool, bool)`

GetVerifySessionActionEnabledOk returns a tuple with the VerifySessionActionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySessionActionEnabled

`func (o *GRPCOptions) SetVerifySessionActionEnabled(v bool)`

SetVerifySessionActionEnabled sets VerifySessionActionEnabled field to given value.

### HasVerifySessionActionEnabled

`func (o *GRPCOptions) HasVerifySessionActionEnabled() bool`

HasVerifySessionActionEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


