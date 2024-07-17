# UsageOrganizationsGet200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | The id of the organization | 
**Host** | **string** | The host of the organization | 
**LastActive** | **time.Time** |  | 

## Methods

### NewUsageOrganizationsGet200ResponseDataInner

`func NewUsageOrganizationsGet200ResponseDataInner(orgId string, host string, lastActive time.Time, ) *UsageOrganizationsGet200ResponseDataInner`

NewUsageOrganizationsGet200ResponseDataInner instantiates a new UsageOrganizationsGet200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageOrganizationsGet200ResponseDataInnerWithDefaults

`func NewUsageOrganizationsGet200ResponseDataInnerWithDefaults() *UsageOrganizationsGet200ResponseDataInner`

NewUsageOrganizationsGet200ResponseDataInnerWithDefaults instantiates a new UsageOrganizationsGet200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *UsageOrganizationsGet200ResponseDataInner) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *UsageOrganizationsGet200ResponseDataInner) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *UsageOrganizationsGet200ResponseDataInner) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetHost

`func (o *UsageOrganizationsGet200ResponseDataInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *UsageOrganizationsGet200ResponseDataInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *UsageOrganizationsGet200ResponseDataInner) SetHost(v string)`

SetHost sets Host field to given value.


### GetLastActive

`func (o *UsageOrganizationsGet200ResponseDataInner) GetLastActive() time.Time`

GetLastActive returns the LastActive field if non-nil, zero value otherwise.

### GetLastActiveOk

`func (o *UsageOrganizationsGet200ResponseDataInner) GetLastActiveOk() (*time.Time, bool)`

GetLastActiveOk returns a tuple with the LastActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActive

`func (o *UsageOrganizationsGet200ResponseDataInner) SetLastActive(v time.Time)`

SetLastActive sets LastActive field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


