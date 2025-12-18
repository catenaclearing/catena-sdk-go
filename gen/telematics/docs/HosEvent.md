# HosEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
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
**CoDriverId** | Pointer to **NullableString** |  | [optional] 
**AdditionalDriverIds** | Pointer to **map[string]interface{}** |  | [optional] 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 
**DriverLastEditAt** | Pointer to **NullableTime** |  | [optional] 
**DriverCertifiedAt** | Pointer to **NullableTime** |  | [optional] 
**HosRulesetCode** | Pointer to [**NullableHosRulesetCodeEnum**](HosRulesetCodeEnum.md) |  | [optional] 
**TimeZoneCode** | Pointer to [**NullableTimezoneCodeEnum**](TimezoneCodeEnum.md) |  | [optional] 
**Annotations** | Pointer to [**NullableHosEventAnnotation**](HosEventAnnotation.md) |  | [optional] 
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
**Location** | Pointer to [**NullableLocation2**](Location2.md) |  | [optional] 
**H3Index11** | Pointer to **NullableInt32** |  | [optional] 
**Odometer** | Pointer to **NullableFloat32** |  | [optional] 
**EngineHours** | Pointer to **NullableFloat32** |  | [optional] 
**Sequence** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewHosEvent

`func NewHosEvent(id string, createdAt time.Time, updatedAt time.Time, fleetId string, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *HosEvent`

NewHosEvent instantiates a new HosEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHosEventWithDefaults

`func NewHosEventWithDefaults() *HosEvent`

NewHosEventWithDefaults instantiates a new HosEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HosEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HosEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HosEvent) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *HosEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HosEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HosEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *HosEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HosEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HosEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *HosEvent) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *HosEvent) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *HosEvent) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *HosEvent) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *HosEvent) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *HosEvent) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetFleetId

`func (o *HosEvent) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *HosEvent) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *HosEvent) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetConnectionId

`func (o *HosEvent) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *HosEvent) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *HosEvent) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceName

`func (o *HosEvent) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *HosEvent) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *HosEvent) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *HosEvent) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *HosEvent) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *HosEvent) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *HosEvent) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *HosEvent) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *HosEvent) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *HosEvent) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *HosEvent) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *HosEvent) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *HosEvent) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *HosEvent) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *HosEvent) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *HosEvent) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *HosEvent) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *HosEvent) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *HosEvent) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *HosEvent) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HosEvent) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HosEvent) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HosEvent) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HosEvent) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HosEvent) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *HosEvent) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *HosEvent) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *HosEvent) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *HosEvent) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *HosEvent) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *HosEvent) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetDriverId

`func (o *HosEvent) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *HosEvent) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *HosEvent) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *HosEvent) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *HosEvent) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *HosEvent) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetVehicleId

`func (o *HosEvent) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *HosEvent) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *HosEvent) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *HosEvent) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *HosEvent) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *HosEvent) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetCoDriverId

`func (o *HosEvent) GetCoDriverId() string`

GetCoDriverId returns the CoDriverId field if non-nil, zero value otherwise.

### GetCoDriverIdOk

`func (o *HosEvent) GetCoDriverIdOk() (*string, bool)`

GetCoDriverIdOk returns a tuple with the CoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDriverId

`func (o *HosEvent) SetCoDriverId(v string)`

SetCoDriverId sets CoDriverId field to given value.

### HasCoDriverId

`func (o *HosEvent) HasCoDriverId() bool`

HasCoDriverId returns a boolean if a field has been set.

### SetCoDriverIdNil

`func (o *HosEvent) SetCoDriverIdNil(b bool)`

 SetCoDriverIdNil sets the value for CoDriverId to be an explicit nil

### UnsetCoDriverId
`func (o *HosEvent) UnsetCoDriverId()`

UnsetCoDriverId ensures that no value is present for CoDriverId, not even an explicit nil
### GetAdditionalDriverIds

`func (o *HosEvent) GetAdditionalDriverIds() map[string]interface{}`

GetAdditionalDriverIds returns the AdditionalDriverIds field if non-nil, zero value otherwise.

### GetAdditionalDriverIdsOk

`func (o *HosEvent) GetAdditionalDriverIdsOk() (*map[string]interface{}, bool)`

GetAdditionalDriverIdsOk returns a tuple with the AdditionalDriverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDriverIds

`func (o *HosEvent) SetAdditionalDriverIds(v map[string]interface{})`

SetAdditionalDriverIds sets AdditionalDriverIds field to given value.

### HasAdditionalDriverIds

`func (o *HosEvent) HasAdditionalDriverIds() bool`

HasAdditionalDriverIds returns a boolean if a field has been set.

### SetAdditionalDriverIdsNil

`func (o *HosEvent) SetAdditionalDriverIdsNil(b bool)`

 SetAdditionalDriverIdsNil sets the value for AdditionalDriverIds to be an explicit nil

### UnsetAdditionalDriverIds
`func (o *HosEvent) UnsetAdditionalDriverIds()`

UnsetAdditionalDriverIds ensures that no value is present for AdditionalDriverIds, not even an explicit nil
### GetStartedAt

`func (o *HosEvent) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *HosEvent) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *HosEvent) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *HosEvent) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *HosEvent) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *HosEvent) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *HosEvent) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *HosEvent) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *HosEvent) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *HosEvent) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *HosEvent) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *HosEvent) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetDriverLastEditAt

`func (o *HosEvent) GetDriverLastEditAt() time.Time`

GetDriverLastEditAt returns the DriverLastEditAt field if non-nil, zero value otherwise.

### GetDriverLastEditAtOk

`func (o *HosEvent) GetDriverLastEditAtOk() (*time.Time, bool)`

GetDriverLastEditAtOk returns a tuple with the DriverLastEditAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLastEditAt

`func (o *HosEvent) SetDriverLastEditAt(v time.Time)`

SetDriverLastEditAt sets DriverLastEditAt field to given value.

### HasDriverLastEditAt

`func (o *HosEvent) HasDriverLastEditAt() bool`

HasDriverLastEditAt returns a boolean if a field has been set.

### SetDriverLastEditAtNil

`func (o *HosEvent) SetDriverLastEditAtNil(b bool)`

 SetDriverLastEditAtNil sets the value for DriverLastEditAt to be an explicit nil

### UnsetDriverLastEditAt
`func (o *HosEvent) UnsetDriverLastEditAt()`

UnsetDriverLastEditAt ensures that no value is present for DriverLastEditAt, not even an explicit nil
### GetDriverCertifiedAt

`func (o *HosEvent) GetDriverCertifiedAt() time.Time`

GetDriverCertifiedAt returns the DriverCertifiedAt field if non-nil, zero value otherwise.

### GetDriverCertifiedAtOk

`func (o *HosEvent) GetDriverCertifiedAtOk() (*time.Time, bool)`

GetDriverCertifiedAtOk returns a tuple with the DriverCertifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverCertifiedAt

`func (o *HosEvent) SetDriverCertifiedAt(v time.Time)`

SetDriverCertifiedAt sets DriverCertifiedAt field to given value.

### HasDriverCertifiedAt

`func (o *HosEvent) HasDriverCertifiedAt() bool`

HasDriverCertifiedAt returns a boolean if a field has been set.

### SetDriverCertifiedAtNil

`func (o *HosEvent) SetDriverCertifiedAtNil(b bool)`

 SetDriverCertifiedAtNil sets the value for DriverCertifiedAt to be an explicit nil

### UnsetDriverCertifiedAt
`func (o *HosEvent) UnsetDriverCertifiedAt()`

UnsetDriverCertifiedAt ensures that no value is present for DriverCertifiedAt, not even an explicit nil
### GetHosRulesetCode

`func (o *HosEvent) GetHosRulesetCode() HosRulesetCodeEnum`

GetHosRulesetCode returns the HosRulesetCode field if non-nil, zero value otherwise.

### GetHosRulesetCodeOk

`func (o *HosEvent) GetHosRulesetCodeOk() (*HosRulesetCodeEnum, bool)`

GetHosRulesetCodeOk returns a tuple with the HosRulesetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosRulesetCode

`func (o *HosEvent) SetHosRulesetCode(v HosRulesetCodeEnum)`

SetHosRulesetCode sets HosRulesetCode field to given value.

### HasHosRulesetCode

`func (o *HosEvent) HasHosRulesetCode() bool`

HasHosRulesetCode returns a boolean if a field has been set.

### SetHosRulesetCodeNil

`func (o *HosEvent) SetHosRulesetCodeNil(b bool)`

 SetHosRulesetCodeNil sets the value for HosRulesetCode to be an explicit nil

### UnsetHosRulesetCode
`func (o *HosEvent) UnsetHosRulesetCode()`

