# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the record at Catena Telematics. | 
**CreatedAt** | **time.Time** | Immutable: The datetime the record was ingested into Catena Telematics. | 
**UpdatedAt** | **time.Time** | The dateime the record was last modified in Catena Telematics. | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**FleetId** | **string** | Unique identifier of the fleet at Catena Telematics. This record belongs to this fleet. **Note: this is not the fleet ID in the TSP system or in your organization&#39;s systems, use &#x60;share_agreements&#x60; to map &#x60;fleet_ids&#x60; to &#x60;fleet_refs&#x60;.** | 
**ConnectionId** | **string** | Unique identifier of the connection at Catena Telematics through which this record was ingested. A connection represents a Fleet/TSP pairing. | 
**SourceName** | [**TspEnum**](TspEnum.md) | An enumeration identifying the TSP from which this record was sourced. | 
**SourceData** | Pointer to **map[string]interface{}** | Raw source payload as ingested from the TSP. **Note: use it for audit/debugging.** | [optional] 
**SourceId** | **string** | Unique identifier of the record in the TSP. **Note: we generate a unique composite key based on available fields if the TSP does not provide an unique ID.** | 
**SourceDataHash** | **string** | SHA-256 hash of the source data payload. **Note: we use it internally for idempotence and deduplication.** | 
**OccurredAt** | Pointer to **NullableTime** |  | [optional] 
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

### NewUser

`func NewUser(id string, createdAt time.Time, updatedAt time.Time, fleetId string, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *User) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *User) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *User) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *User) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *User) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *User) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *User) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetFleetId

`func (o *User) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *User) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *User) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetConnectionId

`func (o *User) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *User) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *User) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceName

`func (o *User) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *User) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *User) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *User) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *User) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *User) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *User) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *User) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *User) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *User) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *User) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *User) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *User) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *User) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *User) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *User) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *User) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *User) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *User) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *User) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *User) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *User) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *User) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *User) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *User) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *User) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *User) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *User) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *User) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *User) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *User) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *User) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *User) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetStartedAt

`func (o *User) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *User) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *User) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *User) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *User) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *User) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *User) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *User) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *User) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *User) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *User) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *User) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetIsActive

`func (o *User) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *User) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *User) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *User) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *User) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *User) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetStatus

`func (o *User) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *User) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *User) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *User) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *User) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *User) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIsDriver

`func (o *User) GetIsDriver() bool`

GetIsDriver returns the IsDriver field if non-nil, zero value otherwise.

### GetIsDriverOk

`func (o *User) GetIsDriverOk() (*bool, bool)`

GetIsDriverOk returns a tuple with the IsDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDriver

`func (o *User) SetIsDriver(v bool)`

SetIsDriver sets IsDriver field to given value.

### HasIsDriver

`func (o *User) HasIsDriver() bool`

HasIsDriver returns a boolean if a field has been set.

### SetIsDriverNil

`func (o *User) SetIsDriverNil(b bool)`

 SetIsDriverNil sets the value for IsDriver to be an explicit nil

### UnsetIsDriver
`func (o *User) UnsetIsDriver()`

UnsetIsDriver ensures that no value is present for IsDriver, not even an explicit nil
### GetUserDesignation

`func (o *User) GetUserDesignation() string`

GetUserDesignation returns the UserDesignation field if non-nil, zero value otherwise.

### GetUserDesignationOk

`func (o *User) GetUserDesignationOk() (*string, bool)`

GetUserDesignationOk returns a tuple with the UserDesignation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDesignation

`func (o *User) SetUserDesignation(v string)`

SetUserDesignation sets UserDesignation field to given value.

### HasUserDesignation

`func (o *User) HasUserDesignation() bool`

HasUserDesignation returns a boolean if a field has been set.

### SetUserDesignationNil

`func (o *User) SetUserDesignationNil(b bool)`

 SetUserDesignationNil sets the value for UserDesignation to be an explicit nil

### UnsetUserDesignation
`func (o *User) UnsetUserDesignation()`

UnsetUserDesignation ensures that no value is present for UserDesignation, not even an explicit nil
### GetUserEmail

`func (o *User) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *User) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *User) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *User) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmailNil

`func (o *User) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *User) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
### GetFirstName

`func (o *User) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *User) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *User) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *User) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *User) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *User) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *User) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *User) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *User) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *User) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *User) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *User) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetPhoneNumber

`func (o *User) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *User) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *User) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *User) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *User) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *User) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetCountryCode

`func (o *User) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *User) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *User) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *User) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *User) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *User) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetLicenseCountry

`func (o *User) GetLicenseCountry() string`

GetLicenseCountry returns the LicenseCountry field if non-nil, zero value otherwise.

### GetLicenseCountryOk

`func (o *User) GetLicenseCountryOk() (*string, bool)`

GetLicenseCountryOk returns a tuple with the LicenseCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCountry

`func (o *User) SetLicenseCountry(v string)`

SetLicenseCountry sets LicenseCountry field to given value.

### HasLicenseCountry

`func (o *User) HasLicenseCountry() bool`

HasLicenseCountry returns a boolean if a field has been set.

### SetLicenseCountryNil

`func (o *User) SetLicenseCountryNil(b bool)`

 SetLicenseCountryNil sets the value for LicenseCountry to be an explicit nil

### UnsetLicenseCountry
`func (o *User) UnsetLicenseCountry()`

UnsetLicenseCountry ensures that no value is present for LicenseCountry, not even an explicit nil
### GetLicenseRegion

`func (o *User) GetLicenseRegion() string`

GetLicenseRegion returns the LicenseRegion field if non-nil, zero value otherwise.

### GetLicenseRegionOk

`func (o *User) GetLicenseRegionOk() (*string, bool)`

GetLicenseRegionOk returns a tuple with the LicenseRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseRegion

`func (o *User) SetLicenseRegion(v string)`

SetLicenseRegion sets LicenseRegion field to given value.

### HasLicenseRegion

`func (o *User) HasLicenseRegion() bool`

HasLicenseRegion returns a boolean if a field has been set.

### SetLicenseRegionNil

`func (o *User) SetLicenseRegionNil(b bool)`

 SetLicenseRegionNil sets the value for LicenseRegion to be an explicit nil

### UnsetLicenseRegion
`func (o *User) UnsetLicenseRegion()`

UnsetLicenseRegion ensures that no value is present for LicenseRegion, not even an explicit nil
### GetLicenseNumber

`func (o *User) GetLicenseNumber() string`

GetLicenseNumber returns the LicenseNumber field if non-nil, zero value otherwise.

### GetLicenseNumberOk

`func (o *User) GetLicenseNumberOk() (*string, bool)`

GetLicenseNumberOk returns a tuple with the LicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNumber

`func (o *User) SetLicenseNumber(v string)`

SetLicenseNumber sets LicenseNumber field to given value.

### HasLicenseNumber

`func (o *User) HasLicenseNumber() bool`

HasLicenseNumber returns a boolean if a field has been set.

### SetLicenseNumberNil

`func (o *User) SetLicenseNumberNil(b bool)`

 SetLicenseNumberNil sets the value for LicenseNumber to be an explicit nil

### UnsetLicenseNumber
`func (o *User) UnsetLicenseNumber()`

UnsetLicenseNumber ensures that no value is present for LicenseNumber, not even an explicit nil
### GetLicenseExpiration

`func (o *User) GetLicenseExpiration() string`

GetLicenseExpiration returns the LicenseExpiration field if non-nil, zero value otherwise.

### GetLicenseExpirationOk

`func (o *User) GetLicenseExpirationOk() (*string, bool)`

GetLicenseExpirationOk returns a tuple with the LicenseExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpiration

`func (o *User) SetLicenseExpiration(v string)`

SetLicenseExpiration sets LicenseExpiration field to given value.

### HasLicenseExpiration

`func (o *User) HasLicenseExpiration() bool`

HasLicenseExpiration returns a boolean if a field has been set.

### SetLicenseExpirationNil

`func (o *User) SetLicenseExpirationNil(b bool)`

 SetLicenseExpirationNil sets the value for LicenseExpiration to be an explicit nil

### UnsetLicenseExpiration
`func (o *User) UnsetLicenseExpiration()`

UnsetLicenseExpiration ensures that no value is present for LicenseExpiration, not even an explicit nil
### GetEmployeeNumber

