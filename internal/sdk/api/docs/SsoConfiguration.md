# SsoConfiguration

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
**IdpMetadataXml** | **string** |  | 
**SamlFirstNameAttribute** | **string** |  | 
**SamlLastNameAttribute** | **string** |  | 
**SamlGroupsAttribute** | Pointer to **string** |  | [optional] 
**SamlSyncGroupClaims** | **bool** |  | 
**LdapSyncGroupClaims** | Pointer to **bool** |  | [optional] 
**LdapRoleMapping** | Pointer to **string** |  | [optional] 
**LdapServerUrl** | Pointer to **string** |  | [optional] 
**LdapBaseDomainComponents** | Pointer to **string** |  | [optional] 
**LdapServerName** | Pointer to **string** |  | [optional] 
**LdapServerKey** | Pointer to **string** |  | [optional] 
**LdapServerCertificate** | Pointer to **string** |  | [optional] 

## Methods

### NewSsoConfiguration

`func NewSsoConfiguration(configType string, googleClientId string, googleClientSecret string, disableEmailPasswordLogin bool, oidcClientId string, oidcClientSecret string, oidcScopes string, oidcAuthUrl string, oidcTokenUrl string, jwtEmailKey string, jwtFirstNameKey string, jwtLastNameKey string, jitEnabled bool, triggerLoginAutomatically bool, idpMetadataXml string, samlFirstNameAttribute string, samlLastNameAttribute string, samlSyncGroupClaims bool, ) *SsoConfiguration`

NewSsoConfiguration instantiates a new SsoConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoConfigurationWithDefaults

`func NewSsoConfigurationWithDefaults() *SsoConfiguration`

NewSsoConfigurationWithDefaults instantiates a new SsoConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *SsoConfiguration) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *SsoConfiguration) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *SsoConfiguration) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetGoogleClientId

`func (o *SsoConfiguration) GetGoogleClientId() string`

GetGoogleClientId returns the GoogleClientId field if non-nil, zero value otherwise.

### GetGoogleClientIdOk

`func (o *SsoConfiguration) GetGoogleClientIdOk() (*string, bool)`

GetGoogleClientIdOk returns a tuple with the GoogleClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleClientId

`func (o *SsoConfiguration) SetGoogleClientId(v string)`

SetGoogleClientId sets GoogleClientId field to given value.


### GetGoogleClientSecret

`func (o *SsoConfiguration) GetGoogleClientSecret() string`

GetGoogleClientSecret returns the GoogleClientSecret field if non-nil, zero value otherwise.

### GetGoogleClientSecretOk

`func (o *SsoConfiguration) GetGoogleClientSecretOk() (*string, bool)`

GetGoogleClientSecretOk returns a tuple with the GoogleClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleClientSecret

`func (o *SsoConfiguration) SetGoogleClientSecret(v string)`

SetGoogleClientSecret sets GoogleClientSecret field to given value.


### GetDisableEmailPasswordLogin

`func (o *SsoConfiguration) GetDisableEmailPasswordLogin() bool`

GetDisableEmailPasswordLogin returns the DisableEmailPasswordLogin field if non-nil, zero value otherwise.

### GetDisableEmailPasswordLoginOk

`func (o *SsoConfiguration) GetDisableEmailPasswordLoginOk() (*bool, bool)`

GetDisableEmailPasswordLoginOk returns a tuple with the DisableEmailPasswordLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmailPasswordLogin

`func (o *SsoConfiguration) SetDisableEmailPasswordLogin(v bool)`

SetDisableEmailPasswordLogin sets DisableEmailPasswordLogin field to given value.


### GetOidcClientId

`func (o *SsoConfiguration) GetOidcClientId() string`

GetOidcClientId returns the OidcClientId field if non-nil, zero value otherwise.

### GetOidcClientIdOk

`func (o *SsoConfiguration) GetOidcClientIdOk() (*string, bool)`

GetOidcClientIdOk returns a tuple with the OidcClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientId

`func (o *SsoConfiguration) SetOidcClientId(v string)`

SetOidcClientId sets OidcClientId field to given value.


### GetOidcClientSecret

`func (o *SsoConfiguration) GetOidcClientSecret() string`

GetOidcClientSecret returns the OidcClientSecret field if non-nil, zero value otherwise.

