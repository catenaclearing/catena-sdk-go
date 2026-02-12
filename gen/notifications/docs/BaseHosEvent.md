# BaseHosEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal unique identifier for the telematics event record (Catena PK). | 
**FleetId** | **string** | The Catena fleet this record belongs to (multi-tenant scope). | 
**FleetRef** | **NullableString** |  | 
**SourceName** | [**TspEnum**](TspEnum.md) | The name of the source | 
**ConnectionId** | **string** | The specific fleet↔TSP connection through which this record was sourced. | 
**SourceId** | **string** | The ID of the record in the TSP or a deterministic ID/Hash generated from a composite unique key | 
**CreatedAt** | **time.Time** | Immutable: first time this record was ingested into our system. | 
**UpdatedAt** | **time.Time** | Last time we modified this record in our system. | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**OccurredAt** | **time.Time** | When the underlying event/observation occurred, as reported by the TSP, or the moment it was ingested by us if not available. | 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 
**DriverId** | Pointer to **NullableString** |  | [optional] 
**VehicleId** | Pointer to **NullableString** |  | [optional] 
**CoDriverId** | Pointer to **NullableString** |  | [optional] 
**AdditionalDriverIds** | Pointer to **map[string]interface{}** |  | [optional] 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 
**DriverLastEditAt** | Pointer to **NullableTime** |  | [optional] 
**DriverCertifiedAt** | Pointer to **NullableTime** |  | [optional] 
**HosRulesetCode** | Pointer to [**NullableHosRulesetCodeEnum**](HosRulesetCodeEnum.md) |  | [optional] 
**TimeZoneCode** | Pointer to [**NullableTimezoneCodeEnum**](TimezoneCodeEnum.md) |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**DutyStatusCode** | Pointer to [**NullableDutyStatusCodeEnum**](DutyStatusCodeEnum.md) |  | [optional] 
**EventTypeCode** | Pointer to [**NullableHosEventTypeCodeEnum**](HosEventTypeCodeEnum.md) |  | [optional] 
**EventCode** | Pointer to [**NullableHosEventCodeEnum**](HosEventCodeEnum.md) |  | [optional] 
**LogStateCode** | Pointer to [**NullableHosRecordStatusCodeEnum**](HosRecordStatusCodeEnum.md) |  | [optional] 
**LogOriginCode** | Pointer to [**NullableHosRecordOriginCodeEnum**](HosRecordOriginCodeEnum.md) |  | [optional] 
**DeferralStatus** | Pointer to **NullableString** |  | [optional] 
**DeferralMinutes** | Pointer to **NullableInt32** |  | [optional] 
**EldMalfunctionCode** | Pointer to [**NullableHosMalfunctionCodeEnum**](HosMalfunctionCodeEnum.md) |  | [optional] 
**IsExcluded** | Pointer to **NullableBool** |  | [optional] 
**IsTransitioning** | Pointer to **NullableBool** |  | [optional] 
**RecordStatus** | Pointer to **NullableString** |  | [optional] 
**RegionCode** | Pointer to [**NullableHosRegionCodeEnum**](HosRegionCodeEnum.md) |  | [optional] 
**LocationName** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to [**NullablePoint**](Point.md) |  | [optional] 
**H3Index11** | Pointer to **NullableInt32** |  | [optional] 
**Odometer** | Pointer to **NullableFloat32** |  | [optional] 
**EngineHours** | Pointer to **NullableFloat32** |  | [optional] 
**Sequence** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBaseHosEvent

`func NewBaseHosEvent(id string, fleetId string, fleetRef NullableString, sourceName TspEnum, connectionId string, sourceId string, createdAt time.Time, updatedAt time.Time, occurredAt time.Time, ) *BaseHosEvent`

NewBaseHosEvent instantiates a new BaseHosEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseHosEventWithDefaults

`func NewBaseHosEventWithDefaults() *BaseHosEvent`

NewBaseHosEventWithDefaults instantiates a new BaseHosEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseHosEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseHosEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseHosEvent) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseHosEvent) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseHosEvent) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseHosEvent) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *BaseHosEvent) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *BaseHosEvent) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *BaseHosEvent) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### SetFleetRefNil

`func (o *BaseHosEvent) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *BaseHosEvent) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetSourceName

`func (o *BaseHosEvent) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseHosEvent) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseHosEvent) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetConnectionId

`func (o *BaseHosEvent) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseHosEvent) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseHosEvent) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceId

`func (o *BaseHosEvent) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseHosEvent) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseHosEvent) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCreatedAt

`func (o *BaseHosEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseHosEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseHosEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseHosEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseHosEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseHosEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BaseHosEvent) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BaseHosEvent) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BaseHosEvent) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BaseHosEvent) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BaseHosEvent) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BaseHosEvent) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOccurredAt

`func (o *BaseHosEvent) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BaseHosEvent) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BaseHosEvent) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetExecutionId

`func (o *BaseHosEvent) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BaseHosEvent) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BaseHosEvent) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *BaseHosEvent) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *BaseHosEvent) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *BaseHosEvent) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *BaseHosEvent) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BaseHosEvent) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BaseHosEvent) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *BaseHosEvent) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *BaseHosEvent) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *BaseHosEvent) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetDriverId

`func (o *BaseHosEvent) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *BaseHosEvent) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *BaseHosEvent) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *BaseHosEvent) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *BaseHosEvent) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *BaseHosEvent) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetVehicleId

`func (o *BaseHosEvent) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *BaseHosEvent) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *BaseHosEvent) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *BaseHosEvent) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *BaseHosEvent) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *BaseHosEvent) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetCoDriverId

`func (o *BaseHosEvent) GetCoDriverId() string`

GetCoDriverId returns the CoDriverId field if non-nil, zero value otherwise.

### GetCoDriverIdOk

`func (o *BaseHosEvent) GetCoDriverIdOk() (*string, bool)`

GetCoDriverIdOk returns a tuple with the CoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDriverId

`func (o *BaseHosEvent) SetCoDriverId(v string)`

SetCoDriverId sets CoDriverId field to given value.

### HasCoDriverId

`func (o *BaseHosEvent) HasCoDriverId() bool`

HasCoDriverId returns a boolean if a field has been set.

### SetCoDriverIdNil

`func (o *BaseHosEvent) SetCoDriverIdNil(b bool)`

 SetCoDriverIdNil sets the value for CoDriverId to be an explicit nil

### UnsetCoDriverId
`func (o *BaseHosEvent) UnsetCoDriverId()`

UnsetCoDriverId ensures that no value is present for CoDriverId, not even an explicit nil
### GetAdditionalDriverIds

`func (o *BaseHosEvent) GetAdditionalDriverIds() map[string]interface{}`

GetAdditionalDriverIds returns the AdditionalDriverIds field if non-nil, zero value otherwise.

### GetAdditionalDriverIdsOk

`func (o *BaseHosEvent) GetAdditionalDriverIdsOk() (*map[string]interface{}, bool)`

GetAdditionalDriverIdsOk returns a tuple with the AdditionalDriverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDriverIds

`func (o *BaseHosEvent) SetAdditionalDriverIds(v map[string]interface{})`

SetAdditionalDriverIds sets AdditionalDriverIds field to given value.

### HasAdditionalDriverIds

`func (o *BaseHosEvent) HasAdditionalDriverIds() bool`

HasAdditionalDriverIds returns a boolean if a field has been set.

### SetAdditionalDriverIdsNil

`func (o *BaseHosEvent) SetAdditionalDriverIdsNil(b bool)`

 SetAdditionalDriverIdsNil sets the value for AdditionalDriverIds to be an explicit nil

### UnsetAdditionalDriverIds
`func (o *BaseHosEvent) UnsetAdditionalDriverIds()`

UnsetAdditionalDriverIds ensures that no value is present for AdditionalDriverIds, not even an explicit nil
### GetStartedAt

`func (o *BaseHosEvent) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BaseHosEvent) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BaseHosEvent) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *BaseHosEvent) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *BaseHosEvent) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *BaseHosEvent) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *BaseHosEvent) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *BaseHosEvent) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *BaseHosEvent) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *BaseHosEvent) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *BaseHosEvent) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *BaseHosEvent) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetDriverLastEditAt

`func (o *BaseHosEvent) GetDriverLastEditAt() time.Time`

GetDriverLastEditAt returns the DriverLastEditAt field if non-nil, zero value otherwise.

### GetDriverLastEditAtOk

`func (o *BaseHosEvent) GetDriverLastEditAtOk() (*time.Time, bool)`

GetDriverLastEditAtOk returns a tuple with the DriverLastEditAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLastEditAt

