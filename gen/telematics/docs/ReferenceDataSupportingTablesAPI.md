# \ReferenceDataSupportingTablesAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListHosEventCodesReference**](ReferenceDataSupportingTablesAPI.md#ListHosEventCodesReference) | **Get** /v2/telematics/ref-hos-event-codes | List HOS Event Codes
[**ListHosMalfunctionCodesReference**](ReferenceDataSupportingTablesAPI.md#ListHosMalfunctionCodesReference) | **Get** /v2/telematics/ref-hos-malfunction-codes | List HOS Malfunction Codes
[**ListHosRecordOriginsReference**](ReferenceDataSupportingTablesAPI.md#ListHosRecordOriginsReference) | **Get** /v2/telematics/ref-hos-record-origins | List HOS Record Origins
[**ListHosRecordStatusesReference**](ReferenceDataSupportingTablesAPI.md#ListHosRecordStatusesReference) | **Get** /v2/telematics/ref-hos-record-statuses | List HOS Record Statuses
[**ListHosRegionsReference**](ReferenceDataSupportingTablesAPI.md#ListHosRegionsReference) | **Get** /v2/telematics/ref-hos-regions | List HOS Regions
[**ListHosViolationCodesReference**](ReferenceDataSupportingTablesAPI.md#ListHosViolationCodesReference) | **Get** /v2/telematics/ref-hos-violation-codes | List HOS Violation Codes
[**ListRulesetsReference**](ReferenceDataSupportingTablesAPI.md#ListRulesetsReference) | **Get** /v2/telematics/ref-hos-rulesets | List HOS Rulesets
[**ListTimezonesReference**](ReferenceDataSupportingTablesAPI.md#ListTimezonesReference) | **Get** /v2/telematics/ref-timezones | List Timezones



## ListHosEventCodesReference

> CursorPageTypeVarCustomizedRefHosEventCode ListHosEventCodesReference(ctx).Cursor(cursor).Size(size).Execute()

List HOS Event Codes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/telematicsapi"
)

func main() {
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataSupportingTablesAPI.ListHosEventCodesReference(context.Background()).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataSupportingTablesAPI.ListHosEventCodesReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosEventCodesReference`: CursorPageTypeVarCustomizedRefHosEventCode
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataSupportingTablesAPI.ListHosEventCodesReference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHosEventCodesReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedRefHosEventCode**](CursorPageTypeVarCustomizedRefHosEventCode.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosMalfunctionCodesReference

> CursorPageTypeVarCustomizedRefHosMalfunctionCode ListHosMalfunctionCodesReference(ctx).Cursor(cursor).Size(size).Execute()

List HOS Malfunction Codes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/telematicsapi"
)

func main() {
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataSupportingTablesAPI.ListHosMalfunctionCodesReference(context.Background()).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataSupportingTablesAPI.ListHosMalfunctionCodesReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosMalfunctionCodesReference`: CursorPageTypeVarCustomizedRefHosMalfunctionCode
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataSupportingTablesAPI.ListHosMalfunctionCodesReference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHosMalfunctionCodesReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedRefHosMalfunctionCode**](CursorPageTypeVarCustomizedRefHosMalfunctionCode.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosRecordOriginsReference

> CursorPageTypeVarCustomizedRefHosRecordOrigin ListHosRecordOriginsReference(ctx).Cursor(cursor).Size(size).Execute()

List HOS Record Origins



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/telematicsapi"
)

func main() {
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataSupportingTablesAPI.ListHosRecordOriginsReference(context.Background()).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataSupportingTablesAPI.ListHosRecordOriginsReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosRecordOriginsReference`: CursorPageTypeVarCustomizedRefHosRecordOrigin
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataSupportingTablesAPI.ListHosRecordOriginsReference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHosRecordOriginsReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedRefHosRecordOrigin**](CursorPageTypeVarCustomizedRefHosRecordOrigin.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosRecordStatusesReference

> CursorPageTypeVarCustomizedRefHosRecordStatus ListHosRecordStatusesReference(ctx).Cursor(cursor).Size(size).Execute()

List HOS Record Statuses



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/telematicsapi"
)

func main() {
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataSupportingTablesAPI.ListHosRecordStatusesReference(context.Background()).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataSupportingTablesAPI.ListHosRecordStatusesReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosRecordStatusesReference`: CursorPageTypeVarCustomizedRefHosRecordStatus
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataSupportingTablesAPI.ListHosRecordStatusesReference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHosRecordStatusesReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedRefHosRecordStatus**](CursorPageTypeVarCustomizedRefHosRecordStatus.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosRegionsReference

> CursorPageTypeVarCustomizedRefHosRegion ListHosRegionsReference(ctx).Cursor(cursor).Size(size).Execute()

List HOS Regions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/telematicsapi"
)

func main() {
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataSupportingTablesAPI.ListHosRegionsReference(context.Background()).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataSupportingTablesAPI.ListHosRegionsReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosRegionsReference`: CursorPageTypeVarCustomizedRefHosRegion
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataSupportingTablesAPI.ListHosRegionsReference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHosRegionsReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedRefHosRegion**](CursorPageTypeVarCustomizedRefHosRegion.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosViolationCodesReference

> CursorPageTypeVarCustomizedRefHosViolationCode ListHosViolationCodesReference(ctx).Cursor(cursor).Size(size).Execute()

List HOS Violation Codes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/telematicsapi"
)

func main() {
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataSupportingTablesAPI.ListHosViolationCodesReference(context.Background()).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataSupportingTablesAPI.ListHosViolationCodesReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosViolationCodesReference`: CursorPageTypeVarCustomizedRefHosViolationCode
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataSupportingTablesAPI.ListHosViolationCodesReference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHosViolationCodesReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedRefHosViolationCode**](CursorPageTypeVarCustomizedRefHosViolationCode.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRulesetsReference

> CursorPageTypeVarCustomizedRefHosRuleset ListRulesetsReference(ctx).Cursor(cursor).Size(size).Execute()

List HOS Rulesets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/telematicsapi"
)

func main() {
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataSupportingTablesAPI.ListRulesetsReference(context.Background()).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataSupportingTablesAPI.ListRulesetsReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRulesetsReference`: CursorPageTypeVarCustomizedRefHosRuleset
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataSupportingTablesAPI.ListRulesetsReference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRulesetsReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedRefHosRuleset**](CursorPageTypeVarCustomizedRefHosRuleset.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTimezonesReference

> CursorPageTypeVarCustomizedRefTimezoneCode ListTimezonesReference(ctx).Cursor(cursor).Size(size).Execute()

List Timezones



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/telematicsapi"
)

func main() {
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReferenceDataSupportingTablesAPI.ListTimezonesReference(context.Background()).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataSupportingTablesAPI.ListTimezonesReference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTimezonesReference`: CursorPageTypeVarCustomizedRefTimezoneCode
	fmt.Fprintf(os.Stdout, "Response from `ReferenceDataSupportingTablesAPI.ListTimezonesReference`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTimezonesReferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedRefTimezoneCode**](CursorPageTypeVarCustomizedRefTimezoneCode.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

