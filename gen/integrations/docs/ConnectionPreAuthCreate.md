# ConnectionPreAuthCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TspId** | **string** | The ID of the Telematics Service Provider (TSP) that the fleet uses. | 
**PreAuthCredentials** | [**ApiBasicCredsInput**](ApiBasicCredsInput.md) | The pre-authentication credentials required to authenticate with the TSP on behalf of the fleet. | 

## Methods

### NewConnectionPreAuthCreate

`func NewConnectionPreAuthCreate(tspId string, preAuthCredentials ApiBasicCredsInput, ) *ConnectionPreAuthCreate`

NewConnectionPreAuthCreate instantiates a new ConnectionPreAuthCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionPreAuthCreateWithDefaults

`func NewConnectionPreAuthCreateWithDefaults() *ConnectionPreAuthCreate`

NewConnectionPreAuthCreateWithDefaults instantiates a new ConnectionPreAuthCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTspId

`func (o *ConnectionPreAuthCreate) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *ConnectionPreAuthCreate) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *ConnectionPreAuthCreate) SetTspId(v string)`

SetTspId sets TspId field to given value.


### GetPreAuthCredentials

`func (o *ConnectionPreAuthCreate) GetPreAuthCredentials() ApiBasicCredsInput`

GetPreAuthCredentials returns the PreAuthCredentials field if non-nil, zero value otherwise.

### GetPreAuthCredentialsOk

`func (o *ConnectionPreAuthCreate) GetPreAuthCredentialsOk() (*ApiBasicCredsInput, bool)`

GetPreAuthCredentialsOk returns a tuple with the PreAuthCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthCredentials

`func (o *ConnectionPreAuthCreate) SetPreAuthCredentials(v ApiBasicCredsInput)`

SetPreAuthCredentials sets PreAuthCredentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


