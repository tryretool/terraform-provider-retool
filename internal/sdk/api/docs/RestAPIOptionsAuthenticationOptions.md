# RestAPIOptionsAuthenticationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | **string** |  | 
**Auth0Domain** | **string** |  | 
**Auth0ClientId** | **string** |  | 
**Auth0ClientSecret** | **string** |  | 
**Auth0CustomAudience** | **string** |  | 
**AmazonAwsRegion** | **string** |  | 
**AmazonServiceName** | **string** |  | 
**AmazonAccessKeyId** | **string** |  | 
**AmazonSecretAccessKey** | **string** |  | 
**AuthWithDefaultCredentialProviderChain** | Pointer to **bool** |  | [optional] 
**EnableAwsv4AuthenticationViaHeaders** | Pointer to **bool** |  | [optional] 
**BasicUsername** | **string** |  | 
**BasicPassword** | Pointer to **string** |  | [optional] 
**BearerToken** | Pointer to **string** |  | [optional] 
**DigestUsername** | **string** |  | 
**DigestPassword** | **string** |  | 
**Oauth1SignatureMethod** | [**RestAPIOptionsAuthenticationOptionsAnyOf5Oauth1SignatureMethod**](RestAPIOptionsAuthenticationOptionsAnyOf5Oauth1SignatureMethod.md) |  | 
**Oauth1ConsumerKey** | **string** |  | 
**Oauth1ConsumerSecret** | **string** |  | 
**Oauth1TokenKey** | **string** |  | 
**Oauth1TokenSecret** | **string** |  | 
**Oauth1RealmParameter** | **string** |  | 
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

### NewRestAPIOptionsAuthenticationOptions

`func NewRestAPIOptionsAuthenticationOptions(authenticationType string, auth0Domain string, auth0ClientId string, auth0ClientSecret string, auth0CustomAudience string, amazonAwsRegion string, amazonServiceName string, amazonAccessKeyId string, amazonSecretAccessKey string, basicUsername string, digestUsername string, digestPassword string, oauth1SignatureMethod RestAPIOptionsAuthenticationOptionsAnyOf5Oauth1SignatureMethod, oauth1ConsumerKey string, oauth1ConsumerSecret string, oauth1TokenKey string, oauth1TokenSecret string, oauth1RealmParameter string, oauth2AccessTokenUrl string, oauth2AuthUrl string, oauth2ClientId string, oauth2ClientSecret string, ) *RestAPIOptionsAuthenticationOptions`

NewRestAPIOptionsAuthenticationOptions instantiates a new RestAPIOptionsAuthenticationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestAPIOptionsAuthenticationOptionsWithDefaults

`func NewRestAPIOptionsAuthenticationOptionsWithDefaults() *RestAPIOptionsAuthenticationOptions`

NewRestAPIOptionsAuthenticationOptionsWithDefaults instantiates a new RestAPIOptionsAuthenticationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *RestAPIOptionsAuthenticationOptions) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *RestAPIOptionsAuthenticationOptions) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *RestAPIOptionsAuthenticationOptions) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetAuth0Domain

`func (o *RestAPIOptionsAuthenticationOptions) GetAuth0Domain() string`

GetAuth0Domain returns the Auth0Domain field if non-nil, zero value otherwise.

### GetAuth0DomainOk

`func (o *RestAPIOptionsAuthenticationOptions) GetAuth0DomainOk() (*string, bool)`

GetAuth0DomainOk returns a tuple with the Auth0Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth0Domain

`func (o *RestAPIOptionsAuthenticationOptions) SetAuth0Domain(v string)`

SetAuth0Domain sets Auth0Domain field to given value.


### GetAuth0ClientId

`func (o *RestAPIOptionsAuthenticationOptions) GetAuth0ClientId() string`

GetAuth0ClientId returns the Auth0ClientId field if non-nil, zero value otherwise.

### GetAuth0ClientIdOk

`func (o *RestAPIOptionsAuthenticationOptions) GetAuth0ClientIdOk() (*string, bool)`

GetAuth0ClientIdOk returns a tuple with the Auth0ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth0ClientId

`func (o *RestAPIOptionsAuthenticationOptions) SetAuth0ClientId(v string)`

SetAuth0ClientId sets Auth0ClientId field to given value.


### GetAuth0ClientSecret

`func (o *RestAPIOptionsAuthenticationOptions) GetAuth0ClientSecret() string`

