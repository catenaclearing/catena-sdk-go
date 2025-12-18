# RetryAfterDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetryAfterSeconds** | **int32** |  | 
**Message** | **string** |  | 

## Methods

### NewRetryAfterDetail

`func NewRetryAfterDetail(retryAfterSeconds int32, message string, ) *RetryAfterDetail`

NewRetryAfterDetail instantiates a new RetryAfterDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryAfterDetailWithDefaults

`func NewRetryAfterDetailWithDefaults() *RetryAfterDetail`

NewRetryAfterDetailWithDefaults instantiates a new RetryAfterDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetryAfterSeconds

`func (o *RetryAfterDetail) GetRetryAfterSeconds() int32`

GetRetryAfterSeconds returns the RetryAfterSeconds field if non-nil, zero value otherwise.

### GetRetryAfterSecondsOk

`func (o *RetryAfterDetail) GetRetryAfterSecondsOk() (*int32, bool)`

GetRetryAfterSecondsOk returns a tuple with the RetryAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfterSeconds

`func (o *RetryAfterDetail) SetRetryAfterSeconds(v int32)`

SetRetryAfterSeconds sets RetryAfterSeconds field to given value.


### GetMessage

`func (o *RetryAfterDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RetryAfterDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RetryAfterDetail) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


