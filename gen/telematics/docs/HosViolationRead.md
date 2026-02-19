# HosViolationRead

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
**HosLogId** | Pointer to **NullableString** |  | [optional] 
**SourceHosLogId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewHosViolationRead

`func NewHosViolationRead(fleetId NullableString, id string, createdAt time.Time, updatedAt time.Time, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *HosViolationRead`

NewHosViolationRead instantiates a new HosViolationRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHosViolationReadWithDefaults

`func NewHosViolationReadWithDefaults() *HosViolationRead`

NewHosViolationReadWithDefaults instantiates a new HosViolationRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *HosViolationRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *HosViolationRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *HosViolationRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### SetFleetIdNil

`func (o *HosViolationRead) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *HosViolationRead) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetFleetRef

`func (o *HosViolationRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *HosViolationRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *HosViolationRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *HosViolationRead) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *HosViolationRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *HosViolationRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetId

`func (o *HosViolationRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HosViolationRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HosViolationRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *HosViolationRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HosViolationRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HosViolationRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *HosViolationRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HosViolationRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HosViolationRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *HosViolationRead) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *HosViolationRead) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *HosViolationRead) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *HosViolationRead) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *HosViolationRead) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *HosViolationRead) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetConnectionId

`func (o *HosViolationRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *HosViolationRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *HosViolationRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetTspId

`func (o *HosViolationRead) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *HosViolationRead) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *HosViolationRead) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *HosViolationRead) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *HosViolationRead) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *HosViolationRead) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *HosViolationRead) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *HosViolationRead) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *HosViolationRead) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *HosViolationRead) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *HosViolationRead) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *HosViolationRead) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *HosViolationRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *HosViolationRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *HosViolationRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *HosViolationRead) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *HosViolationRead) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *HosViolationRead) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *HosViolationRead) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *HosViolationRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *HosViolationRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *HosViolationRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *HosViolationRead) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *HosViolationRead) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *HosViolationRead) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *HosViolationRead) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *HosViolationRead) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *HosViolationRead) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *HosViolationRead) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *HosViolationRead) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *HosViolationRead) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *HosViolationRead) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *HosViolationRead) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *HosViolationRead) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *HosViolationRead) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *HosViolationRead) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *HosViolationRead) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *HosViolationRead) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *HosViolationRead) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *HosViolationRead) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *HosViolationRead) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *HosViolationRead) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *HosViolationRead) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *HosViolationRead) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *HosViolationRead) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *HosViolationRead) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *HosViolationRead) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *HosViolationRead) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *HosViolationRead) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetHosEventId

`func (o *HosViolationRead) GetHosEventId() string`

GetHosEventId returns the HosEventId field if non-nil, zero value otherwise.

### GetHosEventIdOk

`func (o *HosViolationRead) GetHosEventIdOk() (*string, bool)`

GetHosEventIdOk returns a tuple with the HosEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosEventId

`func (o *HosViolationRead) SetHosEventId(v string)`

SetHosEventId sets HosEventId field to given value.

### HasHosEventId

`func (o *HosViolationRead) HasHosEventId() bool`

HasHosEventId returns a boolean if a field has been set.

### SetHosEventIdNil

`func (o *HosViolationRead) SetHosEventIdNil(b bool)`

 SetHosEventIdNil sets the value for HosEventId to be an explicit nil

### UnsetHosEventId
`func (o *HosViolationRead) UnsetHosEventId()`

UnsetHosEventId ensures that no value is present for HosEventId, not even an explicit nil
### GetDriverId

`func (o *HosViolationRead) GetDriverId() string`

GetDriverId returns the DriverId field if non-nil, zero value otherwise.

### GetDriverIdOk

`func (o *HosViolationRead) GetDriverIdOk() (*string, bool)`

GetDriverIdOk returns a tuple with the DriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverId

`func (o *HosViolationRead) SetDriverId(v string)`

SetDriverId sets DriverId field to given value.

### HasDriverId

`func (o *HosViolationRead) HasDriverId() bool`

HasDriverId returns a boolean if a field has been set.

### SetDriverIdNil

`func (o *HosViolationRead) SetDriverIdNil(b bool)`

 SetDriverIdNil sets the value for DriverId to be an explicit nil

### UnsetDriverId
`func (o *HosViolationRead) UnsetDriverId()`

UnsetDriverId ensures that no value is present for DriverId, not even an explicit nil
### GetSourceDriverId

`func (o *HosViolationRead) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *HosViolationRead) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *HosViolationRead) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *HosViolationRead) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *HosViolationRead) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *HosViolationRead) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSourceHosEventId

`func (o *HosViolationRead) GetSourceHosEventId() string`

GetSourceHosEventId returns the SourceHosEventId field if non-nil, zero value otherwise.

### GetSourceHosEventIdOk

`func (o *HosViolationRead) GetSourceHosEventIdOk() (*string, bool)`

GetSourceHosEventIdOk returns a tuple with the SourceHosEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHosEventId

`func (o *HosViolationRead) SetSourceHosEventId(v string)`

SetSourceHosEventId sets SourceHosEventId field to given value.

### HasSourceHosEventId

`func (o *HosViolationRead) HasSourceHosEventId() bool`

HasSourceHosEventId returns a boolean if a field has been set.

### SetSourceHosEventIdNil

`func (o *HosViolationRead) SetSourceHosEventIdNil(b bool)`

 SetSourceHosEventIdNil sets the value for SourceHosEventId to be an explicit nil

### UnsetSourceHosEventId
`func (o *HosViolationRead) UnsetSourceHosEventId()`

UnsetSourceHosEventId ensures that no value is present for SourceHosEventId, not even an explicit nil
### GetViolationCode

`func (o *HosViolationRead) GetViolationCode() HosViolationCodeEnum`

GetViolationCode returns the ViolationCode field if non-nil, zero value otherwise.

### GetViolationCodeOk

`func (o *HosViolationRead) GetViolationCodeOk() (*HosViolationCodeEnum, bool)`

GetViolationCodeOk returns a tuple with the ViolationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCode

`func (o *HosViolationRead) SetViolationCode(v HosViolationCodeEnum)`

SetViolationCode sets ViolationCode field to given value.

### HasViolationCode

`func (o *HosViolationRead) HasViolationCode() bool`

HasViolationCode returns a boolean if a field has been set.

### SetViolationCodeNil

`func (o *HosViolationRead) SetViolationCodeNil(b bool)`

 SetViolationCodeNil sets the value for ViolationCode to be an explicit nil

### UnsetViolationCode
`func (o *HosViolationRead) UnsetViolationCode()`

UnsetViolationCode ensures that no value is present for ViolationCode, not even an explicit nil
### GetViolationCategory

`func (o *HosViolationRead) GetViolationCategory() HosViolationCategoryEnum`

GetViolationCategory returns the ViolationCategory field if non-nil, zero value otherwise.

### GetViolationCategoryOk

`func (o *HosViolationRead) GetViolationCategoryOk() (*HosViolationCategoryEnum, bool)`

GetViolationCategoryOk returns a tuple with the ViolationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCategory

`func (o *HosViolationRead) SetViolationCategory(v HosViolationCategoryEnum)`

SetViolationCategory sets ViolationCategory field to given value.

### HasViolationCategory

`func (o *HosViolationRead) HasViolationCategory() bool`

HasViolationCategory returns a boolean if a field has been set.

### SetViolationCategoryNil

`func (o *HosViolationRead) SetViolationCategoryNil(b bool)`

 SetViolationCategoryNil sets the value for ViolationCategory to be an explicit nil

### UnsetViolationCategory
`func (o *HosViolationRead) UnsetViolationCategory()`

UnsetViolationCategory ensures that no value is present for ViolationCategory, not even an explicit nil
### GetViolationDescription

`func (o *HosViolationRead) GetViolationDescription() string`

GetViolationDescription returns the ViolationDescription field if non-nil, zero value otherwise.

### GetViolationDescriptionOk

`func (o *HosViolationRead) GetViolationDescriptionOk() (*string, bool)`

GetViolationDescriptionOk returns a tuple with the ViolationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationDescription

`func (o *HosViolationRead) SetViolationDescription(v string)`

SetViolationDescription sets ViolationDescription field to given value.

### HasViolationDescription

`func (o *HosViolationRead) HasViolationDescription() bool`

HasViolationDescription returns a boolean if a field has been set.

### SetViolationDescriptionNil

`func (o *HosViolationRead) SetViolationDescriptionNil(b bool)`

 SetViolationDescriptionNil sets the value for ViolationDescription to be an explicit nil

