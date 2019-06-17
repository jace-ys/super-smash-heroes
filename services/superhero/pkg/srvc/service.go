package srvc

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/jace-ys/super-smash-heroes/services/superhero/pkg/psql"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/superhero"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
	table    = "superheroes"
)

type superheroService struct {
	psql *psql.Client
	*grpc.Server
}

func NewService() *superheroService {
	return &superheroService{
		Server: grpc.NewServer(),
	}
}

func (s *superheroService) Init() {
	pb.RegisterSuperheroServiceServer(s.Server, s)
	psqlSourceInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable dbname=%s", host, port, user, password, dbname)
	psqlClient, err := psql.Open(psqlSourceInfo)
	if err != nil {
		log.Fatal(err)
	}
	s.psql = psqlClient
}

func (s *superheroService) Shutdown() {
	s.psql.Close()
	s.Server.Stop()
}
