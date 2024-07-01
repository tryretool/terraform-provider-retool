# SsoConfigPostRequestData

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

### NewSsoConfigPostRequestData

`func NewSsoConfigPostRequestData(configType string, googleClientId string, googleClientSecret string, disableEmailPasswordLogin bool, oidcClientId string, oidcClientSecret string, oidcScopes string, oidcAuthUrl string, oidcTokenUrl string, jwtEmailKey string, jwtFirstNameKey string, jwtLastNameKey string, jitEnabled bool, triggerLoginAutomatically bool, idpMetadataXml string, samlFirstNameAttribute string, samlLastNameAttribute string, samlSyncGroupClaims bool, ) *SsoConfigPostRequestData`

NewSsoConfigPostRequestData instantiates a new SsoConfigPostRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoConfigPostRequestDataWithDefaults

`func NewSsoConfigPostRequestDataWithDefaults() *SsoConfigPostRequestData`

NewSsoConfigPostRequestDataWithDefaults instantiates a new SsoConfigPostRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *SsoConfigPostRequestData) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *SsoConfigPostRequestData) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *SsoConfigPostRequestData) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetGoogleClientId

`func (o *SsoConfigPostRequestData) GetGoogleClientId() string`

GetGoogleClientId returns the GoogleClientId field if non-nil, zero value otherwise.

### GetGoogleClientIdOk

`func (o *SsoConfigPostRequestData) GetGoogleClientIdOk() (*string, bool)`

GetGoogleClientIdOk returns a tuple with the GoogleClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleClientId

`func (o *SsoConfigPostRequestData) SetGoogleClientId(v string)`

SetGoogleClientId sets GoogleClientId field to given value.


### GetGoogleClientSecret

`func (o *SsoConfigPostRequestData) GetGoogleClientSecret() string`

GetGoogleClientSecret returns the GoogleClientSecret field if non-nil, zero value otherwise.

### GetGoogleClientSecretOk

`func (o *SsoConfigPostRequestData) GetGoogleClientSecretOk() (*string, bool)`

GetGoogleClientSecretOk returns a tuple with the GoogleClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleClientSecret

`func (o *SsoConfigPostRequestData) SetGoogleClientSecret(v string)`

SetGoogleClientSecret sets GoogleClientSecret field to given value.


### GetDisableEmailPasswordLogin

`func (o *SsoConfigPostRequestData) GetDisableEmailPasswordLogin() bool`

GetDisableEmailPasswordLogin returns the DisableEmailPasswordLogin field if non-nil, zero value otherwise.

### GetDisableEmailPasswordLoginOk

`func (o *SsoConfigPostRequestData) GetDisableEmailPasswordLoginOk() (*bool, bool)`

GetDisableEmailPasswordLoginOk returns a tuple with the DisableEmailPasswordLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmailPasswordLogin

`func (o *SsoConfigPostRequestData) SetDisableEmailPasswordLogin(v bool)`

SetDisableEmailPasswordLogin sets DisableEmailPasswordLogin field to given value.


### GetOidcClientId

`func (o *SsoConfigPostRequestData) GetOidcClientId() string`

GetOidcClientId returns the OidcClientId field if non-nil, zero value otherwise.

### GetOidcClientIdOk

`func (o *SsoConfigPostRequestData) GetOidcClientIdOk() (*string, bool)`

GetOidcClientIdOk returns a tuple with the OidcClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientId

`func (o *SsoConfigPostRequestData) SetOidcClientId(v string)`

SetOidcClientId sets OidcClientId field to given value.


### GetOidcClientSecret

`func (o *SsoConfigPostRequestData) GetOidcClientSecret() string`

GetOidcClientSecret returns the OidcClientSecret field if non-nil, zero value otherwise.

### GetOidcClientSecretOk

`func (o *SsoConfigPostRequestData) GetOidcClientSecretOk() (*string, bool)`

GetOidcClientSecretOk returns a tuple with the OidcClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientSecret

`func (o *SsoConfigPostRequestData) SetOidcClientSecret(v string)`

SetOidcClientSecret sets OidcClientSecret field to given value.


### GetOidcScopes

`func (o *SsoConfigPostRequestData) GetOidcScopes() string`

GetOidcScopes returns the OidcScopes field if non-nil, zero value otherwise.

### GetOidcScopesOk

