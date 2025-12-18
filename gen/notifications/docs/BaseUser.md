# BaseUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal unique identifier for the telematics event record (Catena PK). | 
**FleetId** | **string** | The Catena fleet this record belongs to (multi-tenant scope). | 
**SourceName** | [**TspEnum**](TspEnum.md) | The name of the source | 
**ConnectionId** | **string** | The specific fleet↔TSP connection through which this record was sourced. | 
**SourceId** | **string** | The ID of the record in the TSP or a deterministic ID/Hash generated from a composite unique key | 
**CreatedAt** | **time.Time** | Immutable: first time this record was ingested into our system. | 
**UpdatedAt** | **time.Time** | Last time we modified this record in our system. | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**OccurredAt** | **time.Time** | When the underlying event/observation occurred, as reported by the TSP, or the moment it was ingested by us if not available. | 
**ExecutionId** | Pointer to **NullableString** |  | [optional] 
**ScheduleId** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**EndedAt** | Pointer to **NullableTime** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
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
**CompanyGroups** | Pointer to **map[string]interface{}** |  | [optional] 
**PrivateUserGroups** | Pointer to **map[string]interface{}** |  | [optional] 
**ReportGroups** | Pointer to **map[string]interface{}** |  | [optional] 
**SecurityGroups** | Pointer to **map[string]interface{}** |  | [optional] 
**AuthorityName** | Pointer to **NullableString** |  | [optional] 
**AuthorityAddress** | Pointer to **NullableString** |  | [optional] 
**CompanyName** | Pointer to **NullableString** |  | [optional] 
**CompanyAddress** | Pointer to **NullableString** |  | [optional] 
**CarrierNumber** | Pointer to **NullableString** |  | [optional] 
**LastTspLogin** | Pointer to **NullableTime** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**HosRulesetCode** | Pointer to [**NullableHosRulesetCodeEnum**](HosRulesetCodeEnum.md) |  | [optional] 
**AllowYardMove** | Pointer to **NullableBool** |  | [optional] 
**AllowPersonalConveyance** | Pointer to **NullableBool** |  | [optional] 
**AllowAdverseDriving** | Pointer to **NullableBool** |  | [optional] 
**DefaultTimeZone** | Pointer to [**NullableTimezoneCodeEnum**](TimezoneCodeEnum.md) |  | [optional] 

## Methods

### NewBaseUser

`func NewBaseUser(id string, fleetId string, sourceName TspEnum, connectionId string, sourceId string, createdAt time.Time, updatedAt time.Time, occurredAt time.Time, ) *BaseUser`

NewBaseUser instantiates a new BaseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseUserWithDefaults

`func NewBaseUserWithDefaults() *BaseUser`

NewBaseUserWithDefaults instantiates a new BaseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseUser) SetId(v string)`

SetId sets Id field to given value.


### GetFleetId

`func (o *BaseUser) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *BaseUser) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *BaseUser) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetSourceName

`func (o *BaseUser) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseUser) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseUser) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetConnectionId

`func (o *BaseUser) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseUser) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseUser) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceId

`func (o *BaseUser) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseUser) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseUser) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetCreatedAt

`func (o *BaseUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *BaseUser) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *BaseUser) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *BaseUser) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *BaseUser) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *BaseUser) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *BaseUser) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetOccurredAt

`func (o *BaseUser) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *BaseUser) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *BaseUser) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.


### GetExecutionId

`func (o *BaseUser) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *BaseUser) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *BaseUser) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *BaseUser) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *BaseUser) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *BaseUser) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *BaseUser) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *BaseUser) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *BaseUser) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *BaseUser) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *BaseUser) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *BaseUser) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetUsername

`func (o *BaseUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BaseUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BaseUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *BaseUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *BaseUser) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *BaseUser) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetStartedAt

`func (o *BaseUser) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *BaseUser) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *BaseUser) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *BaseUser) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *BaseUser) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *BaseUser) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *BaseUser) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *BaseUser) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *BaseUser) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *BaseUser) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *BaseUser) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *BaseUser) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetIsActive

`func (o *BaseUser) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *BaseUser) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *BaseUser) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *BaseUser) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *BaseUser) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *BaseUser) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetStatus

`func (o *BaseUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseUser) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseUser) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BaseUser) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BaseUser) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIsDriver

`func (o *BaseUser) GetIsDriver() bool`

GetIsDriver returns the IsDriver field if non-nil, zero value otherwise.

