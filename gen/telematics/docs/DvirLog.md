# DvirLog

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

### NewDvirLog

`func NewDvirLog(id string, createdAt time.Time, updatedAt time.Time, fleetId string, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *DvirLog`

NewDvirLog instantiates a new DvirLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDvirLogWithDefaults

`func NewDvirLogWithDefaults() *DvirLog`

NewDvirLogWithDefaults instantiates a new DvirLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DvirLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DvirLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DvirLog) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DvirLog) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DvirLog) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DvirLog) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DvirLog) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DvirLog) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DvirLog) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *DvirLog) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DvirLog) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DvirLog) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DvirLog) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *DvirLog) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *DvirLog) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetFleetId

`func (o *DvirLog) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *DvirLog) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *DvirLog) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetConnectionId

`func (o *DvirLog) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DvirLog) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DvirLog) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceName

`func (o *DvirLog) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *DvirLog) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *DvirLog) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *DvirLog) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *DvirLog) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *DvirLog) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *DvirLog) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *DvirLog) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DvirLog) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DvirLog) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *DvirLog) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *DvirLog) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *DvirLog) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *DvirLog) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *DvirLog) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *DvirLog) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *DvirLog) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *DvirLog) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *DvirLog) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *DvirLog) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *DvirLog) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *DvirLog) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *DvirLog) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *DvirLog) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *DvirLog) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *DvirLog) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *DvirLog) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *DvirLog) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *DvirLog) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *DvirLog) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *DvirLog) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetVehicleId

`func (o *DvirLog) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *DvirLog) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *DvirLog) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *DvirLog) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *DvirLog) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *DvirLog) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetDriverId

`func (o *DvirLog) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *DvirLog) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *DvirLog) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *DvirLog) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *DvirLog) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *DvirLog) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetTspTrailerId

`func (o *DvirLog) GetTspTrailerId() string`

GetTspTrailerId returns the TspTrailerId field if non-nil, zero value otherwise.

### GetTspTrailerIdOk

`func (o *DvirLog) GetTspTrailerIdOk() (*string, bool)`

GetTspTrailerIdOk returns a tuple with the TspTrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspTrailerId

`func (o *DvirLog) SetTspTrailerId(v string)`

SetTspTrailerId sets TspTrailerId field to given value.

### HasTspTrailerId

`func (o *DvirLog) HasTspTrailerId() bool`

HasTspTrailerId returns a boolean if a field has been set.

### SetTspTrailerIdNil

`func (o *DvirLog) SetTspTrailerIdNil(b bool)`

 SetTspTrailerIdNil sets the value for TspTrailerId to be an explicit nil

### UnsetTspTrailerId
`func (o *DvirLog) UnsetTspTrailerId()`

UnsetTspTrailerId ensures that no value is present for TspTrailerId, not even an explicit nil
### GetSourceDriverId

`func (o *DvirLog) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *DvirLog) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *DvirLog) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *DvirLog) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *DvirLog) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *DvirLog) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSourceVehicleId

`func (o *DvirLog) GetSourceVehicleId() string`

GetSourceVehicleId returns the SourceVehicleId field if non-nil, zero value otherwise.

### GetSourceVehicleIdOk

`func (o *DvirLog) GetSourceVehicleIdOk() (*string, bool)`

GetSourceVehicleIdOk returns a tuple with the SourceVehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVehicleId

`func (o *DvirLog) SetSourceVehicleId(v string)`

SetSourceVehicleId sets SourceVehicleId field to given value.

### HasSourceVehicleId

`func (o *DvirLog) HasSourceVehicleId() bool`

HasSourceVehicleId returns a boolean if a field has been set.

### SetSourceVehicleIdNil

`func (o *DvirLog) SetSourceVehicleIdNil(b bool)`

 SetSourceVehicleIdNil sets the value for SourceVehicleId to be an explicit nil

### UnsetSourceVehicleId
`func (o *DvirLog) UnsetSourceVehicleId()`

UnsetSourceVehicleId ensures that no value is present for SourceVehicleId, not even an explicit nil
### GetSourceTspTrailerId

`func (o *DvirLog) GetSourceTspTrailerId() string`

GetSourceTspTrailerId returns the SourceTspTrailerId field if non-nil, zero value otherwise.

### GetSourceTspTrailerIdOk