`func (o *BaseHosEvent) SetDriverLastEditAt(v time.Time)`

SetDriverLastEditAt sets DriverLastEditAt field to given value.

### HasDriverLastEditAt

`func (o *BaseHosEvent) HasDriverLastEditAt() bool`

HasDriverLastEditAt returns a boolean if a field has been set.

### SetDriverLastEditAtNil

`func (o *BaseHosEvent) SetDriverLastEditAtNil(b bool)`

 SetDriverLastEditAtNil sets the value for DriverLastEditAt to be an explicit nil

### UnsetDriverLastEditAt
`func (o *BaseHosEvent) UnsetDriverLastEditAt()`

UnsetDriverLastEditAt ensures that no value is present for DriverLastEditAt, not even an explicit nil
### GetDriverCertifiedAt

`func (o *BaseHosEvent) GetDriverCertifiedAt() time.Time`

GetDriverCertifiedAt returns the DriverCertifiedAt field if non-nil, zero value otherwise.

### GetDriverCertifiedAtOk

`func (o *BaseHosEvent) GetDriverCertifiedAtOk() (*time.Time, bool)`

GetDriverCertifiedAtOk returns a tuple with the DriverCertifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverCertifiedAt

`func (o *BaseHosEvent) SetDriverCertifiedAt(v time.Time)`

SetDriverCertifiedAt sets DriverCertifiedAt field to given value.

### HasDriverCertifiedAt

`func (o *BaseHosEvent) HasDriverCertifiedAt() bool`

HasDriverCertifiedAt returns a boolean if a field has been set.

### SetDriverCertifiedAtNil

`func (o *BaseHosEvent) SetDriverCertifiedAtNil(b bool)`

 SetDriverCertifiedAtNil sets the value for DriverCertifiedAt to be an explicit nil

### UnsetDriverCertifiedAt
`func (o *BaseHosEvent) UnsetDriverCertifiedAt()`

UnsetDriverCertifiedAt ensures that no value is present for DriverCertifiedAt, not even an explicit nil
### GetHosRulesetCode

`func (o *BaseHosEvent) GetHosRulesetCode() HosRulesetCodeEnum`

GetHosRulesetCode returns the HosRulesetCode field if non-nil, zero value otherwise.

### GetHosRulesetCodeOk

`func (o *BaseHosEvent) GetHosRulesetCodeOk() (*HosRulesetCodeEnum, bool)`

GetHosRulesetCodeOk returns a tuple with the HosRulesetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosRulesetCode

`func (o *BaseHosEvent) SetHosRulesetCode(v HosRulesetCodeEnum)`

SetHosRulesetCode sets HosRulesetCode field to given value.

### HasHosRulesetCode

`func (o *BaseHosEvent) HasHosRulesetCode() bool`

HasHosRulesetCode returns a boolean if a field has been set.

### SetHosRulesetCodeNil

`func (o *BaseHosEvent) SetHosRulesetCodeNil(b bool)`

 SetHosRulesetCodeNil sets the value for HosRulesetCode to be an explicit nil

### UnsetHosRulesetCode
`func (o *BaseHosEvent) UnsetHosRulesetCode()`

UnsetHosRulesetCode ensures that no value is present for HosRulesetCode, not even an explicit nil
### GetTimeZoneCode

`func (o *BaseHosEvent) GetTimeZoneCode() TimezoneCodeEnum`

GetTimeZoneCode returns the TimeZoneCode field if non-nil, zero value otherwise.

### GetTimeZoneCodeOk

`func (o *BaseHosEvent) GetTimeZoneCodeOk() (*TimezoneCodeEnum, bool)`

GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneCode

`func (o *BaseHosEvent) SetTimeZoneCode(v TimezoneCodeEnum)`

SetTimeZoneCode sets TimeZoneCode field to given value.

### HasTimeZoneCode

`func (o *BaseHosEvent) HasTimeZoneCode() bool`

HasTimeZoneCode returns a boolean if a field has been set.

### SetTimeZoneCodeNil

`func (o *BaseHosEvent) SetTimeZoneCodeNil(b bool)`

 SetTimeZoneCodeNil sets the value for TimeZoneCode to be an explicit nil

### UnsetTimeZoneCode
`func (o *BaseHosEvent) UnsetTimeZoneCode()`