GetAuth0ClientSecret returns the Auth0ClientSecret field if non-nil, zero value otherwise.

### GetAuth0ClientSecretOk

`func (o *RestAPIOptionsAuthenticationOptions) GetAuth0ClientSecretOk() (*string, bool)`

GetAuth0ClientSecretOk returns a tuple with the Auth0ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth0ClientSecret

`func (o *RestAPIOptionsAuthenticationOptions) SetAuth0ClientSecret(v string)`

SetAuth0ClientSecret sets Auth0ClientSecret field to given value.


### GetAuth0CustomAudience

`func (o *RestAPIOptionsAuthenticationOptions) GetAuth0CustomAudience() string`

GetAuth0CustomAudience returns the Auth0CustomAudience field if non-nil, zero value otherwise.

### GetAuth0CustomAudienceOk

`func (o *RestAPIOptionsAuthenticationOptions) GetAuth0CustomAudienceOk() (*string, bool)`

GetAuth0CustomAudienceOk returns a tuple with the Auth0CustomAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth0CustomAudience

`func (o *RestAPIOptionsAuthenticationOptions) SetAuth0CustomAudience(v string)`

SetAuth0CustomAudience sets Auth0CustomAudience field to given value.


### GetAmazonAwsRegion

`func (o *RestAPIOptionsAuthenticationOptions) GetAmazonAwsRegion() string`

GetAmazonAwsRegion returns the AmazonAwsRegion field if non-nil, zero value otherwise.

### GetAmazonAwsRegionOk

`func (o *RestAPIOptionsAuthenticationOptions) GetAmazonAwsRegionOk() (*string, bool)`

GetAmazonAwsRegionOk returns a tuple with the AmazonAwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAwsRegion

`func (o *RestAPIOptionsAuthenticationOptions) SetAmazonAwsRegion(v string)`

SetAmazonAwsRegion sets AmazonAwsRegion field to given value.


### GetAmazonServiceName

`func (o *RestAPIOptionsAuthenticationOptions) GetAmazonServiceName() string`

GetAmazonServiceName returns the AmazonServiceName field if non-nil, zero value otherwise.

### GetAmazonServiceNameOk

`func (o *RestAPIOptionsAuthenticationOptions) GetAmazonServiceNameOk() (*string, bool)`

GetAmazonServiceNameOk returns a tuple with the AmazonServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonServiceName

`func (o *RestAPIOptionsAuthenticationOptions) SetAmazonServiceName(v string)`

SetAmazonServiceName sets AmazonServiceName field to given value.


### GetAmazonAccessKeyId

`func (o *RestAPIOptionsAuthenticationOptions) GetAmazonAccessKeyId() string`

GetAmazonAccessKeyId returns the AmazonAccessKeyId field if non-nil, zero value otherwise.

### GetAmazonAccessKeyIdOk

`func (o *RestAPIOptionsAuthenticationOptions) GetAmazonAccessKeyIdOk() (*string, bool)`

GetAmazonAccessKeyIdOk returns a tuple with the AmazonAccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonAccessKeyId

`func (o *RestAPIOptionsAuthenticationOptions) SetAmazonAccessKeyId(v string)`

SetAmazonAccessKeyId sets AmazonAccessKeyId field to given value.


### GetAmazonSecretAccessKey

`func (o *RestAPIOptionsAuthenticationOptions) GetAmazonSecretAccessKey() string`

GetAmazonSecretAccessKey returns the AmazonSecretAccessKey field if non-nil, zero value otherwise.

### GetAmazonSecretAccessKeyOk

`func (o *RestAPIOptionsAuthenticationOptions) GetAmazonSecretAccessKeyOk() (*string, bool)`

GetAmazonSecretAccessKeyOk returns a tuple with the AmazonSecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmazonSecretAccessKey

`func (o *RestAPIOptionsAuthenticationOptions) SetAmazonSecretAccessKey(v string)`

SetAmazonSecretAccessKey sets AmazonSecretAccessKey field to given value.


### GetAuthWithDefaultCredentialProviderChain

`func (o *RestAPIOptionsAuthenticationOptions) GetAuthWithDefaultCredentialProviderChain() bool`

GetAuthWithDefaultCredentialProviderChain returns the AuthWithDefaultCredentialProviderChain field if non-nil, zero value otherwise.

### GetAuthWithDefaultCredentialProviderChainOk

