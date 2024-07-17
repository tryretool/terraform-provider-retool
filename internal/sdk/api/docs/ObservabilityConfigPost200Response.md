# ObservabilityConfigPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**ObservabilityConfigPost200ResponseData**](ObservabilityConfigPost200ResponseData.md) |  | 

## Methods

### NewObservabilityConfigPost200Response

`func NewObservabilityConfigPost200Response(success bool, data ObservabilityConfigPost200ResponseData, ) *ObservabilityConfigPost200Response`

NewObservabilityConfigPost200Response instantiates a new ObservabilityConfigPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservabilityConfigPost200ResponseWithDefaults

`func NewObservabilityConfigPost200ResponseWithDefaults() *ObservabilityConfigPost200Response`

NewObservabilityConfigPost200ResponseWithDefaults instantiates a new ObservabilityConfigPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ObservabilityConfigPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ObservabilityConfigPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ObservabilityConfigPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *ObservabilityConfigPost200Response) GetData() ObservabilityConfigPost200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ObservabilityConfigPost200Response) GetDataOk() (*ObservabilityConfigPost200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ObservabilityConfigPost200Response) SetData(v ObservabilityConfigPost200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


