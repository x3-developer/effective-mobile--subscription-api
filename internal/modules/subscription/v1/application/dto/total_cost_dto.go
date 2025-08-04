package dto

import (
	"github.com/google/uuid"
)

type TotalCostDTO struct {
	Name      *string    `json:"name" validate:"omitempty,min=1,max=255"`
	UserId    *uuid.UUID `json:"userId" validate:"omitempty,uuid4"`
	StartDate string     `json:"startDate" validate:"required,date"`
	EndDate   string     `json:"endDate" validate:"required,date"`
}
