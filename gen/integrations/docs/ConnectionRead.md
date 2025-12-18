# ConnectionRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the connection. | 
**CreatedAt** | **time.Time** | Timestamp when the connection was created. | 
**UpdatedAt** | **time.Time** | Timestamp when the connection was last updated. | 
**FleetId** | **string** | The Catena ID of the fleet that owns this connection. | 
**TspId** | **string** | The ID of the Telematics Service Provider (TSP). | 
**SourceName** | [**TspEnum**](TspEnum.md) | The name of the TSP integration used for this connection. | 
**Credentials** | [**Credentials1**](Credentials1.md) |  | 
**Status** | [**StatusEnum**](StatusEnum.md) | The current status of the connection. | 
**Description** | **NullableString** |  | 

## Methods

### NewConnectionRead

`func NewConnectionRead(id string, createdAt time.Time, updatedAt time.Time, fleetId string, tspId string, sourceName TspEnum, credentials Credentials1, status StatusEnum, description NullableString, ) *ConnectionRead`

NewConnectionRead instantiates a new ConnectionRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionReadWithDefaults

`func NewConnectionReadWithDefaults() *ConnectionRead`

NewConnectionReadWithDefaults instantiates a new ConnectionRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionRead) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ConnectionRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConnectionRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConnectionRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ConnectionRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConnectionRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConnectionRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetFleetId

`func (o *ConnectionRead) GetFleetId() string`

GetFleetId returns the FleetId field if non-nil, zero value otherwise.

### GetFleetIdOk

`func (o *ConnectionRead) GetFleetIdOk() (*string, bool)`

GetFleetIdOk returns a tuple with the FleetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFleetId

`func (o *ConnectionRead) SetFleetId(v string)`

SetFleetId sets FleetId field to given value.


### GetTspId

`func (o *ConnectionRead) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *ConnectionRead) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *ConnectionRead) SetTspId(v string)`

SetTspId sets TspId field to given value.


### GetSourceName

`func (o *ConnectionRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *ConnectionRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *ConnectionRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.


### GetCredentials

`func (o *ConnectionRead) GetCredentials() Credentials1`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ConnectionRead) GetCredentialsOk() (*Credentials1, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ConnectionRead) SetCredentials(v Credentials1)`

SetCredentials sets Credentials field to given value.


### GetStatus

`func (o *ConnectionRead) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectionRead) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectionRead) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *ConnectionRead) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectionRead) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectionRead) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *ConnectionRead) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ConnectionRead) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


