# Folder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the folder. Currently this is the same as legacy_id but will be different in the future. | 
**LegacyId** | **string** | The legacy id of the folder. | 
**Name** | **string** | The name of the folder | 
**ParentFolderId** | Pointer to **NullableString** | The id of the parent folder | [optional] 
**IsSystemFolder** | **bool** | Whether the folder is a system folder | 
**FolderType** | **string** | The type of the folder | 

## Methods

### NewFolder

`func NewFolder(id string, legacyId string, name string, isSystemFolder bool, folderType string, ) *Folder`

NewFolder instantiates a new Folder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderWithDefaults

`func NewFolderWithDefaults() *Folder`

NewFolderWithDefaults instantiates a new Folder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Folder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Folder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Folder) SetId(v string)`

SetId sets Id field to given value.


### GetLegacyId

`func (o *Folder) GetLegacyId() string`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *Folder) GetLegacyIdOk() (*string, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyId

`func (o *Folder) SetLegacyId(v string)`

SetLegacyId sets LegacyId field to given value.


### GetName

`func (o *Folder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Folder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Folder) SetName(v string)`

SetName sets Name field to given value.


### GetParentFolderId

`func (o *Folder) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *Folder) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *Folder) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *Folder) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *Folder) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *Folder) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetIsSystemFolder

`func (o *Folder) GetIsSystemFolder() bool`

GetIsSystemFolder returns the IsSystemFolder field if non-nil, zero value otherwise.

### GetIsSystemFolderOk

`func (o *Folder) GetIsSystemFolderOk() (*bool, bool)`

GetIsSystemFolderOk returns a tuple with the IsSystemFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemFolder

`func (o *Folder) SetIsSystemFolder(v bool)`

SetIsSystemFolder sets IsSystemFolder field to given value.


### GetFolderType

`func (o *Folder) GetFolderType() string`

GetFolderType returns the FolderType field if non-nil, zero value otherwise.

### GetFolderTypeOk

`func (o *Folder) GetFolderTypeOk() (*string, bool)`

GetFolderTypeOk returns a tuple with the FolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderType

`func (o *Folder) SetFolderType(v string)`

SetFolderType sets FolderType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


