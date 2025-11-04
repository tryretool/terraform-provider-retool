# RolePermissionsRolesRoleIdPatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the role. | [optional] 
**Description** | Pointer to **string** | The description of the role | [optional] 
**ObjectScopes** | Pointer to **[]string** | The object scopes of the role | [optional] 
**OrganizationScopes** | Pointer to **[]string** | The organization scopes of the role | [optional] 

## Methods

### NewRolePermissionsRolesRoleIdPatchRequest

`func NewRolePermissionsRolesRoleIdPatchRequest() *RolePermissionsRolesRoleIdPatchRequest`

NewRolePermissionsRolesRoleIdPatchRequest instantiates a new RolePermissionsRolesRoleIdPatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionsRolesRoleIdPatchRequestWithDefaults

`func NewRolePermissionsRolesRoleIdPatchRequestWithDefaults() *RolePermissionsRolesRoleIdPatchRequest`

NewRolePermissionsRolesRoleIdPatchRequestWithDefaults instantiates a new RolePermissionsRolesRoleIdPatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RolePermissionsRolesRoleIdPatchRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RolePermissionsRolesRoleIdPatchRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RolePermissionsRolesRoleIdPatchRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RolePermissionsRolesRoleIdPatchRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RolePermissionsRolesRoleIdPatchRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RolePermissionsRolesRoleIdPatchRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RolePermissionsRolesRoleIdPatchRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RolePermissionsRolesRoleIdPatchRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetObjectScopes

`func (o *RolePermissionsRolesRoleIdPatchRequest) GetObjectScopes() []string`

GetObjectScopes returns the ObjectScopes field if non-nil, zero value otherwise.

### GetObjectScopesOk

`func (o *RolePermissionsRolesRoleIdPatchRequest) GetObjectScopesOk() (*[]string, bool)`

GetObjectScopesOk returns a tuple with the ObjectScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectScopes

`func (o *RolePermissionsRolesRoleIdPatchRequest) SetObjectScopes(v []string)`

SetObjectScopes sets ObjectScopes field to given value.

### HasObjectScopes

`func (o *RolePermissionsRolesRoleIdPatchRequest) HasObjectScopes() bool`

HasObjectScopes returns a boolean if a field has been set.

### GetOrganizationScopes

`func (o *RolePermissionsRolesRoleIdPatchRequest) GetOrganizationScopes() []string`

GetOrganizationScopes returns the OrganizationScopes field if non-nil, zero value otherwise.

### GetOrganizationScopesOk

`func (o *RolePermissionsRolesRoleIdPatchRequest) GetOrganizationScopesOk() (*[]string, bool)`

GetOrganizationScopesOk returns a tuple with the OrganizationScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationScopes

`func (o *RolePermissionsRolesRoleIdPatchRequest) SetOrganizationScopes(v []string)`

SetOrganizationScopes sets OrganizationScopes field to given value.

### HasOrganizationScopes

`func (o *RolePermissionsRolesRoleIdPatchRequest) HasOrganizationScopes() bool`

HasOrganizationScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


