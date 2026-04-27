# ExecutionRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the execution. Use this ID to trace which execution ingested specific data. | 
**CreatedAt** | **time.Time** | Timestamp when the execution was created. | 
**UpdatedAt** | **time.Time** | Timestamp when the execution was last updated. | 
**ScheduleId** | **string** | The ID of the schedule that this execution belongs to. | 
**ConnectionId** | **string** | The ID of the connection used for this execution. | 
**FleetId** | **string** | The ID of the fleet that owns this execution. | 
**Status** | [**ExecutionStatusEnum**](ExecutionStatusEnum.md) | The current status of the execution. | 
**SourceName** | [**TspEnum**](TspEnum.md) | The name of the TSP integration used for this execution. | 
**Resource** | [**ResourceEnum**](ResourceEnum.md) | The type of resource being fetched (e.g., VEHICLE, DRIVER, HOS, IFTA). | 
**Cursor** | Pointer to **NullableString** |  | [optional] 
**Response** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewExecutionRead

`func NewExecutionRead(id string, createdAt time.Time, updatedAt time.Time, scheduleId string, connectionId string, fleetId string, status ExecutionStatusEnum, sourceName TspEnum, resource ResourceEnum, ) *ExecutionRead`

NewExecutionRead instantiates a new ExecutionRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionReadWithDefaults

`func NewExecutionReadWithDefaults() *ExecutionRead`

NewExecutionReadWithDefaults instantiates a new ExecutionRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExecutionRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExecutionRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExecutionRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ExecutionRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExecutionRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExecutionRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ExecutionRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExecutionRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExecutionRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetScheduleId

`func (o *ExecutionRead) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *ExecutionRead) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *ExecutionRead) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.


### GetConnectionId

`func (o *ExecutionRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ExecutionRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ExecutionRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetFleetId

`func (o *ExecutionRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *ExecutionRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *ExecutionRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetStatus

`func (o *ExecutionRead) GetStatus() ExecutionStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionRead) GetStatusOk() (*ExecutionStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionRead) SetStatus(v ExecutionStatusEnum)`

SetStatus sets Status field to given value.


### GetSourceName

`func (o *ExecutionRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ExecutionRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ExecutionRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetResource

`func (o *ExecutionRead) GetResource() ResourceEnum`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ExecutionRead) GetResourceOk() (*ResourceEnum, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ExecutionRead) SetResource(v ResourceEnum)`

SetResource sets Resource field to given value.


### GetCursor

`func (o *ExecutionRead) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ExecutionRead) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ExecutionRead) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *ExecutionRead) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### SetCursorNil

`func (o *ExecutionRead) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *ExecutionRead) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
### GetResponse

`func (o *ExecutionRead) GetResponse() map[string]interface{}`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *ExecutionRead) GetResponseOk() (*map[string]interface{}, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *ExecutionRead) SetResponse(v map[string]interface{})`

SetResponse sets Response field to given value.

### HasResponse

`func (o *ExecutionRead) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponseNil

`func (o *ExecutionRead) SetResponseNil(b bool)`

 SetResponseNil sets the value for Response to be an explicit nil

### UnsetResponse
`func (o *ExecutionRead) UnsetResponse()`

UnsetResponse ensures that no value is present for Response, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


