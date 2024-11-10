package repositories

import (
	"context"
	"go-live-chat/internal/configs"
	"go-live-chat/internal/infraestructure/databases"
	"go-live-chat/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ChatroomRepository struct {
	client *mongo.Client
	config *configs.Config
}

func NewChatroomRepository(client *databases.MongoDBClient, config *configs.Config) *ChatroomRepository {
	return &ChatroomRepository{client.OpenChat, config}
}

func (c *ChatroomRepository) Create(chatroom model.Chatroom, ctx context.Context) (*model.Chatroom, error) {

	timeNow := time.Now()

	chatroom.CreatedAt = timeNow
	chatroom.UpdatedAt = timeNow

	chat, err := c.client.Database(c.config.OpenChatMongoDB.Database).Collection("chatrooms").InsertOne(ctx, chatroom)
	if err != nil {
		return nil, err

	}
	chatroom.Id = chat.InsertedID.(primitive.ObjectID)
	return &chatroom, nil
}

func (c *ChatroomRepository) GetById(id string, ctx context.Context) (*model.Chatroom, error) {

	bsonId, err := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": bsonId}
	if err != nil {
		return nil, err
	}

	var chatroom model.Chatroom

	err = c.client.Database(c.config.OpenChatMongoDB.Database).Collection("chatrooms").FindOne(ctx, filter).Decode(&chatroom)

	if err != nil {
		return nil, err
	}

	return &chatroom, nil
}

func (c *ChatroomRepository) Update(chatroom model.Chatroom, ctx context.Context) (*model.Chatroom, error) {
	timeNow := time.Now()
	chatroom.UpdatedAt = timeNow
	filter := bson.M{"_id": chatroom.Id}
	update := bson.M{"$set": chatroom}

	_, err := c.client.Database(c.config.OpenChatMongoDB.Database).Collection("chatrooms").UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}

	return &chatroom, nil
}
