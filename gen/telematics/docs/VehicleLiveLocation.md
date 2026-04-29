# VehicleLiveLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VehicleId** | **string** | Unique Catena identifier for the vehicle. | 
**Vin** | **NullableString** |  | 
**VehicleName** | **NullableString** |  | 
**TspId** | **NullableString** |  | 
**TspSlug** | **NullableString** |  | 
**SourceName** | [**NullableTspEnum**](TspEnum.md) |  | 
**DriverId** | **NullableString** |  | 
**DriverName** | **NullableString** |  | 
**H3Index11** | **NullableInt32** |  | 
**Speed** | **NullableFloat32** |  | 
**Odometer** | **NullableFloat32** |  | 
**FuelLevel** | **NullableFloat32** |  | 
**EngineHours** | **NullableFloat32** |  | 
**OilPressure** | **NullableFloat32** |  | 
**CoolantTemperature** | **NullableFloat32** |  | 
**OccurredAt** | **time.Time** | Timestamp (UTC) when this telemetry data was recorded by the vehicle. | 
**Location** | [**NullableLocation7**](Location7.md) |  | 

## Methods

### NewVehicleLiveLocation

`func NewVehicleLiveLocation(vehicleId string, vin NullableString, vehicleName NullableString, tspId NullableString, tspSlug NullableString, sourceName NullableTspEnum, driverId NullableString, driverName NullableString, h3Index11 NullableInt32, speed NullableFloat32, odometer NullableFloat32, fuelLevel NullableFloat32, engineHours NullableFloat32, oilPressure NullableFloat32, coolantTemperature NullableFloat32, occurredAt time.Time, location NullableLocation7, ) *VehicleLiveLocation`

NewVehicleLiveLocation instantiates a new VehicleLiveLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleLiveLocationWithDefaults

`func NewVehicleLiveLocationWithDefaults() *VehicleLiveLocation`

NewVehicleLiveLocationWithDefaults instantiates a new VehicleLiveLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVehicleId

`func (o *VehicleLiveLocation) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *VehicleLiveLocation) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *VehicleLiveLocation) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.


### GetVin

`func (o *VehicleLiveLocation) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *VehicleLiveLocation) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *VehicleLiveLocation) SetVin(v string)`

SetVin sets Vin field to given value.


### SetVinNil

`func (o *VehicleLiveLocation) SetVinNil(b bool)`

 SetVinNil sets the value for Vin to be an explicit nil

### UnsetVin
`func (o *VehicleLiveLocation) UnsetVin()`

UnsetVin ensures that no value is present for Vin, not even an explicit nil
### GetVehicleName

`func (o *VehicleLiveLocation) GetVehicleName() string`

GetVehicleName returns the VehicleName field if non-nil, zero value otherwise.

### GetVehicleNameOk

`func (o *VehicleLiveLocation) GetVehicleNameOk() (*string, bool)`

GetVehicleNameOk returns a tuple with the VehicleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleName

`func (o *VehicleLiveLocation) SetVehicleName(v string)`

SetVehicleName sets VehicleName field to given value.


### SetVehicleNameNil

`func (o *VehicleLiveLocation) SetVehicleNameNil(b bool)`

 SetVehicleNameNil sets the value for VehicleName to be an explicit nil

### UnsetVehicleName
`func (o *VehicleLiveLocation) UnsetVehicleName()`

UnsetVehicleName ensures that no value is present for VehicleName, not even an explicit nil
### GetTspId

`func (o *VehicleLiveLocation) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *VehicleLiveLocation) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *VehicleLiveLocation) SetTspId(v string)`

SetTspId sets TspId field to given value.


### SetTspIdNil

`func (o *VehicleLiveLocation) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *VehicleLiveLocation) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *VehicleLiveLocation) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *VehicleLiveLocation) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *VehicleLiveLocation) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.


### SetTspSlugNil

`func (o *VehicleLiveLocation) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *VehicleLiveLocation) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *VehicleLiveLocation) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *VehicleLiveLocation) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *VehicleLiveLocation) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### SetSourceNameNil

`func (o *VehicleLiveLocation) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *VehicleLiveLocation) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetDriverId

`func (o *VehicleLiveLocation) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *VehicleLiveLocation) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *VehicleLiveLocation) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.


### SetDriverIdNil

`func (o *VehicleLiveLocation) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *VehicleLiveLocation) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetDriverName

