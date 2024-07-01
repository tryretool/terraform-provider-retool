# SsoConfigPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**SsoConfigPost200ResponseData**](SsoConfigPost200ResponseData.md) |  | 

## Methods

### NewSsoConfigPost200Response

`func NewSsoConfigPost200Response(success bool, data SsoConfigPost200ResponseData, ) *SsoConfigPost200Response`

NewSsoConfigPost200Response instantiates a new SsoConfigPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoConfigPost200ResponseWithDefaults

`func NewSsoConfigPost200ResponseWithDefaults() *SsoConfigPost200Response`

NewSsoConfigPost200ResponseWithDefaults instantiates a new SsoConfigPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SsoConfigPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SsoConfigPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SsoConfigPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *SsoConfigPost200Response) GetData() SsoConfigPost200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SsoConfigPost200Response) GetDataOk() (*SsoConfigPost200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SsoConfigPost200Response) SetData(v SsoConfigPost200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


