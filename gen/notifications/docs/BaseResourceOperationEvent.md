# BaseResourceOperationEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the resource operation. | 
**CreatedAt** | **time.Time** | Timestamp when the resource operation was created. | 
**UpdatedAt** | **time.Time** | Timestamp when the resource operation was last updated. | 
**PartnerId** | **string** | The ID of the partner that created the operation. | 
**ConnectionId** | **string** | The ID of the connection associated with this operation. | 
**FleetRef** | Pointer to **NullableString** |  | [optional] 
**SourceName** | [**TspEnum**](TspEnum.md) | The source name of the operation. | 
**Resource** | [**ResourceEnum**](ResourceEnum.md) | The type of resource to operate on (e.g., vehicle, driver, hos_event). | 
**ResourceId** | Pointer to **NullableString** |  | [optional] 
**SourceId** | Pointer to **NullableString** |  | [optional] 
**OperationType** | [**ResourceOperationTypeEnum**](ResourceOperationTypeEnum.md) | The type of operation to perform (e.g., create, update). | 
**Status** | [**ResourceOperationStatusEnum**](ResourceOperationStatusEnum.md) | The current status of the operation. | 
**Payload** | **map[string]interface{}** | The payload of the operation, containing the resource attributes to be created or updated. | 
**Logs** | Pointer to [**[]ResourceOperationLogRead**](ResourceOperationLogRead.md) | Log entries for each attempt to execute this operation. | [optional] 

## Methods

### NewBaseResourceOperationEvent

`func NewBaseResourceOperationEvent(id string, createdAt time.Time, updatedAt time.Time, partnerId string, connectionId string, sourceName TspEnum, resource ResourceEnum, operationType ResourceOperationTypeEnum, status ResourceOperationStatusEnum, payload map[string]interface{}, ) *BaseResourceOperationEvent`

NewBaseResourceOperationEvent instantiates a new BaseResourceOperationEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseResourceOperationEventWithDefaults

`func NewBaseResourceOperationEventWithDefaults() *BaseResourceOperationEvent`

NewBaseResourceOperationEventWithDefaults instantiates a new BaseResourceOperationEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseResourceOperationEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseResourceOperationEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseResourceOperationEvent) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *BaseResourceOperationEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaseResourceOperationEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaseResourceOperationEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaseResourceOperationEvent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaseResourceOperationEvent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaseResourceOperationEvent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPartnerId

`func (o *BaseResourceOperationEvent) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *BaseResourceOperationEvent) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *BaseResourceOperationEvent) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.


### GetConnectionId

`func (o *BaseResourceOperationEvent) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *BaseResourceOperationEvent) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *BaseResourceOperationEvent) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetFleetRef

`func (o *BaseResourceOperationEvent) GetFleetRef() string`

GetFleetRef returns the FleetRef field if non-nil, zero value otherwise.

### GetFleetRefOk

`func (o *BaseResourceOperationEvent) GetFleetRefOk() (*string, bool)`

GetFleetRefOk returns a tuple with the FleetRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetRef

`func (o *BaseResourceOperationEvent) SetFleetRef(v string)`

SetFleetRef sets FleetRef field to given value.

### HasFleetRef

`func (o *BaseResourceOperationEvent) HasFleetRef() bool`

HasFleetRef returns a boolean if a field has been set.

### SetFleetRefNil

`func (o *BaseResourceOperationEvent) SetFleetRefNil(b bool)`

 SetFleetRefNil sets the value for FleetRef to be an explicit nil

### UnsetFleetRef
`func (o *BaseResourceOperationEvent) UnsetFleetRef()`

UnsetFleetRef ensures that no value is present for FleetRef, not even an explicit nil
### GetSourceName

`func (o *BaseResourceOperationEvent) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *BaseResourceOperationEvent) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *BaseResourceOperationEvent) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetResource

`func (o *BaseResourceOperationEvent) GetResource() ResourceEnum`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *BaseResourceOperationEvent) GetResourceOk() (*ResourceEnum, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *BaseResourceOperationEvent) SetResource(v ResourceEnum)`

SetResource sets Resource field to given value.


### GetResourceId

`func (o *BaseResourceOperationEvent) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *BaseResourceOperationEvent) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *BaseResourceOperationEvent) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *BaseResourceOperationEvent) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *BaseResourceOperationEvent) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *BaseResourceOperationEvent) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetSourceId

`func (o *BaseResourceOperationEvent) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *BaseResourceOperationEvent) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *BaseResourceOperationEvent) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *BaseResourceOperationEvent) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *BaseResourceOperationEvent) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *BaseResourceOperationEvent) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetOperationType

`func (o *BaseResourceOperationEvent) GetOperationType() ResourceOperationTypeEnum`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *BaseResourceOperationEvent) GetOperationTypeOk() (*ResourceOperationTypeEnum, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *BaseResourceOperationEvent) SetOperationType(v ResourceOperationTypeEnum)`

SetOperationType sets OperationType field to given value.


### GetStatus

`func (o *BaseResourceOperationEvent) GetStatus() ResourceOperationStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseResourceOperationEvent) GetStatusOk() (*ResourceOperationStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseResourceOperationEvent) SetStatus(v ResourceOperationStatusEnum)`

SetStatus sets Status field to given value.


### GetPayload

`func (o *BaseResourceOperationEvent) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *BaseResourceOperationEvent) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *BaseResourceOperationEvent) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.


### GetLogs

`func (o *BaseResourceOperationEvent) GetLogs() []ResourceOperationLogRead`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *BaseResourceOperationEvent) GetLogsOk() (*[]ResourceOperationLogRead, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *BaseResourceOperationEvent) SetLogs(v []ResourceOperationLogRead)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *BaseResourceOperationEvent) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


