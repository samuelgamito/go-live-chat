package repositories

import (
	"context"
	"fmt"
	"go-live-chat/internal/configs"
	"go-live-chat/internal/infraestructure/databases"
	"go-live-chat/internal/infraestructure/wrappers"
	"go-live-chat/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoClientInterface interface {
	Database(name string, opts ...*options.DatabaseOptions) wrappers.MongoDatabaseInterface
}

type ChatroomRepository struct {
	client MongoClientInterface
	config *configs.Config
}

func NewChatroomRepository(client *databases.MongoDBConnections, config *configs.Config) *ChatroomRepository {
	return &ChatroomRepository{client.OpenChat, config}
}

func (c *ChatroomRepository) Create(chatroom model.Chatroom, ctx context.Context) (*model.Chatroom, error) {

	timeNow := time.Now()

	chatroom.CreatedAt = &timeNow
	chatroom.UpdatedAt = &timeNow

	chat, err := c.client.Database(c.config.OpenChatMongoDB.Database).Collection("chatrooms").InsertOne(ctx, chatroom)
	if err != nil {
		return nil, err

	}
	chatroom.Id = chat.InsertedID.(primitive.ObjectID)
	return &chatroom, nil
}

func (c *ChatroomRepository) GetByFilter(ctx context.Context) ([]model.Chatroom, error) {
	filter := bson.M{}
	projection := bson.D{
		primitive.E{Key: "id", Value: 1},
		primitive.E{Key: "name", Value: 1},
		primitive.E{Key: "owner", Value: 1},
		primitive.E{Key: "description", Value: 1},
		primitive.E{Key: "createdAt", Value: 1},
		primitive.E{Key: "updatedAt", Value: 1},
	}

	opts := options.Find().SetProjection(projection)
	cursor, err := c.client.
		Database(c.config.OpenChatMongoDB.Database).
		Collection("chatrooms").
		Find(ctx, filter, opts)

	if err != nil {
		return nil, err
	}

	var chatrooms []model.Chatroom
	if err = cursor.All(ctx, &chatrooms); err != nil {
		return nil, err
	}

	return chatrooms, nil
}

func (c *ChatroomRepository) GetById(id string, ctx context.Context) (*model.Chatroom, error) {

	bsonId, err := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": bsonId}
	if err != nil {
		return nil, err
	}

	var chatroom model.Chatroom

	singleResult := c.client.
		Database(c.config.OpenChatMongoDB.Database).
		Collection("chatrooms").
		FindOne(ctx, filter)

	if singleResult.Err() != nil {
		fmt.Println(err)
		return nil, err
	}

	err = singleResult.Decode(&chatroom)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &chatroom, nil
}

func (c *ChatroomRepository) Update(chatroom model.Chatroom, ctx context.Context) (*model.Chatroom, error) {
	timeNow := time.Now()
	chatroom.UpdatedAt = &timeNow
	filter := bson.M{"_id": chatroom.Id}
	update := bson.M{"$set": chatroom}

	_, err := c.client.
		Database(c.config.OpenChatMongoDB.Database).
		Collection("chatrooms").
		UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}

	return &chatroom, nil
}
