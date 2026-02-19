# \InvitationsAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptInvitationV2OrgsInvitationsInvitationIdAcceptPost**](InvitationsAPI.md#AcceptInvitationV2OrgsInvitationsInvitationIdAcceptPost) | **Post** /v2/orgs/invitations/{invitation_id}/accept | Accept Invitation
[**CreateInvitation**](InvitationsAPI.md#CreateInvitation) | **Post** /v2/orgs/invitations | Create Invitation
[**DeclineInvitationV2OrgsInvitationsInvitationIdDeclinePost**](InvitationsAPI.md#DeclineInvitationV2OrgsInvitationsInvitationIdDeclinePost) | **Post** /v2/orgs/invitations/{invitation_id}/decline | Decline Invitation
[**DeleteInvitation**](InvitationsAPI.md#DeleteInvitation) | **Delete** /v2/orgs/invitations/{invitation_id} | Delete Invitation
[**GetInvitation**](InvitationsAPI.md#GetInvitation) | **Get** /v2/orgs/invitations/{invitation_id} | Get Invitation
[**ListInvitations**](InvitationsAPI.md#ListInvitations) | **Get** /v2/orgs/invitations | List Invitations



## AcceptInvitationV2OrgsInvitationsInvitationIdAcceptPost

> InvitationRead AcceptInvitationV2OrgsInvitationsInvitationIdAcceptPost(ctx, invitationId).InvitationAccept(invitationAccept).Execute()

Accept Invitation



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
	invitationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invitationAccept := *openapiclient.NewInvitationAccept() // InvitationAccept | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitationsAPI.AcceptInvitationV2OrgsInvitationsInvitationIdAcceptPost(context.Background(), invitationId).InvitationAccept(invitationAccept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.AcceptInvitationV2OrgsInvitationsInvitationIdAcceptPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptInvitationV2OrgsInvitationsInvitationIdAcceptPost`: InvitationRead
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.AcceptInvitationV2OrgsInvitationsInvitationIdAcceptPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptInvitationV2OrgsInvitationsInvitationIdAcceptPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invitationAccept** | [**InvitationAccept**](InvitationAccept.md) |  | 

### Return type

[**InvitationRead**](InvitationRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInvitation

> InvitationRead CreateInvitation(ctx).InvitationCreate(invitationCreate).Execute()

Create Invitation



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
	invitationCreate := *openapiclient.NewInvitationCreate("FleetRef_example") // InvitationCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitationsAPI.CreateInvitation(context.Background()).InvitationCreate(invitationCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.CreateInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInvitation`: InvitationRead
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.CreateInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invitationCreate** | [**InvitationCreate**](InvitationCreate.md) |  | 

### Return type

[**InvitationRead**](InvitationRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeclineInvitationV2OrgsInvitationsInvitationIdDeclinePost

> InvitationRead DeclineInvitationV2OrgsInvitationsInvitationIdDeclinePost(ctx, invitationId).InvitationDecline(invitationDecline).Execute()

Decline Invitation



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
	invitationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	invitationDecline := *openapiclient.NewInvitationDecline() // InvitationDecline | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitationsAPI.DeclineInvitationV2OrgsInvitationsInvitationIdDeclinePost(context.Background(), invitationId).InvitationDecline(invitationDecline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.DeclineInvitationV2OrgsInvitationsInvitationIdDeclinePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeclineInvitationV2OrgsInvitationsInvitationIdDeclinePost`: InvitationRead
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.DeclineInvitationV2OrgsInvitationsInvitationIdDeclinePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeclineInvitationV2OrgsInvitationsInvitationIdDeclinePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invitationDecline** | [**InvitationDecline**](InvitationDecline.md) |  | 

### Return type

[**InvitationRead**](InvitationRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInvitation

> DeleteInvitation(ctx, invitationId).Execute()

Delete Invitation



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
	invitationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InvitationsAPI.DeleteInvitation(context.Background(), invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.DeleteInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInvitationRequest struct via the builder pattern


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


## GetInvitation

> InvitationRead GetInvitation(ctx, invitationId).Execute()

Get Invitation



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
	invitationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitationsAPI.GetInvitation(context.Background(), invitationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.GetInvitation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvitation`: InvitationRead
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.GetInvitation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invitationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InvitationRead**](InvitationRead.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvitations

> CursorPageCustomizedInvitationRead ListInvitations(ctx).FleetRef(fleetRef).Cursor(cursor).Size(size).Execute()

List Invitations



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
	fleetRef := "fleetRef_example" // string | Limit results to specific fleet reference (optional)
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvitationsAPI.ListInvitations(context.Background()).FleetRef(fleetRef).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvitationsAPI.ListInvitations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInvitations`: CursorPageCustomizedInvitationRead
	fmt.Fprintf(os.Stdout, "Response from `InvitationsAPI.ListInvitations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fleetRef** | **string** | Limit results to specific fleet reference | 
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 500]

### Return type

[**CursorPageCustomizedInvitationRead**](CursorPageCustomizedInvitationRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

