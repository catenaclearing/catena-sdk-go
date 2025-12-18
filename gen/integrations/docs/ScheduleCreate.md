# ScheduleCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | [**ResourceEnum**](ResourceEnum.md) | The type of resource to fetch on this schedule (e.g., VEHICLE, DRIVER, HOS, IFTA). | 
**ExecutionIntervalSeconds** | Pointer to **int32** | The interval between scheduled executions in seconds. Defaults to 10 minutes (600 seconds). | [optional] [default to 600]
**ConsecutiveErrorThreshold** | Pointer to **int32** | The number of consecutive errors allowed before the schedule is automatically set to INACTIVE. Defaults to 15. | [optional] [default to 15]
**MaxConcurrentExecutions** | Pointer to **int32** | The maximum number of concurrent executions allowed for this schedule. Defaults to 1. | [optional] [default to 1]
**Status** | Pointer to [**StatusEnum**](StatusEnum.md) | The initial status of the schedule. Defaults to ACTIVE. | [optional] 

## Methods

### NewScheduleCreate

`func NewScheduleCreate(resource ResourceEnum, ) *ScheduleCreate`

NewScheduleCreate instantiates a new ScheduleCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleCreateWithDefaults

`func NewScheduleCreateWithDefaults() *ScheduleCreate`

NewScheduleCreateWithDefaults instantiates a new ScheduleCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *ScheduleCreate) GetResource() ResourceEnum`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ScheduleCreate) GetResourceOk() (*ResourceEnum, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ScheduleCreate) SetResource(v ResourceEnum)`

SetResource sets Resource field to given value.


### GetExecutionIntervalSeconds

`func (o *ScheduleCreate) GetExecutionIntervalSeconds() int32`

GetExecutionIntervalSeconds returns the ExecutionIntervalSeconds field if non-nil, zero value otherwise.

### GetExecutionIntervalSecondsOk

`func (o *ScheduleCreate) GetExecutionIntervalSecondsOk() (*int32, bool)`

GetExecutionIntervalSecondsOk returns a tuple with the ExecutionIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionIntervalSeconds

`func (o *ScheduleCreate) SetExecutionIntervalSeconds(v int32)`

SetExecutionIntervalSeconds sets ExecutionIntervalSeconds field to given value.

### HasExecutionIntervalSeconds

`func (o *ScheduleCreate) HasExecutionIntervalSeconds() bool`

HasExecutionIntervalSeconds returns a boolean if a field has been set.

### GetConsecutiveErrorThreshold

`func (o *ScheduleCreate) GetConsecutiveErrorThreshold() int32`

GetConsecutiveErrorThreshold returns the ConsecutiveErrorThreshold field if non-nil, zero value otherwise.

### GetConsecutiveErrorThresholdOk

`func (o *ScheduleCreate) GetConsecutiveErrorThresholdOk() (*int32, bool)`

GetConsecutiveErrorThresholdOk returns a tuple with the ConsecutiveErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveErrorThreshold

`func (o *ScheduleCreate) SetConsecutiveErrorThreshold(v int32)`

SetConsecutiveErrorThreshold sets ConsecutiveErrorThreshold field to given value.

### HasConsecutiveErrorThreshold

`func (o *ScheduleCreate) HasConsecutiveErrorThreshold() bool`

HasConsecutiveErrorThreshold returns a boolean if a field has been set.

### GetMaxConcurrentExecutions

`func (o *ScheduleCreate) GetMaxConcurrentExecutions() int32`

GetMaxConcurrentExecutions returns the MaxConcurrentExecutions field if non-nil, zero value otherwise.

### GetMaxConcurrentExecutionsOk

`func (o *ScheduleCreate) GetMaxConcurrentExecutionsOk() (*int32, bool)`

GetMaxConcurrentExecutionsOk returns a tuple with the MaxConcurrentExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentExecutions

`func (o *ScheduleCreate) SetMaxConcurrentExecutions(v int32)`

SetMaxConcurrentExecutions sets MaxConcurrentExecutions field to given value.

### HasMaxConcurrentExecutions

`func (o *ScheduleCreate) HasMaxConcurrentExecutions() bool`

HasMaxConcurrentExecutions returns a boolean if a field has been set.

### GetStatus

`func (o *ScheduleCreate) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduleCreate) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduleCreate) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScheduleCreate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


