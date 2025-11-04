# AccessTokensPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**AccessTokensPost200ResponseData**](AccessTokensPost200ResponseData.md) |  | 

## Methods

### NewAccessTokensPost200Response

`func NewAccessTokensPost200Response(success bool, data AccessTokensPost200ResponseData, ) *AccessTokensPost200Response`

NewAccessTokensPost200Response instantiates a new AccessTokensPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokensPost200ResponseWithDefaults

`func NewAccessTokensPost200ResponseWithDefaults() *AccessTokensPost200Response`

NewAccessTokensPost200ResponseWithDefaults instantiates a new AccessTokensPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AccessTokensPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AccessTokensPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AccessTokensPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *AccessTokensPost200Response) GetData() AccessTokensPost200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccessTokensPost200Response) GetDataOk() (*AccessTokensPost200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccessTokensPost200Response) SetData(v AccessTokensPost200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