### GetIsDriverOk

`func (o *BaseUser) GetIsDriverOk() (*bool, bool)`

GetIsDriverOk returns a tuple with the IsDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDriver

`func (o *BaseUser) SetIsDriver(v bool)`

SetIsDriver sets IsDriver field to given value.

### HasIsDriver

`func (o *BaseUser) HasIsDriver() bool`

HasIsDriver returns a boolean if a field has been set.

### SetIsDriverNil

`func (o *BaseUser) SetIsDriverNil(b bool)`

 SetIsDriverNil sets the value for IsDriver to be an explicit nil

### UnsetIsDriver
`func (o *BaseUser) UnsetIsDriver()`

UnsetIsDriver ensures that no value is present for IsDriver, not even an explicit nil
### GetUserDesignation

`func (o *BaseUser) GetUserDesignation() string`

GetUserDesignation returns the UserDesignation field if non-nil, zero value otherwise.

### GetUserDesignationOk

`func (o *BaseUser) GetUserDesignationOk() (*string, bool)`

GetUserDesignationOk returns a tuple with the UserDesignation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDesignation

`func (o *BaseUser) SetUserDesignation(v string)`

SetUserDesignation sets UserDesignation field to given value.

### HasUserDesignation

`func (o *BaseUser) HasUserDesignation() bool`

HasUserDesignation returns a boolean if a field has been set.

### SetUserDesignationNil

`func (o *BaseUser) SetUserDesignationNil(b bool)`

 SetUserDesignationNil sets the value for UserDesignation to be an explicit nil

### UnsetUserDesignation
`func (o *BaseUser) UnsetUserDesignation()`

UnsetUserDesignation ensures that no value is present for UserDesignation, not even an explicit nil
### GetUserEmail

`func (o *BaseUser) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *BaseUser) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *BaseUser) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *BaseUser) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmailNil

`func (o *BaseUser) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *BaseUser) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
### GetFirstName

`func (o *BaseUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BaseUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BaseUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BaseUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *BaseUser) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *BaseUser) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *BaseUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BaseUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BaseUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *BaseUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *BaseUser) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *BaseUser) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetPhoneNumber

`func (o *BaseUser) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *BaseUser) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *BaseUser) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *BaseUser) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *BaseUser) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *BaseUser) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetCountryCode

