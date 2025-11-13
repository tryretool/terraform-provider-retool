# RolePermissionsRolesPost201Response

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

### NewRolePermissionsRolesPost201Response

`func NewRolePermissionsRolesPost201Response(id string, organizationId string, name string, description string, isDefaultRole bool, isLegacyRole bool, objectScopes []string, organizationScopes []string, createdAt string, updatedAt string, ) *RolePermissionsRolesPost201Response`

NewRolePermissionsRolesPost201Response instantiates a new RolePermissionsRolesPost201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionsRolesPost201ResponseWithDefaults

`func NewRolePermissionsRolesPost201ResponseWithDefaults() *RolePermissionsRolesPost201Response`

NewRolePermissionsRolesPost201ResponseWithDefaults instantiates a new RolePermissionsRolesPost201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RolePermissionsRolesPost201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RolePermissionsRolesPost201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RolePermissionsRolesPost201Response) SetId(v string)`

SetId sets Id field to given value.


### GetOrganizationId

`func (o *RolePermissionsRolesPost201Response) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *RolePermissionsRolesPost201Response) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *RolePermissionsRolesPost201Response) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetName

`func (o *RolePermissionsRolesPost201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RolePermissionsRolesPost201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RolePermissionsRolesPost201Response) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RolePermissionsRolesPost201Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RolePermissionsRolesPost201Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RolePermissionsRolesPost201Response) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsDefaultRole

`func (o *RolePermissionsRolesPost201Response) GetIsDefaultRole() bool`

GetIsDefaultRole returns the IsDefaultRole field if non-nil, zero value otherwise.

### GetIsDefaultRoleOk

`func (o *RolePermissionsRolesPost201Response) GetIsDefaultRoleOk() (*bool, bool)`

GetIsDefaultRoleOk returns a tuple with the IsDefaultRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefaultRole

`func (o *RolePermissionsRolesPost201Response) SetIsDefaultRole(v bool)`

SetIsDefaultRole sets IsDefaultRole field to given value.


### GetIsLegacyRole

`func (o *RolePermissionsRolesPost201Response) GetIsLegacyRole() bool`

GetIsLegacyRole returns the IsLegacyRole field if non-nil, zero value otherwise.

### GetIsLegacyRoleOk

`func (o *RolePermissionsRolesPost201Response) GetIsLegacyRoleOk() (*bool, bool)`

GetIsLegacyRoleOk returns a tuple with the IsLegacyRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLegacyRole

`func (o *RolePermissionsRolesPost201Response) SetIsLegacyRole(v bool)`

SetIsLegacyRole sets IsLegacyRole field to given value.


### GetObjectScopes

`func (o *RolePermissionsRolesPost201Response) GetObjectScopes() []string`

GetObjectScopes returns the ObjectScopes field if non-nil, zero value otherwise.

### GetObjectScopesOk

`func (o *RolePermissionsRolesPost201Response) GetObjectScopesOk() (*[]string, bool)`

GetObjectScopesOk returns a tuple with the ObjectScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectScopes

`func (o *RolePermissionsRolesPost201Response) SetObjectScopes(v []string)`

SetObjectScopes sets ObjectScopes field to given value.


### GetOrganizationScopes

`func (o *RolePermissionsRolesPost201Response) GetOrganizationScopes() []string`

GetOrganizationScopes returns the OrganizationScopes field if non-nil, zero value otherwise.

### GetOrganizationScopesOk

`func (o *RolePermissionsRolesPost201Response) GetOrganizationScopesOk() (*[]string, bool)`

GetOrganizationScopesOk returns a tuple with the OrganizationScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationScopes

`func (o *RolePermissionsRolesPost201Response) SetOrganizationScopes(v []string)`

SetOrganizationScopes sets OrganizationScopes field to given value.


### GetCreatedAt

`func (o *RolePermissionsRolesPost201Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RolePermissionsRolesPost201Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RolePermissionsRolesPost201Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *RolePermissionsRolesPost201Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RolePermissionsRolesPost201Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RolePermissionsRolesPost201Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


