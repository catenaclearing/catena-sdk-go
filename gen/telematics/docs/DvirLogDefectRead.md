# DvirLogDefectRead

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
**DvirLogId** | Pointer to **NullableString** |  | [optional] 
**SourceDvirLogId** | Pointer to **NullableString** |  | [optional] 
**DefectType** | Pointer to **NullableString** |  | [optional] 
**DefectCode** | Pointer to **NullableString** |  | [optional] 
**DefectName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**IsSafetyCritical** | Pointer to **NullableBool** |  | [optional] 
**ResolvedAt** | Pointer to **NullableTime** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDvirLogDefectRead

`func NewDvirLogDefectRead(fleetId NullableString, id string, createdAt time.Time, updatedAt time.Time, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *DvirLogDefectRead`

NewDvirLogDefectRead instantiates a new DvirLogDefectRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDvirLogDefectReadWithDefaults

`func NewDvirLogDefectReadWithDefaults() *DvirLogDefectRead`

NewDvirLogDefectReadWithDefaults instantiates a new DvirLogDefectRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *DvirLogDefectRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *DvirLogDefectRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *DvirLogDefectRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### SetFleetIdNil

`func (o *DvirLogDefectRead) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *DvirLogDefectRead) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetFleetRef

`func (o *DvirLogDefectRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *DvirLogDefectRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *DvirLogDefectRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *DvirLogDefectRead) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *DvirLogDefectRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *DvirLogDefectRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetId

`func (o *DvirLogDefectRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DvirLogDefectRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DvirLogDefectRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DvirLogDefectRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DvirLogDefectRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DvirLogDefectRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DvirLogDefectRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DvirLogDefectRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DvirLogDefectRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *DvirLogDefectRead) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DvirLogDefectRead) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DvirLogDefectRead) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DvirLogDefectRead) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *DvirLogDefectRead) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *DvirLogDefectRead) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetConnectionId

`func (o *DvirLogDefectRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DvirLogDefectRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DvirLogDefectRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetTspId

`func (o *DvirLogDefectRead) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *DvirLogDefectRead) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *DvirLogDefectRead) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *DvirLogDefectRead) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *DvirLogDefectRead) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *DvirLogDefectRead) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *DvirLogDefectRead) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *DvirLogDefectRead) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *DvirLogDefectRead) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *DvirLogDefectRead) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *DvirLogDefectRead) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *DvirLogDefectRead) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *DvirLogDefectRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *DvirLogDefectRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *DvirLogDefectRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *DvirLogDefectRead) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *DvirLogDefectRead) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *DvirLogDefectRead) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *DvirLogDefectRead) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *DvirLogDefectRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DvirLogDefectRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DvirLogDefectRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *DvirLogDefectRead) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *DvirLogDefectRead) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *DvirLogDefectRead) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *DvirLogDefectRead) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *DvirLogDefectRead) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *DvirLogDefectRead) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *DvirLogDefectRead) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *DvirLogDefectRead) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *DvirLogDefectRead) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *DvirLogDefectRead) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *DvirLogDefectRead) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *DvirLogDefectRead) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *DvirLogDefectRead) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *DvirLogDefectRead) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *DvirLogDefectRead) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *DvirLogDefectRead) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *DvirLogDefectRead) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *DvirLogDefectRead) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *DvirLogDefectRead) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *DvirLogDefectRead) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *DvirLogDefectRead) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *DvirLogDefectRead) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *DvirLogDefectRead) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *DvirLogDefectRead) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *DvirLogDefectRead) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *DvirLogDefectRead) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *DvirLogDefectRead) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetDvirLogId

`func (o *DvirLogDefectRead) GetDvirLogId() string`

GetDvirLogId returns the DvirLogId field if non-nil, zero value otherwise.

### GetDvirLogIdOk

`func (o *DvirLogDefectRead) GetDvirLogIdOk() (*string, bool)`

GetDvirLogIdOk returns a tuple with the DvirLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvirLogId

`func (o *DvirLogDefectRead) SetDvirLogId(v string)`

SetDvirLogId sets DvirLogId field to given value.

### HasDvirLogId

`func (o *DvirLogDefectRead) HasDvirLogId() bool`

HasDvirLogId returns a boolean if a field has been set.

### SetDvirLogIdNil

`func (o *DvirLogDefectRead) SetDvirLogIdNil(b bool)`

 SetDvirLogIdNil sets the value for DvirLogId to be an explicit nil

### UnsetDvirLogId
`func (o *DvirLogDefectRead) UnsetDvirLogId()`

UnsetDvirLogId ensures that no value is present for DvirLogId, not even an explicit nil
### GetSourceDvirLogId

`func (o *DvirLogDefectRead) GetSourceDvirLogId() string`

GetSourceDvirLogId returns the SourceDvirLogId field if non-nil, zero value otherwise.

### GetSourceDvirLogIdOk

`func (o *DvirLogDefectRead) GetSourceDvirLogIdOk() (*string, bool)`

GetSourceDvirLogIdOk returns a tuple with the SourceDvirLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDvirLogId

`func (o *DvirLogDefectRead) SetSourceDvirLogId(v string)`

SetSourceDvirLogId sets SourceDvirLogId field to given value.

### HasSourceDvirLogId

`func (o *DvirLogDefectRead) HasSourceDvirLogId() bool`

HasSourceDvirLogId returns a boolean if a field has been set.

### SetSourceDvirLogIdNil

`func (o *DvirLogDefectRead) SetSourceDvirLogIdNil(b bool)`

 SetSourceDvirLogIdNil sets the value for SourceDvirLogId to be an explicit nil

### UnsetSourceDvirLogId
`func (o *DvirLogDefectRead) UnsetSourceDvirLogId()`

UnsetSourceDvirLogId ensures that no value is present for SourceDvirLogId, not even an explicit nil
### GetDefectType

`func (o *DvirLogDefectRead) GetDefectType() string`

GetDefectType returns the DefectType field if non-nil, zero value otherwise.

### GetDefectTypeOk

`func (o *DvirLogDefectRead) GetDefectTypeOk() (*string, bool)`

GetDefectTypeOk returns a tuple with the DefectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectType

`func (o *DvirLogDefectRead) SetDefectType(v string)`

SetDefectType sets DefectType field to given value.

### HasDefectType

`func (o *DvirLogDefectRead) HasDefectType() bool`

HasDefectType returns a boolean if a field has been set.

### SetDefectTypeNil

`func (o *DvirLogDefectRead) SetDefectTypeNil(b bool)`

 SetDefectTypeNil sets the value for DefectType to be an explicit nil

### UnsetDefectType
`func (o *DvirLogDefectRead) UnsetDefectType()`

UnsetDefectType ensures that no value is present for DefectType, not even an explicit nil
### GetDefectCode

`func (o *DvirLogDefectRead) GetDefectCode() string`

GetDefectCode returns the DefectCode field if non-nil, zero value otherwise.

### GetDefectCodeOk

`func (o *DvirLogDefectRead) GetDefectCodeOk() (*string, bool)`

GetDefectCodeOk returns a tuple with the DefectCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectCode

`func (o *DvirLogDefectRead) SetDefectCode(v string)`

SetDefectCode sets DefectCode field to given value.

### HasDefectCode

`func (o *DvirLogDefectRead) HasDefectCode() bool`

HasDefectCode returns a boolean if a field has been set.

### SetDefectCodeNil

`func (o *DvirLogDefectRead) SetDefectCodeNil(b bool)`

 SetDefectCodeNil sets the value for DefectCode to be an explicit nil

### UnsetDefectCode
`func (o *DvirLogDefectRead) UnsetDefectCode()`

UnsetDefectCode ensures that no value is present for DefectCode, not even an explicit nil
### GetDefectName

`func (o *DvirLogDefectRead) GetDefectName() string`

GetDefectName returns the DefectName field if non-nil, zero value otherwise.

### GetDefectNameOk

`func (o *DvirLogDefectRead) GetDefectNameOk() (*string, bool)`

GetDefectNameOk returns a tuple with the DefectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectName

`func (o *DvirLogDefectRead) SetDefectName(v string)`

SetDefectName sets DefectName field to given value.

### HasDefectName

`func (o *DvirLogDefectRead) HasDefectName() bool`

HasDefectName returns a boolean if a field has been set.

### SetDefectNameNil

`func (o *DvirLogDefectRead) SetDefectNameNil(b bool)`

 SetDefectNameNil sets the value for DefectName to be an explicit nil

### UnsetDefectName
`func (o *DvirLogDefectRead) UnsetDefectName()`

UnsetDefectName ensures that no value is present for DefectName, not even an explicit nil
### GetDescription

`func (o *DvirLogDefectRead) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DvirLogDefectRead) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DvirLogDefectRead) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DvirLogDefectRead) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DvirLogDefectRead) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DvirLogDefectRead) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *DvirLogDefectRead) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DvirLogDefectRead) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DvirLogDefectRead) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DvirLogDefectRead) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DvirLogDefectRead) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DvirLogDefectRead) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIsSafetyCritical

`func (o *DvirLogDefectRead) GetIsSafetyCritical() bool`

GetIsSafetyCritical returns the IsSafetyCritical field if non-nil, zero value otherwise.

### GetIsSafetyCriticalOk

`func (o *DvirLogDefectRead) GetIsSafetyCriticalOk() (*bool, bool)`

GetIsSafetyCriticalOk returns a tuple with the IsSafetyCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafetyCritical

`func (o *DvirLogDefectRead) SetIsSafetyCritical(v bool)`

SetIsSafetyCritical sets IsSafetyCritical field to given value.

### HasIsSafetyCritical

`func (o *DvirLogDefectRead) HasIsSafetyCritical() bool`

HasIsSafetyCritical returns a boolean if a field has been set.

### SetIsSafetyCriticalNil

`func (o *DvirLogDefectRead) SetIsSafetyCriticalNil(b bool)`

 SetIsSafetyCriticalNil sets the value for IsSafetyCritical to be an explicit nil

### UnsetIsSafetyCritical
`func (o *DvirLogDefectRead) UnsetIsSafetyCritical()`

UnsetIsSafetyCritical ensures that no value is present for IsSafetyCritical, not even an explicit nil
### GetResolvedAt

`func (o *DvirLogDefectRead) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *DvirLogDefectRead) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *DvirLogDefectRead) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *DvirLogDefectRead) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### SetResolvedAtNil

`func (o *DvirLogDefectRead) SetResolvedAtNil(b bool)`

 SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil

### UnsetResolvedAt
`func (o *DvirLogDefectRead) UnsetResolvedAt()`

UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
### GetNotes

`func (o *DvirLogDefectRead) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DvirLogDefectRead) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DvirLogDefectRead) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DvirLogDefectRead) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *DvirLogDefectRead) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *DvirLogDefectRead) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


