# ObservabilityConfigConfigIdPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**ObservabilityConfigGet200ResponseDataInnerConfig**](ObservabilityConfigGet200ResponseDataInnerConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** | When enabled, use this provider for apps observability monitoring. | [optional] 

## Methods

### NewObservabilityConfigConfigIdPutRequest

`func NewObservabilityConfigConfigIdPutRequest() *ObservabilityConfigConfigIdPutRequest`

NewObservabilityConfigConfigIdPutRequest instantiates a new ObservabilityConfigConfigIdPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservabilityConfigConfigIdPutRequestWithDefaults

`func NewObservabilityConfigConfigIdPutRequestWithDefaults() *ObservabilityConfigConfigIdPutRequest`

NewObservabilityConfigConfigIdPutRequestWithDefaults instantiates a new ObservabilityConfigConfigIdPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ObservabilityConfigConfigIdPutRequest) GetConfig() ObservabilityConfigGet200ResponseDataInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ObservabilityConfigConfigIdPutRequest) GetConfigOk() (*ObservabilityConfigGet200ResponseDataInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ObservabilityConfigConfigIdPutRequest) SetConfig(v ObservabilityConfigGet200ResponseDataInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ObservabilityConfigConfigIdPutRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *ObservabilityConfigConfigIdPutRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ObservabilityConfigConfigIdPutRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ObservabilityConfigConfigIdPutRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ObservabilityConfigConfigIdPutRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


