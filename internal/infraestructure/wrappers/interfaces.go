package wrappers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClientInterface interface {
	Database(name string, opts ...*options.DatabaseOptions) MongoDatabaseInterface
	Disconnect(ctx context.Context) error
	Ping(ctx context.Context, rp *readpref.ReadPref) error
}

type MongoDatabaseInterface interface {
	Collection(name string, opts ...*options.CollectionOptions) MongoCollectionInterface
}

type MongoCollectionInterface interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (MongoCursorInterface, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
}

type MongoCursorInterface interface {
	All(ctx context.Context, results interface{}) error
}
