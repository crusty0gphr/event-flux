package db

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

func NewCassandraSession(host string) (createSession *gocql.Session, err error) {
	for retries := 0; retries < 5; retries++ {
		cluster := gocql.NewCluster(host)
		cluster.Consistency = gocql.Quorum
		cluster.Timeout = 10 * time.Second

		createSession, err = cluster.CreateSession()
		if err == nil {
			return createSession, err
		}
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("cassandra driver failed: %w", err)
}
