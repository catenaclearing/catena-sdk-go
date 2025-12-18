# ShareAgreementCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerId** | **string** | Your organization ID receiving access to fleet data | 
**Scopes** | [**map[string]ShareLevelEnum**](ShareLevelEnum.md) | Defines which resources (vehicle, locations, users, etc.) you can access and the permission level (read, write) for each. | 
**FleetId** | Pointer to **NullableString** |  | [optional] 
**FleetRef** | Pointer to **NullableString** |  | [optional] 
**InvitationId** | Pointer to **NullableString** |  | [optional] 
**EffectiveDate** | Pointer to **NullableTime** |  | [optional] 
**ExpirationDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewShareAgreementCreate

`func NewShareAgreementCreate(partnerId string, scopes map[string]ShareLevelEnum, ) *ShareAgreementCreate`

NewShareAgreementCreate instantiates a new ShareAgreementCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareAgreementCreateWithDefaults

`func NewShareAgreementCreateWithDefaults() *ShareAgreementCreate`

NewShareAgreementCreateWithDefaults instantiates a new ShareAgreementCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerId

`func (o *ShareAgreementCreate) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ShareAgreementCreate) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ShareAgreementCreate) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.


### GetScopes

`func (o *ShareAgreementCreate) GetScopes() map[string]ShareLevelEnum`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ShareAgreementCreate) GetScopesOk() (*map[string]ShareLevelEnum, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ShareAgreementCreate) SetScopes(v map[string]ShareLevelEnum)`

SetScopes sets Scopes field to given value.


### GetFleetId

`func (o *ShareAgreementCreate) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *ShareAgreementCreate) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *ShareAgreementCreate) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.

### HasFleetId

`func (o *ShareAgreementCreate) HasFleetId() bool`

HasFleetId returns a boolean if a field has been set.

### SetFleetIdNil

`func (o *ShareAgreementCreate) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *ShareAgreementCreate) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetFleetRef

`func (o *ShareAgreementCreate) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *ShareAgreementCreate) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *ShareAgreementCreate) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *ShareAgreementCreate) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *ShareAgreementCreate) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *ShareAgreementCreate) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetInvitationId

`func (o *ShareAgreementCreate) GetInvitationId() string`

GetInvitationId returns the InvitationId field if non-nil, zero value otherwise.

### GetInvitationIdOk

`func (o *ShareAgreementCreate) GetInvitationIdOk() (*string, bool)`

GetInvitationIdOk returns a tuple with the InvitationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationId

`func (o *ShareAgreementCreate) SetInvitationId(v string)`

SetInvitationId sets InvitationId field to given value.

### HasInvitationId

`func (o *ShareAgreementCreate) HasInvitationId() bool`

HasInvitationId returns a boolean if a field has been set.

### SetInvitationIdNil

`func (o *ShareAgreementCreate) SetInvitationIdNil(b bool)`

 SetInvitationIdNil sets the value for InvitationId to be an explicit nil

### UnsetInvitationId
`func (o *ShareAgreementCreate) UnsetInvitationId()`

UnsetInvitationId ensures that no value is present for InvitationId, not even an explicit nil
### GetEffectiveDate

`func (o *ShareAgreementCreate) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *ShareAgreementCreate) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *ShareAgreementCreate) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *ShareAgreementCreate) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### SetEffectiveDateNil

`func (o *ShareAgreementCreate) SetEffectiveDateNil(b bool)`

 SetEffectiveDateNil sets the value for EffectiveDate to be an explicit nil

### UnsetEffectiveDate
`func (o *ShareAgreementCreate) UnsetEffectiveDate()`

UnsetEffectiveDate ensures that no value is present for EffectiveDate, not even an explicit nil
### GetExpirationDate

`func (o *ShareAgreementCreate) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ShareAgreementCreate) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ShareAgreementCreate) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ShareAgreementCreate) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *ShareAgreementCreate) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *ShareAgreementCreate) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


