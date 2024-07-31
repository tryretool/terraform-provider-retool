# SsoConfigGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** |  | 
**GoogleClientId** | Pointer to **string** |  | [optional] 
**GoogleClientSecret** | Pointer to **string** |  | [optional] 
**DisableEmailPasswordLogin** | **bool** |  | 
**OidcClientId** | Pointer to **string** |  | [optional] 
**OidcClientSecret** | Pointer to **string** |  | [optional] 
**OidcScopes** | Pointer to **string** |  | [optional] 
**OidcAuthUrl** | Pointer to **string** |  | [optional] 
**OidcTokenUrl** | Pointer to **string** |  | [optional] 
**OidcUserinfoUrl** | Pointer to **string** |  | [optional] 
**OidcAudience** | Pointer to **string** |  | [optional] 
**JwtEmailKey** | Pointer to **string** |  | [optional] 
**JwtRolesKey** | Pointer to **string** |  | [optional] 
**JwtFirstNameKey** | Pointer to **string** |  | [optional] 
**JwtLastNameKey** | Pointer to **string** |  | [optional] 
**RolesMapping** | Pointer to **string** |  | [optional] 
**JitEnabled** | Pointer to **bool** |  | [optional] 
**RestrictedDomain** | Pointer to **string** |  | [optional] 
**TriggerLoginAutomatically** | Pointer to **bool** |  | [optional] 
**IdpMetadataXml** | Pointer to **string** |  | [optional] 
**SamlFirstNameAttribute** | Pointer to **string** |  | [optional] 
**SamlLastNameAttribute** | Pointer to **string** |  | [optional] 
**SamlGroupsAttribute** | Pointer to **string** |  | [optional] 
**SamlSyncGroupClaims** | Pointer to **bool** |  | [optional] 
**LdapSyncGroupClaims** | Pointer to **bool** |  | [optional] 
**LdapRoleMapping** | Pointer to **string** |  | [optional] 
**LdapServerUrl** | Pointer to **string** |  | [optional] 
**LdapBaseDomainComponents** | Pointer to **string** |  | [optional] 
**LdapServerName** | Pointer to **string** |  | [optional] 
**LdapServerKey** | Pointer to **string** |  | [optional] 
**LdapServerCertificate** | Pointer to **string** |  | [optional] 

## Methods

### NewSsoConfigGet200ResponseData

`func NewSsoConfigGet200ResponseData(configType string, disableEmailPasswordLogin bool, ) *SsoConfigGet200ResponseData`

NewSsoConfigGet200ResponseData instantiates a new SsoConfigGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoConfigGet200ResponseDataWithDefaults

`func NewSsoConfigGet200ResponseDataWithDefaults() *SsoConfigGet200ResponseData`

NewSsoConfigGet200ResponseDataWithDefaults instantiates a new SsoConfigGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *SsoConfigGet200ResponseData) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *SsoConfigGet200ResponseData) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *SsoConfigGet200ResponseData) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetGoogleClientId

`func (o *SsoConfigGet200ResponseData) GetGoogleClientId() string`

GetGoogleClientId returns the GoogleClientId field if non-nil, zero value otherwise.

### GetGoogleClientIdOk

`func (o *SsoConfigGet200ResponseData) GetGoogleClientIdOk() (*string, bool)`

GetGoogleClientIdOk returns a tuple with the GoogleClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleClientId

`func (o *SsoConfigGet200ResponseData) SetGoogleClientId(v string)`

SetGoogleClientId sets GoogleClientId field to given value.

### HasGoogleClientId

`func (o *SsoConfigGet200ResponseData) HasGoogleClientId() bool`

HasGoogleClientId returns a boolean if a field has been set.

### GetGoogleClientSecret

`func (o *SsoConfigGet200ResponseData) GetGoogleClientSecret() string`

GetGoogleClientSecret returns the GoogleClientSecret field if non-nil, zero value otherwise.

### GetGoogleClientSecretOk

`func (o *SsoConfigGet200ResponseData) GetGoogleClientSecretOk() (*string, bool)`

GetGoogleClientSecretOk returns a tuple with the GoogleClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleClientSecret

`func (o *SsoConfigGet200ResponseData) SetGoogleClientSecret(v string)`

SetGoogleClientSecret sets GoogleClientSecret field to given value.

### HasGoogleClientSecret

`func (o *SsoConfigGet200ResponseData) HasGoogleClientSecret() bool`

HasGoogleClientSecret returns a boolean if a field has been set.

