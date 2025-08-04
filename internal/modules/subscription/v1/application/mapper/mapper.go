package mapper

import (
	"github.com/google/uuid"
	"subscriptions/internal/modules/subscription/v1/application/dto"
	"subscriptions/internal/modules/subscription/v1/domain/model"
	"subscriptions/internal/modules/subscription/v1/domain/vo"
	"time"
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

func ToModelFromCreateDTO(createDto *dto.CreateDTO) *model.Subscription {
	if createDto == nil {
		return nil
	}

	layout := "02.01.2006"
	startDate, err := time.Parse(layout, createDto.StartDate)
	if err != nil {
		return nil
	}

	var endDate *time.Time
	if createDto.EndDate != nil {
		parsedEndDate, err := time.Parse(layout, *createDto.EndDate)
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

func ToModelFromUpdateDTO(updateDto *dto.UpdateDTO, model *model.Subscription) *model.Subscription {
	if updateDto == nil || model == nil {
		return model
	}

	layout := "02.01.2006"

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
		startDate, err := time.Parse(layout, *updateDto.StartDate)
		if err == nil {
			model.StartDate = startDate
		}
	}
	if updateDto.EndDate != nil {
		endDate, err := time.Parse(layout, *updateDto.EndDate)
		if err == nil {
			model.EndDate = &endDate
		}
	}

	return model
}

func ToTotalCostFilterVOFromDTO(dto *dto.TotalCostDTO) *vo.TotalCostFilter {
	if dto == nil {
		return nil
	}

	layout := "02.01.2006"
	startDate, err := time.Parse(layout, dto.StartDate)
	if err != nil {
		return nil
	}
	endDate, err := time.Parse(layout, dto.EndDate)
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
