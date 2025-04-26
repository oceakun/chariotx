package app

import (
    "log"
    "net"

    pb "github.com/oceakun/chariotx/services/location/generated"
    "google.golang.org/grpc"
)

func StartGRPCServer() {
    cassClient, err := NewCassandraClient([]string{"127.0.0.1"}, "location_keyspace")
    if err != nil {
        log.Fatalf("Cassandra connection failed: %v", err)
    }

    kafkaProducer := NewKafkaProducer("localhost:9092", "location-pings")

    server := &LocationServiceServer{
        Cassandra: cassClient,
        Kafka:     kafkaProducer,
    }

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterLocationServiceServer(grpcServer, server)

    log.Println("gRPC server started on :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