`func (o *VehicleLiveLocation) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *VehicleLiveLocation) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *VehicleLiveLocation) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.


### SetDriverNameNil

`func (o *VehicleLiveLocation) SetDriverNameNil(b bool)`

 SetDriverNameNil sets the value for DriverName to be an explicit nil

### UnsetDriverName
`func (o *VehicleLiveLocation) UnsetDriverName()`

UnsetDriverName ensures that no value is present for DriverName, not even an explicit nil
### GetH3Index11

`func (o *VehicleLiveLocation) GetH3Index11() int32`

GetH3Index11 returns the H3Index11 field if non-nil, zero value otherwise.

### GetH3Index11Ok

`func (o *VehicleLiveLocation) GetH3Index11Ok() (*int32, bool)`

GetH3Index11Ok returns a tuple with the H3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH3Index11

`func (o *VehicleLiveLocation) SetH3Index11(v int32)`

SetH3Index11 sets H3Index11 field to given value.


### SetH3Index11Nil

`func (o *VehicleLiveLocation) SetH3Index11Nil(b bool)`

 SetH3Index11Nil sets the value for H3Index11 to be an explicit nil

### UnsetH3Index11
`func (o *VehicleLiveLocation) UnsetH3Index11()`

UnsetH3Index11 ensures that no value is present for H3Index11, not even an explicit nil
### GetSpeed

`func (o *VehicleLiveLocation) GetSpeed() float32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *VehicleLiveLocation) GetSpeedOk() (*float32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *VehicleLiveLocation) SetSpeed(v float32)`

SetSpeed sets Speed field to given value.


### SetSpeedNil

`func (o *VehicleLiveLocation) SetSpeedNil(b bool)`

 SetSpeedNil sets the value for Speed to be an explicit nil

### UnsetSpeed
`func (o *VehicleLiveLocation) UnsetSpeed()`

UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
### GetOdometer

`func (o *VehicleLiveLocation) GetOdometer() float32`

GetOdometer returns the Odometer field if non-nil, zero value otherwise.

### GetOdometerOk

`func (o *VehicleLiveLocation) GetOdometerOk() (*float32, bool)`

GetOdometerOk returns a tuple with the Odometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometer

`func (o *VehicleLiveLocation) SetOdometer(v float32)`

SetOdometer sets Odometer field to given value.


### SetOdometerNil

`func (o *VehicleLiveLocation) SetOdometerNil(b bool)`

 SetOdometerNil sets the value for Odometer to be an explicit nil

### UnsetOdometer
`func (o *VehicleLiveLocation) UnsetOdometer()`

UnsetOdometer ensures that no value is present for Odometer, not even an explicit nil
### GetFuelLevel

`func (o *VehicleLiveLocation) GetFuelLevel() float32`

GetFuelLevel returns the FuelLevel field if non-nil, zero value otherwise.

### GetFuelLevelOk

`func (o *VehicleLiveLocation) GetFuelLevelOk() (*float32, bool)`

GetFuelLevelOk returns a tuple with the FuelLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelLevel

`func (o *VehicleLiveLocation) SetFuelLevel(v float32)`

SetFuelLevel sets FuelLevel field to given value.


### SetFuelLevelNil

`func (o *VehicleLiveLocation) SetFuelLevelNil(b bool)`

 SetFuelLevelNil sets the value for FuelLevel to be an explicit nil

### UnsetFuelLevel
`func (o *VehicleLiveLocation) UnsetFuelLevel()`

UnsetFuelLevel ensures that no value is present for FuelLevel, not even an explicit nil
### GetEngineHours

`func (o *VehicleLiveLocation) GetEngineHours() float32`

GetEngineHours returns the EngineHours field if non-nil, zero value otherwise.

### GetEngineHoursOk

`func (o *VehicleLiveLocation) GetEngineHoursOk() (*float32, bool)`

GetEngineHoursOk returns a tuple with the EngineHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineHours

`func (o *VehicleLiveLocation) SetEngineHours(v float32)`

SetEngineHours sets EngineHours field to given value.


### SetEngineHoursNil

`func (o *VehicleLiveLocation) SetEngineHoursNil(b bool)`

 SetEngineHoursNil sets the value for EngineHours to be an explicit nil

### UnsetEngineHours
`func (o *VehicleLiveLocation) UnsetEngineHours()`

UnsetEngineHours ensures that no value is present for EngineHours, not even an explicit nil
### GetOilPressure

`func (o *VehicleLiveLocation) GetOilPressure() float32`

GetOilPressure returns the OilPressure field if non-nil, zero value otherwise.

### GetOilPressureOk

`func (o *VehicleLiveLocation) GetOilPressureOk() (*float32, bool)`

GetOilPressureOk returns a tuple with the OilPressure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOilPressure

`func (o *VehicleLiveLocation) SetOilPressure(v float32)`

SetOilPressure sets OilPressure field to given value.


### SetOilPressureNil

`func (o *VehicleLiveLocation) SetOilPressureNil(b bool)`

 SetOilPressureNil sets the value for OilPressure to be an explicit nil

### UnsetOilPressure
`func (o *VehicleLiveLocation) UnsetOilPressure()`

UnsetOilPressure ensures that no value is present for OilPressure, not even an explicit nil
### GetCoolantTemperature

`func (o *VehicleLiveLocation) GetCoolantTemperature() float32`

GetCoolantTemperature returns the CoolantTemperature field if non-nil, zero value otherwise.

### GetCoolantTemperatureOk

`func (o *VehicleLiveLocation) GetCoolantTemperatureOk() (*float32, bool)`

GetCoolantTemperatureOk returns a tuple with the CoolantTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolantTemperature

`func (o *VehicleLiveLocation) SetCoolantTemperature(v float32)`

SetCoolantTemperature sets CoolantTemperature field to given value.


### SetCoolantTemperatureNil

`func (o *VehicleLiveLocation) SetCoolantTemperatureNil(b bool)`

 SetCoolantTemperatureNil sets the value for CoolantTemperature to be an explicit nil

### UnsetCoolantTemperature
`func (o *VehicleLiveLocation) UnsetCoolantTemperature()`

UnsetCoolantTemperature ensures that no value is present for CoolantTemperature, not even an explicit nil
### GetOccurredAt

`func (o *VehicleLiveLocation) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *VehicleLiveLocation) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *VehicleLiveLocation) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetLocation

`func (o *VehicleLiveLocation) GetLocation() Location7`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VehicleLiveLocation) GetLocationOk() (*Location7, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VehicleLiveLocation) SetLocation(v Location7)`

SetLocation sets Location field to given value.


### SetLocationNil

`func (o *VehicleLiveLocation) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *VehicleLiveLocation) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


