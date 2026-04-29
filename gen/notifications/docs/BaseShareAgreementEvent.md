# BaseShareAgreementEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the share agreement | 
**FleetId** | **string** | The Catena ID of the fleet. | 
**PartnerId** | **string** | The Catena ID of the partner. | 
**FleetRef** | Pointer to **NullableString** |  | [optional] 
**Status** | [**StatusEnum**](StatusEnum.md) | The status of the share agreement | 
**EffectiveDate** | Pointer to **NullableTime** |  | [optional] 
**ExpirationDate** | Pointer to **NullableTime** |  | [optional] 
**Scopes** | **map[string]string** | The scopes/resources that are shared between the fleet and the partner | 

## Methods

### NewBaseShareAgreementEvent

`func NewBaseShareAgreementEvent(id string, fleetId string, partnerId string, status StatusEnum, scopes map[string]string, ) *BaseShareAgreementEvent`

NewBaseShareAgreementEvent instantiates a new BaseShareAgreementEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseShareAgreementEventWithDefaults

`func NewBaseShareAgreementEventWithDefaults() *BaseShareAgreementEvent`

NewBaseShareAgreementEventWithDefaults instantiates a new BaseShareAgreementEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseShareAgreementEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseShareAgreementEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseShareAgreementEvent) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseShareAgreementEvent) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseShareAgreementEvent) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseShareAgreementEvent) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetPartnerId

`func (o *BaseShareAgreementEvent) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *BaseShareAgreementEvent) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *BaseShareAgreementEvent) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.


### GetFleetRef

`func (o *BaseShareAgreementEvent) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *BaseShareAgreementEvent) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *BaseShareAgreementEvent) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *BaseShareAgreementEvent) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *BaseShareAgreementEvent) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *BaseShareAgreementEvent) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetStatus

`func (o *BaseShareAgreementEvent) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseShareAgreementEvent) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseShareAgreementEvent) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetEffectiveDate

`func (o *BaseShareAgreementEvent) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *BaseShareAgreementEvent) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *BaseShareAgreementEvent) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *BaseShareAgreementEvent) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### SetEffectiveDateNil

`func (o *BaseShareAgreementEvent) SetEffectiveDateNil(b bool)`

 SetEffectiveDateNil sets the value for EffectiveDate to be an explicit nil

### UnsetEffectiveDate
`func (o *BaseShareAgreementEvent) UnsetEffectiveDate()`

UnsetEffectiveDate ensures that no value is present for EffectiveDate, not even an explicit nil
### GetExpirationDate

`func (o *BaseShareAgreementEvent) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *BaseShareAgreementEvent) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *BaseShareAgreementEvent) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *BaseShareAgreementEvent) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *BaseShareAgreementEvent) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *BaseShareAgreementEvent) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetScopes

`func (o *BaseShareAgreementEvent) GetScopes() map[string]string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *BaseShareAgreementEvent) GetScopesOk() (*map[string]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *BaseShareAgreementEvent) SetScopes(v map[string]string)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


