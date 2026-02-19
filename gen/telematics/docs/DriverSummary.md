# DriverSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FleetId** | **string** | Catena fleet identifier. | 
**FleetRef** | **NullableString** |  | 
**ConnectionId** | **string** | Catena connection identifier through which this driver was ingested. | 
**UserId** | **string** | Unique Catena identifier for the driver. | 
**TspId** | **NullableString** |  | 
**TspSlug** | **NullableString** |  | 
**SourceName** | [**NullableTspEnum**](TspEnum.md) |  | 
**SourceId** | **NullableString** |  | 
**EmployeeNumber** | **NullableString** |  | 
**FirstName** | **NullableString** |  | 
**LastName** | **NullableString** |  | 
**Username** | **NullableString** |  | 
**Status** | **NullableString** |  | 
**PhoneNumber** | **NullableString** |  | 
**LicenseCountry** | **NullableString** |  | 
**LicenseRegion** | **NullableString** |  | 
**LicenseNumber** | **NullableString** |  | 
**LicenseExpiration** | **NullableString** |  | 
**HosRulesetCode** | **NullableString** |  | 
**SafetyEvents30d** | Pointer to **int32** | Count of safety events recorded in the last 30 days. | [optional] [default to 0]
**HosViolations30d** | Pointer to **int32** | Count of HOS violations recorded in the last 30 days. | [optional] [default to 0]

## Methods

### NewDriverSummary

`func NewDriverSummary(fleetId string, fleetRef NullableString, connectionId string, userId string, tspId NullableString, tspSlug NullableString, sourceName NullableTspEnum, sourceId NullableString, employeeNumber NullableString, firstName NullableString, lastName NullableString, username NullableString, status NullableString, phoneNumber NullableString, licenseCountry NullableString, licenseRegion NullableString, licenseNumber NullableString, licenseExpiration NullableString, hosRulesetCode NullableString, ) *DriverSummary`

NewDriverSummary instantiates a new DriverSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDriverSummaryWithDefaults

`func NewDriverSummaryWithDefaults() *DriverSummary`

NewDriverSummaryWithDefaults instantiates a new DriverSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *DriverSummary) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *DriverSummary) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *DriverSummary) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *DriverSummary) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *DriverSummary) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *DriverSummary) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.


### SetFleetRefNil

`func (o *DriverSummary) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *DriverSummary) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetConnectionId

`func (o *DriverSummary) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DriverSummary) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DriverSummary) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetUserId

`func (o *DriverSummary) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DriverSummary) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DriverSummary) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTspId

`func (o *DriverSummary) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *DriverSummary) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *DriverSummary) SetTspId(v string)`

SetTspId sets TspId field to given value.


### SetTspIdNil

`func (o *DriverSummary) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *DriverSummary) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetTspSlug

`func (o *DriverSummary) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *DriverSummary) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *DriverSummary) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.


### SetTspSlugNil

`func (o *DriverSummary) SetTspSlugNil(b bool)`

 SetTspSlugNil sets the value for TspSlug to be an explicit nil

### UnsetTspSlug
`func (o *DriverSummary) UnsetTspSlug()`

UnsetTspSlug ensures that no value is present for TspSlug, not even an explicit nil
### GetSourceName

`func (o *DriverSummary) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *DriverSummary) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *DriverSummary) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### SetSourceNameNil

`func (o *DriverSummary) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *DriverSummary) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetSourceId

`func (o *DriverSummary) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *DriverSummary) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *DriverSummary) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### SetSourceIdNil

`func (o *DriverSummary) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *DriverSummary) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetEmployeeNumber

`func (o *DriverSummary) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *DriverSummary) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *DriverSummary) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.


### SetEmployeeNumberNil

`func (o *DriverSummary) SetEmployeeNumberNil(b bool)`

 SetEmployeeNumberNil sets the value for EmployeeNumber to be an explicit nil

### UnsetEmployeeNumber
`func (o *DriverSummary) UnsetEmployeeNumber()`

UnsetEmployeeNumber ensures that no value is present for EmployeeNumber, not even an explicit nil
### GetFirstName

`func (o *DriverSummary) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *DriverSummary) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *DriverSummary) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### SetFirstNameNil

`func (o *DriverSummary) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *DriverSummary) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *DriverSummary) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *DriverSummary) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *DriverSummary) SetLastName(v string)`

SetLastName sets LastName field to given value.


### SetLastNameNil

`func (o *DriverSummary) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *DriverSummary) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetUsername

`func (o *DriverSummary) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DriverSummary) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DriverSummary) SetUsername(v string)`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *DriverSummary) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *DriverSummary) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetStatus

`func (o *DriverSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DriverSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DriverSummary) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *DriverSummary) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *DriverSummary) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetPhoneNumber

`func (o *DriverSummary) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *DriverSummary) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *DriverSummary) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### SetPhoneNumberNil

