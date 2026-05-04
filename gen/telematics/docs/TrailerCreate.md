# TrailerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | Unique identifier of the connection at Catena Telematics which will be used to create this resource. A connection represents a Fleet/TSP pairing. | 
**TrailerName** | Pointer to **NullableString** |  | [optional] 
**Oem** | Pointer to **NullableString** |  | [optional] 
**ModelType** | Pointer to **NullableString** |  | [optional] 
**ModelYear** | Pointer to **NullableInt32** |  | [optional] 
**Vin** | Pointer to **NullableString** |  | [optional] 
**LicensePlateRegion** | Pointer to **NullableString** |  | [optional] 
**LicensePlateCountry** | Pointer to **NullableString** |  | [optional] 
**LicensePlateNumber** | Pointer to **NullableString** |  | [optional] 
**IsActive** | Pointer to **NullableBool** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**EldId** | Pointer to **NullableString** |  | [optional] 
**EldSerialNumber** | Pointer to **NullableString** |  | [optional] 
**EldDeviceType** | Pointer to **NullableString** |  | [optional] 
**EldProductId** | Pointer to **NullableString** |  | [optional] 
**TotalAxles** | Pointer to **NullableInt32** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**TrailerLength** | Pointer to **NullableFloat32** |  | [optional] 
**TrailerType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTrailerCreate

`func NewTrailerCreate(connectionId string, ) *TrailerCreate`

NewTrailerCreate instantiates a new TrailerCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailerCreateWithDefaults

`func NewTrailerCreateWithDefaults() *TrailerCreate`

NewTrailerCreateWithDefaults instantiates a new TrailerCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *TrailerCreate) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *TrailerCreate) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *TrailerCreate) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetTrailerName

`func (o *TrailerCreate) GetTrailerName() string`

GetTrailerName returns the TrailerName field if non-nil, zero value otherwise.

### GetTrailerNameOk

`func (o *TrailerCreate) GetTrailerNameOk() (*string, bool)`

GetTrailerNameOk returns a tuple with the TrailerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerName

`func (o *TrailerCreate) SetTrailerName(v string)`

SetTrailerName sets TrailerName field to given value.

### HasTrailerName

`func (o *TrailerCreate) HasTrailerName() bool`

HasTrailerName returns a boolean if a field has been set.

### SetTrailerNameNil

`func (o *TrailerCreate) SetTrailerNameNil(b bool)`

 SetTrailerNameNil sets the value for TrailerName to be an explicit nil

### UnsetTrailerName
`func (o *TrailerCreate) UnsetTrailerName()`

UnsetTrailerName ensures that no value is present for TrailerName, not even an explicit nil
### GetOem

`func (o *TrailerCreate) GetOem() string`

GetOem returns the Oem field if non-nil, zero value otherwise.

### GetOemOk

`func (o *TrailerCreate) GetOemOk() (*string, bool)`

GetOemOk returns a tuple with the Oem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOem

`func (o *TrailerCreate) SetOem(v string)`

SetOem sets Oem field to given value.

### HasOem

`func (o *TrailerCreate) HasOem() bool`

HasOem returns a boolean if a field has been set.

### SetOemNil

`func (o *TrailerCreate) SetOemNil(b bool)`

 SetOemNil sets the value for Oem to be an explicit nil

### UnsetOem
`func (o *TrailerCreate) UnsetOem()`

UnsetOem ensures that no value is present for Oem, not even an explicit nil
### GetModelType

