# SftpCredsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **interface{}** |  | 
**Port** | Pointer to **NullableInt32** |  | [optional] 
**Username** | **interface{}** |  | 
**Password** | **interface{}** |  | 

## Methods

### NewSftpCredsOutput

`func NewSftpCredsOutput(host interface{}, username interface{}, password interface{}, ) *SftpCredsOutput`

NewSftpCredsOutput instantiates a new SftpCredsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSftpCredsOutputWithDefaults

`func NewSftpCredsOutputWithDefaults() *SftpCredsOutput`

NewSftpCredsOutputWithDefaults instantiates a new SftpCredsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SftpCredsOutput) GetHost() interface{}`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SftpCredsOutput) GetHostOk() (*interface{}, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SftpCredsOutput) SetHost(v interface{})`

SetHost sets Host field to given value.


### SetHostNil

`func (o *SftpCredsOutput) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *SftpCredsOutput) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *SftpCredsOutput) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SftpCredsOutput) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SftpCredsOutput) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SftpCredsOutput) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *SftpCredsOutput) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *SftpCredsOutput) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUsername

`func (o *SftpCredsOutput) GetUsername() interface{}`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SftpCredsOutput) GetUsernameOk() (*interface{}, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SftpCredsOutput) SetUsername(v interface{})`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *SftpCredsOutput) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *SftpCredsOutput) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *SftpCredsOutput) GetPassword() interface{}`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SftpCredsOutput) GetPasswordOk() (*interface{}, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SftpCredsOutput) SetPassword(v interface{})`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *SftpCredsOutput) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *SftpCredsOutput) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


