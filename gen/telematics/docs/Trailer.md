# Trailer

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
**TrailerName** | Pointer to **NullableString** |  | [optional] 
**Oem** | Pointer to **NullableString** |  | [optional] 
**ModelType** | Pointer to **NullableString** |  | [optional] 
**ModelYear** | Pointer to **NullableInt32** |  | [optional] 
**Vin** | Pointer to **NullableString** |  | [optional] 
**LicensePlateRegion** | Pointer to **NullableString** |  | [optional] 
**LicensePlateCountry** | Pointer to **NullableString** |  | [optional] 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**EldId** | Pointer to **NullableString** |  | [optional] 
**EldSerialNumber** | Pointer to **NullableString** |  | [optional] 
**EldDeviceType** | Pointer to **NullableString** |  | [optional] 
**EldProductId** | Pointer to **NullableString** |  | [optional] 
**TotalAxles** | Pointer to **NullableInt32** |  | [optional] 
**TrailerGroups** | Pointer to **map[string]interface{}** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**TrailerLength** | Pointer to **NullableFloat32** |  | [optional] 
**TrailerType** | Pointer to **NullableString** |  | [optional] 
**SpeedUnit** | Pointer to **NullableString** |  | [optional] 
**OdometerUnit** | Pointer to **NullableString** |  | [optional] 
**FuelUnit** | Pointer to **NullableString** |  | [optional] 
**FuelCapacity** | Pointer to **NullableString** |  | [optional] 
**EngineType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTrailer

`func NewTrailer(id string, createdAt time.Time, updatedAt time.Time, fleetId string, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *Trailer`

NewTrailer instantiates a new Trailer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailerWithDefaults

`func NewTrailerWithDefaults() *Trailer`

NewTrailerWithDefaults instantiates a new Trailer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Trailer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Trailer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Trailer) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Trailer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Trailer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Trailer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Trailer) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Trailer) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Trailer) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Trailer) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Trailer) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Trailer) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Trailer) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *Trailer) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *Trailer) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetFleetId

`func (o *Trailer) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *Trailer) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *Trailer) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetConnectionId

`func (o *Trailer) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Trailer) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Trailer) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceName

`func (o *Trailer) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *Trailer) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *Trailer) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *Trailer) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *Trailer) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *Trailer) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *Trailer) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *Trailer) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Trailer) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Trailer) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *Trailer) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *Trailer) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *Trailer) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *Trailer) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *Trailer) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *Trailer) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *Trailer) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *Trailer) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *Trailer) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *Trailer) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *Trailer) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *Trailer) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *Trailer) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *Trailer) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *Trailer) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *Trailer) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *Trailer) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *Trailer) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *Trailer) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *Trailer) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *Trailer) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetTrailerName

`func (o *Trailer) GetTrailerName() string`

GetTrailerName returns the TrailerName field if non-nil, zero value otherwise.

### GetTrailerNameOk

`func (o *Trailer) GetTrailerNameOk() (*string, bool)`

GetTrailerNameOk returns a tuple with the TrailerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerName

`func (o *Trailer) SetTrailerName(v string)`

SetTrailerName sets TrailerName field to given value.

### HasTrailerName

`func (o *Trailer) HasTrailerName() bool`

HasTrailerName returns a boolean if a field has been set.

### SetTrailerNameNil

`func (o *Trailer) SetTrailerNameNil(b bool)`

 SetTrailerNameNil sets the value for TrailerName to be an explicit nil

### UnsetTrailerName
`func (o *Trailer) UnsetTrailerName()`

UnsetTrailerName ensures that no value is present for TrailerName, not even an explicit nil
### GetOem

`func (o *Trailer) GetOem() string`

GetOem returns the Oem field if non-nil, zero value otherwise.

### GetOemOk

`func (o *Trailer) GetOemOk() (*string, bool)`

GetOemOk returns a tuple with the Oem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOem

`func (o *Trailer) SetOem(v string)`

SetOem sets Oem field to given value.

### HasOem

`func (o *Trailer) HasOem() bool`

HasOem returns a boolean if a field has been set.

### SetOemNil

`func (o *Trailer) SetOemNil(b bool)`

 SetOemNil sets the value for Oem to be an explicit nil

### UnsetOem
`func (o *Trailer) UnsetOem()`

UnsetOem ensures that no value is present for Oem, not even an explicit nil
### GetModelType

`func (o *Trailer) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *Trailer) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *Trailer) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *Trailer) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### SetModelTypeNil

