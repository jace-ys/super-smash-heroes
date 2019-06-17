package srvc

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/jace-ys/super-smash-heroes/libraries/go/response"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/superhero"
)

var (
	baseSuperheroAPI = fmt.Sprintf("%s/%s", "https://superheroapi.com/api", os.Getenv("SUPERHERO_API_ACCESS_TOKEN"))
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

func (s *superheroService) AddSuperhero(ctx context.Context, sr *pb.SearchRequest) (*pb.SuperheroResponse, error) {
	resp, err := http.Get(fmt.Sprintf("%s/search/%s", baseSuperheroAPI, sr.GetAlterEgo()))
	if err != nil {
		return nil, status.Error(codes.Internal, errInternalServerError.Error())
	}
	defer resp.Body.Close()

	var res Response
	err = response.Decode(resp.Body, &res)
	if err != nil {
		return nil, status.Error(codes.Internal, errInternalServerError.Error())
	}
	if res.Response == "error" {
		return nil, status.Error(codes.NotFound, errSuperheroDoesNotExist.Error())
	}

	for _, superhero := range res.Results {
		if superhero.AlterEgo == sr.GetAlterEgo() && superhero.Biography.FullName == sr.GetFullName() {
			return &pb.SuperheroResponse{
				FullName: sr.GetFullName(),
				AlterEgo: sr.GetAlterEgo(),
			}, nil
		}
	}
	return nil, status.Error(codes.NotFound, errSuperheroDoesNotExist.Error())
}

func (s *superheroService) DeleteOneSuperhero(ctx context.Context, id *pb.SuperheroIdRequest) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
