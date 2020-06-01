package superhero

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/jace-ys/go-library/postgres"

	pb "github.com/jace-ys/super-smash-heroes/services/superhero/api/superhero"
)

type Server interface {
	Init(ctx context.Context, server pb.SuperheroServiceServer) error
	Serve() error
	Shutdown(ctx context.Context) error
}

type SuperheroService struct {
	pb.UnimplementedSuperheroServiceServer
	logger   log.Logger
	database *postgres.Client
	registry SuperheroRegistry
}

func NewService(logger log.Logger, postgres *postgres.Client, registry SuperheroRegistry) (*SuperheroService, error) {
	return &SuperheroService{
		logger:   logger,
		database: postgres,
		registry: registry,
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
	if err := s.database.Close(); err != nil {
		return err
	}
	return nil
}
