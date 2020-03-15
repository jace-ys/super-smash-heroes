package battle

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *BattleService) Error(code codes.Code, err error) error {
	switch code {
	case codes.Internal:
		return status.Errorf(code, "internal server error")
	case codes.Unknown:
		return status.Errorf(code, "unknown error")
	default:
		s, ok := status.FromError(err)
		if !ok {
			return status.Errorf(code, err.Error())
		}
		return status.Errorf(code, s.Message())
	}

}
