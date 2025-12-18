# Vehicle

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
**VehicleName** | Pointer to **NullableString** |  | [optional] 
**Oem** | Pointer to **NullableString** |  | [optional] 
**ModelType** | Pointer to **NullableString** |  | [optional] 
**ModelYear** | Pointer to **NullableInt32** |  | [optional] 
**Vin** | Pointer to **NullableString** |  | [optional] 
**EngineVin** | Pointer to **NullableString** |  | [optional] 
**LicensePlateRegion** | Pointer to **NullableString** |  | [optional] 
**LicensePlateCountry** | Pointer to **NullableString** |  | [optional] 
**LicensePlateNumber** | Pointer to **NullableString** |  | [optional] 
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
**VehicleGroups** | Pointer to **map[string]interface{}** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**SpeedUnit** | Pointer to **NullableString** |  | [optional] 
**OdometerUnit** | Pointer to **NullableString** |  | [optional] 
**FuelUnit** | Pointer to **NullableString** |  | [optional] 
**FuelCapacity** | Pointer to **NullableString** |  | [optional] 
**EngineType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVehicle

`func NewVehicle(id string, createdAt time.Time, updatedAt time.Time, fleetId string, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *Vehicle`

NewVehicle instantiates a new Vehicle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleWithDefaults

`func NewVehicleWithDefaults() *Vehicle`

NewVehicleWithDefaults instantiates a new Vehicle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Vehicle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vehicle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vehicle) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Vehicle) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Vehicle) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Vehicle) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Vehicle) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Vehicle) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Vehicle) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *Vehicle) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Vehicle) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Vehicle) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Vehicle) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *Vehicle) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *Vehicle) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetFleetId

`func (o *Vehicle) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *Vehicle) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *Vehicle) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetConnectionId

`func (o *Vehicle) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Vehicle) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Vehicle) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceName

`func (o *Vehicle) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *Vehicle) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *Vehicle) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *Vehicle) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *Vehicle) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *Vehicle) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *Vehicle) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *Vehicle) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *Vehicle) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *Vehicle) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *Vehicle) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *Vehicle) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *Vehicle) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *Vehicle) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *Vehicle) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *Vehicle) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *Vehicle) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *Vehicle) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *Vehicle) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *Vehicle) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *Vehicle) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *Vehicle) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *Vehicle) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *Vehicle) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *Vehicle) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *Vehicle) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *Vehicle) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *Vehicle) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *Vehicle) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *Vehicle) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *Vehicle) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetVehicleName

`func (o *Vehicle) GetVehicleName() string`

GetVehicleName returns the VehicleName field if non-nil, zero value otherwise.

### GetVehicleNameOk

`func (o *Vehicle) GetVehicleNameOk() (*string, bool)`

GetVehicleNameOk returns a tuple with the VehicleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleName

`func (o *Vehicle) SetVehicleName(v string)`

SetVehicleName sets VehicleName field to given value.

### HasVehicleName

`func (o *Vehicle) HasVehicleName() bool`

HasVehicleName returns a boolean if a field has been set.

### SetVehicleNameNil

`func (o *Vehicle) SetVehicleNameNil(b bool)`

 SetVehicleNameNil sets the value for VehicleName to be an explicit nil

### UnsetVehicleName
`func (o *Vehicle) UnsetVehicleName()`

UnsetVehicleName ensures that no value is present for VehicleName, not even an explicit nil
### GetOem

`func (o *Vehicle) GetOem() string`

GetOem returns the Oem field if non-nil, zero value otherwise.

### GetOemOk

`func (o *Vehicle) GetOemOk() (*string, bool)`

GetOemOk returns a tuple with the Oem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOem

`func (o *Vehicle) SetOem(v string)`

SetOem sets Oem field to given value.

### HasOem

`func (o *Vehicle) HasOem() bool`

HasOem returns a boolean if a field has been set.

### SetOemNil

`func (o *Vehicle) SetOemNil(b bool)`

 SetOemNil sets the value for Oem to be an explicit nil

### UnsetOem
`func (o *Vehicle) UnsetOem()`

UnsetOem ensures that no value is present for Oem, not even an explicit nil
### GetModelType

`func (o *Vehicle) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *Vehicle) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *Vehicle) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *Vehicle) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### SetModelTypeNil

`func (o *Vehicle) SetModelTypeNil(b bool)`

 SetModelTypeNil sets the value for ModelType to be an explicit nil