### GetDisableEmailPasswordLogin

`func (o *SsoConfigGet200ResponseData) GetDisableEmailPasswordLogin() bool`

GetDisableEmailPasswordLogin returns the DisableEmailPasswordLogin field if non-nil, zero value otherwise.

### GetDisableEmailPasswordLoginOk

`func (o *SsoConfigGet200ResponseData) GetDisableEmailPasswordLoginOk() (*bool, bool)`

GetDisableEmailPasswordLoginOk returns a tuple with the DisableEmailPasswordLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmailPasswordLogin

`func (o *SsoConfigGet200ResponseData) SetDisableEmailPasswordLogin(v bool)`

SetDisableEmailPasswordLogin sets DisableEmailPasswordLogin field to given value.


### GetOidcClientId

`func (o *SsoConfigGet200ResponseData) GetOidcClientId() string`

GetOidcClientId returns the OidcClientId field if non-nil, zero value otherwise.

### GetOidcClientIdOk

`func (o *SsoConfigGet200ResponseData) GetOidcClientIdOk() (*string, bool)`

GetOidcClientIdOk returns a tuple with the OidcClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientId

`func (o *SsoConfigGet200ResponseData) SetOidcClientId(v string)`

SetOidcClientId sets OidcClientId field to given value.

### HasOidcClientId

`func (o *SsoConfigGet200ResponseData) HasOidcClientId() bool`

HasOidcClientId returns a boolean if a field has been set.

### GetOidcClientSecret

`func (o *SsoConfigGet200ResponseData) GetOidcClientSecret() string`

GetOidcClientSecret returns the OidcClientSecret field if non-nil, zero value otherwise.

### GetOidcClientSecretOk

`func (o *SsoConfigGet200ResponseData) GetOidcClientSecretOk() (*string, bool)`

GetOidcClientSecretOk returns a tuple with the OidcClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientSecret

`func (o *SsoConfigGet200ResponseData) SetOidcClientSecret(v string)`

SetOidcClientSecret sets OidcClientSecret field to given value.

### HasOidcClientSecret

`func (o *SsoConfigGet200ResponseData) HasOidcClientSecret() bool`

HasOidcClientSecret returns a boolean if a field has been set.

### GetOidcScopes

`func (o *SsoConfigGet200ResponseData) GetOidcScopes() string`

GetOidcScopes returns the OidcScopes field if non-nil, zero value otherwise.

### GetOidcScopesOk

`func (o *SsoConfigGet200ResponseData) GetOidcScopesOk() (*string, bool)`

GetOidcScopesOk returns a tuple with the OidcScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcScopes

`func (o *SsoConfigGet200ResponseData) SetOidcScopes(v string)`

SetOidcScopes sets OidcScopes field to given value.

### HasOidcScopes

`func (o *SsoConfigGet200ResponseData) HasOidcScopes() bool`

HasOidcScopes returns a boolean if a field has been set.

### GetOidcAuthUrl

`func (o *SsoConfigGet200ResponseData) GetOidcAuthUrl() string`

GetOidcAuthUrl returns the OidcAuthUrl field if non-nil, zero value otherwise.

### GetOidcAuthUrlOk

`func (o *SsoConfigGet200ResponseData) GetOidcAuthUrlOk() (*string, bool)`

GetOidcAuthUrlOk returns a tuple with the OidcAuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAuthUrl

`func (o *SsoConfigGet200ResponseData) SetOidcAuthUrl(v string)`

SetOidcAuthUrl sets OidcAuthUrl field to given value.

### HasOidcAuthUrl

`func (o *SsoConfigGet200ResponseData) HasOidcAuthUrl() bool`

HasOidcAuthUrl returns a boolean if a field has been set.

### GetOidcTokenUrl

`func (o *SsoConfigGet200ResponseData) GetOidcTokenUrl() string`

GetOidcTokenUrl returns the OidcTokenUrl field if non-nil, zero value otherwise.

### GetOidcTokenUrlOk

`func (o *SsoConfigGet200ResponseData) GetOidcTokenUrlOk() (*string, bool)`

GetOidcTokenUrlOk returns a tuple with the OidcTokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcTokenUrl

`func (o *SsoConfigGet200ResponseData) SetOidcTokenUrl(v string)`

SetOidcTokenUrl sets OidcTokenUrl field to given value.

### HasOidcTokenUrl

`func (o *SsoConfigGet200ResponseData) HasOidcTokenUrl() bool`

HasOidcTokenUrl returns a boolean if a field has been set.

