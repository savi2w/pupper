package service

import (
	"github.com/rs/zerolog"
	"github.com/savi2w/pupper/config"
	"github.com/savi2w/pupper/repo"
)

type Service struct {
	User *UserService
}

func New(cfg *config.Config, logger *zerolog.Logger, repo *repo.RepoManager) *Service {
	return &Service{
		User: newUserService(cfg, logger, repo),
	}
}
