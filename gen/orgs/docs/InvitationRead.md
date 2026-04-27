# InvitationRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique invitation identifier | 
**CreatedAt** | **time.Time** | When the invitation was created | 
**MagicLink** | **string** | The magic link that is used to open Catena Connect and accept the invitation. Share this URL with the fleet to begin onboarding. | 
**ExpiresAt** | **time.Time** | The expiration date and time of the invitation | 
**ExpiresInHours** | **int32** | The number of hours the invitation is valid for. | 
**Status** | [**InvitationStatusEnum**](InvitationStatusEnum.md) | The current status of the invitation (active, accepted, declined, or expired) | 
**FleetName** | **NullableString** |  | 
**AcceptedAt** | **NullableTime** |  | 
**PreRegistrationAccessToken** | **NullableString** |  | 
**PreRegistrationRefreshToken** | **NullableString** |  | 
**CallbackUrl** | **NullableString** |  | 
**SuccessRedirectUrl** | **NullableString** |  | 
**FailureRedirectUrl** | **NullableString** |  | 
**LimitTsps** | **[]string** |  | 
**FleetId** | Pointer to **NullableString** |  | [optional] 
**PartnerSlug** | **NullableString** |  | 
**PartnerId** | **string** | Your organization ID requesting access to fleet data | 
**DeclineReason** | **NullableString** |  | 
**DeclinedAt** | **NullableTime** |  | 
**FleetRef** | **string** | Your internal fleet identifier. Use this to map Catena fleets back to your system. This value will be returned in webhooks and redirect URLs. | 
**Permissions** | [**map[string]ShareLevelEnum**](ShareLevelEnum.md) | Defines which resources (vehicle, locations, users, etc.) the fleet must grant access to and the permission level (read, write) for each. | 
**FleetEmail** | **NullableString** |  | 
**FleetRegulatoryId** | **NullableString** |  | 
**FleetRegulatoryIdType** | **NullableString** |  | 
**FleetPhone** | **NullableString** |  | 
**FleetWebsite** | **NullableString** |  | 
**FleetCountryCode** | **NullableString** |  | 
**PartnerProvidedFleetName** | **NullableString** |  | 
**PartnerProvidedFleetEmail** | **NullableString** |  | 
**PartnerProvidedFleetRegulatoryId** | **NullableString** |  | 
**PartnerProvidedFleetRegulatoryIdType** | **NullableString** |  | 
**PartnerProvidedFleetPhone** | **NullableString** |  | 
**PartnerProvidedFleetWebsite** | **NullableString** |  | 
**PartnerProvidedFleetCountryCode** | **NullableString** |  | 

## Methods

### NewInvitationRead

`func NewInvitationRead(id string, createdAt time.Time, magicLink string, expiresAt time.Time, expiresInHours int32, status InvitationStatusEnum, fleetName NullableString, acceptedAt NullableTime, preRegistrationAccessToken NullableString, preRegistrationRefreshToken NullableString, callbackUrl NullableString, successRedirectUrl NullableString, failureRedirectUrl NullableString, limitTsps []string, partnerSlug NullableString, partnerId string, declineReason NullableString, declinedAt NullableTime, fleetRef string, permissions map[string]ShareLevelEnum, fleetEmail NullableString, fleetRegulatoryId NullableString, fleetRegulatoryIdType NullableString, fleetPhone NullableString, fleetWebsite NullableString, fleetCountryCode NullableString, partnerProvidedFleetName NullableString, partnerProvidedFleetEmail NullableString, partnerProvidedFleetRegulatoryId NullableString, partnerProvidedFleetRegulatoryIdType NullableString, partnerProvidedFleetPhone NullableString, partnerProvidedFleetWebsite NullableString, partnerProvidedFleetCountryCode NullableString, ) *InvitationRead`

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

`func (o *InvitationRead) GetStatus() InvitationStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvitationRead) GetStatusOk() (*InvitationStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvitationRead) SetStatus(v InvitationStatusEnum)`

