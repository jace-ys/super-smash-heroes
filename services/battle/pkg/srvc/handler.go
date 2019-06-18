package srvc

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/jace-ys/super-smash-heroes/libraries/go/errors"
	"github.com/jace-ys/super-smash-heroes/libraries/go/response"
	"github.com/jace-ys/super-smash-heroes/libraries/go/service"
	"github.com/jace-ys/super-smash-heroes/libraries/go/superhero"
	"github.com/jace-ys/super-smash-heroes/libraries/go/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/battle"
	superheroPb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/superhero"
)

func (s *battleService) GetBattleResult(ctx context.Context, br *pb.BattleRequest) (*pb.BattleResponse, error) {
	superhero1, err := getOneSuperhero(br.GetId1())
	if err != nil {
		return nil, err
	}
	superhero2, err := getOneSuperhero(br.GetId2())
	if err != nil {
		return nil, err
	}

	power1, err := getPowerStats(superhero1.FullName, superhero1.AlterEgo)
	if err != nil {
		return nil, err
	}
	power2, err := getPowerStats(superhero2.FullName, superhero2.AlterEgo)
	if err != nil {
		return nil, err
	}

	var winnerId int32
	winner := determineWinner(power1, power2)
	switch winner {
	case 1:
		winnerId = br.GetId1()
	case 2:
		winnerId = br.GetId2()
	default:
		return nil, errors.UndeterminedWinner
	}
	return &pb.BattleResponse{WinnerId: winnerId}, nil
}

func getOneSuperhero(id int32) (*superheroPb.SuperheroResponse, error) {
	conn, err := service.CreateClientConn(service.SuperheroServerAddress)
	client := superheroPb.NewSuperheroServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	s, err := client.GetOneSuperhero(ctx, &superheroPb.SuperheroIdRequest{
		Val: id,
	})
	if err != nil {
		return nil, status.Error(codes.NotFound, errors.SuperheroNotFound.Error())
	}
	return s, nil
}

func getPowerStats(fullName, alterEgo string) (*superhero.Powerstats, error) {
	baseUri := superhero.GetBaseUri()
	if baseUri == "" {
		return nil, errors.MissingAccessToken
	}
	resp, err := http.Get(fmt.Sprintf("%s/search/%s", baseUri, alterEgo))
	if err != nil {
		return nil, errors.InternalServerError
	}
	defer resp.Body.Close()

	var r superhero.Response
	err = response.Decode(resp.Body, &r)
	if err != nil {
		return nil, errors.InternalServerError
	}
	if r.Response == "error" {
		return nil, errors.SuperheroDoesNotExist
	}

	for _, superhero := range r.Results {
		if strings.EqualFold(superhero.AlterEgo, alterEgo) && strings.EqualFold(superhero.Biography.FullName, fullName) {
			return &superhero.Powerstats, nil
		}
	}
	return nil, errors.SuperheroDoesNotExist
}

func determineWinner(p1, p2 *superhero.Powerstats) int {
	p1Total := sumWithMultiplier([]string{p1.Intelligence, p1.Strength, p1.Speed, p1.Durability, p1.Power, p1.Combat})
	p2Total := sumWithMultiplier([]string{p2.Intelligence, p2.Strength, p2.Speed, p2.Durability, p2.Power, p2.Combat})
	if p1Total > p2Total {
		return 1
	}
	return 2
}

func sumWithMultiplier(stats []string) float32 {
	rand.Seed(time.Now().UnixNano())
	var total float32
	for i, stat := range stats {
		total += float32(utils.Atoi(stat)) * rand.Float32()
	}
	return total
}
