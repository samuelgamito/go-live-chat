package misc

import (
	"github.com/go-playground/validator"
	"go-live-chat/internal/model"
)

func Validate(body interface{}) *model.Error {

	v := validator.New()
	err := v.Struct(body)

	if err == nil {
		return nil
	}

	errBody := model.Error{
		StatusCode: 400,
		Messages:   make([]string, len(err.(validator.ValidationErrors))),
	}

	for i, e := range err.(validator.ValidationErrors) {
		errBody.Messages[i] = BuildErrorMessage(e)
	}

	return &errBody
}
