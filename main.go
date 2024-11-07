package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"go-live-chat/handlers"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Mount("/", handlers.NoAuthRoutes())
	r.Mount("/api", handlers.AuthenticatedRoutes())
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		fmt.Println(err)
		return
	}
}
