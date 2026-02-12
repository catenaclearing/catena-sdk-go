# BaseHosViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal unique identifier for the telematics event record (Catena PK). | 
**FleetId** | **string** | The Catena fleet this record belongs to (multi-tenant scope). | 
**FleetRef** | **NullableString** |  | 
**SourceName** | [**TspEnum**](TspEnum.md) | The name of the source | 
**ConnectionId** | **string** | The specific fleet↔TSP connection through which this record was sourced. | 
**SourceId** | **string** | The ID of the record in the TSP or a deterministic ID/Hash generated from a composite unique key | 
**CreatedAt** | **time.Time** | Immutable: first time this record was ingested into our system. | 
**UpdatedAt** | **time.Time** | Last time we modified this record in our system. | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**OccurredAt** | **time.Time** | When the underlying event/observation occurred, as reported by the TSP, or the moment it was ingested by us if not available. | 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 
**HosEventId** | Pointer to **NullableString** |  | [optional] 
**DriverId** | Pointer to **NullableString** |  | [optional] 
**SourceDriverId** | Pointer to **NullableString** |  | [optional] 
**SourceHosEventId** | Pointer to **NullableString** |  | [optional] 
**ViolationCode** | Pointer to [**NullableHosViolationCodeEnum**](HosViolationCodeEnum.md) |  | [optional] 
**ViolationCategory** | Pointer to [**NullableHosViolationCategoryEnum**](HosViolationCategoryEnum.md) |  | [optional] 
**ViolationDescription** | Pointer to **NullableString** |  | [optional] 
**StartTime** | Pointer to **NullableTime** |  | [optional] 
**EndTime** | Pointer to **NullableTime** |  | [optional] 
**Duration** | Pointer to **NullableInt32** |  | [optional] 
**HoursLimit** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewBaseHosViolation

`func NewBaseHosViolation(id string, fleetId string, fleetRef NullableString, sourceName TspEnum, connectionId string, sourceId string, createdAt time.Time, updatedAt time.Time, occurredAt time.Time, ) *BaseHosViolation`

NewBaseHosViolation instantiates a new BaseHosViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseHosViolationWithDefaults

`func NewBaseHosViolationWithDefaults() *BaseHosViolation`

NewBaseHosViolationWithDefaults instantiates a new BaseHosViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseHosViolation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseHosViolation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseHosViolation) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseHosViolation) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseHosViolation) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseHosViolation) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *BaseHosViolation) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *BaseHosViolation) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *BaseHosViolation) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### SetFleetRefNil

`func (o *BaseHosViolation) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *BaseHosViolation) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetSourceName

`func (o *BaseHosViolation) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseHosViolation) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseHosViolation) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetConnectionId

`func (o *BaseHosViolation) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseHosViolation) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseHosViolation) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceId

`func (o *BaseHosViolation) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseHosViolation) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseHosViolation) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCreatedAt

`func (o *BaseHosViolation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseHosViolation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseHosViolation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseHosViolation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseHosViolation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseHosViolation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BaseHosViolation) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BaseHosViolation) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BaseHosViolation) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BaseHosViolation) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BaseHosViolation) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BaseHosViolation) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOccurredAt

`func (o *BaseHosViolation) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BaseHosViolation) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BaseHosViolation) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetExecutionId

`func (o *BaseHosViolation) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BaseHosViolation) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BaseHosViolation) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *BaseHosViolation) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *BaseHosViolation) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *BaseHosViolation) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *BaseHosViolation) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BaseHosViolation) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BaseHosViolation) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *BaseHosViolation) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *BaseHosViolation) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *BaseHosViolation) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetHosEventId

`func (o *BaseHosViolation) GetHosEventId() string`

GetHosEventId returns the HosEventId field if non-nil, zero value otherwise.

### GetHosEventIdOk

`func (o *BaseHosViolation) GetHosEventIdOk() (*string, bool)`

GetHosEventIdOk returns a tuple with the HosEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosEventId

`func (o *BaseHosViolation) SetHosEventId(v string)`

SetHosEventId sets HosEventId field to given value.

### HasHosEventId

`func (o *BaseHosViolation) HasHosEventId() bool`

HasHosEventId returns a boolean if a field has been set.

### SetHosEventIdNil

`func (o *BaseHosViolation) SetHosEventIdNil(b bool)`

 SetHosEventIdNil sets the value for HosEventId to be an explicit nil

### UnsetHosEventId
`func (o *BaseHosViolation) UnsetHosEventId()`

UnsetHosEventId ensures that no value is present for HosEventId, not even an explicit nil
### GetDriverId

`func (o *BaseHosViolation) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *BaseHosViolation) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *BaseHosViolation) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *BaseHosViolation) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *BaseHosViolation) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *BaseHosViolation) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetSourceDriverId

`func (o *BaseHosViolation) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *BaseHosViolation) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *BaseHosViolation) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *BaseHosViolation) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *BaseHosViolation) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *BaseHosViolation) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSourceHosEventId

`func (o *BaseHosViolation) GetSourceHosEventId() string`

GetSourceHosEventId returns the SourceHosEventId field if non-nil, zero value otherwise.

### GetSourceHosEventIdOk

`func (o *BaseHosViolation) GetSourceHosEventIdOk() (*string, bool)`

GetSourceHosEventIdOk returns a tuple with the SourceHosEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHosEventId

`func (o *BaseHosViolation) SetSourceHosEventId(v string)`

SetSourceHosEventId sets SourceHosEventId field to given value.

### HasSourceHosEventId

`func (o *BaseHosViolation) HasSourceHosEventId() bool`

HasSourceHosEventId returns a boolean if a field has been set.

### SetSourceHosEventIdNil

`func (o *BaseHosViolation) SetSourceHosEventIdNil(b bool)`

 SetSourceHosEventIdNil sets the value for SourceHosEventId to be an explicit nil

### UnsetSourceHosEventId
`func (o *BaseHosViolation) UnsetSourceHosEventId()`

UnsetSourceHosEventId ensures that no value is present for SourceHosEventId, not even an explicit nil
### GetViolationCode

`func (o *BaseHosViolation) GetViolationCode() HosViolationCodeEnum`

GetViolationCode returns the ViolationCode field if non-nil, zero value otherwise.

### GetViolationCodeOk

`func (o *BaseHosViolation) GetViolationCodeOk() (*HosViolationCodeEnum, bool)`

GetViolationCodeOk returns a tuple with the ViolationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCode

`func (o *BaseHosViolation) SetViolationCode(v HosViolationCodeEnum)`

SetViolationCode sets ViolationCode field to given value.

### HasViolationCode

`func (o *BaseHosViolation) HasViolationCode() bool`

HasViolationCode returns a boolean if a field has been set.

### SetViolationCodeNil

`func (o *BaseHosViolation) SetViolationCodeNil(b bool)`

 SetViolationCodeNil sets the value for ViolationCode to be an explicit nil

### UnsetViolationCode
`func (o *BaseHosViolation) UnsetViolationCode()`

UnsetViolationCode ensures that no value is present for ViolationCode, not even an explicit nil
### GetViolationCategory

`func (o *BaseHosViolation) GetViolationCategory() HosViolationCategoryEnum`

GetViolationCategory returns the ViolationCategory field if non-nil, zero value otherwise.

### GetViolationCategoryOk

`func (o *BaseHosViolation) GetViolationCategoryOk() (*HosViolationCategoryEnum, bool)`

GetViolationCategoryOk returns a tuple with the ViolationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCategory

`func (o *BaseHosViolation) SetViolationCategory(v HosViolationCategoryEnum)`

SetViolationCategory sets ViolationCategory field to given value.

### HasViolationCategory

`func (o *BaseHosViolation) HasViolationCategory() bool`

HasViolationCategory returns a boolean if a field has been set.

### SetViolationCategoryNil

`func (o *BaseHosViolation) SetViolationCategoryNil(b bool)`

 SetViolationCategoryNil sets the value for ViolationCategory to be an explicit nil

### UnsetViolationCategory
`func (o *BaseHosViolation) UnsetViolationCategory()`

UnsetViolationCategory ensures that no value is present for ViolationCategory, not even an explicit nil
### GetViolationDescription

`func (o *BaseHosViolation) GetViolationDescription() string`

GetViolationDescription returns the ViolationDescription field if non-nil, zero value otherwise.

### GetViolationDescriptionOk

`func (o *BaseHosViolation) GetViolationDescriptionOk() (*string, bool)`

GetViolationDescriptionOk returns a tuple with the ViolationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationDescription

`func (o *BaseHosViolation) SetViolationDescription(v string)`

SetViolationDescription sets ViolationDescription field to given value.

### HasViolationDescription

`func (o *BaseHosViolation) HasViolationDescription() bool`

HasViolationDescription returns a boolean if a field has been set.

### SetViolationDescriptionNil

`func (o *BaseHosViolation) SetViolationDescriptionNil(b bool)`

 SetViolationDescriptionNil sets the value for ViolationDescription to be an explicit nil

### UnsetViolationDescription
`func (o *BaseHosViolation) UnsetViolationDescription()`

UnsetViolationDescription ensures that no value is present for ViolationDescription, not even an explicit nil
### GetStartTime

`func (o *BaseHosViolation) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *BaseHosViolation) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *BaseHosViolation) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *BaseHosViolation) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *BaseHosViolation) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *BaseHosViolation) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *BaseHosViolation) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *BaseHosViolation) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *BaseHosViolation) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *BaseHosViolation) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *BaseHosViolation) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *BaseHosViolation) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetDuration

`func (o *BaseHosViolation) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *BaseHosViolation) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *BaseHosViolation) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *BaseHosViolation) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *BaseHosViolation) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *BaseHosViolation) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetHoursLimit

`func (o *BaseHosViolation) GetHoursLimit() float32`

GetHoursLimit returns the HoursLimit field if non-nil, zero value otherwise.

### GetHoursLimitOk

`func (o *BaseHosViolation) GetHoursLimitOk() (*float32, bool)`

GetHoursLimitOk returns a tuple with the HoursLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursLimit

`func (o *BaseHosViolation) SetHoursLimit(v float32)`

SetHoursLimit sets HoursLimit field to given value.

### HasHoursLimit

`func (o *BaseHosViolation) HasHoursLimit() bool`

HasHoursLimit returns a boolean if a field has been set.

### SetHoursLimitNil

`func (o *BaseHosViolation) SetHoursLimitNil(b bool)`

 SetHoursLimitNil sets the value for HoursLimit to be an explicit nil

### UnsetHoursLimit
`func (o *BaseHosViolation) UnsetHoursLimit()`

UnsetHoursLimit ensures that no value is present for HoursLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


