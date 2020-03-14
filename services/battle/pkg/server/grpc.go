package server

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "github.com/jace-ys/super-smash-heroes/services/battle/api/battle"
)

type GRPCServerConfig struct {
	Port int
}

type GRPCServer struct {
	config *GRPCServerConfig
	server *grpc.Server
}

func NewGRPCServer(port int, opt ...grpc.ServerOption) *GRPCServer {
	return &GRPCServer{
		config: &GRPCServerConfig{
			Port: port,
		},
		server: grpc.NewServer(opt...),
	}
}

func (g *GRPCServer) Init(ctx context.Context, s pb.BattleServiceServer) error {
	pb.RegisterBattleServiceServer(g.server, s)
	return nil
}

func (g *GRPCServer) Serve() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.config.Port))
	if err != nil {
		return fmt.Errorf("grpc server failed to serve: %w", err)
	}
	if err := g.server.Serve(lis); err != nil {
		return fmt.Errorf("grpc server failed to serve: %w", err)
	}
	return nil
}

func (g *GRPCServer) Shutdown(ctx context.Context) error {
	g.server.GracefulStop()
	return nil
}
