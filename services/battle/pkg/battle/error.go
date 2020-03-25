package battle

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrRequestInvalid = errors.New("invalid request payload")
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
