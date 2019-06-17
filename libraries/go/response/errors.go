package response

import (
	"log"
	"net/http"

	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
)

func HandleGrpcError(w http.ResponseWriter, err error) {
	if s, ok := status.FromError(err); ok {
		log.Printf("gRPC error - %s: %s\n", s.Code(), s.Message())
		switch s.Code() {
		case codes.NotFound:
			SendError(w, http.StatusNotFound, s.Message())
		case codes.Internal:
			SendError(w, http.StatusInternalServerError, s.Message())
		default:
			SendError(w, http.StatusInternalServerError, err.Error())
		}
	}
}
