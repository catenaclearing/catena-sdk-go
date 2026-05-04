# MessageRead

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
**SourceRecipientId** | Pointer to **NullableString** |  | [optional] 
**SenderId** | Pointer to **NullableString** |  | [optional] 
**RecipientId** | Pointer to **NullableString** |  | [optional] 
**MessageText** | Pointer to **NullableString** |  | [optional] 
**MessageHash** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullableMessageStatusEnum**](MessageStatusEnum.md) |  | [optional] 
**DeliveredAt** | Pointer to **NullableTime** |  | [optional] 
**ReadAt** | Pointer to **NullableTime** |  | [optional] 
**SentByDriver** | Pointer to **NullableBool** |  | [optional] 
**SenderName** | Pointer to **NullableString** |  | [optional] 
**RecipientName** | Pointer to **NullableString** |  | [optional] 
**ThreadId** | Pointer to **NullableString** |  | [optional] 
**SourceThreadId** | Pointer to **NullableString** |  | [optional] 
**ReplyToMessageId** | Pointer to **NullableString** |  | [optional] 
**SourceReplyToMessageId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMessageRead

`func NewMessageRead(fleetId NullableString, id string, createdAt time.Time, updatedAt time.Time, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *MessageRead`

NewMessageRead instantiates a new MessageRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageReadWithDefaults

`func NewMessageReadWithDefaults() *MessageRead`

NewMessageReadWithDefaults instantiates a new MessageRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *MessageRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *MessageRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *MessageRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### SetFleetIdNil

`func (o *MessageRead) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *MessageRead) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetFleetRef

`func (o *MessageRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *MessageRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *MessageRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *MessageRead) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *MessageRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *MessageRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetId

`func (o *MessageRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *MessageRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessageRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MessageRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *MessageRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MessageRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MessageRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *MessageRead) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *MessageRead) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *MessageRead) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *MessageRead) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *MessageRead) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *MessageRead) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetConnectionId

`func (o *MessageRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *MessageRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *MessageRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetTspId

`func (o *MessageRead) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *MessageRead) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *MessageRead) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *MessageRead) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *MessageRead) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *MessageRead) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *MessageRead) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *MessageRead) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *MessageRead) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *MessageRead) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *MessageRead) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *MessageRead) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *MessageRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *MessageRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *MessageRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *MessageRead) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *MessageRead) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *MessageRead) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *MessageRead) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *MessageRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *MessageRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *MessageRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *MessageRead) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *MessageRead) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *MessageRead) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *MessageRead) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *MessageRead) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *MessageRead) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *MessageRead) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *MessageRead) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *MessageRead) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *MessageRead) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *MessageRead) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *MessageRead) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *MessageRead) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *MessageRead) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *MessageRead) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *MessageRead) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *MessageRead) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *MessageRead) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *MessageRead) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *MessageRead) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *MessageRead) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *MessageRead) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *MessageRead) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *MessageRead) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *MessageRead) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *MessageRead) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *MessageRead) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetSourceSenderId

`func (o *MessageRead) GetSourceSenderId() string`

GetSourceSenderId returns the SourceSenderId field if non-nil, zero value otherwise.

### GetSourceSenderIdOk

`func (o *MessageRead) GetSourceSenderIdOk() (*string, bool)`

GetSourceSenderIdOk returns a tuple with the SourceSenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSenderId

`func (o *MessageRead) SetSourceSenderId(v string)`

SetSourceSenderId sets SourceSenderId field to given value.

### HasSourceSenderId

`func (o *MessageRead) HasSourceSenderId() bool`

HasSourceSenderId returns a boolean if a field has been set.

### SetSourceSenderIdNil

`func (o *MessageRead) SetSourceSenderIdNil(b bool)`

 SetSourceSenderIdNil sets the value for SourceSenderId to be an explicit nil

### UnsetSourceSenderId
`func (o *MessageRead) UnsetSourceSenderId()`

UnsetSourceSenderId ensures that no value is present for SourceSenderId, not even an explicit nil
### GetSourceRecipientId

`func (o *MessageRead) GetSourceRecipientId() string`

GetSourceRecipientId returns the SourceRecipientId field if non-nil, zero value otherwise.

### GetSourceRecipientIdOk

`func (o *MessageRead) GetSourceRecipientIdOk() (*string, bool)`

GetSourceRecipientIdOk returns a tuple with the SourceRecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRecipientId

`func (o *MessageRead) SetSourceRecipientId(v string)`

SetSourceRecipientId sets SourceRecipientId field to given value.

### HasSourceRecipientId

`func (o *MessageRead) HasSourceRecipientId() bool`

HasSourceRecipientId returns a boolean if a field has been set.

### SetSourceRecipientIdNil

`func (o *MessageRead) SetSourceRecipientIdNil(b bool)`

 SetSourceRecipientIdNil sets the value for SourceRecipientId to be an explicit nil

### UnsetSourceRecipientId
`func (o *MessageRead) UnsetSourceRecipientId()`

UnsetSourceRecipientId ensures that no value is present for SourceRecipientId, not even an explicit nil
### GetSenderId

`func (o *MessageRead) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *MessageRead) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *MessageRead) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *MessageRead) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### SetSenderIdNil

