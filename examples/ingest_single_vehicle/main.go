package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/catenaclearing/catena-sdk-go"
	telematicsapi "github.com/catenaclearing/catena-sdk-go/gen/telematics"
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

	// 2. Find a Vehicle
	// We'll list vehicles to find one to query.
	fmt.Println("\n--- Finding a Vehicle ---")
	var vehicleID string

	err := pagination.ListVehiclesEach(
		c,
		context.Background(),
		pagination.ListVehiclesPaginationOptions{},
		func(vehicle telematicsapi.Vehicle) error {
			vehicleID = vehicle.GetId()
			fmt.Printf("Found Vehicle: %s (Name: %s)\n", vehicleID, vehicle.GetVehicleName())
			return ErrStop // Stop after finding the first one
		},
	)

	if err != nil && err != ErrStop {
		log.Fatalf("Error listing vehicles: %v", err)
	}

	if vehicleID == "" {
		log.Println("No vehicles found.")
		return
	}

	// 3. Get Vehicle Details
	// Now we'll fetch the full details for this specific vehicle.
	fmt.Printf("\n--- Getting Details for Vehicle %s ---\n", vehicleID)

	vehicle, resp, err := c.Telematics().FleetOperationsTrackingAPI.
		GetVehicle(context.Background(), vehicleID).
		Execute()

	if err != nil {
		log.Fatalf("Error getting vehicle details: %v\nResponse: %v", err, resp)
	}

	fmt.Printf("Vehicle ID: %s\n", vehicle.GetId())
	fmt.Printf("Name: %s\n", vehicle.GetVehicleName())
	fmt.Printf("VIN: %s\n", vehicle.GetVin())
	fmt.Printf("Make/Model/Year: %s %s %d\n", vehicle.GetOem(), vehicle.GetModelType(), vehicle.GetModelYear())
	fmt.Printf("License Plate: %s\n", vehicle.GetLicensePlateNumber())
}
