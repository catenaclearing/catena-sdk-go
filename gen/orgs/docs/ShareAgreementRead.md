# ShareAgreementRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique agreement identifier | 
**FleetId** | **string** | The Catena fleet ID sharing their data with you | 
**PartnerId** | **string** | Your organization ID receiving access to fleet data | 
**FleetRef** | **NullableString** |  | 
**InvitationId** | **NullableString** |  | 
**Status** | [**StatusEnum**](StatusEnum.md) | Current state: ACTIVE (data access enabled), PAUSED (temporarily disabled), CANCELLED (permanently ended), EXPIRED (past expiration_date) | 
**EffectiveDate** | **time.Time** | When data access begins. Check this before attempting to fetch fleet data. | 
**ExpirationDate** | **NullableTime** |  | 
**Scopes** | [**map[string]ShareLevelEnum**](ShareLevelEnum.md) | Defines which resources (vehicle, locations, users, etc.) you can access and the permission level (read, write) for each. | 

## Methods

### NewShareAgreementRead

`func NewShareAgreementRead(id string, fleetId string, partnerId string, fleetRef NullableString, invitationId NullableString, status StatusEnum, effectiveDate time.Time, expirationDate NullableTime, scopes map[string]ShareLevelEnum, ) *ShareAgreementRead`

NewShareAgreementRead instantiates a new ShareAgreementRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareAgreementReadWithDefaults

`func NewShareAgreementReadWithDefaults() *ShareAgreementRead`

NewShareAgreementReadWithDefaults instantiates a new ShareAgreementRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShareAgreementRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShareAgreementRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShareAgreementRead) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *ShareAgreementRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *ShareAgreementRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *ShareAgreementRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetPartnerId

`func (o *ShareAgreementRead) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ShareAgreementRead) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ShareAgreementRead) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.


### GetFleetRef

`func (o *ShareAgreementRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *ShareAgreementRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *ShareAgreementRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### SetFleetRefNil

`func (o *ShareAgreementRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *ShareAgreementRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetInvitationId

`func (o *ShareAgreementRead) GetInvitationId() string`

GetInvitationId returns the InvitationId field if non-nil, zero value otherwise.

### GetInvitationIdOk

`func (o *ShareAgreementRead) GetInvitationIdOk() (*string, bool)`

GetInvitationIdOk returns a tuple with the InvitationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationId

`func (o *ShareAgreementRead) SetInvitationId(v string)`

SetInvitationId sets InvitationId field to given value.


### SetInvitationIdNil

`func (o *ShareAgreementRead) SetInvitationIdNil(b bool)`

 SetInvitationIdNil sets the value for InvitationId to be an explicit nil

### UnsetInvitationId
`func (o *ShareAgreementRead) UnsetInvitationId()`

UnsetInvitationId ensures that no value is present for InvitationId, not even an explicit nil
### GetStatus

`func (o *ShareAgreementRead) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShareAgreementRead) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShareAgreementRead) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetEffectiveDate

`func (o *ShareAgreementRead) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *ShareAgreementRead) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *ShareAgreementRead) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.


### GetExpirationDate

`func (o *ShareAgreementRead) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ShareAgreementRead) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ShareAgreementRead) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.


### SetExpirationDateNil

`func (o *ShareAgreementRead) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *ShareAgreementRead) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetScopes

`func (o *ShareAgreementRead) GetScopes() map[string]ShareLevelEnum`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ShareAgreementRead) GetScopesOk() (*map[string]ShareLevelEnum, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ShareAgreementRead) SetScopes(v map[string]ShareLevelEnum)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


