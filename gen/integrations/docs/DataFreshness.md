# DataFreshness

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | [**ResourceEnum**](ResourceEnum.md) | The telematics resource type this entry describes. | 
**ShareLevel** | Pointer to [**NullableShareLevelEnum**](ShareLevelEnum.md) |  | [optional] 
**FleetId** | **string** | Unique identifier of the fleet this freshness record applies to. | 
**FleetRef** | Pointer to **NullableString** |  | [optional] 
**ConnectionId** | **string** | Unique identifier of the telematics provider connection that supplies this resource&#39;s data. | 
**ConnectionDescription** | Pointer to **NullableString** |  | [optional] 
**ScheduleId** | **string** | Unique identifier of the schedule that controls automatic data fetching for this resource. | 
**SourceName** | **string** | Name of the telematics provider (e.g., geotab, samsara, motive). | 
**TspId** | **string** | Unique identifier of the telematics service provider (TSP) at Catena. | 
**TspSlug** | **string** | Human-readable identifier of the telematics service provider (TSP) at Catena. | 
**LastSuccessfulExecutionAt** | **NullableTime** |  | 
**NextExecutionAt** | **NullableTime** |  | 
**ScheduleStatus** | [**NullableScheduleStatusEnum**](ScheduleStatusEnum.md) |  | 

## Methods

### NewDataFreshness

`func NewDataFreshness(resource ResourceEnum, fleetId string, connectionId string, scheduleId string, sourceName string, tspId string, tspSlug string, lastSuccessfulExecutionAt NullableTime, nextExecutionAt NullableTime, scheduleStatus NullableScheduleStatusEnum, ) *DataFreshness`

NewDataFreshness instantiates a new DataFreshness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataFreshnessWithDefaults

`func NewDataFreshnessWithDefaults() *DataFreshness`

NewDataFreshnessWithDefaults instantiates a new DataFreshness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *DataFreshness) GetResource() ResourceEnum`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *DataFreshness) GetResourceOk() (*ResourceEnum, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *DataFreshness) SetResource(v ResourceEnum)`

SetResource sets Resource field to given value.


### GetShareLevel

`func (o *DataFreshness) GetShareLevel() ShareLevelEnum`

GetShareLevel returns the ShareLevel field if non-nil, zero value otherwise.

### GetShareLevelOk

`func (o *DataFreshness) GetShareLevelOk() (*ShareLevelEnum, bool)`

GetShareLevelOk returns a tuple with the ShareLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareLevel

`func (o *DataFreshness) SetShareLevel(v ShareLevelEnum)`

SetShareLevel sets ShareLevel field to given value.

### HasShareLevel

`func (o *DataFreshness) HasShareLevel() bool`

HasShareLevel returns a boolean if a field has been set.

### SetShareLevelNil

`func (o *DataFreshness) SetShareLevelNil(b bool)`

 SetShareLevelNil sets the value for ShareLevel to be an explicit nil

### UnsetShareLevel
`func (o *DataFreshness) UnsetShareLevel()`

UnsetShareLevel ensures that no value is present for ShareLevel, not even an explicit nil
### GetFleetId

`func (o *DataFreshness) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *DataFreshness) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *DataFreshness) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetFleetRef

`func (o *DataFreshness) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *DataFreshness) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *DataFreshness) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *DataFreshness) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *DataFreshness) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *DataFreshness) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetConnectionId

`func (o *DataFreshness) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DataFreshness) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DataFreshness) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetConnectionDescription

`func (o *DataFreshness) GetConnectionDescription() string`

GetConnectionDescription returns the ConnectionDescription field if non-nil, zero value otherwise.

### GetConnectionDescriptionOk

`func (o *DataFreshness) GetConnectionDescriptionOk() (*string, bool)`

GetConnectionDescriptionOk returns a tuple with the ConnectionDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionDescription

`func (o *DataFreshness) SetConnectionDescription(v string)`

SetConnectionDescription sets ConnectionDescription field to given value.

### HasConnectionDescription

`func (o *DataFreshness) HasConnectionDescription() bool`

HasConnectionDescription returns a boolean if a field has been set.

### SetConnectionDescriptionNil

`func (o *DataFreshness) SetConnectionDescriptionNil(b bool)`

 SetConnectionDescriptionNil sets the value for ConnectionDescription to be an explicit nil

### UnsetConnectionDescription
`func (o *DataFreshness) UnsetConnectionDescription()`

UnsetConnectionDescription ensures that no value is present for ConnectionDescription, not even an explicit nil
### GetScheduleId

`func (o *DataFreshness) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *DataFreshness) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *DataFreshness) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.


