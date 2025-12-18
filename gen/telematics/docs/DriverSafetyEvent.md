# DriverSafetyEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the record at Catena Telematics. | 
**CreatedAt** | **time.Time** | Immutable: The datetime the record was ingested into Catena Telematics. | 
**UpdatedAt** | **time.Time** | The dateime the record was last modified in Catena Telematics. | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**FleetId** | **string** | Unique identifier of the fleet at Catena Telematics. This record belongs to this fleet. **Note: this is not the fleet ID in the TSP system or in your organization&#39;s systems, use &#x60;share_agreements&#x60; to map &#x60;fleet_ids&#x60; to &#x60;fleet_refs&#x60;.** | 
**ConnectionId** | **string** | Unique identifier of the connection at Catena Telematics through which this record was ingested. A connection represents a Fleet/TSP pairing. | 
**SourceName** | [**TspEnum**](TspEnum.md) | An enumeration identifying the TSP from which this record was sourced. | 
**SourceData** | Pointer to **map[string]interface{}** | Raw source payload as ingested from the TSP. **Note: use it for audit/debugging.** | [optional] 
**SourceId** | **string** | Unique identifier of the record in the TSP. **Note: we generate a unique composite key based on available fields if the TSP does not provide an unique ID.** | 
**SourceDataHash** | **string** | SHA-256 hash of the source data payload. **Note: we use it internally for idempotence and deduplication.** | 
**OccurredAt** | Pointer to **NullableTime** |  | [optional] 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 
**DriverId** | Pointer to **NullableString** |  | [optional] 
**VehicleId** | Pointer to **NullableString** |  | [optional] 
**SourceDriverId** | Pointer to **NullableString** |  | [optional] 
**SourceVehicleId** | Pointer to **NullableString** |  | [optional] 
**Event** | [**DriverSafetyEventEnum**](DriverSafetyEventEnum.md) | Normalized safety event type. | 
**Location** | Pointer to [**NullableLocation**](Location.md) |  | [optional] 
**H3Index11** | Pointer to **NullableInt32** |  | [optional] 
**DurationSeconds** | Pointer to **NullableInt32** |  | [optional] 
**EventMetadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDriverSafetyEvent

`func NewDriverSafetyEvent(id string, createdAt time.Time, updatedAt time.Time, fleetId string, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, event DriverSafetyEventEnum, ) *DriverSafetyEvent`

NewDriverSafetyEvent instantiates a new DriverSafetyEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverSafetyEventWithDefaults

`func NewDriverSafetyEventWithDefaults() *DriverSafetyEvent`

NewDriverSafetyEventWithDefaults instantiates a new DriverSafetyEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DriverSafetyEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriverSafetyEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriverSafetyEvent) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DriverSafetyEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DriverSafetyEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DriverSafetyEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DriverSafetyEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DriverSafetyEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DriverSafetyEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *DriverSafetyEvent) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DriverSafetyEvent) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DriverSafetyEvent) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DriverSafetyEvent) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *DriverSafetyEvent) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *DriverSafetyEvent) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetFleetId

`func (o *DriverSafetyEvent) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *DriverSafetyEvent) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *DriverSafetyEvent) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetConnectionId

`func (o *DriverSafetyEvent) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DriverSafetyEvent) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DriverSafetyEvent) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceName

`func (o *DriverSafetyEvent) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *DriverSafetyEvent) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *DriverSafetyEvent) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *DriverSafetyEvent) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *DriverSafetyEvent) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *DriverSafetyEvent) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *DriverSafetyEvent) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *DriverSafetyEvent) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DriverSafetyEvent) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DriverSafetyEvent) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *DriverSafetyEvent) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *DriverSafetyEvent) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *DriverSafetyEvent) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *DriverSafetyEvent) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *DriverSafetyEvent) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *DriverSafetyEvent) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *DriverSafetyEvent) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *DriverSafetyEvent) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *DriverSafetyEvent) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *DriverSafetyEvent) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *DriverSafetyEvent) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *DriverSafetyEvent) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *DriverSafetyEvent) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *DriverSafetyEvent) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *DriverSafetyEvent) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *DriverSafetyEvent) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *DriverSafetyEvent) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *DriverSafetyEvent) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *DriverSafetyEvent) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *DriverSafetyEvent) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *DriverSafetyEvent) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetDriverId

`func (o *DriverSafetyEvent) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *DriverSafetyEvent) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *DriverSafetyEvent) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *DriverSafetyEvent) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *DriverSafetyEvent) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *DriverSafetyEvent) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetVehicleId

`func (o *DriverSafetyEvent) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *DriverSafetyEvent) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *DriverSafetyEvent) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *DriverSafetyEvent) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *DriverSafetyEvent) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *DriverSafetyEvent) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetSourceDriverId

`func (o *DriverSafetyEvent) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *DriverSafetyEvent) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *DriverSafetyEvent) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *DriverSafetyEvent) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *DriverSafetyEvent) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *DriverSafetyEvent) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSourceVehicleId

`func (o *DriverSafetyEvent) GetSourceVehicleId() string`

GetSourceVehicleId returns the SourceVehicleId field if non-nil, zero value otherwise.

### GetSourceVehicleIdOk

`func (o *DriverSafetyEvent) GetSourceVehicleIdOk() (*string, bool)`

GetSourceVehicleIdOk returns a tuple with the SourceVehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVehicleId

`func (o *DriverSafetyEvent) SetSourceVehicleId(v string)`

SetSourceVehicleId sets SourceVehicleId field to given value.

### HasSourceVehicleId

`func (o *DriverSafetyEvent) HasSourceVehicleId() bool`

HasSourceVehicleId returns a boolean if a field has been set.

### SetSourceVehicleIdNil

`func (o *DriverSafetyEvent) SetSourceVehicleIdNil(b bool)`

 SetSourceVehicleIdNil sets the value for SourceVehicleId to be an explicit nil

### UnsetSourceVehicleId
`func (o *DriverSafetyEvent) UnsetSourceVehicleId()`

UnsetSourceVehicleId ensures that no value is present for SourceVehicleId, not even an explicit nil
### GetEvent

`func (o *DriverSafetyEvent) GetEvent() DriverSafetyEventEnum`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *DriverSafetyEvent) GetEventOk() (*DriverSafetyEventEnum, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *DriverSafetyEvent) SetEvent(v DriverSafetyEventEnum)`

SetEvent sets Event field to given value.


### GetLocation

`func (o *DriverSafetyEvent) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DriverSafetyEvent) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DriverSafetyEvent) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DriverSafetyEvent) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *DriverSafetyEvent) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *DriverSafetyEvent) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetH3Index11

`func (o *DriverSafetyEvent) GetH3Index11() int32`

GetH3Index11 returns the H3Index11 field if non-nil, zero value otherwise.

### GetH3Index11Ok

`func (o *DriverSafetyEvent) GetH3Index11Ok() (*int32, bool)`

GetH3Index11Ok returns a tuple with the H3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH3Index11

`func (o *DriverSafetyEvent) SetH3Index11(v int32)`

SetH3Index11 sets H3Index11 field to given value.

### HasH3Index11

`func (o *DriverSafetyEvent) HasH3Index11() bool`

HasH3Index11 returns a boolean if a field has been set.

### SetH3Index11Nil

`func (o *DriverSafetyEvent) SetH3Index11Nil(b bool)`

 SetH3Index11Nil sets the value for H3Index11 to be an explicit nil

### UnsetH3Index11
`func (o *DriverSafetyEvent) UnsetH3Index11()`

UnsetH3Index11 ensures that no value is present for H3Index11, not even an explicit nil
### GetDurationSeconds

`func (o *DriverSafetyEvent) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *DriverSafetyEvent) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *DriverSafetyEvent) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *DriverSafetyEvent) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *DriverSafetyEvent) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *DriverSafetyEvent) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetEventMetadata

`func (o *DriverSafetyEvent) GetEventMetadata() map[string]interface{}`

GetEventMetadata returns the EventMetadata field if non-nil, zero value otherwise.

### GetEventMetadataOk

`func (o *DriverSafetyEvent) GetEventMetadataOk() (*map[string]interface{}, bool)`

GetEventMetadataOk returns a tuple with the EventMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventMetadata

`func (o *DriverSafetyEvent) SetEventMetadata(v map[string]interface{})`

SetEventMetadata sets EventMetadata field to given value.

### HasEventMetadata

`func (o *DriverSafetyEvent) HasEventMetadata() bool`

HasEventMetadata returns a boolean if a field has been set.

### SetEventMetadataNil

`func (o *DriverSafetyEvent) SetEventMetadataNil(b bool)`

 SetEventMetadataNil sets the value for EventMetadata to be an explicit nil

### UnsetEventMetadata
`func (o *DriverSafetyEvent) UnsetEventMetadata()`

UnsetEventMetadata ensures that no value is present for EventMetadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


