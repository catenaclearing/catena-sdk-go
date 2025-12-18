# OAuth2CredsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** |  | 
**ClientSecret** | **interface{}** |  | 
**Scope** | Pointer to **NullableString** |  | [optional] 
**RefreshToken** | Pointer to **interface{}** |  | [optional] 
**TokenUrl** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewOAuth2CredsOutput

`func NewOAuth2CredsOutput(clientId string, clientSecret interface{}, tokenUrl string, url string, ) *OAuth2CredsOutput`

NewOAuth2CredsOutput instantiates a new OAuth2CredsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2CredsOutputWithDefaults

`func NewOAuth2CredsOutputWithDefaults() *OAuth2CredsOutput`

NewOAuth2CredsOutputWithDefaults instantiates a new OAuth2CredsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuth2CredsOutput) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuth2CredsOutput) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuth2CredsOutput) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OAuth2CredsOutput) GetClientSecret() interface{}`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuth2CredsOutput) GetClientSecretOk() (*interface{}, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuth2CredsOutput) SetClientSecret(v interface{})`

SetClientSecret sets ClientSecret field to given value.


### SetClientSecretNil

`func (o *OAuth2CredsOutput) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *OAuth2CredsOutput) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetScope

`func (o *OAuth2CredsOutput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *OAuth2CredsOutput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *OAuth2CredsOutput) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *OAuth2CredsOutput) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *OAuth2CredsOutput) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *OAuth2CredsOutput) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetRefreshToken

`func (o *OAuth2CredsOutput) GetRefreshToken() interface{}`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *OAuth2CredsOutput) GetRefreshTokenOk() (*interface{}, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *OAuth2CredsOutput) SetRefreshToken(v interface{})`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *OAuth2CredsOutput) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *OAuth2CredsOutput) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *OAuth2CredsOutput) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetTokenUrl

`func (o *OAuth2CredsOutput) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *OAuth2CredsOutput) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *OAuth2CredsOutput) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetUrl

`func (o *OAuth2CredsOutput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OAuth2CredsOutput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OAuth2CredsOutput) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


