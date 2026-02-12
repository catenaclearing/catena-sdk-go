# UserRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FleetId** | **NullableString** |  | 
**FleetRef** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | Unique identifier of the record at Catena Telematics. | 
**CreatedAt** | **time.Time** | Immutable: The datetime the record was ingested into Catena Telematics. | 
**UpdatedAt** | **time.Time** | The dateime the record was last modified in Catena Telematics. | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
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

### NewUserRead

`func NewUserRead(fleetId NullableString, id string, createdAt time.Time, updatedAt time.Time, connectionId string, sourceName TspEnum, sourceId string, sourceDataHash string, ) *UserRead`

NewUserRead instantiates a new UserRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserReadWithDefaults

`func NewUserReadWithDefaults() *UserRead`

NewUserReadWithDefaults instantiates a new UserRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFleetId

`func (o *UserRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *UserRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *UserRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### SetFleetIdNil

`func (o *UserRead) SetFleetIdNil(b bool)`

 SetFleetIdNil sets the value for FleetId to be an explicit nil

### UnsetFleetId
`func (o *UserRead) UnsetFleetId()`

UnsetFleetId ensures that no value is present for FleetId, not even an explicit nil
### GetFleetRef

`func (o *UserRead) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *UserRead) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *UserRead) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *UserRead) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *UserRead) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *UserRead) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetId

`func (o *UserRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *UserRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDeletedAt

`func (o *UserRead) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *UserRead) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *UserRead) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *UserRead) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *UserRead) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *UserRead) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetConnectionId

`func (o *UserRead) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *UserRead) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *UserRead) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetSourceName

`func (o *UserRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *UserRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *UserRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetSourceData

`func (o *UserRead) GetSourceData() map[string]interface{}`

GetSourceData returns the SourceData field if non-nil, zero value otherwise.

### GetSourceDataOk

`func (o *UserRead) GetSourceDataOk() (*map[string]interface{}, bool)`

GetSourceDataOk returns a tuple with the SourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceData

`func (o *UserRead) SetSourceData(v map[string]interface{})`

SetSourceData sets SourceData field to given value.

### HasSourceData

`func (o *UserRead) HasSourceData() bool`

HasSourceData returns a boolean if a field has been set.

### GetSourceId

`func (o *UserRead) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *UserRead) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *UserRead) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetSourceDataHash

`func (o *UserRead) GetSourceDataHash() string`

GetSourceDataHash returns the SourceDataHash field if non-nil, zero value otherwise.

### GetSourceDataHashOk

`func (o *UserRead) GetSourceDataHashOk() (*string, bool)`

GetSourceDataHashOk returns a tuple with the SourceDataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDataHash

`func (o *UserRead) SetSourceDataHash(v string)`

SetSourceDataHash sets SourceDataHash field to given value.


### GetOccurredAt

`func (o *UserRead) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *UserRead) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *UserRead) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *UserRead) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *UserRead) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *UserRead) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetExecutionId

`func (o *UserRead) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *UserRead) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *UserRead) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *UserRead) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.

### SetExecutionIdNil

`func (o *UserRead) SetExecutionIdNil(b bool)`

 SetExecutionIdNil sets the value for ExecutionId to be an explicit nil

### UnsetExecutionId
`func (o *UserRead) UnsetExecutionId()`

UnsetExecutionId ensures that no value is present for ExecutionId, not even an explicit nil
### GetScheduleId

`func (o *UserRead) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *UserRead) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *UserRead) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *UserRead) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### SetScheduleIdNil

`func (o *UserRead) SetScheduleIdNil(b bool)`

 SetScheduleIdNil sets the value for ScheduleId to be an explicit nil

### UnsetScheduleId
`func (o *UserRead) UnsetScheduleId()`

UnsetScheduleId ensures that no value is present for ScheduleId, not even an explicit nil
### GetUsername

