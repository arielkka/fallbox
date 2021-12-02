package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func newCookie() *http.Cookie {
	return new(http.Cookie)
}

func writeCookie(c echo.Context, cookie *http.Cookie, key, value string, expiresTime time.Time) {
	cookie.Name = key
	cookie.Value = value
	cookie.Expires = expiresTime
	c.SetCookie(cookie)
}

func readCookie(c echo.Context, key string) (string, error) {
	cookie, err := c.Cookie(key)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
