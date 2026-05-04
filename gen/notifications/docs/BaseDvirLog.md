# BaseDvirLog

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
**VehicleId** | Pointer to **NullableString** |  | [optional] 
**DriverId** | Pointer to **NullableString** |  | [optional] 
**TrailerId** | Pointer to **NullableString** |  | [optional] 
**SourceDriverId** | Pointer to **NullableString** |  | [optional] 
**SourceVehicleId** | Pointer to **NullableString** |  | [optional] 
**SourceTrailerId** | Pointer to **NullableString** |  | [optional] 
**LogType** | Pointer to [**NullableDvirLogTypeEnum**](DvirLogTypeEnum.md) |  | [optional] 
**AuthorityName** | Pointer to **NullableString** |  | [optional] 
**AuthorityAddress** | Pointer to **NullableString** |  | [optional] 
**DurationSeconds** | Pointer to **NullableFloat32** |  | [optional] 
**OdometerKm** | Pointer to **NullableFloat32** |  | [optional] 
**Odometer** | Pointer to **NullableFloat32** |  | [optional] 
**EngineHours** | Pointer to **NullableFloat32** |  | [optional] 
**Location** | Pointer to [**NullablePoint**](Point.md) |  | [optional] 
**H3Index11** | Pointer to **NullableInt32** |  | [optional] 
**InspectedBy** | Pointer to **NullableString** |  | [optional] 
**CertifyComment** | Pointer to **NullableString** |  | [optional] 
**DriverComment** | Pointer to **NullableString** |  | [optional] 
**DefectCount** | Pointer to **NullableInt32** |  | [optional] 
**Version** | Pointer to **NullableString** |  | [optional] 
**Defects** | Pointer to [**[]DvirDefectEnum**](DvirDefectEnum.md) |  | [optional] 
**IsSafetyCritical** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewBaseDvirLog

`func NewBaseDvirLog(id string, fleetId string, fleetRef NullableString, sourceName TspEnum, connectionId string, sourceId string, createdAt time.Time, updatedAt time.Time, occurredAt time.Time, ) *BaseDvirLog`

NewBaseDvirLog instantiates a new BaseDvirLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseDvirLogWithDefaults

`func NewBaseDvirLogWithDefaults() *BaseDvirLog`

NewBaseDvirLogWithDefaults instantiates a new BaseDvirLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseDvirLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseDvirLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseDvirLog) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseDvirLog) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseDvirLog) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseDvirLog) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *BaseDvirLog) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *BaseDvirLog) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *BaseDvirLog) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### SetFleetRefNil

`func (o *BaseDvirLog) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *BaseDvirLog) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetTspId

`func (o *BaseDvirLog) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *BaseDvirLog) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *BaseDvirLog) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *BaseDvirLog) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *BaseDvirLog) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *BaseDvirLog) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *BaseDvirLog) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *BaseDvirLog) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *BaseDvirLog) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *BaseDvirLog) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *BaseDvirLog) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *BaseDvirLog) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *BaseDvirLog) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseDvirLog) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseDvirLog) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetConnectionId

`func (o *BaseDvirLog) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseDvirLog) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseDvirLog) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceId

`func (o *BaseDvirLog) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseDvirLog) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseDvirLog) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCreatedAt

`func (o *BaseDvirLog) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseDvirLog) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseDvirLog) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseDvirLog) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseDvirLog) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseDvirLog) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BaseDvirLog) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BaseDvirLog) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BaseDvirLog) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BaseDvirLog) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BaseDvirLog) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BaseDvirLog) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOccurredAt

`func (o *BaseDvirLog) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BaseDvirLog) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BaseDvirLog) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetExecutionId

`func (o *BaseDvirLog) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BaseDvirLog) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BaseDvirLog) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *BaseDvirLog) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *BaseDvirLog) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *BaseDvirLog) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *BaseDvirLog) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BaseDvirLog) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BaseDvirLog) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *BaseDvirLog) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *BaseDvirLog) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *BaseDvirLog) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *BaseDvirLog) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *BaseDvirLog) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *BaseDvirLog) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *BaseDvirLog) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *BaseDvirLog) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *BaseDvirLog) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetVehicleId

`func (o *BaseDvirLog) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *BaseDvirLog) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *BaseDvirLog) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *BaseDvirLog) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *BaseDvirLog) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *BaseDvirLog) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetDriverId

`func (o *BaseDvirLog) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *BaseDvirLog) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *BaseDvirLog) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *BaseDvirLog) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *BaseDvirLog) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *BaseDvirLog) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetTrailerId

`func (o *BaseDvirLog) GetTrailerId() string`

GetTrailerId returns the TrailerId field if non-nil, zero value otherwise.

### GetTrailerIdOk

`func (o *BaseDvirLog) GetTrailerIdOk() (*string, bool)`

GetTrailerIdOk returns a tuple with the TrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerId

`func (o *BaseDvirLog) SetTrailerId(v string)`

SetTrailerId sets TrailerId field to given value.

### HasTrailerId

`func (o *BaseDvirLog) HasTrailerId() bool`

HasTrailerId returns a boolean if a field has been set.

### SetTrailerIdNil

`func (o *BaseDvirLog) SetTrailerIdNil(b bool)`

 SetTrailerIdNil sets the value for TrailerId to be an explicit nil

### UnsetTrailerId
`func (o *BaseDvirLog) UnsetTrailerId()`

UnsetTrailerId ensures that no value is present for TrailerId, not even an explicit nil
### GetSourceDriverId

`func (o *BaseDvirLog) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *BaseDvirLog) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *BaseDvirLog) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *BaseDvirLog) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *BaseDvirLog) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *BaseDvirLog) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSourceVehicleId

`func (o *BaseDvirLog) GetSourceVehicleId() string`

GetSourceVehicleId returns the SourceVehicleId field if non-nil, zero value otherwise.

### GetSourceVehicleIdOk

`func (o *BaseDvirLog) GetSourceVehicleIdOk() (*string, bool)`

GetSourceVehicleIdOk returns a tuple with the SourceVehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVehicleId

`func (o *BaseDvirLog) SetSourceVehicleId(v string)`

SetSourceVehicleId sets SourceVehicleId field to given value.

### HasSourceVehicleId

`func (o *BaseDvirLog) HasSourceVehicleId() bool`

HasSourceVehicleId returns a boolean if a field has been set.

### SetSourceVehicleIdNil

`func (o *BaseDvirLog) SetSourceVehicleIdNil(b bool)`

 SetSourceVehicleIdNil sets the value for SourceVehicleId to be an explicit nil

### UnsetSourceVehicleId
`func (o *BaseDvirLog) UnsetSourceVehicleId()`

UnsetSourceVehicleId ensures that no value is present for SourceVehicleId, not even an explicit nil
### GetSourceTrailerId

`func (o *BaseDvirLog) GetSourceTrailerId() string`

GetSourceTrailerId returns the SourceTrailerId field if non-nil, zero value otherwise.

### GetSourceTrailerIdOk

`func (o *BaseDvirLog) GetSourceTrailerIdOk() (*string, bool)`

GetSourceTrailerIdOk returns a tuple with the SourceTrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTrailerId

`func (o *BaseDvirLog) SetSourceTrailerId(v string)`

SetSourceTrailerId sets SourceTrailerId field to given value.

### HasSourceTrailerId

`func (o *BaseDvirLog) HasSourceTrailerId() bool`

HasSourceTrailerId returns a boolean if a field has been set.

### SetSourceTrailerIdNil

`func (o *BaseDvirLog) SetSourceTrailerIdNil(b bool)`

 SetSourceTrailerIdNil sets the value for SourceTrailerId to be an explicit nil

### UnsetSourceTrailerId
`func (o *BaseDvirLog) UnsetSourceTrailerId()`

UnsetSourceTrailerId ensures that no value is present for SourceTrailerId, not even an explicit nil
### GetLogType

