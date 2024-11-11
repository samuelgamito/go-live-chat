package usecase

import (
	"context"
	"go-live-chat/internal/model"
)

type ChatroomRepository interface {
	Create(chatroom model.Chatroom, ctx context.Context) (*model.Chatroom, error)
	GetById(id string, ctx context.Context) (*model.Chatroom, error)
	Update(chatroom model.Chatroom, ctx context.Context) (*model.Chatroom, error)
}
