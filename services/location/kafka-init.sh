#!/bin/bash
echo "Creating Kafka topic 'location-pings'..."
docker exec -it kafka kafka-topics --create --topic location-pings --partitions 1 --replication-factor 1 --if-not-exists --bootstrap-server kafka:9092
echo "Kafka topic 'location-pings' created successfully."