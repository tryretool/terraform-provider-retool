# SourceControlManifestsManifestNameAppsAppUuidPutRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Release** | [**SourceControlManifestsManifestNameAppsAppUuidPutRequestRelease**](SourceControlManifestsManifestNameAppsAppUuidPutRequestRelease.md) |  | 
**CommitMessage** | Pointer to **string** | Message to use for the commit that updates the specified manifest. If a message is not provided, a default will be used. | [optional] 

## Methods

### NewSourceControlManifestsManifestNameAppsAppUuidPutRequest

`func NewSourceControlManifestsManifestNameAppsAppUuidPutRequest(release SourceControlManifestsManifestNameAppsAppUuidPutRequestRelease, ) *SourceControlManifestsManifestNameAppsAppUuidPutRequest`

NewSourceControlManifestsManifestNameAppsAppUuidPutRequest instantiates a new SourceControlManifestsManifestNameAppsAppUuidPutRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlManifestsManifestNameAppsAppUuidPutRequestWithDefaults

`func NewSourceControlManifestsManifestNameAppsAppUuidPutRequestWithDefaults() *SourceControlManifestsManifestNameAppsAppUuidPutRequest`

NewSourceControlManifestsManifestNameAppsAppUuidPutRequestWithDefaults instantiates a new SourceControlManifestsManifestNameAppsAppUuidPutRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelease

`func (o *SourceControlManifestsManifestNameAppsAppUuidPutRequest) GetRelease() SourceControlManifestsManifestNameAppsAppUuidPutRequestRelease`

GetRelease returns the Release field if non-nil, zero value otherwise.

### GetReleaseOk

`func (o *SourceControlManifestsManifestNameAppsAppUuidPutRequest) GetReleaseOk() (*SourceControlManifestsManifestNameAppsAppUuidPutRequestRelease, bool)`

GetReleaseOk returns a tuple with the Release field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelease

`func (o *SourceControlManifestsManifestNameAppsAppUuidPutRequest) SetRelease(v SourceControlManifestsManifestNameAppsAppUuidPutRequestRelease)`

SetRelease sets Release field to given value.


### GetCommitMessage

`func (o *SourceControlManifestsManifestNameAppsAppUuidPutRequest) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *SourceControlManifestsManifestNameAppsAppUuidPutRequest) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *SourceControlManifestsManifestNameAppsAppUuidPutRequest) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *SourceControlManifestsManifestNameAppsAppUuidPutRequest) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


