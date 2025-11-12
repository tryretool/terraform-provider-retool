# AppsCloneAppPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** | The app ID. | 
**NewAppName** | **string** | The name of the new app | 
**FolderId** | Pointer to **string** | The id of the folder | [optional] 

## Methods

### NewAppsCloneAppPostRequest

`func NewAppsCloneAppPostRequest(appId string, newAppName string, ) *AppsCloneAppPostRequest`

NewAppsCloneAppPostRequest instantiates a new AppsCloneAppPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsCloneAppPostRequestWithDefaults

`func NewAppsCloneAppPostRequestWithDefaults() *AppsCloneAppPostRequest`

NewAppsCloneAppPostRequestWithDefaults instantiates a new AppsCloneAppPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *AppsCloneAppPostRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AppsCloneAppPostRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AppsCloneAppPostRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetNewAppName

`func (o *AppsCloneAppPostRequest) GetNewAppName() string`

GetNewAppName returns the NewAppName field if non-nil, zero value otherwise.

### GetNewAppNameOk

`func (o *AppsCloneAppPostRequest) GetNewAppNameOk() (*string, bool)`

GetNewAppNameOk returns a tuple with the NewAppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAppName

`func (o *AppsCloneAppPostRequest) SetNewAppName(v string)`

SetNewAppName sets NewAppName field to given value.


### GetFolderId

`func (o *AppsCloneAppPostRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AppsCloneAppPostRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AppsCloneAppPostRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *AppsCloneAppPostRequest) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


