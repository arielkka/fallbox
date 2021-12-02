package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arielkka/fallbox/handler/internal/models"
	"github.com/labstack/echo/v4"
)

func (r *router) GetUserTxt(c echo.Context) error {
	log.Println("Get user excel started")

	req := new(models.Request)

	err := json.NewDecoder(c.Request().Body).Decode(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	cookieUserID, err := c.Cookie(r.cfg.Router.CookieUserID)
	if err != nil {
		return err
	}

	err = r.service.GetUserTxt(cookieUserID.Value, req.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	log.Println("Get user excel finished")

	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("receive your txt by %v id :-)", req.ID),
	})
}

func (r *router) PostUserTxt(c echo.Context) error {
	cookieUserID, err := c.Cookie(r.cfg.Router.CookieUserID)
	if err != nil {
		return err
	}

	path := new(models.FilePath)

	err = json.NewDecoder(c.Request().Body).Decode(path)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	txtID, err := r.service.AddUserTxt(cookieUserID.Value, path.Path)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"txtID": txtID,
	})
}

func (r *router) DeleteUserTxt(c echo.Context) error {
	cookieUserID, err := c.Cookie(r.cfg.Router.CookieUserID)
	if err != nil {
		return err
	}

	req := new(models.Request)

	err = json.NewDecoder(c.Request().Body).Decode(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = r.service.DeleteUserTxt(cookieUserID.Value, req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("txt file №%v was deleted", req.ID),
	})
}
