package handlers

import (
	"github.com/gorilla/websocket"
	"go-live-chat/internal/handlers/dto"
	"go-live-chat/internal/handlers/ws"
	"go-live-chat/internal/infraestructure/databases"
	"go-live-chat/internal/misc"
	"go.uber.org/fx"
	"log"
	"net/http"
)

type ChatWebSocketHandler struct {
	redisClient *databases.RedisClient
	useCases    map[string]ws.ConversationUseCase
}

func NewChatWebSocketHandler(
	redisClient *databases.RedisClient,
	useCases map[string]ws.ConversationUseCase,
) *ChatWebSocketHandler {

	return &ChatWebSocketHandler{
		redisClient: redisClient,
		useCases:    useCases,
	}
}

func registerChatWebSocketHandlers(c *ChatWebSocketHandler, h *Handler) {
	h.Runner.Get("/chat/ws", c.ServeHTTP)
}

var ModuleChatWebSocketHandler = fx.Invoke(registerChatWebSocketHandlers)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *ChatWebSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := r.Header.Get("X-User-ID")
	if user == "" {
		errResp := dto.ErrorResponse{
			Body: dto.ErrorBodyDTO{
				Messages: []string{"Missing User ID"},
			},
		}
		misc.WriteJSONResponse(w, http.StatusUnauthorized, errResp)
		return
	}

	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		errResp := dto.ErrorResponse{
			Body: dto.ErrorBodyDTO{
				Messages: []string{"Could not open WebSocket connection"},
			},
		}
		misc.WriteJSONResponse(w, http.StatusUnauthorized, errResp)
		return
	}
	log.Println("WebSocket connection established")

	client := &ws.ConversationWSUseCase{
		Conn:    conn,
		Channel: user,
		Rdb:     h.redisClient.NotifyClientsRedis,
		UseCase: h.useCases,
	}

	defer client.Close()

	go client.ListenAndForward(ctx)
	go client.PublishFromWebSocket(ctx)

	select {}
}
