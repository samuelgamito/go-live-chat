package handlers

import (
	"github.com/gorilla/websocket"
	"go.uber.org/fx"
	"log"
	"net/http"
)

type ChatWebSocketHandler struct {
}

func NewChatWebSocketHandler() *ChatWebSocketHandler {
	return &ChatWebSocketHandler{}
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

	conn, err := upgrade.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Error upgrading connection:", err)
		http.Error(w, "Could not open WebSocket connection", http.StatusBadRequest)
		return
	}

	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("Error closing WebSocket connection:", err)
		}
	}(conn)

	log.Println("WebSocket connection established")

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		log.Printf("Received: %s\n", message)

		// Echo the message back to the client.
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}
