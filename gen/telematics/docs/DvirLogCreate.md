# DvirLogCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | Unique identifier of the connection at Catena Telematics which will be used to create this resource. A connection represents a Fleet/TSP pairing. | 
**FleetRef** | Pointer to **NullableString** |  | [optional] 
**SourceDriverId** | Pointer to **NullableString** |  | [optional] 
**SourceVehicleId** | Pointer to **NullableString** |  | [optional] 
**SourceTrailerId** | Pointer to **NullableString** |  | [optional] 
**OccurredAt** | Pointer to **NullableTime** |  | [optional] 
**LogType** | Pointer to [**DvirLogTypeEnum**](DvirLogTypeEnum.md) | Type of DVIR log. Defaults to &#x60;unknown&#x60; when omitted. | [optional] 
**InspectedBy** | Pointer to **NullableString** |  | [optional] 
**DefectCount** | Pointer to **NullableInt32** |  | [optional] 
**Defects** | Pointer to [**[]DvirDefectEnum**](DvirDefectEnum.md) |  | [optional] 
**IsSafetyCritical** | Pointer to **NullableBool** |  | [optional] 
**DriverComment** | Pointer to **NullableString** |  | [optional] 
**CertifyComment** | Pointer to **NullableString** |  | [optional] 
**Odometer** | Pointer to **NullableFloat32** |  | [optional] 
**OdometerUnit** | Pointer to [**NullableDistanceUnitEnum**](DistanceUnitEnum.md) |  | [optional] 

## Methods

### NewDvirLogCreate

`func NewDvirLogCreate(connectionId string, ) *DvirLogCreate`

NewDvirLogCreate instantiates a new DvirLogCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDvirLogCreateWithDefaults

`func NewDvirLogCreateWithDefaults() *DvirLogCreate`

NewDvirLogCreateWithDefaults instantiates a new DvirLogCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *DvirLogCreate) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DvirLogCreate) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DvirLogCreate) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetFleetRef

`func (o *DvirLogCreate) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *DvirLogCreate) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *DvirLogCreate) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *DvirLogCreate) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *DvirLogCreate) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *DvirLogCreate) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetSourceDriverId

`func (o *DvirLogCreate) GetSourceDriverId() string`

GetSourceDriverId returns the SourceDriverId field if non-nil, zero value otherwise.

### GetSourceDriverIdOk

`func (o *DvirLogCreate) GetSourceDriverIdOk() (*string, bool)`

GetSourceDriverIdOk returns a tuple with the SourceDriverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDriverId

`func (o *DvirLogCreate) SetSourceDriverId(v string)`

SetSourceDriverId sets SourceDriverId field to given value.

### HasSourceDriverId

`func (o *DvirLogCreate) HasSourceDriverId() bool`

HasSourceDriverId returns a boolean if a field has been set.

### SetSourceDriverIdNil

`func (o *DvirLogCreate) SetSourceDriverIdNil(b bool)`

 SetSourceDriverIdNil sets the value for SourceDriverId to be an explicit nil

### UnsetSourceDriverId
`func (o *DvirLogCreate) UnsetSourceDriverId()`

UnsetSourceDriverId ensures that no value is present for SourceDriverId, not even an explicit nil
### GetSourceVehicleId

`func (o *DvirLogCreate) GetSourceVehicleId() string`

GetSourceVehicleId returns the SourceVehicleId field if non-nil, zero value otherwise.

### GetSourceVehicleIdOk

`func (o *DvirLogCreate) GetSourceVehicleIdOk() (*string, bool)`

GetSourceVehicleIdOk returns a tuple with the SourceVehicleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVehicleId

`func (o *DvirLogCreate) SetSourceVehicleId(v string)`

SetSourceVehicleId sets SourceVehicleId field to given value.

### HasSourceVehicleId

`func (o *DvirLogCreate) HasSourceVehicleId() bool`

HasSourceVehicleId returns a boolean if a field has been set.

### SetSourceVehicleIdNil

`func (o *DvirLogCreate) SetSourceVehicleIdNil(b bool)`

 SetSourceVehicleIdNil sets the value for SourceVehicleId to be an explicit nil

### UnsetSourceVehicleId
`func (o *DvirLogCreate) UnsetSourceVehicleId()`

UnsetSourceVehicleId ensures that no value is present for SourceVehicleId, not even an explicit nil
### GetSourceTrailerId

`func (o *DvirLogCreate) GetSourceTrailerId() string`

GetSourceTrailerId returns the SourceTrailerId field if non-nil, zero value otherwise.

### GetSourceTrailerIdOk

`func (o *DvirLogCreate) GetSourceTrailerIdOk() (*string, bool)`

GetSourceTrailerIdOk returns a tuple with the SourceTrailerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTrailerId

`func (o *DvirLogCreate) SetSourceTrailerId(v string)`

SetSourceTrailerId sets SourceTrailerId field to given value.

### HasSourceTrailerId

`func (o *DvirLogCreate) HasSourceTrailerId() bool`

HasSourceTrailerId returns a boolean if a field has been set.

### SetSourceTrailerIdNil

`func (o *DvirLogCreate) SetSourceTrailerIdNil(b bool)`

 SetSourceTrailerIdNil sets the value for SourceTrailerId to be an explicit nil

### UnsetSourceTrailerId
`func (o *DvirLogCreate) UnsetSourceTrailerId()`

UnsetSourceTrailerId ensures that no value is present for SourceTrailerId, not even an explicit nil
### GetOccurredAt

`func (o *DvirLogCreate) GetOccurredAt() time.Time`

GetOccurredAt returns the OccurredAt field if non-nil, zero value otherwise.

### GetOccurredAtOk

`func (o *DvirLogCreate) GetOccurredAtOk() (*time.Time, bool)`

GetOccurredAtOk returns a tuple with the OccurredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurredAt

`func (o *DvirLogCreate) SetOccurredAt(v time.Time)`

SetOccurredAt sets OccurredAt field to given value.

### HasOccurredAt

`func (o *DvirLogCreate) HasOccurredAt() bool`

HasOccurredAt returns a boolean if a field has been set.

### SetOccurredAtNil

`func (o *DvirLogCreate) SetOccurredAtNil(b bool)`

 SetOccurredAtNil sets the value for OccurredAt to be an explicit nil

### UnsetOccurredAt
`func (o *DvirLogCreate) UnsetOccurredAt()`

UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
### GetLogType

`func (o *DvirLogCreate) GetLogType() DvirLogTypeEnum`

GetLogType returns the LogType field if non-nil, zero value otherwise.

### GetLogTypeOk

`func (o *DvirLogCreate) GetLogTypeOk() (*DvirLogTypeEnum, bool)`

GetLogTypeOk returns a tuple with the LogType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogType

`func (o *DvirLogCreate) SetLogType(v DvirLogTypeEnum)`

SetLogType sets LogType field to given value.

### HasLogType

`func (o *DvirLogCreate) HasLogType() bool`

HasLogType returns a boolean if a field has been set.

### GetInspectedBy

`func (o *DvirLogCreate) GetInspectedBy() string`

GetInspectedBy returns the InspectedBy field if non-nil, zero value otherwise.

### GetInspectedByOk

`func (o *DvirLogCreate) GetInspectedByOk() (*string, bool)`

GetInspectedByOk returns a tuple with the InspectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedBy

`func (o *DvirLogCreate) SetInspectedBy(v string)`

SetInspectedBy sets InspectedBy field to given value.

### HasInspectedBy

`func (o *DvirLogCreate) HasInspectedBy() bool`

HasInspectedBy returns a boolean if a field has been set.

### SetInspectedByNil

`func (o *DvirLogCreate) SetInspectedByNil(b bool)`

 SetInspectedByNil sets the value for InspectedBy to be an explicit nil

### UnsetInspectedBy
`func (o *DvirLogCreate) UnsetInspectedBy()`

UnsetInspectedBy ensures that no value is present for InspectedBy, not even an explicit nil
### GetDefectCount

`func (o *DvirLogCreate) GetDefectCount() int32`

GetDefectCount returns the DefectCount field if non-nil, zero value otherwise.

### GetDefectCountOk

`func (o *DvirLogCreate) GetDefectCountOk() (*int32, bool)`

GetDefectCountOk returns a tuple with the DefectCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefectCount

`func (o *DvirLogCreate) SetDefectCount(v int32)`

SetDefectCount sets DefectCount field to given value.

### HasDefectCount

`func (o *DvirLogCreate) HasDefectCount() bool`

HasDefectCount returns a boolean if a field has been set.

### SetDefectCountNil

`func (o *DvirLogCreate) SetDefectCountNil(b bool)`

 SetDefectCountNil sets the value for DefectCount to be an explicit nil

### UnsetDefectCount
`func (o *DvirLogCreate) UnsetDefectCount()`

UnsetDefectCount ensures that no value is present for DefectCount, not even an explicit nil
### GetDefects

`func (o *DvirLogCreate) GetDefects() []DvirDefectEnum`

GetDefects returns the Defects field if non-nil, zero value otherwise.

### GetDefectsOk

`func (o *DvirLogCreate) GetDefectsOk() (*[]DvirDefectEnum, bool)`

GetDefectsOk returns a tuple with the Defects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefects

`func (o *DvirLogCreate) SetDefects(v []DvirDefectEnum)`

SetDefects sets Defects field to given value.

### HasDefects

`func (o *DvirLogCreate) HasDefects() bool`

HasDefects returns a boolean if a field has been set.

### SetDefectsNil

`func (o *DvirLogCreate) SetDefectsNil(b bool)`

 SetDefectsNil sets the value for Defects to be an explicit nil

### UnsetDefects
`func (o *DvirLogCreate) UnsetDefects()`

UnsetDefects ensures that no value is present for Defects, not even an explicit nil
### GetIsSafetyCritical

`func (o *DvirLogCreate) GetIsSafetyCritical() bool`

GetIsSafetyCritical returns the IsSafetyCritical field if non-nil, zero value otherwise.

### GetIsSafetyCriticalOk

`func (o *DvirLogCreate) GetIsSafetyCriticalOk() (*bool, bool)`

GetIsSafetyCriticalOk returns a tuple with the IsSafetyCritical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSafetyCritical

`func (o *DvirLogCreate) SetIsSafetyCritical(v bool)`

SetIsSafetyCritical sets IsSafetyCritical field to given value.

### HasIsSafetyCritical

`func (o *DvirLogCreate) HasIsSafetyCritical() bool`

HasIsSafetyCritical returns a boolean if a field has been set.

### SetIsSafetyCriticalNil

`func (o *DvirLogCreate) SetIsSafetyCriticalNil(b bool)`

 SetIsSafetyCriticalNil sets the value for IsSafetyCritical to be an explicit nil

### UnsetIsSafetyCritical
`func (o *DvirLogCreate) UnsetIsSafetyCritical()`

UnsetIsSafetyCritical ensures that no value is present for IsSafetyCritical, not even an explicit nil
### GetDriverComment

`func (o *DvirLogCreate) GetDriverComment() string`

GetDriverComment returns the DriverComment field if non-nil, zero value otherwise.

### GetDriverCommentOk

`func (o *DvirLogCreate) GetDriverCommentOk() (*string, bool)`

GetDriverCommentOk returns a tuple with the DriverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverComment

`func (o *DvirLogCreate) SetDriverComment(v string)`

SetDriverComment sets DriverComment field to given value.

### HasDriverComment

`func (o *DvirLogCreate) HasDriverComment() bool`

HasDriverComment returns a boolean if a field has been set.

### SetDriverCommentNil

`func (o *DvirLogCreate) SetDriverCommentNil(b bool)`

 SetDriverCommentNil sets the value for DriverComment to be an explicit nil

### UnsetDriverComment
`func (o *DvirLogCreate) UnsetDriverComment()`

UnsetDriverComment ensures that no value is present for DriverComment, not even an explicit nil
### GetCertifyComment

`func (o *DvirLogCreate) GetCertifyComment() string`

GetCertifyComment returns the CertifyComment field if non-nil, zero value otherwise.

### GetCertifyCommentOk

`func (o *DvirLogCreate) GetCertifyCommentOk() (*string, bool)`

GetCertifyCommentOk returns a tuple with the CertifyComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifyComment

`func (o *DvirLogCreate) SetCertifyComment(v string)`

SetCertifyComment sets CertifyComment field to given value.

### HasCertifyComment

`func (o *DvirLogCreate) HasCertifyComment() bool`

HasCertifyComment returns a boolean if a field has been set.

### SetCertifyCommentNil

`func (o *DvirLogCreate) SetCertifyCommentNil(b bool)`

 SetCertifyCommentNil sets the value for CertifyComment to be an explicit nil

### UnsetCertifyComment
`func (o *DvirLogCreate) UnsetCertifyComment()`

UnsetCertifyComment ensures that no value is present for CertifyComment, not even an explicit nil
### GetOdometer

`func (o *DvirLogCreate) GetOdometer() float32`

GetOdometer returns the Odometer field if non-nil, zero value otherwise.

### GetOdometerOk

`func (o *DvirLogCreate) GetOdometerOk() (*float32, bool)`

GetOdometerOk returns a tuple with the Odometer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometer

`func (o *DvirLogCreate) SetOdometer(v float32)`

SetOdometer sets Odometer field to given value.

### HasOdometer

`func (o *DvirLogCreate) HasOdometer() bool`

HasOdometer returns a boolean if a field has been set.

### SetOdometerNil

`func (o *DvirLogCreate) SetOdometerNil(b bool)`

 SetOdometerNil sets the value for Odometer to be an explicit nil

### UnsetOdometer
`func (o *DvirLogCreate) UnsetOdometer()`

UnsetOdometer ensures that no value is present for Odometer, not even an explicit nil
### GetOdometerUnit

`func (o *DvirLogCreate) GetOdometerUnit() DistanceUnitEnum`

GetOdometerUnit returns the OdometerUnit field if non-nil, zero value otherwise.

### GetOdometerUnitOk

`func (o *DvirLogCreate) GetOdometerUnitOk() (*DistanceUnitEnum, bool)`

GetOdometerUnitOk returns a tuple with the OdometerUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdometerUnit

`func (o *DvirLogCreate) SetOdometerUnit(v DistanceUnitEnum)`

SetOdometerUnit sets OdometerUnit field to given value.

### HasOdometerUnit

`func (o *DvirLogCreate) HasOdometerUnit() bool`

HasOdometerUnit returns a boolean if a field has been set.

### SetOdometerUnitNil

`func (o *DvirLogCreate) SetOdometerUnitNil(b bool)`

 SetOdometerUnitNil sets the value for OdometerUnit to be an explicit nil

### UnsetOdometerUnit
`func (o *DvirLogCreate) UnsetOdometerUnit()`

UnsetOdometerUnit ensures that no value is present for OdometerUnit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


