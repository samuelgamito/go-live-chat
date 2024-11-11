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
			fx.As(new(usecase.ChatroomRepository)),
		),
	)
	UseCaseModule = fx.Options(useCasesFactory)
)
