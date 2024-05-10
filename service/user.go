package service

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/savi2w/pupper/config"
	"github.com/savi2w/pupper/model"
	"github.com/savi2w/pupper/presenter/req"
	"github.com/savi2w/pupper/repo"
)

type UserService struct {
	config *config.Config
	logger *zerolog.Logger
	repo   *repo.RepoManager
}

func newUserService(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *UserService {
	return &UserService{
		config: cfg,
		logger: logger,
		repo:   repo,
	}
}

func (s *UserService) New(ctx context.Context, r *req.NewUser) error {
	user := &model.User{
		FirstName:      r.FirstName,
		LastName:       r.LastName,
		DocumentNumber: r.DocumentNumber,
		Balance:        10000,
	}

	s.logger.Info().Msgf("creating user %s", user.DocumentNumber)

	return s.repo.MySQL.User.Insert(ctx, user)
}