### UnsetModelType
`func (o *Vehicle) UnsetModelType()`

UnsetModelType ensures that no value is present for ModelType, not even an explicit nil
### GetModelYear

`func (o *Vehicle) GetModelYear() int32`

GetModelYear returns the ModelYear field if non-nil, zero value otherwise.

### GetModelYearOk

`func (o *Vehicle) GetModelYearOk() (*int32, bool)`

GetModelYearOk returns a tuple with the ModelYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelYear

`func (o *Vehicle) SetModelYear(v int32)`

SetModelYear sets ModelYear field to given value.

### HasModelYear

`func (o *Vehicle) HasModelYear() bool`

HasModelYear returns a boolean if a field has been set.

### SetModelYearNil

`func (o *Vehicle) SetModelYearNil(b bool)`

 SetModelYearNil sets the value for ModelYear to be an explicit nil

### UnsetModelYear
`func (o *Vehicle) UnsetModelYear()`

UnsetModelYear ensures that no value is present for ModelYear, not even an explicit nil
### GetVin

`func (o *Vehicle) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *Vehicle) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *Vehicle) SetVin(v string)`

SetVin sets Vin field to given value.

### HasVin

`func (o *Vehicle) HasVin() bool`

HasVin returns a boolean if a field has been set.

### SetVinNil

`func (o *Vehicle) SetVinNil(b bool)`

 SetVinNil sets the value for Vin to be an explicit nil

### UnsetVin
`func (o *Vehicle) UnsetVin()`

UnsetVin ensures that no value is present for Vin, not even an explicit nil
### GetEngineVin

`func (o *Vehicle) GetEngineVin() string`

GetEngineVin returns the EngineVin field if non-nil, zero value otherwise.

### GetEngineVinOk

`func (o *Vehicle) GetEngineVinOk() (*string, bool)`

GetEngineVinOk returns a tuple with the EngineVin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVin

`func (o *Vehicle) SetEngineVin(v string)`

SetEngineVin sets EngineVin field to given value.

### HasEngineVin

`func (o *Vehicle) HasEngineVin() bool`

HasEngineVin returns a boolean if a field has been set.

### SetEngineVinNil

`func (o *Vehicle) SetEngineVinNil(b bool)`

 SetEngineVinNil sets the value for EngineVin to be an explicit nil

### UnsetEngineVin
`func (o *Vehicle) UnsetEngineVin()`

UnsetEngineVin ensures that no value is present for EngineVin, not even an explicit nil
### GetLicensePlateRegion

`func (o *Vehicle) GetLicensePlateRegion() string`

GetLicensePlateRegion returns the LicensePlateRegion field if non-nil, zero value otherwise.

### GetLicensePlateRegionOk

`func (o *Vehicle) GetLicensePlateRegionOk() (*string, bool)`

GetLicensePlateRegionOk returns a tuple with the LicensePlateRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateRegion

`func (o *Vehicle) SetLicensePlateRegion(v string)`

SetLicensePlateRegion sets LicensePlateRegion field to given value.

### HasLicensePlateRegion

`func (o *Vehicle) HasLicensePlateRegion() bool`

HasLicensePlateRegion returns a boolean if a field has been set.

### SetLicensePlateRegionNil

`func (o *Vehicle) SetLicensePlateRegionNil(b bool)`

 SetLicensePlateRegionNil sets the value for LicensePlateRegion to be an explicit nil

### UnsetLicensePlateRegion
`func (o *Vehicle) UnsetLicensePlateRegion()`

UnsetLicensePlateRegion ensures that no value is present for LicensePlateRegion, not even an explicit nil
### GetLicensePlateCountry

`func (o *Vehicle) GetLicensePlateCountry() string`

GetLicensePlateCountry returns the LicensePlateCountry field if non-nil, zero value otherwise.

### GetLicensePlateCountryOk

`func (o *Vehicle) GetLicensePlateCountryOk() (*string, bool)`

GetLicensePlateCountryOk returns a tuple with the LicensePlateCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateCountry

`func (o *Vehicle) SetLicensePlateCountry(v string)`

SetLicensePlateCountry sets LicensePlateCountry field to given value.

### HasLicensePlateCountry

`func (o *Vehicle) HasLicensePlateCountry() bool`

HasLicensePlateCountry returns a boolean if a field has been set.

