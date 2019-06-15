package handler

import (
	"net/http"

	"github.com/jace-ys/super-smash-heroes/services/rest-api/pkg/response"
)

type battleRequest struct {
	id1 string
	id2 string
}

func GetBattleResult(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	br, err := verifyBattleRequest(queryParams)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}
	response.SendJSON(w, http.StatusOK, map[string]string{"id1": br.id1, "id2": br.id2})
	return
}

func verifyBattleRequest(q map[string][]string) (*battleRequest, error) {
	if len(q) > 2 {
		return nil, errInvalidRequest
	}
	id1, ok := q["id1"]
	if !ok {
		return nil, errMissingID
	}
	if len(id1) > 1 {
		return nil, errInvalidRequest
	}
	id2, ok := q["id2"]
	if !ok {
		return nil, errMissingID
	}
	if len(id2) > 1 {
		return nil, errInvalidRequest
	}
	return &battleRequest{
		id1: id1[0],
		id2: id2[0],
	}, nil
}
