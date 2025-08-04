package model

import (
	"github.com/google/uuid"
	"time"
)

type Subscription struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Price     int64      `json:"price"`
	UserId    uuid.UUID  `json:"userId"`
	StartDate time.Time  `json:"startDate"`
	EndDate   *time.Time `json:"endDate"`
}
