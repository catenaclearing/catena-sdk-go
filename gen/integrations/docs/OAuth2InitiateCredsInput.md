# OAuth2InitiateCredsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**GrantType** | Pointer to **NullableString** |  | [optional] 
**TokenUrl** | Pointer to **NullableString** |  | [optional] 
**RedirectUri** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOAuth2InitiateCredsInput

`func NewOAuth2InitiateCredsInput(code string, ) *OAuth2InitiateCredsInput`

NewOAuth2InitiateCredsInput instantiates a new OAuth2InitiateCredsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2InitiateCredsInputWithDefaults

`func NewOAuth2InitiateCredsInputWithDefaults() *OAuth2InitiateCredsInput`

NewOAuth2InitiateCredsInputWithDefaults instantiates a new OAuth2InitiateCredsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *OAuth2InitiateCredsInput) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OAuth2InitiateCredsInput) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OAuth2InitiateCredsInput) SetCode(v string)`

SetCode sets Code field to given value.


### GetGrantType

`func (o *OAuth2InitiateCredsInput) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *OAuth2InitiateCredsInput) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *OAuth2InitiateCredsInput) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *OAuth2InitiateCredsInput) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### SetGrantTypeNil

`func (o *OAuth2InitiateCredsInput) SetGrantTypeNil(b bool)`

 SetGrantTypeNil sets the value for GrantType to be an explicit nil

### UnsetGrantType
`func (o *OAuth2InitiateCredsInput) UnsetGrantType()`

UnsetGrantType ensures that no value is present for GrantType, not even an explicit nil
### GetTokenUrl

`func (o *OAuth2InitiateCredsInput) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *OAuth2InitiateCredsInput) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *OAuth2InitiateCredsInput) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *OAuth2InitiateCredsInput) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### SetTokenUrlNil

`func (o *OAuth2InitiateCredsInput) SetTokenUrlNil(b bool)`

 SetTokenUrlNil sets the value for TokenUrl to be an explicit nil

### UnsetTokenUrl
`func (o *OAuth2InitiateCredsInput) UnsetTokenUrl()`

UnsetTokenUrl ensures that no value is present for TokenUrl, not even an explicit nil
### GetRedirectUri

`func (o *OAuth2InitiateCredsInput) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *OAuth2InitiateCredsInput) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *OAuth2InitiateCredsInput) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *OAuth2InitiateCredsInput) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### SetRedirectUriNil

`func (o *OAuth2InitiateCredsInput) SetRedirectUriNil(b bool)`

 SetRedirectUriNil sets the value for RedirectUri to be an explicit nil

### UnsetRedirectUri
`func (o *OAuth2InitiateCredsInput) UnsetRedirectUri()`

UnsetRedirectUri ensures that no value is present for RedirectUri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


