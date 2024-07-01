# AccessRequestsAccessRequestIdPatch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**AccessRequestsGet200ResponseDataInner**](AccessRequestsGet200ResponseDataInner.md) |  | 

## Methods

### NewAccessRequestsAccessRequestIdPatch200Response

`func NewAccessRequestsAccessRequestIdPatch200Response(success bool, data AccessRequestsGet200ResponseDataInner, ) *AccessRequestsAccessRequestIdPatch200Response`

NewAccessRequestsAccessRequestIdPatch200Response instantiates a new AccessRequestsAccessRequestIdPatch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRequestsAccessRequestIdPatch200ResponseWithDefaults

`func NewAccessRequestsAccessRequestIdPatch200ResponseWithDefaults() *AccessRequestsAccessRequestIdPatch200Response`

NewAccessRequestsAccessRequestIdPatch200ResponseWithDefaults instantiates a new AccessRequestsAccessRequestIdPatch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AccessRequestsAccessRequestIdPatch200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AccessRequestsAccessRequestIdPatch200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AccessRequestsAccessRequestIdPatch200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *AccessRequestsAccessRequestIdPatch200Response) GetData() AccessRequestsGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AccessRequestsAccessRequestIdPatch200Response) GetDataOk() (*AccessRequestsGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AccessRequestsAccessRequestIdPatch200Response) SetData(v AccessRequestsGet200ResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


