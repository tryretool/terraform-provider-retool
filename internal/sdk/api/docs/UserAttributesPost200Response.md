# UserAttributesPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**UserAttributesPost200ResponseData**](UserAttributesPost200ResponseData.md) |  | 

## Methods

### NewUserAttributesPost200Response

`func NewUserAttributesPost200Response(success bool, data UserAttributesPost200ResponseData, ) *UserAttributesPost200Response`

NewUserAttributesPost200Response instantiates a new UserAttributesPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAttributesPost200ResponseWithDefaults

`func NewUserAttributesPost200ResponseWithDefaults() *UserAttributesPost200Response`

NewUserAttributesPost200ResponseWithDefaults instantiates a new UserAttributesPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UserAttributesPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UserAttributesPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UserAttributesPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *UserAttributesPost200Response) GetData() UserAttributesPost200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserAttributesPost200Response) GetDataOk() (*UserAttributesPost200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserAttributesPost200Response) SetData(v UserAttributesPost200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