`func (o *RestAPIOptionsAuthenticationOptions) GetAuthWithDefaultCredentialProviderChainOk() (*bool, bool)`

GetAuthWithDefaultCredentialProviderChainOk returns a tuple with the AuthWithDefaultCredentialProviderChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthWithDefaultCredentialProviderChain

`func (o *RestAPIOptionsAuthenticationOptions) SetAuthWithDefaultCredentialProviderChain(v bool)`

SetAuthWithDefaultCredentialProviderChain sets AuthWithDefaultCredentialProviderChain field to given value.

### HasAuthWithDefaultCredentialProviderChain

`func (o *RestAPIOptionsAuthenticationOptions) HasAuthWithDefaultCredentialProviderChain() bool`

HasAuthWithDefaultCredentialProviderChain returns a boolean if a field has been set.

### GetEnableAwsv4AuthenticationViaHeaders

`func (o *RestAPIOptionsAuthenticationOptions) GetEnableAwsv4AuthenticationViaHeaders() bool`

GetEnableAwsv4AuthenticationViaHeaders returns the EnableAwsv4AuthenticationViaHeaders field if non-nil, zero value otherwise.

### GetEnableAwsv4AuthenticationViaHeadersOk

`func (o *RestAPIOptionsAuthenticationOptions) GetEnableAwsv4AuthenticationViaHeadersOk() (*bool, bool)`

GetEnableAwsv4AuthenticationViaHeadersOk returns a tuple with the EnableAwsv4AuthenticationViaHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAwsv4AuthenticationViaHeaders

`func (o *RestAPIOptionsAuthenticationOptions) SetEnableAwsv4AuthenticationViaHeaders(v bool)`

SetEnableAwsv4AuthenticationViaHeaders sets EnableAwsv4AuthenticationViaHeaders field to given value.

### HasEnableAwsv4AuthenticationViaHeaders

`func (o *RestAPIOptionsAuthenticationOptions) HasEnableAwsv4AuthenticationViaHeaders() bool`

HasEnableAwsv4AuthenticationViaHeaders returns a boolean if a field has been set.

### GetBasicUsername

`func (o *RestAPIOptionsAuthenticationOptions) GetBasicUsername() string`

GetBasicUsername returns the BasicUsername field if non-nil, zero value otherwise.

### GetBasicUsernameOk

`func (o *RestAPIOptionsAuthenticationOptions) GetBasicUsernameOk() (*string, bool)`

GetBasicUsernameOk returns a tuple with the BasicUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicUsername

`func (o *RestAPIOptionsAuthenticationOptions) SetBasicUsername(v string)`

SetBasicUsername sets BasicUsername field to given value.


### GetBasicPassword

`func (o *RestAPIOptionsAuthenticationOptions) GetBasicPassword() string`

GetBasicPassword returns the BasicPassword field if non-nil, zero value otherwise.

### GetBasicPasswordOk

`func (o *RestAPIOptionsAuthenticationOptions) GetBasicPasswordOk() (*string, bool)`

GetBasicPasswordOk returns a tuple with the BasicPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicPassword

`func (o *RestAPIOptionsAuthenticationOptions) SetBasicPassword(v string)`

SetBasicPassword sets BasicPassword field to given value.

### HasBasicPassword

`func (o *RestAPIOptionsAuthenticationOptions) HasBasicPassword() bool`

HasBasicPassword returns a boolean if a field has been set.

### GetBearerToken

`func (o *RestAPIOptionsAuthenticationOptions) GetBearerToken() string`

GetBearerToken returns the BearerToken field if non-nil, zero value otherwise.

### GetBearerTokenOk

`func (o *RestAPIOptionsAuthenticationOptions) GetBearerTokenOk() (*string, bool)`

GetBearerTokenOk returns a tuple with the BearerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBearerToken

`func (o *RestAPIOptionsAuthenticationOptions) SetBearerToken(v string)`

SetBearerToken sets BearerToken field to given value.

### HasBearerToken

`func (o *RestAPIOptionsAuthenticationOptions) HasBearerToken() bool`

HasBearerToken returns a boolean if a field has been set.

### GetDigestUsername

`func (o *RestAPIOptionsAuthenticationOptions) GetDigestUsername() string`

GetDigestUsername returns the DigestUsername field if non-nil, zero value otherwise.

### GetDigestUsernameOk