UnsetHosRulesetCode ensures that no value is present for HosRulesetCode, not even an explicit nil
### GetTimeZoneCode

`func (o *HosEvent) GetTimeZoneCode() TimezoneCodeEnum`

GetTimeZoneCode returns the TimeZoneCode field if non-nil, zero value otherwise.

### GetTimeZoneCodeOk

`func (o *HosEvent) GetTimeZoneCodeOk() (*TimezoneCodeEnum, bool)`

GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneCode

`func (o *HosEvent) SetTimeZoneCode(v TimezoneCodeEnum)`

SetTimeZoneCode sets TimeZoneCode field to given value.

### HasTimeZoneCode

`func (o *HosEvent) HasTimeZoneCode() bool`

HasTimeZoneCode returns a boolean if a field has been set.

### SetTimeZoneCodeNil

`func (o *HosEvent) SetTimeZoneCodeNil(b bool)`

 SetTimeZoneCodeNil sets the value for TimeZoneCode to be an explicit nil

### UnsetTimeZoneCode
`func (o *HosEvent) UnsetTimeZoneCode()`

UnsetTimeZoneCode ensures that no value is present for TimeZoneCode, not even an explicit nil
### GetAnnotations

`func (o *HosEvent) GetAnnotations() HosEventAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *HosEvent) GetAnnotationsOk() (*HosEventAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *HosEvent) SetAnnotations(v HosEventAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *HosEvent) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *HosEvent) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *HosEvent) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetDutyStatusCode

`func (o *HosEvent) GetDutyStatusCode() DutyStatusCodeEnum`

GetDutyStatusCode returns the DutyStatusCode field if non-nil, zero value otherwise.

### GetDutyStatusCodeOk

`func (o *HosEvent) GetDutyStatusCodeOk() (*DutyStatusCodeEnum, bool)`

GetDutyStatusCodeOk returns a tuple with the DutyStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyStatusCode

`func (o *HosEvent) SetDutyStatusCode(v DutyStatusCodeEnum)`

SetDutyStatusCode sets DutyStatusCode field to given value.

### HasDutyStatusCode

`func (o *HosEvent) HasDutyStatusCode() bool`

HasDutyStatusCode returns a boolean if a field has been set.

### SetDutyStatusCodeNil

`func (o *HosEvent) SetDutyStatusCodeNil(b bool)`

 SetDutyStatusCodeNil sets the value for DutyStatusCode to be an explicit nil

### UnsetDutyStatusCode
`func (o *HosEvent) UnsetDutyStatusCode()`

UnsetDutyStatusCode ensures that no value is present for DutyStatusCode, not even an explicit nil
### GetEventTypeCode

`func (o *HosEvent) GetEventTypeCode() HosEventTypeCodeEnum`

GetEventTypeCode returns the EventTypeCode field if non-nil, zero value otherwise.

### GetEventTypeCodeOk

`func (o *HosEvent) GetEventTypeCodeOk() (*HosEventTypeCodeEnum, bool)`

GetEventTypeCodeOk returns a tuple with the EventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeCode

`func (o *HosEvent) SetEventTypeCode(v HosEventTypeCodeEnum)`

SetEventTypeCode sets EventTypeCode field to given value.

### HasEventTypeCode

`func (o *HosEvent) HasEventTypeCode() bool`

HasEventTypeCode returns a boolean if a field has been set.

### SetEventTypeCodeNil

`func (o *HosEvent) SetEventTypeCodeNil(b bool)`

 SetEventTypeCodeNil sets the value for EventTypeCode to be an explicit nil

### UnsetEventTypeCode
`func (o *HosEvent) UnsetEventTypeCode()`

UnsetEventTypeCode ensures that no value is present for EventTypeCode, not even an explicit nil
### GetEventCode

`func (o *HosEvent) GetEventCode() HosEventCodeEnum`

GetEventCode returns the EventCode field if non-nil, zero value otherwise.

### GetEventCodeOk

`func (o *HosEvent) GetEventCodeOk() (*HosEventCodeEnum, bool)`

GetEventCodeOk returns a tuple with the EventCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCode

`func (o *HosEvent) SetEventCode(v HosEventCodeEnum)`

SetEventCode sets EventCode field to given value.

### HasEventCode

`func (o *HosEvent) HasEventCode() bool`

HasEventCode returns a boolean if a field has been set.

### SetEventCodeNil

`func (o *HosEvent) SetEventCodeNil(b bool)`

 SetEventCodeNil sets the value for EventCode to be an explicit nil

### UnsetEventCode
`func (o *HosEvent) UnsetEventCode()`

UnsetEventCode ensures that no value is present for EventCode, not even an explicit nil
### GetLogStateCode

`func (o *HosEvent) GetLogStateCode() HosRecordStatusCodeEnum`

GetLogStateCode returns the LogStateCode field if non-nil, zero value otherwise.

### GetLogStateCodeOk

`func (o *HosEvent) GetLogStateCodeOk() (*HosRecordStatusCodeEnum, bool)`

GetLogStateCodeOk returns a tuple with the LogStateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStateCode

`func (o *HosEvent) SetLogStateCode(v HosRecordStatusCodeEnum)`

SetLogStateCode sets LogStateCode field to given value.

### HasLogStateCode

`func (o *HosEvent) HasLogStateCode() bool`

HasLogStateCode returns a boolean if a field has been set.

### SetLogStateCodeNil

`func (o *HosEvent) SetLogStateCodeNil(b bool)`

 SetLogStateCodeNil sets the value for LogStateCode to be an explicit nil

### UnsetLogStateCode
`func (o *HosEvent) UnsetLogStateCode()`

UnsetLogStateCode ensures that no value is present for LogStateCode, not even an explicit nil
### GetLogOriginCode

`func (o *HosEvent) GetLogOriginCode() HosRecordOriginCodeEnum`

GetLogOriginCode returns the LogOriginCode field if non-nil, zero value otherwise.

### GetLogOriginCodeOk

`func (o *HosEvent) GetLogOriginCodeOk() (*HosRecordOriginCodeEnum, bool)`

GetLogOriginCodeOk returns a tuple with the LogOriginCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOriginCode

`func (o *HosEvent) SetLogOriginCode(v HosRecordOriginCodeEnum)`

SetLogOriginCode sets LogOriginCode field to given value.

### HasLogOriginCode

`func (o *HosEvent) HasLogOriginCode() bool`

HasLogOriginCode returns a boolean if a field has been set.

### SetLogOriginCodeNil

`func (o *HosEvent) SetLogOriginCodeNil(b bool)`

 SetLogOriginCodeNil sets the value for LogOriginCode to be an explicit nil

### UnsetLogOriginCode
`func (o *HosEvent) UnsetLogOriginCode()`

UnsetLogOriginCode ensures that no value is present for LogOriginCode, not even an explicit nil
### GetDeferralStatus

`func (o *HosEvent) GetDeferralStatus() string`

GetDeferralStatus returns the DeferralStatus field if non-nil, zero value otherwise.

### GetDeferralStatusOk

`func (o *HosEvent) GetDeferralStatusOk() (*string, bool)`

GetDeferralStatusOk returns a tuple with the DeferralStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferralStatus

`func (o *HosEvent) SetDeferralStatus(v string)`

SetDeferralStatus sets DeferralStatus field to given value.

### HasDeferralStatus

`func (o *HosEvent) HasDeferralStatus() bool`

HasDeferralStatus returns a boolean if a field has been set.

### SetDeferralStatusNil

`func (o *HosEvent) SetDeferralStatusNil(b bool)`

 SetDeferralStatusNil sets the value for DeferralStatus to be an explicit nil

### UnsetDeferralStatus
`func (o *HosEvent) UnsetDeferralStatus()`

UnsetDeferralStatus ensures that no value is present for DeferralStatus, not even an explicit nil
### GetDeferralMinutes

`func (o *HosEvent) GetDeferralMinutes() int32`

GetDeferralMinutes returns the DeferralMinutes field if non-nil, zero value otherwise.

### GetDeferralMinutesOk

`func (o *HosEvent) GetDeferralMinutesOk() (*int32, bool)`

GetDeferralMinutesOk returns a tuple with the DeferralMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferralMinutes

`func (o *HosEvent) SetDeferralMinutes(v int32)`

SetDeferralMinutes sets DeferralMinutes field to given value.

### HasDeferralMinutes

`func (o *HosEvent) HasDeferralMinutes() bool`

HasDeferralMinutes returns a boolean if a field has been set.

### SetDeferralMinutesNil

`func (o *HosEvent) SetDeferralMinutesNil(b bool)`

 SetDeferralMinutesNil sets the value for DeferralMinutes to be an explicit nil

### UnsetDeferralMinutes
`func (o *HosEvent) UnsetDeferralMinutes()`

UnsetDeferralMinutes ensures that no value is present for DeferralMinutes, not even an explicit nil
### GetEldMalfunctionCode

`func (o *HosEvent) GetEldMalfunctionCode() HosMalfunctionCodeEnum`

GetEldMalfunctionCode returns the EldMalfunctionCode field if non-nil, zero value otherwise.

### GetEldMalfunctionCodeOk

`func (o *HosEvent) GetEldMalfunctionCodeOk() (*HosMalfunctionCodeEnum, bool)`

GetEldMalfunctionCodeOk returns a tuple with the EldMalfunctionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldMalfunctionCode

`func (o *HosEvent) SetEldMalfunctionCode(v HosMalfunctionCodeEnum)`

SetEldMalfunctionCode sets EldMalfunctionCode field to given value.

### HasEldMalfunctionCode

`func (o *HosEvent) HasEldMalfunctionCode() bool`

HasEldMalfunctionCode returns a boolean if a field has been set.

### SetEldMalfunctionCodeNil

`func (o *HosEvent) SetEldMalfunctionCodeNil(b bool)`

 SetEldMalfunctionCodeNil sets the value for EldMalfunctionCode to be an explicit nil

### UnsetEldMalfunctionCode
`func (o *HosEvent) UnsetEldMalfunctionCode()`

UnsetEldMalfunctionCode ensures that no value is present for EldMalfunctionCode, not even an explicit nil
### GetIsExcluded

`func (o *HosEvent) GetIsExcluded() bool`

GetIsExcluded returns the IsExcluded field if non-nil, zero value otherwise.

### GetIsExcludedOk

`func (o *HosEvent) GetIsExcludedOk() (*bool, bool)`

GetIsExcludedOk returns a tuple with the IsExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExcluded

`func (o *HosEvent) SetIsExcluded(v bool)`

SetIsExcluded sets IsExcluded field to given value.

### HasIsExcluded

`func (o *HosEvent) HasIsExcluded() bool`

HasIsExcluded returns a boolean if a field has been set.

### SetIsExcludedNil

`func (o *HosEvent) SetIsExcludedNil(b bool)`

 SetIsExcludedNil sets the value for IsExcluded to be an explicit nil

### UnsetIsExcluded
`func (o *HosEvent) UnsetIsExcluded()`

UnsetIsExcluded ensures that no value is present for IsExcluded, not even an explicit nil
### GetIsTransitioning

`func (o *HosEvent) GetIsTransitioning() bool`

GetIsTransitioning returns the IsTransitioning field if non-nil, zero value otherwise.

### GetIsTransitioningOk

`func (o *HosEvent) GetIsTransitioningOk() (*bool, bool)`

GetIsTransitioningOk returns a tuple with the IsTransitioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTransitioning

`func (o *HosEvent) SetIsTransitioning(v bool)`

SetIsTransitioning sets IsTransitioning field to given value.

### HasIsTransitioning

`func (o *HosEvent) HasIsTransitioning() bool`

HasIsTransitioning returns a boolean if a field has been set.

### SetIsTransitioningNil

`func (o *HosEvent) SetIsTransitioningNil(b bool)`

 SetIsTransitioningNil sets the value for IsTransitioning to be an explicit nil

### UnsetIsTransitioning
`func (o *HosEvent) UnsetIsTransitioning()`

UnsetIsTransitioning ensures that no value is present for IsTransitioning, not even an explicit nil
### GetRecordStatus

`func (o *HosEvent) GetRecordStatus() string`

GetRecordStatus returns the RecordStatus field if non-nil, zero value otherwise.

### GetRecordStatusOk

`func (o *HosEvent) GetRecordStatusOk() (*string, bool)`

GetRecordStatusOk returns a tuple with the RecordStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStatus

`func (o *HosEvent) SetRecordStatus(v string)`

SetRecordStatus sets RecordStatus field to given value.

### HasRecordStatus

`func (o *HosEvent) HasRecordStatus() bool`

HasRecordStatus returns a boolean if a field has been set.

### SetRecordStatusNil

`func (o *HosEvent) SetRecordStatusNil(b bool)`

 SetRecordStatusNil sets the value for RecordStatus to be an explicit nil

### UnsetRecordStatus
`func (o *HosEvent) UnsetRecordStatus()`

UnsetRecordStatus ensures that no value is present for RecordStatus, not even an explicit nil
### GetRegionCode

`func (o *HosEvent) GetRegionCode() HosRegionCodeEnum`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *HosEvent) GetRegionCodeOk() (*HosRegionCodeEnum, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *HosEvent) SetRegionCode(v HosRegionCodeEnum)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *HosEvent) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### SetRegionCodeNil

`func (o *HosEvent) SetRegionCodeNil(b bool)`

 SetRegionCodeNil sets the value for RegionCode to be an explicit nil

### UnsetRegionCode
`func (o *HosEvent) UnsetRegionCode()`

UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
### GetLocationName

`func (o *HosEvent) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *HosEvent) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *HosEvent) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *HosEvent) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### SetLocationNameNil