`func (o *MessageRead) SetSenderIdNil(b bool)`

 SetSenderIdNil sets the value for SenderId to be an explicit nil

### UnsetSenderId
`func (o *MessageRead) UnsetSenderId()`

UnsetSenderId ensures that no value is present for SenderId, not even an explicit nil
### GetRecipientId

`func (o *MessageRead) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *MessageRead) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *MessageRead) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *MessageRead) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### SetRecipientIdNil

`func (o *MessageRead) SetRecipientIdNil(b bool)`

 SetRecipientIdNil sets the value for RecipientId to be an explicit nil

### UnsetRecipientId
`func (o *MessageRead) UnsetRecipientId()`

UnsetRecipientId ensures that no value is present for RecipientId, not even an explicit nil
### GetMessageText

`func (o *MessageRead) GetMessageText() string`

GetMessageText returns the MessageText field if non-nil, zero value otherwise.

### GetMessageTextOk

`func (o *MessageRead) GetMessageTextOk() (*string, bool)`

GetMessageTextOk returns a tuple with the MessageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageText

`func (o *MessageRead) SetMessageText(v string)`

SetMessageText sets MessageText field to given value.

### HasMessageText

`func (o *MessageRead) HasMessageText() bool`

HasMessageText returns a boolean if a field has been set.

### SetMessageTextNil

`func (o *MessageRead) SetMessageTextNil(b bool)`

 SetMessageTextNil sets the value for MessageText to be an explicit nil

### UnsetMessageText
`func (o *MessageRead) UnsetMessageText()`

UnsetMessageText ensures that no value is present for MessageText, not even an explicit nil
### GetMessageHash

`func (o *MessageRead) GetMessageHash() string`

GetMessageHash returns the MessageHash field if non-nil, zero value otherwise.

### GetMessageHashOk

`func (o *MessageRead) GetMessageHashOk() (*string, bool)`

GetMessageHashOk returns a tuple with the MessageHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageHash

`func (o *MessageRead) SetMessageHash(v string)`

SetMessageHash sets MessageHash field to given value.

### HasMessageHash

`func (o *MessageRead) HasMessageHash() bool`

HasMessageHash returns a boolean if a field has been set.

### SetMessageHashNil

`func (o *MessageRead) SetMessageHashNil(b bool)`

 SetMessageHashNil sets the value for MessageHash to be an explicit nil

### UnsetMessageHash
`func (o *MessageRead) UnsetMessageHash()`

UnsetMessageHash ensures that no value is present for MessageHash, not even an explicit nil
### GetStatus

`func (o *MessageRead) GetStatus() MessageStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessageRead) GetStatusOk() (*MessageStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessageRead) SetStatus(v MessageStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MessageRead) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *MessageRead) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *MessageRead) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDeliveredAt

`func (o *MessageRead) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *MessageRead) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *MessageRead) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *MessageRead) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### SetDeliveredAtNil

`func (o *MessageRead) SetDeliveredAtNil(b bool)`

 SetDeliveredAtNil sets the value for DeliveredAt to be an explicit nil

### UnsetDeliveredAt
`func (o *MessageRead) UnsetDeliveredAt()`

UnsetDeliveredAt ensures that no value is present for DeliveredAt, not even an explicit nil
### GetReadAt

`func (o *MessageRead) GetReadAt() time.Time`

GetReadAt returns the ReadAt field if non-nil, zero value otherwise.

### GetReadAtOk

`func (o *MessageRead) GetReadAtOk() (*time.Time, bool)`

GetReadAtOk returns a tuple with the ReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAt

`func (o *MessageRead) SetReadAt(v time.Time)`

SetReadAt sets ReadAt field to given value.

### HasReadAt

`func (o *MessageRead) HasReadAt() bool`

HasReadAt returns a boolean if a field has been set.

### SetReadAtNil

`func (o *MessageRead) SetReadAtNil(b bool)`

 SetReadAtNil sets the value for ReadAt to be an explicit nil

### UnsetReadAt
`func (o *MessageRead) UnsetReadAt()`

UnsetReadAt ensures that no value is present for ReadAt, not even an explicit nil
### GetSentByDriver

`func (o *MessageRead) GetSentByDriver() bool`

GetSentByDriver returns the SentByDriver field if non-nil, zero value otherwise.

### GetSentByDriverOk

`func (o *MessageRead) GetSentByDriverOk() (*bool, bool)`

GetSentByDriverOk returns a tuple with the SentByDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentByDriver

`func (o *MessageRead) SetSentByDriver(v bool)`

SetSentByDriver sets SentByDriver field to given value.

### HasSentByDriver

`func (o *MessageRead) HasSentByDriver() bool`

HasSentByDriver returns a boolean if a field has been set.

### SetSentByDriverNil

`func (o *MessageRead) SetSentByDriverNil(b bool)`

 SetSentByDriverNil sets the value for SentByDriver to be an explicit nil

### UnsetSentByDriver
`func (o *MessageRead) UnsetSentByDriver()`

UnsetSentByDriver ensures that no value is present for SentByDriver, not even an explicit nil
### GetSenderName

`func (o *MessageRead) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *MessageRead) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *MessageRead) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *MessageRead) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### SetSenderNameNil

