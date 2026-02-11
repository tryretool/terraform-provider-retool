# RestAPIOptionsAuthenticationOptionsAnyOf6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | **string** |  | 
**Oauth2Audience** | Pointer to **string** |  | [optional] 
**Oauth2AccessToken** | Pointer to **string** |  | [optional] 
**Oauth2AccessTokenUrl** | **string** |  | 
**Oauth2AccessTokenLifespanSeconds** | Pointer to **NullableFloat32** |  | [optional] 
**Oauth2AuthUrl** | **string** |  | 
**Oauth2ClientId** | **string** |  | 
**Oauth2ClientSecret** | **string** |  | 
**Oauth2CallbackUrl** | Pointer to **string** |  | [optional] 
**Oauth2IdToken** | Pointer to **string** |  | [optional] 
**Oauth2RefreshToken** | Pointer to **string** |  | [optional] 
**Oauth2Scope** | Pointer to **string** |  | [optional] 
**Oauth2ShareUserCredentials** | Pointer to **bool** |  | [optional] 
**VerifySessionUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewRestAPIOptionsAuthenticationOptionsAnyOf6

`func NewRestAPIOptionsAuthenticationOptionsAnyOf6(authenticationType string, oauth2AccessTokenUrl string, oauth2AuthUrl string, oauth2ClientId string, oauth2ClientSecret string, ) *RestAPIOptionsAuthenticationOptionsAnyOf6`

NewRestAPIOptionsAuthenticationOptionsAnyOf6 instantiates a new RestAPIOptionsAuthenticationOptionsAnyOf6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestAPIOptionsAuthenticationOptionsAnyOf6WithDefaults

`func NewRestAPIOptionsAuthenticationOptionsAnyOf6WithDefaults() *RestAPIOptionsAuthenticationOptionsAnyOf6`

NewRestAPIOptionsAuthenticationOptionsAnyOf6WithDefaults instantiates a new RestAPIOptionsAuthenticationOptionsAnyOf6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetOauth2Audience

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2Audience() string`

GetOauth2Audience returns the Oauth2Audience field if non-nil, zero value otherwise.

### GetOauth2AudienceOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2AudienceOk() (*string, bool)`

GetOauth2AudienceOk returns a tuple with the Oauth2Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Audience

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) SetOauth2Audience(v string)`

SetOauth2Audience sets Oauth2Audience field to given value.

### HasOauth2Audience

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) HasOauth2Audience() bool`

HasOauth2Audience returns a boolean if a field has been set.

### GetOauth2AccessToken

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2AccessToken() string`

GetOauth2AccessToken returns the Oauth2AccessToken field if non-nil, zero value otherwise.

### GetOauth2AccessTokenOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2AccessTokenOk() (*string, bool)`

GetOauth2AccessTokenOk returns a tuple with the Oauth2AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2AccessToken

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) SetOauth2AccessToken(v string)`

SetOauth2AccessToken sets Oauth2AccessToken field to given value.

### HasOauth2AccessToken

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) HasOauth2AccessToken() bool`

HasOauth2AccessToken returns a boolean if a field has been set.

### GetOauth2AccessTokenUrl

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2AccessTokenUrl() string`

GetOauth2AccessTokenUrl returns the Oauth2AccessTokenUrl field if non-nil, zero value otherwise.

### GetOauth2AccessTokenUrlOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2AccessTokenUrlOk() (*string, bool)`

GetOauth2AccessTokenUrlOk returns a tuple with the Oauth2AccessTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2AccessTokenUrl

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) SetOauth2AccessTokenUrl(v string)`

SetOauth2AccessTokenUrl sets Oauth2AccessTokenUrl field to given value.


### GetOauth2AccessTokenLifespanSeconds

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2AccessTokenLifespanSeconds() float32`

GetOauth2AccessTokenLifespanSeconds returns the Oauth2AccessTokenLifespanSeconds field if non-nil, zero value otherwise.

### GetOauth2AccessTokenLifespanSecondsOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2AccessTokenLifespanSecondsOk() (*float32, bool)`

GetOauth2AccessTokenLifespanSecondsOk returns a tuple with the Oauth2AccessTokenLifespanSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2AccessTokenLifespanSeconds

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) SetOauth2AccessTokenLifespanSeconds(v float32)`

SetOauth2AccessTokenLifespanSeconds sets Oauth2AccessTokenLifespanSeconds field to given value.

### HasOauth2AccessTokenLifespanSeconds

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) HasOauth2AccessTokenLifespanSeconds() bool`

HasOauth2AccessTokenLifespanSeconds returns a boolean if a field has been set.

### SetOauth2AccessTokenLifespanSecondsNil

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) SetOauth2AccessTokenLifespanSecondsNil(b bool)`

 SetOauth2AccessTokenLifespanSecondsNil sets the value for Oauth2AccessTokenLifespanSeconds to be an explicit nil

### UnsetOauth2AccessTokenLifespanSeconds
`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) UnsetOauth2AccessTokenLifespanSeconds()`

UnsetOauth2AccessTokenLifespanSeconds ensures that no value is present for Oauth2AccessTokenLifespanSeconds, not even an explicit nil
### GetOauth2AuthUrl

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2AuthUrl() string`

