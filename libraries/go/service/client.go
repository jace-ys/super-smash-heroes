package service

import (
	"google.golang.org/grpc"
)

const (
	BattleServerAddress    = "service.battle:3000"
	SuperheroServerAddress = "service.superhero:3000"
)

func CreateClientConn(serverAddress string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
