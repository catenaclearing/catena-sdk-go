# BaseConnectionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the connection | 
**SourceName** | [**TspEnum**](TspEnum.md) | The name of the source | 
**FleetId** | **string** | The ID of the fleet | 
**TspId** | **string** | The ID of the Telematics Service Provider (TSP) at Catena | 
**Status** | [**StatusEnum**](StatusEnum.md) | The status of the connection, used to determine if the connection is active or stale. | 

## Methods

### NewBaseConnectionEvent

`func NewBaseConnectionEvent(id string, sourceName TspEnum, fleetId string, tspId string, status StatusEnum, ) *BaseConnectionEvent`

NewBaseConnectionEvent instantiates a new BaseConnectionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseConnectionEventWithDefaults

`func NewBaseConnectionEventWithDefaults() *BaseConnectionEvent`

NewBaseConnectionEventWithDefaults instantiates a new BaseConnectionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseConnectionEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseConnectionEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseConnectionEvent) SetId(v string)`

SetId sets Id field to given value.


### GetSourceName

`func (o *BaseConnectionEvent) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseConnectionEvent) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseConnectionEvent) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetFleetId

`func (o *BaseConnectionEvent) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseConnectionEvent) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseConnectionEvent) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetTspId

`func (o *BaseConnectionEvent) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *BaseConnectionEvent) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *BaseConnectionEvent) SetTspId(v string)`

SetTspId sets TspId field to given value.


### GetStatus

`func (o *BaseConnectionEvent) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseConnectionEvent) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseConnectionEvent) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


