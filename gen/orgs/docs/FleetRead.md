# FleetRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**LegalName** | Pointer to **NullableString** |  | [optional] 
**DbaName** | Pointer to **NullableString** |  | [optional] 
**Websites** | Pointer to **[]string** |  | [optional] 
**RegulatoryId** | Pointer to **NullableString** |  | [optional] 
**RegulatoryIdType** | Pointer to **NullableString** |  | [optional] 
**RegulatoryIdDate** | Pointer to **NullableString** |  | [optional] 
**RegulatoryIdStatus** | Pointer to **NullableString** |  | [optional] 
**RegisteredEmail** | Pointer to **NullableString** |  | [optional] 
**RegisteredPhone** | Pointer to **NullableString** |  | [optional] 
**RegisteredFax** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to **NullableString** |  | [optional] 
**City** | Pointer to **NullableString** |  | [optional] 
**Province** | Pointer to **NullableString** |  | [optional] 
**PostalCode** | Pointer to **NullableString** |  | [optional] 
**CountryCode** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** | Unique Catena fleet identifier | 
**CreatedAt** | **time.Time** | When the fleet was created | 
**UpdatedAt** | **time.Time** | Last modification timestamp | 

## Methods

### NewFleetRead

`func NewFleetRead(name string, id string, createdAt time.Time, updatedAt time.Time, ) *FleetRead`

NewFleetRead instantiates a new FleetRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetReadWithDefaults

`func NewFleetReadWithDefaults() *FleetRead`

NewFleetReadWithDefaults instantiates a new FleetRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FleetRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FleetRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FleetRead) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FleetRead) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FleetRead) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FleetRead) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FleetRead) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FleetRead) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FleetRead) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *FleetRead) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FleetRead) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FleetRead) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FleetRead) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *FleetRead) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *FleetRead) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLegalName

`func (o *FleetRead) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *FleetRead) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *FleetRead) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *FleetRead) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### SetLegalNameNil

`func (o *FleetRead) SetLegalNameNil(b bool)`

 SetLegalNameNil sets the value for LegalName to be an explicit nil

### UnsetLegalName
`func (o *FleetRead) UnsetLegalName()`

UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
### GetDbaName

`func (o *FleetRead) GetDbaName() string`

GetDbaName returns the DbaName field if non-nil, zero value otherwise.

### GetDbaNameOk

`func (o *FleetRead) GetDbaNameOk() (*string, bool)`

GetDbaNameOk returns a tuple with the DbaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaName

`func (o *FleetRead) SetDbaName(v string)`

SetDbaName sets DbaName field to given value.

### HasDbaName

`func (o *FleetRead) HasDbaName() bool`

HasDbaName returns a boolean if a field has been set.

### SetDbaNameNil

`func (o *FleetRead) SetDbaNameNil(b bool)`

 SetDbaNameNil sets the value for DbaName to be an explicit nil

### UnsetDbaName
`func (o *FleetRead) UnsetDbaName()`

UnsetDbaName ensures that no value is present for DbaName, not even an explicit nil
### GetWebsites

`func (o *FleetRead) GetWebsites() []string`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *FleetRead) GetWebsitesOk() (*[]string, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsites

`func (o *FleetRead) SetWebsites(v []string)`

SetWebsites sets Websites field to given value.

### HasWebsites

`func (o *FleetRead) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.

### SetWebsitesNil

`func (o *FleetRead) SetWebsitesNil(b bool)`

 SetWebsitesNil sets the value for Websites to be an explicit nil

### UnsetWebsites
`func (o *FleetRead) UnsetWebsites()`

UnsetWebsites ensures that no value is present for Websites, not even an explicit nil
### GetRegulatoryId

`func (o *FleetRead) GetRegulatoryId() string`

GetRegulatoryId returns the RegulatoryId field if non-nil, zero value otherwise.

### GetRegulatoryIdOk

`func (o *FleetRead) GetRegulatoryIdOk() (*string, bool)`

GetRegulatoryIdOk returns a tuple with the RegulatoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryId

`func (o *FleetRead) SetRegulatoryId(v string)`

SetRegulatoryId sets RegulatoryId field to given value.

### HasRegulatoryId

