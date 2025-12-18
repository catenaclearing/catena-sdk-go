# Catena SDK for Go

[![CI](https://github.com/catenaclearing/catena-sdk-go/workflows/CI/badge.svg)](https://github.com/catenaclearing/catena-sdk-go/actions)
[![Go Reference](https://pkg.go.dev/badge/github.com/catenaclearing/catena-sdk-go.svg)](https://pkg.go.dev/github.com/catenaclearing/catena-sdk-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/catenaclearing/catena-sdk-go)](https://goreportcard.com/report/github.com/catenaclearing/catena-sdk-go)

Official Go SDK for the Catena Telematics API. This SDK provides type-safe Go clients generated from OpenAPI specifications for all Catena API services.

## Features

- 🔄 **Auto-generated from OpenAPI specs** - Always up-to-date with the latest API changes
- 📦 **Modular packages** - Separate clients for integrations, orgs, telematics, and notifications
- 🛡️ **Type-safe** - Full Go type safety with generated models and interfaces
- 🔧 **Deterministic generation** - Reproducible builds using pinned tooling versions
- ✅ **Well-tested** - Comprehensive CI/CD pipeline with linting and testing

## Installation

```bash
go get github.com/catenaclearing/catena-sdk-go
```

Individual packages:

```bash
# For integrations API
go get github.com/catenaclearing/catena-sdk-go/gen/integrations

# For orgs API
go get github.com/catenaclearing/catena-sdk-go/gen/orgs

# For telematics API
go get github.com/catenaclearing/catena-sdk-go/gen/telematics

# For notifications API
go get github.com/catenaclearing/catena-sdk-go/gen/notifications
```

## Quick Start

### Basic Usage

The SDK provides a unified client that handles authentication (OAuth2 Client Credentials), token refreshing, and access to all API services.

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/catenaclearing/catena-sdk-go"
)

func main() {
    ctx := context.Background()

    // Initialize the client
    // WithBaseURL is optional and defaults to "https://api.catenatelematics.com"
    client := catena.NewClient(
        catena.WithBaseURL("https://api.catenatelematics.com"),
    )

    // Authenticate
    // This exchanges your client credentials for an access token,
    // caches it, and sets up automatic refreshing.
    err := client.Authenticate(ctx, "your-client-id", "your-client-secret")
    if err != nil {
        log.Fatal(err)
    }

    // Access identity info from the decoded JWT
    fmt.Printf("Logged in as user: %s, Org: %s\n", client.UserID(), client.OrgName())

    // Make API calls
    // The client automatically injects the Authorization header
    // and refreshes the token if it expires.
    
    // Example: List integrations
    resp, httpResp, err := client.Integrations().DefaultApi.ListIntegrations(ctx).Execute()
    if err != nil {
        log.Fatalf("Error calling API: %v", err)
    }
    defer httpResp.Body.Close()
    
    fmt.Printf("Response: %+v\n", resp)
}
```

### Accessing API Services

The client provides accessors for all generated API services:

```go
// Integrations API
client.Integrations()

// Organizations API
client.Orgs()

## Pagination

The SDK provides high-level wrapper methods for paginated operations that handle cursor management and streaming automatically. These methods are generated and ensure efficient memory usage by yielding items one by one.

### Streaming Usage

Use the `*Each` methods from the `pagination` package to stream items. The callback function is called for each item. If the callback returns an error, iteration stops.

```go
import "github.com/catenaclearing/catena-sdk-go/pagination"

// Helper for pointers
func ptr[T any](v T) *T { return &v }

// List connections with pagination
opts := pagination.ListConnectionsPaginationOptions{
    Limit: ptr(100), // Optional: set page size
}

err := pagination.ListConnectionsEach(client, ctx, opts, func(conn integrationsapi.ConnectionRead) error {
    fmt.Printf("Connection: %s (%s)\n", conn.SourceName, conn.Id)
    return nil
})

if err != nil {
    log.Fatal(err)
}
```

The wrappers automatically:
- Fetch pages using the underlying generated client
- Extract items and the next cursor
- Handle empty pages and end-of-stream conditions
- Respect context cancellation

### Authentication & Security

The SDK handles OAuth2 authentication automatically.

**Note on JWT Claims:** The SDK exposes methods like `UserID()` and `OrgID()` which decode the access token's JWT payload. These claims are decoded **without verification** and are provided for convenience only (e.g., for UI display or logging). **Do not use these claims for security decisions** within your application; always validate against the API or a verified token if needed.

### Configuration Options

```go
client := catena.NewClient(
    // Set a custom base URL (default: https://api.catenatelematics.com)
    catena.WithBaseURL("https://staging-api.catenatelematics.com"),
    
    // Use a custom HTTP client
    catena.WithHTTPClient(myCustomHTTPClient),
)
```

## API Packages

This SDK provides four separate client packages:

| Package | Description | Import Path |
|---------|-------------|-------------|
| `integrationsapi` | Integrations API client | `github.com/catenaclearing/catena-sdk-go/gen/integrations` |
| `orgsapi` | Organizations API client | `github.com/catenaclearing/catena-sdk-go/gen/orgs` |
| `telematicsapi` | Telematics API client | `github.com/catenaclearing/catena-sdk-go/gen/telematics` |
| `notificationsapi` | Notifications API client | `github.com/catenaclearing/catena-sdk-go/gen/notifications` |


## Versioning

This SDK follows [Semantic Versioning](https://semver.org/):

- **Major version** bumps indicate breaking changes in the generated API
- **Minor version** bumps indicate new features or non-breaking API changes
- **Patch version** bumps indicate bug fixes or internal improvements

### API Stability

**Important:** This SDK is auto-generated from OpenAPI specifications. The API structure may change when:

- Catena updates their OpenAPI specs
- New endpoints are added or removed
- Request/response schemas change

**Recommendation:** Pin your dependency to a specific version in your `go.mod`:

```go
require github.com/catenaclearing/catena-sdk-go v1.2.3
```

## Webhooks

The SDK provides a `webhooks` package to easily consume and route webhook events.

### Router Usage

The `Router` handles parsing, validation, and dispatching to typed handlers.

```go
package main

import (
    "context"
    "log"
    "net/http"

    "github.com/catenaclearing/catena-sdk-go/gen/notifications"
    "github.com/catenaclearing/catena-sdk-go/webhooks"
)

func main() {
    router := webhooks.NewRouter()

    // Register typed handler
    router.Handle("webhook.deleted", func(ctx context.Context, env *webhooks.Envelope, data []notificationsapi.WebhookDeleted) error {
        for _, item := range data {
            log.Println("deleted webhook id:", item.Id)
        }
        return nil
    })

    http.Handle("/catena/webhooks", router)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

The router automatically handles:
- **404 Not Found** for unregistered events
- **400 Bad Request** for invalid JSON or decode errors
- **500 Internal Server Error** if a handler returns an error

### Manual Parsing

You can also parse payloads manually if you don't want to use the router.

```go
env, payload, err := webhooks.ParseBytes(body)
if err != nil {
    // Handle error (invalid JSON, etc.)
}

switch p := payload.(type) {
case []notificationsapi.WebhookDeleted:
    // Handle typed slice
case webhooks.UnknownEvent:
    // Handle unknown event (forward compatibility)
    log.Printf("Unknown event: %s", p.EventName)
}
```

Note: `data` in the envelope is always an array, so handlers receive a slice of the generated type.

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Code of Conduct

This project adheres to a [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## Security

For security concerns, please see our [Security Policy](SECURITY.md).

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Links

- [Catena Telematics API Documentation](https://api.catenatelematics.com)
- [OpenAPI Generator](https://openapi-generator.tech/)
- [Redocly CLI](https://redocly.com/docs/cli/)
- [Go Package Documentation](https://pkg.go.dev/github.com/catenaclearing/catena-sdk-go)

## Acknowledgments

This SDK is generated using:
- [OpenAPI Generator](https://openapi-generator.tech/) - v7.2.0
- [Redocly CLI](https://redocly.com/docs/cli/) - for spec validation and bundling
