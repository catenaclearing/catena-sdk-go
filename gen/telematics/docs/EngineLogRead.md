# EngineLogRead

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
**FaultCodeSource** | Pointer to [**NullableFaultCodeSourceEnum**](FaultCodeSourceEnum.md) |  | [optional] 
**FaultCode** | Pointer to **map[string]interface{}** |  | [optional] 
**FaultCodeDescription** | Pointer to **NullableString** |  | [optional] 
**CanonicalCode** | Pointer to [**NullableCanonicalFaultCodeEnum**](CanonicalFaultCodeEnum.md) |  | [optional] 
**Status** | Pointer to [**NullableStatusEnum**](StatusEnum.md) |  | [optional] 
**EngineStatus** | Pointer to **NullableBool** |  | [optional] 
**MilStatus** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewEngineLogRead

`func NewEngineLogRead(fleetId NullableString, id string, createdAt time.Time, updatedAt time.Time, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *EngineLogRead`

NewEngineLogRead instantiates a new EngineLogRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineLogReadWithDefaults

`func NewEngineLogReadWithDefaults() *EngineLogRead`

NewEngineLogReadWithDefaults instantiates a new EngineLogRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *EngineLogRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *EngineLogRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *EngineLogRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### SetFleetIdNil

`func (o *EngineLogRead) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *EngineLogRead) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetFleetRef

`func (o *EngineLogRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *EngineLogRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *EngineLogRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *EngineLogRead) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *EngineLogRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *EngineLogRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetId

`func (o *EngineLogRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EngineLogRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EngineLogRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EngineLogRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EngineLogRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EngineLogRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EngineLogRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EngineLogRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EngineLogRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *EngineLogRead) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *EngineLogRead) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *EngineLogRead) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *EngineLogRead) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *EngineLogRead) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *EngineLogRead) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetConnectionId

`func (o *EngineLogRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *EngineLogRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *EngineLogRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetTspId

`func (o *EngineLogRead) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *EngineLogRead) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *EngineLogRead) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *EngineLogRead) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *EngineLogRead) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *EngineLogRead) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *EngineLogRead) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *EngineLogRead) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *EngineLogRead) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *EngineLogRead) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *EngineLogRead) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *EngineLogRead) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *EngineLogRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *EngineLogRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *EngineLogRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *EngineLogRead) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *EngineLogRead) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *EngineLogRead) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *EngineLogRead) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *EngineLogRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *EngineLogRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *EngineLogRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *EngineLogRead) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *EngineLogRead) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *EngineLogRead) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *EngineLogRead) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *EngineLogRead) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *EngineLogRead) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *EngineLogRead) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *EngineLogRead) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *EngineLogRead) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *EngineLogRead) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *EngineLogRead) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *EngineLogRead) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *EngineLogRead) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *EngineLogRead) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *EngineLogRead) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *EngineLogRead) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *EngineLogRead) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *EngineLogRead) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *EngineLogRead) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *EngineLogRead) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *EngineLogRead) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *EngineLogRead) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *EngineLogRead) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *EngineLogRead) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *EngineLogRead) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *EngineLogRead) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *EngineLogRead) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetVehicleId

`func (o *EngineLogRead) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *EngineLogRead) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *EngineLogRead) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *EngineLogRead) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *EngineLogRead) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *EngineLogRead) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetDriverId

`func (o *EngineLogRead) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *EngineLogRead) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *EngineLogRead) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *EngineLogRead) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *EngineLogRead) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *EngineLogRead) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetCoDriverId

`func (o *EngineLogRead) GetCoDriverId() string`

GetCoDriverId returns the CoDriverId field if non-nil, zero value otherwise.

### GetCoDriverIdOk

`func (o *EngineLogRead) GetCoDriverIdOk() (*string, bool)`

GetCoDriverIdOk returns a tuple with the CoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDriverId

`func (o *EngineLogRead) SetCoDriverId(v string)`

SetCoDriverId sets CoDriverId field to given value.

### HasCoDriverId

`func (o *EngineLogRead) HasCoDriverId() bool`

HasCoDriverId returns a boolean if a field has been set.

### SetCoDriverIdNil

`func (o *EngineLogRead) SetCoDriverIdNil(b bool)`

 SetCoDriverIdNil sets the value for CoDriverId to be an explicit nil

### UnsetCoDriverId
`func (o *EngineLogRead) UnsetCoDriverId()`

UnsetCoDriverId ensures that no value is present for CoDriverId, not even an explicit nil
### GetSourceDriverId

`func (o *EngineLogRead) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *EngineLogRead) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *EngineLogRead) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *EngineLogRead) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *EngineLogRead) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *EngineLogRead) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSourceVehicleId

`func (o *EngineLogRead) GetSourceVehicleId() string`

GetSourceVehicleId returns the SourceVehicleId field if non-nil, zero value otherwise.

### GetSourceVehicleIdOk

`func (o *EngineLogRead) GetSourceVehicleIdOk() (*string, bool)`

GetSourceVehicleIdOk returns a tuple with the SourceVehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVehicleId

`func (o *EngineLogRead) SetSourceVehicleId(v string)`

SetSourceVehicleId sets SourceVehicleId field to given value.

### HasSourceVehicleId

`func (o *EngineLogRead) HasSourceVehicleId() bool`

HasSourceVehicleId returns a boolean if a field has been set.

### SetSourceVehicleIdNil

`func (o *EngineLogRead) SetSourceVehicleIdNil(b bool)`

 SetSourceVehicleIdNil sets the value for SourceVehicleId to be an explicit nil

### UnsetSourceVehicleId
`func (o *EngineLogRead) UnsetSourceVehicleId()`

UnsetSourceVehicleId ensures that no value is present for SourceVehicleId, not even an explicit nil
### GetSourceCoDriverId

`func (o *EngineLogRead) GetSourceCoDriverId() string`

GetSourceCoDriverId returns the SourceCoDriverId field if non-nil, zero value otherwise.

### GetSourceCoDriverIdOk

`func (o *EngineLogRead) GetSourceCoDriverIdOk() (*string, bool)`

GetSourceCoDriverIdOk returns a tuple with the SourceCoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCoDriverId

`func (o *EngineLogRead) SetSourceCoDriverId(v string)`

SetSourceCoDriverId sets SourceCoDriverId field to given value.

### HasSourceCoDriverId

`func (o *EngineLogRead) HasSourceCoDriverId() bool`

HasSourceCoDriverId returns a boolean if a field has been set.

### SetSourceCoDriverIdNil

`func (o *EngineLogRead) SetSourceCoDriverIdNil(b bool)`

 SetSourceCoDriverIdNil sets the value for SourceCoDriverId to be an explicit nil

### UnsetSourceCoDriverId
`func (o *EngineLogRead) UnsetSourceCoDriverId()`

UnsetSourceCoDriverId ensures that no value is present for SourceCoDriverId, not even an explicit nil
### GetFaultCodeSource

`func (o *EngineLogRead) GetFaultCodeSource() FaultCodeSourceEnum`

GetFaultCodeSource returns the FaultCodeSource field if non-nil, zero value otherwise.

### GetFaultCodeSourceOk

`func (o *EngineLogRead) GetFaultCodeSourceOk() (*FaultCodeSourceEnum, bool)`

GetFaultCodeSourceOk returns a tuple with the FaultCodeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultCodeSource

`func (o *EngineLogRead) SetFaultCodeSource(v FaultCodeSourceEnum)`

SetFaultCodeSource sets FaultCodeSource field to given value.

### HasFaultCodeSource

`func (o *EngineLogRead) HasFaultCodeSource() bool`

HasFaultCodeSource returns a boolean if a field has been set.

### SetFaultCodeSourceNil

`func (o *EngineLogRead) SetFaultCodeSourceNil(b bool)`

 SetFaultCodeSourceNil sets the value for FaultCodeSource to be an explicit nil

### UnsetFaultCodeSource
`func (o *EngineLogRead) UnsetFaultCodeSource()`

UnsetFaultCodeSource ensures that no value is present for FaultCodeSource, not even an explicit nil
### GetFaultCode

`func (o *EngineLogRead) GetFaultCode() map[string]interface{}`

GetFaultCode returns the FaultCode field if non-nil, zero value otherwise.

### GetFaultCodeOk

`func (o *EngineLogRead) GetFaultCodeOk() (*map[string]interface{}, bool)`

GetFaultCodeOk returns a tuple with the FaultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultCode

`func (o *EngineLogRead) SetFaultCode(v map[string]interface{})`

SetFaultCode sets FaultCode field to given value.

### HasFaultCode

`func (o *EngineLogRead) HasFaultCode() bool`

HasFaultCode returns a boolean if a field has been set.

### SetFaultCodeNil

`func (o *EngineLogRead) SetFaultCodeNil(b bool)`

 SetFaultCodeNil sets the value for FaultCode to be an explicit nil

### UnsetFaultCode
`func (o *EngineLogRead) UnsetFaultCode()`

