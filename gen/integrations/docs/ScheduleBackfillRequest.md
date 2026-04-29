# ScheduleBackfillRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resource** | [**ResourceEnum**](ResourceEnum.md) | The type of resource to backfill schedules for (e.g., VEHICLE, DRIVER, HOS, IFTA). | 
**TspId** | Pointer to **NullableString** |  | [optional] 
**ConnectionIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewScheduleBackfillRequest

`func NewScheduleBackfillRequest(resource ResourceEnum, ) *ScheduleBackfillRequest`

NewScheduleBackfillRequest instantiates a new ScheduleBackfillRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleBackfillRequestWithDefaults

`func NewScheduleBackfillRequestWithDefaults() *ScheduleBackfillRequest`

NewScheduleBackfillRequestWithDefaults instantiates a new ScheduleBackfillRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResource

`func (o *ScheduleBackfillRequest) GetResource() ResourceEnum`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ScheduleBackfillRequest) GetResourceOk() (*ResourceEnum, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ScheduleBackfillRequest) SetResource(v ResourceEnum)`

SetResource sets Resource field to given value.


### GetTspId

`func (o *ScheduleBackfillRequest) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *ScheduleBackfillRequest) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *ScheduleBackfillRequest) SetTspId(v string)`

SetTspId sets TspId field to given value.

### HasTspId

`func (o *ScheduleBackfillRequest) HasTspId() bool`

HasTspId returns a boolean if a field has been set.

### SetTspIdNil

`func (o *ScheduleBackfillRequest) SetTspIdNil(b bool)`

 SetTspIdNil sets the value for TspId to be an explicit nil

### UnsetTspId
`func (o *ScheduleBackfillRequest) UnsetTspId()`

UnsetTspId ensures that no value is present for TspId, not even an explicit nil
### GetConnectionIds

`func (o *ScheduleBackfillRequest) GetConnectionIds() []string`

GetConnectionIds returns the ConnectionIds field if non-nil, zero value otherwise.

### GetConnectionIdsOk

`func (o *ScheduleBackfillRequest) GetConnectionIdsOk() (*[]string, bool)`

GetConnectionIdsOk returns a tuple with the ConnectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionIds

`func (o *ScheduleBackfillRequest) SetConnectionIds(v []string)`

SetConnectionIds sets ConnectionIds field to given value.

### HasConnectionIds

`func (o *ScheduleBackfillRequest) HasConnectionIds() bool`

HasConnectionIds returns a boolean if a field has been set.

### SetConnectionIdsNil

`func (o *ScheduleBackfillRequest) SetConnectionIdsNil(b bool)`

 SetConnectionIdsNil sets the value for ConnectionIds to be an explicit nil

### UnsetConnectionIds
`func (o *ScheduleBackfillRequest) UnsetConnectionIds()`

UnsetConnectionIds ensures that no value is present for ConnectionIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


