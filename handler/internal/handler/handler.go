package handler

import (
	"fmt"
	"os"

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
	r.e.Use(middleware.Recover())
	r.e.Use(middleware.RequestID())

	config := middleware.JWTConfig{
		Skipper:     r.SkipJWTMiddleware,
		TokenLookup: "cookie:fallbox.jwt_token",
		SigningKey:  []byte(os.Getenv("JWT_SECRET_KEY")),
	}
	r.e.Use(middleware.JWTWithConfig(config))

	r.e.POST(r.cfg.Router.AuthPath, r.auth)
	r.e.POST(r.cfg.Router.Registration, r.registration)

	r.e.GET(r.cfg.Router.GetUserExcel, r.GetUserExcel)
	r.e.POST(r.cfg.Router.PostUserExcel, r.PostUserExcel)
	r.e.DELETE(r.cfg.Router.DeleteUserExcel, r.DeleteUserExcel)

	r.e.GET(r.cfg.Router.GetUserTxt, r.GetUserTxt)
	r.e.POST(r.cfg.Router.PostUserTxt, r.PostUserTxt)
	r.e.DELETE(r.cfg.Router.DeleteUserTxt, r.DeleteUserTxt)
	return r.e.Start(fmt.Sprintf(":%s", r.cfg.Router.Port))
}
