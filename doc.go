// Package catena provides the Go SDK for the Catena API.
//
// The SDK provides a unified client that handles authentication, token management,
// and access to all Catena API services (Integrations, Organizations, Telematics, Notifications).
//
// Basic Usage:
//
//	import "github.com/catenaclearing/catena-sdk-go"
//
//	// Initialize the client
//	client := catena.NewClient()
//
//	// Authenticate
//	err := client.Authenticate(ctx, "client-id", "client-secret")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Access API services
//	integrations, _, err := client.Integrations().DefaultApi.ListIntegrations(ctx).Execute()
package catena
