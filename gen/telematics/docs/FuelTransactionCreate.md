# FuelTransactionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | Unique identifier of the connection at Catena Telematics which will be used to create this resource. A connection represents a Fleet/TSP pairing. | 
**DriverId** | **string** | Identifier for the driver (from the source system) | 
**VehicleId** | **string** | Identifier for the vehicle (from the source system) | 
**CoDriverId** | Pointer to **NullableString** |  | [optional] 
**Latitude** | Pointer to **NullableFloat32** |  | [optional] 
**Longitude** | Pointer to **NullableFloat32** |  | [optional] 
**LocationString** | Pointer to **NullableString** |  | [optional] 
**Odometer** | Pointer to **NullableFloat32** |  | [optional] 
**OdometerUnit** | Pointer to [**NullableDistanceUnitEnum**](DistanceUnitEnum.md) |  | [optional] 
**TransactionTime** | Pointer to **NullableTime** |  | [optional] 
**TransactionJurisdiction** | Pointer to **NullableString** |  | [optional] 
**FuelType** | Pointer to [**NullableEngineType**](EngineType.md) |  | [optional] 
**FuelVolume** | **float32** | Volume of fuel purchased | 
**FuelVolumeUnit** | [**FuelVolumeUnitEnum**](FuelVolumeUnitEnum.md) | Unit for fuel_volume | 
**FuelVendor** | Pointer to **NullableString** |  | [optional] 
**TotalCost** | Pointer to [**NullableTotalCost**](TotalCost.md) |  | [optional] 
**Currency** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFuelTransactionCreate

`func NewFuelTransactionCreate(connectionId string, driverId string, vehicleId string, fuelVolume float32, fuelVolumeUnit FuelVolumeUnitEnum, ) *FuelTransactionCreate`

NewFuelTransactionCreate instantiates a new FuelTransactionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuelTransactionCreateWithDefaults

`func NewFuelTransactionCreateWithDefaults() *FuelTransactionCreate`

NewFuelTransactionCreateWithDefaults instantiates a new FuelTransactionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *FuelTransactionCreate) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *FuelTransactionCreate) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *FuelTransactionCreate) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetDriverId

`func (o *FuelTransactionCreate) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *FuelTransactionCreate) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *FuelTransactionCreate) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.


### GetVehicleId

`func (o *FuelTransactionCreate) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *FuelTransactionCreate) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *FuelTransactionCreate) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.


### GetCoDriverId

`func (o *FuelTransactionCreate) GetCoDriverId() string`

GetCoDriverId returns the CoDriverId field if non-nil, zero value otherwise.

### GetCoDriverIdOk

`func (o *FuelTransactionCreate) GetCoDriverIdOk() (*string, bool)`

GetCoDriverIdOk returns a tuple with the CoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDriverId

`func (o *FuelTransactionCreate) SetCoDriverId(v string)`

SetCoDriverId sets CoDriverId field to given value.

### HasCoDriverId

`func (o *FuelTransactionCreate) HasCoDriverId() bool`

HasCoDriverId returns a boolean if a field has been set.

### SetCoDriverIdNil

`func (o *FuelTransactionCreate) SetCoDriverIdNil(b bool)`

 SetCoDriverIdNil sets the value for CoDriverId to be an explicit nil

### UnsetCoDriverId
`func (o *FuelTransactionCreate) UnsetCoDriverId()`

UnsetCoDriverId ensures that no value is present for CoDriverId, not even an explicit nil
### GetLatitude

