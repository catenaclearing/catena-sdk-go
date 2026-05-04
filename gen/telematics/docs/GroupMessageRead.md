# GroupMessageRead

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
**SourceSenderId** | Pointer to **NullableString** |  | [optional] 
**SourceRecipientIds** | Pointer to **[]string** |  | [optional] 
**SenderId** | Pointer to **NullableString** |  | [optional] 
**RecipientIds** | Pointer to **[]string** |  | [optional] 
**MessageText** | Pointer to **NullableString** |  | [optional] 
**MessageHash** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullableMessageStatusEnum**](MessageStatusEnum.md) |  | [optional] 
**DeliveredAt** | Pointer to **NullableTime** |  | [optional] 
**SentByDriver** | Pointer to **NullableBool** |  | [optional] 
**SenderName** | Pointer to **NullableString** |  | [optional] 
**ThreadId** | Pointer to **NullableString** |  | [optional] 
**SourceThreadId** | Pointer to **NullableString** |  | [optional] 
**GroupId** | Pointer to **NullableString** |  | [optional] 
**SourceGroupId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGroupMessageRead

`func NewGroupMessageRead(fleetId NullableString, id string, createdAt time.Time, updatedAt time.Time, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *GroupMessageRead`

NewGroupMessageRead instantiates a new GroupMessageRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMessageReadWithDefaults

`func NewGroupMessageReadWithDefaults() *GroupMessageRead`

NewGroupMessageReadWithDefaults instantiates a new GroupMessageRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *GroupMessageRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *GroupMessageRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *GroupMessageRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### SetFleetIdNil

`func (o *GroupMessageRead) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *GroupMessageRead) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetFleetRef

`func (o *GroupMessageRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *GroupMessageRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *GroupMessageRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *GroupMessageRead) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *GroupMessageRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *GroupMessageRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetId

`func (o *GroupMessageRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupMessageRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupMessageRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GroupMessageRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupMessageRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupMessageRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GroupMessageRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GroupMessageRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GroupMessageRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *GroupMessageRead) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GroupMessageRead) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GroupMessageRead) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *GroupMessageRead) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *GroupMessageRead) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *GroupMessageRead) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetConnectionId

`func (o *GroupMessageRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *GroupMessageRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *GroupMessageRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetTspId

`func (o *GroupMessageRead) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *GroupMessageRead) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *GroupMessageRead) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *GroupMessageRead) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *GroupMessageRead) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *GroupMessageRead) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *GroupMessageRead) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *GroupMessageRead) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *GroupMessageRead) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *GroupMessageRead) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *GroupMessageRead) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *GroupMessageRead) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *GroupMessageRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *GroupMessageRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *GroupMessageRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *GroupMessageRead) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *GroupMessageRead) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *GroupMessageRead) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *GroupMessageRead) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *GroupMessageRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *GroupMessageRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *GroupMessageRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *GroupMessageRead) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *GroupMessageRead) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *GroupMessageRead) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *GroupMessageRead) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *GroupMessageRead) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *GroupMessageRead) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *GroupMessageRead) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *GroupMessageRead) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *GroupMessageRead) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *GroupMessageRead) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *GroupMessageRead) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *GroupMessageRead) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *GroupMessageRead) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *GroupMessageRead) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *GroupMessageRead) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *GroupMessageRead) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *GroupMessageRead) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *GroupMessageRead) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *GroupMessageRead) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *GroupMessageRead) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *GroupMessageRead) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *GroupMessageRead) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *GroupMessageRead) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *GroupMessageRead) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *GroupMessageRead) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *GroupMessageRead) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *GroupMessageRead) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetSourceSenderId

`func (o *GroupMessageRead) GetSourceSenderId() string`

GetSourceSenderId returns the SourceSenderId field if non-nil, zero value otherwise.

### GetSourceSenderIdOk

`func (o *GroupMessageRead) GetSourceSenderIdOk() (*string, bool)`

GetSourceSenderIdOk returns a tuple with the SourceSenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSenderId

`func (o *GroupMessageRead) SetSourceSenderId(v string)`

SetSourceSenderId sets SourceSenderId field to given value.

### HasSourceSenderId

`func (o *GroupMessageRead) HasSourceSenderId() bool`

HasSourceSenderId returns a boolean if a field has been set.

### SetSourceSenderIdNil

`func (o *GroupMessageRead) SetSourceSenderIdNil(b bool)`

 SetSourceSenderIdNil sets the value for SourceSenderId to be an explicit nil

### UnsetSourceSenderId
`func (o *GroupMessageRead) UnsetSourceSenderId()`

UnsetSourceSenderId ensures that no value is present for SourceSenderId, not even an explicit nil
### GetSourceRecipientIds

`func (o *GroupMessageRead) GetSourceRecipientIds() []string`

GetSourceRecipientIds returns the SourceRecipientIds field if non-nil, zero value otherwise.

### GetSourceRecipientIdsOk

`func (o *GroupMessageRead) GetSourceRecipientIdsOk() (*[]string, bool)`

GetSourceRecipientIdsOk returns a tuple with the SourceRecipientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRecipientIds

`func (o *GroupMessageRead) SetSourceRecipientIds(v []string)`

SetSourceRecipientIds sets SourceRecipientIds field to given value.

### HasSourceRecipientIds

`func (o *GroupMessageRead) HasSourceRecipientIds() bool`

HasSourceRecipientIds returns a boolean if a field has been set.

### SetSourceRecipientIdsNil

`func (o *GroupMessageRead) SetSourceRecipientIdsNil(b bool)`

 SetSourceRecipientIdsNil sets the value for SourceRecipientIds to be an explicit nil

### UnsetSourceRecipientIds
`func (o *GroupMessageRead) UnsetSourceRecipientIds()`

UnsetSourceRecipientIds ensures that no value is present for SourceRecipientIds, not even an explicit nil
### GetSenderId

`func (o *GroupMessageRead) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *GroupMessageRead) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *GroupMessageRead) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *GroupMessageRead) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### SetSenderIdNil

