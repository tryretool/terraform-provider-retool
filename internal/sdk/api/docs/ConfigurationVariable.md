# ConfigurationVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the configuration variable | 
**Name** | **string** | The name of the configuration variable | 
**Description** | Pointer to **string** | The description of the configuration variable | [optional] 
**Secret** | **bool** | Whether the configuration variable is a secret | 
**Values** | [**[]ConfigurationVariablesGet200ResponseDataInnerValuesInner**](ConfigurationVariablesGet200ResponseDataInnerValuesInner.md) |  | 

## Methods

### NewConfigurationVariable

`func NewConfigurationVariable(id string, name string, secret bool, values []ConfigurationVariablesGet200ResponseDataInnerValuesInner, ) *ConfigurationVariable`

NewConfigurationVariable instantiates a new ConfigurationVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationVariableWithDefaults

`func NewConfigurationVariableWithDefaults() *ConfigurationVariable`

NewConfigurationVariableWithDefaults instantiates a new ConfigurationVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationVariable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationVariable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationVariable) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ConfigurationVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationVariable) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ConfigurationVariable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigurationVariable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigurationVariable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigurationVariable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSecret

`func (o *ConfigurationVariable) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ConfigurationVariable) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ConfigurationVariable) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### GetValues

`func (o *ConfigurationVariable) GetValues() []ConfigurationVariablesGet200ResponseDataInnerValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ConfigurationVariable) GetValuesOk() (*[]ConfigurationVariablesGet200ResponseDataInnerValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ConfigurationVariable) SetValues(v []ConfigurationVariablesGet200ResponseDataInnerValuesInner)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


