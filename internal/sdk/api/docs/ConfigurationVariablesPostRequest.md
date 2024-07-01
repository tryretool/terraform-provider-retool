# ConfigurationVariablesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the configuration variable | 
**Description** | Pointer to **string** | The description of the configuration variable | [optional] 
**Secret** | **bool** | Whether the configuration variable is a secret | 
**Values** | [**[]ConfigurationVariablesGet200ResponseDataInnerValuesInner**](ConfigurationVariablesGet200ResponseDataInnerValuesInner.md) |  | 

## Methods

### NewConfigurationVariablesPostRequest

`func NewConfigurationVariablesPostRequest(name string, secret bool, values []ConfigurationVariablesGet200ResponseDataInnerValuesInner, ) *ConfigurationVariablesPostRequest`

NewConfigurationVariablesPostRequest instantiates a new ConfigurationVariablesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationVariablesPostRequestWithDefaults

`func NewConfigurationVariablesPostRequestWithDefaults() *ConfigurationVariablesPostRequest`

NewConfigurationVariablesPostRequestWithDefaults instantiates a new ConfigurationVariablesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConfigurationVariablesPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationVariablesPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationVariablesPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ConfigurationVariablesPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigurationVariablesPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigurationVariablesPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigurationVariablesPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSecret

`func (o *ConfigurationVariablesPostRequest) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ConfigurationVariablesPostRequest) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ConfigurationVariablesPostRequest) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### GetValues

`func (o *ConfigurationVariablesPostRequest) GetValues() []ConfigurationVariablesGet200ResponseDataInnerValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ConfigurationVariablesPostRequest) GetValuesOk() (*[]ConfigurationVariablesGet200ResponseDataInnerValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ConfigurationVariablesPostRequest) SetValues(v []ConfigurationVariablesGet200ResponseDataInnerValuesInner)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


