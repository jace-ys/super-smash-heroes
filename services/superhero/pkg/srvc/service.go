package srvc

import (
	"google.golang.org/grpc"

	"github.com/jace-ys/super-smash-heroes/services/superhero/pkg/db"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/superhero"
)

type superheroService struct {
	db *db.Client
	*grpc.Server
}

func NewService() *superheroService {
	return &superheroService{
		Server: grpc.NewServer(),
	}
}

func (s *superheroService) Init() error {
	pb.RegisterSuperheroServiceServer(s.Server, s)
	dbClient, err := db.NewClient()
	if err != nil {
		return err
	}
	s.db = dbClient
	return nil
}

func (s *superheroService) Shutdown() {
	s.db.Teardown()
	s.Server.Stop()
}
