# InvitationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FleetRef** | Pointer to **NullableString** |  | [optional] 
**PartnerProvidedFleetName** | Pointer to **NullableString** |  | [optional] 
**PartnerProvidedFleetEmail** | Pointer to **NullableString** |  | [optional] 
**PartnerProvidedFleetRegulatoryId** | Pointer to **NullableString** |  | [optional] 
**PartnerProvidedFleetRegulatoryIdType** | Pointer to **NullableString** |  | [optional] 
**PartnerProvidedFleetPhone** | Pointer to **NullableString** |  | [optional] 
**PartnerProvidedFleetWebsite** | Pointer to **NullableString** |  | [optional] 
**PartnerProvidedFleetCountryCode** | Pointer to **NullableString** |  | [optional] 
**SuccessRedirectUrl** | Pointer to **NullableString** |  | [optional] 
**FailureRedirectUrl** | Pointer to **NullableString** |  | [optional] 
**CallbackUrl** | Pointer to **NullableString** |  | [optional] 
**LimitTsps** | Pointer to **[]string** |  | [optional] 
**ExpiresInHours** | Pointer to **int32** | How long the invitation link remains valid (1-672 hours). Default is 24 hours. *Consider longer durations for email campaigns.* | [optional] [default to 24]

## Methods

### NewInvitationCreate

`func NewInvitationCreate() *InvitationCreate`

NewInvitationCreate instantiates a new InvitationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationCreateWithDefaults

`func NewInvitationCreateWithDefaults() *InvitationCreate`

NewInvitationCreateWithDefaults instantiates a new InvitationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetRef

