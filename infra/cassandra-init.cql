-- Location service keyspace
CREATE KEYSPACE IF NOT EXISTS location_keyspace
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

USE location_keyspace;

CREATE TABLE IF NOT EXISTS user_locations (
    user_id   text,
    timestamp bigint,
    latitude  double,
    longitude double,
    PRIMARY KEY (user_id, timestamp)
);

-- Graph-processing / segment keyspace
CREATE KEYSPACE IF NOT EXISTS segment_keyspace
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};

USE segment_keyspace;

CREATE TABLE IF NOT EXISTS road_segments (
    id       text PRIMARY KEY,
    source   text,
    target   text,
    distance double
);
