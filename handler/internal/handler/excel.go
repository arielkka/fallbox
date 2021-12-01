package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arielkka/fallbox/handler/internal/models"
	"github.com/labstack/echo/v4"
)

func (r *router) GetUserExcel(c echo.Context) error {
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

	err = r.service.GetUserExcel(userID.ID, req.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("receive your excel by %v id :-)", req.ID),
	})
}

func (r *router) PostUserExcel(c echo.Context) error {
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

	excelID, err := r.service.AddUserExcel(userID.ID, req.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"excelID": excelID,
	})

}

func (r *router) DeleteUserExcel(c echo.Context) error {
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

	err = r.service.DeleteUserExcel(userID.ID, req.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("excel №%v was deleted", req.ID),
	})
}
