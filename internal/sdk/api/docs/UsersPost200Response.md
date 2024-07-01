# UsersPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**UsersPost200ResponseData**](UsersPost200ResponseData.md) |  | 

## Methods

### NewUsersPost200Response

`func NewUsersPost200Response(success bool, data UsersPost200ResponseData, ) *UsersPost200Response`

NewUsersPost200Response instantiates a new UsersPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersPost200ResponseWithDefaults

`func NewUsersPost200ResponseWithDefaults() *UsersPost200Response`

NewUsersPost200ResponseWithDefaults instantiates a new UsersPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UsersPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UsersPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UsersPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *UsersPost200Response) GetData() UsersPost200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsersPost200Response) GetDataOk() (*UsersPost200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsersPost200Response) SetData(v UsersPost200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