`func (o *SsoConfigPostRequestData) GetOidcScopesOk() (*string, bool)`

GetOidcScopesOk returns a tuple with the OidcScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcScopes

`func (o *SsoConfigPostRequestData) SetOidcScopes(v string)`

SetOidcScopes sets OidcScopes field to given value.


### GetOidcAuthUrl

`func (o *SsoConfigPostRequestData) GetOidcAuthUrl() string`

GetOidcAuthUrl returns the OidcAuthUrl field if non-nil, zero value otherwise.

### GetOidcAuthUrlOk

`func (o *SsoConfigPostRequestData) GetOidcAuthUrlOk() (*string, bool)`

GetOidcAuthUrlOk returns a tuple with the OidcAuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAuthUrl

`func (o *SsoConfigPostRequestData) SetOidcAuthUrl(v string)`

SetOidcAuthUrl sets OidcAuthUrl field to given value.


### GetOidcTokenUrl

`func (o *SsoConfigPostRequestData) GetOidcTokenUrl() string`

GetOidcTokenUrl returns the OidcTokenUrl field if non-nil, zero value otherwise.

### GetOidcTokenUrlOk

`func (o *SsoConfigPostRequestData) GetOidcTokenUrlOk() (*string, bool)`

GetOidcTokenUrlOk returns a tuple with the OidcTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcTokenUrl

`func (o *SsoConfigPostRequestData) SetOidcTokenUrl(v string)`

SetOidcTokenUrl sets OidcTokenUrl field to given value.


### GetOidcUserinfoUrl

`func (o *SsoConfigPostRequestData) GetOidcUserinfoUrl() string`

GetOidcUserinfoUrl returns the OidcUserinfoUrl field if non-nil, zero value otherwise.

### GetOidcUserinfoUrlOk

`func (o *SsoConfigPostRequestData) GetOidcUserinfoUrlOk() (*string, bool)`

GetOidcUserinfoUrlOk returns a tuple with the OidcUserinfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcUserinfoUrl

`func (o *SsoConfigPostRequestData) SetOidcUserinfoUrl(v string)`

SetOidcUserinfoUrl sets OidcUserinfoUrl field to given value.

### HasOidcUserinfoUrl

`func (o *SsoConfigPostRequestData) HasOidcUserinfoUrl() bool`

HasOidcUserinfoUrl returns a boolean if a field has been set.

### GetOidcAudience

`func (o *SsoConfigPostRequestData) GetOidcAudience() string`

GetOidcAudience returns the OidcAudience field if non-nil, zero value otherwise.

### GetOidcAudienceOk

`func (o *SsoConfigPostRequestData) GetOidcAudienceOk() (*string, bool)`

GetOidcAudienceOk returns a tuple with the OidcAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAudience

`func (o *SsoConfigPostRequestData) SetOidcAudience(v string)`

SetOidcAudience sets OidcAudience field to given value.

### HasOidcAudience

`func (o *SsoConfigPostRequestData) HasOidcAudience() bool`

HasOidcAudience returns a boolean if a field has been set.

### GetJwtEmailKey

`func (o *SsoConfigPostRequestData) GetJwtEmailKey() string`

GetJwtEmailKey returns the JwtEmailKey field if non-nil, zero value otherwise.

### GetJwtEmailKeyOk

`func (o *SsoConfigPostRequestData) GetJwtEmailKeyOk() (*string, bool)`

GetJwtEmailKeyOk returns a tuple with the JwtEmailKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtEmailKey

`func (o *SsoConfigPostRequestData) SetJwtEmailKey(v string)`

SetJwtEmailKey sets JwtEmailKey field to given value.


### GetJwtRolesKey

`func (o *SsoConfigPostRequestData) GetJwtRolesKey() string`

GetJwtRolesKey returns the JwtRolesKey field if non-nil, zero value otherwise.

### GetJwtRolesKeyOk

`func (o *SsoConfigPostRequestData) GetJwtRolesKeyOk() (*string, bool)`

GetJwtRolesKeyOk returns a tuple with the JwtRolesKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtRolesKey

`func (o *SsoConfigPostRequestData) SetJwtRolesKey(v string)`

SetJwtRolesKey sets JwtRolesKey field to given value.

### HasJwtRolesKey

`func (o *SsoConfigPostRequestData) HasJwtRolesKey() bool`