`func (o *Trailer) SetModelTypeNil(b bool)`

 SetModelTypeNil sets the value for ModelType to be an explicit nil

### UnsetModelType
`func (o *Trailer) UnsetModelType()`

UnsetModelType ensures that no value is present for ModelType, not even an explicit nil
### GetModelYear

`func (o *Trailer) GetModelYear() int32`

GetModelYear returns the ModelYear field if non-nil, zero value otherwise.

### GetModelYearOk

`func (o *Trailer) GetModelYearOk() (*int32, bool)`

GetModelYearOk returns a tuple with the ModelYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelYear

`func (o *Trailer) SetModelYear(v int32)`

SetModelYear sets ModelYear field to given value.

### HasModelYear

`func (o *Trailer) HasModelYear() bool`

HasModelYear returns a boolean if a field has been set.

### SetModelYearNil

`func (o *Trailer) SetModelYearNil(b bool)`

 SetModelYearNil sets the value for ModelYear to be an explicit nil

### UnsetModelYear
`func (o *Trailer) UnsetModelYear()`

UnsetModelYear ensures that no value is present for ModelYear, not even an explicit nil
### GetVin

`func (o *Trailer) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *Trailer) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *Trailer) SetVin(v string)`

SetVin sets Vin field to given value.

### HasVin

`func (o *Trailer) HasVin() bool`

HasVin returns a boolean if a field has been set.

### SetVinNil

`func (o *Trailer) SetVinNil(b bool)`

 SetVinNil sets the value for Vin to be an explicit nil

### UnsetVin
`func (o *Trailer) UnsetVin()`

UnsetVin ensures that no value is present for Vin, not even an explicit nil
### GetLicensePlateRegion

`func (o *Trailer) GetLicensePlateRegion() string`

GetLicensePlateRegion returns the LicensePlateRegion field if non-nil, zero value otherwise.

### GetLicensePlateRegionOk

`func (o *Trailer) GetLicensePlateRegionOk() (*string, bool)`

GetLicensePlateRegionOk returns a tuple with the LicensePlateRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateRegion

`func (o *Trailer) SetLicensePlateRegion(v string)`

SetLicensePlateRegion sets LicensePlateRegion field to given value.

### HasLicensePlateRegion

`func (o *Trailer) HasLicensePlateRegion() bool`

HasLicensePlateRegion returns a boolean if a field has been set.

### SetLicensePlateRegionNil

`func (o *Trailer) SetLicensePlateRegionNil(b bool)`

 SetLicensePlateRegionNil sets the value for LicensePlateRegion to be an explicit nil

### UnsetLicensePlateRegion
`func (o *Trailer) UnsetLicensePlateRegion()`

UnsetLicensePlateRegion ensures that no value is present for LicensePlateRegion, not even an explicit nil
### GetLicensePlateCountry

`func (o *Trailer) GetLicensePlateCountry() string`

GetLicensePlateCountry returns the LicensePlateCountry field if non-nil, zero value otherwise.

### GetLicensePlateCountryOk

`func (o *Trailer) GetLicensePlateCountryOk() (*string, bool)`

GetLicensePlateCountryOk returns a tuple with the LicensePlateCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateCountry

`func (o *Trailer) SetLicensePlateCountry(v string)`

SetLicensePlateCountry sets LicensePlateCountry field to given value.

### HasLicensePlateCountry

`func (o *Trailer) HasLicensePlateCountry() bool`

HasLicensePlateCountry returns a boolean if a field has been set.

### SetLicensePlateCountryNil

`func (o *Trailer) SetLicensePlateCountryNil(b bool)`

 SetLicensePlateCountryNil sets the value for LicensePlateCountry to be an explicit nil

### UnsetLicensePlateCountry
`func (o *Trailer) UnsetLicensePlateCountry()`

UnsetLicensePlateCountry ensures that no value is present for LicensePlateCountry, not even an explicit nil
### GetStartedAt

`func (o *Trailer) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Trailer) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Trailer) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Trailer) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *Trailer) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *Trailer) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *Trailer) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *Trailer) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *Trailer) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *Trailer) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *Trailer) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *Trailer) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetIsActive

`func (o *Trailer) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Trailer) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Trailer) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Trailer) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *Trailer) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *Trailer) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetStatus

`func (o *Trailer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Trailer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Trailer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Trailer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Trailer) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Trailer) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNotes

`func (o *Trailer) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Trailer) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Trailer) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Trailer) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Trailer) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Trailer) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetEldId

`func (o *Trailer) GetEldId() string`

GetEldId returns the EldId field if non-nil, zero value otherwise.

### GetEldIdOk

`func (o *Trailer) GetEldIdOk() (*string, bool)`

GetEldIdOk returns a tuple with the EldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldId

`func (o *Trailer) SetEldId(v string)`

SetEldId sets EldId field to given value.

### HasEldId

`func (o *Trailer) HasEldId() bool`

HasEldId returns a boolean if a field has been set.

### SetEldIdNil

`func (o *Trailer) SetEldIdNil(b bool)`

 SetEldIdNil sets the value for EldId to be an explicit nil

### UnsetEldId
`func (o *Trailer) UnsetEldId()`

UnsetEldId ensures that no value is present for EldId, not even an explicit nil
### GetEldSerialNumber

`func (o *Trailer) GetEldSerialNumber() string`

GetEldSerialNumber returns the EldSerialNumber field if non-nil, zero value otherwise.

### GetEldSerialNumberOk

`func (o *Trailer) GetEldSerialNumberOk() (*string, bool)`

GetEldSerialNumberOk returns a tuple with the EldSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldSerialNumber

`func (o *Trailer) SetEldSerialNumber(v string)`

SetEldSerialNumber sets EldSerialNumber field to given value.

### HasEldSerialNumber

`func (o *Trailer) HasEldSerialNumber() bool`

HasEldSerialNumber returns a boolean if a field has been set.

### SetEldSerialNumberNil

`func (o *Trailer) SetEldSerialNumberNil(b bool)`

 SetEldSerialNumberNil sets the value for EldSerialNumber to be an explicit nil

### UnsetEldSerialNumber
`func (o *Trailer) UnsetEldSerialNumber()`

UnsetEldSerialNumber ensures that no value is present for EldSerialNumber, not even an explicit nil
### GetEldDeviceType

`func (o *Trailer) GetEldDeviceType() string`

GetEldDeviceType returns the EldDeviceType field if non-nil, zero value otherwise.

### GetEldDeviceTypeOk

`func (o *Trailer) GetEldDeviceTypeOk() (*string, bool)`

GetEldDeviceTypeOk returns a tuple with the EldDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldDeviceType

`func (o *Trailer) SetEldDeviceType(v string)`

SetEldDeviceType sets EldDeviceType field to given value.

### HasEldDeviceType

`func (o *Trailer) HasEldDeviceType() bool`

HasEldDeviceType returns a boolean if a field has been set.

### SetEldDeviceTypeNil

`func (o *Trailer) SetEldDeviceTypeNil(b bool)`

 SetEldDeviceTypeNil sets the value for EldDeviceType to be an explicit nil

### UnsetEldDeviceType
`func (o *Trailer) UnsetEldDeviceType()`

UnsetEldDeviceType ensures that no value is present for EldDeviceType, not even an explicit nil
### GetEldProductId

`func (o *Trailer) GetEldProductId() string`

GetEldProductId returns the EldProductId field if non-nil, zero value otherwise.

### GetEldProductIdOk

`func (o *Trailer) GetEldProductIdOk() (*string, bool)`

GetEldProductIdOk returns a tuple with the EldProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldProductId

`func (o *Trailer) SetEldProductId(v string)`

SetEldProductId sets EldProductId field to given value.

### HasEldProductId

`func (o *Trailer) HasEldProductId() bool`

HasEldProductId returns a boolean if a field has been set.

### SetEldProductIdNil

`func (o *Trailer) SetEldProductIdNil(b bool)`

 SetEldProductIdNil sets the value for EldProductId to be an explicit nil

### UnsetEldProductId
`func (o *Trailer) UnsetEldProductId()`

UnsetEldProductId ensures that no value is present for EldProductId, not even an explicit nil
### GetTotalAxles

`func (o *Trailer) GetTotalAxles() int32`

GetTotalAxles returns the TotalAxles field if non-nil, zero value otherwise.

### GetTotalAxlesOk

