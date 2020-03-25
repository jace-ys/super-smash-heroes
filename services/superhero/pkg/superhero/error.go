package superhero

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrSuperheroExists   = errors.New("superhero already exists")
	ErrSuperheroInvalid  = errors.New("superhero is invalid")
	ErrSuperheroNotFound = errors.New("superhero not found")
	ErrRequestInvalid    = errors.New("invalid request payload")
)

func (s *SuperheroService) Error(code codes.Code, err error) error {
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
