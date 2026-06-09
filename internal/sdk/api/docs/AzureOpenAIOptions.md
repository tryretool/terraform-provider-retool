# AzureOpenAIOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeySource** | **string** | Always &#39;byok&#39; (bring your own key) — this provider requires you to supply credentials. Defaults to &#39;byok&#39; and may be omitted. | [default to "byok"]
**BaseUrl** | **string** |  | 
**ApiKey** | **string** |  | 
**DeploymentNames** | **[]string** |  | 
**ApiVersion** | Pointer to **string** |  | [optional] 
**CustomHeaders** | Pointer to **[][]string** |  | [optional] 

## Methods

### NewAzureOpenAIOptions

`func NewAzureOpenAIOptions(keySource string, baseUrl string, apiKey string, deploymentNames []string, ) *AzureOpenAIOptions`

NewAzureOpenAIOptions instantiates a new AzureOpenAIOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureOpenAIOptionsWithDefaults

`func NewAzureOpenAIOptionsWithDefaults() *AzureOpenAIOptions`

NewAzureOpenAIOptionsWithDefaults instantiates a new AzureOpenAIOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeySource

`func (o *AzureOpenAIOptions) GetKeySource() string`

GetKeySource returns the KeySource field if non-nil, zero value otherwise.

### GetKeySourceOk

`func (o *AzureOpenAIOptions) GetKeySourceOk() (*string, bool)`

GetKeySourceOk returns a tuple with the KeySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySource

`func (o *AzureOpenAIOptions) SetKeySource(v string)`

SetKeySource sets KeySource field to given value.


### GetBaseUrl

`func (o *AzureOpenAIOptions) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *AzureOpenAIOptions) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *AzureOpenAIOptions) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetApiKey

`func (o *AzureOpenAIOptions) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AzureOpenAIOptions) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AzureOpenAIOptions) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetDeploymentNames

`func (o *AzureOpenAIOptions) GetDeploymentNames() []string`

GetDeploymentNames returns the DeploymentNames field if non-nil, zero value otherwise.

### GetDeploymentNamesOk

`func (o *AzureOpenAIOptions) GetDeploymentNamesOk() (*[]string, bool)`

GetDeploymentNamesOk returns a tuple with the DeploymentNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentNames

`func (o *AzureOpenAIOptions) SetDeploymentNames(v []string)`

SetDeploymentNames sets DeploymentNames field to given value.


### GetApiVersion

`func (o *AzureOpenAIOptions) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AzureOpenAIOptions) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AzureOpenAIOptions) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AzureOpenAIOptions) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *AzureOpenAIOptions) GetCustomHeaders() [][]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *AzureOpenAIOptions) GetCustomHeadersOk() (*[][]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *AzureOpenAIOptions) SetCustomHeaders(v [][]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *AzureOpenAIOptions) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


