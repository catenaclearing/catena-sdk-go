# BaseVehicleLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal unique identifier for the telematics event record (Catena PK). | 
**FleetId** | **string** | The Catena fleet this record belongs to (multi-tenant scope). | 
**SourceName** | [**TspEnum**](TspEnum.md) | The name of the source | 
**ConnectionId** | **string** | The specific fleet↔TSP connection through which this record was sourced. | 
**SourceId** | **string** | The ID of the record in the TSP or a deterministic ID/Hash generated from a composite unique key | 
**CreatedAt** | **time.Time** | Immutable: first time this record was ingested into our system. | 
**UpdatedAt** | **time.Time** | Last time we modified this record in our system. | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**OccurredAt** | **time.Time** | When the underlying event/observation occurred, as reported by the TSP, or the moment it was ingested by us if not available. | 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 
**VehicleId** | Pointer to **NullableString** |  | [optional] 
**DriverId** | Pointer to **NullableString** |  | [optional] 
**CoDriverId** | Pointer to **NullableString** |  | [optional] 
**SourceDriverId** | Pointer to **NullableString** |  | [optional] 
**SourceVehicleId** | Pointer to **NullableString** |  | [optional] 
**SourceCoDriverId** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to [**NullablePoint**](Point.md) |  | [optional] 
**H3Index11** | Pointer to **NullableInt32** |  | [optional] 
**Speed** | Pointer to **NullableInt32** |  | [optional] 
**Odometer** | Pointer to **NullableFloat32** |  | [optional] 
**FuelLevel** | Pointer to **NullableFloat32** |  | [optional] 
**EngineHours** | Pointer to **NullableFloat32** |  | [optional] 
**OilPressure** | Pointer to **NullableFloat32** |  | [optional] 
**CoolantTemperature** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewBaseVehicleLocation

`func NewBaseVehicleLocation(id string, fleetId string, sourceName TspEnum, connectionId string, sourceId string, createdAt time.Time, updatedAt time.Time, occurredAt time.Time, ) *BaseVehicleLocation`

NewBaseVehicleLocation instantiates a new BaseVehicleLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseVehicleLocationWithDefaults

`func NewBaseVehicleLocationWithDefaults() *BaseVehicleLocation`

NewBaseVehicleLocationWithDefaults instantiates a new BaseVehicleLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseVehicleLocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseVehicleLocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseVehicleLocation) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseVehicleLocation) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseVehicleLocation) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseVehicleLocation) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetSourceName

`func (o *BaseVehicleLocation) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseVehicleLocation) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseVehicleLocation) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetConnectionId

`func (o *BaseVehicleLocation) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseVehicleLocation) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseVehicleLocation) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceId

`func (o *BaseVehicleLocation) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseVehicleLocation) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseVehicleLocation) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCreatedAt

`func (o *BaseVehicleLocation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseVehicleLocation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseVehicleLocation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseVehicleLocation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseVehicleLocation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseVehicleLocation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BaseVehicleLocation) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BaseVehicleLocation) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BaseVehicleLocation) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BaseVehicleLocation) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BaseVehicleLocation) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BaseVehicleLocation) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOccurredAt

`func (o *BaseVehicleLocation) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BaseVehicleLocation) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BaseVehicleLocation) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetExecutionId

`func (o *BaseVehicleLocation) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BaseVehicleLocation) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BaseVehicleLocation) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *BaseVehicleLocation) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *BaseVehicleLocation) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *BaseVehicleLocation) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *BaseVehicleLocation) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BaseVehicleLocation) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BaseVehicleLocation) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *BaseVehicleLocation) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *BaseVehicleLocation) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *BaseVehicleLocation) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetVehicleId

`func (o *BaseVehicleLocation) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *BaseVehicleLocation) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *BaseVehicleLocation) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *BaseVehicleLocation) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *BaseVehicleLocation) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *BaseVehicleLocation) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetDriverId

`func (o *BaseVehicleLocation) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *BaseVehicleLocation) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *BaseVehicleLocation) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *BaseVehicleLocation) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *BaseVehicleLocation) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *BaseVehicleLocation) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetCoDriverId

`func (o *BaseVehicleLocation) GetCoDriverId() string`

GetCoDriverId returns the CoDriverId field if non-nil, zero value otherwise.

### GetCoDriverIdOk

`func (o *BaseVehicleLocation) GetCoDriverIdOk() (*string, bool)`

GetCoDriverIdOk returns a tuple with the CoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDriverId

`func (o *BaseVehicleLocation) SetCoDriverId(v string)`

SetCoDriverId sets CoDriverId field to given value.

### HasCoDriverId

`func (o *BaseVehicleLocation) HasCoDriverId() bool`

HasCoDriverId returns a boolean if a field has been set.

### SetCoDriverIdNil

