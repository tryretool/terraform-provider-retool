# GRPCOptionsAuthenticationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | **string** |  | 
**Auth0Domain** | **string** |  | 
**Auth0ClientId** | **string** |  | 
**Auth0ClientSecret** | **string** |  | 
**Auth0CustomAudience** | **string** |  | 
**BearerToken** | Pointer to **string** |  | [optional] 
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

### NewGRPCOptionsAuthenticationOptions

`func NewGRPCOptionsAuthenticationOptions(authenticationType string, auth0Domain string, auth0ClientId string, auth0ClientSecret string, auth0CustomAudience string, oauth2AccessTokenUrl string, oauth2AuthUrl string, oauth2ClientId string, oauth2ClientSecret string, ) *GRPCOptionsAuthenticationOptions`

NewGRPCOptionsAuthenticationOptions instantiates a new GRPCOptionsAuthenticationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGRPCOptionsAuthenticationOptionsWithDefaults

`func NewGRPCOptionsAuthenticationOptionsWithDefaults() *GRPCOptionsAuthenticationOptions`

NewGRPCOptionsAuthenticationOptionsWithDefaults instantiates a new GRPCOptionsAuthenticationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *GRPCOptionsAuthenticationOptions) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *GRPCOptionsAuthenticationOptions) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *GRPCOptionsAuthenticationOptions) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetAuth0Domain

`func (o *GRPCOptionsAuthenticationOptions) GetAuth0Domain() string`

GetAuth0Domain returns the Auth0Domain field if non-nil, zero value otherwise.

### GetAuth0DomainOk

`func (o *GRPCOptionsAuthenticationOptions) GetAuth0DomainOk() (*string, bool)`

GetAuth0DomainOk returns a tuple with the Auth0Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth0Domain

`func (o *GRPCOptionsAuthenticationOptions) SetAuth0Domain(v string)`

SetAuth0Domain sets Auth0Domain field to given value.


### GetAuth0ClientId

`func (o *GRPCOptionsAuthenticationOptions) GetAuth0ClientId() string`

GetAuth0ClientId returns the Auth0ClientId field if non-nil, zero value otherwise.

### GetAuth0ClientIdOk

`func (o *GRPCOptionsAuthenticationOptions) GetAuth0ClientIdOk() (*string, bool)`

GetAuth0ClientIdOk returns a tuple with the Auth0ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth0ClientId

`func (o *GRPCOptionsAuthenticationOptions) SetAuth0ClientId(v string)`

SetAuth0ClientId sets Auth0ClientId field to given value.


### GetAuth0ClientSecret

`func (o *GRPCOptionsAuthenticationOptions) GetAuth0ClientSecret() string`

GetAuth0ClientSecret returns the Auth0ClientSecret field if non-nil, zero value otherwise.

### GetAuth0ClientSecretOk

`func (o *GRPCOptionsAuthenticationOptions) GetAuth0ClientSecretOk() (*string, bool)`

GetAuth0ClientSecretOk returns a tuple with the Auth0ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth0ClientSecret

`func (o *GRPCOptionsAuthenticationOptions) SetAuth0ClientSecret(v string)`

SetAuth0ClientSecret sets Auth0ClientSecret field to given value.


### GetAuth0CustomAudience

`func (o *GRPCOptionsAuthenticationOptions) GetAuth0CustomAudience() string`

GetAuth0CustomAudience returns the Auth0CustomAudience field if non-nil, zero value otherwise.

### GetAuth0CustomAudienceOk

`func (o *GRPCOptionsAuthenticationOptions) GetAuth0CustomAudienceOk() (*string, bool)`

GetAuth0CustomAudienceOk returns a tuple with the Auth0CustomAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth0CustomAudience

`func (o *GRPCOptionsAuthenticationOptions) SetAuth0CustomAudience(v string)`

SetAuth0CustomAudience sets Auth0CustomAudience field to given value.


### GetBearerToken

`func (o *GRPCOptionsAuthenticationOptions) GetBearerToken() string`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *GRPCOptionsAuthenticationOptions) GetBearerTokenOk() (*string, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *GRPCOptionsAuthenticationOptions) SetBearerToken(v string)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *GRPCOptionsAuthenticationOptions) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetOauth2Audience

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2Audience() string`

GetOauth2Audience returns the Oauth2Audience field if non-nil, zero value otherwise.

### GetOauth2AudienceOk

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2AudienceOk() (*string, bool)`

GetOauth2AudienceOk returns a tuple with the Oauth2Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Audience

`func (o *GRPCOptionsAuthenticationOptions) SetOauth2Audience(v string)`

SetOauth2Audience sets Oauth2Audience field to given value.

### HasOauth2Audience

