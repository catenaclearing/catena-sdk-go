# Credentials1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **interface{}** |  | 
**TokenType** | Pointer to **string** |  | [optional] 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**RefreshToken** | Pointer to **interface{}** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**AccountId** | **interface{}** |  | 
**Host** | **interface{}** |  | 
**Username** | **interface{}** |  | 
**Password** | **interface{}** |  | 
**ApiKey** | **interface{}** |  | 
**Url** | **string** |  | 
**AuthCode** | **interface{}** |  | 
**Token** | **interface{}** |  | 
**RedirectUri** | **string** |  | 
**ClientId** | **string** |  | 
**ClientSecret** | **interface{}** |  | 
**CompanyId** | **interface{}** |  | 
**Drivername** | [**DatabaseDriverEnum**](DatabaseDriverEnum.md) |  | 
**Port** | **int32** |  | 
**Database** | **string** |  | 
**CarrierIdentifier** | **string** |  | 
**CarrierIdentifierType** | **string** |  | 
**AppId** | **string** |  | 
**AppKey** | **string** |  | 
**ClientKey** | **string** |  | 
**SecretKey** | **interface{}** |  | 
**SessionId** | Pointer to **string** |  | [optional] 
**ResourceOwnerId** | **string** |  | 
**ResourceOwnerSecret** | **interface{}** |  | 
**SignatureMethod** | **string** |  | 
**Realm** | **string** |  | 
**TokenUrl** | **string** |  | 
**Code** | **interface{}** |  | 
**GrantType** | Pointer to **string** |  | [optional] 
**PrivateKey** | **interface{}** |  | 
**ProviderToken** | **interface{}** |  | 
**DotNumber** | **string** |  | 
**ConsumerKey** | **interface{}** |  | 
**AccessKey** | **string** |  | 
**BucketName** | **string** |  | 
**Region** | **string** |  | 
**ApiId** | **string** |  | 
**CarrierId** | **string** |  | 
**RestUsername** | **interface{}** |  | 
**RestPassword** | **interface{}** |  | 
**TkAuth** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewCredentials1

`func NewCredentials1(accessToken interface{}, accountId interface{}, host interface{}, username interface{}, password interface{}, apiKey interface{}, url string, authCode interface{}, token interface{}, redirectUri string, clientId string, clientSecret interface{}, companyId interface{}, drivername DatabaseDriverEnum, port int32, database string, carrierIdentifier string, carrierIdentifierType string, appId string, appKey string, clientKey string, secretKey interface{}, resourceOwnerId string, resourceOwnerSecret interface{}, signatureMethod string, realm string, tokenUrl string, code interface{}, privateKey interface{}, providerToken interface{}, dotNumber string, consumerKey interface{}, accessKey string, bucketName string, region string, apiId string, carrierId string, restUsername interface{}, restPassword interface{}, ) *Credentials1`

NewCredentials1 instantiates a new Credentials1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentials1WithDefaults

`func NewCredentials1WithDefaults() *Credentials1`

NewCredentials1WithDefaults instantiates a new Credentials1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *Credentials1) GetAccessToken() interface{}`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Credentials1) GetAccessTokenOk() (*interface{}, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Credentials1) SetAccessToken(v interface{})`

SetAccessToken sets AccessToken field to given value.


### SetAccessTokenNil

`func (o *Credentials1) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *Credentials1) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetTokenType

`func (o *Credentials1) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *Credentials1) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *Credentials1) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *Credentials1) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetExpiresIn

`func (o *Credentials1) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *Credentials1) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *Credentials1) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *Credentials1) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetRefreshToken

`func (o *Credentials1) GetRefreshToken() interface{}`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *Credentials1) GetRefreshTokenOk() (*interface{}, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *Credentials1) SetRefreshToken(v interface{})`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *Credentials1) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### SetRefreshTokenNil

`func (o *Credentials1) SetRefreshTokenNil(b bool)`

 SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil

### UnsetRefreshToken
`func (o *Credentials1) UnsetRefreshToken()`

UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
### GetScope

`func (o *Credentials1) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Credentials1) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Credentials1) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Credentials1) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetUserId

`func (o *Credentials1) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Credentials1) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Credentials1) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Credentials1) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAccountId

`func (o *Credentials1) GetAccountId() interface{}`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Credentials1) GetAccountIdOk() (*interface{}, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Credentials1) SetAccountId(v interface{})`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *Credentials1) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *Credentials1) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetHost

`func (o *Credentials1) GetHost() interface{}`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Credentials1) GetHostOk() (*interface{}, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Credentials1) SetHost(v interface{})`

SetHost sets Host field to given value.