HasJwtRolesKey returns a boolean if a field has been set.

### GetJwtFirstNameKey

`func (o *SsoConfigPostRequestData) GetJwtFirstNameKey() string`

GetJwtFirstNameKey returns the JwtFirstNameKey field if non-nil, zero value otherwise.

### GetJwtFirstNameKeyOk

`func (o *SsoConfigPostRequestData) GetJwtFirstNameKeyOk() (*string, bool)`

GetJwtFirstNameKeyOk returns a tuple with the JwtFirstNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtFirstNameKey

`func (o *SsoConfigPostRequestData) SetJwtFirstNameKey(v string)`

SetJwtFirstNameKey sets JwtFirstNameKey field to given value.


### GetJwtLastNameKey

`func (o *SsoConfigPostRequestData) GetJwtLastNameKey() string`

GetJwtLastNameKey returns the JwtLastNameKey field if non-nil, zero value otherwise.

### GetJwtLastNameKeyOk

`func (o *SsoConfigPostRequestData) GetJwtLastNameKeyOk() (*string, bool)`

GetJwtLastNameKeyOk returns a tuple with the JwtLastNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtLastNameKey

`func (o *SsoConfigPostRequestData) SetJwtLastNameKey(v string)`

SetJwtLastNameKey sets JwtLastNameKey field to given value.


### GetRolesMapping

`func (o *SsoConfigPostRequestData) GetRolesMapping() string`

GetRolesMapping returns the RolesMapping field if non-nil, zero value otherwise.

### GetRolesMappingOk

`func (o *SsoConfigPostRequestData) GetRolesMappingOk() (*string, bool)`

GetRolesMappingOk returns a tuple with the RolesMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolesMapping

`func (o *SsoConfigPostRequestData) SetRolesMapping(v string)`

SetRolesMapping sets RolesMapping field to given value.

### HasRolesMapping

`func (o *SsoConfigPostRequestData) HasRolesMapping() bool`

HasRolesMapping returns a boolean if a field has been set.

### GetJitEnabled

`func (o *SsoConfigPostRequestData) GetJitEnabled() bool`

GetJitEnabled returns the JitEnabled field if non-nil, zero value otherwise.

### GetJitEnabledOk

`func (o *SsoConfigPostRequestData) GetJitEnabledOk() (*bool, bool)`

GetJitEnabledOk returns a tuple with the JitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitEnabled

`func (o *SsoConfigPostRequestData) SetJitEnabled(v bool)`

SetJitEnabled sets JitEnabled field to given value.


### GetRestrictedDomain

`func (o *SsoConfigPostRequestData) GetRestrictedDomain() string`

GetRestrictedDomain returns the RestrictedDomain field if non-nil, zero value otherwise.

### GetRestrictedDomainOk

`func (o *SsoConfigPostRequestData) GetRestrictedDomainOk() (*string, bool)`

GetRestrictedDomainOk returns a tuple with the RestrictedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDomain

`func (o *SsoConfigPostRequestData) SetRestrictedDomain(v string)`

SetRestrictedDomain sets RestrictedDomain field to given value.

### HasRestrictedDomain

`func (o *SsoConfigPostRequestData) HasRestrictedDomain() bool`

HasRestrictedDomain returns a boolean if a field has been set.

### GetTriggerLoginAutomatically

`func (o *SsoConfigPostRequestData) GetTriggerLoginAutomatically() bool`

GetTriggerLoginAutomatically returns the TriggerLoginAutomatically field if non-nil, zero value otherwise.

### GetTriggerLoginAutomaticallyOk

`func (o *SsoConfigPostRequestData) GetTriggerLoginAutomaticallyOk() (*bool, bool)`

GetTriggerLoginAutomaticallyOk returns a tuple with the TriggerLoginAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerLoginAutomatically

`func (o *SsoConfigPostRequestData) SetTriggerLoginAutomatically(v bool)`

SetTriggerLoginAutomatically sets TriggerLoginAutomatically field to given value.


### GetIdpMetadataXml

`func (o *SsoConfigPostRequestData) GetIdpMetadataXml() string`

GetIdpMetadataXml returns the IdpMetadataXml field if non-nil, zero value otherwise.

### GetIdpMetadataXmlOk

`func (o *SsoConfigPostRequestData) GetIdpMetadataXmlOk() (*string, bool)`

