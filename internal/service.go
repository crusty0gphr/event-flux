package internal

import (
	"context"

	"github.com/event-flux/domain"
)

type Service struct {
	repo Repository
}

func (s Service) GetByID(ctx context.Context, id string) (*domain.Event, error) {
	return s.repo.GetByID(ctx, id)
}

func (s Service) GetAll(ctx context.Context) ([]domain.Event, error) {
	return s.repo.GetAll(ctx)
}

func (s Service) GetByFilter(ctx context.Context, filters map[string]string) ([]domain.Event, error) {
	return s.repo.GetByFilter(ctx, filters)
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
