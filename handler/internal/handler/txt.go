package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arielkka/fallbox/handler/internal/models"
	"github.com/labstack/echo/v4"
)

func (r *router) GetUserTxt(c echo.Context) error {
	req := new(models.Request)

	err := json.NewDecoder(c.Request().Body).Decode(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userID := new(models.UserID)

	err = json.NewDecoder(c.Request().Body).Decode(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = r.service.GetUserTxt(userID.ID, req.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("receive your txt by %v id :-)", req.ID),
	})
}

func (r *router) PostUserTxt(c echo.Context) error {
	userID := new(models.UserID)

	err := json.NewDecoder(c.Request().Body).Decode(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	req := new(models.Request)

	err = json.NewDecoder(c.Request().Body).Decode(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	txtID, err := r.service.AddUserTxt(userID.ID, req.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"txtID": txtID,
	})
}

func (r *router) DeleteUserTxt(c echo.Context) error {
	userID := new(models.UserID)

	err := json.NewDecoder(c.Request().Body).Decode(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	req := new(models.Request)

	err = json.NewDecoder(c.Request().Body).Decode(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = r.service.DeleteUserTxt(userID.ID, req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("txt №%v was deleted", req.ID),
	})
}