`func (o *TrailerCreate) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *TrailerCreate) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *TrailerCreate) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *TrailerCreate) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### SetModelTypeNil

`func (o *TrailerCreate) SetModelTypeNil(b bool)`

 SetModelTypeNil sets the value for ModelType to be an explicit nil

### UnsetModelType
`func (o *TrailerCreate) UnsetModelType()`

UnsetModelType ensures that no value is present for ModelType, not even an explicit nil
### GetModelYear

`func (o *TrailerCreate) GetModelYear() int32`

GetModelYear returns the ModelYear field if non-nil, zero value otherwise.

### GetModelYearOk

`func (o *TrailerCreate) GetModelYearOk() (*int32, bool)`

GetModelYearOk returns a tuple with the ModelYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelYear

`func (o *TrailerCreate) SetModelYear(v int32)`

SetModelYear sets ModelYear field to given value.

### HasModelYear

`func (o *TrailerCreate) HasModelYear() bool`

HasModelYear returns a boolean if a field has been set.

### SetModelYearNil

`func (o *TrailerCreate) SetModelYearNil(b bool)`

 SetModelYearNil sets the value for ModelYear to be an explicit nil

### UnsetModelYear
`func (o *TrailerCreate) UnsetModelYear()`

UnsetModelYear ensures that no value is present for ModelYear, not even an explicit nil
### GetVin

`func (o *TrailerCreate) GetVin() string`

GetVin returns the Vin field if non-nil, zero value otherwise.

### GetVinOk

`func (o *TrailerCreate) GetVinOk() (*string, bool)`

GetVinOk returns a tuple with the Vin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVin

`func (o *TrailerCreate) SetVin(v string)`

SetVin sets Vin field to given value.

### HasVin

`func (o *TrailerCreate) HasVin() bool`

HasVin returns a boolean if a field has been set.

### SetVinNil

`func (o *TrailerCreate) SetVinNil(b bool)`

 SetVinNil sets the value for Vin to be an explicit nil

### UnsetVin
`func (o *TrailerCreate) UnsetVin()`

UnsetVin ensures that no value is present for Vin, not even an explicit nil
### GetLicensePlateRegion

`func (o *TrailerCreate) GetLicensePlateRegion() string`

GetLicensePlateRegion returns the LicensePlateRegion field if non-nil, zero value otherwise.

### GetLicensePlateRegionOk

`func (o *TrailerCreate) GetLicensePlateRegionOk() (*string, bool)`

GetLicensePlateRegionOk returns a tuple with the LicensePlateRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateRegion

`func (o *TrailerCreate) SetLicensePlateRegion(v string)`

SetLicensePlateRegion sets LicensePlateRegion field to given value.

### HasLicensePlateRegion

`func (o *TrailerCreate) HasLicensePlateRegion() bool`

HasLicensePlateRegion returns a boolean if a field has been set.

### SetLicensePlateRegionNil

`func (o *TrailerCreate) SetLicensePlateRegionNil(b bool)`

 SetLicensePlateRegionNil sets the value for LicensePlateRegion to be an explicit nil

### UnsetLicensePlateRegion
`func (o *TrailerCreate) UnsetLicensePlateRegion()`

UnsetLicensePlateRegion ensures that no value is present for LicensePlateRegion, not even an explicit nil
### GetLicensePlateCountry

`func (o *TrailerCreate) GetLicensePlateCountry() string`

GetLicensePlateCountry returns the LicensePlateCountry field if non-nil, zero value otherwise.

### GetLicensePlateCountryOk

`func (o *TrailerCreate) GetLicensePlateCountryOk() (*string, bool)`

GetLicensePlateCountryOk returns a tuple with the LicensePlateCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateCountry

`func (o *TrailerCreate) SetLicensePlateCountry(v string)`

SetLicensePlateCountry sets LicensePlateCountry field to given value.

### HasLicensePlateCountry

`func (o *TrailerCreate) HasLicensePlateCountry() bool`

HasLicensePlateCountry returns a boolean if a field has been set.

### SetLicensePlateCountryNil

`func (o *TrailerCreate) SetLicensePlateCountryNil(b bool)`

 SetLicensePlateCountryNil sets the value for LicensePlateCountry to be an explicit nil

### UnsetLicensePlateCountry
`func (o *TrailerCreate) UnsetLicensePlateCountry()`

UnsetLicensePlateCountry ensures that no value is present for LicensePlateCountry, not even an explicit nil
### GetLicensePlateNumber

`func (o *TrailerCreate) GetLicensePlateNumber() string`

GetLicensePlateNumber returns the LicensePlateNumber field if non-nil, zero value otherwise.

### GetLicensePlateNumberOk

`func (o *TrailerCreate) GetLicensePlateNumberOk() (*string, bool)`

GetLicensePlateNumberOk returns a tuple with the LicensePlateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensePlateNumber

`func (o *TrailerCreate) SetLicensePlateNumber(v string)`

SetLicensePlateNumber sets LicensePlateNumber field to given value.

### HasLicensePlateNumber

`func (o *TrailerCreate) HasLicensePlateNumber() bool`

HasLicensePlateNumber returns a boolean if a field has been set.

### SetLicensePlateNumberNil

`func (o *TrailerCreate) SetLicensePlateNumberNil(b bool)`

 SetLicensePlateNumberNil sets the value for LicensePlateNumber to be an explicit nil

### UnsetLicensePlateNumber
`func (o *TrailerCreate) UnsetLicensePlateNumber()`

UnsetLicensePlateNumber ensures that no value is present for LicensePlateNumber, not even an explicit nil
### GetIsActive

`func (o *TrailerCreate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *TrailerCreate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *TrailerCreate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *TrailerCreate) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### SetIsActiveNil

