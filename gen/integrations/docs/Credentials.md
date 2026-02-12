# Credentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Drivername** | [**DatabaseDriverEnum**](DatabaseDriverEnum.md) |  | 
**Host** | **string** |  | 
**Port** | **int32** |  | 
**Username** | **string** |  | 
**Password** | **string** |  | 
**Database** | **string** |  | 
**ApiKey** | **string** |  | 
**Url** | **string** |  | 
**UserId** | Pointer to **string** |  | [optional] 
**AuthCode** | **string** |  | 
**Token** | **string** |  | 
**RedirectUri** | **string** |  | 
**ClientId** | **string** |  | 
**ClientSecret** | **string** |  | 
**CompanyId** | **string** |  | 
**ResourceOwnerId** | **string** |  | 
**ResourceOwnerSecret** | **string** |  | 
**SignatureMethod** | **string** |  | 
**Realm** | **string** |  | 
**Scope** | Pointer to **string** |  | [optional] 
**RefreshToken** | Pointer to **string** |  | [optional] 
**TokenUrl** | **string** |  | 
**PrivateKey** | **string** |  | 
**ConsumerKey** | **string** |  | 
**AccessKey** | **string** |  | 
**SecretKey** | **string** |  | 
**BucketName** | **string** |  | 
**Region** | **string** |  | 
**AccessToken** | **string** |  | 
**TokenType** | **string** |  | 
**ExpiresIn** | Pointer to **int32** |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**Code** | **string** |  | 
**GrantType** | Pointer to **string** |  | [optional] 
**AppId** | **string** |  | 
**AppKey** | **string** |  | 
**ClientKey** | **string** |  | 
**AccountId** | **string** |  | 
**ProviderToken** | **string** |  | 
**DotNumber** | **string** |  | 
**ApiId** | **string** |  | 
**CarrierId** | **string** |  | 

## Methods

### NewCredentials

`func NewCredentials(drivername DatabaseDriverEnum, host string, port int32, username string, password string, database string, apiKey string, url string, authCode string, token string, redirectUri string, clientId string, clientSecret string, companyId string, resourceOwnerId string, resourceOwnerSecret string, signatureMethod string, realm string, tokenUrl string, privateKey string, consumerKey string, accessKey string, secretKey string, bucketName string, region string, accessToken string, tokenType string, code string, appId string, appKey string, clientKey string, accountId string, providerToken string, dotNumber string, apiId string, carrierId string, ) *Credentials`

NewCredentials instantiates a new Credentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialsWithDefaults

`func NewCredentialsWithDefaults() *Credentials`

NewCredentialsWithDefaults instantiates a new Credentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrivername

`func (o *Credentials) GetDrivername() DatabaseDriverEnum`

GetDrivername returns the Drivername field if non-nil, zero value otherwise.

### GetDrivernameOk

`func (o *Credentials) GetDrivernameOk() (*DatabaseDriverEnum, bool)`

GetDrivernameOk returns a tuple with the Drivername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivername

`func (o *Credentials) SetDrivername(v DatabaseDriverEnum)`

SetDrivername sets Drivername field to given value.


### GetHost

`func (o *Credentials) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Credentials) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Credentials) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *Credentials) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Credentials) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Credentials) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUsername

`func (o *Credentials) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Credentials) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Credentials) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *Credentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Credentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Credentials) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDatabase

`func (o *Credentials) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *Credentials) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *Credentials) SetDatabase(v string)`

SetDatabase sets Database field to given value.


### GetApiKey

`func (o *Credentials) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *Credentials) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *Credentials) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetUrl

`func (o *Credentials) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Credentials) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Credentials) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUserId

`func (o *Credentials) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Credentials) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Credentials) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Credentials) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAuthCode

`func (o *Credentials) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *Credentials) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *Credentials) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.


### GetToken

`func (o *Credentials) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Credentials) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Credentials) SetToken(v string)`

SetToken sets Token field to given value.


### GetRedirectUri

`func (o *Credentials) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *Credentials) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *Credentials) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### GetClientId

`func (o *Credentials) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Credentials) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Credentials) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *Credentials) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Credentials) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Credentials) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetCompanyId

`func (o *Credentials) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Credentials) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Credentials) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetResourceOwnerId

`func (o *Credentials) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *Credentials) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *Credentials) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.


### GetResourceOwnerSecret

`func (o *Credentials) GetResourceOwnerSecret() string`