### GetSourceName

`func (o *DataFreshness) GetSourceName() string`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *DataFreshness) GetSourceNameOk() (*string, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *DataFreshness) SetSourceName(v string)`

SetSourceName sets SourceName field to given value.


### GetTspId

`func (o *DataFreshness) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *DataFreshness) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *DataFreshness) SetTspId(v string)`

SetTspId sets TspId field to given value.


### GetTspSlug

`func (o *DataFreshness) GetTspSlug() string`

GetTspSlug returns the TspSlug field if non-nil, zero value otherwise.

### GetTspSlugOk

`func (o *DataFreshness) GetTspSlugOk() (*string, bool)`

GetTspSlugOk returns a tuple with the TspSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspSlug

`func (o *DataFreshness) SetTspSlug(v string)`

SetTspSlug sets TspSlug field to given value.


### GetLastSuccessfulExecutionAt

`func (o *DataFreshness) GetLastSuccessfulExecutionAt() time.Time`

GetLastSuccessfulExecutionAt returns the LastSuccessfulExecutionAt field if non-nil, zero value otherwise.

### GetLastSuccessfulExecutionAtOk

`func (o *DataFreshness) GetLastSuccessfulExecutionAtOk() (*time.Time, bool)`

GetLastSuccessfulExecutionAtOk returns a tuple with the LastSuccessfulExecutionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulExecutionAt

`func (o *DataFreshness) SetLastSuccessfulExecutionAt(v time.Time)`

SetLastSuccessfulExecutionAt sets LastSuccessfulExecutionAt field to given value.


### SetLastSuccessfulExecutionAtNil

`func (o *DataFreshness) SetLastSuccessfulExecutionAtNil(b bool)`

 SetLastSuccessfulExecutionAtNil sets the value for LastSuccessfulExecutionAt to be an explicit nil

### UnsetLastSuccessfulExecutionAt
`func (o *DataFreshness) UnsetLastSuccessfulExecutionAt()`

UnsetLastSuccessfulExecutionAt ensures that no value is present for LastSuccessfulExecutionAt, not even an explicit nil
### GetNextExecutionAt

`func (o *DataFreshness) GetNextExecutionAt() time.Time`

GetNextExecutionAt returns the NextExecutionAt field if non-nil, zero value otherwise.

### GetNextExecutionAtOk

`func (o *DataFreshness) GetNextExecutionAtOk() (*time.Time, bool)`

GetNextExecutionAtOk returns a tuple with the NextExecutionAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecutionAt

`func (o *DataFreshness) SetNextExecutionAt(v time.Time)`

SetNextExecutionAt sets NextExecutionAt field to given value.


### SetNextExecutionAtNil

`func (o *DataFreshness) SetNextExecutionAtNil(b bool)`

 SetNextExecutionAtNil sets the value for NextExecutionAt to be an explicit nil

### UnsetNextExecutionAt
`func (o *DataFreshness) UnsetNextExecutionAt()`

UnsetNextExecutionAt ensures that no value is present for NextExecutionAt, not even an explicit nil
### GetScheduleStatus

`func (o *DataFreshness) GetScheduleStatus() ScheduleStatusEnum`

GetScheduleStatus returns the ScheduleStatus field if non-nil, zero value otherwise.

### GetScheduleStatusOk

`func (o *DataFreshness) GetScheduleStatusOk() (*ScheduleStatusEnum, bool)`

GetScheduleStatusOk returns a tuple with the ScheduleStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleStatus

`func (o *DataFreshness) SetScheduleStatus(v ScheduleStatusEnum)`

SetScheduleStatus sets ScheduleStatus field to given value.


### SetScheduleStatusNil

`func (o *DataFreshness) SetScheduleStatusNil(b bool)`

 SetScheduleStatusNil sets the value for ScheduleStatus to be an explicit nil

### UnsetScheduleStatus
`func (o *DataFreshness) UnsetScheduleStatus()`

UnsetScheduleStatus ensures that no value is present for ScheduleStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


