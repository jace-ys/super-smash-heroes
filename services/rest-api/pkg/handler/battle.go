package handler

import (
	"context"
	"net/http"

	"google.golang.org/grpc"

	"github.com/jace-ys/super-smash-heroes/libraries/go/errors"
	"github.com/jace-ys/super-smash-heroes/libraries/go/response"
	"github.com/jace-ys/super-smash-heroes/libraries/go/utils"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/battle"
)

type BattleServiceClient struct {
	conn *grpc.ClientConn
}

type battleRequest struct {
	id1 int32
	id2 int32
}

func (c *BattleServiceClient) GetBattleResult(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	battleReq, err := verifyBattleRequest(queryParams)
	if err != nil {
		response.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	client := pb.NewBattleServiceClient(c.conn)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	battleRes, err := client.GetBattleResult(ctx, &pb.BattleRequest{
		Id1: battleReq.id1,
		Id2: battleReq.id2,
	})
	if err != nil {
		response.HandleGrpcError(w, err)
		return
	}

	response.SendJSON(w, http.StatusOK, response.EncodePbToJSON(battleRes))
	return
}

func verifyBattleRequest(q map[string][]string) (*battleRequest, error) {
	if len(q) > 2 {
		return nil, errors.InvalidRequest
	}
	id1, ok := q["id1"]
	if !ok {
		return nil, errors.MissingID
	}
	if len(id1) > 1 {
		return nil, errors.InvalidRequest
	}
	id2, ok := q["id2"]
	if !ok {
		return nil, errors.MissingID
	}
	if len(id2) > 1 {
		return nil, errors.InvalidRequest
	}
	return &battleRequest{
		id1: int32(utils.Atoi(id1[0])),
		id2: int32(utils.Atoi(id2[0])),
	}, nil
}