`func (o *GRPCOptionsAuthenticationOptions) HasOauth2Audience() bool`

HasOauth2Audience returns a boolean if a field has been set.

### GetOauth2AccessToken

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2AccessToken() string`

GetOauth2AccessToken returns the Oauth2AccessToken field if non-nil, zero value otherwise.

### GetOauth2AccessTokenOk

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2AccessTokenOk() (*string, bool)`

GetOauth2AccessTokenOk returns a tuple with the Oauth2AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2AccessToken

`func (o *GRPCOptionsAuthenticationOptions) SetOauth2AccessToken(v string)`

SetOauth2AccessToken sets Oauth2AccessToken field to given value.

### HasOauth2AccessToken

`func (o *GRPCOptionsAuthenticationOptions) HasOauth2AccessToken() bool`

HasOauth2AccessToken returns a boolean if a field has been set.

### GetOauth2AccessTokenUrl

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2AccessTokenUrl() string`

GetOauth2AccessTokenUrl returns the Oauth2AccessTokenUrl field if non-nil, zero value otherwise.

### GetOauth2AccessTokenUrlOk

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2AccessTokenUrlOk() (*string, bool)`

GetOauth2AccessTokenUrlOk returns a tuple with the Oauth2AccessTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2AccessTokenUrl

`func (o *GRPCOptionsAuthenticationOptions) SetOauth2AccessTokenUrl(v string)`

SetOauth2AccessTokenUrl sets Oauth2AccessTokenUrl field to given value.


### GetOauth2AccessTokenLifespanSeconds

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2AccessTokenLifespanSeconds() float32`

GetOauth2AccessTokenLifespanSeconds returns the Oauth2AccessTokenLifespanSeconds field if non-nil, zero value otherwise.

### GetOauth2AccessTokenLifespanSecondsOk

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2AccessTokenLifespanSecondsOk() (*float32, bool)`

GetOauth2AccessTokenLifespanSecondsOk returns a tuple with the Oauth2AccessTokenLifespanSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2AccessTokenLifespanSeconds

`func (o *GRPCOptionsAuthenticationOptions) SetOauth2AccessTokenLifespanSeconds(v float32)`

SetOauth2AccessTokenLifespanSeconds sets Oauth2AccessTokenLifespanSeconds field to given value.

### HasOauth2AccessTokenLifespanSeconds

`func (o *GRPCOptionsAuthenticationOptions) HasOauth2AccessTokenLifespanSeconds() bool`

HasOauth2AccessTokenLifespanSeconds returns a boolean if a field has been set.

### SetOauth2AccessTokenLifespanSecondsNil

`func (o *GRPCOptionsAuthenticationOptions) SetOauth2AccessTokenLifespanSecondsNil(b bool)`

 SetOauth2AccessTokenLifespanSecondsNil sets the value for Oauth2AccessTokenLifespanSeconds to be an explicit nil

### UnsetOauth2AccessTokenLifespanSeconds
`func (o *GRPCOptionsAuthenticationOptions) UnsetOauth2AccessTokenLifespanSeconds()`

UnsetOauth2AccessTokenLifespanSeconds ensures that no value is present for Oauth2AccessTokenLifespanSeconds, not even an explicit nil
### GetOauth2AuthUrl

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2AuthUrl() string`

GetOauth2AuthUrl returns the Oauth2AuthUrl field if non-nil, zero value otherwise.

### GetOauth2AuthUrlOk

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2AuthUrlOk() (*string, bool)`

GetOauth2AuthUrlOk returns a tuple with the Oauth2AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2AuthUrl

`func (o *GRPCOptionsAuthenticationOptions) SetOauth2AuthUrl(v string)`

SetOauth2AuthUrl sets Oauth2AuthUrl field to given value.


### GetOauth2ClientId

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2ClientId() string`

GetOauth2ClientId returns the Oauth2ClientId field if non-nil, zero value otherwise.

### GetOauth2ClientIdOk

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2ClientIdOk() (*string, bool)`

GetOauth2ClientIdOk returns a tuple with the Oauth2ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ClientId

`func (o *GRPCOptionsAuthenticationOptions) SetOauth2ClientId(v string)`

SetOauth2ClientId sets Oauth2ClientId field to given value.


### GetOauth2ClientSecret

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2ClientSecret() string`

GetOauth2ClientSecret returns the Oauth2ClientSecret field if non-nil, zero value otherwise.

### GetOauth2ClientSecretOk

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2ClientSecretOk() (*string, bool)`

GetOauth2ClientSecretOk returns a tuple with the Oauth2ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ClientSecret

`func (o *GRPCOptionsAuthenticationOptions) SetOauth2ClientSecret(v string)`

