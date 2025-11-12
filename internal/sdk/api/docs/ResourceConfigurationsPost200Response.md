# ResourceConfigurationsPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**ResourceConfigurationsPost200ResponseData**](ResourceConfigurationsPost200ResponseData.md) |  | 

## Methods

### NewResourceConfigurationsPost200Response

`func NewResourceConfigurationsPost200Response(success bool, data ResourceConfigurationsPost200ResponseData, ) *ResourceConfigurationsPost200Response`

NewResourceConfigurationsPost200Response instantiates a new ResourceConfigurationsPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConfigurationsPost200ResponseWithDefaults

`func NewResourceConfigurationsPost200ResponseWithDefaults() *ResourceConfigurationsPost200Response`

NewResourceConfigurationsPost200ResponseWithDefaults instantiates a new ResourceConfigurationsPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ResourceConfigurationsPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResourceConfigurationsPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResourceConfigurationsPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *ResourceConfigurationsPost200Response) GetData() ResourceConfigurationsPost200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResourceConfigurationsPost200Response) GetDataOk() (*ResourceConfigurationsPost200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResourceConfigurationsPost200Response) SetData(v ResourceConfigurationsPost200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


