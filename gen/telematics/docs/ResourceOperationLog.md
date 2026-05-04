# ResourceOperationLog

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
**ResponseStatusCode** | Pointer to **NullableInt32** |  | [optional] 
**ResponseHeaders** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewResourceOperationLog

`func NewResourceOperationLog(id string, createdAt time.Time, updatedAt time.Time, resourceOperationId string, status ResourceOperationLogStatusEnum, ) *ResourceOperationLog`

NewResourceOperationLog instantiates a new ResourceOperationLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceOperationLogWithDefaults

`func NewResourceOperationLogWithDefaults() *ResourceOperationLog`

NewResourceOperationLogWithDefaults instantiates a new ResourceOperationLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceOperationLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceOperationLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceOperationLog) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ResourceOperationLog) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceOperationLog) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceOperationLog) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ResourceOperationLog) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourceOperationLog) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourceOperationLog) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetResourceOperationId

`func (o *ResourceOperationLog) GetResourceOperationId() string`

GetResourceOperationId returns the ResourceOperationId field if non-nil, zero value otherwise.

### GetResourceOperationIdOk

`func (o *ResourceOperationLog) GetResourceOperationIdOk() (*string, bool)`

GetResourceOperationIdOk returns a tuple with the ResourceOperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOperationId

`func (o *ResourceOperationLog) SetResourceOperationId(v string)`

SetResourceOperationId sets ResourceOperationId field to given value.


### GetStatus

`func (o *ResourceOperationLog) GetStatus() ResourceOperationLogStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceOperationLog) GetStatusOk() (*ResourceOperationLogStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceOperationLog) SetStatus(v ResourceOperationLogStatusEnum)`

SetStatus sets Status field to given value.


### GetErrorType

`func (o *ResourceOperationLog) GetErrorType() ResourceOperationErrorTypeEnum`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *ResourceOperationLog) GetErrorTypeOk() (*ResourceOperationErrorTypeEnum, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *ResourceOperationLog) SetErrorType(v ResourceOperationErrorTypeEnum)`

SetErrorType sets ErrorType field to given value.

### HasErrorType

`func (o *ResourceOperationLog) HasErrorType() bool`

HasErrorType returns a boolean if a field has been set.

### SetErrorTypeNil

`func (o *ResourceOperationLog) SetErrorTypeNil(b bool)`

 SetErrorTypeNil sets the value for ErrorType to be an explicit nil

### UnsetErrorType
`func (o *ResourceOperationLog) UnsetErrorType()`

UnsetErrorType ensures that no value is present for ErrorType, not even an explicit nil
### GetErrorDetails

`func (o *ResourceOperationLog) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *ResourceOperationLog) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *ResourceOperationLog) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *ResourceOperationLog) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### SetErrorDetailsNil

`func (o *ResourceOperationLog) SetErrorDetailsNil(b bool)`

 SetErrorDetailsNil sets the value for ErrorDetails to be an explicit nil

### UnsetErrorDetails
`func (o *ResourceOperationLog) UnsetErrorDetails()`

UnsetErrorDetails ensures that no value is present for ErrorDetails, not even an explicit nil
### GetResponseStatusCode

`func (o *ResourceOperationLog) GetResponseStatusCode() int32`

GetResponseStatusCode returns the ResponseStatusCode field if non-nil, zero value otherwise.

### GetResponseStatusCodeOk

`func (o *ResourceOperationLog) GetResponseStatusCodeOk() (*int32, bool)`

GetResponseStatusCodeOk returns a tuple with the ResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseStatusCode

`func (o *ResourceOperationLog) SetResponseStatusCode(v int32)`

SetResponseStatusCode sets ResponseStatusCode field to given value.

### HasResponseStatusCode

`func (o *ResourceOperationLog) HasResponseStatusCode() bool`

HasResponseStatusCode returns a boolean if a field has been set.

### SetResponseStatusCodeNil

`func (o *ResourceOperationLog) SetResponseStatusCodeNil(b bool)`

 SetResponseStatusCodeNil sets the value for ResponseStatusCode to be an explicit nil

### UnsetResponseStatusCode
`func (o *ResourceOperationLog) UnsetResponseStatusCode()`

UnsetResponseStatusCode ensures that no value is present for ResponseStatusCode, not even an explicit nil
### GetResponseHeaders

`func (o *ResourceOperationLog) GetResponseHeaders() map[string]string`

GetResponseHeaders returns the ResponseHeaders field if non-nil, zero value otherwise.

### GetResponseHeadersOk

`func (o *ResourceOperationLog) GetResponseHeadersOk() (*map[string]string, bool)`

GetResponseHeadersOk returns a tuple with the ResponseHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseHeaders

`func (o *ResourceOperationLog) SetResponseHeaders(v map[string]string)`

SetResponseHeaders sets ResponseHeaders field to given value.

### HasResponseHeaders

`func (o *ResourceOperationLog) HasResponseHeaders() bool`

HasResponseHeaders returns a boolean if a field has been set.

### SetResponseHeadersNil

`func (o *ResourceOperationLog) SetResponseHeadersNil(b bool)`

 SetResponseHeadersNil sets the value for ResponseHeaders to be an explicit nil

### UnsetResponseHeaders
`func (o *ResourceOperationLog) UnsetResponseHeaders()`

UnsetResponseHeaders ensures that no value is present for ResponseHeaders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