UnsetTimeZoneCode ensures that no value is present for TimeZoneCode, not even an explicit nil
### GetNotes

`func (o *BaseHosEvent) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BaseHosEvent) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BaseHosEvent) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BaseHosEvent) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *BaseHosEvent) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *BaseHosEvent) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetDutyStatusCode

`func (o *BaseHosEvent) GetDutyStatusCode() DutyStatusCodeEnum`

GetDutyStatusCode returns the DutyStatusCode field if non-nil, zero value otherwise.

### GetDutyStatusCodeOk

`func (o *BaseHosEvent) GetDutyStatusCodeOk() (*DutyStatusCodeEnum, bool)`

GetDutyStatusCodeOk returns a tuple with the DutyStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyStatusCode

`func (o *BaseHosEvent) SetDutyStatusCode(v DutyStatusCodeEnum)`

SetDutyStatusCode sets DutyStatusCode field to given value.

### HasDutyStatusCode

`func (o *BaseHosEvent) HasDutyStatusCode() bool`

HasDutyStatusCode returns a boolean if a field has been set.

### SetDutyStatusCodeNil

`func (o *BaseHosEvent) SetDutyStatusCodeNil(b bool)`

 SetDutyStatusCodeNil sets the value for DutyStatusCode to be an explicit nil

### UnsetDutyStatusCode
`func (o *BaseHosEvent) UnsetDutyStatusCode()`

UnsetDutyStatusCode ensures that no value is present for DutyStatusCode, not even an explicit nil
### GetEventTypeCode

`func (o *BaseHosEvent) GetEventTypeCode() HosEventTypeCodeEnum`

GetEventTypeCode returns the EventTypeCode field if non-nil, zero value otherwise.

### GetEventTypeCodeOk

`func (o *BaseHosEvent) GetEventTypeCodeOk() (*HosEventTypeCodeEnum, bool)`

GetEventTypeCodeOk returns a tuple with the EventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeCode

`func (o *BaseHosEvent) SetEventTypeCode(v HosEventTypeCodeEnum)`

SetEventTypeCode sets EventTypeCode field to given value.

### HasEventTypeCode

`func (o *BaseHosEvent) HasEventTypeCode() bool`

HasEventTypeCode returns a boolean if a field has been set.

### SetEventTypeCodeNil

`func (o *BaseHosEvent) SetEventTypeCodeNil(b bool)`

 SetEventTypeCodeNil sets the value for EventTypeCode to be an explicit nil

### UnsetEventTypeCode
`func (o *BaseHosEvent) UnsetEventTypeCode()`

UnsetEventTypeCode ensures that no value is present for EventTypeCode, not even an explicit nil
### GetEventCode

`func (o *BaseHosEvent) GetEventCode() HosEventCodeEnum`

GetEventCode returns the EventCode field if non-nil, zero value otherwise.

### GetEventCodeOk

`func (o *BaseHosEvent) GetEventCodeOk() (*HosEventCodeEnum, bool)`

GetEventCodeOk returns a tuple with the EventCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCode

`func (o *BaseHosEvent) SetEventCode(v HosEventCodeEnum)`

SetEventCode sets EventCode field to given value.

### HasEventCode

`func (o *BaseHosEvent) HasEventCode() bool`

HasEventCode returns a boolean if a field has been set.

### SetEventCodeNil

`func (o *BaseHosEvent) SetEventCodeNil(b bool)`

 SetEventCodeNil sets the value for EventCode to be an explicit nil

### UnsetEventCode
`func (o *BaseHosEvent) UnsetEventCode()`

UnsetEventCode ensures that no value is present for EventCode, not even an explicit nil
### GetLogStateCode

`func (o *BaseHosEvent) GetLogStateCode() HosRecordStatusCodeEnum`

GetLogStateCode returns the LogStateCode field if non-nil, zero value otherwise.

### GetLogStateCodeOk

`func (o *BaseHosEvent) GetLogStateCodeOk() (*HosRecordStatusCodeEnum, bool)`

GetLogStateCodeOk returns a tuple with the LogStateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStateCode

`func (o *BaseHosEvent) SetLogStateCode(v HosRecordStatusCodeEnum)`

SetLogStateCode sets LogStateCode field to given value.

### HasLogStateCode

`func (o *BaseHosEvent) HasLogStateCode() bool`

HasLogStateCode returns a boolean if a field has been set.

### SetLogStateCodeNil

`func (o *BaseHosEvent) SetLogStateCodeNil(b bool)`

 SetLogStateCodeNil sets the value for LogStateCode to be an explicit nil

### UnsetLogStateCode
`func (o *BaseHosEvent) UnsetLogStateCode()`

UnsetLogStateCode ensures that no value is present for LogStateCode, not even an explicit nil
### GetLogOriginCode

`func (o *BaseHosEvent) GetLogOriginCode() HosRecordOriginCodeEnum`

GetLogOriginCode returns the LogOriginCode field if non-nil, zero value otherwise.

### GetLogOriginCodeOk

`func (o *BaseHosEvent) GetLogOriginCodeOk() (*HosRecordOriginCodeEnum, bool)`

GetLogOriginCodeOk returns a tuple with the LogOriginCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOriginCode

`func (o *BaseHosEvent) SetLogOriginCode(v HosRecordOriginCodeEnum)`

SetLogOriginCode sets LogOriginCode field to given value.

### HasLogOriginCode

`func (o *BaseHosEvent) HasLogOriginCode() bool`

HasLogOriginCode returns a boolean if a field has been set.

### SetLogOriginCodeNil

`func (o *BaseHosEvent) SetLogOriginCodeNil(b bool)`

 SetLogOriginCodeNil sets the value for LogOriginCode to be an explicit nil

### UnsetLogOriginCode
`func (o *BaseHosEvent) UnsetLogOriginCode()`

UnsetLogOriginCode ensures that no value is present for LogOriginCode, not even an explicit nil
### GetDeferralStatus

`func (o *BaseHosEvent) GetDeferralStatus() string`

GetDeferralStatus returns the DeferralStatus field if non-nil, zero value otherwise.

### GetDeferralStatusOk

`func (o *BaseHosEvent) GetDeferralStatusOk() (*string, bool)`

GetDeferralStatusOk returns a tuple with the DeferralStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferralStatus

`func (o *BaseHosEvent) SetDeferralStatus(v string)`

SetDeferralStatus sets DeferralStatus field to given value.

### HasDeferralStatus

`func (o *BaseHosEvent) HasDeferralStatus() bool`

HasDeferralStatus returns a boolean if a field has been set.

### SetDeferralStatusNil

`func (o *BaseHosEvent) SetDeferralStatusNil(b bool)`

 SetDeferralStatusNil sets the value for DeferralStatus to be an explicit nil

### UnsetDeferralStatus
`func (o *BaseHosEvent) UnsetDeferralStatus()`

UnsetDeferralStatus ensures that no value is present for DeferralStatus, not even an explicit nil
### GetDeferralMinutes

`func (o *BaseHosEvent) GetDeferralMinutes() int32`

GetDeferralMinutes returns the DeferralMinutes field if non-nil, zero value otherwise.

### GetDeferralMinutesOk

`func (o *BaseHosEvent) GetDeferralMinutesOk() (*int32, bool)`

GetDeferralMinutesOk returns a tuple with the DeferralMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferralMinutes

`func (o *BaseHosEvent) SetDeferralMinutes(v int32)`

SetDeferralMinutes sets DeferralMinutes field to given value.

### HasDeferralMinutes

`func (o *BaseHosEvent) HasDeferralMinutes() bool`

HasDeferralMinutes returns a boolean if a field has been set.

### SetDeferralMinutesNil

`func (o *BaseHosEvent) SetDeferralMinutesNil(b bool)`

 SetDeferralMinutesNil sets the value for DeferralMinutes to be an explicit nil

### UnsetDeferralMinutes
`func (o *BaseHosEvent) UnsetDeferralMinutes()`

UnsetDeferralMinutes ensures that no value is present for DeferralMinutes, not even an explicit nil
### GetEldMalfunctionCode

`func (o *BaseHosEvent) GetEldMalfunctionCode() HosMalfunctionCodeEnum`

GetEldMalfunctionCode returns the EldMalfunctionCode field if non-nil, zero value otherwise.

### GetEldMalfunctionCodeOk

`func (o *BaseHosEvent) GetEldMalfunctionCodeOk() (*HosMalfunctionCodeEnum, bool)`

GetEldMalfunctionCodeOk returns a tuple with the EldMalfunctionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldMalfunctionCode

`func (o *BaseHosEvent) SetEldMalfunctionCode(v HosMalfunctionCodeEnum)`

SetEldMalfunctionCode sets EldMalfunctionCode field to given value.

### HasEldMalfunctionCode

`func (o *BaseHosEvent) HasEldMalfunctionCode() bool`

HasEldMalfunctionCode returns a boolean if a field has been set.

### SetEldMalfunctionCodeNil

`func (o *BaseHosEvent) SetEldMalfunctionCodeNil(b bool)`

 SetEldMalfunctionCodeNil sets the value for EldMalfunctionCode to be an explicit nil

### UnsetEldMalfunctionCode
`func (o *BaseHosEvent) UnsetEldMalfunctionCode()`

UnsetEldMalfunctionCode ensures that no value is present for EldMalfunctionCode, not even an explicit nil
### GetIsExcluded

`func (o *BaseHosEvent) GetIsExcluded() bool`

GetIsExcluded returns the IsExcluded field if non-nil, zero value otherwise.

### GetIsExcludedOk

`func (o *BaseHosEvent) GetIsExcludedOk() (*bool, bool)`

GetIsExcludedOk returns a tuple with the IsExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExcluded

`func (o *BaseHosEvent) SetIsExcluded(v bool)`

SetIsExcluded sets IsExcluded field to given value.

### HasIsExcluded

`func (o *BaseHosEvent) HasIsExcluded() bool`

HasIsExcluded returns a boolean if a field has been set.

### SetIsExcludedNil

`func (o *BaseHosEvent) SetIsExcludedNil(b bool)`

 SetIsExcludedNil sets the value for IsExcluded to be an explicit nil

### UnsetIsExcluded
`func (o *BaseHosEvent) UnsetIsExcluded()`

UnsetIsExcluded ensures that no value is present for IsExcluded, not even an explicit nil
### GetIsTransitioning

`func (o *BaseHosEvent) GetIsTransitioning() bool`

GetIsTransitioning returns the IsTransitioning field if non-nil, zero value otherwise.

### GetIsTransitioningOk

`func (o *BaseHosEvent) GetIsTransitioningOk() (*bool, bool)`

GetIsTransitioningOk returns a tuple with the IsTransitioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTransitioning

`func (o *BaseHosEvent) SetIsTransitioning(v bool)`

SetIsTransitioning sets IsTransitioning field to given value.

### HasIsTransitioning

`func (o *BaseHosEvent) HasIsTransitioning() bool`

HasIsTransitioning returns a boolean if a field has been set.

### SetIsTransitioningNil

`func (o *BaseHosEvent) SetIsTransitioningNil(b bool)`

 SetIsTransitioningNil sets the value for IsTransitioning to be an explicit nil

### UnsetIsTransitioning
`func (o *BaseHosEvent) UnsetIsTransitioning()`

UnsetIsTransitioning ensures that no value is present for IsTransitioning, not even an explicit nil
### GetRecordStatus

`func (o *BaseHosEvent) GetRecordStatus() string`

GetRecordStatus returns the RecordStatus field if non-nil, zero value otherwise.

### GetRecordStatusOk

`func (o *BaseHosEvent) GetRecordStatusOk() (*string, bool)`

GetRecordStatusOk returns a tuple with the RecordStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStatus

`func (o *BaseHosEvent) SetRecordStatus(v string)`

SetRecordStatus sets RecordStatus field to given value.

### HasRecordStatus

`func (o *BaseHosEvent) HasRecordStatus() bool`

HasRecordStatus returns a boolean if a field has been set.

### SetRecordStatusNil

`func (o *BaseHosEvent) SetRecordStatusNil(b bool)`

 SetRecordStatusNil sets the value for RecordStatus to be an explicit nil

### UnsetRecordStatus
`func (o *BaseHosEvent) UnsetRecordStatus()`

UnsetRecordStatus ensures that no value is present for RecordStatus, not even an explicit nil
### GetRegionCode

`func (o *BaseHosEvent) GetRegionCode() HosRegionCodeEnum`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *BaseHosEvent) GetRegionCodeOk() (*HosRegionCodeEnum, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *BaseHosEvent) SetRegionCode(v HosRegionCodeEnum)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *BaseHosEvent) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### SetRegionCodeNil

`func (o *BaseHosEvent) SetRegionCodeNil(b bool)`

 SetRegionCodeNil sets the value for RegionCode to be an explicit nil

### UnsetRegionCode
`func (o *BaseHosEvent) UnsetRegionCode()`

UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
### GetLocationName

`func (o *BaseHosEvent) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *BaseHosEvent) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *BaseHosEvent) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *BaseHosEvent) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### SetLocationNameNil

