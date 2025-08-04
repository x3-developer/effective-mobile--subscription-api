package res

type ErrorCode string

const (
	BadRequest       ErrorCode = "BAD_REQUEST"
	ServerError      ErrorCode = "SERVER_ERROR"
	MinLength        ErrorCode = "MIN_LENGTH"
	MaxLength        ErrorCode = "MAX_LENGTH"
	NotBlank         ErrorCode = "NOT_BLANK"
	NotFound         ErrorCode = "NOT_FOUND"
	MethodNotAllowed ErrorCode = "METHOD_NOT_ALLOWED"
)

func GetErrorCodeByTag(tag string) ErrorCode {
	switch tag {
	case "required":
		return NotBlank
	case "min":
		return MinLength
	case "max":
		return MaxLength
	case "date":
		return BadRequest
	case "uuid4":
		return BadRequest
	default:
		return ServerError
	}
}
