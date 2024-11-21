package wrappers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClientWrapper struct {
	client *mongo.Client
}

func (w *MongoClientWrapper) Database(name string, opts ...*options.DatabaseOptions) MongoDatabaseInterface {
	return &MongoDatabaseWrapper{db: w.client.Database(name, opts...)}
}

func (w *MongoClientWrapper) Disconnect(ctx context.Context) error {
	return w.client.Disconnect(ctx)
}

func (w *MongoClientWrapper) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	return w.client.Ping(ctx, rp)
}

func NewMongoClientWrapper(client *mongo.Client) *MongoClientWrapper {
	return &MongoClientWrapper{client: client}
}
