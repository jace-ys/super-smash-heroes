package superhero

import (
	"context"

	pb "github.com/jace-ys/super-smash-heroes/services/superhero/api/superhero"
)

func (s *SuperheroService) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	return &pb.ListResponse{}, nil
}

func (s *SuperheroService) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	return &pb.GetResponse{}, nil
}

func (s *SuperheroService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{}, nil
}

func (s *SuperheroService) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	return &pb.DeleteResponse{}, nil
}
