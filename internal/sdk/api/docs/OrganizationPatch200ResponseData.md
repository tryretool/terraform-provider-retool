# OrganizationPatch200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the organization. | 
**RequestAccessEnabled** | **bool** | Whether users can request access to join the organization. | 
**AiSupportBotDisabled** | **bool** | Whether the AI support bot is disabled. | 
**RetoolFormsDisabled** | **bool** | Whether Retool forms are disabled. | 
**ReleaseManagementEnabled** | **bool** | Whether versions and releases on apps are disabled. | 
**CacheQueriesPerUser** | **NullableBool** |  | 
**ApplyPreloadedCssToHomepage** | **bool** | Whether preloaded CSS is applied to the Retool homepage. | 
**PreloadedCss** | **NullableString** | Custom CSS rules to apply across Retool. | 
**PreloadedJavascript** | **NullableString** | Preloaded JavaScript that will apply to every Retool app. | 
**JavascriptLinks** | **[]string** | List of custom JavaScript libraries to load in every app. | 
**WorkflowRunRetentionPeriodMins** | **float32** | Number of minutes to store workflow run history data (up to a max of 90 days on cloud, 1 year on-prem) | 
**AppOwnersPermissionsManagement** | **bool** | Whether app owners can manage permissions for their apps directly. | 
**TwoFactorAuthRequired** | **bool** | Whether two-factor authentication is required for all users in the organization. | 
**TwoFactorAuthType** | **NullableString** | Required 2FA type, applies to the whole organization | 
**DisableNewLoginIpNotificationEmail** | **bool** | Whether notification emails for logins from new IPs are disabled. | 

## Methods

### NewOrganizationPatch200ResponseData

`func NewOrganizationPatch200ResponseData(id string, requestAccessEnabled bool, aiSupportBotDisabled bool, retoolFormsDisabled bool, releaseManagementEnabled bool, cacheQueriesPerUser NullableBool, applyPreloadedCssToHomepage bool, preloadedCss NullableString, preloadedJavascript NullableString, javascriptLinks []string, workflowRunRetentionPeriodMins float32, appOwnersPermissionsManagement bool, twoFactorAuthRequired bool, twoFactorAuthType NullableString, disableNewLoginIpNotificationEmail bool, ) *OrganizationPatch200ResponseData`

NewOrganizationPatch200ResponseData instantiates a new OrganizationPatch200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationPatch200ResponseDataWithDefaults

`func NewOrganizationPatch200ResponseDataWithDefaults() *OrganizationPatch200ResponseData`

NewOrganizationPatch200ResponseDataWithDefaults instantiates a new OrganizationPatch200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationPatch200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationPatch200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationPatch200ResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetRequestAccessEnabled

`func (o *OrganizationPatch200ResponseData) GetRequestAccessEnabled() bool`

GetRequestAccessEnabled returns the RequestAccessEnabled field if non-nil, zero value otherwise.

### GetRequestAccessEnabledOk

`func (o *OrganizationPatch200ResponseData) GetRequestAccessEnabledOk() (*bool, bool)`

GetRequestAccessEnabledOk returns a tuple with the RequestAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAccessEnabled

`func (o *OrganizationPatch200ResponseData) SetRequestAccessEnabled(v bool)`

SetRequestAccessEnabled sets RequestAccessEnabled field to given value.


### GetAiSupportBotDisabled

`func (o *OrganizationPatch200ResponseData) GetAiSupportBotDisabled() bool`

GetAiSupportBotDisabled returns the AiSupportBotDisabled field if non-nil, zero value otherwise.

### GetAiSupportBotDisabledOk

`func (o *OrganizationPatch200ResponseData) GetAiSupportBotDisabledOk() (*bool, bool)`

GetAiSupportBotDisabledOk returns a tuple with the AiSupportBotDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSupportBotDisabled

`func (o *OrganizationPatch200ResponseData) SetAiSupportBotDisabled(v bool)`

SetAiSupportBotDisabled sets AiSupportBotDisabled field to given value.


### GetRetoolFormsDisabled

`func (o *OrganizationPatch200ResponseData) GetRetoolFormsDisabled() bool`

GetRetoolFormsDisabled returns the RetoolFormsDisabled field if non-nil, zero value otherwise.

### GetRetoolFormsDisabledOk

`func (o *OrganizationPatch200ResponseData) GetRetoolFormsDisabledOk() (*bool, bool)`

GetRetoolFormsDisabledOk returns a tuple with the RetoolFormsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetoolFormsDisabled

`func (o *OrganizationPatch200ResponseData) SetRetoolFormsDisabled(v bool)`

SetRetoolFormsDisabled sets RetoolFormsDisabled field to given value.


### GetReleaseManagementEnabled

`func (o *OrganizationPatch200ResponseData) GetReleaseManagementEnabled() bool`

