package modules

import (
	"go-live-chat/internal/handlers"
	usecase "go-live-chat/internal/usecase/chatroom"
	"go.uber.org/fx"
)

var (
	handlersFactory = fx.Provide(
		handlers.NewChatRoomHandler,
		handlers.NewMessageHandler,
		handlers.NewUserManagementHandler,
		handlers.NewActuatorHandler,
		fx.Annotate(
			usecase.NewCreateChatroomUseCase,
			fx.As(new(handlers.CreateChatroomUseCase)),
		),
	)
	HandlersModule = fx.Options(
		handlersFactory,
		handlers.ModuleUserManagementHandler,
		handlers.ModuleMessageHandler,
		handlers.ModuleChatRoomHandler,
		handlers.ModuleHandler,
		handlers.ModuleActuatorHandler,
	)
)
