# RolePermissionsRolesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the role. | 
**Description** | Pointer to **string** | The description of the role | [optional] 
**ObjectScopes** | **[]string** | The object scopes of the role | 
**OrganizationScopes** | **[]string** | The organization scopes of the role | 

## Methods

### NewRolePermissionsRolesPostRequest

`func NewRolePermissionsRolesPostRequest(name string, objectScopes []string, organizationScopes []string, ) *RolePermissionsRolesPostRequest`

NewRolePermissionsRolesPostRequest instantiates a new RolePermissionsRolesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionsRolesPostRequestWithDefaults

`func NewRolePermissionsRolesPostRequestWithDefaults() *RolePermissionsRolesPostRequest`

NewRolePermissionsRolesPostRequestWithDefaults instantiates a new RolePermissionsRolesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RolePermissionsRolesPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RolePermissionsRolesPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RolePermissionsRolesPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RolePermissionsRolesPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RolePermissionsRolesPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RolePermissionsRolesPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RolePermissionsRolesPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetObjectScopes

`func (o *RolePermissionsRolesPostRequest) GetObjectScopes() []string`

GetObjectScopes returns the ObjectScopes field if non-nil, zero value otherwise.

### GetObjectScopesOk

`func (o *RolePermissionsRolesPostRequest) GetObjectScopesOk() (*[]string, bool)`

GetObjectScopesOk returns a tuple with the ObjectScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectScopes

`func (o *RolePermissionsRolesPostRequest) SetObjectScopes(v []string)`

SetObjectScopes sets ObjectScopes field to given value.


### GetOrganizationScopes

`func (o *RolePermissionsRolesPostRequest) GetOrganizationScopes() []string`

GetOrganizationScopes returns the OrganizationScopes field if non-nil, zero value otherwise.

### GetOrganizationScopesOk

`func (o *RolePermissionsRolesPostRequest) GetOrganizationScopesOk() (*[]string, bool)`

GetOrganizationScopesOk returns a tuple with the OrganizationScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationScopes

`func (o *RolePermissionsRolesPostRequest) SetOrganizationScopes(v []string)`

SetOrganizationScopes sets OrganizationScopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


