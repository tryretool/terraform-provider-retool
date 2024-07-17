# ObservabilityConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the configuration. | 
**Config** | [**ObservabilityConfigGet200ResponseDataInnerConfig**](ObservabilityConfigGet200ResponseDataInnerConfig.md) |  | 
**Enabled** | **bool** | When enabled, use this provider for apps observability monitoring. | 

## Methods

### NewObservabilityConfiguration

`func NewObservabilityConfiguration(id string, config ObservabilityConfigGet200ResponseDataInnerConfig, enabled bool, ) *ObservabilityConfiguration`

NewObservabilityConfiguration instantiates a new ObservabilityConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservabilityConfigurationWithDefaults

`func NewObservabilityConfigurationWithDefaults() *ObservabilityConfiguration`

NewObservabilityConfigurationWithDefaults instantiates a new ObservabilityConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ObservabilityConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ObservabilityConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ObservabilityConfiguration) SetId(v string)`

SetId sets Id field to given value.


### GetConfig

`func (o *ObservabilityConfiguration) GetConfig() ObservabilityConfigGet200ResponseDataInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ObservabilityConfiguration) GetConfigOk() (*ObservabilityConfigGet200ResponseDataInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ObservabilityConfiguration) SetConfig(v ObservabilityConfigGet200ResponseDataInnerConfig)`

SetConfig sets Config field to given value.


### GetEnabled

`func (o *ObservabilityConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ObservabilityConfiguration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ObservabilityConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


