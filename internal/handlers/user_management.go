package handlers

import (
	"github.com/go-chi/chi"
	"go.uber.org/fx"
	"net/http"
)

type UserManagementHandler struct{}

func NewUserManagementHandler() *UserManagementHandler {
	return &UserManagementHandler{}
}

func registerUserManagementHandlers(u *UserManagementHandler, h *Handler) {
	h.Runner.Post("/auth/login", u.login)
	h.Runner.Post("/auth/signup", u.signup)

	h.Runner.Group(func(r chi.Router) {
		r.Use(tokenValidationMiddleware)
		r.Route("/user", func(r chi.Router) {
			r.Get("/info", u.getUserInfo)
			r.Get("/friends", u.getUserFriends)
		})
		r.Post("/auth/logout", u.logout)
	})
}

func (u *UserManagementHandler) getUserInfo(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Create Chatroom"))
	if err != nil {
		return
	}
}

func (u *UserManagementHandler) getUserFriends(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Create Chatroom"))
	if err != nil {
		return
	}
}

func (u *UserManagementHandler) login(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Create Chatroom"))
	if err != nil {
		return
	}
}

func (u *UserManagementHandler) logout(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Create Chatroom"))
	if err != nil {
		return
	}
}

func (u *UserManagementHandler) signup(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Create Chatroom"))
	if err != nil {
		return
	}
}

var ModuleUserManagementHandler = fx.Invoke(registerUserManagementHandlers)
