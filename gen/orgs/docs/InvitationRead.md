# InvitationRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique invitation identifier | 
**CreatedAt** | **time.Time** | When the invitation was created | 
**MagicLink** | **string** | The magic link that is used to open Catena Connect and accept the invitation. Share this URL with the fleet to begin onboarding. | 
**ExpiresAt** | **time.Time** | The expiration date and time of the invitation | 
**ExpiresInHours** | **int32** | The number of hours the invitation is valid for. | 
**Status** | [**StatusEnum**](StatusEnum.md) | The current status of the invitation (ACTIVE, ACCEPTED, DECLINED, or EXPIRED) | 
**PartnerProvidedFleetName** | Pointer to **NullableString** |  | [optional] 
**AcceptedAt** | Pointer to **NullableTime** |  | [optional] 
**PreRegistrationAccessToken** | Pointer to **NullableString** |  | [optional] 
**PreRegistrationRefreshToken** | Pointer to **NullableString** |  | [optional] 
**CallbackUrl** | Pointer to **NullableString** |  | [optional] 
**SuccessRedirectUrl** | Pointer to **NullableString** |  | [optional] 
**FailureRedirectUrl** | Pointer to **NullableString** |  | [optional] 
**LimitTsps** | Pointer to **[]string** |  | [optional] 
**FleetId** | Pointer to **NullableString** |  | [optional] 
**PartnerSlug** | Pointer to **NullableString** |  | [optional] 
**PartnerId** | Pointer to **NullableString** |  | [optional] 
**DeclineReason** | Pointer to **NullableString** |  | [optional] 
**DeclinedAt** | Pointer to **NullableTime** |  | [optional] 
**FleetRef** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInvitationRead

`func NewInvitationRead(id string, createdAt time.Time, magicLink string, expiresAt time.Time, expiresInHours int32, status StatusEnum, ) *InvitationRead`

NewInvitationRead instantiates a new InvitationRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationReadWithDefaults

`func NewInvitationReadWithDefaults() *InvitationRead`

NewInvitationReadWithDefaults instantiates a new InvitationRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InvitationRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvitationRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvitationRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *InvitationRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InvitationRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InvitationRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMagicLink

`func (o *InvitationRead) GetMagicLink() string`

GetMagicLink returns the MagicLink field if non-nil, zero value otherwise.

### GetMagicLinkOk

`func (o *InvitationRead) GetMagicLinkOk() (*string, bool)`

GetMagicLinkOk returns a tuple with the MagicLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagicLink

`func (o *InvitationRead) SetMagicLink(v string)`

SetMagicLink sets MagicLink field to given value.


### GetExpiresAt

`func (o *InvitationRead) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InvitationRead) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InvitationRead) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetExpiresInHours

`func (o *InvitationRead) GetExpiresInHours() int32`

GetExpiresInHours returns the ExpiresInHours field if non-nil, zero value otherwise.

### GetExpiresInHoursOk

`func (o *InvitationRead) GetExpiresInHoursOk() (*int32, bool)`

GetExpiresInHoursOk returns a tuple with the ExpiresInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInHours

`func (o *InvitationRead) SetExpiresInHours(v int32)`

SetExpiresInHours sets ExpiresInHours field to given value.


### GetStatus

`func (o *InvitationRead) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvitationRead) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvitationRead) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetPartnerProvidedFleetName

`func (o *InvitationRead) GetPartnerProvidedFleetName() string`

GetPartnerProvidedFleetName returns the PartnerProvidedFleetName field if non-nil, zero value otherwise.

### GetPartnerProvidedFleetNameOk

`func (o *InvitationRead) GetPartnerProvidedFleetNameOk() (*string, bool)`

GetPartnerProvidedFleetNameOk returns a tuple with the PartnerProvidedFleetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProvidedFleetName

`func (o *InvitationRead) SetPartnerProvidedFleetName(v string)`

SetPartnerProvidedFleetName sets PartnerProvidedFleetName field to given value.

### HasPartnerProvidedFleetName

`func (o *InvitationRead) HasPartnerProvidedFleetName() bool`

HasPartnerProvidedFleetName returns a boolean if a field has been set.

### SetPartnerProvidedFleetNameNil

`func (o *InvitationRead) SetPartnerProvidedFleetNameNil(b bool)`

 SetPartnerProvidedFleetNameNil sets the value for PartnerProvidedFleetName to be an explicit nil

### UnsetPartnerProvidedFleetName
`func (o *InvitationRead) UnsetPartnerProvidedFleetName()`

UnsetPartnerProvidedFleetName ensures that no value is present for PartnerProvidedFleetName, not even an explicit nil
### GetAcceptedAt

