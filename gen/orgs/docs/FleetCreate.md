# FleetCreate

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
**Properties** | Pointer to [**[]FleetPropertyCreate**](FleetPropertyCreate.md) | Optional custom properties for the fleet | [optional] [default to []]

## Methods

### NewFleetCreate

`func NewFleetCreate(name string, ) *FleetCreate`

NewFleetCreate instantiates a new FleetCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetCreateWithDefaults

`func NewFleetCreateWithDefaults() *FleetCreate`

NewFleetCreateWithDefaults instantiates a new FleetCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FleetCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FleetCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FleetCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FleetCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FleetCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FleetCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FleetCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *FleetCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *FleetCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *FleetCreate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FleetCreate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FleetCreate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FleetCreate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *FleetCreate) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *FleetCreate) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetLegalName

`func (o *FleetCreate) GetLegalName() string`

GetLegalName returns the LegalName field if non-nil, zero value otherwise.

### GetLegalNameOk

`func (o *FleetCreate) GetLegalNameOk() (*string, bool)`

GetLegalNameOk returns a tuple with the LegalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalName

`func (o *FleetCreate) SetLegalName(v string)`

SetLegalName sets LegalName field to given value.

### HasLegalName

`func (o *FleetCreate) HasLegalName() bool`

HasLegalName returns a boolean if a field has been set.

### SetLegalNameNil

`func (o *FleetCreate) SetLegalNameNil(b bool)`

 SetLegalNameNil sets the value for LegalName to be an explicit nil

### UnsetLegalName
`func (o *FleetCreate) UnsetLegalName()`

UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
### GetDbaName

`func (o *FleetCreate) GetDbaName() string`

GetDbaName returns the DbaName field if non-nil, zero value otherwise.

### GetDbaNameOk

`func (o *FleetCreate) GetDbaNameOk() (*string, bool)`

GetDbaNameOk returns a tuple with the DbaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaName

`func (o *FleetCreate) SetDbaName(v string)`

SetDbaName sets DbaName field to given value.

### HasDbaName

`func (o *FleetCreate) HasDbaName() bool`

HasDbaName returns a boolean if a field has been set.

### SetDbaNameNil

`func (o *FleetCreate) SetDbaNameNil(b bool)`

 SetDbaNameNil sets the value for DbaName to be an explicit nil

### UnsetDbaName
`func (o *FleetCreate) UnsetDbaName()`

UnsetDbaName ensures that no value is present for DbaName, not even an explicit nil
### GetWebsites