`func (o *RestAPIOptionsAuthenticationOptions) GetDigestUsernameOk() (*string, bool)`

GetDigestUsernameOk returns a tuple with the DigestUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestUsername

`func (o *RestAPIOptionsAuthenticationOptions) SetDigestUsername(v string)`

SetDigestUsername sets DigestUsername field to given value.


### GetDigestPassword

`func (o *RestAPIOptionsAuthenticationOptions) GetDigestPassword() string`

GetDigestPassword returns the DigestPassword field if non-nil, zero value otherwise.

### GetDigestPasswordOk

`func (o *RestAPIOptionsAuthenticationOptions) GetDigestPasswordOk() (*string, bool)`

GetDigestPasswordOk returns a tuple with the DigestPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestPassword

`func (o *RestAPIOptionsAuthenticationOptions) SetDigestPassword(v string)`

SetDigestPassword sets DigestPassword field to given value.


### GetOauth1SignatureMethod

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth1SignatureMethod() RestAPIOptionsAuthenticationOptionsAnyOf5Oauth1SignatureMethod`

GetOauth1SignatureMethod returns the Oauth1SignatureMethod field if non-nil, zero value otherwise.

### GetOauth1SignatureMethodOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth1SignatureMethodOk() (*RestAPIOptionsAuthenticationOptionsAnyOf5Oauth1SignatureMethod, bool)`

GetOauth1SignatureMethodOk returns a tuple with the Oauth1SignatureMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth1SignatureMethod

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth1SignatureMethod(v RestAPIOptionsAuthenticationOptionsAnyOf5Oauth1SignatureMethod)`

SetOauth1SignatureMethod sets Oauth1SignatureMethod field to given value.


### GetOauth1ConsumerKey

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth1ConsumerKey() string`

GetOauth1ConsumerKey returns the Oauth1ConsumerKey field if non-nil, zero value otherwise.

### GetOauth1ConsumerKeyOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth1ConsumerKeyOk() (*string, bool)`

GetOauth1ConsumerKeyOk returns a tuple with the Oauth1ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth1ConsumerKey

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth1ConsumerKey(v string)`

SetOauth1ConsumerKey sets Oauth1ConsumerKey field to given value.


### GetOauth1ConsumerSecret

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth1ConsumerSecret() string`

GetOauth1ConsumerSecret returns the Oauth1ConsumerSecret field if non-nil, zero value otherwise.

### GetOauth1ConsumerSecretOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth1ConsumerSecretOk() (*string, bool)`

GetOauth1ConsumerSecretOk returns a tuple with the Oauth1ConsumerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth1ConsumerSecret

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth1ConsumerSecret(v string)`

SetOauth1ConsumerSecret sets Oauth1ConsumerSecret field to given value.


### GetOauth1TokenKey

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth1TokenKey() string`

GetOauth1TokenKey returns the Oauth1TokenKey field if non-nil, zero value otherwise.

### GetOauth1TokenKeyOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth1TokenKeyOk() (*string, bool)`

GetOauth1TokenKeyOk returns a tuple with the Oauth1TokenKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth1TokenKey

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth1TokenKey(v string)`

SetOauth1TokenKey sets Oauth1TokenKey field to given value.


### GetOauth1TokenSecret

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth1TokenSecret() string`

GetOauth1TokenSecret returns the Oauth1TokenSecret field if non-nil, zero value otherwise.

### GetOauth1TokenSecretOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth1TokenSecretOk() (*string, bool)`

GetOauth1TokenSecretOk returns a tuple with the Oauth1TokenSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth1TokenSecret

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth1TokenSecret(v string)`

SetOauth1TokenSecret sets Oauth1TokenSecret field to given value.


### GetOauth1RealmParameter

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth1RealmParameter() string`

GetOauth1RealmParameter returns the Oauth1RealmParameter field if non-nil, zero value otherwise.

### GetOauth1RealmParameterOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth1RealmParameterOk() (*string, bool)`

GetOauth1RealmParameterOk returns a tuple with the Oauth1RealmParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth1RealmParameter

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth1RealmParameter(v string)`

SetOauth1RealmParameter sets Oauth1RealmParameter field to given value.


### GetOauth2Audience

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2Audience() string`

GetOauth2Audience returns the Oauth2Audience field if non-nil, zero value otherwise.

### GetOauth2AudienceOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2AudienceOk() (*string, bool)`

