# FleetPropertyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**FleetPropertyKeyEnum**](FleetPropertyKeyEnum.md) | The key of the property | 
**Value** | **string** | The value of the property | 

## Methods

### NewFleetPropertyCreate

`func NewFleetPropertyCreate(key FleetPropertyKeyEnum, value string, ) *FleetPropertyCreate`

NewFleetPropertyCreate instantiates a new FleetPropertyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetPropertyCreateWithDefaults

`func NewFleetPropertyCreateWithDefaults() *FleetPropertyCreate`

NewFleetPropertyCreateWithDefaults instantiates a new FleetPropertyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FleetPropertyCreate) GetKey() FleetPropertyKeyEnum`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FleetPropertyCreate) GetKeyOk() (*FleetPropertyKeyEnum, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FleetPropertyCreate) SetKey(v FleetPropertyKeyEnum)`

SetKey sets Key field to given value.


### GetValue

`func (o *FleetPropertyCreate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FleetPropertyCreate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FleetPropertyCreate) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