`func (o *BaseVehicleLocation) SetCoDriverIdNil(b bool)`

 SetCoDriverIdNil sets the value for CoDriverId to be an explicit nil

### UnsetCoDriverId
`func (o *BaseVehicleLocation) UnsetCoDriverId()`

UnsetCoDriverId ensures that no value is present for CoDriverId, not even an explicit nil
### GetSourceDriverId

`func (o *BaseVehicleLocation) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *BaseVehicleLocation) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *BaseVehicleLocation) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *BaseVehicleLocation) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *BaseVehicleLocation) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *BaseVehicleLocation) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSourceVehicleId

`func (o *BaseVehicleLocation) GetSourceVehicleId() string`

GetSourceVehicleId returns the SourceVehicleId field if non-nil, zero value otherwise.

### GetSourceVehicleIdOk

`func (o *BaseVehicleLocation) GetSourceVehicleIdOk() (*string, bool)`

GetSourceVehicleIdOk returns a tuple with the SourceVehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVehicleId

`func (o *BaseVehicleLocation) SetSourceVehicleId(v string)`

SetSourceVehicleId sets SourceVehicleId field to given value.

### HasSourceVehicleId

`func (o *BaseVehicleLocation) HasSourceVehicleId() bool`

HasSourceVehicleId returns a boolean if a field has been set.

### SetSourceVehicleIdNil

`func (o *BaseVehicleLocation) SetSourceVehicleIdNil(b bool)`

 SetSourceVehicleIdNil sets the value for SourceVehicleId to be an explicit nil

### UnsetSourceVehicleId
`func (o *BaseVehicleLocation) UnsetSourceVehicleId()`

UnsetSourceVehicleId ensures that no value is present for SourceVehicleId, not even an explicit nil
### GetSourceCoDriverId

`func (o *BaseVehicleLocation) GetSourceCoDriverId() string`

GetSourceCoDriverId returns the SourceCoDriverId field if non-nil, zero value otherwise.

### GetSourceCoDriverIdOk

`func (o *BaseVehicleLocation) GetSourceCoDriverIdOk() (*string, bool)`

GetSourceCoDriverIdOk returns a tuple with the SourceCoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCoDriverId

`func (o *BaseVehicleLocation) SetSourceCoDriverId(v string)`

SetSourceCoDriverId sets SourceCoDriverId field to given value.

### HasSourceCoDriverId

`func (o *BaseVehicleLocation) HasSourceCoDriverId() bool`

HasSourceCoDriverId returns a boolean if a field has been set.

### SetSourceCoDriverIdNil

`func (o *BaseVehicleLocation) SetSourceCoDriverIdNil(b bool)`

 SetSourceCoDriverIdNil sets the value for SourceCoDriverId to be an explicit nil

### UnsetSourceCoDriverId
`func (o *BaseVehicleLocation) UnsetSourceCoDriverId()`

UnsetSourceCoDriverId ensures that no value is present for SourceCoDriverId, not even an explicit nil
### GetLocation

`func (o *BaseVehicleLocation) GetLocation() Point`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *BaseVehicleLocation) GetLocationOk() (*Point, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *BaseVehicleLocation) SetLocation(v Point)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *BaseVehicleLocation) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *BaseVehicleLocation) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *BaseVehicleLocation) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetH3Index11

`func (o *BaseVehicleLocation) GetH3Index11() int32`

GetH3Index11 returns the H3Index11 field if non-nil, zero value otherwise.

### GetH3Index11Ok

`func (o *BaseVehicleLocation) GetH3Index11Ok() (*int32, bool)`

GetH3Index11Ok returns a tuple with the H3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH3Index11

`func (o *BaseVehicleLocation) SetH3Index11(v int32)`

SetH3Index11 sets H3Index11 field to given value.

### HasH3Index11

`func (o *BaseVehicleLocation) HasH3Index11() bool`

HasH3Index11 returns a boolean if a field has been set.

### SetH3Index11Nil

`func (o *BaseVehicleLocation) SetH3Index11Nil(b bool)`

 SetH3Index11Nil sets the value for H3Index11 to be an explicit nil

### UnsetH3Index11
`func (o *BaseVehicleLocation) UnsetH3Index11()`

UnsetH3Index11 ensures that no value is present for H3Index11, not even an explicit nil
### GetSpeed

`func (o *BaseVehicleLocation) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *BaseVehicleLocation) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *BaseVehicleLocation) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *BaseVehicleLocation) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### SetSpeedNil

