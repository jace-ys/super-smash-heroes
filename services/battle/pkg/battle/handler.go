package battle

import (
	"context"
	"math/rand"
	"time"

	pb "github.com/jace-ys/super-smash-heroes/services/battle/api/battle"
)

func (s *BattleService) GetResult(ctx context.Context, req *pb.GetResultRequest) (*pb.GetResultResponse, error) {
	winner := determineWinner(req.PlayerOne, req.PlayerTwo)
	return &pb.GetResultResponse{
		Winner: winner,
	}, nil
}

func determineWinner(players ...*pb.Player) int32 {
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
