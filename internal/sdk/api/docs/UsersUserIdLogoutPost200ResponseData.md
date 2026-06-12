# UsersUserIdLogoutPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsoLogout** | **bool** | Whether SSO logout was initiated. | 
**SsoLogoutUrl** | Pointer to **string** | URL to redirect the user browser to for IdP session logout. | [optional] 
**Warning** | Pointer to **string** | Warning message if SSO logout could not be performed. | [optional] 

## Methods

### NewUsersUserIdLogoutPost200ResponseData

`func NewUsersUserIdLogoutPost200ResponseData(ssoLogout bool, ) *UsersUserIdLogoutPost200ResponseData`

NewUsersUserIdLogoutPost200ResponseData instantiates a new UsersUserIdLogoutPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersUserIdLogoutPost200ResponseDataWithDefaults

`func NewUsersUserIdLogoutPost200ResponseDataWithDefaults() *UsersUserIdLogoutPost200ResponseData`

NewUsersUserIdLogoutPost200ResponseDataWithDefaults instantiates a new UsersUserIdLogoutPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsoLogout

`func (o *UsersUserIdLogoutPost200ResponseData) GetSsoLogout() bool`

GetSsoLogout returns the SsoLogout field if non-nil, zero value otherwise.

### GetSsoLogoutOk

`func (o *UsersUserIdLogoutPost200ResponseData) GetSsoLogoutOk() (*bool, bool)`

GetSsoLogoutOk returns a tuple with the SsoLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoLogout

`func (o *UsersUserIdLogoutPost200ResponseData) SetSsoLogout(v bool)`

SetSsoLogout sets SsoLogout field to given value.


### GetSsoLogoutUrl

`func (o *UsersUserIdLogoutPost200ResponseData) GetSsoLogoutUrl() string`

GetSsoLogoutUrl returns the SsoLogoutUrl field if non-nil, zero value otherwise.

### GetSsoLogoutUrlOk

`func (o *UsersUserIdLogoutPost200ResponseData) GetSsoLogoutUrlOk() (*string, bool)`

GetSsoLogoutUrlOk returns a tuple with the SsoLogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoLogoutUrl

`func (o *UsersUserIdLogoutPost200ResponseData) SetSsoLogoutUrl(v string)`

SetSsoLogoutUrl sets SsoLogoutUrl field to given value.

### HasSsoLogoutUrl

`func (o *UsersUserIdLogoutPost200ResponseData) HasSsoLogoutUrl() bool`

HasSsoLogoutUrl returns a boolean if a field has been set.

### GetWarning

`func (o *UsersUserIdLogoutPost200ResponseData) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *UsersUserIdLogoutPost200ResponseData) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *UsersUserIdLogoutPost200ResponseData) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *UsersUserIdLogoutPost200ResponseData) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


