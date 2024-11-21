package wrappers

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabaseWrapper struct {
	db *mongo.Database
}

func (w *MongoDatabaseWrapper) Collection(name string, opts ...*options.CollectionOptions) MongoCollectionInterface {
	return &CollectionWrapper{collection: w.db.Collection(name, opts...)}
}
