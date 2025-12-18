# VehicleLocation

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
**OccurredAt** | Pointer to **time.Time** | Timestamp when this location sample was recorded by the TSP (UTC). | [optional] 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 
**VehicleId** | Pointer to **NullableString** |  | [optional] 
**DriverId** | Pointer to **NullableString** |  | [optional] 
**CoDriverId** | Pointer to **NullableString** |  | [optional] 
**SourceDriverId** | Pointer to **NullableString** |  | [optional] 
**SourceVehicleId** | Pointer to **NullableString** |  | [optional] 
**SourceCoDriverId** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to [**NullableLocation3**](Location3.md) |  | [optional] 
**H3Index11** | Pointer to **NullableInt32** |  | [optional] 
**Speed** | Pointer to **NullableInt32** |  | [optional] 
**Odometer** | Pointer to **NullableFloat32** |  | [optional] 
**FuelLevel** | Pointer to **NullableFloat32** |  | [optional] 
**EngineHours** | Pointer to **NullableFloat32** |  | [optional] 
**OilPressure** | Pointer to **NullableFloat32** |  | [optional] 
**CoolantTemperature** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewVehicleLocation

`func NewVehicleLocation(id string, createdAt time.Time, updatedAt time.Time, fleetId string, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *VehicleLocation`

NewVehicleLocation instantiates a new VehicleLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleLocationWithDefaults

`func NewVehicleLocationWithDefaults() *VehicleLocation`

NewVehicleLocationWithDefaults instantiates a new VehicleLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VehicleLocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VehicleLocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VehicleLocation) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *VehicleLocation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VehicleLocation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VehicleLocation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VehicleLocation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VehicleLocation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VehicleLocation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *VehicleLocation) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *VehicleLocation) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *VehicleLocation) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *VehicleLocation) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *VehicleLocation) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *VehicleLocation) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetFleetId

`func (o *VehicleLocation) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *VehicleLocation) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *VehicleLocation) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetConnectionId

`func (o *VehicleLocation) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *VehicleLocation) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *VehicleLocation) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceName

`func (o *VehicleLocation) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *VehicleLocation) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *VehicleLocation) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *VehicleLocation) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *VehicleLocation) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *VehicleLocation) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *VehicleLocation) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *VehicleLocation) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *VehicleLocation) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *VehicleLocation) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *VehicleLocation) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *VehicleLocation) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *VehicleLocation) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *VehicleLocation) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *VehicleLocation) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *VehicleLocation) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *VehicleLocation) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### GetExecutionId

`func (o *VehicleLocation) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *VehicleLocation) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *VehicleLocation) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *VehicleLocation) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *VehicleLocation) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *VehicleLocation) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *VehicleLocation) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *VehicleLocation) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *VehicleLocation) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *VehicleLocation) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *VehicleLocation) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *VehicleLocation) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetVehicleId

`func (o *VehicleLocation) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *VehicleLocation) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *VehicleLocation) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *VehicleLocation) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *VehicleLocation) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *VehicleLocation) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetDriverId

`func (o *VehicleLocation) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *VehicleLocation) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *VehicleLocation) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *VehicleLocation) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *VehicleLocation) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *VehicleLocation) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetCoDriverId

`func (o *VehicleLocation) GetCoDriverId() string`

GetCoDriverId returns the CoDriverId field if non-nil, zero value otherwise.

### GetCoDriverIdOk

`func (o *VehicleLocation) GetCoDriverIdOk() (*string, bool)`

GetCoDriverIdOk returns a tuple with the CoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDriverId

`func (o *VehicleLocation) SetCoDriverId(v string)`

SetCoDriverId sets CoDriverId field to given value.

### HasCoDriverId

`func (o *VehicleLocation) HasCoDriverId() bool`

HasCoDriverId returns a boolean if a field has been set.

### SetCoDriverIdNil

`func (o *VehicleLocation) SetCoDriverIdNil(b bool)`

 SetCoDriverIdNil sets the value for CoDriverId to be an explicit nil

### UnsetCoDriverId
`func (o *VehicleLocation) UnsetCoDriverId()`

UnsetCoDriverId ensures that no value is present for CoDriverId, not even an explicit nil
### GetSourceDriverId

