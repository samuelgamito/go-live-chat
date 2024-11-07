package handlers

import "net/http"

func getUsers(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Create Chatroom"))
	if err != nil {
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Create Chatroom"))
	if err != nil {
		return
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Create Chatroom"))
	if err != nil {
		return
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Create Chatroom"))
	if err != nil {
		return
	}
}
