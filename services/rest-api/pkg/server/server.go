package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jace-ys/super-smash-heroes/libraries/go/router"
	"github.com/jace-ys/super-smash-heroes/services/rest-api/pkg/handler"
)

type server struct {
	*router.Router
	*handler.Handler
}

func Init() (*server, error) {
	handler, err := handler.InitServiceClients()
	if err != nil {
		return nil, err
	}
	s := &server{
		Router:  router.NewRouter(),
		Handler: handler,
	}
	s.createEndpoints()
	return s, nil
}

func (s *server) createHandlers() {

}

func (s *server) createEndpoints() {
	s.Router.Get("/superheroes", s.Handler.SuperheroServiceClient.GetAllSuperheroes)
	s.Router.Get("/superheroes/{id}", s.Handler.SuperheroServiceClient.GetOneSuperhero)
	s.Router.Post("/superheroes", s.Handler.SuperheroServiceClient.AddSuperhero)
	s.Router.Delete("/superheroes/{id}", s.Handler.SuperheroServiceClient.DeleteOneSuperhero)

	s.Router.Get("/battle", s.Handler.BattleServiceClient.GetBattleResult)
}

func (s *server) Start(port int) {
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: s.Router,
	}
	log.Printf("Server listening on port %d\n", port)
	log.Fatal(server.ListenAndServe())
}

func (s *server) Shutdown() {
	s.Handler.TeardownClients()
}
