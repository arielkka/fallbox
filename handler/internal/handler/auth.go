package handler

import (
	"encoding/json"
	"errors"
	"github.com/arielkka/fallbox/handler/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
	"time"
)

type jwtCustomClaims struct {
	ID    string `json:"id"`
	Login string `json:"name"`
	jwt.StandardClaims
}

func (r *router) SkipJWTMiddleware(c echo.Context) bool {
	if c.Path() == "/registration" || c.Path() == "/auth" {
		return true
	}
	return false
}

func (r *router) registration(c echo.Context) error {
	user := new(models.User)

	err := json.NewDecoder(c.Request().Body).Decode(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = r.service.CreateUser(user.Login, user.Password)
	if err != nil {
		if !strings.Contains(err.Error(), "Error 1062") {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "user was created",
	})
}

func (r *router) auth(c echo.Context) error {
	user := new(models.User)

	err := json.NewDecoder(c.Request().Body).Decode(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	id, err := r.service.GetUser(user.Login, user.Password)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	claims := &jwtCustomClaims{
		id,
		user.Login,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		return c.JSON(http.StatusBadRequest, errors.New("couldn't find jwt key"))
	}
	t, err := token.SignedString([]byte(key))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	cookie := newCookie()
	writeCookie(c, cookie, r.cfg.Router.CookieToken, t, time.Now().Add(15*time.Minute))
	writeCookie(c, cookie, r.cfg.Router.CookieUserID, id, time.Now().Add(15*time.Minute))
	return c.JSON(http.StatusOK, echo.Map{
		"message": "authorization passed",
	})
}