`func (o *VehicleLocation) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *VehicleLocation) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *VehicleLocation) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *VehicleLocation) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *VehicleLocation) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *VehicleLocation) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSourceVehicleId

`func (o *VehicleLocation) GetSourceVehicleId() string`

GetSourceVehicleId returns the SourceVehicleId field if non-nil, zero value otherwise.

### GetSourceVehicleIdOk

`func (o *VehicleLocation) GetSourceVehicleIdOk() (*string, bool)`

GetSourceVehicleIdOk returns a tuple with the SourceVehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVehicleId

`func (o *VehicleLocation) SetSourceVehicleId(v string)`

SetSourceVehicleId sets SourceVehicleId field to given value.

### HasSourceVehicleId

`func (o *VehicleLocation) HasSourceVehicleId() bool`

HasSourceVehicleId returns a boolean if a field has been set.

### SetSourceVehicleIdNil

`func (o *VehicleLocation) SetSourceVehicleIdNil(b bool)`

 SetSourceVehicleIdNil sets the value for SourceVehicleId to be an explicit nil

### UnsetSourceVehicleId
`func (o *VehicleLocation) UnsetSourceVehicleId()`

UnsetSourceVehicleId ensures that no value is present for SourceVehicleId, not even an explicit nil
### GetSourceCoDriverId

`func (o *VehicleLocation) GetSourceCoDriverId() string`

GetSourceCoDriverId returns the SourceCoDriverId field if non-nil, zero value otherwise.

### GetSourceCoDriverIdOk

`func (o *VehicleLocation) GetSourceCoDriverIdOk() (*string, bool)`

GetSourceCoDriverIdOk returns a tuple with the SourceCoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCoDriverId

`func (o *VehicleLocation) SetSourceCoDriverId(v string)`

SetSourceCoDriverId sets SourceCoDriverId field to given value.

### HasSourceCoDriverId

`func (o *VehicleLocation) HasSourceCoDriverId() bool`

HasSourceCoDriverId returns a boolean if a field has been set.

### SetSourceCoDriverIdNil

`func (o *VehicleLocation) SetSourceCoDriverIdNil(b bool)`

 SetSourceCoDriverIdNil sets the value for SourceCoDriverId to be an explicit nil

### UnsetSourceCoDriverId
`func (o *VehicleLocation) UnsetSourceCoDriverId()`

UnsetSourceCoDriverId ensures that no value is present for SourceCoDriverId, not even an explicit nil
### GetLocation

`func (o *VehicleLocation) GetLocation() Location3`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VehicleLocation) GetLocationOk() (*Location3, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VehicleLocation) SetLocation(v Location3)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *VehicleLocation) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *VehicleLocation) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *VehicleLocation) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetH3Index11

`func (o *VehicleLocation) GetH3Index11() int32`

GetH3Index11 returns the H3Index11 field if non-nil, zero value otherwise.

### GetH3Index11Ok

`func (o *VehicleLocation) GetH3Index11Ok() (*int32, bool)`

GetH3Index11Ok returns a tuple with the H3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH3Index11

`func (o *VehicleLocation) SetH3Index11(v int32)`

SetH3Index11 sets H3Index11 field to given value.

### HasH3Index11

`func (o *VehicleLocation) HasH3Index11() bool`

HasH3Index11 returns a boolean if a field has been set.

### SetH3Index11Nil

`func (o *VehicleLocation) SetH3Index11Nil(b bool)`

 SetH3Index11Nil sets the value for H3Index11 to be an explicit nil

### UnsetH3Index11
`func (o *VehicleLocation) UnsetH3Index11()`

UnsetH3Index11 ensures that no value is present for H3Index11, not even an explicit nil
### GetSpeed

`func (o *VehicleLocation) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VehicleLocation) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VehicleLocation) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *VehicleLocation) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *VehicleLocation) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *VehicleLocation) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetOdometer

`func (o *VehicleLocation) GetOdometer() float32`

GetOdometer returns the Odometer field if non-nil, zero value otherwise.

### GetOdometerOk

`func (o *VehicleLocation) GetOdometerOk() (*float32, bool)`

GetOdometerOk returns a tuple with the Odometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometer

`func (o *VehicleLocation) SetOdometer(v float32)`

SetOdometer sets Odometer field to given value.

### HasOdometer

`func (o *VehicleLocation) HasOdometer() bool`

HasOdometer returns a boolean if a field has been set.

### SetOdometerNil

