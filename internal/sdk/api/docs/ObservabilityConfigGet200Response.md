# ObservabilityConfigGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**[]ObservabilityConfigGet200ResponseDataInner**](ObservabilityConfigGet200ResponseDataInner.md) |  | 

## Methods

### NewObservabilityConfigGet200Response

`func NewObservabilityConfigGet200Response(success bool, data []ObservabilityConfigGet200ResponseDataInner, ) *ObservabilityConfigGet200Response`

NewObservabilityConfigGet200Response instantiates a new ObservabilityConfigGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObservabilityConfigGet200ResponseWithDefaults

`func NewObservabilityConfigGet200ResponseWithDefaults() *ObservabilityConfigGet200Response`

NewObservabilityConfigGet200ResponseWithDefaults instantiates a new ObservabilityConfigGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ObservabilityConfigGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ObservabilityConfigGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ObservabilityConfigGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *ObservabilityConfigGet200Response) GetData() []ObservabilityConfigGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ObservabilityConfigGet200Response) GetDataOk() (*[]ObservabilityConfigGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ObservabilityConfigGet200Response) SetData(v []ObservabilityConfigGet200ResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


