# AppsAppIdGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The app ID. | 
**Name** | **string** | The name of the App | 
**Description** | **NullableString** | The description of the App | 
**FolderId** | **string** | The id of the folder | 
**Protected** | **bool** | Whether the App is protected | 
**Synced** | **bool** | Whether the App is synced | 
**Shortlink** | **NullableString** | The shortlink of the App | 
**IsModule** | **bool** | Whether the App is a module | 
**IsMobileApp** | **bool** | Whether the App is a mobile app | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**ReleaseVersion** | **string** | The live release version of the app. &#39;latest&#39; if releases are disabled or the app has no live release yet | 

## Methods

### NewAppsAppIdGet200ResponseData

`func NewAppsAppIdGet200ResponseData(id string, name string, description NullableString, folderId string, protected bool, synced bool, shortlink NullableString, isModule bool, isMobileApp bool, createdAt string, updatedAt string, releaseVersion string, ) *AppsAppIdGet200ResponseData`

NewAppsAppIdGet200ResponseData instantiates a new AppsAppIdGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsAppIdGet200ResponseDataWithDefaults

`func NewAppsAppIdGet200ResponseDataWithDefaults() *AppsAppIdGet200ResponseData`

NewAppsAppIdGet200ResponseDataWithDefaults instantiates a new AppsAppIdGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppsAppIdGet200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppsAppIdGet200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppsAppIdGet200ResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AppsAppIdGet200ResponseData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppsAppIdGet200ResponseData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppsAppIdGet200ResponseData) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppsAppIdGet200ResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppsAppIdGet200ResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppsAppIdGet200ResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *AppsAppIdGet200ResponseData) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AppsAppIdGet200ResponseData) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFolderId

`func (o *AppsAppIdGet200ResponseData) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AppsAppIdGet200ResponseData) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AppsAppIdGet200ResponseData) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.


### GetProtected

`func (o *AppsAppIdGet200ResponseData) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *AppsAppIdGet200ResponseData) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *AppsAppIdGet200ResponseData) SetProtected(v bool)`

SetProtected sets Protected field to given value.


### GetSynced

`func (o *AppsAppIdGet200ResponseData) GetSynced() bool`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *AppsAppIdGet200ResponseData) GetSyncedOk() (*bool, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *AppsAppIdGet200ResponseData) SetSynced(v bool)`

SetSynced sets Synced field to given value.


### GetShortlink

`func (o *AppsAppIdGet200ResponseData) GetShortlink() string`

GetShortlink returns the Shortlink field if non-nil, zero value otherwise.

### GetShortlinkOk

`func (o *AppsAppIdGet200ResponseData) GetShortlinkOk() (*string, bool)`

GetShortlinkOk returns a tuple with the Shortlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortlink

`func (o *AppsAppIdGet200ResponseData) SetShortlink(v string)`

SetShortlink sets Shortlink field to given value.


### SetShortlinkNil

`func (o *AppsAppIdGet200ResponseData) SetShortlinkNil(b bool)`

 SetShortlinkNil sets the value for Shortlink to be an explicit nil

### UnsetShortlink
`func (o *AppsAppIdGet200ResponseData) UnsetShortlink()`

UnsetShortlink ensures that no value is present for Shortlink, not even an explicit nil
### GetIsModule

`func (o *AppsAppIdGet200ResponseData) GetIsModule() bool`

GetIsModule returns the IsModule field if non-nil, zero value otherwise.

### GetIsModuleOk

`func (o *AppsAppIdGet200ResponseData) GetIsModuleOk() (*bool, bool)`

GetIsModuleOk returns a tuple with the IsModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModule

`func (o *AppsAppIdGet200ResponseData) SetIsModule(v bool)`

SetIsModule sets IsModule field to given value.


### GetIsMobileApp

`func (o *AppsAppIdGet200ResponseData) GetIsMobileApp() bool`

GetIsMobileApp returns the IsMobileApp field if non-nil, zero value otherwise.

### GetIsMobileAppOk

`func (o *AppsAppIdGet200ResponseData) GetIsMobileAppOk() (*bool, bool)`

GetIsMobileAppOk returns a tuple with the IsMobileApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMobileApp

`func (o *AppsAppIdGet200ResponseData) SetIsMobileApp(v bool)`

SetIsMobileApp sets IsMobileApp field to given value.


### GetCreatedAt

`func (o *AppsAppIdGet200ResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AppsAppIdGet200ResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AppsAppIdGet200ResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AppsAppIdGet200ResponseData) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AppsAppIdGet200ResponseData) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AppsAppIdGet200ResponseData) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetReleaseVersion

`func (o *AppsAppIdGet200ResponseData) GetReleaseVersion() string`

GetReleaseVersion returns the ReleaseVersion field if non-nil, zero value otherwise.

### GetReleaseVersionOk

`func (o *AppsAppIdGet200ResponseData) GetReleaseVersionOk() (*string, bool)`

GetReleaseVersionOk returns a tuple with the ReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersion

`func (o *AppsAppIdGet200ResponseData) SetReleaseVersion(v string)`

SetReleaseVersion sets ReleaseVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


