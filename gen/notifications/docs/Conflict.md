# Conflict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] [default to 409]
**Message** | Pointer to **string** |  | [optional] [default to "Conflict"]
**Detail** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConflict

`func NewConflict() *Conflict`

NewConflict instantiates a new Conflict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConflictWithDefaults

`func NewConflictWithDefaults() *Conflict`

NewConflictWithDefaults instantiates a new Conflict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Conflict) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Conflict) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Conflict) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *Conflict) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *Conflict) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Conflict) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Conflict) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Conflict) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *Conflict) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Conflict) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Conflict) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *Conflict) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *Conflict) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *Conflict) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


