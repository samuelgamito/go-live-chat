package modules

import (
	"go-live-chat/internal/handlers"
	"go-live-chat/internal/handlers/ws"
	"go-live-chat/internal/infraestructure/databases"
	"go-live-chat/internal/usecase"
	usecase_chatroom "go-live-chat/internal/usecase/chatroom"
	usecase_conversation "go-live-chat/internal/usecase/conversation"
	"go.uber.org/fx"
)

func ConversationModule() fx.Option {
	return fx.Provide(
		func(chatroomRepositorySearch usecase.ChatroomRepositorySearch,
			conversationsRepository usecase.ConversationsRepository,
			rdb *databases.RedisClient) map[string]ws.ConversationUseCase {
			return map[string]ws.ConversationUseCase{
				"chatroom": usecase_conversation.NewChatroomConversationUseCase(chatroomRepositorySearch, conversationsRepository, rdb),
				"user":     usecase_conversation.NewUserConversationUseCase(rdb),
			}
		},
	)
}

var (
	handlersFactory = fx.Provide(
		handlers.NewChatRoomHandler,
		handlers.NewMessageHandler,
		handlers.NewUserManagementHandler,
		handlers.NewActuatorHandler,
		handlers.NewChatWebSocketHandler,
		fx.Annotate(
			usecase_chatroom.NewCreateChatroomUseCase,
			fx.As(new(handlers.CreateChatroomUseCase)),
		),
		fx.Annotate(
			usecase_chatroom.NewRetrieveChatroom,
			fx.As(new(handlers.RetrieveChatroomUseCase)),
		),
		fx.Annotate(
			usecase_chatroom.NewUserManagementChatroomUseCase,
			fx.As(new(handlers.UserManagementChatroomUseCase)),
		),
	)
	HandlersModule = fx.Options(
		handlersFactory,
		ConversationModule(),
		handlers.ModuleChatWebSocketHandler,
		handlers.ModuleUserManagementHandler,
		handlers.ModuleMessageHandler,
		handlers.ModuleChatRoomHandler,
		handlers.ModuleHandler,
		handlers.ModuleActuatorHandler,
	)
)