### SetLicensePlateCountryNil

`func (o *Vehicle) SetLicensePlateCountryNil(b bool)`

 SetLicensePlateCountryNil sets the value for LicensePlateCountry to be an explicit nil

### UnsetLicensePlateCountry
`func (o *Vehicle) UnsetLicensePlateCountry()`

UnsetLicensePlateCountry ensures that no value is present for LicensePlateCountry, not even an explicit nil
### GetLicensePlateNumber

`func (o *Vehicle) GetLicensePlateNumber() string`

GetLicensePlateNumber returns the LicensePlateNumber field if non-nil, zero value otherwise.

### GetLicensePlateNumberOk

`func (o *Vehicle) GetLicensePlateNumberOk() (*string, bool)`

GetLicensePlateNumberOk returns a tuple with the LicensePlateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateNumber

`func (o *Vehicle) SetLicensePlateNumber(v string)`

SetLicensePlateNumber sets LicensePlateNumber field to given value.

### HasLicensePlateNumber

`func (o *Vehicle) HasLicensePlateNumber() bool`

HasLicensePlateNumber returns a boolean if a field has been set.

### SetLicensePlateNumberNil

`func (o *Vehicle) SetLicensePlateNumberNil(b bool)`

 SetLicensePlateNumberNil sets the value for LicensePlateNumber to be an explicit nil

### UnsetLicensePlateNumber
`func (o *Vehicle) UnsetLicensePlateNumber()`

UnsetLicensePlateNumber ensures that no value is present for LicensePlateNumber, not even an explicit nil
### GetStartedAt

`func (o *Vehicle) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Vehicle) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Vehicle) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Vehicle) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *Vehicle) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *Vehicle) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *Vehicle) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *Vehicle) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *Vehicle) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *Vehicle) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *Vehicle) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *Vehicle) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetIsActive

`func (o *Vehicle) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Vehicle) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Vehicle) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Vehicle) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *Vehicle) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *Vehicle) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetStatus

`func (o *Vehicle) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Vehicle) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Vehicle) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Vehicle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *Vehicle) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Vehicle) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNotes

`func (o *Vehicle) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Vehicle) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Vehicle) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Vehicle) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *Vehicle) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Vehicle) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetEldId

`func (o *Vehicle) GetEldId() string`

GetEldId returns the EldId field if non-nil, zero value otherwise.

### GetEldIdOk

`func (o *Vehicle) GetEldIdOk() (*string, bool)`

GetEldIdOk returns a tuple with the EldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldId

`func (o *Vehicle) SetEldId(v string)`

SetEldId sets EldId field to given value.

### HasEldId

`func (o *Vehicle) HasEldId() bool`

HasEldId returns a boolean if a field has been set.

### SetEldIdNil

`func (o *Vehicle) SetEldIdNil(b bool)`

 SetEldIdNil sets the value for EldId to be an explicit nil

### UnsetEldId
`func (o *Vehicle) UnsetEldId()`

UnsetEldId ensures that no value is present for EldId, not even an explicit nil
### GetEldSerialNumber

`func (o *Vehicle) GetEldSerialNumber() string`

GetEldSerialNumber returns the EldSerialNumber field if non-nil, zero value otherwise.

### GetEldSerialNumberOk

`func (o *Vehicle) GetEldSerialNumberOk() (*string, bool)`

GetEldSerialNumberOk returns a tuple with the EldSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldSerialNumber

`func (o *Vehicle) SetEldSerialNumber(v string)`

SetEldSerialNumber sets EldSerialNumber field to given value.

### HasEldSerialNumber

`func (o *Vehicle) HasEldSerialNumber() bool`

HasEldSerialNumber returns a boolean if a field has been set.

### SetEldSerialNumberNil

`func (o *Vehicle) SetEldSerialNumberNil(b bool)`

 SetEldSerialNumberNil sets the value for EldSerialNumber to be an explicit nil

### UnsetEldSerialNumber
`func (o *Vehicle) UnsetEldSerialNumber()`

UnsetEldSerialNumber ensures that no value is present for EldSerialNumber, not even an explicit nil
### GetEldDeviceType

`func (o *Vehicle) GetEldDeviceType() string`

GetEldDeviceType returns the EldDeviceType field if non-nil, zero value otherwise.

### GetEldDeviceTypeOk

`func (o *Vehicle) GetEldDeviceTypeOk() (*string, bool)`

