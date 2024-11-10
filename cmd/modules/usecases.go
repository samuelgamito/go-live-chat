package modules

import (
	usecase "go-live-chat/internal/usecase/chatroom"
	"go.uber.org/fx"
)

var (
	useCasesFactory = fx.Provide(usecase.NewCreateChatroomUseCase)
	UseCaseModule   = fx.Options(useCasesFactory)
)