`func (o *FleetRead) HasRegulatoryId() bool`

HasRegulatoryId returns a boolean if a field has been set.

### SetRegulatoryIdNil

`func (o *FleetRead) SetRegulatoryIdNil(b bool)`

 SetRegulatoryIdNil sets the value for RegulatoryId to be an explicit nil

### UnsetRegulatoryId
`func (o *FleetRead) UnsetRegulatoryId()`

UnsetRegulatoryId ensures that no value is present for RegulatoryId, not even an explicit nil
### GetRegulatoryIdType

`func (o *FleetRead) GetRegulatoryIdType() string`

GetRegulatoryIdType returns the RegulatoryIdType field if non-nil, zero value otherwise.

### GetRegulatoryIdTypeOk

`func (o *FleetRead) GetRegulatoryIdTypeOk() (*string, bool)`

GetRegulatoryIdTypeOk returns a tuple with the RegulatoryIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryIdType

`func (o *FleetRead) SetRegulatoryIdType(v string)`

SetRegulatoryIdType sets RegulatoryIdType field to given value.

### HasRegulatoryIdType

`func (o *FleetRead) HasRegulatoryIdType() bool`

HasRegulatoryIdType returns a boolean if a field has been set.

### SetRegulatoryIdTypeNil

`func (o *FleetRead) SetRegulatoryIdTypeNil(b bool)`

 SetRegulatoryIdTypeNil sets the value for RegulatoryIdType to be an explicit nil

### UnsetRegulatoryIdType
`func (o *FleetRead) UnsetRegulatoryIdType()`

UnsetRegulatoryIdType ensures that no value is present for RegulatoryIdType, not even an explicit nil
### GetRegulatoryIdDate

`func (o *FleetRead) GetRegulatoryIdDate() string`

GetRegulatoryIdDate returns the RegulatoryIdDate field if non-nil, zero value otherwise.

### GetRegulatoryIdDateOk

`func (o *FleetRead) GetRegulatoryIdDateOk() (*string, bool)`

GetRegulatoryIdDateOk returns a tuple with the RegulatoryIdDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryIdDate

`func (o *FleetRead) SetRegulatoryIdDate(v string)`

SetRegulatoryIdDate sets RegulatoryIdDate field to given value.

### HasRegulatoryIdDate

`func (o *FleetRead) HasRegulatoryIdDate() bool`

HasRegulatoryIdDate returns a boolean if a field has been set.

### SetRegulatoryIdDateNil

`func (o *FleetRead) SetRegulatoryIdDateNil(b bool)`

 SetRegulatoryIdDateNil sets the value for RegulatoryIdDate to be an explicit nil

### UnsetRegulatoryIdDate
`func (o *FleetRead) UnsetRegulatoryIdDate()`

UnsetRegulatoryIdDate ensures that no value is present for RegulatoryIdDate, not even an explicit nil
### GetRegulatoryIdStatus

`func (o *FleetRead) GetRegulatoryIdStatus() string`

GetRegulatoryIdStatus returns the RegulatoryIdStatus field if non-nil, zero value otherwise.

### GetRegulatoryIdStatusOk

`func (o *FleetRead) GetRegulatoryIdStatusOk() (*string, bool)`

GetRegulatoryIdStatusOk returns a tuple with the RegulatoryIdStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryIdStatus

`func (o *FleetRead) SetRegulatoryIdStatus(v string)`

SetRegulatoryIdStatus sets RegulatoryIdStatus field to given value.

### HasRegulatoryIdStatus

`func (o *FleetRead) HasRegulatoryIdStatus() bool`

HasRegulatoryIdStatus returns a boolean if a field has been set.

### SetRegulatoryIdStatusNil

`func (o *FleetRead) SetRegulatoryIdStatusNil(b bool)`

 SetRegulatoryIdStatusNil sets the value for RegulatoryIdStatus to be an explicit nil

### UnsetRegulatoryIdStatus
`func (o *FleetRead) UnsetRegulatoryIdStatus()`

UnsetRegulatoryIdStatus ensures that no value is present for RegulatoryIdStatus, not even an explicit nil
### GetRegisteredEmail

`func (o *FleetRead) GetRegisteredEmail() string`

