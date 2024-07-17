# ObservabilityConfigPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**ObservabilityConfigGet200ResponseDataInnerConfig**](ObservabilityConfigGet200ResponseDataInnerConfig.md) |  | 
**Enabled** | **bool** | When enabled, use this provider for apps observability monitoring. | 

## Methods

### NewObservabilityConfigPostRequest

`func NewObservabilityConfigPostRequest(config ObservabilityConfigGet200ResponseDataInnerConfig, enabled bool, ) *ObservabilityConfigPostRequest`

NewObservabilityConfigPostRequest instantiates a new ObservabilityConfigPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservabilityConfigPostRequestWithDefaults

`func NewObservabilityConfigPostRequestWithDefaults() *ObservabilityConfigPostRequest`

NewObservabilityConfigPostRequestWithDefaults instantiates a new ObservabilityConfigPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ObservabilityConfigPostRequest) GetConfig() ObservabilityConfigGet200ResponseDataInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ObservabilityConfigPostRequest) GetConfigOk() (*ObservabilityConfigGet200ResponseDataInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ObservabilityConfigPostRequest) SetConfig(v ObservabilityConfigGet200ResponseDataInnerConfig)`

SetConfig sets Config field to given value.


### GetEnabled

`func (o *ObservabilityConfigPostRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ObservabilityConfigPostRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ObservabilityConfigPostRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


