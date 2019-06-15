package srvc

import (
	"google.golang.org/grpc"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/superhero"
)

type superheroService struct {
	*grpc.Server
}

func NewService() *superheroService {
	return &superheroService{
		Server: grpc.NewServer(),
	}
}

func (s *superheroService) Register() {
	pb.RegisterSuperheroServiceServer(s.Server, s)
}
