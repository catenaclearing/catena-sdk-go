# ResourceOperationLogRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the log entry. | 
**CreatedAt** | **time.Time** | Timestamp when the log entry was created. | 
**UpdatedAt** | **time.Time** | Timestamp when the log entry was last updated. | 
**ResourceOperationId** | **string** | The ID of the resource operation this log entry belongs to. | 
**Status** | [**ResourceOperationLogStatusEnum**](ResourceOperationLogStatusEnum.md) | The status of the operation attempt. | 
**ErrorType** | Pointer to [**NullableResourceOperationErrorTypeEnum**](ResourceOperationErrorTypeEnum.md) |  | [optional] 
**ErrorDetails** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewResourceOperationLogRead

`func NewResourceOperationLogRead(id string, createdAt time.Time, updatedAt time.Time, resourceOperationId string, status ResourceOperationLogStatusEnum, ) *ResourceOperationLogRead`

NewResourceOperationLogRead instantiates a new ResourceOperationLogRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceOperationLogReadWithDefaults

`func NewResourceOperationLogReadWithDefaults() *ResourceOperationLogRead`

NewResourceOperationLogReadWithDefaults instantiates a new ResourceOperationLogRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceOperationLogRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceOperationLogRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceOperationLogRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ResourceOperationLogRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceOperationLogRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceOperationLogRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ResourceOperationLogRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourceOperationLogRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourceOperationLogRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetResourceOperationId

`func (o *ResourceOperationLogRead) GetResourceOperationId() string`

GetResourceOperationId returns the ResourceOperationId field if non-nil, zero value otherwise.

### GetResourceOperationIdOk

`func (o *ResourceOperationLogRead) GetResourceOperationIdOk() (*string, bool)`

GetResourceOperationIdOk returns a tuple with the ResourceOperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOperationId

`func (o *ResourceOperationLogRead) SetResourceOperationId(v string)`

SetResourceOperationId sets ResourceOperationId field to given value.


### GetStatus

`func (o *ResourceOperationLogRead) GetStatus() ResourceOperationLogStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceOperationLogRead) GetStatusOk() (*ResourceOperationLogStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceOperationLogRead) SetStatus(v ResourceOperationLogStatusEnum)`

SetStatus sets Status field to given value.


### GetErrorType

`func (o *ResourceOperationLogRead) GetErrorType() ResourceOperationErrorTypeEnum`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *ResourceOperationLogRead) GetErrorTypeOk() (*ResourceOperationErrorTypeEnum, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *ResourceOperationLogRead) SetErrorType(v ResourceOperationErrorTypeEnum)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *ResourceOperationLogRead) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### SetErrorTypeNil

`func (o *ResourceOperationLogRead) SetErrorTypeNil(b bool)`

 SetErrorTypeNil sets the value for ErrorType to be an explicit nil

### UnsetErrorType
`func (o *ResourceOperationLogRead) UnsetErrorType()`

UnsetErrorType ensures that no value is present for ErrorType, not even an explicit nil
### GetErrorDetails

`func (o *ResourceOperationLogRead) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *ResourceOperationLogRead) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *ResourceOperationLogRead) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *ResourceOperationLogRead) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### SetErrorDetailsNil

`func (o *ResourceOperationLogRead) SetErrorDetailsNil(b bool)`

 SetErrorDetailsNil sets the value for ErrorDetails to be an explicit nil

### UnsetErrorDetails
`func (o *ResourceOperationLogRead) UnsetErrorDetails()`

UnsetErrorDetails ensures that no value is present for ErrorDetails, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


