package user

import (
	"context"
	"proj/pkg/logging"
)

type Service struct {
	starage  Storage
	logger  *logging.Logger
}

func (s *Service) Create(ctx context.Context, dto CreateUserDTO) (u User, err error) {
	// TODO
	return
}