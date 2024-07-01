# FoldersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the folder. | 
**ParentFolderId** | Pointer to **NullableString** | The ID of the parent folder. | [optional] 
**FolderType** | **string** | The type of the folder. | 

## Methods

### NewFoldersPostRequest

`func NewFoldersPostRequest(name string, folderType string, ) *FoldersPostRequest`

NewFoldersPostRequest instantiates a new FoldersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoldersPostRequestWithDefaults

`func NewFoldersPostRequestWithDefaults() *FoldersPostRequest`

NewFoldersPostRequestWithDefaults instantiates a new FoldersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FoldersPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FoldersPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FoldersPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetParentFolderId

`func (o *FoldersPostRequest) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *FoldersPostRequest) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *FoldersPostRequest) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.

### HasParentFolderId

`func (o *FoldersPostRequest) HasParentFolderId() bool`

HasParentFolderId returns a boolean if a field has been set.

### SetParentFolderIdNil

`func (o *FoldersPostRequest) SetParentFolderIdNil(b bool)`

 SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil

### UnsetParentFolderId
`func (o *FoldersPostRequest) UnsetParentFolderId()`

UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
### GetFolderType

`func (o *FoldersPostRequest) GetFolderType() string`

GetFolderType returns the FolderType field if non-nil, zero value otherwise.

### GetFolderTypeOk

`func (o *FoldersPostRequest) GetFolderTypeOk() (*string, bool)`

GetFolderTypeOk returns a tuple with the FolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderType

`func (o *FoldersPostRequest) SetFolderType(v string)`

SetFolderType sets FolderType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