`func (o *DvirLog) GetSourceTspTrailerIdOk() (*string, bool)`

GetSourceTspTrailerIdOk returns a tuple with the SourceTspTrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTspTrailerId

`func (o *DvirLog) SetSourceTspTrailerId(v string)`

SetSourceTspTrailerId sets SourceTspTrailerId field to given value.

### HasSourceTspTrailerId

`func (o *DvirLog) HasSourceTspTrailerId() bool`

HasSourceTspTrailerId returns a boolean if a field has been set.

### SetSourceTspTrailerIdNil

`func (o *DvirLog) SetSourceTspTrailerIdNil(b bool)`

 SetSourceTspTrailerIdNil sets the value for SourceTspTrailerId to be an explicit nil

### UnsetSourceTspTrailerId
`func (o *DvirLog) UnsetSourceTspTrailerId()`

UnsetSourceTspTrailerId ensures that no value is present for SourceTspTrailerId, not even an explicit nil
### GetLogType

`func (o *DvirLog) GetLogType() string`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *DvirLog) GetLogTypeOk() (*string, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *DvirLog) SetLogType(v string)`

SetLogType sets LogType field to given value.

### HasLogType

`func (o *DvirLog) HasLogType() bool`

HasLogType returns a boolean if a field has been set.

### SetLogTypeNil

`func (o *DvirLog) SetLogTypeNil(b bool)`

 SetLogTypeNil sets the value for LogType to be an explicit nil

### UnsetLogType
`func (o *DvirLog) UnsetLogType()`

UnsetLogType ensures that no value is present for LogType, not even an explicit nil
### GetAuthorityName

`func (o *DvirLog) GetAuthorityName() string`

GetAuthorityName returns the AuthorityName field if non-nil, zero value otherwise.

### GetAuthorityNameOk

`func (o *DvirLog) GetAuthorityNameOk() (*string, bool)`

GetAuthorityNameOk returns a tuple with the AuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityName

`func (o *DvirLog) SetAuthorityName(v string)`

SetAuthorityName sets AuthorityName field to given value.

### HasAuthorityName

`func (o *DvirLog) HasAuthorityName() bool`

HasAuthorityName returns a boolean if a field has been set.

### SetAuthorityNameNil

`func (o *DvirLog) SetAuthorityNameNil(b bool)`

 SetAuthorityNameNil sets the value for AuthorityName to be an explicit nil

### UnsetAuthorityName
`func (o *DvirLog) UnsetAuthorityName()`

UnsetAuthorityName ensures that no value is present for AuthorityName, not even an explicit nil
### GetAuthorityAddress

`func (o *DvirLog) GetAuthorityAddress() string`

GetAuthorityAddress returns the AuthorityAddress field if non-nil, zero value otherwise.

### GetAuthorityAddressOk

`func (o *DvirLog) GetAuthorityAddressOk() (*string, bool)`

GetAuthorityAddressOk returns a tuple with the AuthorityAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityAddress

`func (o *DvirLog) SetAuthorityAddress(v string)`

SetAuthorityAddress sets AuthorityAddress field to given value.

### HasAuthorityAddress

`func (o *DvirLog) HasAuthorityAddress() bool`

HasAuthorityAddress returns a boolean if a field has been set.

### SetAuthorityAddressNil

`func (o *DvirLog) SetAuthorityAddressNil(b bool)`

 SetAuthorityAddressNil sets the value for AuthorityAddress to be an explicit nil

### UnsetAuthorityAddress
`func (o *DvirLog) UnsetAuthorityAddress()`

UnsetAuthorityAddress ensures that no value is present for AuthorityAddress, not even an explicit nil
### GetDurationSeconds

`func (o *DvirLog) GetDurationSeconds() float32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *DvirLog) GetDurationSecondsOk() (*float32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *DvirLog) SetDurationSeconds(v float32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *DvirLog) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *DvirLog) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *DvirLog) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetOdometerKm

`func (o *DvirLog) GetOdometerKm() float32`

GetOdometerKm returns the OdometerKm field if non-nil, zero value otherwise.

### GetOdometerKmOk

`func (o *DvirLog) GetOdometerKmOk() (*float32, bool)`

GetOdometerKmOk returns a tuple with the OdometerKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerKm

`func (o *DvirLog) SetOdometerKm(v float32)`

