package repositories

import (
	"go-live-chat/internal/configs"
	"go-live-chat/internal/infraestructure/databases"
	"go-live-chat/internal/model"
)

type ConversationsRepository struct {
	client MongoClientInterface
	config *configs.Config
}

func NewConversationsRepository(client *databases.MongoDBConnections, config *configs.Config) *ConversationsRepository {
	return &ConversationsRepository{client: client.OpenChat, config: config}
}

func (c *ConversationsRepository) RetrieveChatroomHistory(chatroomId string, page int, pageSize int) ([]model.Message, error) {

	return nil, nil
}

func (c *ConversationsRepository) SaveMessageToUser(message model.Message) error {
	return nil
}

func (c *ConversationsRepository) BatchSaveMessage(messages []model.Message) error {
	return nil
}
func (c *ConversationsRepository) RetrieveLatestMessageFromUser(userId string) (*model.Message, error) {
	return nil, nil
}