`func (o *BaseUser) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *BaseUser) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *BaseUser) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *BaseUser) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *BaseUser) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *BaseUser) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetLicenseCountry

`func (o *BaseUser) GetLicenseCountry() string`

GetLicenseCountry returns the LicenseCountry field if non-nil, zero value otherwise.

### GetLicenseCountryOk

`func (o *BaseUser) GetLicenseCountryOk() (*string, bool)`

GetLicenseCountryOk returns a tuple with the LicenseCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCountry

`func (o *BaseUser) SetLicenseCountry(v string)`

SetLicenseCountry sets LicenseCountry field to given value.

### HasLicenseCountry

`func (o *BaseUser) HasLicenseCountry() bool`

HasLicenseCountry returns a boolean if a field has been set.

### SetLicenseCountryNil

`func (o *BaseUser) SetLicenseCountryNil(b bool)`

 SetLicenseCountryNil sets the value for LicenseCountry to be an explicit nil

### UnsetLicenseCountry
`func (o *BaseUser) UnsetLicenseCountry()`

UnsetLicenseCountry ensures that no value is present for LicenseCountry, not even an explicit nil
### GetLicenseRegion

`func (o *BaseUser) GetLicenseRegion() string`

GetLicenseRegion returns the LicenseRegion field if non-nil, zero value otherwise.

### GetLicenseRegionOk

`func (o *BaseUser) GetLicenseRegionOk() (*string, bool)`

GetLicenseRegionOk returns a tuple with the LicenseRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseRegion

`func (o *BaseUser) SetLicenseRegion(v string)`

SetLicenseRegion sets LicenseRegion field to given value.

### HasLicenseRegion

`func (o *BaseUser) HasLicenseRegion() bool`

HasLicenseRegion returns a boolean if a field has been set.

### SetLicenseRegionNil

`func (o *BaseUser) SetLicenseRegionNil(b bool)`

 SetLicenseRegionNil sets the value for LicenseRegion to be an explicit nil

### UnsetLicenseRegion
`func (o *BaseUser) UnsetLicenseRegion()`

UnsetLicenseRegion ensures that no value is present for LicenseRegion, not even an explicit nil
### GetLicenseNumber

`func (o *BaseUser) GetLicenseNumber() string`

GetLicenseNumber returns the LicenseNumber field if non-nil, zero value otherwise.

### GetLicenseNumberOk

`func (o *BaseUser) GetLicenseNumberOk() (*string, bool)`

GetLicenseNumberOk returns a tuple with the LicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNumber

`func (o *BaseUser) SetLicenseNumber(v string)`

SetLicenseNumber sets LicenseNumber field to given value.

### HasLicenseNumber

`func (o *BaseUser) HasLicenseNumber() bool`

HasLicenseNumber returns a boolean if a field has been set.

### SetLicenseNumberNil

`func (o *BaseUser) SetLicenseNumberNil(b bool)`

 SetLicenseNumberNil sets the value for LicenseNumber to be an explicit nil

### UnsetLicenseNumber
`func (o *BaseUser) UnsetLicenseNumber()`

UnsetLicenseNumber ensures that no value is present for LicenseNumber, not even an explicit nil
### GetLicenseExpiration

`func (o *BaseUser) GetLicenseExpiration() string`

GetLicenseExpiration returns the LicenseExpiration field if non-nil, zero value otherwise.

### GetLicenseExpirationOk

`func (o *BaseUser) GetLicenseExpirationOk() (*string, bool)`

GetLicenseExpirationOk returns a tuple with the LicenseExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpiration

`func (o *BaseUser) SetLicenseExpiration(v string)`

SetLicenseExpiration sets LicenseExpiration field to given value.

### HasLicenseExpiration

`func (o *BaseUser) HasLicenseExpiration() bool`

HasLicenseExpiration returns a boolean if a field has been set.

### SetLicenseExpirationNil

`func (o *BaseUser) SetLicenseExpirationNil(b bool)`

 SetLicenseExpirationNil sets the value for LicenseExpiration to be an explicit nil

### UnsetLicenseExpiration
`func (o *BaseUser) UnsetLicenseExpiration()`

UnsetLicenseExpiration ensures that no value is present for LicenseExpiration, not even an explicit nil
### GetEmployeeNumber

`func (o *BaseUser) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *BaseUser) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *BaseUser) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *BaseUser) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### SetEmployeeNumberNil

`func (o *BaseUser) SetEmployeeNumberNil(b bool)`

 SetEmployeeNumberNil sets the value for EmployeeNumber to be an explicit nil

### UnsetEmployeeNumber
`func (o *BaseUser) UnsetEmployeeNumber()`

UnsetEmployeeNumber ensures that no value is present for EmployeeNumber, not even an explicit nil
### GetCompanyGroups

`func (o *BaseUser) GetCompanyGroups() map[string]interface{}`

GetCompanyGroups returns the CompanyGroups field if non-nil, zero value otherwise.

### GetCompanyGroupsOk

`func (o *BaseUser) GetCompanyGroupsOk() (*map[string]interface{}, bool)`

GetCompanyGroupsOk returns a tuple with the CompanyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyGroups

`func (o *BaseUser) SetCompanyGroups(v map[string]interface{})`

SetCompanyGroups sets CompanyGroups field to given value.

### HasCompanyGroups

`func (o *BaseUser) HasCompanyGroups() bool`

HasCompanyGroups returns a boolean if a field has been set.

### SetCompanyGroupsNil

`func (o *BaseUser) SetCompanyGroupsNil(b bool)`

 SetCompanyGroupsNil sets the value for CompanyGroups to be an explicit nil

### UnsetCompanyGroups
`func (o *BaseUser) UnsetCompanyGroups()`

UnsetCompanyGroups ensures that no value is present for CompanyGroups, not even an explicit nil
### GetPrivateUserGroups

`func (o *BaseUser) GetPrivateUserGroups() map[string]interface{}`

GetPrivateUserGroups returns the PrivateUserGroups field if non-nil, zero value otherwise.

### GetPrivateUserGroupsOk

`func (o *BaseUser) GetPrivateUserGroupsOk() (*map[string]interface{}, bool)`

GetPrivateUserGroupsOk returns a tuple with the PrivateUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateUserGroups

`func (o *BaseUser) SetPrivateUserGroups(v map[string]interface{})`

SetPrivateUserGroups sets PrivateUserGroups field to given value.

### HasPrivateUserGroups