GetResourceOwnerSecret returns the ResourceOwnerSecret field if non-nil, zero value otherwise.

### GetResourceOwnerSecretOk

`func (o *Credentials) GetResourceOwnerSecretOk() (*string, bool)`

GetResourceOwnerSecretOk returns a tuple with the ResourceOwnerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerSecret

`func (o *Credentials) SetResourceOwnerSecret(v string)`

SetResourceOwnerSecret sets ResourceOwnerSecret field to given value.


### GetSignatureMethod

`func (o *Credentials) GetSignatureMethod() string`

GetSignatureMethod returns the SignatureMethod field if non-nil, zero value otherwise.

### GetSignatureMethodOk

`func (o *Credentials) GetSignatureMethodOk() (*string, bool)`

GetSignatureMethodOk returns a tuple with the SignatureMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureMethod

`func (o *Credentials) SetSignatureMethod(v string)`

SetSignatureMethod sets SignatureMethod field to given value.


### GetRealm

`func (o *Credentials) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *Credentials) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *Credentials) SetRealm(v string)`

SetRealm sets Realm field to given value.


### GetScope

`func (o *Credentials) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Credentials) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Credentials) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Credentials) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetRefreshToken

`func (o *Credentials) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *Credentials) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *Credentials) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *Credentials) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetTokenUrl

`func (o *Credentials) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *Credentials) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *Credentials) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetPrivateKey

`func (o *Credentials) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *Credentials) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *Credentials) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetConsumerKey

`func (o *Credentials) GetConsumerKey() string`

GetConsumerKey returns the ConsumerKey field if non-nil, zero value otherwise.

### GetConsumerKeyOk

`func (o *Credentials) GetConsumerKeyOk() (*string, bool)`

GetConsumerKeyOk returns a tuple with the ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerKey

`func (o *Credentials) SetConsumerKey(v string)`

SetConsumerKey sets ConsumerKey field to given value.


### GetAccessKey

`func (o *Credentials) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *Credentials) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *Credentials) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *Credentials) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *Credentials) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *Credentials) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetBucketName

`func (o *Credentials) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *Credentials) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *Credentials) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetRegion

`func (o *Credentials) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Credentials) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Credentials) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAccessToken

`func (o *Credentials) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Credentials) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Credentials) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTokenType

`func (o *Credentials) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *Credentials) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *Credentials) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *Credentials) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *Credentials) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *Credentials) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *Credentials) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetSessionId

`func (o *Credentials) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *Credentials) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *Credentials) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *Credentials) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetCode

`func (o *Credentials) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Credentials) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Credentials) SetCode(v string)`

SetCode sets Code field to given value.


### GetGrantType

`func (o *Credentials) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *Credentials) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *Credentials) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *Credentials) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetAppId

`func (o *Credentials) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Credentials) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Credentials) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetAppKey

`func (o *Credentials) GetAppKey() string`

GetAppKey returns the AppKey field if non-nil, zero value otherwise.

### GetAppKeyOk

`func (o *Credentials) GetAppKeyOk() (*string, bool)`

GetAppKeyOk returns a tuple with the AppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppKey

`func (o *Credentials) SetAppKey(v string)`

SetAppKey sets AppKey field to given value.


### GetClientKey

`func (o *Credentials) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *Credentials) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *Credentials) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.


### GetAccountId

`func (o *Credentials) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Credentials) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Credentials) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetProviderToken

`func (o *Credentials) GetProviderToken() string`

GetProviderToken returns the ProviderToken field if non-nil, zero value otherwise.

### GetProviderTokenOk

`func (o *Credentials) GetProviderTokenOk() (*string, bool)`

GetProviderTokenOk returns a tuple with the ProviderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderToken

`func (o *Credentials) SetProviderToken(v string)`

SetProviderToken sets ProviderToken field to given value.


### GetDotNumber

`func (o *Credentials) GetDotNumber() string`

GetDotNumber returns the DotNumber field if non-nil, zero value otherwise.

### GetDotNumberOk

`func (o *Credentials) GetDotNumberOk() (*string, bool)`

GetDotNumberOk returns a tuple with the DotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDotNumber

`func (o *Credentials) SetDotNumber(v string)`

SetDotNumber sets DotNumber field to given value.


### GetApiId

`func (o *Credentials) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *Credentials) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *Credentials) SetApiId(v string)`

SetApiId sets ApiId field to given value.


### GetCarrierId

`func (o *Credentials) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *Credentials) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *Credentials) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


