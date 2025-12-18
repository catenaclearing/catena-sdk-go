# HosDailySnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the record at Catena Telematics. | 
**CreatedAt** | **time.Time** | Immutable: The datetime the record was ingested into Catena Telematics. | 
**UpdatedAt** | **time.Time** | The dateime the record was last modified in Catena Telematics. | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**FleetId** | **string** | Unique identifier of the fleet at Catena Telematics. This record belongs to this fleet. **Note: this is not the fleet ID in the TSP system or in your organization&#39;s systems, use &#x60;share_agreements&#x60; to map &#x60;fleet_ids&#x60; to &#x60;fleet_refs&#x60;.** | 
**ConnectionId** | **string** | Unique identifier of the connection at Catena Telematics through which this record was ingested. A connection represents a Fleet/TSP pairing. | 
**SourceName** | [**TspEnum**](TspEnum.md) | An enumeration identifying the TSP from which this record was sourced. | 
**SourceData** | Pointer to **map[string]interface{}** | Raw source payload as ingested from the TSP. **Note: use it for audit/debugging.** | [optional] 
**SourceId** | **string** | Unique identifier of the record in the TSP. **Note: we generate a unique composite key based on available fields if the TSP does not provide an unique ID.** | 
**SourceDataHash** | **string** | SHA-256 hash of the source data payload. **Note: we use it internally for idempotence and deduplication.** | 
**OccurredAt** | Pointer to **NullableTime** |  | [optional] 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 
**DriverId** | Pointer to **NullableString** |  | [optional] 
**SnapshotDate** | Pointer to **NullableString** |  | [optional] 
**TimeZoneCode** | Pointer to [**NullableTimezoneCodeEnum**](TimezoneCodeEnum.md) |  | [optional] 
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

### NewHosDailySnapshot

`func NewHosDailySnapshot(id string, createdAt time.Time, updatedAt time.Time, fleetId string, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *HosDailySnapshot`

NewHosDailySnapshot instantiates a new HosDailySnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHosDailySnapshotWithDefaults

`func NewHosDailySnapshotWithDefaults() *HosDailySnapshot`

NewHosDailySnapshotWithDefaults instantiates a new HosDailySnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HosDailySnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HosDailySnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HosDailySnapshot) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *HosDailySnapshot) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HosDailySnapshot) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HosDailySnapshot) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *HosDailySnapshot) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HosDailySnapshot) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HosDailySnapshot) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *HosDailySnapshot) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *HosDailySnapshot) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *HosDailySnapshot) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *HosDailySnapshot) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *HosDailySnapshot) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *HosDailySnapshot) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetFleetId

`func (o *HosDailySnapshot) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *HosDailySnapshot) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *HosDailySnapshot) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetConnectionId

`func (o *HosDailySnapshot) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *HosDailySnapshot) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *HosDailySnapshot) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceName

`func (o *HosDailySnapshot) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *HosDailySnapshot) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *HosDailySnapshot) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *HosDailySnapshot) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *HosDailySnapshot) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *HosDailySnapshot) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *HosDailySnapshot) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *HosDailySnapshot) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *HosDailySnapshot) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *HosDailySnapshot) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *HosDailySnapshot) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *HosDailySnapshot) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *HosDailySnapshot) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *HosDailySnapshot) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *HosDailySnapshot) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *HosDailySnapshot) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *HosDailySnapshot) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *HosDailySnapshot) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *HosDailySnapshot) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *HosDailySnapshot) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HosDailySnapshot) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HosDailySnapshot) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HosDailySnapshot) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HosDailySnapshot) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HosDailySnapshot) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *HosDailySnapshot) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *HosDailySnapshot) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *HosDailySnapshot) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *HosDailySnapshot) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *HosDailySnapshot) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *HosDailySnapshot) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetDriverId

`func (o *HosDailySnapshot) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *HosDailySnapshot) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *HosDailySnapshot) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *HosDailySnapshot) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *HosDailySnapshot) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *HosDailySnapshot) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetSnapshotDate

`func (o *HosDailySnapshot) GetSnapshotDate() string`

GetSnapshotDate returns the SnapshotDate field if non-nil, zero value otherwise.

### GetSnapshotDateOk

`func (o *HosDailySnapshot) GetSnapshotDateOk() (*string, bool)`

GetSnapshotDateOk returns a tuple with the SnapshotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotDate

`func (o *HosDailySnapshot) SetSnapshotDate(v string)`

SetSnapshotDate sets SnapshotDate field to given value.

### HasSnapshotDate

`func (o *HosDailySnapshot) HasSnapshotDate() bool`

HasSnapshotDate returns a boolean if a field has been set.

### SetSnapshotDateNil

