# RefHosRecordStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordStatusCode** | **string** | Canonical HOS record status code (Catena-normalized). Primary key. | 
**EldCode** | **string** | ELD specification code for the record status (vendor/ELD spec value). | 
**Description** | **string** | Human-readable description of the record status (e.g., &#39;ACTIVE&#39;, &#39;INACTIVE&#39;, &#39;CHANGED&#39;). | 

## Methods

### NewRefHosRecordStatus

`func NewRefHosRecordStatus(recordStatusCode string, eldCode string, description string, ) *RefHosRecordStatus`

NewRefHosRecordStatus instantiates a new RefHosRecordStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefHosRecordStatusWithDefaults

`func NewRefHosRecordStatusWithDefaults() *RefHosRecordStatus`

NewRefHosRecordStatusWithDefaults instantiates a new RefHosRecordStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordStatusCode

`func (o *RefHosRecordStatus) GetRecordStatusCode() string`

GetRecordStatusCode returns the RecordStatusCode field if non-nil, zero value otherwise.

### GetRecordStatusCodeOk

`func (o *RefHosRecordStatus) GetRecordStatusCodeOk() (*string, bool)`

GetRecordStatusCodeOk returns a tuple with the RecordStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordStatusCode

`func (o *RefHosRecordStatus) SetRecordStatusCode(v string)`

SetRecordStatusCode sets RecordStatusCode field to given value.


### GetEldCode

`func (o *RefHosRecordStatus) GetEldCode() string`

GetEldCode returns the EldCode field if non-nil, zero value otherwise.

### GetEldCodeOk

`func (o *RefHosRecordStatus) GetEldCodeOk() (*string, bool)`

GetEldCodeOk returns a tuple with the EldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldCode

`func (o *RefHosRecordStatus) SetEldCode(v string)`

SetEldCode sets EldCode field to given value.


### GetDescription

`func (o *RefHosRecordStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RefHosRecordStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RefHosRecordStatus) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


