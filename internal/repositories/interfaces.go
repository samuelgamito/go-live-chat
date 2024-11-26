package repositories

import (
	"go-live-chat/internal/infraestructure/wrappers"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClientInterface interface {
	Database(name string, opts ...*options.DatabaseOptions) wrappers.MongoDatabaseInterface
}
