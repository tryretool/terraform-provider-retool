# SsoConfigPostRequestDataOneOf4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** |  | 
**GoogleClientId** | **string** |  | 
**GoogleClientSecret** | **string** |  | 
**DisableEmailPasswordLogin** | **bool** |  | 
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

## Methods

### NewSsoConfigPostRequestDataOneOf4

`func NewSsoConfigPostRequestDataOneOf4(configType string, googleClientId string, googleClientSecret string, disableEmailPasswordLogin bool, idpMetadataXml string, samlFirstNameAttribute string, samlLastNameAttribute string, samlSyncGroupClaims bool, jitEnabled bool, triggerLoginAutomatically bool, ) *SsoConfigPostRequestDataOneOf4`

NewSsoConfigPostRequestDataOneOf4 instantiates a new SsoConfigPostRequestDataOneOf4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoConfigPostRequestDataOneOf4WithDefaults

`func NewSsoConfigPostRequestDataOneOf4WithDefaults() *SsoConfigPostRequestDataOneOf4`

NewSsoConfigPostRequestDataOneOf4WithDefaults instantiates a new SsoConfigPostRequestDataOneOf4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *SsoConfigPostRequestDataOneOf4) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *SsoConfigPostRequestDataOneOf4) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *SsoConfigPostRequestDataOneOf4) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetGoogleClientId

`func (o *SsoConfigPostRequestDataOneOf4) GetGoogleClientId() string`

GetGoogleClientId returns the GoogleClientId field if non-nil, zero value otherwise.

### GetGoogleClientIdOk

`func (o *SsoConfigPostRequestDataOneOf4) GetGoogleClientIdOk() (*string, bool)`

GetGoogleClientIdOk returns a tuple with the GoogleClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleClientId

`func (o *SsoConfigPostRequestDataOneOf4) SetGoogleClientId(v string)`

SetGoogleClientId sets GoogleClientId field to given value.


### GetGoogleClientSecret

`func (o *SsoConfigPostRequestDataOneOf4) GetGoogleClientSecret() string`

GetGoogleClientSecret returns the GoogleClientSecret field if non-nil, zero value otherwise.

### GetGoogleClientSecretOk

`func (o *SsoConfigPostRequestDataOneOf4) GetGoogleClientSecretOk() (*string, bool)`

GetGoogleClientSecretOk returns a tuple with the GoogleClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleClientSecret

`func (o *SsoConfigPostRequestDataOneOf4) SetGoogleClientSecret(v string)`

SetGoogleClientSecret sets GoogleClientSecret field to given value.


### GetDisableEmailPasswordLogin

`func (o *SsoConfigPostRequestDataOneOf4) GetDisableEmailPasswordLogin() bool`

GetDisableEmailPasswordLogin returns the DisableEmailPasswordLogin field if non-nil, zero value otherwise.

### GetDisableEmailPasswordLoginOk

`func (o *SsoConfigPostRequestDataOneOf4) GetDisableEmailPasswordLoginOk() (*bool, bool)`

GetDisableEmailPasswordLoginOk returns a tuple with the DisableEmailPasswordLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmailPasswordLogin

`func (o *SsoConfigPostRequestDataOneOf4) SetDisableEmailPasswordLogin(v bool)`

SetDisableEmailPasswordLogin sets DisableEmailPasswordLogin field to given value.


### GetIdpMetadataXml

`func (o *SsoConfigPostRequestDataOneOf4) GetIdpMetadataXml() string`

GetIdpMetadataXml returns the IdpMetadataXml field if non-nil, zero value otherwise.

### GetIdpMetadataXmlOk

`func (o *SsoConfigPostRequestDataOneOf4) GetIdpMetadataXmlOk() (*string, bool)`

GetIdpMetadataXmlOk returns a tuple with the IdpMetadataXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataXml

`func (o *SsoConfigPostRequestDataOneOf4) SetIdpMetadataXml(v string)`

SetIdpMetadataXml sets IdpMetadataXml field to given value.


### GetSamlFirstNameAttribute

`func (o *SsoConfigPostRequestDataOneOf4) GetSamlFirstNameAttribute() string`

GetSamlFirstNameAttribute returns the SamlFirstNameAttribute field if non-nil, zero value otherwise.

### GetSamlFirstNameAttributeOk

`func (o *SsoConfigPostRequestDataOneOf4) GetSamlFirstNameAttributeOk() (*string, bool)`

GetSamlFirstNameAttributeOk returns a tuple with the SamlFirstNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlFirstNameAttribute

`func (o *SsoConfigPostRequestDataOneOf4) SetSamlFirstNameAttribute(v string)`

SetSamlFirstNameAttribute sets SamlFirstNameAttribute field to given value.


### GetSamlLastNameAttribute

`func (o *SsoConfigPostRequestDataOneOf4) GetSamlLastNameAttribute() string`

GetSamlLastNameAttribute returns the SamlLastNameAttribute field if non-nil, zero value otherwise.

### GetSamlLastNameAttributeOk

