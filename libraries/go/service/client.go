package service

import (
	"fmt"
	"google.golang.org/grpc"
	"log"

	"github.com/jace-ys/super-smash-heroes/libraries/go/config"
)

var (
	BattleServerAddress    = fmt.Sprintf("%s:%d", config.Get("service.battle.host").String("localhost"), config.Get("service.battle.port").Int(3000))
	SuperheroServerAddress = fmt.Sprintf("%s:%d", config.Get("service.superhero.host").String("localhost"), config.Get("service.superhero.port").Int(3000))
)

func CreateClientConn(serverAddress string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	log.Println("Connected to gRPC service at address:", serverAddress)
	return conn, nil
}
