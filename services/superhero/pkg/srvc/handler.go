package srvc

import (
	"context"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/superhero"
)

func (s *superheroService) AddSuperhero(ctx context.Context, kv *pb.SearchKeyVal) (*pb.Superhero, error) {
	return nil, nil
}

func (s *superheroService) DeleteOneSuperhero(ctx context.Context, id *pb.SuperheroID) (*pb.Empty, error) {
	return nil, nil
}

func (s *superheroService) GetAllSuperheroes(empty *pb.Empty, stream pb.SuperheroService_GetAllSuperheroesServer) error {
	return nil
}

func (s *superheroService) GetOneSuperhero(ctx context.Context, id *pb.SuperheroID) (*pb.Superhero, error) {
	return nil, nil
}
