# PartnerPropertyRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**PartnerPropertyKeyEnum**](PartnerPropertyKeyEnum.md) | The key of the property | 
**Value** | **string** | The value of the property | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**PartnerId** | **string** |  | 

## Methods

### NewPartnerPropertyRead

`func NewPartnerPropertyRead(key PartnerPropertyKeyEnum, value string, id string, createdAt time.Time, partnerId string, ) *PartnerPropertyRead`

NewPartnerPropertyRead instantiates a new PartnerPropertyRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerPropertyReadWithDefaults

`func NewPartnerPropertyReadWithDefaults() *PartnerPropertyRead`

NewPartnerPropertyReadWithDefaults instantiates a new PartnerPropertyRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PartnerPropertyRead) GetKey() PartnerPropertyKeyEnum`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PartnerPropertyRead) GetKeyOk() (*PartnerPropertyKeyEnum, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PartnerPropertyRead) SetKey(v PartnerPropertyKeyEnum)`

SetKey sets Key field to given value.


### GetValue

`func (o *PartnerPropertyRead) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PartnerPropertyRead) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PartnerPropertyRead) SetValue(v string)`

SetValue sets Value field to given value.


### GetId

`func (o *PartnerPropertyRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartnerPropertyRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartnerPropertyRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *PartnerPropertyRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PartnerPropertyRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PartnerPropertyRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPartnerId

`func (o *PartnerPropertyRead) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *PartnerPropertyRead) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *PartnerPropertyRead) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