### GetOidcClientSecretOk

`func (o *SsoConfiguration) GetOidcClientSecretOk() (*string, bool)`

GetOidcClientSecretOk returns a tuple with the OidcClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientSecret

`func (o *SsoConfiguration) SetOidcClientSecret(v string)`

SetOidcClientSecret sets OidcClientSecret field to given value.


### GetOidcScopes

`func (o *SsoConfiguration) GetOidcScopes() string`

GetOidcScopes returns the OidcScopes field if non-nil, zero value otherwise.

### GetOidcScopesOk

`func (o *SsoConfiguration) GetOidcScopesOk() (*string, bool)`

GetOidcScopesOk returns a tuple with the OidcScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcScopes

`func (o *SsoConfiguration) SetOidcScopes(v string)`

SetOidcScopes sets OidcScopes field to given value.


### GetOidcAuthUrl

`func (o *SsoConfiguration) GetOidcAuthUrl() string`

GetOidcAuthUrl returns the OidcAuthUrl field if non-nil, zero value otherwise.

### GetOidcAuthUrlOk

`func (o *SsoConfiguration) GetOidcAuthUrlOk() (*string, bool)`

GetOidcAuthUrlOk returns a tuple with the OidcAuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAuthUrl

`func (o *SsoConfiguration) SetOidcAuthUrl(v string)`

SetOidcAuthUrl sets OidcAuthUrl field to given value.


### GetOidcTokenUrl

`func (o *SsoConfiguration) GetOidcTokenUrl() string`

GetOidcTokenUrl returns the OidcTokenUrl field if non-nil, zero value otherwise.

### GetOidcTokenUrlOk

`func (o *SsoConfiguration) GetOidcTokenUrlOk() (*string, bool)`

GetOidcTokenUrlOk returns a tuple with the OidcTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcTokenUrl

`func (o *SsoConfiguration) SetOidcTokenUrl(v string)`

SetOidcTokenUrl sets OidcTokenUrl field to given value.


### GetOidcUserinfoUrl

`func (o *SsoConfiguration) GetOidcUserinfoUrl() string`

GetOidcUserinfoUrl returns the OidcUserinfoUrl field if non-nil, zero value otherwise.

### GetOidcUserinfoUrlOk

`func (o *SsoConfiguration) GetOidcUserinfoUrlOk() (*string, bool)`

GetOidcUserinfoUrlOk returns a tuple with the OidcUserinfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcUserinfoUrl

`func (o *SsoConfiguration) SetOidcUserinfoUrl(v string)`

SetOidcUserinfoUrl sets OidcUserinfoUrl field to given value.

### HasOidcUserinfoUrl

`func (o *SsoConfiguration) HasOidcUserinfoUrl() bool`

HasOidcUserinfoUrl returns a boolean if a field has been set.

### GetOidcAudience

`func (o *SsoConfiguration) GetOidcAudience() string`

GetOidcAudience returns the OidcAudience field if non-nil, zero value otherwise.

### GetOidcAudienceOk

`func (o *SsoConfiguration) GetOidcAudienceOk() (*string, bool)`

GetOidcAudienceOk returns a tuple with the OidcAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAudience

`func (o *SsoConfiguration) SetOidcAudience(v string)`

SetOidcAudience sets OidcAudience field to given value.

### HasOidcAudience

`func (o *SsoConfiguration) HasOidcAudience() bool`

HasOidcAudience returns a boolean if a field has been set.

### GetJwtEmailKey

`func (o *SsoConfiguration) GetJwtEmailKey() string`

GetJwtEmailKey returns the JwtEmailKey field if non-nil, zero value otherwise.

### GetJwtEmailKeyOk

`func (o *SsoConfiguration) GetJwtEmailKeyOk() (*string, bool)`

GetJwtEmailKeyOk returns a tuple with the JwtEmailKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtEmailKey

`func (o *SsoConfiguration) SetJwtEmailKey(v string)`

SetJwtEmailKey sets JwtEmailKey field to given value.


### GetJwtRolesKey

`func (o *SsoConfiguration) GetJwtRolesKey() string`

GetJwtRolesKey returns the JwtRolesKey field if non-nil, zero value otherwise.

### GetJwtRolesKeyOk

