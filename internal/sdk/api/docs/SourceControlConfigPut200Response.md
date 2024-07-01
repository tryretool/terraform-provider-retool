# SourceControlConfigPut200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**SourceControlConfigPut200ResponseData**](SourceControlConfigPut200ResponseData.md) |  | 

## Methods

### NewSourceControlConfigPut200Response

`func NewSourceControlConfigPut200Response(success bool, data SourceControlConfigPut200ResponseData, ) *SourceControlConfigPut200Response`

NewSourceControlConfigPut200Response instantiates a new SourceControlConfigPut200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlConfigPut200ResponseWithDefaults

`func NewSourceControlConfigPut200ResponseWithDefaults() *SourceControlConfigPut200Response`

NewSourceControlConfigPut200ResponseWithDefaults instantiates a new SourceControlConfigPut200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SourceControlConfigPut200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SourceControlConfigPut200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SourceControlConfigPut200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *SourceControlConfigPut200Response) GetData() SourceControlConfigPut200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SourceControlConfigPut200Response) GetDataOk() (*SourceControlConfigPut200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SourceControlConfigPut200Response) SetData(v SourceControlConfigPut200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


