# ScheduleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**NullableScheduleStatusEnum**](ScheduleStatusEnum.md) |  | [optional] 
**ExecutionIntervalSeconds** | Pointer to **NullableInt32** |  | [optional] 
**ConsecutiveErrorThreshold** | Pointer to **NullableInt32** |  | [optional] 
**ConsecutiveErrorCount** | Pointer to **NullableInt32** |  | [optional] 
**MaxConcurrentExecutions** | Pointer to **NullableInt32** |  | [optional] 
**NextExecutionAt** | Pointer to **NullableTime** |  | [optional] 
**Cursor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewScheduleUpdate

`func NewScheduleUpdate() *ScheduleUpdate`

NewScheduleUpdate instantiates a new ScheduleUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleUpdateWithDefaults

`func NewScheduleUpdateWithDefaults() *ScheduleUpdate`

NewScheduleUpdateWithDefaults instantiates a new ScheduleUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ScheduleUpdate) GetStatus() ScheduleStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ScheduleUpdate) GetStatusOk() (*ScheduleStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ScheduleUpdate) SetStatus(v ScheduleStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ScheduleUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ScheduleUpdate) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ScheduleUpdate) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetExecutionIntervalSeconds

`func (o *ScheduleUpdate) GetExecutionIntervalSeconds() int32`

GetExecutionIntervalSeconds returns the ExecutionIntervalSeconds field if non-nil, zero value otherwise.

### GetExecutionIntervalSecondsOk

`func (o *ScheduleUpdate) GetExecutionIntervalSecondsOk() (*int32, bool)`

GetExecutionIntervalSecondsOk returns a tuple with the ExecutionIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionIntervalSeconds

`func (o *ScheduleUpdate) SetExecutionIntervalSeconds(v int32)`

SetExecutionIntervalSeconds sets ExecutionIntervalSeconds field to given value.

### HasExecutionIntervalSeconds

`func (o *ScheduleUpdate) HasExecutionIntervalSeconds() bool`

HasExecutionIntervalSeconds returns a boolean if a field has been set.

### SetExecutionIntervalSecondsNil

`func (o *ScheduleUpdate) SetExecutionIntervalSecondsNil(b bool)`

 SetExecutionIntervalSecondsNil sets the value for ExecutionIntervalSeconds to be an explicit nil

### UnsetExecutionIntervalSeconds
`func (o *ScheduleUpdate) UnsetExecutionIntervalSeconds()`

UnsetExecutionIntervalSeconds ensures that no value is present for ExecutionIntervalSeconds, not even an explicit nil
### GetConsecutiveErrorThreshold

`func (o *ScheduleUpdate) GetConsecutiveErrorThreshold() int32`

GetConsecutiveErrorThreshold returns the ConsecutiveErrorThreshold field if non-nil, zero value otherwise.

### GetConsecutiveErrorThresholdOk

`func (o *ScheduleUpdate) GetConsecutiveErrorThresholdOk() (*int32, bool)`

GetConsecutiveErrorThresholdOk returns a tuple with the ConsecutiveErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveErrorThreshold

`func (o *ScheduleUpdate) SetConsecutiveErrorThreshold(v int32)`

SetConsecutiveErrorThreshold sets ConsecutiveErrorThreshold field to given value.

### HasConsecutiveErrorThreshold

`func (o *ScheduleUpdate) HasConsecutiveErrorThreshold() bool`

HasConsecutiveErrorThreshold returns a boolean if a field has been set.

### SetConsecutiveErrorThresholdNil

`func (o *ScheduleUpdate) SetConsecutiveErrorThresholdNil(b bool)`

 SetConsecutiveErrorThresholdNil sets the value for ConsecutiveErrorThreshold to be an explicit nil

### UnsetConsecutiveErrorThreshold
`func (o *ScheduleUpdate) UnsetConsecutiveErrorThreshold()`

UnsetConsecutiveErrorThreshold ensures that no value is present for ConsecutiveErrorThreshold, not even an explicit nil
### GetConsecutiveErrorCount

