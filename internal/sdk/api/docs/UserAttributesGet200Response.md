# UserAttributesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**[]UserAttributesGet200ResponseDataInner**](UserAttributesGet200ResponseDataInner.md) | An array of requested items | 

## Methods

### NewUserAttributesGet200Response

`func NewUserAttributesGet200Response(success bool, data []UserAttributesGet200ResponseDataInner, ) *UserAttributesGet200Response`

NewUserAttributesGet200Response instantiates a new UserAttributesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAttributesGet200ResponseWithDefaults

`func NewUserAttributesGet200ResponseWithDefaults() *UserAttributesGet200Response`

NewUserAttributesGet200ResponseWithDefaults instantiates a new UserAttributesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UserAttributesGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UserAttributesGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UserAttributesGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *UserAttributesGet200Response) GetData() []UserAttributesGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UserAttributesGet200Response) GetDataOk() (*[]UserAttributesGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UserAttributesGet200Response) SetData(v []UserAttributesGet200ResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


