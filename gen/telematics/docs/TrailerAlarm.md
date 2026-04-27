# TrailerAlarm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to [**NullableReeferAlarmCodeEnum**](ReeferAlarmCodeEnum.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTrailerAlarm

`func NewTrailerAlarm() *TrailerAlarm`

NewTrailerAlarm instantiates a new TrailerAlarm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailerAlarmWithDefaults

`func NewTrailerAlarmWithDefaults() *TrailerAlarm`

NewTrailerAlarmWithDefaults instantiates a new TrailerAlarm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TrailerAlarm) GetCode() ReeferAlarmCodeEnum`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TrailerAlarm) GetCodeOk() (*ReeferAlarmCodeEnum, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TrailerAlarm) SetCode(v ReeferAlarmCodeEnum)`

SetCode sets Code field to given value.

### HasCode

`func (o *TrailerAlarm) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *TrailerAlarm) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *TrailerAlarm) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDescription

`func (o *TrailerAlarm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TrailerAlarm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TrailerAlarm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TrailerAlarm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *TrailerAlarm) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *TrailerAlarm) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSeverity

`func (o *TrailerAlarm) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *TrailerAlarm) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *TrailerAlarm) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *TrailerAlarm) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### SetSeverityNil

`func (o *TrailerAlarm) SetSeverityNil(b bool)`

 SetSeverityNil sets the value for Severity to be an explicit nil

### UnsetSeverity
`func (o *TrailerAlarm) UnsetSeverity()`

UnsetSeverity ensures that no value is present for Severity, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


