# UsageUserDetailsGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**UsageUserDetailsGet200ResponseData**](UsageUserDetailsGet200ResponseData.md) |  | 

## Methods

### NewUsageUserDetailsGet200Response

`func NewUsageUserDetailsGet200Response(success bool, data UsageUserDetailsGet200ResponseData, ) *UsageUserDetailsGet200Response`

NewUsageUserDetailsGet200Response instantiates a new UsageUserDetailsGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageUserDetailsGet200ResponseWithDefaults

`func NewUsageUserDetailsGet200ResponseWithDefaults() *UsageUserDetailsGet200Response`

NewUsageUserDetailsGet200ResponseWithDefaults instantiates a new UsageUserDetailsGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UsageUserDetailsGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UsageUserDetailsGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UsageUserDetailsGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *UsageUserDetailsGet200Response) GetData() UsageUserDetailsGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UsageUserDetailsGet200Response) GetDataOk() (*UsageUserDetailsGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UsageUserDetailsGet200Response) SetData(v UsageUserDetailsGet200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


