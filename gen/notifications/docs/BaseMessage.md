# BaseMessage

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
**SourceThreadId** | Pointer to **NullableString** |  | [optional] 
**SourceReplyToMessageId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBaseMessage

`func NewBaseMessage(id string, fleetId string, fleetRef NullableString, sourceName TspEnum, connectionId string, sourceId string, createdAt time.Time, updatedAt time.Time, occurredAt time.Time, ) *BaseMessage`

NewBaseMessage instantiates a new BaseMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseMessageWithDefaults

`func NewBaseMessageWithDefaults() *BaseMessage`

NewBaseMessageWithDefaults instantiates a new BaseMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseMessage) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseMessage) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseMessage) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseMessage) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *BaseMessage) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *BaseMessage) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *BaseMessage) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### SetFleetRefNil

`func (o *BaseMessage) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *BaseMessage) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetTspId

`func (o *BaseMessage) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *BaseMessage) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *BaseMessage) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *BaseMessage) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *BaseMessage) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *BaseMessage) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *BaseMessage) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *BaseMessage) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *BaseMessage) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *BaseMessage) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *BaseMessage) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *BaseMessage) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *BaseMessage) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseMessage) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseMessage) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetConnectionId

`func (o *BaseMessage) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseMessage) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseMessage) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceId

`func (o *BaseMessage) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseMessage) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseMessage) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCreatedAt

`func (o *BaseMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseMessage) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseMessage) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseMessage) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BaseMessage) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BaseMessage) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BaseMessage) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BaseMessage) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BaseMessage) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BaseMessage) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOccurredAt

`func (o *BaseMessage) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BaseMessage) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BaseMessage) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetExecutionId

`func (o *BaseMessage) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BaseMessage) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BaseMessage) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *BaseMessage) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *BaseMessage) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *BaseMessage) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *BaseMessage) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BaseMessage) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BaseMessage) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *BaseMessage) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *BaseMessage) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *BaseMessage) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *BaseMessage) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *BaseMessage) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *BaseMessage) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *BaseMessage) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *BaseMessage) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *BaseMessage) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetSourceSenderId

`func (o *BaseMessage) GetSourceSenderId() string`

GetSourceSenderId returns the SourceSenderId field if non-nil, zero value otherwise.

### GetSourceSenderIdOk

`func (o *BaseMessage) GetSourceSenderIdOk() (*string, bool)`

GetSourceSenderIdOk returns a tuple with the SourceSenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSenderId

`func (o *BaseMessage) SetSourceSenderId(v string)`

SetSourceSenderId sets SourceSenderId field to given value.

### HasSourceSenderId

`func (o *BaseMessage) HasSourceSenderId() bool`

HasSourceSenderId returns a boolean if a field has been set.

### SetSourceSenderIdNil

`func (o *BaseMessage) SetSourceSenderIdNil(b bool)`

 SetSourceSenderIdNil sets the value for SourceSenderId to be an explicit nil

### UnsetSourceSenderId
`func (o *BaseMessage) UnsetSourceSenderId()`

UnsetSourceSenderId ensures that no value is present for SourceSenderId, not even an explicit nil
### GetSourceRecipientId

`func (o *BaseMessage) GetSourceRecipientId() string`

GetSourceRecipientId returns the SourceRecipientId field if non-nil, zero value otherwise.

### GetSourceRecipientIdOk

`func (o *BaseMessage) GetSourceRecipientIdOk() (*string, bool)`

GetSourceRecipientIdOk returns a tuple with the SourceRecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRecipientId

`func (o *BaseMessage) SetSourceRecipientId(v string)`

SetSourceRecipientId sets SourceRecipientId field to given value.

### HasSourceRecipientId

`func (o *BaseMessage) HasSourceRecipientId() bool`

HasSourceRecipientId returns a boolean if a field has been set.

### SetSourceRecipientIdNil

`func (o *BaseMessage) SetSourceRecipientIdNil(b bool)`

 SetSourceRecipientIdNil sets the value for SourceRecipientId to be an explicit nil

### UnsetSourceRecipientId
`func (o *BaseMessage) UnsetSourceRecipientId()`

UnsetSourceRecipientId ensures that no value is present for SourceRecipientId, not even an explicit nil
### GetSenderId

`func (o *BaseMessage) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *BaseMessage) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *BaseMessage) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *BaseMessage) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### SetSenderIdNil