`func (o *UserRead) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserRead) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserRead) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserRead) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *UserRead) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *UserRead) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetStartedAt

`func (o *UserRead) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *UserRead) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *UserRead) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *UserRead) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *UserRead) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *UserRead) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetEndedAt

`func (o *UserRead) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *UserRead) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *UserRead) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *UserRead) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *UserRead) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *UserRead) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetIsActive

`func (o *UserRead) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UserRead) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UserRead) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UserRead) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *UserRead) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *UserRead) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetStatus

`func (o *UserRead) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserRead) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserRead) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserRead) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *UserRead) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *UserRead) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetIsDriver

`func (o *UserRead) GetIsDriver() bool`

GetIsDriver returns the IsDriver field if non-nil, zero value otherwise.

### GetIsDriverOk

`func (o *UserRead) GetIsDriverOk() (*bool, bool)`

GetIsDriverOk returns a tuple with the IsDriver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDriver

`func (o *UserRead) SetIsDriver(v bool)`

SetIsDriver sets IsDriver field to given value.

### HasIsDriver

`func (o *UserRead) HasIsDriver() bool`

HasIsDriver returns a boolean if a field has been set.

### SetIsDriverNil

`func (o *UserRead) SetIsDriverNil(b bool)`

 SetIsDriverNil sets the value for IsDriver to be an explicit nil

### UnsetIsDriver
`func (o *UserRead) UnsetIsDriver()`

UnsetIsDriver ensures that no value is present for IsDriver, not even an explicit nil
### GetUserDesignation

`func (o *UserRead) GetUserDesignation() string`

GetUserDesignation returns the UserDesignation field if non-nil, zero value otherwise.

### GetUserDesignationOk

`func (o *UserRead) GetUserDesignationOk() (*string, bool)`

GetUserDesignationOk returns a tuple with the UserDesignation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDesignation

`func (o *UserRead) SetUserDesignation(v string)`

SetUserDesignation sets UserDesignation field to given value.

### HasUserDesignation

`func (o *UserRead) HasUserDesignation() bool`

HasUserDesignation returns a boolean if a field has been set.

### SetUserDesignationNil

`func (o *UserRead) SetUserDesignationNil(b bool)`

 SetUserDesignationNil sets the value for UserDesignation to be an explicit nil

### UnsetUserDesignation
`func (o *UserRead) UnsetUserDesignation()`

UnsetUserDesignation ensures that no value is present for UserDesignation, not even an explicit nil
### GetUserEmail

`func (o *UserRead) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *UserRead) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *UserRead) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.

### HasUserEmail

`func (o *UserRead) HasUserEmail() bool`

HasUserEmail returns a boolean if a field has been set.

### SetUserEmailNil

`func (o *UserRead) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *UserRead) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
### GetFirstName

`func (o *UserRead) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserRead) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserRead) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserRead) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *UserRead) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UserRead) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UserRead) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserRead) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserRead) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserRead) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UserRead) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UserRead) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetPhoneNumber

`func (o *UserRead) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UserRead) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UserRead) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UserRead) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *UserRead) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *UserRead) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetCountryCode

`func (o *UserRead) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *UserRead) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *UserRead) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *UserRead) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *UserRead) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *UserRead) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetLicenseCountry

`func (o *UserRead) GetLicenseCountry() string`

GetLicenseCountry returns the LicenseCountry field if non-nil, zero value otherwise.

### GetLicenseCountryOk

`func (o *UserRead) GetLicenseCountryOk() (*string, bool)`

GetLicenseCountryOk returns a tuple with the LicenseCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseCountry

`func (o *UserRead) SetLicenseCountry(v string)`

SetLicenseCountry sets LicenseCountry field to given value.

### HasLicenseCountry

`func (o *UserRead) HasLicenseCountry() bool`

HasLicenseCountry returns a boolean if a field has been set.

### SetLicenseCountryNil

