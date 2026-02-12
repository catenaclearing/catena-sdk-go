package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/catenaclearing/catena-sdk-go"
	telematicsapi "github.com/catenaclearing/catena-sdk-go/gen/telematics"
	"github.com/catenaclearing/catena-sdk-go/pagination"
)

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
		log.Printf("Failed to authenticate: %v", err)
		return
	}

	fmt.Println("Successfully authenticated!")

	// 2. List Vehicle Locations (Paginated)
	// We'll fetch vehicle locations for the last 24 hours.
	fmt.Println("\n--- Ingesting Vehicle Locations ---")

	startTime := time.Now().Add(-24 * time.Hour)
	opts := pagination.ListVehicleLocationsPaginationOptions{
		FromDatetime: &startTime,
	}

	count := 0
	err := pagination.ListVehicleLocationsEach(
		c,
		context.Background(),
		opts,
		func(loc telematicsapi.VehicleLocationRead) error {
			count++

			// Print details for every 100th location to avoid spamming the console
			if count%100 == 0 || count == 1 {
				vehicleID := loc.GetVehicleId()
				occurredAt := loc.GetOccurredAt()

				fmt.Printf("Location %d: Vehicle %s @ %s\n",
					count,
					vehicleID,
					occurredAt,
				)
			}

			return nil // Continue processing
		},
	)

	if err != nil {
		log.Printf("Error listing vehicle locations: %v", err)
		return
	}

	fmt.Printf("\nTotal locations ingested: %d\n", count)
}
