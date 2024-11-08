package handlers

import (
	"github.com/go-chi/chi"
	"go.uber.org/fx"
	"net/http"
)

type ChatRoomHandler struct{}

func NewChatRoomHandler() *ChatRoomHandler {
	return &ChatRoomHandler{}
}

func registerChatRoomHandlers(c *ChatRoomHandler, h *Handler) {
	h.Runner.Group(func(r chi.Router) {
		r.Use(tokenValidationMiddleware)
		r.Route("/api/chatrooms", func(r chi.Router) {
			r.Post("/", c.createChatroom)

			r.Route("/{roomId}", func(r chi.Router) {

				r.Use(chatroomCtx)

				r.Post("/leave", c.leaveChatroom)
				r.Post("/join", c.joinChatroom)
			})
		})
	})
}

func (c *ChatRoomHandler) createChatroom(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Create Chatroom"))
	if err != nil {
		return
	}
}

func (c *ChatRoomHandler) leaveChatroom(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Leave Chatroom"))
	if err != nil {
		return
	}
}

func (c *ChatRoomHandler) joinChatroom(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Join Chatroom"))
	if err != nil {
		return
	}
}

var ModuleChatRoomHandler = fx.Invoke(registerChatRoomHandlers)
