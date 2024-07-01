# CustomComponentLibrariesPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**CustomComponentLibrariesLibraryIdGet200ResponseData**](CustomComponentLibrariesLibraryIdGet200ResponseData.md) |  | 

## Methods

### NewCustomComponentLibrariesPost200Response

`func NewCustomComponentLibrariesPost200Response(success bool, data CustomComponentLibrariesLibraryIdGet200ResponseData, ) *CustomComponentLibrariesPost200Response`

NewCustomComponentLibrariesPost200Response instantiates a new CustomComponentLibrariesPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomComponentLibrariesPost200ResponseWithDefaults

`func NewCustomComponentLibrariesPost200ResponseWithDefaults() *CustomComponentLibrariesPost200Response`

NewCustomComponentLibrariesPost200ResponseWithDefaults instantiates a new CustomComponentLibrariesPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CustomComponentLibrariesPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CustomComponentLibrariesPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CustomComponentLibrariesPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *CustomComponentLibrariesPost200Response) GetData() CustomComponentLibrariesLibraryIdGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomComponentLibrariesPost200Response) GetDataOk() (*CustomComponentLibrariesLibraryIdGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomComponentLibrariesPost200Response) SetData(v CustomComponentLibrariesLibraryIdGet200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