GetReleaseManagementEnabled returns the ReleaseManagementEnabled field if non-nil, zero value otherwise.

### GetReleaseManagementEnabledOk

`func (o *OrganizationPatch200ResponseData) GetReleaseManagementEnabledOk() (*bool, bool)`

GetReleaseManagementEnabledOk returns a tuple with the ReleaseManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseManagementEnabled

`func (o *OrganizationPatch200ResponseData) SetReleaseManagementEnabled(v bool)`

SetReleaseManagementEnabled sets ReleaseManagementEnabled field to given value.


### GetCacheQueriesPerUser

`func (o *OrganizationPatch200ResponseData) GetCacheQueriesPerUser() bool`

GetCacheQueriesPerUser returns the CacheQueriesPerUser field if non-nil, zero value otherwise.

### GetCacheQueriesPerUserOk

`func (o *OrganizationPatch200ResponseData) GetCacheQueriesPerUserOk() (*bool, bool)`

GetCacheQueriesPerUserOk returns a tuple with the CacheQueriesPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheQueriesPerUser

`func (o *OrganizationPatch200ResponseData) SetCacheQueriesPerUser(v bool)`

SetCacheQueriesPerUser sets CacheQueriesPerUser field to given value.


### SetCacheQueriesPerUserNil

`func (o *OrganizationPatch200ResponseData) SetCacheQueriesPerUserNil(b bool)`

 SetCacheQueriesPerUserNil sets the value for CacheQueriesPerUser to be an explicit nil

### UnsetCacheQueriesPerUser
`func (o *OrganizationPatch200ResponseData) UnsetCacheQueriesPerUser()`

UnsetCacheQueriesPerUser ensures that no value is present for CacheQueriesPerUser, not even an explicit nil
### GetApplyPreloadedCssToHomepage

`func (o *OrganizationPatch200ResponseData) GetApplyPreloadedCssToHomepage() bool`

GetApplyPreloadedCssToHomepage returns the ApplyPreloadedCssToHomepage field if non-nil, zero value otherwise.

### GetApplyPreloadedCssToHomepageOk

`func (o *OrganizationPatch200ResponseData) GetApplyPreloadedCssToHomepageOk() (*bool, bool)`

GetApplyPreloadedCssToHomepageOk returns a tuple with the ApplyPreloadedCssToHomepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPreloadedCssToHomepage

`func (o *OrganizationPatch200ResponseData) SetApplyPreloadedCssToHomepage(v bool)`

SetApplyPreloadedCssToHomepage sets ApplyPreloadedCssToHomepage field to given value.


### GetPreloadedCss

`func (o *OrganizationPatch200ResponseData) GetPreloadedCss() string`

GetPreloadedCss returns the PreloadedCss field if non-nil, zero value otherwise.

### GetPreloadedCssOk

`func (o *OrganizationPatch200ResponseData) GetPreloadedCssOk() (*string, bool)`

GetPreloadedCssOk returns a tuple with the PreloadedCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreloadedCss

`func (o *OrganizationPatch200ResponseData) SetPreloadedCss(v string)`

SetPreloadedCss sets PreloadedCss field to given value.


### SetPreloadedCssNil

`func (o *OrganizationPatch200ResponseData) SetPreloadedCssNil(b bool)`

 SetPreloadedCssNil sets the value for PreloadedCss to be an explicit nil

### UnsetPreloadedCss
`func (o *OrganizationPatch200ResponseData) UnsetPreloadedCss()`

UnsetPreloadedCss ensures that no value is present for PreloadedCss, not even an explicit nil
### GetPreloadedJavascript

`func (o *OrganizationPatch200ResponseData) GetPreloadedJavascript() string`

GetPreloadedJavascript returns the PreloadedJavascript field if non-nil, zero value otherwise.

### GetPreloadedJavascriptOk

`func (o *OrganizationPatch200ResponseData) GetPreloadedJavascriptOk() (*string, bool)`

GetPreloadedJavascriptOk returns a tuple with the PreloadedJavascript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreloadedJavascript

`func (o *OrganizationPatch200ResponseData) SetPreloadedJavascript(v string)`

SetPreloadedJavascript sets PreloadedJavascript field to given value.


### SetPreloadedJavascriptNil

`func (o *OrganizationPatch200ResponseData) SetPreloadedJavascriptNil(b bool)`

 SetPreloadedJavascriptNil sets the value for PreloadedJavascript to be an explicit nil

### UnsetPreloadedJavascript
`func (o *OrganizationPatch200ResponseData) UnsetPreloadedJavascript()`

UnsetPreloadedJavascript ensures that no value is present for PreloadedJavascript, not even an explicit nil
### GetJavascriptLinks

`func (o *OrganizationPatch200ResponseData) GetJavascriptLinks() []string`

