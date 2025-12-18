# UnprocessableEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] [default to 422]
**Message** | Pointer to **string** |  | [optional] [default to "Invalid Request Body"]
**Detail** | Pointer to [**[]ValidationErrorDetail**](ValidationErrorDetail.md) |  | [optional] 

## Methods

### NewUnprocessableEntity

`func NewUnprocessableEntity() *UnprocessableEntity`

NewUnprocessableEntity instantiates a new UnprocessableEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnprocessableEntityWithDefaults

`func NewUnprocessableEntityWithDefaults() *UnprocessableEntity`

NewUnprocessableEntityWithDefaults instantiates a new UnprocessableEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UnprocessableEntity) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnprocessableEntity) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnprocessableEntity) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *UnprocessableEntity) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *UnprocessableEntity) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnprocessableEntity) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnprocessableEntity) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UnprocessableEntity) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *UnprocessableEntity) GetDetail() []ValidationErrorDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *UnprocessableEntity) GetDetailOk() (*[]ValidationErrorDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *UnprocessableEntity) SetDetail(v []ValidationErrorDetail)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *UnprocessableEntity) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *UnprocessableEntity) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *UnprocessableEntity) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