`func (o *BaseUser) HasPrivateUserGroups() bool`

HasPrivateUserGroups returns a boolean if a field has been set.

### SetPrivateUserGroupsNil

`func (o *BaseUser) SetPrivateUserGroupsNil(b bool)`

 SetPrivateUserGroupsNil sets the value for PrivateUserGroups to be an explicit nil

### UnsetPrivateUserGroups
`func (o *BaseUser) UnsetPrivateUserGroups()`

UnsetPrivateUserGroups ensures that no value is present for PrivateUserGroups, not even an explicit nil
### GetReportGroups

`func (o *BaseUser) GetReportGroups() map[string]interface{}`

GetReportGroups returns the ReportGroups field if non-nil, zero value otherwise.

### GetReportGroupsOk

`func (o *BaseUser) GetReportGroupsOk() (*map[string]interface{}, bool)`

GetReportGroupsOk returns a tuple with the ReportGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportGroups

`func (o *BaseUser) SetReportGroups(v map[string]interface{})`

SetReportGroups sets ReportGroups field to given value.

### HasReportGroups

`func (o *BaseUser) HasReportGroups() bool`

HasReportGroups returns a boolean if a field has been set.

### SetReportGroupsNil

`func (o *BaseUser) SetReportGroupsNil(b bool)`

 SetReportGroupsNil sets the value for ReportGroups to be an explicit nil

### UnsetReportGroups
`func (o *BaseUser) UnsetReportGroups()`

UnsetReportGroups ensures that no value is present for ReportGroups, not even an explicit nil
### GetSecurityGroups

`func (o *BaseUser) GetSecurityGroups() map[string]interface{}`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *BaseUser) GetSecurityGroupsOk() (*map[string]interface{}, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *BaseUser) SetSecurityGroups(v map[string]interface{})`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *BaseUser) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *BaseUser) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *BaseUser) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetAuthorityName

`func (o *BaseUser) GetAuthorityName() string`

GetAuthorityName returns the AuthorityName field if non-nil, zero value otherwise.

### GetAuthorityNameOk

`func (o *BaseUser) GetAuthorityNameOk() (*string, bool)`

GetAuthorityNameOk returns a tuple with the AuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityName

`func (o *BaseUser) SetAuthorityName(v string)`

SetAuthorityName sets AuthorityName field to given value.

### HasAuthorityName

`func (o *BaseUser) HasAuthorityName() bool`

HasAuthorityName returns a boolean if a field has been set.

### SetAuthorityNameNil

`func (o *BaseUser) SetAuthorityNameNil(b bool)`

 SetAuthorityNameNil sets the value for AuthorityName to be an explicit nil

### UnsetAuthorityName
`func (o *BaseUser) UnsetAuthorityName()`

UnsetAuthorityName ensures that no value is present for AuthorityName, not even an explicit nil
### GetAuthorityAddress

`func (o *BaseUser) GetAuthorityAddress() string`

GetAuthorityAddress returns the AuthorityAddress field if non-nil, zero value otherwise.

### GetAuthorityAddressOk

`func (o *BaseUser) GetAuthorityAddressOk() (*string, bool)`

GetAuthorityAddressOk returns a tuple with the AuthorityAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityAddress

`func (o *BaseUser) SetAuthorityAddress(v string)`

SetAuthorityAddress sets AuthorityAddress field to given value.

### HasAuthorityAddress

`func (o *BaseUser) HasAuthorityAddress() bool`

HasAuthorityAddress returns a boolean if a field has been set.

### SetAuthorityAddressNil

`func (o *BaseUser) SetAuthorityAddressNil(b bool)`

 SetAuthorityAddressNil sets the value for AuthorityAddress to be an explicit nil

### UnsetAuthorityAddress
`func (o *BaseUser) UnsetAuthorityAddress()`

UnsetAuthorityAddress ensures that no value is present for AuthorityAddress, not even an explicit nil
### GetCompanyName

`func (o *BaseUser) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *BaseUser) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *BaseUser) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *BaseUser) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *BaseUser) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *BaseUser) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetCompanyAddress

`func (o *BaseUser) GetCompanyAddress() string`

GetCompanyAddress returns the CompanyAddress field if non-nil, zero value otherwise.

### GetCompanyAddressOk

`func (o *BaseUser) GetCompanyAddressOk() (*string, bool)`

GetCompanyAddressOk returns a tuple with the CompanyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress

`func (o *BaseUser) SetCompanyAddress(v string)`

SetCompanyAddress sets CompanyAddress field to given value.

### HasCompanyAddress

`func (o *BaseUser) HasCompanyAddress() bool`

HasCompanyAddress returns a boolean if a field has been set.

### SetCompanyAddressNil

`func (o *BaseUser) SetCompanyAddressNil(b bool)`

 SetCompanyAddressNil sets the value for CompanyAddress to be an explicit nil

### UnsetCompanyAddress
`func (o *BaseUser) UnsetCompanyAddress()`

UnsetCompanyAddress ensures that no value is present for CompanyAddress, not even an explicit nil
### GetCarrierNumber

`func (o *BaseUser) GetCarrierNumber() string`

GetCarrierNumber returns the CarrierNumber field if non-nil, zero value otherwise.

### GetCarrierNumberOk

`func (o *BaseUser) GetCarrierNumberOk() (*string, bool)`

GetCarrierNumberOk returns a tuple with the CarrierNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierNumber

`func (o *BaseUser) SetCarrierNumber(v string)`

SetCarrierNumber sets CarrierNumber field to given value.

### HasCarrierNumber

`func (o *BaseUser) HasCarrierNumber() bool`

HasCarrierNumber returns a boolean if a field has been set.

### SetCarrierNumberNil

`func (o *BaseUser) SetCarrierNumberNil(b bool)`

 SetCarrierNumberNil sets the value for CarrierNumber to be an explicit nil

### UnsetCarrierNumber
`func (o *BaseUser) UnsetCarrierNumber()`

UnsetCarrierNumber ensures that no value is present for CarrierNumber, not even an explicit nil
### GetLastTspLogin

`func (o *BaseUser) GetLastTspLogin() time.Time`

GetLastTspLogin returns the LastTspLogin field if non-nil, zero value otherwise.

### GetLastTspLoginOk

`func (o *BaseUser) GetLastTspLoginOk() (*time.Time, bool)`

GetLastTspLoginOk returns a tuple with the LastTspLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTspLogin

`func (o *BaseUser) SetLastTspLogin(v time.Time)`

SetLastTspLogin sets LastTspLogin field to given value.

### HasLastTspLogin

`func (o *BaseUser) HasLastTspLogin() bool`

HasLastTspLogin returns a boolean if a field has been set.

### SetLastTspLoginNil

`func (o *BaseUser) SetLastTspLoginNil(b bool)`

 SetLastTspLoginNil sets the value for LastTspLogin to be an explicit nil

### UnsetLastTspLogin
`func (o *BaseUser) UnsetLastTspLogin()`

UnsetLastTspLogin ensures that no value is present for LastTspLogin, not even an explicit nil
### GetNotes

