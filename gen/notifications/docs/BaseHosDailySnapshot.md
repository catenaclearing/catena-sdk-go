# BaseHosDailySnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal unique identifier for the telematics event record (Catena PK). | 
**FleetId** | **string** | The Catena fleet this record belongs to (multi-tenant scope). | 
**FleetRef** | **NullableString** |  | 
**TspId** | Pointer to **NullableString** |  | [optional] 
**TspSlug** | Pointer to **NullableString** |  | [optional] 
**SourceName** | [**TspEnum**](TspEnum.md) | The underlying telematics platform that provided this data (e.g., &#x60;samsara&#x60;, &#x60;motive&#x60;, &#x60;hos247&#x60;). Note: Some platforms like &#x60;hos247&#x60; offer white-labeling, so multiple TSPs may share the same source_name — use &#x60;tsp_id&#x60; or &#x60;tsp_slug&#x60; to identify the specific ELD provider. | 
**ConnectionId** | **string** | The specific fleet↔TSP connection through which this record was sourced. | 
**SourceId** | **string** | The ID of the record in the TSP or a deterministic ID/Hash generated from a composite unique key | 
**CreatedAt** | **time.Time** | Immutable: first time this record was ingested into our system. | 
**UpdatedAt** | **time.Time** | Last time we modified this record in our system. | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**OccurredAt** | **time.Time** | When the underlying event/observation occurred, as reported by the TSP, or the moment it was ingested by us if not available. | 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**DriverId** | Pointer to **NullableString** |  | [optional] 
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

### NewBaseHosDailySnapshot

`func NewBaseHosDailySnapshot(id string, fleetId string, fleetRef NullableString, sourceName TspEnum, connectionId string, sourceId string, createdAt time.Time, updatedAt time.Time, occurredAt time.Time, ) *BaseHosDailySnapshot`

NewBaseHosDailySnapshot instantiates a new BaseHosDailySnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseHosDailySnapshotWithDefaults

`func NewBaseHosDailySnapshotWithDefaults() *BaseHosDailySnapshot`

NewBaseHosDailySnapshotWithDefaults instantiates a new BaseHosDailySnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseHosDailySnapshot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseHosDailySnapshot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseHosDailySnapshot) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseHosDailySnapshot) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseHosDailySnapshot) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseHosDailySnapshot) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *BaseHosDailySnapshot) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *BaseHosDailySnapshot) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *BaseHosDailySnapshot) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### SetFleetRefNil

`func (o *BaseHosDailySnapshot) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *BaseHosDailySnapshot) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetTspId

`func (o *BaseHosDailySnapshot) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *BaseHosDailySnapshot) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *BaseHosDailySnapshot) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *BaseHosDailySnapshot) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *BaseHosDailySnapshot) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *BaseHosDailySnapshot) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *BaseHosDailySnapshot) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *BaseHosDailySnapshot) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *BaseHosDailySnapshot) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *BaseHosDailySnapshot) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *BaseHosDailySnapshot) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *BaseHosDailySnapshot) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *BaseHosDailySnapshot) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseHosDailySnapshot) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseHosDailySnapshot) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetConnectionId

`func (o *BaseHosDailySnapshot) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseHosDailySnapshot) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseHosDailySnapshot) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceId

`func (o *BaseHosDailySnapshot) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseHosDailySnapshot) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseHosDailySnapshot) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCreatedAt

`func (o *BaseHosDailySnapshot) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseHosDailySnapshot) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseHosDailySnapshot) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseHosDailySnapshot) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseHosDailySnapshot) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseHosDailySnapshot) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BaseHosDailySnapshot) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BaseHosDailySnapshot) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BaseHosDailySnapshot) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BaseHosDailySnapshot) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BaseHosDailySnapshot) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BaseHosDailySnapshot) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOccurredAt

`func (o *BaseHosDailySnapshot) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BaseHosDailySnapshot) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BaseHosDailySnapshot) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetExecutionId

`func (o *BaseHosDailySnapshot) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BaseHosDailySnapshot) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BaseHosDailySnapshot) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *BaseHosDailySnapshot) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *BaseHosDailySnapshot) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *BaseHosDailySnapshot) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *BaseHosDailySnapshot) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BaseHosDailySnapshot) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BaseHosDailySnapshot) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *BaseHosDailySnapshot) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *BaseHosDailySnapshot) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *BaseHosDailySnapshot) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *BaseHosDailySnapshot) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *BaseHosDailySnapshot) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *BaseHosDailySnapshot) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *BaseHosDailySnapshot) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *BaseHosDailySnapshot) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *BaseHosDailySnapshot) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetDriverId

`func (o *BaseHosDailySnapshot) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *BaseHosDailySnapshot) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *BaseHosDailySnapshot) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *BaseHosDailySnapshot) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *BaseHosDailySnapshot) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *BaseHosDailySnapshot) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetSnapshotDate

`func (o *BaseHosDailySnapshot) GetSnapshotDate() string`

GetSnapshotDate returns the SnapshotDate field if non-nil, zero value otherwise.

### GetSnapshotDateOk

`func (o *BaseHosDailySnapshot) GetSnapshotDateOk() (*string, bool)`

GetSnapshotDateOk returns a tuple with the SnapshotDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotDate

`func (o *BaseHosDailySnapshot) SetSnapshotDate(v string)`

SetSnapshotDate sets SnapshotDate field to given value.

### HasSnapshotDate

`func (o *BaseHosDailySnapshot) HasSnapshotDate() bool`

HasSnapshotDate returns a boolean if a field has been set.

### SetSnapshotDateNil

`func (o *BaseHosDailySnapshot) SetSnapshotDateNil(b bool)`

 SetSnapshotDateNil sets the value for SnapshotDate to be an explicit nil

### UnsetSnapshotDate
`func (o *BaseHosDailySnapshot) UnsetSnapshotDate()`

UnsetSnapshotDate ensures that no value is present for SnapshotDate, not even an explicit nil
### GetDurationOffDutySeconds

`func (o *BaseHosDailySnapshot) GetDurationOffDutySeconds() int32`

GetDurationOffDutySeconds returns the DurationOffDutySeconds field if non-nil, zero value otherwise.

### GetDurationOffDutySecondsOk

`func (o *BaseHosDailySnapshot) GetDurationOffDutySecondsOk() (*int32, bool)`

GetDurationOffDutySecondsOk returns a tuple with the DurationOffDutySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationOffDutySeconds

`func (o *BaseHosDailySnapshot) SetDurationOffDutySeconds(v int32)`

SetDurationOffDutySeconds sets DurationOffDutySeconds field to given value.

### HasDurationOffDutySeconds

`func (o *BaseHosDailySnapshot) HasDurationOffDutySeconds() bool`

HasDurationOffDutySeconds returns a boolean if a field has been set.

### SetDurationOffDutySecondsNil

`func (o *BaseHosDailySnapshot) SetDurationOffDutySecondsNil(b bool)`

 SetDurationOffDutySecondsNil sets the value for DurationOffDutySeconds to be an explicit nil

### UnsetDurationOffDutySeconds
`func (o *BaseHosDailySnapshot) UnsetDurationOffDutySeconds()`

UnsetDurationOffDutySeconds ensures that no value is present for DurationOffDutySeconds, not even an explicit nil
### GetDurationSleeperBerthSeconds

`func (o *BaseHosDailySnapshot) GetDurationSleeperBerthSeconds() int32`

GetDurationSleeperBerthSeconds returns the DurationSleeperBerthSeconds field if non-nil, zero value otherwise.

### GetDurationSleeperBerthSecondsOk

`func (o *BaseHosDailySnapshot) GetDurationSleeperBerthSecondsOk() (*int32, bool)`

GetDurationSleeperBerthSecondsOk returns a tuple with the DurationSleeperBerthSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSleeperBerthSeconds

`func (o *BaseHosDailySnapshot) SetDurationSleeperBerthSeconds(v int32)`

SetDurationSleeperBerthSeconds sets DurationSleeperBerthSeconds field to given value.

### HasDurationSleeperBerthSeconds

`func (o *BaseHosDailySnapshot) HasDurationSleeperBerthSeconds() bool`

HasDurationSleeperBerthSeconds returns a boolean if a field has been set.

### SetDurationSleeperBerthSecondsNil

`func (o *BaseHosDailySnapshot) SetDurationSleeperBerthSecondsNil(b bool)`

 SetDurationSleeperBerthSecondsNil sets the value for DurationSleeperBerthSeconds to be an explicit nil

### UnsetDurationSleeperBerthSeconds
`func (o *BaseHosDailySnapshot) UnsetDurationSleeperBerthSeconds()`

UnsetDurationSleeperBerthSeconds ensures that no value is present for DurationSleeperBerthSeconds, not even an explicit nil
### GetDurationOnDutySeconds

`func (o *BaseHosDailySnapshot) GetDurationOnDutySeconds() int32`

GetDurationOnDutySeconds returns the DurationOnDutySeconds field if non-nil, zero value otherwise.

### GetDurationOnDutySecondsOk

`func (o *BaseHosDailySnapshot) GetDurationOnDutySecondsOk() (*int32, bool)`

GetDurationOnDutySecondsOk returns a tuple with the DurationOnDutySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationOnDutySeconds

`func (o *BaseHosDailySnapshot) SetDurationOnDutySeconds(v int32)`

