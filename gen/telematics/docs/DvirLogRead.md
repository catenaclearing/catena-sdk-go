# DvirLogRead

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
**VehicleId** | Pointer to **NullableString** |  | [optional] 
**DriverId** | Pointer to **NullableString** |  | [optional] 
**TspTrailerId** | Pointer to **NullableString** |  | [optional] 
**SourceDriverId** | Pointer to **NullableString** |  | [optional] 
**SourceVehicleId** | Pointer to **NullableString** |  | [optional] 
**SourceTspTrailerId** | Pointer to **NullableString** |  | [optional] 
**LogType** | Pointer to **NullableString** |  | [optional] 
**AuthorityName** | Pointer to **NullableString** |  | [optional] 
**AuthorityAddress** | Pointer to **NullableString** |  | [optional] 
**DurationSeconds** | Pointer to **NullableFloat32** |  | [optional] 
**OdometerKm** | Pointer to **NullableFloat32** |  | [optional] 
**EngineHours** | Pointer to **NullableFloat32** |  | [optional] 
**Location** | Pointer to [**NullableLocation1**](Location1.md) |  | [optional] 
**H3Index11** | Pointer to **NullableInt32** |  | [optional] 
**InspectedBy** | Pointer to **NullableString** |  | [optional] 
**CertifyComment** | Pointer to **NullableString** |  | [optional] 
**DriverComment** | Pointer to **NullableString** |  | [optional] 
**DefectCount** | Pointer to **NullableInt32** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Defects** | Pointer to [**[]DvirDefectEnum**](DvirDefectEnum.md) |  | [optional] 
**IsSafetyCritical** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewDvirLogRead

`func NewDvirLogRead(fleetId NullableString, id string, createdAt time.Time, updatedAt time.Time, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *DvirLogRead`

NewDvirLogRead instantiates a new DvirLogRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDvirLogReadWithDefaults

`func NewDvirLogReadWithDefaults() *DvirLogRead`

NewDvirLogReadWithDefaults instantiates a new DvirLogRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *DvirLogRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *DvirLogRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *DvirLogRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### SetFleetIdNil

`func (o *DvirLogRead) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *DvirLogRead) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetFleetRef

`func (o *DvirLogRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *DvirLogRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *DvirLogRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *DvirLogRead) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *DvirLogRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *DvirLogRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetId

`func (o *DvirLogRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DvirLogRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DvirLogRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DvirLogRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DvirLogRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DvirLogRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DvirLogRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DvirLogRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DvirLogRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *DvirLogRead) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DvirLogRead) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DvirLogRead) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DvirLogRead) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *DvirLogRead) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *DvirLogRead) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetConnectionId

`func (o *DvirLogRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DvirLogRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DvirLogRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetTspId

`func (o *DvirLogRead) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *DvirLogRead) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *DvirLogRead) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *DvirLogRead) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *DvirLogRead) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *DvirLogRead) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *DvirLogRead) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *DvirLogRead) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *DvirLogRead) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *DvirLogRead) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *DvirLogRead) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *DvirLogRead) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *DvirLogRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *DvirLogRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *DvirLogRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *DvirLogRead) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *DvirLogRead) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *DvirLogRead) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *DvirLogRead) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *DvirLogRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DvirLogRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DvirLogRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *DvirLogRead) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *DvirLogRead) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *DvirLogRead) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *DvirLogRead) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *DvirLogRead) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *DvirLogRead) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *DvirLogRead) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *DvirLogRead) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *DvirLogRead) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *DvirLogRead) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *DvirLogRead) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *DvirLogRead) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *DvirLogRead) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *DvirLogRead) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *DvirLogRead) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *DvirLogRead) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *DvirLogRead) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *DvirLogRead) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *DvirLogRead) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *DvirLogRead) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *DvirLogRead) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *DvirLogRead) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *DvirLogRead) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *DvirLogRead) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *DvirLogRead) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *DvirLogRead) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *DvirLogRead) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetVehicleId

`func (o *DvirLogRead) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *DvirLogRead) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *DvirLogRead) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *DvirLogRead) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *DvirLogRead) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *DvirLogRead) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetDriverId

`func (o *DvirLogRead) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *DvirLogRead) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *DvirLogRead) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *DvirLogRead) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *DvirLogRead) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *DvirLogRead) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetTspTrailerId

`func (o *DvirLogRead) GetTspTrailerId() string`

GetTspTrailerId returns the TspTrailerId field if non-nil, zero value otherwise.

### GetTspTrailerIdOk

`func (o *DvirLogRead) GetTspTrailerIdOk() (*string, bool)`

GetTspTrailerIdOk returns a tuple with the TspTrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspTrailerId

`func (o *DvirLogRead) SetTspTrailerId(v string)`

SetTspTrailerId sets TspTrailerId field to given value.

### HasTspTrailerId

`func (o *DvirLogRead) HasTspTrailerId() bool`

HasTspTrailerId returns a boolean if a field has been set.

### SetTspTrailerIdNil

`func (o *DvirLogRead) SetTspTrailerIdNil(b bool)`

 SetTspTrailerIdNil sets the value for TspTrailerId to be an explicit nil

### UnsetTspTrailerId
`func (o *DvirLogRead) UnsetTspTrailerId()`

UnsetTspTrailerId ensures that no value is present for TspTrailerId, not even an explicit nil
### GetSourceDriverId

`func (o *DvirLogRead) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *DvirLogRead) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *DvirLogRead) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *DvirLogRead) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *DvirLogRead) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *DvirLogRead) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSourceVehicleId

`func (o *DvirLogRead) GetSourceVehicleId() string`

GetSourceVehicleId returns the SourceVehicleId field if non-nil, zero value otherwise.

### GetSourceVehicleIdOk

`func (o *DvirLogRead) GetSourceVehicleIdOk() (*string, bool)`

GetSourceVehicleIdOk returns a tuple with the SourceVehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVehicleId

`func (o *DvirLogRead) SetSourceVehicleId(v string)`

SetSourceVehicleId sets SourceVehicleId field to given value.

### HasSourceVehicleId

`func (o *DvirLogRead) HasSourceVehicleId() bool`

HasSourceVehicleId returns a boolean if a field has been set.

### SetSourceVehicleIdNil

`func (o *DvirLogRead) SetSourceVehicleIdNil(b bool)`

 SetSourceVehicleIdNil sets the value for SourceVehicleId to be an explicit nil

### UnsetSourceVehicleId
`func (o *DvirLogRead) UnsetSourceVehicleId()`

UnsetSourceVehicleId ensures that no value is present for SourceVehicleId, not even an explicit nil
### GetSourceTspTrailerId

`func (o *DvirLogRead) GetSourceTspTrailerId() string`

GetSourceTspTrailerId returns the SourceTspTrailerId field if non-nil, zero value otherwise.

### GetSourceTspTrailerIdOk

`func (o *DvirLogRead) GetSourceTspTrailerIdOk() (*string, bool)`

GetSourceTspTrailerIdOk returns a tuple with the SourceTspTrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTspTrailerId

`func (o *DvirLogRead) SetSourceTspTrailerId(v string)`

SetSourceTspTrailerId sets SourceTspTrailerId field to given value.

### HasSourceTspTrailerId

`func (o *DvirLogRead) HasSourceTspTrailerId() bool`

HasSourceTspTrailerId returns a boolean if a field has been set.

### SetSourceTspTrailerIdNil

`func (o *DvirLogRead) SetSourceTspTrailerIdNil(b bool)`

 SetSourceTspTrailerIdNil sets the value for SourceTspTrailerId to be an explicit nil

### UnsetSourceTspTrailerId
`func (o *DvirLogRead) UnsetSourceTspTrailerId()`

UnsetSourceTspTrailerId ensures that no value is present for SourceTspTrailerId, not even an explicit nil
### GetLogType

`func (o *DvirLogRead) GetLogType() string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *DvirLogRead) GetLogTypeOk() (*string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *DvirLogRead) SetLogType(v string)`

SetLogType sets LogType field to given value.

