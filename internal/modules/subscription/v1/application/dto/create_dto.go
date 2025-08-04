package dto

type CreateDTO struct {
	Name      string  `json:"name" validate:"required,min=1,max=255"`
	Price     int64   `json:"price" validate:"required,min=0"`
	UserId    string  `json:"userId" validate:"required,uuid4"`
	StartDate string  `json:"startDate" validate:"required,date"`
	EndDate   *string `json:"endDate" validate:"omitempty,date"`
}
