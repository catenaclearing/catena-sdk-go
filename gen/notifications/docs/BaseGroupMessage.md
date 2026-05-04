# BaseGroupMessage

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

### NewBaseGroupMessage

`func NewBaseGroupMessage(id string, fleetId string, fleetRef NullableString, sourceName TspEnum, connectionId string, sourceId string, createdAt time.Time, updatedAt time.Time, occurredAt time.Time, ) *BaseGroupMessage`

NewBaseGroupMessage instantiates a new BaseGroupMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseGroupMessageWithDefaults

`func NewBaseGroupMessageWithDefaults() *BaseGroupMessage`

NewBaseGroupMessageWithDefaults instantiates a new BaseGroupMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseGroupMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseGroupMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseGroupMessage) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseGroupMessage) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseGroupMessage) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseGroupMessage) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *BaseGroupMessage) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *BaseGroupMessage) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *BaseGroupMessage) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### SetFleetRefNil

`func (o *BaseGroupMessage) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *BaseGroupMessage) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetTspId

`func (o *BaseGroupMessage) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *BaseGroupMessage) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *BaseGroupMessage) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *BaseGroupMessage) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *BaseGroupMessage) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *BaseGroupMessage) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *BaseGroupMessage) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *BaseGroupMessage) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *BaseGroupMessage) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.

### HasTspSlug

`func (o *BaseGroupMessage) HasTspSlug() bool`

HasTspSlug returns a boolean if a field has been set.

### SetTspSlugNil

`func (o *BaseGroupMessage) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *BaseGroupMessage) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *BaseGroupMessage) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseGroupMessage) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseGroupMessage) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetConnectionId

`func (o *BaseGroupMessage) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseGroupMessage) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseGroupMessage) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceId

`func (o *BaseGroupMessage) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseGroupMessage) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseGroupMessage) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCreatedAt

`func (o *BaseGroupMessage) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseGroupMessage) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseGroupMessage) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseGroupMessage) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseGroupMessage) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseGroupMessage) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BaseGroupMessage) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BaseGroupMessage) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BaseGroupMessage) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BaseGroupMessage) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BaseGroupMessage) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BaseGroupMessage) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOccurredAt

`func (o *BaseGroupMessage) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BaseGroupMessage) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BaseGroupMessage) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetExecutionId

`func (o *BaseGroupMessage) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BaseGroupMessage) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BaseGroupMessage) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *BaseGroupMessage) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *BaseGroupMessage) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *BaseGroupMessage) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *BaseGroupMessage) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BaseGroupMessage) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BaseGroupMessage) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *BaseGroupMessage) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *BaseGroupMessage) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *BaseGroupMessage) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetExtras

`func (o *BaseGroupMessage) GetExtras() map[string]interface{}`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *BaseGroupMessage) GetExtrasOk() (*map[string]interface{}, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *BaseGroupMessage) SetExtras(v map[string]interface{})`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *BaseGroupMessage) HasExtras() bool`

HasExtras returns a boolean if a field has been set.

### SetExtrasNil

`func (o *BaseGroupMessage) SetExtrasNil(b bool)`

 SetExtrasNil sets the value for Extras to be an explicit nil

### UnsetExtras
`func (o *BaseGroupMessage) UnsetExtras()`

UnsetExtras ensures that no value is present for Extras, not even an explicit nil
### GetSourceSenderId

`func (o *BaseGroupMessage) GetSourceSenderId() string`

GetSourceSenderId returns the SourceSenderId field if non-nil, zero value otherwise.

### GetSourceSenderIdOk

`func (o *BaseGroupMessage) GetSourceSenderIdOk() (*string, bool)`

GetSourceSenderIdOk returns a tuple with the SourceSenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSenderId

`func (o *BaseGroupMessage) SetSourceSenderId(v string)`

SetSourceSenderId sets SourceSenderId field to given value.

### HasSourceSenderId

`func (o *BaseGroupMessage) HasSourceSenderId() bool`

HasSourceSenderId returns a boolean if a field has been set.

### SetSourceSenderIdNil

`func (o *BaseGroupMessage) SetSourceSenderIdNil(b bool)`

 SetSourceSenderIdNil sets the value for SourceSenderId to be an explicit nil

### UnsetSourceSenderId
`func (o *BaseGroupMessage) UnsetSourceSenderId()`

UnsetSourceSenderId ensures that no value is present for SourceSenderId, not even an explicit nil
### GetSourceRecipientIds

`func (o *BaseGroupMessage) GetSourceRecipientIds() []string`

GetSourceRecipientIds returns the SourceRecipientIds field if non-nil, zero value otherwise.

### GetSourceRecipientIdsOk

`func (o *BaseGroupMessage) GetSourceRecipientIdsOk() (*[]string, bool)`

GetSourceRecipientIdsOk returns a tuple with the SourceRecipientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRecipientIds

`func (o *BaseGroupMessage) SetSourceRecipientIds(v []string)`

