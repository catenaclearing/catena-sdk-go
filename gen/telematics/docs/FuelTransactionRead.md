# FuelTransactionRead

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
**TspId** | Pointer to **NullableString** |  | [optional] 
**TspSlug** | Pointer to **NullableString** |  | [optional] 
**SourceName** | [**TspEnum**](TspEnum.md) | The underlying telematics platform that provided this data (e.g., &#x60;samsara&#x60;, &#x60;motive&#x60;, &#x60;hos247&#x60;). Note: Some platforms like &#x60;hos247&#x60; offer white-labeling, so multiple TSPs may share the same source_name — use &#x60;tsp_id&#x60; or &#x60;tsp_slug&#x60; to identify the specific ELD provider. | 
**SourceData** | Pointer to **map[string]interface{}** | Raw source payload as ingested from the TSP. **Note: use it for audit/debugging.** | [optional] 
**SourceId** | **string** | Unique identifier of the record in the TSP. **Note: we generate a unique composite key based on available fields if the TSP does not provide an unique ID.** | 
**SourceDataHash** | **string** | SHA-256 hash of the source data payload. **Note: we use it internally for idempotence and deduplication.** | 
**OccurredAt** | Pointer to **NullableTime** |  | [optional] 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 
**Extras** | Pointer to **map[string]interface{}** |  | [optional] 
**VehicleId** | Pointer to **NullableString** |  | [optional] 
**DriverId** | Pointer to **NullableString** |  | [optional] 
**CoDriverId** | Pointer to **NullableString** |  | [optional] 
**SourceDriverId** | Pointer to **NullableString** |  | [optional] 
**SourceVehicleId** | Pointer to **NullableString** |  | [optional] 
**SourceCoDriverId** | Pointer to **NullableString** |  | [optional] 
**Location** | Pointer to [**NullableLocation3**](Location3.md) |  | [optional] 
**H3Index11** | Pointer to **NullableInt32** |  | [optional] 
**InferredAddress** | Pointer to **map[string]interface{}** |  | [optional] 
**Odometer** | Pointer to **NullableFloat32** |  | [optional] 
**FuelType** | Pointer to [**NullableEngineType**](EngineType.md) |  | [optional] 
**FuelVolume** | Pointer to **NullableFloat32** |  | [optional] 
**FuelVendor** | Pointer to **NullableString** |  | [optional] 
**TotalCost** | Pointer to **NullableString** |  | [optional] 
**Currency** | Pointer to **NullableString** |  | [optional] 
**OdometerUnit** | Pointer to [**DistanceUnit**](DistanceUnit.md) | Unit for odometer. | [optional] 

## Methods

### NewFuelTransactionRead

`func NewFuelTransactionRead(fleetId NullableString, id string, createdAt time.Time, updatedAt time.Time, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *FuelTransactionRead`

NewFuelTransactionRead instantiates a new FuelTransactionRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuelTransactionReadWithDefaults

`func NewFuelTransactionReadWithDefaults() *FuelTransactionRead`

NewFuelTransactionReadWithDefaults instantiates a new FuelTransactionRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *FuelTransactionRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *FuelTransactionRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *FuelTransactionRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### SetFleetIdNil

`func (o *FuelTransactionRead) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *FuelTransactionRead) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetFleetRef

`func (o *FuelTransactionRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *FuelTransactionRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *FuelTransactionRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *FuelTransactionRead) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *FuelTransactionRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *FuelTransactionRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetId

`func (o *FuelTransactionRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FuelTransactionRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FuelTransactionRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *FuelTransactionRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FuelTransactionRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FuelTransactionRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *FuelTransactionRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FuelTransactionRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FuelTransactionRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *FuelTransactionRead) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *FuelTransactionRead) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *FuelTransactionRead) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *FuelTransactionRead) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *FuelTransactionRead) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *FuelTransactionRead) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetConnectionId

`func (o *FuelTransactionRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *FuelTransactionRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *FuelTransactionRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetTspId

`func (o *FuelTransactionRead) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *FuelTransactionRead) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *FuelTransactionRead) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *FuelTransactionRead) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *FuelTransactionRead) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *FuelTransactionRead) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *FuelTransactionRead) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *FuelTransactionRead) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *FuelTransactionRead) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *FuelTransactionRead) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *FuelTransactionRead) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *FuelTransactionRead) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *FuelTransactionRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *FuelTransactionRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *FuelTransactionRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *FuelTransactionRead) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *FuelTransactionRead) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *FuelTransactionRead) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *FuelTransactionRead) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *FuelTransactionRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *FuelTransactionRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *FuelTransactionRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *FuelTransactionRead) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *FuelTransactionRead) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *FuelTransactionRead) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *FuelTransactionRead) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *FuelTransactionRead) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *FuelTransactionRead) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *FuelTransactionRead) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *FuelTransactionRead) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *FuelTransactionRead) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *FuelTransactionRead) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *FuelTransactionRead) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *FuelTransactionRead) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *FuelTransactionRead) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *FuelTransactionRead) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *FuelTransactionRead) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *FuelTransactionRead) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *FuelTransactionRead) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *FuelTransactionRead) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *FuelTransactionRead) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *FuelTransactionRead) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *FuelTransactionRead) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *FuelTransactionRead) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *FuelTransactionRead) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *FuelTransactionRead) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *FuelTransactionRead) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *FuelTransactionRead) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *FuelTransactionRead) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetVehicleId

`func (o *FuelTransactionRead) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *FuelTransactionRead) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *FuelTransactionRead) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *FuelTransactionRead) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *FuelTransactionRead) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *FuelTransactionRead) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetDriverId

`func (o *FuelTransactionRead) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *FuelTransactionRead) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *FuelTransactionRead) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *FuelTransactionRead) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *FuelTransactionRead) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *FuelTransactionRead) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetCoDriverId

`func (o *FuelTransactionRead) GetCoDriverId() string`

GetCoDriverId returns the CoDriverId field if non-nil, zero value otherwise.

### GetCoDriverIdOk

`func (o *FuelTransactionRead) GetCoDriverIdOk() (*string, bool)`

GetCoDriverIdOk returns a tuple with the CoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDriverId

`func (o *FuelTransactionRead) SetCoDriverId(v string)`

SetCoDriverId sets CoDriverId field to given value.

### HasCoDriverId

`func (o *FuelTransactionRead) HasCoDriverId() bool`

HasCoDriverId returns a boolean if a field has been set.

### SetCoDriverIdNil

`func (o *FuelTransactionRead) SetCoDriverIdNil(b bool)`

 SetCoDriverIdNil sets the value for CoDriverId to be an explicit nil

### UnsetCoDriverId
`func (o *FuelTransactionRead) UnsetCoDriverId()`

UnsetCoDriverId ensures that no value is present for CoDriverId, not even an explicit nil
### GetSourceDriverId

`func (o *FuelTransactionRead) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *FuelTransactionRead) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *FuelTransactionRead) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *FuelTransactionRead) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *FuelTransactionRead) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *FuelTransactionRead) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSourceVehicleId

`func (o *FuelTransactionRead) GetSourceVehicleId() string`

GetSourceVehicleId returns the SourceVehicleId field if non-nil, zero value otherwise.

### GetSourceVehicleIdOk

`func (o *FuelTransactionRead) GetSourceVehicleIdOk() (*string, bool)`

GetSourceVehicleIdOk returns a tuple with the SourceVehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVehicleId

`func (o *FuelTransactionRead) SetSourceVehicleId(v string)`

SetSourceVehicleId sets SourceVehicleId field to given value.

### HasSourceVehicleId

`func (o *FuelTransactionRead) HasSourceVehicleId() bool`

HasSourceVehicleId returns a boolean if a field has been set.

### SetSourceVehicleIdNil

`func (o *FuelTransactionRead) SetSourceVehicleIdNil(b bool)`

 SetSourceVehicleIdNil sets the value for SourceVehicleId to be an explicit nil

### UnsetSourceVehicleId
`func (o *FuelTransactionRead) UnsetSourceVehicleId()`

UnsetSourceVehicleId ensures that no value is present for SourceVehicleId, not even an explicit nil
### GetSourceCoDriverId

`func (o *FuelTransactionRead) GetSourceCoDriverId() string`

GetSourceCoDriverId returns the SourceCoDriverId field if non-nil, zero value otherwise.

### GetSourceCoDriverIdOk

`func (o *FuelTransactionRead) GetSourceCoDriverIdOk() (*string, bool)`

GetSourceCoDriverIdOk returns a tuple with the SourceCoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCoDriverId

`func (o *FuelTransactionRead) SetSourceCoDriverId(v string)`

SetSourceCoDriverId sets SourceCoDriverId field to given value.

### HasSourceCoDriverId

`func (o *FuelTransactionRead) HasSourceCoDriverId() bool`

HasSourceCoDriverId returns a boolean if a field has been set.

### SetSourceCoDriverIdNil

`func (o *FuelTransactionRead) SetSourceCoDriverIdNil(b bool)`

 SetSourceCoDriverIdNil sets the value for SourceCoDriverId to be an explicit nil

### UnsetSourceCoDriverId
`func (o *FuelTransactionRead) UnsetSourceCoDriverId()`

UnsetSourceCoDriverId ensures that no value is present for SourceCoDriverId, not even an explicit nil
### GetLocation

`func (o *FuelTransactionRead) GetLocation() Location3`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *FuelTransactionRead) GetLocationOk() (*Location3, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *FuelTransactionRead) SetLocation(v Location3)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *FuelTransactionRead) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *FuelTransactionRead) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *FuelTransactionRead) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetH3Index11

`func (o *FuelTransactionRead) GetH3Index11() int32`

GetH3Index11 returns the H3Index11 field if non-nil, zero value otherwise.

### GetH3Index11Ok

`func (o *FuelTransactionRead) GetH3Index11Ok() (*int32, bool)`

GetH3Index11Ok returns a tuple with the H3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH3Index11

`func (o *FuelTransactionRead) SetH3Index11(v int32)`

SetH3Index11 sets H3Index11 field to given value.

### HasH3Index11

`func (o *FuelTransactionRead) HasH3Index11() bool`

HasH3Index11 returns a boolean if a field has been set.

### SetH3Index11Nil

`func (o *FuelTransactionRead) SetH3Index11Nil(b bool)`

 SetH3Index11Nil sets the value for H3Index11 to be an explicit nil

### UnsetH3Index11
`func (o *FuelTransactionRead) UnsetH3Index11()`

UnsetH3Index11 ensures that no value is present for H3Index11, not even an explicit nil
### GetInferredAddress

`func (o *FuelTransactionRead) GetInferredAddress() map[string]interface{}`

GetInferredAddress returns the InferredAddress field if non-nil, zero value otherwise.

### GetInferredAddressOk

`func (o *FuelTransactionRead) GetInferredAddressOk() (*map[string]interface{}, bool)`

GetInferredAddressOk returns a tuple with the InferredAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferredAddress

`func (o *FuelTransactionRead) SetInferredAddress(v map[string]interface{})`

SetInferredAddress sets InferredAddress field to given value.

### HasInferredAddress

`func (o *FuelTransactionRead) HasInferredAddress() bool`

HasInferredAddress returns a boolean if a field has been set.

### SetInferredAddressNil

`func (o *FuelTransactionRead) SetInferredAddressNil(b bool)`

 SetInferredAddressNil sets the value for InferredAddress to be an explicit nil

### UnsetInferredAddress
`func (o *FuelTransactionRead) UnsetInferredAddress()`

UnsetInferredAddress ensures that no value is present for InferredAddress, not even an explicit nil
### GetOdometer

`func (o *FuelTransactionRead) GetOdometer() float32`

GetOdometer returns the Odometer field if non-nil, zero value otherwise.

### GetOdometerOk

`func (o *FuelTransactionRead) GetOdometerOk() (*float32, bool)`

GetOdometerOk returns a tuple with the Odometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometer

`func (o *FuelTransactionRead) SetOdometer(v float32)`

SetOdometer sets Odometer field to given value.

### HasOdometer

`func (o *FuelTransactionRead) HasOdometer() bool`

HasOdometer returns a boolean if a field has been set.

### SetOdometerNil

`func (o *FuelTransactionRead) SetOdometerNil(b bool)`

 SetOdometerNil sets the value for Odometer to be an explicit nil

### UnsetOdometer
`func (o *FuelTransactionRead) UnsetOdometer()`

UnsetOdometer ensures that no value is present for Odometer, not even an explicit nil
### GetFuelType

`func (o *FuelTransactionRead) GetFuelType() EngineType`

GetFuelType returns the FuelType field if non-nil, zero value otherwise.

### GetFuelTypeOk

