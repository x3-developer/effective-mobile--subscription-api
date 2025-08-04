package usecase

import (
	"context"
	"subscriptions/internal/modules/subscription/v1/domain/repo"
	"subscriptions/internal/modules/subscription/v1/domain/vo"
)

type GetTotalCostUseCase interface {
	Execute(ctx context.Context, filter *vo.TotalCostFilter) (float64, error)
}

type getTotalCostUseCase struct {
	repo repo.Repository
}

func NewGetTotalCostUseCase(repo repo.Repository) GetTotalCostUseCase {
	return &getTotalCostUseCase{
		repo: repo,
	}
}

func (u *getTotalCostUseCase) Execute(ctx context.Context, filter *vo.TotalCostFilter) (float64, error) {
	return u.repo.CalculateTotalCost(ctx, filter.StartDate, filter.EndDate, filter.UserId, filter.Name)
}
