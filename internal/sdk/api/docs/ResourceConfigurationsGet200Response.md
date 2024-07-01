# ResourceConfigurationsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**[]ResourceConfigurationsGet200ResponseDataInner**](ResourceConfigurationsGet200ResponseDataInner.md) | An array of requested items | 
**TotalCount** | **float32** | Total number of items in the response | 
**NextToken** | **NullableString** | A token to retrieve the next page of items in the collection | 
**HasMore** | **bool** | Whether there are more items in the collection | 

## Methods

### NewResourceConfigurationsGet200Response

`func NewResourceConfigurationsGet200Response(success bool, data []ResourceConfigurationsGet200ResponseDataInner, totalCount float32, nextToken NullableString, hasMore bool, ) *ResourceConfigurationsGet200Response`

NewResourceConfigurationsGet200Response instantiates a new ResourceConfigurationsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConfigurationsGet200ResponseWithDefaults

`func NewResourceConfigurationsGet200ResponseWithDefaults() *ResourceConfigurationsGet200Response`

NewResourceConfigurationsGet200ResponseWithDefaults instantiates a new ResourceConfigurationsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ResourceConfigurationsGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResourceConfigurationsGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResourceConfigurationsGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *ResourceConfigurationsGet200Response) GetData() []ResourceConfigurationsGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResourceConfigurationsGet200Response) GetDataOk() (*[]ResourceConfigurationsGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResourceConfigurationsGet200Response) SetData(v []ResourceConfigurationsGet200ResponseDataInner)`

SetData sets Data field to given value.


### GetTotalCount

`func (o *ResourceConfigurationsGet200Response) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ResourceConfigurationsGet200Response) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ResourceConfigurationsGet200Response) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.


### GetNextToken

`func (o *ResourceConfigurationsGet200Response) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ResourceConfigurationsGet200Response) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ResourceConfigurationsGet200Response) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.


### SetNextTokenNil

`func (o *ResourceConfigurationsGet200Response) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *ResourceConfigurationsGet200Response) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil
### GetHasMore

`func (o *ResourceConfigurationsGet200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ResourceConfigurationsGet200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ResourceConfigurationsGet200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


