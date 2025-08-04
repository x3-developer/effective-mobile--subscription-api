package usecase

import (
	"context"
	"subscriptions/internal/modules/subscription/v1/domain/model"
	"subscriptions/internal/modules/subscription/v1/domain/repo"
)

type DeleteUseCase interface {
	Execute(ctx context.Context, id uint) (*model.Subscription, error)
}

type deleteUseCase struct {
	repo repo.Repository
}

func NewDeleteUseCase(repo repo.Repository) DeleteUseCase {
	return &deleteUseCase{
		repo: repo,
	}
}

func (u *deleteUseCase) Execute(ctx context.Context, id uint) (*model.Subscription, error) {
	existsModel, err := u.repo.GetById(ctx, id)
	if existsModel == nil {
		return nil, err
	}

	err = u.repo.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return existsModel, nil
}
