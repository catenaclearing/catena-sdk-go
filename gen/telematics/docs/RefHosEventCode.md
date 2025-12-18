# RefHosEventCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventCode** | **string** | Canonical HOS event code (Catena-normalized). Primary key. | 
**EldEventTypeCode** | **string** | ELD &#39;Event Type&#39; code per vendor/ELD spec (normalized, e.g., 1-9). | 
**EldEventTypeName** | **string** | Human-readable name of the ELD event type (e.g., &#39;Duty Status Change&#39;). | 
**EldEventTypeDescription** | **string** | Detailed description of the ELD event type per spec. | 
**EldCode** | **string** | ELD &#39;Event Code&#39; within the event type (normalized vendor/ELD spec value). | 
**Description** | **string** | Human-readable description of the specific event code (display). | 

## Methods

### NewRefHosEventCode

`func NewRefHosEventCode(eventCode string, eldEventTypeCode string, eldEventTypeName string, eldEventTypeDescription string, eldCode string, description string, ) *RefHosEventCode`

NewRefHosEventCode instantiates a new RefHosEventCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefHosEventCodeWithDefaults

`func NewRefHosEventCodeWithDefaults() *RefHosEventCode`

NewRefHosEventCodeWithDefaults instantiates a new RefHosEventCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventCode

`func (o *RefHosEventCode) GetEventCode() string`

GetEventCode returns the EventCode field if non-nil, zero value otherwise.

### GetEventCodeOk

`func (o *RefHosEventCode) GetEventCodeOk() (*string, bool)`

GetEventCodeOk returns a tuple with the EventCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCode

`func (o *RefHosEventCode) SetEventCode(v string)`

SetEventCode sets EventCode field to given value.


### GetEldEventTypeCode

`func (o *RefHosEventCode) GetEldEventTypeCode() string`

GetEldEventTypeCode returns the EldEventTypeCode field if non-nil, zero value otherwise.

### GetEldEventTypeCodeOk

`func (o *RefHosEventCode) GetEldEventTypeCodeOk() (*string, bool)`

GetEldEventTypeCodeOk returns a tuple with the EldEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldEventTypeCode

`func (o *RefHosEventCode) SetEldEventTypeCode(v string)`

SetEldEventTypeCode sets EldEventTypeCode field to given value.


### GetEldEventTypeName

`func (o *RefHosEventCode) GetEldEventTypeName() string`

GetEldEventTypeName returns the EldEventTypeName field if non-nil, zero value otherwise.

### GetEldEventTypeNameOk

`func (o *RefHosEventCode) GetEldEventTypeNameOk() (*string, bool)`

GetEldEventTypeNameOk returns a tuple with the EldEventTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldEventTypeName

`func (o *RefHosEventCode) SetEldEventTypeName(v string)`

SetEldEventTypeName sets EldEventTypeName field to given value.


### GetEldEventTypeDescription

`func (o *RefHosEventCode) GetEldEventTypeDescription() string`

GetEldEventTypeDescription returns the EldEventTypeDescription field if non-nil, zero value otherwise.

### GetEldEventTypeDescriptionOk

`func (o *RefHosEventCode) GetEldEventTypeDescriptionOk() (*string, bool)`

GetEldEventTypeDescriptionOk returns a tuple with the EldEventTypeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldEventTypeDescription

`func (o *RefHosEventCode) SetEldEventTypeDescription(v string)`

SetEldEventTypeDescription sets EldEventTypeDescription field to given value.


### GetEldCode

`func (o *RefHosEventCode) GetEldCode() string`

GetEldCode returns the EldCode field if non-nil, zero value otherwise.

### GetEldCodeOk

`func (o *RefHosEventCode) GetEldCodeOk() (*string, bool)`

GetEldCodeOk returns a tuple with the EldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEldCode

`func (o *RefHosEventCode) SetEldCode(v string)`

SetEldCode sets EldCode field to given value.


### GetDescription

`func (o *RefHosEventCode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RefHosEventCode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RefHosEventCode) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


