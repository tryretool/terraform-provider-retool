# AzureBlobOptionsAuthenticationOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | **string** |  | 
**AzureidScope** | Pointer to **string** |  | [optional] 
**AzureidTenantId** | Pointer to **string** |  | [optional] 
**AzureidClientId** | Pointer to **string** |  | [optional] 
**AzureidClientSecret** | Pointer to **string** |  | [optional] 
**AuthWithDefaultCredentialProviderChain** | **bool** |  | 
**AzureBlobAccountName** | **string** |  | 
**AzureBlobDefaultContainerName** | Pointer to **string** |  | [optional] 
**AzureStorageUrl** | **string** |  | 
**AzureStorageAccountName** | **string** |  | 
**AzureStorageAccountKey** | **string** |  | 

## Methods

### NewAzureBlobOptionsAuthenticationOptions

`func NewAzureBlobOptionsAuthenticationOptions(authenticationType string, authWithDefaultCredentialProviderChain bool, azureBlobAccountName string, azureStorageUrl string, azureStorageAccountName string, azureStorageAccountKey string, ) *AzureBlobOptionsAuthenticationOptions`

NewAzureBlobOptionsAuthenticationOptions instantiates a new AzureBlobOptionsAuthenticationOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureBlobOptionsAuthenticationOptionsWithDefaults

`func NewAzureBlobOptionsAuthenticationOptionsWithDefaults() *AzureBlobOptionsAuthenticationOptions`

NewAzureBlobOptionsAuthenticationOptionsWithDefaults instantiates a new AzureBlobOptionsAuthenticationOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *AzureBlobOptionsAuthenticationOptions) GetAuthenticationType() string`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *AzureBlobOptionsAuthenticationOptions) GetAuthenticationTypeOk() (*string, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *AzureBlobOptionsAuthenticationOptions) SetAuthenticationType(v string)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetAzureidScope

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureidScope() string`

GetAzureidScope returns the AzureidScope field if non-nil, zero value otherwise.

### GetAzureidScopeOk

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureidScopeOk() (*string, bool)`

GetAzureidScopeOk returns a tuple with the AzureidScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureidScope

`func (o *AzureBlobOptionsAuthenticationOptions) SetAzureidScope(v string)`

SetAzureidScope sets AzureidScope field to given value.

### HasAzureidScope

`func (o *AzureBlobOptionsAuthenticationOptions) HasAzureidScope() bool`

HasAzureidScope returns a boolean if a field has been set.

### GetAzureidTenantId

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureidTenantId() string`

GetAzureidTenantId returns the AzureidTenantId field if non-nil, zero value otherwise.

### GetAzureidTenantIdOk

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureidTenantIdOk() (*string, bool)`

GetAzureidTenantIdOk returns a tuple with the AzureidTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureidTenantId

`func (o *AzureBlobOptionsAuthenticationOptions) SetAzureidTenantId(v string)`

SetAzureidTenantId sets AzureidTenantId field to given value.

### HasAzureidTenantId

`func (o *AzureBlobOptionsAuthenticationOptions) HasAzureidTenantId() bool`

HasAzureidTenantId returns a boolean if a field has been set.

### GetAzureidClientId

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureidClientId() string`

GetAzureidClientId returns the AzureidClientId field if non-nil, zero value otherwise.

### GetAzureidClientIdOk

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureidClientIdOk() (*string, bool)`

GetAzureidClientIdOk returns a tuple with the AzureidClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureidClientId

`func (o *AzureBlobOptionsAuthenticationOptions) SetAzureidClientId(v string)`

SetAzureidClientId sets AzureidClientId field to given value.

### HasAzureidClientId

`func (o *AzureBlobOptionsAuthenticationOptions) HasAzureidClientId() bool`

HasAzureidClientId returns a boolean if a field has been set.

### GetAzureidClientSecret

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureidClientSecret() string`

GetAzureidClientSecret returns the AzureidClientSecret field if non-nil, zero value otherwise.

