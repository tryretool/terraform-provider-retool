# SAML

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** |  | 
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
**JitEnabled** | **bool** |  | 
**RestrictedDomain** | Pointer to **string** |  | [optional] 
**TriggerLoginAutomatically** | **bool** |  | 
**DisableEmailPasswordLogin** | **bool** |  | 

## Methods

### NewSAML

`func NewSAML(configType string, idpMetadataXml string, samlFirstNameAttribute string, samlLastNameAttribute string, samlSyncGroupClaims bool, jitEnabled bool, triggerLoginAutomatically bool, disableEmailPasswordLogin bool, ) *SAML`

NewSAML instantiates a new SAML object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSAMLWithDefaults

`func NewSAMLWithDefaults() *SAML`

NewSAMLWithDefaults instantiates a new SAML object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *SAML) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *SAML) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *SAML) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetIdpMetadataXml

`func (o *SAML) GetIdpMetadataXml() string`

GetIdpMetadataXml returns the IdpMetadataXml field if non-nil, zero value otherwise.

### GetIdpMetadataXmlOk

`func (o *SAML) GetIdpMetadataXmlOk() (*string, bool)`

GetIdpMetadataXmlOk returns a tuple with the IdpMetadataXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataXml

`func (o *SAML) SetIdpMetadataXml(v string)`

SetIdpMetadataXml sets IdpMetadataXml field to given value.


### GetSamlFirstNameAttribute

`func (o *SAML) GetSamlFirstNameAttribute() string`

GetSamlFirstNameAttribute returns the SamlFirstNameAttribute field if non-nil, zero value otherwise.

### GetSamlFirstNameAttributeOk

`func (o *SAML) GetSamlFirstNameAttributeOk() (*string, bool)`

GetSamlFirstNameAttributeOk returns a tuple with the SamlFirstNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlFirstNameAttribute

`func (o *SAML) SetSamlFirstNameAttribute(v string)`

SetSamlFirstNameAttribute sets SamlFirstNameAttribute field to given value.


### GetSamlLastNameAttribute

`func (o *SAML) GetSamlLastNameAttribute() string`

GetSamlLastNameAttribute returns the SamlLastNameAttribute field if non-nil, zero value otherwise.

### GetSamlLastNameAttributeOk

`func (o *SAML) GetSamlLastNameAttributeOk() (*string, bool)`

GetSamlLastNameAttributeOk returns a tuple with the SamlLastNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlLastNameAttribute

`func (o *SAML) SetSamlLastNameAttribute(v string)`

SetSamlLastNameAttribute sets SamlLastNameAttribute field to given value.


### GetSamlGroupsAttribute

`func (o *SAML) GetSamlGroupsAttribute() string`

GetSamlGroupsAttribute returns the SamlGroupsAttribute field if non-nil, zero value otherwise.

### GetSamlGroupsAttributeOk

`func (o *SAML) GetSamlGroupsAttributeOk() (*string, bool)`

GetSamlGroupsAttributeOk returns a tuple with the SamlGroupsAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlGroupsAttribute

`func (o *SAML) SetSamlGroupsAttribute(v string)`

SetSamlGroupsAttribute sets SamlGroupsAttribute field to given value.

### HasSamlGroupsAttribute

`func (o *SAML) HasSamlGroupsAttribute() bool`

HasSamlGroupsAttribute returns a boolean if a field has been set.

### GetSamlSyncGroupClaims

`func (o *SAML) GetSamlSyncGroupClaims() bool`

GetSamlSyncGroupClaims returns the SamlSyncGroupClaims field if non-nil, zero value otherwise.

### GetSamlSyncGroupClaimsOk

`func (o *SAML) GetSamlSyncGroupClaimsOk() (*bool, bool)`

GetSamlSyncGroupClaimsOk returns a tuple with the SamlSyncGroupClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSyncGroupClaims

`func (o *SAML) SetSamlSyncGroupClaims(v bool)`

SetSamlSyncGroupClaims sets SamlSyncGroupClaims field to given value.


### GetLdapSyncGroupClaims

`func (o *SAML) GetLdapSyncGroupClaims() bool`

GetLdapSyncGroupClaims returns the LdapSyncGroupClaims field if non-nil, zero value otherwise.

### GetLdapSyncGroupClaimsOk

`func (o *SAML) GetLdapSyncGroupClaimsOk() (*bool, bool)`

GetLdapSyncGroupClaimsOk returns a tuple with the LdapSyncGroupClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapSyncGroupClaims

`func (o *SAML) SetLdapSyncGroupClaims(v bool)`

SetLdapSyncGroupClaims sets LdapSyncGroupClaims field to given value.

### HasLdapSyncGroupClaims

`func (o *SAML) HasLdapSyncGroupClaims() bool`

HasLdapSyncGroupClaims returns a boolean if a field has been set.

### GetLdapRoleMapping

`func (o *SAML) GetLdapRoleMapping() string`

GetLdapRoleMapping returns the LdapRoleMapping field if non-nil, zero value otherwise.

### GetLdapRoleMappingOk

`func (o *SAML) GetLdapRoleMappingOk() (*string, bool)`

GetLdapRoleMappingOk returns a tuple with the LdapRoleMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapRoleMapping

`func (o *SAML) SetLdapRoleMapping(v string)`

SetLdapRoleMapping sets LdapRoleMapping field to given value.

### HasLdapRoleMapping

`func (o *SAML) HasLdapRoleMapping() bool`

HasLdapRoleMapping returns a boolean if a field has been set.

### GetLdapServerUrl

`func (o *SAML) GetLdapServerUrl() string`