### GetOidcUserinfoUrl

`func (o *SsoConfigGet200ResponseData) GetOidcUserinfoUrl() string`

GetOidcUserinfoUrl returns the OidcUserinfoUrl field if non-nil, zero value otherwise.

### GetOidcUserinfoUrlOk

`func (o *SsoConfigGet200ResponseData) GetOidcUserinfoUrlOk() (*string, bool)`

GetOidcUserinfoUrlOk returns a tuple with the OidcUserinfoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcUserinfoUrl

`func (o *SsoConfigGet200ResponseData) SetOidcUserinfoUrl(v string)`

SetOidcUserinfoUrl sets OidcUserinfoUrl field to given value.

### HasOidcUserinfoUrl

`func (o *SsoConfigGet200ResponseData) HasOidcUserinfoUrl() bool`

HasOidcUserinfoUrl returns a boolean if a field has been set.

### GetOidcAudience

`func (o *SsoConfigGet200ResponseData) GetOidcAudience() string`

GetOidcAudience returns the OidcAudience field if non-nil, zero value otherwise.

### GetOidcAudienceOk

`func (o *SsoConfigGet200ResponseData) GetOidcAudienceOk() (*string, bool)`

GetOidcAudienceOk returns a tuple with the OidcAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAudience

`func (o *SsoConfigGet200ResponseData) SetOidcAudience(v string)`

SetOidcAudience sets OidcAudience field to given value.

### HasOidcAudience

`func (o *SsoConfigGet200ResponseData) HasOidcAudience() bool`

HasOidcAudience returns a boolean if a field has been set.

### GetJwtEmailKey

`func (o *SsoConfigGet200ResponseData) GetJwtEmailKey() string`

GetJwtEmailKey returns the JwtEmailKey field if non-nil, zero value otherwise.

### GetJwtEmailKeyOk

`func (o *SsoConfigGet200ResponseData) GetJwtEmailKeyOk() (*string, bool)`

GetJwtEmailKeyOk returns a tuple with the JwtEmailKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtEmailKey

`func (o *SsoConfigGet200ResponseData) SetJwtEmailKey(v string)`

SetJwtEmailKey sets JwtEmailKey field to given value.

### HasJwtEmailKey

`func (o *SsoConfigGet200ResponseData) HasJwtEmailKey() bool`

HasJwtEmailKey returns a boolean if a field has been set.

### GetJwtRolesKey

`func (o *SsoConfigGet200ResponseData) GetJwtRolesKey() string`

GetJwtRolesKey returns the JwtRolesKey field if non-nil, zero value otherwise.

### GetJwtRolesKeyOk

`func (o *SsoConfigGet200ResponseData) GetJwtRolesKeyOk() (*string, bool)`

GetJwtRolesKeyOk returns a tuple with the JwtRolesKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtRolesKey

`func (o *SsoConfigGet200ResponseData) SetJwtRolesKey(v string)`

SetJwtRolesKey sets JwtRolesKey field to given value.

### HasJwtRolesKey

`func (o *SsoConfigGet200ResponseData) HasJwtRolesKey() bool`

HasJwtRolesKey returns a boolean if a field has been set.

### GetJwtFirstNameKey

`func (o *SsoConfigGet200ResponseData) GetJwtFirstNameKey() string`

GetJwtFirstNameKey returns the JwtFirstNameKey field if non-nil, zero value otherwise.

### GetJwtFirstNameKeyOk

`func (o *SsoConfigGet200ResponseData) GetJwtFirstNameKeyOk() (*string, bool)`

GetJwtFirstNameKeyOk returns a tuple with the JwtFirstNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtFirstNameKey

`func (o *SsoConfigGet200ResponseData) SetJwtFirstNameKey(v string)`

SetJwtFirstNameKey sets JwtFirstNameKey field to given value.

### HasJwtFirstNameKey

`func (o *SsoConfigGet200ResponseData) HasJwtFirstNameKey() bool`

HasJwtFirstNameKey returns a boolean if a field has been set.

### GetJwtLastNameKey

`func (o *SsoConfigGet200ResponseData) GetJwtLastNameKey() string`

GetJwtLastNameKey returns the JwtLastNameKey field if non-nil, zero value otherwise.

### GetJwtLastNameKeyOk

`func (o *SsoConfigGet200ResponseData) GetJwtLastNameKeyOk() (*string, bool)`

GetJwtLastNameKeyOk returns a tuple with the JwtLastNameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtLastNameKey

