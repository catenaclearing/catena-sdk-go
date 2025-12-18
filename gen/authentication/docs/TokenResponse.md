# TokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** |  | 
**ExpiresIn** | **int32** |  | 
**RefreshExpiresIn** | Pointer to **int32** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**TokenType** | **string** |  | 
**SessionState** | Pointer to **string** |  | [optional] 
**NotBeforePolicy** | Pointer to **int32** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 

## Methods

### NewTokenResponse

`func NewTokenResponse(accessToken string, expiresIn int32, tokenType string, ) *TokenResponse`

NewTokenResponse instantiates a new TokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenResponseWithDefaults

`func NewTokenResponseWithDefaults() *TokenResponse`

NewTokenResponseWithDefaults instantiates a new TokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *TokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetExpiresIn

`func (o *TokenResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *TokenResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *TokenResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.


### GetRefreshExpiresIn

`func (o *TokenResponse) GetRefreshExpiresIn() int32`

GetRefreshExpiresIn returns the RefreshExpiresIn field if non-nil, zero value otherwise.

### GetRefreshExpiresInOk

`func (o *TokenResponse) GetRefreshExpiresInOk() (*int32, bool)`

GetRefreshExpiresInOk returns a tuple with the RefreshExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshExpiresIn

`func (o *TokenResponse) SetRefreshExpiresIn(v int32)`

SetRefreshExpiresIn sets RefreshExpiresIn field to given value.

### HasRefreshExpiresIn

`func (o *TokenResponse) HasRefreshExpiresIn() bool`

HasRefreshExpiresIn returns a boolean if a field has been set.

### GetRefreshToken

`func (o *TokenResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *TokenResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *TokenResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *TokenResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetTokenType

`func (o *TokenResponse) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TokenResponse) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TokenResponse) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetSessionState

`func (o *TokenResponse) GetSessionState() string`

GetSessionState returns the SessionState field if non-nil, zero value otherwise.

### GetSessionStateOk

`func (o *TokenResponse) GetSessionStateOk() (*string, bool)`

GetSessionStateOk returns a tuple with the SessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionState

`func (o *TokenResponse) SetSessionState(v string)`

SetSessionState sets SessionState field to given value.

### HasSessionState

`func (o *TokenResponse) HasSessionState() bool`

HasSessionState returns a boolean if a field has been set.

### GetNotBeforePolicy

`func (o *TokenResponse) GetNotBeforePolicy() int32`

GetNotBeforePolicy returns the NotBeforePolicy field if non-nil, zero value otherwise.

### GetNotBeforePolicyOk

`func (o *TokenResponse) GetNotBeforePolicyOk() (*int32, bool)`

GetNotBeforePolicyOk returns a tuple with the NotBeforePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforePolicy

`func (o *TokenResponse) SetNotBeforePolicy(v int32)`

SetNotBeforePolicy sets NotBeforePolicy field to given value.

### HasNotBeforePolicy

`func (o *TokenResponse) HasNotBeforePolicy() bool`

HasNotBeforePolicy returns a boolean if a field has been set.

### GetScope

`func (o *TokenResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *TokenResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *TokenResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *TokenResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