GetIdpMetadataXmlOk returns a tuple with the IdpMetadataXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataXml

`func (o *SsoConfigPostRequestData) SetIdpMetadataXml(v string)`

SetIdpMetadataXml sets IdpMetadataXml field to given value.


### GetSamlFirstNameAttribute

`func (o *SsoConfigPostRequestData) GetSamlFirstNameAttribute() string`

GetSamlFirstNameAttribute returns the SamlFirstNameAttribute field if non-nil, zero value otherwise.

### GetSamlFirstNameAttributeOk

`func (o *SsoConfigPostRequestData) GetSamlFirstNameAttributeOk() (*string, bool)`

GetSamlFirstNameAttributeOk returns a tuple with the SamlFirstNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlFirstNameAttribute

`func (o *SsoConfigPostRequestData) SetSamlFirstNameAttribute(v string)`

SetSamlFirstNameAttribute sets SamlFirstNameAttribute field to given value.


### GetSamlLastNameAttribute

`func (o *SsoConfigPostRequestData) GetSamlLastNameAttribute() string`

GetSamlLastNameAttribute returns the SamlLastNameAttribute field if non-nil, zero value otherwise.

### GetSamlLastNameAttributeOk

`func (o *SsoConfigPostRequestData) GetSamlLastNameAttributeOk() (*string, bool)`

GetSamlLastNameAttributeOk returns a tuple with the SamlLastNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlLastNameAttribute

`func (o *SsoConfigPostRequestData) SetSamlLastNameAttribute(v string)`

SetSamlLastNameAttribute sets SamlLastNameAttribute field to given value.


### GetSamlGroupsAttribute

`func (o *SsoConfigPostRequestData) GetSamlGroupsAttribute() string`

GetSamlGroupsAttribute returns the SamlGroupsAttribute field if non-nil, zero value otherwise.

### GetSamlGroupsAttributeOk

`func (o *SsoConfigPostRequestData) GetSamlGroupsAttributeOk() (*string, bool)`

GetSamlGroupsAttributeOk returns a tuple with the SamlGroupsAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlGroupsAttribute

`func (o *SsoConfigPostRequestData) SetSamlGroupsAttribute(v string)`

SetSamlGroupsAttribute sets SamlGroupsAttribute field to given value.

### HasSamlGroupsAttribute

`func (o *SsoConfigPostRequestData) HasSamlGroupsAttribute() bool`

HasSamlGroupsAttribute returns a boolean if a field has been set.

### GetSamlSyncGroupClaims

`func (o *SsoConfigPostRequestData) GetSamlSyncGroupClaims() bool`

GetSamlSyncGroupClaims returns the SamlSyncGroupClaims field if non-nil, zero value otherwise.

### GetSamlSyncGroupClaimsOk

`func (o *SsoConfigPostRequestData) GetSamlSyncGroupClaimsOk() (*bool, bool)`

GetSamlSyncGroupClaimsOk returns a tuple with the SamlSyncGroupClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSyncGroupClaims

`func (o *SsoConfigPostRequestData) SetSamlSyncGroupClaims(v bool)`

SetSamlSyncGroupClaims sets SamlSyncGroupClaims field to given value.


### GetLdapSyncGroupClaims

`func (o *SsoConfigPostRequestData) GetLdapSyncGroupClaims() bool`

GetLdapSyncGroupClaims returns the LdapSyncGroupClaims field if non-nil, zero value otherwise.

### GetLdapSyncGroupClaimsOk

`func (o *SsoConfigPostRequestData) GetLdapSyncGroupClaimsOk() (*bool, bool)`

GetLdapSyncGroupClaimsOk returns a tuple with the LdapSyncGroupClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapSyncGroupClaims

`func (o *SsoConfigPostRequestData) SetLdapSyncGroupClaims(v bool)`

SetLdapSyncGroupClaims sets LdapSyncGroupClaims field to given value.

### HasLdapSyncGroupClaims

`func (o *SsoConfigPostRequestData) HasLdapSyncGroupClaims() bool`

HasLdapSyncGroupClaims returns a boolean if a field has been set.

### GetLdapRoleMapping

`func (o *SsoConfigPostRequestData) GetLdapRoleMapping() string`

GetLdapRoleMapping returns the LdapRoleMapping field if non-nil, zero value otherwise.

### GetLdapRoleMappingOk

