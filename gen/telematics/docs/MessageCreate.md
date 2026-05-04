# MessageCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | Unique identifier of the connection at Catena Telematics which will be used to create this resource. A connection represents a Fleet/TSP pairing. | 
**SenderId** | **string** | Identifier for the sender | 
**RecipientId** | **string** | Identifier for the recipient | 
**MessageText** | Pointer to **NullableString** |  | [optional] 
**Priority** | Pointer to [**NullableMessagePriorityEnum**](MessagePriorityEnum.md) |  | [optional] 
**ReplyToMessageId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewMessageCreate

`func NewMessageCreate(connectionId string, senderId string, recipientId string, ) *MessageCreate`

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


### GetSenderId

`func (o *MessageCreate) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *MessageCreate) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *MessageCreate) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.


### GetRecipientId

`func (o *MessageCreate) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *MessageCreate) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *MessageCreate) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.


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
### GetReplyToMessageId

`func (o *MessageCreate) GetReplyToMessageId() string`

GetReplyToMessageId returns the ReplyToMessageId field if non-nil, zero value otherwise.

### GetReplyToMessageIdOk

`func (o *MessageCreate) GetReplyToMessageIdOk() (*string, bool)`

GetReplyToMessageIdOk returns a tuple with the ReplyToMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToMessageId

`func (o *MessageCreate) SetReplyToMessageId(v string)`

SetReplyToMessageId sets ReplyToMessageId field to given value.

### HasReplyToMessageId

`func (o *MessageCreate) HasReplyToMessageId() bool`

HasReplyToMessageId returns a boolean if a field has been set.

### SetReplyToMessageIdNil

`func (o *MessageCreate) SetReplyToMessageIdNil(b bool)`

 SetReplyToMessageIdNil sets the value for ReplyToMessageId to be an explicit nil

### UnsetReplyToMessageId
`func (o *MessageCreate) UnsetReplyToMessageId()`

UnsetReplyToMessageId ensures that no value is present for ReplyToMessageId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


