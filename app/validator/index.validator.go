package validator

import (
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

//Formatting message validation Error
func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, fmt.Sprintf("Field '%s' failed validation with tag '%s'", e.Field(), e.Tag()))
	}

	return errors
} 


