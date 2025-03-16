package helpers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"time"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	_ = validate.RegisterValidation("date", validateDate)
}

func validateDate(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	if _, err := time.Parse(time.RFC3339, dateStr); err == nil {
		return true
	}

	if _, err := time.Parse("2006-01-02", dateStr); err == nil {
		return true
	}

	return false
}

func ValidateStruct(i interface{}) map[string]string {
	err := validate.Struct(i)

	if err != nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	errors.As(err, &validationErrors)

	errorsMaps := make(map[string]string)

	for _, ve := range validationErrors {
		field := ve.Field()
		tag := ve.Tag()

		switch tag {
		case "required":
			errorsMaps[field] = "This field is required"

		case "email":
			errorsMaps[field] = "This field must be a valid email address"

		case "oneof":
			errorsMaps[field] = "This field must be a valid oneof type"

		case "min":
			errorsMaps[field] = "This field must be a valid minimum"
		case "max":
			errorsMaps[field] = "This field must be a valid maximum"

		case "alphanum":
			errorsMaps[field] = "This field must be a valid alphanumeric string"

		case "alpha":
			errorsMaps[field] = "This field must be a valid alpha string"

		case "containsany":
			errorsMaps[field] = "This field must be a valid containsany string"

		case "date":
			errorsMaps[field] = "This field must be a valid date"

		default:
			errorsMaps[field] = "Validation failed for " + tag
		}
	}
	return errorsMaps
}
