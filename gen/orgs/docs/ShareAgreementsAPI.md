# \ShareAgreementsAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateShareAgreement**](ShareAgreementsAPI.md#ActivateShareAgreement) | **Post** /v2/orgs/share-agreements/{share_agreement_id}/activate | Activate Share Agreement
[**BackfillShareAgreements**](ShareAgreementsAPI.md#BackfillShareAgreements) | **Post** /v2/orgs/share-agreements/backfill | Backfill Share Agreements
[**CancelShareAgreement**](ShareAgreementsAPI.md#CancelShareAgreement) | **Post** /v2/orgs/share-agreements/{share_agreement_id}/cancel | Cancel Share Agreement
[**DeleteShareAgreement**](ShareAgreementsAPI.md#DeleteShareAgreement) | **Delete** /v2/orgs/share-agreements/{share_agreement_id} | Delete Share Agreement
[**GetShareAgreement**](ShareAgreementsAPI.md#GetShareAgreement) | **Get** /v2/orgs/share-agreements/{share_agreement_id} | Get Share Agreement
[**ListShareAgreements**](ShareAgreementsAPI.md#ListShareAgreements) | **Get** /v2/orgs/share-agreements | List Share Agreements
[**PauseShareAgreement**](ShareAgreementsAPI.md#PauseShareAgreement) | **Post** /v2/orgs/share-agreements/{share_agreement_id}/pause | Pause Share Agreement
[**UpdateShareAgreement**](ShareAgreementsAPI.md#UpdateShareAgreement) | **Patch** /v2/orgs/share-agreements/{share_agreement_id} | Update Share Agreement



## ActivateShareAgreement

> ShareAgreementRead ActivateShareAgreement(ctx, shareAgreementId).Execute()

Activate Share Agreement



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
	shareAgreementId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAgreementsAPI.ActivateShareAgreement(context.Background(), shareAgreementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAgreementsAPI.ActivateShareAgreement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateShareAgreement`: ShareAgreementRead
	fmt.Fprintf(os.Stdout, "Response from `ShareAgreementsAPI.ActivateShareAgreement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareAgreementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateShareAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShareAgreementRead**](ShareAgreementRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackfillShareAgreements

> ShareAgreementBackfillResult BackfillShareAgreements(ctx).ShareAgreementBackfill(shareAgreementBackfill).Execute()

Backfill Share Agreements



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
	shareAgreementBackfill := *openapiclient.NewShareAgreementBackfill(openapiclient.ResourceEnum("asset")) // ShareAgreementBackfill | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAgreementsAPI.BackfillShareAgreements(context.Background()).ShareAgreementBackfill(shareAgreementBackfill).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAgreementsAPI.BackfillShareAgreements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BackfillShareAgreements`: ShareAgreementBackfillResult
	fmt.Fprintf(os.Stdout, "Response from `ShareAgreementsAPI.BackfillShareAgreements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBackfillShareAgreementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shareAgreementBackfill** | [**ShareAgreementBackfill**](ShareAgreementBackfill.md) |  | 

### Return type

[**ShareAgreementBackfillResult**](ShareAgreementBackfillResult.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelShareAgreement

> ShareAgreementRead CancelShareAgreement(ctx, shareAgreementId).Execute()

Cancel Share Agreement



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
	shareAgreementId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAgreementsAPI.CancelShareAgreement(context.Background(), shareAgreementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAgreementsAPI.CancelShareAgreement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelShareAgreement`: ShareAgreementRead
	fmt.Fprintf(os.Stdout, "Response from `ShareAgreementsAPI.CancelShareAgreement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareAgreementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelShareAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShareAgreementRead**](ShareAgreementRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShareAgreement

> DeleteShareAgreement(ctx, shareAgreementId).Execute()

Delete Share Agreement



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
	shareAgreementId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ShareAgreementsAPI.DeleteShareAgreement(context.Background(), shareAgreementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAgreementsAPI.DeleteShareAgreement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareAgreementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteShareAgreementRequest struct via the builder pattern


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


## GetShareAgreement

> ShareAgreementRead GetShareAgreement(ctx, shareAgreementId).Execute()

Get Share Agreement



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
	shareAgreementId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAgreementsAPI.GetShareAgreement(context.Background(), shareAgreementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAgreementsAPI.GetShareAgreement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShareAgreement`: ShareAgreementRead
	fmt.Fprintf(os.Stdout, "Response from `ShareAgreementsAPI.GetShareAgreement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareAgreementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShareAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShareAgreementRead**](ShareAgreementRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShareAgreements

> CursorPageShareAgreementRead ListShareAgreements(ctx).FleetId(fleetId).FleetRef(fleetRef).ShareAgreementStatus(shareAgreementStatus).Cursor(cursor).Size(size).Execute()

List Share Agreements



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
	fleetId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Limit results to specific fleet (optional)
	fleetRef := "fleetRef_example" // string | Limit results to specific fleet reference (optional)
	shareAgreementStatus := openapiclient.ShareAgreementStatusEnum("active") // ShareAgreementStatusEnum | Limit results to specific Share Agreement status (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAgreementsAPI.ListShareAgreements(context.Background()).FleetId(fleetId).FleetRef(fleetRef).ShareAgreementStatus(shareAgreementStatus).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAgreementsAPI.ListShareAgreements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListShareAgreements`: CursorPageShareAgreementRead
	fmt.Fprintf(os.Stdout, "Response from `ShareAgreementsAPI.ListShareAgreements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListShareAgreementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetId** | **string** | Limit results to specific fleet | 
 **fleetRef** | **string** | Limit results to specific fleet reference | 
 **shareAgreementStatus** | [**ShareAgreementStatusEnum**](ShareAgreementStatusEnum.md) | Limit results to specific Share Agreement status | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageShareAgreementRead**](CursorPageShareAgreementRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseShareAgreement

> ShareAgreementRead PauseShareAgreement(ctx, shareAgreementId).Execute()

Pause Share Agreement



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
	shareAgreementId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAgreementsAPI.PauseShareAgreement(context.Background(), shareAgreementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAgreementsAPI.PauseShareAgreement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PauseShareAgreement`: ShareAgreementRead
	fmt.Fprintf(os.Stdout, "Response from `ShareAgreementsAPI.PauseShareAgreement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareAgreementId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseShareAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ShareAgreementRead**](ShareAgreementRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShareAgreement

> ShareAgreementRead UpdateShareAgreement(ctx, shareAgreementId).ShareAgreementUpdate(shareAgreementUpdate).Execute()

Update Share Agreement



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
	shareAgreementId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the Share Agreement to update
	shareAgreementUpdate := *openapiclient.NewShareAgreementUpdate() // ShareAgreementUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShareAgreementsAPI.UpdateShareAgreement(context.Background(), shareAgreementId).ShareAgreementUpdate(shareAgreementUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShareAgreementsAPI.UpdateShareAgreement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateShareAgreement`: ShareAgreementRead
	fmt.Fprintf(os.Stdout, "Response from `ShareAgreementsAPI.UpdateShareAgreement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shareAgreementId** | **string** | The ID of the Share Agreement to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShareAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shareAgreementUpdate** | [**ShareAgreementUpdate**](ShareAgreementUpdate.md) |  | 

### Return type

[**ShareAgreementRead**](ShareAgreementRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

