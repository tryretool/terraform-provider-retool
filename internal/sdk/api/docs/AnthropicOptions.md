# AnthropicOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeySource** | **string** | &#39;retoolManaged&#39; — uses Retool&#39;s managed credentials; api_key, base_url, and custom_headers must not be provided. | 
**BaseUrl** | **string** |  | [default to "https://api.anthropic.com"]
**ApiKey** | **string** |  | 
**CustomHeaders** | Pointer to **[][]string** |  | [optional] 

## Methods

### NewAnthropicOptions

`func NewAnthropicOptions(keySource string, baseUrl string, apiKey string, ) *AnthropicOptions`

NewAnthropicOptions instantiates a new AnthropicOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnthropicOptionsWithDefaults

`func NewAnthropicOptionsWithDefaults() *AnthropicOptions`

NewAnthropicOptionsWithDefaults instantiates a new AnthropicOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeySource

`func (o *AnthropicOptions) GetKeySource() string`

GetKeySource returns the KeySource field if non-nil, zero value otherwise.

### GetKeySourceOk

`func (o *AnthropicOptions) GetKeySourceOk() (*string, bool)`

GetKeySourceOk returns a tuple with the KeySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySource

`func (o *AnthropicOptions) SetKeySource(v string)`

SetKeySource sets KeySource field to given value.


### GetBaseUrl

`func (o *AnthropicOptions) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *AnthropicOptions) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *AnthropicOptions) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetApiKey

`func (o *AnthropicOptions) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AnthropicOptions) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AnthropicOptions) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetCustomHeaders

`func (o *AnthropicOptions) GetCustomHeaders() [][]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *AnthropicOptions) GetCustomHeadersOk() (*[][]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *AnthropicOptions) SetCustomHeaders(v [][]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *AnthropicOptions) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


