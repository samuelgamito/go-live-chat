package dto

import "go-live-chat/internal/model"

type (
	ErrorResponse struct {
		StatusCode int
		Body       ErrorBodyDTO
	}

	ErrorBodyDTO struct {
		Messages []string `json:"messages"`
	}
)

func (e *ErrorResponse) Error() string {
	return "has validation error"
}

func (e *ErrorResponse) FromModel(err *model.Error) {
	e.StatusCode = err.StatusCode
	e.Body.Messages = err.Messages
}
