package server

import (
	"fmt"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/savi2w/pupper/config"
	"github.com/savi2w/pupper/server/controller"
	"github.com/savi2w/pupper/server/middleware"
	"github.com/savi2w/pupper/server/router"
)

var (
	instance *Server
	once     sync.Once
)

type Server struct {
	cfg    *config.Config
	svr    *echo.Echo
	logger *zerolog.Logger
	ctrl   *controller.Controller
}

func New(cfg *config.Config, logger *zerolog.Logger, ctrl *controller.Controller) *Server {
	once.Do(func() {
		svr := echo.New()

		svr.HideBanner = true
		svr.HidePort = true

		middleware.SetMiddlewares(svr, cfg)
		router.Register(cfg, svr, ctrl)

		instance = &Server{
			cfg:    cfg,
			svr:    svr,
			logger: logger,
			ctrl:   ctrl,
		}
	})

	return instance
}

func (s *Server) Start() error {
	s.logger.Info().Msg("starting server")

	if err := s.svr.Start(fmt.Sprintf(":%d", s.cfg.InternalConfig.ServerPort)); err != nil {
		return err
	}

	return nil
}