`func (o *TrailerCreate) SetIsActiveNil(b bool)`

 SetIsActiveNil sets the value for IsActive to be an explicit nil

### UnsetIsActive
`func (o *TrailerCreate) UnsetIsActive()`

UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
### GetStatus

`func (o *TrailerCreate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrailerCreate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrailerCreate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrailerCreate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *TrailerCreate) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *TrailerCreate) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNotes

`func (o *TrailerCreate) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *TrailerCreate) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *TrailerCreate) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *TrailerCreate) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *TrailerCreate) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *TrailerCreate) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetEldId

`func (o *TrailerCreate) GetEldId() string`

GetEldId returns the EldId field if non-nil, zero value otherwise.

### GetEldIdOk

`func (o *TrailerCreate) GetEldIdOk() (*string, bool)`

GetEldIdOk returns a tuple with the EldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldId

`func (o *TrailerCreate) SetEldId(v string)`

SetEldId sets EldId field to given value.

### HasEldId

`func (o *TrailerCreate) HasEldId() bool`

HasEldId returns a boolean if a field has been set.

### SetEldIdNil

`func (o *TrailerCreate) SetEldIdNil(b bool)`

 SetEldIdNil sets the value for EldId to be an explicit nil

### UnsetEldId
`func (o *TrailerCreate) UnsetEldId()`

UnsetEldId ensures that no value is present for EldId, not even an explicit nil
### GetEldSerialNumber

`func (o *TrailerCreate) GetEldSerialNumber() string`

GetEldSerialNumber returns the EldSerialNumber field if non-nil, zero value otherwise.

### GetEldSerialNumberOk

`func (o *TrailerCreate) GetEldSerialNumberOk() (*string, bool)`

GetEldSerialNumberOk returns a tuple with the EldSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldSerialNumber

`func (o *TrailerCreate) SetEldSerialNumber(v string)`

SetEldSerialNumber sets EldSerialNumber field to given value.

### HasEldSerialNumber

`func (o *TrailerCreate) HasEldSerialNumber() bool`

HasEldSerialNumber returns a boolean if a field has been set.

### SetEldSerialNumberNil

`func (o *TrailerCreate) SetEldSerialNumberNil(b bool)`

 SetEldSerialNumberNil sets the value for EldSerialNumber to be an explicit nil

### UnsetEldSerialNumber
`func (o *TrailerCreate) UnsetEldSerialNumber()`

UnsetEldSerialNumber ensures that no value is present for EldSerialNumber, not even an explicit nil
### GetEldDeviceType

`func (o *TrailerCreate) GetEldDeviceType() string`

GetEldDeviceType returns the EldDeviceType field if non-nil, zero value otherwise.

### GetEldDeviceTypeOk

`func (o *TrailerCreate) GetEldDeviceTypeOk() (*string, bool)`

GetEldDeviceTypeOk returns a tuple with the EldDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldDeviceType

`func (o *TrailerCreate) SetEldDeviceType(v string)`

SetEldDeviceType sets EldDeviceType field to given value.

### HasEldDeviceType

`func (o *TrailerCreate) HasEldDeviceType() bool`

HasEldDeviceType returns a boolean if a field has been set.

### SetEldDeviceTypeNil

`func (o *TrailerCreate) SetEldDeviceTypeNil(b bool)`

 SetEldDeviceTypeNil sets the value for EldDeviceType to be an explicit nil

### UnsetEldDeviceType
`func (o *TrailerCreate) UnsetEldDeviceType()`