`func (o *HosDailySnapshot) SetSnapshotDateNil(b bool)`

 SetSnapshotDateNil sets the value for SnapshotDate to be an explicit nil

### UnsetSnapshotDate
`func (o *HosDailySnapshot) UnsetSnapshotDate()`

UnsetSnapshotDate ensures that no value is present for SnapshotDate, not even an explicit nil
### GetTimeZoneCode

`func (o *HosDailySnapshot) GetTimeZoneCode() TimezoneCodeEnum`

GetTimeZoneCode returns the TimeZoneCode field if non-nil, zero value otherwise.

### GetTimeZoneCodeOk

`func (o *HosDailySnapshot) GetTimeZoneCodeOk() (*TimezoneCodeEnum, bool)`

GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneCode

`func (o *HosDailySnapshot) SetTimeZoneCode(v TimezoneCodeEnum)`

SetTimeZoneCode sets TimeZoneCode field to given value.

### HasTimeZoneCode

`func (o *HosDailySnapshot) HasTimeZoneCode() bool`

HasTimeZoneCode returns a boolean if a field has been set.

### SetTimeZoneCodeNil

`func (o *HosDailySnapshot) SetTimeZoneCodeNil(b bool)`

 SetTimeZoneCodeNil sets the value for TimeZoneCode to be an explicit nil

### UnsetTimeZoneCode
`func (o *HosDailySnapshot) UnsetTimeZoneCode()`

UnsetTimeZoneCode ensures that no value is present for TimeZoneCode, not even an explicit nil
### GetDurationOffDutySeconds

`func (o *HosDailySnapshot) GetDurationOffDutySeconds() int32`

GetDurationOffDutySeconds returns the DurationOffDutySeconds field if non-nil, zero value otherwise.

### GetDurationOffDutySecondsOk

`func (o *HosDailySnapshot) GetDurationOffDutySecondsOk() (*int32, bool)`

GetDurationOffDutySecondsOk returns a tuple with the DurationOffDutySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationOffDutySeconds

`func (o *HosDailySnapshot) SetDurationOffDutySeconds(v int32)`

SetDurationOffDutySeconds sets DurationOffDutySeconds field to given value.

### HasDurationOffDutySeconds

`func (o *HosDailySnapshot) HasDurationOffDutySeconds() bool`

HasDurationOffDutySeconds returns a boolean if a field has been set.

### SetDurationOffDutySecondsNil

`func (o *HosDailySnapshot) SetDurationOffDutySecondsNil(b bool)`

 SetDurationOffDutySecondsNil sets the value for DurationOffDutySeconds to be an explicit nil

### UnsetDurationOffDutySeconds
`func (o *HosDailySnapshot) UnsetDurationOffDutySeconds()`

UnsetDurationOffDutySeconds ensures that no value is present for DurationOffDutySeconds, not even an explicit nil
### GetDurationSleeperBerthSeconds

`func (o *HosDailySnapshot) GetDurationSleeperBerthSeconds() int32`

GetDurationSleeperBerthSeconds returns the DurationSleeperBerthSeconds field if non-nil, zero value otherwise.

### GetDurationSleeperBerthSecondsOk

`func (o *HosDailySnapshot) GetDurationSleeperBerthSecondsOk() (*int32, bool)`

GetDurationSleeperBerthSecondsOk returns a tuple with the DurationSleeperBerthSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSleeperBerthSeconds

`func (o *HosDailySnapshot) SetDurationSleeperBerthSeconds(v int32)`

SetDurationSleeperBerthSeconds sets DurationSleeperBerthSeconds field to given value.

### HasDurationSleeperBerthSeconds

`func (o *HosDailySnapshot) HasDurationSleeperBerthSeconds() bool`

HasDurationSleeperBerthSeconds returns a boolean if a field has been set.

### SetDurationSleeperBerthSecondsNil

`func (o *HosDailySnapshot) SetDurationSleeperBerthSecondsNil(b bool)`

 SetDurationSleeperBerthSecondsNil sets the value for DurationSleeperBerthSeconds to be an explicit nil

### UnsetDurationSleeperBerthSeconds
`func (o *HosDailySnapshot) UnsetDurationSleeperBerthSeconds()`

UnsetDurationSleeperBerthSeconds ensures that no value is present for DurationSleeperBerthSeconds, not even an explicit nil
### GetDurationOnDutySeconds

`func (o *HosDailySnapshot) GetDurationOnDutySeconds() int32`

GetDurationOnDutySeconds returns the DurationOnDutySeconds field if non-nil, zero value otherwise.

### GetDurationOnDutySecondsOk

`func (o *HosDailySnapshot) GetDurationOnDutySecondsOk() (*int32, bool)`

