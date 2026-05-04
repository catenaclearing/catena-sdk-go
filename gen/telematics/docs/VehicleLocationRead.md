# VehicleLocationRead

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
**CoDriverId** | Pointer to **NullableString** |  | [optional] 
**SourceDriverId** | Pointer to **NullableString** |  | [optional] 
**SourceVehicleId** | Pointer to **NullableString** |  | [optional] 
**SourceCoDriverId** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to [**NullableLocation5**](Location5.md) |  | [optional] 
**H3Index11** | Pointer to **NullableInt32** |  | [optional] 
**Speed** | Pointer to **NullableFloat32** |  | [optional] 
**Odometer** | Pointer to **NullableFloat32** |  | [optional] 
**FuelLevel** | Pointer to **NullableFloat32** |  | [optional] 
**FuelValue** | Pointer to **NullableFloat32** |  | [optional] 
**EngineHours** | Pointer to **NullableFloat32** |  | [optional] 
**OilPressure** | Pointer to **NullableFloat32** |  | [optional] 
**CoolantTemperature** | Pointer to **NullableFloat32** |  | [optional] 
**InferredAddress** | Pointer to [**NullableInferredAddress**](InferredAddress.md) |  | [optional] 
**SpeedUnit** | Pointer to [**SpeedUnit**](SpeedUnit.md) | Unit for speed. | [optional] 
**OdometerUnit** | Pointer to [**DistanceUnit**](DistanceUnit.md) | Unit for odometer. | [optional] 
**FuelValueUnit** | Pointer to [**VolumeUnit**](VolumeUnit.md) | Unit for fuel_value. | [optional] 
**OilPressureUnit** | Pointer to [**PressureUnit**](PressureUnit.md) | Unit for oil_pressure. | [optional] 
**CoolantTemperatureUnit** | Pointer to [**TemperatureUnit**](TemperatureUnit.md) | Unit for coolant_temperature. | [optional] 

## Methods

### NewVehicleLocationRead

`func NewVehicleLocationRead(fleetId NullableString, id string, createdAt time.Time, updatedAt time.Time, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *VehicleLocationRead`

NewVehicleLocationRead instantiates a new VehicleLocationRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleLocationReadWithDefaults

`func NewVehicleLocationReadWithDefaults() *VehicleLocationRead`

NewVehicleLocationReadWithDefaults instantiates a new VehicleLocationRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *VehicleLocationRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *VehicleLocationRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *VehicleLocationRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### SetFleetIdNil

`func (o *VehicleLocationRead) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *VehicleLocationRead) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetFleetRef

`func (o *VehicleLocationRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *VehicleLocationRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *VehicleLocationRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *VehicleLocationRead) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *VehicleLocationRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *VehicleLocationRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetId

`func (o *VehicleLocationRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VehicleLocationRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VehicleLocationRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *VehicleLocationRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VehicleLocationRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VehicleLocationRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VehicleLocationRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VehicleLocationRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VehicleLocationRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *VehicleLocationRead) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *VehicleLocationRead) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *VehicleLocationRead) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *VehicleLocationRead) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *VehicleLocationRead) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *VehicleLocationRead) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetConnectionId

`func (o *VehicleLocationRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *VehicleLocationRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *VehicleLocationRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetTspId

`func (o *VehicleLocationRead) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *VehicleLocationRead) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *VehicleLocationRead) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *VehicleLocationRead) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *VehicleLocationRead) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *VehicleLocationRead) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *VehicleLocationRead) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *VehicleLocationRead) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *VehicleLocationRead) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *VehicleLocationRead) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *VehicleLocationRead) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *VehicleLocationRead) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *VehicleLocationRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *VehicleLocationRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *VehicleLocationRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *VehicleLocationRead) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *VehicleLocationRead) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *VehicleLocationRead) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *VehicleLocationRead) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *VehicleLocationRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *VehicleLocationRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *VehicleLocationRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *VehicleLocationRead) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *VehicleLocationRead) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *VehicleLocationRead) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *VehicleLocationRead) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *VehicleLocationRead) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *VehicleLocationRead) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *VehicleLocationRead) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *VehicleLocationRead) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *VehicleLocationRead) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *VehicleLocationRead) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *VehicleLocationRead) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *VehicleLocationRead) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *VehicleLocationRead) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *VehicleLocationRead) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *VehicleLocationRead) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *VehicleLocationRead) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *VehicleLocationRead) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *VehicleLocationRead) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *VehicleLocationRead) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *VehicleLocationRead) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *VehicleLocationRead) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *VehicleLocationRead) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *VehicleLocationRead) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *VehicleLocationRead) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *VehicleLocationRead) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *VehicleLocationRead) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *VehicleLocationRead) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetVehicleId

`func (o *VehicleLocationRead) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *VehicleLocationRead) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *VehicleLocationRead) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *VehicleLocationRead) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *VehicleLocationRead) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *VehicleLocationRead) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetDriverId

`func (o *VehicleLocationRead) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *VehicleLocationRead) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *VehicleLocationRead) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *VehicleLocationRead) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *VehicleLocationRead) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *VehicleLocationRead) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetCoDriverId

`func (o *VehicleLocationRead) GetCoDriverId() string`

GetCoDriverId returns the CoDriverId field if non-nil, zero value otherwise.

### GetCoDriverIdOk

`func (o *VehicleLocationRead) GetCoDriverIdOk() (*string, bool)`

GetCoDriverIdOk returns a tuple with the CoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDriverId

`func (o *VehicleLocationRead) SetCoDriverId(v string)`

SetCoDriverId sets CoDriverId field to given value.

### HasCoDriverId

`func (o *VehicleLocationRead) HasCoDriverId() bool`

HasCoDriverId returns a boolean if a field has been set.

### SetCoDriverIdNil

`func (o *VehicleLocationRead) SetCoDriverIdNil(b bool)`

 SetCoDriverIdNil sets the value for CoDriverId to be an explicit nil

### UnsetCoDriverId
`func (o *VehicleLocationRead) UnsetCoDriverId()`

UnsetCoDriverId ensures that no value is present for CoDriverId, not even an explicit nil
### GetSourceDriverId

`func (o *VehicleLocationRead) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *VehicleLocationRead) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *VehicleLocationRead) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *VehicleLocationRead) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *VehicleLocationRead) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *VehicleLocationRead) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSourceVehicleId

`func (o *VehicleLocationRead) GetSourceVehicleId() string`

GetSourceVehicleId returns the SourceVehicleId field if non-nil, zero value otherwise.

### GetSourceVehicleIdOk

`func (o *VehicleLocationRead) GetSourceVehicleIdOk() (*string, bool)`

GetSourceVehicleIdOk returns a tuple with the SourceVehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVehicleId

`func (o *VehicleLocationRead) SetSourceVehicleId(v string)`

SetSourceVehicleId sets SourceVehicleId field to given value.

### HasSourceVehicleId

`func (o *VehicleLocationRead) HasSourceVehicleId() bool`

HasSourceVehicleId returns a boolean if a field has been set.

### SetSourceVehicleIdNil

`func (o *VehicleLocationRead) SetSourceVehicleIdNil(b bool)`

 SetSourceVehicleIdNil sets the value for SourceVehicleId to be an explicit nil

### UnsetSourceVehicleId
`func (o *VehicleLocationRead) UnsetSourceVehicleId()`

UnsetSourceVehicleId ensures that no value is present for SourceVehicleId, not even an explicit nil
### GetSourceCoDriverId

`func (o *VehicleLocationRead) GetSourceCoDriverId() string`

GetSourceCoDriverId returns the SourceCoDriverId field if non-nil, zero value otherwise.

### GetSourceCoDriverIdOk

`func (o *VehicleLocationRead) GetSourceCoDriverIdOk() (*string, bool)`

GetSourceCoDriverIdOk returns a tuple with the SourceCoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCoDriverId

`func (o *VehicleLocationRead) SetSourceCoDriverId(v string)`

SetSourceCoDriverId sets SourceCoDriverId field to given value.

### HasSourceCoDriverId

`func (o *VehicleLocationRead) HasSourceCoDriverId() bool`

HasSourceCoDriverId returns a boolean if a field has been set.

### SetSourceCoDriverIdNil

`func (o *VehicleLocationRead) SetSourceCoDriverIdNil(b bool)`

 SetSourceCoDriverIdNil sets the value for SourceCoDriverId to be an explicit nil