`func (o *SsoConfigGet200ResponseData) SetJwtLastNameKey(v string)`

SetJwtLastNameKey sets JwtLastNameKey field to given value.

### HasJwtLastNameKey

`func (o *SsoConfigGet200ResponseData) HasJwtLastNameKey() bool`

HasJwtLastNameKey returns a boolean if a field has been set.

### GetRolesMapping

`func (o *SsoConfigGet200ResponseData) GetRolesMapping() string`

GetRolesMapping returns the RolesMapping field if non-nil, zero value otherwise.

### GetRolesMappingOk

`func (o *SsoConfigGet200ResponseData) GetRolesMappingOk() (*string, bool)`

GetRolesMappingOk returns a tuple with the RolesMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolesMapping

`func (o *SsoConfigGet200ResponseData) SetRolesMapping(v string)`

SetRolesMapping sets RolesMapping field to given value.

### HasRolesMapping

`func (o *SsoConfigGet200ResponseData) HasRolesMapping() bool`

HasRolesMapping returns a boolean if a field has been set.

### GetJitEnabled

`func (o *SsoConfigGet200ResponseData) GetJitEnabled() bool`

GetJitEnabled returns the JitEnabled field if non-nil, zero value otherwise.

### GetJitEnabledOk

`func (o *SsoConfigGet200ResponseData) GetJitEnabledOk() (*bool, bool)`

GetJitEnabledOk returns a tuple with the JitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitEnabled

`func (o *SsoConfigGet200ResponseData) SetJitEnabled(v bool)`

SetJitEnabled sets JitEnabled field to given value.

### HasJitEnabled

`func (o *SsoConfigGet200ResponseData) HasJitEnabled() bool`

HasJitEnabled returns a boolean if a field has been set.

### GetRestrictedDomain

`func (o *SsoConfigGet200ResponseData) GetRestrictedDomain() string`

GetRestrictedDomain returns the RestrictedDomain field if non-nil, zero value otherwise.

### GetRestrictedDomainOk

`func (o *SsoConfigGet200ResponseData) GetRestrictedDomainOk() (*string, bool)`

GetRestrictedDomainOk returns a tuple with the RestrictedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDomain

`func (o *SsoConfigGet200ResponseData) SetRestrictedDomain(v string)`

SetRestrictedDomain sets RestrictedDomain field to given value.

### HasRestrictedDomain

`func (o *SsoConfigGet200ResponseData) HasRestrictedDomain() bool`

HasRestrictedDomain returns a boolean if a field has been set.

### GetTriggerLoginAutomatically

`func (o *SsoConfigGet200ResponseData) GetTriggerLoginAutomatically() bool`

GetTriggerLoginAutomatically returns the TriggerLoginAutomatically field if non-nil, zero value otherwise.

### GetTriggerLoginAutomaticallyOk

`func (o *SsoConfigGet200ResponseData) GetTriggerLoginAutomaticallyOk() (*bool, bool)`

GetTriggerLoginAutomaticallyOk returns a tuple with the TriggerLoginAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerLoginAutomatically

`func (o *SsoConfigGet200ResponseData) SetTriggerLoginAutomatically(v bool)`

SetTriggerLoginAutomatically sets TriggerLoginAutomatically field to given value.

### HasTriggerLoginAutomatically

`func (o *SsoConfigGet200ResponseData) HasTriggerLoginAutomatically() bool`

HasTriggerLoginAutomatically returns a boolean if a field has been set.

### GetIdpMetadataXml

`func (o *SsoConfigGet200ResponseData) GetIdpMetadataXml() string`

GetIdpMetadataXml returns the IdpMetadataXml field if non-nil, zero value otherwise.

### GetIdpMetadataXmlOk

`func (o *SsoConfigGet200ResponseData) GetIdpMetadataXmlOk() (*string, bool)`

GetIdpMetadataXmlOk returns a tuple with the IdpMetadataXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataXml

`func (o *SsoConfigGet200ResponseData) SetIdpMetadataXml(v string)`

SetIdpMetadataXml sets IdpMetadataXml field to given value.

### HasIdpMetadataXml

`func (o *SsoConfigGet200ResponseData) HasIdpMetadataXml() bool`

HasIdpMetadataXml returns a boolean if a field has been set.

### GetSamlFirstNameAttribute

`func (o *SsoConfigGet200ResponseData) GetSamlFirstNameAttribute() string`

GetSamlFirstNameAttribute returns the SamlFirstNameAttribute field if non-nil, zero value otherwise.