`func (o *GroupMessageRead) SetSenderIdNil(b bool)`

 SetSenderIdNil sets the value for SenderId to be an explicit nil

### UnsetSenderId
`func (o *GroupMessageRead) UnsetSenderId()`

UnsetSenderId ensures that no value is present for SenderId, not even an explicit nil
### GetRecipientIds

`func (o *GroupMessageRead) GetRecipientIds() []string`

GetRecipientIds returns the RecipientIds field if non-nil, zero value otherwise.

### GetRecipientIdsOk

`func (o *GroupMessageRead) GetRecipientIdsOk() (*[]string, bool)`

GetRecipientIdsOk returns a tuple with the RecipientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIds

`func (o *GroupMessageRead) SetRecipientIds(v []string)`

SetRecipientIds sets RecipientIds field to given value.

### HasRecipientIds

`func (o *GroupMessageRead) HasRecipientIds() bool`

HasRecipientIds returns a boolean if a field has been set.

### SetRecipientIdsNil

`func (o *GroupMessageRead) SetRecipientIdsNil(b bool)`

 SetRecipientIdsNil sets the value for RecipientIds to be an explicit nil

### UnsetRecipientIds
`func (o *GroupMessageRead) UnsetRecipientIds()`

UnsetRecipientIds ensures that no value is present for RecipientIds, not even an explicit nil
### GetMessageText

`func (o *GroupMessageRead) GetMessageText() string`

GetMessageText returns the MessageText field if non-nil, zero value otherwise.

### GetMessageTextOk

`func (o *GroupMessageRead) GetMessageTextOk() (*string, bool)`

GetMessageTextOk returns a tuple with the MessageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageText

`func (o *GroupMessageRead) SetMessageText(v string)`

SetMessageText sets MessageText field to given value.

### HasMessageText

`func (o *GroupMessageRead) HasMessageText() bool`

HasMessageText returns a boolean if a field has been set.

### SetMessageTextNil

`func (o *GroupMessageRead) SetMessageTextNil(b bool)`

 SetMessageTextNil sets the value for MessageText to be an explicit nil

### UnsetMessageText
`func (o *GroupMessageRead) UnsetMessageText()`

UnsetMessageText ensures that no value is present for MessageText, not even an explicit nil
### GetMessageHash

`func (o *GroupMessageRead) GetMessageHash() string`

GetMessageHash returns the MessageHash field if non-nil, zero value otherwise.

### GetMessageHashOk

`func (o *GroupMessageRead) GetMessageHashOk() (*string, bool)`

GetMessageHashOk returns a tuple with the MessageHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageHash

`func (o *GroupMessageRead) SetMessageHash(v string)`

SetMessageHash sets MessageHash field to given value.

### HasMessageHash

`func (o *GroupMessageRead) HasMessageHash() bool`

HasMessageHash returns a boolean if a field has been set.

### SetMessageHashNil

`func (o *GroupMessageRead) SetMessageHashNil(b bool)`

 SetMessageHashNil sets the value for MessageHash to be an explicit nil

### UnsetMessageHash
`func (o *GroupMessageRead) UnsetMessageHash()`

UnsetMessageHash ensures that no value is present for MessageHash, not even an explicit nil
### GetStatus

`func (o *GroupMessageRead) GetStatus() MessageStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GroupMessageRead) GetStatusOk() (*MessageStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GroupMessageRead) SetStatus(v MessageStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GroupMessageRead) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GroupMessageRead) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GroupMessageRead) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDeliveredAt

`func (o *GroupMessageRead) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *GroupMessageRead) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *GroupMessageRead) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *GroupMessageRead) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### SetDeliveredAtNil

