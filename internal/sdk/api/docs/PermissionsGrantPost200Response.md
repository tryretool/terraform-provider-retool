# PermissionsGrantPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**[]PermissionsListObjectsPost200ResponseDataInner**](PermissionsListObjectsPost200ResponseDataInner.md) | An array of requested items | 

## Methods

### NewPermissionsGrantPost200Response

`func NewPermissionsGrantPost200Response(success bool, data []PermissionsListObjectsPost200ResponseDataInner, ) *PermissionsGrantPost200Response`

NewPermissionsGrantPost200Response instantiates a new PermissionsGrantPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsGrantPost200ResponseWithDefaults

`func NewPermissionsGrantPost200ResponseWithDefaults() *PermissionsGrantPost200Response`

NewPermissionsGrantPost200ResponseWithDefaults instantiates a new PermissionsGrantPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *PermissionsGrantPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PermissionsGrantPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PermissionsGrantPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *PermissionsGrantPost200Response) GetData() []PermissionsListObjectsPost200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PermissionsGrantPost200Response) GetDataOk() (*[]PermissionsListObjectsPost200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PermissionsGrantPost200Response) SetData(v []PermissionsListObjectsPost200ResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