### GetSamlFirstNameAttributeOk

`func (o *SsoConfigGet200ResponseData) GetSamlFirstNameAttributeOk() (*string, bool)`

GetSamlFirstNameAttributeOk returns a tuple with the SamlFirstNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlFirstNameAttribute

`func (o *SsoConfigGet200ResponseData) SetSamlFirstNameAttribute(v string)`

SetSamlFirstNameAttribute sets SamlFirstNameAttribute field to given value.

### HasSamlFirstNameAttribute

`func (o *SsoConfigGet200ResponseData) HasSamlFirstNameAttribute() bool`

HasSamlFirstNameAttribute returns a boolean if a field has been set.

### GetSamlLastNameAttribute

`func (o *SsoConfigGet200ResponseData) GetSamlLastNameAttribute() string`

GetSamlLastNameAttribute returns the SamlLastNameAttribute field if non-nil, zero value otherwise.

### GetSamlLastNameAttributeOk

`func (o *SsoConfigGet200ResponseData) GetSamlLastNameAttributeOk() (*string, bool)`

GetSamlLastNameAttributeOk returns a tuple with the SamlLastNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlLastNameAttribute

`func (o *SsoConfigGet200ResponseData) SetSamlLastNameAttribute(v string)`

SetSamlLastNameAttribute sets SamlLastNameAttribute field to given value.

### HasSamlLastNameAttribute

`func (o *SsoConfigGet200ResponseData) HasSamlLastNameAttribute() bool`

HasSamlLastNameAttribute returns a boolean if a field has been set.

### GetSamlGroupsAttribute

`func (o *SsoConfigGet200ResponseData) GetSamlGroupsAttribute() string`

GetSamlGroupsAttribute returns the SamlGroupsAttribute field if non-nil, zero value otherwise.

### GetSamlGroupsAttributeOk

`func (o *SsoConfigGet200ResponseData) GetSamlGroupsAttributeOk() (*string, bool)`

GetSamlGroupsAttributeOk returns a tuple with the SamlGroupsAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlGroupsAttribute

`func (o *SsoConfigGet200ResponseData) SetSamlGroupsAttribute(v string)`

SetSamlGroupsAttribute sets SamlGroupsAttribute field to given value.

### HasSamlGroupsAttribute

`func (o *SsoConfigGet200ResponseData) HasSamlGroupsAttribute() bool`

HasSamlGroupsAttribute returns a boolean if a field has been set.

### GetSamlSyncGroupClaims

`func (o *SsoConfigGet200ResponseData) GetSamlSyncGroupClaims() bool`

GetSamlSyncGroupClaims returns the SamlSyncGroupClaims field if non-nil, zero value otherwise.

### GetSamlSyncGroupClaimsOk

`func (o *SsoConfigGet200ResponseData) GetSamlSyncGroupClaimsOk() (*bool, bool)`

GetSamlSyncGroupClaimsOk returns a tuple with the SamlSyncGroupClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSyncGroupClaims

`func (o *SsoConfigGet200ResponseData) SetSamlSyncGroupClaims(v bool)`

SetSamlSyncGroupClaims sets SamlSyncGroupClaims field to given value.

### HasSamlSyncGroupClaims

`func (o *SsoConfigGet200ResponseData) HasSamlSyncGroupClaims() bool`

HasSamlSyncGroupClaims returns a boolean if a field has been set.

### GetLdapSyncGroupClaims

`func (o *SsoConfigGet200ResponseData) GetLdapSyncGroupClaims() bool`

GetLdapSyncGroupClaims returns the LdapSyncGroupClaims field if non-nil, zero value otherwise.

### GetLdapSyncGroupClaimsOk

`func (o *SsoConfigGet200ResponseData) GetLdapSyncGroupClaimsOk() (*bool, bool)`

GetLdapSyncGroupClaimsOk returns a tuple with the LdapSyncGroupClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapSyncGroupClaims

`func (o *SsoConfigGet200ResponseData) SetLdapSyncGroupClaims(v bool)`

SetLdapSyncGroupClaims sets LdapSyncGroupClaims field to given value.

### HasLdapSyncGroupClaims

`func (o *SsoConfigGet200ResponseData) HasLdapSyncGroupClaims() bool`

HasLdapSyncGroupClaims returns a boolean if a field has been set.

### GetLdapRoleMapping

`func (o *SsoConfigGet200ResponseData) GetLdapRoleMapping() string`

GetLdapRoleMapping returns the LdapRoleMapping field if non-nil, zero value otherwise.

