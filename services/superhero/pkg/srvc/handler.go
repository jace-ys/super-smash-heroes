package srvc

import (
	"context"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/superhero"
)

var mock = []*pb.SuperheroResponse{
	&pb.SuperheroResponse{
		Id:       "1",
		FullName: "Oliver Queen",
	},
	&pb.SuperheroResponse{
		Id:       "2",
		FullName: "Kara Zor-El",
	},
	&pb.SuperheroResponse{
		Id:       "3",
		FullName: "Barry Allen",
	},
	&pb.SuperheroResponse{
		Id:       "4",
		FullName: "Peter Parker",
	},
}

func (s *superheroService) GetAllSuperheroes(empty *pb.Empty, stream pb.SuperheroService_GetAllSuperheroesServer) error {
	for _, superhero := range mock {
		if err := stream.Send(superhero); err != nil {
			return err
		}
	}
	return nil
}

func (s *superheroService) GetOneSuperhero(ctx context.Context, id *pb.SuperheroIdRequest) (*pb.SuperheroResponse, error) {
	return &pb.SuperheroResponse{
		Id: id.GetVal(),
	}, nil
}

func (s *superheroService) AddSuperhero(ctx context.Context, kv *pb.SearchRequest) (*pb.SuperheroResponse, error) {
	key := pb.SearchRequest_Key_name[int32(kv.GetKey())]
	val := kv.GetVal()
	return &pb.SuperheroResponse{
		Id: key + val,
	}, nil
}

func (s *superheroService) DeleteOneSuperhero(ctx context.Context, id *pb.SuperheroIdRequest) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
