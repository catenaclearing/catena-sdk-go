package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/catenaclearing/catena-sdk-go"
	notificationsapi "github.com/catenaclearing/catena-sdk-go/gen/notifications"
	"github.com/catenaclearing/catena-sdk-go/pagination"
)

var ErrStop = errors.New("stop iteration")

func main() {
	// 1. Initialize the client
	clientID := os.Getenv("CATENA_CLIENT_ID")
	clientSecret := os.Getenv("CATENA_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Please set CATENA_CLIENT_ID and CATENA_CLIENT_SECRET environment variables")
	}

	c := catena.NewClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.Authenticate(ctx, clientID, clientSecret); err != nil {
		log.Fatalf("Failed to authenticate: %v", err)
	}

	fmt.Println("Successfully authenticated!")

	// 2. Find a Webhook
	fmt.Println("\n--- Finding a Webhook ---")
	var webhookID string

	err := pagination.ListWebhookSubscriptionsEach(
		c,
		context.Background(),
		pagination.ListWebhookSubscriptionsPaginationOptions{},
		func(webhook notificationsapi.WebhookRead) error {
			webhookID = webhook.GetId()
			fmt.Printf("Found Webhook: %s (URL: %s)\n", webhookID, webhook.GetUrl())
			return ErrStop // Stop after finding the first one
		},
	)

	if err != nil && err != ErrStop {
		log.Fatalf("Error listing webhooks: %v", err)
	}

	if webhookID == "" {
		log.Println("No webhooks found. Please create one first.")
		return
	}

	// 3. Get Webhook Metrics
	fmt.Printf("\n--- Getting Metrics for Webhook %s ---\n", webhookID)

	metrics, resp, err := c.Notifications().WebhookSubscriptionsAPI.
		GetWebhookSubscriptionMetrics(context.Background(), webhookID).
		Execute()

	if err != nil {
		log.Fatalf("Error getting webhook metrics: %v\nResponse: %v", err, resp)
	}

	successAttempts := metrics.GetHttpSuccessAttempts()
	var success24h int32
	if val := successAttempts.Var24h.Get(); val != nil {
		success24h = *val
	}

	failureAttempts := metrics.GetHttpFailureAttempts()
	var failed24h int32
	if val := failureAttempts.Var24h.Get(); val != nil {
		failed24h = *val
	}

	fmt.Printf("24h Success: %d\n", success24h)
	fmt.Printf("24h Failed: %d\n", failed24h)
}
