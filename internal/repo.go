package internal

import (
	"context"
	"fmt"

	eventflux "github.com/event-flux"
	"github.com/event-flux/db"
	"github.com/event-flux/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (*domain.Event, error)
	GetAll(ctx context.Context) ([]domain.Event, error)
	GetByFilter(ctx context.Context, filters map[string]string) ([]domain.Event, error)
}

type repoBuilder func(config *eventflux.Config) (Repository, error)

var availableTypes = map[string]repoBuilder{
	"cassandra": makeCassandra,
	"scylla":    makeScylla,
}

func BuildRepository(config *eventflux.Config) (Repository, error) {
	builderFunc, ok := availableTypes[config.DbDriverType]
	if !ok {
		return nil, fmt.Errorf("repo factory: invaid driver type %s", config.DbDriverType)
	}

	return builderFunc(config)
}

func makeCassandra(config *eventflux.Config) (Repository, error) {
	return makeGOCQL(config.CassandraHost)
}

func makeScylla(config *eventflux.Config) (Repository, error) {
	return makeGOCQL(config.ScyllaHost)
}

func makeGOCQL(host string) (Repository, error) {
	session, err := db.NewGOCQLSession(host)
	if err != nil {
		return nil, err
	}
	return NewCQLRepo(session), nil
}
