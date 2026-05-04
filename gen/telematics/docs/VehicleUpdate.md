# VehicleUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | Unique identifier of the connection at Catena Telematics which will be used to create this resource. A connection represents a Fleet/TSP pairing. | 
**VehicleName** | Pointer to **NullableString** |  | [optional] 
**Oem** | Pointer to **NullableString** |  | [optional] 
**ModelType** | Pointer to **NullableString** |  | [optional] 
**ModelYear** | Pointer to **NullableInt32** |  | [optional] 
**Vin** | Pointer to **NullableString** |  | [optional] 
**EngineVin** | Pointer to **NullableString** |  | [optional] 
**LicensePlateRegion** | Pointer to **NullableString** |  | [optional] 
**LicensePlateCountry** | Pointer to **NullableString** |  | [optional] 
**LicensePlateNumber** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**Status** | Pointer to [**NullableVehicleStatusEnum**](VehicleStatusEnum.md) |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**EldId** | Pointer to **NullableString** |  | [optional] 
**EldSerialNumber** | Pointer to **NullableString** |  | [optional] 
**EldDeviceType** | Pointer to **NullableString** |  | [optional] 
**EldProductId** | Pointer to **NullableString** |  | [optional] 
**TotalAxles** | Pointer to **NullableInt32** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**SpeedUnit** | Pointer to **NullableString** |  | [optional] 
**OdometerUnit** | Pointer to **NullableString** |  | [optional] 
**FuelUnit** | Pointer to **NullableString** |  | [optional] 
**FuelCapacity** | Pointer to **NullableString** |  | [optional] 
**EngineType** | Pointer to [**NullableEngineType**](EngineType.md) |  | [optional] 

## Methods

### NewVehicleUpdate

`func NewVehicleUpdate(connectionId string, ) *VehicleUpdate`

NewVehicleUpdate instantiates a new VehicleUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleUpdateWithDefaults

`func NewVehicleUpdateWithDefaults() *VehicleUpdate`

NewVehicleUpdateWithDefaults instantiates a new VehicleUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *VehicleUpdate) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *VehicleUpdate) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *VehicleUpdate) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetVehicleName

`func (o *VehicleUpdate) GetVehicleName() string`

GetVehicleName returns the VehicleName field if non-nil, zero value otherwise.

### GetVehicleNameOk

`func (o *VehicleUpdate) GetVehicleNameOk() (*string, bool)`

GetVehicleNameOk returns a tuple with the VehicleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleName

`func (o *VehicleUpdate) SetVehicleName(v string)`

SetVehicleName sets VehicleName field to given value.

### HasVehicleName

`func (o *VehicleUpdate) HasVehicleName() bool`

HasVehicleName returns a boolean if a field has been set.

### SetVehicleNameNil

`func (o *VehicleUpdate) SetVehicleNameNil(b bool)`

 SetVehicleNameNil sets the value for VehicleName to be an explicit nil

### UnsetVehicleName
`func (o *VehicleUpdate) UnsetVehicleName()`

UnsetVehicleName ensures that no value is present for VehicleName, not even an explicit nil
### GetOem

`func (o *VehicleUpdate) GetOem() string`

GetOem returns the Oem field if non-nil, zero value otherwise.

### GetOemOk

`func (o *VehicleUpdate) GetOemOk() (*string, bool)`

GetOemOk returns a tuple with the Oem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOem

`func (o *VehicleUpdate) SetOem(v string)`

SetOem sets Oem field to given value.

### HasOem

`func (o *VehicleUpdate) HasOem() bool`

HasOem returns a boolean if a field has been set.

### SetOemNil

`func (o *VehicleUpdate) SetOemNil(b bool)`

 SetOemNil sets the value for Oem to be an explicit nil

### UnsetOem
`func (o *VehicleUpdate) UnsetOem()`

UnsetOem ensures that no value is present for Oem, not even an explicit nil
### GetModelType