### UnsetSourceCoDriverId
`func (o *VehicleLocationRead) UnsetSourceCoDriverId()`

UnsetSourceCoDriverId ensures that no value is present for SourceCoDriverId, not even an explicit nil
### GetLocation

`func (o *VehicleLocationRead) GetLocation() Location5`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VehicleLocationRead) GetLocationOk() (*Location5, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VehicleLocationRead) SetLocation(v Location5)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *VehicleLocationRead) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *VehicleLocationRead) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *VehicleLocationRead) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetH3Index11

`func (o *VehicleLocationRead) GetH3Index11() int32`

GetH3Index11 returns the H3Index11 field if non-nil, zero value otherwise.

### GetH3Index11Ok

`func (o *VehicleLocationRead) GetH3Index11Ok() (*int32, bool)`

GetH3Index11Ok returns a tuple with the H3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH3Index11

`func (o *VehicleLocationRead) SetH3Index11(v int32)`

SetH3Index11 sets H3Index11 field to given value.

### HasH3Index11

`func (o *VehicleLocationRead) HasH3Index11() bool`

HasH3Index11 returns a boolean if a field has been set.

### SetH3Index11Nil

`func (o *VehicleLocationRead) SetH3Index11Nil(b bool)`

 SetH3Index11Nil sets the value for H3Index11 to be an explicit nil

### UnsetH3Index11
`func (o *VehicleLocationRead) UnsetH3Index11()`

UnsetH3Index11 ensures that no value is present for H3Index11, not even an explicit nil
### GetSpeed