`func (o *UserRead) SetLicenseCountryNil(b bool)`

 SetLicenseCountryNil sets the value for LicenseCountry to be an explicit nil

### UnsetLicenseCountry
`func (o *UserRead) UnsetLicenseCountry()`

UnsetLicenseCountry ensures that no value is present for LicenseCountry, not even an explicit nil
### GetLicenseRegion

`func (o *UserRead) GetLicenseRegion() string`

GetLicenseRegion returns the LicenseRegion field if non-nil, zero value otherwise.

### GetLicenseRegionOk

`func (o *UserRead) GetLicenseRegionOk() (*string, bool)`

GetLicenseRegionOk returns a tuple with the LicenseRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseRegion

`func (o *UserRead) SetLicenseRegion(v string)`

SetLicenseRegion sets LicenseRegion field to given value.

### HasLicenseRegion

`func (o *UserRead) HasLicenseRegion() bool`

HasLicenseRegion returns a boolean if a field has been set.

### SetLicenseRegionNil

`func (o *UserRead) SetLicenseRegionNil(b bool)`

 SetLicenseRegionNil sets the value for LicenseRegion to be an explicit nil

### UnsetLicenseRegion
`func (o *UserRead) UnsetLicenseRegion()`

UnsetLicenseRegion ensures that no value is present for LicenseRegion, not even an explicit nil
### GetLicenseNumber

`func (o *UserRead) GetLicenseNumber() string`

GetLicenseNumber returns the LicenseNumber field if non-nil, zero value otherwise.

### GetLicenseNumberOk

`func (o *UserRead) GetLicenseNumberOk() (*string, bool)`

GetLicenseNumberOk returns a tuple with the LicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNumber

`func (o *UserRead) SetLicenseNumber(v string)`

SetLicenseNumber sets LicenseNumber field to given value.

### HasLicenseNumber

`func (o *UserRead) HasLicenseNumber() bool`

HasLicenseNumber returns a boolean if a field has been set.

### SetLicenseNumberNil

`func (o *UserRead) SetLicenseNumberNil(b bool)`

 SetLicenseNumberNil sets the value for LicenseNumber to be an explicit nil

### UnsetLicenseNumber
`func (o *UserRead) UnsetLicenseNumber()`

UnsetLicenseNumber ensures that no value is present for LicenseNumber, not even an explicit nil
### GetLicenseExpiration

`func (o *UserRead) GetLicenseExpiration() string`

GetLicenseExpiration returns the LicenseExpiration field if non-nil, zero value otherwise.

### GetLicenseExpirationOk

`func (o *UserRead) GetLicenseExpirationOk() (*string, bool)`

GetLicenseExpirationOk returns a tuple with the LicenseExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseExpiration

`func (o *UserRead) SetLicenseExpiration(v string)`

SetLicenseExpiration sets LicenseExpiration field to given value.

### HasLicenseExpiration

`func (o *UserRead) HasLicenseExpiration() bool`

HasLicenseExpiration returns a boolean if a field has been set.

### SetLicenseExpirationNil

`func (o *UserRead) SetLicenseExpirationNil(b bool)`

 SetLicenseExpirationNil sets the value for LicenseExpiration to be an explicit nil

### UnsetLicenseExpiration
`func (o *UserRead) UnsetLicenseExpiration()`

UnsetLicenseExpiration ensures that no value is present for LicenseExpiration, not even an explicit nil
### GetEmployeeNumber

`func (o *UserRead) GetEmployeeNumber() string`

GetEmployeeNumber returns the EmployeeNumber field if non-nil, zero value otherwise.

### GetEmployeeNumberOk

`func (o *UserRead) GetEmployeeNumberOk() (*string, bool)`

GetEmployeeNumberOk returns a tuple with the EmployeeNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeNumber

`func (o *UserRead) SetEmployeeNumber(v string)`

SetEmployeeNumber sets EmployeeNumber field to given value.

