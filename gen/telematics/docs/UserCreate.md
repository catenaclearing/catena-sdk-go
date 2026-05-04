# UserCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | Unique identifier of the connection at Catena Telematics which will be used to create this resource. A connection represents a Fleet/TSP pairing. | 
**Username** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**Status** | Pointer to [**NullableUserStatusEnum**](UserStatusEnum.md) |  | [optional] 
**IsDriver** | Pointer to **NullableBool** |  | [optional] 
**UserDesignation** | Pointer to **NullableString** |  | [optional] 
**UserEmail** | Pointer to **NullableString** |  | [optional] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**PhoneNumber** | Pointer to **NullableString** |  | [optional] 
**CountryCode** | Pointer to **NullableString** |  | [optional] 
**LicenseCountry** | Pointer to **NullableString** |  | [optional] 
**LicenseRegion** | Pointer to **NullableString** |  | [optional] 
**LicenseNumber** | Pointer to **NullableString** |  | [optional] 
**LicenseExpiration** | Pointer to **NullableString** |  | [optional] 
**EmployeeNumber** | Pointer to **NullableString** |  | [optional] 
**CarrierNumber** | Pointer to **NullableString** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**HosRulesetCode** | Pointer to [**NullableHosRulesetCodeEnum**](HosRulesetCodeEnum.md) |  | [optional] 
**AllowYardMove** | Pointer to **NullableBool** |  | [optional] 
**AllowPersonalConveyance** | Pointer to **NullableBool** |  | [optional] 
**AllowAdverseDriving** | Pointer to **NullableBool** |  | [optional] 
**DefaultTimeZone** | Pointer to [**NullableTimezoneCodeEnum**](TimezoneCodeEnum.md) |  | [optional] 
**HosComplianceType** | Pointer to [**NullableHosComplianceTypeEnum**](HosComplianceTypeEnum.md) |  | [optional] 
**HosTrackingMethod** | Pointer to [**NullableHosTrackingMethodEnum**](HosTrackingMethodEnum.md) |  | [optional] 

## Methods

### NewUserCreate

`func NewUserCreate(connectionId string, ) *UserCreate`

NewUserCreate instantiates a new UserCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateWithDefaults

`func NewUserCreateWithDefaults() *UserCreate`

NewUserCreateWithDefaults instantiates a new UserCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *UserCreate) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *UserCreate) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *UserCreate) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetUsername

`func (o *UserCreate) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserCreate) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserCreate) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserCreate) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UserCreate) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UserCreate) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetIsActive

`func (o *UserCreate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UserCreate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UserCreate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UserCreate) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *UserCreate) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *UserCreate) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetStatus

`func (o *UserCreate) GetStatus() UserStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserCreate) GetStatusOk() (*UserStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserCreate) SetStatus(v UserStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserCreate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *UserCreate) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *UserCreate) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIsDriver

`func (o *UserCreate) GetIsDriver() bool`

GetIsDriver returns the IsDriver field if non-nil, zero value otherwise.

### GetIsDriverOk

`func (o *UserCreate) GetIsDriverOk() (*bool, bool)`

GetIsDriverOk returns a tuple with the IsDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDriver

`func (o *UserCreate) SetIsDriver(v bool)`

SetIsDriver sets IsDriver field to given value.

### HasIsDriver

`func (o *UserCreate) HasIsDriver() bool`

HasIsDriver returns a boolean if a field has been set.

### SetIsDriverNil

`func (o *UserCreate) SetIsDriverNil(b bool)`

 SetIsDriverNil sets the value for IsDriver to be an explicit nil

### UnsetIsDriver
`func (o *UserCreate) UnsetIsDriver()`

UnsetIsDriver ensures that no value is present for IsDriver, not even an explicit nil
### GetUserDesignation

`func (o *UserCreate) GetUserDesignation() string`

GetUserDesignation returns the UserDesignation field if non-nil, zero value otherwise.

### GetUserDesignationOk

`func (o *UserCreate) GetUserDesignationOk() (*string, bool)`

GetUserDesignationOk returns a tuple with the UserDesignation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDesignation

`func (o *UserCreate) SetUserDesignation(v string)`

SetUserDesignation sets UserDesignation field to given value.

### HasUserDesignation

`func (o *UserCreate) HasUserDesignation() bool`

HasUserDesignation returns a boolean if a field has been set.

### SetUserDesignationNil

`func (o *UserCreate) SetUserDesignationNil(b bool)`

 SetUserDesignationNil sets the value for UserDesignation to be an explicit nil

### UnsetUserDesignation
`func (o *UserCreate) UnsetUserDesignation()`

UnsetUserDesignation ensures that no value is present for UserDesignation, not even an explicit nil
### GetUserEmail

`func (o *UserCreate) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *UserCreate) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *UserCreate) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *UserCreate) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmailNil

