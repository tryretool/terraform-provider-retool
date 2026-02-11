# AppThemesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**[]AppThemesGet200ResponseDataInner**](AppThemesGet200ResponseDataInner.md) | An array of requested items | 

## Methods

### NewAppThemesGet200Response

`func NewAppThemesGet200Response(success bool, data []AppThemesGet200ResponseDataInner, ) *AppThemesGet200Response`

NewAppThemesGet200Response instantiates a new AppThemesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppThemesGet200ResponseWithDefaults

`func NewAppThemesGet200ResponseWithDefaults() *AppThemesGet200Response`

NewAppThemesGet200ResponseWithDefaults instantiates a new AppThemesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AppThemesGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AppThemesGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AppThemesGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *AppThemesGet200Response) GetData() []AppThemesGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppThemesGet200Response) GetDataOk() (*[]AppThemesGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppThemesGet200Response) SetData(v []AppThemesGet200ResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


