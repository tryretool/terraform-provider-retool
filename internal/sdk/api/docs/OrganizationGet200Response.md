# OrganizationGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**OrganizationGet200ResponseData**](OrganizationGet200ResponseData.md) |  | 

## Methods

### NewOrganizationGet200Response

`func NewOrganizationGet200Response(success bool, data OrganizationGet200ResponseData, ) *OrganizationGet200Response`

NewOrganizationGet200Response instantiates a new OrganizationGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationGet200ResponseWithDefaults

`func NewOrganizationGet200ResponseWithDefaults() *OrganizationGet200Response`

NewOrganizationGet200ResponseWithDefaults instantiates a new OrganizationGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *OrganizationGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *OrganizationGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *OrganizationGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *OrganizationGet200Response) GetData() OrganizationGet200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrganizationGet200Response) GetDataOk() (*OrganizationGet200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrganizationGet200Response) SetData(v OrganizationGet200ResponseData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


