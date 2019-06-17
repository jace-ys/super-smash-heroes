package handler

import (
	"time"

	"google.golang.org/grpc"
)

var (
	timeout                 = time.Second * 10
	battleServerAddress     = "localhost:3000"
	supherheroServerAddress = "localhost:3001"
)

type Handler struct {
	*BattleServiceClient
	*SuperheroServiceClient
}

func InitServiceClients() (*Handler, error) {
	battleClient, err := createBattleClient()
	if err != nil {
		return nil, err
	}
	superheroClient, err := createSuperheroClient()
	if err != nil {
		return nil, err
	}
	return &Handler{
		battleClient,
		superheroClient,
	}, nil
}

func (h *Handler) TeardownClients() {
	h.BattleServiceClient.conn.Close()
	h.SuperheroServiceClient.conn.Close()
}

func createBattleClient() (*BattleServiceClient, error) {
	conn, err := grpc.Dial(battleServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &BattleServiceClient{conn}, nil
}

func createSuperheroClient() (*SuperheroServiceClient, error) {
	conn, err := grpc.Dial(supherheroServerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &SuperheroServiceClient{conn}, nil
}
