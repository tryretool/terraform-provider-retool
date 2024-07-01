# SsoConfigPostRequestDataOneOf3

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

### NewSsoConfigPostRequestDataOneOf3

`func NewSsoConfigPostRequestDataOneOf3(configType string, idpMetadataXml string, samlFirstNameAttribute string, samlLastNameAttribute string, samlSyncGroupClaims bool, jitEnabled bool, triggerLoginAutomatically bool, disableEmailPasswordLogin bool, ) *SsoConfigPostRequestDataOneOf3`

NewSsoConfigPostRequestDataOneOf3 instantiates a new SsoConfigPostRequestDataOneOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoConfigPostRequestDataOneOf3WithDefaults

`func NewSsoConfigPostRequestDataOneOf3WithDefaults() *SsoConfigPostRequestDataOneOf3`

NewSsoConfigPostRequestDataOneOf3WithDefaults instantiates a new SsoConfigPostRequestDataOneOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *SsoConfigPostRequestDataOneOf3) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *SsoConfigPostRequestDataOneOf3) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *SsoConfigPostRequestDataOneOf3) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetIdpMetadataXml

`func (o *SsoConfigPostRequestDataOneOf3) GetIdpMetadataXml() string`

GetIdpMetadataXml returns the IdpMetadataXml field if non-nil, zero value otherwise.

### GetIdpMetadataXmlOk

`func (o *SsoConfigPostRequestDataOneOf3) GetIdpMetadataXmlOk() (*string, bool)`

GetIdpMetadataXmlOk returns a tuple with the IdpMetadataXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpMetadataXml

`func (o *SsoConfigPostRequestDataOneOf3) SetIdpMetadataXml(v string)`

SetIdpMetadataXml sets IdpMetadataXml field to given value.


### GetSamlFirstNameAttribute

`func (o *SsoConfigPostRequestDataOneOf3) GetSamlFirstNameAttribute() string`

GetSamlFirstNameAttribute returns the SamlFirstNameAttribute field if non-nil, zero value otherwise.

### GetSamlFirstNameAttributeOk

`func (o *SsoConfigPostRequestDataOneOf3) GetSamlFirstNameAttributeOk() (*string, bool)`

GetSamlFirstNameAttributeOk returns a tuple with the SamlFirstNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlFirstNameAttribute

`func (o *SsoConfigPostRequestDataOneOf3) SetSamlFirstNameAttribute(v string)`

SetSamlFirstNameAttribute sets SamlFirstNameAttribute field to given value.


### GetSamlLastNameAttribute

`func (o *SsoConfigPostRequestDataOneOf3) GetSamlLastNameAttribute() string`

GetSamlLastNameAttribute returns the SamlLastNameAttribute field if non-nil, zero value otherwise.

### GetSamlLastNameAttributeOk

`func (o *SsoConfigPostRequestDataOneOf3) GetSamlLastNameAttributeOk() (*string, bool)`

GetSamlLastNameAttributeOk returns a tuple with the SamlLastNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlLastNameAttribute

`func (o *SsoConfigPostRequestDataOneOf3) SetSamlLastNameAttribute(v string)`

SetSamlLastNameAttribute sets SamlLastNameAttribute field to given value.


### GetSamlGroupsAttribute

`func (o *SsoConfigPostRequestDataOneOf3) GetSamlGroupsAttribute() string`

GetSamlGroupsAttribute returns the SamlGroupsAttribute field if non-nil, zero value otherwise.

### GetSamlGroupsAttributeOk

`func (o *SsoConfigPostRequestDataOneOf3) GetSamlGroupsAttributeOk() (*string, bool)`

GetSamlGroupsAttributeOk returns a tuple with the SamlGroupsAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlGroupsAttribute

`func (o *SsoConfigPostRequestDataOneOf3) SetSamlGroupsAttribute(v string)`

SetSamlGroupsAttribute sets SamlGroupsAttribute field to given value.

### HasSamlGroupsAttribute

`func (o *SsoConfigPostRequestDataOneOf3) HasSamlGroupsAttribute() bool`

HasSamlGroupsAttribute returns a boolean if a field has been set.

### GetSamlSyncGroupClaims

`func (o *SsoConfigPostRequestDataOneOf3) GetSamlSyncGroupClaims() bool`

GetSamlSyncGroupClaims returns the SamlSyncGroupClaims field if non-nil, zero value otherwise.

### GetSamlSyncGroupClaimsOk

`func (o *SsoConfigPostRequestDataOneOf3) GetSamlSyncGroupClaimsOk() (*bool, bool)`

