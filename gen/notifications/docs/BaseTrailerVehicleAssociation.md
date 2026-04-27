# BaseTrailerVehicleAssociation

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
**TrailerId** | Pointer to **NullableString** |  | [optional] 
**VehicleId** | Pointer to **NullableString** |  | [optional] 
**SourceTrailerId** | Pointer to **NullableString** |  | [optional] 
**SourceVehicleId** | Pointer to **NullableString** |  | [optional] 
**TrailerSequence** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBaseTrailerVehicleAssociation

`func NewBaseTrailerVehicleAssociation(id string, fleetId string, fleetRef NullableString, sourceName TspEnum, connectionId string, sourceId string, createdAt time.Time, updatedAt time.Time, occurredAt time.Time, ) *BaseTrailerVehicleAssociation`

NewBaseTrailerVehicleAssociation instantiates a new BaseTrailerVehicleAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseTrailerVehicleAssociationWithDefaults

`func NewBaseTrailerVehicleAssociationWithDefaults() *BaseTrailerVehicleAssociation`

NewBaseTrailerVehicleAssociationWithDefaults instantiates a new BaseTrailerVehicleAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseTrailerVehicleAssociation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseTrailerVehicleAssociation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseTrailerVehicleAssociation) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseTrailerVehicleAssociation) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseTrailerVehicleAssociation) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseTrailerVehicleAssociation) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *BaseTrailerVehicleAssociation) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *BaseTrailerVehicleAssociation) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *BaseTrailerVehicleAssociation) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### SetFleetRefNil

`func (o *BaseTrailerVehicleAssociation) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *BaseTrailerVehicleAssociation) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetTspId

`func (o *BaseTrailerVehicleAssociation) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *BaseTrailerVehicleAssociation) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *BaseTrailerVehicleAssociation) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *BaseTrailerVehicleAssociation) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *BaseTrailerVehicleAssociation) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *BaseTrailerVehicleAssociation) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *BaseTrailerVehicleAssociation) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *BaseTrailerVehicleAssociation) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *BaseTrailerVehicleAssociation) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *BaseTrailerVehicleAssociation) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *BaseTrailerVehicleAssociation) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *BaseTrailerVehicleAssociation) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *BaseTrailerVehicleAssociation) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseTrailerVehicleAssociation) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseTrailerVehicleAssociation) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetConnectionId

`func (o *BaseTrailerVehicleAssociation) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseTrailerVehicleAssociation) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseTrailerVehicleAssociation) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceId

`func (o *BaseTrailerVehicleAssociation) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseTrailerVehicleAssociation) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseTrailerVehicleAssociation) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCreatedAt

`func (o *BaseTrailerVehicleAssociation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseTrailerVehicleAssociation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseTrailerVehicleAssociation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseTrailerVehicleAssociation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseTrailerVehicleAssociation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseTrailerVehicleAssociation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BaseTrailerVehicleAssociation) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BaseTrailerVehicleAssociation) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BaseTrailerVehicleAssociation) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BaseTrailerVehicleAssociation) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BaseTrailerVehicleAssociation) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BaseTrailerVehicleAssociation) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOccurredAt

`func (o *BaseTrailerVehicleAssociation) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BaseTrailerVehicleAssociation) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BaseTrailerVehicleAssociation) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetExecutionId

`func (o *BaseTrailerVehicleAssociation) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BaseTrailerVehicleAssociation) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BaseTrailerVehicleAssociation) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *BaseTrailerVehicleAssociation) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *BaseTrailerVehicleAssociation) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *BaseTrailerVehicleAssociation) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *BaseTrailerVehicleAssociation) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BaseTrailerVehicleAssociation) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BaseTrailerVehicleAssociation) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *BaseTrailerVehicleAssociation) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *BaseTrailerVehicleAssociation) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *BaseTrailerVehicleAssociation) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *BaseTrailerVehicleAssociation) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *BaseTrailerVehicleAssociation) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *BaseTrailerVehicleAssociation) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *BaseTrailerVehicleAssociation) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *BaseTrailerVehicleAssociation) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *BaseTrailerVehicleAssociation) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetTrailerId

`func (o *BaseTrailerVehicleAssociation) GetTrailerId() string`

GetTrailerId returns the TrailerId field if non-nil, zero value otherwise.

### GetTrailerIdOk

`func (o *BaseTrailerVehicleAssociation) GetTrailerIdOk() (*string, bool)`

GetTrailerIdOk returns a tuple with the TrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerId

`func (o *BaseTrailerVehicleAssociation) SetTrailerId(v string)`

SetTrailerId sets TrailerId field to given value.

### HasTrailerId

`func (o *BaseTrailerVehicleAssociation) HasTrailerId() bool`

HasTrailerId returns a boolean if a field has been set.

### SetTrailerIdNil

`func (o *BaseTrailerVehicleAssociation) SetTrailerIdNil(b bool)`

 SetTrailerIdNil sets the value for TrailerId to be an explicit nil

### UnsetTrailerId
`func (o *BaseTrailerVehicleAssociation) UnsetTrailerId()`

UnsetTrailerId ensures that no value is present for TrailerId, not even an explicit nil
### GetVehicleId

`func (o *BaseTrailerVehicleAssociation) GetVehicleId() string`

GetVehicleId returns the VehicleId field if non-nil, zero value otherwise.

### GetVehicleIdOk

`func (o *BaseTrailerVehicleAssociation) GetVehicleIdOk() (*string, bool)`

GetVehicleIdOk returns a tuple with the VehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleId

`func (o *BaseTrailerVehicleAssociation) SetVehicleId(v string)`

SetVehicleId sets VehicleId field to given value.

### HasVehicleId

`func (o *BaseTrailerVehicleAssociation) HasVehicleId() bool`

HasVehicleId returns a boolean if a field has been set.

### SetVehicleIdNil

`func (o *BaseTrailerVehicleAssociation) SetVehicleIdNil(b bool)`

 SetVehicleIdNil sets the value for VehicleId to be an explicit nil

### UnsetVehicleId
`func (o *BaseTrailerVehicleAssociation) UnsetVehicleId()`

UnsetVehicleId ensures that no value is present for VehicleId, not even an explicit nil
### GetSourceTrailerId

`func (o *BaseTrailerVehicleAssociation) GetSourceTrailerId() string`

GetSourceTrailerId returns the SourceTrailerId field if non-nil, zero value otherwise.

### GetSourceTrailerIdOk

`func (o *BaseTrailerVehicleAssociation) GetSourceTrailerIdOk() (*string, bool)`

GetSourceTrailerIdOk returns a tuple with the SourceTrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTrailerId

`func (o *BaseTrailerVehicleAssociation) SetSourceTrailerId(v string)`

SetSourceTrailerId sets SourceTrailerId field to given value.

### HasSourceTrailerId

`func (o *BaseTrailerVehicleAssociation) HasSourceTrailerId() bool`

HasSourceTrailerId returns a boolean if a field has been set.

### SetSourceTrailerIdNil

`func (o *BaseTrailerVehicleAssociation) SetSourceTrailerIdNil(b bool)`

 SetSourceTrailerIdNil sets the value for SourceTrailerId to be an explicit nil

### UnsetSourceTrailerId
`func (o *BaseTrailerVehicleAssociation) UnsetSourceTrailerId()`

UnsetSourceTrailerId ensures that no value is present for SourceTrailerId, not even an explicit nil
### GetSourceVehicleId

`func (o *BaseTrailerVehicleAssociation) GetSourceVehicleId() string`

GetSourceVehicleId returns the SourceVehicleId field if non-nil, zero value otherwise.

### GetSourceVehicleIdOk

`func (o *BaseTrailerVehicleAssociation) GetSourceVehicleIdOk() (*string, bool)`

GetSourceVehicleIdOk returns a tuple with the SourceVehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVehicleId

`func (o *BaseTrailerVehicleAssociation) SetSourceVehicleId(v string)`

SetSourceVehicleId sets SourceVehicleId field to given value.

### HasSourceVehicleId

`func (o *BaseTrailerVehicleAssociation) HasSourceVehicleId() bool`

HasSourceVehicleId returns a boolean if a field has been set.

### SetSourceVehicleIdNil

`func (o *BaseTrailerVehicleAssociation) SetSourceVehicleIdNil(b bool)`

 SetSourceVehicleIdNil sets the value for SourceVehicleId to be an explicit nil

### UnsetSourceVehicleId
`func (o *BaseTrailerVehicleAssociation) UnsetSourceVehicleId()`

UnsetSourceVehicleId ensures that no value is present for SourceVehicleId, not even an explicit nil
### GetTrailerSequence

`func (o *BaseTrailerVehicleAssociation) GetTrailerSequence() int32`

GetTrailerSequence returns the TrailerSequence field if non-nil, zero value otherwise.

### GetTrailerSequenceOk

`func (o *BaseTrailerVehicleAssociation) GetTrailerSequenceOk() (*int32, bool)`

GetTrailerSequenceOk returns a tuple with the TrailerSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerSequence

`func (o *BaseTrailerVehicleAssociation) SetTrailerSequence(v int32)`

SetTrailerSequence sets TrailerSequence field to given value.

### HasTrailerSequence

`func (o *BaseTrailerVehicleAssociation) HasTrailerSequence() bool`

HasTrailerSequence returns a boolean if a field has been set.

### SetTrailerSequenceNil

`func (o *BaseTrailerVehicleAssociation) SetTrailerSequenceNil(b bool)`

 SetTrailerSequenceNil sets the value for TrailerSequence to be an explicit nil

### UnsetTrailerSequence
`func (o *BaseTrailerVehicleAssociation) UnsetTrailerSequence()`

UnsetTrailerSequence ensures that no value is present for TrailerSequence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


