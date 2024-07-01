# FoldersPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**FoldersPost200ResponseData**](FoldersPost200ResponseData.md) |  | 

## Methods

### NewFoldersPost200Response

`func NewFoldersPost200Response(success bool, data FoldersPost200ResponseData, ) *FoldersPost200Response`

NewFoldersPost200Response instantiates a new FoldersPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoldersPost200ResponseWithDefaults

`func NewFoldersPost200ResponseWithDefaults() *FoldersPost200Response`

NewFoldersPost200ResponseWithDefaults instantiates a new FoldersPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *FoldersPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *FoldersPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *FoldersPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *FoldersPost200Response) GetData() FoldersPost200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FoldersPost200Response) GetDataOk() (*FoldersPost200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FoldersPost200Response) SetData(v FoldersPost200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


