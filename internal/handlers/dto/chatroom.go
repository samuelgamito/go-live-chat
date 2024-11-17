package dto

import (
	"go-live-chat/internal/misc"
	"go-live-chat/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
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

	ChatRoomDTO struct {
		Id          string      `json:"id,omitempty"`
		Name        string      `json:"name,omitempty"`
		Owner       string      `json:"owner,omitempty"`
		Members     []MemberDTO `json:"members,omitempty"`
		Description string      `json:"description,omitempty"`
		CreatedAt   *time.Time  `json:"created,omitempty"`
		UpdatedAt   *time.Time  `json:"updated,omitempty"`
	}

	MemberDTO struct {
		Id      string    `json:"id"`
		SinceAt time.Time `json:"since_at"`
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

func GetChatroomResponse(chatroom *model.Chatroom) ChatRoomDTO {
	var membersDto []MemberDTO

	if chatroom.Members != nil {
		for _, member := range chatroom.Members {
			membersDto = append(membersDto, MemberDTO{
				Id:      member.Id,
				SinceAt: member.SinceAt,
			})
		}
	} else {
		membersDto = nil
	}

	return ChatRoomDTO{
		Id:          chatroom.Id.Hex(),
		Name:        chatroom.Name,
		Owner:       chatroom.Owner,
		Description: chatroom.Description,
		Members:     membersDto,
		CreatedAt:   chatroom.CreatedAt,
		UpdatedAt:   chatroom.UpdatedAt,
	}
}
