package usecase_conversation

import (
	"context"
	"encoding/json"
	"go-live-chat/internal/infraestructure/databases"
	"go-live-chat/internal/model"
	"go-live-chat/internal/usecase"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type UserConversationUseCase struct {
	rdb usecase.RedisClient
}

func NewUserConversationUseCase(rdb *databases.RedisClient) *UserConversationUseCase {
	return &UserConversationUseCase{
		rdb: rdb.NotifyClientsRedis,
	}
}

func (c *UserConversationUseCase) FindMembers(input interface{}, ctx context.Context) ([]model.Member, *model.Error) {
	m := make([]model.Member, 1)
	m[0] = model.Member{
		Id:      input.(string),
		SinceAt: time.Now(),
	}
	return m, nil
}

func (c *UserConversationUseCase) StoreMessage(messages []model.Message, ctx context.Context) *model.Error {
	return nil
}
func (c *UserConversationUseCase) PublishMessage(messages []model.Message, ctx context.Context) *model.Error {
	message := messages[0]

	if jsonData, err := json.Marshal(message); err == nil {
		err = c.rdb.Publish(ctx, message.To, jsonData).Err()
		if err != nil {
			log.Printf("Failed to publish to Redis channel %s: %v\n", message.To, err)
		}

		log.Printf("Message published to channel %s: %s\n", message.To, message)
	} else {
		log.Printf("Failed to publish to Redis channel %s: %v\n", message.To, err)
	}
	return nil
}

func (c *UserConversationUseCase) PrepareMessage(
	members []model.Member,
	message string,
	from string) []model.Message {

	m := make([]model.Message, 1)
	m[0] = model.Message{
		Id:      primitive.NewObjectID(),
		To:      members[0].Id,
		From:    from,
		Content: message,
		Type:    "user",
	}
	return m
}
