# OrganizationCodeExecutorSettingsGet200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemoryLimitMbs** | **NullableInt32** | Memory limit in MB for Code Executor processes. null uses the service default (WORKFLOW_MEMORY_LIMIT_MBS). | 
**MaxConcurrentProcesses** | **NullableInt32** | Max concurrent Code Executor processes for this org. null uses the service deployment default. | 
**CircuitBreakerEnabled** | **bool** | Whether the utilization-based circuit breaker is enforced for this org. This overrides the service default (UTILIZATION_CIRCUIT_BREAKER_ENABLED). | 

## Methods

### NewOrganizationCodeExecutorSettingsGet200ResponseData

`func NewOrganizationCodeExecutorSettingsGet200ResponseData(memoryLimitMbs NullableInt32, maxConcurrentProcesses NullableInt32, circuitBreakerEnabled bool, ) *OrganizationCodeExecutorSettingsGet200ResponseData`

NewOrganizationCodeExecutorSettingsGet200ResponseData instantiates a new OrganizationCodeExecutorSettingsGet200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationCodeExecutorSettingsGet200ResponseDataWithDefaults

`func NewOrganizationCodeExecutorSettingsGet200ResponseDataWithDefaults() *OrganizationCodeExecutorSettingsGet200ResponseData`

NewOrganizationCodeExecutorSettingsGet200ResponseDataWithDefaults instantiates a new OrganizationCodeExecutorSettingsGet200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemoryLimitMbs

`func (o *OrganizationCodeExecutorSettingsGet200ResponseData) GetMemoryLimitMbs() int32`

GetMemoryLimitMbs returns the MemoryLimitMbs field if non-nil, zero value otherwise.

### GetMemoryLimitMbsOk

`func (o *OrganizationCodeExecutorSettingsGet200ResponseData) GetMemoryLimitMbsOk() (*int32, bool)`

GetMemoryLimitMbsOk returns a tuple with the MemoryLimitMbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimitMbs

`func (o *OrganizationCodeExecutorSettingsGet200ResponseData) SetMemoryLimitMbs(v int32)`

SetMemoryLimitMbs sets MemoryLimitMbs field to given value.


### SetMemoryLimitMbsNil

`func (o *OrganizationCodeExecutorSettingsGet200ResponseData) SetMemoryLimitMbsNil(b bool)`

 SetMemoryLimitMbsNil sets the value for MemoryLimitMbs to be an explicit nil

### UnsetMemoryLimitMbs
`func (o *OrganizationCodeExecutorSettingsGet200ResponseData) UnsetMemoryLimitMbs()`

UnsetMemoryLimitMbs ensures that no value is present for MemoryLimitMbs, not even an explicit nil
### GetMaxConcurrentProcesses

`func (o *OrganizationCodeExecutorSettingsGet200ResponseData) GetMaxConcurrentProcesses() int32`

GetMaxConcurrentProcesses returns the MaxConcurrentProcesses field if non-nil, zero value otherwise.

### GetMaxConcurrentProcessesOk

`func (o *OrganizationCodeExecutorSettingsGet200ResponseData) GetMaxConcurrentProcessesOk() (*int32, bool)`

GetMaxConcurrentProcessesOk returns a tuple with the MaxConcurrentProcesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentProcesses

`func (o *OrganizationCodeExecutorSettingsGet200ResponseData) SetMaxConcurrentProcesses(v int32)`

SetMaxConcurrentProcesses sets MaxConcurrentProcesses field to given value.


### SetMaxConcurrentProcessesNil

`func (o *OrganizationCodeExecutorSettingsGet200ResponseData) SetMaxConcurrentProcessesNil(b bool)`

 SetMaxConcurrentProcessesNil sets the value for MaxConcurrentProcesses to be an explicit nil

### UnsetMaxConcurrentProcesses
`func (o *OrganizationCodeExecutorSettingsGet200ResponseData) UnsetMaxConcurrentProcesses()`

UnsetMaxConcurrentProcesses ensures that no value is present for MaxConcurrentProcesses, not even an explicit nil
### GetCircuitBreakerEnabled

`func (o *OrganizationCodeExecutorSettingsGet200ResponseData) GetCircuitBreakerEnabled() bool`

GetCircuitBreakerEnabled returns the CircuitBreakerEnabled field if non-nil, zero value otherwise.

### GetCircuitBreakerEnabledOk

`func (o *OrganizationCodeExecutorSettingsGet200ResponseData) GetCircuitBreakerEnabledOk() (*bool, bool)`

GetCircuitBreakerEnabledOk returns a tuple with the CircuitBreakerEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitBreakerEnabled

`func (o *OrganizationCodeExecutorSettingsGet200ResponseData) SetCircuitBreakerEnabled(v bool)`

SetCircuitBreakerEnabled sets CircuitBreakerEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


