# InternalServerError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] [default to "Internal Server Error"]

## Methods

### NewInternalServerError

`func NewInternalServerError() *InternalServerError`

NewInternalServerError instantiates a new InternalServerError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalServerErrorWithDefaults

`func NewInternalServerErrorWithDefaults() *InternalServerError`

NewInternalServerErrorWithDefaults instantiates a new InternalServerError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InternalServerError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InternalServerError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InternalServerError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InternalServerError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


