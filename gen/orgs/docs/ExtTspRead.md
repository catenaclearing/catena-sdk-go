# ExtTspRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique external TSP identifier | 
**Name** | **string** | Company name as listed in the registry | 
**Registry** | [**ExtTspRegistryEnum**](ExtTspRegistryEnum.md) | The regulatory body that certified this provider | 
**Email** | **NullableString** |  | 
**Phone** | **NullableString** |  | 
**Website** | **NullableString** |  | 
**CreatedAt** | **time.Time** | When the external TSP was created in our system | 
**UpdatedAt** | **time.Time** | Last modification timestamp for the external TSP record in our system | 
**IntegratedTspIds** | Pointer to **[]string** | List of TSP IDs that are integrated for this external TSP. Empty if not yet integrated. | [optional] 

## Methods

### NewExtTspRead

`func NewExtTspRead(id string, name string, registry ExtTspRegistryEnum, email NullableString, phone NullableString, website NullableString, createdAt time.Time, updatedAt time.Time, ) *ExtTspRead`

NewExtTspRead instantiates a new ExtTspRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtTspReadWithDefaults

`func NewExtTspReadWithDefaults() *ExtTspRead`

NewExtTspReadWithDefaults instantiates a new ExtTspRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtTspRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtTspRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtTspRead) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ExtTspRead) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtTspRead) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtTspRead) SetName(v string)`

SetName sets Name field to given value.


### GetRegistry

`func (o *ExtTspRead) GetRegistry() ExtTspRegistryEnum`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *ExtTspRead) GetRegistryOk() (*ExtTspRegistryEnum, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *ExtTspRead) SetRegistry(v ExtTspRegistryEnum)`

SetRegistry sets Registry field to given value.


### GetEmail

`func (o *ExtTspRead) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ExtTspRead) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ExtTspRead) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *ExtTspRead) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ExtTspRead) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *ExtTspRead) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ExtTspRead) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ExtTspRead) SetPhone(v string)`

SetPhone sets Phone field to given value.


### SetPhoneNil

`func (o *ExtTspRead) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *ExtTspRead) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetWebsite

`func (o *ExtTspRead) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ExtTspRead) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ExtTspRead) SetWebsite(v string)`

SetWebsite sets Website field to given value.


### SetWebsiteNil

`func (o *ExtTspRead) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *ExtTspRead) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetCreatedAt

`func (o *ExtTspRead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExtTspRead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExtTspRead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ExtTspRead) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExtTspRead) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExtTspRead) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIntegratedTspIds

`func (o *ExtTspRead) GetIntegratedTspIds() []string`

GetIntegratedTspIds returns the IntegratedTspIds field if non-nil, zero value otherwise.

### GetIntegratedTspIdsOk

`func (o *ExtTspRead) GetIntegratedTspIdsOk() (*[]string, bool)`

GetIntegratedTspIdsOk returns a tuple with the IntegratedTspIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegratedTspIds

`func (o *ExtTspRead) SetIntegratedTspIds(v []string)`

SetIntegratedTspIds sets IntegratedTspIds field to given value.

### HasIntegratedTspIds

`func (o *ExtTspRead) HasIntegratedTspIds() bool`

HasIntegratedTspIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


