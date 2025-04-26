#!/bin/bash
echo "Creating keyspace and table..."
docker exec -it cassandra cqlsh -e "CREATE KEYSPACE IF NOT EXISTS location_keyspace WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}; USE location_keyspace; CREATE TABLE IF NOT EXISTS user_locations (user_id text, timestamp bigint, latitude double, longitude double, PRIMARY KEY (user_id, timestamp));"
echo "Keyspace and table created successfully."