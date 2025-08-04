package mapper

import (
	"subscriptions/internal/modules/subscription/v1/application/dto"
	"subscriptions/internal/modules/subscription/v1/domain/model"
)

func ToResponseDTOFromModel(model *model.Subscription) *dto.ResponseDTO {
	if model == nil {
		return nil
	}

	return &dto.ResponseDTO{
		Id:        model.ID,
		Name:      model.Name,
		Price:     model.Price,
		UserId:    model.UserId,
		StartDate: model.StartDate,
		EndDate:   model.EndDate,
	}
}
