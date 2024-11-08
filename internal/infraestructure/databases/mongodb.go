package databases

import (
	"context"
	"go-live-chat/internal/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	OpenChat *mongo.Client
}

func (m *MongoDBClient) CloseAll() {
	err := m.OpenChat.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
}

func NewMongoDBClient(configs *configs.Config) *MongoDBClient {
	return &MongoDBClient{
		OpenChat: getOpenChatMongoClient(configs.OpenChatMongoDB),
	}
}

func getOpenChatMongoClient(conf *configs.MongoDBConfig) *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(conf.GetConnectionString()).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	bsonCommand := bson.D{{Key: "ping", Value: 1}}
	if err := client.Database("admin").RunCommand(context.TODO(), bsonCommand).Err(); err != nil {
		panic(err)
	}

	return client
}
