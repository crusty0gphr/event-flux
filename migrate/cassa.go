package migrate

import (
	"fmt"

	"github.com/gocql/gocql"
)

const (
	queryDown = "DROP TABLE IF EXISTS events;"
)

func CassandraMigrateUP(session *gocql.Session) error {
	keyspaceCreate := `
		CREATE KEYSPACE IF NOT EXISTS eventflux
		WITH replication = {
		  'class': 'SimpleStrategy',
		  'replication_factor': 3
		};
		`
	if err := session.Query(keyspaceCreate).Exec(); err != nil {
		return fmt.Errorf("cassandra migrate failed: %w :keyspaceCreate", err)
	}

	tableCreate := `
		CREATE TABLE IF NOT EXISTS eventflux.events (
			id UUID PRIMARY KEY,
			browser_fingerprint BIGINT,
			canvas_fingerprint BIGINT,
			created_at TIMESTAMP,
			device_language TEXT,
			device_timezone INT,
			event_name TEXT,
			font_fingerprint BIGINT,
			incognito BOOLEAN,
			ip TEXT,
			periodic_wave BIGINT,
			processed BOOLEAN,
			screen_resolution TEXT,
			session TEXT,
			status TEXT,
			storage BIGINT,
			updated_at TIMESTAMP,
			user_agent TEXT,
			user_id BIGINT,
			utm TEXT,
			webgl_fingerprint BIGINT
		);
		`
	if err := session.Query(tableCreate).Exec(); err != nil {
		return fmt.Errorf("cassandra migrate failed: %w :tableCreate", err)
	}
	return nil
}

func CassandraMigrateDOWN(session *gocql.Session) error {
	tableDrop := "DROP TABLE IF EXISTS eventflux.events;"

	if err := session.Query(tableDrop).Exec(); err != nil {
		return fmt.Errorf("cassandra migrate failed: %w", err)
	}
	return nil
}