`func (o *User) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *User) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *User) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *User) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### SetEmployeeNumberNil

`func (o *User) SetEmployeeNumberNil(b bool)`

 SetEmployeeNumberNil sets the value for EmployeeNumber to be an explicit nil

### UnsetEmployeeNumber
`func (o *User) UnsetEmployeeNumber()`

UnsetEmployeeNumber ensures that no value is present for EmployeeNumber, not even an explicit nil
### GetCompanyGroups

`func (o *User) GetCompanyGroups() map[string]interface{}`

GetCompanyGroups returns the CompanyGroups field if non-nil, zero value otherwise.

### GetCompanyGroupsOk

`func (o *User) GetCompanyGroupsOk() (*map[string]interface{}, bool)`

GetCompanyGroupsOk returns a tuple with the CompanyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyGroups

`func (o *User) SetCompanyGroups(v map[string]interface{})`

SetCompanyGroups sets CompanyGroups field to given value.

### HasCompanyGroups

`func (o *User) HasCompanyGroups() bool`

HasCompanyGroups returns a boolean if a field has been set.

### SetCompanyGroupsNil

`func (o *User) SetCompanyGroupsNil(b bool)`

 SetCompanyGroupsNil sets the value for CompanyGroups to be an explicit nil

### UnsetCompanyGroups
`func (o *User) UnsetCompanyGroups()`

UnsetCompanyGroups ensures that no value is present for CompanyGroups, not even an explicit nil
### GetPrivateUserGroups

`func (o *User) GetPrivateUserGroups() map[string]interface{}`

GetPrivateUserGroups returns the PrivateUserGroups field if non-nil, zero value otherwise.

### GetPrivateUserGroupsOk

`func (o *User) GetPrivateUserGroupsOk() (*map[string]interface{}, bool)`

GetPrivateUserGroupsOk returns a tuple with the PrivateUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateUserGroups

`func (o *User) SetPrivateUserGroups(v map[string]interface{})`

SetPrivateUserGroups sets PrivateUserGroups field to given value.

### HasPrivateUserGroups

`func (o *User) HasPrivateUserGroups() bool`

HasPrivateUserGroups returns a boolean if a field has been set.

### SetPrivateUserGroupsNil

`func (o *User) SetPrivateUserGroupsNil(b bool)`

 SetPrivateUserGroupsNil sets the value for PrivateUserGroups to be an explicit nil

### UnsetPrivateUserGroups
`func (o *User) UnsetPrivateUserGroups()`

UnsetPrivateUserGroups ensures that no value is present for PrivateUserGroups, not even an explicit nil
### GetReportGroups

`func (o *User) GetReportGroups() map[string]interface{}`

GetReportGroups returns the ReportGroups field if non-nil, zero value otherwise.

### GetReportGroupsOk

`func (o *User) GetReportGroupsOk() (*map[string]interface{}, bool)`

GetReportGroupsOk returns a tuple with the ReportGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportGroups

`func (o *User) SetReportGroups(v map[string]interface{})`

SetReportGroups sets ReportGroups field to given value.

### HasReportGroups

`func (o *User) HasReportGroups() bool`

HasReportGroups returns a boolean if a field has been set.

### SetReportGroupsNil

`func (o *User) SetReportGroupsNil(b bool)`

 SetReportGroupsNil sets the value for ReportGroups to be an explicit nil

### UnsetReportGroups
`func (o *User) UnsetReportGroups()`

UnsetReportGroups ensures that no value is present for ReportGroups, not even an explicit nil
### GetSecurityGroups

`func (o *User) GetSecurityGroups() map[string]interface{}`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *User) GetSecurityGroupsOk() (*map[string]interface{}, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *User) SetSecurityGroups(v map[string]interface{})`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *User) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *User) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *User) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetAuthorityName

`func (o *User) GetAuthorityName() string`

GetAuthorityName returns the AuthorityName field if non-nil, zero value otherwise.

### GetAuthorityNameOk

`func (o *User) GetAuthorityNameOk() (*string, bool)`

GetAuthorityNameOk returns a tuple with the AuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityName

`func (o *User) SetAuthorityName(v string)`

SetAuthorityName sets AuthorityName field to given value.

### HasAuthorityName

`func (o *User) HasAuthorityName() bool`

HasAuthorityName returns a boolean if a field has been set.

### SetAuthorityNameNil

`func (o *User) SetAuthorityNameNil(b bool)`

 SetAuthorityNameNil sets the value for AuthorityName to be an explicit nil

### UnsetAuthorityName
`func (o *User) UnsetAuthorityName()`

UnsetAuthorityName ensures that no value is present for AuthorityName, not even an explicit nil
### GetAuthorityAddress

`func (o *User) GetAuthorityAddress() string`

GetAuthorityAddress returns the AuthorityAddress field if non-nil, zero value otherwise.

### GetAuthorityAddressOk

`func (o *User) GetAuthorityAddressOk() (*string, bool)`

GetAuthorityAddressOk returns a tuple with the AuthorityAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityAddress

`func (o *User) SetAuthorityAddress(v string)`

SetAuthorityAddress sets AuthorityAddress field to given value.

### HasAuthorityAddress

`func (o *User) HasAuthorityAddress() bool`