GetOauth2AudienceOk returns a tuple with the Oauth2Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Audience

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth2Audience(v string)`

SetOauth2Audience sets Oauth2Audience field to given value.

### HasOauth2Audience

`func (o *RestAPIOptionsAuthenticationOptions) HasOauth2Audience() bool`

HasOauth2Audience returns a boolean if a field has been set.

### GetOauth2AccessToken

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2AccessToken() string`

GetOauth2AccessToken returns the Oauth2AccessToken field if non-nil, zero value otherwise.

### GetOauth2AccessTokenOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2AccessTokenOk() (*string, bool)`

GetOauth2AccessTokenOk returns a tuple with the Oauth2AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2AccessToken

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth2AccessToken(v string)`

SetOauth2AccessToken sets Oauth2AccessToken field to given value.

### HasOauth2AccessToken

`func (o *RestAPIOptionsAuthenticationOptions) HasOauth2AccessToken() bool`

HasOauth2AccessToken returns a boolean if a field has been set.

### GetOauth2AccessTokenUrl

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2AccessTokenUrl() string`

GetOauth2AccessTokenUrl returns the Oauth2AccessTokenUrl field if non-nil, zero value otherwise.

### GetOauth2AccessTokenUrlOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2AccessTokenUrlOk() (*string, bool)`

GetOauth2AccessTokenUrlOk returns a tuple with the Oauth2AccessTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2AccessTokenUrl

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth2AccessTokenUrl(v string)`

SetOauth2AccessTokenUrl sets Oauth2AccessTokenUrl field to given value.


### GetOauth2AccessTokenLifespanSeconds

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2AccessTokenLifespanSeconds() float32`

GetOauth2AccessTokenLifespanSeconds returns the Oauth2AccessTokenLifespanSeconds field if non-nil, zero value otherwise.

### GetOauth2AccessTokenLifespanSecondsOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2AccessTokenLifespanSecondsOk() (*float32, bool)`

GetOauth2AccessTokenLifespanSecondsOk returns a tuple with the Oauth2AccessTokenLifespanSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2AccessTokenLifespanSeconds

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth2AccessTokenLifespanSeconds(v float32)`

SetOauth2AccessTokenLifespanSeconds sets Oauth2AccessTokenLifespanSeconds field to given value.

### HasOauth2AccessTokenLifespanSeconds

`func (o *RestAPIOptionsAuthenticationOptions) HasOauth2AccessTokenLifespanSeconds() bool`

HasOauth2AccessTokenLifespanSeconds returns a boolean if a field has been set.

### SetOauth2AccessTokenLifespanSecondsNil

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth2AccessTokenLifespanSecondsNil(b bool)`

 SetOauth2AccessTokenLifespanSecondsNil sets the value for Oauth2AccessTokenLifespanSeconds to be an explicit nil

### UnsetOauth2AccessTokenLifespanSeconds
`func (o *RestAPIOptionsAuthenticationOptions) UnsetOauth2AccessTokenLifespanSeconds()`

UnsetOauth2AccessTokenLifespanSeconds ensures that no value is present for Oauth2AccessTokenLifespanSeconds, not even an explicit nil
### GetOauth2AuthUrl

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2AuthUrl() string`

GetOauth2AuthUrl returns the Oauth2AuthUrl field if non-nil, zero value otherwise.

### GetOauth2AuthUrlOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2AuthUrlOk() (*string, bool)`

GetOauth2AuthUrlOk returns a tuple with the Oauth2AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2AuthUrl

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth2AuthUrl(v string)`

SetOauth2AuthUrl sets Oauth2AuthUrl field to given value.


### GetOauth2ClientId

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2ClientId() string`

GetOauth2ClientId returns the Oauth2ClientId field if non-nil, zero value otherwise.

### GetOauth2ClientIdOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2ClientIdOk() (*string, bool)`

GetOauth2ClientIdOk returns a tuple with the Oauth2ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ClientId

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth2ClientId(v string)`

SetOauth2ClientId sets Oauth2ClientId field to given value.


### GetOauth2ClientSecret

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2ClientSecret() string`

GetOauth2ClientSecret returns the Oauth2ClientSecret field if non-nil, zero value otherwise.

### GetOauth2ClientSecretOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2ClientSecretOk() (*string, bool)`

GetOauth2ClientSecretOk returns a tuple with the Oauth2ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ClientSecret

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth2ClientSecret(v string)`

SetOauth2ClientSecret sets Oauth2ClientSecret field to given value.


### GetOauth2CallbackUrl

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2CallbackUrl() string`

