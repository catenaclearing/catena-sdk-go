# BaseScheduleEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the schedule | 
**ConnectionId** | **string** | The ID of the connection used for this schedule | 
**Resource** | [**ResourceEnum**](ResourceEnum.md) | The resource type being scheduled for processing | 
**ExecutionIntervalSeconds** | **int32** | The interval in seconds between executions | 
**Status** | [**StatusEnum**](StatusEnum.md) | The status of the schedule | 
**ConsecutiveErrorThreshold** | **int32** | The threshold for consecutive errors before marking as failed | 
**MaxConcurrentExecutions** | **int32** | The maximum number of concurrent executions allowed | 

## Methods

### NewBaseScheduleEvent

`func NewBaseScheduleEvent(id string, connectionId string, resource ResourceEnum, executionIntervalSeconds int32, status StatusEnum, consecutiveErrorThreshold int32, maxConcurrentExecutions int32, ) *BaseScheduleEvent`

NewBaseScheduleEvent instantiates a new BaseScheduleEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseScheduleEventWithDefaults

`func NewBaseScheduleEventWithDefaults() *BaseScheduleEvent`

NewBaseScheduleEventWithDefaults instantiates a new BaseScheduleEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseScheduleEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseScheduleEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseScheduleEvent) SetId(v string)`

SetId sets Id field to given value.


### GetConnectionId

`func (o *BaseScheduleEvent) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseScheduleEvent) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseScheduleEvent) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetResource

`func (o *BaseScheduleEvent) GetResource() ResourceEnum`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BaseScheduleEvent) GetResourceOk() (*ResourceEnum, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BaseScheduleEvent) SetResource(v ResourceEnum)`

SetResource sets Resource field to given value.


### GetExecutionIntervalSeconds

`func (o *BaseScheduleEvent) GetExecutionIntervalSeconds() int32`

GetExecutionIntervalSeconds returns the ExecutionIntervalSeconds field if non-nil, zero value otherwise.

### GetExecutionIntervalSecondsOk

`func (o *BaseScheduleEvent) GetExecutionIntervalSecondsOk() (*int32, bool)`

GetExecutionIntervalSecondsOk returns a tuple with the ExecutionIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionIntervalSeconds

`func (o *BaseScheduleEvent) SetExecutionIntervalSeconds(v int32)`

SetExecutionIntervalSeconds sets ExecutionIntervalSeconds field to given value.


### GetStatus

`func (o *BaseScheduleEvent) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseScheduleEvent) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseScheduleEvent) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetConsecutiveErrorThreshold

`func (o *BaseScheduleEvent) GetConsecutiveErrorThreshold() int32`

GetConsecutiveErrorThreshold returns the ConsecutiveErrorThreshold field if non-nil, zero value otherwise.

### GetConsecutiveErrorThresholdOk

`func (o *BaseScheduleEvent) GetConsecutiveErrorThresholdOk() (*int32, bool)`

GetConsecutiveErrorThresholdOk returns a tuple with the ConsecutiveErrorThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveErrorThreshold

`func (o *BaseScheduleEvent) SetConsecutiveErrorThreshold(v int32)`

SetConsecutiveErrorThreshold sets ConsecutiveErrorThreshold field to given value.


### GetMaxConcurrentExecutions

`func (o *BaseScheduleEvent) GetMaxConcurrentExecutions() int32`

GetMaxConcurrentExecutions returns the MaxConcurrentExecutions field if non-nil, zero value otherwise.

### GetMaxConcurrentExecutionsOk

`func (o *BaseScheduleEvent) GetMaxConcurrentExecutionsOk() (*int32, bool)`

GetMaxConcurrentExecutionsOk returns a tuple with the MaxConcurrentExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentExecutions

`func (o *BaseScheduleEvent) SetMaxConcurrentExecutions(v int32)`

SetMaxConcurrentExecutions sets MaxConcurrentExecutions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