`func (o *SsoConfiguration) GetJwtRolesKeyOk() (*string, bool)`

GetJwtRolesKeyOk returns a tuple with the JwtRolesKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtRolesKey

`func (o *SsoConfiguration) SetJwtRolesKey(v string)`

SetJwtRolesKey sets JwtRolesKey field to given value.

### HasJwtRolesKey

`func (o *SsoConfiguration) HasJwtRolesKey() bool`

HasJwtRolesKey returns a boolean if a field has been set.

### GetJwtFirstNameKey

`func (o *SsoConfiguration) GetJwtFirstNameKey() string`

GetJwtFirstNameKey returns the JwtFirstNameKey field if non-nil, zero value otherwise.

### GetJwtFirstNameKeyOk

`func (o *SsoConfiguration) GetJwtFirstNameKeyOk() (*string, bool)`

GetJwtFirstNameKeyOk returns a tuple with the JwtFirstNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtFirstNameKey

`func (o *SsoConfiguration) SetJwtFirstNameKey(v string)`

SetJwtFirstNameKey sets JwtFirstNameKey field to given value.


### GetJwtLastNameKey

`func (o *SsoConfiguration) GetJwtLastNameKey() string`

GetJwtLastNameKey returns the JwtLastNameKey field if non-nil, zero value otherwise.

### GetJwtLastNameKeyOk

`func (o *SsoConfiguration) GetJwtLastNameKeyOk() (*string, bool)`

GetJwtLastNameKeyOk returns a tuple with the JwtLastNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtLastNameKey

`func (o *SsoConfiguration) SetJwtLastNameKey(v string)`

SetJwtLastNameKey sets JwtLastNameKey field to given value.


### GetRolesMapping

`func (o *SsoConfiguration) GetRolesMapping() string`

GetRolesMapping returns the RolesMapping field if non-nil, zero value otherwise.

### GetRolesMappingOk

`func (o *SsoConfiguration) GetRolesMappingOk() (*string, bool)`

GetRolesMappingOk returns a tuple with the RolesMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolesMapping

`func (o *SsoConfiguration) SetRolesMapping(v string)`

SetRolesMapping sets RolesMapping field to given value.

### HasRolesMapping

`func (o *SsoConfiguration) HasRolesMapping() bool`

HasRolesMapping returns a boolean if a field has been set.

### GetJitEnabled

`func (o *SsoConfiguration) GetJitEnabled() bool`

GetJitEnabled returns the JitEnabled field if non-nil, zero value otherwise.

### GetJitEnabledOk

`func (o *SsoConfiguration) GetJitEnabledOk() (*bool, bool)`

GetJitEnabledOk returns a tuple with the JitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitEnabled

`func (o *SsoConfiguration) SetJitEnabled(v bool)`

SetJitEnabled sets JitEnabled field to given value.


### GetRestrictedDomain

`func (o *SsoConfiguration) GetRestrictedDomain() string`

GetRestrictedDomain returns the RestrictedDomain field if non-nil, zero value otherwise.

### GetRestrictedDomainOk

`func (o *SsoConfiguration) GetRestrictedDomainOk() (*string, bool)`

GetRestrictedDomainOk returns a tuple with the RestrictedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDomain

`func (o *SsoConfiguration) SetRestrictedDomain(v string)`

SetRestrictedDomain sets RestrictedDomain field to given value.

### HasRestrictedDomain

`func (o *SsoConfiguration) HasRestrictedDomain() bool`

HasRestrictedDomain returns a boolean if a field has been set.

### GetTriggerLoginAutomatically

`func (o *SsoConfiguration) GetTriggerLoginAutomatically() bool`

GetTriggerLoginAutomatically returns the TriggerLoginAutomatically field if non-nil, zero value otherwise.

### GetTriggerLoginAutomaticallyOk

`func (o *SsoConfiguration) GetTriggerLoginAutomaticallyOk() (*bool, bool)`

GetTriggerLoginAutomaticallyOk returns a tuple with the TriggerLoginAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerLoginAutomatically

`func (o *SsoConfiguration) SetTriggerLoginAutomatically(v bool)`

SetTriggerLoginAutomatically sets TriggerLoginAutomatically field to given value.


### GetIdpMetadataXml

`func (o *SsoConfiguration) GetIdpMetadataXml() string`

