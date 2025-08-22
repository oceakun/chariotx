#!/bin/bash
echo "Creating segment keyspace and table..."

docker exec -i cassandra cqlsh -e "
CREATE KEYSPACE IF NOT EXISTS segment_keyspace
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

USE segment_keyspace;

CREATE TABLE IF NOT EXISTS road_segments (
    segment_id bigint PRIMARY KEY,
    start_lat double,
    start_lon double,
    end_lat double,
    end_lon double,
    road_type text,
    length_km double
);"

echo "Segment keyspace and table created successfully."