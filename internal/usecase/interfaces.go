package usecase

import (
	"context"
	"github.com/redis/go-redis/v9"
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

type ConversationsRepository interface {
	SaveMessageToUser(message model.Message) error
	BatchSaveMessage(messages []model.Message) error
}

type RedisClient interface {
	Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
}
