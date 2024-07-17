# ResourcesPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**ResourcesPost200ResponseData**](ResourcesPost200ResponseData.md) |  | 

## Methods

### NewResourcesPost200Response

`func NewResourcesPost200Response(success bool, data ResourcesPost200ResponseData, ) *ResourcesPost200Response`

NewResourcesPost200Response instantiates a new ResourcesPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesPost200ResponseWithDefaults

`func NewResourcesPost200ResponseWithDefaults() *ResourcesPost200Response`

NewResourcesPost200ResponseWithDefaults instantiates a new ResourcesPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ResourcesPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResourcesPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResourcesPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *ResourcesPost200Response) GetData() ResourcesPost200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResourcesPost200Response) GetDataOk() (*ResourcesPost200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResourcesPost200Response) SetData(v ResourcesPost200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


