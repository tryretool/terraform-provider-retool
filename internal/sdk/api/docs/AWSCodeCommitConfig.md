# AWSCodeCommitConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The domain used to access your self-hosted AWS CodeCommit instance. | 
**Region** | **string** | The region of the CodeCommit repository. | 
**AccessKeyId** | Pointer to **string** | The Access key ID from your AWSCodeCommitFullAccess policy. Optional when &#x60;auth_with_default_credential_provider_chain&#x60; is true (gated behind the &#x60;sourceControlCodeCommitIamRole&#x60; feature flag). | [optional] 
**SecretAccessKey** | Pointer to **string** | The Secret Access Key from your AWSCodeCommitFullAccess policy. Optional when &#x60;auth_with_default_credential_provider_chain&#x60; is true. | [optional] 
**HttpsUsername** | Pointer to **string** | The HTTPS username from your security credentials. Optional when &#x60;auth_with_default_credential_provider_chain&#x60; is true (Retool will SigV4-sign the git URL). | [optional] 
**HttpsPassword** | Pointer to **string** | The HTTPS password from your security credentials. Optional when &#x60;auth_with_default_credential_provider_chain&#x60; is true. | [optional] 
**AuthWithDefaultCredentialProviderChain** | Pointer to **bool** | When true, the backend resolves AWS credentials from the default provider chain instead of the static keys. Requires the &#x60;sourceControlCodeCommitIamRole&#x60; feature flag to be enabled. | [optional] 
**AssumeRole** | Pointer to **string** | Optional STS role ARN to assume on top of the resolved credentials when using &#x60;auth_with_default_credential_provider_chain&#x60;. | [optional] 

## Methods

### NewAWSCodeCommitConfig

`func NewAWSCodeCommitConfig(url string, region string, ) *AWSCodeCommitConfig`

NewAWSCodeCommitConfig instantiates a new AWSCodeCommitConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSCodeCommitConfigWithDefaults

`func NewAWSCodeCommitConfigWithDefaults() *AWSCodeCommitConfig`

NewAWSCodeCommitConfigWithDefaults instantiates a new AWSCodeCommitConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AWSCodeCommitConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AWSCodeCommitConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AWSCodeCommitConfig) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRegion

`func (o *AWSCodeCommitConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AWSCodeCommitConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AWSCodeCommitConfig) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAccessKeyId

`func (o *AWSCodeCommitConfig) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AWSCodeCommitConfig) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AWSCodeCommitConfig) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AWSCodeCommitConfig) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *AWSCodeCommitConfig) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *AWSCodeCommitConfig) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *AWSCodeCommitConfig) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *AWSCodeCommitConfig) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetHttpsUsername

`func (o *AWSCodeCommitConfig) GetHttpsUsername() string`

GetHttpsUsername returns the HttpsUsername field if non-nil, zero value otherwise.

### GetHttpsUsernameOk

`func (o *AWSCodeCommitConfig) GetHttpsUsernameOk() (*string, bool)`

GetHttpsUsernameOk returns a tuple with the HttpsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsUsername

`func (o *AWSCodeCommitConfig) SetHttpsUsername(v string)`

SetHttpsUsername sets HttpsUsername field to given value.

### HasHttpsUsername

`func (o *AWSCodeCommitConfig) HasHttpsUsername() bool`

HasHttpsUsername returns a boolean if a field has been set.

### GetHttpsPassword

`func (o *AWSCodeCommitConfig) GetHttpsPassword() string`

GetHttpsPassword returns the HttpsPassword field if non-nil, zero value otherwise.

### GetHttpsPasswordOk

`func (o *AWSCodeCommitConfig) GetHttpsPasswordOk() (*string, bool)`

GetHttpsPasswordOk returns a tuple with the HttpsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsPassword

`func (o *AWSCodeCommitConfig) SetHttpsPassword(v string)`

SetHttpsPassword sets HttpsPassword field to given value.

### HasHttpsPassword

`func (o *AWSCodeCommitConfig) HasHttpsPassword() bool`

HasHttpsPassword returns a boolean if a field has been set.

### GetAuthWithDefaultCredentialProviderChain

`func (o *AWSCodeCommitConfig) GetAuthWithDefaultCredentialProviderChain() bool`

GetAuthWithDefaultCredentialProviderChain returns the AuthWithDefaultCredentialProviderChain field if non-nil, zero value otherwise.

### GetAuthWithDefaultCredentialProviderChainOk

`func (o *AWSCodeCommitConfig) GetAuthWithDefaultCredentialProviderChainOk() (*bool, bool)`

GetAuthWithDefaultCredentialProviderChainOk returns a tuple with the AuthWithDefaultCredentialProviderChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthWithDefaultCredentialProviderChain

`func (o *AWSCodeCommitConfig) SetAuthWithDefaultCredentialProviderChain(v bool)`

SetAuthWithDefaultCredentialProviderChain sets AuthWithDefaultCredentialProviderChain field to given value.

### HasAuthWithDefaultCredentialProviderChain

`func (o *AWSCodeCommitConfig) HasAuthWithDefaultCredentialProviderChain() bool`

HasAuthWithDefaultCredentialProviderChain returns a boolean if a field has been set.

### GetAssumeRole

`func (o *AWSCodeCommitConfig) GetAssumeRole() string`

GetAssumeRole returns the AssumeRole field if non-nil, zero value otherwise.

### GetAssumeRoleOk

`func (o *AWSCodeCommitConfig) GetAssumeRoleOk() (*string, bool)`

GetAssumeRoleOk returns a tuple with the AssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumeRole

`func (o *AWSCodeCommitConfig) SetAssumeRole(v string)`

SetAssumeRole sets AssumeRole field to given value.

### HasAssumeRole

`func (o *AWSCodeCommitConfig) HasAssumeRole() bool`

HasAssumeRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


