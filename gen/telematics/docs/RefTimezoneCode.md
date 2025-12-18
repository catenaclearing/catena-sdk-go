# RefTimezoneCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimezoneCode** | **string** | Canonical timezone code used by Catena (normalized, e.g., &#39;AMER_CHICAGO&#39;). Primary key. | 
**Timezone** | **string** | IANA time zone identifier (e.g., &#39;America/Chicago&#39;). | 
**StandardUtcOffsetMinutes** | **int32** | Standard (non-DST) UTC offset in minutes (e.g., -360 for UTC-06:00). | 
**ObservesDst** | **bool** | Indicates whether this timezone observes daylight saving time (DST). | 
**DstUtcOffsetMinutes** | Pointer to **NullableInt32** |  | [optional] 
**DstStartRule** | Pointer to **NullableString** |  | [optional] 
**DstEndRule** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRefTimezoneCode

`func NewRefTimezoneCode(timezoneCode string, timezone string, standardUtcOffsetMinutes int32, observesDst bool, ) *RefTimezoneCode`

NewRefTimezoneCode instantiates a new RefTimezoneCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefTimezoneCodeWithDefaults

`func NewRefTimezoneCodeWithDefaults() *RefTimezoneCode`

NewRefTimezoneCodeWithDefaults instantiates a new RefTimezoneCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezoneCode

`func (o *RefTimezoneCode) GetTimezoneCode() string`

GetTimezoneCode returns the TimezoneCode field if non-nil, zero value otherwise.

### GetTimezoneCodeOk

`func (o *RefTimezoneCode) GetTimezoneCodeOk() (*string, bool)`

GetTimezoneCodeOk returns a tuple with the TimezoneCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneCode

`func (o *RefTimezoneCode) SetTimezoneCode(v string)`

SetTimezoneCode sets TimezoneCode field to given value.


### GetTimezone

`func (o *RefTimezoneCode) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *RefTimezoneCode) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *RefTimezoneCode) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetStandardUtcOffsetMinutes

`func (o *RefTimezoneCode) GetStandardUtcOffsetMinutes() int32`

GetStandardUtcOffsetMinutes returns the StandardUtcOffsetMinutes field if non-nil, zero value otherwise.

### GetStandardUtcOffsetMinutesOk

`func (o *RefTimezoneCode) GetStandardUtcOffsetMinutesOk() (*int32, bool)`

GetStandardUtcOffsetMinutesOk returns a tuple with the StandardUtcOffsetMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardUtcOffsetMinutes

`func (o *RefTimezoneCode) SetStandardUtcOffsetMinutes(v int32)`

SetStandardUtcOffsetMinutes sets StandardUtcOffsetMinutes field to given value.


### GetObservesDst

`func (o *RefTimezoneCode) GetObservesDst() bool`

GetObservesDst returns the ObservesDst field if non-nil, zero value otherwise.

### GetObservesDstOk

`func (o *RefTimezoneCode) GetObservesDstOk() (*bool, bool)`

GetObservesDstOk returns a tuple with the ObservesDst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservesDst

`func (o *RefTimezoneCode) SetObservesDst(v bool)`

SetObservesDst sets ObservesDst field to given value.


### GetDstUtcOffsetMinutes

`func (o *RefTimezoneCode) GetDstUtcOffsetMinutes() int32`

GetDstUtcOffsetMinutes returns the DstUtcOffsetMinutes field if non-nil, zero value otherwise.

### GetDstUtcOffsetMinutesOk

`func (o *RefTimezoneCode) GetDstUtcOffsetMinutesOk() (*int32, bool)`

GetDstUtcOffsetMinutesOk returns a tuple with the DstUtcOffsetMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstUtcOffsetMinutes

`func (o *RefTimezoneCode) SetDstUtcOffsetMinutes(v int32)`

SetDstUtcOffsetMinutes sets DstUtcOffsetMinutes field to given value.

### HasDstUtcOffsetMinutes

`func (o *RefTimezoneCode) HasDstUtcOffsetMinutes() bool`

HasDstUtcOffsetMinutes returns a boolean if a field has been set.

### SetDstUtcOffsetMinutesNil

`func (o *RefTimezoneCode) SetDstUtcOffsetMinutesNil(b bool)`

 SetDstUtcOffsetMinutesNil sets the value for DstUtcOffsetMinutes to be an explicit nil

### UnsetDstUtcOffsetMinutes
`func (o *RefTimezoneCode) UnsetDstUtcOffsetMinutes()`

UnsetDstUtcOffsetMinutes ensures that no value is present for DstUtcOffsetMinutes, not even an explicit nil
### GetDstStartRule

`func (o *RefTimezoneCode) GetDstStartRule() string`

GetDstStartRule returns the DstStartRule field if non-nil, zero value otherwise.

### GetDstStartRuleOk

`func (o *RefTimezoneCode) GetDstStartRuleOk() (*string, bool)`

GetDstStartRuleOk returns a tuple with the DstStartRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstStartRule

`func (o *RefTimezoneCode) SetDstStartRule(v string)`

SetDstStartRule sets DstStartRule field to given value.

### HasDstStartRule

`func (o *RefTimezoneCode) HasDstStartRule() bool`

HasDstStartRule returns a boolean if a field has been set.

### SetDstStartRuleNil

`func (o *RefTimezoneCode) SetDstStartRuleNil(b bool)`

 SetDstStartRuleNil sets the value for DstStartRule to be an explicit nil

### UnsetDstStartRule
`func (o *RefTimezoneCode) UnsetDstStartRule()`

UnsetDstStartRule ensures that no value is present for DstStartRule, not even an explicit nil
### GetDstEndRule

`func (o *RefTimezoneCode) GetDstEndRule() string`

GetDstEndRule returns the DstEndRule field if non-nil, zero value otherwise.

### GetDstEndRuleOk

`func (o *RefTimezoneCode) GetDstEndRuleOk() (*string, bool)`

GetDstEndRuleOk returns a tuple with the DstEndRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstEndRule

`func (o *RefTimezoneCode) SetDstEndRule(v string)`

SetDstEndRule sets DstEndRule field to given value.

### HasDstEndRule

`func (o *RefTimezoneCode) HasDstEndRule() bool`

HasDstEndRule returns a boolean if a field has been set.

### SetDstEndRuleNil

`func (o *RefTimezoneCode) SetDstEndRuleNil(b bool)`

 SetDstEndRuleNil sets the value for DstEndRule to be an explicit nil

### UnsetDstEndRule
`func (o *RefTimezoneCode) UnsetDstEndRule()`

UnsetDstEndRule ensures that no value is present for DstEndRule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