`func (o *VehicleUpdate) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *VehicleUpdate) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *VehicleUpdate) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *VehicleUpdate) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### SetModelTypeNil

`func (o *VehicleUpdate) SetModelTypeNil(b bool)`

 SetModelTypeNil sets the value for ModelType to be an explicit nil

### UnsetModelType
`func (o *VehicleUpdate) UnsetModelType()`

UnsetModelType ensures that no value is present for ModelType, not even an explicit nil
### GetModelYear

`func (o *VehicleUpdate) GetModelYear() int32`

GetModelYear returns the ModelYear field if non-nil, zero value otherwise.

### GetModelYearOk

`func (o *VehicleUpdate) GetModelYearOk() (*int32, bool)`

GetModelYearOk returns a tuple with the ModelYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelYear

`func (o *VehicleUpdate) SetModelYear(v int32)`

SetModelYear sets ModelYear field to given value.

### HasModelYear

`func (o *VehicleUpdate) HasModelYear() bool`

HasModelYear returns a boolean if a field has been set.

### SetModelYearNil

`func (o *VehicleUpdate) SetModelYearNil(b bool)`

 SetModelYearNil sets the value for ModelYear to be an explicit nil

### UnsetModelYear
`func (o *VehicleUpdate) UnsetModelYear()`

UnsetModelYear ensures that no value is present for ModelYear, not even an explicit nil
### GetVin

`func (o *VehicleUpdate) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *VehicleUpdate) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *VehicleUpdate) SetVin(v string)`

SetVin sets Vin field to given value.

### HasVin

`func (o *VehicleUpdate) HasVin() bool`

HasVin returns a boolean if a field has been set.

### SetVinNil

`func (o *VehicleUpdate) SetVinNil(b bool)`

 SetVinNil sets the value for Vin to be an explicit nil

### UnsetVin
`func (o *VehicleUpdate) UnsetVin()`

UnsetVin ensures that no value is present for Vin, not even an explicit nil
### GetEngineVin

`func (o *VehicleUpdate) GetEngineVin() string`

GetEngineVin returns the EngineVin field if non-nil, zero value otherwise.

### GetEngineVinOk

`func (o *VehicleUpdate) GetEngineVinOk() (*string, bool)`

GetEngineVinOk returns a tuple with the EngineVin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVin

`func (o *VehicleUpdate) SetEngineVin(v string)`

SetEngineVin sets EngineVin field to given value.

### HasEngineVin

`func (o *VehicleUpdate) HasEngineVin() bool`

HasEngineVin returns a boolean if a field has been set.

### SetEngineVinNil

`func (o *VehicleUpdate) SetEngineVinNil(b bool)`

 SetEngineVinNil sets the value for EngineVin to be an explicit nil

### UnsetEngineVin
`func (o *VehicleUpdate) UnsetEngineVin()`

UnsetEngineVin ensures that no value is present for EngineVin, not even an explicit nil
### GetLicensePlateRegion

`func (o *VehicleUpdate) GetLicensePlateRegion() string`

GetLicensePlateRegion returns the LicensePlateRegion field if non-nil, zero value otherwise.

### GetLicensePlateRegionOk

`func (o *VehicleUpdate) GetLicensePlateRegionOk() (*string, bool)`

GetLicensePlateRegionOk returns a tuple with the LicensePlateRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateRegion

`func (o *VehicleUpdate) SetLicensePlateRegion(v string)`

SetLicensePlateRegion sets LicensePlateRegion field to given value.

### HasLicensePlateRegion

`func (o *VehicleUpdate) HasLicensePlateRegion() bool`

HasLicensePlateRegion returns a boolean if a field has been set.

### SetLicensePlateRegionNil

`func (o *VehicleUpdate) SetLicensePlateRegionNil(b bool)`

 SetLicensePlateRegionNil sets the value for LicensePlateRegion to be an explicit nil

### UnsetLicensePlateRegion
`func (o *VehicleUpdate) UnsetLicensePlateRegion()`

UnsetLicensePlateRegion ensures that no value is present for LicensePlateRegion, not even an explicit nil
### GetLicensePlateCountry

`func (o *VehicleUpdate) GetLicensePlateCountry() string`

GetLicensePlateCountry returns the LicensePlateCountry field if non-nil, zero value otherwise.

### GetLicensePlateCountryOk

`func (o *VehicleUpdate) GetLicensePlateCountryOk() (*string, bool)`

GetLicensePlateCountryOk returns a tuple with the LicensePlateCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateCountry

`func (o *VehicleUpdate) SetLicensePlateCountry(v string)`

SetLicensePlateCountry sets LicensePlateCountry field to given value.

### HasLicensePlateCountry

`func (o *VehicleUpdate) HasLicensePlateCountry() bool`

HasLicensePlateCountry returns a boolean if a field has been set.

### SetLicensePlateCountryNil

`func (o *VehicleUpdate) SetLicensePlateCountryNil(b bool)`

 SetLicensePlateCountryNil sets the value for LicensePlateCountry to be an explicit nil

### UnsetLicensePlateCountry
`func (o *VehicleUpdate) UnsetLicensePlateCountry()`

UnsetLicensePlateCountry ensures that no value is present for LicensePlateCountry, not even an explicit nil
### GetLicensePlateNumber

`func (o *VehicleUpdate) GetLicensePlateNumber() string`

GetLicensePlateNumber returns the LicensePlateNumber field if non-nil, zero value otherwise.

### GetLicensePlateNumberOk

`func (o *VehicleUpdate) GetLicensePlateNumberOk() (*string, bool)`

GetLicensePlateNumberOk returns a tuple with the LicensePlateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateNumber

`func (o *VehicleUpdate) SetLicensePlateNumber(v string)`

SetLicensePlateNumber sets LicensePlateNumber field to given value.

### HasLicensePlateNumber

`func (o *VehicleUpdate) HasLicensePlateNumber() bool`

HasLicensePlateNumber returns a boolean if a field has been set.

### SetLicensePlateNumberNil

`func (o *VehicleUpdate) SetLicensePlateNumberNil(b bool)`

 SetLicensePlateNumberNil sets the value for LicensePlateNumber to be an explicit nil

### UnsetLicensePlateNumber
`func (o *VehicleUpdate) UnsetLicensePlateNumber()`

UnsetLicensePlateNumber ensures that no value is present for LicensePlateNumber, not even an explicit nil
### GetIsActive

`func (o *VehicleUpdate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *VehicleUpdate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *VehicleUpdate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *VehicleUpdate) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *VehicleUpdate) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *VehicleUpdate) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetStatus

`func (o *VehicleUpdate) GetStatus() VehicleStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VehicleUpdate) GetStatusOk() (*VehicleStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VehicleUpdate) SetStatus(v VehicleStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VehicleUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *VehicleUpdate) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *VehicleUpdate) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNotes

`func (o *VehicleUpdate) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *VehicleUpdate) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *VehicleUpdate) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *VehicleUpdate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *VehicleUpdate) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *VehicleUpdate) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetEldId

`func (o *VehicleUpdate) GetEldId() string`

GetEldId returns the EldId field if non-nil, zero value otherwise.

### GetEldIdOk

`func (o *VehicleUpdate) GetEldIdOk() (*string, bool)`

GetEldIdOk returns a tuple with the EldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldId

`func (o *VehicleUpdate) SetEldId(v string)`

SetEldId sets EldId field to given value.

### HasEldId

`func (o *VehicleUpdate) HasEldId() bool`

HasEldId returns a boolean if a field has been set.

### SetEldIdNil

`func (o *VehicleUpdate) SetEldIdNil(b bool)`

 SetEldIdNil sets the value for EldId to be an explicit nil

### UnsetEldId
`func (o *VehicleUpdate) UnsetEldId()`

UnsetEldId ensures that no value is present for EldId, not even an explicit nil
### GetEldSerialNumber

`func (o *VehicleUpdate) GetEldSerialNumber() string`

GetEldSerialNumber returns the EldSerialNumber field if non-nil, zero value otherwise.

### GetEldSerialNumberOk

`func (o *VehicleUpdate) GetEldSerialNumberOk() (*string, bool)`

GetEldSerialNumberOk returns a tuple with the EldSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldSerialNumber

`func (o *VehicleUpdate) SetEldSerialNumber(v string)`

SetEldSerialNumber sets EldSerialNumber field to given value.

### HasEldSerialNumber

`func (o *VehicleUpdate) HasEldSerialNumber() bool`

HasEldSerialNumber returns a boolean if a field has been set.

### SetEldSerialNumberNil

`func (o *VehicleUpdate) SetEldSerialNumberNil(b bool)`

 SetEldSerialNumberNil sets the value for EldSerialNumber to be an explicit nil

### UnsetEldSerialNumber
`func (o *VehicleUpdate) UnsetEldSerialNumber()`

