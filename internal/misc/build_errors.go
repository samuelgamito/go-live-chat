package misc

import (
	"fmt"
	"github.com/go-playground/validator"
	"go-live-chat/internal/model"
)

const (
	defaultErrorMessage           = "Not able to process the request, contact the admin"
	maxErrorMessage               = "%s max values is %s, actual value is %d."
	minErrorMessage               = "%s min values is %s, actual value is %d."
	missingRequiredFieldMessage   = "Field %s is required."
	defaultValidationErrorMessage = "Validation fails on %s"
)

func DefaultError() *model.Error {
	return &model.Error{
		StatusCode: 500,
		Messages:   []string{defaultErrorMessage},
	}
}

func BuildErrorMessage(e validator.FieldError) string {
	switch errType := e.Tag(); errType {
	case "min":
		return fmt.Sprintf(minErrorMessage, e.Field(), e.Param(), e.Value())
	case "max":
		return fmt.Sprintf(maxErrorMessage, e.Field(), e.Param(), e.Value())
	case "required":
		return fmt.Sprintf(missingRequiredFieldMessage, e.Field())
	default:
		return fmt.Sprintf(defaultValidationErrorMessage, e.Field())
	}
}
