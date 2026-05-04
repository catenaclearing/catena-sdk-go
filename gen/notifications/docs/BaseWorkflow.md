# BaseWorkflow

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
**SourceVehicleId** | Pointer to **NullableString** |  | [optional] 
**SourceDriverId** | Pointer to **NullableString** |  | [optional] 
**WorkflowStatus** | Pointer to [**NullableWorkflowStatusEnum**](WorkflowStatusEnum.md) |  | [optional] 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 
**TotalDistance** | Pointer to **NullableFloat32** |  | [optional] 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**Stops** | Pointer to [**[]WorkflowStop**](WorkflowStop.md) |  | [optional] 

## Methods

### NewBaseWorkflow

`func NewBaseWorkflow(id string, fleetId string, fleetRef NullableString, sourceName TspEnum, connectionId string, sourceId string, createdAt time.Time, updatedAt time.Time, occurredAt time.Time, ) *BaseWorkflow`

NewBaseWorkflow instantiates a new BaseWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseWorkflowWithDefaults

`func NewBaseWorkflowWithDefaults() *BaseWorkflow`

NewBaseWorkflowWithDefaults instantiates a new BaseWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseWorkflow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseWorkflow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseWorkflow) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseWorkflow) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseWorkflow) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseWorkflow) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *BaseWorkflow) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *BaseWorkflow) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *BaseWorkflow) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### SetFleetRefNil

`func (o *BaseWorkflow) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *BaseWorkflow) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetTspId

`func (o *BaseWorkflow) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *BaseWorkflow) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *BaseWorkflow) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *BaseWorkflow) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *BaseWorkflow) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *BaseWorkflow) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *BaseWorkflow) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *BaseWorkflow) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *BaseWorkflow) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *BaseWorkflow) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *BaseWorkflow) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *BaseWorkflow) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *BaseWorkflow) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseWorkflow) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseWorkflow) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetConnectionId

`func (o *BaseWorkflow) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseWorkflow) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseWorkflow) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceId

`func (o *BaseWorkflow) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseWorkflow) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseWorkflow) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCreatedAt

`func (o *BaseWorkflow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseWorkflow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseWorkflow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseWorkflow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseWorkflow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseWorkflow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BaseWorkflow) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BaseWorkflow) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BaseWorkflow) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BaseWorkflow) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BaseWorkflow) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BaseWorkflow) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOccurredAt

`func (o *BaseWorkflow) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BaseWorkflow) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BaseWorkflow) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetExecutionId

`func (o *BaseWorkflow) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BaseWorkflow) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BaseWorkflow) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *BaseWorkflow) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *BaseWorkflow) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *BaseWorkflow) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *BaseWorkflow) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BaseWorkflow) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BaseWorkflow) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *BaseWorkflow) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *BaseWorkflow) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *BaseWorkflow) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *BaseWorkflow) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *BaseWorkflow) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *BaseWorkflow) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *BaseWorkflow) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *BaseWorkflow) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *BaseWorkflow) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetVehicleId

`func (o *BaseWorkflow) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *BaseWorkflow) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *BaseWorkflow) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *BaseWorkflow) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *BaseWorkflow) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *BaseWorkflow) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetDriverId

`func (o *BaseWorkflow) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *BaseWorkflow) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *BaseWorkflow) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *BaseWorkflow) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *BaseWorkflow) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *BaseWorkflow) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetSourceVehicleId

`func (o *BaseWorkflow) GetSourceVehicleId() string`

GetSourceVehicleId returns the SourceVehicleId field if non-nil, zero value otherwise.

### GetSourceVehicleIdOk

`func (o *BaseWorkflow) GetSourceVehicleIdOk() (*string, bool)`

GetSourceVehicleIdOk returns a tuple with the SourceVehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVehicleId

`func (o *BaseWorkflow) SetSourceVehicleId(v string)`

SetSourceVehicleId sets SourceVehicleId field to given value.

### HasSourceVehicleId

`func (o *BaseWorkflow) HasSourceVehicleId() bool`

HasSourceVehicleId returns a boolean if a field has been set.

### SetSourceVehicleIdNil

`func (o *BaseWorkflow) SetSourceVehicleIdNil(b bool)`

 SetSourceVehicleIdNil sets the value for SourceVehicleId to be an explicit nil

### UnsetSourceVehicleId
`func (o *BaseWorkflow) UnsetSourceVehicleId()`

UnsetSourceVehicleId ensures that no value is present for SourceVehicleId, not even an explicit nil
### GetSourceDriverId

`func (o *BaseWorkflow) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *BaseWorkflow) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *BaseWorkflow) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *BaseWorkflow) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *BaseWorkflow) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *BaseWorkflow) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetWorkflowStatus

`func (o *BaseWorkflow) GetWorkflowStatus() WorkflowStatusEnum`

GetWorkflowStatus returns the WorkflowStatus field if non-nil, zero value otherwise.

### GetWorkflowStatusOk

`func (o *BaseWorkflow) GetWorkflowStatusOk() (*WorkflowStatusEnum, bool)`

GetWorkflowStatusOk returns a tuple with the WorkflowStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowStatus

`func (o *BaseWorkflow) SetWorkflowStatus(v WorkflowStatusEnum)`

SetWorkflowStatus sets WorkflowStatus field to given value.

### HasWorkflowStatus

`func (o *BaseWorkflow) HasWorkflowStatus() bool`

HasWorkflowStatus returns a boolean if a field has been set.

### SetWorkflowStatusNil

`func (o *BaseWorkflow) SetWorkflowStatusNil(b bool)`

 SetWorkflowStatusNil sets the value for WorkflowStatus to be an explicit nil

### UnsetWorkflowStatus
`func (o *BaseWorkflow) UnsetWorkflowStatus()`

UnsetWorkflowStatus ensures that no value is present for WorkflowStatus, not even an explicit nil
### GetStartedAt

`func (o *BaseWorkflow) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BaseWorkflow) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BaseWorkflow) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *BaseWorkflow) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *BaseWorkflow) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *BaseWorkflow) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *BaseWorkflow) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *BaseWorkflow) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *BaseWorkflow) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *BaseWorkflow) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *BaseWorkflow) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *BaseWorkflow) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetTotalDistance

`func (o *BaseWorkflow) GetTotalDistance() float32`

GetTotalDistance returns the TotalDistance field if non-nil, zero value otherwise.

### GetTotalDistanceOk

`func (o *BaseWorkflow) GetTotalDistanceOk() (*float32, bool)`

GetTotalDistanceOk returns a tuple with the TotalDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDistance

`func (o *BaseWorkflow) SetTotalDistance(v float32)`

SetTotalDistance sets TotalDistance field to given value.

### HasTotalDistance

`func (o *BaseWorkflow) HasTotalDistance() bool`

HasTotalDistance returns a boolean if a field has been set.

### SetTotalDistanceNil

`func (o *BaseWorkflow) SetTotalDistanceNil(b bool)`

 SetTotalDistanceNil sets the value for TotalDistance to be an explicit nil

### UnsetTotalDistance
`func (o *BaseWorkflow) UnsetTotalDistance()`

UnsetTotalDistance ensures that no value is present for TotalDistance, not even an explicit nil
### GetInstructions

`func (o *BaseWorkflow) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *BaseWorkflow) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *BaseWorkflow) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *BaseWorkflow) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *BaseWorkflow) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *BaseWorkflow) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetStops

`func (o *BaseWorkflow) GetStops() []WorkflowStop`

GetStops returns the Stops field if non-nil, zero value otherwise.

### GetStopsOk

`func (o *BaseWorkflow) GetStopsOk() (*[]WorkflowStop, bool)`

GetStopsOk returns a tuple with the Stops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStops

`func (o *BaseWorkflow) SetStops(v []WorkflowStop)`

SetStops sets Stops field to given value.

### HasStops

`func (o *BaseWorkflow) HasStops() bool`

HasStops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


