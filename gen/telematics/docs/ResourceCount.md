# ResourceCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fleets** | **int32** | Current count of active fleets. | 
**Vehicles** | **int32** | Current count of active vehicles. | 
**Drivers** | **int32** | Current count of active drivers. | 
**Trailers** | **int32** | Current count of active trailers. | 

## Methods

### NewResourceCount

`func NewResourceCount(fleets int32, vehicles int32, drivers int32, trailers int32, ) *ResourceCount`

NewResourceCount instantiates a new ResourceCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceCountWithDefaults

`func NewResourceCountWithDefaults() *ResourceCount`

NewResourceCountWithDefaults instantiates a new ResourceCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleets

`func (o *ResourceCount) GetFleets() int32`

GetFleets returns the Fleets field if non-nil, zero value otherwise.

### GetFleetsOk

`func (o *ResourceCount) GetFleetsOk() (*int32, bool)`

GetFleetsOk returns a tuple with the Fleets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleets

`func (o *ResourceCount) SetFleets(v int32)`

SetFleets sets Fleets field to given value.


### GetVehicles

`func (o *ResourceCount) GetVehicles() int32`

GetVehicles returns the Vehicles field if non-nil, zero value otherwise.

### GetVehiclesOk

`func (o *ResourceCount) GetVehiclesOk() (*int32, bool)`

GetVehiclesOk returns a tuple with the Vehicles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicles

`func (o *ResourceCount) SetVehicles(v int32)`

SetVehicles sets Vehicles field to given value.


### GetDrivers

`func (o *ResourceCount) GetDrivers() int32`

GetDrivers returns the Drivers field if non-nil, zero value otherwise.

### GetDriversOk

`func (o *ResourceCount) GetDriversOk() (*int32, bool)`

GetDriversOk returns a tuple with the Drivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivers

`func (o *ResourceCount) SetDrivers(v int32)`

SetDrivers sets Drivers field to given value.


### GetTrailers

`func (o *ResourceCount) GetTrailers() int32`

GetTrailers returns the Trailers field if non-nil, zero value otherwise.

### GetTrailersOk

`func (o *ResourceCount) GetTrailersOk() (*int32, bool)`

GetTrailersOk returns a tuple with the Trailers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailers

`func (o *ResourceCount) SetTrailers(v int32)`

SetTrailers sets Trailers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


