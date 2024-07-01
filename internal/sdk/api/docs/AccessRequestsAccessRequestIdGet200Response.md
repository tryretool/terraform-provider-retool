# AccessRequestsAccessRequestIdGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**AccessRequestsGet200ResponseDataInner**](AccessRequestsGet200ResponseDataInner.md) |  | 

## Methods

### NewAccessRequestsAccessRequestIdGet200Response

`func NewAccessRequestsAccessRequestIdGet200Response(success bool, data AccessRequestsGet200ResponseDataInner, ) *AccessRequestsAccessRequestIdGet200Response`

NewAccessRequestsAccessRequestIdGet200Response instantiates a new AccessRequestsAccessRequestIdGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRequestsAccessRequestIdGet200ResponseWithDefaults

`func NewAccessRequestsAccessRequestIdGet200ResponseWithDefaults() *AccessRequestsAccessRequestIdGet200Response`

NewAccessRequestsAccessRequestIdGet200ResponseWithDefaults instantiates a new AccessRequestsAccessRequestIdGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AccessRequestsAccessRequestIdGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AccessRequestsAccessRequestIdGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AccessRequestsAccessRequestIdGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *AccessRequestsAccessRequestIdGet200Response) GetData() AccessRequestsGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccessRequestsAccessRequestIdGet200Response) GetDataOk() (*AccessRequestsGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccessRequestsAccessRequestIdGet200Response) SetData(v AccessRequestsGet200ResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


