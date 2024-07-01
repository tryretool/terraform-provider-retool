# SsoConfigPostRequestDataOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** |  | 
**GoogleClientId** | **string** |  | 
**GoogleClientSecret** | **string** |  | 
**DisableEmailPasswordLogin** | **bool** |  | 
**OidcClientId** | **string** |  | 
**OidcClientSecret** | **string** |  | 
**OidcScopes** | **string** |  | 
**OidcAuthUrl** | **string** |  | 
**OidcTokenUrl** | **string** |  | 
**OidcUserinfoUrl** | Pointer to **string** |  | [optional] 
**OidcAudience** | Pointer to **string** |  | [optional] 
**JwtEmailKey** | **string** |  | 
**JwtRolesKey** | Pointer to **string** |  | [optional] 
**JwtFirstNameKey** | **string** |  | 
**JwtLastNameKey** | **string** |  | 
**RolesMapping** | Pointer to **string** |  | [optional] 
**JitEnabled** | **bool** |  | 
**RestrictedDomain** | Pointer to **string** |  | [optional] 
**TriggerLoginAutomatically** | **bool** |  | 

## Methods

### NewSsoConfigPostRequestDataOneOf2

`func NewSsoConfigPostRequestDataOneOf2(configType string, googleClientId string, googleClientSecret string, disableEmailPasswordLogin bool, oidcClientId string, oidcClientSecret string, oidcScopes string, oidcAuthUrl string, oidcTokenUrl string, jwtEmailKey string, jwtFirstNameKey string, jwtLastNameKey string, jitEnabled bool, triggerLoginAutomatically bool, ) *SsoConfigPostRequestDataOneOf2`

NewSsoConfigPostRequestDataOneOf2 instantiates a new SsoConfigPostRequestDataOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoConfigPostRequestDataOneOf2WithDefaults

`func NewSsoConfigPostRequestDataOneOf2WithDefaults() *SsoConfigPostRequestDataOneOf2`

NewSsoConfigPostRequestDataOneOf2WithDefaults instantiates a new SsoConfigPostRequestDataOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *SsoConfigPostRequestDataOneOf2) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *SsoConfigPostRequestDataOneOf2) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *SsoConfigPostRequestDataOneOf2) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetGoogleClientId

`func (o *SsoConfigPostRequestDataOneOf2) GetGoogleClientId() string`

GetGoogleClientId returns the GoogleClientId field if non-nil, zero value otherwise.

### GetGoogleClientIdOk

`func (o *SsoConfigPostRequestDataOneOf2) GetGoogleClientIdOk() (*string, bool)`

GetGoogleClientIdOk returns a tuple with the GoogleClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleClientId

`func (o *SsoConfigPostRequestDataOneOf2) SetGoogleClientId(v string)`

SetGoogleClientId sets GoogleClientId field to given value.


### GetGoogleClientSecret

`func (o *SsoConfigPostRequestDataOneOf2) GetGoogleClientSecret() string`

GetGoogleClientSecret returns the GoogleClientSecret field if non-nil, zero value otherwise.

### GetGoogleClientSecretOk

`func (o *SsoConfigPostRequestDataOneOf2) GetGoogleClientSecretOk() (*string, bool)`

GetGoogleClientSecretOk returns a tuple with the GoogleClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleClientSecret

`func (o *SsoConfigPostRequestDataOneOf2) SetGoogleClientSecret(v string)`

SetGoogleClientSecret sets GoogleClientSecret field to given value.


### GetDisableEmailPasswordLogin

`func (o *SsoConfigPostRequestDataOneOf2) GetDisableEmailPasswordLogin() bool`

GetDisableEmailPasswordLogin returns the DisableEmailPasswordLogin field if non-nil, zero value otherwise.

### GetDisableEmailPasswordLoginOk

`func (o *SsoConfigPostRequestDataOneOf2) GetDisableEmailPasswordLoginOk() (*bool, bool)`

GetDisableEmailPasswordLoginOk returns a tuple with the DisableEmailPasswordLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmailPasswordLogin

`func (o *SsoConfigPostRequestDataOneOf2) SetDisableEmailPasswordLogin(v bool)`

SetDisableEmailPasswordLogin sets DisableEmailPasswordLogin field to given value.


### GetOidcClientId

`func (o *SsoConfigPostRequestDataOneOf2) GetOidcClientId() string`

GetOidcClientId returns the OidcClientId field if non-nil, zero value otherwise.

### GetOidcClientIdOk

`func (o *SsoConfigPostRequestDataOneOf2) GetOidcClientIdOk() (*string, bool)`

GetOidcClientIdOk returns a tuple with the OidcClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientId

`func (o *SsoConfigPostRequestDataOneOf2) SetOidcClientId(v string)`

SetOidcClientId sets OidcClientId field to given value.


### GetOidcClientSecret

