package usecase

import (
	"context"
	"subscriptions/internal/modules/subscription/v1/domain/model"
	"subscriptions/internal/modules/subscription/v1/domain/repo"
	"subscriptions/internal/shared/lib/res"
)

type CreateUseCase interface {
	Execute(ctx context.Context, model *model.Subscription) (*model.Subscription, []res.ErrorField, error)
}

type createUseCase struct {
	repo repo.Repository
}

func NewCreateUseCase(repo repo.Repository) CreateUseCase {
	return &createUseCase{
		repo: repo,
	}
}

func (u *createUseCase) Execute(ctx context.Context, model *model.Subscription) (*model.Subscription, []res.ErrorField, error) {
	createdModel, err := u.repo.Create(ctx, model)
	if err != nil {
		return nil, nil, err
	}

	return createdModel, nil, nil
}
