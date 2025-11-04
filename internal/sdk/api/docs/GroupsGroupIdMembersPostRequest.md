# GroupsGroupIdMembersPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]GroupsGroupIdPutRequestMembersInner**](GroupsGroupIdPutRequestMembersInner.md) | Users to add to the group. | 

## Methods

### NewGroupsGroupIdMembersPostRequest

`func NewGroupsGroupIdMembersPostRequest(members []GroupsGroupIdPutRequestMembersInner, ) *GroupsGroupIdMembersPostRequest`

NewGroupsGroupIdMembersPostRequest instantiates a new GroupsGroupIdMembersPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupsGroupIdMembersPostRequestWithDefaults

`func NewGroupsGroupIdMembersPostRequestWithDefaults() *GroupsGroupIdMembersPostRequest`

NewGroupsGroupIdMembersPostRequestWithDefaults instantiates a new GroupsGroupIdMembersPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *GroupsGroupIdMembersPostRequest) GetMembers() []GroupsGroupIdPutRequestMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GroupsGroupIdMembersPostRequest) GetMembersOk() (*[]GroupsGroupIdPutRequestMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GroupsGroupIdMembersPostRequest) SetMembers(v []GroupsGroupIdPutRequestMembersInner)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


