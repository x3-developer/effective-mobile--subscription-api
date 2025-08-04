package mapper

import (
	"subscriptions/internal/modules/subscription/v1/application/dto"
	"subscriptions/internal/modules/subscription/v1/domain/vo"
	"subscriptions/internal/shared/lib/val"
	"time"
)

func ToTotalCostFilterVOFromDTO(dto *dto.TotalCostDTO) *vo.TotalCostFilter {
	if dto == nil {
		return nil
	}

	startDate, err := time.Parse(val.MonthYearLayout, dto.StartDate)
	if err != nil {
		return nil
	}
	endDate, err := time.Parse(val.MonthYearLayout, dto.EndDate)
	if err != nil {
		return nil
	}

	return &vo.TotalCostFilter{
		Name:      dto.Name,
		UserId:    dto.UserId,
		StartDate: startDate,
		EndDate:   endDate,
	}
}
