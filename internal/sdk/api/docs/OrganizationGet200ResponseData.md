# OrganizationGet200ResponseData

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

### NewOrganizationGet200ResponseData

`func NewOrganizationGet200ResponseData(id string, requestAccessEnabled bool, aiSupportBotDisabled bool, retoolFormsDisabled bool, releaseManagementEnabled bool, cacheQueriesPerUser NullableBool, applyPreloadedCssToHomepage bool, preloadedCss NullableString, preloadedJavascript NullableString, javascriptLinks []string, workflowRunRetentionPeriodMins float32, appOwnersPermissionsManagement bool, twoFactorAuthRequired bool, twoFactorAuthType NullableString, disableNewLoginIpNotificationEmail bool, ) *OrganizationGet200ResponseData`

NewOrganizationGet200ResponseData instantiates a new OrganizationGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationGet200ResponseDataWithDefaults

`func NewOrganizationGet200ResponseDataWithDefaults() *OrganizationGet200ResponseData`

NewOrganizationGet200ResponseDataWithDefaults instantiates a new OrganizationGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationGet200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationGet200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationGet200ResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetRequestAccessEnabled

`func (o *OrganizationGet200ResponseData) GetRequestAccessEnabled() bool`

GetRequestAccessEnabled returns the RequestAccessEnabled field if non-nil, zero value otherwise.

### GetRequestAccessEnabledOk

`func (o *OrganizationGet200ResponseData) GetRequestAccessEnabledOk() (*bool, bool)`

GetRequestAccessEnabledOk returns a tuple with the RequestAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestAccessEnabled

`func (o *OrganizationGet200ResponseData) SetRequestAccessEnabled(v bool)`

SetRequestAccessEnabled sets RequestAccessEnabled field to given value.


### GetAiSupportBotDisabled

`func (o *OrganizationGet200ResponseData) GetAiSupportBotDisabled() bool`

GetAiSupportBotDisabled returns the AiSupportBotDisabled field if non-nil, zero value otherwise.

### GetAiSupportBotDisabledOk

`func (o *OrganizationGet200ResponseData) GetAiSupportBotDisabledOk() (*bool, bool)`

GetAiSupportBotDisabledOk returns a tuple with the AiSupportBotDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSupportBotDisabled

`func (o *OrganizationGet200ResponseData) SetAiSupportBotDisabled(v bool)`

SetAiSupportBotDisabled sets AiSupportBotDisabled field to given value.


### GetRetoolFormsDisabled

`func (o *OrganizationGet200ResponseData) GetRetoolFormsDisabled() bool`

GetRetoolFormsDisabled returns the RetoolFormsDisabled field if non-nil, zero value otherwise.

### GetRetoolFormsDisabledOk

`func (o *OrganizationGet200ResponseData) GetRetoolFormsDisabledOk() (*bool, bool)`

GetRetoolFormsDisabledOk returns a tuple with the RetoolFormsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetoolFormsDisabled

`func (o *OrganizationGet200ResponseData) SetRetoolFormsDisabled(v bool)`

SetRetoolFormsDisabled sets RetoolFormsDisabled field to given value.


### GetReleaseManagementEnabled

`func (o *OrganizationGet200ResponseData) GetReleaseManagementEnabled() bool`

GetReleaseManagementEnabled returns the ReleaseManagementEnabled field if non-nil, zero value otherwise.

### GetReleaseManagementEnabledOk

`func (o *OrganizationGet200ResponseData) GetReleaseManagementEnabledOk() (*bool, bool)`

GetReleaseManagementEnabledOk returns a tuple with the ReleaseManagementEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseManagementEnabled

`func (o *OrganizationGet200ResponseData) SetReleaseManagementEnabled(v bool)`

SetReleaseManagementEnabled sets ReleaseManagementEnabled field to given value.


### GetCacheQueriesPerUser

`func (o *OrganizationGet200ResponseData) GetCacheQueriesPerUser() bool`

GetCacheQueriesPerUser returns the CacheQueriesPerUser field if non-nil, zero value otherwise.

### GetCacheQueriesPerUserOk

`func (o *OrganizationGet200ResponseData) GetCacheQueriesPerUserOk() (*bool, bool)`

GetCacheQueriesPerUserOk returns a tuple with the CacheQueriesPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheQueriesPerUser

`func (o *OrganizationGet200ResponseData) SetCacheQueriesPerUser(v bool)`

SetCacheQueriesPerUser sets CacheQueriesPerUser field to given value.


### SetCacheQueriesPerUserNil

`func (o *OrganizationGet200ResponseData) SetCacheQueriesPerUserNil(b bool)`

 SetCacheQueriesPerUserNil sets the value for CacheQueriesPerUser to be an explicit nil

### UnsetCacheQueriesPerUser
`func (o *OrganizationGet200ResponseData) UnsetCacheQueriesPerUser()`

UnsetCacheQueriesPerUser ensures that no value is present for CacheQueriesPerUser, not even an explicit nil
### GetApplyPreloadedCssToHomepage

`func (o *OrganizationGet200ResponseData) GetApplyPreloadedCssToHomepage() bool`

GetApplyPreloadedCssToHomepage returns the ApplyPreloadedCssToHomepage field if non-nil, zero value otherwise.

### GetApplyPreloadedCssToHomepageOk

`func (o *OrganizationGet200ResponseData) GetApplyPreloadedCssToHomepageOk() (*bool, bool)`

GetApplyPreloadedCssToHomepageOk returns a tuple with the ApplyPreloadedCssToHomepage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyPreloadedCssToHomepage

`func (o *OrganizationGet200ResponseData) SetApplyPreloadedCssToHomepage(v bool)`

SetApplyPreloadedCssToHomepage sets ApplyPreloadedCssToHomepage field to given value.


### GetPreloadedCss

`func (o *OrganizationGet200ResponseData) GetPreloadedCss() string`

GetPreloadedCss returns the PreloadedCss field if non-nil, zero value otherwise.

### GetPreloadedCssOk

`func (o *OrganizationGet200ResponseData) GetPreloadedCssOk() (*string, bool)`

GetPreloadedCssOk returns a tuple with the PreloadedCss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreloadedCss

`func (o *OrganizationGet200ResponseData) SetPreloadedCss(v string)`

SetPreloadedCss sets PreloadedCss field to given value.


### SetPreloadedCssNil

`func (o *OrganizationGet200ResponseData) SetPreloadedCssNil(b bool)`

 SetPreloadedCssNil sets the value for PreloadedCss to be an explicit nil

### UnsetPreloadedCss
`func (o *OrganizationGet200ResponseData) UnsetPreloadedCss()`

UnsetPreloadedCss ensures that no value is present for PreloadedCss, not even an explicit nil
### GetPreloadedJavascript

`func (o *OrganizationGet200ResponseData) GetPreloadedJavascript() string`

GetPreloadedJavascript returns the PreloadedJavascript field if non-nil, zero value otherwise.

### GetPreloadedJavascriptOk

`func (o *OrganizationGet200ResponseData) GetPreloadedJavascriptOk() (*string, bool)`

GetPreloadedJavascriptOk returns a tuple with the PreloadedJavascript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreloadedJavascript

`func (o *OrganizationGet200ResponseData) SetPreloadedJavascript(v string)`

SetPreloadedJavascript sets PreloadedJavascript field to given value.


### SetPreloadedJavascriptNil

`func (o *OrganizationGet200ResponseData) SetPreloadedJavascriptNil(b bool)`

 SetPreloadedJavascriptNil sets the value for PreloadedJavascript to be an explicit nil

### UnsetPreloadedJavascript
`func (o *OrganizationGet200ResponseData) UnsetPreloadedJavascript()`

UnsetPreloadedJavascript ensures that no value is present for PreloadedJavascript, not even an explicit nil
### GetJavascriptLinks

`func (o *OrganizationGet200ResponseData) GetJavascriptLinks() []string`

GetJavascriptLinks returns the JavascriptLinks field if non-nil, zero value otherwise.

### GetJavascriptLinksOk

`func (o *OrganizationGet200ResponseData) GetJavascriptLinksOk() (*[]string, bool)`

GetJavascriptLinksOk returns a tuple with the JavascriptLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavascriptLinks

`func (o *OrganizationGet200ResponseData) SetJavascriptLinks(v []string)`

SetJavascriptLinks sets JavascriptLinks field to given value.


### GetWorkflowRunRetentionPeriodMins