GetIdpMetadataXml returns the IdpMetadataXml field if non-nil, zero value otherwise.

### GetIdpMetadataXmlOk

`func (o *SsoConfiguration) GetIdpMetadataXmlOk() (*string, bool)`

GetIdpMetadataXmlOk returns a tuple with the IdpMetadataXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataXml

`func (o *SsoConfiguration) SetIdpMetadataXml(v string)`

SetIdpMetadataXml sets IdpMetadataXml field to given value.


### GetSamlFirstNameAttribute

`func (o *SsoConfiguration) GetSamlFirstNameAttribute() string`

GetSamlFirstNameAttribute returns the SamlFirstNameAttribute field if non-nil, zero value otherwise.

### GetSamlFirstNameAttributeOk

`func (o *SsoConfiguration) GetSamlFirstNameAttributeOk() (*string, bool)`

GetSamlFirstNameAttributeOk returns a tuple with the SamlFirstNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlFirstNameAttribute

`func (o *SsoConfiguration) SetSamlFirstNameAttribute(v string)`

SetSamlFirstNameAttribute sets SamlFirstNameAttribute field to given value.


### GetSamlLastNameAttribute

`func (o *SsoConfiguration) GetSamlLastNameAttribute() string`

GetSamlLastNameAttribute returns the SamlLastNameAttribute field if non-nil, zero value otherwise.

### GetSamlLastNameAttributeOk

`func (o *SsoConfiguration) GetSamlLastNameAttributeOk() (*string, bool)`

GetSamlLastNameAttributeOk returns a tuple with the SamlLastNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlLastNameAttribute

`func (o *SsoConfiguration) SetSamlLastNameAttribute(v string)`

SetSamlLastNameAttribute sets SamlLastNameAttribute field to given value.


### GetSamlGroupsAttribute

`func (o *SsoConfiguration) GetSamlGroupsAttribute() string`

GetSamlGroupsAttribute returns the SamlGroupsAttribute field if non-nil, zero value otherwise.

### GetSamlGroupsAttributeOk

`func (o *SsoConfiguration) GetSamlGroupsAttributeOk() (*string, bool)`

GetSamlGroupsAttributeOk returns a tuple with the SamlGroupsAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlGroupsAttribute

`func (o *SsoConfiguration) SetSamlGroupsAttribute(v string)`

SetSamlGroupsAttribute sets SamlGroupsAttribute field to given value.

### HasSamlGroupsAttribute

`func (o *SsoConfiguration) HasSamlGroupsAttribute() bool`

HasSamlGroupsAttribute returns a boolean if a field has been set.

### GetSamlSyncGroupClaims

`func (o *SsoConfiguration) GetSamlSyncGroupClaims() bool`

GetSamlSyncGroupClaims returns the SamlSyncGroupClaims field if non-nil, zero value otherwise.

### GetSamlSyncGroupClaimsOk

`func (o *SsoConfiguration) GetSamlSyncGroupClaimsOk() (*bool, bool)`

GetSamlSyncGroupClaimsOk returns a tuple with the SamlSyncGroupClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSyncGroupClaims

`func (o *SsoConfiguration) SetSamlSyncGroupClaims(v bool)`

SetSamlSyncGroupClaims sets SamlSyncGroupClaims field to given value.


### GetLdapSyncGroupClaims

`func (o *SsoConfiguration) GetLdapSyncGroupClaims() bool`

GetLdapSyncGroupClaims returns the LdapSyncGroupClaims field if non-nil, zero value otherwise.

### GetLdapSyncGroupClaimsOk

`func (o *SsoConfiguration) GetLdapSyncGroupClaimsOk() (*bool, bool)`

GetLdapSyncGroupClaimsOk returns a tuple with the LdapSyncGroupClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapSyncGroupClaims

`func (o *SsoConfiguration) SetLdapSyncGroupClaims(v bool)`

SetLdapSyncGroupClaims sets LdapSyncGroupClaims field to given value.

### HasLdapSyncGroupClaims

`func (o *SsoConfiguration) HasLdapSyncGroupClaims() bool`

HasLdapSyncGroupClaims returns a boolean if a field has been set.

### GetLdapRoleMapping

`func (o *SsoConfiguration) GetLdapRoleMapping() string`