`func (o *BaseDvirLog) GetLogType() DvirLogTypeEnum`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *BaseDvirLog) GetLogTypeOk() (*DvirLogTypeEnum, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *BaseDvirLog) SetLogType(v DvirLogTypeEnum)`

SetLogType sets LogType field to given value.

### HasLogType

`func (o *BaseDvirLog) HasLogType() bool`

HasLogType returns a boolean if a field has been set.

### SetLogTypeNil

`func (o *BaseDvirLog) SetLogTypeNil(b bool)`

 SetLogTypeNil sets the value for LogType to be an explicit nil

### UnsetLogType
`func (o *BaseDvirLog) UnsetLogType()`

UnsetLogType ensures that no value is present for LogType, not even an explicit nil
### GetAuthorityName

`func (o *BaseDvirLog) GetAuthorityName() string`

GetAuthorityName returns the AuthorityName field if non-nil, zero value otherwise.

### GetAuthorityNameOk

`func (o *BaseDvirLog) GetAuthorityNameOk() (*string, bool)`

GetAuthorityNameOk returns a tuple with the AuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityName

`func (o *BaseDvirLog) SetAuthorityName(v string)`

SetAuthorityName sets AuthorityName field to given value.

### HasAuthorityName

`func (o *BaseDvirLog) HasAuthorityName() bool`

HasAuthorityName returns a boolean if a field has been set.

### SetAuthorityNameNil

`func (o *BaseDvirLog) SetAuthorityNameNil(b bool)`

 SetAuthorityNameNil sets the value for AuthorityName to be an explicit nil

### UnsetAuthorityName
`func (o *BaseDvirLog) UnsetAuthorityName()`

UnsetAuthorityName ensures that no value is present for AuthorityName, not even an explicit nil
### GetAuthorityAddress

`func (o *BaseDvirLog) GetAuthorityAddress() string`

GetAuthorityAddress returns the AuthorityAddress field if non-nil, zero value otherwise.

### GetAuthorityAddressOk

`func (o *BaseDvirLog) GetAuthorityAddressOk() (*string, bool)`

GetAuthorityAddressOk returns a tuple with the AuthorityAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityAddress

`func (o *BaseDvirLog) SetAuthorityAddress(v string)`

SetAuthorityAddress sets AuthorityAddress field to given value.

### HasAuthorityAddress

`func (o *BaseDvirLog) HasAuthorityAddress() bool`

HasAuthorityAddress returns a boolean if a field has been set.

### SetAuthorityAddressNil

`func (o *BaseDvirLog) SetAuthorityAddressNil(b bool)`

 SetAuthorityAddressNil sets the value for AuthorityAddress to be an explicit nil

### UnsetAuthorityAddress
`func (o *BaseDvirLog) UnsetAuthorityAddress()`

UnsetAuthorityAddress ensures that no value is present for AuthorityAddress, not even an explicit nil
### GetDurationSeconds

`func (o *BaseDvirLog) GetDurationSeconds() float32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *BaseDvirLog) GetDurationSecondsOk() (*float32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *BaseDvirLog) SetDurationSeconds(v float32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *BaseDvirLog) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *BaseDvirLog) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *BaseDvirLog) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetOdometerKm

`func (o *BaseDvirLog) GetOdometerKm() float32`

GetOdometerKm returns the OdometerKm field if non-nil, zero value otherwise.

### GetOdometerKmOk

`func (o *BaseDvirLog) GetOdometerKmOk() (*float32, bool)`

GetOdometerKmOk returns a tuple with the OdometerKm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerKm

`func (o *BaseDvirLog) SetOdometerKm(v float32)`

SetOdometerKm sets OdometerKm field to given value.

### HasOdometerKm

`func (o *BaseDvirLog) HasOdometerKm() bool`

HasOdometerKm returns a boolean if a field has been set.

### SetOdometerKmNil

`func (o *BaseDvirLog) SetOdometerKmNil(b bool)`

 SetOdometerKmNil sets the value for OdometerKm to be an explicit nil

### UnsetOdometerKm
`func (o *BaseDvirLog) UnsetOdometerKm()`

UnsetOdometerKm ensures that no value is present for OdometerKm, not even an explicit nil
### GetOdometer

`func (o *BaseDvirLog) GetOdometer() float32`

GetOdometer returns the Odometer field if non-nil, zero value otherwise.

### GetOdometerOk

`func (o *BaseDvirLog) GetOdometerOk() (*float32, bool)`

GetOdometerOk returns a tuple with the Odometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometer

`func (o *BaseDvirLog) SetOdometer(v float32)`

SetOdometer sets Odometer field to given value.

### HasOdometer

`func (o *BaseDvirLog) HasOdometer() bool`

HasOdometer returns a boolean if a field has been set.

### SetOdometerNil

`func (o *BaseDvirLog) SetOdometerNil(b bool)`

 SetOdometerNil sets the value for Odometer to be an explicit nil

