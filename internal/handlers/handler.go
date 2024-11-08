package handlers

import (
	"github.com/go-chi/chi"
	"go.uber.org/fx"
)

type Handler struct {
	Runner *chi.Mux
}

func NewHandler() *Handler {

	r := chi.NewRouter()
	r.Get("/health", healthCheck)
	handler := &Handler{Runner: r}

	return handler
}

var ModuleHandler = fx.Provide(NewHandler)
