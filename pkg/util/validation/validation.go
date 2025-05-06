package validationutil

import (
	validationtype "learnyscape-backend-mono/pkg/util/validation/type"

	"github.com/go-playground/validator/v10"
)

func Password(fl validator.FieldLevel) bool {
	data, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}

	password := validationtype.NewPassword(data)
	return password.Validate()
}
