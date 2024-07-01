# ConfigurationVariablesGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the configuration variable | 
**Name** | **string** | The name of the configuration variable | 
**Description** | Pointer to **string** | The description of the configuration variable | [optional] 
**Secret** | **bool** | Whether the configuration variable is a secret | 
**Values** | [**[]ConfigurationVariablesGet200ResponseDataInnerValuesInner**](ConfigurationVariablesGet200ResponseDataInnerValuesInner.md) |  | 

## Methods

### NewConfigurationVariablesGet200ResponseDataInner

`func NewConfigurationVariablesGet200ResponseDataInner(id string, name string, secret bool, values []ConfigurationVariablesGet200ResponseDataInnerValuesInner, ) *ConfigurationVariablesGet200ResponseDataInner`

NewConfigurationVariablesGet200ResponseDataInner instantiates a new ConfigurationVariablesGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationVariablesGet200ResponseDataInnerWithDefaults

`func NewConfigurationVariablesGet200ResponseDataInnerWithDefaults() *ConfigurationVariablesGet200ResponseDataInner`

NewConfigurationVariablesGet200ResponseDataInnerWithDefaults instantiates a new ConfigurationVariablesGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConfigurationVariablesGet200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigurationVariablesGet200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigurationVariablesGet200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ConfigurationVariablesGet200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigurationVariablesGet200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigurationVariablesGet200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ConfigurationVariablesGet200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConfigurationVariablesGet200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConfigurationVariablesGet200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConfigurationVariablesGet200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSecret

`func (o *ConfigurationVariablesGet200ResponseDataInner) GetSecret() bool`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ConfigurationVariablesGet200ResponseDataInner) GetSecretOk() (*bool, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ConfigurationVariablesGet200ResponseDataInner) SetSecret(v bool)`

SetSecret sets Secret field to given value.


### GetValues

`func (o *ConfigurationVariablesGet200ResponseDataInner) GetValues() []ConfigurationVariablesGet200ResponseDataInnerValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ConfigurationVariablesGet200ResponseDataInner) GetValuesOk() (*[]ConfigurationVariablesGet200ResponseDataInnerValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ConfigurationVariablesGet200ResponseDataInner) SetValues(v []ConfigurationVariablesGet200ResponseDataInnerValuesInner)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