SetOdometerKm sets OdometerKm field to given value.

### HasOdometerKm

`func (o *DvirLog) HasOdometerKm() bool`

HasOdometerKm returns a boolean if a field has been set.

### SetOdometerKmNil

`func (o *DvirLog) SetOdometerKmNil(b bool)`

 SetOdometerKmNil sets the value for OdometerKm to be an explicit nil

### UnsetOdometerKm
`func (o *DvirLog) UnsetOdometerKm()`

UnsetOdometerKm ensures that no value is present for OdometerKm, not even an explicit nil
### GetEngineHours

`func (o *DvirLog) GetEngineHours() float32`

GetEngineHours returns the EngineHours field if non-nil, zero value otherwise.

### GetEngineHoursOk

`func (o *DvirLog) GetEngineHoursOk() (*float32, bool)`

GetEngineHoursOk returns a tuple with the EngineHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineHours

`func (o *DvirLog) SetEngineHours(v float32)`

SetEngineHours sets EngineHours field to given value.

### HasEngineHours

`func (o *DvirLog) HasEngineHours() bool`

HasEngineHours returns a boolean if a field has been set.

### SetEngineHoursNil

`func (o *DvirLog) SetEngineHoursNil(b bool)`

 SetEngineHoursNil sets the value for EngineHours to be an explicit nil

### UnsetEngineHours
`func (o *DvirLog) UnsetEngineHours()`

UnsetEngineHours ensures that no value is present for EngineHours, not even an explicit nil
### GetLocation

`func (o *DvirLog) GetLocation() Location1`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DvirLog) GetLocationOk() (*Location1, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DvirLog) SetLocation(v Location1)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DvirLog) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *DvirLog) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DvirLog) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetH3Index11

`func (o *DvirLog) GetH3Index11() int32`

GetH3Index11 returns the H3Index11 field if non-nil, zero value otherwise.

### GetH3Index11Ok

`func (o *DvirLog) GetH3Index11Ok() (*int32, bool)`

GetH3Index11Ok returns a tuple with the H3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH3Index11

`func (o *DvirLog) SetH3Index11(v int32)`

SetH3Index11 sets H3Index11 field to given value.

### HasH3Index11

`func (o *DvirLog) HasH3Index11() bool`

HasH3Index11 returns a boolean if a field has been set.

### SetH3Index11Nil

`func (o *DvirLog) SetH3Index11Nil(b bool)`

 SetH3Index11Nil sets the value for H3Index11 to be an explicit nil

### UnsetH3Index11
`func (o *DvirLog) UnsetH3Index11()`

UnsetH3Index11 ensures that no value is present for H3Index11, not even an explicit nil
### GetInspectedBy

`func (o *DvirLog) GetInspectedBy() string`

GetInspectedBy returns the InspectedBy field if non-nil, zero value otherwise.

### GetInspectedByOk

`func (o *DvirLog) GetInspectedByOk() (*string, bool)`

GetInspectedByOk returns a tuple with the InspectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedBy

`func (o *DvirLog) SetInspectedBy(v string)`

SetInspectedBy sets InspectedBy field to given value.

### HasInspectedBy

`func (o *DvirLog) HasInspectedBy() bool`

HasInspectedBy returns a boolean if a field has been set.

### SetInspectedByNil

`func (o *DvirLog) SetInspectedByNil(b bool)`

 SetInspectedByNil sets the value for InspectedBy to be an explicit nil

### UnsetInspectedBy
`func (o *DvirLog) UnsetInspectedBy()`

UnsetInspectedBy ensures that no value is present for InspectedBy, not even an explicit nil
### GetCertifyComment

`func (o *DvirLog) GetCertifyComment() string`

GetCertifyComment returns the CertifyComment field if non-nil, zero value otherwise.

### GetCertifyCommentOk

`func (o *DvirLog) GetCertifyCommentOk() (*string, bool)`

GetCertifyCommentOk returns a tuple with the CertifyComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifyComment

`func (o *DvirLog) SetCertifyComment(v string)`

SetCertifyComment sets CertifyComment field to given value.

### HasCertifyComment

`func (o *DvirLog) HasCertifyComment() bool`

HasCertifyComment returns a boolean if a field has been set.

### SetCertifyCommentNil

