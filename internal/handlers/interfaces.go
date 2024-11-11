package handlers

import (
	"context"
	"go-live-chat/internal/model"
)

type CreateChatroomUseCase interface {
	Execute(chatroom model.Chatroom, ctx context.Context) (string, *model.Error)
}
