package handlers

import (
	"context"
	"github.com/go-chi/chi"
	"go-live-chat/internal/constants"
	"log"
	"net/http"
	"runtime/debug"
)

func tokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(r.Context()))
	})
}

func chatroomCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chatroomId := chi.URLParam(r, string(constants.RoomIDKey))
		ctx := context.WithValue(r.Context(), constants.RoomIDKey, chatroomId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic and stack trace
				log.Printf("panic recovered: %v\n", err)
				log.Println(string(debug.Stack()))

				// Respond with a 500 Internal Server Error
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
