# RefHosMalfunctionCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MalfunctionCode** | **string** | Canonical malfunction/diagnostic code (Catena-normalized). Primary key. | 
**EldCode** | **string** | ELD specification code corresponding to the malfunction/diagnostic (vendor/ELD spec value). | 
**Category** | **string** | Normalized category (e.g., &#39;MALFUNCTION&#39;, &#39;DIAGNOSTIC&#39;). | 
**Description** | **string** | Human-readable description of the malfunction/diagnostic (display). | 

## Methods

### NewRefHosMalfunctionCode

`func NewRefHosMalfunctionCode(malfunctionCode string, eldCode string, category string, description string, ) *RefHosMalfunctionCode`

NewRefHosMalfunctionCode instantiates a new RefHosMalfunctionCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefHosMalfunctionCodeWithDefaults

`func NewRefHosMalfunctionCodeWithDefaults() *RefHosMalfunctionCode`

NewRefHosMalfunctionCodeWithDefaults instantiates a new RefHosMalfunctionCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMalfunctionCode

`func (o *RefHosMalfunctionCode) GetMalfunctionCode() string`

GetMalfunctionCode returns the MalfunctionCode field if non-nil, zero value otherwise.

### GetMalfunctionCodeOk

`func (o *RefHosMalfunctionCode) GetMalfunctionCodeOk() (*string, bool)`

GetMalfunctionCodeOk returns a tuple with the MalfunctionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMalfunctionCode

`func (o *RefHosMalfunctionCode) SetMalfunctionCode(v string)`

SetMalfunctionCode sets MalfunctionCode field to given value.


### GetEldCode

`func (o *RefHosMalfunctionCode) GetEldCode() string`

GetEldCode returns the EldCode field if non-nil, zero value otherwise.

### GetEldCodeOk

`func (o *RefHosMalfunctionCode) GetEldCodeOk() (*string, bool)`

GetEldCodeOk returns a tuple with the EldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldCode

`func (o *RefHosMalfunctionCode) SetEldCode(v string)`

SetEldCode sets EldCode field to given value.


### GetCategory

`func (o *RefHosMalfunctionCode) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *RefHosMalfunctionCode) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *RefHosMalfunctionCode) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetDescription

`func (o *RefHosMalfunctionCode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RefHosMalfunctionCode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RefHosMalfunctionCode) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