`func (o *OrganizationGet200ResponseData) GetWorkflowRunRetentionPeriodMins() float32`

GetWorkflowRunRetentionPeriodMins returns the WorkflowRunRetentionPeriodMins field if non-nil, zero value otherwise.

### GetWorkflowRunRetentionPeriodMinsOk

`func (o *OrganizationGet200ResponseData) GetWorkflowRunRetentionPeriodMinsOk() (*float32, bool)`

GetWorkflowRunRetentionPeriodMinsOk returns a tuple with the WorkflowRunRetentionPeriodMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRunRetentionPeriodMins

`func (o *OrganizationGet200ResponseData) SetWorkflowRunRetentionPeriodMins(v float32)`

SetWorkflowRunRetentionPeriodMins sets WorkflowRunRetentionPeriodMins field to given value.


### GetAppOwnersPermissionsManagement

`func (o *OrganizationGet200ResponseData) GetAppOwnersPermissionsManagement() bool`

GetAppOwnersPermissionsManagement returns the AppOwnersPermissionsManagement field if non-nil, zero value otherwise.

### GetAppOwnersPermissionsManagementOk

`func (o *OrganizationGet200ResponseData) GetAppOwnersPermissionsManagementOk() (*bool, bool)`

GetAppOwnersPermissionsManagementOk returns a tuple with the AppOwnersPermissionsManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppOwnersPermissionsManagement

`func (o *OrganizationGet200ResponseData) SetAppOwnersPermissionsManagement(v bool)`

SetAppOwnersPermissionsManagement sets AppOwnersPermissionsManagement field to given value.


### GetTwoFactorAuthRequired

`func (o *OrganizationGet200ResponseData) GetTwoFactorAuthRequired() bool`

GetTwoFactorAuthRequired returns the TwoFactorAuthRequired field if non-nil, zero value otherwise.

### GetTwoFactorAuthRequiredOk

`func (o *OrganizationGet200ResponseData) GetTwoFactorAuthRequiredOk() (*bool, bool)`

GetTwoFactorAuthRequiredOk returns a tuple with the TwoFactorAuthRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthRequired

`func (o *OrganizationGet200ResponseData) SetTwoFactorAuthRequired(v bool)`

SetTwoFactorAuthRequired sets TwoFactorAuthRequired field to given value.


### GetTwoFactorAuthType

`func (o *OrganizationGet200ResponseData) GetTwoFactorAuthType() string`

GetTwoFactorAuthType returns the TwoFactorAuthType field if non-nil, zero value otherwise.

### GetTwoFactorAuthTypeOk

`func (o *OrganizationGet200ResponseData) GetTwoFactorAuthTypeOk() (*string, bool)`

GetTwoFactorAuthTypeOk returns a tuple with the TwoFactorAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorAuthType

`func (o *OrganizationGet200ResponseData) SetTwoFactorAuthType(v string)`

SetTwoFactorAuthType sets TwoFactorAuthType field to given value.


### SetTwoFactorAuthTypeNil

`func (o *OrganizationGet200ResponseData) SetTwoFactorAuthTypeNil(b bool)`

 SetTwoFactorAuthTypeNil sets the value for TwoFactorAuthType to be an explicit nil

### UnsetTwoFactorAuthType
`func (o *OrganizationGet200ResponseData) UnsetTwoFactorAuthType()`

UnsetTwoFactorAuthType ensures that no value is present for TwoFactorAuthType, not even an explicit nil
### GetDisableNewLoginIpNotificationEmail

`func (o *OrganizationGet200ResponseData) GetDisableNewLoginIpNotificationEmail() bool`

GetDisableNewLoginIpNotificationEmail returns the DisableNewLoginIpNotificationEmail field if non-nil, zero value otherwise.

### GetDisableNewLoginIpNotificationEmailOk

`func (o *OrganizationGet200ResponseData) GetDisableNewLoginIpNotificationEmailOk() (*bool, bool)`

GetDisableNewLoginIpNotificationEmailOk returns a tuple with the DisableNewLoginIpNotificationEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNewLoginIpNotificationEmail

`func (o *OrganizationGet200ResponseData) SetDisableNewLoginIpNotificationEmail(v bool)`

SetDisableNewLoginIpNotificationEmail sets DisableNewLoginIpNotificationEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


