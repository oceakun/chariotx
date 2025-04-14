package app

import (
    "github.com/gocql/gocql"
)

type CassandraClient struct {
    Session *gocql.Session
}

func NewCassandraClient(hosts []string, keyspace string) (*CassandraClient, error) {
    cluster := gocql.NewCluster(hosts...)
    cluster.Keyspace = keyspace
    cluster.Consistency = gocql.Quorum

    session, err := cluster.CreateSession()
    if err != nil {
        return nil, err
    }

    return &CassandraClient{Session: session}, nil
}

func (c *CassandraClient) SaveLocation(userID string, timestamp int64, lat, long float64) error {
    query := `INSERT INTO user_locations (user_id, timestamp, latitude, longitude) VALUES (?, ?, ?, ?)`
    return c.Session.Query(query, userID, timestamp, lat, long).Exec()
}
