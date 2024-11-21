package wrappers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionWrapper struct {
	collection *mongo.Collection
}

func (w *CollectionWrapper) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return w.collection.InsertOne(ctx, document, opts...)
}

func (w *CollectionWrapper) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return w.collection.UpdateOne(ctx, filter, update, opts...)
}

func (w *CollectionWrapper) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (MongoCursorInterface, error) {
	cursor, err := w.collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	return &MongoCursorWrapper{cursor: cursor}, nil
}

func (w *CollectionWrapper) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return w.collection.FindOne(ctx, filter, opts...)
}
