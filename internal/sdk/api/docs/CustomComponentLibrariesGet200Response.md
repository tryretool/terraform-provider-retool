# CustomComponentLibrariesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**[]CustomComponentLibrariesLibraryIdGet200ResponseData**](CustomComponentLibrariesLibraryIdGet200ResponseData.md) | An array of requested items | 

## Methods

### NewCustomComponentLibrariesGet200Response

`func NewCustomComponentLibrariesGet200Response(success bool, data []CustomComponentLibrariesLibraryIdGet200ResponseData, ) *CustomComponentLibrariesGet200Response`

NewCustomComponentLibrariesGet200Response instantiates a new CustomComponentLibrariesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomComponentLibrariesGet200ResponseWithDefaults

`func NewCustomComponentLibrariesGet200ResponseWithDefaults() *CustomComponentLibrariesGet200Response`

NewCustomComponentLibrariesGet200ResponseWithDefaults instantiates a new CustomComponentLibrariesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CustomComponentLibrariesGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CustomComponentLibrariesGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CustomComponentLibrariesGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *CustomComponentLibrariesGet200Response) GetData() []CustomComponentLibrariesLibraryIdGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomComponentLibrariesGet200Response) GetDataOk() (*[]CustomComponentLibrariesLibraryIdGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomComponentLibrariesGet200Response) SetData(v []CustomComponentLibrariesLibraryIdGet200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


