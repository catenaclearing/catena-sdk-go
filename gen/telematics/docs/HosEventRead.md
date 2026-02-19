# HosEventRead

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
**TspId** | Pointer to **NullableString** |  | [optional] 
**TspSlug** | Pointer to **NullableString** |  | [optional] 
**SourceName** | [**TspEnum**](TspEnum.md) | The underlying telematics platform that provided this data (e.g., &#x60;samsara&#x60;, &#x60;motive&#x60;, &#x60;hos247&#x60;). Note: Some platforms like &#x60;hos247&#x60; offer white-labeling, so multiple TSPs may share the same source_name — use &#x60;tsp_id&#x60; or &#x60;tsp_slug&#x60; to identify the specific ELD provider. | 
**SourceData** | Pointer to **map[string]interface{}** | Raw source payload as ingested from the TSP. **Note: use it for audit/debugging.** | [optional] 
**SourceId** | **string** | Unique identifier of the record in the TSP. **Note: we generate a unique composite key based on available fields if the TSP does not provide an unique ID.** | 
**SourceDataHash** | **string** | SHA-256 hash of the source data payload. **Note: we use it internally for idempotence and deduplication.** | 
**OccurredAt** | Pointer to **NullableTime** |  | [optional] 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
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

### NewHosEventRead

`func NewHosEventRead(fleetId NullableString, id string, createdAt time.Time, updatedAt time.Time, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *HosEventRead`

NewHosEventRead instantiates a new HosEventRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHosEventReadWithDefaults

`func NewHosEventReadWithDefaults() *HosEventRead`

NewHosEventReadWithDefaults instantiates a new HosEventRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *HosEventRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *HosEventRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *HosEventRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### SetFleetIdNil

`func (o *HosEventRead) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *HosEventRead) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetFleetRef

`func (o *HosEventRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *HosEventRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *HosEventRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *HosEventRead) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *HosEventRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *HosEventRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetId

`func (o *HosEventRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HosEventRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HosEventRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *HosEventRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HosEventRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HosEventRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *HosEventRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HosEventRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HosEventRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *HosEventRead) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *HosEventRead) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *HosEventRead) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *HosEventRead) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *HosEventRead) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *HosEventRead) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetConnectionId

`func (o *HosEventRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *HosEventRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *HosEventRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetTspId

`func (o *HosEventRead) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *HosEventRead) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *HosEventRead) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *HosEventRead) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *HosEventRead) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *HosEventRead) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *HosEventRead) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *HosEventRead) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *HosEventRead) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *HosEventRead) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *HosEventRead) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *HosEventRead) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *HosEventRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *HosEventRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *HosEventRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *HosEventRead) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *HosEventRead) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *HosEventRead) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *HosEventRead) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *HosEventRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *HosEventRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *HosEventRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *HosEventRead) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *HosEventRead) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *HosEventRead) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *HosEventRead) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *HosEventRead) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *HosEventRead) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *HosEventRead) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *HosEventRead) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *HosEventRead) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *HosEventRead) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HosEventRead) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HosEventRead) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HosEventRead) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HosEventRead) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HosEventRead) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *HosEventRead) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *HosEventRead) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *HosEventRead) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *HosEventRead) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *HosEventRead) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *HosEventRead) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *HosEventRead) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *HosEventRead) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *HosEventRead) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *HosEventRead) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *HosEventRead) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *HosEventRead) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetDriverId

`func (o *HosEventRead) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *HosEventRead) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *HosEventRead) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *HosEventRead) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *HosEventRead) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *HosEventRead) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetVehicleId

