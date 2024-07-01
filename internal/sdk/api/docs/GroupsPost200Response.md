# GroupsPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**GroupsPost200ResponseData**](GroupsPost200ResponseData.md) |  | 

## Methods

### NewGroupsPost200Response

`func NewGroupsPost200Response(success bool, data GroupsPost200ResponseData, ) *GroupsPost200Response`

NewGroupsPost200Response instantiates a new GroupsPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsPost200ResponseWithDefaults

`func NewGroupsPost200ResponseWithDefaults() *GroupsPost200Response`

NewGroupsPost200ResponseWithDefaults instantiates a new GroupsPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GroupsPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GroupsPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GroupsPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *GroupsPost200Response) GetData() GroupsPost200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GroupsPost200Response) GetDataOk() (*GroupsPost200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GroupsPost200Response) SetData(v GroupsPost200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


