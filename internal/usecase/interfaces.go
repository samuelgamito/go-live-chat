package usecase

import (
	"context"
	"go-live-chat/internal/model"
)

type ChatroomRepositoryUpdate interface {
	GetById(id string, ctx context.Context) (*model.Chatroom, error)
	Update(chatroom model.Chatroom, ctx context.Context) (*model.Chatroom, error)
}

type ChatroomRepositoryCreate interface {
	Create(chatroom model.Chatroom, ctx context.Context) (*model.Chatroom, error)
}

type ChatroomRepositorySearch interface {
	GetById(id string, ctx context.Context) (*model.Chatroom, error)
	GetByFilter(ctx context.Context) ([]model.Chatroom, error)
}