`func (o *BaseHosEvent) SetLocationNameNil(b bool)`

 SetLocationNameNil sets the value for LocationName to be an explicit nil

### UnsetLocationName
`func (o *BaseHosEvent) UnsetLocationName()`

UnsetLocationName ensures that no value is present for LocationName, not even an explicit nil
### GetLocation

`func (o *BaseHosEvent) GetLocation() Point`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BaseHosEvent) GetLocationOk() (*Point, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BaseHosEvent) SetLocation(v Point)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BaseHosEvent) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *BaseHosEvent) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *BaseHosEvent) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetH3Index11

`func (o *BaseHosEvent) GetH3Index11() int32`

GetH3Index11 returns the H3Index11 field if non-nil, zero value otherwise.

### GetH3Index11Ok

`func (o *BaseHosEvent) GetH3Index11Ok() (*int32, bool)`

GetH3Index11Ok returns a tuple with the H3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH3Index11

`func (o *BaseHosEvent) SetH3Index11(v int32)`

SetH3Index11 sets H3Index11 field to given value.

### HasH3Index11

`func (o *BaseHosEvent) HasH3Index11() bool`

HasH3Index11 returns a boolean if a field has been set.

### SetH3Index11Nil

`func (o *BaseHosEvent) SetH3Index11Nil(b bool)`

 SetH3Index11Nil sets the value for H3Index11 to be an explicit nil

