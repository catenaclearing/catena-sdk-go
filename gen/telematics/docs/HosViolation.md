# HosViolation

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
**HosLogId** | Pointer to **NullableString** |  | [optional] 
**DriverId** | Pointer to **NullableString** |  | [optional] 
**SourceDriverId** | Pointer to **NullableString** |  | [optional] 
**SourceHosLogId** | Pointer to **NullableString** |  | [optional] 
**ViolationCode** | Pointer to [**NullableHosViolationCodeEnum**](HosViolationCodeEnum.md) |  | [optional] 
**ViolationCategory** | Pointer to [**NullableHosViolationCategoryEnum**](HosViolationCategoryEnum.md) |  | [optional] 
**ViolationDescription** | Pointer to **NullableString** |  | [optional] 
**StartTime** | Pointer to **NullableTime** |  | [optional] 
**EndTime** | Pointer to **NullableTime** |  | [optional] 
**Duration** | Pointer to **NullableInt32** |  | [optional] 
**HoursLimit** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewHosViolation

`func NewHosViolation(id string, createdAt time.Time, updatedAt time.Time, fleetId string, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *HosViolation`

NewHosViolation instantiates a new HosViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHosViolationWithDefaults

`func NewHosViolationWithDefaults() *HosViolation`

NewHosViolationWithDefaults instantiates a new HosViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HosViolation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HosViolation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HosViolation) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *HosViolation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HosViolation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HosViolation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *HosViolation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HosViolation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HosViolation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *HosViolation) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *HosViolation) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *HosViolation) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *HosViolation) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *HosViolation) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *HosViolation) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetFleetId

`func (o *HosViolation) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *HosViolation) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *HosViolation) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetConnectionId

`func (o *HosViolation) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *HosViolation) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *HosViolation) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceName

`func (o *HosViolation) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *HosViolation) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *HosViolation) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *HosViolation) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *HosViolation) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *HosViolation) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *HosViolation) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *HosViolation) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *HosViolation) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *HosViolation) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *HosViolation) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *HosViolation) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *HosViolation) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *HosViolation) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *HosViolation) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *HosViolation) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *HosViolation) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *HosViolation) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *HosViolation) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *HosViolation) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HosViolation) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HosViolation) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HosViolation) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HosViolation) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HosViolation) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *HosViolation) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *HosViolation) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *HosViolation) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *HosViolation) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *HosViolation) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *HosViolation) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetHosLogId

`func (o *HosViolation) GetHosLogId() string`

GetHosLogId returns the HosLogId field if non-nil, zero value otherwise.

### GetHosLogIdOk

`func (o *HosViolation) GetHosLogIdOk() (*string, bool)`

GetHosLogIdOk returns a tuple with the HosLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosLogId

`func (o *HosViolation) SetHosLogId(v string)`

SetHosLogId sets HosLogId field to given value.

### HasHosLogId

`func (o *HosViolation) HasHosLogId() bool`

HasHosLogId returns a boolean if a field has been set.

### SetHosLogIdNil

`func (o *HosViolation) SetHosLogIdNil(b bool)`

 SetHosLogIdNil sets the value for HosLogId to be an explicit nil

### UnsetHosLogId
`func (o *HosViolation) UnsetHosLogId()`

UnsetHosLogId ensures that no value is present for HosLogId, not even an explicit nil
### GetDriverId

`func (o *HosViolation) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *HosViolation) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *HosViolation) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *HosViolation) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *HosViolation) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *HosViolation) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetSourceDriverId

`func (o *HosViolation) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *HosViolation) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *HosViolation) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *HosViolation) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *HosViolation) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *HosViolation) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSourceHosLogId

`func (o *HosViolation) GetSourceHosLogId() string`

GetSourceHosLogId returns the SourceHosLogId field if non-nil, zero value otherwise.

### GetSourceHosLogIdOk

`func (o *HosViolation) GetSourceHosLogIdOk() (*string, bool)`

GetSourceHosLogIdOk returns a tuple with the SourceHosLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHosLogId

`func (o *HosViolation) SetSourceHosLogId(v string)`

SetSourceHosLogId sets SourceHosLogId field to given value.

### HasSourceHosLogId

`func (o *HosViolation) HasSourceHosLogId() bool`

HasSourceHosLogId returns a boolean if a field has been set.

### SetSourceHosLogIdNil

`func (o *HosViolation) SetSourceHosLogIdNil(b bool)`

 SetSourceHosLogIdNil sets the value for SourceHosLogId to be an explicit nil

### UnsetSourceHosLogId
`func (o *HosViolation) UnsetSourceHosLogId()`

UnsetSourceHosLogId ensures that no value is present for SourceHosLogId, not even an explicit nil
### GetViolationCode

`func (o *HosViolation) GetViolationCode() HosViolationCodeEnum`

GetViolationCode returns the ViolationCode field if non-nil, zero value otherwise.

### GetViolationCodeOk

`func (o *HosViolation) GetViolationCodeOk() (*HosViolationCodeEnum, bool)`

GetViolationCodeOk returns a tuple with the ViolationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCode

`func (o *HosViolation) SetViolationCode(v HosViolationCodeEnum)`

SetViolationCode sets ViolationCode field to given value.

### HasViolationCode

`func (o *HosViolation) HasViolationCode() bool`

HasViolationCode returns a boolean if a field has been set.

### SetViolationCodeNil

`func (o *HosViolation) SetViolationCodeNil(b bool)`

 SetViolationCodeNil sets the value for ViolationCode to be an explicit nil

### UnsetViolationCode
`func (o *HosViolation) UnsetViolationCode()`

UnsetViolationCode ensures that no value is present for ViolationCode, not even an explicit nil
### GetViolationCategory

`func (o *HosViolation) GetViolationCategory() HosViolationCategoryEnum`

GetViolationCategory returns the ViolationCategory field if non-nil, zero value otherwise.

### GetViolationCategoryOk

`func (o *HosViolation) GetViolationCategoryOk() (*HosViolationCategoryEnum, bool)`

GetViolationCategoryOk returns a tuple with the ViolationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCategory

`func (o *HosViolation) SetViolationCategory(v HosViolationCategoryEnum)`

SetViolationCategory sets ViolationCategory field to given value.

### HasViolationCategory

`func (o *HosViolation) HasViolationCategory() bool`

HasViolationCategory returns a boolean if a field has been set.

### SetViolationCategoryNil

`func (o *HosViolation) SetViolationCategoryNil(b bool)`

 SetViolationCategoryNil sets the value for ViolationCategory to be an explicit nil

### UnsetViolationCategory
`func (o *HosViolation) UnsetViolationCategory()`

UnsetViolationCategory ensures that no value is present for ViolationCategory, not even an explicit nil
### GetViolationDescription

`func (o *HosViolation) GetViolationDescription() string`

GetViolationDescription returns the ViolationDescription field if non-nil, zero value otherwise.

### GetViolationDescriptionOk

`func (o *HosViolation) GetViolationDescriptionOk() (*string, bool)`

GetViolationDescriptionOk returns a tuple with the ViolationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationDescription

`func (o *HosViolation) SetViolationDescription(v string)`

SetViolationDescription sets ViolationDescription field to given value.

### HasViolationDescription

`func (o *HosViolation) HasViolationDescription() bool`

HasViolationDescription returns a boolean if a field has been set.

### SetViolationDescriptionNil

`func (o *HosViolation) SetViolationDescriptionNil(b bool)`

 SetViolationDescriptionNil sets the value for ViolationDescription to be an explicit nil

### UnsetViolationDescription
`func (o *HosViolation) UnsetViolationDescription()`

UnsetViolationDescription ensures that no value is present for ViolationDescription, not even an explicit nil
### GetStartTime

`func (o *HosViolation) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HosViolation) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HosViolation) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HosViolation) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *HosViolation) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *HosViolation) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *HosViolation) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *HosViolation) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *HosViolation) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *HosViolation) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *HosViolation) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *HosViolation) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetDuration

`func (o *HosViolation) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *HosViolation) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *HosViolation) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *HosViolation) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *HosViolation) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *HosViolation) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetHoursLimit

`func (o *HosViolation) GetHoursLimit() float32`

GetHoursLimit returns the HoursLimit field if non-nil, zero value otherwise.

### GetHoursLimitOk

`func (o *HosViolation) GetHoursLimitOk() (*float32, bool)`

GetHoursLimitOk returns a tuple with the HoursLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursLimit

`func (o *HosViolation) SetHoursLimit(v float32)`

SetHoursLimit sets HoursLimit field to given value.

### HasHoursLimit

`func (o *HosViolation) HasHoursLimit() bool`

HasHoursLimit returns a boolean if a field has been set.

### SetHoursLimitNil

`func (o *HosViolation) SetHoursLimitNil(b bool)`

 SetHoursLimitNil sets the value for HoursLimit to be an explicit nil

### UnsetHoursLimit
`func (o *HosViolation) UnsetHoursLimit()`

UnsetHoursLimit ensures that no value is present for HoursLimit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