`func (o *Trailer) GetTotalAxlesOk() (*int32, bool)`

GetTotalAxlesOk returns a tuple with the TotalAxles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAxles

`func (o *Trailer) SetTotalAxles(v int32)`

SetTotalAxles sets TotalAxles field to given value.

### HasTotalAxles

`func (o *Trailer) HasTotalAxles() bool`

HasTotalAxles returns a boolean if a field has been set.

### SetTotalAxlesNil

`func (o *Trailer) SetTotalAxlesNil(b bool)`

 SetTotalAxlesNil sets the value for TotalAxles to be an explicit nil

### UnsetTotalAxles
`func (o *Trailer) UnsetTotalAxles()`

UnsetTotalAxles ensures that no value is present for TotalAxles, not even an explicit nil
### GetTrailerGroups

`func (o *Trailer) GetTrailerGroups() map[string]interface{}`

GetTrailerGroups returns the TrailerGroups field if non-nil, zero value otherwise.

### GetTrailerGroupsOk

`func (o *Trailer) GetTrailerGroupsOk() (*map[string]interface{}, bool)`

GetTrailerGroupsOk returns a tuple with the TrailerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerGroups

`func (o *Trailer) SetTrailerGroups(v map[string]interface{})`

SetTrailerGroups sets TrailerGroups field to given value.

### HasTrailerGroups

`func (o *Trailer) HasTrailerGroups() bool`

HasTrailerGroups returns a boolean if a field has been set.

### SetTrailerGroupsNil

`func (o *Trailer) SetTrailerGroupsNil(b bool)`

 SetTrailerGroupsNil sets the value for TrailerGroups to be an explicit nil

### UnsetTrailerGroups
`func (o *Trailer) UnsetTrailerGroups()`

UnsetTrailerGroups ensures that no value is present for TrailerGroups, not even an explicit nil
### GetExternalId

