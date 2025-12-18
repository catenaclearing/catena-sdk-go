# DvirLogDefect

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

### NewDvirLogDefect

`func NewDvirLogDefect(id string, createdAt time.Time, updatedAt time.Time, fleetId string, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *DvirLogDefect`

NewDvirLogDefect instantiates a new DvirLogDefect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDvirLogDefectWithDefaults

`func NewDvirLogDefectWithDefaults() *DvirLogDefect`

NewDvirLogDefectWithDefaults instantiates a new DvirLogDefect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DvirLogDefect) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DvirLogDefect) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DvirLogDefect) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DvirLogDefect) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DvirLogDefect) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DvirLogDefect) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DvirLogDefect) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DvirLogDefect) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DvirLogDefect) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *DvirLogDefect) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DvirLogDefect) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DvirLogDefect) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DvirLogDefect) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *DvirLogDefect) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *DvirLogDefect) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetFleetId

`func (o *DvirLogDefect) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *DvirLogDefect) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *DvirLogDefect) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetConnectionId

`func (o *DvirLogDefect) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DvirLogDefect) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DvirLogDefect) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceName

`func (o *DvirLogDefect) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *DvirLogDefect) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *DvirLogDefect) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *DvirLogDefect) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *DvirLogDefect) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *DvirLogDefect) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *DvirLogDefect) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *DvirLogDefect) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DvirLogDefect) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DvirLogDefect) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *DvirLogDefect) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *DvirLogDefect) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *DvirLogDefect) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *DvirLogDefect) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *DvirLogDefect) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *DvirLogDefect) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *DvirLogDefect) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *DvirLogDefect) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *DvirLogDefect) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *DvirLogDefect) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *DvirLogDefect) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *DvirLogDefect) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *DvirLogDefect) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *DvirLogDefect) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *DvirLogDefect) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *DvirLogDefect) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *DvirLogDefect) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *DvirLogDefect) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *DvirLogDefect) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *DvirLogDefect) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *DvirLogDefect) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetDvirLogId

`func (o *DvirLogDefect) GetDvirLogId() string`

GetDvirLogId returns the DvirLogId field if non-nil, zero value otherwise.

### GetDvirLogIdOk

`func (o *DvirLogDefect) GetDvirLogIdOk() (*string, bool)`

GetDvirLogIdOk returns a tuple with the DvirLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvirLogId

`func (o *DvirLogDefect) SetDvirLogId(v string)`

SetDvirLogId sets DvirLogId field to given value.

### HasDvirLogId

`func (o *DvirLogDefect) HasDvirLogId() bool`

HasDvirLogId returns a boolean if a field has been set.

### SetDvirLogIdNil

`func (o *DvirLogDefect) SetDvirLogIdNil(b bool)`

 SetDvirLogIdNil sets the value for DvirLogId to be an explicit nil

### UnsetDvirLogId
`func (o *DvirLogDefect) UnsetDvirLogId()`

UnsetDvirLogId ensures that no value is present for DvirLogId, not even an explicit nil
### GetSourceDvirLogId

`func (o *DvirLogDefect) GetSourceDvirLogId() string`

GetSourceDvirLogId returns the SourceDvirLogId field if non-nil, zero value otherwise.

### GetSourceDvirLogIdOk

`func (o *DvirLogDefect) GetSourceDvirLogIdOk() (*string, bool)`

GetSourceDvirLogIdOk returns a tuple with the SourceDvirLogId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDvirLogId

`func (o *DvirLogDefect) SetSourceDvirLogId(v string)`

SetSourceDvirLogId sets SourceDvirLogId field to given value.

### HasSourceDvirLogId

`func (o *DvirLogDefect) HasSourceDvirLogId() bool`

HasSourceDvirLogId returns a boolean if a field has been set.

### SetSourceDvirLogIdNil

`func (o *DvirLogDefect) SetSourceDvirLogIdNil(b bool)`

 SetSourceDvirLogIdNil sets the value for SourceDvirLogId to be an explicit nil

### UnsetSourceDvirLogId
`func (o *DvirLogDefect) UnsetSourceDvirLogId()`

UnsetSourceDvirLogId ensures that no value is present for SourceDvirLogId, not even an explicit nil
### GetDefectType

`func (o *DvirLogDefect) GetDefectType() string`

GetDefectType returns the DefectType field if non-nil, zero value otherwise.

### GetDefectTypeOk

`func (o *DvirLogDefect) GetDefectTypeOk() (*string, bool)`

GetDefectTypeOk returns a tuple with the DefectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectType

`func (o *DvirLogDefect) SetDefectType(v string)`

SetDefectType sets DefectType field to given value.

### HasDefectType

`func (o *DvirLogDefect) HasDefectType() bool`

HasDefectType returns a boolean if a field has been set.

### SetDefectTypeNil

`func (o *DvirLogDefect) SetDefectTypeNil(b bool)`

 SetDefectTypeNil sets the value for DefectType to be an explicit nil

### UnsetDefectType
`func (o *DvirLogDefect) UnsetDefectType()`

UnsetDefectType ensures that no value is present for DefectType, not even an explicit nil
### GetDefectCode

`func (o *DvirLogDefect) GetDefectCode() string`

GetDefectCode returns the DefectCode field if non-nil, zero value otherwise.

### GetDefectCodeOk

`func (o *DvirLogDefect) GetDefectCodeOk() (*string, bool)`

GetDefectCodeOk returns a tuple with the DefectCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectCode

`func (o *DvirLogDefect) SetDefectCode(v string)`

SetDefectCode sets DefectCode field to given value.

### HasDefectCode

`func (o *DvirLogDefect) HasDefectCode() bool`

HasDefectCode returns a boolean if a field has been set.

### SetDefectCodeNil

`func (o *DvirLogDefect) SetDefectCodeNil(b bool)`

 SetDefectCodeNil sets the value for DefectCode to be an explicit nil

### UnsetDefectCode
`func (o *DvirLogDefect) UnsetDefectCode()`

UnsetDefectCode ensures that no value is present for DefectCode, not even an explicit nil
### GetDefectName

`func (o *DvirLogDefect) GetDefectName() string`

GetDefectName returns the DefectName field if non-nil, zero value otherwise.

### GetDefectNameOk

`func (o *DvirLogDefect) GetDefectNameOk() (*string, bool)`

GetDefectNameOk returns a tuple with the DefectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectName

`func (o *DvirLogDefect) SetDefectName(v string)`

SetDefectName sets DefectName field to given value.

### HasDefectName

`func (o *DvirLogDefect) HasDefectName() bool`

HasDefectName returns a boolean if a field has been set.

### SetDefectNameNil

`func (o *DvirLogDefect) SetDefectNameNil(b bool)`

 SetDefectNameNil sets the value for DefectName to be an explicit nil

### UnsetDefectName
`func (o *DvirLogDefect) UnsetDefectName()`

UnsetDefectName ensures that no value is present for DefectName, not even an explicit nil
### GetDescription

`func (o *DvirLogDefect) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DvirLogDefect) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DvirLogDefect) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DvirLogDefect) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DvirLogDefect) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DvirLogDefect) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *DvirLogDefect) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DvirLogDefect) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DvirLogDefect) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DvirLogDefect) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *DvirLogDefect) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DvirLogDefect) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIsSafetyCritical

`func (o *DvirLogDefect) GetIsSafetyCritical() bool`

GetIsSafetyCritical returns the IsSafetyCritical field if non-nil, zero value otherwise.

### GetIsSafetyCriticalOk

`func (o *DvirLogDefect) GetIsSafetyCriticalOk() (*bool, bool)`

GetIsSafetyCriticalOk returns a tuple with the IsSafetyCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafetyCritical

`func (o *DvirLogDefect) SetIsSafetyCritical(v bool)`

SetIsSafetyCritical sets IsSafetyCritical field to given value.

### HasIsSafetyCritical

`func (o *DvirLogDefect) HasIsSafetyCritical() bool`

HasIsSafetyCritical returns a boolean if a field has been set.

### SetIsSafetyCriticalNil

`func (o *DvirLogDefect) SetIsSafetyCriticalNil(b bool)`

 SetIsSafetyCriticalNil sets the value for IsSafetyCritical to be an explicit nil

### UnsetIsSafetyCritical
`func (o *DvirLogDefect) UnsetIsSafetyCritical()`

UnsetIsSafetyCritical ensures that no value is present for IsSafetyCritical, not even an explicit nil
### GetResolvedAt

`func (o *DvirLogDefect) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *DvirLogDefect) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *DvirLogDefect) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *DvirLogDefect) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### SetResolvedAtNil

`func (o *DvirLogDefect) SetResolvedAtNil(b bool)`

 SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil

### UnsetResolvedAt
`func (o *DvirLogDefect) UnsetResolvedAt()`

UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
### GetNotes

`func (o *DvirLogDefect) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *DvirLogDefect) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *DvirLogDefect) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *DvirLogDefect) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *DvirLogDefect) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *DvirLogDefect) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