GetOauth2CallbackUrl returns the Oauth2CallbackUrl field if non-nil, zero value otherwise.

### GetOauth2CallbackUrlOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2CallbackUrlOk() (*string, bool)`

GetOauth2CallbackUrlOk returns a tuple with the Oauth2CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2CallbackUrl

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth2CallbackUrl(v string)`

SetOauth2CallbackUrl sets Oauth2CallbackUrl field to given value.

### HasOauth2CallbackUrl

`func (o *RestAPIOptionsAuthenticationOptions) HasOauth2CallbackUrl() bool`

HasOauth2CallbackUrl returns a boolean if a field has been set.

### GetOauth2IdToken

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2IdToken() string`

GetOauth2IdToken returns the Oauth2IdToken field if non-nil, zero value otherwise.

### GetOauth2IdTokenOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2IdTokenOk() (*string, bool)`

GetOauth2IdTokenOk returns a tuple with the Oauth2IdToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2IdToken

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth2IdToken(v string)`

SetOauth2IdToken sets Oauth2IdToken field to given value.

### HasOauth2IdToken

`func (o *RestAPIOptionsAuthenticationOptions) HasOauth2IdToken() bool`

HasOauth2IdToken returns a boolean if a field has been set.

### GetOauth2RefreshToken

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2RefreshToken() string`

GetOauth2RefreshToken returns the Oauth2RefreshToken field if non-nil, zero value otherwise.

### GetOauth2RefreshTokenOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2RefreshTokenOk() (*string, bool)`

GetOauth2RefreshTokenOk returns a tuple with the Oauth2RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2RefreshToken

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth2RefreshToken(v string)`

SetOauth2RefreshToken sets Oauth2RefreshToken field to given value.

### HasOauth2RefreshToken

`func (o *RestAPIOptionsAuthenticationOptions) HasOauth2RefreshToken() bool`

HasOauth2RefreshToken returns a boolean if a field has been set.

### GetOauth2Scope

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2Scope() string`

GetOauth2Scope returns the Oauth2Scope field if non-nil, zero value otherwise.

### GetOauth2ScopeOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2ScopeOk() (*string, bool)`

GetOauth2ScopeOk returns a tuple with the Oauth2Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Scope

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth2Scope(v string)`

SetOauth2Scope sets Oauth2Scope field to given value.

### HasOauth2Scope

`func (o *RestAPIOptionsAuthenticationOptions) HasOauth2Scope() bool`

HasOauth2Scope returns a boolean if a field has been set.

### GetOauth2ShareUserCredentials

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2ShareUserCredentials() bool`

GetOauth2ShareUserCredentials returns the Oauth2ShareUserCredentials field if non-nil, zero value otherwise.

### GetOauth2ShareUserCredentialsOk

`func (o *RestAPIOptionsAuthenticationOptions) GetOauth2ShareUserCredentialsOk() (*bool, bool)`

GetOauth2ShareUserCredentialsOk returns a tuple with the Oauth2ShareUserCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2ShareUserCredentials

`func (o *RestAPIOptionsAuthenticationOptions) SetOauth2ShareUserCredentials(v bool)`

SetOauth2ShareUserCredentials sets Oauth2ShareUserCredentials field to given value.

### HasOauth2ShareUserCredentials

`func (o *RestAPIOptionsAuthenticationOptions) HasOauth2ShareUserCredentials() bool`

HasOauth2ShareUserCredentials returns a boolean if a field has been set.

### GetVerifySessionUrl

`func (o *RestAPIOptionsAuthenticationOptions) GetVerifySessionUrl() string`

GetVerifySessionUrl returns the VerifySessionUrl field if non-nil, zero value otherwise.

### GetVerifySessionUrlOk

`func (o *RestAPIOptionsAuthenticationOptions) GetVerifySessionUrlOk() (*string, bool)`

GetVerifySessionUrlOk returns a tuple with the VerifySessionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifySessionUrl

`func (o *RestAPIOptionsAuthenticationOptions) SetVerifySessionUrl(v string)`

SetVerifySessionUrl sets VerifySessionUrl field to given value.

### HasVerifySessionUrl

`func (o *RestAPIOptionsAuthenticationOptions) HasVerifySessionUrl() bool`

HasVerifySessionUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


