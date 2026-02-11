# SourceControlDeploymentsGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployments** | [**[]SourceControlDeploymentsGet200ResponseDataDeploymentsInner**](SourceControlDeploymentsGet200ResponseDataDeploymentsInner.md) | List of deployments, ordered by creation time (most recent first) | 

## Methods

### NewSourceControlDeploymentsGet200ResponseData

`func NewSourceControlDeploymentsGet200ResponseData(deployments []SourceControlDeploymentsGet200ResponseDataDeploymentsInner, ) *SourceControlDeploymentsGet200ResponseData`

NewSourceControlDeploymentsGet200ResponseData instantiates a new SourceControlDeploymentsGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceControlDeploymentsGet200ResponseDataWithDefaults

`func NewSourceControlDeploymentsGet200ResponseDataWithDefaults() *SourceControlDeploymentsGet200ResponseData`

NewSourceControlDeploymentsGet200ResponseDataWithDefaults instantiates a new SourceControlDeploymentsGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployments

`func (o *SourceControlDeploymentsGet200ResponseData) GetDeployments() []SourceControlDeploymentsGet200ResponseDataDeploymentsInner`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *SourceControlDeploymentsGet200ResponseData) GetDeploymentsOk() (*[]SourceControlDeploymentsGet200ResponseDataDeploymentsInner, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *SourceControlDeploymentsGet200ResponseData) SetDeployments(v []SourceControlDeploymentsGet200ResponseDataDeploymentsInner)`

SetDeployments sets Deployments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


