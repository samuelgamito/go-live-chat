package handlers

import (
	"github.com/go-chi/chi"
	"go.uber.org/fx"
	"net/http"
)

type MessageHandler struct{}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{}
}

func registerMessageHandlers(m *MessageHandler, handlers *Handler) {

	handlers.Runner.Group(func(r chi.Router) {
		r.Use(tokenValidationMiddleware)
		r.Get("/api/message/history", m.getMessageHistory)
	})
}

func (m *MessageHandler) getMessageHistory(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		return
	}
}

var ModuleMessageHandler = fx.Invoke(registerMessageHandlers)