`func (o *FuelTransactionRead) GetFuelTypeOk() (*EngineType, bool)`

GetFuelTypeOk returns a tuple with the FuelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelType

`func (o *FuelTransactionRead) SetFuelType(v EngineType)`

SetFuelType sets FuelType field to given value.

### HasFuelType

`func (o *FuelTransactionRead) HasFuelType() bool`

HasFuelType returns a boolean if a field has been set.

### SetFuelTypeNil

`func (o *FuelTransactionRead) SetFuelTypeNil(b bool)`

 SetFuelTypeNil sets the value for FuelType to be an explicit nil

### UnsetFuelType
`func (o *FuelTransactionRead) UnsetFuelType()`

UnsetFuelType ensures that no value is present for FuelType, not even an explicit nil
### GetFuelVolume

`func (o *FuelTransactionRead) GetFuelVolume() float32`

GetFuelVolume returns the FuelVolume field if non-nil, zero value otherwise.

### GetFuelVolumeOk

`func (o *FuelTransactionRead) GetFuelVolumeOk() (*float32, bool)`

GetFuelVolumeOk returns a tuple with the FuelVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelVolume

`func (o *FuelTransactionRead) SetFuelVolume(v float32)`

SetFuelVolume sets FuelVolume field to given value.

### HasFuelVolume

`func (o *FuelTransactionRead) HasFuelVolume() bool`

HasFuelVolume returns a boolean if a field has been set.

### SetFuelVolumeNil

`func (o *FuelTransactionRead) SetFuelVolumeNil(b bool)`

 SetFuelVolumeNil sets the value for FuelVolume to be an explicit nil

### UnsetFuelVolume
`func (o *FuelTransactionRead) UnsetFuelVolume()`

UnsetFuelVolume ensures that no value is present for FuelVolume, not even an explicit nil
### GetFuelVendor

`func (o *FuelTransactionRead) GetFuelVendor() string`

GetFuelVendor returns the FuelVendor field if non-nil, zero value otherwise.

### GetFuelVendorOk

`func (o *FuelTransactionRead) GetFuelVendorOk() (*string, bool)`

GetFuelVendorOk returns a tuple with the FuelVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuelVendor

`func (o *FuelTransactionRead) SetFuelVendor(v string)`

SetFuelVendor sets FuelVendor field to given value.

### HasFuelVendor

`func (o *FuelTransactionRead) HasFuelVendor() bool`

HasFuelVendor returns a boolean if a field has been set.

### SetFuelVendorNil

`func (o *FuelTransactionRead) SetFuelVendorNil(b bool)`

 SetFuelVendorNil sets the value for FuelVendor to be an explicit nil

### UnsetFuelVendor
`func (o *FuelTransactionRead) UnsetFuelVendor()`

UnsetFuelVendor ensures that no value is present for FuelVendor, not even an explicit nil
### GetTotalCost

`func (o *FuelTransactionRead) GetTotalCost() string`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *FuelTransactionRead) GetTotalCostOk() (*string, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *FuelTransactionRead) SetTotalCost(v string)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *FuelTransactionRead) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### SetTotalCostNil

`func (o *FuelTransactionRead) SetTotalCostNil(b bool)`

 SetTotalCostNil sets the value for TotalCost to be an explicit nil

### UnsetTotalCost
`func (o *FuelTransactionRead) UnsetTotalCost()`

UnsetTotalCost ensures that no value is present for TotalCost, not even an explicit nil
### GetCurrency

`func (o *FuelTransactionRead) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *FuelTransactionRead) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *FuelTransactionRead) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *FuelTransactionRead) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *FuelTransactionRead) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *FuelTransactionRead) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetOdometerUnit

`func (o *FuelTransactionRead) GetOdometerUnit() DistanceUnit`

GetOdometerUnit returns the OdometerUnit field if non-nil, zero value otherwise.

### GetOdometerUnitOk

`func (o *FuelTransactionRead) GetOdometerUnitOk() (*DistanceUnit, bool)`

GetOdometerUnitOk returns a tuple with the OdometerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerUnit

`func (o *FuelTransactionRead) SetOdometerUnit(v DistanceUnit)`

SetOdometerUnit sets OdometerUnit field to given value.

### HasOdometerUnit

`func (o *FuelTransactionRead) HasOdometerUnit() bool`

HasOdometerUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


