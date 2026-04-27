# TrailerZoneStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Zone** | **int32** | 1-based index of the cooling zone. | 
**SetpointTemperature** | Pointer to **NullableFloat32** |  | [optional] 
**ReturnAirTemperature** | Pointer to **NullableFloat32** |  | [optional] 
**DischargeAirTemperature** | Pointer to **NullableFloat32** |  | [optional] 
**ControlSensor** | Pointer to [**NullableReeferControlSensorEnum**](ReeferControlSensorEnum.md) |  | [optional] 
**DoorOpen** | Pointer to **NullableBool** |  | [optional] 
**ThermalStatus** | Pointer to **NullableString** |  | [optional] 
**CargoStatus** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTrailerZoneStatus

`func NewTrailerZoneStatus(zone int32, ) *TrailerZoneStatus`

NewTrailerZoneStatus instantiates a new TrailerZoneStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailerZoneStatusWithDefaults

`func NewTrailerZoneStatusWithDefaults() *TrailerZoneStatus`

NewTrailerZoneStatusWithDefaults instantiates a new TrailerZoneStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZone

`func (o *TrailerZoneStatus) GetZone() int32`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *TrailerZoneStatus) GetZoneOk() (*int32, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *TrailerZoneStatus) SetZone(v int32)`

SetZone sets Zone field to given value.


### GetSetpointTemperature

`func (o *TrailerZoneStatus) GetSetpointTemperature() float32`

GetSetpointTemperature returns the SetpointTemperature field if non-nil, zero value otherwise.

### GetSetpointTemperatureOk

`func (o *TrailerZoneStatus) GetSetpointTemperatureOk() (*float32, bool)`

GetSetpointTemperatureOk returns a tuple with the SetpointTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetpointTemperature

`func (o *TrailerZoneStatus) SetSetpointTemperature(v float32)`

SetSetpointTemperature sets SetpointTemperature field to given value.

### HasSetpointTemperature

`func (o *TrailerZoneStatus) HasSetpointTemperature() bool`

HasSetpointTemperature returns a boolean if a field has been set.

### SetSetpointTemperatureNil

`func (o *TrailerZoneStatus) SetSetpointTemperatureNil(b bool)`

 SetSetpointTemperatureNil sets the value for SetpointTemperature to be an explicit nil

### UnsetSetpointTemperature
`func (o *TrailerZoneStatus) UnsetSetpointTemperature()`

UnsetSetpointTemperature ensures that no value is present for SetpointTemperature, not even an explicit nil
### GetReturnAirTemperature

`func (o *TrailerZoneStatus) GetReturnAirTemperature() float32`

GetReturnAirTemperature returns the ReturnAirTemperature field if non-nil, zero value otherwise.

### GetReturnAirTemperatureOk

`func (o *TrailerZoneStatus) GetReturnAirTemperatureOk() (*float32, bool)`

GetReturnAirTemperatureOk returns a tuple with the ReturnAirTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnAirTemperature

`func (o *TrailerZoneStatus) SetReturnAirTemperature(v float32)`

SetReturnAirTemperature sets ReturnAirTemperature field to given value.

### HasReturnAirTemperature

`func (o *TrailerZoneStatus) HasReturnAirTemperature() bool`

HasReturnAirTemperature returns a boolean if a field has been set.

### SetReturnAirTemperatureNil

`func (o *TrailerZoneStatus) SetReturnAirTemperatureNil(b bool)`

 SetReturnAirTemperatureNil sets the value for ReturnAirTemperature to be an explicit nil

### UnsetReturnAirTemperature
`func (o *TrailerZoneStatus) UnsetReturnAirTemperature()`

UnsetReturnAirTemperature ensures that no value is present for ReturnAirTemperature, not even an explicit nil
### GetDischargeAirTemperature

`func (o *TrailerZoneStatus) GetDischargeAirTemperature() float32`

GetDischargeAirTemperature returns the DischargeAirTemperature field if non-nil, zero value otherwise.

### GetDischargeAirTemperatureOk

`func (o *TrailerZoneStatus) GetDischargeAirTemperatureOk() (*float32, bool)`

GetDischargeAirTemperatureOk returns a tuple with the DischargeAirTemperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDischargeAirTemperature

`func (o *TrailerZoneStatus) SetDischargeAirTemperature(v float32)`

SetDischargeAirTemperature sets DischargeAirTemperature field to given value.

### HasDischargeAirTemperature

`func (o *TrailerZoneStatus) HasDischargeAirTemperature() bool`

HasDischargeAirTemperature returns a boolean if a field has been set.

### SetDischargeAirTemperatureNil

`func (o *TrailerZoneStatus) SetDischargeAirTemperatureNil(b bool)`

 SetDischargeAirTemperatureNil sets the value for DischargeAirTemperature to be an explicit nil

### UnsetDischargeAirTemperature
`func (o *TrailerZoneStatus) UnsetDischargeAirTemperature()`

UnsetDischargeAirTemperature ensures that no value is present for DischargeAirTemperature, not even an explicit nil
### GetControlSensor

`func (o *TrailerZoneStatus) GetControlSensor() ReeferControlSensorEnum`

GetControlSensor returns the ControlSensor field if non-nil, zero value otherwise.

### GetControlSensorOk

`func (o *TrailerZoneStatus) GetControlSensorOk() (*ReeferControlSensorEnum, bool)`

GetControlSensorOk returns a tuple with the ControlSensor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlSensor

`func (o *TrailerZoneStatus) SetControlSensor(v ReeferControlSensorEnum)`

SetControlSensor sets ControlSensor field to given value.

### HasControlSensor

`func (o *TrailerZoneStatus) HasControlSensor() bool`

HasControlSensor returns a boolean if a field has been set.

### SetControlSensorNil

`func (o *TrailerZoneStatus) SetControlSensorNil(b bool)`

 SetControlSensorNil sets the value for ControlSensor to be an explicit nil

### UnsetControlSensor
`func (o *TrailerZoneStatus) UnsetControlSensor()`

UnsetControlSensor ensures that no value is present for ControlSensor, not even an explicit nil
### GetDoorOpen

`func (o *TrailerZoneStatus) GetDoorOpen() bool`

GetDoorOpen returns the DoorOpen field if non-nil, zero value otherwise.

### GetDoorOpenOk

`func (o *TrailerZoneStatus) GetDoorOpenOk() (*bool, bool)`

GetDoorOpenOk returns a tuple with the DoorOpen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoorOpen

`func (o *TrailerZoneStatus) SetDoorOpen(v bool)`

SetDoorOpen sets DoorOpen field to given value.

### HasDoorOpen

`func (o *TrailerZoneStatus) HasDoorOpen() bool`

HasDoorOpen returns a boolean if a field has been set.

### SetDoorOpenNil

`func (o *TrailerZoneStatus) SetDoorOpenNil(b bool)`

 SetDoorOpenNil sets the value for DoorOpen to be an explicit nil

### UnsetDoorOpen
`func (o *TrailerZoneStatus) UnsetDoorOpen()`

UnsetDoorOpen ensures that no value is present for DoorOpen, not even an explicit nil
### GetThermalStatus

`func (o *TrailerZoneStatus) GetThermalStatus() string`

GetThermalStatus returns the ThermalStatus field if non-nil, zero value otherwise.

### GetThermalStatusOk

`func (o *TrailerZoneStatus) GetThermalStatusOk() (*string, bool)`

GetThermalStatusOk returns a tuple with the ThermalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThermalStatus

`func (o *TrailerZoneStatus) SetThermalStatus(v string)`

SetThermalStatus sets ThermalStatus field to given value.

### HasThermalStatus

`func (o *TrailerZoneStatus) HasThermalStatus() bool`

HasThermalStatus returns a boolean if a field has been set.

### SetThermalStatusNil

`func (o *TrailerZoneStatus) SetThermalStatusNil(b bool)`

 SetThermalStatusNil sets the value for ThermalStatus to be an explicit nil

### UnsetThermalStatus
`func (o *TrailerZoneStatus) UnsetThermalStatus()`

UnsetThermalStatus ensures that no value is present for ThermalStatus, not even an explicit nil
### GetCargoStatus

`func (o *TrailerZoneStatus) GetCargoStatus() string`

GetCargoStatus returns the CargoStatus field if non-nil, zero value otherwise.

### GetCargoStatusOk

`func (o *TrailerZoneStatus) GetCargoStatusOk() (*string, bool)`

GetCargoStatusOk returns a tuple with the CargoStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCargoStatus

`func (o *TrailerZoneStatus) SetCargoStatus(v string)`

SetCargoStatus sets CargoStatus field to given value.

### HasCargoStatus

`func (o *TrailerZoneStatus) HasCargoStatus() bool`

HasCargoStatus returns a boolean if a field has been set.

### SetCargoStatusNil

`func (o *TrailerZoneStatus) SetCargoStatusNil(b bool)`

 SetCargoStatusNil sets the value for CargoStatus to be an explicit nil

### UnsetCargoStatus
`func (o *TrailerZoneStatus) UnsetCargoStatus()`

UnsetCargoStatus ensures that no value is present for CargoStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


