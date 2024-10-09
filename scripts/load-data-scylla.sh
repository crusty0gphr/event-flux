#!/bin/bash

# Wait for Cassandra to be available
until echo 'SELECT now() FROM system.local;' | cqlsh scylla; do
  echo "ScyllaDB is unavailable - waiting..."
  sleep 5
done

# Create the keyspace
cqlsh scylla -e "
CREATE KEYSPACE IF NOT EXISTS eventflux
WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}
"

# Create the table inside the keyspace
cqlsh scylla -e "
USE eventflux;

DROP TABLE IF EXISTS events;

CREATE TABLE IF NOT EXISTS events (
    id UUID,
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
    webgl_fingerprint BIGINT,
    PRIMARY KEY ((event_name), created_at)
) WITH CLUSTERING ORDER BY (created_at DESC);

CREATE INDEX IF NOT EXISTS idx_id ON events (id)
"

# Load data from the CSV
cqlsh scylla -e "
COPY eventflux.events (
    id, browser_fingerprint, canvas_fingerprint, created_at, device_language,
    device_timezone, event_name, font_fingerprint, incognito, ip, periodic_wave,
    processed, screen_resolution, session, status, storage, updated_at, user_agent,
    user_id, utm, webgl_fingerprint
) FROM '/data/events.csv' WITH HEADER = TRUE
"

echo "Data loaded successfully into ScyllaDB"
