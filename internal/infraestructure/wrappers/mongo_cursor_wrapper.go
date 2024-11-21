package wrappers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoCursorWrapper struct {
	cursor *mongo.Cursor
}

func (w *MongoCursorWrapper) All(ctx context.Context, results interface{}) error {
	return w.cursor.All(ctx, results)
}