UnsetFaultCode ensures that no value is present for FaultCode, not even an explicit nil
### GetFaultCodeDescription

`func (o *EngineLogRead) GetFaultCodeDescription() string`

GetFaultCodeDescription returns the FaultCodeDescription field if non-nil, zero value otherwise.

### GetFaultCodeDescriptionOk

`func (o *EngineLogRead) GetFaultCodeDescriptionOk() (*string, bool)`

GetFaultCodeDescriptionOk returns a tuple with the FaultCodeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultCodeDescription

`func (o *EngineLogRead) SetFaultCodeDescription(v string)`

SetFaultCodeDescription sets FaultCodeDescription field to given value.

### HasFaultCodeDescription

`func (o *EngineLogRead) HasFaultCodeDescription() bool`

HasFaultCodeDescription returns a boolean if a field has been set.

### SetFaultCodeDescriptionNil

`func (o *EngineLogRead) SetFaultCodeDescriptionNil(b bool)`

 SetFaultCodeDescriptionNil sets the value for FaultCodeDescription to be an explicit nil

### UnsetFaultCodeDescription
`func (o *EngineLogRead) UnsetFaultCodeDescription()`

UnsetFaultCodeDescription ensures that no value is present for FaultCodeDescription, not even an explicit nil
### GetCanonicalCode

`func (o *EngineLogRead) GetCanonicalCode() CanonicalFaultCodeEnum`

GetCanonicalCode returns the CanonicalCode field if non-nil, zero value otherwise.

### GetCanonicalCodeOk

`func (o *EngineLogRead) GetCanonicalCodeOk() (*CanonicalFaultCodeEnum, bool)`

GetCanonicalCodeOk returns a tuple with the CanonicalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalCode

`func (o *EngineLogRead) SetCanonicalCode(v CanonicalFaultCodeEnum)`

SetCanonicalCode sets CanonicalCode field to given value.

### HasCanonicalCode

`func (o *EngineLogRead) HasCanonicalCode() bool`

HasCanonicalCode returns a boolean if a field has been set.

### SetCanonicalCodeNil

`func (o *EngineLogRead) SetCanonicalCodeNil(b bool)`

 SetCanonicalCodeNil sets the value for CanonicalCode to be an explicit nil

### UnsetCanonicalCode
`func (o *EngineLogRead) UnsetCanonicalCode()`

UnsetCanonicalCode ensures that no value is present for CanonicalCode, not even an explicit nil
### GetStatus

`func (o *EngineLogRead) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EngineLogRead) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EngineLogRead) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EngineLogRead) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *EngineLogRead) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *EngineLogRead) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetEngineStatus

`func (o *EngineLogRead) GetEngineStatus() bool`

GetEngineStatus returns the EngineStatus field if non-nil, zero value otherwise.

### GetEngineStatusOk

`func (o *EngineLogRead) GetEngineStatusOk() (*bool, bool)`

GetEngineStatusOk returns a tuple with the EngineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineStatus

`func (o *EngineLogRead) SetEngineStatus(v bool)`

SetEngineStatus sets EngineStatus field to given value.

### HasEngineStatus

`func (o *EngineLogRead) HasEngineStatus() bool`

HasEngineStatus returns a boolean if a field has been set.

### SetEngineStatusNil

`func (o *EngineLogRead) SetEngineStatusNil(b bool)`

 SetEngineStatusNil sets the value for EngineStatus to be an explicit nil

### UnsetEngineStatus
`func (o *EngineLogRead) UnsetEngineStatus()`

UnsetEngineStatus ensures that no value is present for EngineStatus, not even an explicit nil
### GetMilStatus

`func (o *EngineLogRead) GetMilStatus() bool`

GetMilStatus returns the MilStatus field if non-nil, zero value otherwise.

### GetMilStatusOk

`func (o *EngineLogRead) GetMilStatusOk() (*bool, bool)`

GetMilStatusOk returns a tuple with the MilStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilStatus

`func (o *EngineLogRead) SetMilStatus(v bool)`

SetMilStatus sets MilStatus field to given value.

### HasMilStatus

`func (o *EngineLogRead) HasMilStatus() bool`

HasMilStatus returns a boolean if a field has been set.

### SetMilStatusNil

`func (o *EngineLogRead) SetMilStatusNil(b bool)`

 SetMilStatusNil sets the value for MilStatus to be an explicit nil

### UnsetMilStatus
`func (o *EngineLogRead) UnsetMilStatus()`

UnsetMilStatus ensures that no value is present for MilStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


