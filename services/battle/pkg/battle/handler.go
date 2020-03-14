package battle

import (
	"context"

	pb "github.com/jace-ys/super-smash-heroes/services/battle/api/battle"
)

func (s *BattleService) GetResult(ctx context.Context, req *pb.GetResultRequest) (*pb.GetResultResponse, error) {
	return &pb.GetResultResponse{}, nil
}
