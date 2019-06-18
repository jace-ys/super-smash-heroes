package response

import (
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func HandleGrpcError(w http.ResponseWriter, err error) {
	if s, ok := status.FromError(err); ok {
		log.Printf("gRPC error - %s: %s\n", s.Code(), s.Message())
		switch s.Code() {
		case codes.NotFound:
			SendError(w, http.StatusNotFound, s.Message())
		case codes.AlreadyExists:
			SendError(w, http.StatusBadRequest, err.Error())
		case codes.Internal:
			SendError(w, http.StatusInternalServerError, s.Message())
		default:
			SendError(w, http.StatusInternalServerError, err.Error())
		}
	}
}
