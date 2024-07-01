# SnowflakeOptionsAuthenticationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | **string** |  | 
**Username** | **string** |  | 
**Password** | **string** |  | 
**PrivateKey** | **string** |  | 
**PrivateKeyPassphrase** | Pointer to **string** |  | [optional] 
**Oauth2ClientId** | **string** |  | 
**Oauth2ClientSecret** | **string** |  | 
**Oauth2CallbackUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewSnowflakeOptionsAuthenticationOptions

`func NewSnowflakeOptionsAuthenticationOptions(authenticationType string, username string, password string, privateKey string, oauth2ClientId string, oauth2ClientSecret string, ) *SnowflakeOptionsAuthenticationOptions`

NewSnowflakeOptionsAuthenticationOptions instantiates a new SnowflakeOptionsAuthenticationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnowflakeOptionsAuthenticationOptionsWithDefaults

`func NewSnowflakeOptionsAuthenticationOptionsWithDefaults() *SnowflakeOptionsAuthenticationOptions`

NewSnowflakeOptionsAuthenticationOptionsWithDefaults instantiates a new SnowflakeOptionsAuthenticationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *SnowflakeOptionsAuthenticationOptions) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *SnowflakeOptionsAuthenticationOptions) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *SnowflakeOptionsAuthenticationOptions) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetUsername

`func (o *SnowflakeOptionsAuthenticationOptions) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SnowflakeOptionsAuthenticationOptions) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SnowflakeOptionsAuthenticationOptions) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *SnowflakeOptionsAuthenticationOptions) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SnowflakeOptionsAuthenticationOptions) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SnowflakeOptionsAuthenticationOptions) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPrivateKey

`func (o *SnowflakeOptionsAuthenticationOptions) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *SnowflakeOptionsAuthenticationOptions) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *SnowflakeOptionsAuthenticationOptions) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetPrivateKeyPassphrase

`func (o *SnowflakeOptionsAuthenticationOptions) GetPrivateKeyPassphrase() string`

GetPrivateKeyPassphrase returns the PrivateKeyPassphrase field if non-nil, zero value otherwise.

### GetPrivateKeyPassphraseOk

`func (o *SnowflakeOptionsAuthenticationOptions) GetPrivateKeyPassphraseOk() (*string, bool)`

GetPrivateKeyPassphraseOk returns a tuple with the PrivateKeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassphrase

`func (o *SnowflakeOptionsAuthenticationOptions) SetPrivateKeyPassphrase(v string)`

SetPrivateKeyPassphrase sets PrivateKeyPassphrase field to given value.

### HasPrivateKeyPassphrase

`func (o *SnowflakeOptionsAuthenticationOptions) HasPrivateKeyPassphrase() bool`

HasPrivateKeyPassphrase returns a boolean if a field has been set.

### GetOauth2ClientId

`func (o *SnowflakeOptionsAuthenticationOptions) GetOauth2ClientId() string`

GetOauth2ClientId returns the Oauth2ClientId field if non-nil, zero value otherwise.

### GetOauth2ClientIdOk

`func (o *SnowflakeOptionsAuthenticationOptions) GetOauth2ClientIdOk() (*string, bool)`

GetOauth2ClientIdOk returns a tuple with the Oauth2ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ClientId

`func (o *SnowflakeOptionsAuthenticationOptions) SetOauth2ClientId(v string)`

SetOauth2ClientId sets Oauth2ClientId field to given value.


### GetOauth2ClientSecret

`func (o *SnowflakeOptionsAuthenticationOptions) GetOauth2ClientSecret() string`

GetOauth2ClientSecret returns the Oauth2ClientSecret field if non-nil, zero value otherwise.

### GetOauth2ClientSecretOk

`func (o *SnowflakeOptionsAuthenticationOptions) GetOauth2ClientSecretOk() (*string, bool)`

GetOauth2ClientSecretOk returns a tuple with the Oauth2ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ClientSecret

`func (o *SnowflakeOptionsAuthenticationOptions) SetOauth2ClientSecret(v string)`

SetOauth2ClientSecret sets Oauth2ClientSecret field to given value.


### GetOauth2CallbackUrl

`func (o *SnowflakeOptionsAuthenticationOptions) GetOauth2CallbackUrl() string`

GetOauth2CallbackUrl returns the Oauth2CallbackUrl field if non-nil, zero value otherwise.

### GetOauth2CallbackUrlOk

`func (o *SnowflakeOptionsAuthenticationOptions) GetOauth2CallbackUrlOk() (*string, bool)`

GetOauth2CallbackUrlOk returns a tuple with the Oauth2CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2CallbackUrl

`func (o *SnowflakeOptionsAuthenticationOptions) SetOauth2CallbackUrl(v string)`

SetOauth2CallbackUrl sets Oauth2CallbackUrl field to given value.

### HasOauth2CallbackUrl

`func (o *SnowflakeOptionsAuthenticationOptions) HasOauth2CallbackUrl() bool`

HasOauth2CallbackUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


