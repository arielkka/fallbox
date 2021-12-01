package handler

import (
	"encoding/json"
	"github.com/arielkka/fallbox/handler/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

type jwtCustomClaims struct {
	ID    string `json:"id"`
	Login string `json:"name"`
	jwt.StandardClaims
}

func (r *router) registration(c echo.Context) error {
	user := new(models.User)

	err := json.NewDecoder(c.Request().Body).Decode(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	id, err := r.service.CreateUser(user.Login, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = c.JSON(http.StatusCreated, echo.Map{
		"id": id,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *router) auth(c echo.Context) error {
	user := new(models.User)

	err := json.NewDecoder(c.Request().Body).Decode(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	id, err := r.service.GetUser(user.Login, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	claims := &jwtCustomClaims{
		id,
		user.Login,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
