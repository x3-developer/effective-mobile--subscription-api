package val

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"subscriptions/internal/shared/lib/res"
)

func ValidateDTO(dto interface{}) []res.ErrorField {
	var errorFields []res.ErrorField

	err := validate.Struct(dto)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, ve := range validationErrors {
				fieldName := ve.Field()
				tag := ve.Tag()
				errorCode := res.GetErrorCodeByTag(tag)

				errorFields = append(errorFields, res.NewErrorField(fieldName, string(errorCode)))
			}
		} else {
			errorFields = append(errorFields, res.NewErrorField("validation", string(res.BadRequest)))
		}
	}

	return errorFields
}