`func (o *InvitationRead) GetAcceptedAt() time.Time`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *InvitationRead) GetAcceptedAtOk() (*time.Time, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *InvitationRead) SetAcceptedAt(v time.Time)`

SetAcceptedAt sets AcceptedAt field to given value.

### HasAcceptedAt

`func (o *InvitationRead) HasAcceptedAt() bool`

HasAcceptedAt returns a boolean if a field has been set.

### SetAcceptedAtNil

`func (o *InvitationRead) SetAcceptedAtNil(b bool)`

 SetAcceptedAtNil sets the value for AcceptedAt to be an explicit nil

### UnsetAcceptedAt
`func (o *InvitationRead) UnsetAcceptedAt()`

UnsetAcceptedAt ensures that no value is present for AcceptedAt, not even an explicit nil
### GetPreRegistrationAccessToken

`func (o *InvitationRead) GetPreRegistrationAccessToken() string`

GetPreRegistrationAccessToken returns the PreRegistrationAccessToken field if non-nil, zero value otherwise.

### GetPreRegistrationAccessTokenOk

`func (o *InvitationRead) GetPreRegistrationAccessTokenOk() (*string, bool)`

GetPreRegistrationAccessTokenOk returns a tuple with the PreRegistrationAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRegistrationAccessToken

`func (o *InvitationRead) SetPreRegistrationAccessToken(v string)`

SetPreRegistrationAccessToken sets PreRegistrationAccessToken field to given value.

### HasPreRegistrationAccessToken

`func (o *InvitationRead) HasPreRegistrationAccessToken() bool`

HasPreRegistrationAccessToken returns a boolean if a field has been set.

### SetPreRegistrationAccessTokenNil

`func (o *InvitationRead) SetPreRegistrationAccessTokenNil(b bool)`

 SetPreRegistrationAccessTokenNil sets the value for PreRegistrationAccessToken to be an explicit nil

### UnsetPreRegistrationAccessToken
`func (o *InvitationRead) UnsetPreRegistrationAccessToken()`

UnsetPreRegistrationAccessToken ensures that no value is present for PreRegistrationAccessToken, not even an explicit nil
### GetPreRegistrationRefreshToken

`func (o *InvitationRead) GetPreRegistrationRefreshToken() string`

GetPreRegistrationRefreshToken returns the PreRegistrationRefreshToken field if non-nil, zero value otherwise.

### GetPreRegistrationRefreshTokenOk

`func (o *InvitationRead) GetPreRegistrationRefreshTokenOk() (*string, bool)`

GetPreRegistrationRefreshTokenOk returns a tuple with the PreRegistrationRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRegistrationRefreshToken

`func (o *InvitationRead) SetPreRegistrationRefreshToken(v string)`

SetPreRegistrationRefreshToken sets PreRegistrationRefreshToken field to given value.

### HasPreRegistrationRefreshToken

`func (o *InvitationRead) HasPreRegistrationRefreshToken() bool`

HasPreRegistrationRefreshToken returns a boolean if a field has been set.

### SetPreRegistrationRefreshTokenNil

`func (o *InvitationRead) SetPreRegistrationRefreshTokenNil(b bool)`

 SetPreRegistrationRefreshTokenNil sets the value for PreRegistrationRefreshToken to be an explicit nil

### UnsetPreRegistrationRefreshToken
`func (o *InvitationRead) UnsetPreRegistrationRefreshToken()`

UnsetPreRegistrationRefreshToken ensures that no value is present for PreRegistrationRefreshToken, not even an explicit nil
### GetCallbackUrl

`func (o *InvitationRead) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *InvitationRead) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *InvitationRead) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *InvitationRead) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *InvitationRead) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *InvitationRead) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetSuccessRedirectUrl

`func (o *InvitationRead) GetSuccessRedirectUrl() string`

GetSuccessRedirectUrl returns the SuccessRedirectUrl field if non-nil, zero value otherwise.

### GetSuccessRedirectUrlOk

`func (o *InvitationRead) GetSuccessRedirectUrlOk() (*string, bool)`

GetSuccessRedirectUrlOk returns a tuple with the SuccessRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRedirectUrl

`func (o *InvitationRead) SetSuccessRedirectUrl(v string)`

SetSuccessRedirectUrl sets SuccessRedirectUrl field to given value.

### HasSuccessRedirectUrl

`func (o *InvitationRead) HasSuccessRedirectUrl() bool`

HasSuccessRedirectUrl returns a boolean if a field has been set.

### SetSuccessRedirectUrlNil

`func (o *InvitationRead) SetSuccessRedirectUrlNil(b bool)`

 SetSuccessRedirectUrlNil sets the value for SuccessRedirectUrl to be an explicit nil

### UnsetSuccessRedirectUrl
`func (o *InvitationRead) UnsetSuccessRedirectUrl()`

UnsetSuccessRedirectUrl ensures that no value is present for SuccessRedirectUrl, not even an explicit nil
### GetFailureRedirectUrl

`func (o *InvitationRead) GetFailureRedirectUrl() string`

GetFailureRedirectUrl returns the FailureRedirectUrl field if non-nil, zero value otherwise.

### GetFailureRedirectUrlOk

`func (o *InvitationRead) GetFailureRedirectUrlOk() (*string, bool)`

GetFailureRedirectUrlOk returns a tuple with the FailureRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureRedirectUrl

`func (o *InvitationRead) SetFailureRedirectUrl(v string)`

SetFailureRedirectUrl sets FailureRedirectUrl field to given value.