`func (o *VehicleLocationRead) GetSpeed() float32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VehicleLocationRead) GetSpeedOk() (*float32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VehicleLocationRead) SetSpeed(v float32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VehicleLocationRead) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *VehicleLocationRead) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *VehicleLocationRead) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetOdometer

`func (o *VehicleLocationRead) GetOdometer() float32`

GetOdometer returns the Odometer field if non-nil, zero value otherwise.

### GetOdometerOk

`func (o *VehicleLocationRead) GetOdometerOk() (*float32, bool)`

GetOdometerOk returns a tuple with the Odometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometer

`func (o *VehicleLocationRead) SetOdometer(v float32)`

SetOdometer sets Odometer field to given value.

### HasOdometer

`func (o *VehicleLocationRead) HasOdometer() bool`

HasOdometer returns a boolean if a field has been set.

### SetOdometerNil

`func (o *VehicleLocationRead) SetOdometerNil(b bool)`

 SetOdometerNil sets the value for Odometer to be an explicit nil

### UnsetOdometer
`func (o *VehicleLocationRead) UnsetOdometer()`

UnsetOdometer ensures that no value is present for Odometer, not even an explicit nil
### GetFuelLevel

`func (o *VehicleLocationRead) GetFuelLevel() float32`

GetFuelLevel returns the FuelLevel field if non-nil, zero value otherwise.

### GetFuelLevelOk

`func (o *VehicleLocationRead) GetFuelLevelOk() (*float32, bool)`

GetFuelLevelOk returns a tuple with the FuelLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelLevel

`func (o *VehicleLocationRead) SetFuelLevel(v float32)`

SetFuelLevel sets FuelLevel field to given value.

### HasFuelLevel

`func (o *VehicleLocationRead) HasFuelLevel() bool`

HasFuelLevel returns a boolean if a field has been set.

### SetFuelLevelNil

`func (o *VehicleLocationRead) SetFuelLevelNil(b bool)`

 SetFuelLevelNil sets the value for FuelLevel to be an explicit nil

### UnsetFuelLevel
`func (o *VehicleLocationRead) UnsetFuelLevel()`

UnsetFuelLevel ensures that no value is present for FuelLevel, not even an explicit nil
### GetFuelValue

`func (o *VehicleLocationRead) GetFuelValue() float32`

GetFuelValue returns the FuelValue field if non-nil, zero value otherwise.

### GetFuelValueOk

`func (o *VehicleLocationRead) GetFuelValueOk() (*float32, bool)`

GetFuelValueOk returns a tuple with the FuelValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelValue

`func (o *VehicleLocationRead) SetFuelValue(v float32)`

SetFuelValue sets FuelValue field to given value.

### HasFuelValue

`func (o *VehicleLocationRead) HasFuelValue() bool`

HasFuelValue returns a boolean if a field has been set.

### SetFuelValueNil

`func (o *VehicleLocationRead) SetFuelValueNil(b bool)`

 SetFuelValueNil sets the value for FuelValue to be an explicit nil

### UnsetFuelValue
`func (o *VehicleLocationRead) UnsetFuelValue()`

UnsetFuelValue ensures that no value is present for FuelValue, not even an explicit nil
### GetEngineHours

`func (o *VehicleLocationRead) GetEngineHours() float32`

GetEngineHours returns the EngineHours field if non-nil, zero value otherwise.

### GetEngineHoursOk

`func (o *VehicleLocationRead) GetEngineHoursOk() (*float32, bool)`

GetEngineHoursOk returns a tuple with the EngineHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineHours

`func (o *VehicleLocationRead) SetEngineHours(v float32)`

SetEngineHours sets EngineHours field to given value.

### HasEngineHours

`func (o *VehicleLocationRead) HasEngineHours() bool`

HasEngineHours returns a boolean if a field has been set.

### SetEngineHoursNil

`func (o *VehicleLocationRead) SetEngineHoursNil(b bool)`

 SetEngineHoursNil sets the value for EngineHours to be an explicit nil

### UnsetEngineHours
`func (o *VehicleLocationRead) UnsetEngineHours()`

UnsetEngineHours ensures that no value is present for EngineHours, not even an explicit nil
### GetOilPressure

`func (o *VehicleLocationRead) GetOilPressure() float32`

GetOilPressure returns the OilPressure field if non-nil, zero value otherwise.

### GetOilPressureOk

`func (o *VehicleLocationRead) GetOilPressureOk() (*float32, bool)`

GetOilPressureOk returns a tuple with the OilPressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOilPressure

`func (o *VehicleLocationRead) SetOilPressure(v float32)`

SetOilPressure sets OilPressure field to given value.

### HasOilPressure

`func (o *VehicleLocationRead) HasOilPressure() bool`

HasOilPressure returns a boolean if a field has been set.

### SetOilPressureNil

`func (o *VehicleLocationRead) SetOilPressureNil(b bool)`

 SetOilPressureNil sets the value for OilPressure to be an explicit nil

### UnsetOilPressure
`func (o *VehicleLocationRead) UnsetOilPressure()`

UnsetOilPressure ensures that no value is present for OilPressure, not even an explicit nil
### GetCoolantTemperature

`func (o *VehicleLocationRead) GetCoolantTemperature() float32`

GetCoolantTemperature returns the CoolantTemperature field if non-nil, zero value otherwise.

### GetCoolantTemperatureOk

`func (o *VehicleLocationRead) GetCoolantTemperatureOk() (*float32, bool)`

GetCoolantTemperatureOk returns a tuple with the CoolantTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolantTemperature

`func (o *VehicleLocationRead) SetCoolantTemperature(v float32)`

SetCoolantTemperature sets CoolantTemperature field to given value.

### HasCoolantTemperature

`func (o *VehicleLocationRead) HasCoolantTemperature() bool`

HasCoolantTemperature returns a boolean if a field has been set.

### SetCoolantTemperatureNil

`func (o *VehicleLocationRead) SetCoolantTemperatureNil(b bool)`

 SetCoolantTemperatureNil sets the value for CoolantTemperature to be an explicit nil

### UnsetCoolantTemperature
`func (o *VehicleLocationRead) UnsetCoolantTemperature()`

UnsetCoolantTemperature ensures that no value is present for CoolantTemperature, not even an explicit nil
### GetInferredAddress

`func (o *VehicleLocationRead) GetInferredAddress() InferredAddress`

GetInferredAddress returns the InferredAddress field if non-nil, zero value otherwise.

### GetInferredAddressOk

`func (o *VehicleLocationRead) GetInferredAddressOk() (*InferredAddress, bool)`

GetInferredAddressOk returns a tuple with the InferredAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferredAddress

`func (o *VehicleLocationRead) SetInferredAddress(v InferredAddress)`

SetInferredAddress sets InferredAddress field to given value.

### HasInferredAddress

`func (o *VehicleLocationRead) HasInferredAddress() bool`

HasInferredAddress returns a boolean if a field has been set.

### SetInferredAddressNil

`func (o *VehicleLocationRead) SetInferredAddressNil(b bool)`

 SetInferredAddressNil sets the value for InferredAddress to be an explicit nil

### UnsetInferredAddress
`func (o *VehicleLocationRead) UnsetInferredAddress()`

UnsetInferredAddress ensures that no value is present for InferredAddress, not even an explicit nil
### GetSpeedUnit

`func (o *VehicleLocationRead) GetSpeedUnit() SpeedUnit`

GetSpeedUnit returns the SpeedUnit field if non-nil, zero value otherwise.

### GetSpeedUnitOk

`func (o *VehicleLocationRead) GetSpeedUnitOk() (*SpeedUnit, bool)`

GetSpeedUnitOk returns a tuple with the SpeedUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedUnit

`func (o *VehicleLocationRead) SetSpeedUnit(v SpeedUnit)`

SetSpeedUnit sets SpeedUnit field to given value.

### HasSpeedUnit

`func (o *VehicleLocationRead) HasSpeedUnit() bool`

HasSpeedUnit returns a boolean if a field has been set.

### GetOdometerUnit

`func (o *VehicleLocationRead) GetOdometerUnit() DistanceUnit`

GetOdometerUnit returns the OdometerUnit field if non-nil, zero value otherwise.

### GetOdometerUnitOk

`func (o *VehicleLocationRead) GetOdometerUnitOk() (*DistanceUnit, bool)`

GetOdometerUnitOk returns a tuple with the OdometerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerUnit

`func (o *VehicleLocationRead) SetOdometerUnit(v DistanceUnit)`

SetOdometerUnit sets OdometerUnit field to given value.

### HasOdometerUnit

`func (o *VehicleLocationRead) HasOdometerUnit() bool`

HasOdometerUnit returns a boolean if a field has been set.

### GetFuelValueUnit

`func (o *VehicleLocationRead) GetFuelValueUnit() VolumeUnit`

GetFuelValueUnit returns the FuelValueUnit field if non-nil, zero value otherwise.

### GetFuelValueUnitOk

`func (o *VehicleLocationRead) GetFuelValueUnitOk() (*VolumeUnit, bool)`

GetFuelValueUnitOk returns a tuple with the FuelValueUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelValueUnit

`func (o *VehicleLocationRead) SetFuelValueUnit(v VolumeUnit)`

SetFuelValueUnit sets FuelValueUnit field to given value.

### HasFuelValueUnit

`func (o *VehicleLocationRead) HasFuelValueUnit() bool`

HasFuelValueUnit returns a boolean if a field has been set.

### GetOilPressureUnit

`func (o *VehicleLocationRead) GetOilPressureUnit() PressureUnit`

GetOilPressureUnit returns the OilPressureUnit field if non-nil, zero value otherwise.

### GetOilPressureUnitOk

`func (o *VehicleLocationRead) GetOilPressureUnitOk() (*PressureUnit, bool)`

GetOilPressureUnitOk returns a tuple with the OilPressureUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOilPressureUnit

`func (o *VehicleLocationRead) SetOilPressureUnit(v PressureUnit)`

SetOilPressureUnit sets OilPressureUnit field to given value.

### HasOilPressureUnit

`func (o *VehicleLocationRead) HasOilPressureUnit() bool`

HasOilPressureUnit returns a boolean if a field has been set.

### GetCoolantTemperatureUnit

`func (o *VehicleLocationRead) GetCoolantTemperatureUnit() TemperatureUnit`

GetCoolantTemperatureUnit returns the CoolantTemperatureUnit field if non-nil, zero value otherwise.

### GetCoolantTemperatureUnitOk

`func (o *VehicleLocationRead) GetCoolantTemperatureUnitOk() (*TemperatureUnit, bool)`

GetCoolantTemperatureUnitOk returns a tuple with the CoolantTemperatureUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolantTemperatureUnit

`func (o *VehicleLocationRead) SetCoolantTemperatureUnit(v TemperatureUnit)`

SetCoolantTemperatureUnit sets CoolantTemperatureUnit field to given value.

### HasCoolantTemperatureUnit

`func (o *VehicleLocationRead) HasCoolantTemperatureUnit() bool`

HasCoolantTemperatureUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


