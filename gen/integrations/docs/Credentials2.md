# Credentials2

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
**CarrierId** | **string** |  | 

## Methods

### NewCredentials2

`func NewCredentials2(drivername DatabaseDriverEnum, host string, port int32, username string, password string, database string, apiKey string, url string, authCode string, token string, redirectUri string, clientId string, clientSecret string, companyId string, resourceOwnerId string, resourceOwnerSecret string, signatureMethod string, realm string, tokenUrl string, privateKey string, consumerKey string, accessKey string, secretKey string, bucketName string, region string, accessToken string, tokenType string, code string, appId string, appKey string, clientKey string, accountId string, providerToken string, dotNumber string, carrierId string, ) *Credentials2`

NewCredentials2 instantiates a new Credentials2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentials2WithDefaults

`func NewCredentials2WithDefaults() *Credentials2`

NewCredentials2WithDefaults instantiates a new Credentials2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrivername

`func (o *Credentials2) GetDrivername() DatabaseDriverEnum`

GetDrivername returns the Drivername field if non-nil, zero value otherwise.

### GetDrivernameOk

`func (o *Credentials2) GetDrivernameOk() (*DatabaseDriverEnum, bool)`

GetDrivernameOk returns a tuple with the Drivername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrivername

`func (o *Credentials2) SetDrivername(v DatabaseDriverEnum)`

SetDrivername sets Drivername field to given value.


### GetHost

`func (o *Credentials2) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *Credentials2) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *Credentials2) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *Credentials2) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Credentials2) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Credentials2) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUsername

`func (o *Credentials2) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Credentials2) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Credentials2) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *Credentials2) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Credentials2) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Credentials2) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetDatabase

`func (o *Credentials2) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *Credentials2) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *Credentials2) SetDatabase(v string)`

SetDatabase sets Database field to given value.


### GetApiKey

`func (o *Credentials2) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *Credentials2) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *Credentials2) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetUrl

`func (o *Credentials2) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Credentials2) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Credentials2) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUserId

`func (o *Credentials2) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Credentials2) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Credentials2) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Credentials2) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAuthCode

`func (o *Credentials2) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *Credentials2) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *Credentials2) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.


### GetToken

`func (o *Credentials2) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Credentials2) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Credentials2) SetToken(v string)`

SetToken sets Token field to given value.


### GetRedirectUri

`func (o *Credentials2) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *Credentials2) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *Credentials2) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### GetClientId

`func (o *Credentials2) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Credentials2) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Credentials2) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *Credentials2) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Credentials2) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Credentials2) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetCompanyId

`func (o *Credentials2) GetCompanyId() string`

GetCompanyId returns the CompanyId field if non-nil, zero value otherwise.

### GetCompanyIdOk

`func (o *Credentials2) GetCompanyIdOk() (*string, bool)`

GetCompanyIdOk returns a tuple with the CompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyId

`func (o *Credentials2) SetCompanyId(v string)`

SetCompanyId sets CompanyId field to given value.


### GetResourceOwnerId

`func (o *Credentials2) GetResourceOwnerId() string`

GetResourceOwnerId returns the ResourceOwnerId field if non-nil, zero value otherwise.

### GetResourceOwnerIdOk

`func (o *Credentials2) GetResourceOwnerIdOk() (*string, bool)`

GetResourceOwnerIdOk returns a tuple with the ResourceOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerId

`func (o *Credentials2) SetResourceOwnerId(v string)`

SetResourceOwnerId sets ResourceOwnerId field to given value.


### GetResourceOwnerSecret

`func (o *Credentials2) GetResourceOwnerSecret() string`

GetResourceOwnerSecret returns the ResourceOwnerSecret field if non-nil, zero value otherwise.

### GetResourceOwnerSecretOk

`func (o *Credentials2) GetResourceOwnerSecretOk() (*string, bool)`

GetResourceOwnerSecretOk returns a tuple with the ResourceOwnerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOwnerSecret

`func (o *Credentials2) SetResourceOwnerSecret(v string)`

SetResourceOwnerSecret sets ResourceOwnerSecret field to given value.


### GetSignatureMethod

`func (o *Credentials2) GetSignatureMethod() string`

GetSignatureMethod returns the SignatureMethod field if non-nil, zero value otherwise.

### GetSignatureMethodOk

`func (o *Credentials2) GetSignatureMethodOk() (*string, bool)`

GetSignatureMethodOk returns a tuple with the SignatureMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureMethod

`func (o *Credentials2) SetSignatureMethod(v string)`

SetSignatureMethod sets SignatureMethod field to given value.


### GetRealm

`func (o *Credentials2) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *Credentials2) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *Credentials2) SetRealm(v string)`

SetRealm sets Realm field to given value.


### GetScope

`func (o *Credentials2) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Credentials2) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Credentials2) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Credentials2) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetRefreshToken

`func (o *Credentials2) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *Credentials2) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *Credentials2) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *Credentials2) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetTokenUrl

`func (o *Credentials2) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *Credentials2) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *Credentials2) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.


### GetPrivateKey

`func (o *Credentials2) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *Credentials2) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *Credentials2) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetConsumerKey

`func (o *Credentials2) GetConsumerKey() string`

GetConsumerKey returns the ConsumerKey field if non-nil, zero value otherwise.

### GetConsumerKeyOk

`func (o *Credentials2) GetConsumerKeyOk() (*string, bool)`

GetConsumerKeyOk returns a tuple with the ConsumerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumerKey

`func (o *Credentials2) SetConsumerKey(v string)`

SetConsumerKey sets ConsumerKey field to given value.


### GetAccessKey

`func (o *Credentials2) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *Credentials2) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *Credentials2) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.


### GetSecretKey

`func (o *Credentials2) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *Credentials2) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *Credentials2) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.


### GetBucketName

`func (o *Credentials2) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *Credentials2) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *Credentials2) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetRegion

`func (o *Credentials2) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Credentials2) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Credentials2) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAccessToken

`func (o *Credentials2) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Credentials2) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Credentials2) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTokenType

`func (o *Credentials2) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *Credentials2) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *Credentials2) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *Credentials2) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *Credentials2) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *Credentials2) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *Credentials2) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetSessionId

`func (o *Credentials2) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *Credentials2) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *Credentials2) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *Credentials2) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetCode

`func (o *Credentials2) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Credentials2) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Credentials2) SetCode(v string)`

SetCode sets Code field to given value.


### GetGrantType

`func (o *Credentials2) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *Credentials2) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *Credentials2) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *Credentials2) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetAppId

`func (o *Credentials2) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *Credentials2) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *Credentials2) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetAppKey

`func (o *Credentials2) GetAppKey() string`

GetAppKey returns the AppKey field if non-nil, zero value otherwise.

### GetAppKeyOk

`func (o *Credentials2) GetAppKeyOk() (*string, bool)`

GetAppKeyOk returns a tuple with the AppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppKey

`func (o *Credentials2) SetAppKey(v string)`

SetAppKey sets AppKey field to given value.


### GetClientKey

`func (o *Credentials2) GetClientKey() string`

GetClientKey returns the ClientKey field if non-nil, zero value otherwise.

### GetClientKeyOk

`func (o *Credentials2) GetClientKeyOk() (*string, bool)`

GetClientKeyOk returns a tuple with the ClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientKey

`func (o *Credentials2) SetClientKey(v string)`

SetClientKey sets ClientKey field to given value.


### GetAccountId

`func (o *Credentials2) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Credentials2) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Credentials2) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetProviderToken

`func (o *Credentials2) GetProviderToken() string`

GetProviderToken returns the ProviderToken field if non-nil, zero value otherwise.

### GetProviderTokenOk

`func (o *Credentials2) GetProviderTokenOk() (*string, bool)`

GetProviderTokenOk returns a tuple with the ProviderToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderToken

`func (o *Credentials2) SetProviderToken(v string)`

SetProviderToken sets ProviderToken field to given value.


### GetDotNumber

`func (o *Credentials2) GetDotNumber() string`

GetDotNumber returns the DotNumber field if non-nil, zero value otherwise.

### GetDotNumberOk

`func (o *Credentials2) GetDotNumberOk() (*string, bool)`

GetDotNumberOk returns a tuple with the DotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDotNumber

`func (o *Credentials2) SetDotNumber(v string)`

SetDotNumber sets DotNumber field to given value.


### GetCarrierId

`func (o *Credentials2) GetCarrierId() string`

GetCarrierId returns the CarrierId field if non-nil, zero value otherwise.

### GetCarrierIdOk

`func (o *Credentials2) GetCarrierIdOk() (*string, bool)`

GetCarrierIdOk returns a tuple with the CarrierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierId

`func (o *Credentials2) SetCarrierId(v string)`

SetCarrierId sets CarrierId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


