# BaseInvitationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the invitation | 
**CreatedAt** | **time.Time** | The date and time the invitation was created | 
**MagicLink** | **string** | The magic link to accept the invitation. | 
**ExpiresAt** | **time.Time** | The date and time the magic link expires. | 
**ExpiresInHours** | **int32** | The number of hours until the magic link expires. | 
**Status** | [**StatusEnum**](StatusEnum.md) | The status of the invitation. | 
**PartnerProvidedFleetName** | Pointer to **NullableString** |  | [optional] 
**AcceptedAt** | Pointer to **NullableTime** |  | [optional] 
**PreRegistrationAccessToken** | Pointer to **NullableString** |  | [optional] 
**PreRegistrationRefreshToken** | Pointer to **NullableString** |  | [optional] 
**SuccessRedirectUrl** | Pointer to **NullableString** |  | [optional] 
**FailureRedirectUrl** | Pointer to **NullableString** |  | [optional] 
**CallbackUrl** | Pointer to **NullableString** |  | [optional] 
**LimitTsps** | Pointer to [**[]TspEnum**](TspEnum.md) |  | [optional] 
**FleetId** | Pointer to **NullableString** |  | [optional] 
**PartnerSlug** | Pointer to **NullableString** |  | [optional] 
**PartnerId** | Pointer to **NullableString** |  | [optional] 
**DeclineReason** | Pointer to **NullableString** |  | [optional] 
**DeclinedAt** | Pointer to **NullableTime** |  | [optional] 
**FleetRef** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBaseInvitationEvent

`func NewBaseInvitationEvent(id string, createdAt time.Time, magicLink string, expiresAt time.Time, expiresInHours int32, status StatusEnum, ) *BaseInvitationEvent`

NewBaseInvitationEvent instantiates a new BaseInvitationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseInvitationEventWithDefaults

`func NewBaseInvitationEventWithDefaults() *BaseInvitationEvent`

NewBaseInvitationEventWithDefaults instantiates a new BaseInvitationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseInvitationEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseInvitationEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseInvitationEvent) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *BaseInvitationEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseInvitationEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseInvitationEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetMagicLink

`func (o *BaseInvitationEvent) GetMagicLink() string`

GetMagicLink returns the MagicLink field if non-nil, zero value otherwise.

### GetMagicLinkOk

`func (o *BaseInvitationEvent) GetMagicLinkOk() (*string, bool)`

GetMagicLinkOk returns a tuple with the MagicLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMagicLink

`func (o *BaseInvitationEvent) SetMagicLink(v string)`

SetMagicLink sets MagicLink field to given value.


### GetExpiresAt

`func (o *BaseInvitationEvent) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *BaseInvitationEvent) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *BaseInvitationEvent) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetExpiresInHours

`func (o *BaseInvitationEvent) GetExpiresInHours() int32`

GetExpiresInHours returns the ExpiresInHours field if non-nil, zero value otherwise.

### GetExpiresInHoursOk

`func (o *BaseInvitationEvent) GetExpiresInHoursOk() (*int32, bool)`

GetExpiresInHoursOk returns a tuple with the ExpiresInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInHours

`func (o *BaseInvitationEvent) SetExpiresInHours(v int32)`

SetExpiresInHours sets ExpiresInHours field to given value.


### GetStatus

`func (o *BaseInvitationEvent) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseInvitationEvent) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseInvitationEvent) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetPartnerProvidedFleetName

`func (o *BaseInvitationEvent) GetPartnerProvidedFleetName() string`

GetPartnerProvidedFleetName returns the PartnerProvidedFleetName field if non-nil, zero value otherwise.

### GetPartnerProvidedFleetNameOk

`func (o *BaseInvitationEvent) GetPartnerProvidedFleetNameOk() (*string, bool)`

GetPartnerProvidedFleetNameOk returns a tuple with the PartnerProvidedFleetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProvidedFleetName

`func (o *BaseInvitationEvent) SetPartnerProvidedFleetName(v string)`

SetPartnerProvidedFleetName sets PartnerProvidedFleetName field to given value.

### HasPartnerProvidedFleetName

