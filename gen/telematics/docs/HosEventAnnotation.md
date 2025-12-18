# HosEventAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | Timestamp when the annotation was created. | [optional] 
**Remarks** | **string** | Free-text remarks or notes for the HOS event. | 

## Methods

### NewHosEventAnnotation

`func NewHosEventAnnotation(remarks string, ) *HosEventAnnotation`

NewHosEventAnnotation instantiates a new HosEventAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHosEventAnnotationWithDefaults

`func NewHosEventAnnotationWithDefaults() *HosEventAnnotation`

NewHosEventAnnotationWithDefaults instantiates a new HosEventAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *HosEventAnnotation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HosEventAnnotation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HosEventAnnotation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HosEventAnnotation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetRemarks

`func (o *HosEventAnnotation) GetRemarks() string`

GetRemarks returns the Remarks field if non-nil, zero value otherwise.

### GetRemarksOk

`func (o *HosEventAnnotation) GetRemarksOk() (*string, bool)`

GetRemarksOk returns a tuple with the Remarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemarks

`func (o *HosEventAnnotation) SetRemarks(v string)`

SetRemarks sets Remarks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