`func (o *SsoConfigPostRequestData) GetLdapRoleMappingOk() (*string, bool)`

GetLdapRoleMappingOk returns a tuple with the LdapRoleMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapRoleMapping

`func (o *SsoConfigPostRequestData) SetLdapRoleMapping(v string)`

SetLdapRoleMapping sets LdapRoleMapping field to given value.

### HasLdapRoleMapping

`func (o *SsoConfigPostRequestData) HasLdapRoleMapping() bool`

HasLdapRoleMapping returns a boolean if a field has been set.

### GetLdapServerUrl

`func (o *SsoConfigPostRequestData) GetLdapServerUrl() string`

GetLdapServerUrl returns the LdapServerUrl field if non-nil, zero value otherwise.

### GetLdapServerUrlOk

`func (o *SsoConfigPostRequestData) GetLdapServerUrlOk() (*string, bool)`

GetLdapServerUrlOk returns a tuple with the LdapServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerUrl

`func (o *SsoConfigPostRequestData) SetLdapServerUrl(v string)`

SetLdapServerUrl sets LdapServerUrl field to given value.

### HasLdapServerUrl

`func (o *SsoConfigPostRequestData) HasLdapServerUrl() bool`

HasLdapServerUrl returns a boolean if a field has been set.

### GetLdapBaseDomainComponents

`func (o *SsoConfigPostRequestData) GetLdapBaseDomainComponents() string`

GetLdapBaseDomainComponents returns the LdapBaseDomainComponents field if non-nil, zero value otherwise.

### GetLdapBaseDomainComponentsOk

`func (o *SsoConfigPostRequestData) GetLdapBaseDomainComponentsOk() (*string, bool)`

GetLdapBaseDomainComponentsOk returns a tuple with the LdapBaseDomainComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBaseDomainComponents

`func (o *SsoConfigPostRequestData) SetLdapBaseDomainComponents(v string)`

SetLdapBaseDomainComponents sets LdapBaseDomainComponents field to given value.

### HasLdapBaseDomainComponents

`func (o *SsoConfigPostRequestData) HasLdapBaseDomainComponents() bool`

HasLdapBaseDomainComponents returns a boolean if a field has been set.

### GetLdapServerName

`func (o *SsoConfigPostRequestData) GetLdapServerName() string`

GetLdapServerName returns the LdapServerName field if non-nil, zero value otherwise.

### GetLdapServerNameOk

`func (o *SsoConfigPostRequestData) GetLdapServerNameOk() (*string, bool)`

GetLdapServerNameOk returns a tuple with the LdapServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerName

`func (o *SsoConfigPostRequestData) SetLdapServerName(v string)`

SetLdapServerName sets LdapServerName field to given value.

### HasLdapServerName

`func (o *SsoConfigPostRequestData) HasLdapServerName() bool`

HasLdapServerName returns a boolean if a field has been set.

### GetLdapServerKey

`func (o *SsoConfigPostRequestData) GetLdapServerKey() string`

GetLdapServerKey returns the LdapServerKey field if non-nil, zero value otherwise.

### GetLdapServerKeyOk

`func (o *SsoConfigPostRequestData) GetLdapServerKeyOk() (*string, bool)`

GetLdapServerKeyOk returns a tuple with the LdapServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerKey

`func (o *SsoConfigPostRequestData) SetLdapServerKey(v string)`

SetLdapServerKey sets LdapServerKey field to given value.

### HasLdapServerKey

`func (o *SsoConfigPostRequestData) HasLdapServerKey() bool`

HasLdapServerKey returns a boolean if a field has been set.

### GetLdapServerCertificate

`func (o *SsoConfigPostRequestData) GetLdapServerCertificate() string`

GetLdapServerCertificate returns the LdapServerCertificate field if non-nil, zero value otherwise.

### GetLdapServerCertificateOk

`func (o *SsoConfigPostRequestData) GetLdapServerCertificateOk() (*string, bool)`

GetLdapServerCertificateOk returns a tuple with the LdapServerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerCertificate

`func (o *SsoConfigPostRequestData) SetLdapServerCertificate(v string)`

SetLdapServerCertificate sets LdapServerCertificate field to given value.

### HasLdapServerCertificate

`func (o *SsoConfigPostRequestData) HasLdapServerCertificate() bool`

HasLdapServerCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