`func (o *HosEvent) SetLocationNameNil(b bool)`

 SetLocationNameNil sets the value for LocationName to be an explicit nil

### UnsetLocationName
`func (o *HosEvent) UnsetLocationName()`

UnsetLocationName ensures that no value is present for LocationName, not even an explicit nil
### GetLocation

`func (o *HosEvent) GetLocation() Location2`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *HosEvent) GetLocationOk() (*Location2, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *HosEvent) SetLocation(v Location2)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *HosEvent) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *HosEvent) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *HosEvent) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetH3Index11

`func (o *HosEvent) GetH3Index11() int32`

GetH3Index11 returns the H3Index11 field if non-nil, zero value otherwise.

### GetH3Index11Ok

`func (o *HosEvent) GetH3Index11Ok() (*int32, bool)`

GetH3Index11Ok returns a tuple with the H3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH3Index11

`func (o *HosEvent) SetH3Index11(v int32)`

SetH3Index11 sets H3Index11 field to given value.

### HasH3Index11

`func (o *HosEvent) HasH3Index11() bool`

HasH3Index11 returns a boolean if a field has been set.

### SetH3Index11Nil

`func (o *HosEvent) SetH3Index11Nil(b bool)`

 SetH3Index11Nil sets the value for H3Index11 to be an explicit nil

