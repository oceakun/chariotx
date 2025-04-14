package app

import (
	"context"
    "encoding/json"
    "github.com/segmentio/kafka-go"
    pb "github.com/oceakun/chariotx/services/location/generated"
    "time"
)

type KafkaProducer struct {
    Writer *kafka.Writer
}

func NewKafkaProducer(broker, topic string) *KafkaProducer {
    writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{broker},
        Topic:   topic,
        Balancer: &kafka.LeastBytes{},
        WriteTimeout: 5 * time.Second,
    })

    return &KafkaProducer{Writer: writer}
}

func (k *KafkaProducer) PublishLocation(loc *pb.LocationData) error {
    payload, err := json.Marshal(loc)
    if err != nil {
        return err
    }

    msg := kafka.Message{
        Key:   []byte(loc.UserId),
        Value: payload,
    }

    return k.Writer.WriteMessages(context.Background(), msg)
}
