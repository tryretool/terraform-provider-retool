# SpacesPostRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopySsoSettings** | Pointer to **bool** | Copy SSO settings from the admin space. | [optional] 
**CopyBrandingAndThemesSettings** | Pointer to **bool** | Copy Branding and Themes settings from the admin space. | [optional] 
**UsersToCopyAsAdmins** | Pointer to **[]string** | List of emails of users from the admin space that need to be added to the new space as admins. | [optional] 
**CreateAdminUser** | Pointer to **bool** | Create an admin user in the new space for the creator instead of just sending out an invite. | [optional] 

## Methods

### NewSpacesPostRequestOptions

`func NewSpacesPostRequestOptions() *SpacesPostRequestOptions`

NewSpacesPostRequestOptions instantiates a new SpacesPostRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpacesPostRequestOptionsWithDefaults

`func NewSpacesPostRequestOptionsWithDefaults() *SpacesPostRequestOptions`

NewSpacesPostRequestOptionsWithDefaults instantiates a new SpacesPostRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopySsoSettings

`func (o *SpacesPostRequestOptions) GetCopySsoSettings() bool`

GetCopySsoSettings returns the CopySsoSettings field if non-nil, zero value otherwise.

### GetCopySsoSettingsOk

`func (o *SpacesPostRequestOptions) GetCopySsoSettingsOk() (*bool, bool)`

GetCopySsoSettingsOk returns a tuple with the CopySsoSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopySsoSettings

`func (o *SpacesPostRequestOptions) SetCopySsoSettings(v bool)`

SetCopySsoSettings sets CopySsoSettings field to given value.

### HasCopySsoSettings

`func (o *SpacesPostRequestOptions) HasCopySsoSettings() bool`

HasCopySsoSettings returns a boolean if a field has been set.

### GetCopyBrandingAndThemesSettings

`func (o *SpacesPostRequestOptions) GetCopyBrandingAndThemesSettings() bool`

GetCopyBrandingAndThemesSettings returns the CopyBrandingAndThemesSettings field if non-nil, zero value otherwise.

### GetCopyBrandingAndThemesSettingsOk

`func (o *SpacesPostRequestOptions) GetCopyBrandingAndThemesSettingsOk() (*bool, bool)`

GetCopyBrandingAndThemesSettingsOk returns a tuple with the CopyBrandingAndThemesSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyBrandingAndThemesSettings

`func (o *SpacesPostRequestOptions) SetCopyBrandingAndThemesSettings(v bool)`

SetCopyBrandingAndThemesSettings sets CopyBrandingAndThemesSettings field to given value.

### HasCopyBrandingAndThemesSettings

`func (o *SpacesPostRequestOptions) HasCopyBrandingAndThemesSettings() bool`

HasCopyBrandingAndThemesSettings returns a boolean if a field has been set.

### GetUsersToCopyAsAdmins

`func (o *SpacesPostRequestOptions) GetUsersToCopyAsAdmins() []string`

GetUsersToCopyAsAdmins returns the UsersToCopyAsAdmins field if non-nil, zero value otherwise.

### GetUsersToCopyAsAdminsOk

`func (o *SpacesPostRequestOptions) GetUsersToCopyAsAdminsOk() (*[]string, bool)`

GetUsersToCopyAsAdminsOk returns a tuple with the UsersToCopyAsAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersToCopyAsAdmins

`func (o *SpacesPostRequestOptions) SetUsersToCopyAsAdmins(v []string)`

SetUsersToCopyAsAdmins sets UsersToCopyAsAdmins field to given value.

### HasUsersToCopyAsAdmins

`func (o *SpacesPostRequestOptions) HasUsersToCopyAsAdmins() bool`

HasUsersToCopyAsAdmins returns a boolean if a field has been set.

### GetCreateAdminUser

`func (o *SpacesPostRequestOptions) GetCreateAdminUser() bool`

GetCreateAdminUser returns the CreateAdminUser field if non-nil, zero value otherwise.

### GetCreateAdminUserOk

`func (o *SpacesPostRequestOptions) GetCreateAdminUserOk() (*bool, bool)`

GetCreateAdminUserOk returns a tuple with the CreateAdminUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateAdminUser

`func (o *SpacesPostRequestOptions) SetCreateAdminUser(v bool)`

SetCreateAdminUser sets CreateAdminUser field to given value.

### HasCreateAdminUser

`func (o *SpacesPostRequestOptions) HasCreateAdminUser() bool`

HasCreateAdminUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


