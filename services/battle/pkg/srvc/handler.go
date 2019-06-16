package srvc

import (
	"context"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/battle"
)

func (s *battleService) GetBattleResult(ctx context.Context, br *pb.BattleRequest) (*pb.BattleResult, error) {
	return &pb.BattleResult{VictorId: br.GetId1()}, nil
}
