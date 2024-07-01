# SourceControlConfigGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**SourceControlConfigGet200ResponseData**](SourceControlConfigGet200ResponseData.md) |  | 

## Methods

### NewSourceControlConfigGet200Response

`func NewSourceControlConfigGet200Response(success bool, data SourceControlConfigGet200ResponseData, ) *SourceControlConfigGet200Response`

NewSourceControlConfigGet200Response instantiates a new SourceControlConfigGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlConfigGet200ResponseWithDefaults

`func NewSourceControlConfigGet200ResponseWithDefaults() *SourceControlConfigGet200Response`

NewSourceControlConfigGet200ResponseWithDefaults instantiates a new SourceControlConfigGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SourceControlConfigGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SourceControlConfigGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SourceControlConfigGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *SourceControlConfigGet200Response) GetData() SourceControlConfigGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SourceControlConfigGet200Response) GetDataOk() (*SourceControlConfigGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SourceControlConfigGet200Response) SetData(v SourceControlConfigGet200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


