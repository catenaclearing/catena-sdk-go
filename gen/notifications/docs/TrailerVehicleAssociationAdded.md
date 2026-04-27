# TrailerVehicleAssociationAdded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Version of the schema used for this event | 
**EventName** | Pointer to **string** |  | [optional] [default to "trailer_vehicle_association.added"]
**Data** | [**[]BaseTrailerVehicleAssociation**](BaseTrailerVehicleAssociation.md) | The event payload, with one or more records of the specified type | 
**WebhookId** | **string** | Unique identifier for the webhook subscription that triggered this delivery | 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**DeliveryAttempt** | Pointer to **int32** | The number of times the event has been attempted to be delivered | [optional] [default to 0]

## Methods

### NewTrailerVehicleAssociationAdded

`func NewTrailerVehicleAssociationAdded(version string, data []BaseTrailerVehicleAssociation, webhookId string, ) *TrailerVehicleAssociationAdded`

NewTrailerVehicleAssociationAdded instantiates a new TrailerVehicleAssociationAdded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailerVehicleAssociationAddedWithDefaults

`func NewTrailerVehicleAssociationAddedWithDefaults() *TrailerVehicleAssociationAdded`

NewTrailerVehicleAssociationAddedWithDefaults instantiates a new TrailerVehicleAssociationAdded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *TrailerVehicleAssociationAdded) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TrailerVehicleAssociationAdded) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TrailerVehicleAssociationAdded) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetEventName

`func (o *TrailerVehicleAssociationAdded) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *TrailerVehicleAssociationAdded) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *TrailerVehicleAssociationAdded) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *TrailerVehicleAssociationAdded) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetData

`func (o *TrailerVehicleAssociationAdded) GetData() []BaseTrailerVehicleAssociation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TrailerVehicleAssociationAdded) GetDataOk() (*[]BaseTrailerVehicleAssociation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TrailerVehicleAssociationAdded) SetData(v []BaseTrailerVehicleAssociation)`

SetData sets Data field to given value.


### GetWebhookId

`func (o *TrailerVehicleAssociationAdded) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *TrailerVehicleAssociationAdded) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *TrailerVehicleAssociationAdded) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetTimestamp

`func (o *TrailerVehicleAssociationAdded) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TrailerVehicleAssociationAdded) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TrailerVehicleAssociationAdded) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *TrailerVehicleAssociationAdded) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *TrailerVehicleAssociationAdded) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *TrailerVehicleAssociationAdded) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetId

`func (o *TrailerVehicleAssociationAdded) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrailerVehicleAssociationAdded) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrailerVehicleAssociationAdded) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrailerVehicleAssociationAdded) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *TrailerVehicleAssociationAdded) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *TrailerVehicleAssociationAdded) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDeliveryAttempt

`func (o *TrailerVehicleAssociationAdded) GetDeliveryAttempt() int32`

GetDeliveryAttempt returns the DeliveryAttempt field if non-nil, zero value otherwise.

### GetDeliveryAttemptOk

`func (o *TrailerVehicleAssociationAdded) GetDeliveryAttemptOk() (*int32, bool)`

GetDeliveryAttemptOk returns a tuple with the DeliveryAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAttempt

`func (o *TrailerVehicleAssociationAdded) SetDeliveryAttempt(v int32)`

SetDeliveryAttempt sets DeliveryAttempt field to given value.

### HasDeliveryAttempt

`func (o *TrailerVehicleAssociationAdded) HasDeliveryAttempt() bool`

HasDeliveryAttempt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