GetLdapServerUrl returns the LdapServerUrl field if non-nil, zero value otherwise.

### GetLdapServerUrlOk

`func (o *SAML) GetLdapServerUrlOk() (*string, bool)`

GetLdapServerUrlOk returns a tuple with the LdapServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerUrl

`func (o *SAML) SetLdapServerUrl(v string)`

SetLdapServerUrl sets LdapServerUrl field to given value.

### HasLdapServerUrl

`func (o *SAML) HasLdapServerUrl() bool`

HasLdapServerUrl returns a boolean if a field has been set.

### GetLdapBaseDomainComponents

`func (o *SAML) GetLdapBaseDomainComponents() string`

GetLdapBaseDomainComponents returns the LdapBaseDomainComponents field if non-nil, zero value otherwise.

### GetLdapBaseDomainComponentsOk

`func (o *SAML) GetLdapBaseDomainComponentsOk() (*string, bool)`

GetLdapBaseDomainComponentsOk returns a tuple with the LdapBaseDomainComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBaseDomainComponents

`func (o *SAML) SetLdapBaseDomainComponents(v string)`

SetLdapBaseDomainComponents sets LdapBaseDomainComponents field to given value.

### HasLdapBaseDomainComponents

`func (o *SAML) HasLdapBaseDomainComponents() bool`

HasLdapBaseDomainComponents returns a boolean if a field has been set.

### GetLdapServerName

`func (o *SAML) GetLdapServerName() string`

GetLdapServerName returns the LdapServerName field if non-nil, zero value otherwise.

### GetLdapServerNameOk

`func (o *SAML) GetLdapServerNameOk() (*string, bool)`

GetLdapServerNameOk returns a tuple with the LdapServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerName

`func (o *SAML) SetLdapServerName(v string)`

SetLdapServerName sets LdapServerName field to given value.

### HasLdapServerName

`func (o *SAML) HasLdapServerName() bool`

HasLdapServerName returns a boolean if a field has been set.

### GetLdapServerKey

`func (o *SAML) GetLdapServerKey() string`

GetLdapServerKey returns the LdapServerKey field if non-nil, zero value otherwise.

### GetLdapServerKeyOk

`func (o *SAML) GetLdapServerKeyOk() (*string, bool)`

GetLdapServerKeyOk returns a tuple with the LdapServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerKey

`func (o *SAML) SetLdapServerKey(v string)`

SetLdapServerKey sets LdapServerKey field to given value.

### HasLdapServerKey

`func (o *SAML) HasLdapServerKey() bool`

HasLdapServerKey returns a boolean if a field has been set.

### GetLdapServerCertificate

`func (o *SAML) GetLdapServerCertificate() string`

GetLdapServerCertificate returns the LdapServerCertificate field if non-nil, zero value otherwise.

### GetLdapServerCertificateOk

`func (o *SAML) GetLdapServerCertificateOk() (*string, bool)`

GetLdapServerCertificateOk returns a tuple with the LdapServerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerCertificate

`func (o *SAML) SetLdapServerCertificate(v string)`

SetLdapServerCertificate sets LdapServerCertificate field to given value.

### HasLdapServerCertificate

`func (o *SAML) HasLdapServerCertificate() bool`

HasLdapServerCertificate returns a boolean if a field has been set.

### GetJitEnabled

`func (o *SAML) GetJitEnabled() bool`

GetJitEnabled returns the JitEnabled field if non-nil, zero value otherwise.

### GetJitEnabledOk

`func (o *SAML) GetJitEnabledOk() (*bool, bool)`

GetJitEnabledOk returns a tuple with the JitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitEnabled

`func (o *SAML) SetJitEnabled(v bool)`

SetJitEnabled sets JitEnabled field to given value.


### GetRestrictedDomain

`func (o *SAML) GetRestrictedDomain() string`

GetRestrictedDomain returns the RestrictedDomain field if non-nil, zero value otherwise.

### GetRestrictedDomainOk

`func (o *SAML) GetRestrictedDomainOk() (*string, bool)`

GetRestrictedDomainOk returns a tuple with the RestrictedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDomain

`func (o *SAML) SetRestrictedDomain(v string)`

SetRestrictedDomain sets RestrictedDomain field to given value.

### HasRestrictedDomain

`func (o *SAML) HasRestrictedDomain() bool`

HasRestrictedDomain returns a boolean if a field has been set.

### GetTriggerLoginAutomatically

`func (o *SAML) GetTriggerLoginAutomatically() bool`

GetTriggerLoginAutomatically returns the TriggerLoginAutomatically field if non-nil, zero value otherwise.

### GetTriggerLoginAutomaticallyOk

`func (o *SAML) GetTriggerLoginAutomaticallyOk() (*bool, bool)`

GetTriggerLoginAutomaticallyOk returns a tuple with the TriggerLoginAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerLoginAutomatically

`func (o *SAML) SetTriggerLoginAutomatically(v bool)`

SetTriggerLoginAutomatically sets TriggerLoginAutomatically field to given value.


### GetDisableEmailPasswordLogin

`func (o *SAML) GetDisableEmailPasswordLogin() bool`

GetDisableEmailPasswordLogin returns the DisableEmailPasswordLogin field if non-nil, zero value otherwise.

### GetDisableEmailPasswordLoginOk

`func (o *SAML) GetDisableEmailPasswordLoginOk() (*bool, bool)`

GetDisableEmailPasswordLoginOk returns a tuple with the DisableEmailPasswordLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmailPasswordLogin

`func (o *SAML) SetDisableEmailPasswordLogin(v bool)`

SetDisableEmailPasswordLogin sets DisableEmailPasswordLogin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


