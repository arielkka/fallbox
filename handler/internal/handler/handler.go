package handler

import (
	"fmt"
	"github.com/arielkka/fallbox/handler/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type router struct {
	e       *echo.Echo
	cfg     *config.Config
	service Service
}

func NewRouter(cfg *config.Config, service Service) *router {
	return &router{
		e:       echo.New(),
		cfg:     cfg,
		service: service,
	}
}

func (r *router) Run() error {
	r.e.Use(middleware.Logger())
	r.e.Use(middleware.RequestID())

	r.e.POST(r.cfg.Router.AuthPath, r.auth)

	r.e.GET(r.cfg.Router.GetAllUserPNG, r.GetAllUserPNG)
	r.e.GET(r.cfg.Router.GetUserPNG, r.GetUserPNG)
	r.e.POST(r.cfg.Router.PostUserPNG, r.PostUserPNG)
	r.e.DELETE(r.cfg.Router.DeleteUserPNG, r.DeleteUserPNG)

	r.e.GET(r.cfg.Router.GetAllUserJPG, r.GetAllUserJPG)
	r.e.GET(r.cfg.Router.GetUserJPG, r.GetUserJPG)
	r.e.POST(r.cfg.Router.PostUserJPG, r.PostUserJPG)
	r.e.DELETE(r.cfg.Router.DeleteUserJPG, r.DeleteUserJPG)
	return r.e.Start(fmt.Sprintf(":%s", r.cfg.Router.Port))
}
