# ConnectionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**NullableCredentials2**](Credentials2.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullableStatusEnum**](StatusEnum.md) |  | [optional] 

## Methods

### NewConnectionUpdate

`func NewConnectionUpdate() *ConnectionUpdate`

NewConnectionUpdate instantiates a new ConnectionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionUpdateWithDefaults

`func NewConnectionUpdateWithDefaults() *ConnectionUpdate`

NewConnectionUpdateWithDefaults instantiates a new ConnectionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ConnectionUpdate) GetCredentials() Credentials2`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ConnectionUpdate) GetCredentialsOk() (*Credentials2, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ConnectionUpdate) SetCredentials(v Credentials2)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ConnectionUpdate) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *ConnectionUpdate) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *ConnectionUpdate) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetDescription

`func (o *ConnectionUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectionUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectionUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectionUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConnectionUpdate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConnectionUpdate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStatus

`func (o *ConnectionUpdate) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectionUpdate) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectionUpdate) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectionUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ConnectionUpdate) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ConnectionUpdate) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


