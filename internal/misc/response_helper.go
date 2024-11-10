package misc

import (
	"encoding/json"
	"net/http"
)

func WriteJSONResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if body == nil {
		return
	}
	if jsonData, err := json.Marshal(body); err == nil {
		_, _ = w.Write(jsonData)
	} else {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
