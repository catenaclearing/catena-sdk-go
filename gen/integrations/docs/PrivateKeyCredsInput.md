# PrivateKeyCredsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | **string** |  | 
**ConsumerKey** | **string** |  | 
**Username** | **string** |  | 
**Url** | Pointer to **NullableString** |  | [optional] 
**TokenUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPrivateKeyCredsInput

`func NewPrivateKeyCredsInput(privateKey string, consumerKey string, username string, ) *PrivateKeyCredsInput`

NewPrivateKeyCredsInput instantiates a new PrivateKeyCredsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateKeyCredsInputWithDefaults

`func NewPrivateKeyCredsInputWithDefaults() *PrivateKeyCredsInput`

NewPrivateKeyCredsInputWithDefaults instantiates a new PrivateKeyCredsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *PrivateKeyCredsInput) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *PrivateKeyCredsInput) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *PrivateKeyCredsInput) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetConsumerKey

`func (o *PrivateKeyCredsInput) GetConsumerKey() string`

GetConsumerKey returns the ConsumerKey field if non-nil, zero value otherwise.

### GetConsumerKeyOk

`func (o *PrivateKeyCredsInput) GetConsumerKeyOk() (*string, bool)`

GetConsumerKeyOk returns a tuple with the ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerKey

`func (o *PrivateKeyCredsInput) SetConsumerKey(v string)`

SetConsumerKey sets ConsumerKey field to given value.


### GetUsername

`func (o *PrivateKeyCredsInput) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PrivateKeyCredsInput) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PrivateKeyCredsInput) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetUrl

`func (o *PrivateKeyCredsInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PrivateKeyCredsInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PrivateKeyCredsInput) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PrivateKeyCredsInput) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PrivateKeyCredsInput) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PrivateKeyCredsInput) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetTokenUrl

`func (o *PrivateKeyCredsInput) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *PrivateKeyCredsInput) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *PrivateKeyCredsInput) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *PrivateKeyCredsInput) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### SetTokenUrlNil

`func (o *PrivateKeyCredsInput) SetTokenUrlNil(b bool)`

 SetTokenUrlNil sets the value for TokenUrl to be an explicit nil

### UnsetTokenUrl
`func (o *PrivateKeyCredsInput) UnsetTokenUrl()`

UnsetTokenUrl ensures that no value is present for TokenUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


