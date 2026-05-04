# VehicleRegionSegmentRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FleetId** | **NullableString** |  | 
**FleetRef** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | Unique identifier of the record at Catena Telematics. | 
**CreatedAt** | **time.Time** | Immutable: The datetime the record was ingested into Catena Telematics. | 
**UpdatedAt** | **time.Time** | The dateime the record was last modified in Catena Telematics. | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**ConnectionId** | **string** | Unique identifier of the connection at Catena Telematics through which this record was ingested. A connection represents a Fleet/TSP pairing. | 
**ProcessingDate** | **string** | UTC day for which the region segments are listed, bounded from 00:00 to 00:00 UTC. | 
**VehicleId** | **string** | Internal vehicle identifier (Catena FK). | 
**DriverId** | Pointer to **NullableString** |  | [optional] 
**CountryCode** | Pointer to **NullableString** |  | [optional] 
**Region** | Pointer to **NullableString** |  | [optional] 
**SegmentStartTime** | **time.Time** | UTC timestamp of the first location sample in the segment. | 
**SegmentEndTime** | **time.Time** | UTC timestamp of the last location sample in the segment. | 
**OdometerStart** | Pointer to **NullableFloat32** |  | [optional] 
**OdometerEnd** | Pointer to **NullableFloat32** |  | [optional] 
**Distance** | Pointer to **NullableFloat32** |  | [optional] 
**LocationCount** | **int32** | Number of location samples that belong to the segment. | 
**OdometerStartUnit** | Pointer to [**DistanceUnit**](DistanceUnit.md) | Unit for odometer_start. | [optional] 
**OdometerEndUnit** | Pointer to [**DistanceUnit**](DistanceUnit.md) | Unit for odometer_end. | [optional] 
**DistanceUnit** | Pointer to [**DistanceUnit**](DistanceUnit.md) | Unit for distance. | [optional] 

## Methods

### NewVehicleRegionSegmentRead

`func NewVehicleRegionSegmentRead(fleetId NullableString, id string, createdAt time.Time, updatedAt time.Time, connectionId string, processingDate string, vehicleId string, segmentStartTime time.Time, segmentEndTime time.Time, locationCount int32, ) *VehicleRegionSegmentRead`

NewVehicleRegionSegmentRead instantiates a new VehicleRegionSegmentRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVehicleRegionSegmentReadWithDefaults

`func NewVehicleRegionSegmentReadWithDefaults() *VehicleRegionSegmentRead`

NewVehicleRegionSegmentReadWithDefaults instantiates a new VehicleRegionSegmentRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *VehicleRegionSegmentRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *VehicleRegionSegmentRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *VehicleRegionSegmentRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### SetFleetIdNil

`func (o *VehicleRegionSegmentRead) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *VehicleRegionSegmentRead) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetFleetRef

`func (o *VehicleRegionSegmentRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *VehicleRegionSegmentRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *VehicleRegionSegmentRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *VehicleRegionSegmentRead) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *VehicleRegionSegmentRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *VehicleRegionSegmentRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetId

`func (o *VehicleRegionSegmentRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VehicleRegionSegmentRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VehicleRegionSegmentRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *VehicleRegionSegmentRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VehicleRegionSegmentRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VehicleRegionSegmentRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *VehicleRegionSegmentRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VehicleRegionSegmentRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VehicleRegionSegmentRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *VehicleRegionSegmentRead) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *VehicleRegionSegmentRead) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *VehicleRegionSegmentRead) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *VehicleRegionSegmentRead) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *VehicleRegionSegmentRead) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *VehicleRegionSegmentRead) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetConnectionId

`func (o *VehicleRegionSegmentRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *VehicleRegionSegmentRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *VehicleRegionSegmentRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetProcessingDate

`func (o *VehicleRegionSegmentRead) GetProcessingDate() string`

GetProcessingDate returns the ProcessingDate field if non-nil, zero value otherwise.

### GetProcessingDateOk

`func (o *VehicleRegionSegmentRead) GetProcessingDateOk() (*string, bool)`

GetProcessingDateOk returns a tuple with the ProcessingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingDate

`func (o *VehicleRegionSegmentRead) SetProcessingDate(v string)`

SetProcessingDate sets ProcessingDate field to given value.


### GetVehicleId

`func (o *VehicleRegionSegmentRead) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *VehicleRegionSegmentRead) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *VehicleRegionSegmentRead) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.


### GetDriverId

`func (o *VehicleRegionSegmentRead) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *VehicleRegionSegmentRead) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *VehicleRegionSegmentRead) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *VehicleRegionSegmentRead) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *VehicleRegionSegmentRead) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *VehicleRegionSegmentRead) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetCountryCode

`func (o *VehicleRegionSegmentRead) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *VehicleRegionSegmentRead) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *VehicleRegionSegmentRead) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *VehicleRegionSegmentRead) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *VehicleRegionSegmentRead) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *VehicleRegionSegmentRead) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetRegion

`func (o *VehicleRegionSegmentRead) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VehicleRegionSegmentRead) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VehicleRegionSegmentRead) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *VehicleRegionSegmentRead) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *VehicleRegionSegmentRead) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *VehicleRegionSegmentRead) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetSegmentStartTime

`func (o *VehicleRegionSegmentRead) GetSegmentStartTime() time.Time`

GetSegmentStartTime returns the SegmentStartTime field if non-nil, zero value otherwise.

### GetSegmentStartTimeOk

`func (o *VehicleRegionSegmentRead) GetSegmentStartTimeOk() (*time.Time, bool)`

GetSegmentStartTimeOk returns a tuple with the SegmentStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentStartTime

`func (o *VehicleRegionSegmentRead) SetSegmentStartTime(v time.Time)`

SetSegmentStartTime sets SegmentStartTime field to given value.


### GetSegmentEndTime

`func (o *VehicleRegionSegmentRead) GetSegmentEndTime() time.Time`

GetSegmentEndTime returns the SegmentEndTime field if non-nil, zero value otherwise.

### GetSegmentEndTimeOk

`func (o *VehicleRegionSegmentRead) GetSegmentEndTimeOk() (*time.Time, bool)`

GetSegmentEndTimeOk returns a tuple with the SegmentEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegmentEndTime

`func (o *VehicleRegionSegmentRead) SetSegmentEndTime(v time.Time)`

SetSegmentEndTime sets SegmentEndTime field to given value.


### GetOdometerStart

`func (o *VehicleRegionSegmentRead) GetOdometerStart() float32`

GetOdometerStart returns the OdometerStart field if non-nil, zero value otherwise.

### GetOdometerStartOk

`func (o *VehicleRegionSegmentRead) GetOdometerStartOk() (*float32, bool)`

GetOdometerStartOk returns a tuple with the OdometerStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerStart

`func (o *VehicleRegionSegmentRead) SetOdometerStart(v float32)`

SetOdometerStart sets OdometerStart field to given value.

### HasOdometerStart

`func (o *VehicleRegionSegmentRead) HasOdometerStart() bool`

HasOdometerStart returns a boolean if a field has been set.

### SetOdometerStartNil

`func (o *VehicleRegionSegmentRead) SetOdometerStartNil(b bool)`

 SetOdometerStartNil sets the value for OdometerStart to be an explicit nil

### UnsetOdometerStart
`func (o *VehicleRegionSegmentRead) UnsetOdometerStart()`

UnsetOdometerStart ensures that no value is present for OdometerStart, not even an explicit nil
### GetOdometerEnd

`func (o *VehicleRegionSegmentRead) GetOdometerEnd() float32`

GetOdometerEnd returns the OdometerEnd field if non-nil, zero value otherwise.

### GetOdometerEndOk

`func (o *VehicleRegionSegmentRead) GetOdometerEndOk() (*float32, bool)`

GetOdometerEndOk returns a tuple with the OdometerEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerEnd

`func (o *VehicleRegionSegmentRead) SetOdometerEnd(v float32)`

SetOdometerEnd sets OdometerEnd field to given value.

