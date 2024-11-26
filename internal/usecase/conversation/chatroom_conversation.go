package usecase_conversation

import (
	"context"
	"encoding/json"
	"go-live-chat/internal/infraestructure/databases"
	"go-live-chat/internal/misc"
	"go-live-chat/internal/model"
	"go-live-chat/internal/usecase"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type ChatroomConversationUseCase struct {
	chatroomRepositorySearch usecase.ChatroomRepositorySearch
	rdb                      usecase.RedisClient
	conversationsRepository  usecase.ConversationsRepository
}

func NewChatroomConversationUseCase(
	chatroomRepositorySearch usecase.ChatroomRepositorySearch,
	conversationsRepository usecase.ConversationsRepository,
	rdb *databases.RedisClient) *ChatroomConversationUseCase {
	return &ChatroomConversationUseCase{
		chatroomRepositorySearch: chatroomRepositorySearch,
		conversationsRepository:  conversationsRepository,
		rdb:                      rdb.NotifyClientsRedis,
	}
}

func (c *ChatroomConversationUseCase) FindMembers(
	input interface{},
	ctx context.Context,
) ([]model.Member, *model.Error) {

	chatroomId := input.(string)

	chatroom, err := c.chatroomRepositorySearch.GetById(chatroomId, ctx)

	if err != nil {
		return nil, misc.DefaultError()
	}

	return chatroom.Members, nil
}

func (c *ChatroomConversationUseCase) StoreMessage(messages []model.Message, ctx context.Context) *model.Error {

	_ = c.conversationsRepository.BatchSaveMessage(messages)
	return nil
}

func (c *ChatroomConversationUseCase) PublishMessage(messages []model.Message, ctx context.Context) *model.Error {

	for _, message := range messages {
		jsonData, err := json.Marshal(message)
		if err != nil {
			log.Printf("Failed to publish to Redis channel %s: %v\n", message.To, err)
			return misc.DefaultError()
		}

		err = c.rdb.Publish(ctx, message.To, jsonData).Err()
		if err != nil {
			log.Printf("Failed to publish to Redis channel %s: %v\n", message.To, err)
			return misc.DefaultError()
		}
		log.Printf("Message published to channel %s: %s\n", message.To, message)

	}

	return nil
}

func (c *ChatroomConversationUseCase) PrepareMessage(
	members []model.Member,
	message string,
	from string) []model.Message {

	messages := make([]model.Message, 0)

	for _, member := range members {
		messages = append(messages, model.Message{

			Id:      primitive.NewObjectID(),
			To:      member.Id,
			From:    from,
			Content: message,
			Type:    "chatroom",
		})
	}

	return messages
}
