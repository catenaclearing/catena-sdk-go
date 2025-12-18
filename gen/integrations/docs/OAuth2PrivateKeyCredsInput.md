# OAuth2PrivateKeyCredsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**PrivateKey** | **string** |  | 
**Scope** | Pointer to **NullableString** |  | [optional] 
**TokenUrl** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewOAuth2PrivateKeyCredsInput

`func NewOAuth2PrivateKeyCredsInput(clientId string, privateKey string, tokenUrl string, url string, ) *OAuth2PrivateKeyCredsInput`

NewOAuth2PrivateKeyCredsInput instantiates a new OAuth2PrivateKeyCredsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2PrivateKeyCredsInputWithDefaults

`func NewOAuth2PrivateKeyCredsInputWithDefaults() *OAuth2PrivateKeyCredsInput`

NewOAuth2PrivateKeyCredsInputWithDefaults instantiates a new OAuth2PrivateKeyCredsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuth2PrivateKeyCredsInput) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2PrivateKeyCredsInput) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2PrivateKeyCredsInput) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetPrivateKey

`func (o *OAuth2PrivateKeyCredsInput) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *OAuth2PrivateKeyCredsInput) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *OAuth2PrivateKeyCredsInput) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetScope

`func (o *OAuth2PrivateKeyCredsInput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OAuth2PrivateKeyCredsInput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OAuth2PrivateKeyCredsInput) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OAuth2PrivateKeyCredsInput) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *OAuth2PrivateKeyCredsInput) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *OAuth2PrivateKeyCredsInput) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetTokenUrl

`func (o *OAuth2PrivateKeyCredsInput) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *OAuth2PrivateKeyCredsInput) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *OAuth2PrivateKeyCredsInput) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetUrl

`func (o *OAuth2PrivateKeyCredsInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OAuth2PrivateKeyCredsInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OAuth2PrivateKeyCredsInput) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


