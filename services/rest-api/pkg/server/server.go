package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jace-ys/super-smash-heroes/services/rest-api/pkg/handler"
	"github.com/jace-ys/super-smash-heroes/services/rest-api/pkg/router"
)

type server struct {
	*router.Router
}

func Init() *server {
	s := &server{
		Router: router.NewRouter(),
	}
	s.createEndpoints()
	return s
}

func (s *server) createEndpoints() {
	s.Router.Get("/superheroes", handler.GetAllSuperheroes)
	s.Router.Get("/superheroes/{id}", handler.GetOneSuperhero)
	s.Router.Post("/superheroes", handler.AddSuperhero)
	s.Router.Delete("/superheroes/{id}", handler.DeleteOneSuperhero)
	s.Router.Get("/battle", handler.GetBattleResult)
}

func (s *server) Start(port int) {
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: s.Router,
	}
	log.Printf("Server listening on port %d\n", port)
	log.Fatal(server.ListenAndServe())
}
