package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jace-ys/go-library/postgres"
	"golang.org/x/sync/errgroup"
	"gopkg.in/alecthomas/kingpin.v2"

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

	postgres, err := postgres.NewClient(c.database.connectionURL)
	if err != nil {
		exit(err)
	}

	registry := superhero.NewSuperheroAPIClient(c.registry.Token)

	superheroService, err := superhero.NewService(logger, postgres, registry)
	if err != nil {
		exit(err)
	}
	defer superheroService.Teardown()

	grpcServer := server.NewGRPCServer(c.server.Port)
	gatewayProxy := server.NewGatewayProxy(c.proxy.Port, c.proxy.Endpoint)

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		level.Info(logger).Log("event", "server.grpc.started", "port", c.server.Port)
		defer level.Info(logger).Log("event", "server.grpc.stopped")
		return superheroService.StartServer(ctx, grpcServer)
	})
	g.Go(func() error {
		level.Info(logger).Log("event", "gateway.proxy.started", "port", c.proxy.Port)
		defer level.Info(logger).Log("event", "gateway.proxy.stopped")
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
	registry superhero.SuperheroAPIClientConfig
	database struct {
		connectionURL string
	}
}

func parseCommand() *config {
	var c config

	kingpin.Flag("port", "port for the gRPC server").Envar("PORT").Default("8081").IntVar(&c.server.Port)
	kingpin.Flag("gateway-port", "port for the REST gateway proxy").Envar("GATEWAY_PORT").Default("8080").IntVar(&c.proxy.Port)
	kingpin.Flag("database-url", "URL for connecting to Postgres.").Envar("DATABASE_URL").Default("postgres://nintendo:nintendo@127.0.0.1:5432/nintendo").StringVar(&c.database.connectionURL)
	kingpin.Flag("superhero-api-token", "token for authenticating with the Superhero API").Envar("SUPERHERO_API_TOKEN").Required().StringVar(&c.registry.Token)
	kingpin.Parse()

	c.proxy.Endpoint = fmt.Sprintf(":%d", c.server.Port)
	return &c
}

func exit(err error) {
	level.Error(logger).Log("event", "service.fatal", "msg", err)
	os.Exit(1)
}
