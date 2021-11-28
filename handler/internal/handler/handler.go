package handler

import (
	"github.com/arielkka/fallbox/handler/config"
	"github.com/labstack/echo/v4"
)

type Router struct {
	e   *echo.Echo
	cfg *config.Config
}

func NewRouter(cfg *config.Config) *Router {
	return &Router{
		e:   echo.New(),
		cfg: cfg,
	}
}
