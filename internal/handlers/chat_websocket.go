package handlers

import (
	"github.com/gorilla/websocket"
	"go-live-chat/internal/handlers/ws"
	"go-live-chat/internal/infraestructure/databases"
	"go.uber.org/fx"
	"log"
	"net/http"
)

type ChatWebSocketHandler struct {
	redisClient *databases.RedisClient
}

func NewChatWebSocketHandler(redisClient *databases.RedisClient) *ChatWebSocketHandler {
	return &ChatWebSocketHandler{
		redisClient: redisClient,
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
		http.Error(w, "Missing User ID", http.StatusUnauthorized)
		return
	}

	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		http.Error(w, "Could not open WebSocket connection", http.StatusBadRequest)
		return
	}
	log.Println("WebSocket connection established")

	client := &ws.ConversationWSUseCase{
		Conn:    conn,
		Channel: user,
		Rdb:     h.redisClient.NotifyClientsRedis,
	}

	defer client.Close()

	go client.ListenAndForward(ctx)
	go client.PublishFromWebSocket(ctx)

	select {}
}
