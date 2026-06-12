# UsageAppDetailsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | Whether the API request was successful. | 
**Data** | [**UsageAppDetailsGet200ResponseData**](UsageAppDetailsGet200ResponseData.md) |  | 
**TotalViewers** | **float32** | Total number of viewer entries across all pages. | 
**TotalEditors** | **float32** | Total number of editor entries across all pages. | 
**TotalCount** | **float32** | Convenience count equal to the larger of total_viewers and total_editors. | 
**NextToken** | **NullableString** | A token to retrieve the next page of items in the collection. | 
**HasMore** | **bool** | Whether there are more items in the collection. | 

## Methods

### NewUsageAppDetailsGet200Response

`func NewUsageAppDetailsGet200Response(success bool, data UsageAppDetailsGet200ResponseData, totalViewers float32, totalEditors float32, totalCount float32, nextToken NullableString, hasMore bool, ) *UsageAppDetailsGet200Response`

NewUsageAppDetailsGet200Response instantiates a new UsageAppDetailsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageAppDetailsGet200ResponseWithDefaults

`func NewUsageAppDetailsGet200ResponseWithDefaults() *UsageAppDetailsGet200Response`

NewUsageAppDetailsGet200ResponseWithDefaults instantiates a new UsageAppDetailsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UsageAppDetailsGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UsageAppDetailsGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UsageAppDetailsGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *UsageAppDetailsGet200Response) GetData() UsageAppDetailsGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsageAppDetailsGet200Response) GetDataOk() (*UsageAppDetailsGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsageAppDetailsGet200Response) SetData(v UsageAppDetailsGet200ResponseData)`

SetData sets Data field to given value.


### GetTotalViewers

`func (o *UsageAppDetailsGet200Response) GetTotalViewers() float32`

GetTotalViewers returns the TotalViewers field if non-nil, zero value otherwise.

### GetTotalViewersOk

`func (o *UsageAppDetailsGet200Response) GetTotalViewersOk() (*float32, bool)`

GetTotalViewersOk returns a tuple with the TotalViewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalViewers

`func (o *UsageAppDetailsGet200Response) SetTotalViewers(v float32)`

SetTotalViewers sets TotalViewers field to given value.


### GetTotalEditors

`func (o *UsageAppDetailsGet200Response) GetTotalEditors() float32`

GetTotalEditors returns the TotalEditors field if non-nil, zero value otherwise.

### GetTotalEditorsOk

`func (o *UsageAppDetailsGet200Response) GetTotalEditorsOk() (*float32, bool)`

GetTotalEditorsOk returns a tuple with the TotalEditors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEditors

`func (o *UsageAppDetailsGet200Response) SetTotalEditors(v float32)`

SetTotalEditors sets TotalEditors field to given value.


### GetTotalCount

`func (o *UsageAppDetailsGet200Response) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *UsageAppDetailsGet200Response) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *UsageAppDetailsGet200Response) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.


### GetNextToken

`func (o *UsageAppDetailsGet200Response) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *UsageAppDetailsGet200Response) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *UsageAppDetailsGet200Response) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.


### SetNextTokenNil

`func (o *UsageAppDetailsGet200Response) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *UsageAppDetailsGet200Response) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil
### GetHasMore

`func (o *UsageAppDetailsGet200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *UsageAppDetailsGet200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *UsageAppDetailsGet200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