### HasFailureRedirectUrl

`func (o *InvitationRead) HasFailureRedirectUrl() bool`

HasFailureRedirectUrl returns a boolean if a field has been set.

### SetFailureRedirectUrlNil

`func (o *InvitationRead) SetFailureRedirectUrlNil(b bool)`

 SetFailureRedirectUrlNil sets the value for FailureRedirectUrl to be an explicit nil

### UnsetFailureRedirectUrl
`func (o *InvitationRead) UnsetFailureRedirectUrl()`

UnsetFailureRedirectUrl ensures that no value is present for FailureRedirectUrl, not even an explicit nil
### GetLimitTsps

`func (o *InvitationRead) GetLimitTsps() []string`

GetLimitTsps returns the LimitTsps field if non-nil, zero value otherwise.

### GetLimitTspsOk

`func (o *InvitationRead) GetLimitTspsOk() (*[]string, bool)`

GetLimitTspsOk returns a tuple with the LimitTsps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitTsps

`func (o *InvitationRead) SetLimitTsps(v []string)`

SetLimitTsps sets LimitTsps field to given value.

### HasLimitTsps

`func (o *InvitationRead) HasLimitTsps() bool`

HasLimitTsps returns a boolean if a field has been set.

### SetLimitTspsNil

`func (o *InvitationRead) SetLimitTspsNil(b bool)`

 SetLimitTspsNil sets the value for LimitTsps to be an explicit nil

### UnsetLimitTsps
`func (o *InvitationRead) UnsetLimitTsps()`

UnsetLimitTsps ensures that no value is present for LimitTsps, not even an explicit nil
### GetFleetId

`func (o *InvitationRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *InvitationRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *InvitationRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.

### HasFleetId

`func (o *InvitationRead) HasFleetId() bool`

HasFleetId returns a boolean if a field has been set.

### SetFleetIdNil

`func (o *InvitationRead) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *InvitationRead) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetPartnerSlug

`func (o *InvitationRead) GetPartnerSlug() string`

GetPartnerSlug returns the PartnerSlug field if non-nil, zero value otherwise.

### GetPartnerSlugOk

`func (o *InvitationRead) GetPartnerSlugOk() (*string, bool)`

GetPartnerSlugOk returns a tuple with the PartnerSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerSlug

`func (o *InvitationRead) SetPartnerSlug(v string)`

SetPartnerSlug sets PartnerSlug field to given value.

### HasPartnerSlug

`func (o *InvitationRead) HasPartnerSlug() bool`

HasPartnerSlug returns a boolean if a field has been set.

### SetPartnerSlugNil

`func (o *InvitationRead) SetPartnerSlugNil(b bool)`

 SetPartnerSlugNil sets the value for PartnerSlug to be an explicit nil

### UnsetPartnerSlug
`func (o *InvitationRead) UnsetPartnerSlug()`

UnsetPartnerSlug ensures that no value is present for PartnerSlug, not even an explicit nil
### GetPartnerId

`func (o *InvitationRead) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *InvitationRead) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *InvitationRead) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *InvitationRead) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### SetPartnerIdNil

`func (o *InvitationRead) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *InvitationRead) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetDeclineReason

`func (o *InvitationRead) GetDeclineReason() string`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *InvitationRead) GetDeclineReasonOk() (*string, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *InvitationRead) SetDeclineReason(v string)`

SetDeclineReason sets DeclineReason field to given value.

### HasDeclineReason

`func (o *InvitationRead) HasDeclineReason() bool`

HasDeclineReason returns a boolean if a field has been set.

### SetDeclineReasonNil

`func (o *InvitationRead) SetDeclineReasonNil(b bool)`

 SetDeclineReasonNil sets the value for DeclineReason to be an explicit nil

### UnsetDeclineReason
`func (o *InvitationRead) UnsetDeclineReason()`

UnsetDeclineReason ensures that no value is present for DeclineReason, not even an explicit nil
### GetDeclinedAt

`func (o *InvitationRead) GetDeclinedAt() time.Time`

GetDeclinedAt returns the DeclinedAt field if non-nil, zero value otherwise.

### GetDeclinedAtOk

`func (o *InvitationRead) GetDeclinedAtOk() (*time.Time, bool)`

GetDeclinedAtOk returns a tuple with the DeclinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclinedAt

`func (o *InvitationRead) SetDeclinedAt(v time.Time)`

SetDeclinedAt sets DeclinedAt field to given value.

### HasDeclinedAt

`func (o *InvitationRead) HasDeclinedAt() bool`

HasDeclinedAt returns a boolean if a field has been set.

### SetDeclinedAtNil

`func (o *InvitationRead) SetDeclinedAtNil(b bool)`

 SetDeclinedAtNil sets the value for DeclinedAt to be an explicit nil

### UnsetDeclinedAt
`func (o *InvitationRead) UnsetDeclinedAt()`

UnsetDeclinedAt ensures that no value is present for DeclinedAt, not even an explicit nil
### GetFleetRef

`func (o *InvitationRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *InvitationRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *InvitationRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *InvitationRead) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *InvitationRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *InvitationRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