`func (o *FleetCreate) GetWebsites() []string`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *FleetCreate) GetWebsitesOk() (*[]string, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsites

`func (o *FleetCreate) SetWebsites(v []string)`

SetWebsites sets Websites field to given value.

### HasWebsites

`func (o *FleetCreate) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.

### SetWebsitesNil

`func (o *FleetCreate) SetWebsitesNil(b bool)`

 SetWebsitesNil sets the value for Websites to be an explicit nil

### UnsetWebsites
`func (o *FleetCreate) UnsetWebsites()`

UnsetWebsites ensures that no value is present for Websites, not even an explicit nil
### GetRegulatoryId

`func (o *FleetCreate) GetRegulatoryId() string`

GetRegulatoryId returns the RegulatoryId field if non-nil, zero value otherwise.

### GetRegulatoryIdOk

`func (o *FleetCreate) GetRegulatoryIdOk() (*string, bool)`

GetRegulatoryIdOk returns a tuple with the RegulatoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryId

`func (o *FleetCreate) SetRegulatoryId(v string)`

SetRegulatoryId sets RegulatoryId field to given value.

### HasRegulatoryId

`func (o *FleetCreate) HasRegulatoryId() bool`

HasRegulatoryId returns a boolean if a field has been set.

### SetRegulatoryIdNil

`func (o *FleetCreate) SetRegulatoryIdNil(b bool)`

 SetRegulatoryIdNil sets the value for RegulatoryId to be an explicit nil

### UnsetRegulatoryId
`func (o *FleetCreate) UnsetRegulatoryId()`

UnsetRegulatoryId ensures that no value is present for RegulatoryId, not even an explicit nil
### GetRegulatoryIdType

`func (o *FleetCreate) GetRegulatoryIdType() string`

GetRegulatoryIdType returns the RegulatoryIdType field if non-nil, zero value otherwise.

### GetRegulatoryIdTypeOk

`func (o *FleetCreate) GetRegulatoryIdTypeOk() (*string, bool)`

GetRegulatoryIdTypeOk returns a tuple with the RegulatoryIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryIdType

`func (o *FleetCreate) SetRegulatoryIdType(v string)`

SetRegulatoryIdType sets RegulatoryIdType field to given value.

### HasRegulatoryIdType

`func (o *FleetCreate) HasRegulatoryIdType() bool`

HasRegulatoryIdType returns a boolean if a field has been set.

### SetRegulatoryIdTypeNil

`func (o *FleetCreate) SetRegulatoryIdTypeNil(b bool)`

 SetRegulatoryIdTypeNil sets the value for RegulatoryIdType to be an explicit nil

### UnsetRegulatoryIdType
`func (o *FleetCreate) UnsetRegulatoryIdType()`

UnsetRegulatoryIdType ensures that no value is present for RegulatoryIdType, not even an explicit nil
### GetRegulatoryIdDate

`func (o *FleetCreate) GetRegulatoryIdDate() string`

GetRegulatoryIdDate returns the RegulatoryIdDate field if non-nil, zero value otherwise.

### GetRegulatoryIdDateOk

`func (o *FleetCreate) GetRegulatoryIdDateOk() (*string, bool)`

GetRegulatoryIdDateOk returns a tuple with the RegulatoryIdDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryIdDate

`func (o *FleetCreate) SetRegulatoryIdDate(v string)`

SetRegulatoryIdDate sets RegulatoryIdDate field to given value.

### HasRegulatoryIdDate

`func (o *FleetCreate) HasRegulatoryIdDate() bool`

HasRegulatoryIdDate returns a boolean if a field has been set.

### SetRegulatoryIdDateNil

`func (o *FleetCreate) SetRegulatoryIdDateNil(b bool)`

 SetRegulatoryIdDateNil sets the value for RegulatoryIdDate to be an explicit nil

### UnsetRegulatoryIdDate
`func (o *FleetCreate) UnsetRegulatoryIdDate()`

UnsetRegulatoryIdDate ensures that no value is present for RegulatoryIdDate, not even an explicit nil
### GetRegulatoryIdStatus

`func (o *FleetCreate) GetRegulatoryIdStatus() string`

GetRegulatoryIdStatus returns the RegulatoryIdStatus field if non-nil, zero value otherwise.

### GetRegulatoryIdStatusOk

`func (o *FleetCreate) GetRegulatoryIdStatusOk() (*string, bool)`

GetRegulatoryIdStatusOk returns a tuple with the RegulatoryIdStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryIdStatus

`func (o *FleetCreate) SetRegulatoryIdStatus(v string)`

SetRegulatoryIdStatus sets RegulatoryIdStatus field to given value.

### HasRegulatoryIdStatus

`func (o *FleetCreate) HasRegulatoryIdStatus() bool`

HasRegulatoryIdStatus returns a boolean if a field has been set.

### SetRegulatoryIdStatusNil

`func (o *FleetCreate) SetRegulatoryIdStatusNil(b bool)`

 SetRegulatoryIdStatusNil sets the value for RegulatoryIdStatus to be an explicit nil

### UnsetRegulatoryIdStatus
`func (o *FleetCreate) UnsetRegulatoryIdStatus()`

UnsetRegulatoryIdStatus ensures that no value is present for RegulatoryIdStatus, not even an explicit nil
### GetRegisteredEmail

`func (o *FleetCreate) GetRegisteredEmail() string`

GetRegisteredEmail returns the RegisteredEmail field if non-nil, zero value otherwise.

### GetRegisteredEmailOk

`func (o *FleetCreate) GetRegisteredEmailOk() (*string, bool)`

GetRegisteredEmailOk returns a tuple with the RegisteredEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEmail

`func (o *FleetCreate) SetRegisteredEmail(v string)`

SetRegisteredEmail sets RegisteredEmail field to given value.

### HasRegisteredEmail

`func (o *FleetCreate) HasRegisteredEmail() bool`

HasRegisteredEmail returns a boolean if a field has been set.

### SetRegisteredEmailNil

`func (o *FleetCreate) SetRegisteredEmailNil(b bool)`

 SetRegisteredEmailNil sets the value for RegisteredEmail to be an explicit nil

### UnsetRegisteredEmail
`func (o *FleetCreate) UnsetRegisteredEmail()`

UnsetRegisteredEmail ensures that no value is present for RegisteredEmail, not even an explicit nil
### GetRegisteredPhone

`func (o *FleetCreate) GetRegisteredPhone() string`

GetRegisteredPhone returns the RegisteredPhone field if non-nil, zero value otherwise.

### GetRegisteredPhoneOk

`func (o *FleetCreate) GetRegisteredPhoneOk() (*string, bool)`

GetRegisteredPhoneOk returns a tuple with the RegisteredPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredPhone

`func (o *FleetCreate) SetRegisteredPhone(v string)`

SetRegisteredPhone sets RegisteredPhone field to given value.

### HasRegisteredPhone

`func (o *FleetCreate) HasRegisteredPhone() bool`

HasRegisteredPhone returns a boolean if a field has been set.

### SetRegisteredPhoneNil

`func (o *FleetCreate) SetRegisteredPhoneNil(b bool)`

 SetRegisteredPhoneNil sets the value for RegisteredPhone to be an explicit nil

### UnsetRegisteredPhone
`func (o *FleetCreate) UnsetRegisteredPhone()`

UnsetRegisteredPhone ensures that no value is present for RegisteredPhone, not even an explicit nil
### GetRegisteredFax

`func (o *FleetCreate) GetRegisteredFax() string`

GetRegisteredFax returns the RegisteredFax field if non-nil, zero value otherwise.

### GetRegisteredFaxOk

`func (o *FleetCreate) GetRegisteredFaxOk() (*string, bool)`

GetRegisteredFaxOk returns a tuple with the RegisteredFax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredFax

`func (o *FleetCreate) SetRegisteredFax(v string)`

SetRegisteredFax sets RegisteredFax field to given value.

### HasRegisteredFax

`func (o *FleetCreate) HasRegisteredFax() bool`

HasRegisteredFax returns a boolean if a field has been set.

### SetRegisteredFaxNil

`func (o *FleetCreate) SetRegisteredFaxNil(b bool)`

 SetRegisteredFaxNil sets the value for RegisteredFax to be an explicit nil

### UnsetRegisteredFax
`func (o *FleetCreate) UnsetRegisteredFax()`

UnsetRegisteredFax ensures that no value is present for RegisteredFax, not even an explicit nil
### GetAddress

`func (o *FleetCreate) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FleetCreate) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FleetCreate) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FleetCreate) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *FleetCreate) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *FleetCreate) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCity

`func (o *FleetCreate) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *FleetCreate) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *FleetCreate) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *FleetCreate) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *FleetCreate) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *FleetCreate) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetProvince

`func (o *FleetCreate) GetProvince() string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *FleetCreate) GetProvinceOk() (*string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *FleetCreate) SetProvince(v string)`

SetProvince sets Province field to given value.

### HasProvince

`func (o *FleetCreate) HasProvince() bool`

HasProvince returns a boolean if a field has been set.

### SetProvinceNil

`func (o *FleetCreate) SetProvinceNil(b bool)`

 SetProvinceNil sets the value for Province to be an explicit nil

### UnsetProvince
`func (o *FleetCreate) UnsetProvince()`

UnsetProvince ensures that no value is present for Province, not even an explicit nil
### GetPostalCode

`func (o *FleetCreate) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *FleetCreate) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *FleetCreate) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *FleetCreate) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *FleetCreate) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *FleetCreate) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountryCode

`func (o *FleetCreate) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *FleetCreate) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *FleetCreate) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *FleetCreate) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### SetCountryCodeNil

`func (o *FleetCreate) SetCountryCodeNil(b bool)`

 SetCountryCodeNil sets the value for CountryCode to be an explicit nil

### UnsetCountryCode
`func (o *FleetCreate) UnsetCountryCode()`

UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
### GetProperties

`func (o *FleetCreate) GetProperties() []FleetPropertyCreate`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *FleetCreate) GetPropertiesOk() (*[]FleetPropertyCreate, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *FleetCreate) SetProperties(v []FleetPropertyCreate)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *FleetCreate) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