### HasEmployeeNumber

`func (o *UserRead) HasEmployeeNumber() bool`

HasEmployeeNumber returns a boolean if a field has been set.

### SetEmployeeNumberNil

`func (o *UserRead) SetEmployeeNumberNil(b bool)`

 SetEmployeeNumberNil sets the value for EmployeeNumber to be an explicit nil

### UnsetEmployeeNumber
`func (o *UserRead) UnsetEmployeeNumber()`

UnsetEmployeeNumber ensures that no value is present for EmployeeNumber, not even an explicit nil
### GetCompanyGroups

`func (o *UserRead) GetCompanyGroups() map[string]interface{}`

GetCompanyGroups returns the CompanyGroups field if non-nil, zero value otherwise.

### GetCompanyGroupsOk

`func (o *UserRead) GetCompanyGroupsOk() (*map[string]interface{}, bool)`

GetCompanyGroupsOk returns a tuple with the CompanyGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyGroups

`func (o *UserRead) SetCompanyGroups(v map[string]interface{})`

SetCompanyGroups sets CompanyGroups field to given value.

### HasCompanyGroups

`func (o *UserRead) HasCompanyGroups() bool`

HasCompanyGroups returns a boolean if a field has been set.

### SetCompanyGroupsNil

`func (o *UserRead) SetCompanyGroupsNil(b bool)`

 SetCompanyGroupsNil sets the value for CompanyGroups to be an explicit nil

### UnsetCompanyGroups
`func (o *UserRead) UnsetCompanyGroups()`

UnsetCompanyGroups ensures that no value is present for CompanyGroups, not even an explicit nil
### GetPrivateUserGroups

`func (o *UserRead) GetPrivateUserGroups() map[string]interface{}`

GetPrivateUserGroups returns the PrivateUserGroups field if non-nil, zero value otherwise.

### GetPrivateUserGroupsOk

`func (o *UserRead) GetPrivateUserGroupsOk() (*map[string]interface{}, bool)`

GetPrivateUserGroupsOk returns a tuple with the PrivateUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateUserGroups

`func (o *UserRead) SetPrivateUserGroups(v map[string]interface{})`

SetPrivateUserGroups sets PrivateUserGroups field to given value.

### HasPrivateUserGroups

`func (o *UserRead) HasPrivateUserGroups() bool`

HasPrivateUserGroups returns a boolean if a field has been set.

### SetPrivateUserGroupsNil

`func (o *UserRead) SetPrivateUserGroupsNil(b bool)`

 SetPrivateUserGroupsNil sets the value for PrivateUserGroups to be an explicit nil

### UnsetPrivateUserGroups
`func (o *UserRead) UnsetPrivateUserGroups()`

UnsetPrivateUserGroups ensures that no value is present for PrivateUserGroups, not even an explicit nil
### GetReportGroups

`func (o *UserRead) GetReportGroups() map[string]interface{}`

GetReportGroups returns the ReportGroups field if non-nil, zero value otherwise.

### GetReportGroupsOk

`func (o *UserRead) GetReportGroupsOk() (*map[string]interface{}, bool)`

GetReportGroupsOk returns a tuple with the ReportGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportGroups

`func (o *UserRead) SetReportGroups(v map[string]interface{})`

SetReportGroups sets ReportGroups field to given value.

### HasReportGroups

`func (o *UserRead) HasReportGroups() bool`

HasReportGroups returns a boolean if a field has been set.

### SetReportGroupsNil

`func (o *UserRead) SetReportGroupsNil(b bool)`

 SetReportGroupsNil sets the value for ReportGroups to be an explicit nil

### UnsetReportGroups
`func (o *UserRead) UnsetReportGroups()`

UnsetReportGroups ensures that no value is present for ReportGroups, not even an explicit nil
### GetSecurityGroups

