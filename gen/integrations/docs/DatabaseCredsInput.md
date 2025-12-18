# DatabaseCredsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Drivername** | [**DatabaseDriverEnum**](DatabaseDriverEnum.md) |  | 
**Host** | **string** |  | 
**Port** | **int32** |  | 
**Username** | **string** |  | 
**Password** | **string** |  | 
**Database** | **string** |  | 

## Methods

### NewDatabaseCredsInput

`func NewDatabaseCredsInput(drivername DatabaseDriverEnum, host string, port int32, username string, password string, database string, ) *DatabaseCredsInput`

NewDatabaseCredsInput instantiates a new DatabaseCredsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseCredsInputWithDefaults

`func NewDatabaseCredsInputWithDefaults() *DatabaseCredsInput`

NewDatabaseCredsInputWithDefaults instantiates a new DatabaseCredsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrivername

`func (o *DatabaseCredsInput) GetDrivername() DatabaseDriverEnum`

GetDrivername returns the Drivername field if non-nil, zero value otherwise.

### GetDrivernameOk

`func (o *DatabaseCredsInput) GetDrivernameOk() (*DatabaseDriverEnum, bool)`

GetDrivernameOk returns a tuple with the Drivername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivername

`func (o *DatabaseCredsInput) SetDrivername(v DatabaseDriverEnum)`

SetDrivername sets Drivername field to given value.


### GetHost

`func (o *DatabaseCredsInput) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *DatabaseCredsInput) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *DatabaseCredsInput) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *DatabaseCredsInput) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DatabaseCredsInput) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DatabaseCredsInput) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUsername

`func (o *DatabaseCredsInput) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DatabaseCredsInput) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DatabaseCredsInput) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *DatabaseCredsInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DatabaseCredsInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DatabaseCredsInput) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDatabase

`func (o *DatabaseCredsInput) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *DatabaseCredsInput) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *DatabaseCredsInput) SetDatabase(v string)`

SetDatabase sets Database field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


