package handlers

import (
	"context"
	"go-live-chat/internal/model"
)

type CreateChatroomUseCase interface {
	Execute(chatroom model.Chatroom, ctx context.Context) (string, *model.Error)
}

type RetrieveChatroomUseCase interface {
	ExecuteById(id string, ctx context.Context) (*model.Chatroom, *model.Error)
	ExecuteByFilter(ctx context.Context) ([]model.Chatroom, *model.Error)
}
