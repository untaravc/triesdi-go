package validator

import (
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
