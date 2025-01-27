package validator

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Initialize the validator
func InitValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
	}
	return validate
}

// ValidateStruct validates any struct based on its tags
func ValidateStruct(s interface{}) error {
	v := InitValidator()
	return v.Struct(s)
}

// Formatting message validation Error
func FormatValidationError(err error) []string {
	var errors []string

	// Handle validation errors
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrs {
			errors = append(errors, fmt.Sprintf("Field '%s' failed validation with tag '%s'", e.Field(), e.Tag()))
		}
	}

	// Handle JSON unmarshaling errors
	if unmarshalErr, ok := err.(*json.UnmarshalTypeError); ok {
		errors = append(errors, fmt.Sprintf(
			"Field '%s' has an invalid value. Expected '%s' but got '%s'",
			unmarshalErr.Field, unmarshalErr.Type.String(), unmarshalErr.Value,
		))
	}

	// Handle generic error if it's neither validation nor unmarshaling
	if len(errors) == 0 {
		errors = append(errors, err.Error())
	}

	return errors
}
