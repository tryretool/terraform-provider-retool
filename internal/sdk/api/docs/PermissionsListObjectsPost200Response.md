# PermissionsListObjectsPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**[]PermissionsListObjectsPost200ResponseDataInner**](PermissionsListObjectsPost200ResponseDataInner.md) | An array of requested items | 
**TotalCount** | **float32** | Total number of items in the response | 
**NextToken** | **NullableString** | A token to retrieve the next page of items in the collection | 
**HasMore** | **bool** | Whether there are more items in the collection | 

## Methods

### NewPermissionsListObjectsPost200Response

`func NewPermissionsListObjectsPost200Response(success bool, data []PermissionsListObjectsPost200ResponseDataInner, totalCount float32, nextToken NullableString, hasMore bool, ) *PermissionsListObjectsPost200Response`

NewPermissionsListObjectsPost200Response instantiates a new PermissionsListObjectsPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsListObjectsPost200ResponseWithDefaults

`func NewPermissionsListObjectsPost200ResponseWithDefaults() *PermissionsListObjectsPost200Response`

NewPermissionsListObjectsPost200ResponseWithDefaults instantiates a new PermissionsListObjectsPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *PermissionsListObjectsPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PermissionsListObjectsPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PermissionsListObjectsPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *PermissionsListObjectsPost200Response) GetData() []PermissionsListObjectsPost200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PermissionsListObjectsPost200Response) GetDataOk() (*[]PermissionsListObjectsPost200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PermissionsListObjectsPost200Response) SetData(v []PermissionsListObjectsPost200ResponseDataInner)`

SetData sets Data field to given value.


### GetTotalCount

`func (o *PermissionsListObjectsPost200Response) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PermissionsListObjectsPost200Response) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PermissionsListObjectsPost200Response) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.


### GetNextToken

`func (o *PermissionsListObjectsPost200Response) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *PermissionsListObjectsPost200Response) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *PermissionsListObjectsPost200Response) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.


### SetNextTokenNil

`func (o *PermissionsListObjectsPost200Response) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *PermissionsListObjectsPost200Response) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil
### GetHasMore

`func (o *PermissionsListObjectsPost200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *PermissionsListObjectsPost200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *PermissionsListObjectsPost200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


