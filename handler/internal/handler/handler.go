package handler

import (
	"fmt"
	"github.com/arielkka/fallbox/handler/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type router struct {
	e   *echo.Echo
	cfg *config.Config
}

func NewRouter(cfg *config.Config) *router {
	return &router{
		e:   echo.New(),
		cfg: cfg,
	}
}

func (r *router) Run() error {
	r.e.Use(middleware.Logger())
	r.e.Use(middleware.RequestID())

	return r.e.Start(fmt.Sprintf(":%s", r.cfg.Router.Port))
}
