# ConfigurationVariablesPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**ConfigurationVariablesGet200ResponseDataInner**](ConfigurationVariablesGet200ResponseDataInner.md) |  | 

## Methods

### NewConfigurationVariablesPost200Response

`func NewConfigurationVariablesPost200Response(success bool, data ConfigurationVariablesGet200ResponseDataInner, ) *ConfigurationVariablesPost200Response`

NewConfigurationVariablesPost200Response instantiates a new ConfigurationVariablesPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationVariablesPost200ResponseWithDefaults

`func NewConfigurationVariablesPost200ResponseWithDefaults() *ConfigurationVariablesPost200Response`

NewConfigurationVariablesPost200ResponseWithDefaults instantiates a new ConfigurationVariablesPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ConfigurationVariablesPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ConfigurationVariablesPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ConfigurationVariablesPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *ConfigurationVariablesPost200Response) GetData() ConfigurationVariablesGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConfigurationVariablesPost200Response) GetDataOk() (*ConfigurationVariablesGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConfigurationVariablesPost200Response) SetData(v ConfigurationVariablesGet200ResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