`func (o *GroupMessageRead) SetDeliveredAtNil(b bool)`

 SetDeliveredAtNil sets the value for DeliveredAt to be an explicit nil

### UnsetDeliveredAt
`func (o *GroupMessageRead) UnsetDeliveredAt()`

UnsetDeliveredAt ensures that no value is present for DeliveredAt, not even an explicit nil
### GetSentByDriver

`func (o *GroupMessageRead) GetSentByDriver() bool`

GetSentByDriver returns the SentByDriver field if non-nil, zero value otherwise.

### GetSentByDriverOk

`func (o *GroupMessageRead) GetSentByDriverOk() (*bool, bool)`

GetSentByDriverOk returns a tuple with the SentByDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentByDriver

`func (o *GroupMessageRead) SetSentByDriver(v bool)`

SetSentByDriver sets SentByDriver field to given value.

### HasSentByDriver

`func (o *GroupMessageRead) HasSentByDriver() bool`

HasSentByDriver returns a boolean if a field has been set.

### SetSentByDriverNil

`func (o *GroupMessageRead) SetSentByDriverNil(b bool)`

 SetSentByDriverNil sets the value for SentByDriver to be an explicit nil

### UnsetSentByDriver
`func (o *GroupMessageRead) UnsetSentByDriver()`

UnsetSentByDriver ensures that no value is present for SentByDriver, not even an explicit nil
### GetSenderName

`func (o *GroupMessageRead) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *GroupMessageRead) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *GroupMessageRead) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *GroupMessageRead) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### SetSenderNameNil

`func (o *GroupMessageRead) SetSenderNameNil(b bool)`

 SetSenderNameNil sets the value for SenderName to be an explicit nil

### UnsetSenderName
`func (o *GroupMessageRead) UnsetSenderName()`

UnsetSenderName ensures that no value is present for SenderName, not even an explicit nil
### GetThreadId

`func (o *GroupMessageRead) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *GroupMessageRead) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *GroupMessageRead) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *GroupMessageRead) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### SetThreadIdNil

`func (o *GroupMessageRead) SetThreadIdNil(b bool)`

 SetThreadIdNil sets the value for ThreadId to be an explicit nil

### UnsetThreadId
`func (o *GroupMessageRead) UnsetThreadId()`

UnsetThreadId ensures that no value is present for ThreadId, not even an explicit nil
### GetSourceThreadId

`func (o *GroupMessageRead) GetSourceThreadId() string`

GetSourceThreadId returns the SourceThreadId field if non-nil, zero value otherwise.

### GetSourceThreadIdOk

`func (o *GroupMessageRead) GetSourceThreadIdOk() (*string, bool)`

GetSourceThreadIdOk returns a tuple with the SourceThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceThreadId

`func (o *GroupMessageRead) SetSourceThreadId(v string)`

SetSourceThreadId sets SourceThreadId field to given value.

### HasSourceThreadId

`func (o *GroupMessageRead) HasSourceThreadId() bool`

HasSourceThreadId returns a boolean if a field has been set.

### SetSourceThreadIdNil

`func (o *GroupMessageRead) SetSourceThreadIdNil(b bool)`

 SetSourceThreadIdNil sets the value for SourceThreadId to be an explicit nil

### UnsetSourceThreadId
`func (o *GroupMessageRead) UnsetSourceThreadId()`

UnsetSourceThreadId ensures that no value is present for SourceThreadId, not even an explicit nil
### GetGroupId

`func (o *GroupMessageRead) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupMessageRead) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupMessageRead) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *GroupMessageRead) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *GroupMessageRead) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *GroupMessageRead) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetSourceGroupId

`func (o *GroupMessageRead) GetSourceGroupId() string`

GetSourceGroupId returns the SourceGroupId field if non-nil, zero value otherwise.

### GetSourceGroupIdOk

`func (o *GroupMessageRead) GetSourceGroupIdOk() (*string, bool)`

GetSourceGroupIdOk returns a tuple with the SourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroupId

`func (o *GroupMessageRead) SetSourceGroupId(v string)`

SetSourceGroupId sets SourceGroupId field to given value.

### HasSourceGroupId

`func (o *GroupMessageRead) HasSourceGroupId() bool`

HasSourceGroupId returns a boolean if a field has been set.

### SetSourceGroupIdNil

`func (o *GroupMessageRead) SetSourceGroupIdNil(b bool)`

 SetSourceGroupIdNil sets the value for SourceGroupId to be an explicit nil

### UnsetSourceGroupId
`func (o *GroupMessageRead) UnsetSourceGroupId()`

UnsetSourceGroupId ensures that no value is present for SourceGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


