package handlers

import "net/http"

func createChatroom(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Create Chatroom"))
	if err != nil {
		return
	}
}

func leaveChatroom(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Leave Chatroom"))
	if err != nil {
		return
	}
}

func joinChatroom(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Join Chatroom"))
	if err != nil {
		return
	}
}
