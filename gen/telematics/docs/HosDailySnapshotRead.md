# HosDailySnapshotRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FleetId** | **NullableString** |  | 
**FleetRef** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | Unique identifier of the record at Catena Telematics. | 
**CreatedAt** | **time.Time** | Immutable: The datetime the record was ingested into Catena Telematics. | 
**UpdatedAt** | **time.Time** | The dateime the record was last modified in Catena Telematics. | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**ConnectionId** | **string** | Unique identifier of the connection at Catena Telematics through which this record was ingested. A connection represents a Fleet/TSP pairing. | 
**SourceName** | [**TspEnum**](TspEnum.md) | An enumeration identifying the TSP from which this record was sourced. | 
**SourceData** | Pointer to **map[string]interface{}** | Raw source payload as ingested from the TSP. **Note: use it for audit/debugging.** | [optional] 
**SourceId** | **string** | Unique identifier of the record in the TSP. **Note: we generate a unique composite key based on available fields if the TSP does not provide an unique ID.** | 
**SourceDataHash** | **string** | SHA-256 hash of the source data payload. **Note: we use it internally for idempotence and deduplication.** | 
**OccurredAt** | Pointer to **NullableTime** |  | [optional] 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 
**DriverId** | Pointer to **NullableString** |  | [optional] 
**SourceDriverId** | Pointer to **NullableString** |  | [optional] 
**SnapshotDate** | Pointer to **NullableString** |  | [optional] 
**DurationOffDutySeconds** | Pointer to **NullableInt32** |  | [optional] 
**DurationSleeperBerthSeconds** | Pointer to **NullableInt32** |  | [optional] 
**DurationOnDutySeconds** | Pointer to **NullableInt32** |  | [optional] 
**DurationDrivingSeconds** | Pointer to **NullableInt32** |  | [optional] 
**DurationPersonalConveyanceSeconds** | Pointer to **NullableInt32** |  | [optional] 
**DurationYardMoveSeconds** | Pointer to **NullableInt32** |  | [optional] 
**DurationWaitingSeconds** | Pointer to **NullableInt32** |  | [optional] 
**DurationUnknownSeconds** | Pointer to **NullableInt32** |  | [optional] 
**DurationCycleOnDutySeconds** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewHosDailySnapshotRead

`func NewHosDailySnapshotRead(fleetId NullableString, id string, createdAt time.Time, updatedAt time.Time, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *HosDailySnapshotRead`

NewHosDailySnapshotRead instantiates a new HosDailySnapshotRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHosDailySnapshotReadWithDefaults

`func NewHosDailySnapshotReadWithDefaults() *HosDailySnapshotRead`

NewHosDailySnapshotReadWithDefaults instantiates a new HosDailySnapshotRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *HosDailySnapshotRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *HosDailySnapshotRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *HosDailySnapshotRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### SetFleetIdNil

`func (o *HosDailySnapshotRead) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *HosDailySnapshotRead) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetFleetRef

`func (o *HosDailySnapshotRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *HosDailySnapshotRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *HosDailySnapshotRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *HosDailySnapshotRead) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *HosDailySnapshotRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *HosDailySnapshotRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetId

`func (o *HosDailySnapshotRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HosDailySnapshotRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HosDailySnapshotRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *HosDailySnapshotRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HosDailySnapshotRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HosDailySnapshotRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *HosDailySnapshotRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HosDailySnapshotRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HosDailySnapshotRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *HosDailySnapshotRead) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *HosDailySnapshotRead) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *HosDailySnapshotRead) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *HosDailySnapshotRead) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *HosDailySnapshotRead) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *HosDailySnapshotRead) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetConnectionId

`func (o *HosDailySnapshotRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *HosDailySnapshotRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *HosDailySnapshotRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceName

`func (o *HosDailySnapshotRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *HosDailySnapshotRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *HosDailySnapshotRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *HosDailySnapshotRead) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *HosDailySnapshotRead) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *HosDailySnapshotRead) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *HosDailySnapshotRead) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *HosDailySnapshotRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *HosDailySnapshotRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *HosDailySnapshotRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *HosDailySnapshotRead) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *HosDailySnapshotRead) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *HosDailySnapshotRead) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *HosDailySnapshotRead) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *HosDailySnapshotRead) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *HosDailySnapshotRead) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *HosDailySnapshotRead) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *HosDailySnapshotRead) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *HosDailySnapshotRead) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *HosDailySnapshotRead) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HosDailySnapshotRead) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HosDailySnapshotRead) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HosDailySnapshotRead) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HosDailySnapshotRead) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HosDailySnapshotRead) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *HosDailySnapshotRead) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *HosDailySnapshotRead) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *HosDailySnapshotRead) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *HosDailySnapshotRead) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *HosDailySnapshotRead) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *HosDailySnapshotRead) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetDriverId

`func (o *HosDailySnapshotRead) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *HosDailySnapshotRead) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *HosDailySnapshotRead) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *HosDailySnapshotRead) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *HosDailySnapshotRead) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *HosDailySnapshotRead) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetSourceDriverId

`func (o *HosDailySnapshotRead) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *HosDailySnapshotRead) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *HosDailySnapshotRead) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *HosDailySnapshotRead) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *HosDailySnapshotRead) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *HosDailySnapshotRead) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSnapshotDate

`func (o *HosDailySnapshotRead) GetSnapshotDate() string`

GetSnapshotDate returns the SnapshotDate field if non-nil, zero value otherwise.

### GetSnapshotDateOk

`func (o *HosDailySnapshotRead) GetSnapshotDateOk() (*string, bool)`

GetSnapshotDateOk returns a tuple with the SnapshotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotDate

`func (o *HosDailySnapshotRead) SetSnapshotDate(v string)`

SetSnapshotDate sets SnapshotDate field to given value.

### HasSnapshotDate

`func (o *HosDailySnapshotRead) HasSnapshotDate() bool`

HasSnapshotDate returns a boolean if a field has been set.

### SetSnapshotDateNil

`func (o *HosDailySnapshotRead) SetSnapshotDateNil(b bool)`

 SetSnapshotDateNil sets the value for SnapshotDate to be an explicit nil

### UnsetSnapshotDate
`func (o *HosDailySnapshotRead) UnsetSnapshotDate()`

UnsetSnapshotDate ensures that no value is present for SnapshotDate, not even an explicit nil
### GetDurationOffDutySeconds

`func (o *HosDailySnapshotRead) GetDurationOffDutySeconds() int32`

GetDurationOffDutySeconds returns the DurationOffDutySeconds field if non-nil, zero value otherwise.

### GetDurationOffDutySecondsOk

`func (o *HosDailySnapshotRead) GetDurationOffDutySecondsOk() (*int32, bool)`

GetDurationOffDutySecondsOk returns a tuple with the DurationOffDutySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationOffDutySeconds

`func (o *HosDailySnapshotRead) SetDurationOffDutySeconds(v int32)`

SetDurationOffDutySeconds sets DurationOffDutySeconds field to given value.

### HasDurationOffDutySeconds

`func (o *HosDailySnapshotRead) HasDurationOffDutySeconds() bool`

HasDurationOffDutySeconds returns a boolean if a field has been set.

### SetDurationOffDutySecondsNil

`func (o *HosDailySnapshotRead) SetDurationOffDutySecondsNil(b bool)`

 SetDurationOffDutySecondsNil sets the value for DurationOffDutySeconds to be an explicit nil

### UnsetDurationOffDutySeconds
`func (o *HosDailySnapshotRead) UnsetDurationOffDutySeconds()`

UnsetDurationOffDutySeconds ensures that no value is present for DurationOffDutySeconds, not even an explicit nil
### GetDurationSleeperBerthSeconds

`func (o *HosDailySnapshotRead) GetDurationSleeperBerthSeconds() int32`

GetDurationSleeperBerthSeconds returns the DurationSleeperBerthSeconds field if non-nil, zero value otherwise.

### GetDurationSleeperBerthSecondsOk

`func (o *HosDailySnapshotRead) GetDurationSleeperBerthSecondsOk() (*int32, bool)`

GetDurationSleeperBerthSecondsOk returns a tuple with the DurationSleeperBerthSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSleeperBerthSeconds

`func (o *HosDailySnapshotRead) SetDurationSleeperBerthSeconds(v int32)`

SetDurationSleeperBerthSeconds sets DurationSleeperBerthSeconds field to given value.

### HasDurationSleeperBerthSeconds

`func (o *HosDailySnapshotRead) HasDurationSleeperBerthSeconds() bool`

HasDurationSleeperBerthSeconds returns a boolean if a field has been set.

### SetDurationSleeperBerthSecondsNil

`func (o *HosDailySnapshotRead) SetDurationSleeperBerthSecondsNil(b bool)`

 SetDurationSleeperBerthSecondsNil sets the value for DurationSleeperBerthSeconds to be an explicit nil

### UnsetDurationSleeperBerthSeconds
`func (o *HosDailySnapshotRead) UnsetDurationSleeperBerthSeconds()`

UnsetDurationSleeperBerthSeconds ensures that no value is present for DurationSleeperBerthSeconds, not even an explicit nil
### GetDurationOnDutySeconds

`func (o *HosDailySnapshotRead) GetDurationOnDutySeconds() int32`

GetDurationOnDutySeconds returns the DurationOnDutySeconds field if non-nil, zero value otherwise.

### GetDurationOnDutySecondsOk

`func (o *HosDailySnapshotRead) GetDurationOnDutySecondsOk() (*int32, bool)`

GetDurationOnDutySecondsOk returns a tuple with the DurationOnDutySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationOnDutySeconds

`func (o *HosDailySnapshotRead) SetDurationOnDutySeconds(v int32)`

SetDurationOnDutySeconds sets DurationOnDutySeconds field to given value.

### HasDurationOnDutySeconds

`func (o *HosDailySnapshotRead) HasDurationOnDutySeconds() bool`

HasDurationOnDutySeconds returns a boolean if a field has been set.

### SetDurationOnDutySecondsNil

`func (o *HosDailySnapshotRead) SetDurationOnDutySecondsNil(b bool)`

 SetDurationOnDutySecondsNil sets the value for DurationOnDutySeconds to be an explicit nil

### UnsetDurationOnDutySeconds
`func (o *HosDailySnapshotRead) UnsetDurationOnDutySeconds()`

UnsetDurationOnDutySeconds ensures that no value is present for DurationOnDutySeconds, not even an explicit nil
### GetDurationDrivingSeconds

`func (o *HosDailySnapshotRead) GetDurationDrivingSeconds() int32`

GetDurationDrivingSeconds returns the DurationDrivingSeconds field if non-nil, zero value otherwise.

### GetDurationDrivingSecondsOk

`func (o *HosDailySnapshotRead) GetDurationDrivingSecondsOk() (*int32, bool)`

GetDurationDrivingSecondsOk returns a tuple with the DurationDrivingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationDrivingSeconds

`func (o *HosDailySnapshotRead) SetDurationDrivingSeconds(v int32)`

SetDurationDrivingSeconds sets DurationDrivingSeconds field to given value.

### HasDurationDrivingSeconds

`func (o *HosDailySnapshotRead) HasDurationDrivingSeconds() bool`

HasDurationDrivingSeconds returns a boolean if a field has been set.

### SetDurationDrivingSecondsNil

`func (o *HosDailySnapshotRead) SetDurationDrivingSecondsNil(b bool)`

 SetDurationDrivingSecondsNil sets the value for DurationDrivingSeconds to be an explicit nil

### UnsetDurationDrivingSeconds
`func (o *HosDailySnapshotRead) UnsetDurationDrivingSeconds()`

UnsetDurationDrivingSeconds ensures that no value is present for DurationDrivingSeconds, not even an explicit nil
### GetDurationPersonalConveyanceSeconds

`func (o *HosDailySnapshotRead) GetDurationPersonalConveyanceSeconds() int32`

GetDurationPersonalConveyanceSeconds returns the DurationPersonalConveyanceSeconds field if non-nil, zero value otherwise.

### GetDurationPersonalConveyanceSecondsOk

`func (o *HosDailySnapshotRead) GetDurationPersonalConveyanceSecondsOk() (*int32, bool)`

GetDurationPersonalConveyanceSecondsOk returns a tuple with the DurationPersonalConveyanceSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationPersonalConveyanceSeconds

`func (o *HosDailySnapshotRead) SetDurationPersonalConveyanceSeconds(v int32)`

SetDurationPersonalConveyanceSeconds sets DurationPersonalConveyanceSeconds field to given value.

### HasDurationPersonalConveyanceSeconds

`func (o *HosDailySnapshotRead) HasDurationPersonalConveyanceSeconds() bool`

HasDurationPersonalConveyanceSeconds returns a boolean if a field has been set.

### SetDurationPersonalConveyanceSecondsNil

`func (o *HosDailySnapshotRead) SetDurationPersonalConveyanceSecondsNil(b bool)`

 SetDurationPersonalConveyanceSecondsNil sets the value for DurationPersonalConveyanceSeconds to be an explicit nil

### UnsetDurationPersonalConveyanceSeconds
`func (o *HosDailySnapshotRead) UnsetDurationPersonalConveyanceSeconds()`

UnsetDurationPersonalConveyanceSeconds ensures that no value is present for DurationPersonalConveyanceSeconds, not even an explicit nil
### GetDurationYardMoveSeconds

`func (o *HosDailySnapshotRead) GetDurationYardMoveSeconds() int32`

GetDurationYardMoveSeconds returns the DurationYardMoveSeconds field if non-nil, zero value otherwise.

### GetDurationYardMoveSecondsOk

`func (o *HosDailySnapshotRead) GetDurationYardMoveSecondsOk() (*int32, bool)`

GetDurationYardMoveSecondsOk returns a tuple with the DurationYardMoveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationYardMoveSeconds

`func (o *HosDailySnapshotRead) SetDurationYardMoveSeconds(v int32)`

SetDurationYardMoveSeconds sets DurationYardMoveSeconds field to given value.

### HasDurationYardMoveSeconds

`func (o *HosDailySnapshotRead) HasDurationYardMoveSeconds() bool`

HasDurationYardMoveSeconds returns a boolean if a field has been set.

### SetDurationYardMoveSecondsNil

`func (o *HosDailySnapshotRead) SetDurationYardMoveSecondsNil(b bool)`

 SetDurationYardMoveSecondsNil sets the value for DurationYardMoveSeconds to be an explicit nil

### UnsetDurationYardMoveSeconds
`func (o *HosDailySnapshotRead) UnsetDurationYardMoveSeconds()`

UnsetDurationYardMoveSeconds ensures that no value is present for DurationYardMoveSeconds, not even an explicit nil
### GetDurationWaitingSeconds

`func (o *HosDailySnapshotRead) GetDurationWaitingSeconds() int32`

GetDurationWaitingSeconds returns the DurationWaitingSeconds field if non-nil, zero value otherwise.

### GetDurationWaitingSecondsOk

`func (o *HosDailySnapshotRead) GetDurationWaitingSecondsOk() (*int32, bool)`

GetDurationWaitingSecondsOk returns a tuple with the DurationWaitingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationWaitingSeconds

`func (o *HosDailySnapshotRead) SetDurationWaitingSeconds(v int32)`

SetDurationWaitingSeconds sets DurationWaitingSeconds field to given value.

### HasDurationWaitingSeconds

`func (o *HosDailySnapshotRead) HasDurationWaitingSeconds() bool`

HasDurationWaitingSeconds returns a boolean if a field has been set.

### SetDurationWaitingSecondsNil

`func (o *HosDailySnapshotRead) SetDurationWaitingSecondsNil(b bool)`

 SetDurationWaitingSecondsNil sets the value for DurationWaitingSeconds to be an explicit nil

### UnsetDurationWaitingSeconds
`func (o *HosDailySnapshotRead) UnsetDurationWaitingSeconds()`

UnsetDurationWaitingSeconds ensures that no value is present for DurationWaitingSeconds, not even an explicit nil
### GetDurationUnknownSeconds

`func (o *HosDailySnapshotRead) GetDurationUnknownSeconds() int32`

GetDurationUnknownSeconds returns the DurationUnknownSeconds field if non-nil, zero value otherwise.

### GetDurationUnknownSecondsOk

`func (o *HosDailySnapshotRead) GetDurationUnknownSecondsOk() (*int32, bool)`

GetDurationUnknownSecondsOk returns a tuple with the DurationUnknownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationUnknownSeconds

`func (o *HosDailySnapshotRead) SetDurationUnknownSeconds(v int32)`

SetDurationUnknownSeconds sets DurationUnknownSeconds field to given value.

### HasDurationUnknownSeconds

`func (o *HosDailySnapshotRead) HasDurationUnknownSeconds() bool`

HasDurationUnknownSeconds returns a boolean if a field has been set.

### SetDurationUnknownSecondsNil

`func (o *HosDailySnapshotRead) SetDurationUnknownSecondsNil(b bool)`

 SetDurationUnknownSecondsNil sets the value for DurationUnknownSeconds to be an explicit nil

### UnsetDurationUnknownSeconds
`func (o *HosDailySnapshotRead) UnsetDurationUnknownSeconds()`

UnsetDurationUnknownSeconds ensures that no value is present for DurationUnknownSeconds, not even an explicit nil
### GetDurationCycleOnDutySeconds

`func (o *HosDailySnapshotRead) GetDurationCycleOnDutySeconds() int32`

GetDurationCycleOnDutySeconds returns the DurationCycleOnDutySeconds field if non-nil, zero value otherwise.

### GetDurationCycleOnDutySecondsOk

`func (o *HosDailySnapshotRead) GetDurationCycleOnDutySecondsOk() (*int32, bool)`

GetDurationCycleOnDutySecondsOk returns a tuple with the DurationCycleOnDutySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationCycleOnDutySeconds

`func (o *HosDailySnapshotRead) SetDurationCycleOnDutySeconds(v int32)`

SetDurationCycleOnDutySeconds sets DurationCycleOnDutySeconds field to given value.

### HasDurationCycleOnDutySeconds

`func (o *HosDailySnapshotRead) HasDurationCycleOnDutySeconds() bool`

HasDurationCycleOnDutySeconds returns a boolean if a field has been set.

### SetDurationCycleOnDutySecondsNil

`func (o *HosDailySnapshotRead) SetDurationCycleOnDutySecondsNil(b bool)`

 SetDurationCycleOnDutySecondsNil sets the value for DurationCycleOnDutySeconds to be an explicit nil

### UnsetDurationCycleOnDutySeconds
`func (o *HosDailySnapshotRead) UnsetDurationCycleOnDutySeconds()`

UnsetDurationCycleOnDutySeconds ensures that no value is present for DurationCycleOnDutySeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


