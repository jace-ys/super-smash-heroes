package integration

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"

	"github.com/jace-ys/super-smash-heroes/test/api/battle"
	"github.com/jace-ys/super-smash-heroes/test/api/superhero"
)

var (
	battleService    battle.BattleServiceClient
	superheroService superhero.SuperheroServiceClient
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integration Suite")
}

var _ = BeforeSuite(func() {
	conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
	Expect(err).NotTo(HaveOccurred())
	battleService = battle.NewBattleServiceClient(conn)

	conn, err = grpc.Dial("localhost:5001", grpc.WithInsecure())
	Expect(err).NotTo(HaveOccurred())
	superheroService = superhero.NewSuperheroServiceClient(conn)
})
