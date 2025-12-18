# FleetPropertyRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**FleetPropertyKeyEnum**](FleetPropertyKeyEnum.md) | The key of the property | 
**Value** | **string** | The value of the property | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**FleetId** | **string** |  | 

## Methods

### NewFleetPropertyRead

`func NewFleetPropertyRead(key FleetPropertyKeyEnum, value string, id string, createdAt time.Time, fleetId string, ) *FleetPropertyRead`

NewFleetPropertyRead instantiates a new FleetPropertyRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetPropertyReadWithDefaults

`func NewFleetPropertyReadWithDefaults() *FleetPropertyRead`

NewFleetPropertyReadWithDefaults instantiates a new FleetPropertyRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FleetPropertyRead) GetKey() FleetPropertyKeyEnum`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FleetPropertyRead) GetKeyOk() (*FleetPropertyKeyEnum, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FleetPropertyRead) SetKey(v FleetPropertyKeyEnum)`

SetKey sets Key field to given value.


### GetValue

`func (o *FleetPropertyRead) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FleetPropertyRead) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FleetPropertyRead) SetValue(v string)`

SetValue sets Value field to given value.


### GetId

`func (o *FleetPropertyRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetPropertyRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetPropertyRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *FleetPropertyRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FleetPropertyRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FleetPropertyRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFleetId

`func (o *FleetPropertyRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *FleetPropertyRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *FleetPropertyRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