HasAuthorityAddress returns a boolean if a field has been set.

### SetAuthorityAddressNil

`func (o *User) SetAuthorityAddressNil(b bool)`

 SetAuthorityAddressNil sets the value for AuthorityAddress to be an explicit nil

### UnsetAuthorityAddress
`func (o *User) UnsetAuthorityAddress()`

UnsetAuthorityAddress ensures that no value is present for AuthorityAddress, not even an explicit nil
### GetCompanyName

`func (o *User) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *User) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *User) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *User) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *User) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *User) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetCompanyAddress

`func (o *User) GetCompanyAddress() string`

GetCompanyAddress returns the CompanyAddress field if non-nil, zero value otherwise.

### GetCompanyAddressOk

`func (o *User) GetCompanyAddressOk() (*string, bool)`

GetCompanyAddressOk returns a tuple with the CompanyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress

`func (o *User) SetCompanyAddress(v string)`

SetCompanyAddress sets CompanyAddress field to given value.

### HasCompanyAddress

`func (o *User) HasCompanyAddress() bool`

HasCompanyAddress returns a boolean if a field has been set.

### SetCompanyAddressNil

`func (o *User) SetCompanyAddressNil(b bool)`

 SetCompanyAddressNil sets the value for CompanyAddress to be an explicit nil

### UnsetCompanyAddress
`func (o *User) UnsetCompanyAddress()`

UnsetCompanyAddress ensures that no value is present for CompanyAddress, not even an explicit nil
### GetCarrierNumber

`func (o *User) GetCarrierNumber() string`

GetCarrierNumber returns the CarrierNumber field if non-nil, zero value otherwise.

### GetCarrierNumberOk

`func (o *User) GetCarrierNumberOk() (*string, bool)`

GetCarrierNumberOk returns a tuple with the CarrierNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierNumber

`func (o *User) SetCarrierNumber(v string)`

SetCarrierNumber sets CarrierNumber field to given value.

### HasCarrierNumber

`func (o *User) HasCarrierNumber() bool`

HasCarrierNumber returns a boolean if a field has been set.

### SetCarrierNumberNil

`func (o *User) SetCarrierNumberNil(b bool)`

 SetCarrierNumberNil sets the value for CarrierNumber to be an explicit nil

### UnsetCarrierNumber
`func (o *User) UnsetCarrierNumber()`

UnsetCarrierNumber ensures that no value is present for CarrierNumber, not even an explicit nil
### GetLastTspLogin

`func (o *User) GetLastTspLogin() time.Time`

GetLastTspLogin returns the LastTspLogin field if non-nil, zero value otherwise.

### GetLastTspLoginOk

`func (o *User) GetLastTspLoginOk() (*time.Time, bool)`

GetLastTspLoginOk returns a tuple with the LastTspLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTspLogin

`func (o *User) SetLastTspLogin(v time.Time)`

SetLastTspLogin sets LastTspLogin field to given value.

### HasLastTspLogin

`func (o *User) HasLastTspLogin() bool`

HasLastTspLogin returns a boolean if a field has been set.

### SetLastTspLoginNil

`func (o *User) SetLastTspLoginNil(b bool)`

 SetLastTspLoginNil sets the value for LastTspLogin to be an explicit nil

### UnsetLastTspLogin
`func (o *User) UnsetLastTspLogin()`

UnsetLastTspLogin ensures that no value is present for LastTspLogin, not even an explicit nil
### GetNotes

`func (o *User) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *User) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *User) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *User) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *User) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *User) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetHosRulesetCode

`func (o *User) GetHosRulesetCode() HosRulesetCodeEnum`

GetHosRulesetCode returns the HosRulesetCode field if non-nil, zero value otherwise.

### GetHosRulesetCodeOk

`func (o *User) GetHosRulesetCodeOk() (*HosRulesetCodeEnum, bool)`

GetHosRulesetCodeOk returns a tuple with the HosRulesetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosRulesetCode

`func (o *User) SetHosRulesetCode(v HosRulesetCodeEnum)`

SetHosRulesetCode sets HosRulesetCode field to given value.

### HasHosRulesetCode

`func (o *User) HasHosRulesetCode() bool`

HasHosRulesetCode returns a boolean if a field has been set.

### SetHosRulesetCodeNil

`func (o *User) SetHosRulesetCodeNil(b bool)`

 SetHosRulesetCodeNil sets the value for HosRulesetCode to be an explicit nil

### UnsetHosRulesetCode
`func (o *User) UnsetHosRulesetCode()`

