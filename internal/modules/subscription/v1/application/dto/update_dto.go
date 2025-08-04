package dto

type UpdateDTO struct {
	Name      *string `json:"name" validate:"omitempty,min=1,max=255"`
	Price     *int64  `json:"price" validate:"omitempty,min=0"`
	UserId    *string `json:"userId" validate:"omitempty,uuid4"`
	StartDate *string `json:"startDate" validate:"omitempty,monthYear"`
	EndDate   *string `json:"endDate,omitempty" validate:"omitempty,monthYear"`
}