`func (o *Trailer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Trailer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Trailer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Trailer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *Trailer) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Trailer) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetTrailerLength

`func (o *Trailer) GetTrailerLength() float32`

GetTrailerLength returns the TrailerLength field if non-nil, zero value otherwise.

### GetTrailerLengthOk

`func (o *Trailer) GetTrailerLengthOk() (*float32, bool)`

GetTrailerLengthOk returns a tuple with the TrailerLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerLength

`func (o *Trailer) SetTrailerLength(v float32)`

SetTrailerLength sets TrailerLength field to given value.

### HasTrailerLength

`func (o *Trailer) HasTrailerLength() bool`

HasTrailerLength returns a boolean if a field has been set.

### SetTrailerLengthNil

`func (o *Trailer) SetTrailerLengthNil(b bool)`

 SetTrailerLengthNil sets the value for TrailerLength to be an explicit nil

### UnsetTrailerLength
`func (o *Trailer) UnsetTrailerLength()`

UnsetTrailerLength ensures that no value is present for TrailerLength, not even an explicit nil
### GetTrailerType

`func (o *Trailer) GetTrailerType() string`

GetTrailerType returns the TrailerType field if non-nil, zero value otherwise.

### GetTrailerTypeOk

`func (o *Trailer) GetTrailerTypeOk() (*string, bool)`

GetTrailerTypeOk returns a tuple with the TrailerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerType

`func (o *Trailer) SetTrailerType(v string)`

SetTrailerType sets TrailerType field to given value.

### HasTrailerType

`func (o *Trailer) HasTrailerType() bool`

HasTrailerType returns a boolean if a field has been set.

### SetTrailerTypeNil

`func (o *Trailer) SetTrailerTypeNil(b bool)`

 SetTrailerTypeNil sets the value for TrailerType to be an explicit nil

### UnsetTrailerType
`func (o *Trailer) UnsetTrailerType()`

UnsetTrailerType ensures that no value is present for TrailerType, not even an explicit nil
### GetSpeedUnit

`func (o *Trailer) GetSpeedUnit() string`

GetSpeedUnit returns the SpeedUnit field if non-nil, zero value otherwise.

### GetSpeedUnitOk

`func (o *Trailer) GetSpeedUnitOk() (*string, bool)`

GetSpeedUnitOk returns a tuple with the SpeedUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedUnit

`func (o *Trailer) SetSpeedUnit(v string)`

SetSpeedUnit sets SpeedUnit field to given value.

### HasSpeedUnit

`func (o *Trailer) HasSpeedUnit() bool`

HasSpeedUnit returns a boolean if a field has been set.

### SetSpeedUnitNil

`func (o *Trailer) SetSpeedUnitNil(b bool)`

 SetSpeedUnitNil sets the value for SpeedUnit to be an explicit nil

### UnsetSpeedUnit
`func (o *Trailer) UnsetSpeedUnit()`

UnsetSpeedUnit ensures that no value is present for SpeedUnit, not even an explicit nil
### GetOdometerUnit

`func (o *Trailer) GetOdometerUnit() string`

GetOdometerUnit returns the OdometerUnit field if non-nil, zero value otherwise.

### GetOdometerUnitOk

`func (o *Trailer) GetOdometerUnitOk() (*string, bool)`

GetOdometerUnitOk returns a tuple with the OdometerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerUnit

`func (o *Trailer) SetOdometerUnit(v string)`

SetOdometerUnit sets OdometerUnit field to given value.

### HasOdometerUnit

`func (o *Trailer) HasOdometerUnit() bool`

HasOdometerUnit returns a boolean if a field has been set.

### SetOdometerUnitNil

`func (o *Trailer) SetOdometerUnitNil(b bool)`

 SetOdometerUnitNil sets the value for OdometerUnit to be an explicit nil

### UnsetOdometerUnit
`func (o *Trailer) UnsetOdometerUnit()`

UnsetOdometerUnit ensures that no value is present for OdometerUnit, not even an explicit nil
### GetFuelUnit

`func (o *Trailer) GetFuelUnit() string`

GetFuelUnit returns the FuelUnit field if non-nil, zero value otherwise.

### GetFuelUnitOk

`func (o *Trailer) GetFuelUnitOk() (*string, bool)`

GetFuelUnitOk returns a tuple with the FuelUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelUnit

`func (o *Trailer) SetFuelUnit(v string)`

SetFuelUnit sets FuelUnit field to given value.

### HasFuelUnit

`func (o *Trailer) HasFuelUnit() bool`

HasFuelUnit returns a boolean if a field has been set.

### SetFuelUnitNil

`func (o *Trailer) SetFuelUnitNil(b bool)`

 SetFuelUnitNil sets the value for FuelUnit to be an explicit nil

### UnsetFuelUnit
`func (o *Trailer) UnsetFuelUnit()`

UnsetFuelUnit ensures that no value is present for FuelUnit, not even an explicit nil
### GetFuelCapacity

`func (o *Trailer) GetFuelCapacity() string`

GetFuelCapacity returns the FuelCapacity field if non-nil, zero value otherwise.

### GetFuelCapacityOk

`func (o *Trailer) GetFuelCapacityOk() (*string, bool)`

GetFuelCapacityOk returns a tuple with the FuelCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelCapacity

`func (o *Trailer) SetFuelCapacity(v string)`

SetFuelCapacity sets FuelCapacity field to given value.

### HasFuelCapacity

`func (o *Trailer) HasFuelCapacity() bool`

HasFuelCapacity returns a boolean if a field has been set.

### SetFuelCapacityNil

`func (o *Trailer) SetFuelCapacityNil(b bool)`

 SetFuelCapacityNil sets the value for FuelCapacity to be an explicit nil

### UnsetFuelCapacity
`func (o *Trailer) UnsetFuelCapacity()`

UnsetFuelCapacity ensures that no value is present for FuelCapacity, not even an explicit nil
### GetEngineType

`func (o *Trailer) GetEngineType() string`

GetEngineType returns the EngineType field if non-nil, zero value otherwise.

### GetEngineTypeOk

`func (o *Trailer) GetEngineTypeOk() (*string, bool)`

GetEngineTypeOk returns a tuple with the EngineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineType

`func (o *Trailer) SetEngineType(v string)`

SetEngineType sets EngineType field to given value.

### HasEngineType

`func (o *Trailer) HasEngineType() bool`

HasEngineType returns a boolean if a field has been set.

### SetEngineTypeNil

`func (o *Trailer) SetEngineTypeNil(b bool)`

 SetEngineTypeNil sets the value for EngineType to be an explicit nil

### UnsetEngineType
`func (o *Trailer) UnsetEngineType()`

UnsetEngineType ensures that no value is present for EngineType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


