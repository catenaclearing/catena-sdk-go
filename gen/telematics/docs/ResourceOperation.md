# ResourceOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the resource operation. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Immutable: The datetime the record was created in Catena Telematics. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The datetime the record was last modified in Catena Telematics. | [optional] 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**PartnerId** | **string** | The ID of the partner that created the operation. | 
**ConnectionId** | **string** | The ID of the connection | 
**Resource** | [**ResourceEnum**](ResourceEnum.md) | The type of resource to operate on (e.g., vehicle, driver). | 
**ResourceId** | Pointer to **NullableString** |  | [optional] 
**SourceId** | Pointer to **NullableString** |  | [optional] 
**OperationType** | [**ResourceOperationTypeEnum**](ResourceOperationTypeEnum.md) | The type of operation to perform (e.g., create, update). | 
**Payload** | **map[string]interface{}** | The payload of the operation, containing the resource attributes to be created or updated. | 
**Status** | Pointer to [**ResourceOperationStatusEnum**](ResourceOperationStatusEnum.md) | The status of the operation | [optional] 
**Logs** | Pointer to [**[]ResourceOperationLog**](ResourceOperationLog.md) | Log entries for each attempt to execute this operation. | [optional] 

## Methods

### NewResourceOperation

`func NewResourceOperation(partnerId string, connectionId string, resource ResourceEnum, operationType ResourceOperationTypeEnum, payload map[string]interface{}, ) *ResourceOperation`

NewResourceOperation instantiates a new ResourceOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceOperationWithDefaults

`func NewResourceOperationWithDefaults() *ResourceOperation`

NewResourceOperationWithDefaults instantiates a new ResourceOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceOperation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceOperation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResourceOperation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceOperation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceOperation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResourceOperation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ResourceOperation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourceOperation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourceOperation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ResourceOperation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetDeletedAt

`func (o *ResourceOperation) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ResourceOperation) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ResourceOperation) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ResourceOperation) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *ResourceOperation) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *ResourceOperation) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetPartnerId

`func (o *ResourceOperation) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ResourceOperation) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ResourceOperation) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.


### GetConnectionId

`func (o *ResourceOperation) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ResourceOperation) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ResourceOperation) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetResource

`func (o *ResourceOperation) GetResource() ResourceEnum`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceOperation) GetResourceOk() (*ResourceEnum, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceOperation) SetResource(v ResourceEnum)`

SetResource sets Resource field to given value.


### GetResourceId

`func (o *ResourceOperation) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceOperation) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceOperation) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ResourceOperation) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *ResourceOperation) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *ResourceOperation) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetSourceId

`func (o *ResourceOperation) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ResourceOperation) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ResourceOperation) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *ResourceOperation) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### SetSourceIdNil

`func (o *ResourceOperation) SetSourceIdNil(b bool)`

 SetSourceIdNil sets the value for SourceId to be an explicit nil

### UnsetSourceId
`func (o *ResourceOperation) UnsetSourceId()`

UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
### GetOperationType

`func (o *ResourceOperation) GetOperationType() ResourceOperationTypeEnum`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *ResourceOperation) GetOperationTypeOk() (*ResourceOperationTypeEnum, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *ResourceOperation) SetOperationType(v ResourceOperationTypeEnum)`

SetOperationType sets OperationType field to given value.


### GetPayload

`func (o *ResourceOperation) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ResourceOperation) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ResourceOperation) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.


### GetStatus

`func (o *ResourceOperation) GetStatus() ResourceOperationStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceOperation) GetStatusOk() (*ResourceOperationStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceOperation) SetStatus(v ResourceOperationStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceOperation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLogs

`func (o *ResourceOperation) GetLogs() []ResourceOperationLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ResourceOperation) GetLogsOk() (*[]ResourceOperationLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ResourceOperation) SetLogs(v []ResourceOperationLog)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *ResourceOperation) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