`func (o *DvirLog) SetCertifyCommentNil(b bool)`

 SetCertifyCommentNil sets the value for CertifyComment to be an explicit nil

### UnsetCertifyComment
`func (o *DvirLog) UnsetCertifyComment()`

UnsetCertifyComment ensures that no value is present for CertifyComment, not even an explicit nil
### GetDriverComment

`func (o *DvirLog) GetDriverComment() string`

GetDriverComment returns the DriverComment field if non-nil, zero value otherwise.

### GetDriverCommentOk

`func (o *DvirLog) GetDriverCommentOk() (*string, bool)`

GetDriverCommentOk returns a tuple with the DriverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverComment

`func (o *DvirLog) SetDriverComment(v string)`

SetDriverComment sets DriverComment field to given value.

### HasDriverComment

`func (o *DvirLog) HasDriverComment() bool`

HasDriverComment returns a boolean if a field has been set.

### SetDriverCommentNil

`func (o *DvirLog) SetDriverCommentNil(b bool)`

 SetDriverCommentNil sets the value for DriverComment to be an explicit nil

### UnsetDriverComment
`func (o *DvirLog) UnsetDriverComment()`

UnsetDriverComment ensures that no value is present for DriverComment, not even an explicit nil
### GetDefectCount

`func (o *DvirLog) GetDefectCount() int32`

GetDefectCount returns the DefectCount field if non-nil, zero value otherwise.

### GetDefectCountOk

`func (o *DvirLog) GetDefectCountOk() (*int32, bool)`

GetDefectCountOk returns a tuple with the DefectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectCount

`func (o *DvirLog) SetDefectCount(v int32)`

SetDefectCount sets DefectCount field to given value.

### HasDefectCount

`func (o *DvirLog) HasDefectCount() bool`

HasDefectCount returns a boolean if a field has been set.

### SetDefectCountNil

`func (o *DvirLog) SetDefectCountNil(b bool)`

 SetDefectCountNil sets the value for DefectCount to be an explicit nil

### UnsetDefectCount
`func (o *DvirLog) UnsetDefectCount()`

UnsetDefectCount ensures that no value is present for DefectCount, not even an explicit nil
### GetVersion

`func (o *DvirLog) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DvirLog) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DvirLog) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DvirLog) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *DvirLog) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *DvirLog) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetDefects

`func (o *DvirLog) GetDefects() []DvirDefectEnum`

GetDefects returns the Defects field if non-nil, zero value otherwise.

### GetDefectsOk

`func (o *DvirLog) GetDefectsOk() (*[]DvirDefectEnum, bool)`

GetDefectsOk returns a tuple with the Defects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefects

`func (o *DvirLog) SetDefects(v []DvirDefectEnum)`

SetDefects sets Defects field to given value.

### HasDefects

`func (o *DvirLog) HasDefects() bool`

HasDefects returns a boolean if a field has been set.

### SetDefectsNil

`func (o *DvirLog) SetDefectsNil(b bool)`

 SetDefectsNil sets the value for Defects to be an explicit nil

### UnsetDefects
`func (o *DvirLog) UnsetDefects()`

UnsetDefects ensures that no value is present for Defects, not even an explicit nil
### GetIsSafetyCritical

`func (o *DvirLog) GetIsSafetyCritical() bool`

GetIsSafetyCritical returns the IsSafetyCritical field if non-nil, zero value otherwise.

### GetIsSafetyCriticalOk

`func (o *DvirLog) GetIsSafetyCriticalOk() (*bool, bool)`

GetIsSafetyCriticalOk returns a tuple with the IsSafetyCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafetyCritical

`func (o *DvirLog) SetIsSafetyCritical(v bool)`

SetIsSafetyCritical sets IsSafetyCritical field to given value.

### HasIsSafetyCritical

`func (o *DvirLog) HasIsSafetyCritical() bool`

HasIsSafetyCritical returns a boolean if a field has been set.

### SetIsSafetyCriticalNil

`func (o *DvirLog) SetIsSafetyCriticalNil(b bool)`

 SetIsSafetyCriticalNil sets the value for IsSafetyCritical to be an explicit nil

### UnsetIsSafetyCritical
`func (o *DvirLog) UnsetIsSafetyCritical()`

UnsetIsSafetyCritical ensures that no value is present for IsSafetyCritical, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