UnsetHosRulesetCode ensures that no value is present for HosRulesetCode, not even an explicit nil
### GetAllowYardMove

`func (o *User) GetAllowYardMove() bool`

GetAllowYardMove returns the AllowYardMove field if non-nil, zero value otherwise.

### GetAllowYardMoveOk

`func (o *User) GetAllowYardMoveOk() (*bool, bool)`

GetAllowYardMoveOk returns a tuple with the AllowYardMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowYardMove

`func (o *User) SetAllowYardMove(v bool)`

SetAllowYardMove sets AllowYardMove field to given value.

### HasAllowYardMove

`func (o *User) HasAllowYardMove() bool`

HasAllowYardMove returns a boolean if a field has been set.

### SetAllowYardMoveNil

`func (o *User) SetAllowYardMoveNil(b bool)`

 SetAllowYardMoveNil sets the value for AllowYardMove to be an explicit nil

### UnsetAllowYardMove
`func (o *User) UnsetAllowYardMove()`

UnsetAllowYardMove ensures that no value is present for AllowYardMove, not even an explicit nil
### GetAllowPersonalConveyance

`func (o *User) GetAllowPersonalConveyance() bool`

GetAllowPersonalConveyance returns the AllowPersonalConveyance field if non-nil, zero value otherwise.

### GetAllowPersonalConveyanceOk

`func (o *User) GetAllowPersonalConveyanceOk() (*bool, bool)`

GetAllowPersonalConveyanceOk returns a tuple with the AllowPersonalConveyance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPersonalConveyance

`func (o *User) SetAllowPersonalConveyance(v bool)`

SetAllowPersonalConveyance sets AllowPersonalConveyance field to given value.

### HasAllowPersonalConveyance

`func (o *User) HasAllowPersonalConveyance() bool`

HasAllowPersonalConveyance returns a boolean if a field has been set.

### SetAllowPersonalConveyanceNil

`func (o *User) SetAllowPersonalConveyanceNil(b bool)`

 SetAllowPersonalConveyanceNil sets the value for AllowPersonalConveyance to be an explicit nil

### UnsetAllowPersonalConveyance
`func (o *User) UnsetAllowPersonalConveyance()`

UnsetAllowPersonalConveyance ensures that no value is present for AllowPersonalConveyance, not even an explicit nil
### GetAllowAdverseDriving

`func (o *User) GetAllowAdverseDriving() bool`

GetAllowAdverseDriving returns the AllowAdverseDriving field if non-nil, zero value otherwise.

### GetAllowAdverseDrivingOk

`func (o *User) GetAllowAdverseDrivingOk() (*bool, bool)`

GetAllowAdverseDrivingOk returns a tuple with the AllowAdverseDriving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAdverseDriving

`func (o *User) SetAllowAdverseDriving(v bool)`

SetAllowAdverseDriving sets AllowAdverseDriving field to given value.

### HasAllowAdverseDriving

`func (o *User) HasAllowAdverseDriving() bool`

HasAllowAdverseDriving returns a boolean if a field has been set.

### SetAllowAdverseDrivingNil

`func (o *User) SetAllowAdverseDrivingNil(b bool)`

 SetAllowAdverseDrivingNil sets the value for AllowAdverseDriving to be an explicit nil

### UnsetAllowAdverseDriving
`func (o *User) UnsetAllowAdverseDriving()`

UnsetAllowAdverseDriving ensures that no value is present for AllowAdverseDriving, not even an explicit nil
### GetDefaultTimeZone

`func (o *User) GetDefaultTimeZone() TimezoneCodeEnum`

GetDefaultTimeZone returns the DefaultTimeZone field if non-nil, zero value otherwise.

### GetDefaultTimeZoneOk

`func (o *User) GetDefaultTimeZoneOk() (*TimezoneCodeEnum, bool)`

GetDefaultTimeZoneOk returns a tuple with the DefaultTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeZone

`func (o *User) SetDefaultTimeZone(v TimezoneCodeEnum)`

SetDefaultTimeZone sets DefaultTimeZone field to given value.

### HasDefaultTimeZone

`func (o *User) HasDefaultTimeZone() bool`

HasDefaultTimeZone returns a boolean if a field has been set.

### SetDefaultTimeZoneNil

`func (o *User) SetDefaultTimeZoneNil(b bool)`

 SetDefaultTimeZoneNil sets the value for DefaultTimeZone to be an explicit nil

### UnsetDefaultTimeZone
`func (o *User) UnsetDefaultTimeZone()`

UnsetDefaultTimeZone ensures that no value is present for DefaultTimeZone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


