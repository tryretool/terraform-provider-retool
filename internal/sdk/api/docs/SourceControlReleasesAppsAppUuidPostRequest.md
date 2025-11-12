# SourceControlReleasesAppsAppUuidPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReleaseVersion** | **string** | The version of the release. | 
**ReleaseDescription** | Pointer to **string** | The description of the release. | [optional] 
**CommitMessage** | Pointer to **string** | Message to use for the commit that updates the specified manifest. If a message is not provided, a default will be used. | [optional] 

## Methods

### NewSourceControlReleasesAppsAppUuidPostRequest

`func NewSourceControlReleasesAppsAppUuidPostRequest(releaseVersion string, ) *SourceControlReleasesAppsAppUuidPostRequest`

NewSourceControlReleasesAppsAppUuidPostRequest instantiates a new SourceControlReleasesAppsAppUuidPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlReleasesAppsAppUuidPostRequestWithDefaults

`func NewSourceControlReleasesAppsAppUuidPostRequestWithDefaults() *SourceControlReleasesAppsAppUuidPostRequest`

NewSourceControlReleasesAppsAppUuidPostRequestWithDefaults instantiates a new SourceControlReleasesAppsAppUuidPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReleaseVersion

`func (o *SourceControlReleasesAppsAppUuidPostRequest) GetReleaseVersion() string`

GetReleaseVersion returns the ReleaseVersion field if non-nil, zero value otherwise.

### GetReleaseVersionOk

`func (o *SourceControlReleasesAppsAppUuidPostRequest) GetReleaseVersionOk() (*string, bool)`

GetReleaseVersionOk returns a tuple with the ReleaseVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseVersion

`func (o *SourceControlReleasesAppsAppUuidPostRequest) SetReleaseVersion(v string)`

SetReleaseVersion sets ReleaseVersion field to given value.


### GetReleaseDescription

`func (o *SourceControlReleasesAppsAppUuidPostRequest) GetReleaseDescription() string`

GetReleaseDescription returns the ReleaseDescription field if non-nil, zero value otherwise.

### GetReleaseDescriptionOk

`func (o *SourceControlReleasesAppsAppUuidPostRequest) GetReleaseDescriptionOk() (*string, bool)`

GetReleaseDescriptionOk returns a tuple with the ReleaseDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDescription

`func (o *SourceControlReleasesAppsAppUuidPostRequest) SetReleaseDescription(v string)`

SetReleaseDescription sets ReleaseDescription field to given value.

### HasReleaseDescription

`func (o *SourceControlReleasesAppsAppUuidPostRequest) HasReleaseDescription() bool`

HasReleaseDescription returns a boolean if a field has been set.

### GetCommitMessage

`func (o *SourceControlReleasesAppsAppUuidPostRequest) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *SourceControlReleasesAppsAppUuidPostRequest) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *SourceControlReleasesAppsAppUuidPostRequest) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.

### HasCommitMessage

`func (o *SourceControlReleasesAppsAppUuidPostRequest) HasCommitMessage() bool`

HasCommitMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


