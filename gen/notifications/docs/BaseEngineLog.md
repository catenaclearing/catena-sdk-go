# BaseEngineLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal unique identifier for the telematics event record (Catena PK). | 
**FleetId** | **string** | The Catena fleet this record belongs to (multi-tenant scope). | 
**FleetRef** | **NullableString** |  | 
**TspId** | Pointer to **NullableString** |  | [optional] 
**TspSlug** | Pointer to **NullableString** |  | [optional] 
**SourceName** | [**TspEnum**](TspEnum.md) | The underlying telematics platform that provided this data (e.g., &#x60;samsara&#x60;, &#x60;motive&#x60;, &#x60;hos247&#x60;). Note: Some platforms like &#x60;hos247&#x60; offer white-labeling, so multiple TSPs may share the same source_name — use &#x60;tsp_id&#x60; or &#x60;tsp_slug&#x60; to identify the specific ELD provider. | 
**ConnectionId** | **string** | The specific fleet↔TSP connection through which this record was sourced. | 
**SourceId** | **string** | The ID of the record in the TSP or a deterministic ID/Hash generated from a composite unique key | 
**CreatedAt** | **time.Time** | Immutable: first time this record was ingested into our system. | 
**UpdatedAt** | **time.Time** | Last time we modified this record in our system. | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**OccurredAt** | **time.Time** | When the underlying event/observation occurred, as reported by the TSP, or the moment it was ingested by us if not available. | 
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

### NewBaseEngineLog

`func NewBaseEngineLog(id string, fleetId string, fleetRef NullableString, sourceName TspEnum, connectionId string, sourceId string, createdAt time.Time, updatedAt time.Time, occurredAt time.Time, ) *BaseEngineLog`

NewBaseEngineLog instantiates a new BaseEngineLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseEngineLogWithDefaults

`func NewBaseEngineLogWithDefaults() *BaseEngineLog`

NewBaseEngineLogWithDefaults instantiates a new BaseEngineLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseEngineLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseEngineLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseEngineLog) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseEngineLog) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseEngineLog) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseEngineLog) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *BaseEngineLog) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *BaseEngineLog) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *BaseEngineLog) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### SetFleetRefNil

`func (o *BaseEngineLog) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *BaseEngineLog) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetTspId

`func (o *BaseEngineLog) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *BaseEngineLog) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *BaseEngineLog) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *BaseEngineLog) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *BaseEngineLog) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *BaseEngineLog) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *BaseEngineLog) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *BaseEngineLog) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *BaseEngineLog) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *BaseEngineLog) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *BaseEngineLog) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *BaseEngineLog) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *BaseEngineLog) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseEngineLog) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseEngineLog) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetConnectionId

`func (o *BaseEngineLog) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseEngineLog) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseEngineLog) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceId

`func (o *BaseEngineLog) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseEngineLog) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseEngineLog) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCreatedAt

`func (o *BaseEngineLog) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseEngineLog) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseEngineLog) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseEngineLog) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseEngineLog) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseEngineLog) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BaseEngineLog) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BaseEngineLog) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BaseEngineLog) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BaseEngineLog) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BaseEngineLog) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BaseEngineLog) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOccurredAt

`func (o *BaseEngineLog) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BaseEngineLog) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BaseEngineLog) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetExecutionId

`func (o *BaseEngineLog) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BaseEngineLog) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BaseEngineLog) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *BaseEngineLog) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *BaseEngineLog) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *BaseEngineLog) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *BaseEngineLog) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BaseEngineLog) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BaseEngineLog) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *BaseEngineLog) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *BaseEngineLog) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *BaseEngineLog) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *BaseEngineLog) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *BaseEngineLog) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *BaseEngineLog) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *BaseEngineLog) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *BaseEngineLog) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *BaseEngineLog) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetVehicleId

`func (o *BaseEngineLog) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *BaseEngineLog) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *BaseEngineLog) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *BaseEngineLog) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *BaseEngineLog) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *BaseEngineLog) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetDriverId

`func (o *BaseEngineLog) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *BaseEngineLog) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *BaseEngineLog) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *BaseEngineLog) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *BaseEngineLog) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *BaseEngineLog) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetCoDriverId

`func (o *BaseEngineLog) GetCoDriverId() string`

GetCoDriverId returns the CoDriverId field if non-nil, zero value otherwise.

### GetCoDriverIdOk

`func (o *BaseEngineLog) GetCoDriverIdOk() (*string, bool)`

GetCoDriverIdOk returns a tuple with the CoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoDriverId

`func (o *BaseEngineLog) SetCoDriverId(v string)`

SetCoDriverId sets CoDriverId field to given value.

### HasCoDriverId

`func (o *BaseEngineLog) HasCoDriverId() bool`

HasCoDriverId returns a boolean if a field has been set.

### SetCoDriverIdNil

`func (o *BaseEngineLog) SetCoDriverIdNil(b bool)`

 SetCoDriverIdNil sets the value for CoDriverId to be an explicit nil

### UnsetCoDriverId
`func (o *BaseEngineLog) UnsetCoDriverId()`

UnsetCoDriverId ensures that no value is present for CoDriverId, not even an explicit nil
### GetSourceDriverId

`func (o *BaseEngineLog) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *BaseEngineLog) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *BaseEngineLog) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *BaseEngineLog) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *BaseEngineLog) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *BaseEngineLog) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSourceVehicleId

`func (o *BaseEngineLog) GetSourceVehicleId() string`

GetSourceVehicleId returns the SourceVehicleId field if non-nil, zero value otherwise.

### GetSourceVehicleIdOk

`func (o *BaseEngineLog) GetSourceVehicleIdOk() (*string, bool)`

GetSourceVehicleIdOk returns a tuple with the SourceVehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVehicleId

`func (o *BaseEngineLog) SetSourceVehicleId(v string)`

SetSourceVehicleId sets SourceVehicleId field to given value.

### HasSourceVehicleId

`func (o *BaseEngineLog) HasSourceVehicleId() bool`

HasSourceVehicleId returns a boolean if a field has been set.

### SetSourceVehicleIdNil

`func (o *BaseEngineLog) SetSourceVehicleIdNil(b bool)`

 SetSourceVehicleIdNil sets the value for SourceVehicleId to be an explicit nil

### UnsetSourceVehicleId
`func (o *BaseEngineLog) UnsetSourceVehicleId()`

UnsetSourceVehicleId ensures that no value is present for SourceVehicleId, not even an explicit nil
### GetSourceCoDriverId

`func (o *BaseEngineLog) GetSourceCoDriverId() string`

GetSourceCoDriverId returns the SourceCoDriverId field if non-nil, zero value otherwise.

### GetSourceCoDriverIdOk

`func (o *BaseEngineLog) GetSourceCoDriverIdOk() (*string, bool)`

GetSourceCoDriverIdOk returns a tuple with the SourceCoDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCoDriverId

`func (o *BaseEngineLog) SetSourceCoDriverId(v string)`

SetSourceCoDriverId sets SourceCoDriverId field to given value.

### HasSourceCoDriverId

`func (o *BaseEngineLog) HasSourceCoDriverId() bool`

HasSourceCoDriverId returns a boolean if a field has been set.

### SetSourceCoDriverIdNil

`func (o *BaseEngineLog) SetSourceCoDriverIdNil(b bool)`

 SetSourceCoDriverIdNil sets the value for SourceCoDriverId to be an explicit nil

### UnsetSourceCoDriverId
`func (o *BaseEngineLog) UnsetSourceCoDriverId()`

UnsetSourceCoDriverId ensures that no value is present for SourceCoDriverId, not even an explicit nil
### GetFaultCodeSource

`func (o *BaseEngineLog) GetFaultCodeSource() FaultCodeSourceEnum`

GetFaultCodeSource returns the FaultCodeSource field if non-nil, zero value otherwise.

### GetFaultCodeSourceOk

`func (o *BaseEngineLog) GetFaultCodeSourceOk() (*FaultCodeSourceEnum, bool)`

GetFaultCodeSourceOk returns a tuple with the FaultCodeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultCodeSource

`func (o *BaseEngineLog) SetFaultCodeSource(v FaultCodeSourceEnum)`

SetFaultCodeSource sets FaultCodeSource field to given value.

### HasFaultCodeSource

`func (o *BaseEngineLog) HasFaultCodeSource() bool`

HasFaultCodeSource returns a boolean if a field has been set.

### SetFaultCodeSourceNil

`func (o *BaseEngineLog) SetFaultCodeSourceNil(b bool)`

 SetFaultCodeSourceNil sets the value for FaultCodeSource to be an explicit nil

### UnsetFaultCodeSource
`func (o *BaseEngineLog) UnsetFaultCodeSource()`

UnsetFaultCodeSource ensures that no value is present for FaultCodeSource, not even an explicit nil
### GetFaultCode

`func (o *BaseEngineLog) GetFaultCode() map[string]interface{}`

GetFaultCode returns the FaultCode field if non-nil, zero value otherwise.

### GetFaultCodeOk

`func (o *BaseEngineLog) GetFaultCodeOk() (*map[string]interface{}, bool)`

GetFaultCodeOk returns a tuple with the FaultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultCode

`func (o *BaseEngineLog) SetFaultCode(v map[string]interface{})`

SetFaultCode sets FaultCode field to given value.

### HasFaultCode

`func (o *BaseEngineLog) HasFaultCode() bool`

HasFaultCode returns a boolean if a field has been set.

### SetFaultCodeNil

