# GatewayTimeout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] [default to 504]
**Message** | Pointer to **string** |  | [optional] [default to "Gateway Timeout"]
**Detail** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGatewayTimeout

`func NewGatewayTimeout() *GatewayTimeout`

NewGatewayTimeout instantiates a new GatewayTimeout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTimeoutWithDefaults

`func NewGatewayTimeoutWithDefaults() *GatewayTimeout`

NewGatewayTimeoutWithDefaults instantiates a new GatewayTimeout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GatewayTimeout) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GatewayTimeout) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GatewayTimeout) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GatewayTimeout) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GatewayTimeout) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GatewayTimeout) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GatewayTimeout) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GatewayTimeout) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetDetail

`func (o *GatewayTimeout) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GatewayTimeout) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GatewayTimeout) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *GatewayTimeout) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### SetDetailNil

`func (o *GatewayTimeout) SetDetailNil(b bool)`

 SetDetailNil sets the value for Detail to be an explicit nil

### UnsetDetail
`func (o *GatewayTimeout) UnsetDetail()`

UnsetDetail ensures that no value is present for Detail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


