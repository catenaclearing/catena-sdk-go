# TspIntegrationsRead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TspId** | **string** | Unique identifier for the Telematics Service Provider (TSP). | 
**Slug** | Pointer to **NullableString** |  | [optional] 
**SourceName** | Pointer to [**NullableTspEnum**](TspEnum.md) |  | [optional] 
**Integrations** | [**map[string]IntegrationStatusEnum**](IntegrationStatusEnum.md) | Map of every resource to the current integration implementation status. | 
**ImplementedCount** | **int32** | Count of integrations supported by the TSP and implemented by Catena | 
**NotImplementedCount** | **int32** | Count of integrations supported by the TSP, but not implemented by Catena | 
**NotSupportedCount** | **int32** | Count of integrations not supported by the TSP | 

## Methods

### NewTspIntegrationsRead

`func NewTspIntegrationsRead(tspId string, integrations map[string]IntegrationStatusEnum, implementedCount int32, notImplementedCount int32, notSupportedCount int32, ) *TspIntegrationsRead`

NewTspIntegrationsRead instantiates a new TspIntegrationsRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTspIntegrationsReadWithDefaults

`func NewTspIntegrationsReadWithDefaults() *TspIntegrationsRead`

NewTspIntegrationsReadWithDefaults instantiates a new TspIntegrationsRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTspId

`func (o *TspIntegrationsRead) GetTspId() string`

GetTspId returns the TspId field if non-nil, zero value otherwise.

### GetTspIdOk

`func (o *TspIntegrationsRead) GetTspIdOk() (*string, bool)`

GetTspIdOk returns a tuple with the TspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTspId

`func (o *TspIntegrationsRead) SetTspId(v string)`

SetTspId sets TspId field to given value.


### GetSlug

`func (o *TspIntegrationsRead) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TspIntegrationsRead) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TspIntegrationsRead) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *TspIntegrationsRead) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *TspIntegrationsRead) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *TspIntegrationsRead) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetSourceName

`func (o *TspIntegrationsRead) GetSourceName() TspEnum`

GetSourceName returns the SourceName field if non-nil, zero value otherwise.

### GetSourceNameOk

`func (o *TspIntegrationsRead) GetSourceNameOk() (*TspEnum, bool)`

GetSourceNameOk returns a tuple with the SourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceName

`func (o *TspIntegrationsRead) SetSourceName(v TspEnum)`

SetSourceName sets SourceName field to given value.

### HasSourceName

`func (o *TspIntegrationsRead) HasSourceName() bool`

HasSourceName returns a boolean if a field has been set.

### SetSourceNameNil

`func (o *TspIntegrationsRead) SetSourceNameNil(b bool)`

 SetSourceNameNil sets the value for SourceName to be an explicit nil

### UnsetSourceName
`func (o *TspIntegrationsRead) UnsetSourceName()`

UnsetSourceName ensures that no value is present for SourceName, not even an explicit nil
### GetIntegrations

`func (o *TspIntegrationsRead) GetIntegrations() map[string]IntegrationStatusEnum`

GetIntegrations returns the Integrations field if non-nil, zero value otherwise.

### GetIntegrationsOk

`func (o *TspIntegrationsRead) GetIntegrationsOk() (*map[string]IntegrationStatusEnum, bool)`

GetIntegrationsOk returns a tuple with the Integrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrations

`func (o *TspIntegrationsRead) SetIntegrations(v map[string]IntegrationStatusEnum)`

SetIntegrations sets Integrations field to given value.


### GetImplementedCount

`func (o *TspIntegrationsRead) GetImplementedCount() int32`

GetImplementedCount returns the ImplementedCount field if non-nil, zero value otherwise.

### GetImplementedCountOk

`func (o *TspIntegrationsRead) GetImplementedCountOk() (*int32, bool)`

GetImplementedCountOk returns a tuple with the ImplementedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementedCount

`func (o *TspIntegrationsRead) SetImplementedCount(v int32)`

SetImplementedCount sets ImplementedCount field to given value.


### GetNotImplementedCount

`func (o *TspIntegrationsRead) GetNotImplementedCount() int32`

GetNotImplementedCount returns the NotImplementedCount field if non-nil, zero value otherwise.

### GetNotImplementedCountOk

`func (o *TspIntegrationsRead) GetNotImplementedCountOk() (*int32, bool)`

GetNotImplementedCountOk returns a tuple with the NotImplementedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotImplementedCount

`func (o *TspIntegrationsRead) SetNotImplementedCount(v int32)`

SetNotImplementedCount sets NotImplementedCount field to given value.


### GetNotSupportedCount

`func (o *TspIntegrationsRead) GetNotSupportedCount() int32`

GetNotSupportedCount returns the NotSupportedCount field if non-nil, zero value otherwise.

### GetNotSupportedCountOk

`func (o *TspIntegrationsRead) GetNotSupportedCountOk() (*int32, bool)`

GetNotSupportedCountOk returns a tuple with the NotSupportedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotSupportedCount

`func (o *TspIntegrationsRead) SetNotSupportedCount(v int32)`

SetNotSupportedCount sets NotSupportedCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