### UnsetOdometer
`func (o *BaseDvirLog) UnsetOdometer()`

UnsetOdometer ensures that no value is present for Odometer, not even an explicit nil
### GetEngineHours

`func (o *BaseDvirLog) GetEngineHours() float32`

GetEngineHours returns the EngineHours field if non-nil, zero value otherwise.

### GetEngineHoursOk

`func (o *BaseDvirLog) GetEngineHoursOk() (*float32, bool)`

GetEngineHoursOk returns a tuple with the EngineHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineHours

`func (o *BaseDvirLog) SetEngineHours(v float32)`

SetEngineHours sets EngineHours field to given value.

### HasEngineHours

`func (o *BaseDvirLog) HasEngineHours() bool`

HasEngineHours returns a boolean if a field has been set.

### SetEngineHoursNil

`func (o *BaseDvirLog) SetEngineHoursNil(b bool)`

 SetEngineHoursNil sets the value for EngineHours to be an explicit nil

### UnsetEngineHours
`func (o *BaseDvirLog) UnsetEngineHours()`

UnsetEngineHours ensures that no value is present for EngineHours, not even an explicit nil
### GetLocation

`func (o *BaseDvirLog) GetLocation() Point`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BaseDvirLog) GetLocationOk() (*Point, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BaseDvirLog) SetLocation(v Point)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BaseDvirLog) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *BaseDvirLog) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *BaseDvirLog) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetH3Index11

`func (o *BaseDvirLog) GetH3Index11() int32`

GetH3Index11 returns the H3Index11 field if non-nil, zero value otherwise.

### GetH3Index11Ok

`func (o *BaseDvirLog) GetH3Index11Ok() (*int32, bool)`

GetH3Index11Ok returns a tuple with the H3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH3Index11

`func (o *BaseDvirLog) SetH3Index11(v int32)`

SetH3Index11 sets H3Index11 field to given value.

### HasH3Index11

`func (o *BaseDvirLog) HasH3Index11() bool`

HasH3Index11 returns a boolean if a field has been set.

### SetH3Index11Nil

`func (o *BaseDvirLog) SetH3Index11Nil(b bool)`

 SetH3Index11Nil sets the value for H3Index11 to be an explicit nil

### UnsetH3Index11
`func (o *BaseDvirLog) UnsetH3Index11()`

UnsetH3Index11 ensures that no value is present for H3Index11, not even an explicit nil
### GetInspectedBy

`func (o *BaseDvirLog) GetInspectedBy() string`

GetInspectedBy returns the InspectedBy field if non-nil, zero value otherwise.

### GetInspectedByOk

`func (o *BaseDvirLog) GetInspectedByOk() (*string, bool)`

GetInspectedByOk returns a tuple with the InspectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedBy

`func (o *BaseDvirLog) SetInspectedBy(v string)`

SetInspectedBy sets InspectedBy field to given value.

### HasInspectedBy

`func (o *BaseDvirLog) HasInspectedBy() bool`

HasInspectedBy returns a boolean if a field has been set.

### SetInspectedByNil

`func (o *BaseDvirLog) SetInspectedByNil(b bool)`

 SetInspectedByNil sets the value for InspectedBy to be an explicit nil

### UnsetInspectedBy
`func (o *BaseDvirLog) UnsetInspectedBy()`

UnsetInspectedBy ensures that no value is present for InspectedBy, not even an explicit nil
### GetCertifyComment

`func (o *BaseDvirLog) GetCertifyComment() string`

GetCertifyComment returns the CertifyComment field if non-nil, zero value otherwise.

### GetCertifyCommentOk

`func (o *BaseDvirLog) GetCertifyCommentOk() (*string, bool)`

GetCertifyCommentOk returns a tuple with the CertifyComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifyComment

`func (o *BaseDvirLog) SetCertifyComment(v string)`

SetCertifyComment sets CertifyComment field to given value.

### HasCertifyComment

`func (o *BaseDvirLog) HasCertifyComment() bool`

HasCertifyComment returns a boolean if a field has been set.

### SetCertifyCommentNil

`func (o *BaseDvirLog) SetCertifyCommentNil(b bool)`

 SetCertifyCommentNil sets the value for CertifyComment to be an explicit nil

### UnsetCertifyComment
`func (o *BaseDvirLog) UnsetCertifyComment()`