### UnsetViolationDescription
`func (o *HosViolationRead) UnsetViolationDescription()`

UnsetViolationDescription ensures that no value is present for ViolationDescription, not even an explicit nil
### GetStartTime

`func (o *HosViolationRead) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HosViolationRead) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HosViolationRead) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HosViolationRead) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### SetStartTimeNil

`func (o *HosViolationRead) SetStartTimeNil(b bool)`

 SetStartTimeNil sets the value for StartTime to be an explicit nil

### UnsetStartTime
`func (o *HosViolationRead) UnsetStartTime()`

UnsetStartTime ensures that no value is present for StartTime, not even an explicit nil
### GetEndTime

`func (o *HosViolationRead) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *HosViolationRead) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *HosViolationRead) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *HosViolationRead) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### SetEndTimeNil

`func (o *HosViolationRead) SetEndTimeNil(b bool)`

 SetEndTimeNil sets the value for EndTime to be an explicit nil

### UnsetEndTime
`func (o *HosViolationRead) UnsetEndTime()`

UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil
### GetDuration

`func (o *HosViolationRead) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *HosViolationRead) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *HosViolationRead) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *HosViolationRead) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *HosViolationRead) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *HosViolationRead) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetHoursLimit

`func (o *HosViolationRead) GetHoursLimit() float32`

GetHoursLimit returns the HoursLimit field if non-nil, zero value otherwise.

### GetHoursLimitOk

`func (o *HosViolationRead) GetHoursLimitOk() (*float32, bool)`

GetHoursLimitOk returns a tuple with the HoursLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursLimit

`func (o *HosViolationRead) SetHoursLimit(v float32)`

SetHoursLimit sets HoursLimit field to given value.

### HasHoursLimit

`func (o *HosViolationRead) HasHoursLimit() bool`

HasHoursLimit returns a boolean if a field has been set.

### SetHoursLimitNil

`func (o *HosViolationRead) SetHoursLimitNil(b bool)`

 SetHoursLimitNil sets the value for HoursLimit to be an explicit nil

### UnsetHoursLimit
`func (o *HosViolationRead) UnsetHoursLimit()`

UnsetHoursLimit ensures that no value is present for HoursLimit, not even an explicit nil
### GetHosLogId

`func (o *HosViolationRead) GetHosLogId() string`

GetHosLogId returns the HosLogId field if non-nil, zero value otherwise.

### GetHosLogIdOk

`func (o *HosViolationRead) GetHosLogIdOk() (*string, bool)`

GetHosLogIdOk returns a tuple with the HosLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosLogId

`func (o *HosViolationRead) SetHosLogId(v string)`

SetHosLogId sets HosLogId field to given value.

### HasHosLogId

`func (o *HosViolationRead) HasHosLogId() bool`

HasHosLogId returns a boolean if a field has been set.

### SetHosLogIdNil

`func (o *HosViolationRead) SetHosLogIdNil(b bool)`

 SetHosLogIdNil sets the value for HosLogId to be an explicit nil

### UnsetHosLogId
`func (o *HosViolationRead) UnsetHosLogId()`

UnsetHosLogId ensures that no value is present for HosLogId, not even an explicit nil
### GetSourceHosLogId

`func (o *HosViolationRead) GetSourceHosLogId() string`

GetSourceHosLogId returns the SourceHosLogId field if non-nil, zero value otherwise.

### GetSourceHosLogIdOk

`func (o *HosViolationRead) GetSourceHosLogIdOk() (*string, bool)`

GetSourceHosLogIdOk returns a tuple with the SourceHosLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHosLogId

`func (o *HosViolationRead) SetSourceHosLogId(v string)`

SetSourceHosLogId sets SourceHosLogId field to given value.

### HasSourceHosLogId

`func (o *HosViolationRead) HasSourceHosLogId() bool`

HasSourceHosLogId returns a boolean if a field has been set.

### SetSourceHosLogIdNil

`func (o *HosViolationRead) SetSourceHosLogIdNil(b bool)`

 SetSourceHosLogIdNil sets the value for SourceHosLogId to be an explicit nil

### UnsetSourceHosLogId
`func (o *HosViolationRead) UnsetSourceHosLogId()`

UnsetSourceHosLogId ensures that no value is present for SourceHosLogId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


