package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arielkka/fallbox/handler/internal/models"
	"github.com/labstack/echo/v4"
)

func (r *router) GetUserExcel(c echo.Context) error {
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

	err = r.service.GetUserExcel(cookieUserID.Value, req.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	log.Println("Get user excel finished")

	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("receive your excel by %v id :-)", req.ID),
	})
}

func (r *router) PostUserExcel(c echo.Context) error {
	cookieUserID, err := c.Cookie(r.cfg.Router.CookieUserID)
	if err != nil {
		return err
	}

	path := new(models.FilePath)

	err = json.NewDecoder(c.Request().Body).Decode(path)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	excelID, err := r.service.AddUserExcel(cookieUserID.Value, path.Path)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"excelID": excelID,
	})

}

func (r *router) DeleteUserExcel(c echo.Context) error {
	cookieUserID, err := c.Cookie(r.cfg.Router.CookieUserID)
	if err != nil {
		return err
	}

	req := new(models.Request)

	err = json.NewDecoder(c.Request().Body).Decode(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = r.service.DeleteUserExcel(cookieUserID.Value, req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("excel №%v was deleted", req.ID),
	})
}
