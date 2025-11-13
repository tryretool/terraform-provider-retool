# UsageAnalyticsOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | The id of the organization | 
**Host** | **string** | The host of the organization | 
**LastActive** | **time.Time** |  | 

## Methods

### NewUsageAnalyticsOrganization

`func NewUsageAnalyticsOrganization(orgId string, host string, lastActive time.Time, ) *UsageAnalyticsOrganization`

NewUsageAnalyticsOrganization instantiates a new UsageAnalyticsOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageAnalyticsOrganizationWithDefaults

`func NewUsageAnalyticsOrganizationWithDefaults() *UsageAnalyticsOrganization`

NewUsageAnalyticsOrganizationWithDefaults instantiates a new UsageAnalyticsOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *UsageAnalyticsOrganization) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *UsageAnalyticsOrganization) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *UsageAnalyticsOrganization) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetHost

`func (o *UsageAnalyticsOrganization) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UsageAnalyticsOrganization) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UsageAnalyticsOrganization) SetHost(v string)`

SetHost sets Host field to given value.


### GetLastActive

`func (o *UsageAnalyticsOrganization) GetLastActive() time.Time`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *UsageAnalyticsOrganization) GetLastActiveOk() (*time.Time, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *UsageAnalyticsOrganization) SetLastActive(v time.Time)`

SetLastActive sets LastActive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