UnsetCertifyComment ensures that no value is present for CertifyComment, not even an explicit nil
### GetDriverComment

`func (o *BaseDvirLog) GetDriverComment() string`

GetDriverComment returns the DriverComment field if non-nil, zero value otherwise.

### GetDriverCommentOk

`func (o *BaseDvirLog) GetDriverCommentOk() (*string, bool)`

GetDriverCommentOk returns a tuple with the DriverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverComment

`func (o *BaseDvirLog) SetDriverComment(v string)`

SetDriverComment sets DriverComment field to given value.

### HasDriverComment

`func (o *BaseDvirLog) HasDriverComment() bool`

HasDriverComment returns a boolean if a field has been set.

### SetDriverCommentNil

`func (o *BaseDvirLog) SetDriverCommentNil(b bool)`

 SetDriverCommentNil sets the value for DriverComment to be an explicit nil

### UnsetDriverComment
`func (o *BaseDvirLog) UnsetDriverComment()`

UnsetDriverComment ensures that no value is present for DriverComment, not even an explicit nil
### GetDefectCount

`func (o *BaseDvirLog) GetDefectCount() int32`

GetDefectCount returns the DefectCount field if non-nil, zero value otherwise.

### GetDefectCountOk

`func (o *BaseDvirLog) GetDefectCountOk() (*int32, bool)`

GetDefectCountOk returns a tuple with the DefectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectCount

`func (o *BaseDvirLog) SetDefectCount(v int32)`

SetDefectCount sets DefectCount field to given value.

### HasDefectCount

`func (o *BaseDvirLog) HasDefectCount() bool`

HasDefectCount returns a boolean if a field has been set.

### SetDefectCountNil

`func (o *BaseDvirLog) SetDefectCountNil(b bool)`

 SetDefectCountNil sets the value for DefectCount to be an explicit nil

### UnsetDefectCount
`func (o *BaseDvirLog) UnsetDefectCount()`

UnsetDefectCount ensures that no value is present for DefectCount, not even an explicit nil
### GetVersion

`func (o *BaseDvirLog) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BaseDvirLog) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BaseDvirLog) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BaseDvirLog) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *BaseDvirLog) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *BaseDvirLog) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetDefects

`func (o *BaseDvirLog) GetDefects() []DvirDefectEnum`

GetDefects returns the Defects field if non-nil, zero value otherwise.

### GetDefectsOk

`func (o *BaseDvirLog) GetDefectsOk() (*[]DvirDefectEnum, bool)`

GetDefectsOk returns a tuple with the Defects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefects

`func (o *BaseDvirLog) SetDefects(v []DvirDefectEnum)`

SetDefects sets Defects field to given value.

### HasDefects

`func (o *BaseDvirLog) HasDefects() bool`

HasDefects returns a boolean if a field has been set.

### SetDefectsNil

`func (o *BaseDvirLog) SetDefectsNil(b bool)`

 SetDefectsNil sets the value for Defects to be an explicit nil

### UnsetDefects
`func (o *BaseDvirLog) UnsetDefects()`

UnsetDefects ensures that no value is present for Defects, not even an explicit nil
### GetIsSafetyCritical

`func (o *BaseDvirLog) GetIsSafetyCritical() bool`

GetIsSafetyCritical returns the IsSafetyCritical field if non-nil, zero value otherwise.

### GetIsSafetyCriticalOk

`func (o *BaseDvirLog) GetIsSafetyCriticalOk() (*bool, bool)`

GetIsSafetyCriticalOk returns a tuple with the IsSafetyCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafetyCritical

`func (o *BaseDvirLog) SetIsSafetyCritical(v bool)`

SetIsSafetyCritical sets IsSafetyCritical field to given value.

### HasIsSafetyCritical

`func (o *BaseDvirLog) HasIsSafetyCritical() bool`

HasIsSafetyCritical returns a boolean if a field has been set.

### SetIsSafetyCriticalNil

`func (o *BaseDvirLog) SetIsSafetyCriticalNil(b bool)`

 SetIsSafetyCriticalNil sets the value for IsSafetyCritical to be an explicit nil

### UnsetIsSafetyCritical
`func (o *BaseDvirLog) UnsetIsSafetyCritical()`

UnsetIsSafetyCritical ensures that no value is present for IsSafetyCritical, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


