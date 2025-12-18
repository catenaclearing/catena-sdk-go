# GeotabCredsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Password** | **string** |  | 
**Database** | **string** |  | 
**SessionId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGeotabCredsInput

`func NewGeotabCredsInput(username string, password string, database string, ) *GeotabCredsInput`

NewGeotabCredsInput instantiates a new GeotabCredsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeotabCredsInputWithDefaults

`func NewGeotabCredsInputWithDefaults() *GeotabCredsInput`

NewGeotabCredsInputWithDefaults instantiates a new GeotabCredsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *GeotabCredsInput) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GeotabCredsInput) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GeotabCredsInput) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *GeotabCredsInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GeotabCredsInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GeotabCredsInput) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDatabase

`func (o *GeotabCredsInput) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *GeotabCredsInput) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *GeotabCredsInput) SetDatabase(v string)`

SetDatabase sets Database field to given value.


### GetSessionId

`func (o *GeotabCredsInput) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *GeotabCredsInput) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *GeotabCredsInput) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *GeotabCredsInput) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### SetSessionIdNil

`func (o *GeotabCredsInput) SetSessionIdNil(b bool)`

 SetSessionIdNil sets the value for SessionId to be an explicit nil

### UnsetSessionId
`func (o *GeotabCredsInput) UnsetSessionId()`

UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


