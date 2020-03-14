package superhero

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/jace-ys/super-smash-heroes/services/superhero/pkg/postgres"

	pb "github.com/jace-ys/super-smash-heroes/services/superhero/api/superhero"
)

type Server interface {
	Init(ctx context.Context, server pb.SuperheroServiceServer) error
	Serve() error
	Shutdown(ctx context.Context) error
}

type SuperheroService struct {
	logger log.Logger
	db     postgres.Client
}

func NewService(logger log.Logger, dbClient postgres.Client) (*SuperheroService, error) {
	return &SuperheroService{
		logger: logger,
		db:     dbClient,
	}, nil
}

func (s *SuperheroService) StartServer(ctx context.Context, server Server) error {
	if err := server.Init(ctx, s); err != nil {
		return err
	}
	if err := server.Serve(); err != nil {
		return err
	}
	return nil
}

func (s *SuperheroService) Teardown() error {
	if err := s.db.Close(); err != nil {
		return err
	}
	return nil
}
