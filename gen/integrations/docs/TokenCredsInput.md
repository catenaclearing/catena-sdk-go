# TokenCredsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewTokenCredsInput

`func NewTokenCredsInput(token string, ) *TokenCredsInput`

NewTokenCredsInput instantiates a new TokenCredsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCredsInputWithDefaults

`func NewTokenCredsInputWithDefaults() *TokenCredsInput`

NewTokenCredsInputWithDefaults instantiates a new TokenCredsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *TokenCredsInput) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenCredsInput) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenCredsInput) SetToken(v string)`

SetToken sets Token field to given value.


### GetUrl

`func (o *TokenCredsInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TokenCredsInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TokenCredsInput) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TokenCredsInput) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *TokenCredsInput) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *TokenCredsInput) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


