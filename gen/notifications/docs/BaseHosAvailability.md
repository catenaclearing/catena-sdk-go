# BaseHosAvailability

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

### NewBaseHosAvailability

`func NewBaseHosAvailability(id string, fleetId string, fleetRef NullableString, sourceName TspEnum, connectionId string, sourceId string, createdAt time.Time, updatedAt time.Time, occurredAt time.Time, ) *BaseHosAvailability`

NewBaseHosAvailability instantiates a new BaseHosAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseHosAvailabilityWithDefaults

`func NewBaseHosAvailabilityWithDefaults() *BaseHosAvailability`

NewBaseHosAvailabilityWithDefaults instantiates a new BaseHosAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseHosAvailability) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseHosAvailability) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseHosAvailability) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseHosAvailability) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseHosAvailability) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseHosAvailability) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *BaseHosAvailability) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *BaseHosAvailability) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *BaseHosAvailability) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### SetFleetRefNil

`func (o *BaseHosAvailability) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *BaseHosAvailability) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetTspId

`func (o *BaseHosAvailability) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *BaseHosAvailability) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *BaseHosAvailability) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *BaseHosAvailability) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *BaseHosAvailability) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *BaseHosAvailability) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *BaseHosAvailability) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *BaseHosAvailability) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *BaseHosAvailability) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *BaseHosAvailability) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *BaseHosAvailability) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *BaseHosAvailability) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *BaseHosAvailability) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseHosAvailability) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseHosAvailability) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetConnectionId

`func (o *BaseHosAvailability) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseHosAvailability) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseHosAvailability) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceId

`func (o *BaseHosAvailability) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseHosAvailability) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseHosAvailability) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCreatedAt

`func (o *BaseHosAvailability) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseHosAvailability) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseHosAvailability) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseHosAvailability) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseHosAvailability) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseHosAvailability) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BaseHosAvailability) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BaseHosAvailability) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BaseHosAvailability) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BaseHosAvailability) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BaseHosAvailability) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BaseHosAvailability) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOccurredAt

`func (o *BaseHosAvailability) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BaseHosAvailability) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BaseHosAvailability) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetExecutionId

`func (o *BaseHosAvailability) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BaseHosAvailability) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BaseHosAvailability) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *BaseHosAvailability) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *BaseHosAvailability) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *BaseHosAvailability) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *BaseHosAvailability) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BaseHosAvailability) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BaseHosAvailability) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *BaseHosAvailability) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *BaseHosAvailability) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *BaseHosAvailability) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *BaseHosAvailability) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *BaseHosAvailability) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *BaseHosAvailability) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *BaseHosAvailability) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *BaseHosAvailability) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *BaseHosAvailability) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetDriverId

`func (o *BaseHosAvailability) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *BaseHosAvailability) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *BaseHosAvailability) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *BaseHosAvailability) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *BaseHosAvailability) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *BaseHosAvailability) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetVehicleId

`func (o *BaseHosAvailability) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *BaseHosAvailability) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *BaseHosAvailability) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *BaseHosAvailability) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *BaseHosAvailability) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *BaseHosAvailability) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetHosRulesetCode

`func (o *BaseHosAvailability) GetHosRulesetCode() HosRulesetCodeEnum`

GetHosRulesetCode returns the HosRulesetCode field if non-nil, zero value otherwise.

### GetHosRulesetCodeOk

`func (o *BaseHosAvailability) GetHosRulesetCodeOk() (*HosRulesetCodeEnum, bool)`

GetHosRulesetCodeOk returns a tuple with the HosRulesetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosRulesetCode

`func (o *BaseHosAvailability) SetHosRulesetCode(v HosRulesetCodeEnum)`

SetHosRulesetCode sets HosRulesetCode field to given value.

### HasHosRulesetCode

`func (o *BaseHosAvailability) HasHosRulesetCode() bool`

HasHosRulesetCode returns a boolean if a field has been set.

### SetHosRulesetCodeNil

`func (o *BaseHosAvailability) SetHosRulesetCodeNil(b bool)`

 SetHosRulesetCodeNil sets the value for HosRulesetCode to be an explicit nil

### UnsetHosRulesetCode
`func (o *BaseHosAvailability) UnsetHosRulesetCode()`

UnsetHosRulesetCode ensures that no value is present for HosRulesetCode, not even an explicit nil
### GetDutyStatusCode

`func (o *BaseHosAvailability) GetDutyStatusCode() DutyStatusCodeEnum`

GetDutyStatusCode returns the DutyStatusCode field if non-nil, zero value otherwise.

### GetDutyStatusCodeOk

`func (o *BaseHosAvailability) GetDutyStatusCodeOk() (*DutyStatusCodeEnum, bool)`

GetDutyStatusCodeOk returns a tuple with the DutyStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyStatusCode

`func (o *BaseHosAvailability) SetDutyStatusCode(v DutyStatusCodeEnum)`

SetDutyStatusCode sets DutyStatusCode field to given value.

### HasDutyStatusCode

`func (o *BaseHosAvailability) HasDutyStatusCode() bool`

HasDutyStatusCode returns a boolean if a field has been set.

### SetDutyStatusCodeNil

`func (o *BaseHosAvailability) SetDutyStatusCodeNil(b bool)`

 SetDutyStatusCodeNil sets the value for DutyStatusCode to be an explicit nil

### UnsetDutyStatusCode
`func (o *BaseHosAvailability) UnsetDutyStatusCode()`

UnsetDutyStatusCode ensures that no value is present for DutyStatusCode, not even an explicit nil
### GetCycleStartedAt

`func (o *BaseHosAvailability) GetCycleStartedAt() time.Time`

GetCycleStartedAt returns the CycleStartedAt field if non-nil, zero value otherwise.

### GetCycleStartedAtOk

`func (o *BaseHosAvailability) GetCycleStartedAtOk() (*time.Time, bool)`

GetCycleStartedAtOk returns a tuple with the CycleStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleStartedAt

`func (o *BaseHosAvailability) SetCycleStartedAt(v time.Time)`

SetCycleStartedAt sets CycleStartedAt field to given value.

### HasCycleStartedAt

`func (o *BaseHosAvailability) HasCycleStartedAt() bool`

HasCycleStartedAt returns a boolean if a field has been set.

### SetCycleStartedAtNil

`func (o *BaseHosAvailability) SetCycleStartedAtNil(b bool)`

 SetCycleStartedAtNil sets the value for CycleStartedAt to be an explicit nil

### UnsetCycleStartedAt
`func (o *BaseHosAvailability) UnsetCycleStartedAt()`

UnsetCycleStartedAt ensures that no value is present for CycleStartedAt, not even an explicit nil
### GetCycleEndsAt

`func (o *BaseHosAvailability) GetCycleEndsAt() time.Time`

GetCycleEndsAt returns the CycleEndsAt field if non-nil, zero value otherwise.

### GetCycleEndsAtOk

`func (o *BaseHosAvailability) GetCycleEndsAtOk() (*time.Time, bool)`

GetCycleEndsAtOk returns a tuple with the CycleEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleEndsAt

`func (o *BaseHosAvailability) SetCycleEndsAt(v time.Time)`

SetCycleEndsAt sets CycleEndsAt field to given value.

### HasCycleEndsAt

`func (o *BaseHosAvailability) HasCycleEndsAt() bool`

HasCycleEndsAt returns a boolean if a field has been set.

### SetCycleEndsAtNil

`func (o *BaseHosAvailability) SetCycleEndsAtNil(b bool)`

 SetCycleEndsAtNil sets the value for CycleEndsAt to be an explicit nil

### UnsetCycleEndsAt
`func (o *BaseHosAvailability) UnsetCycleEndsAt()`

UnsetCycleEndsAt ensures that no value is present for CycleEndsAt, not even an explicit nil
### GetAvailableDriveSeconds

`func (o *BaseHosAvailability) GetAvailableDriveSeconds() int32`

GetAvailableDriveSeconds returns the AvailableDriveSeconds field if non-nil, zero value otherwise.

### GetAvailableDriveSecondsOk

`func (o *BaseHosAvailability) GetAvailableDriveSecondsOk() (*int32, bool)`

GetAvailableDriveSecondsOk returns a tuple with the AvailableDriveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDriveSeconds

`func (o *BaseHosAvailability) SetAvailableDriveSeconds(v int32)`

SetAvailableDriveSeconds sets AvailableDriveSeconds field to given value.

### HasAvailableDriveSeconds

`func (o *BaseHosAvailability) HasAvailableDriveSeconds() bool`

HasAvailableDriveSeconds returns a boolean if a field has been set.

### SetAvailableDriveSecondsNil

`func (o *BaseHosAvailability) SetAvailableDriveSecondsNil(b bool)`

 SetAvailableDriveSecondsNil sets the value for AvailableDriveSeconds to be an explicit nil

### UnsetAvailableDriveSeconds
`func (o *BaseHosAvailability) UnsetAvailableDriveSeconds()`

UnsetAvailableDriveSeconds ensures that no value is present for AvailableDriveSeconds, not even an explicit nil
### GetAvailableShiftSeconds

`func (o *BaseHosAvailability) GetAvailableShiftSeconds() int32`

GetAvailableShiftSeconds returns the AvailableShiftSeconds field if non-nil, zero value otherwise.

### GetAvailableShiftSecondsOk

`func (o *BaseHosAvailability) GetAvailableShiftSecondsOk() (*int32, bool)`

GetAvailableShiftSecondsOk returns a tuple with the AvailableShiftSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShiftSeconds

`func (o *BaseHosAvailability) SetAvailableShiftSeconds(v int32)`

SetAvailableShiftSeconds sets AvailableShiftSeconds field to given value.

### HasAvailableShiftSeconds

`func (o *BaseHosAvailability) HasAvailableShiftSeconds() bool`

HasAvailableShiftSeconds returns a boolean if a field has been set.

### SetAvailableShiftSecondsNil

`func (o *BaseHosAvailability) SetAvailableShiftSecondsNil(b bool)`

 SetAvailableShiftSecondsNil sets the value for AvailableShiftSeconds to be an explicit nil

### UnsetAvailableShiftSeconds
`func (o *BaseHosAvailability) UnsetAvailableShiftSeconds()`

UnsetAvailableShiftSeconds ensures that no value is present for AvailableShiftSeconds, not even an explicit nil
### GetAvailableCycleSeconds

`func (o *BaseHosAvailability) GetAvailableCycleSeconds() int32`

GetAvailableCycleSeconds returns the AvailableCycleSeconds field if non-nil, zero value otherwise.

### GetAvailableCycleSecondsOk

`func (o *BaseHosAvailability) GetAvailableCycleSecondsOk() (*int32, bool)`

GetAvailableCycleSecondsOk returns a tuple with the AvailableCycleSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCycleSeconds

`func (o *BaseHosAvailability) SetAvailableCycleSeconds(v int32)`

SetAvailableCycleSeconds sets AvailableCycleSeconds field to given value.

### HasAvailableCycleSeconds

`func (o *BaseHosAvailability) HasAvailableCycleSeconds() bool`

HasAvailableCycleSeconds returns a boolean if a field has been set.

### SetAvailableCycleSecondsNil

`func (o *BaseHosAvailability) SetAvailableCycleSecondsNil(b bool)`

 SetAvailableCycleSecondsNil sets the value for AvailableCycleSeconds to be an explicit nil

### UnsetAvailableCycleSeconds
`func (o *BaseHosAvailability) UnsetAvailableCycleSeconds()`

UnsetAvailableCycleSeconds ensures that no value is present for AvailableCycleSeconds, not even an explicit nil
### GetAvailableTomorrowSeconds

`func (o *BaseHosAvailability) GetAvailableTomorrowSeconds() int32`

GetAvailableTomorrowSeconds returns the AvailableTomorrowSeconds field if non-nil, zero value otherwise.

### GetAvailableTomorrowSecondsOk

`func (o *BaseHosAvailability) GetAvailableTomorrowSecondsOk() (*int32, bool)`

GetAvailableTomorrowSecondsOk returns a tuple with the AvailableTomorrowSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTomorrowSeconds

`func (o *BaseHosAvailability) SetAvailableTomorrowSeconds(v int32)`

SetAvailableTomorrowSeconds sets AvailableTomorrowSeconds field to given value.

### HasAvailableTomorrowSeconds

`func (o *BaseHosAvailability) HasAvailableTomorrowSeconds() bool`

HasAvailableTomorrowSeconds returns a boolean if a field has been set.

### SetAvailableTomorrowSecondsNil

`func (o *BaseHosAvailability) SetAvailableTomorrowSecondsNil(b bool)`

 SetAvailableTomorrowSecondsNil sets the value for AvailableTomorrowSeconds to be an explicit nil

### UnsetAvailableTomorrowSeconds
`func (o *BaseHosAvailability) UnsetAvailableTomorrowSeconds()`

UnsetAvailableTomorrowSeconds ensures that no value is present for AvailableTomorrowSeconds, not even an explicit nil
### GetAvailableDay2Seconds

`func (o *BaseHosAvailability) GetAvailableDay2Seconds() int32`

GetAvailableDay2Seconds returns the AvailableDay2Seconds field if non-nil, zero value otherwise.

### GetAvailableDay2SecondsOk

`func (o *BaseHosAvailability) GetAvailableDay2SecondsOk() (*int32, bool)`

GetAvailableDay2SecondsOk returns a tuple with the AvailableDay2Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDay2Seconds

`func (o *BaseHosAvailability) SetAvailableDay2Seconds(v int32)`

SetAvailableDay2Seconds sets AvailableDay2Seconds field to given value.

### HasAvailableDay2Seconds

`func (o *BaseHosAvailability) HasAvailableDay2Seconds() bool`

HasAvailableDay2Seconds returns a boolean if a field has been set.

### SetAvailableDay2SecondsNil

`func (o *BaseHosAvailability) SetAvailableDay2SecondsNil(b bool)`

 SetAvailableDay2SecondsNil sets the value for AvailableDay2Seconds to be an explicit nil

### UnsetAvailableDay2Seconds
`func (o *BaseHosAvailability) UnsetAvailableDay2Seconds()`

UnsetAvailableDay2Seconds ensures that no value is present for AvailableDay2Seconds, not even an explicit nil
### GetAvailableDay3Seconds

`func (o *BaseHosAvailability) GetAvailableDay3Seconds() int32`

GetAvailableDay3Seconds returns the AvailableDay3Seconds field if non-nil, zero value otherwise.

### GetAvailableDay3SecondsOk

`func (o *BaseHosAvailability) GetAvailableDay3SecondsOk() (*int32, bool)`

GetAvailableDay3SecondsOk returns a tuple with the AvailableDay3Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDay3Seconds

`func (o *BaseHosAvailability) SetAvailableDay3Seconds(v int32)`

SetAvailableDay3Seconds sets AvailableDay3Seconds field to given value.

### HasAvailableDay3Seconds

`func (o *BaseHosAvailability) HasAvailableDay3Seconds() bool`

HasAvailableDay3Seconds returns a boolean if a field has been set.

### SetAvailableDay3SecondsNil

`func (o *BaseHosAvailability) SetAvailableDay3SecondsNil(b bool)`

 SetAvailableDay3SecondsNil sets the value for AvailableDay3Seconds to be an explicit nil

### UnsetAvailableDay3Seconds
`func (o *BaseHosAvailability) UnsetAvailableDay3Seconds()`

UnsetAvailableDay3Seconds ensures that no value is present for AvailableDay3Seconds, not even an explicit nil
### GetForecastHorizonDays

`func (o *BaseHosAvailability) GetForecastHorizonDays() int32`

GetForecastHorizonDays returns the ForecastHorizonDays field if non-nil, zero value otherwise.

### GetForecastHorizonDaysOk

`func (o *BaseHosAvailability) GetForecastHorizonDaysOk() (*int32, bool)`

GetForecastHorizonDaysOk returns a tuple with the ForecastHorizonDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastHorizonDays

`func (o *BaseHosAvailability) SetForecastHorizonDays(v int32)`

SetForecastHorizonDays sets ForecastHorizonDays field to given value.

### HasForecastHorizonDays

`func (o *BaseHosAvailability) HasForecastHorizonDays() bool`

HasForecastHorizonDays returns a boolean if a field has been set.

### SetForecastHorizonDaysNil

`func (o *BaseHosAvailability) SetForecastHorizonDaysNil(b bool)`

 SetForecastHorizonDaysNil sets the value for ForecastHorizonDays to be an explicit nil

### UnsetForecastHorizonDays
`func (o *BaseHosAvailability) UnsetForecastHorizonDays()`

UnsetForecastHorizonDays ensures that no value is present for ForecastHorizonDays, not even an explicit nil
### GetTimeUntilBreakSeconds

`func (o *BaseHosAvailability) GetTimeUntilBreakSeconds() int32`

GetTimeUntilBreakSeconds returns the TimeUntilBreakSeconds field if non-nil, zero value otherwise.

### GetTimeUntilBreakSecondsOk

`func (o *BaseHosAvailability) GetTimeUntilBreakSecondsOk() (*int32, bool)`

GetTimeUntilBreakSecondsOk returns a tuple with the TimeUntilBreakSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUntilBreakSeconds

`func (o *BaseHosAvailability) SetTimeUntilBreakSeconds(v int32)`

SetTimeUntilBreakSeconds sets TimeUntilBreakSeconds field to given value.

### HasTimeUntilBreakSeconds

`func (o *BaseHosAvailability) HasTimeUntilBreakSeconds() bool`

HasTimeUntilBreakSeconds returns a boolean if a field has been set.

### SetTimeUntilBreakSecondsNil

`func (o *BaseHosAvailability) SetTimeUntilBreakSecondsNil(b bool)`

 SetTimeUntilBreakSecondsNil sets the value for TimeUntilBreakSeconds to be an explicit nil

### UnsetTimeUntilBreakSeconds
`func (o *BaseHosAvailability) UnsetTimeUntilBreakSeconds()`

UnsetTimeUntilBreakSeconds ensures that no value is present for TimeUntilBreakSeconds, not even an explicit nil
### GetRestRemainingSeconds

`func (o *BaseHosAvailability) GetRestRemainingSeconds() int32`

GetRestRemainingSeconds returns the RestRemainingSeconds field if non-nil, zero value otherwise.

### GetRestRemainingSecondsOk

`func (o *BaseHosAvailability) GetRestRemainingSecondsOk() (*int32, bool)`

GetRestRemainingSecondsOk returns a tuple with the RestRemainingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestRemainingSeconds

`func (o *BaseHosAvailability) SetRestRemainingSeconds(v int32)`

SetRestRemainingSeconds sets RestRemainingSeconds field to given value.

### HasRestRemainingSeconds

`func (o *BaseHosAvailability) HasRestRemainingSeconds() bool`

HasRestRemainingSeconds returns a boolean if a field has been set.

### SetRestRemainingSecondsNil

`func (o *BaseHosAvailability) SetRestRemainingSecondsNil(b bool)`

 SetRestRemainingSecondsNil sets the value for RestRemainingSeconds to be an explicit nil

### UnsetRestRemainingSeconds
`func (o *BaseHosAvailability) UnsetRestRemainingSeconds()`

UnsetRestRemainingSeconds ensures that no value is present for RestRemainingSeconds, not even an explicit nil
### GetCycleViolationDurationSeconds

`func (o *BaseHosAvailability) GetCycleViolationDurationSeconds() int32`

GetCycleViolationDurationSeconds returns the CycleViolationDurationSeconds field if non-nil, zero value otherwise.

### GetCycleViolationDurationSecondsOk

`func (o *BaseHosAvailability) GetCycleViolationDurationSecondsOk() (*int32, bool)`

GetCycleViolationDurationSecondsOk returns a tuple with the CycleViolationDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleViolationDurationSeconds

`func (o *BaseHosAvailability) SetCycleViolationDurationSeconds(v int32)`

SetCycleViolationDurationSeconds sets CycleViolationDurationSeconds field to given value.

### HasCycleViolationDurationSeconds

`func (o *BaseHosAvailability) HasCycleViolationDurationSeconds() bool`

HasCycleViolationDurationSeconds returns a boolean if a field has been set.

### SetCycleViolationDurationSecondsNil

`func (o *BaseHosAvailability) SetCycleViolationDurationSecondsNil(b bool)`

 SetCycleViolationDurationSecondsNil sets the value for CycleViolationDurationSeconds to be an explicit nil

### UnsetCycleViolationDurationSeconds
`func (o *BaseHosAvailability) UnsetCycleViolationDurationSeconds()`

UnsetCycleViolationDurationSeconds ensures that no value is present for CycleViolationDurationSeconds, not even an explicit nil
### GetShiftDrivingViolationDurationSeconds

`func (o *BaseHosAvailability) GetShiftDrivingViolationDurationSeconds() int32`

GetShiftDrivingViolationDurationSeconds returns the ShiftDrivingViolationDurationSeconds field if non-nil, zero value otherwise.

### GetShiftDrivingViolationDurationSecondsOk

`func (o *BaseHosAvailability) GetShiftDrivingViolationDurationSecondsOk() (*int32, bool)`

GetShiftDrivingViolationDurationSecondsOk returns a tuple with the ShiftDrivingViolationDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftDrivingViolationDurationSeconds

`func (o *BaseHosAvailability) SetShiftDrivingViolationDurationSeconds(v int32)`

SetShiftDrivingViolationDurationSeconds sets ShiftDrivingViolationDurationSeconds field to given value.

### HasShiftDrivingViolationDurationSeconds

`func (o *BaseHosAvailability) HasShiftDrivingViolationDurationSeconds() bool`

HasShiftDrivingViolationDurationSeconds returns a boolean if a field has been set.

### SetShiftDrivingViolationDurationSecondsNil

`func (o *BaseHosAvailability) SetShiftDrivingViolationDurationSecondsNil(b bool)`

 SetShiftDrivingViolationDurationSecondsNil sets the value for ShiftDrivingViolationDurationSeconds to be an explicit nil

### UnsetShiftDrivingViolationDurationSeconds
`func (o *BaseHosAvailability) UnsetShiftDrivingViolationDurationSeconds()`

UnsetShiftDrivingViolationDurationSeconds ensures that no value is present for ShiftDrivingViolationDurationSeconds, not even an explicit nil
### GetShiftEndsAt

`func (o *BaseHosAvailability) GetShiftEndsAt() time.Time`

GetShiftEndsAt returns the ShiftEndsAt field if non-nil, zero value otherwise.

### GetShiftEndsAtOk

`func (o *BaseHosAvailability) GetShiftEndsAtOk() (*time.Time, bool)`

GetShiftEndsAtOk returns a tuple with the ShiftEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftEndsAt

`func (o *BaseHosAvailability) SetShiftEndsAt(v time.Time)`

SetShiftEndsAt sets ShiftEndsAt field to given value.

### HasShiftEndsAt

`func (o *BaseHosAvailability) HasShiftEndsAt() bool`

HasShiftEndsAt returns a boolean if a field has been set.

### SetShiftEndsAtNil

`func (o *BaseHosAvailability) SetShiftEndsAtNil(b bool)`

 SetShiftEndsAtNil sets the value for ShiftEndsAt to be an explicit nil

### UnsetShiftEndsAt
`func (o *BaseHosAvailability) UnsetShiftEndsAt()`

UnsetShiftEndsAt ensures that no value is present for ShiftEndsAt, not even an explicit nil
### GetNextBreakDueAt

`func (o *BaseHosAvailability) GetNextBreakDueAt() time.Time`

GetNextBreakDueAt returns the NextBreakDueAt field if non-nil, zero value otherwise.

### GetNextBreakDueAtOk

`func (o *BaseHosAvailability) GetNextBreakDueAtOk() (*time.Time, bool)`

GetNextBreakDueAtOk returns a tuple with the NextBreakDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBreakDueAt

`func (o *BaseHosAvailability) SetNextBreakDueAt(v time.Time)`

SetNextBreakDueAt sets NextBreakDueAt field to given value.

### HasNextBreakDueAt

`func (o *BaseHosAvailability) HasNextBreakDueAt() bool`

HasNextBreakDueAt returns a boolean if a field has been set.

### SetNextBreakDueAtNil

`func (o *BaseHosAvailability) SetNextBreakDueAtNil(b bool)`

 SetNextBreakDueAtNil sets the value for NextBreakDueAt to be an explicit nil

### UnsetNextBreakDueAt
`func (o *BaseHosAvailability) UnsetNextBreakDueAt()`

UnsetNextBreakDueAt ensures that no value is present for NextBreakDueAt, not even an explicit nil
### GetNext10hrResetEligibleAt

`func (o *BaseHosAvailability) GetNext10hrResetEligibleAt() time.Time`

GetNext10hrResetEligibleAt returns the Next10hrResetEligibleAt field if non-nil, zero value otherwise.

### GetNext10hrResetEligibleAtOk

`func (o *BaseHosAvailability) GetNext10hrResetEligibleAtOk() (*time.Time, bool)`

GetNext10hrResetEligibleAtOk returns a tuple with the Next10hrResetEligibleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext10hrResetEligibleAt

`func (o *BaseHosAvailability) SetNext10hrResetEligibleAt(v time.Time)`

SetNext10hrResetEligibleAt sets Next10hrResetEligibleAt field to given value.

### HasNext10hrResetEligibleAt

`func (o *BaseHosAvailability) HasNext10hrResetEligibleAt() bool`

HasNext10hrResetEligibleAt returns a boolean if a field has been set.

### SetNext10hrResetEligibleAtNil

`func (o *BaseHosAvailability) SetNext10hrResetEligibleAtNil(b bool)`

 SetNext10hrResetEligibleAtNil sets the value for Next10hrResetEligibleAt to be an explicit nil

### UnsetNext10hrResetEligibleAt
`func (o *BaseHosAvailability) UnsetNext10hrResetEligibleAt()`

UnsetNext10hrResetEligibleAt ensures that no value is present for Next10hrResetEligibleAt, not even an explicit nil
### GetNext34hrResetEligibleAt

`func (o *BaseHosAvailability) GetNext34hrResetEligibleAt() time.Time`

GetNext34hrResetEligibleAt returns the Next34hrResetEligibleAt field if non-nil, zero value otherwise.

### GetNext34hrResetEligibleAtOk

`func (o *BaseHosAvailability) GetNext34hrResetEligibleAtOk() (*time.Time, bool)`

GetNext34hrResetEligibleAtOk returns a tuple with the Next34hrResetEligibleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext34hrResetEligibleAt

`func (o *BaseHosAvailability) SetNext34hrResetEligibleAt(v time.Time)`

SetNext34hrResetEligibleAt sets Next34hrResetEligibleAt field to given value.

### HasNext34hrResetEligibleAt

`func (o *BaseHosAvailability) HasNext34hrResetEligibleAt() bool`

HasNext34hrResetEligibleAt returns a boolean if a field has been set.

### SetNext34hrResetEligibleAtNil

`func (o *BaseHosAvailability) SetNext34hrResetEligibleAtNil(b bool)`

 SetNext34hrResetEligibleAtNil sets the value for Next34hrResetEligibleAt to be an explicit nil

### UnsetNext34hrResetEligibleAt
`func (o *BaseHosAvailability) UnsetNext34hrResetEligibleAt()`

UnsetNext34hrResetEligibleAt ensures that no value is present for Next34hrResetEligibleAt, not even an explicit nil
### GetIsPersonalConveyanceApplied

`func (o *BaseHosAvailability) GetIsPersonalConveyanceApplied() bool`

GetIsPersonalConveyanceApplied returns the IsPersonalConveyanceApplied field if non-nil, zero value otherwise.

### GetIsPersonalConveyanceAppliedOk

`func (o *BaseHosAvailability) GetIsPersonalConveyanceAppliedOk() (*bool, bool)`

GetIsPersonalConveyanceAppliedOk returns a tuple with the IsPersonalConveyanceApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersonalConveyanceApplied

`func (o *BaseHosAvailability) SetIsPersonalConveyanceApplied(v bool)`

SetIsPersonalConveyanceApplied sets IsPersonalConveyanceApplied field to given value.

### HasIsPersonalConveyanceApplied

`func (o *BaseHosAvailability) HasIsPersonalConveyanceApplied() bool`

HasIsPersonalConveyanceApplied returns a boolean if a field has been set.

### SetIsPersonalConveyanceAppliedNil

`func (o *BaseHosAvailability) SetIsPersonalConveyanceAppliedNil(b bool)`

 SetIsPersonalConveyanceAppliedNil sets the value for IsPersonalConveyanceApplied to be an explicit nil

### UnsetIsPersonalConveyanceApplied
`func (o *BaseHosAvailability) UnsetIsPersonalConveyanceApplied()`

UnsetIsPersonalConveyanceApplied ensures that no value is present for IsPersonalConveyanceApplied, not even an explicit nil
### GetIsYardMoveApplied

`func (o *BaseHosAvailability) GetIsYardMoveApplied() bool`

GetIsYardMoveApplied returns the IsYardMoveApplied field if non-nil, zero value otherwise.

### GetIsYardMoveAppliedOk

`func (o *BaseHosAvailability) GetIsYardMoveAppliedOk() (*bool, bool)`

GetIsYardMoveAppliedOk returns a tuple with the IsYardMoveApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYardMoveApplied

`func (o *BaseHosAvailability) SetIsYardMoveApplied(v bool)`

SetIsYardMoveApplied sets IsYardMoveApplied field to given value.

### HasIsYardMoveApplied

`func (o *BaseHosAvailability) HasIsYardMoveApplied() bool`

HasIsYardMoveApplied returns a boolean if a field has been set.

### SetIsYardMoveAppliedNil

`func (o *BaseHosAvailability) SetIsYardMoveAppliedNil(b bool)`

 SetIsYardMoveAppliedNil sets the value for IsYardMoveApplied to be an explicit nil

### UnsetIsYardMoveApplied
`func (o *BaseHosAvailability) UnsetIsYardMoveApplied()`

UnsetIsYardMoveApplied ensures that no value is present for IsYardMoveApplied, not even an explicit nil
### GetIsAdverseDrivingExemptionAvailable

`func (o *BaseHosAvailability) GetIsAdverseDrivingExemptionAvailable() bool`

GetIsAdverseDrivingExemptionAvailable returns the IsAdverseDrivingExemptionAvailable field if non-nil, zero value otherwise.

### GetIsAdverseDrivingExemptionAvailableOk

`func (o *BaseHosAvailability) GetIsAdverseDrivingExemptionAvailableOk() (*bool, bool)`

GetIsAdverseDrivingExemptionAvailableOk returns a tuple with the IsAdverseDrivingExemptionAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdverseDrivingExemptionAvailable

`func (o *BaseHosAvailability) SetIsAdverseDrivingExemptionAvailable(v bool)`

SetIsAdverseDrivingExemptionAvailable sets IsAdverseDrivingExemptionAvailable field to given value.

### HasIsAdverseDrivingExemptionAvailable

`func (o *BaseHosAvailability) HasIsAdverseDrivingExemptionAvailable() bool`

HasIsAdverseDrivingExemptionAvailable returns a boolean if a field has been set.

### SetIsAdverseDrivingExemptionAvailableNil

`func (o *BaseHosAvailability) SetIsAdverseDrivingExemptionAvailableNil(b bool)`

 SetIsAdverseDrivingExemptionAvailableNil sets the value for IsAdverseDrivingExemptionAvailable to be an explicit nil

### UnsetIsAdverseDrivingExemptionAvailable
`func (o *BaseHosAvailability) UnsetIsAdverseDrivingExemptionAvailable()`

UnsetIsAdverseDrivingExemptionAvailable ensures that no value is present for IsAdverseDrivingExemptionAvailable, not even an explicit nil
### GetIsAdverseDrivingApplied

`func (o *BaseHosAvailability) GetIsAdverseDrivingApplied() bool`

GetIsAdverseDrivingApplied returns the IsAdverseDrivingApplied field if non-nil, zero value otherwise.

### GetIsAdverseDrivingAppliedOk

`func (o *BaseHosAvailability) GetIsAdverseDrivingAppliedOk() (*bool, bool)`

GetIsAdverseDrivingAppliedOk returns a tuple with the IsAdverseDrivingApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdverseDrivingApplied

`func (o *BaseHosAvailability) SetIsAdverseDrivingApplied(v bool)`

SetIsAdverseDrivingApplied sets IsAdverseDrivingApplied field to given value.

### HasIsAdverseDrivingApplied

`func (o *BaseHosAvailability) HasIsAdverseDrivingApplied() bool`

HasIsAdverseDrivingApplied returns a boolean if a field has been set.

### SetIsAdverseDrivingAppliedNil

`func (o *BaseHosAvailability) SetIsAdverseDrivingAppliedNil(b bool)`

 SetIsAdverseDrivingAppliedNil sets the value for IsAdverseDrivingApplied to be an explicit nil

### UnsetIsAdverseDrivingApplied
`func (o *BaseHosAvailability) UnsetIsAdverseDrivingApplied()`

UnsetIsAdverseDrivingApplied ensures that no value is present for IsAdverseDrivingApplied, not even an explicit nil
### GetIsMealBreakRequired

`func (o *BaseHosAvailability) GetIsMealBreakRequired() bool`

GetIsMealBreakRequired returns the IsMealBreakRequired field if non-nil, zero value otherwise.

### GetIsMealBreakRequiredOk

`func (o *BaseHosAvailability) GetIsMealBreakRequiredOk() (*bool, bool)`

GetIsMealBreakRequiredOk returns a tuple with the IsMealBreakRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMealBreakRequired

`func (o *BaseHosAvailability) SetIsMealBreakRequired(v bool)`

SetIsMealBreakRequired sets IsMealBreakRequired field to given value.

### HasIsMealBreakRequired

`func (o *BaseHosAvailability) HasIsMealBreakRequired() bool`

HasIsMealBreakRequired returns a boolean if a field has been set.

### SetIsMealBreakRequiredNil

`func (o *BaseHosAvailability) SetIsMealBreakRequiredNil(b bool)`

 SetIsMealBreakRequiredNil sets the value for IsMealBreakRequired to be an explicit nil

### UnsetIsMealBreakRequired
`func (o *BaseHosAvailability) UnsetIsMealBreakRequired()`

UnsetIsMealBreakRequired ensures that no value is present for IsMealBreakRequired, not even an explicit nil
### GetMealBreakDueAt

`func (o *BaseHosAvailability) GetMealBreakDueAt() time.Time`

GetMealBreakDueAt returns the MealBreakDueAt field if non-nil, zero value otherwise.

### GetMealBreakDueAtOk

`func (o *BaseHosAvailability) GetMealBreakDueAtOk() (*time.Time, bool)`

GetMealBreakDueAtOk returns a tuple with the MealBreakDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMealBreakDueAt

`func (o *BaseHosAvailability) SetMealBreakDueAt(v time.Time)`

SetMealBreakDueAt sets MealBreakDueAt field to given value.

### HasMealBreakDueAt

`func (o *BaseHosAvailability) HasMealBreakDueAt() bool`

HasMealBreakDueAt returns a boolean if a field has been set.

### SetMealBreakDueAtNil

`func (o *BaseHosAvailability) SetMealBreakDueAtNil(b bool)`

 SetMealBreakDueAtNil sets the value for MealBreakDueAt to be an explicit nil

### UnsetMealBreakDueAt
`func (o *BaseHosAvailability) UnsetMealBreakDueAt()`

UnsetMealBreakDueAt ensures that no value is present for MealBreakDueAt, not even an explicit nil
### GetMealBreakMinDurationSeconds

`func (o *BaseHosAvailability) GetMealBreakMinDurationSeconds() int32`

GetMealBreakMinDurationSeconds returns the MealBreakMinDurationSeconds field if non-nil, zero value otherwise.

### GetMealBreakMinDurationSecondsOk

`func (o *BaseHosAvailability) GetMealBreakMinDurationSecondsOk() (*int32, bool)`

GetMealBreakMinDurationSecondsOk returns a tuple with the MealBreakMinDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMealBreakMinDurationSeconds

`func (o *BaseHosAvailability) SetMealBreakMinDurationSeconds(v int32)`

SetMealBreakMinDurationSeconds sets MealBreakMinDurationSeconds field to given value.

### HasMealBreakMinDurationSeconds

`func (o *BaseHosAvailability) HasMealBreakMinDurationSeconds() bool`

HasMealBreakMinDurationSeconds returns a boolean if a field has been set.

### SetMealBreakMinDurationSecondsNil

`func (o *BaseHosAvailability) SetMealBreakMinDurationSecondsNil(b bool)`

 SetMealBreakMinDurationSecondsNil sets the value for MealBreakMinDurationSeconds to be an explicit nil

### UnsetMealBreakMinDurationSeconds
`func (o *BaseHosAvailability) UnsetMealBreakMinDurationSeconds()`

UnsetMealBreakMinDurationSeconds ensures that no value is present for MealBreakMinDurationSeconds, not even an explicit nil
### GetMealBreakTimeUntilDueSeconds

`func (o *BaseHosAvailability) GetMealBreakTimeUntilDueSeconds() int32`

GetMealBreakTimeUntilDueSeconds returns the MealBreakTimeUntilDueSeconds field if non-nil, zero value otherwise.

### GetMealBreakTimeUntilDueSecondsOk

`func (o *BaseHosAvailability) GetMealBreakTimeUntilDueSecondsOk() (*int32, bool)`

GetMealBreakTimeUntilDueSecondsOk returns a tuple with the MealBreakTimeUntilDueSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMealBreakTimeUntilDueSeconds

`func (o *BaseHosAvailability) SetMealBreakTimeUntilDueSeconds(v int32)`

SetMealBreakTimeUntilDueSeconds sets MealBreakTimeUntilDueSeconds field to given value.

### HasMealBreakTimeUntilDueSeconds

`func (o *BaseHosAvailability) HasMealBreakTimeUntilDueSeconds() bool`

HasMealBreakTimeUntilDueSeconds returns a boolean if a field has been set.

### SetMealBreakTimeUntilDueSecondsNil

`func (o *BaseHosAvailability) SetMealBreakTimeUntilDueSecondsNil(b bool)`

 SetMealBreakTimeUntilDueSecondsNil sets the value for MealBreakTimeUntilDueSeconds to be an explicit nil

### UnsetMealBreakTimeUntilDueSeconds
`func (o *BaseHosAvailability) UnsetMealBreakTimeUntilDueSeconds()`

UnsetMealBreakTimeUntilDueSeconds ensures that no value is present for MealBreakTimeUntilDueSeconds, not even an explicit nil
### GetIsSplitSleepApplied

`func (o *BaseHosAvailability) GetIsSplitSleepApplied() bool`

GetIsSplitSleepApplied returns the IsSplitSleepApplied field if non-nil, zero value otherwise.

### GetIsSplitSleepAppliedOk

`func (o *BaseHosAvailability) GetIsSplitSleepAppliedOk() (*bool, bool)`

GetIsSplitSleepAppliedOk returns a tuple with the IsSplitSleepApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSplitSleepApplied

`func (o *BaseHosAvailability) SetIsSplitSleepApplied(v bool)`

SetIsSplitSleepApplied sets IsSplitSleepApplied field to given value.

### HasIsSplitSleepApplied

`func (o *BaseHosAvailability) HasIsSplitSleepApplied() bool`

HasIsSplitSleepApplied returns a boolean if a field has been set.

### SetIsSplitSleepAppliedNil

`func (o *BaseHosAvailability) SetIsSplitSleepAppliedNil(b bool)`

 SetIsSplitSleepAppliedNil sets the value for IsSplitSleepApplied to be an explicit nil

### UnsetIsSplitSleepApplied
`func (o *BaseHosAvailability) UnsetIsSplitSleepApplied()`

UnsetIsSplitSleepApplied ensures that no value is present for IsSplitSleepApplied, not even an explicit nil
### GetIsSleeperEligible

`func (o *BaseHosAvailability) GetIsSleeperEligible() bool`

GetIsSleeperEligible returns the IsSleeperEligible field if non-nil, zero value otherwise.

### GetIsSleeperEligibleOk

`func (o *BaseHosAvailability) GetIsSleeperEligibleOk() (*bool, bool)`

GetIsSleeperEligibleOk returns a tuple with the IsSleeperEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSleeperEligible

`func (o *BaseHosAvailability) SetIsSleeperEligible(v bool)`

SetIsSleeperEligible sets IsSleeperEligible field to given value.

### HasIsSleeperEligible

`func (o *BaseHosAvailability) HasIsSleeperEligible() bool`

HasIsSleeperEligible returns a boolean if a field has been set.

### SetIsSleeperEligibleNil

`func (o *BaseHosAvailability) SetIsSleeperEligibleNil(b bool)`

 SetIsSleeperEligibleNil sets the value for IsSleeperEligible to be an explicit nil

### UnsetIsSleeperEligible
`func (o *BaseHosAvailability) UnsetIsSleeperEligible()`

UnsetIsSleeperEligible ensures that no value is present for IsSleeperEligible, not even an explicit nil
### GetSleeperRequiredRemainingSeconds

`func (o *BaseHosAvailability) GetSleeperRequiredRemainingSeconds() int32`

GetSleeperRequiredRemainingSeconds returns the SleeperRequiredRemainingSeconds field if non-nil, zero value otherwise.

### GetSleeperRequiredRemainingSecondsOk

`func (o *BaseHosAvailability) GetSleeperRequiredRemainingSecondsOk() (*int32, bool)`

GetSleeperRequiredRemainingSecondsOk returns a tuple with the SleeperRequiredRemainingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleeperRequiredRemainingSeconds

`func (o *BaseHosAvailability) SetSleeperRequiredRemainingSeconds(v int32)`

SetSleeperRequiredRemainingSeconds sets SleeperRequiredRemainingSeconds field to given value.

### HasSleeperRequiredRemainingSeconds

`func (o *BaseHosAvailability) HasSleeperRequiredRemainingSeconds() bool`

HasSleeperRequiredRemainingSeconds returns a boolean if a field has been set.

### SetSleeperRequiredRemainingSecondsNil

`func (o *BaseHosAvailability) SetSleeperRequiredRemainingSecondsNil(b bool)`

 SetSleeperRequiredRemainingSecondsNil sets the value for SleeperRequiredRemainingSeconds to be an explicit nil

### UnsetSleeperRequiredRemainingSeconds
`func (o *BaseHosAvailability) UnsetSleeperRequiredRemainingSeconds()`

UnsetSleeperRequiredRemainingSeconds ensures that no value is present for SleeperRequiredRemainingSeconds, not even an explicit nil
### GetSleeperSplitWindowEndsAt

`func (o *BaseHosAvailability) GetSleeperSplitWindowEndsAt() time.Time`

GetSleeperSplitWindowEndsAt returns the SleeperSplitWindowEndsAt field if non-nil, zero value otherwise.

### GetSleeperSplitWindowEndsAtOk

`func (o *BaseHosAvailability) GetSleeperSplitWindowEndsAtOk() (*time.Time, bool)`

GetSleeperSplitWindowEndsAtOk returns a tuple with the SleeperSplitWindowEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleeperSplitWindowEndsAt

`func (o *BaseHosAvailability) SetSleeperSplitWindowEndsAt(v time.Time)`

SetSleeperSplitWindowEndsAt sets SleeperSplitWindowEndsAt field to given value.

### HasSleeperSplitWindowEndsAt

`func (o *BaseHosAvailability) HasSleeperSplitWindowEndsAt() bool`

HasSleeperSplitWindowEndsAt returns a boolean if a field has been set.

### SetSleeperSplitWindowEndsAtNil

`func (o *BaseHosAvailability) SetSleeperSplitWindowEndsAtNil(b bool)`

 SetSleeperSplitWindowEndsAtNil sets the value for SleeperSplitWindowEndsAt to be an explicit nil

### UnsetSleeperSplitWindowEndsAt
`func (o *BaseHosAvailability) UnsetSleeperSplitWindowEndsAt()`

UnsetSleeperSplitWindowEndsAt ensures that no value is present for SleeperSplitWindowEndsAt, not even an explicit nil
### GetIsCycleApplicable

`func (o *BaseHosAvailability) GetIsCycleApplicable() bool`

GetIsCycleApplicable returns the IsCycleApplicable field if non-nil, zero value otherwise.

### GetIsCycleApplicableOk

`func (o *BaseHosAvailability) GetIsCycleApplicableOk() (*bool, bool)`

GetIsCycleApplicableOk returns a tuple with the IsCycleApplicable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCycleApplicable

`func (o *BaseHosAvailability) SetIsCycleApplicable(v bool)`

SetIsCycleApplicable sets IsCycleApplicable field to given value.

### HasIsCycleApplicable

`func (o *BaseHosAvailability) HasIsCycleApplicable() bool`

HasIsCycleApplicable returns a boolean if a field has been set.

### SetIsCycleApplicableNil

`func (o *BaseHosAvailability) SetIsCycleApplicableNil(b bool)`

 SetIsCycleApplicableNil sets the value for IsCycleApplicable to be an explicit nil

### UnsetIsCycleApplicable
`func (o *BaseHosAvailability) UnsetIsCycleApplicable()`

UnsetIsCycleApplicable ensures that no value is present for IsCycleApplicable, not even an explicit nil
### GetExceptionCodes

`func (o *BaseHosAvailability) GetExceptionCodes() []string`

GetExceptionCodes returns the ExceptionCodes field if non-nil, zero value otherwise.

### GetExceptionCodesOk

`func (o *BaseHosAvailability) GetExceptionCodesOk() (*[]string, bool)`

GetExceptionCodesOk returns a tuple with the ExceptionCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionCodes

`func (o *BaseHosAvailability) SetExceptionCodes(v []string)`

SetExceptionCodes sets ExceptionCodes field to given value.

### HasExceptionCodes

`func (o *BaseHosAvailability) HasExceptionCodes() bool`

HasExceptionCodes returns a boolean if a field has been set.

### SetExceptionCodesNil

`func (o *BaseHosAvailability) SetExceptionCodesNil(b bool)`

 SetExceptionCodesNil sets the value for ExceptionCodes to be an explicit nil

### UnsetExceptionCodes
`func (o *BaseHosAvailability) UnsetExceptionCodes()`

UnsetExceptionCodes ensures that no value is present for ExceptionCodes, not even an explicit nil
### GetNotes

`func (o *BaseHosAvailability) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BaseHosAvailability) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BaseHosAvailability) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BaseHosAvailability) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *BaseHosAvailability) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *BaseHosAvailability) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


