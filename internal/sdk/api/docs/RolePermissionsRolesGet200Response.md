# RolePermissionsRolesGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** | API request succeeded | 
**Data** | [**[]RolePermissionsRolesGet200ResponseDataInner**](RolePermissionsRolesGet200ResponseDataInner.md) | An array of requested items | 

## Methods

### NewRolePermissionsRolesGet200Response

`func NewRolePermissionsRolesGet200Response(success bool, data []RolePermissionsRolesGet200ResponseDataInner, ) *RolePermissionsRolesGet200Response`

NewRolePermissionsRolesGet200Response instantiates a new RolePermissionsRolesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionsRolesGet200ResponseWithDefaults

`func NewRolePermissionsRolesGet200ResponseWithDefaults() *RolePermissionsRolesGet200Response`

NewRolePermissionsRolesGet200ResponseWithDefaults instantiates a new RolePermissionsRolesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RolePermissionsRolesGet200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RolePermissionsRolesGet200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RolePermissionsRolesGet200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetData

`func (o *RolePermissionsRolesGet200Response) GetData() []RolePermissionsRolesGet200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RolePermissionsRolesGet200Response) GetDataOk() (*[]RolePermissionsRolesGet200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RolePermissionsRolesGet200Response) SetData(v []RolePermissionsRolesGet200ResponseDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


