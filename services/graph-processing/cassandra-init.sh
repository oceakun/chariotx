#!/bin/bash
echo "Creating keyspace and table..."

docker exec -it cassandra cqlsh -e "
CREATE KEYSPACE IF NOT EXISTS segment_keyspace 
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

USE segment_keyspace;

CREATE TABLE IF NOT EXISTS segments (
    id UUID PRIMARY KEY,
    start_latitude DOUBLE,
    start_longitude DOUBLE,
    end_latitude DOUBLE,
    end_longitude DOUBLE,
    metadata TEXT
);"

echo "Keyspace and table created successfully."