`func (o *VehicleLocation) SetOdometerNil(b bool)`

 SetOdometerNil sets the value for Odometer to be an explicit nil

### UnsetOdometer
`func (o *VehicleLocation) UnsetOdometer()`

UnsetOdometer ensures that no value is present for Odometer, not even an explicit nil
### GetFuelLevel

`func (o *VehicleLocation) GetFuelLevel() float32`

GetFuelLevel returns the FuelLevel field if non-nil, zero value otherwise.

### GetFuelLevelOk

`func (o *VehicleLocation) GetFuelLevelOk() (*float32, bool)`

GetFuelLevelOk returns a tuple with the FuelLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelLevel

`func (o *VehicleLocation) SetFuelLevel(v float32)`

SetFuelLevel sets FuelLevel field to given value.

### HasFuelLevel

`func (o *VehicleLocation) HasFuelLevel() bool`

HasFuelLevel returns a boolean if a field has been set.

### SetFuelLevelNil

`func (o *VehicleLocation) SetFuelLevelNil(b bool)`

 SetFuelLevelNil sets the value for FuelLevel to be an explicit nil

### UnsetFuelLevel
`func (o *VehicleLocation) UnsetFuelLevel()`

UnsetFuelLevel ensures that no value is present for FuelLevel, not even an explicit nil
### GetEngineHours

`func (o *VehicleLocation) GetEngineHours() float32`

GetEngineHours returns the EngineHours field if non-nil, zero value otherwise.

### GetEngineHoursOk

`func (o *VehicleLocation) GetEngineHoursOk() (*float32, bool)`

GetEngineHoursOk returns a tuple with the EngineHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineHours

`func (o *VehicleLocation) SetEngineHours(v float32)`

SetEngineHours sets EngineHours field to given value.

### HasEngineHours

`func (o *VehicleLocation) HasEngineHours() bool`

HasEngineHours returns a boolean if a field has been set.

### SetEngineHoursNil

`func (o *VehicleLocation) SetEngineHoursNil(b bool)`

 SetEngineHoursNil sets the value for EngineHours to be an explicit nil

### UnsetEngineHours
`func (o *VehicleLocation) UnsetEngineHours()`

UnsetEngineHours ensures that no value is present for EngineHours, not even an explicit nil
### GetOilPressure

`func (o *VehicleLocation) GetOilPressure() float32`

GetOilPressure returns the OilPressure field if non-nil, zero value otherwise.

### GetOilPressureOk

`func (o *VehicleLocation) GetOilPressureOk() (*float32, bool)`

GetOilPressureOk returns a tuple with the OilPressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOilPressure

`func (o *VehicleLocation) SetOilPressure(v float32)`

SetOilPressure sets OilPressure field to given value.

### HasOilPressure

`func (o *VehicleLocation) HasOilPressure() bool`

HasOilPressure returns a boolean if a field has been set.

### SetOilPressureNil

`func (o *VehicleLocation) SetOilPressureNil(b bool)`

 SetOilPressureNil sets the value for OilPressure to be an explicit nil

### UnsetOilPressure
`func (o *VehicleLocation) UnsetOilPressure()`

UnsetOilPressure ensures that no value is present for OilPressure, not even an explicit nil
### GetCoolantTemperature

`func (o *VehicleLocation) GetCoolantTemperature() float32`

GetCoolantTemperature returns the CoolantTemperature field if non-nil, zero value otherwise.

### GetCoolantTemperatureOk

`func (o *VehicleLocation) GetCoolantTemperatureOk() (*float32, bool)`

GetCoolantTemperatureOk returns a tuple with the CoolantTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolantTemperature

`func (o *VehicleLocation) SetCoolantTemperature(v float32)`

SetCoolantTemperature sets CoolantTemperature field to given value.

### HasCoolantTemperature

`func (o *VehicleLocation) HasCoolantTemperature() bool`

HasCoolantTemperature returns a boolean if a field has been set.

### SetCoolantTemperatureNil

`func (o *VehicleLocation) SetCoolantTemperatureNil(b bool)`

 SetCoolantTemperatureNil sets the value for CoolantTemperature to be an explicit nil

### UnsetCoolantTemperature
`func (o *VehicleLocation) UnsetCoolantTemperature()`

UnsetCoolantTemperature ensures that no value is present for CoolantTemperature, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