### UnsetH3Index11
`func (o *BaseHosEvent) UnsetH3Index11()`

UnsetH3Index11 ensures that no value is present for H3Index11, not even an explicit nil
### GetOdometer

`func (o *BaseHosEvent) GetOdometer() float32`

GetOdometer returns the Odometer field if non-nil, zero value otherwise.

### GetOdometerOk

`func (o *BaseHosEvent) GetOdometerOk() (*float32, bool)`

GetOdometerOk returns a tuple with the Odometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometer

`func (o *BaseHosEvent) SetOdometer(v float32)`

SetOdometer sets Odometer field to given value.

### HasOdometer

`func (o *BaseHosEvent) HasOdometer() bool`

HasOdometer returns a boolean if a field has been set.

### SetOdometerNil

`func (o *BaseHosEvent) SetOdometerNil(b bool)`

 SetOdometerNil sets the value for Odometer to be an explicit nil

### UnsetOdometer
`func (o *BaseHosEvent) UnsetOdometer()`

UnsetOdometer ensures that no value is present for Odometer, not even an explicit nil
### GetEngineHours

`func (o *BaseHosEvent) GetEngineHours() float32`

GetEngineHours returns the EngineHours field if non-nil, zero value otherwise.

### GetEngineHoursOk

`func (o *BaseHosEvent) GetEngineHoursOk() (*float32, bool)`