`func (o *InvitationCreate) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *InvitationCreate) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *InvitationCreate) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *InvitationCreate) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *InvitationCreate) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *InvitationCreate) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetPartnerProvidedFleetName

`func (o *InvitationCreate) GetPartnerProvidedFleetName() string`

GetPartnerProvidedFleetName returns the PartnerProvidedFleetName field if non-nil, zero value otherwise.

### GetPartnerProvidedFleetNameOk

`func (o *InvitationCreate) GetPartnerProvidedFleetNameOk() (*string, bool)`

GetPartnerProvidedFleetNameOk returns a tuple with the PartnerProvidedFleetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProvidedFleetName

`func (o *InvitationCreate) SetPartnerProvidedFleetName(v string)`

SetPartnerProvidedFleetName sets PartnerProvidedFleetName field to given value.

### HasPartnerProvidedFleetName

`func (o *InvitationCreate) HasPartnerProvidedFleetName() bool`

HasPartnerProvidedFleetName returns a boolean if a field has been set.

### SetPartnerProvidedFleetNameNil

`func (o *InvitationCreate) SetPartnerProvidedFleetNameNil(b bool)`

 SetPartnerProvidedFleetNameNil sets the value for PartnerProvidedFleetName to be an explicit nil

### UnsetPartnerProvidedFleetName
`func (o *InvitationCreate) UnsetPartnerProvidedFleetName()`

UnsetPartnerProvidedFleetName ensures that no value is present for PartnerProvidedFleetName, not even an explicit nil
### GetPartnerProvidedFleetEmail

`func (o *InvitationCreate) GetPartnerProvidedFleetEmail() string`

GetPartnerProvidedFleetEmail returns the PartnerProvidedFleetEmail field if non-nil, zero value otherwise.

### GetPartnerProvidedFleetEmailOk

`func (o *InvitationCreate) GetPartnerProvidedFleetEmailOk() (*string, bool)`

GetPartnerProvidedFleetEmailOk returns a tuple with the PartnerProvidedFleetEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProvidedFleetEmail

`func (o *InvitationCreate) SetPartnerProvidedFleetEmail(v string)`

SetPartnerProvidedFleetEmail sets PartnerProvidedFleetEmail field to given value.

### HasPartnerProvidedFleetEmail

`func (o *InvitationCreate) HasPartnerProvidedFleetEmail() bool`

HasPartnerProvidedFleetEmail returns a boolean if a field has been set.

### SetPartnerProvidedFleetEmailNil

`func (o *InvitationCreate) SetPartnerProvidedFleetEmailNil(b bool)`

 SetPartnerProvidedFleetEmailNil sets the value for PartnerProvidedFleetEmail to be an explicit nil

### UnsetPartnerProvidedFleetEmail
`func (o *InvitationCreate) UnsetPartnerProvidedFleetEmail()`

UnsetPartnerProvidedFleetEmail ensures that no value is present for PartnerProvidedFleetEmail, not even an explicit nil
### GetPartnerProvidedFleetRegulatoryId

`func (o *InvitationCreate) GetPartnerProvidedFleetRegulatoryId() string`

GetPartnerProvidedFleetRegulatoryId returns the PartnerProvidedFleetRegulatoryId field if non-nil, zero value otherwise.

### GetPartnerProvidedFleetRegulatoryIdOk

`func (o *InvitationCreate) GetPartnerProvidedFleetRegulatoryIdOk() (*string, bool)`

GetPartnerProvidedFleetRegulatoryIdOk returns a tuple with the PartnerProvidedFleetRegulatoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProvidedFleetRegulatoryId

`func (o *InvitationCreate) SetPartnerProvidedFleetRegulatoryId(v string)`

SetPartnerProvidedFleetRegulatoryId sets PartnerProvidedFleetRegulatoryId field to given value.

### HasPartnerProvidedFleetRegulatoryId

`func (o *InvitationCreate) HasPartnerProvidedFleetRegulatoryId() bool`

HasPartnerProvidedFleetRegulatoryId returns a boolean if a field has been set.

### SetPartnerProvidedFleetRegulatoryIdNil

`func (o *InvitationCreate) SetPartnerProvidedFleetRegulatoryIdNil(b bool)`

 SetPartnerProvidedFleetRegulatoryIdNil sets the value for PartnerProvidedFleetRegulatoryId to be an explicit nil

### UnsetPartnerProvidedFleetRegulatoryId
`func (o *InvitationCreate) UnsetPartnerProvidedFleetRegulatoryId()`

UnsetPartnerProvidedFleetRegulatoryId ensures that no value is present for PartnerProvidedFleetRegulatoryId, not even an explicit nil
### GetPartnerProvidedFleetRegulatoryIdType

`func (o *InvitationCreate) GetPartnerProvidedFleetRegulatoryIdType() string`

GetPartnerProvidedFleetRegulatoryIdType returns the PartnerProvidedFleetRegulatoryIdType field if non-nil, zero value otherwise.

### GetPartnerProvidedFleetRegulatoryIdTypeOk

`func (o *InvitationCreate) GetPartnerProvidedFleetRegulatoryIdTypeOk() (*string, bool)`

GetPartnerProvidedFleetRegulatoryIdTypeOk returns a tuple with the PartnerProvidedFleetRegulatoryIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProvidedFleetRegulatoryIdType

`func (o *InvitationCreate) SetPartnerProvidedFleetRegulatoryIdType(v string)`

SetPartnerProvidedFleetRegulatoryIdType sets PartnerProvidedFleetRegulatoryIdType field to given value.

### HasPartnerProvidedFleetRegulatoryIdType

`func (o *InvitationCreate) HasPartnerProvidedFleetRegulatoryIdType() bool`

HasPartnerProvidedFleetRegulatoryIdType returns a boolean if a field has been set.

### SetPartnerProvidedFleetRegulatoryIdTypeNil

`func (o *InvitationCreate) SetPartnerProvidedFleetRegulatoryIdTypeNil(b bool)`

 SetPartnerProvidedFleetRegulatoryIdTypeNil sets the value for PartnerProvidedFleetRegulatoryIdType to be an explicit nil

### UnsetPartnerProvidedFleetRegulatoryIdType
`func (o *InvitationCreate) UnsetPartnerProvidedFleetRegulatoryIdType()`

UnsetPartnerProvidedFleetRegulatoryIdType ensures that no value is present for PartnerProvidedFleetRegulatoryIdType, not even an explicit nil
### GetPartnerProvidedFleetPhone

`func (o *InvitationCreate) GetPartnerProvidedFleetPhone() string`

GetPartnerProvidedFleetPhone returns the PartnerProvidedFleetPhone field if non-nil, zero value otherwise.

### GetPartnerProvidedFleetPhoneOk

`func (o *InvitationCreate) GetPartnerProvidedFleetPhoneOk() (*string, bool)`

GetPartnerProvidedFleetPhoneOk returns a tuple with the PartnerProvidedFleetPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProvidedFleetPhone

`func (o *InvitationCreate) SetPartnerProvidedFleetPhone(v string)`

SetPartnerProvidedFleetPhone sets PartnerProvidedFleetPhone field to given value.

### HasPartnerProvidedFleetPhone

`func (o *InvitationCreate) HasPartnerProvidedFleetPhone() bool`

HasPartnerProvidedFleetPhone returns a boolean if a field has been set.

### SetPartnerProvidedFleetPhoneNil

`func (o *InvitationCreate) SetPartnerProvidedFleetPhoneNil(b bool)`

 SetPartnerProvidedFleetPhoneNil sets the value for PartnerProvidedFleetPhone to be an explicit nil

### UnsetPartnerProvidedFleetPhone
`func (o *InvitationCreate) UnsetPartnerProvidedFleetPhone()`

UnsetPartnerProvidedFleetPhone ensures that no value is present for PartnerProvidedFleetPhone, not even an explicit nil
### GetPartnerProvidedFleetWebsite

`func (o *InvitationCreate) GetPartnerProvidedFleetWebsite() string`

GetPartnerProvidedFleetWebsite returns the PartnerProvidedFleetWebsite field if non-nil, zero value otherwise.

### GetPartnerProvidedFleetWebsiteOk

`func (o *InvitationCreate) GetPartnerProvidedFleetWebsiteOk() (*string, bool)`

GetPartnerProvidedFleetWebsiteOk returns a tuple with the PartnerProvidedFleetWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProvidedFleetWebsite

`func (o *InvitationCreate) SetPartnerProvidedFleetWebsite(v string)`

SetPartnerProvidedFleetWebsite sets PartnerProvidedFleetWebsite field to given value.

### HasPartnerProvidedFleetWebsite

`func (o *InvitationCreate) HasPartnerProvidedFleetWebsite() bool`

HasPartnerProvidedFleetWebsite returns a boolean if a field has been set.

### SetPartnerProvidedFleetWebsiteNil

`func (o *InvitationCreate) SetPartnerProvidedFleetWebsiteNil(b bool)`

 SetPartnerProvidedFleetWebsiteNil sets the value for PartnerProvidedFleetWebsite to be an explicit nil

### UnsetPartnerProvidedFleetWebsite
`func (o *InvitationCreate) UnsetPartnerProvidedFleetWebsite()`

UnsetPartnerProvidedFleetWebsite ensures that no value is present for PartnerProvidedFleetWebsite, not even an explicit nil
### GetPartnerProvidedFleetCountryCode

`func (o *InvitationCreate) GetPartnerProvidedFleetCountryCode() string`

GetPartnerProvidedFleetCountryCode returns the PartnerProvidedFleetCountryCode field if non-nil, zero value otherwise.

### GetPartnerProvidedFleetCountryCodeOk

`func (o *InvitationCreate) GetPartnerProvidedFleetCountryCodeOk() (*string, bool)`

GetPartnerProvidedFleetCountryCodeOk returns a tuple with the PartnerProvidedFleetCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerProvidedFleetCountryCode

`func (o *InvitationCreate) SetPartnerProvidedFleetCountryCode(v string)`

SetPartnerProvidedFleetCountryCode sets PartnerProvidedFleetCountryCode field to given value.

### HasPartnerProvidedFleetCountryCode

`func (o *InvitationCreate) HasPartnerProvidedFleetCountryCode() bool`

HasPartnerProvidedFleetCountryCode returns a boolean if a field has been set.

### SetPartnerProvidedFleetCountryCodeNil

`func (o *InvitationCreate) SetPartnerProvidedFleetCountryCodeNil(b bool)`

 SetPartnerProvidedFleetCountryCodeNil sets the value for PartnerProvidedFleetCountryCode to be an explicit nil

### UnsetPartnerProvidedFleetCountryCode
`func (o *InvitationCreate) UnsetPartnerProvidedFleetCountryCode()`

UnsetPartnerProvidedFleetCountryCode ensures that no value is present for PartnerProvidedFleetCountryCode, not even an explicit nil
### GetSuccessRedirectUrl

`func (o *InvitationCreate) GetSuccessRedirectUrl() string`

GetSuccessRedirectUrl returns the SuccessRedirectUrl field if non-nil, zero value otherwise.

### GetSuccessRedirectUrlOk

`func (o *InvitationCreate) GetSuccessRedirectUrlOk() (*string, bool)`

GetSuccessRedirectUrlOk returns a tuple with the SuccessRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRedirectUrl

`func (o *InvitationCreate) SetSuccessRedirectUrl(v string)`

SetSuccessRedirectUrl sets SuccessRedirectUrl field to given value.

### HasSuccessRedirectUrl

`func (o *InvitationCreate) HasSuccessRedirectUrl() bool`

HasSuccessRedirectUrl returns a boolean if a field has been set.

### SetSuccessRedirectUrlNil

`func (o *InvitationCreate) SetSuccessRedirectUrlNil(b bool)`

 SetSuccessRedirectUrlNil sets the value for SuccessRedirectUrl to be an explicit nil

### UnsetSuccessRedirectUrl
`func (o *InvitationCreate) UnsetSuccessRedirectUrl()`

UnsetSuccessRedirectUrl ensures that no value is present for SuccessRedirectUrl, not even an explicit nil
### GetFailureRedirectUrl

`func (o *InvitationCreate) GetFailureRedirectUrl() string`

GetFailureRedirectUrl returns the FailureRedirectUrl field if non-nil, zero value otherwise.

### GetFailureRedirectUrlOk

`func (o *InvitationCreate) GetFailureRedirectUrlOk() (*string, bool)`

GetFailureRedirectUrlOk returns a tuple with the FailureRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureRedirectUrl

`func (o *InvitationCreate) SetFailureRedirectUrl(v string)`

SetFailureRedirectUrl sets FailureRedirectUrl field to given value.

### HasFailureRedirectUrl

`func (o *InvitationCreate) HasFailureRedirectUrl() bool`

HasFailureRedirectUrl returns a boolean if a field has been set.

### SetFailureRedirectUrlNil

`func (o *InvitationCreate) SetFailureRedirectUrlNil(b bool)`

 SetFailureRedirectUrlNil sets the value for FailureRedirectUrl to be an explicit nil

### UnsetFailureRedirectUrl
`func (o *InvitationCreate) UnsetFailureRedirectUrl()`

UnsetFailureRedirectUrl ensures that no value is present for FailureRedirectUrl, not even an explicit nil
### GetCallbackUrl

`func (o *InvitationCreate) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *InvitationCreate) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *InvitationCreate) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *InvitationCreate) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### SetCallbackUrlNil

