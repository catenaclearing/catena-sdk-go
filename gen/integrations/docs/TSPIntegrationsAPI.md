# \TSPIntegrationsAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIntegrationsPerTsp**](TSPIntegrationsAPI.md#GetIntegrationsPerTsp) | **Get** /v2/integrations/tsps/{tsp_id} | Get Integrations Per Tsp
[**ListIntegrationsPerTsp**](TSPIntegrationsAPI.md#ListIntegrationsPerTsp) | **Get** /v2/integrations/tsps | List Integrations Per Tsp



## GetIntegrationsPerTsp

> TspIntegrationsRead GetIntegrationsPerTsp(ctx, tspId).Execute()

Get Integrations Per Tsp



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/integrations"
)

func main() {
	tspId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TSPIntegrationsAPI.GetIntegrationsPerTsp(context.Background(), tspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TSPIntegrationsAPI.GetIntegrationsPerTsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIntegrationsPerTsp`: TspIntegrationsRead
	fmt.Fprintf(os.Stdout, "Response from `TSPIntegrationsAPI.GetIntegrationsPerTsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIntegrationsPerTspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TspIntegrationsRead**](TspIntegrationsRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIntegrationsPerTsp

> CursorPageTypeVarCustomizedTspIntegrationsRead ListIntegrationsPerTsp(ctx).Cursor(cursor).Size(size).Execute()

List Integrations Per Tsp



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/integrations"
)

func main() {
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TSPIntegrationsAPI.ListIntegrationsPerTsp(context.Background()).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TSPIntegrationsAPI.ListIntegrationsPerTsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIntegrationsPerTsp`: CursorPageTypeVarCustomizedTspIntegrationsRead
	fmt.Fprintf(os.Stdout, "Response from `TSPIntegrationsAPI.ListIntegrationsPerTsp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIntegrationsPerTspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedTspIntegrationsRead**](CursorPageTypeVarCustomizedTspIntegrationsRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

