# SsoConfigPostRequestDataOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** |  | 
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
**DisableEmailPasswordLogin** | **bool** |  | 

## Methods

### NewSsoConfigPostRequestDataOneOf1

`func NewSsoConfigPostRequestDataOneOf1(configType string, oidcClientId string, oidcClientSecret string, oidcScopes string, oidcAuthUrl string, oidcTokenUrl string, jwtEmailKey string, jwtFirstNameKey string, jwtLastNameKey string, jitEnabled bool, triggerLoginAutomatically bool, disableEmailPasswordLogin bool, ) *SsoConfigPostRequestDataOneOf1`

NewSsoConfigPostRequestDataOneOf1 instantiates a new SsoConfigPostRequestDataOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoConfigPostRequestDataOneOf1WithDefaults

`func NewSsoConfigPostRequestDataOneOf1WithDefaults() *SsoConfigPostRequestDataOneOf1`

NewSsoConfigPostRequestDataOneOf1WithDefaults instantiates a new SsoConfigPostRequestDataOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *SsoConfigPostRequestDataOneOf1) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *SsoConfigPostRequestDataOneOf1) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *SsoConfigPostRequestDataOneOf1) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetOidcClientId

`func (o *SsoConfigPostRequestDataOneOf1) GetOidcClientId() string`

GetOidcClientId returns the OidcClientId field if non-nil, zero value otherwise.

### GetOidcClientIdOk

`func (o *SsoConfigPostRequestDataOneOf1) GetOidcClientIdOk() (*string, bool)`

GetOidcClientIdOk returns a tuple with the OidcClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientId

`func (o *SsoConfigPostRequestDataOneOf1) SetOidcClientId(v string)`

SetOidcClientId sets OidcClientId field to given value.


### GetOidcClientSecret

`func (o *SsoConfigPostRequestDataOneOf1) GetOidcClientSecret() string`

GetOidcClientSecret returns the OidcClientSecret field if non-nil, zero value otherwise.

### GetOidcClientSecretOk

`func (o *SsoConfigPostRequestDataOneOf1) GetOidcClientSecretOk() (*string, bool)`

GetOidcClientSecretOk returns a tuple with the OidcClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientSecret

`func (o *SsoConfigPostRequestDataOneOf1) SetOidcClientSecret(v string)`

SetOidcClientSecret sets OidcClientSecret field to given value.


### GetOidcScopes

`func (o *SsoConfigPostRequestDataOneOf1) GetOidcScopes() string`

GetOidcScopes returns the OidcScopes field if non-nil, zero value otherwise.

### GetOidcScopesOk

`func (o *SsoConfigPostRequestDataOneOf1) GetOidcScopesOk() (*string, bool)`

GetOidcScopesOk returns a tuple with the OidcScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcScopes

`func (o *SsoConfigPostRequestDataOneOf1) SetOidcScopes(v string)`

SetOidcScopes sets OidcScopes field to given value.


### GetOidcAuthUrl

`func (o *SsoConfigPostRequestDataOneOf1) GetOidcAuthUrl() string`

GetOidcAuthUrl returns the OidcAuthUrl field if non-nil, zero value otherwise.

### GetOidcAuthUrlOk

`func (o *SsoConfigPostRequestDataOneOf1) GetOidcAuthUrlOk() (*string, bool)`

GetOidcAuthUrlOk returns a tuple with the OidcAuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAuthUrl

`func (o *SsoConfigPostRequestDataOneOf1) SetOidcAuthUrl(v string)`

SetOidcAuthUrl sets OidcAuthUrl field to given value.


### GetOidcTokenUrl

`func (o *SsoConfigPostRequestDataOneOf1) GetOidcTokenUrl() string`

GetOidcTokenUrl returns the OidcTokenUrl field if non-nil, zero value otherwise.

### GetOidcTokenUrlOk

`func (o *SsoConfigPostRequestDataOneOf1) GetOidcTokenUrlOk() (*string, bool)`

GetOidcTokenUrlOk returns a tuple with the OidcTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcTokenUrl

`func (o *SsoConfigPostRequestDataOneOf1) SetOidcTokenUrl(v string)`

