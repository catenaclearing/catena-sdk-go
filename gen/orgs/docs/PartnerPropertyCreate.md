# PartnerPropertyCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**PartnerPropertyKeyEnum**](PartnerPropertyKeyEnum.md) | The key of the property | 
**Value** | **string** | The value of the property | 

## Methods

### NewPartnerPropertyCreate

`func NewPartnerPropertyCreate(key PartnerPropertyKeyEnum, value string, ) *PartnerPropertyCreate`

NewPartnerPropertyCreate instantiates a new PartnerPropertyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerPropertyCreateWithDefaults

`func NewPartnerPropertyCreateWithDefaults() *PartnerPropertyCreate`

NewPartnerPropertyCreateWithDefaults instantiates a new PartnerPropertyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PartnerPropertyCreate) GetKey() PartnerPropertyKeyEnum`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PartnerPropertyCreate) GetKeyOk() (*PartnerPropertyKeyEnum, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PartnerPropertyCreate) SetKey(v PartnerPropertyKeyEnum)`

SetKey sets Key field to given value.


### GetValue

`func (o *PartnerPropertyCreate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PartnerPropertyCreate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PartnerPropertyCreate) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


