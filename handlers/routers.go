package handlers

import (
	"github.com/go-chi/chi"
	"net/http"
)

func NoAuthRoutes() http.Handler {
	r := chi.NewRouter()

	r.Get("/ping", healthCheck)

	r.Post("/auth/login", login)

	r.Post("/auth/signup", signup)

	return r
}

func AuthenticatedRoutes() http.Handler {
	r := chi.NewRouter()

	r.Use(tokenValidationMiddleware)

	r.Get("/users", getUsers)
	r.Get("/message/history", getMessageHistory)

	r.Route("/chatrooms", func(r chi.Router) {
		r.Post("/", createChatroom)

		r.Route("/{roomId}", func(r chi.Router) {

			r.Use(chatroomCtx)

			r.Post("/leave", leaveChatroom)
			r.Post("/join", joinChatroom)
		})
	})

	r.Post("/logout", logout)

	return r
}
