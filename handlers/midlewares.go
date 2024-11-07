package handlers

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"go-live-chat/constants"
	"net/http"
)

func tokenValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("passou")
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
