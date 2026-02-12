# HosAvailabilityRead

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
**VehicleId** | Pointer to **NullableString** |  | [optional] 
**HosRulesetCode** | Pointer to [**NullableHosRulesetCodeEnum**](HosRulesetCodeEnum.md) |  | [optional] 
**DutyStatusCode** | Pointer to [**NullableDutyStatusCodeEnum**](DutyStatusCodeEnum.md) |  | [optional] 
**CycleStartedAt** | Pointer to **NullableTime** |  | [optional] 
**CycleEndsAt** | Pointer to **NullableTime** |  | [optional] 
**AvailableDriveSeconds** | Pointer to **NullableInt32** |  | [optional] 
**AvailableShiftSeconds** | Pointer to **NullableInt32** |  | [optional] 
**AvailableCycleSeconds** | Pointer to **NullableInt32** |  | [optional] 
**AvailableTomorrowSeconds** | Pointer to **NullableInt32** |  | [optional] 
**AvailableDay2Seconds** | Pointer to **NullableInt32** |  | [optional] 
**AvailableDay3Seconds** | Pointer to **NullableInt32** |  | [optional] 
**ForecastHorizonDays** | Pointer to **NullableInt32** |  | [optional] 
**TimeUntilBreakSeconds** | Pointer to **NullableInt32** |  | [optional] 
**RestRemainingSeconds** | Pointer to **NullableInt32** |  | [optional] 
**CycleViolationDurationSeconds** | Pointer to **NullableInt32** |  | [optional] 
**ShiftDrivingViolationDurationSeconds** | Pointer to **NullableInt32** |  | [optional] 
**ShiftEndsAt** | Pointer to **NullableTime** |  | [optional] 
**NextBreakDueAt** | Pointer to **NullableTime** |  | [optional] 
**Next10hrResetEligibleAt** | Pointer to **NullableTime** |  | [optional] 
**Next34hrResetEligibleAt** | Pointer to **NullableTime** |  | [optional] 
**IsPersonalConveyanceApplied** | Pointer to **NullableBool** |  | [optional] 
**IsYardMoveApplied** | Pointer to **NullableBool** |  | [optional] 
**IsAdverseDrivingExemptionAvailable** | Pointer to **NullableBool** |  | [optional] 
**IsAdverseDrivingApplied** | Pointer to **NullableBool** |  | [optional] 
**IsMealBreakRequired** | Pointer to **NullableBool** |  | [optional] 
**MealBreakDueAt** | Pointer to **NullableTime** |  | [optional] 
**MealBreakMinDurationSeconds** | Pointer to **NullableInt32** |  | [optional] 
**MealBreakTimeUntilDueSeconds** | Pointer to **NullableInt32** |  | [optional] 
**IsSplitSleepApplied** | Pointer to **NullableBool** |  | [optional] 
**IsSleeperEligible** | Pointer to **NullableBool** |  | [optional] 
**SleeperRequiredRemainingSeconds** | Pointer to **NullableInt32** |  | [optional] 
**SleeperSplitWindowEndsAt** | Pointer to **NullableTime** |  | [optional] 
**IsCycleApplicable** | Pointer to **NullableBool** |  | [optional] 
**ExceptionCodes** | Pointer to **[]string** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewHosAvailabilityRead

`func NewHosAvailabilityRead(fleetId NullableString, id string, createdAt time.Time, updatedAt time.Time, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *HosAvailabilityRead`

NewHosAvailabilityRead instantiates a new HosAvailabilityRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHosAvailabilityReadWithDefaults

`func NewHosAvailabilityReadWithDefaults() *HosAvailabilityRead`

NewHosAvailabilityReadWithDefaults instantiates a new HosAvailabilityRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *HosAvailabilityRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *HosAvailabilityRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *HosAvailabilityRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### SetFleetIdNil

`func (o *HosAvailabilityRead) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *HosAvailabilityRead) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetFleetRef

`func (o *HosAvailabilityRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *HosAvailabilityRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *HosAvailabilityRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *HosAvailabilityRead) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *HosAvailabilityRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *HosAvailabilityRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetId

`func (o *HosAvailabilityRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HosAvailabilityRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HosAvailabilityRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *HosAvailabilityRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HosAvailabilityRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HosAvailabilityRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *HosAvailabilityRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HosAvailabilityRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HosAvailabilityRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *HosAvailabilityRead) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *HosAvailabilityRead) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *HosAvailabilityRead) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *HosAvailabilityRead) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *HosAvailabilityRead) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *HosAvailabilityRead) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetConnectionId

`func (o *HosAvailabilityRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *HosAvailabilityRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *HosAvailabilityRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceName

`func (o *HosAvailabilityRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *HosAvailabilityRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *HosAvailabilityRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *HosAvailabilityRead) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *HosAvailabilityRead) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *HosAvailabilityRead) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *HosAvailabilityRead) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *HosAvailabilityRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *HosAvailabilityRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *HosAvailabilityRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *HosAvailabilityRead) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *HosAvailabilityRead) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *HosAvailabilityRead) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *HosAvailabilityRead) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *HosAvailabilityRead) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *HosAvailabilityRead) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *HosAvailabilityRead) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *HosAvailabilityRead) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *HosAvailabilityRead) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *HosAvailabilityRead) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HosAvailabilityRead) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HosAvailabilityRead) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HosAvailabilityRead) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HosAvailabilityRead) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HosAvailabilityRead) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *HosAvailabilityRead) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *HosAvailabilityRead) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *HosAvailabilityRead) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *HosAvailabilityRead) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *HosAvailabilityRead) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *HosAvailabilityRead) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetDriverId

`func (o *HosAvailabilityRead) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *HosAvailabilityRead) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *HosAvailabilityRead) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *HosAvailabilityRead) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *HosAvailabilityRead) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *HosAvailabilityRead) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetVehicleId

`func (o *HosAvailabilityRead) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *HosAvailabilityRead) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *HosAvailabilityRead) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *HosAvailabilityRead) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *HosAvailabilityRead) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *HosAvailabilityRead) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetHosRulesetCode

`func (o *HosAvailabilityRead) GetHosRulesetCode() HosRulesetCodeEnum`

GetHosRulesetCode returns the HosRulesetCode field if non-nil, zero value otherwise.

### GetHosRulesetCodeOk

`func (o *HosAvailabilityRead) GetHosRulesetCodeOk() (*HosRulesetCodeEnum, bool)`

GetHosRulesetCodeOk returns a tuple with the HosRulesetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosRulesetCode

`func (o *HosAvailabilityRead) SetHosRulesetCode(v HosRulesetCodeEnum)`

SetHosRulesetCode sets HosRulesetCode field to given value.

### HasHosRulesetCode

`func (o *HosAvailabilityRead) HasHosRulesetCode() bool`

HasHosRulesetCode returns a boolean if a field has been set.

### SetHosRulesetCodeNil

`func (o *HosAvailabilityRead) SetHosRulesetCodeNil(b bool)`

 SetHosRulesetCodeNil sets the value for HosRulesetCode to be an explicit nil

### UnsetHosRulesetCode
`func (o *HosAvailabilityRead) UnsetHosRulesetCode()`

UnsetHosRulesetCode ensures that no value is present for HosRulesetCode, not even an explicit nil
### GetDutyStatusCode

`func (o *HosAvailabilityRead) GetDutyStatusCode() DutyStatusCodeEnum`

GetDutyStatusCode returns the DutyStatusCode field if non-nil, zero value otherwise.

### GetDutyStatusCodeOk

`func (o *HosAvailabilityRead) GetDutyStatusCodeOk() (*DutyStatusCodeEnum, bool)`

GetDutyStatusCodeOk returns a tuple with the DutyStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyStatusCode

`func (o *HosAvailabilityRead) SetDutyStatusCode(v DutyStatusCodeEnum)`

SetDutyStatusCode sets DutyStatusCode field to given value.

### HasDutyStatusCode

`func (o *HosAvailabilityRead) HasDutyStatusCode() bool`

HasDutyStatusCode returns a boolean if a field has been set.

### SetDutyStatusCodeNil

`func (o *HosAvailabilityRead) SetDutyStatusCodeNil(b bool)`

 SetDutyStatusCodeNil sets the value for DutyStatusCode to be an explicit nil

### UnsetDutyStatusCode
`func (o *HosAvailabilityRead) UnsetDutyStatusCode()`

UnsetDutyStatusCode ensures that no value is present for DutyStatusCode, not even an explicit nil
### GetCycleStartedAt

`func (o *HosAvailabilityRead) GetCycleStartedAt() time.Time`

GetCycleStartedAt returns the CycleStartedAt field if non-nil, zero value otherwise.

### GetCycleStartedAtOk

`func (o *HosAvailabilityRead) GetCycleStartedAtOk() (*time.Time, bool)`

GetCycleStartedAtOk returns a tuple with the CycleStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleStartedAt

`func (o *HosAvailabilityRead) SetCycleStartedAt(v time.Time)`

SetCycleStartedAt sets CycleStartedAt field to given value.

### HasCycleStartedAt

`func (o *HosAvailabilityRead) HasCycleStartedAt() bool`

HasCycleStartedAt returns a boolean if a field has been set.

### SetCycleStartedAtNil

`func (o *HosAvailabilityRead) SetCycleStartedAtNil(b bool)`

 SetCycleStartedAtNil sets the value for CycleStartedAt to be an explicit nil

### UnsetCycleStartedAt
`func (o *HosAvailabilityRead) UnsetCycleStartedAt()`

UnsetCycleStartedAt ensures that no value is present for CycleStartedAt, not even an explicit nil
### GetCycleEndsAt

`func (o *HosAvailabilityRead) GetCycleEndsAt() time.Time`

GetCycleEndsAt returns the CycleEndsAt field if non-nil, zero value otherwise.

### GetCycleEndsAtOk

`func (o *HosAvailabilityRead) GetCycleEndsAtOk() (*time.Time, bool)`

GetCycleEndsAtOk returns a tuple with the CycleEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleEndsAt

`func (o *HosAvailabilityRead) SetCycleEndsAt(v time.Time)`

SetCycleEndsAt sets CycleEndsAt field to given value.

### HasCycleEndsAt

`func (o *HosAvailabilityRead) HasCycleEndsAt() bool`

HasCycleEndsAt returns a boolean if a field has been set.

### SetCycleEndsAtNil

`func (o *HosAvailabilityRead) SetCycleEndsAtNil(b bool)`

 SetCycleEndsAtNil sets the value for CycleEndsAt to be an explicit nil

### UnsetCycleEndsAt
`func (o *HosAvailabilityRead) UnsetCycleEndsAt()`

UnsetCycleEndsAt ensures that no value is present for CycleEndsAt, not even an explicit nil
### GetAvailableDriveSeconds

`func (o *HosAvailabilityRead) GetAvailableDriveSeconds() int32`

GetAvailableDriveSeconds returns the AvailableDriveSeconds field if non-nil, zero value otherwise.

### GetAvailableDriveSecondsOk

`func (o *HosAvailabilityRead) GetAvailableDriveSecondsOk() (*int32, bool)`

GetAvailableDriveSecondsOk returns a tuple with the AvailableDriveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDriveSeconds

`func (o *HosAvailabilityRead) SetAvailableDriveSeconds(v int32)`

SetAvailableDriveSeconds sets AvailableDriveSeconds field to given value.

### HasAvailableDriveSeconds

`func (o *HosAvailabilityRead) HasAvailableDriveSeconds() bool`

HasAvailableDriveSeconds returns a boolean if a field has been set.

### SetAvailableDriveSecondsNil

`func (o *HosAvailabilityRead) SetAvailableDriveSecondsNil(b bool)`

 SetAvailableDriveSecondsNil sets the value for AvailableDriveSeconds to be an explicit nil

### UnsetAvailableDriveSeconds
`func (o *HosAvailabilityRead) UnsetAvailableDriveSeconds()`

UnsetAvailableDriveSeconds ensures that no value is present for AvailableDriveSeconds, not even an explicit nil
### GetAvailableShiftSeconds

`func (o *HosAvailabilityRead) GetAvailableShiftSeconds() int32`

GetAvailableShiftSeconds returns the AvailableShiftSeconds field if non-nil, zero value otherwise.

### GetAvailableShiftSecondsOk

`func (o *HosAvailabilityRead) GetAvailableShiftSecondsOk() (*int32, bool)`

GetAvailableShiftSecondsOk returns a tuple with the AvailableShiftSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShiftSeconds

`func (o *HosAvailabilityRead) SetAvailableShiftSeconds(v int32)`

SetAvailableShiftSeconds sets AvailableShiftSeconds field to given value.

### HasAvailableShiftSeconds

`func (o *HosAvailabilityRead) HasAvailableShiftSeconds() bool`

HasAvailableShiftSeconds returns a boolean if a field has been set.

### SetAvailableShiftSecondsNil

`func (o *HosAvailabilityRead) SetAvailableShiftSecondsNil(b bool)`

 SetAvailableShiftSecondsNil sets the value for AvailableShiftSeconds to be an explicit nil

### UnsetAvailableShiftSeconds
`func (o *HosAvailabilityRead) UnsetAvailableShiftSeconds()`

UnsetAvailableShiftSeconds ensures that no value is present for AvailableShiftSeconds, not even an explicit nil
### GetAvailableCycleSeconds

`func (o *HosAvailabilityRead) GetAvailableCycleSeconds() int32`

GetAvailableCycleSeconds returns the AvailableCycleSeconds field if non-nil, zero value otherwise.

### GetAvailableCycleSecondsOk

`func (o *HosAvailabilityRead) GetAvailableCycleSecondsOk() (*int32, bool)`

GetAvailableCycleSecondsOk returns a tuple with the AvailableCycleSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCycleSeconds

`func (o *HosAvailabilityRead) SetAvailableCycleSeconds(v int32)`

SetAvailableCycleSeconds sets AvailableCycleSeconds field to given value.

### HasAvailableCycleSeconds

`func (o *HosAvailabilityRead) HasAvailableCycleSeconds() bool`

HasAvailableCycleSeconds returns a boolean if a field has been set.

### SetAvailableCycleSecondsNil

`func (o *HosAvailabilityRead) SetAvailableCycleSecondsNil(b bool)`

 SetAvailableCycleSecondsNil sets the value for AvailableCycleSeconds to be an explicit nil

### UnsetAvailableCycleSeconds
`func (o *HosAvailabilityRead) UnsetAvailableCycleSeconds()`

UnsetAvailableCycleSeconds ensures that no value is present for AvailableCycleSeconds, not even an explicit nil
### GetAvailableTomorrowSeconds

`func (o *HosAvailabilityRead) GetAvailableTomorrowSeconds() int32`

GetAvailableTomorrowSeconds returns the AvailableTomorrowSeconds field if non-nil, zero value otherwise.

### GetAvailableTomorrowSecondsOk

`func (o *HosAvailabilityRead) GetAvailableTomorrowSecondsOk() (*int32, bool)`

GetAvailableTomorrowSecondsOk returns a tuple with the AvailableTomorrowSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTomorrowSeconds

`func (o *HosAvailabilityRead) SetAvailableTomorrowSeconds(v int32)`

SetAvailableTomorrowSeconds sets AvailableTomorrowSeconds field to given value.

### HasAvailableTomorrowSeconds

`func (o *HosAvailabilityRead) HasAvailableTomorrowSeconds() bool`

HasAvailableTomorrowSeconds returns a boolean if a field has been set.

### SetAvailableTomorrowSecondsNil

`func (o *HosAvailabilityRead) SetAvailableTomorrowSecondsNil(b bool)`

 SetAvailableTomorrowSecondsNil sets the value for AvailableTomorrowSeconds to be an explicit nil

### UnsetAvailableTomorrowSeconds
`func (o *HosAvailabilityRead) UnsetAvailableTomorrowSeconds()`

UnsetAvailableTomorrowSeconds ensures that no value is present for AvailableTomorrowSeconds, not even an explicit nil
### GetAvailableDay2Seconds

`func (o *HosAvailabilityRead) GetAvailableDay2Seconds() int32`

GetAvailableDay2Seconds returns the AvailableDay2Seconds field if non-nil, zero value otherwise.

### GetAvailableDay2SecondsOk

`func (o *HosAvailabilityRead) GetAvailableDay2SecondsOk() (*int32, bool)`

GetAvailableDay2SecondsOk returns a tuple with the AvailableDay2Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDay2Seconds

`func (o *HosAvailabilityRead) SetAvailableDay2Seconds(v int32)`

SetAvailableDay2Seconds sets AvailableDay2Seconds field to given value.

### HasAvailableDay2Seconds

`func (o *HosAvailabilityRead) HasAvailableDay2Seconds() bool`

HasAvailableDay2Seconds returns a boolean if a field has been set.

### SetAvailableDay2SecondsNil

`func (o *HosAvailabilityRead) SetAvailableDay2SecondsNil(b bool)`

 SetAvailableDay2SecondsNil sets the value for AvailableDay2Seconds to be an explicit nil

### UnsetAvailableDay2Seconds
`func (o *HosAvailabilityRead) UnsetAvailableDay2Seconds()`

UnsetAvailableDay2Seconds ensures that no value is present for AvailableDay2Seconds, not even an explicit nil
### GetAvailableDay3Seconds

`func (o *HosAvailabilityRead) GetAvailableDay3Seconds() int32`

GetAvailableDay3Seconds returns the AvailableDay3Seconds field if non-nil, zero value otherwise.

### GetAvailableDay3SecondsOk

`func (o *HosAvailabilityRead) GetAvailableDay3SecondsOk() (*int32, bool)`

GetAvailableDay3SecondsOk returns a tuple with the AvailableDay3Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDay3Seconds

`func (o *HosAvailabilityRead) SetAvailableDay3Seconds(v int32)`

SetAvailableDay3Seconds sets AvailableDay3Seconds field to given value.

### HasAvailableDay3Seconds

`func (o *HosAvailabilityRead) HasAvailableDay3Seconds() bool`

HasAvailableDay3Seconds returns a boolean if a field has been set.

### SetAvailableDay3SecondsNil

`func (o *HosAvailabilityRead) SetAvailableDay3SecondsNil(b bool)`

 SetAvailableDay3SecondsNil sets the value for AvailableDay3Seconds to be an explicit nil

### UnsetAvailableDay3Seconds
`func (o *HosAvailabilityRead) UnsetAvailableDay3Seconds()`

UnsetAvailableDay3Seconds ensures that no value is present for AvailableDay3Seconds, not even an explicit nil
### GetForecastHorizonDays

`func (o *HosAvailabilityRead) GetForecastHorizonDays() int32`

GetForecastHorizonDays returns the ForecastHorizonDays field if non-nil, zero value otherwise.

### GetForecastHorizonDaysOk

`func (o *HosAvailabilityRead) GetForecastHorizonDaysOk() (*int32, bool)`

GetForecastHorizonDaysOk returns a tuple with the ForecastHorizonDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastHorizonDays

`func (o *HosAvailabilityRead) SetForecastHorizonDays(v int32)`

SetForecastHorizonDays sets ForecastHorizonDays field to given value.

### HasForecastHorizonDays

`func (o *HosAvailabilityRead) HasForecastHorizonDays() bool`

HasForecastHorizonDays returns a boolean if a field has been set.

### SetForecastHorizonDaysNil

`func (o *HosAvailabilityRead) SetForecastHorizonDaysNil(b bool)`

 SetForecastHorizonDaysNil sets the value for ForecastHorizonDays to be an explicit nil

### UnsetForecastHorizonDays
`func (o *HosAvailabilityRead) UnsetForecastHorizonDays()`

UnsetForecastHorizonDays ensures that no value is present for ForecastHorizonDays, not even an explicit nil
### GetTimeUntilBreakSeconds

`func (o *HosAvailabilityRead) GetTimeUntilBreakSeconds() int32`

GetTimeUntilBreakSeconds returns the TimeUntilBreakSeconds field if non-nil, zero value otherwise.

### GetTimeUntilBreakSecondsOk

`func (o *HosAvailabilityRead) GetTimeUntilBreakSecondsOk() (*int32, bool)`

GetTimeUntilBreakSecondsOk returns a tuple with the TimeUntilBreakSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUntilBreakSeconds

`func (o *HosAvailabilityRead) SetTimeUntilBreakSeconds(v int32)`

SetTimeUntilBreakSeconds sets TimeUntilBreakSeconds field to given value.

### HasTimeUntilBreakSeconds

`func (o *HosAvailabilityRead) HasTimeUntilBreakSeconds() bool`

HasTimeUntilBreakSeconds returns a boolean if a field has been set.

### SetTimeUntilBreakSecondsNil

`func (o *HosAvailabilityRead) SetTimeUntilBreakSecondsNil(b bool)`

 SetTimeUntilBreakSecondsNil sets the value for TimeUntilBreakSeconds to be an explicit nil

### UnsetTimeUntilBreakSeconds
`func (o *HosAvailabilityRead) UnsetTimeUntilBreakSeconds()`

UnsetTimeUntilBreakSeconds ensures that no value is present for TimeUntilBreakSeconds, not even an explicit nil
### GetRestRemainingSeconds

`func (o *HosAvailabilityRead) GetRestRemainingSeconds() int32`

GetRestRemainingSeconds returns the RestRemainingSeconds field if non-nil, zero value otherwise.

### GetRestRemainingSecondsOk

`func (o *HosAvailabilityRead) GetRestRemainingSecondsOk() (*int32, bool)`

GetRestRemainingSecondsOk returns a tuple with the RestRemainingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestRemainingSeconds

`func (o *HosAvailabilityRead) SetRestRemainingSeconds(v int32)`

SetRestRemainingSeconds sets RestRemainingSeconds field to given value.

### HasRestRemainingSeconds

`func (o *HosAvailabilityRead) HasRestRemainingSeconds() bool`

HasRestRemainingSeconds returns a boolean if a field has been set.

### SetRestRemainingSecondsNil

`func (o *HosAvailabilityRead) SetRestRemainingSecondsNil(b bool)`

 SetRestRemainingSecondsNil sets the value for RestRemainingSeconds to be an explicit nil

### UnsetRestRemainingSeconds
`func (o *HosAvailabilityRead) UnsetRestRemainingSeconds()`

UnsetRestRemainingSeconds ensures that no value is present for RestRemainingSeconds, not even an explicit nil
### GetCycleViolationDurationSeconds

`func (o *HosAvailabilityRead) GetCycleViolationDurationSeconds() int32`

GetCycleViolationDurationSeconds returns the CycleViolationDurationSeconds field if non-nil, zero value otherwise.

### GetCycleViolationDurationSecondsOk

`func (o *HosAvailabilityRead) GetCycleViolationDurationSecondsOk() (*int32, bool)`

GetCycleViolationDurationSecondsOk returns a tuple with the CycleViolationDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleViolationDurationSeconds

`func (o *HosAvailabilityRead) SetCycleViolationDurationSeconds(v int32)`

SetCycleViolationDurationSeconds sets CycleViolationDurationSeconds field to given value.

### HasCycleViolationDurationSeconds

`func (o *HosAvailabilityRead) HasCycleViolationDurationSeconds() bool`

HasCycleViolationDurationSeconds returns a boolean if a field has been set.

### SetCycleViolationDurationSecondsNil

`func (o *HosAvailabilityRead) SetCycleViolationDurationSecondsNil(b bool)`

 SetCycleViolationDurationSecondsNil sets the value for CycleViolationDurationSeconds to be an explicit nil

### UnsetCycleViolationDurationSeconds
`func (o *HosAvailabilityRead) UnsetCycleViolationDurationSeconds()`

UnsetCycleViolationDurationSeconds ensures that no value is present for CycleViolationDurationSeconds, not even an explicit nil
### GetShiftDrivingViolationDurationSeconds

`func (o *HosAvailabilityRead) GetShiftDrivingViolationDurationSeconds() int32`

GetShiftDrivingViolationDurationSeconds returns the ShiftDrivingViolationDurationSeconds field if non-nil, zero value otherwise.

### GetShiftDrivingViolationDurationSecondsOk

`func (o *HosAvailabilityRead) GetShiftDrivingViolationDurationSecondsOk() (*int32, bool)`

GetShiftDrivingViolationDurationSecondsOk returns a tuple with the ShiftDrivingViolationDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftDrivingViolationDurationSeconds

`func (o *HosAvailabilityRead) SetShiftDrivingViolationDurationSeconds(v int32)`

SetShiftDrivingViolationDurationSeconds sets ShiftDrivingViolationDurationSeconds field to given value.

### HasShiftDrivingViolationDurationSeconds

`func (o *HosAvailabilityRead) HasShiftDrivingViolationDurationSeconds() bool`

HasShiftDrivingViolationDurationSeconds returns a boolean if a field has been set.

### SetShiftDrivingViolationDurationSecondsNil

`func (o *HosAvailabilityRead) SetShiftDrivingViolationDurationSecondsNil(b bool)`

 SetShiftDrivingViolationDurationSecondsNil sets the value for ShiftDrivingViolationDurationSeconds to be an explicit nil

### UnsetShiftDrivingViolationDurationSeconds
`func (o *HosAvailabilityRead) UnsetShiftDrivingViolationDurationSeconds()`

UnsetShiftDrivingViolationDurationSeconds ensures that no value is present for ShiftDrivingViolationDurationSeconds, not even an explicit nil
### GetShiftEndsAt

`func (o *HosAvailabilityRead) GetShiftEndsAt() time.Time`

GetShiftEndsAt returns the ShiftEndsAt field if non-nil, zero value otherwise.

### GetShiftEndsAtOk

`func (o *HosAvailabilityRead) GetShiftEndsAtOk() (*time.Time, bool)`

GetShiftEndsAtOk returns a tuple with the ShiftEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftEndsAt

`func (o *HosAvailabilityRead) SetShiftEndsAt(v time.Time)`

SetShiftEndsAt sets ShiftEndsAt field to given value.

### HasShiftEndsAt

`func (o *HosAvailabilityRead) HasShiftEndsAt() bool`

HasShiftEndsAt returns a boolean if a field has been set.

### SetShiftEndsAtNil

`func (o *HosAvailabilityRead) SetShiftEndsAtNil(b bool)`

 SetShiftEndsAtNil sets the value for ShiftEndsAt to be an explicit nil

### UnsetShiftEndsAt
`func (o *HosAvailabilityRead) UnsetShiftEndsAt()`

UnsetShiftEndsAt ensures that no value is present for ShiftEndsAt, not even an explicit nil
### GetNextBreakDueAt

`func (o *HosAvailabilityRead) GetNextBreakDueAt() time.Time`

GetNextBreakDueAt returns the NextBreakDueAt field if non-nil, zero value otherwise.

### GetNextBreakDueAtOk

`func (o *HosAvailabilityRead) GetNextBreakDueAtOk() (*time.Time, bool)`

GetNextBreakDueAtOk returns a tuple with the NextBreakDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBreakDueAt

`func (o *HosAvailabilityRead) SetNextBreakDueAt(v time.Time)`

SetNextBreakDueAt sets NextBreakDueAt field to given value.

### HasNextBreakDueAt

`func (o *HosAvailabilityRead) HasNextBreakDueAt() bool`

HasNextBreakDueAt returns a boolean if a field has been set.

### SetNextBreakDueAtNil

`func (o *HosAvailabilityRead) SetNextBreakDueAtNil(b bool)`

 SetNextBreakDueAtNil sets the value for NextBreakDueAt to be an explicit nil

### UnsetNextBreakDueAt
`func (o *HosAvailabilityRead) UnsetNextBreakDueAt()`

UnsetNextBreakDueAt ensures that no value is present for NextBreakDueAt, not even an explicit nil
### GetNext10hrResetEligibleAt

`func (o *HosAvailabilityRead) GetNext10hrResetEligibleAt() time.Time`

GetNext10hrResetEligibleAt returns the Next10hrResetEligibleAt field if non-nil, zero value otherwise.

### GetNext10hrResetEligibleAtOk

`func (o *HosAvailabilityRead) GetNext10hrResetEligibleAtOk() (*time.Time, bool)`

GetNext10hrResetEligibleAtOk returns a tuple with the Next10hrResetEligibleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext10hrResetEligibleAt

`func (o *HosAvailabilityRead) SetNext10hrResetEligibleAt(v time.Time)`

SetNext10hrResetEligibleAt sets Next10hrResetEligibleAt field to given value.

### HasNext10hrResetEligibleAt

`func (o *HosAvailabilityRead) HasNext10hrResetEligibleAt() bool`

HasNext10hrResetEligibleAt returns a boolean if a field has been set.

### SetNext10hrResetEligibleAtNil

`func (o *HosAvailabilityRead) SetNext10hrResetEligibleAtNil(b bool)`

 SetNext10hrResetEligibleAtNil sets the value for Next10hrResetEligibleAt to be an explicit nil

### UnsetNext10hrResetEligibleAt
`func (o *HosAvailabilityRead) UnsetNext10hrResetEligibleAt()`

UnsetNext10hrResetEligibleAt ensures that no value is present for Next10hrResetEligibleAt, not even an explicit nil
### GetNext34hrResetEligibleAt

`func (o *HosAvailabilityRead) GetNext34hrResetEligibleAt() time.Time`

GetNext34hrResetEligibleAt returns the Next34hrResetEligibleAt field if non-nil, zero value otherwise.

### GetNext34hrResetEligibleAtOk

`func (o *HosAvailabilityRead) GetNext34hrResetEligibleAtOk() (*time.Time, bool)`

GetNext34hrResetEligibleAtOk returns a tuple with the Next34hrResetEligibleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext34hrResetEligibleAt

`func (o *HosAvailabilityRead) SetNext34hrResetEligibleAt(v time.Time)`

SetNext34hrResetEligibleAt sets Next34hrResetEligibleAt field to given value.

### HasNext34hrResetEligibleAt

`func (o *HosAvailabilityRead) HasNext34hrResetEligibleAt() bool`

HasNext34hrResetEligibleAt returns a boolean if a field has been set.

### SetNext34hrResetEligibleAtNil

`func (o *HosAvailabilityRead) SetNext34hrResetEligibleAtNil(b bool)`

 SetNext34hrResetEligibleAtNil sets the value for Next34hrResetEligibleAt to be an explicit nil

### UnsetNext34hrResetEligibleAt
`func (o *HosAvailabilityRead) UnsetNext34hrResetEligibleAt()`

UnsetNext34hrResetEligibleAt ensures that no value is present for Next34hrResetEligibleAt, not even an explicit nil
### GetIsPersonalConveyanceApplied

`func (o *HosAvailabilityRead) GetIsPersonalConveyanceApplied() bool`

GetIsPersonalConveyanceApplied returns the IsPersonalConveyanceApplied field if non-nil, zero value otherwise.

### GetIsPersonalConveyanceAppliedOk

`func (o *HosAvailabilityRead) GetIsPersonalConveyanceAppliedOk() (*bool, bool)`

GetIsPersonalConveyanceAppliedOk returns a tuple with the IsPersonalConveyanceApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersonalConveyanceApplied

`func (o *HosAvailabilityRead) SetIsPersonalConveyanceApplied(v bool)`

SetIsPersonalConveyanceApplied sets IsPersonalConveyanceApplied field to given value.

### HasIsPersonalConveyanceApplied

`func (o *HosAvailabilityRead) HasIsPersonalConveyanceApplied() bool`

HasIsPersonalConveyanceApplied returns a boolean if a field has been set.

### SetIsPersonalConveyanceAppliedNil

`func (o *HosAvailabilityRead) SetIsPersonalConveyanceAppliedNil(b bool)`

 SetIsPersonalConveyanceAppliedNil sets the value for IsPersonalConveyanceApplied to be an explicit nil

### UnsetIsPersonalConveyanceApplied
`func (o *HosAvailabilityRead) UnsetIsPersonalConveyanceApplied()`

UnsetIsPersonalConveyanceApplied ensures that no value is present for IsPersonalConveyanceApplied, not even an explicit nil
### GetIsYardMoveApplied

`func (o *HosAvailabilityRead) GetIsYardMoveApplied() bool`

GetIsYardMoveApplied returns the IsYardMoveApplied field if non-nil, zero value otherwise.

### GetIsYardMoveAppliedOk

`func (o *HosAvailabilityRead) GetIsYardMoveAppliedOk() (*bool, bool)`

GetIsYardMoveAppliedOk returns a tuple with the IsYardMoveApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYardMoveApplied

`func (o *HosAvailabilityRead) SetIsYardMoveApplied(v bool)`

SetIsYardMoveApplied sets IsYardMoveApplied field to given value.

### HasIsYardMoveApplied

`func (o *HosAvailabilityRead) HasIsYardMoveApplied() bool`

HasIsYardMoveApplied returns a boolean if a field has been set.

### SetIsYardMoveAppliedNil

`func (o *HosAvailabilityRead) SetIsYardMoveAppliedNil(b bool)`

 SetIsYardMoveAppliedNil sets the value for IsYardMoveApplied to be an explicit nil

### UnsetIsYardMoveApplied
`func (o *HosAvailabilityRead) UnsetIsYardMoveApplied()`

UnsetIsYardMoveApplied ensures that no value is present for IsYardMoveApplied, not even an explicit nil
### GetIsAdverseDrivingExemptionAvailable

`func (o *HosAvailabilityRead) GetIsAdverseDrivingExemptionAvailable() bool`

GetIsAdverseDrivingExemptionAvailable returns the IsAdverseDrivingExemptionAvailable field if non-nil, zero value otherwise.

### GetIsAdverseDrivingExemptionAvailableOk

`func (o *HosAvailabilityRead) GetIsAdverseDrivingExemptionAvailableOk() (*bool, bool)`

GetIsAdverseDrivingExemptionAvailableOk returns a tuple with the IsAdverseDrivingExemptionAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdverseDrivingExemptionAvailable

`func (o *HosAvailabilityRead) SetIsAdverseDrivingExemptionAvailable(v bool)`

SetIsAdverseDrivingExemptionAvailable sets IsAdverseDrivingExemptionAvailable field to given value.

### HasIsAdverseDrivingExemptionAvailable

`func (o *HosAvailabilityRead) HasIsAdverseDrivingExemptionAvailable() bool`

HasIsAdverseDrivingExemptionAvailable returns a boolean if a field has been set.

### SetIsAdverseDrivingExemptionAvailableNil

`func (o *HosAvailabilityRead) SetIsAdverseDrivingExemptionAvailableNil(b bool)`

 SetIsAdverseDrivingExemptionAvailableNil sets the value for IsAdverseDrivingExemptionAvailable to be an explicit nil

### UnsetIsAdverseDrivingExemptionAvailable
`func (o *HosAvailabilityRead) UnsetIsAdverseDrivingExemptionAvailable()`

UnsetIsAdverseDrivingExemptionAvailable ensures that no value is present for IsAdverseDrivingExemptionAvailable, not even an explicit nil
### GetIsAdverseDrivingApplied

`func (o *HosAvailabilityRead) GetIsAdverseDrivingApplied() bool`

GetIsAdverseDrivingApplied returns the IsAdverseDrivingApplied field if non-nil, zero value otherwise.

### GetIsAdverseDrivingAppliedOk

`func (o *HosAvailabilityRead) GetIsAdverseDrivingAppliedOk() (*bool, bool)`

GetIsAdverseDrivingAppliedOk returns a tuple with the IsAdverseDrivingApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdverseDrivingApplied

`func (o *HosAvailabilityRead) SetIsAdverseDrivingApplied(v bool)`

SetIsAdverseDrivingApplied sets IsAdverseDrivingApplied field to given value.

### HasIsAdverseDrivingApplied

`func (o *HosAvailabilityRead) HasIsAdverseDrivingApplied() bool`

HasIsAdverseDrivingApplied returns a boolean if a field has been set.

### SetIsAdverseDrivingAppliedNil

`func (o *HosAvailabilityRead) SetIsAdverseDrivingAppliedNil(b bool)`

 SetIsAdverseDrivingAppliedNil sets the value for IsAdverseDrivingApplied to be an explicit nil

### UnsetIsAdverseDrivingApplied
`func (o *HosAvailabilityRead) UnsetIsAdverseDrivingApplied()`

UnsetIsAdverseDrivingApplied ensures that no value is present for IsAdverseDrivingApplied, not even an explicit nil
### GetIsMealBreakRequired

`func (o *HosAvailabilityRead) GetIsMealBreakRequired() bool`

GetIsMealBreakRequired returns the IsMealBreakRequired field if non-nil, zero value otherwise.

### GetIsMealBreakRequiredOk

`func (o *HosAvailabilityRead) GetIsMealBreakRequiredOk() (*bool, bool)`

GetIsMealBreakRequiredOk returns a tuple with the IsMealBreakRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMealBreakRequired

`func (o *HosAvailabilityRead) SetIsMealBreakRequired(v bool)`

SetIsMealBreakRequired sets IsMealBreakRequired field to given value.

### HasIsMealBreakRequired

`func (o *HosAvailabilityRead) HasIsMealBreakRequired() bool`

HasIsMealBreakRequired returns a boolean if a field has been set.

### SetIsMealBreakRequiredNil

`func (o *HosAvailabilityRead) SetIsMealBreakRequiredNil(b bool)`

 SetIsMealBreakRequiredNil sets the value for IsMealBreakRequired to be an explicit nil

### UnsetIsMealBreakRequired
`func (o *HosAvailabilityRead) UnsetIsMealBreakRequired()`

UnsetIsMealBreakRequired ensures that no value is present for IsMealBreakRequired, not even an explicit nil
### GetMealBreakDueAt

`func (o *HosAvailabilityRead) GetMealBreakDueAt() time.Time`

GetMealBreakDueAt returns the MealBreakDueAt field if non-nil, zero value otherwise.

### GetMealBreakDueAtOk

`func (o *HosAvailabilityRead) GetMealBreakDueAtOk() (*time.Time, bool)`

GetMealBreakDueAtOk returns a tuple with the MealBreakDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMealBreakDueAt

`func (o *HosAvailabilityRead) SetMealBreakDueAt(v time.Time)`

SetMealBreakDueAt sets MealBreakDueAt field to given value.

### HasMealBreakDueAt

`func (o *HosAvailabilityRead) HasMealBreakDueAt() bool`

HasMealBreakDueAt returns a boolean if a field has been set.

### SetMealBreakDueAtNil

`func (o *HosAvailabilityRead) SetMealBreakDueAtNil(b bool)`

 SetMealBreakDueAtNil sets the value for MealBreakDueAt to be an explicit nil

### UnsetMealBreakDueAt
`func (o *HosAvailabilityRead) UnsetMealBreakDueAt()`

UnsetMealBreakDueAt ensures that no value is present for MealBreakDueAt, not even an explicit nil
### GetMealBreakMinDurationSeconds

`func (o *HosAvailabilityRead) GetMealBreakMinDurationSeconds() int32`

GetMealBreakMinDurationSeconds returns the MealBreakMinDurationSeconds field if non-nil, zero value otherwise.

### GetMealBreakMinDurationSecondsOk

`func (o *HosAvailabilityRead) GetMealBreakMinDurationSecondsOk() (*int32, bool)`

GetMealBreakMinDurationSecondsOk returns a tuple with the MealBreakMinDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMealBreakMinDurationSeconds

`func (o *HosAvailabilityRead) SetMealBreakMinDurationSeconds(v int32)`

SetMealBreakMinDurationSeconds sets MealBreakMinDurationSeconds field to given value.

### HasMealBreakMinDurationSeconds

`func (o *HosAvailabilityRead) HasMealBreakMinDurationSeconds() bool`

HasMealBreakMinDurationSeconds returns a boolean if a field has been set.

### SetMealBreakMinDurationSecondsNil

`func (o *HosAvailabilityRead) SetMealBreakMinDurationSecondsNil(b bool)`

 SetMealBreakMinDurationSecondsNil sets the value for MealBreakMinDurationSeconds to be an explicit nil

### UnsetMealBreakMinDurationSeconds
`func (o *HosAvailabilityRead) UnsetMealBreakMinDurationSeconds()`

UnsetMealBreakMinDurationSeconds ensures that no value is present for MealBreakMinDurationSeconds, not even an explicit nil
### GetMealBreakTimeUntilDueSeconds

`func (o *HosAvailabilityRead) GetMealBreakTimeUntilDueSeconds() int32`

GetMealBreakTimeUntilDueSeconds returns the MealBreakTimeUntilDueSeconds field if non-nil, zero value otherwise.

### GetMealBreakTimeUntilDueSecondsOk

`func (o *HosAvailabilityRead) GetMealBreakTimeUntilDueSecondsOk() (*int32, bool)`

GetMealBreakTimeUntilDueSecondsOk returns a tuple with the MealBreakTimeUntilDueSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMealBreakTimeUntilDueSeconds

`func (o *HosAvailabilityRead) SetMealBreakTimeUntilDueSeconds(v int32)`

SetMealBreakTimeUntilDueSeconds sets MealBreakTimeUntilDueSeconds field to given value.

### HasMealBreakTimeUntilDueSeconds

`func (o *HosAvailabilityRead) HasMealBreakTimeUntilDueSeconds() bool`

HasMealBreakTimeUntilDueSeconds returns a boolean if a field has been set.

### SetMealBreakTimeUntilDueSecondsNil

`func (o *HosAvailabilityRead) SetMealBreakTimeUntilDueSecondsNil(b bool)`

 SetMealBreakTimeUntilDueSecondsNil sets the value for MealBreakTimeUntilDueSeconds to be an explicit nil

### UnsetMealBreakTimeUntilDueSeconds
`func (o *HosAvailabilityRead) UnsetMealBreakTimeUntilDueSeconds()`

UnsetMealBreakTimeUntilDueSeconds ensures that no value is present for MealBreakTimeUntilDueSeconds, not even an explicit nil
### GetIsSplitSleepApplied

`func (o *HosAvailabilityRead) GetIsSplitSleepApplied() bool`

GetIsSplitSleepApplied returns the IsSplitSleepApplied field if non-nil, zero value otherwise.

### GetIsSplitSleepAppliedOk

`func (o *HosAvailabilityRead) GetIsSplitSleepAppliedOk() (*bool, bool)`

GetIsSplitSleepAppliedOk returns a tuple with the IsSplitSleepApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSplitSleepApplied

`func (o *HosAvailabilityRead) SetIsSplitSleepApplied(v bool)`

SetIsSplitSleepApplied sets IsSplitSleepApplied field to given value.

### HasIsSplitSleepApplied

`func (o *HosAvailabilityRead) HasIsSplitSleepApplied() bool`

HasIsSplitSleepApplied returns a boolean if a field has been set.

### SetIsSplitSleepAppliedNil

`func (o *HosAvailabilityRead) SetIsSplitSleepAppliedNil(b bool)`

 SetIsSplitSleepAppliedNil sets the value for IsSplitSleepApplied to be an explicit nil

### UnsetIsSplitSleepApplied
`func (o *HosAvailabilityRead) UnsetIsSplitSleepApplied()`

UnsetIsSplitSleepApplied ensures that no value is present for IsSplitSleepApplied, not even an explicit nil
### GetIsSleeperEligible

`func (o *HosAvailabilityRead) GetIsSleeperEligible() bool`

GetIsSleeperEligible returns the IsSleeperEligible field if non-nil, zero value otherwise.

### GetIsSleeperEligibleOk

`func (o *HosAvailabilityRead) GetIsSleeperEligibleOk() (*bool, bool)`

GetIsSleeperEligibleOk returns a tuple with the IsSleeperEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSleeperEligible

`func (o *HosAvailabilityRead) SetIsSleeperEligible(v bool)`

SetIsSleeperEligible sets IsSleeperEligible field to given value.

### HasIsSleeperEligible

`func (o *HosAvailabilityRead) HasIsSleeperEligible() bool`

HasIsSleeperEligible returns a boolean if a field has been set.

### SetIsSleeperEligibleNil

`func (o *HosAvailabilityRead) SetIsSleeperEligibleNil(b bool)`

 SetIsSleeperEligibleNil sets the value for IsSleeperEligible to be an explicit nil

### UnsetIsSleeperEligible
`func (o *HosAvailabilityRead) UnsetIsSleeperEligible()`

UnsetIsSleeperEligible ensures that no value is present for IsSleeperEligible, not even an explicit nil
### GetSleeperRequiredRemainingSeconds

`func (o *HosAvailabilityRead) GetSleeperRequiredRemainingSeconds() int32`

GetSleeperRequiredRemainingSeconds returns the SleeperRequiredRemainingSeconds field if non-nil, zero value otherwise.

### GetSleeperRequiredRemainingSecondsOk

`func (o *HosAvailabilityRead) GetSleeperRequiredRemainingSecondsOk() (*int32, bool)`

GetSleeperRequiredRemainingSecondsOk returns a tuple with the SleeperRequiredRemainingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleeperRequiredRemainingSeconds

`func (o *HosAvailabilityRead) SetSleeperRequiredRemainingSeconds(v int32)`

SetSleeperRequiredRemainingSeconds sets SleeperRequiredRemainingSeconds field to given value.

### HasSleeperRequiredRemainingSeconds

`func (o *HosAvailabilityRead) HasSleeperRequiredRemainingSeconds() bool`

HasSleeperRequiredRemainingSeconds returns a boolean if a field has been set.

### SetSleeperRequiredRemainingSecondsNil

`func (o *HosAvailabilityRead) SetSleeperRequiredRemainingSecondsNil(b bool)`

 SetSleeperRequiredRemainingSecondsNil sets the value for SleeperRequiredRemainingSeconds to be an explicit nil

### UnsetSleeperRequiredRemainingSeconds
`func (o *HosAvailabilityRead) UnsetSleeperRequiredRemainingSeconds()`

UnsetSleeperRequiredRemainingSeconds ensures that no value is present for SleeperRequiredRemainingSeconds, not even an explicit nil
### GetSleeperSplitWindowEndsAt

`func (o *HosAvailabilityRead) GetSleeperSplitWindowEndsAt() time.Time`

GetSleeperSplitWindowEndsAt returns the SleeperSplitWindowEndsAt field if non-nil, zero value otherwise.

### GetSleeperSplitWindowEndsAtOk

`func (o *HosAvailabilityRead) GetSleeperSplitWindowEndsAtOk() (*time.Time, bool)`

GetSleeperSplitWindowEndsAtOk returns a tuple with the SleeperSplitWindowEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleeperSplitWindowEndsAt

`func (o *HosAvailabilityRead) SetSleeperSplitWindowEndsAt(v time.Time)`

SetSleeperSplitWindowEndsAt sets SleeperSplitWindowEndsAt field to given value.

### HasSleeperSplitWindowEndsAt

`func (o *HosAvailabilityRead) HasSleeperSplitWindowEndsAt() bool`

HasSleeperSplitWindowEndsAt returns a boolean if a field has been set.

### SetSleeperSplitWindowEndsAtNil

`func (o *HosAvailabilityRead) SetSleeperSplitWindowEndsAtNil(b bool)`

 SetSleeperSplitWindowEndsAtNil sets the value for SleeperSplitWindowEndsAt to be an explicit nil

### UnsetSleeperSplitWindowEndsAt
`func (o *HosAvailabilityRead) UnsetSleeperSplitWindowEndsAt()`

UnsetSleeperSplitWindowEndsAt ensures that no value is present for SleeperSplitWindowEndsAt, not even an explicit nil
### GetIsCycleApplicable

`func (o *HosAvailabilityRead) GetIsCycleApplicable() bool`

GetIsCycleApplicable returns the IsCycleApplicable field if non-nil, zero value otherwise.

### GetIsCycleApplicableOk

`func (o *HosAvailabilityRead) GetIsCycleApplicableOk() (*bool, bool)`

GetIsCycleApplicableOk returns a tuple with the IsCycleApplicable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCycleApplicable

`func (o *HosAvailabilityRead) SetIsCycleApplicable(v bool)`

SetIsCycleApplicable sets IsCycleApplicable field to given value.

### HasIsCycleApplicable

`func (o *HosAvailabilityRead) HasIsCycleApplicable() bool`

HasIsCycleApplicable returns a boolean if a field has been set.

### SetIsCycleApplicableNil

`func (o *HosAvailabilityRead) SetIsCycleApplicableNil(b bool)`

 SetIsCycleApplicableNil sets the value for IsCycleApplicable to be an explicit nil

### UnsetIsCycleApplicable
`func (o *HosAvailabilityRead) UnsetIsCycleApplicable()`

UnsetIsCycleApplicable ensures that no value is present for IsCycleApplicable, not even an explicit nil
### GetExceptionCodes

`func (o *HosAvailabilityRead) GetExceptionCodes() []string`

GetExceptionCodes returns the ExceptionCodes field if non-nil, zero value otherwise.

### GetExceptionCodesOk

`func (o *HosAvailabilityRead) GetExceptionCodesOk() (*[]string, bool)`

GetExceptionCodesOk returns a tuple with the ExceptionCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionCodes

`func (o *HosAvailabilityRead) SetExceptionCodes(v []string)`

SetExceptionCodes sets ExceptionCodes field to given value.

### HasExceptionCodes

`func (o *HosAvailabilityRead) HasExceptionCodes() bool`

HasExceptionCodes returns a boolean if a field has been set.

### SetExceptionCodesNil

`func (o *HosAvailabilityRead) SetExceptionCodesNil(b bool)`

 SetExceptionCodesNil sets the value for ExceptionCodes to be an explicit nil

### UnsetExceptionCodes
`func (o *HosAvailabilityRead) UnsetExceptionCodes()`

UnsetExceptionCodes ensures that no value is present for ExceptionCodes, not even an explicit nil
### GetNotes

`func (o *HosAvailabilityRead) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *HosAvailabilityRead) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *HosAvailabilityRead) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *HosAvailabilityRead) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *HosAvailabilityRead) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *HosAvailabilityRead) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


