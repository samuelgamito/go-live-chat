package modules

import (
	"go-live-chat/internal/repositories"
	"go-live-chat/internal/usecase"
	"go.uber.org/fx"
)

var (
	useCasesFactory = fx.Provide(
		fx.Annotate(
			repositories.NewChatroomRepository,
			fx.As(new(usecase.ChatroomRepositoryUpdate)),
			fx.As(new(usecase.ChatroomRepositoryCreate)),
			fx.As(new(usecase.ChatroomRepositorySearch)),
		),
	)
	UseCaseModule = fx.Options(useCasesFactory)
)
