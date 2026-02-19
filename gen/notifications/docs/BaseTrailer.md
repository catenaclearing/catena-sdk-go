# BaseTrailer

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
**VehicleName** | Pointer to **NullableString** |  | [optional] 
**Oem** | Pointer to **NullableString** |  | [optional] 
**ModelType** | Pointer to **NullableString** |  | [optional] 
**ModelYear** | Pointer to **NullableInt32** |  | [optional] 
**Vin** | Pointer to **NullableString** |  | [optional] 
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
**TrailerGroups** | Pointer to **map[string]interface{}** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**TrailerLength** | Pointer to **NullableFloat32** |  | [optional] 
**TrailerType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBaseTrailer

`func NewBaseTrailer(id string, fleetId string, fleetRef NullableString, sourceName TspEnum, connectionId string, sourceId string, createdAt time.Time, updatedAt time.Time, occurredAt time.Time, ) *BaseTrailer`

NewBaseTrailer instantiates a new BaseTrailer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseTrailerWithDefaults

`func NewBaseTrailerWithDefaults() *BaseTrailer`

NewBaseTrailerWithDefaults instantiates a new BaseTrailer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseTrailer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseTrailer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseTrailer) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseTrailer) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseTrailer) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseTrailer) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *BaseTrailer) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *BaseTrailer) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *BaseTrailer) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### SetFleetRefNil

`func (o *BaseTrailer) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *BaseTrailer) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetTspId

`func (o *BaseTrailer) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *BaseTrailer) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *BaseTrailer) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *BaseTrailer) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *BaseTrailer) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *BaseTrailer) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *BaseTrailer) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *BaseTrailer) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *BaseTrailer) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *BaseTrailer) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *BaseTrailer) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *BaseTrailer) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *BaseTrailer) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseTrailer) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseTrailer) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetConnectionId

`func (o *BaseTrailer) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseTrailer) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseTrailer) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceId

`func (o *BaseTrailer) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseTrailer) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseTrailer) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCreatedAt

`func (o *BaseTrailer) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseTrailer) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseTrailer) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseTrailer) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseTrailer) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseTrailer) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BaseTrailer) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BaseTrailer) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BaseTrailer) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BaseTrailer) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BaseTrailer) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BaseTrailer) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOccurredAt

`func (o *BaseTrailer) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BaseTrailer) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BaseTrailer) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetExecutionId

`func (o *BaseTrailer) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BaseTrailer) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BaseTrailer) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *BaseTrailer) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *BaseTrailer) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *BaseTrailer) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *BaseTrailer) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BaseTrailer) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BaseTrailer) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *BaseTrailer) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *BaseTrailer) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *BaseTrailer) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *BaseTrailer) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *BaseTrailer) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *BaseTrailer) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *BaseTrailer) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *BaseTrailer) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *BaseTrailer) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetVehicleName

`func (o *BaseTrailer) GetVehicleName() string`

GetVehicleName returns the VehicleName field if non-nil, zero value otherwise.

### GetVehicleNameOk

`func (o *BaseTrailer) GetVehicleNameOk() (*string, bool)`

GetVehicleNameOk returns a tuple with the VehicleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleName

`func (o *BaseTrailer) SetVehicleName(v string)`

SetVehicleName sets VehicleName field to given value.

### HasVehicleName

`func (o *BaseTrailer) HasVehicleName() bool`

HasVehicleName returns a boolean if a field has been set.

### SetVehicleNameNil

`func (o *BaseTrailer) SetVehicleNameNil(b bool)`

 SetVehicleNameNil sets the value for VehicleName to be an explicit nil

### UnsetVehicleName
`func (o *BaseTrailer) UnsetVehicleName()`

UnsetVehicleName ensures that no value is present for VehicleName, not even an explicit nil
### GetOem

`func (o *BaseTrailer) GetOem() string`

GetOem returns the Oem field if non-nil, zero value otherwise.

### GetOemOk

`func (o *BaseTrailer) GetOemOk() (*string, bool)`

GetOemOk returns a tuple with the Oem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOem

`func (o *BaseTrailer) SetOem(v string)`

SetOem sets Oem field to given value.

### HasOem

`func (o *BaseTrailer) HasOem() bool`

HasOem returns a boolean if a field has been set.

### SetOemNil

`func (o *BaseTrailer) SetOemNil(b bool)`

 SetOemNil sets the value for Oem to be an explicit nil

### UnsetOem
`func (o *BaseTrailer) UnsetOem()`

UnsetOem ensures that no value is present for Oem, not even an explicit nil
### GetModelType

`func (o *BaseTrailer) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *BaseTrailer) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *BaseTrailer) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *BaseTrailer) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### SetModelTypeNil