GetSamlSyncGroupClaimsOk returns a tuple with the SamlSyncGroupClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSyncGroupClaims

`func (o *SsoConfigPostRequestDataOneOf3) SetSamlSyncGroupClaims(v bool)`

SetSamlSyncGroupClaims sets SamlSyncGroupClaims field to given value.


### GetLdapSyncGroupClaims

`func (o *SsoConfigPostRequestDataOneOf3) GetLdapSyncGroupClaims() bool`

GetLdapSyncGroupClaims returns the LdapSyncGroupClaims field if non-nil, zero value otherwise.

### GetLdapSyncGroupClaimsOk

`func (o *SsoConfigPostRequestDataOneOf3) GetLdapSyncGroupClaimsOk() (*bool, bool)`

GetLdapSyncGroupClaimsOk returns a tuple with the LdapSyncGroupClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapSyncGroupClaims

`func (o *SsoConfigPostRequestDataOneOf3) SetLdapSyncGroupClaims(v bool)`

SetLdapSyncGroupClaims sets LdapSyncGroupClaims field to given value.

### HasLdapSyncGroupClaims

`func (o *SsoConfigPostRequestDataOneOf3) HasLdapSyncGroupClaims() bool`

HasLdapSyncGroupClaims returns a boolean if a field has been set.

### GetLdapRoleMapping

`func (o *SsoConfigPostRequestDataOneOf3) GetLdapRoleMapping() string`

GetLdapRoleMapping returns the LdapRoleMapping field if non-nil, zero value otherwise.

### GetLdapRoleMappingOk

`func (o *SsoConfigPostRequestDataOneOf3) GetLdapRoleMappingOk() (*string, bool)`

GetLdapRoleMappingOk returns a tuple with the LdapRoleMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapRoleMapping

`func (o *SsoConfigPostRequestDataOneOf3) SetLdapRoleMapping(v string)`

SetLdapRoleMapping sets LdapRoleMapping field to given value.

### HasLdapRoleMapping

`func (o *SsoConfigPostRequestDataOneOf3) HasLdapRoleMapping() bool`

HasLdapRoleMapping returns a boolean if a field has been set.

### GetLdapServerUrl

`func (o *SsoConfigPostRequestDataOneOf3) GetLdapServerUrl() string`

GetLdapServerUrl returns the LdapServerUrl field if non-nil, zero value otherwise.

### GetLdapServerUrlOk

`func (o *SsoConfigPostRequestDataOneOf3) GetLdapServerUrlOk() (*string, bool)`

GetLdapServerUrlOk returns a tuple with the LdapServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerUrl

`func (o *SsoConfigPostRequestDataOneOf3) SetLdapServerUrl(v string)`

SetLdapServerUrl sets LdapServerUrl field to given value.

### HasLdapServerUrl

`func (o *SsoConfigPostRequestDataOneOf3) HasLdapServerUrl() bool`

HasLdapServerUrl returns a boolean if a field has been set.

### GetLdapBaseDomainComponents

`func (o *SsoConfigPostRequestDataOneOf3) GetLdapBaseDomainComponents() string`

GetLdapBaseDomainComponents returns the LdapBaseDomainComponents field if non-nil, zero value otherwise.

### GetLdapBaseDomainComponentsOk

`func (o *SsoConfigPostRequestDataOneOf3) GetLdapBaseDomainComponentsOk() (*string, bool)`

GetLdapBaseDomainComponentsOk returns a tuple with the LdapBaseDomainComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBaseDomainComponents

`func (o *SsoConfigPostRequestDataOneOf3) SetLdapBaseDomainComponents(v string)`

SetLdapBaseDomainComponents sets LdapBaseDomainComponents field to given value.

### HasLdapBaseDomainComponents

`func (o *SsoConfigPostRequestDataOneOf3) HasLdapBaseDomainComponents() bool`

HasLdapBaseDomainComponents returns a boolean if a field has been set.

### GetLdapServerName

`func (o *SsoConfigPostRequestDataOneOf3) GetLdapServerName() string`

GetLdapServerName returns the LdapServerName field if non-nil, zero value otherwise.

### GetLdapServerNameOk

`func (o *SsoConfigPostRequestDataOneOf3) GetLdapServerNameOk() (*string, bool)`

GetLdapServerNameOk returns a tuple with the LdapServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerName

`func (o *SsoConfigPostRequestDataOneOf3) SetLdapServerName(v string)`

SetLdapServerName sets LdapServerName field to given value.

### HasLdapServerName

`func (o *SsoConfigPostRequestDataOneOf3) HasLdapServerName() bool`

