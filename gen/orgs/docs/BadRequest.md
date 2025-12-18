# BadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] [default to 400]
**Message** | Pointer to **string** |  | [optional] [default to "Bad Request"]
**Detail** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBadRequest

`func NewBadRequest() *BadRequest`

NewBadRequest instantiates a new BadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestWithDefaults

`func NewBadRequestWithDefaults() *BadRequest`

NewBadRequestWithDefaults instantiates a new BadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *BadRequest) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BadRequest) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BadRequest) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *BadRequest) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *BadRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BadRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BadRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *BadRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *BadRequest) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *BadRequest) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *BadRequest) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *BadRequest) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *BadRequest) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *BadRequest) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