`func (o *UserRead) GetSecurityGroups() map[string]interface{}`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *UserRead) GetSecurityGroupsOk() (*map[string]interface{}, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *UserRead) SetSecurityGroups(v map[string]interface{})`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *UserRead) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *UserRead) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *UserRead) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetAuthorityName

`func (o *UserRead) GetAuthorityName() string`

GetAuthorityName returns the AuthorityName field if non-nil, zero value otherwise.

### GetAuthorityNameOk

`func (o *UserRead) GetAuthorityNameOk() (*string, bool)`

GetAuthorityNameOk returns a tuple with the AuthorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityName

`func (o *UserRead) SetAuthorityName(v string)`

SetAuthorityName sets AuthorityName field to given value.

### HasAuthorityName

`func (o *UserRead) HasAuthorityName() bool`

HasAuthorityName returns a boolean if a field has been set.

### SetAuthorityNameNil

`func (o *UserRead) SetAuthorityNameNil(b bool)`

 SetAuthorityNameNil sets the value for AuthorityName to be an explicit nil

### UnsetAuthorityName
`func (o *UserRead) UnsetAuthorityName()`

UnsetAuthorityName ensures that no value is present for AuthorityName, not even an explicit nil
### GetAuthorityAddress

`func (o *UserRead) GetAuthorityAddress() string`

GetAuthorityAddress returns the AuthorityAddress field if non-nil, zero value otherwise.

### GetAuthorityAddressOk

`func (o *UserRead) GetAuthorityAddressOk() (*string, bool)`

GetAuthorityAddressOk returns a tuple with the AuthorityAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityAddress

`func (o *UserRead) SetAuthorityAddress(v string)`

SetAuthorityAddress sets AuthorityAddress field to given value.

### HasAuthorityAddress

`func (o *UserRead) HasAuthorityAddress() bool`

HasAuthorityAddress returns a boolean if a field has been set.

### SetAuthorityAddressNil

`func (o *UserRead) SetAuthorityAddressNil(b bool)`

 SetAuthorityAddressNil sets the value for AuthorityAddress to be an explicit nil

### UnsetAuthorityAddress
`func (o *UserRead) UnsetAuthorityAddress()`

UnsetAuthorityAddress ensures that no value is present for AuthorityAddress, not even an explicit nil
### GetCompanyName

`func (o *UserRead) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UserRead) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UserRead) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UserRead) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### SetCompanyNameNil

`func (o *UserRead) SetCompanyNameNil(b bool)`

 SetCompanyNameNil sets the value for CompanyName to be an explicit nil

### UnsetCompanyName
`func (o *UserRead) UnsetCompanyName()`

UnsetCompanyName ensures that no value is present for CompanyName, not even an explicit nil
### GetCompanyAddress

`func (o *UserRead) GetCompanyAddress() string`

GetCompanyAddress returns the CompanyAddress field if non-nil, zero value otherwise.

### GetCompanyAddressOk

`func (o *UserRead) GetCompanyAddressOk() (*string, bool)`

GetCompanyAddressOk returns a tuple with the CompanyAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyAddress

`func (o *UserRead) SetCompanyAddress(v string)`

SetCompanyAddress sets CompanyAddress field to given value.

### HasCompanyAddress

`func (o *UserRead) HasCompanyAddress() bool`

HasCompanyAddress returns a boolean if a field has been set.

### SetCompanyAddressNil

`func (o *UserRead) SetCompanyAddressNil(b bool)`

 SetCompanyAddressNil sets the value for CompanyAddress to be an explicit nil

### UnsetCompanyAddress
`func (o *UserRead) UnsetCompanyAddress()`

UnsetCompanyAddress ensures that no value is present for CompanyAddress, not even an explicit nil
### GetCarrierNumber

`func (o *UserRead) GetCarrierNumber() string`

GetCarrierNumber returns the CarrierNumber field if non-nil, zero value otherwise.

### GetCarrierNumberOk