`func (o *SsoConfigPostRequestDataOneOf2) GetOidcClientSecret() string`

GetOidcClientSecret returns the OidcClientSecret field if non-nil, zero value otherwise.

### GetOidcClientSecretOk

`func (o *SsoConfigPostRequestDataOneOf2) GetOidcClientSecretOk() (*string, bool)`

GetOidcClientSecretOk returns a tuple with the OidcClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientSecret

`func (o *SsoConfigPostRequestDataOneOf2) SetOidcClientSecret(v string)`

SetOidcClientSecret sets OidcClientSecret field to given value.


### GetOidcScopes

`func (o *SsoConfigPostRequestDataOneOf2) GetOidcScopes() string`

GetOidcScopes returns the OidcScopes field if non-nil, zero value otherwise.

### GetOidcScopesOk

`func (o *SsoConfigPostRequestDataOneOf2) GetOidcScopesOk() (*string, bool)`

GetOidcScopesOk returns a tuple with the OidcScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcScopes

`func (o *SsoConfigPostRequestDataOneOf2) SetOidcScopes(v string)`

SetOidcScopes sets OidcScopes field to given value.


### GetOidcAuthUrl

`func (o *SsoConfigPostRequestDataOneOf2) GetOidcAuthUrl() string`

GetOidcAuthUrl returns the OidcAuthUrl field if non-nil, zero value otherwise.

### GetOidcAuthUrlOk

`func (o *SsoConfigPostRequestDataOneOf2) GetOidcAuthUrlOk() (*string, bool)`

GetOidcAuthUrlOk returns a tuple with the OidcAuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAuthUrl

`func (o *SsoConfigPostRequestDataOneOf2) SetOidcAuthUrl(v string)`

SetOidcAuthUrl sets OidcAuthUrl field to given value.


### GetOidcTokenUrl

`func (o *SsoConfigPostRequestDataOneOf2) GetOidcTokenUrl() string`

GetOidcTokenUrl returns the OidcTokenUrl field if non-nil, zero value otherwise.

### GetOidcTokenUrlOk

`func (o *SsoConfigPostRequestDataOneOf2) GetOidcTokenUrlOk() (*string, bool)`

GetOidcTokenUrlOk returns a tuple with the OidcTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcTokenUrl

`func (o *SsoConfigPostRequestDataOneOf2) SetOidcTokenUrl(v string)`

SetOidcTokenUrl sets OidcTokenUrl field to given value.


### GetOidcUserinfoUrl

`func (o *SsoConfigPostRequestDataOneOf2) GetOidcUserinfoUrl() string`

GetOidcUserinfoUrl returns the OidcUserinfoUrl field if non-nil, zero value otherwise.

### GetOidcUserinfoUrlOk

`func (o *SsoConfigPostRequestDataOneOf2) GetOidcUserinfoUrlOk() (*string, bool)`

GetOidcUserinfoUrlOk returns a tuple with the OidcUserinfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcUserinfoUrl

`func (o *SsoConfigPostRequestDataOneOf2) SetOidcUserinfoUrl(v string)`

SetOidcUserinfoUrl sets OidcUserinfoUrl field to given value.

### HasOidcUserinfoUrl

`func (o *SsoConfigPostRequestDataOneOf2) HasOidcUserinfoUrl() bool`

HasOidcUserinfoUrl returns a boolean if a field has been set.

### GetOidcAudience

`func (o *SsoConfigPostRequestDataOneOf2) GetOidcAudience() string`

GetOidcAudience returns the OidcAudience field if non-nil, zero value otherwise.

### GetOidcAudienceOk

`func (o *SsoConfigPostRequestDataOneOf2) GetOidcAudienceOk() (*string, bool)`

GetOidcAudienceOk returns a tuple with the OidcAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAudience

`func (o *SsoConfigPostRequestDataOneOf2) SetOidcAudience(v string)`

SetOidcAudience sets OidcAudience field to given value.

### HasOidcAudience

`func (o *SsoConfigPostRequestDataOneOf2) HasOidcAudience() bool`

HasOidcAudience returns a boolean if a field has been set.

### GetJwtEmailKey

`func (o *SsoConfigPostRequestDataOneOf2) GetJwtEmailKey() string`

GetJwtEmailKey returns the JwtEmailKey field if non-nil, zero value otherwise.

### GetJwtEmailKeyOk

`func (o *SsoConfigPostRequestDataOneOf2) GetJwtEmailKeyOk() (*string, bool)`

GetJwtEmailKeyOk returns a tuple with the JwtEmailKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtEmailKey

`func (o *SsoConfigPostRequestDataOneOf2) SetJwtEmailKey(v string)`

SetJwtEmailKey sets JwtEmailKey field to given value.


### GetJwtRolesKey

`func (o *SsoConfigPostRequestDataOneOf2) GetJwtRolesKey() string`

