package srvc

import (
	"context"

	pb "github.com/jace-ys/super-smash-heroes/api/proto/generated/go/battle"
)

func (s *battleService) GetBattleResult(ctx context.Context, kv *pb.BattleRequest) (*pb.BattleResult, error) {
	return nil, nil
}
