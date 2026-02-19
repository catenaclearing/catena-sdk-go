# FleetSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FleetRef** | **NullableString** |  | 
**FleetIds** | Pointer to **[]string** | List of Catena fleet IDs associated with this reference (multiple if the same fleet reconnected). | [optional] 
**Connections** | Pointer to **int32** | Number of active data connections for this fleet. | [optional] [default to 0]
**Drivers** | Pointer to **int32** | Total number of drivers associated with this fleet. | [optional] [default to 0]
**Vehicles** | Pointer to **int32** | Total number of vehicles associated with this fleet. | [optional] [default to 0]
**VehiclesWithLocations** | Pointer to **int32** | Number of vehicles that have reported at least one location update (indicates active telematics). | [optional] [default to 0]

## Methods

### NewFleetSummary

`func NewFleetSummary(fleetRef NullableString, ) *FleetSummary`

NewFleetSummary instantiates a new FleetSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetSummaryWithDefaults

`func NewFleetSummaryWithDefaults() *FleetSummary`

NewFleetSummaryWithDefaults instantiates a new FleetSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetRef

`func (o *FleetSummary) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *FleetSummary) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *FleetSummary) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### SetFleetRefNil

`func (o *FleetSummary) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *FleetSummary) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetFleetIds

`func (o *FleetSummary) GetFleetIds() []string`

GetFleetIds returns the FleetIds field if non-nil, zero value otherwise.

### GetFleetIdsOk

`func (o *FleetSummary) GetFleetIdsOk() (*[]string, bool)`

GetFleetIdsOk returns a tuple with the FleetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetIds

`func (o *FleetSummary) SetFleetIds(v []string)`

SetFleetIds sets FleetIds field to given value.

### HasFleetIds

`func (o *FleetSummary) HasFleetIds() bool`

HasFleetIds returns a boolean if a field has been set.

### GetConnections

`func (o *FleetSummary) GetConnections() int32`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *FleetSummary) GetConnectionsOk() (*int32, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *FleetSummary) SetConnections(v int32)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *FleetSummary) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetDrivers

`func (o *FleetSummary) GetDrivers() int32`

GetDrivers returns the Drivers field if non-nil, zero value otherwise.

### GetDriversOk

`func (o *FleetSummary) GetDriversOk() (*int32, bool)`

GetDriversOk returns a tuple with the Drivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivers

`func (o *FleetSummary) SetDrivers(v int32)`

SetDrivers sets Drivers field to given value.

### HasDrivers

`func (o *FleetSummary) HasDrivers() bool`

HasDrivers returns a boolean if a field has been set.

### GetVehicles

`func (o *FleetSummary) GetVehicles() int32`

GetVehicles returns the Vehicles field if non-nil, zero value otherwise.

### GetVehiclesOk

`func (o *FleetSummary) GetVehiclesOk() (*int32, bool)`

GetVehiclesOk returns a tuple with the Vehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicles

`func (o *FleetSummary) SetVehicles(v int32)`

SetVehicles sets Vehicles field to given value.

### HasVehicles

`func (o *FleetSummary) HasVehicles() bool`

HasVehicles returns a boolean if a field has been set.

### GetVehiclesWithLocations

`func (o *FleetSummary) GetVehiclesWithLocations() int32`

GetVehiclesWithLocations returns the VehiclesWithLocations field if non-nil, zero value otherwise.

### GetVehiclesWithLocationsOk

`func (o *FleetSummary) GetVehiclesWithLocationsOk() (*int32, bool)`

GetVehiclesWithLocationsOk returns a tuple with the VehiclesWithLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehiclesWithLocations

`func (o *FleetSummary) SetVehiclesWithLocations(v int32)`

SetVehiclesWithLocations sets VehiclesWithLocations field to given value.

### HasVehiclesWithLocations

`func (o *FleetSummary) HasVehiclesWithLocations() bool`

HasVehiclesWithLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


