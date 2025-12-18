# Unauthorized

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] [default to 401]
**Message** | Pointer to **string** |  | [optional] [default to "Unauthorized"]
**Detail** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUnauthorized

`func NewUnauthorized() *Unauthorized`

NewUnauthorized instantiates a new Unauthorized object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnauthorizedWithDefaults

`func NewUnauthorizedWithDefaults() *Unauthorized`

NewUnauthorizedWithDefaults instantiates a new Unauthorized object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Unauthorized) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Unauthorized) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Unauthorized) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *Unauthorized) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *Unauthorized) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Unauthorized) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Unauthorized) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Unauthorized) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *Unauthorized) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *Unauthorized) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *Unauthorized) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *Unauthorized) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *Unauthorized) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *Unauthorized) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