SetOauth2ClientSecret sets Oauth2ClientSecret field to given value.


### GetOauth2CallbackUrl

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2CallbackUrl() string`

GetOauth2CallbackUrl returns the Oauth2CallbackUrl field if non-nil, zero value otherwise.

### GetOauth2CallbackUrlOk

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2CallbackUrlOk() (*string, bool)`

GetOauth2CallbackUrlOk returns a tuple with the Oauth2CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2CallbackUrl

`func (o *GRPCOptionsAuthenticationOptions) SetOauth2CallbackUrl(v string)`

SetOauth2CallbackUrl sets Oauth2CallbackUrl field to given value.

### HasOauth2CallbackUrl

`func (o *GRPCOptionsAuthenticationOptions) HasOauth2CallbackUrl() bool`

HasOauth2CallbackUrl returns a boolean if a field has been set.

### GetOauth2IdToken

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2IdToken() string`

GetOauth2IdToken returns the Oauth2IdToken field if non-nil, zero value otherwise.

### GetOauth2IdTokenOk

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2IdTokenOk() (*string, bool)`

GetOauth2IdTokenOk returns a tuple with the Oauth2IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2IdToken

`func (o *GRPCOptionsAuthenticationOptions) SetOauth2IdToken(v string)`

SetOauth2IdToken sets Oauth2IdToken field to given value.

### HasOauth2IdToken

`func (o *GRPCOptionsAuthenticationOptions) HasOauth2IdToken() bool`

HasOauth2IdToken returns a boolean if a field has been set.

### GetOauth2RefreshToken

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2RefreshToken() string`

GetOauth2RefreshToken returns the Oauth2RefreshToken field if non-nil, zero value otherwise.

### GetOauth2RefreshTokenOk

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2RefreshTokenOk() (*string, bool)`

GetOauth2RefreshTokenOk returns a tuple with the Oauth2RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2RefreshToken

`func (o *GRPCOptionsAuthenticationOptions) SetOauth2RefreshToken(v string)`

SetOauth2RefreshToken sets Oauth2RefreshToken field to given value.

### HasOauth2RefreshToken

`func (o *GRPCOptionsAuthenticationOptions) HasOauth2RefreshToken() bool`

HasOauth2RefreshToken returns a boolean if a field has been set.

### GetOauth2Scope

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2Scope() string`

GetOauth2Scope returns the Oauth2Scope field if non-nil, zero value otherwise.

### GetOauth2ScopeOk

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2ScopeOk() (*string, bool)`

GetOauth2ScopeOk returns a tuple with the Oauth2Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Scope

`func (o *GRPCOptionsAuthenticationOptions) SetOauth2Scope(v string)`

SetOauth2Scope sets Oauth2Scope field to given value.

### HasOauth2Scope

`func (o *GRPCOptionsAuthenticationOptions) HasOauth2Scope() bool`

HasOauth2Scope returns a boolean if a field has been set.

### GetOauth2ShareUserCredentials

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2ShareUserCredentials() bool`

GetOauth2ShareUserCredentials returns the Oauth2ShareUserCredentials field if non-nil, zero value otherwise.

### GetOauth2ShareUserCredentialsOk

`func (o *GRPCOptionsAuthenticationOptions) GetOauth2ShareUserCredentialsOk() (*bool, bool)`

GetOauth2ShareUserCredentialsOk returns a tuple with the Oauth2ShareUserCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ShareUserCredentials

`func (o *GRPCOptionsAuthenticationOptions) SetOauth2ShareUserCredentials(v bool)`

SetOauth2ShareUserCredentials sets Oauth2ShareUserCredentials field to given value.

### HasOauth2ShareUserCredentials

`func (o *GRPCOptionsAuthenticationOptions) HasOauth2ShareUserCredentials() bool`

HasOauth2ShareUserCredentials returns a boolean if a field has been set.

### GetVerifySessionUrl

`func (o *GRPCOptionsAuthenticationOptions) GetVerifySessionUrl() string`

GetVerifySessionUrl returns the VerifySessionUrl field if non-nil, zero value otherwise.

### GetVerifySessionUrlOk

`func (o *GRPCOptionsAuthenticationOptions) GetVerifySessionUrlOk() (*string, bool)`

GetVerifySessionUrlOk returns a tuple with the VerifySessionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySessionUrl

`func (o *GRPCOptionsAuthenticationOptions) SetVerifySessionUrl(v string)`

SetVerifySessionUrl sets VerifySessionUrl field to given value.

### HasVerifySessionUrl

`func (o *GRPCOptionsAuthenticationOptions) HasVerifySessionUrl() bool`

HasVerifySessionUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


