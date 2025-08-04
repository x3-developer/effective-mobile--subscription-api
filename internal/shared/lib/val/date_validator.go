package val

import (
	"github.com/go-playground/validator/v10"
	"time"
)

const MonthYearLayout = "01-2006"

var validate *validator.Validate

func init() {
	validate = validator.New()
	_ = validate.RegisterValidation("monthYear", validateMonthYear)
}

func validateMonthYear(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	_, err := time.Parse(MonthYearLayout, value)
	return err == nil
}