### HasLogType

`func (o *DvirLogRead) HasLogType() bool`

HasLogType returns a boolean if a field has been set.

### SetLogTypeNil

`func (o *DvirLogRead) SetLogTypeNil(b bool)`

 SetLogTypeNil sets the value for LogType to be an explicit nil

### UnsetLogType
`func (o *DvirLogRead) UnsetLogType()`

UnsetLogType ensures that no value is present for LogType, not even an explicit nil
### GetAuthorityName

`func (o *DvirLogRead) GetAuthorityName() string`

GetAuthorityName returns the AuthorityName field if non-nil, zero value otherwise.

### GetAuthorityNameOk

`func (o *DvirLogRead) GetAuthorityNameOk() (*string, bool)`

GetAuthorityNameOk returns a tuple with the AuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityName

`func (o *DvirLogRead) SetAuthorityName(v string)`

SetAuthorityName sets AuthorityName field to given value.

### HasAuthorityName

`func (o *DvirLogRead) HasAuthorityName() bool`

HasAuthorityName returns a boolean if a field has been set.

### SetAuthorityNameNil

`func (o *DvirLogRead) SetAuthorityNameNil(b bool)`

 SetAuthorityNameNil sets the value for AuthorityName to be an explicit nil

### UnsetAuthorityName
`func (o *DvirLogRead) UnsetAuthorityName()`

UnsetAuthorityName ensures that no value is present for AuthorityName, not even an explicit nil
### GetAuthorityAddress

`func (o *DvirLogRead) GetAuthorityAddress() string`

GetAuthorityAddress returns the AuthorityAddress field if non-nil, zero value otherwise.

### GetAuthorityAddressOk

`func (o *DvirLogRead) GetAuthorityAddressOk() (*string, bool)`

GetAuthorityAddressOk returns a tuple with the AuthorityAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityAddress

`func (o *DvirLogRead) SetAuthorityAddress(v string)`

SetAuthorityAddress sets AuthorityAddress field to given value.

### HasAuthorityAddress

`func (o *DvirLogRead) HasAuthorityAddress() bool`

HasAuthorityAddress returns a boolean if a field has been set.

### SetAuthorityAddressNil

`func (o *DvirLogRead) SetAuthorityAddressNil(b bool)`

 SetAuthorityAddressNil sets the value for AuthorityAddress to be an explicit nil

### UnsetAuthorityAddress
`func (o *DvirLogRead) UnsetAuthorityAddress()`

UnsetAuthorityAddress ensures that no value is present for AuthorityAddress, not even an explicit nil
### GetDurationSeconds