GetDurationOnDutySecondsOk returns a tuple with the DurationOnDutySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationOnDutySeconds

`func (o *HosDailySnapshot) SetDurationOnDutySeconds(v int32)`

SetDurationOnDutySeconds sets DurationOnDutySeconds field to given value.

### HasDurationOnDutySeconds

`func (o *HosDailySnapshot) HasDurationOnDutySeconds() bool`

HasDurationOnDutySeconds returns a boolean if a field has been set.

### SetDurationOnDutySecondsNil

`func (o *HosDailySnapshot) SetDurationOnDutySecondsNil(b bool)`

 SetDurationOnDutySecondsNil sets the value for DurationOnDutySeconds to be an explicit nil

### UnsetDurationOnDutySeconds
`func (o *HosDailySnapshot) UnsetDurationOnDutySeconds()`

UnsetDurationOnDutySeconds ensures that no value is present for DurationOnDutySeconds, not even an explicit nil
### GetDurationDrivingSeconds

`func (o *HosDailySnapshot) GetDurationDrivingSeconds() int32`

GetDurationDrivingSeconds returns the DurationDrivingSeconds field if non-nil, zero value otherwise.

### GetDurationDrivingSecondsOk

`func (o *HosDailySnapshot) GetDurationDrivingSecondsOk() (*int32, bool)`

GetDurationDrivingSecondsOk returns a tuple with the DurationDrivingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationDrivingSeconds

`func (o *HosDailySnapshot) SetDurationDrivingSeconds(v int32)`

SetDurationDrivingSeconds sets DurationDrivingSeconds field to given value.

### HasDurationDrivingSeconds

`func (o *HosDailySnapshot) HasDurationDrivingSeconds() bool`

HasDurationDrivingSeconds returns a boolean if a field has been set.

### SetDurationDrivingSecondsNil

`func (o *HosDailySnapshot) SetDurationDrivingSecondsNil(b bool)`

 SetDurationDrivingSecondsNil sets the value for DurationDrivingSeconds to be an explicit nil

### UnsetDurationDrivingSeconds
`func (o *HosDailySnapshot) UnsetDurationDrivingSeconds()`

UnsetDurationDrivingSeconds ensures that no value is present for DurationDrivingSeconds, not even an explicit nil
### GetDurationPersonalConveyanceSeconds

`func (o *HosDailySnapshot) GetDurationPersonalConveyanceSeconds() int32`

GetDurationPersonalConveyanceSeconds returns the DurationPersonalConveyanceSeconds field if non-nil, zero value otherwise.

### GetDurationPersonalConveyanceSecondsOk

`func (o *HosDailySnapshot) GetDurationPersonalConveyanceSecondsOk() (*int32, bool)`

GetDurationPersonalConveyanceSecondsOk returns a tuple with the DurationPersonalConveyanceSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationPersonalConveyanceSeconds

`func (o *HosDailySnapshot) SetDurationPersonalConveyanceSeconds(v int32)`

SetDurationPersonalConveyanceSeconds sets DurationPersonalConveyanceSeconds field to given value.

### HasDurationPersonalConveyanceSeconds

`func (o *HosDailySnapshot) HasDurationPersonalConveyanceSeconds() bool`

HasDurationPersonalConveyanceSeconds returns a boolean if a field has been set.

### SetDurationPersonalConveyanceSecondsNil

`func (o *HosDailySnapshot) SetDurationPersonalConveyanceSecondsNil(b bool)`

 SetDurationPersonalConveyanceSecondsNil sets the value for DurationPersonalConveyanceSeconds to be an explicit nil

### UnsetDurationPersonalConveyanceSeconds
`func (o *HosDailySnapshot) UnsetDurationPersonalConveyanceSeconds()`

UnsetDurationPersonalConveyanceSeconds ensures that no value is present for DurationPersonalConveyanceSeconds, not even an explicit nil
### GetDurationYardMoveSeconds

`func (o *HosDailySnapshot) GetDurationYardMoveSeconds() int32`

GetDurationYardMoveSeconds returns the DurationYardMoveSeconds field if non-nil, zero value otherwise.

### GetDurationYardMoveSecondsOk

`func (o *HosDailySnapshot) GetDurationYardMoveSecondsOk() (*int32, bool)`

GetDurationYardMoveSecondsOk returns a tuple with the DurationYardMoveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationYardMoveSeconds

`func (o *HosDailySnapshot) SetDurationYardMoveSeconds(v int32)`

SetDurationYardMoveSeconds sets DurationYardMoveSeconds field to given value.

### HasDurationYardMoveSeconds

`func (o *HosDailySnapshot) HasDurationYardMoveSeconds() bool`

