package config

import (
	"fmt"

	"simple-ecommerce/src/commons/structs"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator *validator.Validate
}

func messageMap(validatorTag string, fieldName string, validatorParam string) string {
	switch validatorTag {
	case "required":
		return fmt.Sprintf("%s field is required", fieldName)
	case "numeric":
		return fmt.Sprintf("%s field must be numeric only", fieldName)
	case "alpha":
		return fmt.Sprintf("%s field must be alphabet only", fieldName)
	case "email":
		return fmt.Sprintf("%s field must be valid email format", fieldName)
	case "alphanum":
		return fmt.Sprintf("%s field must be alphabet and numeric only", fieldName)
	case "uuid4":
		return fmt.Sprintf("%s field be valid uuid v4 format", fieldName)
	case "min":
		return fmt.Sprintf("%s field minimum is %s", fieldName, validatorParam)
	case "max":
		return fmt.Sprintf("%s field maximum is %s", fieldName, validatorParam)
	case "printascii":
		return fmt.Sprintf("%s field must in ascii character", fieldName)
	default:
		return ""
	}
}

func (v *Validator) Validate(data interface{}) *structs.Response {
	if fieldErrors := v.Validator.Struct(data); fieldErrors != nil {
		response := &structs.Response{
			Status:  false,
			Message: "Failed to POST data",
			Data:    nil,
		}
		errors := make([]string, len(fieldErrors.(validator.ValidationErrors)))
		for index, err := range fieldErrors.(validator.ValidationErrors) {
			errors[index] = messageMap(err.Tag(), err.Field(), err.Param())
		}
		response.Errors = &errors
		return response
	}
	return nil
}