UnsetEldDeviceType ensures that no value is present for EldDeviceType, not even an explicit nil
### GetEldProductId

`func (o *TrailerCreate) GetEldProductId() string`

GetEldProductId returns the EldProductId field if non-nil, zero value otherwise.

### GetEldProductIdOk

`func (o *TrailerCreate) GetEldProductIdOk() (*string, bool)`

GetEldProductIdOk returns a tuple with the EldProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldProductId

`func (o *TrailerCreate) SetEldProductId(v string)`

SetEldProductId sets EldProductId field to given value.

### HasEldProductId

`func (o *TrailerCreate) HasEldProductId() bool`

HasEldProductId returns a boolean if a field has been set.

### SetEldProductIdNil

`func (o *TrailerCreate) SetEldProductIdNil(b bool)`

 SetEldProductIdNil sets the value for EldProductId to be an explicit nil

### UnsetEldProductId
`func (o *TrailerCreate) UnsetEldProductId()`

UnsetEldProductId ensures that no value is present for EldProductId, not even an explicit nil
### GetTotalAxles

`func (o *TrailerCreate) GetTotalAxles() int32`

GetTotalAxles returns the TotalAxles field if non-nil, zero value otherwise.

### GetTotalAxlesOk

`func (o *TrailerCreate) GetTotalAxlesOk() (*int32, bool)`

GetTotalAxlesOk returns a tuple with the TotalAxles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAxles

`func (o *TrailerCreate) SetTotalAxles(v int32)`

SetTotalAxles sets TotalAxles field to given value.

### HasTotalAxles

`func (o *TrailerCreate) HasTotalAxles() bool`

HasTotalAxles returns a boolean if a field has been set.

### SetTotalAxlesNil

`func (o *TrailerCreate) SetTotalAxlesNil(b bool)`

 SetTotalAxlesNil sets the value for TotalAxles to be an explicit nil

### UnsetTotalAxles
`func (o *TrailerCreate) UnsetTotalAxles()`

UnsetTotalAxles ensures that no value is present for TotalAxles, not even an explicit nil
### GetExternalId

`func (o *TrailerCreate) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *TrailerCreate) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *TrailerCreate) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *TrailerCreate) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *TrailerCreate) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *TrailerCreate) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetTrailerLength

`func (o *TrailerCreate) GetTrailerLength() float32`

GetTrailerLength returns the TrailerLength field if non-nil, zero value otherwise.

### GetTrailerLengthOk

`func (o *TrailerCreate) GetTrailerLengthOk() (*float32, bool)`

GetTrailerLengthOk returns a tuple with the TrailerLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerLength

`func (o *TrailerCreate) SetTrailerLength(v float32)`

SetTrailerLength sets TrailerLength field to given value.

### HasTrailerLength

`func (o *TrailerCreate) HasTrailerLength() bool`

HasTrailerLength returns a boolean if a field has been set.

### SetTrailerLengthNil

`func (o *TrailerCreate) SetTrailerLengthNil(b bool)`

 SetTrailerLengthNil sets the value for TrailerLength to be an explicit nil

### UnsetTrailerLength
`func (o *TrailerCreate) UnsetTrailerLength()`

UnsetTrailerLength ensures that no value is present for TrailerLength, not even an explicit nil
### GetTrailerType

`func (o *TrailerCreate) GetTrailerType() string`

GetTrailerType returns the TrailerType field if non-nil, zero value otherwise.

### GetTrailerTypeOk

`func (o *TrailerCreate) GetTrailerTypeOk() (*string, bool)`

GetTrailerTypeOk returns a tuple with the TrailerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailerType

`func (o *TrailerCreate) SetTrailerType(v string)`

SetTrailerType sets TrailerType field to given value.

### HasTrailerType

`func (o *TrailerCreate) HasTrailerType() bool`

HasTrailerType returns a boolean if a field has been set.

### SetTrailerTypeNil

`func (o *TrailerCreate) SetTrailerTypeNil(b bool)`

 SetTrailerTypeNil sets the value for TrailerType to be an explicit nil

### UnsetTrailerType
`func (o *TrailerCreate) UnsetTrailerType()`

UnsetTrailerType ensures that no value is present for TrailerType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


