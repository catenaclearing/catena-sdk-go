package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	notificationsapi "github.com/catenaclearing/catena-sdk-go/gen/notifications"
	"github.com/catenaclearing/catena-sdk-go/webhooks"
)

func main() {
	// 1. Create a new Webhook Router
	router := webhooks.NewRouter()

	// 2. Register Handlers
	// We'll register a handler for vehicle location events.
	// The event name must match exactly what Catena sends.
	// Common patterns are "resource.action" (e.g., "vehicle_location.added").

	router.Handle("vehicle_location.added", func(ctx context.Context, env *webhooks.Envelope, data []notificationsapi.BaseVehicleLocation) error {
		fmt.Printf("Received %d vehicle locations (Webhook ID: %s)\n", len(data), env.WebhookID)

		for _, loc := range data {
			fmt.Printf(" - Vehicle %s at %s\n", loc.GetVehicleId(), loc.GetOccurredAt())
		}

		return nil
	})

	// You can register multiple handlers for different events
	router.Handle("connection.created", func(ctx context.Context, env *webhooks.Envelope, data []notificationsapi.BaseConnectionEvent) error {
		fmt.Printf("New connection created! ID: %s\n", data[0].GetId())
		return nil
	})

	// 3. Start the Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Starting webhook receiver on port %s...\n", port)

	// The router implements http.Handler, so you can pass it directly to ListenAndServe
	// or mount it on a subpath using http.Handle("/webhooks", router).
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