### HasOdometerEnd

`func (o *VehicleRegionSegmentRead) HasOdometerEnd() bool`

HasOdometerEnd returns a boolean if a field has been set.

### SetOdometerEndNil

`func (o *VehicleRegionSegmentRead) SetOdometerEndNil(b bool)`

 SetOdometerEndNil sets the value for OdometerEnd to be an explicit nil

### UnsetOdometerEnd
`func (o *VehicleRegionSegmentRead) UnsetOdometerEnd()`

UnsetOdometerEnd ensures that no value is present for OdometerEnd, not even an explicit nil
### GetDistance

`func (o *VehicleRegionSegmentRead) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *VehicleRegionSegmentRead) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *VehicleRegionSegmentRead) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *VehicleRegionSegmentRead) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### SetDistanceNil

`func (o *VehicleRegionSegmentRead) SetDistanceNil(b bool)`

 SetDistanceNil sets the value for Distance to be an explicit nil

### UnsetDistance
`func (o *VehicleRegionSegmentRead) UnsetDistance()`

UnsetDistance ensures that no value is present for Distance, not even an explicit nil
### GetLocationCount

`func (o *VehicleRegionSegmentRead) GetLocationCount() int32`

GetLocationCount returns the LocationCount field if non-nil, zero value otherwise.

### GetLocationCountOk

`func (o *VehicleRegionSegmentRead) GetLocationCountOk() (*int32, bool)`

GetLocationCountOk returns a tuple with the LocationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationCount

`func (o *VehicleRegionSegmentRead) SetLocationCount(v int32)`

SetLocationCount sets LocationCount field to given value.


### GetOdometerStartUnit

`func (o *VehicleRegionSegmentRead) GetOdometerStartUnit() DistanceUnit`

GetOdometerStartUnit returns the OdometerStartUnit field if non-nil, zero value otherwise.

### GetOdometerStartUnitOk

`func (o *VehicleRegionSegmentRead) GetOdometerStartUnitOk() (*DistanceUnit, bool)`

GetOdometerStartUnitOk returns a tuple with the OdometerStartUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerStartUnit

`func (o *VehicleRegionSegmentRead) SetOdometerStartUnit(v DistanceUnit)`

SetOdometerStartUnit sets OdometerStartUnit field to given value.

### HasOdometerStartUnit

`func (o *VehicleRegionSegmentRead) HasOdometerStartUnit() bool`

HasOdometerStartUnit returns a boolean if a field has been set.

### GetOdometerEndUnit

`func (o *VehicleRegionSegmentRead) GetOdometerEndUnit() DistanceUnit`

GetOdometerEndUnit returns the OdometerEndUnit field if non-nil, zero value otherwise.

### GetOdometerEndUnitOk

`func (o *VehicleRegionSegmentRead) GetOdometerEndUnitOk() (*DistanceUnit, bool)`

GetOdometerEndUnitOk returns a tuple with the OdometerEndUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerEndUnit

`func (o *VehicleRegionSegmentRead) SetOdometerEndUnit(v DistanceUnit)`

SetOdometerEndUnit sets OdometerEndUnit field to given value.

### HasOdometerEndUnit

`func (o *VehicleRegionSegmentRead) HasOdometerEndUnit() bool`

HasOdometerEndUnit returns a boolean if a field has been set.

### GetDistanceUnit

`func (o *VehicleRegionSegmentRead) GetDistanceUnit() DistanceUnit`

GetDistanceUnit returns the DistanceUnit field if non-nil, zero value otherwise.

### GetDistanceUnitOk

`func (o *VehicleRegionSegmentRead) GetDistanceUnitOk() (*DistanceUnit, bool)`

GetDistanceUnitOk returns a tuple with the DistanceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceUnit

`func (o *VehicleRegionSegmentRead) SetDistanceUnit(v DistanceUnit)`

SetDistanceUnit sets DistanceUnit field to given value.

### HasDistanceUnit

`func (o *VehicleRegionSegmentRead) HasDistanceUnit() bool`

HasDistanceUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


