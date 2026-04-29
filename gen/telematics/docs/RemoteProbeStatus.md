# RemoteProbeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProbeId** | **string** |  | 
**Temperature** | Pointer to **NullableFloat32** |  | [optional] 
**HumidityPercent** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewRemoteProbeStatus

`func NewRemoteProbeStatus(probeId string, ) *RemoteProbeStatus`

NewRemoteProbeStatus instantiates a new RemoteProbeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteProbeStatusWithDefaults

`func NewRemoteProbeStatusWithDefaults() *RemoteProbeStatus`

NewRemoteProbeStatusWithDefaults instantiates a new RemoteProbeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProbeId

`func (o *RemoteProbeStatus) GetProbeId() string`

GetProbeId returns the ProbeId field if non-nil, zero value otherwise.

### GetProbeIdOk

`func (o *RemoteProbeStatus) GetProbeIdOk() (*string, bool)`

GetProbeIdOk returns a tuple with the ProbeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbeId

`func (o *RemoteProbeStatus) SetProbeId(v string)`

SetProbeId sets ProbeId field to given value.


### GetTemperature

`func (o *RemoteProbeStatus) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *RemoteProbeStatus) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *RemoteProbeStatus) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *RemoteProbeStatus) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### SetTemperatureNil

`func (o *RemoteProbeStatus) SetTemperatureNil(b bool)`

 SetTemperatureNil sets the value for Temperature to be an explicit nil

### UnsetTemperature
`func (o *RemoteProbeStatus) UnsetTemperature()`

UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
### GetHumidityPercent

`func (o *RemoteProbeStatus) GetHumidityPercent() float32`

GetHumidityPercent returns the HumidityPercent field if non-nil, zero value otherwise.

### GetHumidityPercentOk

`func (o *RemoteProbeStatus) GetHumidityPercentOk() (*float32, bool)`

GetHumidityPercentOk returns a tuple with the HumidityPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHumidityPercent

`func (o *RemoteProbeStatus) SetHumidityPercent(v float32)`

SetHumidityPercent sets HumidityPercent field to given value.

### HasHumidityPercent

`func (o *RemoteProbeStatus) HasHumidityPercent() bool`

HasHumidityPercent returns a boolean if a field has been set.

### SetHumidityPercentNil

`func (o *RemoteProbeStatus) SetHumidityPercentNil(b bool)`

 SetHumidityPercentNil sets the value for HumidityPercent to be an explicit nil

### UnsetHumidityPercent
`func (o *RemoteProbeStatus) UnsetHumidityPercent()`

UnsetHumidityPercent ensures that no value is present for HumidityPercent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


