# CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**[]CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200ResponseDataInner**](CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200ResponseDataInner.md) | An array of requested items | 
**TotalCount** | **float32** | Total number of items in the response | 
**NextToken** | **NullableString** | A token to retrieve the next page of items in the collection | 
**HasMore** | **bool** | Whether there are more items in the collection | 

## Methods

### NewCustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response

`func NewCustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response(success bool, data []CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200ResponseDataInner, totalCount float32, nextToken NullableString, hasMore bool, ) *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response`

NewCustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response instantiates a new CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200ResponseWithDefaults

`func NewCustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200ResponseWithDefaults() *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response`

NewCustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200ResponseWithDefaults instantiates a new CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) GetData() []CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) GetDataOk() (*[]CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) SetData(v []CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200ResponseDataInner)`

SetData sets Data field to given value.


### GetTotalCount

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.


### GetNextToken

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.


### SetNextTokenNil

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) SetNextTokenNil(b bool)`

 SetNextTokenNil sets the value for NextToken to be an explicit nil

### UnsetNextToken
`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) UnsetNextToken()`

UnsetNextToken ensures that no value is present for NextToken, not even an explicit nil
### GetHasMore

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *CustomComponentLibrariesLibraryIdRevisionsRevisionIdFilesGet200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


