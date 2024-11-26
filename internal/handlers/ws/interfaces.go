package ws

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-live-chat/internal/model"
)

type ConversationUseCase interface {
	FindMembers(input interface{}, ctx context.Context) ([]model.Member, *model.Error)
	StoreMessage(messages []model.Message, ctx context.Context) *model.Error
	PublishMessage(messages []model.Message, ctx context.Context) *model.Error
	PrepareMessage(
		members []model.Member,
		message string,
		from string) []model.Message
}

type RetrieveChatroomUseCase interface {
	ExecuteById(id string, ctx context.Context) (*model.Chatroom, *model.Error)
}

type RedisClient interface {
	Publish(ctx context.Context, channel string, message interface{}) *redis.IntCmd
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
}
