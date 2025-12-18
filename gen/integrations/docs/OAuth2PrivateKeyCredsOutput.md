# OAuth2PrivateKeyCredsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**PrivateKey** | **interface{}** |  | 
**Scope** | Pointer to **NullableString** |  | [optional] 
**TokenUrl** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewOAuth2PrivateKeyCredsOutput

`func NewOAuth2PrivateKeyCredsOutput(clientId string, privateKey interface{}, tokenUrl string, url string, ) *OAuth2PrivateKeyCredsOutput`

NewOAuth2PrivateKeyCredsOutput instantiates a new OAuth2PrivateKeyCredsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2PrivateKeyCredsOutputWithDefaults

`func NewOAuth2PrivateKeyCredsOutputWithDefaults() *OAuth2PrivateKeyCredsOutput`

NewOAuth2PrivateKeyCredsOutputWithDefaults instantiates a new OAuth2PrivateKeyCredsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuth2PrivateKeyCredsOutput) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2PrivateKeyCredsOutput) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2PrivateKeyCredsOutput) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetPrivateKey

`func (o *OAuth2PrivateKeyCredsOutput) GetPrivateKey() interface{}`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *OAuth2PrivateKeyCredsOutput) GetPrivateKeyOk() (*interface{}, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *OAuth2PrivateKeyCredsOutput) SetPrivateKey(v interface{})`

SetPrivateKey sets PrivateKey field to given value.


### SetPrivateKeyNil

`func (o *OAuth2PrivateKeyCredsOutput) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *OAuth2PrivateKeyCredsOutput) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetScope

`func (o *OAuth2PrivateKeyCredsOutput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OAuth2PrivateKeyCredsOutput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OAuth2PrivateKeyCredsOutput) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OAuth2PrivateKeyCredsOutput) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *OAuth2PrivateKeyCredsOutput) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *OAuth2PrivateKeyCredsOutput) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetTokenUrl

`func (o *OAuth2PrivateKeyCredsOutput) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *OAuth2PrivateKeyCredsOutput) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *OAuth2PrivateKeyCredsOutput) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetUrl

`func (o *OAuth2PrivateKeyCredsOutput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OAuth2PrivateKeyCredsOutput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OAuth2PrivateKeyCredsOutput) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