`func (o *HosEventRead) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *HosEventRead) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *HosEventRead) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *HosEventRead) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *HosEventRead) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *HosEventRead) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetCoDriverId

`func (o *HosEventRead) GetCoDriverId() string`

GetCoDriverId returns the CoDriverId field if non-nil, zero value otherwise.

### GetCoDriverIdOk

`func (o *HosEventRead) GetCoDriverIdOk() (*string, bool)`

GetCoDriverIdOk returns a tuple with the CoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDriverId

`func (o *HosEventRead) SetCoDriverId(v string)`

SetCoDriverId sets CoDriverId field to given value.

### HasCoDriverId

`func (o *HosEventRead) HasCoDriverId() bool`

HasCoDriverId returns a boolean if a field has been set.

### SetCoDriverIdNil

`func (o *HosEventRead) SetCoDriverIdNil(b bool)`

 SetCoDriverIdNil sets the value for CoDriverId to be an explicit nil

### UnsetCoDriverId
`func (o *HosEventRead) UnsetCoDriverId()`

UnsetCoDriverId ensures that no value is present for CoDriverId, not even an explicit nil
### GetAdditionalDriverIds

`func (o *HosEventRead) GetAdditionalDriverIds() map[string]interface{}`

GetAdditionalDriverIds returns the AdditionalDriverIds field if non-nil, zero value otherwise.

### GetAdditionalDriverIdsOk

`func (o *HosEventRead) GetAdditionalDriverIdsOk() (*map[string]interface{}, bool)`

GetAdditionalDriverIdsOk returns a tuple with the AdditionalDriverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDriverIds

`func (o *HosEventRead) SetAdditionalDriverIds(v map[string]interface{})`

SetAdditionalDriverIds sets AdditionalDriverIds field to given value.

### HasAdditionalDriverIds

`func (o *HosEventRead) HasAdditionalDriverIds() bool`

HasAdditionalDriverIds returns a boolean if a field has been set.

### SetAdditionalDriverIdsNil

`func (o *HosEventRead) SetAdditionalDriverIdsNil(b bool)`

 SetAdditionalDriverIdsNil sets the value for AdditionalDriverIds to be an explicit nil

### UnsetAdditionalDriverIds
`func (o *HosEventRead) UnsetAdditionalDriverIds()`

UnsetAdditionalDriverIds ensures that no value is present for AdditionalDriverIds, not even an explicit nil
### GetStartedAt

`func (o *HosEventRead) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *HosEventRead) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *HosEventRead) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *HosEventRead) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *HosEventRead) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *HosEventRead) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *HosEventRead) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *HosEventRead) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *HosEventRead) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *HosEventRead) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *HosEventRead) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *HosEventRead) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetDriverLastEditAt

`func (o *HosEventRead) GetDriverLastEditAt() time.Time`

GetDriverLastEditAt returns the DriverLastEditAt field if non-nil, zero value otherwise.

### GetDriverLastEditAtOk

`func (o *HosEventRead) GetDriverLastEditAtOk() (*time.Time, bool)`

GetDriverLastEditAtOk returns a tuple with the DriverLastEditAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLastEditAt

`func (o *HosEventRead) SetDriverLastEditAt(v time.Time)`

SetDriverLastEditAt sets DriverLastEditAt field to given value.

### HasDriverLastEditAt

`func (o *HosEventRead) HasDriverLastEditAt() bool`

HasDriverLastEditAt returns a boolean if a field has been set.

### SetDriverLastEditAtNil

`func (o *HosEventRead) SetDriverLastEditAtNil(b bool)`

 SetDriverLastEditAtNil sets the value for DriverLastEditAt to be an explicit nil

### UnsetDriverLastEditAt
`func (o *HosEventRead) UnsetDriverLastEditAt()`

UnsetDriverLastEditAt ensures that no value is present for DriverLastEditAt, not even an explicit nil
### GetDriverCertifiedAt

`func (o *HosEventRead) GetDriverCertifiedAt() time.Time`

GetDriverCertifiedAt returns the DriverCertifiedAt field if non-nil, zero value otherwise.

### GetDriverCertifiedAtOk

`func (o *HosEventRead) GetDriverCertifiedAtOk() (*time.Time, bool)`

GetDriverCertifiedAtOk returns a tuple with the DriverCertifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverCertifiedAt

`func (o *HosEventRead) SetDriverCertifiedAt(v time.Time)`

SetDriverCertifiedAt sets DriverCertifiedAt field to given value.

### HasDriverCertifiedAt

