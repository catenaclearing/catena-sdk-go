# RefHosRecordOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordOriginCode** | **string** | Canonical HOS record origin code (Catena-normalized). Primary key. | 
**EldCode** | **string** | ELD specification code for the record origin (vendor/ELD spec value). | 
**Description** | **string** | Human-readable description of the record origin (e.g., &#39;AUTOMATIC&#39;, &#39;MANUAL&#39;). | 

## Methods

### NewRefHosRecordOrigin

`func NewRefHosRecordOrigin(recordOriginCode string, eldCode string, description string, ) *RefHosRecordOrigin`

NewRefHosRecordOrigin instantiates a new RefHosRecordOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefHosRecordOriginWithDefaults

`func NewRefHosRecordOriginWithDefaults() *RefHosRecordOrigin`

NewRefHosRecordOriginWithDefaults instantiates a new RefHosRecordOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordOriginCode

`func (o *RefHosRecordOrigin) GetRecordOriginCode() string`

GetRecordOriginCode returns the RecordOriginCode field if non-nil, zero value otherwise.

### GetRecordOriginCodeOk

`func (o *RefHosRecordOrigin) GetRecordOriginCodeOk() (*string, bool)`

GetRecordOriginCodeOk returns a tuple with the RecordOriginCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordOriginCode

`func (o *RefHosRecordOrigin) SetRecordOriginCode(v string)`

SetRecordOriginCode sets RecordOriginCode field to given value.


### GetEldCode

`func (o *RefHosRecordOrigin) GetEldCode() string`

GetEldCode returns the EldCode field if non-nil, zero value otherwise.

### GetEldCodeOk

`func (o *RefHosRecordOrigin) GetEldCodeOk() (*string, bool)`

GetEldCodeOk returns a tuple with the EldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldCode

`func (o *RefHosRecordOrigin) SetEldCode(v string)`

SetEldCode sets EldCode field to given value.


### GetDescription

`func (o *RefHosRecordOrigin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RefHosRecordOrigin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RefHosRecordOrigin) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