`func (o *BaseUser) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BaseUser) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BaseUser) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *BaseUser) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *BaseUser) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *BaseUser) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetHosRulesetCode

`func (o *BaseUser) GetHosRulesetCode() HosRulesetCodeEnum`

GetHosRulesetCode returns the HosRulesetCode field if non-nil, zero value otherwise.

### GetHosRulesetCodeOk

`func (o *BaseUser) GetHosRulesetCodeOk() (*HosRulesetCodeEnum, bool)`

GetHosRulesetCodeOk returns a tuple with the HosRulesetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosRulesetCode

`func (o *BaseUser) SetHosRulesetCode(v HosRulesetCodeEnum)`

SetHosRulesetCode sets HosRulesetCode field to given value.

### HasHosRulesetCode

`func (o *BaseUser) HasHosRulesetCode() bool`

HasHosRulesetCode returns a boolean if a field has been set.

### SetHosRulesetCodeNil

`func (o *BaseUser) SetHosRulesetCodeNil(b bool)`

 SetHosRulesetCodeNil sets the value for HosRulesetCode to be an explicit nil

### UnsetHosRulesetCode
`func (o *BaseUser) UnsetHosRulesetCode()`

UnsetHosRulesetCode ensures that no value is present for HosRulesetCode, not even an explicit nil
### GetAllowYardMove

`func (o *BaseUser) GetAllowYardMove() bool`

GetAllowYardMove returns the AllowYardMove field if non-nil, zero value otherwise.

### GetAllowYardMoveOk

`func (o *BaseUser) GetAllowYardMoveOk() (*bool, bool)`

GetAllowYardMoveOk returns a tuple with the AllowYardMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowYardMove

`func (o *BaseUser) SetAllowYardMove(v bool)`

SetAllowYardMove sets AllowYardMove field to given value.

### HasAllowYardMove

`func (o *BaseUser) HasAllowYardMove() bool`

HasAllowYardMove returns a boolean if a field has been set.

### SetAllowYardMoveNil

`func (o *BaseUser) SetAllowYardMoveNil(b bool)`

 SetAllowYardMoveNil sets the value for AllowYardMove to be an explicit nil

### UnsetAllowYardMove
`func (o *BaseUser) UnsetAllowYardMove()`

UnsetAllowYardMove ensures that no value is present for AllowYardMove, not even an explicit nil
### GetAllowPersonalConveyance

`func (o *BaseUser) GetAllowPersonalConveyance() bool`

GetAllowPersonalConveyance returns the AllowPersonalConveyance field if non-nil, zero value otherwise.

### GetAllowPersonalConveyanceOk

`func (o *BaseUser) GetAllowPersonalConveyanceOk() (*bool, bool)`

GetAllowPersonalConveyanceOk returns a tuple with the AllowPersonalConveyance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPersonalConveyance

`func (o *BaseUser) SetAllowPersonalConveyance(v bool)`

SetAllowPersonalConveyance sets AllowPersonalConveyance field to given value.

### HasAllowPersonalConveyance

`func (o *BaseUser) HasAllowPersonalConveyance() bool`

HasAllowPersonalConveyance returns a boolean if a field has been set.

### SetAllowPersonalConveyanceNil

`func (o *BaseUser) SetAllowPersonalConveyanceNil(b bool)`

 SetAllowPersonalConveyanceNil sets the value for AllowPersonalConveyance to be an explicit nil

### UnsetAllowPersonalConveyance
`func (o *BaseUser) UnsetAllowPersonalConveyance()`

UnsetAllowPersonalConveyance ensures that no value is present for AllowPersonalConveyance, not even an explicit nil
### GetAllowAdverseDriving

`func (o *BaseUser) GetAllowAdverseDriving() bool`

GetAllowAdverseDriving returns the AllowAdverseDriving field if non-nil, zero value otherwise.

### GetAllowAdverseDrivingOk

`func (o *BaseUser) GetAllowAdverseDrivingOk() (*bool, bool)`

GetAllowAdverseDrivingOk returns a tuple with the AllowAdverseDriving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAdverseDriving

`func (o *BaseUser) SetAllowAdverseDriving(v bool)`

SetAllowAdverseDriving sets AllowAdverseDriving field to given value.

### HasAllowAdverseDriving

`func (o *BaseUser) HasAllowAdverseDriving() bool`

HasAllowAdverseDriving returns a boolean if a field has been set.

### SetAllowAdverseDrivingNil

`func (o *BaseUser) SetAllowAdverseDrivingNil(b bool)`

 SetAllowAdverseDrivingNil sets the value for AllowAdverseDriving to be an explicit nil

### UnsetAllowAdverseDriving
`func (o *BaseUser) UnsetAllowAdverseDriving()`

UnsetAllowAdverseDriving ensures that no value is present for AllowAdverseDriving, not even an explicit nil
### GetDefaultTimeZone

`func (o *BaseUser) GetDefaultTimeZone() TimezoneCodeEnum`

GetDefaultTimeZone returns the DefaultTimeZone field if non-nil, zero value otherwise.

### GetDefaultTimeZoneOk

`func (o *BaseUser) GetDefaultTimeZoneOk() (*TimezoneCodeEnum, bool)`

GetDefaultTimeZoneOk returns a tuple with the DefaultTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeZone

`func (o *BaseUser) SetDefaultTimeZone(v TimezoneCodeEnum)`

SetDefaultTimeZone sets DefaultTimeZone field to given value.

### HasDefaultTimeZone

`func (o *BaseUser) HasDefaultTimeZone() bool`

HasDefaultTimeZone returns a boolean if a field has been set.

### SetDefaultTimeZoneNil

`func (o *BaseUser) SetDefaultTimeZoneNil(b bool)`

 SetDefaultTimeZoneNil sets the value for DefaultTimeZone to be an explicit nil

### UnsetDefaultTimeZone
`func (o *BaseUser) UnsetDefaultTimeZone()`

UnsetDefaultTimeZone ensures that no value is present for DefaultTimeZone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