`func (o *HosEventRead) HasDriverCertifiedAt() bool`

HasDriverCertifiedAt returns a boolean if a field has been set.

### SetDriverCertifiedAtNil

`func (o *HosEventRead) SetDriverCertifiedAtNil(b bool)`

 SetDriverCertifiedAtNil sets the value for DriverCertifiedAt to be an explicit nil

### UnsetDriverCertifiedAt
`func (o *HosEventRead) UnsetDriverCertifiedAt()`

UnsetDriverCertifiedAt ensures that no value is present for DriverCertifiedAt, not even an explicit nil
### GetHosRulesetCode

`func (o *HosEventRead) GetHosRulesetCode() HosRulesetCodeEnum`

GetHosRulesetCode returns the HosRulesetCode field if non-nil, zero value otherwise.

### GetHosRulesetCodeOk

`func (o *HosEventRead) GetHosRulesetCodeOk() (*HosRulesetCodeEnum, bool)`

GetHosRulesetCodeOk returns a tuple with the HosRulesetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosRulesetCode

`func (o *HosEventRead) SetHosRulesetCode(v HosRulesetCodeEnum)`

SetHosRulesetCode sets HosRulesetCode field to given value.

### HasHosRulesetCode

`func (o *HosEventRead) HasHosRulesetCode() bool`

HasHosRulesetCode returns a boolean if a field has been set.

### SetHosRulesetCodeNil

`func (o *HosEventRead) SetHosRulesetCodeNil(b bool)`

 SetHosRulesetCodeNil sets the value for HosRulesetCode to be an explicit nil

### UnsetHosRulesetCode
`func (o *HosEventRead) UnsetHosRulesetCode()`

UnsetHosRulesetCode ensures that no value is present for HosRulesetCode, not even an explicit nil
### GetTimeZoneCode

`func (o *HosEventRead) GetTimeZoneCode() TimezoneCodeEnum`

GetTimeZoneCode returns the TimeZoneCode field if non-nil, zero value otherwise.

### GetTimeZoneCodeOk

`func (o *HosEventRead) GetTimeZoneCodeOk() (*TimezoneCodeEnum, bool)`

GetTimeZoneCodeOk returns a tuple with the TimeZoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneCode

`func (o *HosEventRead) SetTimeZoneCode(v TimezoneCodeEnum)`

SetTimeZoneCode sets TimeZoneCode field to given value.

### HasTimeZoneCode

`func (o *HosEventRead) HasTimeZoneCode() bool`

HasTimeZoneCode returns a boolean if a field has been set.

### SetTimeZoneCodeNil

`func (o *HosEventRead) SetTimeZoneCodeNil(b bool)`

 SetTimeZoneCodeNil sets the value for TimeZoneCode to be an explicit nil

### UnsetTimeZoneCode
`func (o *HosEventRead) UnsetTimeZoneCode()`

UnsetTimeZoneCode ensures that no value is present for TimeZoneCode, not even an explicit nil
### GetAnnotations

`func (o *HosEventRead) GetAnnotations() HosEventAnnotation`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *HosEventRead) GetAnnotationsOk() (*HosEventAnnotation, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *HosEventRead) SetAnnotations(v HosEventAnnotation)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *HosEventRead) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.

### SetAnnotationsNil

`func (o *HosEventRead) SetAnnotationsNil(b bool)`

 SetAnnotationsNil sets the value for Annotations to be an explicit nil

### UnsetAnnotations
`func (o *HosEventRead) UnsetAnnotations()`

UnsetAnnotations ensures that no value is present for Annotations, not even an explicit nil
### GetDutyStatusCode

`func (o *HosEventRead) GetDutyStatusCode() DutyStatusCodeEnum`

GetDutyStatusCode returns the DutyStatusCode field if non-nil, zero value otherwise.

### GetDutyStatusCodeOk

`func (o *HosEventRead) GetDutyStatusCodeOk() (*DutyStatusCodeEnum, bool)`

GetDutyStatusCodeOk returns a tuple with the DutyStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyStatusCode

`func (o *HosEventRead) SetDutyStatusCode(v DutyStatusCodeEnum)`

SetDutyStatusCode sets DutyStatusCode field to given value.

### HasDutyStatusCode

`func (o *HosEventRead) HasDutyStatusCode() bool`

HasDutyStatusCode returns a boolean if a field has been set.

### SetDutyStatusCodeNil

`func (o *HosEventRead) SetDutyStatusCodeNil(b bool)`

 SetDutyStatusCodeNil sets the value for DutyStatusCode to be an explicit nil

### UnsetDutyStatusCode
`func (o *HosEventRead) UnsetDutyStatusCode()`

UnsetDutyStatusCode ensures that no value is present for DutyStatusCode, not even an explicit nil
### GetEventTypeCode

`func (o *HosEventRead) GetEventTypeCode() HosEventTypeCodeEnum`

GetEventTypeCode returns the EventTypeCode field if non-nil, zero value otherwise.

### GetEventTypeCodeOk

`func (o *HosEventRead) GetEventTypeCodeOk() (*HosEventTypeCodeEnum, bool)`

GetEventTypeCodeOk returns a tuple with the EventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeCode

`func (o *HosEventRead) SetEventTypeCode(v HosEventTypeCodeEnum)`

SetEventTypeCode sets EventTypeCode field to given value.

### HasEventTypeCode

`func (o *HosEventRead) HasEventTypeCode() bool`

HasEventTypeCode returns a boolean if a field has been set.

### SetEventTypeCodeNil

`func (o *HosEventRead) SetEventTypeCodeNil(b bool)`

 SetEventTypeCodeNil sets the value for EventTypeCode to be an explicit nil

### UnsetEventTypeCode
`func (o *HosEventRead) UnsetEventTypeCode()`

UnsetEventTypeCode ensures that no value is present for EventTypeCode, not even an explicit nil
### GetEventCode

`func (o *HosEventRead) GetEventCode() HosEventCodeEnum`

GetEventCode returns the EventCode field if non-nil, zero value otherwise.

### GetEventCodeOk

`func (o *HosEventRead) GetEventCodeOk() (*HosEventCodeEnum, bool)`

GetEventCodeOk returns a tuple with the EventCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCode

`func (o *HosEventRead) SetEventCode(v HosEventCodeEnum)`

SetEventCode sets EventCode field to given value.

### HasEventCode

`func (o *HosEventRead) HasEventCode() bool`

HasEventCode returns a boolean if a field has been set.

### SetEventCodeNil

`func (o *HosEventRead) SetEventCodeNil(b bool)`

 SetEventCodeNil sets the value for EventCode to be an explicit nil

### UnsetEventCode
`func (o *HosEventRead) UnsetEventCode()`

UnsetEventCode ensures that no value is present for EventCode, not even an explicit nil
### GetLogStateCode

`func (o *HosEventRead) GetLogStateCode() HosRecordStatusCodeEnum`

GetLogStateCode returns the LogStateCode field if non-nil, zero value otherwise.

### GetLogStateCodeOk

`func (o *HosEventRead) GetLogStateCodeOk() (*HosRecordStatusCodeEnum, bool)`

GetLogStateCodeOk returns a tuple with the LogStateCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogStateCode

`func (o *HosEventRead) SetLogStateCode(v HosRecordStatusCodeEnum)`

SetLogStateCode sets LogStateCode field to given value.

### HasLogStateCode

`func (o *HosEventRead) HasLogStateCode() bool`

HasLogStateCode returns a boolean if a field has been set.

### SetLogStateCodeNil

`func (o *HosEventRead) SetLogStateCodeNil(b bool)`

 SetLogStateCodeNil sets the value for LogStateCode to be an explicit nil

### UnsetLogStateCode
`func (o *HosEventRead) UnsetLogStateCode()`

UnsetLogStateCode ensures that no value is present for LogStateCode, not even an explicit nil
### GetLogOriginCode

`func (o *HosEventRead) GetLogOriginCode() HosRecordOriginCodeEnum`

GetLogOriginCode returns the LogOriginCode field if non-nil, zero value otherwise.

