# ValidationErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** |  | 
**Input** | **string** |  | 
**Message** | **string** |  | 
**ErrorType** | **string** |  | 

## Methods

### NewValidationErrorDetail

`func NewValidationErrorDetail(path string, input string, message string, errorType string, ) *ValidationErrorDetail`

NewValidationErrorDetail instantiates a new ValidationErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorDetailWithDefaults

`func NewValidationErrorDetailWithDefaults() *ValidationErrorDetail`

NewValidationErrorDetailWithDefaults instantiates a new ValidationErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *ValidationErrorDetail) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ValidationErrorDetail) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ValidationErrorDetail) SetPath(v string)`

SetPath sets Path field to given value.


### GetInput

`func (o *ValidationErrorDetail) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *ValidationErrorDetail) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *ValidationErrorDetail) SetInput(v string)`

SetInput sets Input field to given value.


### GetMessage

`func (o *ValidationErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrorType

`func (o *ValidationErrorDetail) GetErrorType() string`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *ValidationErrorDetail) GetErrorTypeOk() (*string, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *ValidationErrorDetail) SetErrorType(v string)`

SetErrorType sets ErrorType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