HasLdapServerName returns a boolean if a field has been set.

### GetLdapServerKey

`func (o *SsoConfigPostRequestDataOneOf3) GetLdapServerKey() string`

GetLdapServerKey returns the LdapServerKey field if non-nil, zero value otherwise.

### GetLdapServerKeyOk

`func (o *SsoConfigPostRequestDataOneOf3) GetLdapServerKeyOk() (*string, bool)`

GetLdapServerKeyOk returns a tuple with the LdapServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerKey

`func (o *SsoConfigPostRequestDataOneOf3) SetLdapServerKey(v string)`

SetLdapServerKey sets LdapServerKey field to given value.

### HasLdapServerKey

`func (o *SsoConfigPostRequestDataOneOf3) HasLdapServerKey() bool`

HasLdapServerKey returns a boolean if a field has been set.

### GetLdapServerCertificate

`func (o *SsoConfigPostRequestDataOneOf3) GetLdapServerCertificate() string`

GetLdapServerCertificate returns the LdapServerCertificate field if non-nil, zero value otherwise.

### GetLdapServerCertificateOk

`func (o *SsoConfigPostRequestDataOneOf3) GetLdapServerCertificateOk() (*string, bool)`

GetLdapServerCertificateOk returns a tuple with the LdapServerCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapServerCertificate

`func (o *SsoConfigPostRequestDataOneOf3) SetLdapServerCertificate(v string)`

SetLdapServerCertificate sets LdapServerCertificate field to given value.

### HasLdapServerCertificate

`func (o *SsoConfigPostRequestDataOneOf3) HasLdapServerCertificate() bool`

HasLdapServerCertificate returns a boolean if a field has been set.

### GetJitEnabled

`func (o *SsoConfigPostRequestDataOneOf3) GetJitEnabled() bool`

GetJitEnabled returns the JitEnabled field if non-nil, zero value otherwise.

### GetJitEnabledOk

`func (o *SsoConfigPostRequestDataOneOf3) GetJitEnabledOk() (*bool, bool)`

GetJitEnabledOk returns a tuple with the JitEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJitEnabled

`func (o *SsoConfigPostRequestDataOneOf3) SetJitEnabled(v bool)`

SetJitEnabled sets JitEnabled field to given value.


### GetRestrictedDomain

`func (o *SsoConfigPostRequestDataOneOf3) GetRestrictedDomain() string`

GetRestrictedDomain returns the RestrictedDomain field if non-nil, zero value otherwise.

### GetRestrictedDomainOk

`func (o *SsoConfigPostRequestDataOneOf3) GetRestrictedDomainOk() (*string, bool)`

GetRestrictedDomainOk returns a tuple with the RestrictedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedDomain

`func (o *SsoConfigPostRequestDataOneOf3) SetRestrictedDomain(v string)`

SetRestrictedDomain sets RestrictedDomain field to given value.

### HasRestrictedDomain

`func (o *SsoConfigPostRequestDataOneOf3) HasRestrictedDomain() bool`

HasRestrictedDomain returns a boolean if a field has been set.

### GetTriggerLoginAutomatically

`func (o *SsoConfigPostRequestDataOneOf3) GetTriggerLoginAutomatically() bool`

GetTriggerLoginAutomatically returns the TriggerLoginAutomatically field if non-nil, zero value otherwise.

### GetTriggerLoginAutomaticallyOk

`func (o *SsoConfigPostRequestDataOneOf3) GetTriggerLoginAutomaticallyOk() (*bool, bool)`

GetTriggerLoginAutomaticallyOk returns a tuple with the TriggerLoginAutomatically field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerLoginAutomatically

`func (o *SsoConfigPostRequestDataOneOf3) SetTriggerLoginAutomatically(v bool)`

SetTriggerLoginAutomatically sets TriggerLoginAutomatically field to given value.


### GetDisableEmailPasswordLogin

`func (o *SsoConfigPostRequestDataOneOf3) GetDisableEmailPasswordLogin() bool`

GetDisableEmailPasswordLogin returns the DisableEmailPasswordLogin field if non-nil, zero value otherwise.

### GetDisableEmailPasswordLoginOk

`func (o *SsoConfigPostRequestDataOneOf3) GetDisableEmailPasswordLoginOk() (*bool, bool)`

GetDisableEmailPasswordLoginOk returns a tuple with the DisableEmailPasswordLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmailPasswordLogin

`func (o *SsoConfigPostRequestDataOneOf3) SetDisableEmailPasswordLogin(v bool)`

SetDisableEmailPasswordLogin sets DisableEmailPasswordLogin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