`func (o *DvirLogRead) GetDurationSeconds() float32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *DvirLogRead) GetDurationSecondsOk() (*float32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *DvirLogRead) SetDurationSeconds(v float32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *DvirLogRead) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *DvirLogRead) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *DvirLogRead) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetOdometerKm

`func (o *DvirLogRead) GetOdometerKm() float32`

GetOdometerKm returns the OdometerKm field if non-nil, zero value otherwise.

### GetOdometerKmOk

`func (o *DvirLogRead) GetOdometerKmOk() (*float32, bool)`

GetOdometerKmOk returns a tuple with the OdometerKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerKm

`func (o *DvirLogRead) SetOdometerKm(v float32)`

SetOdometerKm sets OdometerKm field to given value.

### HasOdometerKm

`func (o *DvirLogRead) HasOdometerKm() bool`

HasOdometerKm returns a boolean if a field has been set.

### SetOdometerKmNil

`func (o *DvirLogRead) SetOdometerKmNil(b bool)`

 SetOdometerKmNil sets the value for OdometerKm to be an explicit nil

### UnsetOdometerKm
`func (o *DvirLogRead) UnsetOdometerKm()`

UnsetOdometerKm ensures that no value is present for OdometerKm, not even an explicit nil
### GetEngineHours

`func (o *DvirLogRead) GetEngineHours() float32`

GetEngineHours returns the EngineHours field if non-nil, zero value otherwise.

### GetEngineHoursOk

`func (o *DvirLogRead) GetEngineHoursOk() (*float32, bool)`

GetEngineHoursOk returns a tuple with the EngineHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineHours

`func (o *DvirLogRead) SetEngineHours(v float32)`

SetEngineHours sets EngineHours field to given value.

### HasEngineHours

`func (o *DvirLogRead) HasEngineHours() bool`

HasEngineHours returns a boolean if a field has been set.

### SetEngineHoursNil

`func (o *DvirLogRead) SetEngineHoursNil(b bool)`

 SetEngineHoursNil sets the value for EngineHours to be an explicit nil

### UnsetEngineHours
`func (o *DvirLogRead) UnsetEngineHours()`

UnsetEngineHours ensures that no value is present for EngineHours, not even an explicit nil
### GetLocation

`func (o *DvirLogRead) GetLocation() Location1`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DvirLogRead) GetLocationOk() (*Location1, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DvirLogRead) SetLocation(v Location1)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DvirLogRead) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *DvirLogRead) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DvirLogRead) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetH3Index11

`func (o *DvirLogRead) GetH3Index11() int32`

GetH3Index11 returns the H3Index11 field if non-nil, zero value otherwise.

### GetH3Index11Ok

`func (o *DvirLogRead) GetH3Index11Ok() (*int32, bool)`

GetH3Index11Ok returns a tuple with the H3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH3Index11

`func (o *DvirLogRead) SetH3Index11(v int32)`

SetH3Index11 sets H3Index11 field to given value.

### HasH3Index11

`func (o *DvirLogRead) HasH3Index11() bool`

HasH3Index11 returns a boolean if a field has been set.

### SetH3Index11Nil

`func (o *DvirLogRead) SetH3Index11Nil(b bool)`

 SetH3Index11Nil sets the value for H3Index11 to be an explicit nil

### UnsetH3Index11
`func (o *DvirLogRead) UnsetH3Index11()`

UnsetH3Index11 ensures that no value is present for H3Index11, not even an explicit nil
### GetInspectedBy

`func (o *DvirLogRead) GetInspectedBy() string`

GetInspectedBy returns the InspectedBy field if non-nil, zero value otherwise.

### GetInspectedByOk

`func (o *DvirLogRead) GetInspectedByOk() (*string, bool)`

GetInspectedByOk returns a tuple with the InspectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedBy

`func (o *DvirLogRead) SetInspectedBy(v string)`

SetInspectedBy sets InspectedBy field to given value.

### HasInspectedBy

`func (o *DvirLogRead) HasInspectedBy() bool`

HasInspectedBy returns a boolean if a field has been set.

### SetInspectedByNil

`func (o *DvirLogRead) SetInspectedByNil(b bool)`

 SetInspectedByNil sets the value for InspectedBy to be an explicit nil

### UnsetInspectedBy
`func (o *DvirLogRead) UnsetInspectedBy()`

UnsetInspectedBy ensures that no value is present for InspectedBy, not even an explicit nil
### GetCertifyComment

`func (o *DvirLogRead) GetCertifyComment() string`

GetCertifyComment returns the CertifyComment field if non-nil, zero value otherwise.

### GetCertifyCommentOk

`func (o *DvirLogRead) GetCertifyCommentOk() (*string, bool)`

GetCertifyCommentOk returns a tuple with the CertifyComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifyComment

`func (o *DvirLogRead) SetCertifyComment(v string)`

SetCertifyComment sets CertifyComment field to given value.

### HasCertifyComment

`func (o *DvirLogRead) HasCertifyComment() bool`

HasCertifyComment returns a boolean if a field has been set.

### SetCertifyCommentNil

`func (o *DvirLogRead) SetCertifyCommentNil(b bool)`

 SetCertifyCommentNil sets the value for CertifyComment to be an explicit nil

### UnsetCertifyComment
`func (o *DvirLogRead) UnsetCertifyComment()`

UnsetCertifyComment ensures that no value is present for CertifyComment, not even an explicit nil
### GetDriverComment

`func (o *DvirLogRead) GetDriverComment() string`

GetDriverComment returns the DriverComment field if non-nil, zero value otherwise.

### GetDriverCommentOk

`func (o *DvirLogRead) GetDriverCommentOk() (*string, bool)`

GetDriverCommentOk returns a tuple with the DriverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverComment

`func (o *DvirLogRead) SetDriverComment(v string)`

SetDriverComment sets DriverComment field to given value.

### HasDriverComment

`func (o *DvirLogRead) HasDriverComment() bool`

HasDriverComment returns a boolean if a field has been set.

### SetDriverCommentNil

`func (o *DvirLogRead) SetDriverCommentNil(b bool)`

 SetDriverCommentNil sets the value for DriverComment to be an explicit nil

### UnsetDriverComment
`func (o *DvirLogRead) UnsetDriverComment()`

UnsetDriverComment ensures that no value is present for DriverComment, not even an explicit nil
### GetDefectCount

`func (o *DvirLogRead) GetDefectCount() int32`

GetDefectCount returns the DefectCount field if non-nil, zero value otherwise.

### GetDefectCountOk

`func (o *DvirLogRead) GetDefectCountOk() (*int32, bool)`

GetDefectCountOk returns a tuple with the DefectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectCount

`func (o *DvirLogRead) SetDefectCount(v int32)`

SetDefectCount sets DefectCount field to given value.

### HasDefectCount

`func (o *DvirLogRead) HasDefectCount() bool`

HasDefectCount returns a boolean if a field has been set.

### SetDefectCountNil

`func (o *DvirLogRead) SetDefectCountNil(b bool)`

 SetDefectCountNil sets the value for DefectCount to be an explicit nil

### UnsetDefectCount
`func (o *DvirLogRead) UnsetDefectCount()`

UnsetDefectCount ensures that no value is present for DefectCount, not even an explicit nil
### GetVersion

`func (o *DvirLogRead) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DvirLogRead) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DvirLogRead) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DvirLogRead) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DvirLogRead) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DvirLogRead) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetDefects

`func (o *DvirLogRead) GetDefects() []DvirDefectEnum`

GetDefects returns the Defects field if non-nil, zero value otherwise.

### GetDefectsOk

`func (o *DvirLogRead) GetDefectsOk() (*[]DvirDefectEnum, bool)`

GetDefectsOk returns a tuple with the Defects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefects

`func (o *DvirLogRead) SetDefects(v []DvirDefectEnum)`

SetDefects sets Defects field to given value.

### HasDefects

`func (o *DvirLogRead) HasDefects() bool`

HasDefects returns a boolean if a field has been set.

### SetDefectsNil

`func (o *DvirLogRead) SetDefectsNil(b bool)`

 SetDefectsNil sets the value for Defects to be an explicit nil

### UnsetDefects
`func (o *DvirLogRead) UnsetDefects()`

UnsetDefects ensures that no value is present for Defects, not even an explicit nil
### GetIsSafetyCritical

`func (o *DvirLogRead) GetIsSafetyCritical() bool`

GetIsSafetyCritical returns the IsSafetyCritical field if non-nil, zero value otherwise.

### GetIsSafetyCriticalOk

`func (o *DvirLogRead) GetIsSafetyCriticalOk() (*bool, bool)`

GetIsSafetyCriticalOk returns a tuple with the IsSafetyCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafetyCritical

`func (o *DvirLogRead) SetIsSafetyCritical(v bool)`

SetIsSafetyCritical sets IsSafetyCritical field to given value.

### HasIsSafetyCritical

`func (o *DvirLogRead) HasIsSafetyCritical() bool`

HasIsSafetyCritical returns a boolean if a field has been set.

### SetIsSafetyCriticalNil

`func (o *DvirLogRead) SetIsSafetyCriticalNil(b bool)`

 SetIsSafetyCriticalNil sets the value for IsSafetyCritical to be an explicit nil

### UnsetIsSafetyCritical
`func (o *DvirLogRead) UnsetIsSafetyCritical()`

UnsetIsSafetyCritical ensures that no value is present for IsSafetyCritical, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


