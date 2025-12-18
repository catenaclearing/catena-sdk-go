# ShareAgreementUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**NullableStatusEnum**](StatusEnum.md) |  | [optional] 
**FleetRef** | Pointer to **NullableString** |  | [optional] 
**EffectiveDate** | Pointer to **NullableTime** |  | [optional] 
**ExpirationDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewShareAgreementUpdate

`func NewShareAgreementUpdate() *ShareAgreementUpdate`

NewShareAgreementUpdate instantiates a new ShareAgreementUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareAgreementUpdateWithDefaults

`func NewShareAgreementUpdateWithDefaults() *ShareAgreementUpdate`

NewShareAgreementUpdateWithDefaults instantiates a new ShareAgreementUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ShareAgreementUpdate) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShareAgreementUpdate) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShareAgreementUpdate) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ShareAgreementUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ShareAgreementUpdate) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ShareAgreementUpdate) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetFleetRef

`func (o *ShareAgreementUpdate) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *ShareAgreementUpdate) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *ShareAgreementUpdate) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *ShareAgreementUpdate) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *ShareAgreementUpdate) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *ShareAgreementUpdate) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetEffectiveDate

`func (o *ShareAgreementUpdate) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *ShareAgreementUpdate) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *ShareAgreementUpdate) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *ShareAgreementUpdate) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### SetEffectiveDateNil

`func (o *ShareAgreementUpdate) SetEffectiveDateNil(b bool)`

 SetEffectiveDateNil sets the value for EffectiveDate to be an explicit nil

### UnsetEffectiveDate
`func (o *ShareAgreementUpdate) UnsetEffectiveDate()`

UnsetEffectiveDate ensures that no value is present for EffectiveDate, not even an explicit nil
### GetExpirationDate

`func (o *ShareAgreementUpdate) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ShareAgreementUpdate) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ShareAgreementUpdate) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ShareAgreementUpdate) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *ShareAgreementUpdate) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *ShareAgreementUpdate) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


