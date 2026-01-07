// go get github.com/catenaclearing/catena-sdk-go@latest

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	catena "github.com/catenaclearing/catena-sdk-go"
)

func main() {
	ctx := context.Background()

	clientID := os.Getenv("CATENA_CLIENT_ID")
	if clientID == "" {
		clientID = os.Getenv("CLIENT_ID")
	}
	clientSecret := os.Getenv("CATENA_CLIENT_SECRET")
	if clientSecret == "" {
		clientSecret = os.Getenv("CLIENT_SECRET")
	}
	if clientID == "" || clientSecret == "" {
		log.Fatal("Please set CLIENT_ID/CLIENT_SECRET or CATENA_CLIENT_ID/CATENA_CLIENT_SECRET")
	}

	client := catena.NewClient(
		catena.WithBaseURL("https://api.catenatelematics.com"),
	)

	if err := client.Authenticate(ctx, clientID, clientSecret); err != nil {
		log.Fatal(err)
	}

	// Example: list connections
	resp, httpResp, err := client.Integrations().ConnectionsAPI.ListConnections(ctx).Execute()
	if err != nil {
		log.Printf("API error: %v", err)
		return
	}
	defer httpResp.Body.Close()

	fmt.Printf("Connections: %+v\n", resp)
}