GetEldDeviceTypeOk returns a tuple with the EldDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldDeviceType

`func (o *Vehicle) SetEldDeviceType(v string)`

SetEldDeviceType sets EldDeviceType field to given value.

### HasEldDeviceType

`func (o *Vehicle) HasEldDeviceType() bool`

HasEldDeviceType returns a boolean if a field has been set.

### SetEldDeviceTypeNil

`func (o *Vehicle) SetEldDeviceTypeNil(b bool)`

 SetEldDeviceTypeNil sets the value for EldDeviceType to be an explicit nil

### UnsetEldDeviceType
`func (o *Vehicle) UnsetEldDeviceType()`

UnsetEldDeviceType ensures that no value is present for EldDeviceType, not even an explicit nil
### GetEldProductId

`func (o *Vehicle) GetEldProductId() string`

GetEldProductId returns the EldProductId field if non-nil, zero value otherwise.

### GetEldProductIdOk

`func (o *Vehicle) GetEldProductIdOk() (*string, bool)`

GetEldProductIdOk returns a tuple with the EldProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldProductId

`func (o *Vehicle) SetEldProductId(v string)`

SetEldProductId sets EldProductId field to given value.

### HasEldProductId

`func (o *Vehicle) HasEldProductId() bool`

HasEldProductId returns a boolean if a field has been set.

### SetEldProductIdNil

`func (o *Vehicle) SetEldProductIdNil(b bool)`

 SetEldProductIdNil sets the value for EldProductId to be an explicit nil

### UnsetEldProductId
`func (o *Vehicle) UnsetEldProductId()`

UnsetEldProductId ensures that no value is present for EldProductId, not even an explicit nil
### GetTotalAxles

`func (o *Vehicle) GetTotalAxles() int32`

GetTotalAxles returns the TotalAxles field if non-nil, zero value otherwise.

### GetTotalAxlesOk

`func (o *Vehicle) GetTotalAxlesOk() (*int32, bool)`

GetTotalAxlesOk returns a tuple with the TotalAxles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAxles

`func (o *Vehicle) SetTotalAxles(v int32)`

SetTotalAxles sets TotalAxles field to given value.

### HasTotalAxles

`func (o *Vehicle) HasTotalAxles() bool`

HasTotalAxles returns a boolean if a field has been set.

### SetTotalAxlesNil

`func (o *Vehicle) SetTotalAxlesNil(b bool)`

 SetTotalAxlesNil sets the value for TotalAxles to be an explicit nil

### UnsetTotalAxles
`func (o *Vehicle) UnsetTotalAxles()`

UnsetTotalAxles ensures that no value is present for TotalAxles, not even an explicit nil
### GetVehicleGroups

`func (o *Vehicle) GetVehicleGroups() map[string]interface{}`

GetVehicleGroups returns the VehicleGroups field if non-nil, zero value otherwise.

### GetVehicleGroupsOk

`func (o *Vehicle) GetVehicleGroupsOk() (*map[string]interface{}, bool)`

GetVehicleGroupsOk returns a tuple with the VehicleGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleGroups

`func (o *Vehicle) SetVehicleGroups(v map[string]interface{})`

SetVehicleGroups sets VehicleGroups field to given value.

### HasVehicleGroups

`func (o *Vehicle) HasVehicleGroups() bool`

HasVehicleGroups returns a boolean if a field has been set.

### SetVehicleGroupsNil

`func (o *Vehicle) SetVehicleGroupsNil(b bool)`

 SetVehicleGroupsNil sets the value for VehicleGroups to be an explicit nil

### UnsetVehicleGroups
`func (o *Vehicle) UnsetVehicleGroups()`

UnsetVehicleGroups ensures that no value is present for VehicleGroups, not even an explicit nil
### GetExternalId

