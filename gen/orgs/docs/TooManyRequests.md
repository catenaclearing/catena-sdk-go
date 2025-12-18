# TooManyRequests

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] [default to 429]
**Message** | Pointer to **string** |  | [optional] [default to "Too Many Requests"]
**Detail** | Pointer to [**NullableRetryAfterDetail**](RetryAfterDetail.md) |  | [optional] 

## Methods

### NewTooManyRequests

`func NewTooManyRequests() *TooManyRequests`

NewTooManyRequests instantiates a new TooManyRequests object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTooManyRequestsWithDefaults

`func NewTooManyRequestsWithDefaults() *TooManyRequests`

NewTooManyRequestsWithDefaults instantiates a new TooManyRequests object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TooManyRequests) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TooManyRequests) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TooManyRequests) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *TooManyRequests) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *TooManyRequests) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TooManyRequests) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TooManyRequests) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TooManyRequests) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *TooManyRequests) GetDetail() RetryAfterDetail`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *TooManyRequests) GetDetailOk() (*RetryAfterDetail, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *TooManyRequests) SetDetail(v RetryAfterDetail)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *TooManyRequests) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *TooManyRequests) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *TooManyRequests) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


