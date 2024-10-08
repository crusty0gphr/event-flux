package internal

import "github.com/gocql/gocql"

type CassandraRepo struct {
	session *gocql.Session
}

func NewCassandraRepo(session *gocql.Session) *CassandraRepo {
	return &CassandraRepo{session: session}
}