`func (o *BaseTrailer) SetModelTypeNil(b bool)`

 SetModelTypeNil sets the value for ModelType to be an explicit nil

### UnsetModelType
`func (o *BaseTrailer) UnsetModelType()`

UnsetModelType ensures that no value is present for ModelType, not even an explicit nil
### GetModelYear

`func (o *BaseTrailer) GetModelYear() int32`

GetModelYear returns the ModelYear field if non-nil, zero value otherwise.

### GetModelYearOk

`func (o *BaseTrailer) GetModelYearOk() (*int32, bool)`

GetModelYearOk returns a tuple with the ModelYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelYear

`func (o *BaseTrailer) SetModelYear(v int32)`

SetModelYear sets ModelYear field to given value.

### HasModelYear

`func (o *BaseTrailer) HasModelYear() bool`

HasModelYear returns a boolean if a field has been set.

### SetModelYearNil

`func (o *BaseTrailer) SetModelYearNil(b bool)`

 SetModelYearNil sets the value for ModelYear to be an explicit nil

### UnsetModelYear
`func (o *BaseTrailer) UnsetModelYear()`

UnsetModelYear ensures that no value is present for ModelYear, not even an explicit nil
### GetVin

`func (o *BaseTrailer) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *BaseTrailer) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *BaseTrailer) SetVin(v string)`

SetVin sets Vin field to given value.

### HasVin

`func (o *BaseTrailer) HasVin() bool`

HasVin returns a boolean if a field has been set.

### SetVinNil

`func (o *BaseTrailer) SetVinNil(b bool)`

 SetVinNil sets the value for Vin to be an explicit nil

### UnsetVin
`func (o *BaseTrailer) UnsetVin()`

UnsetVin ensures that no value is present for Vin, not even an explicit nil
### GetLicensePlateRegion

`func (o *BaseTrailer) GetLicensePlateRegion() string`

GetLicensePlateRegion returns the LicensePlateRegion field if non-nil, zero value otherwise.

### GetLicensePlateRegionOk

`func (o *BaseTrailer) GetLicensePlateRegionOk() (*string, bool)`

GetLicensePlateRegionOk returns a tuple with the LicensePlateRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateRegion

`func (o *BaseTrailer) SetLicensePlateRegion(v string)`

SetLicensePlateRegion sets LicensePlateRegion field to given value.

### HasLicensePlateRegion

`func (o *BaseTrailer) HasLicensePlateRegion() bool`

HasLicensePlateRegion returns a boolean if a field has been set.

### SetLicensePlateRegionNil

`func (o *BaseTrailer) SetLicensePlateRegionNil(b bool)`

 SetLicensePlateRegionNil sets the value for LicensePlateRegion to be an explicit nil

### UnsetLicensePlateRegion
`func (o *BaseTrailer) UnsetLicensePlateRegion()`

UnsetLicensePlateRegion ensures that no value is present for LicensePlateRegion, not even an explicit nil
### GetLicensePlateCountry

`func (o *BaseTrailer) GetLicensePlateCountry() string`

GetLicensePlateCountry returns the LicensePlateCountry field if non-nil, zero value otherwise.

### GetLicensePlateCountryOk

`func (o *BaseTrailer) GetLicensePlateCountryOk() (*string, bool)`

GetLicensePlateCountryOk returns a tuple with the LicensePlateCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateCountry

`func (o *BaseTrailer) SetLicensePlateCountry(v string)`

SetLicensePlateCountry sets LicensePlateCountry field to given value.

### HasLicensePlateCountry

`func (o *BaseTrailer) HasLicensePlateCountry() bool`

HasLicensePlateCountry returns a boolean if a field has been set.

### SetLicensePlateCountryNil

`func (o *BaseTrailer) SetLicensePlateCountryNil(b bool)`

 SetLicensePlateCountryNil sets the value for LicensePlateCountry to be an explicit nil

### UnsetLicensePlateCountry
`func (o *BaseTrailer) UnsetLicensePlateCountry()`

UnsetLicensePlateCountry ensures that no value is present for LicensePlateCountry, not even an explicit nil
### GetLicensePlateNumber

`func (o *BaseTrailer) GetLicensePlateNumber() string`

GetLicensePlateNumber returns the LicensePlateNumber field if non-nil, zero value otherwise.

### GetLicensePlateNumberOk

`func (o *BaseTrailer) GetLicensePlateNumberOk() (*string, bool)`

GetLicensePlateNumberOk returns a tuple with the LicensePlateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateNumber

`func (o *BaseTrailer) SetLicensePlateNumber(v string)`

SetLicensePlateNumber sets LicensePlateNumber field to given value.

### HasLicensePlateNumber

`func (o *BaseTrailer) HasLicensePlateNumber() bool`

