# ObservabilityConfigGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the configuration. | 
**Config** | [**ObservabilityConfigGet200ResponseDataInnerConfig**](ObservabilityConfigGet200ResponseDataInnerConfig.md) |  | 
**Enabled** | **bool** | When enabled, use this provider for apps observability monitoring. | 

## Methods

### NewObservabilityConfigGet200ResponseDataInner

`func NewObservabilityConfigGet200ResponseDataInner(id string, config ObservabilityConfigGet200ResponseDataInnerConfig, enabled bool, ) *ObservabilityConfigGet200ResponseDataInner`

NewObservabilityConfigGet200ResponseDataInner instantiates a new ObservabilityConfigGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservabilityConfigGet200ResponseDataInnerWithDefaults

`func NewObservabilityConfigGet200ResponseDataInnerWithDefaults() *ObservabilityConfigGet200ResponseDataInner`

NewObservabilityConfigGet200ResponseDataInnerWithDefaults instantiates a new ObservabilityConfigGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObservabilityConfigGet200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObservabilityConfigGet200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObservabilityConfigGet200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetConfig

`func (o *ObservabilityConfigGet200ResponseDataInner) GetConfig() ObservabilityConfigGet200ResponseDataInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ObservabilityConfigGet200ResponseDataInner) GetConfigOk() (*ObservabilityConfigGet200ResponseDataInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ObservabilityConfigGet200ResponseDataInner) SetConfig(v ObservabilityConfigGet200ResponseDataInnerConfig)`

SetConfig sets Config field to given value.


### GetEnabled

`func (o *ObservabilityConfigGet200ResponseDataInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ObservabilityConfigGet200ResponseDataInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ObservabilityConfigGet200ResponseDataInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