`func (o *DriverSummary) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *DriverSummary) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetLicenseCountry

`func (o *DriverSummary) GetLicenseCountry() string`

GetLicenseCountry returns the LicenseCountry field if non-nil, zero value otherwise.

### GetLicenseCountryOk

`func (o *DriverSummary) GetLicenseCountryOk() (*string, bool)`

GetLicenseCountryOk returns a tuple with the LicenseCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCountry

`func (o *DriverSummary) SetLicenseCountry(v string)`

SetLicenseCountry sets LicenseCountry field to given value.


### SetLicenseCountryNil

`func (o *DriverSummary) SetLicenseCountryNil(b bool)`

 SetLicenseCountryNil sets the value for LicenseCountry to be an explicit nil

### UnsetLicenseCountry
`func (o *DriverSummary) UnsetLicenseCountry()`

UnsetLicenseCountry ensures that no value is present for LicenseCountry, not even an explicit nil
### GetLicenseRegion

`func (o *DriverSummary) GetLicenseRegion() string`

GetLicenseRegion returns the LicenseRegion field if non-nil, zero value otherwise.

### GetLicenseRegionOk

`func (o *DriverSummary) GetLicenseRegionOk() (*string, bool)`

GetLicenseRegionOk returns a tuple with the LicenseRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseRegion

`func (o *DriverSummary) SetLicenseRegion(v string)`

SetLicenseRegion sets LicenseRegion field to given value.


### SetLicenseRegionNil

`func (o *DriverSummary) SetLicenseRegionNil(b bool)`

 SetLicenseRegionNil sets the value for LicenseRegion to be an explicit nil

### UnsetLicenseRegion
`func (o *DriverSummary) UnsetLicenseRegion()`

UnsetLicenseRegion ensures that no value is present for LicenseRegion, not even an explicit nil
### GetLicenseNumber

`func (o *DriverSummary) GetLicenseNumber() string`

GetLicenseNumber returns the LicenseNumber field if non-nil, zero value otherwise.

### GetLicenseNumberOk

`func (o *DriverSummary) GetLicenseNumberOk() (*string, bool)`

GetLicenseNumberOk returns a tuple with the LicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNumber

`func (o *DriverSummary) SetLicenseNumber(v string)`

SetLicenseNumber sets LicenseNumber field to given value.


### SetLicenseNumberNil

`func (o *DriverSummary) SetLicenseNumberNil(b bool)`

 SetLicenseNumberNil sets the value for LicenseNumber to be an explicit nil

### UnsetLicenseNumber
`func (o *DriverSummary) UnsetLicenseNumber()`

UnsetLicenseNumber ensures that no value is present for LicenseNumber, not even an explicit nil
### GetLicenseExpiration

`func (o *DriverSummary) GetLicenseExpiration() string`

GetLicenseExpiration returns the LicenseExpiration field if non-nil, zero value otherwise.

### GetLicenseExpirationOk

`func (o *DriverSummary) GetLicenseExpirationOk() (*string, bool)`

GetLicenseExpirationOk returns a tuple with the LicenseExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpiration

`func (o *DriverSummary) SetLicenseExpiration(v string)`

SetLicenseExpiration sets LicenseExpiration field to given value.


### SetLicenseExpirationNil

`func (o *DriverSummary) SetLicenseExpirationNil(b bool)`

 SetLicenseExpirationNil sets the value for LicenseExpiration to be an explicit nil

### UnsetLicenseExpiration
`func (o *DriverSummary) UnsetLicenseExpiration()`

UnsetLicenseExpiration ensures that no value is present for LicenseExpiration, not even an explicit nil
### GetHosRulesetCode

`func (o *DriverSummary) GetHosRulesetCode() string`

GetHosRulesetCode returns the HosRulesetCode field if non-nil, zero value otherwise.

### GetHosRulesetCodeOk

`func (o *DriverSummary) GetHosRulesetCodeOk() (*string, bool)`

GetHosRulesetCodeOk returns a tuple with the HosRulesetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosRulesetCode

`func (o *DriverSummary) SetHosRulesetCode(v string)`

SetHosRulesetCode sets HosRulesetCode field to given value.


### SetHosRulesetCodeNil

`func (o *DriverSummary) SetHosRulesetCodeNil(b bool)`

 SetHosRulesetCodeNil sets the value for HosRulesetCode to be an explicit nil

### UnsetHosRulesetCode
`func (o *DriverSummary) UnsetHosRulesetCode()`

UnsetHosRulesetCode ensures that no value is present for HosRulesetCode, not even an explicit nil
### GetSafetyEvents30d

`func (o *DriverSummary) GetSafetyEvents30d() int32`

GetSafetyEvents30d returns the SafetyEvents30d field if non-nil, zero value otherwise.

### GetSafetyEvents30dOk

`func (o *DriverSummary) GetSafetyEvents30dOk() (*int32, bool)`

GetSafetyEvents30dOk returns a tuple with the SafetyEvents30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafetyEvents30d

`func (o *DriverSummary) SetSafetyEvents30d(v int32)`

SetSafetyEvents30d sets SafetyEvents30d field to given value.

### HasSafetyEvents30d

`func (o *DriverSummary) HasSafetyEvents30d() bool`

HasSafetyEvents30d returns a boolean if a field has been set.

### GetHosViolations30d

`func (o *DriverSummary) GetHosViolations30d() int32`

GetHosViolations30d returns the HosViolations30d field if non-nil, zero value otherwise.

### GetHosViolations30dOk

`func (o *DriverSummary) GetHosViolations30dOk() (*int32, bool)`

GetHosViolations30dOk returns a tuple with the HosViolations30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosViolations30d

`func (o *DriverSummary) SetHosViolations30d(v int32)`

SetHosViolations30d sets HosViolations30d field to given value.

### HasHosViolations30d

`func (o *DriverSummary) HasHosViolations30d() bool`

HasHosViolations30d returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


