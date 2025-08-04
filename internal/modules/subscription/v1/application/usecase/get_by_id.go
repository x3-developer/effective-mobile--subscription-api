package usecase

import (
	"context"
	"subscriptions/internal/modules/subscription/v1/domain/model"
	"subscriptions/internal/modules/subscription/v1/domain/repo"
)

type GetByIdUseCase interface {
	Execute(ctx context.Context, id uint) (*model.Subscription, error)
}

type getByIdUseCase struct {
	repo repo.Repository
}

func NewGetByIdUseCase(repo repo.Repository) GetByIdUseCase {
	return &getByIdUseCase{
		repo: repo,
	}
}

func (u *getByIdUseCase) Execute(ctx context.Context, id uint) (*model.Subscription, error) {
	existsModel, err := u.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return existsModel, nil
}