`func (o *Vehicle) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Vehicle) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Vehicle) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Vehicle) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *Vehicle) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *Vehicle) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetSpeedUnit

`func (o *Vehicle) GetSpeedUnit() string`

GetSpeedUnit returns the SpeedUnit field if non-nil, zero value otherwise.

### GetSpeedUnitOk

`func (o *Vehicle) GetSpeedUnitOk() (*string, bool)`

GetSpeedUnitOk returns a tuple with the SpeedUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedUnit

`func (o *Vehicle) SetSpeedUnit(v string)`

SetSpeedUnit sets SpeedUnit field to given value.

### HasSpeedUnit

`func (o *Vehicle) HasSpeedUnit() bool`

HasSpeedUnit returns a boolean if a field has been set.

### SetSpeedUnitNil

`func (o *Vehicle) SetSpeedUnitNil(b bool)`

 SetSpeedUnitNil sets the value for SpeedUnit to be an explicit nil

### UnsetSpeedUnit
`func (o *Vehicle) UnsetSpeedUnit()`

UnsetSpeedUnit ensures that no value is present for SpeedUnit, not even an explicit nil
### GetOdometerUnit

`func (o *Vehicle) GetOdometerUnit() string`

GetOdometerUnit returns the OdometerUnit field if non-nil, zero value otherwise.

### GetOdometerUnitOk

`func (o *Vehicle) GetOdometerUnitOk() (*string, bool)`

GetOdometerUnitOk returns a tuple with the OdometerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerUnit

`func (o *Vehicle) SetOdometerUnit(v string)`

SetOdometerUnit sets OdometerUnit field to given value.

### HasOdometerUnit

`func (o *Vehicle) HasOdometerUnit() bool`

HasOdometerUnit returns a boolean if a field has been set.

### SetOdometerUnitNil

`func (o *Vehicle) SetOdometerUnitNil(b bool)`

 SetOdometerUnitNil sets the value for OdometerUnit to be an explicit nil

### UnsetOdometerUnit
`func (o *Vehicle) UnsetOdometerUnit()`

UnsetOdometerUnit ensures that no value is present for OdometerUnit, not even an explicit nil
### GetFuelUnit

`func (o *Vehicle) GetFuelUnit() string`

GetFuelUnit returns the FuelUnit field if non-nil, zero value otherwise.

### GetFuelUnitOk

`func (o *Vehicle) GetFuelUnitOk() (*string, bool)`

GetFuelUnitOk returns a tuple with the FuelUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelUnit

`func (o *Vehicle) SetFuelUnit(v string)`

SetFuelUnit sets FuelUnit field to given value.

### HasFuelUnit

`func (o *Vehicle) HasFuelUnit() bool`

HasFuelUnit returns a boolean if a field has been set.

### SetFuelUnitNil

`func (o *Vehicle) SetFuelUnitNil(b bool)`

 SetFuelUnitNil sets the value for FuelUnit to be an explicit nil

### UnsetFuelUnit
`func (o *Vehicle) UnsetFuelUnit()`

UnsetFuelUnit ensures that no value is present for FuelUnit, not even an explicit nil
### GetFuelCapacity

`func (o *Vehicle) GetFuelCapacity() string`

GetFuelCapacity returns the FuelCapacity field if non-nil, zero value otherwise.

### GetFuelCapacityOk

`func (o *Vehicle) GetFuelCapacityOk() (*string, bool)`

GetFuelCapacityOk returns a tuple with the FuelCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelCapacity

`func (o *Vehicle) SetFuelCapacity(v string)`

SetFuelCapacity sets FuelCapacity field to given value.

### HasFuelCapacity

`func (o *Vehicle) HasFuelCapacity() bool`

HasFuelCapacity returns a boolean if a field has been set.

### SetFuelCapacityNil

`func (o *Vehicle) SetFuelCapacityNil(b bool)`

 SetFuelCapacityNil sets the value for FuelCapacity to be an explicit nil

### UnsetFuelCapacity
`func (o *Vehicle) UnsetFuelCapacity()`

UnsetFuelCapacity ensures that no value is present for FuelCapacity, not even an explicit nil
### GetEngineType

`func (o *Vehicle) GetEngineType() string`

GetEngineType returns the EngineType field if non-nil, zero value otherwise.

### GetEngineTypeOk

`func (o *Vehicle) GetEngineTypeOk() (*string, bool)`

GetEngineTypeOk returns a tuple with the EngineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineType

`func (o *Vehicle) SetEngineType(v string)`

SetEngineType sets EngineType field to given value.

### HasEngineType

`func (o *Vehicle) HasEngineType() bool`

HasEngineType returns a boolean if a field has been set.

### SetEngineTypeNil

`func (o *Vehicle) SetEngineTypeNil(b bool)`

 SetEngineTypeNil sets the value for EngineType to be an explicit nil

### UnsetEngineType
`func (o *Vehicle) UnsetEngineType()`

UnsetEngineType ensures that no value is present for EngineType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


