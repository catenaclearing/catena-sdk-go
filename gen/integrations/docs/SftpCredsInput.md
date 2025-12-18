# SftpCredsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**Port** | Pointer to **NullableInt32** |  | [optional] 
**Username** | **string** |  | 
**Password** | **string** |  | 

## Methods

### NewSftpCredsInput

`func NewSftpCredsInput(host string, username string, password string, ) *SftpCredsInput`

NewSftpCredsInput instantiates a new SftpCredsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSftpCredsInputWithDefaults

`func NewSftpCredsInputWithDefaults() *SftpCredsInput`

NewSftpCredsInputWithDefaults instantiates a new SftpCredsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SftpCredsInput) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SftpCredsInput) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SftpCredsInput) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *SftpCredsInput) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SftpCredsInput) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SftpCredsInput) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SftpCredsInput) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *SftpCredsInput) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *SftpCredsInput) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUsername

`func (o *SftpCredsInput) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SftpCredsInput) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SftpCredsInput) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *SftpCredsInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SftpCredsInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SftpCredsInput) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


