# VehicleSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FleetId** | **string** | Catena fleet identifier. | 
**FleetRef** | Pointer to **NullableString** |  | [optional] 
**ConnectionId** | **string** | Catena connection identifier through which this vehicle was ingested. | 
**VehicleId** | **string** | Unique Catena identifier for the vehicle. | 
**SourceName** | Pointer to **NullableString** |  | [optional] 
**SourceId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**VehicleName** | Pointer to **NullableString** |  | [optional] 
**Oem** | Pointer to **NullableString** |  | [optional] 
**ModelType** | Pointer to **NullableString** |  | [optional] 
**ModelYear** | Pointer to **NullableInt32** |  | [optional] 
**Vin** | Pointer to **NullableString** |  | [optional] 
**EngineVin** | Pointer to **NullableString** |  | [optional] 
**LicensePlateCountry** | Pointer to **NullableString** |  | [optional] 
**LicensePlateRegion** | Pointer to **NullableString** |  | [optional] 
**LicensePlateNumber** | Pointer to **NullableString** |  | [optional] 
**LastLocationPing** | Pointer to **NullableTime** |  | [optional] 
**LastLocationH3Index11** | Pointer to **NullableInt32** |  | [optional] 
**LastLocation** | Pointer to [**NullableLastLocation**](LastLocation.md) |  | [optional] 
**LastFuelLevel** | Pointer to **NullableFloat32** |  | [optional] 
**LastOdometerReading** | Pointer to **NullableFloat32** |  | [optional] 
**LastEngineHours** | Pointer to **NullableFloat32** |  | [optional] 
**LastSpeedReading** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewVehicleSummary

`func NewVehicleSummary(fleetId string, connectionId string, vehicleId string, ) *VehicleSummary`

NewVehicleSummary instantiates a new VehicleSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleSummaryWithDefaults

`func NewVehicleSummaryWithDefaults() *VehicleSummary`

NewVehicleSummaryWithDefaults instantiates a new VehicleSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *VehicleSummary) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *VehicleSummary) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *VehicleSummary) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *VehicleSummary) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *VehicleSummary) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *VehicleSummary) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *VehicleSummary) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *VehicleSummary) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *VehicleSummary) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetConnectionId

`func (o *VehicleSummary) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *VehicleSummary) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *VehicleSummary) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetVehicleId

`func (o *VehicleSummary) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *VehicleSummary) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *VehicleSummary) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.


### GetSourceName

`func (o *VehicleSummary) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *VehicleSummary) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *VehicleSummary) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *VehicleSummary) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *VehicleSummary) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *VehicleSummary) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetSourceId

`func (o *VehicleSummary) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *VehicleSummary) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *VehicleSummary) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *VehicleSummary) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *VehicleSummary) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *VehicleSummary) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetStatus

`func (o *VehicleSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VehicleSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VehicleSummary) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VehicleSummary) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VehicleSummary) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VehicleSummary) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetVehicleName

`func (o *VehicleSummary) GetVehicleName() string`

GetVehicleName returns the VehicleName field if non-nil, zero value otherwise.

### GetVehicleNameOk

`func (o *VehicleSummary) GetVehicleNameOk() (*string, bool)`

GetVehicleNameOk returns a tuple with the VehicleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleName

`func (o *VehicleSummary) SetVehicleName(v string)`

SetVehicleName sets VehicleName field to given value.

### HasVehicleName

`func (o *VehicleSummary) HasVehicleName() bool`

HasVehicleName returns a boolean if a field has been set.

### SetVehicleNameNil

`func (o *VehicleSummary) SetVehicleNameNil(b bool)`

 SetVehicleNameNil sets the value for VehicleName to be an explicit nil

### UnsetVehicleName
`func (o *VehicleSummary) UnsetVehicleName()`

UnsetVehicleName ensures that no value is present for VehicleName, not even an explicit nil
### GetOem

`func (o *VehicleSummary) GetOem() string`

GetOem returns the Oem field if non-nil, zero value otherwise.

### GetOemOk

`func (o *VehicleSummary) GetOemOk() (*string, bool)`

GetOemOk returns a tuple with the Oem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOem

`func (o *VehicleSummary) SetOem(v string)`

SetOem sets Oem field to given value.

### HasOem

`func (o *VehicleSummary) HasOem() bool`

HasOem returns a boolean if a field has been set.

### SetOemNil

`func (o *VehicleSummary) SetOemNil(b bool)`

 SetOemNil sets the value for Oem to be an explicit nil

### UnsetOem
`func (o *VehicleSummary) UnsetOem()`

UnsetOem ensures that no value is present for Oem, not even an explicit nil
### GetModelType

`func (o *VehicleSummary) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *VehicleSummary) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *VehicleSummary) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *VehicleSummary) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### SetModelTypeNil

