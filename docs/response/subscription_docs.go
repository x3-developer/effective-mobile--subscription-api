package docsResponse

import (
	"subscriptions/internal/modules/subscription/v1/application/dto"
)

type ErrorField struct {
	Field     string `json:"field" enums:"name,price,userId,startDate,endDate"`
	ErrorCode string `json:"errorCode" enums:"NOT_BLANK,NOT_FOUND,MIN_LENGTH,MAX_LENGTH"`
}

type SubscriptionCreate201 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type SubscriptionCreate400 struct {
	Response400
	Fields []ErrorField `json:"fields,omitempty"`
}

type SubscriptionList200 struct {
	IsSuccess bool              `json:"isSuccess" example:"true"`
	Data      []dto.ResponseDTO `json:"data"`
}

type SubscriptionGetById200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type SubscriptionUpdate200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type SubscriptionUpdate400 struct {
	Response400
	Fields []ErrorField `json:"fields,omitempty"`
}

type SubscriptionDelete200 struct {
	IsSuccess bool            `json:"isSuccess" example:"true"`
	Data      dto.ResponseDTO `json:"data"`
}

type SubscriptionGetTotalCost200 struct {
	IsSuccess bool  `json:"isSuccess" example:"true"`
	Data      int64 `json:"data"`
}