`func (o *InvitationCreate) SetCallbackUrlNil(b bool)`

 SetCallbackUrlNil sets the value for CallbackUrl to be an explicit nil

### UnsetCallbackUrl
`func (o *InvitationCreate) UnsetCallbackUrl()`

UnsetCallbackUrl ensures that no value is present for CallbackUrl, not even an explicit nil
### GetLimitTsps

`func (o *InvitationCreate) GetLimitTsps() []string`

GetLimitTsps returns the LimitTsps field if non-nil, zero value otherwise.

### GetLimitTspsOk

`func (o *InvitationCreate) GetLimitTspsOk() (*[]string, bool)`

GetLimitTspsOk returns a tuple with the LimitTsps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitTsps

`func (o *InvitationCreate) SetLimitTsps(v []string)`

SetLimitTsps sets LimitTsps field to given value.

### HasLimitTsps

`func (o *InvitationCreate) HasLimitTsps() bool`

HasLimitTsps returns a boolean if a field has been set.

### SetLimitTspsNil

`func (o *InvitationCreate) SetLimitTspsNil(b bool)`

 SetLimitTspsNil sets the value for LimitTsps to be an explicit nil

### UnsetLimitTsps
`func (o *InvitationCreate) UnsetLimitTsps()`

UnsetLimitTsps ensures that no value is present for LimitTsps, not even an explicit nil
### GetExpiresInHours

`func (o *InvitationCreate) GetExpiresInHours() int32`

GetExpiresInHours returns the ExpiresInHours field if non-nil, zero value otherwise.

### GetExpiresInHoursOk

`func (o *InvitationCreate) GetExpiresInHoursOk() (*int32, bool)`

GetExpiresInHoursOk returns a tuple with the ExpiresInHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInHours

`func (o *InvitationCreate) SetExpiresInHours(v int32)`

SetExpiresInHours sets ExpiresInHours field to given value.

### HasExpiresInHours

`func (o *InvitationCreate) HasExpiresInHours() bool`

HasExpiresInHours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


