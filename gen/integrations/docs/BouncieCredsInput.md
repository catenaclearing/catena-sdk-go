# BouncieCredsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCode** | **string** |  | 
**RedirectUri** | **string** |  | 
**ClientId** | Pointer to **NullableString** |  | [optional] 
**ClientSecret** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBouncieCredsInput

`func NewBouncieCredsInput(authCode string, redirectUri string, ) *BouncieCredsInput`

NewBouncieCredsInput instantiates a new BouncieCredsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBouncieCredsInputWithDefaults

`func NewBouncieCredsInputWithDefaults() *BouncieCredsInput`

NewBouncieCredsInputWithDefaults instantiates a new BouncieCredsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthCode

`func (o *BouncieCredsInput) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *BouncieCredsInput) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *BouncieCredsInput) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.


### GetRedirectUri

`func (o *BouncieCredsInput) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *BouncieCredsInput) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *BouncieCredsInput) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### GetClientId

`func (o *BouncieCredsInput) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *BouncieCredsInput) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *BouncieCredsInput) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *BouncieCredsInput) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *BouncieCredsInput) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *BouncieCredsInput) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientSecret

`func (o *BouncieCredsInput) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *BouncieCredsInput) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *BouncieCredsInput) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *BouncieCredsInput) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *BouncieCredsInput) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *BouncieCredsInput) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


