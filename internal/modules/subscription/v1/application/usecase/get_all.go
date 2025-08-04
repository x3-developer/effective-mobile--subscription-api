package usecase

import (
	"context"
	"subscriptions/internal/modules/subscription/v1/domain/model"
	"subscriptions/internal/modules/subscription/v1/domain/repo"
)

type GetAllUseCase interface {
	Execute(ctx context.Context) ([]*model.Subscription, error)
}

type getAllUseCase struct {
	repo repo.Repository
}

func NewGetAllUseCase(repo repo.Repository) GetAllUseCase {
	return &getAllUseCase{
		repo: repo,
	}
}

func (u *getAllUseCase) Execute(ctx context.Context) ([]*model.Subscription, error) {
	models, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return models, nil
}
