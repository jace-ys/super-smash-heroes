package handler

import (
	"time"

	"github.com/jace-ys/super-smash-heroes/libraries/go/service"
)

var (
	timeout = time.Second * 10
)

type Handler struct {
	*BattleServiceClient
	*SuperheroServiceClient
}

func InitServiceClients() (*Handler, error) {
	battleServiceConn, err := service.CreateClientConn(service.BattleServerAddress)
	if err != nil {
		return nil, err
	}
	superheroServiceConn, err := service.CreateClientConn(service.SuperheroServerAddress)
	if err != nil {
		return nil, err
	}
	return &Handler{
		&BattleServiceClient{battleServiceConn},
		&SuperheroServiceClient{superheroServiceConn},
	}, nil
}

func (h *Handler) TeardownClients() {
	h.BattleServiceClient.conn.Close()
	h.SuperheroServiceClient.conn.Close()
}
