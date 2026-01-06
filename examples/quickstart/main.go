// go get github.com/catenaclearing/catena-sdk-go@latest

package main

import (
	"context"
	"fmt"
	"log"

	catena "github.com/catenaclearing/catena-sdk-go"
)

func main() {
	ctx := context.Background()

	client := catena.NewClient(
		catena.WithBaseURL("https://api.catenatelematics.com"),
	)

	if err := client.Authenticate(ctx, "your-client-id", "your-client-secret"); err != nil {
		log.Fatal(err)
	}

	// Example: list connections
	resp, httpResp, err := client.Integrations().ConnectionsAPI.ListConnections(ctx).Execute()
	if err != nil {
		log.Fatalf("API error: %v", err)
	}
	defer httpResp.Body.Close()

	fmt.Printf("Connections: %+v\n", resp)
}
