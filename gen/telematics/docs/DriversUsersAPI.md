# \DriversUsersAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUser**](DriversUsersAPI.md#GetUser) | **Get** /v2/telematics/users/{user_id} | Get User
[**ListUsers**](DriversUsersAPI.md#ListUsers) | **Get** /v2/telematics/users | List Users



## GetUser

> User GetUser(ctx, userId).IncludeSourceData(includeSourceData).Execute()

Get User



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
	userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the user.
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriversUsersAPI.GetUser(context.Background(), userId).IncludeSourceData(includeSourceData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriversUsersAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: User
	fmt.Fprintf(os.Stdout, "Response from `DriversUsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The unique identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]

### Return type

[**User**](User.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> CursorPageTypeVarCustomizedUser ListUsers(ctx).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).UserIds(userIds).Cursor(cursor).Size(size).Execute()

List Users



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
	fleetIds := []string{"Inner_example"} // []string | Limit results to specific fleets using Catena's fleet IDs. *For your own fleet identifiers, use `fleet_refs` instead* (optional)
	fleetRefs := []string{"Inner_example"} // []string | Limit results to specific fleets using your organization's fleet reference identifiers (optional)
	includeSourceData := true // bool | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* (optional) (default to false)
	userIds := []string{"Inner_example"} // []string | Limit results to specific users. **Maximum:** 100 IDs (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 300)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DriversUsersAPI.ListUsers(context.Background()).FleetIds(fleetIds).FleetRefs(fleetRefs).IncludeSourceData(includeSourceData).UserIds(userIds).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DriversUsersAPI.ListUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsers`: CursorPageTypeVarCustomizedUser
	fmt.Fprintf(os.Stdout, "Response from `DriversUsersAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetIds** | **[]string** | Limit results to specific fleets using Catena&#39;s fleet IDs. *For your own fleet identifiers, use &#x60;fleet_refs&#x60; instead* | 
 **fleetRefs** | **[]string** | Limit results to specific fleets using your organization&#39;s fleet reference identifiers | 
 **includeSourceData** | **bool** | Include the raw data from the telematics provider. *Useful for auditing or accessing fields not normalized by Catena* | [default to false]
 **userIds** | **[]string** | Limit results to specific users. **Maximum:** 100 IDs | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 300]

### Return type

[**CursorPageTypeVarCustomizedUser**](CursorPageTypeVarCustomizedUser.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