`func (o *MessageRead) SetSenderNameNil(b bool)`

 SetSenderNameNil sets the value for SenderName to be an explicit nil

### UnsetSenderName
`func (o *MessageRead) UnsetSenderName()`

UnsetSenderName ensures that no value is present for SenderName, not even an explicit nil
### GetRecipientName

`func (o *MessageRead) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *MessageRead) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *MessageRead) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.

### HasRecipientName

`func (o *MessageRead) HasRecipientName() bool`

HasRecipientName returns a boolean if a field has been set.

### SetRecipientNameNil

`func (o *MessageRead) SetRecipientNameNil(b bool)`

 SetRecipientNameNil sets the value for RecipientName to be an explicit nil

### UnsetRecipientName
`func (o *MessageRead) UnsetRecipientName()`

UnsetRecipientName ensures that no value is present for RecipientName, not even an explicit nil
### GetThreadId

`func (o *MessageRead) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *MessageRead) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *MessageRead) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *MessageRead) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### SetThreadIdNil

`func (o *MessageRead) SetThreadIdNil(b bool)`

 SetThreadIdNil sets the value for ThreadId to be an explicit nil

### UnsetThreadId
`func (o *MessageRead) UnsetThreadId()`

UnsetThreadId ensures that no value is present for ThreadId, not even an explicit nil
### GetSourceThreadId

`func (o *MessageRead) GetSourceThreadId() string`

GetSourceThreadId returns the SourceThreadId field if non-nil, zero value otherwise.

### GetSourceThreadIdOk

`func (o *MessageRead) GetSourceThreadIdOk() (*string, bool)`

GetSourceThreadIdOk returns a tuple with the SourceThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceThreadId

`func (o *MessageRead) SetSourceThreadId(v string)`

SetSourceThreadId sets SourceThreadId field to given value.

### HasSourceThreadId

`func (o *MessageRead) HasSourceThreadId() bool`

HasSourceThreadId returns a boolean if a field has been set.

### SetSourceThreadIdNil

`func (o *MessageRead) SetSourceThreadIdNil(b bool)`

 SetSourceThreadIdNil sets the value for SourceThreadId to be an explicit nil

### UnsetSourceThreadId
`func (o *MessageRead) UnsetSourceThreadId()`

UnsetSourceThreadId ensures that no value is present for SourceThreadId, not even an explicit nil
### GetReplyToMessageId

`func (o *MessageRead) GetReplyToMessageId() string`

GetReplyToMessageId returns the ReplyToMessageId field if non-nil, zero value otherwise.

### GetReplyToMessageIdOk

`func (o *MessageRead) GetReplyToMessageIdOk() (*string, bool)`

GetReplyToMessageIdOk returns a tuple with the ReplyToMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToMessageId

`func (o *MessageRead) SetReplyToMessageId(v string)`

SetReplyToMessageId sets ReplyToMessageId field to given value.

### HasReplyToMessageId

`func (o *MessageRead) HasReplyToMessageId() bool`

HasReplyToMessageId returns a boolean if a field has been set.

### SetReplyToMessageIdNil

`func (o *MessageRead) SetReplyToMessageIdNil(b bool)`

 SetReplyToMessageIdNil sets the value for ReplyToMessageId to be an explicit nil

### UnsetReplyToMessageId
`func (o *MessageRead) UnsetReplyToMessageId()`

UnsetReplyToMessageId ensures that no value is present for ReplyToMessageId, not even an explicit nil
### GetSourceReplyToMessageId

`func (o *MessageRead) GetSourceReplyToMessageId() string`

GetSourceReplyToMessageId returns the SourceReplyToMessageId field if non-nil, zero value otherwise.

### GetSourceReplyToMessageIdOk

`func (o *MessageRead) GetSourceReplyToMessageIdOk() (*string, bool)`

GetSourceReplyToMessageIdOk returns a tuple with the SourceReplyToMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceReplyToMessageId

`func (o *MessageRead) SetSourceReplyToMessageId(v string)`

SetSourceReplyToMessageId sets SourceReplyToMessageId field to given value.

### HasSourceReplyToMessageId

`func (o *MessageRead) HasSourceReplyToMessageId() bool`

HasSourceReplyToMessageId returns a boolean if a field has been set.

### SetSourceReplyToMessageIdNil

`func (o *MessageRead) SetSourceReplyToMessageIdNil(b bool)`

 SetSourceReplyToMessageIdNil sets the value for SourceReplyToMessageId to be an explicit nil

### UnsetSourceReplyToMessageId
`func (o *MessageRead) UnsetSourceReplyToMessageId()`

UnsetSourceReplyToMessageId ensures that no value is present for SourceReplyToMessageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