SetOidcTokenUrl sets OidcTokenUrl field to given value.


### GetOidcUserinfoUrl

`func (o *SsoConfigPostRequestDataOneOf1) GetOidcUserinfoUrl() string`

GetOidcUserinfoUrl returns the OidcUserinfoUrl field if non-nil, zero value otherwise.

### GetOidcUserinfoUrlOk

`func (o *SsoConfigPostRequestDataOneOf1) GetOidcUserinfoUrlOk() (*string, bool)`

GetOidcUserinfoUrlOk returns a tuple with the OidcUserinfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcUserinfoUrl

`func (o *SsoConfigPostRequestDataOneOf1) SetOidcUserinfoUrl(v string)`

SetOidcUserinfoUrl sets OidcUserinfoUrl field to given value.

### HasOidcUserinfoUrl

`func (o *SsoConfigPostRequestDataOneOf1) HasOidcUserinfoUrl() bool`

HasOidcUserinfoUrl returns a boolean if a field has been set.

### GetOidcAudience

`func (o *SsoConfigPostRequestDataOneOf1) GetOidcAudience() string`

GetOidcAudience returns the OidcAudience field if non-nil, zero value otherwise.

### GetOidcAudienceOk

`func (o *SsoConfigPostRequestDataOneOf1) GetOidcAudienceOk() (*string, bool)`

GetOidcAudienceOk returns a tuple with the OidcAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAudience

`func (o *SsoConfigPostRequestDataOneOf1) SetOidcAudience(v string)`

SetOidcAudience sets OidcAudience field to given value.

### HasOidcAudience

`func (o *SsoConfigPostRequestDataOneOf1) HasOidcAudience() bool`

HasOidcAudience returns a boolean if a field has been set.

### GetJwtEmailKey

`func (o *SsoConfigPostRequestDataOneOf1) GetJwtEmailKey() string`

GetJwtEmailKey returns the JwtEmailKey field if non-nil, zero value otherwise.

### GetJwtEmailKeyOk

`func (o *SsoConfigPostRequestDataOneOf1) GetJwtEmailKeyOk() (*string, bool)`

GetJwtEmailKeyOk returns a tuple with the JwtEmailKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtEmailKey

`func (o *SsoConfigPostRequestDataOneOf1) SetJwtEmailKey(v string)`

SetJwtEmailKey sets JwtEmailKey field to given value.


### GetJwtRolesKey

`func (o *SsoConfigPostRequestDataOneOf1) GetJwtRolesKey() string`

GetJwtRolesKey returns the JwtRolesKey field if non-nil, zero value otherwise.

### GetJwtRolesKeyOk

`func (o *SsoConfigPostRequestDataOneOf1) GetJwtRolesKeyOk() (*string, bool)`

GetJwtRolesKeyOk returns a tuple with the JwtRolesKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtRolesKey

`func (o *SsoConfigPostRequestDataOneOf1) SetJwtRolesKey(v string)`

SetJwtRolesKey sets JwtRolesKey field to given value.

### HasJwtRolesKey

`func (o *SsoConfigPostRequestDataOneOf1) HasJwtRolesKey() bool`

HasJwtRolesKey returns a boolean if a field has been set.

### GetJwtFirstNameKey

`func (o *SsoConfigPostRequestDataOneOf1) GetJwtFirstNameKey() string`

GetJwtFirstNameKey returns the JwtFirstNameKey field if non-nil, zero value otherwise.

### GetJwtFirstNameKeyOk

`func (o *SsoConfigPostRequestDataOneOf1) GetJwtFirstNameKeyOk() (*string, bool)`

GetJwtFirstNameKeyOk returns a tuple with the JwtFirstNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtFirstNameKey

`func (o *SsoConfigPostRequestDataOneOf1) SetJwtFirstNameKey(v string)`

SetJwtFirstNameKey sets JwtFirstNameKey field to given value.


### GetJwtLastNameKey

`func (o *SsoConfigPostRequestDataOneOf1) GetJwtLastNameKey() string`

GetJwtLastNameKey returns the JwtLastNameKey field if non-nil, zero value otherwise.

