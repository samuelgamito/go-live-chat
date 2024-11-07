package handlers

import "net/http"

func getMessageHistory(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		return
	}
}
