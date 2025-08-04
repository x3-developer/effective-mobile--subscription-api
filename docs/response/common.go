package docsResponse

type Response500 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message" example:"Internal server error"`
	ErrorCode string `json:"errorCode" enums:"SERVER_ERROR"`
}

type Response400 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message" example:"Bad request or validation error"`
	ErrorCode string `json:"errorCode" enums:"BAD_REQUEST"`
}

type Response404 struct {
	IsSuccess bool   `json:"isSuccess" example:"false"`
	Message   string `json:"message" example:"Resource not found"`
	ErrorCode string `json:"errorCode" enums:"NOT_FOUND"`
}
