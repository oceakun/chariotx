package app

import (
    "log"
    "net"
    "os"

    pb "github.com/oceakun/chariotx/services/location/generated"
    "google.golang.org/grpc"
)

func getEnv(key, fallback string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return fallback
}

func StartGRPCServer() {
    cassHost := getEnv("CASSANDRA_HOST", "127.0.0.1")
    kafkaBroker := getEnv("KAFKA_BROKER", "localhost:9092")

    cassClient, err := NewCassandraClient([]string{cassHost}, "location_keyspace")
    if err != nil {
        log.Fatalf("Cassandra connection failed: %v", err)
    }

    kafkaProducer := NewKafkaProducer(kafkaBroker, "location-pings")

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
