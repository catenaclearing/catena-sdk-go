# FuelTransactionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | Unique identifier of the connection at Catena Telematics which will be used to create this resource. A connection represents a Fleet/TSP pairing. | 
**FleetRef** | Pointer to **NullableString** |  | [optional] 
**SourceDriverId** | Pointer to **NullableString** |  | [optional] 
**SourceVehicleId** | Pointer to **NullableString** |  | [optional] 
**SourceCoDriverId** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to [**NullablePoint**](Point.md) |  | [optional] 
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

`func NewFuelTransactionCreate(connectionId string, fuelVolume float32, fuelVolumeUnit FuelVolumeUnitEnum, ) *FuelTransactionCreate`

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


### GetFleetRef

`func (o *FuelTransactionCreate) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *FuelTransactionCreate) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *FuelTransactionCreate) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *FuelTransactionCreate) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *FuelTransactionCreate) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *FuelTransactionCreate) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetSourceDriverId

`func (o *FuelTransactionCreate) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *FuelTransactionCreate) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *FuelTransactionCreate) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *FuelTransactionCreate) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *FuelTransactionCreate) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *FuelTransactionCreate) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSourceVehicleId

`func (o *FuelTransactionCreate) GetSourceVehicleId() string`

GetSourceVehicleId returns the SourceVehicleId field if non-nil, zero value otherwise.

### GetSourceVehicleIdOk

`func (o *FuelTransactionCreate) GetSourceVehicleIdOk() (*string, bool)`

GetSourceVehicleIdOk returns a tuple with the SourceVehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVehicleId

`func (o *FuelTransactionCreate) SetSourceVehicleId(v string)`

SetSourceVehicleId sets SourceVehicleId field to given value.

### HasSourceVehicleId

`func (o *FuelTransactionCreate) HasSourceVehicleId() bool`

HasSourceVehicleId returns a boolean if a field has been set.

### SetSourceVehicleIdNil

`func (o *FuelTransactionCreate) SetSourceVehicleIdNil(b bool)`

 SetSourceVehicleIdNil sets the value for SourceVehicleId to be an explicit nil

### UnsetSourceVehicleId
`func (o *FuelTransactionCreate) UnsetSourceVehicleId()`

UnsetSourceVehicleId ensures that no value is present for SourceVehicleId, not even an explicit nil
### GetSourceCoDriverId

`func (o *FuelTransactionCreate) GetSourceCoDriverId() string`

GetSourceCoDriverId returns the SourceCoDriverId field if non-nil, zero value otherwise.

### GetSourceCoDriverIdOk

`func (o *FuelTransactionCreate) GetSourceCoDriverIdOk() (*string, bool)`

GetSourceCoDriverIdOk returns a tuple with the SourceCoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCoDriverId

`func (o *FuelTransactionCreate) SetSourceCoDriverId(v string)`

SetSourceCoDriverId sets SourceCoDriverId field to given value.

### HasSourceCoDriverId

`func (o *FuelTransactionCreate) HasSourceCoDriverId() bool`

HasSourceCoDriverId returns a boolean if a field has been set.

### SetSourceCoDriverIdNil

`func (o *FuelTransactionCreate) SetSourceCoDriverIdNil(b bool)`

 SetSourceCoDriverIdNil sets the value for SourceCoDriverId to be an explicit nil

### UnsetSourceCoDriverId
`func (o *FuelTransactionCreate) UnsetSourceCoDriverId()`

UnsetSourceCoDriverId ensures that no value is present for SourceCoDriverId, not even an explicit nil
### GetLocation

`func (o *FuelTransactionCreate) GetLocation() Point`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FuelTransactionCreate) GetLocationOk() (*Point, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FuelTransactionCreate) SetLocation(v Point)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FuelTransactionCreate) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *FuelTransactionCreate) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *FuelTransactionCreate) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
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