`func (o *BaseEngineLog) SetFaultCodeNil(b bool)`

 SetFaultCodeNil sets the value for FaultCode to be an explicit nil

### UnsetFaultCode
`func (o *BaseEngineLog) UnsetFaultCode()`

UnsetFaultCode ensures that no value is present for FaultCode, not even an explicit nil
### GetFaultCodeDescription

`func (o *BaseEngineLog) GetFaultCodeDescription() string`

GetFaultCodeDescription returns the FaultCodeDescription field if non-nil, zero value otherwise.

### GetFaultCodeDescriptionOk

`func (o *BaseEngineLog) GetFaultCodeDescriptionOk() (*string, bool)`

GetFaultCodeDescriptionOk returns a tuple with the FaultCodeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaultCodeDescription

`func (o *BaseEngineLog) SetFaultCodeDescription(v string)`

SetFaultCodeDescription sets FaultCodeDescription field to given value.

### HasFaultCodeDescription

`func (o *BaseEngineLog) HasFaultCodeDescription() bool`

HasFaultCodeDescription returns a boolean if a field has been set.

### SetFaultCodeDescriptionNil

`func (o *BaseEngineLog) SetFaultCodeDescriptionNil(b bool)`

 SetFaultCodeDescriptionNil sets the value for FaultCodeDescription to be an explicit nil

### UnsetFaultCodeDescription
`func (o *BaseEngineLog) UnsetFaultCodeDescription()`

UnsetFaultCodeDescription ensures that no value is present for FaultCodeDescription, not even an explicit nil
### GetCanonicalCode

`func (o *BaseEngineLog) GetCanonicalCode() CanonicalFaultCodeEnum`

GetCanonicalCode returns the CanonicalCode field if non-nil, zero value otherwise.

### GetCanonicalCodeOk

`func (o *BaseEngineLog) GetCanonicalCodeOk() (*CanonicalFaultCodeEnum, bool)`

GetCanonicalCodeOk returns a tuple with the CanonicalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalCode

`func (o *BaseEngineLog) SetCanonicalCode(v CanonicalFaultCodeEnum)`

SetCanonicalCode sets CanonicalCode field to given value.

### HasCanonicalCode

`func (o *BaseEngineLog) HasCanonicalCode() bool`

HasCanonicalCode returns a boolean if a field has been set.

### SetCanonicalCodeNil

`func (o *BaseEngineLog) SetCanonicalCodeNil(b bool)`

 SetCanonicalCodeNil sets the value for CanonicalCode to be an explicit nil

### UnsetCanonicalCode
`func (o *BaseEngineLog) UnsetCanonicalCode()`

UnsetCanonicalCode ensures that no value is present for CanonicalCode, not even an explicit nil
### GetStatus

`func (o *BaseEngineLog) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseEngineLog) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseEngineLog) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseEngineLog) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BaseEngineLog) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BaseEngineLog) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetEngineStatus

`func (o *BaseEngineLog) GetEngineStatus() bool`

GetEngineStatus returns the EngineStatus field if non-nil, zero value otherwise.

### GetEngineStatusOk

`func (o *BaseEngineLog) GetEngineStatusOk() (*bool, bool)`

GetEngineStatusOk returns a tuple with the EngineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineStatus

`func (o *BaseEngineLog) SetEngineStatus(v bool)`

SetEngineStatus sets EngineStatus field to given value.

### HasEngineStatus

`func (o *BaseEngineLog) HasEngineStatus() bool`

HasEngineStatus returns a boolean if a field has been set.

### SetEngineStatusNil

`func (o *BaseEngineLog) SetEngineStatusNil(b bool)`

 SetEngineStatusNil sets the value for EngineStatus to be an explicit nil

### UnsetEngineStatus
`func (o *BaseEngineLog) UnsetEngineStatus()`

UnsetEngineStatus ensures that no value is present for EngineStatus, not even an explicit nil
### GetMilStatus

`func (o *BaseEngineLog) GetMilStatus() bool`

GetMilStatus returns the MilStatus field if non-nil, zero value otherwise.

### GetMilStatusOk

`func (o *BaseEngineLog) GetMilStatusOk() (*bool, bool)`

GetMilStatusOk returns a tuple with the MilStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilStatus

`func (o *BaseEngineLog) SetMilStatus(v bool)`

SetMilStatus sets MilStatus field to given value.

### HasMilStatus

`func (o *BaseEngineLog) HasMilStatus() bool`

HasMilStatus returns a boolean if a field has been set.

### SetMilStatusNil

`func (o *BaseEngineLog) SetMilStatusNil(b bool)`

 SetMilStatusNil sets the value for MilStatus to be an explicit nil

### UnsetMilStatus
`func (o *BaseEngineLog) UnsetMilStatus()`

UnsetMilStatus ensures that no value is present for MilStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