`func (o *UserCreate) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *UserCreate) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
### GetFirstName

`func (o *UserCreate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserCreate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserCreate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserCreate) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *UserCreate) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UserCreate) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UserCreate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserCreate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserCreate) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserCreate) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UserCreate) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UserCreate) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetPhoneNumber

`func (o *UserCreate) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserCreate) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserCreate) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserCreate) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *UserCreate) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *UserCreate) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetCountryCode

`func (o *UserCreate) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UserCreate) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UserCreate) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UserCreate) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *UserCreate) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *UserCreate) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetLicenseCountry

`func (o *UserCreate) GetLicenseCountry() string`

GetLicenseCountry returns the LicenseCountry field if non-nil, zero value otherwise.

### GetLicenseCountryOk

`func (o *UserCreate) GetLicenseCountryOk() (*string, bool)`

GetLicenseCountryOk returns a tuple with the LicenseCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCountry

`func (o *UserCreate) SetLicenseCountry(v string)`

SetLicenseCountry sets LicenseCountry field to given value.

### HasLicenseCountry

`func (o *UserCreate) HasLicenseCountry() bool`

HasLicenseCountry returns a boolean if a field has been set.

### SetLicenseCountryNil

`func (o *UserCreate) SetLicenseCountryNil(b bool)`

 SetLicenseCountryNil sets the value for LicenseCountry to be an explicit nil

### UnsetLicenseCountry
`func (o *UserCreate) UnsetLicenseCountry()`

UnsetLicenseCountry ensures that no value is present for LicenseCountry, not even an explicit nil
### GetLicenseRegion

`func (o *UserCreate) GetLicenseRegion() string`

GetLicenseRegion returns the LicenseRegion field if non-nil, zero value otherwise.

### GetLicenseRegionOk

`func (o *UserCreate) GetLicenseRegionOk() (*string, bool)`

GetLicenseRegionOk returns a tuple with the LicenseRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseRegion

`func (o *UserCreate) SetLicenseRegion(v string)`

SetLicenseRegion sets LicenseRegion field to given value.

### HasLicenseRegion

`func (o *UserCreate) HasLicenseRegion() bool`

HasLicenseRegion returns a boolean if a field has been set.

### SetLicenseRegionNil

`func (o *UserCreate) SetLicenseRegionNil(b bool)`

 SetLicenseRegionNil sets the value for LicenseRegion to be an explicit nil

### UnsetLicenseRegion
`func (o *UserCreate) UnsetLicenseRegion()`

UnsetLicenseRegion ensures that no value is present for LicenseRegion, not even an explicit nil
### GetLicenseNumber

`func (o *UserCreate) GetLicenseNumber() string`

GetLicenseNumber returns the LicenseNumber field if non-nil, zero value otherwise.

### GetLicenseNumberOk

`func (o *UserCreate) GetLicenseNumberOk() (*string, bool)`

GetLicenseNumberOk returns a tuple with the LicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNumber

`func (o *UserCreate) SetLicenseNumber(v string)`

SetLicenseNumber sets LicenseNumber field to given value.

### HasLicenseNumber

`func (o *UserCreate) HasLicenseNumber() bool`

HasLicenseNumber returns a boolean if a field has been set.

### SetLicenseNumberNil

`func (o *UserCreate) SetLicenseNumberNil(b bool)`

 SetLicenseNumberNil sets the value for LicenseNumber to be an explicit nil

### UnsetLicenseNumber
`func (o *UserCreate) UnsetLicenseNumber()`

UnsetLicenseNumber ensures that no value is present for LicenseNumber, not even an explicit nil
### GetLicenseExpiration

`func (o *UserCreate) GetLicenseExpiration() string`

GetLicenseExpiration returns the LicenseExpiration field if non-nil, zero value otherwise.

### GetLicenseExpirationOk

`func (o *UserCreate) GetLicenseExpirationOk() (*string, bool)`

GetLicenseExpirationOk returns a tuple with the LicenseExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpiration

`func (o *UserCreate) SetLicenseExpiration(v string)`

SetLicenseExpiration sets LicenseExpiration field to given value.

### HasLicenseExpiration

`func (o *UserCreate) HasLicenseExpiration() bool`

HasLicenseExpiration returns a boolean if a field has been set.

### SetLicenseExpirationNil

`func (o *UserCreate) SetLicenseExpirationNil(b bool)`

 SetLicenseExpirationNil sets the value for LicenseExpiration to be an explicit nil

### UnsetLicenseExpiration
`func (o *UserCreate) UnsetLicenseExpiration()`

UnsetLicenseExpiration ensures that no value is present for LicenseExpiration, not even an explicit nil
### GetEmployeeNumber

`func (o *UserCreate) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *UserCreate) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *UserCreate) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *UserCreate) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### SetEmployeeNumberNil

`func (o *UserCreate) SetEmployeeNumberNil(b bool)`

 SetEmployeeNumberNil sets the value for EmployeeNumber to be an explicit nil

### UnsetEmployeeNumber
`func (o *UserCreate) UnsetEmployeeNumber()`

UnsetEmployeeNumber ensures that no value is present for EmployeeNumber, not even an explicit nil
### GetCarrierNumber

`func (o *UserCreate) GetCarrierNumber() string`

GetCarrierNumber returns the CarrierNumber field if non-nil, zero value otherwise.

### GetCarrierNumberOk

`func (o *UserCreate) GetCarrierNumberOk() (*string, bool)`

GetCarrierNumberOk returns a tuple with the CarrierNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierNumber

`func (o *UserCreate) SetCarrierNumber(v string)`

SetCarrierNumber sets CarrierNumber field to given value.

### HasCarrierNumber

`func (o *UserCreate) HasCarrierNumber() bool`

HasCarrierNumber returns a boolean if a field has been set.

### SetCarrierNumberNil

`func (o *UserCreate) SetCarrierNumberNil(b bool)`

 SetCarrierNumberNil sets the value for CarrierNumber to be an explicit nil

### UnsetCarrierNumber
`func (o *UserCreate) UnsetCarrierNumber()`

UnsetCarrierNumber ensures that no value is present for CarrierNumber, not even an explicit nil
### GetNotes

`func (o *UserCreate) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UserCreate) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UserCreate) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UserCreate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *UserCreate) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *UserCreate) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetHosRulesetCode

`func (o *UserCreate) GetHosRulesetCode() HosRulesetCodeEnum`

GetHosRulesetCode returns the HosRulesetCode field if non-nil, zero value otherwise.

### GetHosRulesetCodeOk

`func (o *UserCreate) GetHosRulesetCodeOk() (*HosRulesetCodeEnum, bool)`

GetHosRulesetCodeOk returns a tuple with the HosRulesetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosRulesetCode

`func (o *UserCreate) SetHosRulesetCode(v HosRulesetCodeEnum)`

SetHosRulesetCode sets HosRulesetCode field to given value.

### HasHosRulesetCode

`func (o *UserCreate) HasHosRulesetCode() bool`

HasHosRulesetCode returns a boolean if a field has been set.

### SetHosRulesetCodeNil

`func (o *UserCreate) SetHosRulesetCodeNil(b bool)`

 SetHosRulesetCodeNil sets the value for HosRulesetCode to be an explicit nil

### UnsetHosRulesetCode
`func (o *UserCreate) UnsetHosRulesetCode()`

UnsetHosRulesetCode ensures that no value is present for HosRulesetCode, not even an explicit nil
### GetAllowYardMove

`func (o *UserCreate) GetAllowYardMove() bool`

GetAllowYardMove returns the AllowYardMove field if non-nil, zero value otherwise.

### GetAllowYardMoveOk

`func (o *UserCreate) GetAllowYardMoveOk() (*bool, bool)`

GetAllowYardMoveOk returns a tuple with the AllowYardMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowYardMove

`func (o *UserCreate) SetAllowYardMove(v bool)`

SetAllowYardMove sets AllowYardMove field to given value.

### HasAllowYardMove

`func (o *UserCreate) HasAllowYardMove() bool`

HasAllowYardMove returns a boolean if a field has been set.

### SetAllowYardMoveNil

`func (o *UserCreate) SetAllowYardMoveNil(b bool)`

 SetAllowYardMoveNil sets the value for AllowYardMove to be an explicit nil

### UnsetAllowYardMove
`func (o *UserCreate) UnsetAllowYardMove()`

UnsetAllowYardMove ensures that no value is present for AllowYardMove, not even an explicit nil
### GetAllowPersonalConveyance

`func (o *UserCreate) GetAllowPersonalConveyance() bool`

GetAllowPersonalConveyance returns the AllowPersonalConveyance field if non-nil, zero value otherwise.

### GetAllowPersonalConveyanceOk

`func (o *UserCreate) GetAllowPersonalConveyanceOk() (*bool, bool)`

GetAllowPersonalConveyanceOk returns a tuple with the AllowPersonalConveyance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPersonalConveyance

`func (o *UserCreate) SetAllowPersonalConveyance(v bool)`

SetAllowPersonalConveyance sets AllowPersonalConveyance field to given value.

### HasAllowPersonalConveyance

`func (o *UserCreate) HasAllowPersonalConveyance() bool`

HasAllowPersonalConveyance returns a boolean if a field has been set.

### SetAllowPersonalConveyanceNil

`func (o *UserCreate) SetAllowPersonalConveyanceNil(b bool)`

 SetAllowPersonalConveyanceNil sets the value for AllowPersonalConveyance to be an explicit nil

### UnsetAllowPersonalConveyance
`func (o *UserCreate) UnsetAllowPersonalConveyance()`

UnsetAllowPersonalConveyance ensures that no value is present for AllowPersonalConveyance, not even an explicit nil
### GetAllowAdverseDriving

`func (o *UserCreate) GetAllowAdverseDriving() bool`

GetAllowAdverseDriving returns the AllowAdverseDriving field if non-nil, zero value otherwise.

### GetAllowAdverseDrivingOk

`func (o *UserCreate) GetAllowAdverseDrivingOk() (*bool, bool)`

GetAllowAdverseDrivingOk returns a tuple with the AllowAdverseDriving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAdverseDriving

`func (o *UserCreate) SetAllowAdverseDriving(v bool)`

SetAllowAdverseDriving sets AllowAdverseDriving field to given value.

### HasAllowAdverseDriving

`func (o *UserCreate) HasAllowAdverseDriving() bool`

HasAllowAdverseDriving returns a boolean if a field has been set.

### SetAllowAdverseDrivingNil

