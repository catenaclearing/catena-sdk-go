# HosAvailability

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
**VehicleId** | Pointer to **NullableString** |  | [optional] 
**HosRulesetCode** | Pointer to [**NullableHosRulesetCodeEnum**](HosRulesetCodeEnum.md) |  | [optional] 
**RegionCode** | Pointer to [**NullableHosRegionCodeEnum**](HosRegionCodeEnum.md) |  | [optional] 
**TimezoneCode** | Pointer to [**NullableTimezoneCodeEnum**](TimezoneCodeEnum.md) |  | [optional] 
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

### NewHosAvailability

`func NewHosAvailability(id string, createdAt time.Time, updatedAt time.Time, fleetId string, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *HosAvailability`

NewHosAvailability instantiates a new HosAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHosAvailabilityWithDefaults

`func NewHosAvailabilityWithDefaults() *HosAvailability`

NewHosAvailabilityWithDefaults instantiates a new HosAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HosAvailability) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HosAvailability) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HosAvailability) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *HosAvailability) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HosAvailability) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HosAvailability) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *HosAvailability) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HosAvailability) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HosAvailability) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *HosAvailability) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *HosAvailability) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *HosAvailability) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *HosAvailability) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *HosAvailability) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *HosAvailability) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetFleetId

`func (o *HosAvailability) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *HosAvailability) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *HosAvailability) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetConnectionId

`func (o *HosAvailability) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *HosAvailability) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *HosAvailability) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceName

`func (o *HosAvailability) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *HosAvailability) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *HosAvailability) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *HosAvailability) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *HosAvailability) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *HosAvailability) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *HosAvailability) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *HosAvailability) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *HosAvailability) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *HosAvailability) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *HosAvailability) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *HosAvailability) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *HosAvailability) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *HosAvailability) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *HosAvailability) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *HosAvailability) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *HosAvailability) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *HosAvailability) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *HosAvailability) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *HosAvailability) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HosAvailability) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HosAvailability) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HosAvailability) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HosAvailability) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HosAvailability) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *HosAvailability) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *HosAvailability) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *HosAvailability) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *HosAvailability) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *HosAvailability) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *HosAvailability) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetDriverId

`func (o *HosAvailability) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *HosAvailability) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *HosAvailability) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *HosAvailability) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *HosAvailability) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *HosAvailability) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetVehicleId

`func (o *HosAvailability) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *HosAvailability) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *HosAvailability) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *HosAvailability) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *HosAvailability) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *HosAvailability) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetHosRulesetCode

`func (o *HosAvailability) GetHosRulesetCode() HosRulesetCodeEnum`

GetHosRulesetCode returns the HosRulesetCode field if non-nil, zero value otherwise.

### GetHosRulesetCodeOk

`func (o *HosAvailability) GetHosRulesetCodeOk() (*HosRulesetCodeEnum, bool)`

GetHosRulesetCodeOk returns a tuple with the HosRulesetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosRulesetCode

`func (o *HosAvailability) SetHosRulesetCode(v HosRulesetCodeEnum)`

SetHosRulesetCode sets HosRulesetCode field to given value.

### HasHosRulesetCode

`func (o *HosAvailability) HasHosRulesetCode() bool`

HasHosRulesetCode returns a boolean if a field has been set.

### SetHosRulesetCodeNil

`func (o *HosAvailability) SetHosRulesetCodeNil(b bool)`

 SetHosRulesetCodeNil sets the value for HosRulesetCode to be an explicit nil

### UnsetHosRulesetCode
`func (o *HosAvailability) UnsetHosRulesetCode()`

UnsetHosRulesetCode ensures that no value is present for HosRulesetCode, not even an explicit nil
### GetRegionCode

`func (o *HosAvailability) GetRegionCode() HosRegionCodeEnum`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *HosAvailability) GetRegionCodeOk() (*HosRegionCodeEnum, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *HosAvailability) SetRegionCode(v HosRegionCodeEnum)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *HosAvailability) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### SetRegionCodeNil

`func (o *HosAvailability) SetRegionCodeNil(b bool)`

 SetRegionCodeNil sets the value for RegionCode to be an explicit nil

### UnsetRegionCode
`func (o *HosAvailability) UnsetRegionCode()`

UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
### GetTimezoneCode

`func (o *HosAvailability) GetTimezoneCode() TimezoneCodeEnum`

GetTimezoneCode returns the TimezoneCode field if non-nil, zero value otherwise.

### GetTimezoneCodeOk

`func (o *HosAvailability) GetTimezoneCodeOk() (*TimezoneCodeEnum, bool)`

GetTimezoneCodeOk returns a tuple with the TimezoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneCode

`func (o *HosAvailability) SetTimezoneCode(v TimezoneCodeEnum)`

SetTimezoneCode sets TimezoneCode field to given value.

### HasTimezoneCode

`func (o *HosAvailability) HasTimezoneCode() bool`

HasTimezoneCode returns a boolean if a field has been set.

### SetTimezoneCodeNil

`func (o *HosAvailability) SetTimezoneCodeNil(b bool)`

 SetTimezoneCodeNil sets the value for TimezoneCode to be an explicit nil

### UnsetTimezoneCode
`func (o *HosAvailability) UnsetTimezoneCode()`

UnsetTimezoneCode ensures that no value is present for TimezoneCode, not even an explicit nil
### GetDutyStatusCode

`func (o *HosAvailability) GetDutyStatusCode() DutyStatusCodeEnum`

GetDutyStatusCode returns the DutyStatusCode field if non-nil, zero value otherwise.

### GetDutyStatusCodeOk

`func (o *HosAvailability) GetDutyStatusCodeOk() (*DutyStatusCodeEnum, bool)`

GetDutyStatusCodeOk returns a tuple with the DutyStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyStatusCode

`func (o *HosAvailability) SetDutyStatusCode(v DutyStatusCodeEnum)`

SetDutyStatusCode sets DutyStatusCode field to given value.

### HasDutyStatusCode

`func (o *HosAvailability) HasDutyStatusCode() bool`

HasDutyStatusCode returns a boolean if a field has been set.

### SetDutyStatusCodeNil

`func (o *HosAvailability) SetDutyStatusCodeNil(b bool)`

 SetDutyStatusCodeNil sets the value for DutyStatusCode to be an explicit nil

### UnsetDutyStatusCode
`func (o *HosAvailability) UnsetDutyStatusCode()`

UnsetDutyStatusCode ensures that no value is present for DutyStatusCode, not even an explicit nil
### GetCycleStartedAt

`func (o *HosAvailability) GetCycleStartedAt() time.Time`

GetCycleStartedAt returns the CycleStartedAt field if non-nil, zero value otherwise.

### GetCycleStartedAtOk

`func (o *HosAvailability) GetCycleStartedAtOk() (*time.Time, bool)`

GetCycleStartedAtOk returns a tuple with the CycleStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleStartedAt

`func (o *HosAvailability) SetCycleStartedAt(v time.Time)`

SetCycleStartedAt sets CycleStartedAt field to given value.

### HasCycleStartedAt

`func (o *HosAvailability) HasCycleStartedAt() bool`

HasCycleStartedAt returns a boolean if a field has been set.

### SetCycleStartedAtNil

`func (o *HosAvailability) SetCycleStartedAtNil(b bool)`

 SetCycleStartedAtNil sets the value for CycleStartedAt to be an explicit nil

### UnsetCycleStartedAt
`func (o *HosAvailability) UnsetCycleStartedAt()`

UnsetCycleStartedAt ensures that no value is present for CycleStartedAt, not even an explicit nil
### GetCycleEndsAt

`func (o *HosAvailability) GetCycleEndsAt() time.Time`

GetCycleEndsAt returns the CycleEndsAt field if non-nil, zero value otherwise.

### GetCycleEndsAtOk

`func (o *HosAvailability) GetCycleEndsAtOk() (*time.Time, bool)`

GetCycleEndsAtOk returns a tuple with the CycleEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleEndsAt

`func (o *HosAvailability) SetCycleEndsAt(v time.Time)`

SetCycleEndsAt sets CycleEndsAt field to given value.

### HasCycleEndsAt

`func (o *HosAvailability) HasCycleEndsAt() bool`

HasCycleEndsAt returns a boolean if a field has been set.

### SetCycleEndsAtNil

`func (o *HosAvailability) SetCycleEndsAtNil(b bool)`

 SetCycleEndsAtNil sets the value for CycleEndsAt to be an explicit nil

### UnsetCycleEndsAt
`func (o *HosAvailability) UnsetCycleEndsAt()`

UnsetCycleEndsAt ensures that no value is present for CycleEndsAt, not even an explicit nil
### GetAvailableDriveSeconds

`func (o *HosAvailability) GetAvailableDriveSeconds() int32`

GetAvailableDriveSeconds returns the AvailableDriveSeconds field if non-nil, zero value otherwise.

### GetAvailableDriveSecondsOk

`func (o *HosAvailability) GetAvailableDriveSecondsOk() (*int32, bool)`

GetAvailableDriveSecondsOk returns a tuple with the AvailableDriveSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDriveSeconds

`func (o *HosAvailability) SetAvailableDriveSeconds(v int32)`

SetAvailableDriveSeconds sets AvailableDriveSeconds field to given value.

### HasAvailableDriveSeconds

`func (o *HosAvailability) HasAvailableDriveSeconds() bool`

HasAvailableDriveSeconds returns a boolean if a field has been set.

### SetAvailableDriveSecondsNil

`func (o *HosAvailability) SetAvailableDriveSecondsNil(b bool)`

 SetAvailableDriveSecondsNil sets the value for AvailableDriveSeconds to be an explicit nil

### UnsetAvailableDriveSeconds
`func (o *HosAvailability) UnsetAvailableDriveSeconds()`

UnsetAvailableDriveSeconds ensures that no value is present for AvailableDriveSeconds, not even an explicit nil
### GetAvailableShiftSeconds

`func (o *HosAvailability) GetAvailableShiftSeconds() int32`

GetAvailableShiftSeconds returns the AvailableShiftSeconds field if non-nil, zero value otherwise.

### GetAvailableShiftSecondsOk

`func (o *HosAvailability) GetAvailableShiftSecondsOk() (*int32, bool)`

GetAvailableShiftSecondsOk returns a tuple with the AvailableShiftSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableShiftSeconds

`func (o *HosAvailability) SetAvailableShiftSeconds(v int32)`

SetAvailableShiftSeconds sets AvailableShiftSeconds field to given value.

### HasAvailableShiftSeconds

`func (o *HosAvailability) HasAvailableShiftSeconds() bool`

HasAvailableShiftSeconds returns a boolean if a field has been set.

### SetAvailableShiftSecondsNil

`func (o *HosAvailability) SetAvailableShiftSecondsNil(b bool)`

 SetAvailableShiftSecondsNil sets the value for AvailableShiftSeconds to be an explicit nil

### UnsetAvailableShiftSeconds
`func (o *HosAvailability) UnsetAvailableShiftSeconds()`

UnsetAvailableShiftSeconds ensures that no value is present for AvailableShiftSeconds, not even an explicit nil
### GetAvailableCycleSeconds

`func (o *HosAvailability) GetAvailableCycleSeconds() int32`

GetAvailableCycleSeconds returns the AvailableCycleSeconds field if non-nil, zero value otherwise.

### GetAvailableCycleSecondsOk

`func (o *HosAvailability) GetAvailableCycleSecondsOk() (*int32, bool)`

GetAvailableCycleSecondsOk returns a tuple with the AvailableCycleSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableCycleSeconds

`func (o *HosAvailability) SetAvailableCycleSeconds(v int32)`

SetAvailableCycleSeconds sets AvailableCycleSeconds field to given value.

### HasAvailableCycleSeconds

`func (o *HosAvailability) HasAvailableCycleSeconds() bool`

HasAvailableCycleSeconds returns a boolean if a field has been set.

### SetAvailableCycleSecondsNil

`func (o *HosAvailability) SetAvailableCycleSecondsNil(b bool)`

 SetAvailableCycleSecondsNil sets the value for AvailableCycleSeconds to be an explicit nil

### UnsetAvailableCycleSeconds
`func (o *HosAvailability) UnsetAvailableCycleSeconds()`

UnsetAvailableCycleSeconds ensures that no value is present for AvailableCycleSeconds, not even an explicit nil
### GetAvailableTomorrowSeconds

`func (o *HosAvailability) GetAvailableTomorrowSeconds() int32`

GetAvailableTomorrowSeconds returns the AvailableTomorrowSeconds field if non-nil, zero value otherwise.

### GetAvailableTomorrowSecondsOk

`func (o *HosAvailability) GetAvailableTomorrowSecondsOk() (*int32, bool)`

GetAvailableTomorrowSecondsOk returns a tuple with the AvailableTomorrowSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTomorrowSeconds

`func (o *HosAvailability) SetAvailableTomorrowSeconds(v int32)`

SetAvailableTomorrowSeconds sets AvailableTomorrowSeconds field to given value.

### HasAvailableTomorrowSeconds

`func (o *HosAvailability) HasAvailableTomorrowSeconds() bool`

HasAvailableTomorrowSeconds returns a boolean if a field has been set.

### SetAvailableTomorrowSecondsNil

`func (o *HosAvailability) SetAvailableTomorrowSecondsNil(b bool)`

 SetAvailableTomorrowSecondsNil sets the value for AvailableTomorrowSeconds to be an explicit nil

### UnsetAvailableTomorrowSeconds
`func (o *HosAvailability) UnsetAvailableTomorrowSeconds()`

UnsetAvailableTomorrowSeconds ensures that no value is present for AvailableTomorrowSeconds, not even an explicit nil
### GetAvailableDay2Seconds

`func (o *HosAvailability) GetAvailableDay2Seconds() int32`

GetAvailableDay2Seconds returns the AvailableDay2Seconds field if non-nil, zero value otherwise.

### GetAvailableDay2SecondsOk

`func (o *HosAvailability) GetAvailableDay2SecondsOk() (*int32, bool)`

GetAvailableDay2SecondsOk returns a tuple with the AvailableDay2Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDay2Seconds

`func (o *HosAvailability) SetAvailableDay2Seconds(v int32)`

SetAvailableDay2Seconds sets AvailableDay2Seconds field to given value.

### HasAvailableDay2Seconds

`func (o *HosAvailability) HasAvailableDay2Seconds() bool`

HasAvailableDay2Seconds returns a boolean if a field has been set.

### SetAvailableDay2SecondsNil

`func (o *HosAvailability) SetAvailableDay2SecondsNil(b bool)`

 SetAvailableDay2SecondsNil sets the value for AvailableDay2Seconds to be an explicit nil

### UnsetAvailableDay2Seconds
`func (o *HosAvailability) UnsetAvailableDay2Seconds()`

UnsetAvailableDay2Seconds ensures that no value is present for AvailableDay2Seconds, not even an explicit nil
### GetAvailableDay3Seconds

`func (o *HosAvailability) GetAvailableDay3Seconds() int32`

GetAvailableDay3Seconds returns the AvailableDay3Seconds field if non-nil, zero value otherwise.

### GetAvailableDay3SecondsOk

`func (o *HosAvailability) GetAvailableDay3SecondsOk() (*int32, bool)`

GetAvailableDay3SecondsOk returns a tuple with the AvailableDay3Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableDay3Seconds

`func (o *HosAvailability) SetAvailableDay3Seconds(v int32)`

SetAvailableDay3Seconds sets AvailableDay3Seconds field to given value.

### HasAvailableDay3Seconds

`func (o *HosAvailability) HasAvailableDay3Seconds() bool`

HasAvailableDay3Seconds returns a boolean if a field has been set.

### SetAvailableDay3SecondsNil

`func (o *HosAvailability) SetAvailableDay3SecondsNil(b bool)`

 SetAvailableDay3SecondsNil sets the value for AvailableDay3Seconds to be an explicit nil

### UnsetAvailableDay3Seconds
`func (o *HosAvailability) UnsetAvailableDay3Seconds()`

UnsetAvailableDay3Seconds ensures that no value is present for AvailableDay3Seconds, not even an explicit nil
### GetForecastHorizonDays

`func (o *HosAvailability) GetForecastHorizonDays() int32`

GetForecastHorizonDays returns the ForecastHorizonDays field if non-nil, zero value otherwise.

### GetForecastHorizonDaysOk

`func (o *HosAvailability) GetForecastHorizonDaysOk() (*int32, bool)`

GetForecastHorizonDaysOk returns a tuple with the ForecastHorizonDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastHorizonDays

`func (o *HosAvailability) SetForecastHorizonDays(v int32)`

SetForecastHorizonDays sets ForecastHorizonDays field to given value.

### HasForecastHorizonDays

`func (o *HosAvailability) HasForecastHorizonDays() bool`

HasForecastHorizonDays returns a boolean if a field has been set.

### SetForecastHorizonDaysNil

`func (o *HosAvailability) SetForecastHorizonDaysNil(b bool)`

 SetForecastHorizonDaysNil sets the value for ForecastHorizonDays to be an explicit nil

### UnsetForecastHorizonDays
`func (o *HosAvailability) UnsetForecastHorizonDays()`

UnsetForecastHorizonDays ensures that no value is present for ForecastHorizonDays, not even an explicit nil
### GetTimeUntilBreakSeconds

`func (o *HosAvailability) GetTimeUntilBreakSeconds() int32`

GetTimeUntilBreakSeconds returns the TimeUntilBreakSeconds field if non-nil, zero value otherwise.

### GetTimeUntilBreakSecondsOk

`func (o *HosAvailability) GetTimeUntilBreakSecondsOk() (*int32, bool)`

GetTimeUntilBreakSecondsOk returns a tuple with the TimeUntilBreakSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUntilBreakSeconds

`func (o *HosAvailability) SetTimeUntilBreakSeconds(v int32)`

SetTimeUntilBreakSeconds sets TimeUntilBreakSeconds field to given value.

### HasTimeUntilBreakSeconds

`func (o *HosAvailability) HasTimeUntilBreakSeconds() bool`

HasTimeUntilBreakSeconds returns a boolean if a field has been set.

### SetTimeUntilBreakSecondsNil

`func (o *HosAvailability) SetTimeUntilBreakSecondsNil(b bool)`

 SetTimeUntilBreakSecondsNil sets the value for TimeUntilBreakSeconds to be an explicit nil

### UnsetTimeUntilBreakSeconds
`func (o *HosAvailability) UnsetTimeUntilBreakSeconds()`

UnsetTimeUntilBreakSeconds ensures that no value is present for TimeUntilBreakSeconds, not even an explicit nil
### GetRestRemainingSeconds

`func (o *HosAvailability) GetRestRemainingSeconds() int32`

GetRestRemainingSeconds returns the RestRemainingSeconds field if non-nil, zero value otherwise.

### GetRestRemainingSecondsOk

`func (o *HosAvailability) GetRestRemainingSecondsOk() (*int32, bool)`

GetRestRemainingSecondsOk returns a tuple with the RestRemainingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestRemainingSeconds

`func (o *HosAvailability) SetRestRemainingSeconds(v int32)`

SetRestRemainingSeconds sets RestRemainingSeconds field to given value.

### HasRestRemainingSeconds

`func (o *HosAvailability) HasRestRemainingSeconds() bool`

HasRestRemainingSeconds returns a boolean if a field has been set.

### SetRestRemainingSecondsNil

`func (o *HosAvailability) SetRestRemainingSecondsNil(b bool)`

 SetRestRemainingSecondsNil sets the value for RestRemainingSeconds to be an explicit nil

### UnsetRestRemainingSeconds
`func (o *HosAvailability) UnsetRestRemainingSeconds()`

UnsetRestRemainingSeconds ensures that no value is present for RestRemainingSeconds, not even an explicit nil
### GetCycleViolationDurationSeconds

`func (o *HosAvailability) GetCycleViolationDurationSeconds() int32`

GetCycleViolationDurationSeconds returns the CycleViolationDurationSeconds field if non-nil, zero value otherwise.

### GetCycleViolationDurationSecondsOk

`func (o *HosAvailability) GetCycleViolationDurationSecondsOk() (*int32, bool)`

GetCycleViolationDurationSecondsOk returns a tuple with the CycleViolationDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleViolationDurationSeconds

`func (o *HosAvailability) SetCycleViolationDurationSeconds(v int32)`

SetCycleViolationDurationSeconds sets CycleViolationDurationSeconds field to given value.

### HasCycleViolationDurationSeconds

`func (o *HosAvailability) HasCycleViolationDurationSeconds() bool`

HasCycleViolationDurationSeconds returns a boolean if a field has been set.

### SetCycleViolationDurationSecondsNil

`func (o *HosAvailability) SetCycleViolationDurationSecondsNil(b bool)`

 SetCycleViolationDurationSecondsNil sets the value for CycleViolationDurationSeconds to be an explicit nil

### UnsetCycleViolationDurationSeconds
`func (o *HosAvailability) UnsetCycleViolationDurationSeconds()`

UnsetCycleViolationDurationSeconds ensures that no value is present for CycleViolationDurationSeconds, not even an explicit nil
### GetShiftDrivingViolationDurationSeconds

`func (o *HosAvailability) GetShiftDrivingViolationDurationSeconds() int32`

GetShiftDrivingViolationDurationSeconds returns the ShiftDrivingViolationDurationSeconds field if non-nil, zero value otherwise.

### GetShiftDrivingViolationDurationSecondsOk

`func (o *HosAvailability) GetShiftDrivingViolationDurationSecondsOk() (*int32, bool)`

GetShiftDrivingViolationDurationSecondsOk returns a tuple with the ShiftDrivingViolationDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftDrivingViolationDurationSeconds

`func (o *HosAvailability) SetShiftDrivingViolationDurationSeconds(v int32)`

SetShiftDrivingViolationDurationSeconds sets ShiftDrivingViolationDurationSeconds field to given value.

### HasShiftDrivingViolationDurationSeconds

`func (o *HosAvailability) HasShiftDrivingViolationDurationSeconds() bool`

HasShiftDrivingViolationDurationSeconds returns a boolean if a field has been set.

### SetShiftDrivingViolationDurationSecondsNil

`func (o *HosAvailability) SetShiftDrivingViolationDurationSecondsNil(b bool)`

 SetShiftDrivingViolationDurationSecondsNil sets the value for ShiftDrivingViolationDurationSeconds to be an explicit nil

### UnsetShiftDrivingViolationDurationSeconds
`func (o *HosAvailability) UnsetShiftDrivingViolationDurationSeconds()`

UnsetShiftDrivingViolationDurationSeconds ensures that no value is present for ShiftDrivingViolationDurationSeconds, not even an explicit nil
### GetShiftEndsAt

`func (o *HosAvailability) GetShiftEndsAt() time.Time`

GetShiftEndsAt returns the ShiftEndsAt field if non-nil, zero value otherwise.

### GetShiftEndsAtOk

`func (o *HosAvailability) GetShiftEndsAtOk() (*time.Time, bool)`

GetShiftEndsAtOk returns a tuple with the ShiftEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShiftEndsAt

`func (o *HosAvailability) SetShiftEndsAt(v time.Time)`

SetShiftEndsAt sets ShiftEndsAt field to given value.

### HasShiftEndsAt

`func (o *HosAvailability) HasShiftEndsAt() bool`

HasShiftEndsAt returns a boolean if a field has been set.

### SetShiftEndsAtNil

`func (o *HosAvailability) SetShiftEndsAtNil(b bool)`

 SetShiftEndsAtNil sets the value for ShiftEndsAt to be an explicit nil

### UnsetShiftEndsAt
`func (o *HosAvailability) UnsetShiftEndsAt()`

UnsetShiftEndsAt ensures that no value is present for ShiftEndsAt, not even an explicit nil
### GetNextBreakDueAt

`func (o *HosAvailability) GetNextBreakDueAt() time.Time`

GetNextBreakDueAt returns the NextBreakDueAt field if non-nil, zero value otherwise.

### GetNextBreakDueAtOk

`func (o *HosAvailability) GetNextBreakDueAtOk() (*time.Time, bool)`

GetNextBreakDueAtOk returns a tuple with the NextBreakDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBreakDueAt

`func (o *HosAvailability) SetNextBreakDueAt(v time.Time)`

SetNextBreakDueAt sets NextBreakDueAt field to given value.

### HasNextBreakDueAt

`func (o *HosAvailability) HasNextBreakDueAt() bool`

HasNextBreakDueAt returns a boolean if a field has been set.

### SetNextBreakDueAtNil

`func (o *HosAvailability) SetNextBreakDueAtNil(b bool)`

 SetNextBreakDueAtNil sets the value for NextBreakDueAt to be an explicit nil

### UnsetNextBreakDueAt
`func (o *HosAvailability) UnsetNextBreakDueAt()`

UnsetNextBreakDueAt ensures that no value is present for NextBreakDueAt, not even an explicit nil
### GetNext10hrResetEligibleAt

`func (o *HosAvailability) GetNext10hrResetEligibleAt() time.Time`

GetNext10hrResetEligibleAt returns the Next10hrResetEligibleAt field if non-nil, zero value otherwise.

### GetNext10hrResetEligibleAtOk

`func (o *HosAvailability) GetNext10hrResetEligibleAtOk() (*time.Time, bool)`

GetNext10hrResetEligibleAtOk returns a tuple with the Next10hrResetEligibleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext10hrResetEligibleAt

`func (o *HosAvailability) SetNext10hrResetEligibleAt(v time.Time)`

SetNext10hrResetEligibleAt sets Next10hrResetEligibleAt field to given value.

### HasNext10hrResetEligibleAt

`func (o *HosAvailability) HasNext10hrResetEligibleAt() bool`

HasNext10hrResetEligibleAt returns a boolean if a field has been set.

### SetNext10hrResetEligibleAtNil

`func (o *HosAvailability) SetNext10hrResetEligibleAtNil(b bool)`

 SetNext10hrResetEligibleAtNil sets the value for Next10hrResetEligibleAt to be an explicit nil

### UnsetNext10hrResetEligibleAt
`func (o *HosAvailability) UnsetNext10hrResetEligibleAt()`

UnsetNext10hrResetEligibleAt ensures that no value is present for Next10hrResetEligibleAt, not even an explicit nil
### GetNext34hrResetEligibleAt

`func (o *HosAvailability) GetNext34hrResetEligibleAt() time.Time`

GetNext34hrResetEligibleAt returns the Next34hrResetEligibleAt field if non-nil, zero value otherwise.

### GetNext34hrResetEligibleAtOk

`func (o *HosAvailability) GetNext34hrResetEligibleAtOk() (*time.Time, bool)`

GetNext34hrResetEligibleAtOk returns a tuple with the Next34hrResetEligibleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext34hrResetEligibleAt

`func (o *HosAvailability) SetNext34hrResetEligibleAt(v time.Time)`

SetNext34hrResetEligibleAt sets Next34hrResetEligibleAt field to given value.

### HasNext34hrResetEligibleAt

`func (o *HosAvailability) HasNext34hrResetEligibleAt() bool`

HasNext34hrResetEligibleAt returns a boolean if a field has been set.

### SetNext34hrResetEligibleAtNil

`func (o *HosAvailability) SetNext34hrResetEligibleAtNil(b bool)`

 SetNext34hrResetEligibleAtNil sets the value for Next34hrResetEligibleAt to be an explicit nil

### UnsetNext34hrResetEligibleAt
`func (o *HosAvailability) UnsetNext34hrResetEligibleAt()`

UnsetNext34hrResetEligibleAt ensures that no value is present for Next34hrResetEligibleAt, not even an explicit nil
### GetIsPersonalConveyanceApplied

`func (o *HosAvailability) GetIsPersonalConveyanceApplied() bool`

GetIsPersonalConveyanceApplied returns the IsPersonalConveyanceApplied field if non-nil, zero value otherwise.

### GetIsPersonalConveyanceAppliedOk

`func (o *HosAvailability) GetIsPersonalConveyanceAppliedOk() (*bool, bool)`

GetIsPersonalConveyanceAppliedOk returns a tuple with the IsPersonalConveyanceApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPersonalConveyanceApplied

`func (o *HosAvailability) SetIsPersonalConveyanceApplied(v bool)`

SetIsPersonalConveyanceApplied sets IsPersonalConveyanceApplied field to given value.

### HasIsPersonalConveyanceApplied

`func (o *HosAvailability) HasIsPersonalConveyanceApplied() bool`

HasIsPersonalConveyanceApplied returns a boolean if a field has been set.

### SetIsPersonalConveyanceAppliedNil

`func (o *HosAvailability) SetIsPersonalConveyanceAppliedNil(b bool)`

 SetIsPersonalConveyanceAppliedNil sets the value for IsPersonalConveyanceApplied to be an explicit nil

### UnsetIsPersonalConveyanceApplied
`func (o *HosAvailability) UnsetIsPersonalConveyanceApplied()`

UnsetIsPersonalConveyanceApplied ensures that no value is present for IsPersonalConveyanceApplied, not even an explicit nil
### GetIsYardMoveApplied

`func (o *HosAvailability) GetIsYardMoveApplied() bool`

GetIsYardMoveApplied returns the IsYardMoveApplied field if non-nil, zero value otherwise.

### GetIsYardMoveAppliedOk

`func (o *HosAvailability) GetIsYardMoveAppliedOk() (*bool, bool)`

GetIsYardMoveAppliedOk returns a tuple with the IsYardMoveApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYardMoveApplied

`func (o *HosAvailability) SetIsYardMoveApplied(v bool)`

SetIsYardMoveApplied sets IsYardMoveApplied field to given value.

### HasIsYardMoveApplied

`func (o *HosAvailability) HasIsYardMoveApplied() bool`

HasIsYardMoveApplied returns a boolean if a field has been set.

### SetIsYardMoveAppliedNil

`func (o *HosAvailability) SetIsYardMoveAppliedNil(b bool)`

 SetIsYardMoveAppliedNil sets the value for IsYardMoveApplied to be an explicit nil

### UnsetIsYardMoveApplied
`func (o *HosAvailability) UnsetIsYardMoveApplied()`

UnsetIsYardMoveApplied ensures that no value is present for IsYardMoveApplied, not even an explicit nil
### GetIsAdverseDrivingExemptionAvailable

`func (o *HosAvailability) GetIsAdverseDrivingExemptionAvailable() bool`

GetIsAdverseDrivingExemptionAvailable returns the IsAdverseDrivingExemptionAvailable field if non-nil, zero value otherwise.

### GetIsAdverseDrivingExemptionAvailableOk

`func (o *HosAvailability) GetIsAdverseDrivingExemptionAvailableOk() (*bool, bool)`

GetIsAdverseDrivingExemptionAvailableOk returns a tuple with the IsAdverseDrivingExemptionAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdverseDrivingExemptionAvailable

`func (o *HosAvailability) SetIsAdverseDrivingExemptionAvailable(v bool)`

SetIsAdverseDrivingExemptionAvailable sets IsAdverseDrivingExemptionAvailable field to given value.

### HasIsAdverseDrivingExemptionAvailable

`func (o *HosAvailability) HasIsAdverseDrivingExemptionAvailable() bool`

HasIsAdverseDrivingExemptionAvailable returns a boolean if a field has been set.

### SetIsAdverseDrivingExemptionAvailableNil

`func (o *HosAvailability) SetIsAdverseDrivingExemptionAvailableNil(b bool)`

 SetIsAdverseDrivingExemptionAvailableNil sets the value for IsAdverseDrivingExemptionAvailable to be an explicit nil

### UnsetIsAdverseDrivingExemptionAvailable
`func (o *HosAvailability) UnsetIsAdverseDrivingExemptionAvailable()`

UnsetIsAdverseDrivingExemptionAvailable ensures that no value is present for IsAdverseDrivingExemptionAvailable, not even an explicit nil
### GetIsAdverseDrivingApplied

`func (o *HosAvailability) GetIsAdverseDrivingApplied() bool`

GetIsAdverseDrivingApplied returns the IsAdverseDrivingApplied field if non-nil, zero value otherwise.

### GetIsAdverseDrivingAppliedOk

`func (o *HosAvailability) GetIsAdverseDrivingAppliedOk() (*bool, bool)`

GetIsAdverseDrivingAppliedOk returns a tuple with the IsAdverseDrivingApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdverseDrivingApplied

`func (o *HosAvailability) SetIsAdverseDrivingApplied(v bool)`

SetIsAdverseDrivingApplied sets IsAdverseDrivingApplied field to given value.

### HasIsAdverseDrivingApplied

`func (o *HosAvailability) HasIsAdverseDrivingApplied() bool`

HasIsAdverseDrivingApplied returns a boolean if a field has been set.

### SetIsAdverseDrivingAppliedNil

`func (o *HosAvailability) SetIsAdverseDrivingAppliedNil(b bool)`

 SetIsAdverseDrivingAppliedNil sets the value for IsAdverseDrivingApplied to be an explicit nil

### UnsetIsAdverseDrivingApplied
`func (o *HosAvailability) UnsetIsAdverseDrivingApplied()`

UnsetIsAdverseDrivingApplied ensures that no value is present for IsAdverseDrivingApplied, not even an explicit nil
### GetIsMealBreakRequired

`func (o *HosAvailability) GetIsMealBreakRequired() bool`

GetIsMealBreakRequired returns the IsMealBreakRequired field if non-nil, zero value otherwise.

### GetIsMealBreakRequiredOk

`func (o *HosAvailability) GetIsMealBreakRequiredOk() (*bool, bool)`

GetIsMealBreakRequiredOk returns a tuple with the IsMealBreakRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMealBreakRequired

`func (o *HosAvailability) SetIsMealBreakRequired(v bool)`

SetIsMealBreakRequired sets IsMealBreakRequired field to given value.

### HasIsMealBreakRequired

`func (o *HosAvailability) HasIsMealBreakRequired() bool`

HasIsMealBreakRequired returns a boolean if a field has been set.

### SetIsMealBreakRequiredNil

`func (o *HosAvailability) SetIsMealBreakRequiredNil(b bool)`

 SetIsMealBreakRequiredNil sets the value for IsMealBreakRequired to be an explicit nil

### UnsetIsMealBreakRequired
`func (o *HosAvailability) UnsetIsMealBreakRequired()`

UnsetIsMealBreakRequired ensures that no value is present for IsMealBreakRequired, not even an explicit nil
### GetMealBreakDueAt

`func (o *HosAvailability) GetMealBreakDueAt() time.Time`

GetMealBreakDueAt returns the MealBreakDueAt field if non-nil, zero value otherwise.

### GetMealBreakDueAtOk

`func (o *HosAvailability) GetMealBreakDueAtOk() (*time.Time, bool)`

GetMealBreakDueAtOk returns a tuple with the MealBreakDueAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMealBreakDueAt

`func (o *HosAvailability) SetMealBreakDueAt(v time.Time)`

SetMealBreakDueAt sets MealBreakDueAt field to given value.

### HasMealBreakDueAt

`func (o *HosAvailability) HasMealBreakDueAt() bool`

HasMealBreakDueAt returns a boolean if a field has been set.

### SetMealBreakDueAtNil

`func (o *HosAvailability) SetMealBreakDueAtNil(b bool)`

 SetMealBreakDueAtNil sets the value for MealBreakDueAt to be an explicit nil

### UnsetMealBreakDueAt
`func (o *HosAvailability) UnsetMealBreakDueAt()`

UnsetMealBreakDueAt ensures that no value is present for MealBreakDueAt, not even an explicit nil
### GetMealBreakMinDurationSeconds

`func (o *HosAvailability) GetMealBreakMinDurationSeconds() int32`

GetMealBreakMinDurationSeconds returns the MealBreakMinDurationSeconds field if non-nil, zero value otherwise.

### GetMealBreakMinDurationSecondsOk

`func (o *HosAvailability) GetMealBreakMinDurationSecondsOk() (*int32, bool)`

GetMealBreakMinDurationSecondsOk returns a tuple with the MealBreakMinDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMealBreakMinDurationSeconds

`func (o *HosAvailability) SetMealBreakMinDurationSeconds(v int32)`

SetMealBreakMinDurationSeconds sets MealBreakMinDurationSeconds field to given value.

### HasMealBreakMinDurationSeconds

`func (o *HosAvailability) HasMealBreakMinDurationSeconds() bool`

HasMealBreakMinDurationSeconds returns a boolean if a field has been set.

### SetMealBreakMinDurationSecondsNil

`func (o *HosAvailability) SetMealBreakMinDurationSecondsNil(b bool)`

 SetMealBreakMinDurationSecondsNil sets the value for MealBreakMinDurationSeconds to be an explicit nil

### UnsetMealBreakMinDurationSeconds
`func (o *HosAvailability) UnsetMealBreakMinDurationSeconds()`

UnsetMealBreakMinDurationSeconds ensures that no value is present for MealBreakMinDurationSeconds, not even an explicit nil
### GetMealBreakTimeUntilDueSeconds

`func (o *HosAvailability) GetMealBreakTimeUntilDueSeconds() int32`

GetMealBreakTimeUntilDueSeconds returns the MealBreakTimeUntilDueSeconds field if non-nil, zero value otherwise.

### GetMealBreakTimeUntilDueSecondsOk

`func (o *HosAvailability) GetMealBreakTimeUntilDueSecondsOk() (*int32, bool)`

GetMealBreakTimeUntilDueSecondsOk returns a tuple with the MealBreakTimeUntilDueSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMealBreakTimeUntilDueSeconds

`func (o *HosAvailability) SetMealBreakTimeUntilDueSeconds(v int32)`

SetMealBreakTimeUntilDueSeconds sets MealBreakTimeUntilDueSeconds field to given value.

### HasMealBreakTimeUntilDueSeconds

`func (o *HosAvailability) HasMealBreakTimeUntilDueSeconds() bool`

HasMealBreakTimeUntilDueSeconds returns a boolean if a field has been set.

### SetMealBreakTimeUntilDueSecondsNil

`func (o *HosAvailability) SetMealBreakTimeUntilDueSecondsNil(b bool)`

 SetMealBreakTimeUntilDueSecondsNil sets the value for MealBreakTimeUntilDueSeconds to be an explicit nil

### UnsetMealBreakTimeUntilDueSeconds
`func (o *HosAvailability) UnsetMealBreakTimeUntilDueSeconds()`

UnsetMealBreakTimeUntilDueSeconds ensures that no value is present for MealBreakTimeUntilDueSeconds, not even an explicit nil
### GetIsSplitSleepApplied

`func (o *HosAvailability) GetIsSplitSleepApplied() bool`

GetIsSplitSleepApplied returns the IsSplitSleepApplied field if non-nil, zero value otherwise.

### GetIsSplitSleepAppliedOk

`func (o *HosAvailability) GetIsSplitSleepAppliedOk() (*bool, bool)`

GetIsSplitSleepAppliedOk returns a tuple with the IsSplitSleepApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSplitSleepApplied

`func (o *HosAvailability) SetIsSplitSleepApplied(v bool)`

SetIsSplitSleepApplied sets IsSplitSleepApplied field to given value.

### HasIsSplitSleepApplied

`func (o *HosAvailability) HasIsSplitSleepApplied() bool`

HasIsSplitSleepApplied returns a boolean if a field has been set.

### SetIsSplitSleepAppliedNil

`func (o *HosAvailability) SetIsSplitSleepAppliedNil(b bool)`

 SetIsSplitSleepAppliedNil sets the value for IsSplitSleepApplied to be an explicit nil

### UnsetIsSplitSleepApplied
`func (o *HosAvailability) UnsetIsSplitSleepApplied()`

UnsetIsSplitSleepApplied ensures that no value is present for IsSplitSleepApplied, not even an explicit nil
### GetIsSleeperEligible

`func (o *HosAvailability) GetIsSleeperEligible() bool`

GetIsSleeperEligible returns the IsSleeperEligible field if non-nil, zero value otherwise.

### GetIsSleeperEligibleOk

`func (o *HosAvailability) GetIsSleeperEligibleOk() (*bool, bool)`

GetIsSleeperEligibleOk returns a tuple with the IsSleeperEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSleeperEligible

`func (o *HosAvailability) SetIsSleeperEligible(v bool)`

SetIsSleeperEligible sets IsSleeperEligible field to given value.

### HasIsSleeperEligible

`func (o *HosAvailability) HasIsSleeperEligible() bool`

HasIsSleeperEligible returns a boolean if a field has been set.

### SetIsSleeperEligibleNil

`func (o *HosAvailability) SetIsSleeperEligibleNil(b bool)`

 SetIsSleeperEligibleNil sets the value for IsSleeperEligible to be an explicit nil

### UnsetIsSleeperEligible
`func (o *HosAvailability) UnsetIsSleeperEligible()`

UnsetIsSleeperEligible ensures that no value is present for IsSleeperEligible, not even an explicit nil
### GetSleeperRequiredRemainingSeconds

`func (o *HosAvailability) GetSleeperRequiredRemainingSeconds() int32`

GetSleeperRequiredRemainingSeconds returns the SleeperRequiredRemainingSeconds field if non-nil, zero value otherwise.

### GetSleeperRequiredRemainingSecondsOk

`func (o *HosAvailability) GetSleeperRequiredRemainingSecondsOk() (*int32, bool)`

GetSleeperRequiredRemainingSecondsOk returns a tuple with the SleeperRequiredRemainingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleeperRequiredRemainingSeconds

`func (o *HosAvailability) SetSleeperRequiredRemainingSeconds(v int32)`

SetSleeperRequiredRemainingSeconds sets SleeperRequiredRemainingSeconds field to given value.

### HasSleeperRequiredRemainingSeconds

`func (o *HosAvailability) HasSleeperRequiredRemainingSeconds() bool`

HasSleeperRequiredRemainingSeconds returns a boolean if a field has been set.

### SetSleeperRequiredRemainingSecondsNil

`func (o *HosAvailability) SetSleeperRequiredRemainingSecondsNil(b bool)`

 SetSleeperRequiredRemainingSecondsNil sets the value for SleeperRequiredRemainingSeconds to be an explicit nil

### UnsetSleeperRequiredRemainingSeconds
`func (o *HosAvailability) UnsetSleeperRequiredRemainingSeconds()`

UnsetSleeperRequiredRemainingSeconds ensures that no value is present for SleeperRequiredRemainingSeconds, not even an explicit nil
### GetSleeperSplitWindowEndsAt

`func (o *HosAvailability) GetSleeperSplitWindowEndsAt() time.Time`

GetSleeperSplitWindowEndsAt returns the SleeperSplitWindowEndsAt field if non-nil, zero value otherwise.

### GetSleeperSplitWindowEndsAtOk

`func (o *HosAvailability) GetSleeperSplitWindowEndsAtOk() (*time.Time, bool)`

GetSleeperSplitWindowEndsAtOk returns a tuple with the SleeperSplitWindowEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleeperSplitWindowEndsAt

`func (o *HosAvailability) SetSleeperSplitWindowEndsAt(v time.Time)`

SetSleeperSplitWindowEndsAt sets SleeperSplitWindowEndsAt field to given value.

### HasSleeperSplitWindowEndsAt

`func (o *HosAvailability) HasSleeperSplitWindowEndsAt() bool`

HasSleeperSplitWindowEndsAt returns a boolean if a field has been set.

### SetSleeperSplitWindowEndsAtNil

`func (o *HosAvailability) SetSleeperSplitWindowEndsAtNil(b bool)`

 SetSleeperSplitWindowEndsAtNil sets the value for SleeperSplitWindowEndsAt to be an explicit nil

### UnsetSleeperSplitWindowEndsAt
`func (o *HosAvailability) UnsetSleeperSplitWindowEndsAt()`

UnsetSleeperSplitWindowEndsAt ensures that no value is present for SleeperSplitWindowEndsAt, not even an explicit nil
### GetIsCycleApplicable

`func (o *HosAvailability) GetIsCycleApplicable() bool`

GetIsCycleApplicable returns the IsCycleApplicable field if non-nil, zero value otherwise.

### GetIsCycleApplicableOk

`func (o *HosAvailability) GetIsCycleApplicableOk() (*bool, bool)`

GetIsCycleApplicableOk returns a tuple with the IsCycleApplicable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCycleApplicable

`func (o *HosAvailability) SetIsCycleApplicable(v bool)`

SetIsCycleApplicable sets IsCycleApplicable field to given value.

### HasIsCycleApplicable

`func (o *HosAvailability) HasIsCycleApplicable() bool`

HasIsCycleApplicable returns a boolean if a field has been set.

### SetIsCycleApplicableNil

`func (o *HosAvailability) SetIsCycleApplicableNil(b bool)`

 SetIsCycleApplicableNil sets the value for IsCycleApplicable to be an explicit nil

### UnsetIsCycleApplicable
`func (o *HosAvailability) UnsetIsCycleApplicable()`

UnsetIsCycleApplicable ensures that no value is present for IsCycleApplicable, not even an explicit nil
### GetExceptionCodes

`func (o *HosAvailability) GetExceptionCodes() []string`

GetExceptionCodes returns the ExceptionCodes field if non-nil, zero value otherwise.

### GetExceptionCodesOk

`func (o *HosAvailability) GetExceptionCodesOk() (*[]string, bool)`

GetExceptionCodesOk returns a tuple with the ExceptionCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionCodes

`func (o *HosAvailability) SetExceptionCodes(v []string)`

SetExceptionCodes sets ExceptionCodes field to given value.

### HasExceptionCodes

`func (o *HosAvailability) HasExceptionCodes() bool`

HasExceptionCodes returns a boolean if a field has been set.

### SetExceptionCodesNil

`func (o *HosAvailability) SetExceptionCodesNil(b bool)`

 SetExceptionCodesNil sets the value for ExceptionCodes to be an explicit nil

### UnsetExceptionCodes
`func (o *HosAvailability) UnsetExceptionCodes()`

UnsetExceptionCodes ensures that no value is present for ExceptionCodes, not even an explicit nil
### GetNotes

`func (o *HosAvailability) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *HosAvailability) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *HosAvailability) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *HosAvailability) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *HosAvailability) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *HosAvailability) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


