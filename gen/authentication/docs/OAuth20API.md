# \OAuth20API

All URIs are relative to *https://auth.catenatelematics.com/realms/catena*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RevokeToken**](OAuth20API.md#RevokeToken) | **Post** /protocol/openid-connect/revoke | Revoke Token
[**Token**](OAuth20API.md#Token) | **Post** /protocol/openid-connect/token | Get Token



## RevokeToken

> RevokeToken(ctx).Token(token).ClientId(clientId).ClientSecret(clientSecret).Execute()

Revoke Token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/authentication"
)

func main() {
	token := "token_example" // string | The token to revoke.
	clientId := "clientId_example" // string | Your client identifier.
	clientSecret := "clientSecret_example" // string | Your client secret.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OAuth20API.RevokeToken(context.Background()).Token(token).ClientId(clientId).ClientSecret(clientSecret).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20API.RevokeToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | The token to revoke. | 
 **clientId** | **string** | Your client identifier. | 
 **clientSecret** | **string** | Your client secret. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Token

> TokenResponse Token(ctx).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Scope(scope).Execute()

Get Token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/authentication"
)

func main() {
	grantType := "grantType_example" // string | Grant type for token request. Must be \\\"client_credentials\\\" (default to "client_credentials")
	clientId := "clientId_example" // string | Your client identifier
	clientSecret := "clientSecret_example" // string | Your client secret
	scope := "scope_example" // string | Required scope for API access. Must be \\\"organization\\\" (default to "organization")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuth20API.Token(context.Background()).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuth20API.Token``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Token`: TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `OAuth20API.Token`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | Grant type for token request. Must be \\\&quot;client_credentials\\\&quot; | [default to &quot;client_credentials&quot;]
 **clientId** | **string** | Your client identifier | 
 **clientSecret** | **string** | Your client secret | 
 **scope** | **string** | Required scope for API access. Must be \\\&quot;organization\\\&quot; | [default to &quot;organization&quot;]

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

