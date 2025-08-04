package repo

import (
	"context"
	"github.com/google/uuid"
	"subscriptions/internal/modules/subscription/v1/domain/model"
	"time"
)

type Repository interface {
	Create(ctx context.Context, model *model.Subscription) (*model.Subscription, error)
	GetAll(ctx context.Context) ([]*model.Subscription, error)
	GetById(ctx context.Context, id uint) (*model.Subscription, error)
	Update(ctx context.Context, model *model.Subscription) (*model.Subscription, error)
	Delete(ctx context.Context, id uint) error
	CalculateTotalCost(ctx context.Context, startDate, endDate time.Time, userId *uuid.UUID, subscriptionName *string) (float64, error)
}