`func (o *SsoConfigPostRequestDataOneOf4) GetSamlLastNameAttributeOk() (*string, bool)`

GetSamlLastNameAttributeOk returns a tuple with the SamlLastNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlLastNameAttribute

`func (o *SsoConfigPostRequestDataOneOf4) SetSamlLastNameAttribute(v string)`

SetSamlLastNameAttribute sets SamlLastNameAttribute field to given value.


### GetSamlGroupsAttribute

`func (o *SsoConfigPostRequestDataOneOf4) GetSamlGroupsAttribute() string`

GetSamlGroupsAttribute returns the SamlGroupsAttribute field if non-nil, zero value otherwise.

### GetSamlGroupsAttributeOk

`func (o *SsoConfigPostRequestDataOneOf4) GetSamlGroupsAttributeOk() (*string, bool)`

GetSamlGroupsAttributeOk returns a tuple with the SamlGroupsAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlGroupsAttribute

`func (o *SsoConfigPostRequestDataOneOf4) SetSamlGroupsAttribute(v string)`

SetSamlGroupsAttribute sets SamlGroupsAttribute field to given value.

### HasSamlGroupsAttribute

`func (o *SsoConfigPostRequestDataOneOf4) HasSamlGroupsAttribute() bool`

HasSamlGroupsAttribute returns a boolean if a field has been set.

### GetSamlSyncGroupClaims

`func (o *SsoConfigPostRequestDataOneOf4) GetSamlSyncGroupClaims() bool`

GetSamlSyncGroupClaims returns the SamlSyncGroupClaims field if non-nil, zero value otherwise.

### GetSamlSyncGroupClaimsOk

`func (o *SsoConfigPostRequestDataOneOf4) GetSamlSyncGroupClaimsOk() (*bool, bool)`

GetSamlSyncGroupClaimsOk returns a tuple with the SamlSyncGroupClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSyncGroupClaims

`func (o *SsoConfigPostRequestDataOneOf4) SetSamlSyncGroupClaims(v bool)`

SetSamlSyncGroupClaims sets SamlSyncGroupClaims field to given value.


### GetLdapSyncGroupClaims

`func (o *SsoConfigPostRequestDataOneOf4) GetLdapSyncGroupClaims() bool`

GetLdapSyncGroupClaims returns the LdapSyncGroupClaims field if non-nil, zero value otherwise.

### GetLdapSyncGroupClaimsOk

`func (o *SsoConfigPostRequestDataOneOf4) GetLdapSyncGroupClaimsOk() (*bool, bool)`

GetLdapSyncGroupClaimsOk returns a tuple with the LdapSyncGroupClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapSyncGroupClaims

`func (o *SsoConfigPostRequestDataOneOf4) SetLdapSyncGroupClaims(v bool)`

SetLdapSyncGroupClaims sets LdapSyncGroupClaims field to given value.

### HasLdapSyncGroupClaims

`func (o *SsoConfigPostRequestDataOneOf4) HasLdapSyncGroupClaims() bool`

HasLdapSyncGroupClaims returns a boolean if a field has been set.

### GetLdapRoleMapping

`func (o *SsoConfigPostRequestDataOneOf4) GetLdapRoleMapping() string`

GetLdapRoleMapping returns the LdapRoleMapping field if non-nil, zero value otherwise.

### GetLdapRoleMappingOk

`func (o *SsoConfigPostRequestDataOneOf4) GetLdapRoleMappingOk() (*string, bool)`

GetLdapRoleMappingOk returns a tuple with the LdapRoleMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapRoleMapping

`func (o *SsoConfigPostRequestDataOneOf4) SetLdapRoleMapping(v string)`

SetLdapRoleMapping sets LdapRoleMapping field to given value.

### HasLdapRoleMapping

`func (o *SsoConfigPostRequestDataOneOf4) HasLdapRoleMapping() bool`

HasLdapRoleMapping returns a boolean if a field has been set.

### GetLdapServerUrl

`func (o *SsoConfigPostRequestDataOneOf4) GetLdapServerUrl() string`

GetLdapServerUrl returns the LdapServerUrl field if non-nil, zero value otherwise.

### GetLdapServerUrlOk

`func (o *SsoConfigPostRequestDataOneOf4) GetLdapServerUrlOk() (*string, bool)`

GetLdapServerUrlOk returns a tuple with the LdapServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerUrl

`func (o *SsoConfigPostRequestDataOneOf4) SetLdapServerUrl(v string)`

SetLdapServerUrl sets LdapServerUrl field to given value.

### HasLdapServerUrl

`func (o *SsoConfigPostRequestDataOneOf4) HasLdapServerUrl() bool`

HasLdapServerUrl returns a boolean if a field has been set.

### GetLdapBaseDomainComponents

`func (o *SsoConfigPostRequestDataOneOf4) GetLdapBaseDomainComponents() string`

GetLdapBaseDomainComponents returns the LdapBaseDomainComponents field if non-nil, zero value otherwise.

### GetLdapBaseDomainComponentsOk

`func (o *SsoConfigPostRequestDataOneOf4) GetLdapBaseDomainComponentsOk() (*string, bool)`

