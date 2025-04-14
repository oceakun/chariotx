package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/oceakun/chariotx/services/location/generated"
)

func main() {
	// Create a new gRPC client connection using the new API
	clientConn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	defer clientConn.Close()

	// Create client from generated proto
	client := pb.NewLocationServiceClient(clientConn)

	// Prepare test data
	req := &pb.LocationData{
		UserId:    "user456",
		Lat:       37.7749,
		Lng:       -122.4194,
		Timestamp: time.Now().Unix(),
	}

	// Call the RPC
	resp, err := client.SendLocation(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling SendLocation: %v", err)
	}

	log.Printf("Ack received: %v", resp.Success)
}
