# SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**PersonalAccessToken** | **string** | The GitHub project access token to authenticate to the GitHub API.  | 
**Url** | Pointer to **string** | The domain used to access your self-hosted GitHub instance. | [optional] 
**EnterpriseApiUrl** | Pointer to **string** | The REST API route for your self-hosted GitHub instance. Defaults to https://[hostname]/api/v3. | [optional] 

## Methods

### NewSourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1

`func NewSourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1(type_ string, personalAccessToken string, ) *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1`

NewSourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1 instantiates a new SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1WithDefaults

`func NewSourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1WithDefaults() *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1`

NewSourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1WithDefaults instantiates a new SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1) SetType(v string)`

SetType sets Type field to given value.


### GetPersonalAccessToken

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1) GetPersonalAccessToken() string`

GetPersonalAccessToken returns the PersonalAccessToken field if non-nil, zero value otherwise.

### GetPersonalAccessTokenOk

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1) GetPersonalAccessTokenOk() (*string, bool)`

GetPersonalAccessTokenOk returns a tuple with the PersonalAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalAccessToken

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1) SetPersonalAccessToken(v string)`

SetPersonalAccessToken sets PersonalAccessToken field to given value.


### GetUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnterpriseApiUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1) GetEnterpriseApiUrl() string`

GetEnterpriseApiUrl returns the EnterpriseApiUrl field if non-nil, zero value otherwise.

### GetEnterpriseApiUrlOk

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1) GetEnterpriseApiUrlOk() (*string, bool)`

GetEnterpriseApiUrlOk returns a tuple with the EnterpriseApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseApiUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1) SetEnterpriseApiUrl(v string)`

SetEnterpriseApiUrl sets EnterpriseApiUrl field to given value.

### HasEnterpriseApiUrl

`func (o *SourceControlConfigGet200ResponseDataAnyOfConfigAnyOf1) HasEnterpriseApiUrl() bool`

HasEnterpriseApiUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