`func (o *ScheduleUpdate) GetConsecutiveErrorCount() int32`

GetConsecutiveErrorCount returns the ConsecutiveErrorCount field if non-nil, zero value otherwise.

### GetConsecutiveErrorCountOk

`func (o *ScheduleUpdate) GetConsecutiveErrorCountOk() (*int32, bool)`

GetConsecutiveErrorCountOk returns a tuple with the ConsecutiveErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveErrorCount

`func (o *ScheduleUpdate) SetConsecutiveErrorCount(v int32)`

SetConsecutiveErrorCount sets ConsecutiveErrorCount field to given value.

### HasConsecutiveErrorCount

`func (o *ScheduleUpdate) HasConsecutiveErrorCount() bool`

HasConsecutiveErrorCount returns a boolean if a field has been set.

### SetConsecutiveErrorCountNil

`func (o *ScheduleUpdate) SetConsecutiveErrorCountNil(b bool)`

 SetConsecutiveErrorCountNil sets the value for ConsecutiveErrorCount to be an explicit nil

### UnsetConsecutiveErrorCount
`func (o *ScheduleUpdate) UnsetConsecutiveErrorCount()`

UnsetConsecutiveErrorCount ensures that no value is present for ConsecutiveErrorCount, not even an explicit nil
### GetMaxConcurrentExecutions

`func (o *ScheduleUpdate) GetMaxConcurrentExecutions() int32`

GetMaxConcurrentExecutions returns the MaxConcurrentExecutions field if non-nil, zero value otherwise.

### GetMaxConcurrentExecutionsOk

`func (o *ScheduleUpdate) GetMaxConcurrentExecutionsOk() (*int32, bool)`

GetMaxConcurrentExecutionsOk returns a tuple with the MaxConcurrentExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentExecutions

`func (o *ScheduleUpdate) SetMaxConcurrentExecutions(v int32)`

SetMaxConcurrentExecutions sets MaxConcurrentExecutions field to given value.

### HasMaxConcurrentExecutions

`func (o *ScheduleUpdate) HasMaxConcurrentExecutions() bool`

HasMaxConcurrentExecutions returns a boolean if a field has been set.

### SetMaxConcurrentExecutionsNil

`func (o *ScheduleUpdate) SetMaxConcurrentExecutionsNil(b bool)`

 SetMaxConcurrentExecutionsNil sets the value for MaxConcurrentExecutions to be an explicit nil

### UnsetMaxConcurrentExecutions
`func (o *ScheduleUpdate) UnsetMaxConcurrentExecutions()`

UnsetMaxConcurrentExecutions ensures that no value is present for MaxConcurrentExecutions, not even an explicit nil
### GetNextExecutionAt

`func (o *ScheduleUpdate) GetNextExecutionAt() time.Time`

GetNextExecutionAt returns the NextExecutionAt field if non-nil, zero value otherwise.

### GetNextExecutionAtOk

`func (o *ScheduleUpdate) GetNextExecutionAtOk() (*time.Time, bool)`

GetNextExecutionAtOk returns a tuple with the NextExecutionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecutionAt

`func (o *ScheduleUpdate) SetNextExecutionAt(v time.Time)`

SetNextExecutionAt sets NextExecutionAt field to given value.

### HasNextExecutionAt

`func (o *ScheduleUpdate) HasNextExecutionAt() bool`

HasNextExecutionAt returns a boolean if a field has been set.

### SetNextExecutionAtNil

`func (o *ScheduleUpdate) SetNextExecutionAtNil(b bool)`

 SetNextExecutionAtNil sets the value for NextExecutionAt to be an explicit nil

### UnsetNextExecutionAt
`func (o *ScheduleUpdate) UnsetNextExecutionAt()`

UnsetNextExecutionAt ensures that no value is present for NextExecutionAt, not even an explicit nil
### GetCursor

`func (o *ScheduleUpdate) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ScheduleUpdate) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ScheduleUpdate) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ScheduleUpdate) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### SetCursorNil

`func (o *ScheduleUpdate) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *ScheduleUpdate) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