SetDurationOnDutySeconds sets DurationOnDutySeconds field to given value.

### HasDurationOnDutySeconds

`func (o *BaseHosDailySnapshot) HasDurationOnDutySeconds() bool`

HasDurationOnDutySeconds returns a boolean if a field has been set.

### SetDurationOnDutySecondsNil

`func (o *BaseHosDailySnapshot) SetDurationOnDutySecondsNil(b bool)`

 SetDurationOnDutySecondsNil sets the value for DurationOnDutySeconds to be an explicit nil

### UnsetDurationOnDutySeconds
`func (o *BaseHosDailySnapshot) UnsetDurationOnDutySeconds()`

UnsetDurationOnDutySeconds ensures that no value is present for DurationOnDutySeconds, not even an explicit nil
### GetDurationDrivingSeconds

`func (o *BaseHosDailySnapshot) GetDurationDrivingSeconds() int32`

GetDurationDrivingSeconds returns the DurationDrivingSeconds field if non-nil, zero value otherwise.

### GetDurationDrivingSecondsOk

`func (o *BaseHosDailySnapshot) GetDurationDrivingSecondsOk() (*int32, bool)`

GetDurationDrivingSecondsOk returns a tuple with the DurationDrivingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationDrivingSeconds

`func (o *BaseHosDailySnapshot) SetDurationDrivingSeconds(v int32)`

SetDurationDrivingSeconds sets DurationDrivingSeconds field to given value.

### HasDurationDrivingSeconds

`func (o *BaseHosDailySnapshot) HasDurationDrivingSeconds() bool`

HasDurationDrivingSeconds returns a boolean if a field has been set.

### SetDurationDrivingSecondsNil

`func (o *BaseHosDailySnapshot) SetDurationDrivingSecondsNil(b bool)`

 SetDurationDrivingSecondsNil sets the value for DurationDrivingSeconds to be an explicit nil

### UnsetDurationDrivingSeconds
`func (o *BaseHosDailySnapshot) UnsetDurationDrivingSeconds()`

UnsetDurationDrivingSeconds ensures that no value is present for DurationDrivingSeconds, not even an explicit nil
### GetDurationPersonalConveyanceSeconds

`func (o *BaseHosDailySnapshot) GetDurationPersonalConveyanceSeconds() int32`

GetDurationPersonalConveyanceSeconds returns the DurationPersonalConveyanceSeconds field if non-nil, zero value otherwise.

### GetDurationPersonalConveyanceSecondsOk

`func (o *BaseHosDailySnapshot) GetDurationPersonalConveyanceSecondsOk() (*int32, bool)`

GetDurationPersonalConveyanceSecondsOk returns a tuple with the DurationPersonalConveyanceSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationPersonalConveyanceSeconds

`func (o *BaseHosDailySnapshot) SetDurationPersonalConveyanceSeconds(v int32)`

SetDurationPersonalConveyanceSeconds sets DurationPersonalConveyanceSeconds field to given value.

### HasDurationPersonalConveyanceSeconds

`func (o *BaseHosDailySnapshot) HasDurationPersonalConveyanceSeconds() bool`

HasDurationPersonalConveyanceSeconds returns a boolean if a field has been set.

### SetDurationPersonalConveyanceSecondsNil

`func (o *BaseHosDailySnapshot) SetDurationPersonalConveyanceSecondsNil(b bool)`

 SetDurationPersonalConveyanceSecondsNil sets the value for DurationPersonalConveyanceSeconds to be an explicit nil

### UnsetDurationPersonalConveyanceSeconds
`func (o *BaseHosDailySnapshot) UnsetDurationPersonalConveyanceSeconds()`

UnsetDurationPersonalConveyanceSeconds ensures that no value is present for DurationPersonalConveyanceSeconds, not even an explicit nil
### GetDurationYardMoveSeconds

`func (o *BaseHosDailySnapshot) GetDurationYardMoveSeconds() int32`

GetDurationYardMoveSeconds returns the DurationYardMoveSeconds field if non-nil, zero value otherwise.

### GetDurationYardMoveSecondsOk

`func (o *BaseHosDailySnapshot) GetDurationYardMoveSecondsOk() (*int32, bool)`

GetDurationYardMoveSecondsOk returns a tuple with the DurationYardMoveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationYardMoveSeconds

`func (o *BaseHosDailySnapshot) SetDurationYardMoveSeconds(v int32)`

SetDurationYardMoveSeconds sets DurationYardMoveSeconds field to given value.

### HasDurationYardMoveSeconds

`func (o *BaseHosDailySnapshot) HasDurationYardMoveSeconds() bool`

HasDurationYardMoveSeconds returns a boolean if a field has been set.

### SetDurationYardMoveSecondsNil

