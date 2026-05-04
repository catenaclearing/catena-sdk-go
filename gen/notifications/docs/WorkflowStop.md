# WorkflowStop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceStopId** | Pointer to **NullableString** |  | [optional] 
**SequenceNumber** | Pointer to **NullableInt32** |  | [optional] 
**StopType** | Pointer to [**NullableStopTypeEnum**](StopTypeEnum.md) |  | [optional] 
**LocationName** | Pointer to **NullableString** |  | [optional] 
**Activities** | Pointer to [**[]StopActivityEnum**](StopActivityEnum.md) |  | [optional] 
**Location** | Pointer to [**NullablePoint**](Point.md) |  | [optional] 
**H3Index11** | Pointer to **NullableInt32** |  | [optional] 
**InferredAddress** | Pointer to [**NullableInferredAddress**](InferredAddress.md) |  | [optional] 
**Instructions** | Pointer to **NullableString** |  | [optional] 
**ScheduledArrival** | Pointer to **NullableTime** |  | [optional] 
**ScheduledDeparture** | Pointer to **NullableTime** |  | [optional] 
**ActualArrival** | Pointer to **NullableTime** |  | [optional] 
**ActualDeparture** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewWorkflowStop

`func NewWorkflowStop() *WorkflowStop`

NewWorkflowStop instantiates a new WorkflowStop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStopWithDefaults

`func NewWorkflowStopWithDefaults() *WorkflowStop`

NewWorkflowStopWithDefaults instantiates a new WorkflowStop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceStopId

`func (o *WorkflowStop) GetSourceStopId() string`

GetSourceStopId returns the SourceStopId field if non-nil, zero value otherwise.

### GetSourceStopIdOk

`func (o *WorkflowStop) GetSourceStopIdOk() (*string, bool)`

GetSourceStopIdOk returns a tuple with the SourceStopId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStopId

`func (o *WorkflowStop) SetSourceStopId(v string)`

SetSourceStopId sets SourceStopId field to given value.

### HasSourceStopId

`func (o *WorkflowStop) HasSourceStopId() bool`

HasSourceStopId returns a boolean if a field has been set.

### SetSourceStopIdNil

`func (o *WorkflowStop) SetSourceStopIdNil(b bool)`

 SetSourceStopIdNil sets the value for SourceStopId to be an explicit nil

### UnsetSourceStopId
`func (o *WorkflowStop) UnsetSourceStopId()`

UnsetSourceStopId ensures that no value is present for SourceStopId, not even an explicit nil
### GetSequenceNumber

`func (o *WorkflowStop) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *WorkflowStop) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *WorkflowStop) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *WorkflowStop) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### SetSequenceNumberNil

`func (o *WorkflowStop) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *WorkflowStop) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetStopType

`func (o *WorkflowStop) GetStopType() StopTypeEnum`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *WorkflowStop) GetStopTypeOk() (*StopTypeEnum, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *WorkflowStop) SetStopType(v StopTypeEnum)`

SetStopType sets StopType field to given value.

### HasStopType

`func (o *WorkflowStop) HasStopType() bool`

HasStopType returns a boolean if a field has been set.

### SetStopTypeNil

`func (o *WorkflowStop) SetStopTypeNil(b bool)`

 SetStopTypeNil sets the value for StopType to be an explicit nil

### UnsetStopType
`func (o *WorkflowStop) UnsetStopType()`

UnsetStopType ensures that no value is present for StopType, not even an explicit nil
### GetLocationName

`func (o *WorkflowStop) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *WorkflowStop) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *WorkflowStop) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *WorkflowStop) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### SetLocationNameNil

`func (o *WorkflowStop) SetLocationNameNil(b bool)`

 SetLocationNameNil sets the value for LocationName to be an explicit nil

### UnsetLocationName
`func (o *WorkflowStop) UnsetLocationName()`

UnsetLocationName ensures that no value is present for LocationName, not even an explicit nil
### GetActivities

`func (o *WorkflowStop) GetActivities() []StopActivityEnum`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *WorkflowStop) GetActivitiesOk() (*[]StopActivityEnum, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *WorkflowStop) SetActivities(v []StopActivityEnum)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *WorkflowStop) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### SetActivitiesNil

`func (o *WorkflowStop) SetActivitiesNil(b bool)`

 SetActivitiesNil sets the value for Activities to be an explicit nil

### UnsetActivities
`func (o *WorkflowStop) UnsetActivities()`

UnsetActivities ensures that no value is present for Activities, not even an explicit nil
### GetLocation

`func (o *WorkflowStop) GetLocation() Point`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WorkflowStop) GetLocationOk() (*Point, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WorkflowStop) SetLocation(v Point)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *WorkflowStop) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *WorkflowStop) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *WorkflowStop) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetH3Index11

`func (o *WorkflowStop) GetH3Index11() int32`

GetH3Index11 returns the H3Index11 field if non-nil, zero value otherwise.

### GetH3Index11Ok

`func (o *WorkflowStop) GetH3Index11Ok() (*int32, bool)`

GetH3Index11Ok returns a tuple with the H3Index11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH3Index11

`func (o *WorkflowStop) SetH3Index11(v int32)`

SetH3Index11 sets H3Index11 field to given value.

### HasH3Index11

`func (o *WorkflowStop) HasH3Index11() bool`

HasH3Index11 returns a boolean if a field has been set.

### SetH3Index11Nil

`func (o *WorkflowStop) SetH3Index11Nil(b bool)`

 SetH3Index11Nil sets the value for H3Index11 to be an explicit nil

### UnsetH3Index11
`func (o *WorkflowStop) UnsetH3Index11()`

UnsetH3Index11 ensures that no value is present for H3Index11, not even an explicit nil
### GetInferredAddress

