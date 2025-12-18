# IftaSummaryAdded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Version of the schema used for this event | 
**EventName** | Pointer to **string** |  | [optional] [default to "ifta_summary.added"]
**Data** | [**[]BaseIftaSummary**](BaseIftaSummary.md) | The event payload, with one or more records of the specified type | 
**WebhookId** | **string** | Unique identifier for the webhook subscription that triggered this delivery | 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**DeliveryAttempt** | Pointer to **int32** | The number of times the event has been attempted to be delivered | [optional] [default to 0]

## Methods

### NewIftaSummaryAdded

`func NewIftaSummaryAdded(version string, data []BaseIftaSummary, webhookId string, ) *IftaSummaryAdded`

NewIftaSummaryAdded instantiates a new IftaSummaryAdded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIftaSummaryAddedWithDefaults

`func NewIftaSummaryAddedWithDefaults() *IftaSummaryAdded`

NewIftaSummaryAddedWithDefaults instantiates a new IftaSummaryAdded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *IftaSummaryAdded) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IftaSummaryAdded) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IftaSummaryAdded) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetEventName

`func (o *IftaSummaryAdded) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *IftaSummaryAdded) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *IftaSummaryAdded) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *IftaSummaryAdded) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetData

`func (o *IftaSummaryAdded) GetData() []BaseIftaSummary`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IftaSummaryAdded) GetDataOk() (*[]BaseIftaSummary, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IftaSummaryAdded) SetData(v []BaseIftaSummary)`

SetData sets Data field to given value.


### GetWebhookId

`func (o *IftaSummaryAdded) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *IftaSummaryAdded) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *IftaSummaryAdded) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetTimestamp

`func (o *IftaSummaryAdded) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *IftaSummaryAdded) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *IftaSummaryAdded) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *IftaSummaryAdded) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *IftaSummaryAdded) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *IftaSummaryAdded) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetId

`func (o *IftaSummaryAdded) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IftaSummaryAdded) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IftaSummaryAdded) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IftaSummaryAdded) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IftaSummaryAdded) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IftaSummaryAdded) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDeliveryAttempt

`func (o *IftaSummaryAdded) GetDeliveryAttempt() int32`

GetDeliveryAttempt returns the DeliveryAttempt field if non-nil, zero value otherwise.

### GetDeliveryAttemptOk

`func (o *IftaSummaryAdded) GetDeliveryAttemptOk() (*int32, bool)`

GetDeliveryAttemptOk returns a tuple with the DeliveryAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAttempt

`func (o *IftaSummaryAdded) SetDeliveryAttempt(v int32)`

SetDeliveryAttempt sets DeliveryAttempt field to given value.

### HasDeliveryAttempt

`func (o *IftaSummaryAdded) HasDeliveryAttempt() bool`

HasDeliveryAttempt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