`func (o *VehicleSummary) SetModelTypeNil(b bool)`

 SetModelTypeNil sets the value for ModelType to be an explicit nil

### UnsetModelType
`func (o *VehicleSummary) UnsetModelType()`

UnsetModelType ensures that no value is present for ModelType, not even an explicit nil
### GetModelYear

`func (o *VehicleSummary) GetModelYear() int32`

GetModelYear returns the ModelYear field if non-nil, zero value otherwise.

### GetModelYearOk

`func (o *VehicleSummary) GetModelYearOk() (*int32, bool)`

GetModelYearOk returns a tuple with the ModelYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelYear

`func (o *VehicleSummary) SetModelYear(v int32)`

SetModelYear sets ModelYear field to given value.

### HasModelYear

`func (o *VehicleSummary) HasModelYear() bool`

HasModelYear returns a boolean if a field has been set.

### SetModelYearNil

`func (o *VehicleSummary) SetModelYearNil(b bool)`

 SetModelYearNil sets the value for ModelYear to be an explicit nil

### UnsetModelYear
`func (o *VehicleSummary) UnsetModelYear()`

UnsetModelYear ensures that no value is present for ModelYear, not even an explicit nil
### GetVin

`func (o *VehicleSummary) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *VehicleSummary) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *VehicleSummary) SetVin(v string)`

SetVin sets Vin field to given value.

### HasVin

`func (o *VehicleSummary) HasVin() bool`

HasVin returns a boolean if a field has been set.

### SetVinNil

`func (o *VehicleSummary) SetVinNil(b bool)`

 SetVinNil sets the value for Vin to be an explicit nil

### UnsetVin
`func (o *VehicleSummary) UnsetVin()`

UnsetVin ensures that no value is present for Vin, not even an explicit nil
### GetEngineVin

`func (o *VehicleSummary) GetEngineVin() string`

GetEngineVin returns the EngineVin field if non-nil, zero value otherwise.

### GetEngineVinOk

`func (o *VehicleSummary) GetEngineVinOk() (*string, bool)`

GetEngineVinOk returns a tuple with the EngineVin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVin

`func (o *VehicleSummary) SetEngineVin(v string)`

SetEngineVin sets EngineVin field to given value.

### HasEngineVin

`func (o *VehicleSummary) HasEngineVin() bool`

HasEngineVin returns a boolean if a field has been set.

### SetEngineVinNil

`func (o *VehicleSummary) SetEngineVinNil(b bool)`

 SetEngineVinNil sets the value for EngineVin to be an explicit nil

### UnsetEngineVin
`func (o *VehicleSummary) UnsetEngineVin()`

UnsetEngineVin ensures that no value is present for EngineVin, not even an explicit nil
### GetLicensePlateCountry

`func (o *VehicleSummary) GetLicensePlateCountry() string`

GetLicensePlateCountry returns the LicensePlateCountry field if non-nil, zero value otherwise.

### GetLicensePlateCountryOk

`func (o *VehicleSummary) GetLicensePlateCountryOk() (*string, bool)`

GetLicensePlateCountryOk returns a tuple with the LicensePlateCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateCountry

`func (o *VehicleSummary) SetLicensePlateCountry(v string)`

SetLicensePlateCountry sets LicensePlateCountry field to given value.

### HasLicensePlateCountry

`func (o *VehicleSummary) HasLicensePlateCountry() bool`

HasLicensePlateCountry returns a boolean if a field has been set.

### SetLicensePlateCountryNil

`func (o *VehicleSummary) SetLicensePlateCountryNil(b bool)`

 SetLicensePlateCountryNil sets the value for LicensePlateCountry to be an explicit nil

### UnsetLicensePlateCountry
`func (o *VehicleSummary) UnsetLicensePlateCountry()`

UnsetLicensePlateCountry ensures that no value is present for LicensePlateCountry, not even an explicit nil
### GetLicensePlateRegion

`func (o *VehicleSummary) GetLicensePlateRegion() string`

GetLicensePlateRegion returns the LicensePlateRegion field if non-nil, zero value otherwise.

### GetLicensePlateRegionOk

`func (o *VehicleSummary) GetLicensePlateRegionOk() (*string, bool)`

GetLicensePlateRegionOk returns a tuple with the LicensePlateRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateRegion

`func (o *VehicleSummary) SetLicensePlateRegion(v string)`

SetLicensePlateRegion sets LicensePlateRegion field to given value.

### HasLicensePlateRegion

`func (o *VehicleSummary) HasLicensePlateRegion() bool`