HasDurationYardMoveSeconds returns a boolean if a field has been set.

### SetDurationYardMoveSecondsNil

`func (o *HosDailySnapshot) SetDurationYardMoveSecondsNil(b bool)`

 SetDurationYardMoveSecondsNil sets the value for DurationYardMoveSeconds to be an explicit nil

### UnsetDurationYardMoveSeconds
`func (o *HosDailySnapshot) UnsetDurationYardMoveSeconds()`

UnsetDurationYardMoveSeconds ensures that no value is present for DurationYardMoveSeconds, not even an explicit nil
### GetDurationWaitingSeconds

`func (o *HosDailySnapshot) GetDurationWaitingSeconds() int32`

GetDurationWaitingSeconds returns the DurationWaitingSeconds field if non-nil, zero value otherwise.

### GetDurationWaitingSecondsOk

`func (o *HosDailySnapshot) GetDurationWaitingSecondsOk() (*int32, bool)`

GetDurationWaitingSecondsOk returns a tuple with the DurationWaitingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationWaitingSeconds

`func (o *HosDailySnapshot) SetDurationWaitingSeconds(v int32)`

SetDurationWaitingSeconds sets DurationWaitingSeconds field to given value.

### HasDurationWaitingSeconds

`func (o *HosDailySnapshot) HasDurationWaitingSeconds() bool`

HasDurationWaitingSeconds returns a boolean if a field has been set.

### SetDurationWaitingSecondsNil

`func (o *HosDailySnapshot) SetDurationWaitingSecondsNil(b bool)`

 SetDurationWaitingSecondsNil sets the value for DurationWaitingSeconds to be an explicit nil

### UnsetDurationWaitingSeconds
`func (o *HosDailySnapshot) UnsetDurationWaitingSeconds()`

UnsetDurationWaitingSeconds ensures that no value is present for DurationWaitingSeconds, not even an explicit nil
### GetDurationUnknownSeconds

`func (o *HosDailySnapshot) GetDurationUnknownSeconds() int32`

GetDurationUnknownSeconds returns the DurationUnknownSeconds field if non-nil, zero value otherwise.

### GetDurationUnknownSecondsOk

`func (o *HosDailySnapshot) GetDurationUnknownSecondsOk() (*int32, bool)`

GetDurationUnknownSecondsOk returns a tuple with the DurationUnknownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationUnknownSeconds

`func (o *HosDailySnapshot) SetDurationUnknownSeconds(v int32)`

SetDurationUnknownSeconds sets DurationUnknownSeconds field to given value.

### HasDurationUnknownSeconds

`func (o *HosDailySnapshot) HasDurationUnknownSeconds() bool`

HasDurationUnknownSeconds returns a boolean if a field has been set.

### SetDurationUnknownSecondsNil

`func (o *HosDailySnapshot) SetDurationUnknownSecondsNil(b bool)`

 SetDurationUnknownSecondsNil sets the value for DurationUnknownSeconds to be an explicit nil

### UnsetDurationUnknownSeconds
`func (o *HosDailySnapshot) UnsetDurationUnknownSeconds()`

UnsetDurationUnknownSeconds ensures that no value is present for DurationUnknownSeconds, not even an explicit nil
### GetDurationCycleOnDutySeconds

`func (o *HosDailySnapshot) GetDurationCycleOnDutySeconds() int32`

GetDurationCycleOnDutySeconds returns the DurationCycleOnDutySeconds field if non-nil, zero value otherwise.

### GetDurationCycleOnDutySecondsOk

`func (o *HosDailySnapshot) GetDurationCycleOnDutySecondsOk() (*int32, bool)`

GetDurationCycleOnDutySecondsOk returns a tuple with the DurationCycleOnDutySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationCycleOnDutySeconds

`func (o *HosDailySnapshot) SetDurationCycleOnDutySeconds(v int32)`

SetDurationCycleOnDutySeconds sets DurationCycleOnDutySeconds field to given value.

### HasDurationCycleOnDutySeconds

`func (o *HosDailySnapshot) HasDurationCycleOnDutySeconds() bool`

HasDurationCycleOnDutySeconds returns a boolean if a field has been set.

### SetDurationCycleOnDutySecondsNil

`func (o *HosDailySnapshot) SetDurationCycleOnDutySecondsNil(b bool)`

 SetDurationCycleOnDutySecondsNil sets the value for DurationCycleOnDutySeconds to be an explicit nil

### UnsetDurationCycleOnDutySeconds
`func (o *HosDailySnapshot) UnsetDurationCycleOnDutySeconds()`

UnsetDurationCycleOnDutySeconds ensures that no value is present for DurationCycleOnDutySeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