GetEngineHoursOk returns a tuple with the EngineHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineHours

`func (o *BaseHosEvent) SetEngineHours(v float32)`

SetEngineHours sets EngineHours field to given value.

### HasEngineHours

`func (o *BaseHosEvent) HasEngineHours() bool`

HasEngineHours returns a boolean if a field has been set.

### SetEngineHoursNil

`func (o *BaseHosEvent) SetEngineHoursNil(b bool)`

 SetEngineHoursNil sets the value for EngineHours to be an explicit nil

### UnsetEngineHours
`func (o *BaseHosEvent) UnsetEngineHours()`

UnsetEngineHours ensures that no value is present for EngineHours, not even an explicit nil
### GetSequence

`func (o *BaseHosEvent) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *BaseHosEvent) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *BaseHosEvent) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *BaseHosEvent) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### SetSequenceNil

`func (o *BaseHosEvent) SetSequenceNil(b bool)`

 SetSequenceNil sets the value for Sequence to be an explicit nil

### UnsetSequence
`func (o *BaseHosEvent) UnsetSequence()`

UnsetSequence ensures that no value is present for Sequence, not even an explicit nil
### GetVersion

`func (o *BaseHosEvent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BaseHosEvent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BaseHosEvent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BaseHosEvent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *BaseHosEvent) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *BaseHosEvent) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