`func (o *BaseVehicleLocation) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *BaseVehicleLocation) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetOdometer

`func (o *BaseVehicleLocation) GetOdometer() float32`

GetOdometer returns the Odometer field if non-nil, zero value otherwise.

### GetOdometerOk

`func (o *BaseVehicleLocation) GetOdometerOk() (*float32, bool)`

GetOdometerOk returns a tuple with the Odometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometer

`func (o *BaseVehicleLocation) SetOdometer(v float32)`

SetOdometer sets Odometer field to given value.

### HasOdometer

`func (o *BaseVehicleLocation) HasOdometer() bool`

HasOdometer returns a boolean if a field has been set.

### SetOdometerNil

`func (o *BaseVehicleLocation) SetOdometerNil(b bool)`

 SetOdometerNil sets the value for Odometer to be an explicit nil

### UnsetOdometer
`func (o *BaseVehicleLocation) UnsetOdometer()`

UnsetOdometer ensures that no value is present for Odometer, not even an explicit nil
### GetFuelLevel

`func (o *BaseVehicleLocation) GetFuelLevel() float32`

GetFuelLevel returns the FuelLevel field if non-nil, zero value otherwise.

### GetFuelLevelOk

`func (o *BaseVehicleLocation) GetFuelLevelOk() (*float32, bool)`

GetFuelLevelOk returns a tuple with the FuelLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelLevel

`func (o *BaseVehicleLocation) SetFuelLevel(v float32)`

SetFuelLevel sets FuelLevel field to given value.

### HasFuelLevel

`func (o *BaseVehicleLocation) HasFuelLevel() bool`

HasFuelLevel returns a boolean if a field has been set.

### SetFuelLevelNil

`func (o *BaseVehicleLocation) SetFuelLevelNil(b bool)`

 SetFuelLevelNil sets the value for FuelLevel to be an explicit nil

### UnsetFuelLevel
`func (o *BaseVehicleLocation) UnsetFuelLevel()`

UnsetFuelLevel ensures that no value is present for FuelLevel, not even an explicit nil
### GetEngineHours

`func (o *BaseVehicleLocation) GetEngineHours() float32`

GetEngineHours returns the EngineHours field if non-nil, zero value otherwise.

### GetEngineHoursOk

`func (o *BaseVehicleLocation) GetEngineHoursOk() (*float32, bool)`

GetEngineHoursOk returns a tuple with the EngineHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineHours

`func (o *BaseVehicleLocation) SetEngineHours(v float32)`

SetEngineHours sets EngineHours field to given value.

### HasEngineHours

`func (o *BaseVehicleLocation) HasEngineHours() bool`

HasEngineHours returns a boolean if a field has been set.

### SetEngineHoursNil

`func (o *BaseVehicleLocation) SetEngineHoursNil(b bool)`

 SetEngineHoursNil sets the value for EngineHours to be an explicit nil

### UnsetEngineHours
`func (o *BaseVehicleLocation) UnsetEngineHours()`

UnsetEngineHours ensures that no value is present for EngineHours, not even an explicit nil
### GetOilPressure

`func (o *BaseVehicleLocation) GetOilPressure() float32`

GetOilPressure returns the OilPressure field if non-nil, zero value otherwise.

### GetOilPressureOk

`func (o *BaseVehicleLocation) GetOilPressureOk() (*float32, bool)`

GetOilPressureOk returns a tuple with the OilPressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOilPressure

`func (o *BaseVehicleLocation) SetOilPressure(v float32)`

SetOilPressure sets OilPressure field to given value.

### HasOilPressure

`func (o *BaseVehicleLocation) HasOilPressure() bool`

HasOilPressure returns a boolean if a field has been set.

### SetOilPressureNil

`func (o *BaseVehicleLocation) SetOilPressureNil(b bool)`

 SetOilPressureNil sets the value for OilPressure to be an explicit nil

### UnsetOilPressure
`func (o *BaseVehicleLocation) UnsetOilPressure()`

UnsetOilPressure ensures that no value is present for OilPressure, not even an explicit nil
### GetCoolantTemperature

`func (o *BaseVehicleLocation) GetCoolantTemperature() float32`

GetCoolantTemperature returns the CoolantTemperature field if non-nil, zero value otherwise.

### GetCoolantTemperatureOk

`func (o *BaseVehicleLocation) GetCoolantTemperatureOk() (*float32, bool)`

GetCoolantTemperatureOk returns a tuple with the CoolantTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolantTemperature

`func (o *BaseVehicleLocation) SetCoolantTemperature(v float32)`

SetCoolantTemperature sets CoolantTemperature field to given value.

### HasCoolantTemperature

`func (o *BaseVehicleLocation) HasCoolantTemperature() bool`

HasCoolantTemperature returns a boolean if a field has been set.

### SetCoolantTemperatureNil

`func (o *BaseVehicleLocation) SetCoolantTemperatureNil(b bool)`

 SetCoolantTemperatureNil sets the value for CoolantTemperature to be an explicit nil

### UnsetCoolantTemperature
`func (o *BaseVehicleLocation) UnsetCoolantTemperature()`

UnsetCoolantTemperature ensures that no value is present for CoolantTemperature, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


