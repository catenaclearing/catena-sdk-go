# RefHosViolationCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViolationCode** | **string** | Canonical HOS violation code (Catena-normalized). Primary key. | 
**Category** | **string** | Normalized violation category (e.g., &#39;DRIVING_HOURS&#39;, &#39;WORKDAY&#39;, &#39;REST&#39;). | 
**Name** | **string** | Short display name for the violation (human-readable). | 
**Description** | **string** | Detailed description of the violation rule/condition (display). | 

## Methods

### NewRefHosViolationCode

`func NewRefHosViolationCode(violationCode string, category string, name string, description string, ) *RefHosViolationCode`

NewRefHosViolationCode instantiates a new RefHosViolationCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefHosViolationCodeWithDefaults

`func NewRefHosViolationCodeWithDefaults() *RefHosViolationCode`

NewRefHosViolationCodeWithDefaults instantiates a new RefHosViolationCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViolationCode

`func (o *RefHosViolationCode) GetViolationCode() string`

GetViolationCode returns the ViolationCode field if non-nil, zero value otherwise.

### GetViolationCodeOk

`func (o *RefHosViolationCode) GetViolationCodeOk() (*string, bool)`

GetViolationCodeOk returns a tuple with the ViolationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCode

`func (o *RefHosViolationCode) SetViolationCode(v string)`

SetViolationCode sets ViolationCode field to given value.


### GetCategory

`func (o *RefHosViolationCode) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *RefHosViolationCode) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *RefHosViolationCode) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetName

`func (o *RefHosViolationCode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RefHosViolationCode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RefHosViolationCode) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RefHosViolationCode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RefHosViolationCode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RefHosViolationCode) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