`func (o *FuelTransactionCreate) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *FuelTransactionCreate) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *FuelTransactionCreate) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *FuelTransactionCreate) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *FuelTransactionCreate) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *FuelTransactionCreate) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *FuelTransactionCreate) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *FuelTransactionCreate) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *FuelTransactionCreate) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *FuelTransactionCreate) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *FuelTransactionCreate) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *FuelTransactionCreate) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetLocationString

`func (o *FuelTransactionCreate) GetLocationString() string`

GetLocationString returns the LocationString field if non-nil, zero value otherwise.

### GetLocationStringOk

`func (o *FuelTransactionCreate) GetLocationStringOk() (*string, bool)`

GetLocationStringOk returns a tuple with the LocationString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationString

`func (o *FuelTransactionCreate) SetLocationString(v string)`

SetLocationString sets LocationString field to given value.

### HasLocationString

`func (o *FuelTransactionCreate) HasLocationString() bool`

HasLocationString returns a boolean if a field has been set.

### SetLocationStringNil

`func (o *FuelTransactionCreate) SetLocationStringNil(b bool)`

 SetLocationStringNil sets the value for LocationString to be an explicit nil

### UnsetLocationString
`func (o *FuelTransactionCreate) UnsetLocationString()`

UnsetLocationString ensures that no value is present for LocationString, not even an explicit nil
### GetOdometer

`func (o *FuelTransactionCreate) GetOdometer() float32`

GetOdometer returns the Odometer field if non-nil, zero value otherwise.

### GetOdometerOk

`func (o *FuelTransactionCreate) GetOdometerOk() (*float32, bool)`

GetOdometerOk returns a tuple with the Odometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometer

`func (o *FuelTransactionCreate) SetOdometer(v float32)`

SetOdometer sets Odometer field to given value.

### HasOdometer

`func (o *FuelTransactionCreate) HasOdometer() bool`

HasOdometer returns a boolean if a field has been set.

### SetOdometerNil

`func (o *FuelTransactionCreate) SetOdometerNil(b bool)`

 SetOdometerNil sets the value for Odometer to be an explicit nil

### UnsetOdometer
`func (o *FuelTransactionCreate) UnsetOdometer()`

UnsetOdometer ensures that no value is present for Odometer, not even an explicit nil
### GetOdometerUnit

`func (o *FuelTransactionCreate) GetOdometerUnit() DistanceUnitEnum`

GetOdometerUnit returns the OdometerUnit field if non-nil, zero value otherwise.

### GetOdometerUnitOk

`func (o *FuelTransactionCreate) GetOdometerUnitOk() (*DistanceUnitEnum, bool)`

GetOdometerUnitOk returns a tuple with the OdometerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerUnit

`func (o *FuelTransactionCreate) SetOdometerUnit(v DistanceUnitEnum)`

SetOdometerUnit sets OdometerUnit field to given value.

### HasOdometerUnit

`func (o *FuelTransactionCreate) HasOdometerUnit() bool`

HasOdometerUnit returns a boolean if a field has been set.

### SetOdometerUnitNil

`func (o *FuelTransactionCreate) SetOdometerUnitNil(b bool)`

 SetOdometerUnitNil sets the value for OdometerUnit to be an explicit nil

### UnsetOdometerUnit
`func (o *FuelTransactionCreate) UnsetOdometerUnit()`

UnsetOdometerUnit ensures that no value is present for OdometerUnit, not even an explicit nil
### GetTransactionTime

`func (o *FuelTransactionCreate) GetTransactionTime() time.Time`

GetTransactionTime returns the TransactionTime field if non-nil, zero value otherwise.

### GetTransactionTimeOk

`func (o *FuelTransactionCreate) GetTransactionTimeOk() (*time.Time, bool)`

GetTransactionTimeOk returns a tuple with the TransactionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTime

`func (o *FuelTransactionCreate) SetTransactionTime(v time.Time)`

SetTransactionTime sets TransactionTime field to given value.

### HasTransactionTime

`func (o *FuelTransactionCreate) HasTransactionTime() bool`

HasTransactionTime returns a boolean if a field has been set.

### SetTransactionTimeNil

`func (o *FuelTransactionCreate) SetTransactionTimeNil(b bool)`

 SetTransactionTimeNil sets the value for TransactionTime to be an explicit nil

### UnsetTransactionTime
`func (o *FuelTransactionCreate) UnsetTransactionTime()`

UnsetTransactionTime ensures that no value is present for TransactionTime, not even an explicit nil
### GetTransactionJurisdiction

`func (o *FuelTransactionCreate) GetTransactionJurisdiction() string`

GetTransactionJurisdiction returns the TransactionJurisdiction field if non-nil, zero value otherwise.

### GetTransactionJurisdictionOk

`func (o *FuelTransactionCreate) GetTransactionJurisdictionOk() (*string, bool)`

GetTransactionJurisdictionOk returns a tuple with the TransactionJurisdiction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionJurisdiction

`func (o *FuelTransactionCreate) SetTransactionJurisdiction(v string)`

SetTransactionJurisdiction sets TransactionJurisdiction field to given value.

### HasTransactionJurisdiction

`func (o *FuelTransactionCreate) HasTransactionJurisdiction() bool`

HasTransactionJurisdiction returns a boolean if a field has been set.

### SetTransactionJurisdictionNil

`func (o *FuelTransactionCreate) SetTransactionJurisdictionNil(b bool)`

 SetTransactionJurisdictionNil sets the value for TransactionJurisdiction to be an explicit nil

### UnsetTransactionJurisdiction
`func (o *FuelTransactionCreate) UnsetTransactionJurisdiction()`

UnsetTransactionJurisdiction ensures that no value is present for TransactionJurisdiction, not even an explicit nil
### GetFuelType

`func (o *FuelTransactionCreate) GetFuelType() EngineType`

GetFuelType returns the FuelType field if non-nil, zero value otherwise.

### GetFuelTypeOk

`func (o *FuelTransactionCreate) GetFuelTypeOk() (*EngineType, bool)`

GetFuelTypeOk returns a tuple with the FuelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelType

`func (o *FuelTransactionCreate) SetFuelType(v EngineType)`

SetFuelType sets FuelType field to given value.

### HasFuelType

`func (o *FuelTransactionCreate) HasFuelType() bool`

HasFuelType returns a boolean if a field has been set.

### SetFuelTypeNil

`func (o *FuelTransactionCreate) SetFuelTypeNil(b bool)`

 SetFuelTypeNil sets the value for FuelType to be an explicit nil

### UnsetFuelType
`func (o *FuelTransactionCreate) UnsetFuelType()`

UnsetFuelType ensures that no value is present for FuelType, not even an explicit nil
### GetFuelVolume

`func (o *FuelTransactionCreate) GetFuelVolume() float32`

GetFuelVolume returns the FuelVolume field if non-nil, zero value otherwise.

### GetFuelVolumeOk

`func (o *FuelTransactionCreate) GetFuelVolumeOk() (*float32, bool)`

GetFuelVolumeOk returns a tuple with the FuelVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelVolume

`func (o *FuelTransactionCreate) SetFuelVolume(v float32)`

SetFuelVolume sets FuelVolume field to given value.


### GetFuelVolumeUnit

`func (o *FuelTransactionCreate) GetFuelVolumeUnit() FuelVolumeUnitEnum`

GetFuelVolumeUnit returns the FuelVolumeUnit field if non-nil, zero value otherwise.

### GetFuelVolumeUnitOk

`func (o *FuelTransactionCreate) GetFuelVolumeUnitOk() (*FuelVolumeUnitEnum, bool)`

GetFuelVolumeUnitOk returns a tuple with the FuelVolumeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelVolumeUnit

`func (o *FuelTransactionCreate) SetFuelVolumeUnit(v FuelVolumeUnitEnum)`

SetFuelVolumeUnit sets FuelVolumeUnit field to given value.


### GetFuelVendor

`func (o *FuelTransactionCreate) GetFuelVendor() string`

GetFuelVendor returns the FuelVendor field if non-nil, zero value otherwise.

### GetFuelVendorOk

`func (o *FuelTransactionCreate) GetFuelVendorOk() (*string, bool)`

GetFuelVendorOk returns a tuple with the FuelVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelVendor

`func (o *FuelTransactionCreate) SetFuelVendor(v string)`

SetFuelVendor sets FuelVendor field to given value.

### HasFuelVendor

`func (o *FuelTransactionCreate) HasFuelVendor() bool`

HasFuelVendor returns a boolean if a field has been set.

### SetFuelVendorNil

`func (o *FuelTransactionCreate) SetFuelVendorNil(b bool)`

 SetFuelVendorNil sets the value for FuelVendor to be an explicit nil

### UnsetFuelVendor
`func (o *FuelTransactionCreate) UnsetFuelVendor()`

UnsetFuelVendor ensures that no value is present for FuelVendor, not even an explicit nil
### GetTotalCost

`func (o *FuelTransactionCreate) GetTotalCost() TotalCost`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *FuelTransactionCreate) GetTotalCostOk() (*TotalCost, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *FuelTransactionCreate) SetTotalCost(v TotalCost)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *FuelTransactionCreate) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### SetTotalCostNil

`func (o *FuelTransactionCreate) SetTotalCostNil(b bool)`

 SetTotalCostNil sets the value for TotalCost to be an explicit nil

### UnsetTotalCost
`func (o *FuelTransactionCreate) UnsetTotalCost()`

UnsetTotalCost ensures that no value is present for TotalCost, not even an explicit nil
### GetCurrency

`func (o *FuelTransactionCreate) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FuelTransactionCreate) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FuelTransactionCreate) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FuelTransactionCreate) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *FuelTransactionCreate) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *FuelTransactionCreate) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


