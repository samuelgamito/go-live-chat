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
	handler := &Handler{Runner: r}
	r.Use(Recovery)
	return handler
}

var ModuleHandler = fx.Provide(NewHandler)
