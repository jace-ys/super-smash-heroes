package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"

	"github.com/jace-ys/super-smash-heroes/services/rest-api/pkg/response"
)

type SuperheroServiceClient struct {
	conn *grpc.ClientConn
}

func (c *SuperheroServiceClient) GetAllSuperheroes(w http.ResponseWriter, r *http.Request) {
	response.SendJSON(w, http.StatusOK, map[string]string{"name": "Jace Tan"})
}

func (c *SuperheroServiceClient) GetOneSuperhero(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	response.SendJSON(w, http.StatusOK, map[string]string{"id": id})
}

func (c *SuperheroServiceClient) AddSuperhero(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var s map[string]string
	err := decoder.Decode(&s)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}
	if len(s) > 1 {
		response.SendError(w, http.StatusBadRequest, errInvalidRequest.Error())
		return
	}
	response.SendJSON(w, http.StatusOK, s)
}

func (c *SuperheroServiceClient) DeleteOneSuperhero(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	response.SendJSON(w, http.StatusOK, map[string]string{"id": id})
}
