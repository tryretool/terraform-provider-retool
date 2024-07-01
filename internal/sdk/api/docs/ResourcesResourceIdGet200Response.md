# ResourcesResourceIdGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**ResourcesGet200ResponseDataInner**](ResourcesGet200ResponseDataInner.md) |  | 

## Methods

### NewResourcesResourceIdGet200Response

`func NewResourcesResourceIdGet200Response(success bool, data ResourcesGet200ResponseDataInner, ) *ResourcesResourceIdGet200Response`

NewResourcesResourceIdGet200Response instantiates a new ResourcesResourceIdGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesResourceIdGet200ResponseWithDefaults

`func NewResourcesResourceIdGet200ResponseWithDefaults() *ResourcesResourceIdGet200Response`

NewResourcesResourceIdGet200ResponseWithDefaults instantiates a new ResourcesResourceIdGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ResourcesResourceIdGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResourcesResourceIdGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResourcesResourceIdGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *ResourcesResourceIdGet200Response) GetData() ResourcesGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResourcesResourceIdGet200Response) GetDataOk() (*ResourcesGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResourcesResourceIdGet200Response) SetData(v ResourcesGet200ResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


