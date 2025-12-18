package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/catenaclearing/catena-sdk-go"
	notificationsapi "github.com/catenaclearing/catena-sdk-go/gen/notifications"
)

func stringPtr(s string) *string {
	return &s
}

func main() {
	// 1. Initialize the client
	clientID := os.Getenv("CATENA_CLIENT_ID")
	clientSecret := os.Getenv("CATENA_CLIENT_SECRET")
	webhookURL := os.Getenv("WEBHOOK_URL")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Please set CATENA_CLIENT_ID and CATENA_CLIENT_SECRET environment variables")
	}
	if webhookURL == "" {
		log.Fatal("Please set WEBHOOK_URL environment variable")
	}

	c := catena.NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.Authenticate(ctx, clientID, clientSecret); err != nil {
		log.Fatalf("Failed to authenticate: %v", err)
	}

	fmt.Println("Successfully authenticated!")

	// 2. Create Webhook Subscription
	// We'll subscribe to all vehicle location events.
	fmt.Println("\n--- Creating Webhook Subscription ---")

	// Define the event we want to subscribe to.
	// We use the catch-all event for vehicle locations ("vehicle_location.*").
	catchAllEvent := notificationsapi.WEBHOOKCATCHALLEVENTS_VEHICLE_LOCATION
	eventName := notificationsapi.WebhookEventName{
		WebhookCatchAllEvents: &catchAllEvent,
	}

	// Define the subscription details
	newWebhook := notificationsapi.WebhookCreate{
		Url:       webhookURL,
		EventName: eventName,
		Secret:    *notificationsapi.NewNullableString(stringPtr("my-secret-signing-key")),
	}

	// Execute the request
	createdWebhook, resp, err := c.Notifications().WebhookSubscriptionsAPI.
		CreateWebhookSubscription(context.Background()).
		WebhookCreate(newWebhook).
		Execute()

	if err != nil {
		log.Fatalf("Error creating webhook: %v\nResponse: %v", err, resp)
	}

	if createdWebhook.Webhook == nil {
		log.Fatalf("Unexpected response format: %v", createdWebhook)
	}

	fmt.Printf("Created Webhook ID: %s\n", createdWebhook.Webhook.GetId())
	fmt.Printf("Target URL: %s\n", createdWebhook.Webhook.GetUrl())
	fmt.Printf("Secret: %s\n", createdWebhook.Webhook.GetSecret())
}
