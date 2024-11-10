package modules

import (
	"go-live-chat/internal/repositories"
	"go.uber.org/fx"
)

var (
	repositoriesFactory = fx.Provide(
		repositories.NewChatroomRepository,
	)

	RepositoriesModule = fx.Options(repositoriesFactory)
)
