package usecase_chatroom

import (
	"context"
	"go-live-chat/internal/model"
	"go-live-chat/internal/usecase"
	"log"
	"time"
)

type ChatroomRepositoryCreate struct {
	repo usecase.ChatroomRepositoryCreate
}

func NewCreateChatroomUseCase(repo usecase.ChatroomRepositoryCreate) *ChatroomRepositoryCreate {
	return &ChatroomRepositoryCreate{repo: repo}
}

func (c *ChatroomRepositoryCreate) Execute(chatroom model.Chatroom, ctx context.Context) (string, *model.Error) {

	chatroom.Members = append(chatroom.Members, model.Member{
		Id:      chatroom.Owner,
		SinceAt: time.Now(),
	})

	chatroomResp, err := c.repo.Create(chatroom, ctx)

	if err != nil {
		log.Printf("Error creating chatroom: %v", err)
		var e model.Error
		e.Messages = []string{"Error creating chatroom"}

		return "", &e
	}

	return chatroomResp.Id.Hex(), nil
}