`func (o *BaseInvitationEvent) HasPartnerProvidedFleetName() bool`

HasPartnerProvidedFleetName returns a boolean if a field has been set.

### SetPartnerProvidedFleetNameNil

`func (o *BaseInvitationEvent) SetPartnerProvidedFleetNameNil(b bool)`

 SetPartnerProvidedFleetNameNil sets the value for PartnerProvidedFleetName to be an explicit nil

### UnsetPartnerProvidedFleetName
`func (o *BaseInvitationEvent) UnsetPartnerProvidedFleetName()`

UnsetPartnerProvidedFleetName ensures that no value is present for PartnerProvidedFleetName, not even an explicit nil
### GetAcceptedAt

`func (o *BaseInvitationEvent) GetAcceptedAt() time.Time`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *BaseInvitationEvent) GetAcceptedAtOk() (*time.Time, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *BaseInvitationEvent) SetAcceptedAt(v time.Time)`

SetAcceptedAt sets AcceptedAt field to given value.

### HasAcceptedAt

`func (o *BaseInvitationEvent) HasAcceptedAt() bool`

HasAcceptedAt returns a boolean if a field has been set.

### SetAcceptedAtNil

`func (o *BaseInvitationEvent) SetAcceptedAtNil(b bool)`

 SetAcceptedAtNil sets the value for AcceptedAt to be an explicit nil

### UnsetAcceptedAt
`func (o *BaseInvitationEvent) UnsetAcceptedAt()`

UnsetAcceptedAt ensures that no value is present for AcceptedAt, not even an explicit nil
### GetPreRegistrationAccessToken

`func (o *BaseInvitationEvent) GetPreRegistrationAccessToken() string`

GetPreRegistrationAccessToken returns the PreRegistrationAccessToken field if non-nil, zero value otherwise.

### GetPreRegistrationAccessTokenOk

`func (o *BaseInvitationEvent) GetPreRegistrationAccessTokenOk() (*string, bool)`

GetPreRegistrationAccessTokenOk returns a tuple with the PreRegistrationAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRegistrationAccessToken

`func (o *BaseInvitationEvent) SetPreRegistrationAccessToken(v string)`

SetPreRegistrationAccessToken sets PreRegistrationAccessToken field to given value.

### HasPreRegistrationAccessToken

`func (o *BaseInvitationEvent) HasPreRegistrationAccessToken() bool`

HasPreRegistrationAccessToken returns a boolean if a field has been set.

### SetPreRegistrationAccessTokenNil

`func (o *BaseInvitationEvent) SetPreRegistrationAccessTokenNil(b bool)`

 SetPreRegistrationAccessTokenNil sets the value for PreRegistrationAccessToken to be an explicit nil

### UnsetPreRegistrationAccessToken
`func (o *BaseInvitationEvent) UnsetPreRegistrationAccessToken()`

UnsetPreRegistrationAccessToken ensures that no value is present for PreRegistrationAccessToken, not even an explicit nil
### GetPreRegistrationRefreshToken

`func (o *BaseInvitationEvent) GetPreRegistrationRefreshToken() string`

GetPreRegistrationRefreshToken returns the PreRegistrationRefreshToken field if non-nil, zero value otherwise.

### GetPreRegistrationRefreshTokenOk

`func (o *BaseInvitationEvent) GetPreRegistrationRefreshTokenOk() (*string, bool)`

GetPreRegistrationRefreshTokenOk returns a tuple with the PreRegistrationRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreRegistrationRefreshToken

`func (o *BaseInvitationEvent) SetPreRegistrationRefreshToken(v string)`

SetPreRegistrationRefreshToken sets PreRegistrationRefreshToken field to given value.

### HasPreRegistrationRefreshToken

`func (o *BaseInvitationEvent) HasPreRegistrationRefreshToken() bool`

HasPreRegistrationRefreshToken returns a boolean if a field has been set.

### SetPreRegistrationRefreshTokenNil

`func (o *BaseInvitationEvent) SetPreRegistrationRefreshTokenNil(b bool)`

 SetPreRegistrationRefreshTokenNil sets the value for PreRegistrationRefreshToken to be an explicit nil

### UnsetPreRegistrationRefreshToken
`func (o *BaseInvitationEvent) UnsetPreRegistrationRefreshToken()`

UnsetPreRegistrationRefreshToken ensures that no value is present for PreRegistrationRefreshToken, not even an explicit nil
### GetSuccessRedirectUrl

`func (o *BaseInvitationEvent) GetSuccessRedirectUrl() string`

GetSuccessRedirectUrl returns the SuccessRedirectUrl field if non-nil, zero value otherwise.

### GetSuccessRedirectUrlOk

`func (o *BaseInvitationEvent) GetSuccessRedirectUrlOk() (*string, bool)`

GetSuccessRedirectUrlOk returns a tuple with the SuccessRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRedirectUrl

`func (o *BaseInvitationEvent) SetSuccessRedirectUrl(v string)`

SetSuccessRedirectUrl sets SuccessRedirectUrl field to given value.

### HasSuccessRedirectUrl

`func (o *BaseInvitationEvent) HasSuccessRedirectUrl() bool`

HasSuccessRedirectUrl returns a boolean if a field has been set.

### SetSuccessRedirectUrlNil

`func (o *BaseInvitationEvent) SetSuccessRedirectUrlNil(b bool)`

 SetSuccessRedirectUrlNil sets the value for SuccessRedirectUrl to be an explicit nil

### UnsetSuccessRedirectUrl
`func (o *BaseInvitationEvent) UnsetSuccessRedirectUrl()`

UnsetSuccessRedirectUrl ensures that no value is present for SuccessRedirectUrl, not even an explicit nil
### GetFailureRedirectUrl

`func (o *BaseInvitationEvent) GetFailureRedirectUrl() string`

GetFailureRedirectUrl returns the FailureRedirectUrl field if non-nil, zero value otherwise.

### GetFailureRedirectUrlOk

`func (o *BaseInvitationEvent) GetFailureRedirectUrlOk() (*string, bool)`

GetFailureRedirectUrlOk returns a tuple with the FailureRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureRedirectUrl

`func (o *BaseInvitationEvent) SetFailureRedirectUrl(v string)`

SetFailureRedirectUrl sets FailureRedirectUrl field to given value.

### HasFailureRedirectUrl

`func (o *BaseInvitationEvent) HasFailureRedirectUrl() bool`

HasFailureRedirectUrl returns a boolean if a field has been set.

### SetFailureRedirectUrlNil

`func (o *BaseInvitationEvent) SetFailureRedirectUrlNil(b bool)`

 SetFailureRedirectUrlNil sets the value for FailureRedirectUrl to be an explicit nil

### UnsetFailureRedirectUrl
`func (o *BaseInvitationEvent) UnsetFailureRedirectUrl()`

UnsetFailureRedirectUrl ensures that no value is present for FailureRedirectUrl, not even an explicit nil
### GetCallbackUrl

`func (o *BaseInvitationEvent) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *BaseInvitationEvent) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *BaseInvitationEvent) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *BaseInvitationEvent) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *BaseInvitationEvent) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *BaseInvitationEvent) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetLimitTsps

`func (o *BaseInvitationEvent) GetLimitTsps() []TspEnum`

GetLimitTsps returns the LimitTsps field if non-nil, zero value otherwise.

### GetLimitTspsOk

`func (o *BaseInvitationEvent) GetLimitTspsOk() (*[]TspEnum, bool)`

GetLimitTspsOk returns a tuple with the LimitTsps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitTsps

`func (o *BaseInvitationEvent) SetLimitTsps(v []TspEnum)`

SetLimitTsps sets LimitTsps field to given value.

### HasLimitTsps

`func (o *BaseInvitationEvent) HasLimitTsps() bool`

HasLimitTsps returns a boolean if a field has been set.

### SetLimitTspsNil

`func (o *BaseInvitationEvent) SetLimitTspsNil(b bool)`

 SetLimitTspsNil sets the value for LimitTsps to be an explicit nil

### UnsetLimitTsps
`func (o *BaseInvitationEvent) UnsetLimitTsps()`

UnsetLimitTsps ensures that no value is present for LimitTsps, not even an explicit nil
### GetFleetId

`func (o *BaseInvitationEvent) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseInvitationEvent) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseInvitationEvent) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.

### HasFleetId

`func (o *BaseInvitationEvent) HasFleetId() bool`

HasFleetId returns a boolean if a field has been set.

### SetFleetIdNil

`func (o *BaseInvitationEvent) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *BaseInvitationEvent) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetPartnerSlug

`func (o *BaseInvitationEvent) GetPartnerSlug() string`

GetPartnerSlug returns the PartnerSlug field if non-nil, zero value otherwise.

### GetPartnerSlugOk

`func (o *BaseInvitationEvent) GetPartnerSlugOk() (*string, bool)`

GetPartnerSlugOk returns a tuple with the PartnerSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerSlug

`func (o *BaseInvitationEvent) SetPartnerSlug(v string)`

SetPartnerSlug sets PartnerSlug field to given value.

### HasPartnerSlug

`func (o *BaseInvitationEvent) HasPartnerSlug() bool`

HasPartnerSlug returns a boolean if a field has been set.

### SetPartnerSlugNil

`func (o *BaseInvitationEvent) SetPartnerSlugNil(b bool)`

 SetPartnerSlugNil sets the value for PartnerSlug to be an explicit nil

### UnsetPartnerSlug
`func (o *BaseInvitationEvent) UnsetPartnerSlug()`

UnsetPartnerSlug ensures that no value is present for PartnerSlug, not even an explicit nil
### GetPartnerId

`func (o *BaseInvitationEvent) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *BaseInvitationEvent) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *BaseInvitationEvent) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *BaseInvitationEvent) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### SetPartnerIdNil

`func (o *BaseInvitationEvent) SetPartnerIdNil(b bool)`

 SetPartnerIdNil sets the value for PartnerId to be an explicit nil

### UnsetPartnerId
`func (o *BaseInvitationEvent) UnsetPartnerId()`

UnsetPartnerId ensures that no value is present for PartnerId, not even an explicit nil
### GetDeclineReason

`func (o *BaseInvitationEvent) GetDeclineReason() string`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *BaseInvitationEvent) GetDeclineReasonOk() (*string, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *BaseInvitationEvent) SetDeclineReason(v string)`

SetDeclineReason sets DeclineReason field to given value.

### HasDeclineReason

`func (o *BaseInvitationEvent) HasDeclineReason() bool`

HasDeclineReason returns a boolean if a field has been set.

### SetDeclineReasonNil

`func (o *BaseInvitationEvent) SetDeclineReasonNil(b bool)`

 SetDeclineReasonNil sets the value for DeclineReason to be an explicit nil

### UnsetDeclineReason
`func (o *BaseInvitationEvent) UnsetDeclineReason()`

UnsetDeclineReason ensures that no value is present for DeclineReason, not even an explicit nil
### GetDeclinedAt

`func (o *BaseInvitationEvent) GetDeclinedAt() time.Time`

GetDeclinedAt returns the DeclinedAt field if non-nil, zero value otherwise.

### GetDeclinedAtOk

`func (o *BaseInvitationEvent) GetDeclinedAtOk() (*time.Time, bool)`

GetDeclinedAtOk returns a tuple with the DeclinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclinedAt

`func (o *BaseInvitationEvent) SetDeclinedAt(v time.Time)`

SetDeclinedAt sets DeclinedAt field to given value.

### HasDeclinedAt

`func (o *BaseInvitationEvent) HasDeclinedAt() bool`

HasDeclinedAt returns a boolean if a field has been set.

### SetDeclinedAtNil

`func (o *BaseInvitationEvent) SetDeclinedAtNil(b bool)`

 SetDeclinedAtNil sets the value for DeclinedAt to be an explicit nil

### UnsetDeclinedAt
`func (o *BaseInvitationEvent) UnsetDeclinedAt()`

UnsetDeclinedAt ensures that no value is present for DeclinedAt, not even an explicit nil
### GetFleetRef

`func (o *BaseInvitationEvent) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *BaseInvitationEvent) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *BaseInvitationEvent) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *BaseInvitationEvent) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *BaseInvitationEvent) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *BaseInvitationEvent) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