`func (o *UserCreate) SetAllowAdverseDrivingNil(b bool)`

 SetAllowAdverseDrivingNil sets the value for AllowAdverseDriving to be an explicit nil

### UnsetAllowAdverseDriving
`func (o *UserCreate) UnsetAllowAdverseDriving()`

UnsetAllowAdverseDriving ensures that no value is present for AllowAdverseDriving, not even an explicit nil
### GetDefaultTimeZone

`func (o *UserCreate) GetDefaultTimeZone() TimezoneCodeEnum`

GetDefaultTimeZone returns the DefaultTimeZone field if non-nil, zero value otherwise.

### GetDefaultTimeZoneOk

`func (o *UserCreate) GetDefaultTimeZoneOk() (*TimezoneCodeEnum, bool)`

GetDefaultTimeZoneOk returns a tuple with the DefaultTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeZone

`func (o *UserCreate) SetDefaultTimeZone(v TimezoneCodeEnum)`

SetDefaultTimeZone sets DefaultTimeZone field to given value.

### HasDefaultTimeZone

`func (o *UserCreate) HasDefaultTimeZone() bool`

HasDefaultTimeZone returns a boolean if a field has been set.

### SetDefaultTimeZoneNil

`func (o *UserCreate) SetDefaultTimeZoneNil(b bool)`

 SetDefaultTimeZoneNil sets the value for DefaultTimeZone to be an explicit nil

### UnsetDefaultTimeZone
`func (o *UserCreate) UnsetDefaultTimeZone()`

UnsetDefaultTimeZone ensures that no value is present for DefaultTimeZone, not even an explicit nil
### GetHosComplianceType

`func (o *UserCreate) GetHosComplianceType() HosComplianceTypeEnum`

GetHosComplianceType returns the HosComplianceType field if non-nil, zero value otherwise.

### GetHosComplianceTypeOk

`func (o *UserCreate) GetHosComplianceTypeOk() (*HosComplianceTypeEnum, bool)`

GetHosComplianceTypeOk returns a tuple with the HosComplianceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosComplianceType

`func (o *UserCreate) SetHosComplianceType(v HosComplianceTypeEnum)`

SetHosComplianceType sets HosComplianceType field to given value.

### HasHosComplianceType

`func (o *UserCreate) HasHosComplianceType() bool`

HasHosComplianceType returns a boolean if a field has been set.

### SetHosComplianceTypeNil

`func (o *UserCreate) SetHosComplianceTypeNil(b bool)`

 SetHosComplianceTypeNil sets the value for HosComplianceType to be an explicit nil

### UnsetHosComplianceType
`func (o *UserCreate) UnsetHosComplianceType()`

UnsetHosComplianceType ensures that no value is present for HosComplianceType, not even an explicit nil
### GetHosTrackingMethod

`func (o *UserCreate) GetHosTrackingMethod() HosTrackingMethodEnum`

GetHosTrackingMethod returns the HosTrackingMethod field if non-nil, zero value otherwise.

### GetHosTrackingMethodOk

`func (o *UserCreate) GetHosTrackingMethodOk() (*HosTrackingMethodEnum, bool)`

GetHosTrackingMethodOk returns a tuple with the HosTrackingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosTrackingMethod

`func (o *UserCreate) SetHosTrackingMethod(v HosTrackingMethodEnum)`

SetHosTrackingMethod sets HosTrackingMethod field to given value.

### HasHosTrackingMethod

`func (o *UserCreate) HasHosTrackingMethod() bool`

HasHosTrackingMethod returns a boolean if a field has been set.

### SetHosTrackingMethodNil

`func (o *UserCreate) SetHosTrackingMethodNil(b bool)`

 SetHosTrackingMethodNil sets the value for HosTrackingMethod to be an explicit nil

### UnsetHosTrackingMethod
`func (o *UserCreate) UnsetHosTrackingMethod()`

UnsetHosTrackingMethod ensures that no value is present for HosTrackingMethod, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


