package modules

import (
	"go-live-chat/internal/repositories"
	"go.uber.org/fx"
)

var (
	repositoriesFactory = fx.Provide(
		repositories.NewChatroomRepository,
		repositories.NewConversationsRepository,
	)

	RepositoriesModule = fx.Options(repositoriesFactory)
)
