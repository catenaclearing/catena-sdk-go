# BouncieCredsOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCode** | **interface{}** |  | 
**RedirectUri** | **string** |  | 
**ClientId** | Pointer to **NullableString** |  | [optional] 
**ClientSecret** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewBouncieCredsOutput

`func NewBouncieCredsOutput(authCode interface{}, redirectUri string, ) *BouncieCredsOutput`

NewBouncieCredsOutput instantiates a new BouncieCredsOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBouncieCredsOutputWithDefaults

`func NewBouncieCredsOutputWithDefaults() *BouncieCredsOutput`

NewBouncieCredsOutputWithDefaults instantiates a new BouncieCredsOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthCode

`func (o *BouncieCredsOutput) GetAuthCode() interface{}`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *BouncieCredsOutput) GetAuthCodeOk() (*interface{}, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *BouncieCredsOutput) SetAuthCode(v interface{})`

SetAuthCode sets AuthCode field to given value.


### SetAuthCodeNil

`func (o *BouncieCredsOutput) SetAuthCodeNil(b bool)`

 SetAuthCodeNil sets the value for AuthCode to be an explicit nil

### UnsetAuthCode
`func (o *BouncieCredsOutput) UnsetAuthCode()`

UnsetAuthCode ensures that no value is present for AuthCode, not even an explicit nil
### GetRedirectUri

`func (o *BouncieCredsOutput) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *BouncieCredsOutput) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *BouncieCredsOutput) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### GetClientId

`func (o *BouncieCredsOutput) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BouncieCredsOutput) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BouncieCredsOutput) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BouncieCredsOutput) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *BouncieCredsOutput) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *BouncieCredsOutput) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientSecret

`func (o *BouncieCredsOutput) GetClientSecret() interface{}`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *BouncieCredsOutput) GetClientSecretOk() (*interface{}, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *BouncieCredsOutput) SetClientSecret(v interface{})`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *BouncieCredsOutput) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *BouncieCredsOutput) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *BouncieCredsOutput) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


