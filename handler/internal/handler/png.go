package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/arielkka/fallbox/handler/internal/models"
	myerrors "github.com/arielkka/fallbox/handler/pkg/errors"
	"github.com/labstack/echo/v4"
)

func (r *router) GetUserPNG(c echo.Context) error {
	p := new(models.PNG)

	err := json.NewDecoder(c.Request().Body).Decode(p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	userID := new(models.UserID)

	err = json.NewDecoder(c.Request().Body).Decode(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	body, err := r.service.GetUserJPG(userID.ID, p.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	err = convertBytesToPNG(body, p.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("receive your image by %s id :-)", p.ID),
	})

}

func (r *router) PostUserPNG(c echo.Context) error {
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

	pngBytes := new(models.PNG)
	pngBytes.Body = bytes

	pngID, err := r.service.AddUserPNG(userID.ID, pngBytes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"pngID": pngID,
	})

}

func (r *router) DeleteUserPNG(c echo.Context) error {
	userID := new(models.UserID)

	err := json.NewDecoder(c.Request().Body).Decode(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	pngID := new(models.PngID)

	err = json.NewDecoder(c.Request().Body).Decode(pngID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err = r.service.DeleteUserPNG(userID.ID, pngID.ID)
	if errors.Is(err, myerrors.NotFound()) {
		return c.JSON(http.StatusNotFound, err)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("png №%s was deleted", pngID),
	})
}