UnsetEldSerialNumber ensures that no value is present for EldSerialNumber, not even an explicit nil
### GetEldDeviceType

`func (o *VehicleUpdate) GetEldDeviceType() string`

GetEldDeviceType returns the EldDeviceType field if non-nil, zero value otherwise.

### GetEldDeviceTypeOk

`func (o *VehicleUpdate) GetEldDeviceTypeOk() (*string, bool)`

GetEldDeviceTypeOk returns a tuple with the EldDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldDeviceType

`func (o *VehicleUpdate) SetEldDeviceType(v string)`

SetEldDeviceType sets EldDeviceType field to given value.

### HasEldDeviceType

`func (o *VehicleUpdate) HasEldDeviceType() bool`

HasEldDeviceType returns a boolean if a field has been set.

### SetEldDeviceTypeNil

`func (o *VehicleUpdate) SetEldDeviceTypeNil(b bool)`

 SetEldDeviceTypeNil sets the value for EldDeviceType to be an explicit nil

### UnsetEldDeviceType
`func (o *VehicleUpdate) UnsetEldDeviceType()`

UnsetEldDeviceType ensures that no value is present for EldDeviceType, not even an explicit nil
### GetEldProductId

`func (o *VehicleUpdate) GetEldProductId() string`

GetEldProductId returns the EldProductId field if non-nil, zero value otherwise.

### GetEldProductIdOk

`func (o *VehicleUpdate) GetEldProductIdOk() (*string, bool)`

GetEldProductIdOk returns a tuple with the EldProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldProductId

`func (o *VehicleUpdate) SetEldProductId(v string)`

SetEldProductId sets EldProductId field to given value.

### HasEldProductId

`func (o *VehicleUpdate) HasEldProductId() bool`

HasEldProductId returns a boolean if a field has been set.

### SetEldProductIdNil

`func (o *VehicleUpdate) SetEldProductIdNil(b bool)`

 SetEldProductIdNil sets the value for EldProductId to be an explicit nil

### UnsetEldProductId
`func (o *VehicleUpdate) UnsetEldProductId()`

UnsetEldProductId ensures that no value is present for EldProductId, not even an explicit nil
### GetTotalAxles

`func (o *VehicleUpdate) GetTotalAxles() int32`

GetTotalAxles returns the TotalAxles field if non-nil, zero value otherwise.

### GetTotalAxlesOk

`func (o *VehicleUpdate) GetTotalAxlesOk() (*int32, bool)`

GetTotalAxlesOk returns a tuple with the TotalAxles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAxles

`func (o *VehicleUpdate) SetTotalAxles(v int32)`

SetTotalAxles sets TotalAxles field to given value.

### HasTotalAxles

`func (o *VehicleUpdate) HasTotalAxles() bool`

HasTotalAxles returns a boolean if a field has been set.

### SetTotalAxlesNil

`func (o *VehicleUpdate) SetTotalAxlesNil(b bool)`

 SetTotalAxlesNil sets the value for TotalAxles to be an explicit nil

### UnsetTotalAxles
`func (o *VehicleUpdate) UnsetTotalAxles()`

UnsetTotalAxles ensures that no value is present for TotalAxles, not even an explicit nil
### GetExternalId

