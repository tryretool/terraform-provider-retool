# GoogleVertexAIOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeySource** | **string** | Always &#39;byok&#39; (bring your own key) — this provider requires you to supply credentials. Defaults to &#39;byok&#39; and may be omitted. | [default to "byok"]
**ProjectId** | **string** |  | 
**Location** | **string** |  | 
**ServiceAccountKey** | **string** |  | 

## Methods

### NewGoogleVertexAIOptions

`func NewGoogleVertexAIOptions(keySource string, projectId string, location string, serviceAccountKey string, ) *GoogleVertexAIOptions`

NewGoogleVertexAIOptions instantiates a new GoogleVertexAIOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleVertexAIOptionsWithDefaults

`func NewGoogleVertexAIOptionsWithDefaults() *GoogleVertexAIOptions`

NewGoogleVertexAIOptionsWithDefaults instantiates a new GoogleVertexAIOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeySource

`func (o *GoogleVertexAIOptions) GetKeySource() string`

GetKeySource returns the KeySource field if non-nil, zero value otherwise.

### GetKeySourceOk

`func (o *GoogleVertexAIOptions) GetKeySourceOk() (*string, bool)`

GetKeySourceOk returns a tuple with the KeySource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySource

`func (o *GoogleVertexAIOptions) SetKeySource(v string)`

SetKeySource sets KeySource field to given value.


### GetProjectId

`func (o *GoogleVertexAIOptions) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GoogleVertexAIOptions) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GoogleVertexAIOptions) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetLocation

`func (o *GoogleVertexAIOptions) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GoogleVertexAIOptions) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GoogleVertexAIOptions) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetServiceAccountKey

`func (o *GoogleVertexAIOptions) GetServiceAccountKey() string`

GetServiceAccountKey returns the ServiceAccountKey field if non-nil, zero value otherwise.

### GetServiceAccountKeyOk

`func (o *GoogleVertexAIOptions) GetServiceAccountKeyOk() (*string, bool)`

GetServiceAccountKeyOk returns a tuple with the ServiceAccountKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountKey

`func (o *GoogleVertexAIOptions) SetServiceAccountKey(v string)`

SetServiceAccountKey sets ServiceAccountKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


