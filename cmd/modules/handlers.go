package modules

import (
	"go-live-chat/internal/handlers"
	"go.uber.org/fx"
)

var (
	handlersFactory = fx.Provide(
		handlers.NewChatRoomHandler,
		handlers.NewMessageHandler,
		handlers.NewUserManagementHandler,
	)
	HandlersModule = fx.Options(
		handlersFactory,
		handlers.ModuleUserManagementHandler,
		handlers.ModuleMessageHandler,
		handlers.ModuleChatRoomHandler,
		handlers.ModuleHandler,
	)
)
