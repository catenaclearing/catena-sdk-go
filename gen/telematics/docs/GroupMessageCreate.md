# GroupMessageCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | Unique identifier of the connection at Catena Telematics which will be used to create this resource. A connection represents a Fleet/TSP pairing. | 
**SenderId** | **string** | Identifier for the sender | 
**RecipientIds** | **[]string** | Identifiers for the recipients | 
**MessageText** | Pointer to **NullableString** |  | [optional] 
**Priority** | Pointer to [**NullableMessagePriorityEnum**](MessagePriorityEnum.md) |  | [optional] 

## Methods

### NewGroupMessageCreate

`func NewGroupMessageCreate(connectionId string, senderId string, recipientIds []string, ) *GroupMessageCreate`

NewGroupMessageCreate instantiates a new GroupMessageCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMessageCreateWithDefaults

`func NewGroupMessageCreateWithDefaults() *GroupMessageCreate`

NewGroupMessageCreateWithDefaults instantiates a new GroupMessageCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *GroupMessageCreate) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *GroupMessageCreate) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *GroupMessageCreate) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSenderId

`func (o *GroupMessageCreate) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *GroupMessageCreate) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *GroupMessageCreate) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.


### GetRecipientIds

`func (o *GroupMessageCreate) GetRecipientIds() []string`

GetRecipientIds returns the RecipientIds field if non-nil, zero value otherwise.

### GetRecipientIdsOk

`func (o *GroupMessageCreate) GetRecipientIdsOk() (*[]string, bool)`

GetRecipientIdsOk returns a tuple with the RecipientIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIds

`func (o *GroupMessageCreate) SetRecipientIds(v []string)`

SetRecipientIds sets RecipientIds field to given value.


### GetMessageText

`func (o *GroupMessageCreate) GetMessageText() string`

GetMessageText returns the MessageText field if non-nil, zero value otherwise.

### GetMessageTextOk

`func (o *GroupMessageCreate) GetMessageTextOk() (*string, bool)`

GetMessageTextOk returns a tuple with the MessageText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageText

`func (o *GroupMessageCreate) SetMessageText(v string)`

SetMessageText sets MessageText field to given value.

### HasMessageText

`func (o *GroupMessageCreate) HasMessageText() bool`

HasMessageText returns a boolean if a field has been set.

### SetMessageTextNil

`func (o *GroupMessageCreate) SetMessageTextNil(b bool)`

 SetMessageTextNil sets the value for MessageText to be an explicit nil

### UnsetMessageText
`func (o *GroupMessageCreate) UnsetMessageText()`

UnsetMessageText ensures that no value is present for MessageText, not even an explicit nil
### GetPriority

`func (o *GroupMessageCreate) GetPriority() MessagePriorityEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GroupMessageCreate) GetPriorityOk() (*MessagePriorityEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GroupMessageCreate) SetPriority(v MessagePriorityEnum)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GroupMessageCreate) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *GroupMessageCreate) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *GroupMessageCreate) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