### SetHostNil

`func (o *Credentials1) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *Credentials1) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetUsername

`func (o *Credentials1) GetUsername() interface{}`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Credentials1) GetUsernameOk() (*interface{}, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Credentials1) SetUsername(v interface{})`

SetUsername sets Username field to given value.


### SetUsernameNil

`func (o *Credentials1) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *Credentials1) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *Credentials1) GetPassword() interface{}`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Credentials1) GetPasswordOk() (*interface{}, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Credentials1) SetPassword(v interface{})`

SetPassword sets Password field to given value.


### SetPasswordNil

`func (o *Credentials1) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *Credentials1) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetApiKey

`func (o *Credentials1) GetApiKey() interface{}`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *Credentials1) GetApiKeyOk() (*interface{}, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *Credentials1) SetApiKey(v interface{})`

SetApiKey sets ApiKey field to given value.


### SetApiKeyNil

`func (o *Credentials1) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *Credentials1) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetUrl

`func (o *Credentials1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Credentials1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Credentials1) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAuthCode

`func (o *Credentials1) GetAuthCode() interface{}`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *Credentials1) GetAuthCodeOk() (*interface{}, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *Credentials1) SetAuthCode(v interface{})`

SetAuthCode sets AuthCode field to given value.


### SetAuthCodeNil

`func (o *Credentials1) SetAuthCodeNil(b bool)`

 SetAuthCodeNil sets the value for AuthCode to be an explicit nil

### UnsetAuthCode
`func (o *Credentials1) UnsetAuthCode()`

UnsetAuthCode ensures that no value is present for AuthCode, not even an explicit nil
### GetToken

`func (o *Credentials1) GetToken() interface{}`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Credentials1) GetTokenOk() (*interface{}, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Credentials1) SetToken(v interface{})`

SetToken sets Token field to given value.


### SetTokenNil

`func (o *Credentials1) SetTokenNil(b bool)`

 SetTokenNil sets the value for Token to be an explicit nil

### UnsetToken
`func (o *Credentials1) UnsetToken()`

UnsetToken ensures that no value is present for Token, not even an explicit nil
### GetRedirectUri

`func (o *Credentials1) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *Credentials1) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *Credentials1) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### GetClientId

`func (o *Credentials1) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Credentials1) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Credentials1) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *Credentials1) GetClientSecret() interface{}`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Credentials1) GetClientSecretOk() (*interface{}, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Credentials1) SetClientSecret(v interface{})`

SetClientSecret sets ClientSecret field to given value.


### SetClientSecretNil

`func (o *Credentials1) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *Credentials1) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetCompanyId

`func (o *Credentials1) GetCompanyId() interface{}`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Credentials1) GetCompanyIdOk() (*interface{}, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Credentials1) SetCompanyId(v interface{})`

SetCompanyId sets CompanyId field to given value.


### SetCompanyIdNil

`func (o *Credentials1) SetCompanyIdNil(b bool)`

 SetCompanyIdNil sets the value for CompanyId to be an explicit nil

### UnsetCompanyId
`func (o *Credentials1) UnsetCompanyId()`

UnsetCompanyId ensures that no value is present for CompanyId, not even an explicit nil
### GetDrivername

`func (o *Credentials1) GetDrivername() DatabaseDriverEnum`

GetDrivername returns the Drivername field if non-nil, zero value otherwise.

### GetDrivernameOk

`func (o *Credentials1) GetDrivernameOk() (*DatabaseDriverEnum, bool)`

GetDrivernameOk returns a tuple with the Drivername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivername

`func (o *Credentials1) SetDrivername(v DatabaseDriverEnum)`

SetDrivername sets Drivername field to given value.


### GetPort

`func (o *Credentials1) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Credentials1) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Credentials1) SetPort(v int32)`

SetPort sets Port field to given value.


### GetDatabase

`func (o *Credentials1) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *Credentials1) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *Credentials1) SetDatabase(v string)`

SetDatabase sets Database field to given value.


### GetCarrierIdentifier

`func (o *Credentials1) GetCarrierIdentifier() string`

GetCarrierIdentifier returns the CarrierIdentifier field if non-nil, zero value otherwise.

### GetCarrierIdentifierOk

`func (o *Credentials1) GetCarrierIdentifierOk() (*string, bool)`

GetCarrierIdentifierOk returns a tuple with the CarrierIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierIdentifier

`func (o *Credentials1) SetCarrierIdentifier(v string)`

SetCarrierIdentifier sets CarrierIdentifier field to given value.


### GetCarrierIdentifierType

`func (o *Credentials1) GetCarrierIdentifierType() string`

GetCarrierIdentifierType returns the CarrierIdentifierType field if non-nil, zero value otherwise.

### GetCarrierIdentifierTypeOk

`func (o *Credentials1) GetCarrierIdentifierTypeOk() (*string, bool)`

GetCarrierIdentifierTypeOk returns a tuple with the CarrierIdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierIdentifierType

`func (o *Credentials1) SetCarrierIdentifierType(v string)`

SetCarrierIdentifierType sets CarrierIdentifierType field to given value.


### GetAppId

`func (o *Credentials1) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Credentials1) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Credentials1) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetAppKey

`func (o *Credentials1) GetAppKey() string`

GetAppKey returns the AppKey field if non-nil, zero value otherwise.

### GetAppKeyOk

`func (o *Credentials1) GetAppKeyOk() (*string, bool)`

GetAppKeyOk returns a tuple with the AppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppKey

`func (o *Credentials1) SetAppKey(v string)`

SetAppKey sets AppKey field to given value.


### GetClientKey

`func (o *Credentials1) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *Credentials1) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *Credentials1) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.


### GetSecretKey

`func (o *Credentials1) GetSecretKey() interface{}`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *Credentials1) GetSecretKeyOk() (*interface{}, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *Credentials1) SetSecretKey(v interface{})`

SetSecretKey sets SecretKey field to given value.


### SetSecretKeyNil

`func (o *Credentials1) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *Credentials1) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
### GetSessionId

`func (o *Credentials1) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *Credentials1) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *Credentials1) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *Credentials1) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetResourceOwnerId

`func (o *Credentials1) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *Credentials1) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *Credentials1) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.


### GetResourceOwnerSecret

`func (o *Credentials1) GetResourceOwnerSecret() interface{}`

GetResourceOwnerSecret returns the ResourceOwnerSecret field if non-nil, zero value otherwise.

### GetResourceOwnerSecretOk

`func (o *Credentials1) GetResourceOwnerSecretOk() (*interface{}, bool)`

GetResourceOwnerSecretOk returns a tuple with the ResourceOwnerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerSecret

`func (o *Credentials1) SetResourceOwnerSecret(v interface{})`

SetResourceOwnerSecret sets ResourceOwnerSecret field to given value.


### SetResourceOwnerSecretNil

`func (o *Credentials1) SetResourceOwnerSecretNil(b bool)`

 SetResourceOwnerSecretNil sets the value for ResourceOwnerSecret to be an explicit nil

### UnsetResourceOwnerSecret
`func (o *Credentials1) UnsetResourceOwnerSecret()`

UnsetResourceOwnerSecret ensures that no value is present for ResourceOwnerSecret, not even an explicit nil
### GetSignatureMethod

`func (o *Credentials1) GetSignatureMethod() string`

GetSignatureMethod returns the SignatureMethod field if non-nil, zero value otherwise.

### GetSignatureMethodOk

`func (o *Credentials1) GetSignatureMethodOk() (*string, bool)`

GetSignatureMethodOk returns a tuple with the SignatureMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureMethod

`func (o *Credentials1) SetSignatureMethod(v string)`

SetSignatureMethod sets SignatureMethod field to given value.


### GetRealm

`func (o *Credentials1) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *Credentials1) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *Credentials1) SetRealm(v string)`

SetRealm sets Realm field to given value.


### GetTokenUrl

`func (o *Credentials1) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *Credentials1) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *Credentials1) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetCode

`func (o *Credentials1) GetCode() interface{}`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Credentials1) GetCodeOk() (*interface{}, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Credentials1) SetCode(v interface{})`

SetCode sets Code field to given value.


### SetCodeNil

`func (o *Credentials1) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *Credentials1) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetGrantType

`func (o *Credentials1) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *Credentials1) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *Credentials1) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *Credentials1) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetPrivateKey

`func (o *Credentials1) GetPrivateKey() interface{}`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *Credentials1) GetPrivateKeyOk() (*interface{}, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *Credentials1) SetPrivateKey(v interface{})`

SetPrivateKey sets PrivateKey field to given value.


### SetPrivateKeyNil

