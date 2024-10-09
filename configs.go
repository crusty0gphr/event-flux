package event_flux

import (
	"fmt"
	"os"
)

type Config struct {
	Host          string
	Port          string
	CassandraHost string
	ScyllaHost    string
	DbDriverType  string
}

func (c *Config) BuildAppHostUrl() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func LoadConfigs() *Config {
	return &Config{
		Host:          os.Getenv("APP_HOST"),
		Port:          os.Getenv("APP_PORT"),
		CassandraHost: os.Getenv("CASSANDRA_HOST"),
		ScyllaHost:    os.Getenv("SCYLLA_HOST"),
		DbDriverType:  os.Getenv("DB_DRIVER_TYPE"),
	}
}
