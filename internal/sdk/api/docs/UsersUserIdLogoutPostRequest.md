# UsersUserIdLogoutPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsoLogout** | Pointer to **bool** | If true, also initiates an RP-initiated logout with the configured SSO provider. | [optional] [default to false]
**PostLogoutRedirectUri** | Pointer to **string** | URI to redirect the user to after SSO logout completes. Only valid when sso_logout is true. Must be registered in your SSO provider. | [optional] 

## Methods

### NewUsersUserIdLogoutPostRequest

`func NewUsersUserIdLogoutPostRequest() *UsersUserIdLogoutPostRequest`

NewUsersUserIdLogoutPostRequest instantiates a new UsersUserIdLogoutPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersUserIdLogoutPostRequestWithDefaults

`func NewUsersUserIdLogoutPostRequestWithDefaults() *UsersUserIdLogoutPostRequest`

NewUsersUserIdLogoutPostRequestWithDefaults instantiates a new UsersUserIdLogoutPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsoLogout

`func (o *UsersUserIdLogoutPostRequest) GetSsoLogout() bool`

GetSsoLogout returns the SsoLogout field if non-nil, zero value otherwise.

### GetSsoLogoutOk

`func (o *UsersUserIdLogoutPostRequest) GetSsoLogoutOk() (*bool, bool)`

GetSsoLogoutOk returns a tuple with the SsoLogout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoLogout

`func (o *UsersUserIdLogoutPostRequest) SetSsoLogout(v bool)`

SetSsoLogout sets SsoLogout field to given value.

### HasSsoLogout

`func (o *UsersUserIdLogoutPostRequest) HasSsoLogout() bool`

HasSsoLogout returns a boolean if a field has been set.

### GetPostLogoutRedirectUri

`func (o *UsersUserIdLogoutPostRequest) GetPostLogoutRedirectUri() string`

GetPostLogoutRedirectUri returns the PostLogoutRedirectUri field if non-nil, zero value otherwise.

### GetPostLogoutRedirectUriOk

`func (o *UsersUserIdLogoutPostRequest) GetPostLogoutRedirectUriOk() (*string, bool)`

GetPostLogoutRedirectUriOk returns a tuple with the PostLogoutRedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostLogoutRedirectUri

`func (o *UsersUserIdLogoutPostRequest) SetPostLogoutRedirectUri(v string)`

SetPostLogoutRedirectUri sets PostLogoutRedirectUri field to given value.

### HasPostLogoutRedirectUri

`func (o *UsersUserIdLogoutPostRequest) HasPostLogoutRedirectUri() bool`

HasPostLogoutRedirectUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


