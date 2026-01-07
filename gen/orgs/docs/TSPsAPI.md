# \TSPsAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTsp**](TSPsAPI.md#CreateTsp) | **Post** /v2/orgs/tsps | Create Tsp
[**DeleteTsp**](TSPsAPI.md#DeleteTsp) | **Delete** /v2/orgs/tsps/{tsp_id} | Delete Tsp
[**GetTsp**](TSPsAPI.md#GetTsp) | **Get** /v2/orgs/tsps/{tsp_id} | Get TSP
[**ListTsps**](TSPsAPI.md#ListTsps) | **Get** /v2/orgs/tsps | List TSPs
[**UpdateTsp**](TSPsAPI.md#UpdateTsp) | **Patch** /v2/orgs/tsps/{tsp_id} | Update Tsp



## CreateTsp

> TspRead CreateTsp(ctx).TspCreate(tspCreate).Execute()

Create Tsp



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/orgs"
)

func main() {
	tspCreate := *openapiclient.NewTspCreate("Name_example", openapiclient.TspEnum("ada"), openapiclient.ConnectionTypeEnum("access_token")) // TspCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TSPsAPI.CreateTsp(context.Background()).TspCreate(tspCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TSPsAPI.CreateTsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTsp`: TspRead
	fmt.Fprintf(os.Stdout, "Response from `TSPsAPI.CreateTsp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tspCreate** | [**TspCreate**](TspCreate.md) |  | 

### Return type

[**TspRead**](TspRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTsp

> DeleteTsp(ctx, tspId).Execute()

Delete Tsp



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/orgs"
)

func main() {
	tspId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TSPsAPI.DeleteTsp(context.Background(), tspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TSPsAPI.DeleteTsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTsp

> TspRead GetTsp(ctx, tspId).Execute()

Get TSP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/orgs"
)

func main() {
	tspId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TSPsAPI.GetTsp(context.Background(), tspId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TSPsAPI.GetTsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTsp`: TspRead
	fmt.Fprintf(os.Stdout, "Response from `TSPsAPI.GetTsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TspRead**](TspRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTsps

> []TspRead ListTsps(ctx).Slug(slug).SourceName(sourceName).Execute()

List TSPs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/orgs"
)

func main() {
	slug := "slug_example" // string |  (optional)
	sourceName := openapiclient.TspEnum("ada") // TspEnum |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TSPsAPI.ListTsps(context.Background()).Slug(slug).SourceName(sourceName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TSPsAPI.ListTsps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTsps`: []TspRead
	fmt.Fprintf(os.Stdout, "Response from `TSPsAPI.ListTsps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTspsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **slug** | **string** |  | 
 **sourceName** | [**TspEnum**](TspEnum.md) |  | 

### Return type

[**[]TspRead**](TspRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTsp

> TspRead UpdateTsp(ctx, tspId).TspUpdate(tspUpdate).Execute()

Update Tsp



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/orgs"
)

func main() {
	tspId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	tspUpdate := *openapiclient.NewTspUpdate() // TspUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TSPsAPI.UpdateTsp(context.Background(), tspId).TspUpdate(tspUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TSPsAPI.UpdateTsp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTsp`: TspRead
	fmt.Fprintf(os.Stdout, "Response from `TSPsAPI.UpdateTsp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tspId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTspRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tspUpdate** | [**TspUpdate**](TspUpdate.md) |  | 

### Return type

[**TspRead**](TspRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

