# \ComplianceRegulationAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDvirLogDefects**](ComplianceRegulationAPI.md#GetDvirLogDefects) | **Get** /v2/telematics/dvir-logs/{dvir_log_id}/defects | Get Dvir Log Defects
[**GetHosEventAttachments**](ComplianceRegulationAPI.md#GetHosEventAttachments) | **Get** /v2/telematics/hos-events/{hos_event_id}/attachments | Get Hos Event Attachments
[**ListDvirLogDefects**](ComplianceRegulationAPI.md#ListDvirLogDefects) | **Get** /v2/telematics/dvir-defects | List Dvir Log Defects
[**ListDvirLogs**](ComplianceRegulationAPI.md#ListDvirLogs) | **Get** /v2/telematics/dvir-logs | List Dvir Logs
[**ListHosAvailabilities**](ComplianceRegulationAPI.md#ListHosAvailabilities) | **Get** /v2/telematics/hos-availabilities | List HOS Availabilities
[**ListHosDailySnapshots**](ComplianceRegulationAPI.md#ListHosDailySnapshots) | **Get** /v2/telematics/hos-daily-snapshots | List Hos Daily Snapshots
[**ListHosEvents**](ComplianceRegulationAPI.md#ListHosEvents) | **Get** /v2/telematics/hos-events | List HOS Events
[**ListHosViolations**](ComplianceRegulationAPI.md#ListHosViolations) | **Get** /v2/telematics/hos-violations | List HOS Violations
[**ListIftaSummaries**](ComplianceRegulationAPI.md#ListIftaSummaries) | **Get** /v2/telematics/ifta-summaries | List Ifta Summaries



## GetDvirLogDefects

> CursorPageDvirLogDefect GetDvirLogDefects(ctx, dvirLogId).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()

Get Dvir Log Defects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	dvirLogId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the DVIR log
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceRegulationAPI.GetDvirLogDefects(context.Background(), dvirLogId).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceRegulationAPI.GetDvirLogDefects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDvirLogDefects`: CursorPageDvirLogDefect
	fmt.Fprintf(os.Stdout, "Response from `ComplianceRegulationAPI.GetDvirLogDefects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dvirLogId** | **string** | The unique identifier of the DVIR log | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDvirLogDefectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 50]

### Return type

[**CursorPageDvirLogDefect**](CursorPageDvirLogDefect.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHosEventAttachments

> CursorPageTypeVarCustomizedHosEventAttachment GetHosEventAttachments(ctx, hosEventId).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()

Get Hos Event Attachments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	hosEventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the HOS event
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceRegulationAPI.GetHosEventAttachments(context.Background(), hosEventId).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceRegulationAPI.GetHosEventAttachments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHosEventAttachments`: CursorPageTypeVarCustomizedHosEventAttachment
	fmt.Fprintf(os.Stdout, "Response from `ComplianceRegulationAPI.GetHosEventAttachments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hosEventId** | **string** | The ID of the HOS event | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHosEventAttachmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedHosEventAttachment**](CursorPageTypeVarCustomizedHosEventAttachment.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDvirLogDefects

> CursorPageDvirLogDefect ListDvirLogDefects(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).DvirLogIds(dvirLogIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()

List Dvir Log Defects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 15 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 15 days (optional)
	dvirLogIds := []string{"Inner_example"} // []string | Limit results to specific DVIR logs. **Maximum:** 100 IDs (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceRegulationAPI.ListDvirLogDefects(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).DvirLogIds(dvirLogIds).IncludeSourceData(includeSourceData).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceRegulationAPI.ListDvirLogDefects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDvirLogDefects`: CursorPageDvirLogDefect
	fmt.Fprintf(os.Stdout, "Response from `ComplianceRegulationAPI.ListDvirLogDefects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDvirLogDefectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 15 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 15 days | 
 **dvirLogIds** | **[]string** | Limit results to specific DVIR logs. **Maximum:** 100 IDs | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 50]

### Return type

[**CursorPageDvirLogDefect**](CursorPageDvirLogDefect.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDvirLogs

> CursorPageDvirLog ListDvirLogs(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).DriverIds(driverIds).VehicleIds(vehicleIds).Cursor(cursor).Size(size).Execute()

List Dvir Logs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 15 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 15 days (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	driverIds := []string{"Inner_example"} // []string | Limit results to specific drivers. **Maximum:** 100 IDs (optional)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceRegulationAPI.ListDvirLogs(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).DriverIds(driverIds).VehicleIds(vehicleIds).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceRegulationAPI.ListDvirLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDvirLogs`: CursorPageDvirLog
	fmt.Fprintf(os.Stdout, "Response from `ComplianceRegulationAPI.ListDvirLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDvirLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 15 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 15 days | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **driverIds** | **[]string** | Limit results to specific drivers. **Maximum:** 100 IDs | 
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 50]

### Return type

[**CursorPageDvirLog**](CursorPageDvirLog.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosAvailabilities

> CursorPageTypeVarCustomizedHosAvailability ListHosAvailabilities(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).DriverIds(driverIds).VehicleIds(vehicleIds).Cursor(cursor).Size(size).Execute()

List HOS Availabilities



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	driverIds := []string{"Inner_example"} // []string | Limit results to specific drivers. **Maximum:** 100 IDs (optional)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceRegulationAPI.ListHosAvailabilities(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).DriverIds(driverIds).VehicleIds(vehicleIds).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceRegulationAPI.ListHosAvailabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosAvailabilities`: CursorPageTypeVarCustomizedHosAvailability
	fmt.Fprintf(os.Stdout, "Response from `ComplianceRegulationAPI.ListHosAvailabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHosAvailabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **driverIds** | **[]string** | Limit results to specific drivers. **Maximum:** 100 IDs | 
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedHosAvailability**](CursorPageTypeVarCustomizedHosAvailability.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosDailySnapshots

> CursorPageTypeVarCustomizedHosDailySnapshot ListHosDailySnapshots(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).DriverIds(driverIds).Cursor(cursor).Size(size).Execute()

List Hos Daily Snapshots



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 15 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 15 days (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	driverIds := []string{"Inner_example"} // []string | Limit results to specific drivers. **Maximum:** 100 IDs (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceRegulationAPI.ListHosDailySnapshots(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).DriverIds(driverIds).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceRegulationAPI.ListHosDailySnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosDailySnapshots`: CursorPageTypeVarCustomizedHosDailySnapshot
	fmt.Fprintf(os.Stdout, "Response from `ComplianceRegulationAPI.ListHosDailySnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHosDailySnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 15 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 15 days | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **driverIds** | **[]string** | Limit results to specific drivers. **Maximum:** 100 IDs | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedHosDailySnapshot**](CursorPageTypeVarCustomizedHosDailySnapshot.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosEvents

> CursorPageTypeVarCustomizedHosEvent ListHosEvents(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).DriverIds(driverIds).VehicleIds(vehicleIds).Cursor(cursor).Size(size).Execute()

List HOS Events



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 15 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 15 days (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	driverIds := []string{"Inner_example"} // []string | Limit results to specific drivers. **Maximum:** 100 IDs (optional)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceRegulationAPI.ListHosEvents(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).DriverIds(driverIds).VehicleIds(vehicleIds).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceRegulationAPI.ListHosEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosEvents`: CursorPageTypeVarCustomizedHosEvent
	fmt.Fprintf(os.Stdout, "Response from `ComplianceRegulationAPI.ListHosEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHosEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 15 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 15 days | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **driverIds** | **[]string** | Limit results to specific drivers. **Maximum:** 100 IDs | 
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedHosEvent**](CursorPageTypeVarCustomizedHosEvent.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHosViolations

> CursorPageTypeVarCustomizedHosViolation ListHosViolations(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).DriverIds(driverIds).Cursor(cursor).Size(size).Execute()

List HOS Violations



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	fromDatetime := time.Now() // time.Time | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at >= from_datetime` **Default value:** `now() - 1 day` **Restriction:** `to_datetime - from_datetime` cannot exceed 15 days (optional)
	toDatetime := time.Now() // time.Time | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** `occurred_at < to_datetime` **Default value:** `now()` **Restriction:** `to_datetime - from_datetime` cannot exceed 15 days (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	driverIds := []string{"Inner_example"} // []string | Limit results to specific drivers. **Maximum:** 100 IDs (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceRegulationAPI.ListHosViolations(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).FromDatetime(fromDatetime).ToDatetime(toDatetime).IncludeSourceData(includeSourceData).DriverIds(driverIds).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceRegulationAPI.ListHosViolations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHosViolations`: CursorPageTypeVarCustomizedHosViolation
	fmt.Fprintf(os.Stdout, "Response from `ComplianceRegulationAPI.ListHosViolations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHosViolationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **fromDatetime** | **time.Time** | Return only records that occurred on or after this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &gt;&#x3D; from_datetime&#x60; **Default value:** &#x60;now() - 1 day&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 15 days | 
 **toDatetime** | **time.Time** | Return only records that occurred before this date and time. **Format:** ISO 8601 (UTC) **Applies filter:** &#x60;occurred_at &lt; to_datetime&#x60; **Default value:** &#x60;now()&#x60; **Restriction:** &#x60;to_datetime - from_datetime&#x60; cannot exceed 15 days | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **driverIds** | **[]string** | Limit results to specific drivers. **Maximum:** 100 IDs | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedHosViolation**](CursorPageTypeVarCustomizedHosViolation.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIftaSummaries

> CursorPageTypeVarCustomizedIftaSummary ListIftaSummaries(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).VehicleIds(vehicleIds).Cursor(cursor).Size(size).Execute()

List Ifta Summaries



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/catenaclearing/catena-sdk-go/gen/telematics"
)

func main() {
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	vehicleIds := []string{"Inner_example"} // []string | Limit results to specific vehicles. **Maximum:** 100 IDs (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComplianceRegulationAPI.ListIftaSummaries(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).VehicleIds(vehicleIds).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComplianceRegulationAPI.ListIftaSummaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIftaSummaries`: CursorPageTypeVarCustomizedIftaSummary
	fmt.Fprintf(os.Stdout, "Response from `ComplianceRegulationAPI.ListIftaSummaries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIftaSummariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **vehicleIds** | **[]string** | Limit results to specific vehicles. **Maximum:** 100 IDs | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedIftaSummary**](CursorPageTypeVarCustomizedIftaSummary.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