GetLdapBaseDomainComponentsOk returns a tuple with the LdapBaseDomainComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBaseDomainComponents

`func (o *SsoConfigPostRequestDataOneOf4) SetLdapBaseDomainComponents(v string)`

SetLdapBaseDomainComponents sets LdapBaseDomainComponents field to given value.

### HasLdapBaseDomainComponents

`func (o *SsoConfigPostRequestDataOneOf4) HasLdapBaseDomainComponents() bool`

HasLdapBaseDomainComponents returns a boolean if a field has been set.

### GetLdapServerName

`func (o *SsoConfigPostRequestDataOneOf4) GetLdapServerName() string`

GetLdapServerName returns the LdapServerName field if non-nil, zero value otherwise.

### GetLdapServerNameOk

`func (o *SsoConfigPostRequestDataOneOf4) GetLdapServerNameOk() (*string, bool)`

GetLdapServerNameOk returns a tuple with the LdapServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerName

`func (o *SsoConfigPostRequestDataOneOf4) SetLdapServerName(v string)`

SetLdapServerName sets LdapServerName field to given value.

### HasLdapServerName

`func (o *SsoConfigPostRequestDataOneOf4) HasLdapServerName() bool`

HasLdapServerName returns a boolean if a field has been set.

### GetLdapServerKey

`func (o *SsoConfigPostRequestDataOneOf4) GetLdapServerKey() string`

GetLdapServerKey returns the LdapServerKey field if non-nil, zero value otherwise.

### GetLdapServerKeyOk

`func (o *SsoConfigPostRequestDataOneOf4) GetLdapServerKeyOk() (*string, bool)`

GetLdapServerKeyOk returns a tuple with the LdapServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerKey

`func (o *SsoConfigPostRequestDataOneOf4) SetLdapServerKey(v string)`

SetLdapServerKey sets LdapServerKey field to given value.

### HasLdapServerKey

`func (o *SsoConfigPostRequestDataOneOf4) HasLdapServerKey() bool`

HasLdapServerKey returns a boolean if a field has been set.

### GetLdapServerCertificate

`func (o *SsoConfigPostRequestDataOneOf4) GetLdapServerCertificate() string`

GetLdapServerCertificate returns the LdapServerCertificate field if non-nil, zero value otherwise.

### GetLdapServerCertificateOk

`func (o *SsoConfigPostRequestDataOneOf4) GetLdapServerCertificateOk() (*string, bool)`

GetLdapServerCertificateOk returns a tuple with the LdapServerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerCertificate

`func (o *SsoConfigPostRequestDataOneOf4) SetLdapServerCertificate(v string)`

SetLdapServerCertificate sets LdapServerCertificate field to given value.

### HasLdapServerCertificate

`func (o *SsoConfigPostRequestDataOneOf4) HasLdapServerCertificate() bool`

HasLdapServerCertificate returns a boolean if a field has been set.

### GetJitEnabled

`func (o *SsoConfigPostRequestDataOneOf4) GetJitEnabled() bool`

GetJitEnabled returns the JitEnabled field if non-nil, zero value otherwise.

### GetJitEnabledOk

`func (o *SsoConfigPostRequestDataOneOf4) GetJitEnabledOk() (*bool, bool)`

GetJitEnabledOk returns a tuple with the JitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitEnabled

`func (o *SsoConfigPostRequestDataOneOf4) SetJitEnabled(v bool)`

SetJitEnabled sets JitEnabled field to given value.


### GetRestrictedDomain

`func (o *SsoConfigPostRequestDataOneOf4) GetRestrictedDomain() string`

GetRestrictedDomain returns the RestrictedDomain field if non-nil, zero value otherwise.

### GetRestrictedDomainOk

`func (o *SsoConfigPostRequestDataOneOf4) GetRestrictedDomainOk() (*string, bool)`

GetRestrictedDomainOk returns a tuple with the RestrictedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDomain

`func (o *SsoConfigPostRequestDataOneOf4) SetRestrictedDomain(v string)`

SetRestrictedDomain sets RestrictedDomain field to given value.

### HasRestrictedDomain

`func (o *SsoConfigPostRequestDataOneOf4) HasRestrictedDomain() bool`

HasRestrictedDomain returns a boolean if a field has been set.

### GetTriggerLoginAutomatically

`func (o *SsoConfigPostRequestDataOneOf4) GetTriggerLoginAutomatically() bool`

GetTriggerLoginAutomatically returns the TriggerLoginAutomatically field if non-nil, zero value otherwise.

### GetTriggerLoginAutomaticallyOk

`func (o *SsoConfigPostRequestDataOneOf4) GetTriggerLoginAutomaticallyOk() (*bool, bool)`

GetTriggerLoginAutomaticallyOk returns a tuple with the TriggerLoginAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerLoginAutomatically

`func (o *SsoConfigPostRequestDataOneOf4) SetTriggerLoginAutomatically(v bool)`

SetTriggerLoginAutomatically sets TriggerLoginAutomatically field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


