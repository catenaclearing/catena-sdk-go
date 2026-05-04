# ScheduleRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the schedule. | 
**ConnectionId** | **string** | The ID of the connection that this schedule belongs to. | 
**Resource** | [**ResourceEnum**](ResourceEnum.md) | The type of resource being fetched on this schedule (e.g., VEHICLE, DRIVER, HOS, IFTA). | 
**ExecutionIntervalSeconds** | **int32** | The interval between scheduled executions in seconds. | 
**Status** | [**ScheduleStatusEnum**](ScheduleStatusEnum.md) | The current status of the schedule (ACTIVE, INACTIVE). | 
**ConsecutiveErrorCount** | **int32** | The number of consecutive errors that have occurred for this schedule. | 
**ConsecutiveErrorThreshold** | **int32** | The number of consecutive errors allowed before the schedule is automatically set to INACTIVE. | 
**MaxConcurrentExecutions** | **int32** | The maximum number of concurrent executions allowed for this schedule. | 
**Cursor** | Pointer to **NullableString** |  | [optional] 
**LastDataAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewScheduleRead

`func NewScheduleRead(id string, connectionId string, resource ResourceEnum, executionIntervalSeconds int32, status ScheduleStatusEnum, consecutiveErrorCount int32, consecutiveErrorThreshold int32, maxConcurrentExecutions int32, ) *ScheduleRead`

NewScheduleRead instantiates a new ScheduleRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleReadWithDefaults

`func NewScheduleReadWithDefaults() *ScheduleRead`

NewScheduleReadWithDefaults instantiates a new ScheduleRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleRead) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionId

`func (o *ScheduleRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ScheduleRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ScheduleRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetResource

`func (o *ScheduleRead) GetResource() ResourceEnum`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ScheduleRead) GetResourceOk() (*ResourceEnum, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ScheduleRead) SetResource(v ResourceEnum)`

SetResource sets Resource field to given value.


### GetExecutionIntervalSeconds

`func (o *ScheduleRead) GetExecutionIntervalSeconds() int32`

GetExecutionIntervalSeconds returns the ExecutionIntervalSeconds field if non-nil, zero value otherwise.

### GetExecutionIntervalSecondsOk

`func (o *ScheduleRead) GetExecutionIntervalSecondsOk() (*int32, bool)`

GetExecutionIntervalSecondsOk returns a tuple with the ExecutionIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionIntervalSeconds

`func (o *ScheduleRead) SetExecutionIntervalSeconds(v int32)`

SetExecutionIntervalSeconds sets ExecutionIntervalSeconds field to given value.


### GetStatus

`func (o *ScheduleRead) GetStatus() ScheduleStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduleRead) GetStatusOk() (*ScheduleStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduleRead) SetStatus(v ScheduleStatusEnum)`

SetStatus sets Status field to given value.


### GetConsecutiveErrorCount

`func (o *ScheduleRead) GetConsecutiveErrorCount() int32`

GetConsecutiveErrorCount returns the ConsecutiveErrorCount field if non-nil, zero value otherwise.

### GetConsecutiveErrorCountOk

`func (o *ScheduleRead) GetConsecutiveErrorCountOk() (*int32, bool)`

GetConsecutiveErrorCountOk returns a tuple with the ConsecutiveErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveErrorCount

`func (o *ScheduleRead) SetConsecutiveErrorCount(v int32)`

SetConsecutiveErrorCount sets ConsecutiveErrorCount field to given value.


### GetConsecutiveErrorThreshold

`func (o *ScheduleRead) GetConsecutiveErrorThreshold() int32`

GetConsecutiveErrorThreshold returns the ConsecutiveErrorThreshold field if non-nil, zero value otherwise.

### GetConsecutiveErrorThresholdOk

`func (o *ScheduleRead) GetConsecutiveErrorThresholdOk() (*int32, bool)`

GetConsecutiveErrorThresholdOk returns a tuple with the ConsecutiveErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveErrorThreshold

`func (o *ScheduleRead) SetConsecutiveErrorThreshold(v int32)`

SetConsecutiveErrorThreshold sets ConsecutiveErrorThreshold field to given value.


### GetMaxConcurrentExecutions

`func (o *ScheduleRead) GetMaxConcurrentExecutions() int32`

GetMaxConcurrentExecutions returns the MaxConcurrentExecutions field if non-nil, zero value otherwise.

### GetMaxConcurrentExecutionsOk

`func (o *ScheduleRead) GetMaxConcurrentExecutionsOk() (*int32, bool)`

GetMaxConcurrentExecutionsOk returns a tuple with the MaxConcurrentExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentExecutions

`func (o *ScheduleRead) SetMaxConcurrentExecutions(v int32)`

SetMaxConcurrentExecutions sets MaxConcurrentExecutions field to given value.


### GetCursor

`func (o *ScheduleRead) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ScheduleRead) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ScheduleRead) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ScheduleRead) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### SetCursorNil

`func (o *ScheduleRead) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *ScheduleRead) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
### GetLastDataAt

`func (o *ScheduleRead) GetLastDataAt() time.Time`

GetLastDataAt returns the LastDataAt field if non-nil, zero value otherwise.

### GetLastDataAtOk

`func (o *ScheduleRead) GetLastDataAtOk() (*time.Time, bool)`

GetLastDataAtOk returns a tuple with the LastDataAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDataAt

`func (o *ScheduleRead) SetLastDataAt(v time.Time)`

SetLastDataAt sets LastDataAt field to given value.

### HasLastDataAt

`func (o *ScheduleRead) HasLastDataAt() bool`

HasLastDataAt returns a boolean if a field has been set.

### SetLastDataAtNil

`func (o *ScheduleRead) SetLastDataAtNil(b bool)`

 SetLastDataAtNil sets the value for LastDataAt to be an explicit nil

### UnsetLastDataAt
`func (o *ScheduleRead) UnsetLastDataAt()`

UnsetLastDataAt ensures that no value is present for LastDataAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


