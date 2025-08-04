package mapper

import (
	"github.com/google/uuid"
	"subscriptions/internal/modules/subscription/v1/application/dto"
	"subscriptions/internal/modules/subscription/v1/domain/model"
	"subscriptions/internal/shared/lib/val"
	"time"
)

func ToModelFromUpdateDTO(updateDto *dto.UpdateDTO, model *model.Subscription) *model.Subscription {
	if updateDto == nil || model == nil {
		return model
	}

	if updateDto.Name != nil && *updateDto.Name != "" {
		model.Name = *updateDto.Name
	}
	if updateDto.Price != nil {
		model.Price = *updateDto.Price
	}
	if updateDto.UserId != nil {
		userId, err := uuid.Parse(*updateDto.UserId)
		if err == nil {
			model.UserId = userId
		}

	}
	if updateDto.StartDate != nil {
		startDate, err := time.Parse(val.MonthYearLayout, *updateDto.StartDate)
		if err == nil {
			model.StartDate = startDate
		}
	}
	if updateDto.EndDate != nil {
		endDate, err := time.Parse(val.MonthYearLayout, *updateDto.EndDate)
		if err == nil {
			model.EndDate = &endDate
		}
	}

	return model
}
