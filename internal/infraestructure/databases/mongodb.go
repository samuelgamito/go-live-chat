package databases

import (
	"context"
	"go-live-chat/internal/configs"
	"go-live-chat/internal/infraestructure/wrappers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBConnections struct {
	OpenChat wrappers.MongoClientInterface
}

func (m *MongoDBConnections) CloseAll() {
	err := m.OpenChat.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
}

func NewMongoDBClient(configs *configs.Config) *MongoDBConnections {
	client := getOpenChatMongoClient(configs.OpenChatMongoDB)
	return &MongoDBConnections{
		OpenChat: wrappers.NewMongoClientWrapper(client),
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
