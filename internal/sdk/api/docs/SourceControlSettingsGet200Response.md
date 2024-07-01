# SourceControlSettingsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**SourceControlSettingsGet200ResponseData**](SourceControlSettingsGet200ResponseData.md) |  | 

## Methods

### NewSourceControlSettingsGet200Response

`func NewSourceControlSettingsGet200Response(success bool, data SourceControlSettingsGet200ResponseData, ) *SourceControlSettingsGet200Response`

NewSourceControlSettingsGet200Response instantiates a new SourceControlSettingsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlSettingsGet200ResponseWithDefaults

`func NewSourceControlSettingsGet200ResponseWithDefaults() *SourceControlSettingsGet200Response`

NewSourceControlSettingsGet200ResponseWithDefaults instantiates a new SourceControlSettingsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SourceControlSettingsGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SourceControlSettingsGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SourceControlSettingsGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *SourceControlSettingsGet200Response) GetData() SourceControlSettingsGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SourceControlSettingsGet200Response) GetDataOk() (*SourceControlSettingsGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SourceControlSettingsGet200Response) SetData(v SourceControlSettingsGet200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


