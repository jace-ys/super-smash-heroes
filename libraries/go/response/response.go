package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func SendJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func SendError(w http.ResponseWriter, code int, message string) {
	SendJSON(w, code, map[string]Error{"error": Error{code, message}})
}
