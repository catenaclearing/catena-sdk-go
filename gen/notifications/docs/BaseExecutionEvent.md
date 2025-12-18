# BaseExecutionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the execution | 
**CreatedAt** | **time.Time** | The date and time the execution was created | 
**UpdatedAt** | **time.Time** | The date and time the execution was last updated | 
**ScheduleId** | **string** | The ID of the schedule that triggered this execution | 
**ConnectionId** | **string** | The ID of the connection used for this execution | 
**FleetId** | **string** | The Catena ID of the fleet. | 
**Status** | [**StatusEnum**](StatusEnum.md) | The status of the execution | 
**SourceName** | [**TspEnum**](TspEnum.md) | The TSP that is the source for this execution | 
**Resource** | [**ResourceEnum**](ResourceEnum.md) | The resource type being processed in this execution | 

## Methods

### NewBaseExecutionEvent

`func NewBaseExecutionEvent(id string, createdAt time.Time, updatedAt time.Time, scheduleId string, connectionId string, fleetId string, status StatusEnum, sourceName TspEnum, resource ResourceEnum, ) *BaseExecutionEvent`

NewBaseExecutionEvent instantiates a new BaseExecutionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseExecutionEventWithDefaults

`func NewBaseExecutionEventWithDefaults() *BaseExecutionEvent`

NewBaseExecutionEventWithDefaults instantiates a new BaseExecutionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseExecutionEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseExecutionEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseExecutionEvent) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *BaseExecutionEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseExecutionEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseExecutionEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseExecutionEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseExecutionEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseExecutionEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetScheduleId

`func (o *BaseExecutionEvent) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BaseExecutionEvent) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BaseExecutionEvent) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.


### GetConnectionId

`func (o *BaseExecutionEvent) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseExecutionEvent) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseExecutionEvent) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetFleetId

`func (o *BaseExecutionEvent) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseExecutionEvent) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseExecutionEvent) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetStatus

`func (o *BaseExecutionEvent) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseExecutionEvent) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseExecutionEvent) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetSourceName

`func (o *BaseExecutionEvent) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseExecutionEvent) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseExecutionEvent) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetResource

`func (o *BaseExecutionEvent) GetResource() ResourceEnum`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BaseExecutionEvent) GetResourceOk() (*ResourceEnum, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BaseExecutionEvent) SetResource(v ResourceEnum)`

SetResource sets Resource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


