# SourceControlSettingsPutRequestSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoBranchNamingEnabled** | Pointer to **bool** | When enabled, Retool automatically suggests a branch name on branch creation. Defaults to true. | [optional] 
**CustomPullRequestTemplateEnabled** | Pointer to **bool** | When enabled, Retool will use the template specified to create pull requests. Defaults to false. | [optional] 
**CustomPullRequestTemplate** | Pointer to **string** | Pull requests created from Retool will use the template specified. | [optional] 
**VersionControlLocked** | Pointer to **bool** | When set to true, creates a read-only instance of Retool, where app editing is disabled. Defaults to false. | [optional] 
**ForceUuidMapping** | Pointer to **bool** | When set to true, creates a uuid mapping for protected elements to be used in the source control repo. Defaults to false. | [optional] 

## Methods

### NewSourceControlSettingsPutRequestSettings

`func NewSourceControlSettingsPutRequestSettings() *SourceControlSettingsPutRequestSettings`

NewSourceControlSettingsPutRequestSettings instantiates a new SourceControlSettingsPutRequestSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlSettingsPutRequestSettingsWithDefaults

`func NewSourceControlSettingsPutRequestSettingsWithDefaults() *SourceControlSettingsPutRequestSettings`

NewSourceControlSettingsPutRequestSettingsWithDefaults instantiates a new SourceControlSettingsPutRequestSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoBranchNamingEnabled

`func (o *SourceControlSettingsPutRequestSettings) GetAutoBranchNamingEnabled() bool`

GetAutoBranchNamingEnabled returns the AutoBranchNamingEnabled field if non-nil, zero value otherwise.

### GetAutoBranchNamingEnabledOk

`func (o *SourceControlSettingsPutRequestSettings) GetAutoBranchNamingEnabledOk() (*bool, bool)`

GetAutoBranchNamingEnabledOk returns a tuple with the AutoBranchNamingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoBranchNamingEnabled

`func (o *SourceControlSettingsPutRequestSettings) SetAutoBranchNamingEnabled(v bool)`

SetAutoBranchNamingEnabled sets AutoBranchNamingEnabled field to given value.

### HasAutoBranchNamingEnabled

`func (o *SourceControlSettingsPutRequestSettings) HasAutoBranchNamingEnabled() bool`

HasAutoBranchNamingEnabled returns a boolean if a field has been set.

### GetCustomPullRequestTemplateEnabled

`func (o *SourceControlSettingsPutRequestSettings) GetCustomPullRequestTemplateEnabled() bool`

GetCustomPullRequestTemplateEnabled returns the CustomPullRequestTemplateEnabled field if non-nil, zero value otherwise.

### GetCustomPullRequestTemplateEnabledOk

`func (o *SourceControlSettingsPutRequestSettings) GetCustomPullRequestTemplateEnabledOk() (*bool, bool)`

GetCustomPullRequestTemplateEnabledOk returns a tuple with the CustomPullRequestTemplateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPullRequestTemplateEnabled

`func (o *SourceControlSettingsPutRequestSettings) SetCustomPullRequestTemplateEnabled(v bool)`

SetCustomPullRequestTemplateEnabled sets CustomPullRequestTemplateEnabled field to given value.

### HasCustomPullRequestTemplateEnabled

`func (o *SourceControlSettingsPutRequestSettings) HasCustomPullRequestTemplateEnabled() bool`

HasCustomPullRequestTemplateEnabled returns a boolean if a field has been set.

### GetCustomPullRequestTemplate

`func (o *SourceControlSettingsPutRequestSettings) GetCustomPullRequestTemplate() string`

GetCustomPullRequestTemplate returns the CustomPullRequestTemplate field if non-nil, zero value otherwise.

### GetCustomPullRequestTemplateOk

`func (o *SourceControlSettingsPutRequestSettings) GetCustomPullRequestTemplateOk() (*string, bool)`

GetCustomPullRequestTemplateOk returns a tuple with the CustomPullRequestTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPullRequestTemplate

`func (o *SourceControlSettingsPutRequestSettings) SetCustomPullRequestTemplate(v string)`

SetCustomPullRequestTemplate sets CustomPullRequestTemplate field to given value.

### HasCustomPullRequestTemplate

`func (o *SourceControlSettingsPutRequestSettings) HasCustomPullRequestTemplate() bool`

HasCustomPullRequestTemplate returns a boolean if a field has been set.

### GetVersionControlLocked

`func (o *SourceControlSettingsPutRequestSettings) GetVersionControlLocked() bool`

GetVersionControlLocked returns the VersionControlLocked field if non-nil, zero value otherwise.

### GetVersionControlLockedOk

`func (o *SourceControlSettingsPutRequestSettings) GetVersionControlLockedOk() (*bool, bool)`

GetVersionControlLockedOk returns a tuple with the VersionControlLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionControlLocked

`func (o *SourceControlSettingsPutRequestSettings) SetVersionControlLocked(v bool)`

SetVersionControlLocked sets VersionControlLocked field to given value.

### HasVersionControlLocked

`func (o *SourceControlSettingsPutRequestSettings) HasVersionControlLocked() bool`

HasVersionControlLocked returns a boolean if a field has been set.

### GetForceUuidMapping

`func (o *SourceControlSettingsPutRequestSettings) GetForceUuidMapping() bool`

GetForceUuidMapping returns the ForceUuidMapping field if non-nil, zero value otherwise.

### GetForceUuidMappingOk

`func (o *SourceControlSettingsPutRequestSettings) GetForceUuidMappingOk() (*bool, bool)`

GetForceUuidMappingOk returns a tuple with the ForceUuidMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceUuidMapping

`func (o *SourceControlSettingsPutRequestSettings) SetForceUuidMapping(v bool)`

SetForceUuidMapping sets ForceUuidMapping field to given value.

### HasForceUuidMapping

`func (o *SourceControlSettingsPutRequestSettings) HasForceUuidMapping() bool`

HasForceUuidMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