GetOauth2AuthUrl returns the Oauth2AuthUrl field if non-nil, zero value otherwise.

### GetOauth2AuthUrlOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2AuthUrlOk() (*string, bool)`

GetOauth2AuthUrlOk returns a tuple with the Oauth2AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2AuthUrl

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) SetOauth2AuthUrl(v string)`

SetOauth2AuthUrl sets Oauth2AuthUrl field to given value.


### GetOauth2ClientId

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2ClientId() string`

GetOauth2ClientId returns the Oauth2ClientId field if non-nil, zero value otherwise.

### GetOauth2ClientIdOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2ClientIdOk() (*string, bool)`

GetOauth2ClientIdOk returns a tuple with the Oauth2ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ClientId

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) SetOauth2ClientId(v string)`

SetOauth2ClientId sets Oauth2ClientId field to given value.


### GetOauth2ClientSecret

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2ClientSecret() string`

GetOauth2ClientSecret returns the Oauth2ClientSecret field if non-nil, zero value otherwise.

### GetOauth2ClientSecretOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2ClientSecretOk() (*string, bool)`

GetOauth2ClientSecretOk returns a tuple with the Oauth2ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ClientSecret

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) SetOauth2ClientSecret(v string)`

SetOauth2ClientSecret sets Oauth2ClientSecret field to given value.


### GetOauth2CallbackUrl

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2CallbackUrl() string`

GetOauth2CallbackUrl returns the Oauth2CallbackUrl field if non-nil, zero value otherwise.

### GetOauth2CallbackUrlOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2CallbackUrlOk() (*string, bool)`

GetOauth2CallbackUrlOk returns a tuple with the Oauth2CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2CallbackUrl

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) SetOauth2CallbackUrl(v string)`

SetOauth2CallbackUrl sets Oauth2CallbackUrl field to given value.

### HasOauth2CallbackUrl

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) HasOauth2CallbackUrl() bool`

HasOauth2CallbackUrl returns a boolean if a field has been set.

### GetOauth2IdToken

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2IdToken() string`

GetOauth2IdToken returns the Oauth2IdToken field if non-nil, zero value otherwise.

### GetOauth2IdTokenOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2IdTokenOk() (*string, bool)`

GetOauth2IdTokenOk returns a tuple with the Oauth2IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2IdToken

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) SetOauth2IdToken(v string)`

SetOauth2IdToken sets Oauth2IdToken field to given value.

### HasOauth2IdToken

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) HasOauth2IdToken() bool`

HasOauth2IdToken returns a boolean if a field has been set.

### GetOauth2RefreshToken

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2RefreshToken() string`

GetOauth2RefreshToken returns the Oauth2RefreshToken field if non-nil, zero value otherwise.

### GetOauth2RefreshTokenOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2RefreshTokenOk() (*string, bool)`

GetOauth2RefreshTokenOk returns a tuple with the Oauth2RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2RefreshToken

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) SetOauth2RefreshToken(v string)`

SetOauth2RefreshToken sets Oauth2RefreshToken field to given value.

### HasOauth2RefreshToken

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) HasOauth2RefreshToken() bool`

HasOauth2RefreshToken returns a boolean if a field has been set.

### GetOauth2Scope

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2Scope() string`

GetOauth2Scope returns the Oauth2Scope field if non-nil, zero value otherwise.

### GetOauth2ScopeOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2ScopeOk() (*string, bool)`

GetOauth2ScopeOk returns a tuple with the Oauth2Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Scope

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) SetOauth2Scope(v string)`

SetOauth2Scope sets Oauth2Scope field to given value.

### HasOauth2Scope

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) HasOauth2Scope() bool`

HasOauth2Scope returns a boolean if a field has been set.

### GetOauth2ShareUserCredentials

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2ShareUserCredentials() bool`

GetOauth2ShareUserCredentials returns the Oauth2ShareUserCredentials field if non-nil, zero value otherwise.

### GetOauth2ShareUserCredentialsOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetOauth2ShareUserCredentialsOk() (*bool, bool)`

GetOauth2ShareUserCredentialsOk returns a tuple with the Oauth2ShareUserCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ShareUserCredentials

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) SetOauth2ShareUserCredentials(v bool)`

SetOauth2ShareUserCredentials sets Oauth2ShareUserCredentials field to given value.

### HasOauth2ShareUserCredentials

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) HasOauth2ShareUserCredentials() bool`

HasOauth2ShareUserCredentials returns a boolean if a field has been set.

### GetVerifySessionUrl

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetVerifySessionUrl() string`

GetVerifySessionUrl returns the VerifySessionUrl field if non-nil, zero value otherwise.

### GetVerifySessionUrlOk

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) GetVerifySessionUrlOk() (*string, bool)`

GetVerifySessionUrlOk returns a tuple with the VerifySessionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySessionUrl

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) SetVerifySessionUrl(v string)`

SetVerifySessionUrl sets VerifySessionUrl field to given value.

### HasVerifySessionUrl

`func (o *RestAPIOptionsAuthenticationOptionsAnyOf6) HasVerifySessionUrl() bool`

HasVerifySessionUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


