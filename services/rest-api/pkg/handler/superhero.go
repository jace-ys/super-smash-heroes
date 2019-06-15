package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jace-ys/super-smash-heroes/services/rest-api/pkg/response"
)

func GetAllSuperheroes(w http.ResponseWriter, r *http.Request) {
	response.SendJSON(w, http.StatusOK, map[string]string{"name": "Jace Tan"})
}

func GetOneSuperhero(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	response.SendJSON(w, http.StatusOK, map[string]string{"id": id})
}

func AddSuperhero(w http.ResponseWriter, r *http.Request) {
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

func DeleteOneSuperhero(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	response.SendJSON(w, http.StatusOK, map[string]string{"id": id})
}
