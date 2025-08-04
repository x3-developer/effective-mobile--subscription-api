package vo

import (
	"github.com/google/uuid"
	"time"
)

type TotalCostFilter struct {
	Name      *string
	UserId    *uuid.UUID
	StartDate time.Time
	EndDate   time.Time
}
