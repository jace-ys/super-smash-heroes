package integration

import (
	"google.golang.org/grpc"

	"github.com/jace-ys/super-smash-heroes/test/api/battle"
	"github.com/jace-ys/super-smash-heroes/test/api/superhero"
)

func NewSuperheroServiceClient(address string) (superhero.SuperheroServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return superhero.NewSuperheroServiceClient(conn), nil
}

func NewBattleServiceClient(address string) (battle.BattleServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return battle.NewBattleServiceClient(conn), nil
}
