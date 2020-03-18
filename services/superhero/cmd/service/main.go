package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"golang.org/x/sync/errgroup"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/jace-ys/super-smash-heroes/services/superhero/pkg/postgres"
	"github.com/jace-ys/super-smash-heroes/services/superhero/pkg/server"
	"github.com/jace-ys/super-smash-heroes/services/superhero/pkg/superhero"
)

var logger log.Logger

func main() {
	c := parseCommand()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "source", log.DefaultCaller)

	postgresClient, err := postgres.NewPostgresClient(c.database.Host, c.database.User, c.database.Password, c.database.Database)
	if err != nil {
		exit(err)
	}

	superheroRegistry := superhero.NewSuperheroAPIClient(c.registry.Token)

	superheroService, err := superhero.NewService(logger, postgresClient, superheroRegistry)
	if err != nil {
		exit(err)
	}
	defer superheroService.Teardown()

	grpcServer := server.NewGRPCServer(c.server.Port)
	gatewayProxy := server.NewGatewayProxy(c.proxy.Port, c.proxy.Endpoint)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		level.Info(logger).Log("event", "grpc_server.started", "port", c.server.Port)
		defer level.Info(logger).Log("event", "grpc_server.stopped")
		return superheroService.StartServer(ctx, grpcServer)
	})
	g.Go(func() error {
		level.Info(logger).Log("event", "gateway_proxy.started", "port", c.proxy.Port)
		defer level.Info(logger).Log("event", "gateway_proxy.stopped")
		return superheroService.StartServer(ctx, gatewayProxy)
	})
	g.Go(func() error {
		select {
		case <-ctx.Done():
			grpcServer.Shutdown(ctx)
			gatewayProxy.Shutdown(ctx)
			return ctx.Err()
		}
	})

	if err := g.Wait(); err != nil {
		exit(err)
	}
}

type config struct {
	server   server.GRPCServerConfig
	proxy    server.GatewayProxyConfig
	database postgres.PostgresClientConfig
	registry superhero.SuperheroAPIClientConfig
}

func parseCommand() *config {
	var c config

	kingpin.Flag("port", "port for the gRPC server").Default("8080").IntVar(&c.server.Port)
	kingpin.Flag("gateway-port", "port for the REST gateway proxy").Default("8081").IntVar(&c.proxy.Port)
	kingpin.Flag("postgres-host", "host for connecting to Postgres").Default("127.0.0.1:5432").StringVar(&c.database.Host)
	kingpin.Flag("postgres-user", "user for connecting to Postgres").Default("postgres").StringVar(&c.database.User)
	kingpin.Flag("postgres-password", "password for connecting to Postgres").Required().StringVar(&c.database.Password)
	kingpin.Flag("postgres-db", "database for connecting to Postgres").Default("postgres").StringVar(&c.database.Database)
	kingpin.Flag("superhero-api-token", "token for authenticating with the Superhero API").Required().StringVar(&c.registry.Token)
	kingpin.Parse()

	c.proxy.Endpoint = fmt.Sprintf(":%d", c.server.Port)
	return &c
}

func exit(err error) {
	level.Error(logger).Log("event", "service.fatal", "msg", err)
	os.Exit(1)
}