### GetAzureidClientSecretOk

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureidClientSecretOk() (*string, bool)`

GetAzureidClientSecretOk returns a tuple with the AzureidClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureidClientSecret

`func (o *AzureBlobOptionsAuthenticationOptions) SetAzureidClientSecret(v string)`

SetAzureidClientSecret sets AzureidClientSecret field to given value.

### HasAzureidClientSecret

`func (o *AzureBlobOptionsAuthenticationOptions) HasAzureidClientSecret() bool`

HasAzureidClientSecret returns a boolean if a field has been set.

### GetAuthWithDefaultCredentialProviderChain

`func (o *AzureBlobOptionsAuthenticationOptions) GetAuthWithDefaultCredentialProviderChain() bool`

GetAuthWithDefaultCredentialProviderChain returns the AuthWithDefaultCredentialProviderChain field if non-nil, zero value otherwise.

### GetAuthWithDefaultCredentialProviderChainOk

`func (o *AzureBlobOptionsAuthenticationOptions) GetAuthWithDefaultCredentialProviderChainOk() (*bool, bool)`

GetAuthWithDefaultCredentialProviderChainOk returns a tuple with the AuthWithDefaultCredentialProviderChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthWithDefaultCredentialProviderChain

`func (o *AzureBlobOptionsAuthenticationOptions) SetAuthWithDefaultCredentialProviderChain(v bool)`

SetAuthWithDefaultCredentialProviderChain sets AuthWithDefaultCredentialProviderChain field to given value.


### GetAzureBlobAccountName

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureBlobAccountName() string`

GetAzureBlobAccountName returns the AzureBlobAccountName field if non-nil, zero value otherwise.

### GetAzureBlobAccountNameOk

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureBlobAccountNameOk() (*string, bool)`

GetAzureBlobAccountNameOk returns a tuple with the AzureBlobAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureBlobAccountName

`func (o *AzureBlobOptionsAuthenticationOptions) SetAzureBlobAccountName(v string)`

SetAzureBlobAccountName sets AzureBlobAccountName field to given value.


### GetAzureBlobDefaultContainerName

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureBlobDefaultContainerName() string`

GetAzureBlobDefaultContainerName returns the AzureBlobDefaultContainerName field if non-nil, zero value otherwise.

### GetAzureBlobDefaultContainerNameOk

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureBlobDefaultContainerNameOk() (*string, bool)`

GetAzureBlobDefaultContainerNameOk returns a tuple with the AzureBlobDefaultContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureBlobDefaultContainerName

`func (o *AzureBlobOptionsAuthenticationOptions) SetAzureBlobDefaultContainerName(v string)`

SetAzureBlobDefaultContainerName sets AzureBlobDefaultContainerName field to given value.

### HasAzureBlobDefaultContainerName

`func (o *AzureBlobOptionsAuthenticationOptions) HasAzureBlobDefaultContainerName() bool`

HasAzureBlobDefaultContainerName returns a boolean if a field has been set.

### GetAzureStorageUrl

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureStorageUrl() string`

GetAzureStorageUrl returns the AzureStorageUrl field if non-nil, zero value otherwise.

### GetAzureStorageUrlOk

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureStorageUrlOk() (*string, bool)`

GetAzureStorageUrlOk returns a tuple with the AzureStorageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureStorageUrl

`func (o *AzureBlobOptionsAuthenticationOptions) SetAzureStorageUrl(v string)`

SetAzureStorageUrl sets AzureStorageUrl field to given value.


### GetAzureStorageAccountName

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureStorageAccountName() string`

GetAzureStorageAccountName returns the AzureStorageAccountName field if non-nil, zero value otherwise.

### GetAzureStorageAccountNameOk

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureStorageAccountNameOk() (*string, bool)`

GetAzureStorageAccountNameOk returns a tuple with the AzureStorageAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureStorageAccountName

`func (o *AzureBlobOptionsAuthenticationOptions) SetAzureStorageAccountName(v string)`

SetAzureStorageAccountName sets AzureStorageAccountName field to given value.


### GetAzureStorageAccountKey

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureStorageAccountKey() string`

GetAzureStorageAccountKey returns the AzureStorageAccountKey field if non-nil, zero value otherwise.

### GetAzureStorageAccountKeyOk

`func (o *AzureBlobOptionsAuthenticationOptions) GetAzureStorageAccountKeyOk() (*string, bool)`

GetAzureStorageAccountKeyOk returns a tuple with the AzureStorageAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureStorageAccountKey

`func (o *AzureBlobOptionsAuthenticationOptions) SetAzureStorageAccountKey(v string)`

SetAzureStorageAccountKey sets AzureStorageAccountKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