SetSourceRecipientIds sets SourceRecipientIds field to given value.

### HasSourceRecipientIds

`func (o *BaseGroupMessage) HasSourceRecipientIds() bool`

HasSourceRecipientIds returns a boolean if a field has been set.

### SetSourceRecipientIdsNil

`func (o *BaseGroupMessage) SetSourceRecipientIdsNil(b bool)`

 SetSourceRecipientIdsNil sets the value for SourceRecipientIds to be an explicit nil

### UnsetSourceRecipientIds
`func (o *BaseGroupMessage) UnsetSourceRecipientIds()`

UnsetSourceRecipientIds ensures that no value is present for SourceRecipientIds, not even an explicit nil
### GetSenderId

`func (o *BaseGroupMessage) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *BaseGroupMessage) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *BaseGroupMessage) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *BaseGroupMessage) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### SetSenderIdNil

`func (o *BaseGroupMessage) SetSenderIdNil(b bool)`

 SetSenderIdNil sets the value for SenderId to be an explicit nil

### UnsetSenderId
`func (o *BaseGroupMessage) UnsetSenderId()`

UnsetSenderId ensures that no value is present for SenderId, not even an explicit nil
### GetRecipientIds

`func (o *BaseGroupMessage) GetRecipientIds() []string`

GetRecipientIds returns the RecipientIds field if non-nil, zero value otherwise.

### GetRecipientIdsOk

`func (o *BaseGroupMessage) GetRecipientIdsOk() (*[]string, bool)`

GetRecipientIdsOk returns a tuple with the RecipientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIds

`func (o *BaseGroupMessage) SetRecipientIds(v []string)`

SetRecipientIds sets RecipientIds field to given value.

### HasRecipientIds

`func (o *BaseGroupMessage) HasRecipientIds() bool`

HasRecipientIds returns a boolean if a field has been set.

### SetRecipientIdsNil

`func (o *BaseGroupMessage) SetRecipientIdsNil(b bool)`

 SetRecipientIdsNil sets the value for RecipientIds to be an explicit nil

### UnsetRecipientIds
`func (o *BaseGroupMessage) UnsetRecipientIds()`

UnsetRecipientIds ensures that no value is present for RecipientIds, not even an explicit nil
### GetMessageText

`func (o *BaseGroupMessage) GetMessageText() string`

GetMessageText returns the MessageText field if non-nil, zero value otherwise.

### GetMessageTextOk

`func (o *BaseGroupMessage) GetMessageTextOk() (*string, bool)`

GetMessageTextOk returns a tuple with the MessageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageText

`func (o *BaseGroupMessage) SetMessageText(v string)`

SetMessageText sets MessageText field to given value.

### HasMessageText

`func (o *BaseGroupMessage) HasMessageText() bool`

HasMessageText returns a boolean if a field has been set.

### SetMessageTextNil

`func (o *BaseGroupMessage) SetMessageTextNil(b bool)`

 SetMessageTextNil sets the value for MessageText to be an explicit nil

### UnsetMessageText
`func (o *BaseGroupMessage) UnsetMessageText()`

UnsetMessageText ensures that no value is present for MessageText, not even an explicit nil
### GetMessageHash

`func (o *BaseGroupMessage) GetMessageHash() string`

GetMessageHash returns the MessageHash field if non-nil, zero value otherwise.

### GetMessageHashOk

`func (o *BaseGroupMessage) GetMessageHashOk() (*string, bool)`

GetMessageHashOk returns a tuple with the MessageHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageHash

`func (o *BaseGroupMessage) SetMessageHash(v string)`

SetMessageHash sets MessageHash field to given value.

### HasMessageHash

`func (o *BaseGroupMessage) HasMessageHash() bool`

HasMessageHash returns a boolean if a field has been set.

### SetMessageHashNil

`func (o *BaseGroupMessage) SetMessageHashNil(b bool)`

 SetMessageHashNil sets the value for MessageHash to be an explicit nil

### UnsetMessageHash
`func (o *BaseGroupMessage) UnsetMessageHash()`

UnsetMessageHash ensures that no value is present for MessageHash, not even an explicit nil
### GetStatus

`func (o *BaseGroupMessage) GetStatus() MessageStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseGroupMessage) GetStatusOk() (*MessageStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseGroupMessage) SetStatus(v MessageStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseGroupMessage) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BaseGroupMessage) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BaseGroupMessage) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDeliveredAt

`func (o *BaseGroupMessage) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *BaseGroupMessage) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *BaseGroupMessage) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *BaseGroupMessage) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### SetDeliveredAtNil

`func (o *BaseGroupMessage) SetDeliveredAtNil(b bool)`

 SetDeliveredAtNil sets the value for DeliveredAt to be an explicit nil

### UnsetDeliveredAt
`func (o *BaseGroupMessage) UnsetDeliveredAt()`

UnsetDeliveredAt ensures that no value is present for DeliveredAt, not even an explicit nil
### GetSentByDriver

`func (o *BaseGroupMessage) GetSentByDriver() bool`

GetSentByDriver returns the SentByDriver field if non-nil, zero value otherwise.

### GetSentByDriverOk

`func (o *BaseGroupMessage) GetSentByDriverOk() (*bool, bool)`

GetSentByDriverOk returns a tuple with the SentByDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentByDriver

`func (o *BaseGroupMessage) SetSentByDriver(v bool)`

SetSentByDriver sets SentByDriver field to given value.

### HasSentByDriver

`func (o *BaseGroupMessage) HasSentByDriver() bool`

HasSentByDriver returns a boolean if a field has been set.

### SetSentByDriverNil

`func (o *BaseGroupMessage) SetSentByDriverNil(b bool)`

 SetSentByDriverNil sets the value for SentByDriver to be an explicit nil

### UnsetSentByDriver
`func (o *BaseGroupMessage) UnsetSentByDriver()`

UnsetSentByDriver ensures that no value is present for SentByDriver, not even an explicit nil
### GetSenderName

`func (o *BaseGroupMessage) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *BaseGroupMessage) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *BaseGroupMessage) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *BaseGroupMessage) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### SetSenderNameNil

`func (o *BaseGroupMessage) SetSenderNameNil(b bool)`

 SetSenderNameNil sets the value for SenderName to be an explicit nil

### UnsetSenderName
`func (o *BaseGroupMessage) UnsetSenderName()`

UnsetSenderName ensures that no value is present for SenderName, not even an explicit nil
### GetThreadId

`func (o *BaseGroupMessage) GetThreadId() string`

GetThreadId returns the ThreadId field if non-nil, zero value otherwise.

### GetThreadIdOk

`func (o *BaseGroupMessage) GetThreadIdOk() (*string, bool)`

GetThreadIdOk returns a tuple with the ThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadId

`func (o *BaseGroupMessage) SetThreadId(v string)`

SetThreadId sets ThreadId field to given value.

### HasThreadId

`func (o *BaseGroupMessage) HasThreadId() bool`

HasThreadId returns a boolean if a field has been set.

### SetThreadIdNil

`func (o *BaseGroupMessage) SetThreadIdNil(b bool)`

 SetThreadIdNil sets the value for ThreadId to be an explicit nil

### UnsetThreadId
`func (o *BaseGroupMessage) UnsetThreadId()`

UnsetThreadId ensures that no value is present for ThreadId, not even an explicit nil
### GetSourceThreadId

`func (o *BaseGroupMessage) GetSourceThreadId() string`

GetSourceThreadId returns the SourceThreadId field if non-nil, zero value otherwise.

### GetSourceThreadIdOk

`func (o *BaseGroupMessage) GetSourceThreadIdOk() (*string, bool)`

GetSourceThreadIdOk returns a tuple with the SourceThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceThreadId

`func (o *BaseGroupMessage) SetSourceThreadId(v string)`

SetSourceThreadId sets SourceThreadId field to given value.

### HasSourceThreadId

`func (o *BaseGroupMessage) HasSourceThreadId() bool`

HasSourceThreadId returns a boolean if a field has been set.

### SetSourceThreadIdNil

`func (o *BaseGroupMessage) SetSourceThreadIdNil(b bool)`

 SetSourceThreadIdNil sets the value for SourceThreadId to be an explicit nil

### UnsetSourceThreadId
`func (o *BaseGroupMessage) UnsetSourceThreadId()`

UnsetSourceThreadId ensures that no value is present for SourceThreadId, not even an explicit nil
### GetGroupId

`func (o *BaseGroupMessage) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *BaseGroupMessage) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *BaseGroupMessage) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *BaseGroupMessage) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *BaseGroupMessage) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *BaseGroupMessage) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetSourceGroupId

`func (o *BaseGroupMessage) GetSourceGroupId() string`

GetSourceGroupId returns the SourceGroupId field if non-nil, zero value otherwise.

### GetSourceGroupIdOk

`func (o *BaseGroupMessage) GetSourceGroupIdOk() (*string, bool)`

GetSourceGroupIdOk returns a tuple with the SourceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceGroupId

`func (o *BaseGroupMessage) SetSourceGroupId(v string)`

SetSourceGroupId sets SourceGroupId field to given value.

### HasSourceGroupId

`func (o *BaseGroupMessage) HasSourceGroupId() bool`

HasSourceGroupId returns a boolean if a field has been set.

### SetSourceGroupIdNil

`func (o *BaseGroupMessage) SetSourceGroupIdNil(b bool)`

 SetSourceGroupIdNil sets the value for SourceGroupId to be an explicit nil

### UnsetSourceGroupId
`func (o *BaseGroupMessage) UnsetSourceGroupId()`

UnsetSourceGroupId ensures that no value is present for SourceGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