### GetLdapRoleMappingOk

`func (o *SsoConfigGet200ResponseData) GetLdapRoleMappingOk() (*string, bool)`

GetLdapRoleMappingOk returns a tuple with the LdapRoleMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapRoleMapping

`func (o *SsoConfigGet200ResponseData) SetLdapRoleMapping(v string)`

SetLdapRoleMapping sets LdapRoleMapping field to given value.

### HasLdapRoleMapping

`func (o *SsoConfigGet200ResponseData) HasLdapRoleMapping() bool`

HasLdapRoleMapping returns a boolean if a field has been set.

### GetLdapServerUrl

`func (o *SsoConfigGet200ResponseData) GetLdapServerUrl() string`

GetLdapServerUrl returns the LdapServerUrl field if non-nil, zero value otherwise.

### GetLdapServerUrlOk

`func (o *SsoConfigGet200ResponseData) GetLdapServerUrlOk() (*string, bool)`

GetLdapServerUrlOk returns a tuple with the LdapServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerUrl

`func (o *SsoConfigGet200ResponseData) SetLdapServerUrl(v string)`

SetLdapServerUrl sets LdapServerUrl field to given value.

### HasLdapServerUrl

`func (o *SsoConfigGet200ResponseData) HasLdapServerUrl() bool`

HasLdapServerUrl returns a boolean if a field has been set.

### GetLdapBaseDomainComponents

`func (o *SsoConfigGet200ResponseData) GetLdapBaseDomainComponents() string`

GetLdapBaseDomainComponents returns the LdapBaseDomainComponents field if non-nil, zero value otherwise.

### GetLdapBaseDomainComponentsOk

`func (o *SsoConfigGet200ResponseData) GetLdapBaseDomainComponentsOk() (*string, bool)`

GetLdapBaseDomainComponentsOk returns a tuple with the LdapBaseDomainComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBaseDomainComponents

`func (o *SsoConfigGet200ResponseData) SetLdapBaseDomainComponents(v string)`

SetLdapBaseDomainComponents sets LdapBaseDomainComponents field to given value.

### HasLdapBaseDomainComponents

`func (o *SsoConfigGet200ResponseData) HasLdapBaseDomainComponents() bool`

HasLdapBaseDomainComponents returns a boolean if a field has been set.

### GetLdapServerName

`func (o *SsoConfigGet200ResponseData) GetLdapServerName() string`

GetLdapServerName returns the LdapServerName field if non-nil, zero value otherwise.

### GetLdapServerNameOk

`func (o *SsoConfigGet200ResponseData) GetLdapServerNameOk() (*string, bool)`

GetLdapServerNameOk returns a tuple with the LdapServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerName

`func (o *SsoConfigGet200ResponseData) SetLdapServerName(v string)`

SetLdapServerName sets LdapServerName field to given value.

### HasLdapServerName

`func (o *SsoConfigGet200ResponseData) HasLdapServerName() bool`

HasLdapServerName returns a boolean if a field has been set.

### GetLdapServerKey

`func (o *SsoConfigGet200ResponseData) GetLdapServerKey() string`

GetLdapServerKey returns the LdapServerKey field if non-nil, zero value otherwise.

### GetLdapServerKeyOk

`func (o *SsoConfigGet200ResponseData) GetLdapServerKeyOk() (*string, bool)`

GetLdapServerKeyOk returns a tuple with the LdapServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerKey

`func (o *SsoConfigGet200ResponseData) SetLdapServerKey(v string)`

SetLdapServerKey sets LdapServerKey field to given value.

### HasLdapServerKey

`func (o *SsoConfigGet200ResponseData) HasLdapServerKey() bool`

HasLdapServerKey returns a boolean if a field has been set.

### GetLdapServerCertificate

`func (o *SsoConfigGet200ResponseData) GetLdapServerCertificate() string`

GetLdapServerCertificate returns the LdapServerCertificate field if non-nil, zero value otherwise.

### GetLdapServerCertificateOk

`func (o *SsoConfigGet200ResponseData) GetLdapServerCertificateOk() (*string, bool)`

GetLdapServerCertificateOk returns a tuple with the LdapServerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerCertificate

`func (o *SsoConfigGet200ResponseData) SetLdapServerCertificate(v string)`

SetLdapServerCertificate sets LdapServerCertificate field to given value.

### HasLdapServerCertificate

`func (o *SsoConfigGet200ResponseData) HasLdapServerCertificate() bool`

HasLdapServerCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


