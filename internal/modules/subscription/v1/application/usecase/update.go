package usecase

import (
	"context"
	"errors"
	"subscriptions/internal/modules/subscription/v1/domain/model"
	"subscriptions/internal/modules/subscription/v1/domain/repo"
	"subscriptions/internal/shared/lib/res"
)

type UpdateUseCase interface {
	Execute(ctx context.Context, id uint, model *model.Subscription) (*model.Subscription, []res.ErrorField, error)
}

type updateUseCase struct {
	repo repo.Repository
}

func NewUpdateUseCase(repo repo.Repository) UpdateUseCase {
	return &updateUseCase{
		repo: repo,
	}
}

func (u *updateUseCase) Execute(ctx context.Context, id uint, model *model.Subscription) (*model.Subscription, []res.ErrorField, error) {
	//var validationErrors []res.ErrorField
	existsModel, err := u.repo.GetById(ctx, id)
	if err != nil {
		return nil, nil, err

	}
	if existsModel == nil {
		return nil, nil, errors.New("subscription not found")
	}

	updatedModel, err := u.repo.Update(ctx, model)
	if err != nil {
		return nil, nil, err
	}

	return updatedModel, nil, nil
}
