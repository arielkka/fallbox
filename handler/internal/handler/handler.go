package handler

import (
	"fmt"
	"github.com/arielkka/fallbox/handler/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
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
	r.e.Use(middleware.Recover())
	r.e.Use(middleware.RequestID())

	config := middleware.JWTConfig{
		Claims: &jwtCustomClaims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}
	r.e.Use(middleware.JWTWithConfig(config))


	r.e.POST(r.cfg.Router.AuthPath, r.auth)

	r.e.GET(r.cfg.Router.GetUserPNG, r.GetUserPNG)
	r.e.POST(r.cfg.Router.PostUserPNG, r.PostUserPNG)
	r.e.DELETE(r.cfg.Router.DeleteUserPNG, r.DeleteUserPNG)

	r.e.GET(r.cfg.Router.GetUserJPG, r.GetUserJPG)
	r.e.POST(r.cfg.Router.PostUserJPG, r.PostUserJPG)
	r.e.DELETE(r.cfg.Router.DeleteUserJPG, r.DeleteUserJPG)
	return r.e.Start(fmt.Sprintf(":%s", r.cfg.Router.Port))
}