### GetLogOriginCodeOk

`func (o *HosEventRead) GetLogOriginCodeOk() (*HosRecordOriginCodeEnum, bool)`

GetLogOriginCodeOk returns a tuple with the LogOriginCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOriginCode

`func (o *HosEventRead) SetLogOriginCode(v HosRecordOriginCodeEnum)`

SetLogOriginCode sets LogOriginCode field to given value.

### HasLogOriginCode

`func (o *HosEventRead) HasLogOriginCode() bool`

HasLogOriginCode returns a boolean if a field has been set.

### SetLogOriginCodeNil

`func (o *HosEventRead) SetLogOriginCodeNil(b bool)`

 SetLogOriginCodeNil sets the value for LogOriginCode to be an explicit nil

### UnsetLogOriginCode
`func (o *HosEventRead) UnsetLogOriginCode()`

UnsetLogOriginCode ensures that no value is present for LogOriginCode, not even an explicit nil
### GetDeferralStatus

`func (o *HosEventRead) GetDeferralStatus() string`

GetDeferralStatus returns the DeferralStatus field if non-nil, zero value otherwise.

### GetDeferralStatusOk

`func (o *HosEventRead) GetDeferralStatusOk() (*string, bool)`

GetDeferralStatusOk returns a tuple with the DeferralStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferralStatus

`func (o *HosEventRead) SetDeferralStatus(v string)`

SetDeferralStatus sets DeferralStatus field to given value.

### HasDeferralStatus

`func (o *HosEventRead) HasDeferralStatus() bool`

HasDeferralStatus returns a boolean if a field has been set.

### SetDeferralStatusNil

`func (o *HosEventRead) SetDeferralStatusNil(b bool)`

 SetDeferralStatusNil sets the value for DeferralStatus to be an explicit nil

### UnsetDeferralStatus
`func (o *HosEventRead) UnsetDeferralStatus()`

UnsetDeferralStatus ensures that no value is present for DeferralStatus, not even an explicit nil
### GetDeferralMinutes

`func (o *HosEventRead) GetDeferralMinutes() int32`

GetDeferralMinutes returns the DeferralMinutes field if non-nil, zero value otherwise.

### GetDeferralMinutesOk

`func (o *HosEventRead) GetDeferralMinutesOk() (*int32, bool)`

GetDeferralMinutesOk returns a tuple with the DeferralMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeferralMinutes

`func (o *HosEventRead) SetDeferralMinutes(v int32)`

SetDeferralMinutes sets DeferralMinutes field to given value.

### HasDeferralMinutes

`func (o *HosEventRead) HasDeferralMinutes() bool`

HasDeferralMinutes returns a boolean if a field has been set.

### SetDeferralMinutesNil

`func (o *HosEventRead) SetDeferralMinutesNil(b bool)`

 SetDeferralMinutesNil sets the value for DeferralMinutes to be an explicit nil

### UnsetDeferralMinutes
`func (o *HosEventRead) UnsetDeferralMinutes()`

UnsetDeferralMinutes ensures that no value is present for DeferralMinutes, not even an explicit nil
### GetEldMalfunctionCode

`func (o *HosEventRead) GetEldMalfunctionCode() HosMalfunctionCodeEnum`

GetEldMalfunctionCode returns the EldMalfunctionCode field if non-nil, zero value otherwise.

### GetEldMalfunctionCodeOk

`func (o *HosEventRead) GetEldMalfunctionCodeOk() (*HosMalfunctionCodeEnum, bool)`

GetEldMalfunctionCodeOk returns a tuple with the EldMalfunctionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldMalfunctionCode

`func (o *HosEventRead) SetEldMalfunctionCode(v HosMalfunctionCodeEnum)`

SetEldMalfunctionCode sets EldMalfunctionCode field to given value.

### HasEldMalfunctionCode

`func (o *HosEventRead) HasEldMalfunctionCode() bool`

HasEldMalfunctionCode returns a boolean if a field has been set.

### SetEldMalfunctionCodeNil

