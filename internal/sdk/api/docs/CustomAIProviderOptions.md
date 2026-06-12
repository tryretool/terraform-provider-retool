# CustomAIProviderOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeySource** | **string** | Always &#39;byok&#39; (bring your own key) — this provider requires you to supply credentials. Defaults to &#39;byok&#39; and may be omitted. | [default to "byok"]
**BaseUrl** | **string** |  | 
**ApiKey** | **string** |  | 
**CompatibleSchema** | **string** |  | 
**ModelList** | **[]string** |  | 
**CustomHeaders** | Pointer to **[][]string** |  | [optional] 

## Methods

### NewCustomAIProviderOptions

`func NewCustomAIProviderOptions(keySource string, baseUrl string, apiKey string, compatibleSchema string, modelList []string, ) *CustomAIProviderOptions`

NewCustomAIProviderOptions instantiates a new CustomAIProviderOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAIProviderOptionsWithDefaults

`func NewCustomAIProviderOptionsWithDefaults() *CustomAIProviderOptions`

NewCustomAIProviderOptionsWithDefaults instantiates a new CustomAIProviderOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeySource

`func (o *CustomAIProviderOptions) GetKeySource() string`

GetKeySource returns the KeySource field if non-nil, zero value otherwise.

### GetKeySourceOk

`func (o *CustomAIProviderOptions) GetKeySourceOk() (*string, bool)`

GetKeySourceOk returns a tuple with the KeySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySource

`func (o *CustomAIProviderOptions) SetKeySource(v string)`

SetKeySource sets KeySource field to given value.


### GetBaseUrl

`func (o *CustomAIProviderOptions) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *CustomAIProviderOptions) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *CustomAIProviderOptions) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetApiKey

`func (o *CustomAIProviderOptions) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CustomAIProviderOptions) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CustomAIProviderOptions) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetCompatibleSchema

`func (o *CustomAIProviderOptions) GetCompatibleSchema() string`

GetCompatibleSchema returns the CompatibleSchema field if non-nil, zero value otherwise.

### GetCompatibleSchemaOk

`func (o *CustomAIProviderOptions) GetCompatibleSchemaOk() (*string, bool)`

GetCompatibleSchemaOk returns a tuple with the CompatibleSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleSchema

`func (o *CustomAIProviderOptions) SetCompatibleSchema(v string)`

SetCompatibleSchema sets CompatibleSchema field to given value.


### GetModelList

`func (o *CustomAIProviderOptions) GetModelList() []string`

GetModelList returns the ModelList field if non-nil, zero value otherwise.

### GetModelListOk

`func (o *CustomAIProviderOptions) GetModelListOk() (*[]string, bool)`

GetModelListOk returns a tuple with the ModelList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelList

`func (o *CustomAIProviderOptions) SetModelList(v []string)`

SetModelList sets ModelList field to given value.


### GetCustomHeaders

`func (o *CustomAIProviderOptions) GetCustomHeaders() [][]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *CustomAIProviderOptions) GetCustomHeadersOk() (*[][]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *CustomAIProviderOptions) SetCustomHeaders(v [][]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *CustomAIProviderOptions) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