### GetJwtLastNameKeyOk

`func (o *SsoConfigPostRequestDataOneOf1) GetJwtLastNameKeyOk() (*string, bool)`

GetJwtLastNameKeyOk returns a tuple with the JwtLastNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtLastNameKey

`func (o *SsoConfigPostRequestDataOneOf1) SetJwtLastNameKey(v string)`

SetJwtLastNameKey sets JwtLastNameKey field to given value.


### GetRolesMapping

`func (o *SsoConfigPostRequestDataOneOf1) GetRolesMapping() string`

GetRolesMapping returns the RolesMapping field if non-nil, zero value otherwise.

### GetRolesMappingOk

`func (o *SsoConfigPostRequestDataOneOf1) GetRolesMappingOk() (*string, bool)`

GetRolesMappingOk returns a tuple with the RolesMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolesMapping

`func (o *SsoConfigPostRequestDataOneOf1) SetRolesMapping(v string)`

SetRolesMapping sets RolesMapping field to given value.

### HasRolesMapping

`func (o *SsoConfigPostRequestDataOneOf1) HasRolesMapping() bool`

HasRolesMapping returns a boolean if a field has been set.

### GetJitEnabled

`func (o *SsoConfigPostRequestDataOneOf1) GetJitEnabled() bool`

GetJitEnabled returns the JitEnabled field if non-nil, zero value otherwise.

### GetJitEnabledOk

`func (o *SsoConfigPostRequestDataOneOf1) GetJitEnabledOk() (*bool, bool)`

GetJitEnabledOk returns a tuple with the JitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitEnabled

`func (o *SsoConfigPostRequestDataOneOf1) SetJitEnabled(v bool)`

SetJitEnabled sets JitEnabled field to given value.


### GetRestrictedDomain

`func (o *SsoConfigPostRequestDataOneOf1) GetRestrictedDomain() string`

GetRestrictedDomain returns the RestrictedDomain field if non-nil, zero value otherwise.

### GetRestrictedDomainOk

`func (o *SsoConfigPostRequestDataOneOf1) GetRestrictedDomainOk() (*string, bool)`

GetRestrictedDomainOk returns a tuple with the RestrictedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDomain

`func (o *SsoConfigPostRequestDataOneOf1) SetRestrictedDomain(v string)`

SetRestrictedDomain sets RestrictedDomain field to given value.

### HasRestrictedDomain

`func (o *SsoConfigPostRequestDataOneOf1) HasRestrictedDomain() bool`

HasRestrictedDomain returns a boolean if a field has been set.

### GetTriggerLoginAutomatically

`func (o *SsoConfigPostRequestDataOneOf1) GetTriggerLoginAutomatically() bool`

GetTriggerLoginAutomatically returns the TriggerLoginAutomatically field if non-nil, zero value otherwise.

### GetTriggerLoginAutomaticallyOk

`func (o *SsoConfigPostRequestDataOneOf1) GetTriggerLoginAutomaticallyOk() (*bool, bool)`

GetTriggerLoginAutomaticallyOk returns a tuple with the TriggerLoginAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerLoginAutomatically

`func (o *SsoConfigPostRequestDataOneOf1) SetTriggerLoginAutomatically(v bool)`

SetTriggerLoginAutomatically sets TriggerLoginAutomatically field to given value.


### GetDisableEmailPasswordLogin

`func (o *SsoConfigPostRequestDataOneOf1) GetDisableEmailPasswordLogin() bool`

GetDisableEmailPasswordLogin returns the DisableEmailPasswordLogin field if non-nil, zero value otherwise.

### GetDisableEmailPasswordLoginOk

`func (o *SsoConfigPostRequestDataOneOf1) GetDisableEmailPasswordLoginOk() (*bool, bool)`

GetDisableEmailPasswordLoginOk returns a tuple with the DisableEmailPasswordLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmailPasswordLogin

`func (o *SsoConfigPostRequestDataOneOf1) SetDisableEmailPasswordLogin(v bool)`

SetDisableEmailPasswordLogin sets DisableEmailPasswordLogin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