### UnsetH3Index11
`func (o *HosEvent) UnsetH3Index11()`

UnsetH3Index11 ensures that no value is present for H3Index11, not even an explicit nil
### GetOdometer

`func (o *HosEvent) GetOdometer() float32`

GetOdometer returns the Odometer field if non-nil, zero value otherwise.

### GetOdometerOk

`func (o *HosEvent) GetOdometerOk() (*float32, bool)`

GetOdometerOk returns a tuple with the Odometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometer

`func (o *HosEvent) SetOdometer(v float32)`

SetOdometer sets Odometer field to given value.

### HasOdometer

`func (o *HosEvent) HasOdometer() bool`

HasOdometer returns a boolean if a field has been set.

### SetOdometerNil

`func (o *HosEvent) SetOdometerNil(b bool)`

 SetOdometerNil sets the value for Odometer to be an explicit nil

### UnsetOdometer
`func (o *HosEvent) UnsetOdometer()`

UnsetOdometer ensures that no value is present for Odometer, not even an explicit nil
### GetEngineHours

`func (o *HosEvent) GetEngineHours() float32`

GetEngineHours returns the EngineHours field if non-nil, zero value otherwise.

### GetEngineHoursOk

`func (o *HosEvent) GetEngineHoursOk() (*float32, bool)`

GetEngineHoursOk returns a tuple with the EngineHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineHours

`func (o *HosEvent) SetEngineHours(v float32)`

SetEngineHours sets EngineHours field to given value.

### HasEngineHours

`func (o *HosEvent) HasEngineHours() bool`

HasEngineHours returns a boolean if a field has been set.

### SetEngineHoursNil

`func (o *HosEvent) SetEngineHoursNil(b bool)`

 SetEngineHoursNil sets the value for EngineHours to be an explicit nil

### UnsetEngineHours
`func (o *HosEvent) UnsetEngineHours()`

UnsetEngineHours ensures that no value is present for EngineHours, not even an explicit nil
### GetSequence

`func (o *HosEvent) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *HosEvent) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *HosEvent) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *HosEvent) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### SetSequenceNil

`func (o *HosEvent) SetSequenceNil(b bool)`

 SetSequenceNil sets the value for Sequence to be an explicit nil

### UnsetSequence
`func (o *HosEvent) UnsetSequence()`

UnsetSequence ensures that no value is present for Sequence, not even an explicit nil
### GetVersion

`func (o *HosEvent) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HosEvent) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HosEvent) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HosEvent) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *HosEvent) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *HosEvent) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


