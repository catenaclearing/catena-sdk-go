# ConnectionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TspId** | **string** | The ID of the Telematics Service Provider (TSP) that the fleet uses. | 
**Credentials** | [**Credentials**](Credentials.md) |  | 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewConnectionCreate

`func NewConnectionCreate(tspId string, credentials Credentials, ) *ConnectionCreate`

NewConnectionCreate instantiates a new ConnectionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionCreateWithDefaults

`func NewConnectionCreateWithDefaults() *ConnectionCreate`

NewConnectionCreateWithDefaults instantiates a new ConnectionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTspId

`func (o *ConnectionCreate) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *ConnectionCreate) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *ConnectionCreate) SetTspId(v string)`

SetTspId sets TspId field to given value.


### GetCredentials

`func (o *ConnectionCreate) GetCredentials() Credentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ConnectionCreate) GetCredentialsOk() (*Credentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ConnectionCreate) SetCredentials(v Credentials)`

SetCredentials sets Credentials field to given value.


### GetDescription

`func (o *ConnectionCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectionCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectionCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectionCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ConnectionCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConnectionCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


