# App

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

## Methods

### NewApp

`func NewApp(id string, name string, description NullableString, folderId string, protected bool, synced bool, shortlink NullableString, isModule bool, isMobileApp bool, ) *App`

NewApp instantiates a new App object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWithDefaults

`func NewAppWithDefaults() *App`

NewAppWithDefaults instantiates a new App object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *App) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *App) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *App) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *App) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *App) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *App) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *App) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *App) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *App) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *App) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *App) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFolderId

`func (o *App) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *App) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *App) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.


### GetProtected

`func (o *App) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *App) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *App) SetProtected(v bool)`

SetProtected sets Protected field to given value.


### GetSynced

`func (o *App) GetSynced() bool`

GetSynced returns the Synced field if non-nil, zero value otherwise.

### GetSyncedOk

`func (o *App) GetSyncedOk() (*bool, bool)`

GetSyncedOk returns a tuple with the Synced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynced

`func (o *App) SetSynced(v bool)`

SetSynced sets Synced field to given value.


### GetShortlink

`func (o *App) GetShortlink() string`

GetShortlink returns the Shortlink field if non-nil, zero value otherwise.

### GetShortlinkOk

`func (o *App) GetShortlinkOk() (*string, bool)`

GetShortlinkOk returns a tuple with the Shortlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortlink

`func (o *App) SetShortlink(v string)`

SetShortlink sets Shortlink field to given value.


### SetShortlinkNil

`func (o *App) SetShortlinkNil(b bool)`

 SetShortlinkNil sets the value for Shortlink to be an explicit nil

### UnsetShortlink
`func (o *App) UnsetShortlink()`

UnsetShortlink ensures that no value is present for Shortlink, not even an explicit nil
### GetIsModule

`func (o *App) GetIsModule() bool`

GetIsModule returns the IsModule field if non-nil, zero value otherwise.

### GetIsModuleOk

`func (o *App) GetIsModuleOk() (*bool, bool)`

GetIsModuleOk returns a tuple with the IsModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModule

`func (o *App) SetIsModule(v bool)`

SetIsModule sets IsModule field to given value.


### GetIsMobileApp

`func (o *App) GetIsMobileApp() bool`

GetIsMobileApp returns the IsMobileApp field if non-nil, zero value otherwise.

### GetIsMobileAppOk

`func (o *App) GetIsMobileAppOk() (*bool, bool)`

GetIsMobileAppOk returns a tuple with the IsMobileApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMobileApp

`func (o *App) SetIsMobileApp(v bool)`

SetIsMobileApp sets IsMobileApp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


