# DatabaseCredsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Drivername** | [**DatabaseDriverEnum**](DatabaseDriverEnum.md) |  | 
**Host** | **interface{}** |  | 
**Port** | **int32** |  | 
**Username** | **interface{}** |  | 
**Password** | **interface{}** |  | 
**Database** | **string** |  | 

## Methods

### NewDatabaseCredsOutput

`func NewDatabaseCredsOutput(drivername DatabaseDriverEnum, host interface{}, port int32, username interface{}, password interface{}, database string, ) *DatabaseCredsOutput`

NewDatabaseCredsOutput instantiates a new DatabaseCredsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseCredsOutputWithDefaults

`func NewDatabaseCredsOutputWithDefaults() *DatabaseCredsOutput`

NewDatabaseCredsOutputWithDefaults instantiates a new DatabaseCredsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrivername

`func (o *DatabaseCredsOutput) GetDrivername() DatabaseDriverEnum`

GetDrivername returns the Drivername field if non-nil, zero value otherwise.

### GetDrivernameOk

`func (o *DatabaseCredsOutput) GetDrivernameOk() (*DatabaseDriverEnum, bool)`

GetDrivernameOk returns a tuple with the Drivername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivername

`func (o *DatabaseCredsOutput) SetDrivername(v DatabaseDriverEnum)`

SetDrivername sets Drivername field to given value.


### GetHost

`func (o *DatabaseCredsOutput) GetHost() interface{}`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DatabaseCredsOutput) GetHostOk() (*interface{}, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DatabaseCredsOutput) SetHost(v interface{})`

SetHost sets Host field to given value.


### SetHostNil

`func (o *DatabaseCredsOutput) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *DatabaseCredsOutput) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *DatabaseCredsOutput) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DatabaseCredsOutput) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DatabaseCredsOutput) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUsername

`func (o *DatabaseCredsOutput) GetUsername() interface{}`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DatabaseCredsOutput) GetUsernameOk() (*interface{}, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DatabaseCredsOutput) SetUsername(v interface{})`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *DatabaseCredsOutput) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *DatabaseCredsOutput) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *DatabaseCredsOutput) GetPassword() interface{}`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DatabaseCredsOutput) GetPasswordOk() (*interface{}, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DatabaseCredsOutput) SetPassword(v interface{})`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *DatabaseCredsOutput) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *DatabaseCredsOutput) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDatabase

`func (o *DatabaseCredsOutput) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *DatabaseCredsOutput) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *DatabaseCredsOutput) SetDatabase(v string)`

SetDatabase sets Database field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