`func (o *HosEventRead) SetEldMalfunctionCodeNil(b bool)`

 SetEldMalfunctionCodeNil sets the value for EldMalfunctionCode to be an explicit nil

### UnsetEldMalfunctionCode
`func (o *HosEventRead) UnsetEldMalfunctionCode()`

UnsetEldMalfunctionCode ensures that no value is present for EldMalfunctionCode, not even an explicit nil
### GetIsExcluded

`func (o *HosEventRead) GetIsExcluded() bool`

GetIsExcluded returns the IsExcluded field if non-nil, zero value otherwise.

### GetIsExcludedOk

`func (o *HosEventRead) GetIsExcludedOk() (*bool, bool)`

GetIsExcludedOk returns a tuple with the IsExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExcluded

`func (o *HosEventRead) SetIsExcluded(v bool)`

SetIsExcluded sets IsExcluded field to given value.

### HasIsExcluded

`func (o *HosEventRead) HasIsExcluded() bool`

HasIsExcluded returns a boolean if a field has been set.

### SetIsExcludedNil

`func (o *HosEventRead) SetIsExcludedNil(b bool)`

 SetIsExcludedNil sets the value for IsExcluded to be an explicit nil

### UnsetIsExcluded
`func (o *HosEventRead) UnsetIsExcluded()`

UnsetIsExcluded ensures that no value is present for IsExcluded, not even an explicit nil
### GetIsTransitioning

`func (o *HosEventRead) GetIsTransitioning() bool`

GetIsTransitioning returns the IsTransitioning field if non-nil, zero value otherwise.

### GetIsTransitioningOk

`func (o *HosEventRead) GetIsTransitioningOk() (*bool, bool)`

GetIsTransitioningOk returns a tuple with the IsTransitioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTransitioning

`func (o *HosEventRead) SetIsTransitioning(v bool)`

SetIsTransitioning sets IsTransitioning field to given value.

### HasIsTransitioning

`func (o *HosEventRead) HasIsTransitioning() bool`

HasIsTransitioning returns a boolean if a field has been set.

### SetIsTransitioningNil

`func (o *HosEventRead) SetIsTransitioningNil(b bool)`

 SetIsTransitioningNil sets the value for IsTransitioning to be an explicit nil

### UnsetIsTransitioning
`func (o *HosEventRead) UnsetIsTransitioning()`

UnsetIsTransitioning ensures that no value is present for IsTransitioning, not even an explicit nil
### GetRecordStatus

`func (o *HosEventRead) GetRecordStatus() string`

GetRecordStatus returns the RecordStatus field if non-nil, zero value otherwise.

### GetRecordStatusOk

`func (o *HosEventRead) GetRecordStatusOk() (*string, bool)`

GetRecordStatusOk returns a tuple with the RecordStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStatus

`func (o *HosEventRead) SetRecordStatus(v string)`

SetRecordStatus sets RecordStatus field to given value.

### HasRecordStatus

`func (o *HosEventRead) HasRecordStatus() bool`

HasRecordStatus returns a boolean if a field has been set.

### SetRecordStatusNil

`func (o *HosEventRead) SetRecordStatusNil(b bool)`

 SetRecordStatusNil sets the value for RecordStatus to be an explicit nil

### UnsetRecordStatus
`func (o *HosEventRead) UnsetRecordStatus()`

UnsetRecordStatus ensures that no value is present for RecordStatus, not even an explicit nil
### GetRegionCode

`func (o *HosEventRead) GetRegionCode() HosRegionCodeEnum`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *HosEventRead) GetRegionCodeOk() (*HosRegionCodeEnum, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *HosEventRead) SetRegionCode(v HosRegionCodeEnum)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *HosEventRead) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### SetRegionCodeNil

`func (o *HosEventRead) SetRegionCodeNil(b bool)`

 SetRegionCodeNil sets the value for RegionCode to be an explicit nil

### UnsetRegionCode
`func (o *HosEventRead) UnsetRegionCode()`

UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
### GetLocationName

