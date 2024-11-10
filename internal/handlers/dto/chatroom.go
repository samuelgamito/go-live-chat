package dto

import (
	"go-live-chat/internal/misc"
	"go-live-chat/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CreateChatRoomRequestDTO struct {
		Name        string `json:"name" validate:"required"`
		Owner       string `json:"owner" validate:"required"`
		Description string `json:"description"`
	}

	CreatedChatRoomDTO struct {
		Id string `json:"id"`
	}
)

func (c *CreateChatRoomRequestDTO) IsValid() *ErrorResponse {
	if (CreateChatRoomRequestDTO{}) == *c {
		return &ErrorResponse{
			StatusCode: 400,
			Body: ErrorBodyDTO{
				Messages: []string{"Add at least one field to update"},
			},
		}
	}

	errMisc := misc.Validate(c)

	if errMisc != nil {
		return &ErrorResponse{
			StatusCode: errMisc.StatusCode,
			Body: ErrorBodyDTO{
				Messages: errMisc.Messages,
			},
		}
	}

	return nil

}

func (c *CreateChatRoomRequestDTO) ToModel() model.Chatroom {

	return model.Chatroom{
		Id:          primitive.NewObjectID(),
		Name:        c.Name,
		Owner:       c.Owner,
		Description: c.Description,
		Members:     []model.Member{},
	}
}
