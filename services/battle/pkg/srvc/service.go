package srvc

import (
	"google.golang.org/grpc"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/battle"
)

type battleService struct {
	*grpc.Server
}

func NewService() *battleService {
	return &battleService{
		Server: grpc.NewServer(),
	}
}

func (s *battleService) Init() error {
	pb.RegisterBattleServiceServer(s.Server, s)
	return nil
}

func (s *battleService) Shutdown() {
	s.Server.Stop()
}
