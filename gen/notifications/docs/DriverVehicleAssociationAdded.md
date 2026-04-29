# DriverVehicleAssociationAdded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** | Version of the schema used for this event | 
**EventName** | Pointer to **string** |  | [optional] [default to "driver_vehicle_association.added"]
**Data** | [**[]BaseDriverVehicleAssociation**](BaseDriverVehicleAssociation.md) | The event payload, with one or more records of the specified type | 
**WebhookId** | **string** | Unique identifier for the webhook subscription that triggered this delivery | 
**Timestamp** | Pointer to **NullableTime** |  | [optional] 
**Id** | Pointer to **NullableString** |  | [optional] 
**DeliveryAttempt** | Pointer to **int32** | The number of times the event has been attempted to be delivered | [optional] [default to 0]

## Methods

### NewDriverVehicleAssociationAdded

`func NewDriverVehicleAssociationAdded(version string, data []BaseDriverVehicleAssociation, webhookId string, ) *DriverVehicleAssociationAdded`

NewDriverVehicleAssociationAdded instantiates a new DriverVehicleAssociationAdded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverVehicleAssociationAddedWithDefaults

`func NewDriverVehicleAssociationAddedWithDefaults() *DriverVehicleAssociationAdded`

NewDriverVehicleAssociationAddedWithDefaults instantiates a new DriverVehicleAssociationAdded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *DriverVehicleAssociationAdded) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DriverVehicleAssociationAdded) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DriverVehicleAssociationAdded) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetEventName

`func (o *DriverVehicleAssociationAdded) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *DriverVehicleAssociationAdded) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *DriverVehicleAssociationAdded) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *DriverVehicleAssociationAdded) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetData

`func (o *DriverVehicleAssociationAdded) GetData() []BaseDriverVehicleAssociation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DriverVehicleAssociationAdded) GetDataOk() (*[]BaseDriverVehicleAssociation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DriverVehicleAssociationAdded) SetData(v []BaseDriverVehicleAssociation)`

SetData sets Data field to given value.


### GetWebhookId

`func (o *DriverVehicleAssociationAdded) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *DriverVehicleAssociationAdded) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *DriverVehicleAssociationAdded) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetTimestamp

`func (o *DriverVehicleAssociationAdded) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DriverVehicleAssociationAdded) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DriverVehicleAssociationAdded) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DriverVehicleAssociationAdded) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *DriverVehicleAssociationAdded) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *DriverVehicleAssociationAdded) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetId

`func (o *DriverVehicleAssociationAdded) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DriverVehicleAssociationAdded) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DriverVehicleAssociationAdded) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DriverVehicleAssociationAdded) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *DriverVehicleAssociationAdded) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *DriverVehicleAssociationAdded) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDeliveryAttempt

`func (o *DriverVehicleAssociationAdded) GetDeliveryAttempt() int32`

GetDeliveryAttempt returns the DeliveryAttempt field if non-nil, zero value otherwise.

### GetDeliveryAttemptOk

`func (o *DriverVehicleAssociationAdded) GetDeliveryAttemptOk() (*int32, bool)`

GetDeliveryAttemptOk returns a tuple with the DeliveryAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAttempt

`func (o *DriverVehicleAssociationAdded) SetDeliveryAttempt(v int32)`

SetDeliveryAttempt sets DeliveryAttempt field to given value.

### HasDeliveryAttempt

`func (o *DriverVehicleAssociationAdded) HasDeliveryAttempt() bool`

HasDeliveryAttempt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