HasLicensePlateNumber returns a boolean if a field has been set.

### SetLicensePlateNumberNil

`func (o *BaseTrailer) SetLicensePlateNumberNil(b bool)`

 SetLicensePlateNumberNil sets the value for LicensePlateNumber to be an explicit nil

### UnsetLicensePlateNumber
`func (o *BaseTrailer) UnsetLicensePlateNumber()`

UnsetLicensePlateNumber ensures that no value is present for LicensePlateNumber, not even an explicit nil
### GetStartedAt

`func (o *BaseTrailer) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BaseTrailer) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BaseTrailer) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *BaseTrailer) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *BaseTrailer) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *BaseTrailer) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *BaseTrailer) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *BaseTrailer) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *BaseTrailer) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *BaseTrailer) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *BaseTrailer) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *BaseTrailer) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetIsActive

`func (o *BaseTrailer) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BaseTrailer) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BaseTrailer) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *BaseTrailer) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *BaseTrailer) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *BaseTrailer) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetStatus

`func (o *BaseTrailer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseTrailer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseTrailer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseTrailer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BaseTrailer) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BaseTrailer) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNotes

`func (o *BaseTrailer) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BaseTrailer) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BaseTrailer) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BaseTrailer) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *BaseTrailer) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *BaseTrailer) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetEldId

`func (o *BaseTrailer) GetEldId() string`

GetEldId returns the EldId field if non-nil, zero value otherwise.

### GetEldIdOk

`func (o *BaseTrailer) GetEldIdOk() (*string, bool)`

GetEldIdOk returns a tuple with the EldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldId

`func (o *BaseTrailer) SetEldId(v string)`

SetEldId sets EldId field to given value.

### HasEldId

`func (o *BaseTrailer) HasEldId() bool`

HasEldId returns a boolean if a field has been set.

### SetEldIdNil

`func (o *BaseTrailer) SetEldIdNil(b bool)`

 SetEldIdNil sets the value for EldId to be an explicit nil

### UnsetEldId
`func (o *BaseTrailer) UnsetEldId()`

UnsetEldId ensures that no value is present for EldId, not even an explicit nil
### GetEldSerialNumber

`func (o *BaseTrailer) GetEldSerialNumber() string`

GetEldSerialNumber returns the EldSerialNumber field if non-nil, zero value otherwise.

### GetEldSerialNumberOk

`func (o *BaseTrailer) GetEldSerialNumberOk() (*string, bool)`

GetEldSerialNumberOk returns a tuple with the EldSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldSerialNumber

`func (o *BaseTrailer) SetEldSerialNumber(v string)`

SetEldSerialNumber sets EldSerialNumber field to given value.

### HasEldSerialNumber

`func (o *BaseTrailer) HasEldSerialNumber() bool`

HasEldSerialNumber returns a boolean if a field has been set.

### SetEldSerialNumberNil

`func (o *BaseTrailer) SetEldSerialNumberNil(b bool)`

 SetEldSerialNumberNil sets the value for EldSerialNumber to be an explicit nil

### UnsetEldSerialNumber
`func (o *BaseTrailer) UnsetEldSerialNumber()`

UnsetEldSerialNumber ensures that no value is present for EldSerialNumber, not even an explicit nil
### GetEldDeviceType

`func (o *BaseTrailer) GetEldDeviceType() string`

GetEldDeviceType returns the EldDeviceType field if non-nil, zero value otherwise.

### GetEldDeviceTypeOk

`func (o *BaseTrailer) GetEldDeviceTypeOk() (*string, bool)`

GetEldDeviceTypeOk returns a tuple with the EldDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldDeviceType

`func (o *BaseTrailer) SetEldDeviceType(v string)`

SetEldDeviceType sets EldDeviceType field to given value.

### HasEldDeviceType

`func (o *BaseTrailer) HasEldDeviceType() bool`

HasEldDeviceType returns a boolean if a field has been set.

### SetEldDeviceTypeNil

`func (o *BaseTrailer) SetEldDeviceTypeNil(b bool)`

 SetEldDeviceTypeNil sets the value for EldDeviceType to be an explicit nil

### UnsetEldDeviceType
`func (o *BaseTrailer) UnsetEldDeviceType()`

UnsetEldDeviceType ensures that no value is present for EldDeviceType, not even an explicit nil
### GetEldProductId

`func (o *BaseTrailer) GetEldProductId() string`

GetEldProductId returns the EldProductId field if non-nil, zero value otherwise.

### GetEldProductIdOk

`func (o *BaseTrailer) GetEldProductIdOk() (*string, bool)`

GetEldProductIdOk returns a tuple with the EldProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldProductId

`func (o *BaseTrailer) SetEldProductId(v string)`

SetEldProductId sets EldProductId field to given value.