GetRegisteredEmail returns the RegisteredEmail field if non-nil, zero value otherwise.

### GetRegisteredEmailOk

`func (o *FleetRead) GetRegisteredEmailOk() (*string, bool)`

GetRegisteredEmailOk returns a tuple with the RegisteredEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEmail

`func (o *FleetRead) SetRegisteredEmail(v string)`

SetRegisteredEmail sets RegisteredEmail field to given value.

### HasRegisteredEmail

`func (o *FleetRead) HasRegisteredEmail() bool`

HasRegisteredEmail returns a boolean if a field has been set.

### SetRegisteredEmailNil

`func (o *FleetRead) SetRegisteredEmailNil(b bool)`

 SetRegisteredEmailNil sets the value for RegisteredEmail to be an explicit nil

### UnsetRegisteredEmail
`func (o *FleetRead) UnsetRegisteredEmail()`

UnsetRegisteredEmail ensures that no value is present for RegisteredEmail, not even an explicit nil
### GetRegisteredPhone

`func (o *FleetRead) GetRegisteredPhone() string`

GetRegisteredPhone returns the RegisteredPhone field if non-nil, zero value otherwise.

### GetRegisteredPhoneOk

`func (o *FleetRead) GetRegisteredPhoneOk() (*string, bool)`

GetRegisteredPhoneOk returns a tuple with the RegisteredPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredPhone

`func (o *FleetRead) SetRegisteredPhone(v string)`

SetRegisteredPhone sets RegisteredPhone field to given value.

### HasRegisteredPhone

`func (o *FleetRead) HasRegisteredPhone() bool`

HasRegisteredPhone returns a boolean if a field has been set.

### SetRegisteredPhoneNil

`func (o *FleetRead) SetRegisteredPhoneNil(b bool)`

 SetRegisteredPhoneNil sets the value for RegisteredPhone to be an explicit nil

### UnsetRegisteredPhone
`func (o *FleetRead) UnsetRegisteredPhone()`

UnsetRegisteredPhone ensures that no value is present for RegisteredPhone, not even an explicit nil
### GetRegisteredFax

`func (o *FleetRead) GetRegisteredFax() string`

GetRegisteredFax returns the RegisteredFax field if non-nil, zero value otherwise.

### GetRegisteredFaxOk

`func (o *FleetRead) GetRegisteredFaxOk() (*string, bool)`

GetRegisteredFaxOk returns a tuple with the RegisteredFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredFax

`func (o *FleetRead) SetRegisteredFax(v string)`

SetRegisteredFax sets RegisteredFax field to given value.

### HasRegisteredFax

`func (o *FleetRead) HasRegisteredFax() bool`

HasRegisteredFax returns a boolean if a field has been set.

### SetRegisteredFaxNil

`func (o *FleetRead) SetRegisteredFaxNil(b bool)`

 SetRegisteredFaxNil sets the value for RegisteredFax to be an explicit nil

### UnsetRegisteredFax
`func (o *FleetRead) UnsetRegisteredFax()`

UnsetRegisteredFax ensures that no value is present for RegisteredFax, not even an explicit nil
### GetAddress

`func (o *FleetRead) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FleetRead) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FleetRead) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FleetRead) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *FleetRead) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *FleetRead) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCity

`func (o *FleetRead) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *FleetRead) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *FleetRead) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *FleetRead) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *FleetRead) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *FleetRead) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetProvince

`func (o *FleetRead) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *FleetRead) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *FleetRead) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *FleetRead) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### SetProvinceNil

`func (o *FleetRead) SetProvinceNil(b bool)`

 SetProvinceNil sets the value for Province to be an explicit nil

### UnsetProvince
`func (o *FleetRead) UnsetProvince()`

UnsetProvince ensures that no value is present for Province, not even an explicit nil
### GetPostalCode

`func (o *FleetRead) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *FleetRead) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *FleetRead) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *FleetRead) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *FleetRead) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *FleetRead) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryCode

`func (o *FleetRead) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *FleetRead) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *FleetRead) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *FleetRead) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *FleetRead) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *FleetRead) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetId

`func (o *FleetRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *FleetRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FleetRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FleetRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *FleetRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FleetRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FleetRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