`func (o *BaseMessage) SetSenderIdNil(b bool)`

 SetSenderIdNil sets the value for SenderId to be an explicit nil

### UnsetSenderId
`func (o *BaseMessage) UnsetSenderId()`

UnsetSenderId ensures that no value is present for SenderId, not even an explicit nil
### GetRecipientId

`func (o *BaseMessage) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *BaseMessage) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *BaseMessage) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.

### HasRecipientId

`func (o *BaseMessage) HasRecipientId() bool`

HasRecipientId returns a boolean if a field has been set.

### SetRecipientIdNil

`func (o *BaseMessage) SetRecipientIdNil(b bool)`

 SetRecipientIdNil sets the value for RecipientId to be an explicit nil

### UnsetRecipientId
`func (o *BaseMessage) UnsetRecipientId()`

UnsetRecipientId ensures that no value is present for RecipientId, not even an explicit nil
### GetMessageText

`func (o *BaseMessage) GetMessageText() string`

GetMessageText returns the MessageText field if non-nil, zero value otherwise.

### GetMessageTextOk

`func (o *BaseMessage) GetMessageTextOk() (*string, bool)`

GetMessageTextOk returns a tuple with the MessageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageText

`func (o *BaseMessage) SetMessageText(v string)`

SetMessageText sets MessageText field to given value.

### HasMessageText

`func (o *BaseMessage) HasMessageText() bool`

HasMessageText returns a boolean if a field has been set.

### SetMessageTextNil

`func (o *BaseMessage) SetMessageTextNil(b bool)`

 SetMessageTextNil sets the value for MessageText to be an explicit nil

### UnsetMessageText
`func (o *BaseMessage) UnsetMessageText()`

UnsetMessageText ensures that no value is present for MessageText, not even an explicit nil
### GetMessageHash

`func (o *BaseMessage) GetMessageHash() string`

GetMessageHash returns the MessageHash field if non-nil, zero value otherwise.

### GetMessageHashOk

`func (o *BaseMessage) GetMessageHashOk() (*string, bool)`

GetMessageHashOk returns a tuple with the MessageHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageHash

`func (o *BaseMessage) SetMessageHash(v string)`

SetMessageHash sets MessageHash field to given value.

### HasMessageHash

`func (o *BaseMessage) HasMessageHash() bool`

HasMessageHash returns a boolean if a field has been set.

### SetMessageHashNil

`func (o *BaseMessage) SetMessageHashNil(b bool)`

 SetMessageHashNil sets the value for MessageHash to be an explicit nil

### UnsetMessageHash
`func (o *BaseMessage) UnsetMessageHash()`

UnsetMessageHash ensures that no value is present for MessageHash, not even an explicit nil
### GetStatus

`func (o *BaseMessage) GetStatus() MessageStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseMessage) GetStatusOk() (*MessageStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseMessage) SetStatus(v MessageStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BaseMessage) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BaseMessage) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDeliveredAt

`func (o *BaseMessage) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *BaseMessage) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *BaseMessage) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *BaseMessage) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### SetDeliveredAtNil

`func (o *BaseMessage) SetDeliveredAtNil(b bool)`

 SetDeliveredAtNil sets the value for DeliveredAt to be an explicit nil

### UnsetDeliveredAt
`func (o *BaseMessage) UnsetDeliveredAt()`

UnsetDeliveredAt ensures that no value is present for DeliveredAt, not even an explicit nil
### GetReadAt

`func (o *BaseMessage) GetReadAt() time.Time`

GetReadAt returns the ReadAt field if non-nil, zero value otherwise.

### GetReadAtOk

`func (o *BaseMessage) GetReadAtOk() (*time.Time, bool)`

GetReadAtOk returns a tuple with the ReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAt

`func (o *BaseMessage) SetReadAt(v time.Time)`

SetReadAt sets ReadAt field to given value.

### HasReadAt

`func (o *BaseMessage) HasReadAt() bool`

HasReadAt returns a boolean if a field has been set.

### SetReadAtNil

`func (o *BaseMessage) SetReadAtNil(b bool)`

 SetReadAtNil sets the value for ReadAt to be an explicit nil

### UnsetReadAt
`func (o *BaseMessage) UnsetReadAt()`

UnsetReadAt ensures that no value is present for ReadAt, not even an explicit nil
### GetSentByDriver

`func (o *BaseMessage) GetSentByDriver() bool`

GetSentByDriver returns the SentByDriver field if non-nil, zero value otherwise.

### GetSentByDriverOk

`func (o *BaseMessage) GetSentByDriverOk() (*bool, bool)`

GetSentByDriverOk returns a tuple with the SentByDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentByDriver

`func (o *BaseMessage) SetSentByDriver(v bool)`

SetSentByDriver sets SentByDriver field to given value.

### HasSentByDriver

`func (o *BaseMessage) HasSentByDriver() bool`

HasSentByDriver returns a boolean if a field has been set.

### SetSentByDriverNil

`func (o *BaseMessage) SetSentByDriverNil(b bool)`

 SetSentByDriverNil sets the value for SentByDriver to be an explicit nil

### UnsetSentByDriver
`func (o *BaseMessage) UnsetSentByDriver()`

UnsetSentByDriver ensures that no value is present for SentByDriver, not even an explicit nil
### GetSenderName

`func (o *BaseMessage) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *BaseMessage) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *BaseMessage) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *BaseMessage) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### SetSenderNameNil

