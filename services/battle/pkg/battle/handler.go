package battle

import (
	"context"
	"math/rand"
	"time"

	"github.com/go-kit/log/level"
	"google.golang.org/grpc/codes"

	pb "github.com/jace-ys/super-smash-heroes/services/battle/api/battle"
)

func (s *BattleService) GetResult(ctx context.Context, req *pb.GetResultRequest) (*pb.GetResultResponse, error) {
	level.Info(s.logger).Log("event", "result.get.started")
	defer level.Info(s.logger).Log("event", "result.get.finished")

	err := s.validateGetResultRequest(req)
	if err != nil {
		level.Error(s.logger).Log("event", "result.get.failed", "msg", err)
		return nil, s.Error(codes.InvalidArgument, err)
	}

	return &pb.GetResultResponse{
		Winner: s.determineWinner(req.PlayerOne, req.PlayerTwo),
	}, nil
}

func (s *BattleService) validateGetResultRequest(req *pb.GetResultRequest) error {
	switch {
	case req.PlayerOne == nil:
		return ErrRequestInvalid
	case req.PlayerTwo == nil:
		return ErrRequestInvalid
	}
	return nil
}

func (s *BattleService) determineWinner(players ...*pb.Player) int32 {
	scores := make([]float32, len(players))
	for i, player := range players {
		scores[i] = computeScore([]int32{
			player.Intelligence,
			player.Strength,
			player.Speed,
			player.Durability,
			player.Power,
			player.Combat,
		})
	}
	return highestScore(scores)
}

func computeScore(stats []int32) float32 {
	rand.Seed(time.Now().UnixNano())
	var total float32
	for _, stat := range stats {
		total += float32(stat) * rand.Float32()
	}
	return total
}

func highestScore(scores []float32) int32 {
	index := 1
	max := scores[0]
	for i, score := range scores {
		if score > max {
			max = score
			index = i + 1
		}
	}
	return int32(index)
}