SetStatus sets Status field to given value.


### GetFleetName

`func (o *InvitationRead) GetFleetName() string`

GetFleetName returns the FleetName field if non-nil, zero value otherwise.

### GetFleetNameOk

`func (o *InvitationRead) GetFleetNameOk() (*string, bool)`

GetFleetNameOk returns a tuple with the FleetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetName

`func (o *InvitationRead) SetFleetName(v string)`

SetFleetName sets FleetName field to given value.


### SetFleetNameNil

`func (o *InvitationRead) SetFleetNameNil(b bool)`

 SetFleetNameNil sets the value for FleetName to be an explicit nil

### UnsetFleetName
`func (o *InvitationRead) UnsetFleetName()`

UnsetFleetName ensures that no value is present for FleetName, not even an explicit nil
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


### GetPermissions

`func (o *InvitationRead) GetPermissions() map[string]ShareLevelEnum`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InvitationRead) GetPermissionsOk() (*map[string]ShareLevelEnum, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InvitationRead) SetPermissions(v map[string]ShareLevelEnum)`

SetPermissions sets Permissions field to given value.


### GetFleetEmail

`func (o *InvitationRead) GetFleetEmail() string`

GetFleetEmail returns the FleetEmail field if non-nil, zero value otherwise.

### GetFleetEmailOk

`func (o *InvitationRead) GetFleetEmailOk() (*string, bool)`

GetFleetEmailOk returns a tuple with the FleetEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetEmail

`func (o *InvitationRead) SetFleetEmail(v string)`

SetFleetEmail sets FleetEmail field to given value.


### SetFleetEmailNil

`func (o *InvitationRead) SetFleetEmailNil(b bool)`

 SetFleetEmailNil sets the value for FleetEmail to be an explicit nil

### UnsetFleetEmail
`func (o *InvitationRead) UnsetFleetEmail()`

UnsetFleetEmail ensures that no value is present for FleetEmail, not even an explicit nil
### GetFleetRegulatoryId

`func (o *InvitationRead) GetFleetRegulatoryId() string`

GetFleetRegulatoryId returns the FleetRegulatoryId field if non-nil, zero value otherwise.

### GetFleetRegulatoryIdOk

`func (o *InvitationRead) GetFleetRegulatoryIdOk() (*string, bool)`

GetFleetRegulatoryIdOk returns a tuple with the FleetRegulatoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRegulatoryId

`func (o *InvitationRead) SetFleetRegulatoryId(v string)`

SetFleetRegulatoryId sets FleetRegulatoryId field to given value.


### SetFleetRegulatoryIdNil

`func (o *InvitationRead) SetFleetRegulatoryIdNil(b bool)`

 SetFleetRegulatoryIdNil sets the value for FleetRegulatoryId to be an explicit nil

### UnsetFleetRegulatoryId
`func (o *InvitationRead) UnsetFleetRegulatoryId()`

UnsetFleetRegulatoryId ensures that no value is present for FleetRegulatoryId, not even an explicit nil
### GetFleetRegulatoryIdType

`func (o *InvitationRead) GetFleetRegulatoryIdType() string`

GetFleetRegulatoryIdType returns the FleetRegulatoryIdType field if non-nil, zero value otherwise.

### GetFleetRegulatoryIdTypeOk

`func (o *InvitationRead) GetFleetRegulatoryIdTypeOk() (*string, bool)`

GetFleetRegulatoryIdTypeOk returns a tuple with the FleetRegulatoryIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRegulatoryIdType

`func (o *InvitationRead) SetFleetRegulatoryIdType(v string)`

SetFleetRegulatoryIdType sets FleetRegulatoryIdType field to given value.


### SetFleetRegulatoryIdTypeNil

`func (o *InvitationRead) SetFleetRegulatoryIdTypeNil(b bool)`

 SetFleetRegulatoryIdTypeNil sets the value for FleetRegulatoryIdType to be an explicit nil

### UnsetFleetRegulatoryIdType
`func (o *InvitationRead) UnsetFleetRegulatoryIdType()`

UnsetFleetRegulatoryIdType ensures that no value is present for FleetRegulatoryIdType, not even an explicit nil
### GetFleetPhone

`func (o *InvitationRead) GetFleetPhone() string`

GetFleetPhone returns the FleetPhone field if non-nil, zero value otherwise.

### GetFleetPhoneOk

`func (o *InvitationRead) GetFleetPhoneOk() (*string, bool)`

GetFleetPhoneOk returns a tuple with the FleetPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetPhone

`func (o *InvitationRead) SetFleetPhone(v string)`

SetFleetPhone sets FleetPhone field to given value.


### SetFleetPhoneNil

`func (o *InvitationRead) SetFleetPhoneNil(b bool)`

 SetFleetPhoneNil sets the value for FleetPhone to be an explicit nil

### UnsetFleetPhone
`func (o *InvitationRead) UnsetFleetPhone()`

UnsetFleetPhone ensures that no value is present for FleetPhone, not even an explicit nil
### GetFleetWebsite

`func (o *InvitationRead) GetFleetWebsite() string`

GetFleetWebsite returns the FleetWebsite field if non-nil, zero value otherwise.

### GetFleetWebsiteOk

`func (o *InvitationRead) GetFleetWebsiteOk() (*string, bool)`

GetFleetWebsiteOk returns a tuple with the FleetWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetWebsite

`func (o *InvitationRead) SetFleetWebsite(v string)`

SetFleetWebsite sets FleetWebsite field to given value.


### SetFleetWebsiteNil

`func (o *InvitationRead) SetFleetWebsiteNil(b bool)`

 SetFleetWebsiteNil sets the value for FleetWebsite to be an explicit nil

### UnsetFleetWebsite
`func (o *InvitationRead) UnsetFleetWebsite()`

UnsetFleetWebsite ensures that no value is present for FleetWebsite, not even an explicit nil
### GetFleetCountryCode

`func (o *InvitationRead) GetFleetCountryCode() string`

GetFleetCountryCode returns the FleetCountryCode field if non-nil, zero value otherwise.

### GetFleetCountryCodeOk

`func (o *InvitationRead) GetFleetCountryCodeOk() (*string, bool)`

GetFleetCountryCodeOk returns a tuple with the FleetCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetCountryCode

`func (o *InvitationRead) SetFleetCountryCode(v string)`

SetFleetCountryCode sets FleetCountryCode field to given value.


### SetFleetCountryCodeNil

`func (o *InvitationRead) SetFleetCountryCodeNil(b bool)`

 SetFleetCountryCodeNil sets the value for FleetCountryCode to be an explicit nil

### UnsetFleetCountryCode
`func (o *InvitationRead) UnsetFleetCountryCode()`

UnsetFleetCountryCode ensures that no value is present for FleetCountryCode, not even an explicit nil
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


### SetPartnerProvidedFleetNameNil

`func (o *InvitationRead) SetPartnerProvidedFleetNameNil(b bool)`

 SetPartnerProvidedFleetNameNil sets the value for PartnerProvidedFleetName to be an explicit nil

### UnsetPartnerProvidedFleetName
`func (o *InvitationRead) UnsetPartnerProvidedFleetName()`

UnsetPartnerProvidedFleetName ensures that no value is present for PartnerProvidedFleetName, not even an explicit nil
### GetPartnerProvidedFleetEmail

`func (o *InvitationRead) GetPartnerProvidedFleetEmail() string`

GetPartnerProvidedFleetEmail returns the PartnerProvidedFleetEmail field if non-nil, zero value otherwise.

### GetPartnerProvidedFleetEmailOk

`func (o *InvitationRead) GetPartnerProvidedFleetEmailOk() (*string, bool)`

GetPartnerProvidedFleetEmailOk returns a tuple with the PartnerProvidedFleetEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProvidedFleetEmail

`func (o *InvitationRead) SetPartnerProvidedFleetEmail(v string)`

SetPartnerProvidedFleetEmail sets PartnerProvidedFleetEmail field to given value.


### SetPartnerProvidedFleetEmailNil

`func (o *InvitationRead) SetPartnerProvidedFleetEmailNil(b bool)`

 SetPartnerProvidedFleetEmailNil sets the value for PartnerProvidedFleetEmail to be an explicit nil

### UnsetPartnerProvidedFleetEmail
`func (o *InvitationRead) UnsetPartnerProvidedFleetEmail()`

UnsetPartnerProvidedFleetEmail ensures that no value is present for PartnerProvidedFleetEmail, not even an explicit nil
### GetPartnerProvidedFleetRegulatoryId

`func (o *InvitationRead) GetPartnerProvidedFleetRegulatoryId() string`

GetPartnerProvidedFleetRegulatoryId returns the PartnerProvidedFleetRegulatoryId field if non-nil, zero value otherwise.

### GetPartnerProvidedFleetRegulatoryIdOk

`func (o *InvitationRead) GetPartnerProvidedFleetRegulatoryIdOk() (*string, bool)`

GetPartnerProvidedFleetRegulatoryIdOk returns a tuple with the PartnerProvidedFleetRegulatoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProvidedFleetRegulatoryId

`func (o *InvitationRead) SetPartnerProvidedFleetRegulatoryId(v string)`

SetPartnerProvidedFleetRegulatoryId sets PartnerProvidedFleetRegulatoryId field to given value.


### SetPartnerProvidedFleetRegulatoryIdNil

`func (o *InvitationRead) SetPartnerProvidedFleetRegulatoryIdNil(b bool)`

 SetPartnerProvidedFleetRegulatoryIdNil sets the value for PartnerProvidedFleetRegulatoryId to be an explicit nil

### UnsetPartnerProvidedFleetRegulatoryId
`func (o *InvitationRead) UnsetPartnerProvidedFleetRegulatoryId()`

UnsetPartnerProvidedFleetRegulatoryId ensures that no value is present for PartnerProvidedFleetRegulatoryId, not even an explicit nil
### GetPartnerProvidedFleetRegulatoryIdType

`func (o *InvitationRead) GetPartnerProvidedFleetRegulatoryIdType() string`

GetPartnerProvidedFleetRegulatoryIdType returns the PartnerProvidedFleetRegulatoryIdType field if non-nil, zero value otherwise.

### GetPartnerProvidedFleetRegulatoryIdTypeOk

`func (o *InvitationRead) GetPartnerProvidedFleetRegulatoryIdTypeOk() (*string, bool)`

GetPartnerProvidedFleetRegulatoryIdTypeOk returns a tuple with the PartnerProvidedFleetRegulatoryIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProvidedFleetRegulatoryIdType

`func (o *InvitationRead) SetPartnerProvidedFleetRegulatoryIdType(v string)`

SetPartnerProvidedFleetRegulatoryIdType sets PartnerProvidedFleetRegulatoryIdType field to given value.


### SetPartnerProvidedFleetRegulatoryIdTypeNil

`func (o *InvitationRead) SetPartnerProvidedFleetRegulatoryIdTypeNil(b bool)`

 SetPartnerProvidedFleetRegulatoryIdTypeNil sets the value for PartnerProvidedFleetRegulatoryIdType to be an explicit nil

### UnsetPartnerProvidedFleetRegulatoryIdType
`func (o *InvitationRead) UnsetPartnerProvidedFleetRegulatoryIdType()`

UnsetPartnerProvidedFleetRegulatoryIdType ensures that no value is present for PartnerProvidedFleetRegulatoryIdType, not even an explicit nil
### GetPartnerProvidedFleetPhone

`func (o *InvitationRead) GetPartnerProvidedFleetPhone() string`

GetPartnerProvidedFleetPhone returns the PartnerProvidedFleetPhone field if non-nil, zero value otherwise.

### GetPartnerProvidedFleetPhoneOk

`func (o *InvitationRead) GetPartnerProvidedFleetPhoneOk() (*string, bool)`

GetPartnerProvidedFleetPhoneOk returns a tuple with the PartnerProvidedFleetPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProvidedFleetPhone

`func (o *InvitationRead) SetPartnerProvidedFleetPhone(v string)`

SetPartnerProvidedFleetPhone sets PartnerProvidedFleetPhone field to given value.


### SetPartnerProvidedFleetPhoneNil

`func (o *InvitationRead) SetPartnerProvidedFleetPhoneNil(b bool)`

 SetPartnerProvidedFleetPhoneNil sets the value for PartnerProvidedFleetPhone to be an explicit nil

### UnsetPartnerProvidedFleetPhone
`func (o *InvitationRead) UnsetPartnerProvidedFleetPhone()`

UnsetPartnerProvidedFleetPhone ensures that no value is present for PartnerProvidedFleetPhone, not even an explicit nil
### GetPartnerProvidedFleetWebsite

`func (o *InvitationRead) GetPartnerProvidedFleetWebsite() string`

GetPartnerProvidedFleetWebsite returns the PartnerProvidedFleetWebsite field if non-nil, zero value otherwise.

### GetPartnerProvidedFleetWebsiteOk

`func (o *InvitationRead) GetPartnerProvidedFleetWebsiteOk() (*string, bool)`

GetPartnerProvidedFleetWebsiteOk returns a tuple with the PartnerProvidedFleetWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProvidedFleetWebsite

`func (o *InvitationRead) SetPartnerProvidedFleetWebsite(v string)`

SetPartnerProvidedFleetWebsite sets PartnerProvidedFleetWebsite field to given value.


### SetPartnerProvidedFleetWebsiteNil

`func (o *InvitationRead) SetPartnerProvidedFleetWebsiteNil(b bool)`

 SetPartnerProvidedFleetWebsiteNil sets the value for PartnerProvidedFleetWebsite to be an explicit nil

### UnsetPartnerProvidedFleetWebsite
`func (o *InvitationRead) UnsetPartnerProvidedFleetWebsite()`

UnsetPartnerProvidedFleetWebsite ensures that no value is present for PartnerProvidedFleetWebsite, not even an explicit nil
### GetPartnerProvidedFleetCountryCode

`func (o *InvitationRead) GetPartnerProvidedFleetCountryCode() string`

GetPartnerProvidedFleetCountryCode returns the PartnerProvidedFleetCountryCode field if non-nil, zero value otherwise.

### GetPartnerProvidedFleetCountryCodeOk

`func (o *InvitationRead) GetPartnerProvidedFleetCountryCodeOk() (*string, bool)`

GetPartnerProvidedFleetCountryCodeOk returns a tuple with the PartnerProvidedFleetCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProvidedFleetCountryCode

`func (o *InvitationRead) SetPartnerProvidedFleetCountryCode(v string)`

SetPartnerProvidedFleetCountryCode sets PartnerProvidedFleetCountryCode field to given value.


### SetPartnerProvidedFleetCountryCodeNil

`func (o *InvitationRead) SetPartnerProvidedFleetCountryCodeNil(b bool)`

 SetPartnerProvidedFleetCountryCodeNil sets the value for PartnerProvidedFleetCountryCode to be an explicit nil

### UnsetPartnerProvidedFleetCountryCode
`func (o *InvitationRead) UnsetPartnerProvidedFleetCountryCode()`

UnsetPartnerProvidedFleetCountryCode ensures that no value is present for PartnerProvidedFleetCountryCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


