package app

import (
    "context"
    "log"
    pb "github.com/oceakun/chariotx/services/location/generated"
)

type LocationServiceServer struct {
    pb.UnimplementedLocationServiceServer
    Cassandra *CassandraClient
    Kafka     *KafkaProducer
}

func (s *LocationServiceServer) SendLocation(ctx context.Context, req *pb.LocationData) (*pb.Ack, error) {
    log.Printf("Received location ping: %v", req)

    // Save to Cassandra
    err := s.Cassandra.SaveLocation(req.UserId, req.Timestamp, req.Lat, req.Lng)
    if err != nil {
        log.Printf("Failed to insert into Cassandra: %v", err)
        return &pb.Ack{Success: false}, err
    }
    log.Printf("Location saved: %+v", req)

    // Publish to Kafka
    err = s.Kafka.PublishLocation(req)
    if err != nil {
        log.Printf("Kafka publish failed: %v", err)
        return &pb.Ack{Success: false}, err
    }
    return &pb.Ack{Success: true}, nil
}