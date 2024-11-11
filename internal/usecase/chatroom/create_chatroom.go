package usecase_chatroom

import (
	"context"
	"fmt"
	"go-live-chat/internal/model"
	"go-live-chat/internal/usecase"
	"time"
)

type CreateChatRoomUseCase struct {
	repo usecase.ChatroomRepository
}

func NewCreateChatroomUseCase(repo usecase.ChatroomRepository) *CreateChatRoomUseCase {
	return &CreateChatRoomUseCase{repo: repo}
}

func (c *CreateChatRoomUseCase) Execute(chatroom model.Chatroom, ctx context.Context) (string, *model.Error) {

	chatroom.Members = append(chatroom.Members, model.Member{
		Id:      chatroom.Owner,
		SinceAt: time.Now(),
	})

	chatroomResp, err := c.repo.Create(chatroom, ctx)

	if err != nil {
		fmt.Print(err)

		var e model.Error
		e.Messages = []string{"Error creating chatroom"}

		return "", &e
	}

	return chatroomResp.Id.Hex(), nil
}
