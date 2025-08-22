package config

import (
	"os"
)

type CassandraConfig struct {
	Host     string
	Keyspace string
}

type Config struct {
	Port      string
	Cassandra CassandraConfig
}

func Load() Config {
	return Config{
		Port: getEnv("PORT", "8080"),
		Cassandra: CassandraConfig{
			Host:     getEnv("CASSANDRA_HOST", "localhost"),
			Keyspace: getEnv("CASSANDRA_KEYSPACE", "segment_keyspace"),
		},
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
