# SourceControlDeployPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**SourceControlDeployPost200ResponseData**](SourceControlDeployPost200ResponseData.md) |  | 

## Methods

### NewSourceControlDeployPost200Response

`func NewSourceControlDeployPost200Response(success bool, data SourceControlDeployPost200ResponseData, ) *SourceControlDeployPost200Response`

NewSourceControlDeployPost200Response instantiates a new SourceControlDeployPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlDeployPost200ResponseWithDefaults

`func NewSourceControlDeployPost200ResponseWithDefaults() *SourceControlDeployPost200Response`

NewSourceControlDeployPost200ResponseWithDefaults instantiates a new SourceControlDeployPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SourceControlDeployPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SourceControlDeployPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SourceControlDeployPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *SourceControlDeployPost200Response) GetData() SourceControlDeployPost200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SourceControlDeployPost200Response) GetDataOk() (*SourceControlDeployPost200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SourceControlDeployPost200Response) SetData(v SourceControlDeployPost200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