`func (o *BaseMessage) SetSenderNameNil(b bool)`

 SetSenderNameNil sets the value for SenderName to be an explicit nil

### UnsetSenderName
`func (o *BaseMessage) UnsetSenderName()`

UnsetSenderName ensures that no value is present for SenderName, not even an explicit nil
### GetRecipientName

`func (o *BaseMessage) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *BaseMessage) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *BaseMessage) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.

### HasRecipientName

`func (o *BaseMessage) HasRecipientName() bool`

HasRecipientName returns a boolean if a field has been set.

### SetRecipientNameNil

`func (o *BaseMessage) SetRecipientNameNil(b bool)`

 SetRecipientNameNil sets the value for RecipientName to be an explicit nil

### UnsetRecipientName
`func (o *BaseMessage) UnsetRecipientName()`

UnsetRecipientName ensures that no value is present for RecipientName, not even an explicit nil
### GetSourceThreadId

`func (o *BaseMessage) GetSourceThreadId() string`

GetSourceThreadId returns the SourceThreadId field if non-nil, zero value otherwise.

### GetSourceThreadIdOk

`func (o *BaseMessage) GetSourceThreadIdOk() (*string, bool)`

GetSourceThreadIdOk returns a tuple with the SourceThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceThreadId

`func (o *BaseMessage) SetSourceThreadId(v string)`

SetSourceThreadId sets SourceThreadId field to given value.

### HasSourceThreadId

`func (o *BaseMessage) HasSourceThreadId() bool`

HasSourceThreadId returns a boolean if a field has been set.

### SetSourceThreadIdNil

`func (o *BaseMessage) SetSourceThreadIdNil(b bool)`

 SetSourceThreadIdNil sets the value for SourceThreadId to be an explicit nil

### UnsetSourceThreadId
`func (o *BaseMessage) UnsetSourceThreadId()`

UnsetSourceThreadId ensures that no value is present for SourceThreadId, not even an explicit nil
### GetSourceReplyToMessageId

`func (o *BaseMessage) GetSourceReplyToMessageId() string`

GetSourceReplyToMessageId returns the SourceReplyToMessageId field if non-nil, zero value otherwise.

### GetSourceReplyToMessageIdOk

`func (o *BaseMessage) GetSourceReplyToMessageIdOk() (*string, bool)`

GetSourceReplyToMessageIdOk returns a tuple with the SourceReplyToMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceReplyToMessageId

`func (o *BaseMessage) SetSourceReplyToMessageId(v string)`

SetSourceReplyToMessageId sets SourceReplyToMessageId field to given value.

### HasSourceReplyToMessageId

`func (o *BaseMessage) HasSourceReplyToMessageId() bool`

HasSourceReplyToMessageId returns a boolean if a field has been set.

### SetSourceReplyToMessageIdNil

`func (o *BaseMessage) SetSourceReplyToMessageIdNil(b bool)`

 SetSourceReplyToMessageIdNil sets the value for SourceReplyToMessageId to be an explicit nil

### UnsetSourceReplyToMessageId
`func (o *BaseMessage) UnsetSourceReplyToMessageId()`

UnsetSourceReplyToMessageId ensures that no value is present for SourceReplyToMessageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