`func (o *BaseHosDailySnapshot) SetDurationYardMoveSecondsNil(b bool)`

 SetDurationYardMoveSecondsNil sets the value for DurationYardMoveSeconds to be an explicit nil

### UnsetDurationYardMoveSeconds
`func (o *BaseHosDailySnapshot) UnsetDurationYardMoveSeconds()`

UnsetDurationYardMoveSeconds ensures that no value is present for DurationYardMoveSeconds, not even an explicit nil
### GetDurationWaitingSeconds

`func (o *BaseHosDailySnapshot) GetDurationWaitingSeconds() int32`

GetDurationWaitingSeconds returns the DurationWaitingSeconds field if non-nil, zero value otherwise.

### GetDurationWaitingSecondsOk

`func (o *BaseHosDailySnapshot) GetDurationWaitingSecondsOk() (*int32, bool)`

GetDurationWaitingSecondsOk returns a tuple with the DurationWaitingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationWaitingSeconds

`func (o *BaseHosDailySnapshot) SetDurationWaitingSeconds(v int32)`

SetDurationWaitingSeconds sets DurationWaitingSeconds field to given value.

### HasDurationWaitingSeconds

`func (o *BaseHosDailySnapshot) HasDurationWaitingSeconds() bool`

HasDurationWaitingSeconds returns a boolean if a field has been set.

### SetDurationWaitingSecondsNil

`func (o *BaseHosDailySnapshot) SetDurationWaitingSecondsNil(b bool)`

 SetDurationWaitingSecondsNil sets the value for DurationWaitingSeconds to be an explicit nil

### UnsetDurationWaitingSeconds
`func (o *BaseHosDailySnapshot) UnsetDurationWaitingSeconds()`

UnsetDurationWaitingSeconds ensures that no value is present for DurationWaitingSeconds, not even an explicit nil
### GetDurationUnknownSeconds

`func (o *BaseHosDailySnapshot) GetDurationUnknownSeconds() int32`

GetDurationUnknownSeconds returns the DurationUnknownSeconds field if non-nil, zero value otherwise.

### GetDurationUnknownSecondsOk

`func (o *BaseHosDailySnapshot) GetDurationUnknownSecondsOk() (*int32, bool)`

GetDurationUnknownSecondsOk returns a tuple with the DurationUnknownSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationUnknownSeconds

`func (o *BaseHosDailySnapshot) SetDurationUnknownSeconds(v int32)`

SetDurationUnknownSeconds sets DurationUnknownSeconds field to given value.

### HasDurationUnknownSeconds

`func (o *BaseHosDailySnapshot) HasDurationUnknownSeconds() bool`

HasDurationUnknownSeconds returns a boolean if a field has been set.

### SetDurationUnknownSecondsNil

`func (o *BaseHosDailySnapshot) SetDurationUnknownSecondsNil(b bool)`

 SetDurationUnknownSecondsNil sets the value for DurationUnknownSeconds to be an explicit nil

### UnsetDurationUnknownSeconds
`func (o *BaseHosDailySnapshot) UnsetDurationUnknownSeconds()`

UnsetDurationUnknownSeconds ensures that no value is present for DurationUnknownSeconds, not even an explicit nil
### GetDurationCycleOnDutySeconds

`func (o *BaseHosDailySnapshot) GetDurationCycleOnDutySeconds() int32`

GetDurationCycleOnDutySeconds returns the DurationCycleOnDutySeconds field if non-nil, zero value otherwise.

### GetDurationCycleOnDutySecondsOk

`func (o *BaseHosDailySnapshot) GetDurationCycleOnDutySecondsOk() (*int32, bool)`

GetDurationCycleOnDutySecondsOk returns a tuple with the DurationCycleOnDutySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationCycleOnDutySeconds

`func (o *BaseHosDailySnapshot) SetDurationCycleOnDutySeconds(v int32)`

SetDurationCycleOnDutySeconds sets DurationCycleOnDutySeconds field to given value.

### HasDurationCycleOnDutySeconds

`func (o *BaseHosDailySnapshot) HasDurationCycleOnDutySeconds() bool`

HasDurationCycleOnDutySeconds returns a boolean if a field has been set.

### SetDurationCycleOnDutySecondsNil

`func (o *BaseHosDailySnapshot) SetDurationCycleOnDutySecondsNil(b bool)`

 SetDurationCycleOnDutySecondsNil sets the value for DurationCycleOnDutySeconds to be an explicit nil

### UnsetDurationCycleOnDutySeconds
`func (o *BaseHosDailySnapshot) UnsetDurationCycleOnDutySeconds()`

UnsetDurationCycleOnDutySeconds ensures that no value is present for DurationCycleOnDutySeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