GetLdapRoleMapping returns the LdapRoleMapping field if non-nil, zero value otherwise.

### GetLdapRoleMappingOk

`func (o *SsoConfiguration) GetLdapRoleMappingOk() (*string, bool)`

GetLdapRoleMappingOk returns a tuple with the LdapRoleMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapRoleMapping

`func (o *SsoConfiguration) SetLdapRoleMapping(v string)`

SetLdapRoleMapping sets LdapRoleMapping field to given value.

### HasLdapRoleMapping

`func (o *SsoConfiguration) HasLdapRoleMapping() bool`

HasLdapRoleMapping returns a boolean if a field has been set.

### GetLdapServerUrl

`func (o *SsoConfiguration) GetLdapServerUrl() string`

GetLdapServerUrl returns the LdapServerUrl field if non-nil, zero value otherwise.

### GetLdapServerUrlOk

`func (o *SsoConfiguration) GetLdapServerUrlOk() (*string, bool)`

GetLdapServerUrlOk returns a tuple with the LdapServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerUrl

`func (o *SsoConfiguration) SetLdapServerUrl(v string)`

SetLdapServerUrl sets LdapServerUrl field to given value.

### HasLdapServerUrl

`func (o *SsoConfiguration) HasLdapServerUrl() bool`

HasLdapServerUrl returns a boolean if a field has been set.

### GetLdapBaseDomainComponents

`func (o *SsoConfiguration) GetLdapBaseDomainComponents() string`

GetLdapBaseDomainComponents returns the LdapBaseDomainComponents field if non-nil, zero value otherwise.

### GetLdapBaseDomainComponentsOk

`func (o *SsoConfiguration) GetLdapBaseDomainComponentsOk() (*string, bool)`

GetLdapBaseDomainComponentsOk returns a tuple with the LdapBaseDomainComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBaseDomainComponents

`func (o *SsoConfiguration) SetLdapBaseDomainComponents(v string)`

SetLdapBaseDomainComponents sets LdapBaseDomainComponents field to given value.

### HasLdapBaseDomainComponents

`func (o *SsoConfiguration) HasLdapBaseDomainComponents() bool`

HasLdapBaseDomainComponents returns a boolean if a field has been set.

### GetLdapServerName

`func (o *SsoConfiguration) GetLdapServerName() string`

GetLdapServerName returns the LdapServerName field if non-nil, zero value otherwise.

### GetLdapServerNameOk

`func (o *SsoConfiguration) GetLdapServerNameOk() (*string, bool)`

GetLdapServerNameOk returns a tuple with the LdapServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerName

`func (o *SsoConfiguration) SetLdapServerName(v string)`

SetLdapServerName sets LdapServerName field to given value.

### HasLdapServerName

`func (o *SsoConfiguration) HasLdapServerName() bool`

HasLdapServerName returns a boolean if a field has been set.

### GetLdapServerKey

`func (o *SsoConfiguration) GetLdapServerKey() string`

GetLdapServerKey returns the LdapServerKey field if non-nil, zero value otherwise.

### GetLdapServerKeyOk

`func (o *SsoConfiguration) GetLdapServerKeyOk() (*string, bool)`

GetLdapServerKeyOk returns a tuple with the LdapServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerKey

`func (o *SsoConfiguration) SetLdapServerKey(v string)`

SetLdapServerKey sets LdapServerKey field to given value.

### HasLdapServerKey

`func (o *SsoConfiguration) HasLdapServerKey() bool`

HasLdapServerKey returns a boolean if a field has been set.

### GetLdapServerCertificate

`func (o *SsoConfiguration) GetLdapServerCertificate() string`

GetLdapServerCertificate returns the LdapServerCertificate field if non-nil, zero value otherwise.

### GetLdapServerCertificateOk

`func (o *SsoConfiguration) GetLdapServerCertificateOk() (*string, bool)`

GetLdapServerCertificateOk returns a tuple with the LdapServerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerCertificate

`func (o *SsoConfiguration) SetLdapServerCertificate(v string)`

SetLdapServerCertificate sets LdapServerCertificate field to given value.

### HasLdapServerCertificate

`func (o *SsoConfiguration) HasLdapServerCertificate() bool`

HasLdapServerCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


