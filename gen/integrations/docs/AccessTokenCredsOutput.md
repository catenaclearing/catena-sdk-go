# AccessTokenCredsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **interface{}** |  | 
**TokenType** | **string** |  | 
**ExpiresIn** | Pointer to **NullableInt32** |  | [optional] 
**RefreshToken** | Pointer to **interface{}** |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Host** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAccessTokenCredsOutput

`func NewAccessTokenCredsOutput(accessToken interface{}, tokenType string, ) *AccessTokenCredsOutput`

NewAccessTokenCredsOutput instantiates a new AccessTokenCredsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessTokenCredsOutputWithDefaults

`func NewAccessTokenCredsOutputWithDefaults() *AccessTokenCredsOutput`

NewAccessTokenCredsOutputWithDefaults instantiates a new AccessTokenCredsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AccessTokenCredsOutput) GetAccessToken() interface{}`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AccessTokenCredsOutput) GetAccessTokenOk() (*interface{}, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AccessTokenCredsOutput) SetAccessToken(v interface{})`

SetAccessToken sets AccessToken field to given value.


### SetAccessTokenNil

`func (o *AccessTokenCredsOutput) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *AccessTokenCredsOutput) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetTokenType

`func (o *AccessTokenCredsOutput) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *AccessTokenCredsOutput) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *AccessTokenCredsOutput) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *AccessTokenCredsOutput) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *AccessTokenCredsOutput) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *AccessTokenCredsOutput) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *AccessTokenCredsOutput) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### SetExpiresInNil

`func (o *AccessTokenCredsOutput) SetExpiresInNil(b bool)`

 SetExpiresInNil sets the value for ExpiresIn to be an explicit nil

### UnsetExpiresIn
`func (o *AccessTokenCredsOutput) UnsetExpiresIn()`

UnsetExpiresIn ensures that no value is present for ExpiresIn, not even an explicit nil
### GetRefreshToken

`func (o *AccessTokenCredsOutput) GetRefreshToken() interface{}`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *AccessTokenCredsOutput) GetRefreshTokenOk() (*interface{}, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *AccessTokenCredsOutput) SetRefreshToken(v interface{})`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *AccessTokenCredsOutput) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *AccessTokenCredsOutput) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *AccessTokenCredsOutput) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetScope

`func (o *AccessTokenCredsOutput) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AccessTokenCredsOutput) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AccessTokenCredsOutput) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AccessTokenCredsOutput) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *AccessTokenCredsOutput) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *AccessTokenCredsOutput) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetUserId

`func (o *AccessTokenCredsOutput) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AccessTokenCredsOutput) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AccessTokenCredsOutput) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AccessTokenCredsOutput) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *AccessTokenCredsOutput) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *AccessTokenCredsOutput) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetHost

`func (o *AccessTokenCredsOutput) GetHost() interface{}`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AccessTokenCredsOutput) GetHostOk() (*interface{}, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AccessTokenCredsOutput) SetHost(v interface{})`

SetHost sets Host field to given value.

### HasHost

`func (o *AccessTokenCredsOutput) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *AccessTokenCredsOutput) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *AccessTokenCredsOutput) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


