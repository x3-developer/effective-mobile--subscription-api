package dto

import (
	"github.com/google/uuid"
	"time"
)

type ResponseDTO struct {
	Id        uint       `json:"id"`
	Name      string     `json:"name"`
	Price     int64      `json:"price"`
	UserId    uuid.UUID  `json:"userId"`
	StartDate time.Time  `json:"startDate"`
	EndDate   *time.Time `json:"endDate,omitempty"`
}