GetJavascriptLinks returns the JavascriptLinks field if non-nil, zero value otherwise.

### GetJavascriptLinksOk

`func (o *OrganizationPatch200ResponseData) GetJavascriptLinksOk() (*[]string, bool)`

GetJavascriptLinksOk returns a tuple with the JavascriptLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavascriptLinks

`func (o *OrganizationPatch200ResponseData) SetJavascriptLinks(v []string)`

SetJavascriptLinks sets JavascriptLinks field to given value.


### GetWorkflowRunRetentionPeriodMins

`func (o *OrganizationPatch200ResponseData) GetWorkflowRunRetentionPeriodMins() float32`

GetWorkflowRunRetentionPeriodMins returns the WorkflowRunRetentionPeriodMins field if non-nil, zero value otherwise.

### GetWorkflowRunRetentionPeriodMinsOk

`func (o *OrganizationPatch200ResponseData) GetWorkflowRunRetentionPeriodMinsOk() (*float32, bool)`

GetWorkflowRunRetentionPeriodMinsOk returns a tuple with the WorkflowRunRetentionPeriodMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRunRetentionPeriodMins

`func (o *OrganizationPatch200ResponseData) SetWorkflowRunRetentionPeriodMins(v float32)`

SetWorkflowRunRetentionPeriodMins sets WorkflowRunRetentionPeriodMins field to given value.


### GetAppOwnersPermissionsManagement

`func (o *OrganizationPatch200ResponseData) GetAppOwnersPermissionsManagement() bool`

GetAppOwnersPermissionsManagement returns the AppOwnersPermissionsManagement field if non-nil, zero value otherwise.

### GetAppOwnersPermissionsManagementOk

`func (o *OrganizationPatch200ResponseData) GetAppOwnersPermissionsManagementOk() (*bool, bool)`

GetAppOwnersPermissionsManagementOk returns a tuple with the AppOwnersPermissionsManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppOwnersPermissionsManagement

`func (o *OrganizationPatch200ResponseData) SetAppOwnersPermissionsManagement(v bool)`

SetAppOwnersPermissionsManagement sets AppOwnersPermissionsManagement field to given value.


### GetTwoFactorAuthRequired

`func (o *OrganizationPatch200ResponseData) GetTwoFactorAuthRequired() bool`

GetTwoFactorAuthRequired returns the TwoFactorAuthRequired field if non-nil, zero value otherwise.

### GetTwoFactorAuthRequiredOk

`func (o *OrganizationPatch200ResponseData) GetTwoFactorAuthRequiredOk() (*bool, bool)`

GetTwoFactorAuthRequiredOk returns a tuple with the TwoFactorAuthRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthRequired

`func (o *OrganizationPatch200ResponseData) SetTwoFactorAuthRequired(v bool)`

SetTwoFactorAuthRequired sets TwoFactorAuthRequired field to given value.


### GetTwoFactorAuthType

`func (o *OrganizationPatch200ResponseData) GetTwoFactorAuthType() string`

GetTwoFactorAuthType returns the TwoFactorAuthType field if non-nil, zero value otherwise.

### GetTwoFactorAuthTypeOk

`func (o *OrganizationPatch200ResponseData) GetTwoFactorAuthTypeOk() (*string, bool)`

GetTwoFactorAuthTypeOk returns a tuple with the TwoFactorAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthType

`func (o *OrganizationPatch200ResponseData) SetTwoFactorAuthType(v string)`

SetTwoFactorAuthType sets TwoFactorAuthType field to given value.


### SetTwoFactorAuthTypeNil

`func (o *OrganizationPatch200ResponseData) SetTwoFactorAuthTypeNil(b bool)`

 SetTwoFactorAuthTypeNil sets the value for TwoFactorAuthType to be an explicit nil

### UnsetTwoFactorAuthType
`func (o *OrganizationPatch200ResponseData) UnsetTwoFactorAuthType()`

UnsetTwoFactorAuthType ensures that no value is present for TwoFactorAuthType, not even an explicit nil
### GetDisableNewLoginIpNotificationEmail

`func (o *OrganizationPatch200ResponseData) GetDisableNewLoginIpNotificationEmail() bool`

GetDisableNewLoginIpNotificationEmail returns the DisableNewLoginIpNotificationEmail field if non-nil, zero value otherwise.

### GetDisableNewLoginIpNotificationEmailOk

`func (o *OrganizationPatch200ResponseData) GetDisableNewLoginIpNotificationEmailOk() (*bool, bool)`

GetDisableNewLoginIpNotificationEmailOk returns a tuple with the DisableNewLoginIpNotificationEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNewLoginIpNotificationEmail

`func (o *OrganizationPatch200ResponseData) SetDisableNewLoginIpNotificationEmail(v bool)`

SetDisableNewLoginIpNotificationEmail sets DisableNewLoginIpNotificationEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


