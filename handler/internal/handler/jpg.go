package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/arielkka/fallbox/handler/internal/models"
	"github.com/labstack/echo/v4"
)

func (r *router) GetUserJPG(c echo.Context) error {
	j := new(models.JPG)

	err := json.NewDecoder(c.Request().Body).Decode(j)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userID := new(models.UserID)

	err = json.NewDecoder(c.Request().Body).Decode(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	body, err := r.service.GetUserJPG(userID.ID, j.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	err = convertBytesToJPG(body, j.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("receive your image by %s id :-)", j.ID),
	})
}

func (r *router) PostUserJPG(c echo.Context) error {
	userID := new(models.UserID)

	err := json.NewDecoder(c.Request().Body).Decode(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	path := new(models.ImagePath)

	err = json.NewDecoder(c.Request().Body).Decode(path)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	image, err := getImage(path.Path)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	bytes, err := convertPNGToBytes(image)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	jpgBytes := new(models.JPG)
	jpgBytes.Body = bytes

	jpgID, err := r.service.AddUserJPG(userID.ID, jpgBytes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"jpgID": jpgID,
	})
}

func (r *router) DeleteUserJPG(c echo.Context) error {
	userID := new(models.UserID)

	err := json.NewDecoder(c.Request().Body).Decode(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	jpgID := new(models.JpgID)

	err = json.NewDecoder(c.Request().Body).Decode(jpgID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = r.service.DeleteUserJPG(userID.ID, jpgID.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("jpg №%s was deleted", jpgID),
	})
}
