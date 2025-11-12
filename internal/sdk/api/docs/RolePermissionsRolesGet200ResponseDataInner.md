# RolePermissionsRolesGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the role. | 
**OrganizationId** | **string** | The id of the organization | 
**Name** | **string** | The name of the role | 
**Description** | **string** | The description of the role | 
**IsDefaultRole** | **bool** | Whether the role is the default role | 
**IsLegacyRole** | **bool** | Whether the role is a migrated legacy role | 
**ObjectScopes** | **[]string** | The object scopes of the role | 
**OrganizationScopes** | **[]string** | The organization scopes of the role | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewRolePermissionsRolesGet200ResponseDataInner

`func NewRolePermissionsRolesGet200ResponseDataInner(id string, organizationId string, name string, description string, isDefaultRole bool, isLegacyRole bool, objectScopes []string, organizationScopes []string, createdAt string, updatedAt string, ) *RolePermissionsRolesGet200ResponseDataInner`

NewRolePermissionsRolesGet200ResponseDataInner instantiates a new RolePermissionsRolesGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionsRolesGet200ResponseDataInnerWithDefaults

`func NewRolePermissionsRolesGet200ResponseDataInnerWithDefaults() *RolePermissionsRolesGet200ResponseDataInner`

NewRolePermissionsRolesGet200ResponseDataInnerWithDefaults instantiates a new RolePermissionsRolesGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RolePermissionsRolesGet200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetOrganizationId

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *RolePermissionsRolesGet200ResponseDataInner) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetName

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RolePermissionsRolesGet200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RolePermissionsRolesGet200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsDefaultRole

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetIsDefaultRole() bool`

GetIsDefaultRole returns the IsDefaultRole field if non-nil, zero value otherwise.

### GetIsDefaultRoleOk

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetIsDefaultRoleOk() (*bool, bool)`

GetIsDefaultRoleOk returns a tuple with the IsDefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultRole

`func (o *RolePermissionsRolesGet200ResponseDataInner) SetIsDefaultRole(v bool)`

SetIsDefaultRole sets IsDefaultRole field to given value.


### GetIsLegacyRole

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetIsLegacyRole() bool`

GetIsLegacyRole returns the IsLegacyRole field if non-nil, zero value otherwise.

### GetIsLegacyRoleOk

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetIsLegacyRoleOk() (*bool, bool)`

GetIsLegacyRoleOk returns a tuple with the IsLegacyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLegacyRole

`func (o *RolePermissionsRolesGet200ResponseDataInner) SetIsLegacyRole(v bool)`

SetIsLegacyRole sets IsLegacyRole field to given value.


### GetObjectScopes

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetObjectScopes() []string`

GetObjectScopes returns the ObjectScopes field if non-nil, zero value otherwise.

### GetObjectScopesOk

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetObjectScopesOk() (*[]string, bool)`

GetObjectScopesOk returns a tuple with the ObjectScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectScopes

`func (o *RolePermissionsRolesGet200ResponseDataInner) SetObjectScopes(v []string)`

SetObjectScopes sets ObjectScopes field to given value.


### GetOrganizationScopes

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetOrganizationScopes() []string`

GetOrganizationScopes returns the OrganizationScopes field if non-nil, zero value otherwise.

### GetOrganizationScopesOk

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetOrganizationScopesOk() (*[]string, bool)`

GetOrganizationScopesOk returns a tuple with the OrganizationScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationScopes

`func (o *RolePermissionsRolesGet200ResponseDataInner) SetOrganizationScopes(v []string)`

SetOrganizationScopes sets OrganizationScopes field to given value.


### GetCreatedAt

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RolePermissionsRolesGet200ResponseDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RolePermissionsRolesGet200ResponseDataInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RolePermissionsRolesGet200ResponseDataInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


