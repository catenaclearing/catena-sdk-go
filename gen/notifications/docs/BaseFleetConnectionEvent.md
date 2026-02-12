# BaseFleetConnectionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the fleet connection | 
**FleetId** | **string** | The ID of the fleet | 
**ConnectionId** | **string** | The ID of the connection | 

## Methods

### NewBaseFleetConnectionEvent

`func NewBaseFleetConnectionEvent(id string, fleetId string, connectionId string, ) *BaseFleetConnectionEvent`

NewBaseFleetConnectionEvent instantiates a new BaseFleetConnectionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseFleetConnectionEventWithDefaults

`func NewBaseFleetConnectionEventWithDefaults() *BaseFleetConnectionEvent`

NewBaseFleetConnectionEventWithDefaults instantiates a new BaseFleetConnectionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseFleetConnectionEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseFleetConnectionEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseFleetConnectionEvent) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseFleetConnectionEvent) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseFleetConnectionEvent) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseFleetConnectionEvent) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetConnectionId

`func (o *BaseFleetConnectionEvent) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseFleetConnectionEvent) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseFleetConnectionEvent) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


