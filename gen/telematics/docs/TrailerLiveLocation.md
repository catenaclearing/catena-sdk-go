# TrailerLiveLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrailerId** | **string** | Unique Catena identifier for the trailer. | 
**SourceName** | Pointer to **NullableString** |  | [optional] 
**TrailerName** | Pointer to **NullableString** |  | [optional] 
**VehicleId** | Pointer to **NullableString** |  | [optional] 
**DriverId** | Pointer to **NullableString** |  | [optional] 
**DriverName** | Pointer to **NullableString** |  | [optional] 
**TrailerH3Index11** | Pointer to **NullableInt32** |  | [optional] 
**TrailerLocationOccurredAt** | **time.Time** | Timestamp (UTC) when this telemetry data was recorded by the trailer. | 
**TrailerLocation** | Pointer to [**NullableTrailerLocation**](TrailerLocation.md) |  | [optional] 
**VehicleH3Index11** | Pointer to **NullableInt32** |  | [optional] 
**VehicleLocationOccurredAt** | Pointer to **NullableTime** |  | [optional] 
**VehicleLocation** | Pointer to [**NullableVehicleLocation**](VehicleLocation.md) |  | [optional] 

## Methods

### NewTrailerLiveLocation

`func NewTrailerLiveLocation(trailerId string, trailerLocationOccurredAt time.Time, ) *TrailerLiveLocation`

NewTrailerLiveLocation instantiates a new TrailerLiveLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailerLiveLocationWithDefaults

`func NewTrailerLiveLocationWithDefaults() *TrailerLiveLocation`

NewTrailerLiveLocationWithDefaults instantiates a new TrailerLiveLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrailerId

`func (o *TrailerLiveLocation) GetTrailerId() string`

GetTrailerId returns the TrailerId field if non-nil, zero value otherwise.

### GetTrailerIdOk

`func (o *TrailerLiveLocation) GetTrailerIdOk() (*string, bool)`

GetTrailerIdOk returns a tuple with the TrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerId

`func (o *TrailerLiveLocation) SetTrailerId(v string)`

SetTrailerId sets TrailerId field to given value.


### GetSourceName

`func (o *TrailerLiveLocation) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *TrailerLiveLocation) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *TrailerLiveLocation) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *TrailerLiveLocation) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *TrailerLiveLocation) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *TrailerLiveLocation) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetTrailerName

`func (o *TrailerLiveLocation) GetTrailerName() string`

GetTrailerName returns the TrailerName field if non-nil, zero value otherwise.

### GetTrailerNameOk

`func (o *TrailerLiveLocation) GetTrailerNameOk() (*string, bool)`

GetTrailerNameOk returns a tuple with the TrailerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerName

`func (o *TrailerLiveLocation) SetTrailerName(v string)`

SetTrailerName sets TrailerName field to given value.

### HasTrailerName

`func (o *TrailerLiveLocation) HasTrailerName() bool`

HasTrailerName returns a boolean if a field has been set.

### SetTrailerNameNil

`func (o *TrailerLiveLocation) SetTrailerNameNil(b bool)`

 SetTrailerNameNil sets the value for TrailerName to be an explicit nil

### UnsetTrailerName
`func (o *TrailerLiveLocation) UnsetTrailerName()`

UnsetTrailerName ensures that no value is present for TrailerName, not even an explicit nil
### GetVehicleId

`func (o *TrailerLiveLocation) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *TrailerLiveLocation) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *TrailerLiveLocation) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *TrailerLiveLocation) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *TrailerLiveLocation) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *TrailerLiveLocation) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetDriverId

`func (o *TrailerLiveLocation) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *TrailerLiveLocation) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *TrailerLiveLocation) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *TrailerLiveLocation) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *TrailerLiveLocation) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *TrailerLiveLocation) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetDriverName

`func (o *TrailerLiveLocation) GetDriverName() string`

GetDriverName returns the DriverName field if non-nil, zero value otherwise.

### GetDriverNameOk

`func (o *TrailerLiveLocation) GetDriverNameOk() (*string, bool)`

GetDriverNameOk returns a tuple with the DriverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverName

`func (o *TrailerLiveLocation) SetDriverName(v string)`

SetDriverName sets DriverName field to given value.

### HasDriverName

`func (o *TrailerLiveLocation) HasDriverName() bool`

HasDriverName returns a boolean if a field has been set.

### SetDriverNameNil

`func (o *TrailerLiveLocation) SetDriverNameNil(b bool)`

 SetDriverNameNil sets the value for DriverName to be an explicit nil

### UnsetDriverName
`func (o *TrailerLiveLocation) UnsetDriverName()`

UnsetDriverName ensures that no value is present for DriverName, not even an explicit nil
### GetTrailerH3Index11

`func (o *TrailerLiveLocation) GetTrailerH3Index11() int32`

GetTrailerH3Index11 returns the TrailerH3Index11 field if non-nil, zero value otherwise.

### GetTrailerH3Index11Ok

`func (o *TrailerLiveLocation) GetTrailerH3Index11Ok() (*int32, bool)`

GetTrailerH3Index11Ok returns a tuple with the TrailerH3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerH3Index11

`func (o *TrailerLiveLocation) SetTrailerH3Index11(v int32)`

SetTrailerH3Index11 sets TrailerH3Index11 field to given value.

### HasTrailerH3Index11

`func (o *TrailerLiveLocation) HasTrailerH3Index11() bool`

HasTrailerH3Index11 returns a boolean if a field has been set.

### SetTrailerH3Index11Nil

`func (o *TrailerLiveLocation) SetTrailerH3Index11Nil(b bool)`

 SetTrailerH3Index11Nil sets the value for TrailerH3Index11 to be an explicit nil

### UnsetTrailerH3Index11
`func (o *TrailerLiveLocation) UnsetTrailerH3Index11()`

UnsetTrailerH3Index11 ensures that no value is present for TrailerH3Index11, not even an explicit nil
### GetTrailerLocationOccurredAt

`func (o *TrailerLiveLocation) GetTrailerLocationOccurredAt() time.Time`

GetTrailerLocationOccurredAt returns the TrailerLocationOccurredAt field if non-nil, zero value otherwise.

### GetTrailerLocationOccurredAtOk

`func (o *TrailerLiveLocation) GetTrailerLocationOccurredAtOk() (*time.Time, bool)`

GetTrailerLocationOccurredAtOk returns a tuple with the TrailerLocationOccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerLocationOccurredAt

`func (o *TrailerLiveLocation) SetTrailerLocationOccurredAt(v time.Time)`

SetTrailerLocationOccurredAt sets TrailerLocationOccurredAt field to given value.


### GetTrailerLocation

`func (o *TrailerLiveLocation) GetTrailerLocation() TrailerLocation`

GetTrailerLocation returns the TrailerLocation field if non-nil, zero value otherwise.

### GetTrailerLocationOk

`func (o *TrailerLiveLocation) GetTrailerLocationOk() (*TrailerLocation, bool)`

GetTrailerLocationOk returns a tuple with the TrailerLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerLocation

`func (o *TrailerLiveLocation) SetTrailerLocation(v TrailerLocation)`

SetTrailerLocation sets TrailerLocation field to given value.

### HasTrailerLocation

`func (o *TrailerLiveLocation) HasTrailerLocation() bool`

HasTrailerLocation returns a boolean if a field has been set.

### SetTrailerLocationNil

`func (o *TrailerLiveLocation) SetTrailerLocationNil(b bool)`

 SetTrailerLocationNil sets the value for TrailerLocation to be an explicit nil

### UnsetTrailerLocation
`func (o *TrailerLiveLocation) UnsetTrailerLocation()`

UnsetTrailerLocation ensures that no value is present for TrailerLocation, not even an explicit nil
### GetVehicleH3Index11

`func (o *TrailerLiveLocation) GetVehicleH3Index11() int32`

GetVehicleH3Index11 returns the VehicleH3Index11 field if non-nil, zero value otherwise.

### GetVehicleH3Index11Ok

`func (o *TrailerLiveLocation) GetVehicleH3Index11Ok() (*int32, bool)`

GetVehicleH3Index11Ok returns a tuple with the VehicleH3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleH3Index11

`func (o *TrailerLiveLocation) SetVehicleH3Index11(v int32)`

SetVehicleH3Index11 sets VehicleH3Index11 field to given value.

### HasVehicleH3Index11

`func (o *TrailerLiveLocation) HasVehicleH3Index11() bool`

HasVehicleH3Index11 returns a boolean if a field has been set.

### SetVehicleH3Index11Nil

`func (o *TrailerLiveLocation) SetVehicleH3Index11Nil(b bool)`

 SetVehicleH3Index11Nil sets the value for VehicleH3Index11 to be an explicit nil

### UnsetVehicleH3Index11
`func (o *TrailerLiveLocation) UnsetVehicleH3Index11()`

UnsetVehicleH3Index11 ensures that no value is present for VehicleH3Index11, not even an explicit nil
### GetVehicleLocationOccurredAt

`func (o *TrailerLiveLocation) GetVehicleLocationOccurredAt() time.Time`

GetVehicleLocationOccurredAt returns the VehicleLocationOccurredAt field if non-nil, zero value otherwise.

### GetVehicleLocationOccurredAtOk

`func (o *TrailerLiveLocation) GetVehicleLocationOccurredAtOk() (*time.Time, bool)`

GetVehicleLocationOccurredAtOk returns a tuple with the VehicleLocationOccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleLocationOccurredAt

`func (o *TrailerLiveLocation) SetVehicleLocationOccurredAt(v time.Time)`

SetVehicleLocationOccurredAt sets VehicleLocationOccurredAt field to given value.

### HasVehicleLocationOccurredAt

`func (o *TrailerLiveLocation) HasVehicleLocationOccurredAt() bool`

HasVehicleLocationOccurredAt returns a boolean if a field has been set.

### SetVehicleLocationOccurredAtNil

`func (o *TrailerLiveLocation) SetVehicleLocationOccurredAtNil(b bool)`

 SetVehicleLocationOccurredAtNil sets the value for VehicleLocationOccurredAt to be an explicit nil

### UnsetVehicleLocationOccurredAt
`func (o *TrailerLiveLocation) UnsetVehicleLocationOccurredAt()`

UnsetVehicleLocationOccurredAt ensures that no value is present for VehicleLocationOccurredAt, not even an explicit nil
### GetVehicleLocation

`func (o *TrailerLiveLocation) GetVehicleLocation() VehicleLocation`

GetVehicleLocation returns the VehicleLocation field if non-nil, zero value otherwise.

### GetVehicleLocationOk

`func (o *TrailerLiveLocation) GetVehicleLocationOk() (*VehicleLocation, bool)`

GetVehicleLocationOk returns a tuple with the VehicleLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleLocation

`func (o *TrailerLiveLocation) SetVehicleLocation(v VehicleLocation)`

SetVehicleLocation sets VehicleLocation field to given value.

### HasVehicleLocation

`func (o *TrailerLiveLocation) HasVehicleLocation() bool`

HasVehicleLocation returns a boolean if a field has been set.

### SetVehicleLocationNil

`func (o *TrailerLiveLocation) SetVehicleLocationNil(b bool)`

 SetVehicleLocationNil sets the value for VehicleLocation to be an explicit nil

### UnsetVehicleLocation
`func (o *TrailerLiveLocation) UnsetVehicleLocation()`

UnsetVehicleLocation ensures that no value is present for VehicleLocation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


