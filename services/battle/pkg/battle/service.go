package battle

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/jace-ys/go-library/postgres"

	pb "github.com/jace-ys/super-smash-heroes/services/battle/api/battle"
)

type Server interface {
	Init(ctx context.Context, server pb.BattleServiceServer) error
	Serve() error
	Shutdown(ctx context.Context) error
}

type BattleService struct {
	pb.UnimplementedBattleServiceServer
	logger   log.Logger
	database *postgres.Client
}

func NewService(logger log.Logger, database *postgres.Client) (*BattleService, error) {
	return &BattleService{
		logger:   logger,
		database: database,
	}, nil
}

func (s *BattleService) StartServer(ctx context.Context, server Server) error {
	if err := server.Init(ctx, s); err != nil {
		return err
	}
	if err := server.Serve(); err != nil {
		return err
	}
	return nil
}

func (s *BattleService) Teardown() error {
	if err := s.database.Close(); err != nil {
		return err
	}
	return nil
}
