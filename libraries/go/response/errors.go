package response

import (
	"log"
	"net/http"

	"google.golang.org/grpc/status"
)

func HandleGrpcError(w http.ResponseWriter, err error) {
	if s, ok := status.FromError(err); !ok {
		log.Println("gRPC error:", s.Code(), s.Message())
		switch s.Code() {
		default:
			SendError(w, http.StatusInternalServerError, err.Error())
		}
	}
}