`func (o *Credentials1) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *Credentials1) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetProviderToken

`func (o *Credentials1) GetProviderToken() interface{}`

GetProviderToken returns the ProviderToken field if non-nil, zero value otherwise.

### GetProviderTokenOk

`func (o *Credentials1) GetProviderTokenOk() (*interface{}, bool)`

GetProviderTokenOk returns a tuple with the ProviderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderToken

`func (o *Credentials1) SetProviderToken(v interface{})`

SetProviderToken sets ProviderToken field to given value.


### SetProviderTokenNil

`func (o *Credentials1) SetProviderTokenNil(b bool)`

 SetProviderTokenNil sets the value for ProviderToken to be an explicit nil

### UnsetProviderToken
`func (o *Credentials1) UnsetProviderToken()`

UnsetProviderToken ensures that no value is present for ProviderToken, not even an explicit nil
### GetDotNumber

`func (o *Credentials1) GetDotNumber() string`

GetDotNumber returns the DotNumber field if non-nil, zero value otherwise.

### GetDotNumberOk

`func (o *Credentials1) GetDotNumberOk() (*string, bool)`

GetDotNumberOk returns a tuple with the DotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDotNumber

`func (o *Credentials1) SetDotNumber(v string)`

SetDotNumber sets DotNumber field to given value.


### GetConsumerKey

`func (o *Credentials1) GetConsumerKey() interface{}`

GetConsumerKey returns the ConsumerKey field if non-nil, zero value otherwise.

### GetConsumerKeyOk

`func (o *Credentials1) GetConsumerKeyOk() (*interface{}, bool)`

GetConsumerKeyOk returns a tuple with the ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerKey

`func (o *Credentials1) SetConsumerKey(v interface{})`

SetConsumerKey sets ConsumerKey field to given value.


### SetConsumerKeyNil

`func (o *Credentials1) SetConsumerKeyNil(b bool)`

 SetConsumerKeyNil sets the value for ConsumerKey to be an explicit nil

### UnsetConsumerKey
`func (o *Credentials1) UnsetConsumerKey()`

UnsetConsumerKey ensures that no value is present for ConsumerKey, not even an explicit nil
### GetAccessKey

`func (o *Credentials1) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *Credentials1) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *Credentials1) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetBucketName

`func (o *Credentials1) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *Credentials1) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *Credentials1) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetRegion

`func (o *Credentials1) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Credentials1) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Credentials1) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetApiId

`func (o *Credentials1) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *Credentials1) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *Credentials1) SetApiId(v string)`

SetApiId sets ApiId field to given value.


### GetCarrierId

`func (o *Credentials1) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *Credentials1) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *Credentials1) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.


### GetRestUsername

`func (o *Credentials1) GetRestUsername() interface{}`

GetRestUsername returns the RestUsername field if non-nil, zero value otherwise.

### GetRestUsernameOk

`func (o *Credentials1) GetRestUsernameOk() (*interface{}, bool)`

GetRestUsernameOk returns a tuple with the RestUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestUsername

`func (o *Credentials1) SetRestUsername(v interface{})`

SetRestUsername sets RestUsername field to given value.


### SetRestUsernameNil

`func (o *Credentials1) SetRestUsernameNil(b bool)`

 SetRestUsernameNil sets the value for RestUsername to be an explicit nil

### UnsetRestUsername
`func (o *Credentials1) UnsetRestUsername()`

UnsetRestUsername ensures that no value is present for RestUsername, not even an explicit nil
### GetRestPassword

`func (o *Credentials1) GetRestPassword() interface{}`

GetRestPassword returns the RestPassword field if non-nil, zero value otherwise.

### GetRestPasswordOk

`func (o *Credentials1) GetRestPasswordOk() (*interface{}, bool)`

GetRestPasswordOk returns a tuple with the RestPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestPassword

`func (o *Credentials1) SetRestPassword(v interface{})`

SetRestPassword sets RestPassword field to given value.


### SetRestPasswordNil

`func (o *Credentials1) SetRestPasswordNil(b bool)`

 SetRestPasswordNil sets the value for RestPassword to be an explicit nil

### UnsetRestPassword
`func (o *Credentials1) UnsetRestPassword()`

UnsetRestPassword ensures that no value is present for RestPassword, not even an explicit nil
### GetTkAuth

`func (o *Credentials1) GetTkAuth() interface{}`

GetTkAuth returns the TkAuth field if non-nil, zero value otherwise.

### GetTkAuthOk

`func (o *Credentials1) GetTkAuthOk() (*interface{}, bool)`

GetTkAuthOk returns a tuple with the TkAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTkAuth

`func (o *Credentials1) SetTkAuth(v interface{})`

SetTkAuth sets TkAuth field to given value.

### HasTkAuth

`func (o *Credentials1) HasTkAuth() bool`

HasTkAuth returns a boolean if a field has been set.

### SetTkAuthNil

`func (o *Credentials1) SetTkAuthNil(b bool)`

 SetTkAuthNil sets the value for TkAuth to be an explicit nil

### UnsetTkAuth
`func (o *Credentials1) UnsetTkAuth()`

UnsetTkAuth ensures that no value is present for TkAuth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


