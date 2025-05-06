package validationutil

import (
	"fmt"
	"learnyscape-backend-mono/pkg/constant"
	validationtype "learnyscape-backend-mono/pkg/util/validation/type"

	"github.com/go-playground/validator/v10"
)

func TagToMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "len":
		return fmt.Sprintf("%s length or value must be exactly %v", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s length or value %v must be at most", fe.Field(), fe.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %v", fe.Field(), fe.Param())
	case "lte":
		return fmt.Sprintf("%s must be lower than or equal to %v", fe.Field(), fe.Param())
	case "gtefield":
		return fmt.Sprintf("%s must be greater than or equal to %v", fe.StructField(), fe.Param())
	case "email":
		return fmt.Sprintf("%s has invalid email format", fe.Field())
	case "eq":
		return fmt.Sprintf("%s must be equal to %v", fe.Field(), fe.Param())
	case "min":
		return fmt.Sprintf("%s length or value must be at least %v", fe.Field(), fe.Param())
	case "numeric":
		return fmt.Sprintf("%s must be a number", fe.Field())
	case "boolean":
		return fmt.Sprintf("%s must be a boolean", fe.Field())
	case "time_format":
		return fmt.Sprintf("please send time in format of %s", constant.ConvertGoTimeLayoutToReadable(fe.Param()))
	case "password":
		password := validationtype.NewPassword(fe.Value().(string))
		password.Validate()
		return password.Message()
	default:
		return "invalid input"
	}
}
