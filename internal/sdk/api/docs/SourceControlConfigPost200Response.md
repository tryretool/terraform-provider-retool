# SourceControlConfigPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**SourceControlConfigPost200ResponseData**](SourceControlConfigPost200ResponseData.md) |  | 

## Methods

### NewSourceControlConfigPost200Response

`func NewSourceControlConfigPost200Response(success bool, data SourceControlConfigPost200ResponseData, ) *SourceControlConfigPost200Response`

NewSourceControlConfigPost200Response instantiates a new SourceControlConfigPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlConfigPost200ResponseWithDefaults

`func NewSourceControlConfigPost200ResponseWithDefaults() *SourceControlConfigPost200Response`

NewSourceControlConfigPost200ResponseWithDefaults instantiates a new SourceControlConfigPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SourceControlConfigPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SourceControlConfigPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SourceControlConfigPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *SourceControlConfigPost200Response) GetData() SourceControlConfigPost200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SourceControlConfigPost200Response) GetDataOk() (*SourceControlConfigPost200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SourceControlConfigPost200Response) SetData(v SourceControlConfigPost200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