`func (o *WorkflowStop) GetInferredAddress() InferredAddress`

GetInferredAddress returns the InferredAddress field if non-nil, zero value otherwise.

### GetInferredAddressOk

`func (o *WorkflowStop) GetInferredAddressOk() (*InferredAddress, bool)`

GetInferredAddressOk returns a tuple with the InferredAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferredAddress

`func (o *WorkflowStop) SetInferredAddress(v InferredAddress)`

SetInferredAddress sets InferredAddress field to given value.

### HasInferredAddress

`func (o *WorkflowStop) HasInferredAddress() bool`

HasInferredAddress returns a boolean if a field has been set.

### SetInferredAddressNil

`func (o *WorkflowStop) SetInferredAddressNil(b bool)`

 SetInferredAddressNil sets the value for InferredAddress to be an explicit nil

### UnsetInferredAddress
`func (o *WorkflowStop) UnsetInferredAddress()`

UnsetInferredAddress ensures that no value is present for InferredAddress, not even an explicit nil
### GetInstructions

`func (o *WorkflowStop) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *WorkflowStop) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *WorkflowStop) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *WorkflowStop) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### SetInstructionsNil

`func (o *WorkflowStop) SetInstructionsNil(b bool)`

 SetInstructionsNil sets the value for Instructions to be an explicit nil

### UnsetInstructions
`func (o *WorkflowStop) UnsetInstructions()`

UnsetInstructions ensures that no value is present for Instructions, not even an explicit nil
### GetScheduledArrival

`func (o *WorkflowStop) GetScheduledArrival() time.Time`

GetScheduledArrival returns the ScheduledArrival field if non-nil, zero value otherwise.

### GetScheduledArrivalOk

`func (o *WorkflowStop) GetScheduledArrivalOk() (*time.Time, bool)`

GetScheduledArrivalOk returns a tuple with the ScheduledArrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledArrival

`func (o *WorkflowStop) SetScheduledArrival(v time.Time)`

SetScheduledArrival sets ScheduledArrival field to given value.

### HasScheduledArrival

`func (o *WorkflowStop) HasScheduledArrival() bool`

HasScheduledArrival returns a boolean if a field has been set.

### SetScheduledArrivalNil

`func (o *WorkflowStop) SetScheduledArrivalNil(b bool)`

 SetScheduledArrivalNil sets the value for ScheduledArrival to be an explicit nil

### UnsetScheduledArrival
`func (o *WorkflowStop) UnsetScheduledArrival()`

UnsetScheduledArrival ensures that no value is present for ScheduledArrival, not even an explicit nil
### GetScheduledDeparture

`func (o *WorkflowStop) GetScheduledDeparture() time.Time`

GetScheduledDeparture returns the ScheduledDeparture field if non-nil, zero value otherwise.

### GetScheduledDepartureOk

`func (o *WorkflowStop) GetScheduledDepartureOk() (*time.Time, bool)`

GetScheduledDepartureOk returns a tuple with the ScheduledDeparture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDeparture

`func (o *WorkflowStop) SetScheduledDeparture(v time.Time)`

SetScheduledDeparture sets ScheduledDeparture field to given value.

### HasScheduledDeparture

`func (o *WorkflowStop) HasScheduledDeparture() bool`

HasScheduledDeparture returns a boolean if a field has been set.

### SetScheduledDepartureNil

`func (o *WorkflowStop) SetScheduledDepartureNil(b bool)`

 SetScheduledDepartureNil sets the value for ScheduledDeparture to be an explicit nil

### UnsetScheduledDeparture
`func (o *WorkflowStop) UnsetScheduledDeparture()`

UnsetScheduledDeparture ensures that no value is present for ScheduledDeparture, not even an explicit nil
### GetActualArrival

`func (o *WorkflowStop) GetActualArrival() time.Time`

GetActualArrival returns the ActualArrival field if non-nil, zero value otherwise.

### GetActualArrivalOk

`func (o *WorkflowStop) GetActualArrivalOk() (*time.Time, bool)`

GetActualArrivalOk returns a tuple with the ActualArrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualArrival

`func (o *WorkflowStop) SetActualArrival(v time.Time)`

SetActualArrival sets ActualArrival field to given value.

### HasActualArrival

`func (o *WorkflowStop) HasActualArrival() bool`

HasActualArrival returns a boolean if a field has been set.

### SetActualArrivalNil

`func (o *WorkflowStop) SetActualArrivalNil(b bool)`

 SetActualArrivalNil sets the value for ActualArrival to be an explicit nil

### UnsetActualArrival
`func (o *WorkflowStop) UnsetActualArrival()`

UnsetActualArrival ensures that no value is present for ActualArrival, not even an explicit nil
### GetActualDeparture

`func (o *WorkflowStop) GetActualDeparture() time.Time`

GetActualDeparture returns the ActualDeparture field if non-nil, zero value otherwise.

### GetActualDepartureOk

`func (o *WorkflowStop) GetActualDepartureOk() (*time.Time, bool)`

GetActualDepartureOk returns a tuple with the ActualDeparture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDeparture

`func (o *WorkflowStop) SetActualDeparture(v time.Time)`

SetActualDeparture sets ActualDeparture field to given value.

### HasActualDeparture

`func (o *WorkflowStop) HasActualDeparture() bool`

HasActualDeparture returns a boolean if a field has been set.

### SetActualDepartureNil

`func (o *WorkflowStop) SetActualDepartureNil(b bool)`

 SetActualDepartureNil sets the value for ActualDeparture to be an explicit nil

### UnsetActualDeparture
`func (o *WorkflowStop) UnsetActualDeparture()`

UnsetActualDeparture ensures that no value is present for ActualDeparture, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