### HasEldProductId

`func (o *BaseTrailer) HasEldProductId() bool`

HasEldProductId returns a boolean if a field has been set.

### SetEldProductIdNil

`func (o *BaseTrailer) SetEldProductIdNil(b bool)`

 SetEldProductIdNil sets the value for EldProductId to be an explicit nil

### UnsetEldProductId
`func (o *BaseTrailer) UnsetEldProductId()`

UnsetEldProductId ensures that no value is present for EldProductId, not even an explicit nil
### GetTotalAxles

`func (o *BaseTrailer) GetTotalAxles() int32`

GetTotalAxles returns the TotalAxles field if non-nil, zero value otherwise.

### GetTotalAxlesOk

`func (o *BaseTrailer) GetTotalAxlesOk() (*int32, bool)`

GetTotalAxlesOk returns a tuple with the TotalAxles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAxles

`func (o *BaseTrailer) SetTotalAxles(v int32)`

SetTotalAxles sets TotalAxles field to given value.

### HasTotalAxles

`func (o *BaseTrailer) HasTotalAxles() bool`

HasTotalAxles returns a boolean if a field has been set.

### SetTotalAxlesNil

`func (o *BaseTrailer) SetTotalAxlesNil(b bool)`

 SetTotalAxlesNil sets the value for TotalAxles to be an explicit nil

### UnsetTotalAxles
`func (o *BaseTrailer) UnsetTotalAxles()`

UnsetTotalAxles ensures that no value is present for TotalAxles, not even an explicit nil
### GetTrailerGroups

`func (o *BaseTrailer) GetTrailerGroups() map[string]interface{}`

GetTrailerGroups returns the TrailerGroups field if non-nil, zero value otherwise.

### GetTrailerGroupsOk

`func (o *BaseTrailer) GetTrailerGroupsOk() (*map[string]interface{}, bool)`

GetTrailerGroupsOk returns a tuple with the TrailerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerGroups

`func (o *BaseTrailer) SetTrailerGroups(v map[string]interface{})`

SetTrailerGroups sets TrailerGroups field to given value.

### HasTrailerGroups

`func (o *BaseTrailer) HasTrailerGroups() bool`

HasTrailerGroups returns a boolean if a field has been set.

### SetTrailerGroupsNil

`func (o *BaseTrailer) SetTrailerGroupsNil(b bool)`

 SetTrailerGroupsNil sets the value for TrailerGroups to be an explicit nil

### UnsetTrailerGroups
`func (o *BaseTrailer) UnsetTrailerGroups()`

UnsetTrailerGroups ensures that no value is present for TrailerGroups, not even an explicit nil
### GetExternalId

`func (o *BaseTrailer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *BaseTrailer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *BaseTrailer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *BaseTrailer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *BaseTrailer) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *BaseTrailer) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetTrailerLength

`func (o *BaseTrailer) GetTrailerLength() float32`

GetTrailerLength returns the TrailerLength field if non-nil, zero value otherwise.

### GetTrailerLengthOk

`func (o *BaseTrailer) GetTrailerLengthOk() (*float32, bool)`

GetTrailerLengthOk returns a tuple with the TrailerLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerLength

`func (o *BaseTrailer) SetTrailerLength(v float32)`

SetTrailerLength sets TrailerLength field to given value.

### HasTrailerLength

`func (o *BaseTrailer) HasTrailerLength() bool`

HasTrailerLength returns a boolean if a field has been set.

### SetTrailerLengthNil

`func (o *BaseTrailer) SetTrailerLengthNil(b bool)`

 SetTrailerLengthNil sets the value for TrailerLength to be an explicit nil

### UnsetTrailerLength
`func (o *BaseTrailer) UnsetTrailerLength()`

UnsetTrailerLength ensures that no value is present for TrailerLength, not even an explicit nil
### GetTrailerType

`func (o *BaseTrailer) GetTrailerType() string`

GetTrailerType returns the TrailerType field if non-nil, zero value otherwise.

### GetTrailerTypeOk

`func (o *BaseTrailer) GetTrailerTypeOk() (*string, bool)`

GetTrailerTypeOk returns a tuple with the TrailerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerType

`func (o *BaseTrailer) SetTrailerType(v string)`

SetTrailerType sets TrailerType field to given value.

### HasTrailerType

`func (o *BaseTrailer) HasTrailerType() bool`

HasTrailerType returns a boolean if a field has been set.

### SetTrailerTypeNil

`func (o *BaseTrailer) SetTrailerTypeNil(b bool)`

 SetTrailerTypeNil sets the value for TrailerType to be an explicit nil

### UnsetTrailerType
`func (o *BaseTrailer) UnsetTrailerType()`

UnsetTrailerType ensures that no value is present for TrailerType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


