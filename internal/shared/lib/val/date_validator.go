package val

import (
	"github.com/go-playground/validator/v10"
	"time"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	_ = validate.RegisterValidation("date", validateDate)
}

func validateDate(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	layout := "02.01.2006"
	_, err := time.Parse(layout, value)
	return err == nil
}