`func (o *UserRead) GetCarrierNumberOk() (*string, bool)`

GetCarrierNumberOk returns a tuple with the CarrierNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierNumber

`func (o *UserRead) SetCarrierNumber(v string)`

SetCarrierNumber sets CarrierNumber field to given value.

### HasCarrierNumber

`func (o *UserRead) HasCarrierNumber() bool`

HasCarrierNumber returns a boolean if a field has been set.

### SetCarrierNumberNil

`func (o *UserRead) SetCarrierNumberNil(b bool)`

 SetCarrierNumberNil sets the value for CarrierNumber to be an explicit nil

### UnsetCarrierNumber
`func (o *UserRead) UnsetCarrierNumber()`

UnsetCarrierNumber ensures that no value is present for CarrierNumber, not even an explicit nil
### GetLastTspLogin

`func (o *UserRead) GetLastTspLogin() time.Time`

GetLastTspLogin returns the LastTspLogin field if non-nil, zero value otherwise.

### GetLastTspLoginOk

`func (o *UserRead) GetLastTspLoginOk() (*time.Time, bool)`

GetLastTspLoginOk returns a tuple with the LastTspLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTspLogin

`func (o *UserRead) SetLastTspLogin(v time.Time)`

SetLastTspLogin sets LastTspLogin field to given value.

### HasLastTspLogin

`func (o *UserRead) HasLastTspLogin() bool`

HasLastTspLogin returns a boolean if a field has been set.

### SetLastTspLoginNil

`func (o *UserRead) SetLastTspLoginNil(b bool)`

 SetLastTspLoginNil sets the value for LastTspLogin to be an explicit nil

### UnsetLastTspLogin
`func (o *UserRead) UnsetLastTspLogin()`

UnsetLastTspLogin ensures that no value is present for LastTspLogin, not even an explicit nil
### GetNotes

