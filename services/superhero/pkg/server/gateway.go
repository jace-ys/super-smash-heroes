package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	gw "github.com/jace-ys/super-smash-heroes/services/superhero/api/superhero"
)

type GatewayProxyConfig struct {
	Port     int
	Endpoint string
}

type GatewayProxy struct {
	config *GatewayProxyConfig
	server *http.Server
}

func NewGatewayProxy(port int, endpoint string) *GatewayProxy {
	return &GatewayProxy{
		config: &GatewayProxyConfig{
			Port:     port,
			Endpoint: endpoint,
		},
		server: &http.Server{
			Handler: runtime.NewServeMux(),
			Addr:    fmt.Sprintf(":%d", port),
		},
	}
}

func (g *GatewayProxy) Init(ctx context.Context, s gw.SuperheroServiceServer) error {
	runtime.HTTPError = HTTPError
	err := gw.RegisterSuperheroServiceHandlerFromEndpoint(
		ctx,
		g.server.Handler.(*runtime.ServeMux),
		g.config.Endpoint,
		[]grpc.DialOption{grpc.WithInsecure()},
	)
	if err != nil {
		return fmt.Errorf("gateway proxy failed to initialize: %w", err)
	}
	return nil
}

func (g *GatewayProxy) Serve() error {
	if err := g.server.ListenAndServe(); err != nil {
		return fmt.Errorf("gateway proxy failed to serve: %w", err)
	}
	return nil
}

func (g *GatewayProxy) Shutdown(ctx context.Context) error {
	if err := g.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("gateway proxy failed to shutdown: %w", err)
	}
	return nil
}

type httpError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func HTTPError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	httpErr := convertGRPCError(err)
	w.Header().Set("Content-Type", marshaler.ContentType())
	w.WriteHeader(httpErr.Error.Code)
	json.NewEncoder(w).Encode(httpErr)
}

func convertGRPCError(err error) *httpError {
	var httpErr httpError

	s, ok := status.FromError(err)
	if !ok {
		httpErr.Error.Message = err.Error()
	}
	httpErr.Error.Message = s.Message()
	httpErr.Error.Code = runtime.HTTPStatusFromCode(s.Code())

	return &httpErr
}