GetJwtRolesKey returns the JwtRolesKey field if non-nil, zero value otherwise.

### GetJwtRolesKeyOk

`func (o *SsoConfigPostRequestDataOneOf2) GetJwtRolesKeyOk() (*string, bool)`

GetJwtRolesKeyOk returns a tuple with the JwtRolesKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtRolesKey

`func (o *SsoConfigPostRequestDataOneOf2) SetJwtRolesKey(v string)`

SetJwtRolesKey sets JwtRolesKey field to given value.

### HasJwtRolesKey

`func (o *SsoConfigPostRequestDataOneOf2) HasJwtRolesKey() bool`

HasJwtRolesKey returns a boolean if a field has been set.

### GetJwtFirstNameKey

`func (o *SsoConfigPostRequestDataOneOf2) GetJwtFirstNameKey() string`

GetJwtFirstNameKey returns the JwtFirstNameKey field if non-nil, zero value otherwise.

### GetJwtFirstNameKeyOk

`func (o *SsoConfigPostRequestDataOneOf2) GetJwtFirstNameKeyOk() (*string, bool)`

GetJwtFirstNameKeyOk returns a tuple with the JwtFirstNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtFirstNameKey

`func (o *SsoConfigPostRequestDataOneOf2) SetJwtFirstNameKey(v string)`

SetJwtFirstNameKey sets JwtFirstNameKey field to given value.


### GetJwtLastNameKey

`func (o *SsoConfigPostRequestDataOneOf2) GetJwtLastNameKey() string`

GetJwtLastNameKey returns the JwtLastNameKey field if non-nil, zero value otherwise.

### GetJwtLastNameKeyOk

`func (o *SsoConfigPostRequestDataOneOf2) GetJwtLastNameKeyOk() (*string, bool)`

GetJwtLastNameKeyOk returns a tuple with the JwtLastNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtLastNameKey

`func (o *SsoConfigPostRequestDataOneOf2) SetJwtLastNameKey(v string)`

SetJwtLastNameKey sets JwtLastNameKey field to given value.


### GetRolesMapping

`func (o *SsoConfigPostRequestDataOneOf2) GetRolesMapping() string`

GetRolesMapping returns the RolesMapping field if non-nil, zero value otherwise.

### GetRolesMappingOk

`func (o *SsoConfigPostRequestDataOneOf2) GetRolesMappingOk() (*string, bool)`

GetRolesMappingOk returns a tuple with the RolesMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolesMapping

`func (o *SsoConfigPostRequestDataOneOf2) SetRolesMapping(v string)`

SetRolesMapping sets RolesMapping field to given value.

### HasRolesMapping

`func (o *SsoConfigPostRequestDataOneOf2) HasRolesMapping() bool`

HasRolesMapping returns a boolean if a field has been set.

### GetJitEnabled

`func (o *SsoConfigPostRequestDataOneOf2) GetJitEnabled() bool`

GetJitEnabled returns the JitEnabled field if non-nil, zero value otherwise.

### GetJitEnabledOk

`func (o *SsoConfigPostRequestDataOneOf2) GetJitEnabledOk() (*bool, bool)`

GetJitEnabledOk returns a tuple with the JitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitEnabled

`func (o *SsoConfigPostRequestDataOneOf2) SetJitEnabled(v bool)`

SetJitEnabled sets JitEnabled field to given value.


### GetRestrictedDomain

`func (o *SsoConfigPostRequestDataOneOf2) GetRestrictedDomain() string`

GetRestrictedDomain returns the RestrictedDomain field if non-nil, zero value otherwise.

### GetRestrictedDomainOk

`func (o *SsoConfigPostRequestDataOneOf2) GetRestrictedDomainOk() (*string, bool)`

GetRestrictedDomainOk returns a tuple with the RestrictedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDomain

`func (o *SsoConfigPostRequestDataOneOf2) SetRestrictedDomain(v string)`

SetRestrictedDomain sets RestrictedDomain field to given value.

### HasRestrictedDomain

`func (o *SsoConfigPostRequestDataOneOf2) HasRestrictedDomain() bool`

HasRestrictedDomain returns a boolean if a field has been set.

### GetTriggerLoginAutomatically

`func (o *SsoConfigPostRequestDataOneOf2) GetTriggerLoginAutomatically() bool`

GetTriggerLoginAutomatically returns the TriggerLoginAutomatically field if non-nil, zero value otherwise.

### GetTriggerLoginAutomaticallyOk

`func (o *SsoConfigPostRequestDataOneOf2) GetTriggerLoginAutomaticallyOk() (*bool, bool)`

GetTriggerLoginAutomaticallyOk returns a tuple with the TriggerLoginAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerLoginAutomatically

`func (o *SsoConfigPostRequestDataOneOf2) SetTriggerLoginAutomatically(v bool)`

SetTriggerLoginAutomatically sets TriggerLoginAutomatically field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


