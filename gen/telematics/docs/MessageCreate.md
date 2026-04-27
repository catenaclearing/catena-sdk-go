# MessageCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | Unique identifier of the connection at Catena Telematics which will be used to create this resource. A connection represents a Fleet/TSP pairing. | 
**FleetRef** | Pointer to **NullableString** |  | [optional] 
**SourceSenderId** | Pointer to **NullableString** |  | [optional] 
**SourceRecipientIds** | Pointer to **[]string** |  | [optional] 
**RecipientGroup** | Pointer to [**NullableMessageReceiverGroupEnum**](MessageReceiverGroupEnum.md) |  | [optional] 
**SenderGroup** | Pointer to [**NullableMessageSenderGroupEnum**](MessageSenderGroupEnum.md) |  | [optional] 
**MessageText** | Pointer to **NullableString** |  | [optional] 
**Mode** | Pointer to [**NullableMessageModeEnum**](MessageModeEnum.md) |  | [optional] 
**Priority** | Pointer to [**NullableMessagePriorityEnum**](MessagePriorityEnum.md) |  | [optional] 
**SourceThreadId** | Pointer to **NullableString** |  | [optional] 
**SourceReplyToMessageId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMessageCreate

`func NewMessageCreate(connectionId string, ) *MessageCreate`

NewMessageCreate instantiates a new MessageCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageCreateWithDefaults

`func NewMessageCreateWithDefaults() *MessageCreate`

NewMessageCreateWithDefaults instantiates a new MessageCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *MessageCreate) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *MessageCreate) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *MessageCreate) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetFleetRef

`func (o *MessageCreate) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *MessageCreate) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *MessageCreate) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *MessageCreate) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *MessageCreate) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *MessageCreate) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetSourceSenderId

`func (o *MessageCreate) GetSourceSenderId() string`

GetSourceSenderId returns the SourceSenderId field if non-nil, zero value otherwise.

### GetSourceSenderIdOk

`func (o *MessageCreate) GetSourceSenderIdOk() (*string, bool)`

GetSourceSenderIdOk returns a tuple with the SourceSenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSenderId

`func (o *MessageCreate) SetSourceSenderId(v string)`

SetSourceSenderId sets SourceSenderId field to given value.

### HasSourceSenderId

`func (o *MessageCreate) HasSourceSenderId() bool`

HasSourceSenderId returns a boolean if a field has been set.

### SetSourceSenderIdNil

`func (o *MessageCreate) SetSourceSenderIdNil(b bool)`

 SetSourceSenderIdNil sets the value for SourceSenderId to be an explicit nil

### UnsetSourceSenderId
`func (o *MessageCreate) UnsetSourceSenderId()`

UnsetSourceSenderId ensures that no value is present for SourceSenderId, not even an explicit nil
### GetSourceRecipientIds

`func (o *MessageCreate) GetSourceRecipientIds() []string`

GetSourceRecipientIds returns the SourceRecipientIds field if non-nil, zero value otherwise.

### GetSourceRecipientIdsOk

`func (o *MessageCreate) GetSourceRecipientIdsOk() (*[]string, bool)`

GetSourceRecipientIdsOk returns a tuple with the SourceRecipientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRecipientIds

`func (o *MessageCreate) SetSourceRecipientIds(v []string)`

SetSourceRecipientIds sets SourceRecipientIds field to given value.

### HasSourceRecipientIds

`func (o *MessageCreate) HasSourceRecipientIds() bool`

HasSourceRecipientIds returns a boolean if a field has been set.

### SetSourceRecipientIdsNil

`func (o *MessageCreate) SetSourceRecipientIdsNil(b bool)`

 SetSourceRecipientIdsNil sets the value for SourceRecipientIds to be an explicit nil

### UnsetSourceRecipientIds
`func (o *MessageCreate) UnsetSourceRecipientIds()`

UnsetSourceRecipientIds ensures that no value is present for SourceRecipientIds, not even an explicit nil
### GetRecipientGroup

`func (o *MessageCreate) GetRecipientGroup() MessageReceiverGroupEnum`

GetRecipientGroup returns the RecipientGroup field if non-nil, zero value otherwise.

### GetRecipientGroupOk

`func (o *MessageCreate) GetRecipientGroupOk() (*MessageReceiverGroupEnum, bool)`

GetRecipientGroupOk returns a tuple with the RecipientGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientGroup

`func (o *MessageCreate) SetRecipientGroup(v MessageReceiverGroupEnum)`

SetRecipientGroup sets RecipientGroup field to given value.

### HasRecipientGroup

`func (o *MessageCreate) HasRecipientGroup() bool`

HasRecipientGroup returns a boolean if a field has been set.

### SetRecipientGroupNil

`func (o *MessageCreate) SetRecipientGroupNil(b bool)`

 SetRecipientGroupNil sets the value for RecipientGroup to be an explicit nil

### UnsetRecipientGroup
`func (o *MessageCreate) UnsetRecipientGroup()`

UnsetRecipientGroup ensures that no value is present for RecipientGroup, not even an explicit nil
### GetSenderGroup

`func (o *MessageCreate) GetSenderGroup() MessageSenderGroupEnum`

GetSenderGroup returns the SenderGroup field if non-nil, zero value otherwise.

### GetSenderGroupOk

`func (o *MessageCreate) GetSenderGroupOk() (*MessageSenderGroupEnum, bool)`

GetSenderGroupOk returns a tuple with the SenderGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderGroup

`func (o *MessageCreate) SetSenderGroup(v MessageSenderGroupEnum)`

SetSenderGroup sets SenderGroup field to given value.

### HasSenderGroup

`func (o *MessageCreate) HasSenderGroup() bool`

HasSenderGroup returns a boolean if a field has been set.

### SetSenderGroupNil

`func (o *MessageCreate) SetSenderGroupNil(b bool)`

 SetSenderGroupNil sets the value for SenderGroup to be an explicit nil

### UnsetSenderGroup
`func (o *MessageCreate) UnsetSenderGroup()`

UnsetSenderGroup ensures that no value is present for SenderGroup, not even an explicit nil
### GetMessageText

`func (o *MessageCreate) GetMessageText() string`

GetMessageText returns the MessageText field if non-nil, zero value otherwise.

### GetMessageTextOk

`func (o *MessageCreate) GetMessageTextOk() (*string, bool)`

GetMessageTextOk returns a tuple with the MessageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageText

`func (o *MessageCreate) SetMessageText(v string)`

SetMessageText sets MessageText field to given value.

### HasMessageText

`func (o *MessageCreate) HasMessageText() bool`

HasMessageText returns a boolean if a field has been set.

### SetMessageTextNil

`func (o *MessageCreate) SetMessageTextNil(b bool)`

 SetMessageTextNil sets the value for MessageText to be an explicit nil

### UnsetMessageText
`func (o *MessageCreate) UnsetMessageText()`

UnsetMessageText ensures that no value is present for MessageText, not even an explicit nil
### GetMode

`func (o *MessageCreate) GetMode() MessageModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MessageCreate) GetModeOk() (*MessageModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MessageCreate) SetMode(v MessageModeEnum)`

SetMode sets Mode field to given value.

### HasMode

`func (o *MessageCreate) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *MessageCreate) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *MessageCreate) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetPriority

`func (o *MessageCreate) GetPriority() MessagePriorityEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MessageCreate) GetPriorityOk() (*MessagePriorityEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MessageCreate) SetPriority(v MessagePriorityEnum)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MessageCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *MessageCreate) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *MessageCreate) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetSourceThreadId

`func (o *MessageCreate) GetSourceThreadId() string`

GetSourceThreadId returns the SourceThreadId field if non-nil, zero value otherwise.

### GetSourceThreadIdOk

`func (o *MessageCreate) GetSourceThreadIdOk() (*string, bool)`

GetSourceThreadIdOk returns a tuple with the SourceThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceThreadId

`func (o *MessageCreate) SetSourceThreadId(v string)`

SetSourceThreadId sets SourceThreadId field to given value.

### HasSourceThreadId

`func (o *MessageCreate) HasSourceThreadId() bool`

HasSourceThreadId returns a boolean if a field has been set.

### SetSourceThreadIdNil

`func (o *MessageCreate) SetSourceThreadIdNil(b bool)`

 SetSourceThreadIdNil sets the value for SourceThreadId to be an explicit nil

### UnsetSourceThreadId
`func (o *MessageCreate) UnsetSourceThreadId()`

UnsetSourceThreadId ensures that no value is present for SourceThreadId, not even an explicit nil
### GetSourceReplyToMessageId

`func (o *MessageCreate) GetSourceReplyToMessageId() string`

GetSourceReplyToMessageId returns the SourceReplyToMessageId field if non-nil, zero value otherwise.

### GetSourceReplyToMessageIdOk

`func (o *MessageCreate) GetSourceReplyToMessageIdOk() (*string, bool)`

GetSourceReplyToMessageIdOk returns a tuple with the SourceReplyToMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceReplyToMessageId

`func (o *MessageCreate) SetSourceReplyToMessageId(v string)`

SetSourceReplyToMessageId sets SourceReplyToMessageId field to given value.

### HasSourceReplyToMessageId

`func (o *MessageCreate) HasSourceReplyToMessageId() bool`

HasSourceReplyToMessageId returns a boolean if a field has been set.

### SetSourceReplyToMessageIdNil

`func (o *MessageCreate) SetSourceReplyToMessageIdNil(b bool)`

 SetSourceReplyToMessageIdNil sets the value for SourceReplyToMessageId to be an explicit nil

### UnsetSourceReplyToMessageId
`func (o *MessageCreate) UnsetSourceReplyToMessageId()`

UnsetSourceReplyToMessageId ensures that no value is present for SourceReplyToMessageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