HasLicensePlateRegion returns a boolean if a field has been set.

### SetLicensePlateRegionNil

`func (o *VehicleSummary) SetLicensePlateRegionNil(b bool)`

 SetLicensePlateRegionNil sets the value for LicensePlateRegion to be an explicit nil

### UnsetLicensePlateRegion
`func (o *VehicleSummary) UnsetLicensePlateRegion()`

UnsetLicensePlateRegion ensures that no value is present for LicensePlateRegion, not even an explicit nil
### GetLicensePlateNumber

`func (o *VehicleSummary) GetLicensePlateNumber() string`

GetLicensePlateNumber returns the LicensePlateNumber field if non-nil, zero value otherwise.

### GetLicensePlateNumberOk

`func (o *VehicleSummary) GetLicensePlateNumberOk() (*string, bool)`

GetLicensePlateNumberOk returns a tuple with the LicensePlateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateNumber

`func (o *VehicleSummary) SetLicensePlateNumber(v string)`

SetLicensePlateNumber sets LicensePlateNumber field to given value.

### HasLicensePlateNumber

`func (o *VehicleSummary) HasLicensePlateNumber() bool`

HasLicensePlateNumber returns a boolean if a field has been set.

### SetLicensePlateNumberNil

`func (o *VehicleSummary) SetLicensePlateNumberNil(b bool)`

 SetLicensePlateNumberNil sets the value for LicensePlateNumber to be an explicit nil

### UnsetLicensePlateNumber
`func (o *VehicleSummary) UnsetLicensePlateNumber()`

UnsetLicensePlateNumber ensures that no value is present for LicensePlateNumber, not even an explicit nil
### GetLastLocationPing

`func (o *VehicleSummary) GetLastLocationPing() time.Time`

GetLastLocationPing returns the LastLocationPing field if non-nil, zero value otherwise.

### GetLastLocationPingOk

`func (o *VehicleSummary) GetLastLocationPingOk() (*time.Time, bool)`

GetLastLocationPingOk returns a tuple with the LastLocationPing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocationPing

`func (o *VehicleSummary) SetLastLocationPing(v time.Time)`

SetLastLocationPing sets LastLocationPing field to given value.

### HasLastLocationPing

`func (o *VehicleSummary) HasLastLocationPing() bool`

HasLastLocationPing returns a boolean if a field has been set.

### SetLastLocationPingNil

`func (o *VehicleSummary) SetLastLocationPingNil(b bool)`

 SetLastLocationPingNil sets the value for LastLocationPing to be an explicit nil

### UnsetLastLocationPing
`func (o *VehicleSummary) UnsetLastLocationPing()`

UnsetLastLocationPing ensures that no value is present for LastLocationPing, not even an explicit nil
### GetLastLocationH3Index11

`func (o *VehicleSummary) GetLastLocationH3Index11() int32`

GetLastLocationH3Index11 returns the LastLocationH3Index11 field if non-nil, zero value otherwise.

### GetLastLocationH3Index11Ok

`func (o *VehicleSummary) GetLastLocationH3Index11Ok() (*int32, bool)`

GetLastLocationH3Index11Ok returns a tuple with the LastLocationH3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocationH3Index11

`func (o *VehicleSummary) SetLastLocationH3Index11(v int32)`

SetLastLocationH3Index11 sets LastLocationH3Index11 field to given value.

### HasLastLocationH3Index11

`func (o *VehicleSummary) HasLastLocationH3Index11() bool`

HasLastLocationH3Index11 returns a boolean if a field has been set.

### SetLastLocationH3Index11Nil

`func (o *VehicleSummary) SetLastLocationH3Index11Nil(b bool)`

 SetLastLocationH3Index11Nil sets the value for LastLocationH3Index11 to be an explicit nil

### UnsetLastLocationH3Index11
`func (o *VehicleSummary) UnsetLastLocationH3Index11()`

UnsetLastLocationH3Index11 ensures that no value is present for LastLocationH3Index11, not even an explicit nil
### GetLastLocation

`func (o *VehicleSummary) GetLastLocation() LastLocation`

GetLastLocation returns the LastLocation field if non-nil, zero value otherwise.

### GetLastLocationOk

`func (o *VehicleSummary) GetLastLocationOk() (*LastLocation, bool)`

GetLastLocationOk returns a tuple with the LastLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLocation

`func (o *VehicleSummary) SetLastLocation(v LastLocation)`

SetLastLocation sets LastLocation field to given value.

### HasLastLocation

`func (o *VehicleSummary) HasLastLocation() bool`

HasLastLocation returns a boolean if a field has been set.

### SetLastLocationNil

`func (o *VehicleSummary) SetLastLocationNil(b bool)`

 SetLastLocationNil sets the value for LastLocation to be an explicit nil