`func (o *VehicleUpdate) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *VehicleUpdate) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *VehicleUpdate) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *VehicleUpdate) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *VehicleUpdate) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *VehicleUpdate) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetSpeedUnit

`func (o *VehicleUpdate) GetSpeedUnit() string`

GetSpeedUnit returns the SpeedUnit field if non-nil, zero value otherwise.

### GetSpeedUnitOk

`func (o *VehicleUpdate) GetSpeedUnitOk() (*string, bool)`

GetSpeedUnitOk returns a tuple with the SpeedUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedUnit

`func (o *VehicleUpdate) SetSpeedUnit(v string)`

SetSpeedUnit sets SpeedUnit field to given value.

### HasSpeedUnit

`func (o *VehicleUpdate) HasSpeedUnit() bool`

HasSpeedUnit returns a boolean if a field has been set.

### SetSpeedUnitNil

`func (o *VehicleUpdate) SetSpeedUnitNil(b bool)`

 SetSpeedUnitNil sets the value for SpeedUnit to be an explicit nil

### UnsetSpeedUnit
`func (o *VehicleUpdate) UnsetSpeedUnit()`

UnsetSpeedUnit ensures that no value is present for SpeedUnit, not even an explicit nil
### GetOdometerUnit

`func (o *VehicleUpdate) GetOdometerUnit() string`

GetOdometerUnit returns the OdometerUnit field if non-nil, zero value otherwise.

### GetOdometerUnitOk

`func (o *VehicleUpdate) GetOdometerUnitOk() (*string, bool)`

GetOdometerUnitOk returns a tuple with the OdometerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerUnit

`func (o *VehicleUpdate) SetOdometerUnit(v string)`

SetOdometerUnit sets OdometerUnit field to given value.

### HasOdometerUnit

`func (o *VehicleUpdate) HasOdometerUnit() bool`

HasOdometerUnit returns a boolean if a field has been set.

### SetOdometerUnitNil

`func (o *VehicleUpdate) SetOdometerUnitNil(b bool)`

 SetOdometerUnitNil sets the value for OdometerUnit to be an explicit nil

### UnsetOdometerUnit
`func (o *VehicleUpdate) UnsetOdometerUnit()`

UnsetOdometerUnit ensures that no value is present for OdometerUnit, not even an explicit nil
### GetFuelUnit

`func (o *VehicleUpdate) GetFuelUnit() string`

GetFuelUnit returns the FuelUnit field if non-nil, zero value otherwise.

### GetFuelUnitOk

`func (o *VehicleUpdate) GetFuelUnitOk() (*string, bool)`

GetFuelUnitOk returns a tuple with the FuelUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelUnit

`func (o *VehicleUpdate) SetFuelUnit(v string)`

SetFuelUnit sets FuelUnit field to given value.

### HasFuelUnit

`func (o *VehicleUpdate) HasFuelUnit() bool`

HasFuelUnit returns a boolean if a field has been set.

### SetFuelUnitNil

`func (o *VehicleUpdate) SetFuelUnitNil(b bool)`

 SetFuelUnitNil sets the value for FuelUnit to be an explicit nil

### UnsetFuelUnit
`func (o *VehicleUpdate) UnsetFuelUnit()`

UnsetFuelUnit ensures that no value is present for FuelUnit, not even an explicit nil
### GetFuelCapacity

`func (o *VehicleUpdate) GetFuelCapacity() string`

GetFuelCapacity returns the FuelCapacity field if non-nil, zero value otherwise.

### GetFuelCapacityOk

`func (o *VehicleUpdate) GetFuelCapacityOk() (*string, bool)`

GetFuelCapacityOk returns a tuple with the FuelCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelCapacity

`func (o *VehicleUpdate) SetFuelCapacity(v string)`

SetFuelCapacity sets FuelCapacity field to given value.

### HasFuelCapacity

`func (o *VehicleUpdate) HasFuelCapacity() bool`

HasFuelCapacity returns a boolean if a field has been set.

### SetFuelCapacityNil

`func (o *VehicleUpdate) SetFuelCapacityNil(b bool)`

 SetFuelCapacityNil sets the value for FuelCapacity to be an explicit nil

### UnsetFuelCapacity
`func (o *VehicleUpdate) UnsetFuelCapacity()`

UnsetFuelCapacity ensures that no value is present for FuelCapacity, not even an explicit nil
### GetEngineType

`func (o *VehicleUpdate) GetEngineType() EngineType`

GetEngineType returns the EngineType field if non-nil, zero value otherwise.

### GetEngineTypeOk

`func (o *VehicleUpdate) GetEngineTypeOk() (*EngineType, bool)`

GetEngineTypeOk returns a tuple with the EngineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineType

`func (o *VehicleUpdate) SetEngineType(v EngineType)`

SetEngineType sets EngineType field to given value.

### HasEngineType

`func (o *VehicleUpdate) HasEngineType() bool`

HasEngineType returns a boolean if a field has been set.

### SetEngineTypeNil

`func (o *VehicleUpdate) SetEngineTypeNil(b bool)`

 SetEngineTypeNil sets the value for EngineType to be an explicit nil

### UnsetEngineType
`func (o *VehicleUpdate) UnsetEngineType()`

UnsetEngineType ensures that no value is present for EngineType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


