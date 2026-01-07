# \PartnersAPI

All URIs are relative to *https://api.catenatelematics.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePartner**](PartnersAPI.md#CreatePartner) | **Post** /v2/orgs/partners | Create Partner
[**CreatePartnerProperties**](PartnersAPI.md#CreatePartnerProperties) | **Post** /v2/orgs/partners/{partner_id}/properties | Create Partner Properties
[**DeletePartner**](PartnersAPI.md#DeletePartner) | **Delete** /v2/orgs/partners/{partner_id} | Delete Partner
[**DeletePartnerProperty**](PartnersAPI.md#DeletePartnerProperty) | **Delete** /v2/orgs/partners/{partner_id}/properties/{property_id} | Delete Partner Property
[**GetPartner**](PartnersAPI.md#GetPartner) | **Get** /v2/orgs/partners/{partner_id} | Get Partner
[**ListPartnerProperties**](PartnersAPI.md#ListPartnerProperties) | **Get** /v2/orgs/partners/{partner_id}/properties | List Partner Properties
[**ListPartners**](PartnersAPI.md#ListPartners) | **Get** /v2/orgs/partners | List Partners
[**UpdatePartner**](PartnersAPI.md#UpdatePartner) | **Patch** /v2/orgs/partners/{partner_id} | Update Partner



## CreatePartner

> PartnerRead CreatePartner(ctx).PartnerCreate(partnerCreate).Execute()

Create Partner



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
	partnerCreate := *openapiclient.NewPartnerCreate("Name_example") // PartnerCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PartnersAPI.CreatePartner(context.Background()).PartnerCreate(partnerCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.CreatePartner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePartner`: PartnerRead
	fmt.Fprintf(os.Stdout, "Response from `PartnersAPI.CreatePartner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partnerCreate** | [**PartnerCreate**](PartnerCreate.md) |  | 

### Return type

[**PartnerRead**](PartnerRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePartnerProperties

> []PartnerPropertyRead CreatePartnerProperties(ctx, partnerId).PartnerPropertyCreate(partnerPropertyCreate).Execute()

Create Partner Properties



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
	partnerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	partnerPropertyCreate := []openapiclient.PartnerPropertyCreate{*openapiclient.NewPartnerPropertyCreate(openapiclient.PartnerPropertyKeyEnum("default_failure_redirect_url"), "Value_example")} // []PartnerPropertyCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PartnersAPI.CreatePartnerProperties(context.Background(), partnerId).PartnerPropertyCreate(partnerPropertyCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.CreatePartnerProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePartnerProperties`: []PartnerPropertyRead
	fmt.Fprintf(os.Stdout, "Response from `PartnersAPI.CreatePartnerProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePartnerPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partnerPropertyCreate** | [**[]PartnerPropertyCreate**](PartnerPropertyCreate.md) |  | 

### Return type

[**[]PartnerPropertyRead**](PartnerPropertyRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePartner

> DeletePartner(ctx, partnerId).Execute()

Delete Partner



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
	partnerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PartnersAPI.DeletePartner(context.Background(), partnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.DeletePartner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePartnerRequest struct via the builder pattern


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


## DeletePartnerProperty

> DeletePartnerProperty(ctx, partnerId, propertyId).Execute()

Delete Partner Property



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
	partnerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	propertyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PartnersAPI.DeletePartnerProperty(context.Background(), partnerId, propertyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.DeletePartnerProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** |  | 
**propertyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePartnerPropertyRequest struct via the builder pattern


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


## GetPartner

> PartnerRead GetPartner(ctx, partnerId).Execute()

Get Partner



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
	partnerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PartnersAPI.GetPartner(context.Background(), partnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.GetPartner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPartner`: PartnerRead
	fmt.Fprintf(os.Stdout, "Response from `PartnersAPI.GetPartner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PartnerRead**](PartnerRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPartnerProperties

> []PartnerPropertyRead ListPartnerProperties(ctx, partnerId).Execute()

List Partner Properties



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
	partnerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PartnersAPI.ListPartnerProperties(context.Background(), partnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.ListPartnerProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPartnerProperties`: []PartnerPropertyRead
	fmt.Fprintf(os.Stdout, "Response from `PartnersAPI.ListPartnerProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPartnerPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PartnerPropertyRead**](PartnerPropertyRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPartners

> CursorPageCustomizedPartnerRead ListPartners(ctx).Cursor(cursor).Size(size).Execute()

List Partners



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
	cursor := "cursor_example" // string | Cursor for the next page (optional)
	size := int32(56) // int32 | Page size (optional) (default to 500)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PartnersAPI.ListPartners(context.Background()).Cursor(cursor).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.ListPartners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPartners`: CursorPageCustomizedPartnerRead
	fmt.Fprintf(os.Stdout, "Response from `PartnersAPI.ListPartners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPartnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** | Cursor for the next page | 
 **size** | **int32** | Page size | [default to 500]

### Return type

[**CursorPageCustomizedPartnerRead**](CursorPageCustomizedPartnerRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartner

> PartnerRead UpdatePartner(ctx, partnerId).PartnerUpdate(partnerUpdate).Execute()

Update Partner



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
	partnerId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	partnerUpdate := *openapiclient.NewPartnerUpdate() // PartnerUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PartnersAPI.UpdatePartner(context.Background(), partnerId).PartnerUpdate(partnerUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnersAPI.UpdatePartner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePartner`: PartnerRead
	fmt.Fprintf(os.Stdout, "Response from `PartnersAPI.UpdatePartner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePartnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partnerUpdate** | [**PartnerUpdate**](PartnerUpdate.md) |  | 

### Return type

[**PartnerRead**](PartnerRead.md)

### Authorization

[Bearer](../README.md#Bearer), [Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