`func (o *UserRead) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UserRead) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UserRead) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UserRead) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *UserRead) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *UserRead) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetHosRulesetCode

`func (o *UserRead) GetHosRulesetCode() HosRulesetCodeEnum`

GetHosRulesetCode returns the HosRulesetCode field if non-nil, zero value otherwise.

### GetHosRulesetCodeOk

`func (o *UserRead) GetHosRulesetCodeOk() (*HosRulesetCodeEnum, bool)`

GetHosRulesetCodeOk returns a tuple with the HosRulesetCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosRulesetCode

`func (o *UserRead) SetHosRulesetCode(v HosRulesetCodeEnum)`

SetHosRulesetCode sets HosRulesetCode field to given value.

### HasHosRulesetCode

`func (o *UserRead) HasHosRulesetCode() bool`

HasHosRulesetCode returns a boolean if a field has been set.

### SetHosRulesetCodeNil

`func (o *UserRead) SetHosRulesetCodeNil(b bool)`

 SetHosRulesetCodeNil sets the value for HosRulesetCode to be an explicit nil

### UnsetHosRulesetCode
`func (o *UserRead) UnsetHosRulesetCode()`

UnsetHosRulesetCode ensures that no value is present for HosRulesetCode, not even an explicit nil
### GetAllowYardMove

`func (o *UserRead) GetAllowYardMove() bool`

GetAllowYardMove returns the AllowYardMove field if non-nil, zero value otherwise.

### GetAllowYardMoveOk

`func (o *UserRead) GetAllowYardMoveOk() (*bool, bool)`

GetAllowYardMoveOk returns a tuple with the AllowYardMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowYardMove

`func (o *UserRead) SetAllowYardMove(v bool)`

SetAllowYardMove sets AllowYardMove field to given value.

### HasAllowYardMove

`func (o *UserRead) HasAllowYardMove() bool`

HasAllowYardMove returns a boolean if a field has been set.

### SetAllowYardMoveNil

`func (o *UserRead) SetAllowYardMoveNil(b bool)`

 SetAllowYardMoveNil sets the value for AllowYardMove to be an explicit nil

### UnsetAllowYardMove
`func (o *UserRead) UnsetAllowYardMove()`

UnsetAllowYardMove ensures that no value is present for AllowYardMove, not even an explicit nil
### GetAllowPersonalConveyance

`func (o *UserRead) GetAllowPersonalConveyance() bool`

GetAllowPersonalConveyance returns the AllowPersonalConveyance field if non-nil, zero value otherwise.

### GetAllowPersonalConveyanceOk

`func (o *UserRead) GetAllowPersonalConveyanceOk() (*bool, bool)`

GetAllowPersonalConveyanceOk returns a tuple with the AllowPersonalConveyance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPersonalConveyance

`func (o *UserRead) SetAllowPersonalConveyance(v bool)`

SetAllowPersonalConveyance sets AllowPersonalConveyance field to given value.

### HasAllowPersonalConveyance

`func (o *UserRead) HasAllowPersonalConveyance() bool`

HasAllowPersonalConveyance returns a boolean if a field has been set.

### SetAllowPersonalConveyanceNil

`func (o *UserRead) SetAllowPersonalConveyanceNil(b bool)`

 SetAllowPersonalConveyanceNil sets the value for AllowPersonalConveyance to be an explicit nil

### UnsetAllowPersonalConveyance
`func (o *UserRead) UnsetAllowPersonalConveyance()`

UnsetAllowPersonalConveyance ensures that no value is present for AllowPersonalConveyance, not even an explicit nil
### GetAllowAdverseDriving

`func (o *UserRead) GetAllowAdverseDriving() bool`

GetAllowAdverseDriving returns the AllowAdverseDriving field if non-nil, zero value otherwise.

### GetAllowAdverseDrivingOk

`func (o *UserRead) GetAllowAdverseDrivingOk() (*bool, bool)`

GetAllowAdverseDrivingOk returns a tuple with the AllowAdverseDriving field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAdverseDriving

`func (o *UserRead) SetAllowAdverseDriving(v bool)`

SetAllowAdverseDriving sets AllowAdverseDriving field to given value.

### HasAllowAdverseDriving

`func (o *UserRead) HasAllowAdverseDriving() bool`

HasAllowAdverseDriving returns a boolean if a field has been set.

### SetAllowAdverseDrivingNil

`func (o *UserRead) SetAllowAdverseDrivingNil(b bool)`

 SetAllowAdverseDrivingNil sets the value for AllowAdverseDriving to be an explicit nil

### UnsetAllowAdverseDriving
`func (o *UserRead) UnsetAllowAdverseDriving()`

UnsetAllowAdverseDriving ensures that no value is present for AllowAdverseDriving, not even an explicit nil
### GetDefaultTimeZone

`func (o *UserRead) GetDefaultTimeZone() TimezoneCodeEnum`

GetDefaultTimeZone returns the DefaultTimeZone field if non-nil, zero value otherwise.

### GetDefaultTimeZoneOk

`func (o *UserRead) GetDefaultTimeZoneOk() (*TimezoneCodeEnum, bool)`

GetDefaultTimeZoneOk returns a tuple with the DefaultTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeZone

`func (o *UserRead) SetDefaultTimeZone(v TimezoneCodeEnum)`

SetDefaultTimeZone sets DefaultTimeZone field to given value.

### HasDefaultTimeZone

`func (o *UserRead) HasDefaultTimeZone() bool`

HasDefaultTimeZone returns a boolean if a field has been set.

### SetDefaultTimeZoneNil

`func (o *UserRead) SetDefaultTimeZoneNil(b bool)`

 SetDefaultTimeZoneNil sets the value for DefaultTimeZone to be an explicit nil

### UnsetDefaultTimeZone
`func (o *UserRead) UnsetDefaultTimeZone()`

UnsetDefaultTimeZone ensures that no value is present for DefaultTimeZone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


