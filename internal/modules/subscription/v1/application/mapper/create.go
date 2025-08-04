package mapper

import (
	"github.com/google/uuid"
	"subscriptions/internal/modules/subscription/v1/application/dto"
	"subscriptions/internal/modules/subscription/v1/domain/model"
	"subscriptions/internal/shared/lib/val"
	"time"
)

func ToModelFromCreateDTO(createDto *dto.CreateDTO) *model.Subscription {
	if createDto == nil {
		return nil
	}

	startDate, err := time.Parse(val.MonthYearLayout, createDto.StartDate)
	if err != nil {
		return nil
	}

	var endDate *time.Time
	if createDto.EndDate != nil {
		parsedEndDate, err := time.Parse(val.MonthYearLayout, *createDto.EndDate)
		if err == nil {
			endDate = &parsedEndDate
		}
	}

	userId, err := uuid.Parse(createDto.UserId)
	if err != nil {
		return nil
	}

	return &model.Subscription{
		Name:      createDto.Name,
		Price:     createDto.Price,
		UserId:    userId,
		StartDate: startDate,
		EndDate:   endDate,
	}
}
