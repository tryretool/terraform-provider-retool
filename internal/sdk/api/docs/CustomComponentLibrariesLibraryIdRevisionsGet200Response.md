# CustomComponentLibrariesLibraryIdRevisionsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**[]CustomComponentLibrariesLibraryIdRevisionsGet200ResponseDataInner**](CustomComponentLibrariesLibraryIdRevisionsGet200ResponseDataInner.md) | An array of requested items | 
**TotalCount** | **float32** | Total number of items in the response | 
**NextToken** | **NullableString** | A token to retrieve the next page of items in the collection | 
**HasMore** | **bool** | Whether there are more items in the collection | 

## Methods

### NewCustomComponentLibrariesLibraryIdRevisionsGet200Response

`func NewCustomComponentLibrariesLibraryIdRevisionsGet200Response(success bool, data []CustomComponentLibrariesLibraryIdRevisionsGet200ResponseDataInner, totalCount float32, nextToken NullableString, hasMore bool, ) *CustomComponentLibrariesLibraryIdRevisionsGet200Response`

NewCustomComponentLibrariesLibraryIdRevisionsGet200Response instantiates a new CustomComponentLibrariesLibraryIdRevisionsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomComponentLibrariesLibraryIdRevisionsGet200ResponseWithDefaults

`func NewCustomComponentLibrariesLibraryIdRevisionsGet200ResponseWithDefaults() *CustomComponentLibrariesLibraryIdRevisionsGet200Response`

NewCustomComponentLibrariesLibraryIdRevisionsGet200ResponseWithDefaults instantiates a new CustomComponentLibrariesLibraryIdRevisionsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) GetData() []CustomComponentLibrariesLibraryIdRevisionsGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) GetDataOk() (*[]CustomComponentLibrariesLibraryIdRevisionsGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) SetData(v []CustomComponentLibrariesLibraryIdRevisionsGet200ResponseDataInner)`

SetData sets Data field to given value.


### GetTotalCount

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.


### GetNextToken

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.


### SetNextTokenNil

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil
### GetHasMore

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *CustomComponentLibrariesLibraryIdRevisionsGet200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


