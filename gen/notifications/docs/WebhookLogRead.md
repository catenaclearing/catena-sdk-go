# WebhookLogRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | The date and time the log was created | 
**WebhookId** | **string** | The ID of the webhook that the metrics are for | 
**MessageId** | **string** | The ID of the message | 
**EventName** | [**WebhookEventNameEnum**](WebhookEventNameEnum.md) | The name of the event | 
**Status** | [**WebhookLogStatusEnum**](WebhookLogStatusEnum.md) | The status of the webhook message | 
**StatusCode** | **int32** | The HTTP status code of the webhook delivery | 
**ErrorMessage** | **NullableString** |  | 
**ResponseTimeMs** | **NullableInt32** |  | 

## Methods

### NewWebhookLogRead

`func NewWebhookLogRead(createdAt time.Time, webhookId string, messageId string, eventName WebhookEventNameEnum, status WebhookLogStatusEnum, statusCode int32, errorMessage NullableString, responseTimeMs NullableInt32, ) *WebhookLogRead`

NewWebhookLogRead instantiates a new WebhookLogRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookLogReadWithDefaults

`func NewWebhookLogReadWithDefaults() *WebhookLogRead`

NewWebhookLogReadWithDefaults instantiates a new WebhookLogRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WebhookLogRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookLogRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookLogRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetWebhookId

`func (o *WebhookLogRead) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookLogRead) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookLogRead) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetMessageId

`func (o *WebhookLogRead) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *WebhookLogRead) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *WebhookLogRead) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.


### GetEventName

`func (o *WebhookLogRead) GetEventName() WebhookEventNameEnum`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *WebhookLogRead) GetEventNameOk() (*WebhookEventNameEnum, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *WebhookLogRead) SetEventName(v WebhookEventNameEnum)`

SetEventName sets EventName field to given value.


### GetStatus

`func (o *WebhookLogRead) GetStatus() WebhookLogStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookLogRead) GetStatusOk() (*WebhookLogStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookLogRead) SetStatus(v WebhookLogStatusEnum)`

SetStatus sets Status field to given value.


### GetStatusCode

`func (o *WebhookLogRead) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *WebhookLogRead) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *WebhookLogRead) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetErrorMessage

`func (o *WebhookLogRead) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *WebhookLogRead) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *WebhookLogRead) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *WebhookLogRead) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *WebhookLogRead) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetResponseTimeMs

`func (o *WebhookLogRead) GetResponseTimeMs() int32`

GetResponseTimeMs returns the ResponseTimeMs field if non-nil, zero value otherwise.

### GetResponseTimeMsOk

`func (o *WebhookLogRead) GetResponseTimeMsOk() (*int32, bool)`

GetResponseTimeMsOk returns a tuple with the ResponseTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseTimeMs

`func (o *WebhookLogRead) SetResponseTimeMs(v int32)`

SetResponseTimeMs sets ResponseTimeMs field to given value.


### SetResponseTimeMsNil

`func (o *WebhookLogRead) SetResponseTimeMsNil(b bool)`

 SetResponseTimeMsNil sets the value for ResponseTimeMs to be an explicit nil

### UnsetResponseTimeMs
`func (o *WebhookLogRead) UnsetResponseTimeMs()`

UnsetResponseTimeMs ensures that no value is present for ResponseTimeMs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


