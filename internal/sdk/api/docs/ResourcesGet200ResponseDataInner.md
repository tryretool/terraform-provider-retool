# ResourcesGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The uuid or name for the resource. | 
**Type** | **string** | The type of resource. | 
**DisplayName** | **string** |  | 
**FolderId** | Pointer to **NullableString** | The id of the folder this resource belongs to | [optional] 
**Protected** | **bool** | Whether the resource is protected in source control | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewResourcesGet200ResponseDataInner

`func NewResourcesGet200ResponseDataInner(id string, type_ string, displayName string, protected bool, createdAt string, updatedAt string, ) *ResourcesGet200ResponseDataInner`

NewResourcesGet200ResponseDataInner instantiates a new ResourcesGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesGet200ResponseDataInnerWithDefaults

`func NewResourcesGet200ResponseDataInnerWithDefaults() *ResourcesGet200ResponseDataInner`

NewResourcesGet200ResponseDataInnerWithDefaults instantiates a new ResourcesGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourcesGet200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourcesGet200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourcesGet200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ResourcesGet200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourcesGet200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourcesGet200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.


### GetDisplayName

`func (o *ResourcesGet200ResponseDataInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ResourcesGet200ResponseDataInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ResourcesGet200ResponseDataInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetFolderId

`func (o *ResourcesGet200ResponseDataInner) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *ResourcesGet200ResponseDataInner) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *ResourcesGet200ResponseDataInner) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *ResourcesGet200ResponseDataInner) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *ResourcesGet200ResponseDataInner) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *ResourcesGet200ResponseDataInner) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetProtected

`func (o *ResourcesGet200ResponseDataInner) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *ResourcesGet200ResponseDataInner) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *ResourcesGet200ResponseDataInner) SetProtected(v bool)`

SetProtected sets Protected field to given value.


### GetCreatedAt

`func (o *ResourcesGet200ResponseDataInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourcesGet200ResponseDataInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourcesGet200ResponseDataInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ResourcesGet200ResponseDataInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourcesGet200ResponseDataInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourcesGet200ResponseDataInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


