package internal

import (
	"context"

	"github.com/event-flux/domain"
)

type Repository interface {
	GetByID(ctx context.Context, id string) (*domain.Event, error)
	GetAll(ctx context.Context) ([]domain.Event, error)
	GetByFilter(ctx context.Context, filter string) ([]domain.Event, error)
}

type Service struct {
	repo Repository
}

func (s Service) GetByID(ctx context.Context, id string) (*domain.Event, error) {
	return s.repo.GetByID(ctx, id)
}

func (s Service) GetAll(ctx context.Context) ([]domain.Event, error) {
	return s.repo.GetAll(ctx)
}

func (s Service) GetByFilter(ctx context.Context, filter string) ([]domain.Event, error) {
	return s.repo.GetByFilter(ctx, filter)
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