### UnsetLastLocation
`func (o *VehicleSummary) UnsetLastLocation()`

UnsetLastLocation ensures that no value is present for LastLocation, not even an explicit nil
### GetLastFuelLevel

`func (o *VehicleSummary) GetLastFuelLevel() float32`

GetLastFuelLevel returns the LastFuelLevel field if non-nil, zero value otherwise.

### GetLastFuelLevelOk

`func (o *VehicleSummary) GetLastFuelLevelOk() (*float32, bool)`

GetLastFuelLevelOk returns a tuple with the LastFuelLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFuelLevel

`func (o *VehicleSummary) SetLastFuelLevel(v float32)`

SetLastFuelLevel sets LastFuelLevel field to given value.

### HasLastFuelLevel

`func (o *VehicleSummary) HasLastFuelLevel() bool`

HasLastFuelLevel returns a boolean if a field has been set.

### SetLastFuelLevelNil

`func (o *VehicleSummary) SetLastFuelLevelNil(b bool)`

 SetLastFuelLevelNil sets the value for LastFuelLevel to be an explicit nil

### UnsetLastFuelLevel
`func (o *VehicleSummary) UnsetLastFuelLevel()`

UnsetLastFuelLevel ensures that no value is present for LastFuelLevel, not even an explicit nil
### GetLastOdometerReading

`func (o *VehicleSummary) GetLastOdometerReading() float32`

GetLastOdometerReading returns the LastOdometerReading field if non-nil, zero value otherwise.

### GetLastOdometerReadingOk

`func (o *VehicleSummary) GetLastOdometerReadingOk() (*float32, bool)`

GetLastOdometerReadingOk returns a tuple with the LastOdometerReading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOdometerReading

`func (o *VehicleSummary) SetLastOdometerReading(v float32)`

SetLastOdometerReading sets LastOdometerReading field to given value.

### HasLastOdometerReading

`func (o *VehicleSummary) HasLastOdometerReading() bool`

HasLastOdometerReading returns a boolean if a field has been set.

### SetLastOdometerReadingNil

`func (o *VehicleSummary) SetLastOdometerReadingNil(b bool)`

 SetLastOdometerReadingNil sets the value for LastOdometerReading to be an explicit nil

### UnsetLastOdometerReading
`func (o *VehicleSummary) UnsetLastOdometerReading()`

UnsetLastOdometerReading ensures that no value is present for LastOdometerReading, not even an explicit nil
### GetLastEngineHours

`func (o *VehicleSummary) GetLastEngineHours() float32`

GetLastEngineHours returns the LastEngineHours field if non-nil, zero value otherwise.

### GetLastEngineHoursOk

`func (o *VehicleSummary) GetLastEngineHoursOk() (*float32, bool)`

GetLastEngineHoursOk returns a tuple with the LastEngineHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEngineHours

`func (o *VehicleSummary) SetLastEngineHours(v float32)`

SetLastEngineHours sets LastEngineHours field to given value.

### HasLastEngineHours

`func (o *VehicleSummary) HasLastEngineHours() bool`

HasLastEngineHours returns a boolean if a field has been set.

### SetLastEngineHoursNil

`func (o *VehicleSummary) SetLastEngineHoursNil(b bool)`

 SetLastEngineHoursNil sets the value for LastEngineHours to be an explicit nil

### UnsetLastEngineHours
`func (o *VehicleSummary) UnsetLastEngineHours()`

UnsetLastEngineHours ensures that no value is present for LastEngineHours, not even an explicit nil
### GetLastSpeedReading

`func (o *VehicleSummary) GetLastSpeedReading() float32`

GetLastSpeedReading returns the LastSpeedReading field if non-nil, zero value otherwise.

### GetLastSpeedReadingOk

`func (o *VehicleSummary) GetLastSpeedReadingOk() (*float32, bool)`

GetLastSpeedReadingOk returns a tuple with the LastSpeedReading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSpeedReading

`func (o *VehicleSummary) SetLastSpeedReading(v float32)`

SetLastSpeedReading sets LastSpeedReading field to given value.

### HasLastSpeedReading

`func (o *VehicleSummary) HasLastSpeedReading() bool`

HasLastSpeedReading returns a boolean if a field has been set.

### SetLastSpeedReadingNil

`func (o *VehicleSummary) SetLastSpeedReadingNil(b bool)`

 SetLastSpeedReadingNil sets the value for LastSpeedReading to be an explicit nil

### UnsetLastSpeedReading
`func (o *VehicleSummary) UnsetLastSpeedReading()`

UnsetLastSpeedReading ensures that no value is present for LastSpeedReading, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