`func (o *HosEventRead) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *HosEventRead) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *HosEventRead) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *HosEventRead) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### SetLocationNameNil

`func (o *HosEventRead) SetLocationNameNil(b bool)`

 SetLocationNameNil sets the value for LocationName to be an explicit nil

### UnsetLocationName
`func (o *HosEventRead) UnsetLocationName()`

UnsetLocationName ensures that no value is present for LocationName, not even an explicit nil
### GetLocation

`func (o *HosEventRead) GetLocation() Location2`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *HosEventRead) GetLocationOk() (*Location2, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *HosEventRead) SetLocation(v Location2)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *HosEventRead) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *HosEventRead) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *HosEventRead) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetH3Index11

`func (o *HosEventRead) GetH3Index11() int32`

GetH3Index11 returns the H3Index11 field if non-nil, zero value otherwise.

### GetH3Index11Ok

`func (o *HosEventRead) GetH3Index11Ok() (*int32, bool)`

GetH3Index11Ok returns a tuple with the H3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH3Index11

`func (o *HosEventRead) SetH3Index11(v int32)`

SetH3Index11 sets H3Index11 field to given value.

### HasH3Index11

`func (o *HosEventRead) HasH3Index11() bool`

HasH3Index11 returns a boolean if a field has been set.

### SetH3Index11Nil

`func (o *HosEventRead) SetH3Index11Nil(b bool)`

 SetH3Index11Nil sets the value for H3Index11 to be an explicit nil

### UnsetH3Index11
`func (o *HosEventRead) UnsetH3Index11()`

UnsetH3Index11 ensures that no value is present for H3Index11, not even an explicit nil
### GetOdometer

`func (o *HosEventRead) GetOdometer() float32`

GetOdometer returns the Odometer field if non-nil, zero value otherwise.

### GetOdometerOk

`func (o *HosEventRead) GetOdometerOk() (*float32, bool)`

GetOdometerOk returns a tuple with the Odometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometer

`func (o *HosEventRead) SetOdometer(v float32)`

SetOdometer sets Odometer field to given value.

### HasOdometer

`func (o *HosEventRead) HasOdometer() bool`

HasOdometer returns a boolean if a field has been set.

### SetOdometerNil

`func (o *HosEventRead) SetOdometerNil(b bool)`

 SetOdometerNil sets the value for Odometer to be an explicit nil

### UnsetOdometer
`func (o *HosEventRead) UnsetOdometer()`

UnsetOdometer ensures that no value is present for Odometer, not even an explicit nil
### GetEngineHours

`func (o *HosEventRead) GetEngineHours() float32`

GetEngineHours returns the EngineHours field if non-nil, zero value otherwise.

### GetEngineHoursOk

`func (o *HosEventRead) GetEngineHoursOk() (*float32, bool)`

GetEngineHoursOk returns a tuple with the EngineHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineHours

`func (o *HosEventRead) SetEngineHours(v float32)`

SetEngineHours sets EngineHours field to given value.

### HasEngineHours

`func (o *HosEventRead) HasEngineHours() bool`

HasEngineHours returns a boolean if a field has been set.

### SetEngineHoursNil

`func (o *HosEventRead) SetEngineHoursNil(b bool)`

 SetEngineHoursNil sets the value for EngineHours to be an explicit nil

### UnsetEngineHours
`func (o *HosEventRead) UnsetEngineHours()`

UnsetEngineHours ensures that no value is present for EngineHours, not even an explicit nil
### GetSequence

`func (o *HosEventRead) GetSequence() string`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *HosEventRead) GetSequenceOk() (*string, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *HosEventRead) SetSequence(v string)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *HosEventRead) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### SetSequenceNil

`func (o *HosEventRead) SetSequenceNil(b bool)`

 SetSequenceNil sets the value for Sequence to be an explicit nil

### UnsetSequence
`func (o *HosEventRead) UnsetSequence()`

UnsetSequence ensures that no value is present for Sequence, not even an explicit nil
### GetVersion

`func (o *HosEventRead) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HosEventRead) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HosEventRead) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HosEventRead) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *HosEventRead) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *HosEventRead) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


