# AWSBedrockOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeySource** | **string** | Always &#39;byok&#39; (bring your own key) — this provider requires you to supply credentials. Defaults to &#39;byok&#39; and may be omitted. | [default to "byok"]
**Region** | **string** |  | 
**AccessKeyId** | Pointer to **string** |  | [optional] 
**SecretAccessKey** | Pointer to **string** |  | [optional] 
**SessionToken** | Pointer to **string** |  | [optional] 
**AssumeRole** | Pointer to **string** |  | [optional] 
**AuthWithDefaultCredentialProviderChain** | Pointer to **bool** | Retool will source AWS credentials from the credential provider chain. Use this option to authenticate with credentials provided in environment variables or the underlying instance role. | [optional] 

## Methods

### NewAWSBedrockOptions

`func NewAWSBedrockOptions(keySource string, region string, ) *AWSBedrockOptions`

NewAWSBedrockOptions instantiates a new AWSBedrockOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSBedrockOptionsWithDefaults

`func NewAWSBedrockOptionsWithDefaults() *AWSBedrockOptions`

NewAWSBedrockOptionsWithDefaults instantiates a new AWSBedrockOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeySource

`func (o *AWSBedrockOptions) GetKeySource() string`

GetKeySource returns the KeySource field if non-nil, zero value otherwise.

### GetKeySourceOk

`func (o *AWSBedrockOptions) GetKeySourceOk() (*string, bool)`

GetKeySourceOk returns a tuple with the KeySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySource

`func (o *AWSBedrockOptions) SetKeySource(v string)`

SetKeySource sets KeySource field to given value.


### GetRegion

`func (o *AWSBedrockOptions) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AWSBedrockOptions) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AWSBedrockOptions) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAccessKeyId

`func (o *AWSBedrockOptions) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *AWSBedrockOptions) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *AWSBedrockOptions) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *AWSBedrockOptions) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *AWSBedrockOptions) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *AWSBedrockOptions) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *AWSBedrockOptions) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *AWSBedrockOptions) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetSessionToken

`func (o *AWSBedrockOptions) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *AWSBedrockOptions) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *AWSBedrockOptions) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *AWSBedrockOptions) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetAssumeRole

`func (o *AWSBedrockOptions) GetAssumeRole() string`

GetAssumeRole returns the AssumeRole field if non-nil, zero value otherwise.

### GetAssumeRoleOk

`func (o *AWSBedrockOptions) GetAssumeRoleOk() (*string, bool)`

GetAssumeRoleOk returns a tuple with the AssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumeRole

`func (o *AWSBedrockOptions) SetAssumeRole(v string)`

SetAssumeRole sets AssumeRole field to given value.

### HasAssumeRole

`func (o *AWSBedrockOptions) HasAssumeRole() bool`

HasAssumeRole returns a boolean if a field has been set.

### GetAuthWithDefaultCredentialProviderChain

`func (o *AWSBedrockOptions) GetAuthWithDefaultCredentialProviderChain() bool`

GetAuthWithDefaultCredentialProviderChain returns the AuthWithDefaultCredentialProviderChain field if non-nil, zero value otherwise.

### GetAuthWithDefaultCredentialProviderChainOk

`func (o *AWSBedrockOptions) GetAuthWithDefaultCredentialProviderChainOk() (*bool, bool)`

GetAuthWithDefaultCredentialProviderChainOk returns a tuple with the AuthWithDefaultCredentialProviderChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthWithDefaultCredentialProviderChain

`func (o *AWSBedrockOptions) SetAuthWithDefaultCredentialProviderChain(v bool)`

SetAuthWithDefaultCredentialProviderChain sets AuthWithDefaultCredentialProviderChain field to given value.

### HasAuthWithDefaultCredentialProviderChain

`func (o *AWSBedrockOptions) HasAuthWithDefaultCredentialProviderChain() bool`

HasAuthWithDefaultCredentialProviderChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


