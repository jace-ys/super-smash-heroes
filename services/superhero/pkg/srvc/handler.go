package srvc

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/jace-ys/super-smash-heroes/libraries/go/errors"
	"github.com/jace-ys/super-smash-heroes/libraries/go/response"
	"github.com/jace-ys/super-smash-heroes/libraries/go/superhero"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/superhero"
)

func (s *superheroService) GetAllSuperheroes(empty *pb.Empty, stream pb.SuperheroService_GetAllSuperheroesServer) error {
	superheroes, err := s.psql.GetAll()
	if err != nil {
		return status.Error(codes.NotFound, errors.SuperheroesNotFound.Error())
	}
	for _, superhero := range superheroes {
		if err := stream.Send(superhero); err != nil {
			return err
		}
	}
	return nil
}

func (s *superheroService) GetOneSuperhero(ctx context.Context, id *pb.SuperheroIdRequest) (*pb.SuperheroResponse, error) {
	superhero, err := s.psql.FindById(id.GetVal())
	if err != nil {
		return nil, status.Error(codes.NotFound, errors.SuperheroNotFound.Error())
	}
	return superhero, nil
}

func (s *superheroService) AddSuperhero(ctx context.Context, sr *pb.SearchRequest) (*pb.SuperheroResponse, error) {
	baseUri := superhero.GetBaseUri()
	if baseUri == "" {
		return nil, status.Error(codes.Internal, errors.MissingAccessToken.Error())
	}
	resp, err := http.Get(fmt.Sprintf("%s/search/%s", baseUri, sr.GetAlterEgo()))
	if err != nil {
		return nil, status.Error(codes.Internal, errors.InternalServerError.Error())
	}
	defer resp.Body.Close()

	var r superhero.Response
	err = response.Decode(resp.Body, &r)
	if err != nil {
		return nil, status.Error(codes.Internal, errors.InternalServerError.Error())
	}
	if r.Response == "error" {
		return nil, status.Error(codes.NotFound, errors.SuperheroDoesNotExist.Error())
	}

	for _, superhero := range r.Results {
		if strings.EqualFold(superhero.AlterEgo, sr.GetAlterEgo()) && strings.EqualFold(superhero.Biography.FullName, sr.GetFullName()) {
			new := &pb.SuperheroResponse{
				FullName: superhero.Biography.FullName,
				AlterEgo: superhero.AlterEgo,
				ImageUrl: superhero.Image.URL,
			}
			id, err := s.psql.Insert(new)
			if err != nil {
				switch err {
				case errors.SuperheroExists:
					return nil, status.Error(codes.AlreadyExists, err.Error())
				default:
					return nil, status.Error(codes.Internal, err.Error())
				}
			}
			new.Id = id
			return new, nil
		}
	}
	return nil, status.Error(codes.NotFound, errors.SuperheroDoesNotExist.Error())
}

func (s *superheroService) DeleteOneSuperhero(ctx context.Context, id *pb.SuperheroIdRequest) (*pb.Empty, error) {
	err := s.psql.Delete(id.GetVal())
	if err != nil {
		return nil, status.Error(codes.Internal, errors.InternalServerError.Error())
	}
	return &pb.Empty{}, nil
}
