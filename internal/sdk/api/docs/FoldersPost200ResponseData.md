# FoldersPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the folder. Currently this is the same as legacy_id but will be different in the future. | 
**LegacyId** | **string** | The legacy id of the folder. | 
**Name** | **string** | The name of the folder | 
**ParentFolderId** | Pointer to **NullableString** | The id of the parent folder | [optional] 
**IsSystemFolder** | **bool** | Whether the folder is a system folder | 
**FolderType** | **string** | The type of the folder | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewFoldersPost200ResponseData

`func NewFoldersPost200ResponseData(id string, legacyId string, name string, isSystemFolder bool, folderType string, createdAt string, updatedAt string, ) *FoldersPost200ResponseData`

NewFoldersPost200ResponseData instantiates a new FoldersPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoldersPost200ResponseDataWithDefaults

`func NewFoldersPost200ResponseDataWithDefaults() *FoldersPost200ResponseData`

NewFoldersPost200ResponseDataWithDefaults instantiates a new FoldersPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FoldersPost200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FoldersPost200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FoldersPost200ResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetLegacyId

`func (o *FoldersPost200ResponseData) GetLegacyId() string`

GetLegacyId returns the LegacyId field if non-nil, zero value otherwise.

### GetLegacyIdOk

`func (o *FoldersPost200ResponseData) GetLegacyIdOk() (*string, bool)`

GetLegacyIdOk returns a tuple with the LegacyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyId

`func (o *FoldersPost200ResponseData) SetLegacyId(v string)`

SetLegacyId sets LegacyId field to given value.


### GetName

`func (o *FoldersPost200ResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FoldersPost200ResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FoldersPost200ResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetParentFolderId

`func (o *FoldersPost200ResponseData) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *FoldersPost200ResponseData) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *FoldersPost200ResponseData) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *FoldersPost200ResponseData) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *FoldersPost200ResponseData) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *FoldersPost200ResponseData) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetIsSystemFolder

`func (o *FoldersPost200ResponseData) GetIsSystemFolder() bool`

GetIsSystemFolder returns the IsSystemFolder field if non-nil, zero value otherwise.

### GetIsSystemFolderOk

`func (o *FoldersPost200ResponseData) GetIsSystemFolderOk() (*bool, bool)`

GetIsSystemFolderOk returns a tuple with the IsSystemFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemFolder

`func (o *FoldersPost200ResponseData) SetIsSystemFolder(v bool)`

SetIsSystemFolder sets IsSystemFolder field to given value.


### GetFolderType

`func (o *FoldersPost200ResponseData) GetFolderType() string`

GetFolderType returns the FolderType field if non-nil, zero value otherwise.

### GetFolderTypeOk

`func (o *FoldersPost200ResponseData) GetFolderTypeOk() (*string, bool)`

GetFolderTypeOk returns a tuple with the FolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderType

`func (o *FoldersPost200ResponseData) SetFolderType(v string)`

SetFolderType sets FolderType field to given value.


### GetCreatedAt

`func (o *FoldersPost200ResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FoldersPost200ResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FoldersPost200ResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *FoldersPost200ResponseData) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FoldersPost200ResponseData) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FoldersPost200ResponseData) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


