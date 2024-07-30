# SsoConfigGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**SsoConfigGet200ResponseData**](SsoConfigGet200ResponseData.md) |  | 

## Methods

### NewSsoConfigGet200Response

`func NewSsoConfigGet200Response(success bool, data SsoConfigGet200ResponseData, ) *SsoConfigGet200Response`

NewSsoConfigGet200Response instantiates a new SsoConfigGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoConfigGet200ResponseWithDefaults

`func NewSsoConfigGet200ResponseWithDefaults() *SsoConfigGet200Response`

NewSsoConfigGet200ResponseWithDefaults instantiates a new SsoConfigGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SsoConfigGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SsoConfigGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SsoConfigGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *SsoConfigGet200Response) GetData() SsoConfigGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SsoConfigGet200Response) GetDataOk() (*SsoConfigGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SsoConfigGet200Response) SetData(v SsoConfigGet200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


