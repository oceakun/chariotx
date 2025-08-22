package cassandra

import (
	"log"
	"github.com/gocql/gocql"
	"github.com/oceakun/chariotx/services/graph-processing/config"
)

func Connect(cfg config.CassandraConfig) *gocql.Session {
	cluster := gocql.NewCluster(cfg.Host)
	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = gocql.One

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to Cassandra: %v", err)
	}
	return session
}