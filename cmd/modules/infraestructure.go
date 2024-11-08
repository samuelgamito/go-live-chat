package modules

import (
	"go-live-chat/internal/infraestructure/databases"
	"go.uber.org/fx"
)

var (
	infraFactory = fx.Provide(
		databases.NewMongoDBClient,
		databases.NewRedisClient,
	)

	InfraModule = fx.Options(
		infraFactory,
	)
)
