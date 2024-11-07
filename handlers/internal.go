package handlers

import "net/http"

func healthCheck(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Health Check"))
	if err != nil {
		return
	}
}
