package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/arielkka/fallbox/handler/internal/models"
	"github.com/labstack/echo/v4"
	"image"
	"image/png"
	"net/http"
	"os"
)

func convertToPNG(body []byte, id string) error{
	img, _, err := image.Decode(bytes.NewReader(body))
	if err != nil {
		return err
	}
	out, err := os.Create(fmt.Sprintf("./%s.png", id))
	if err != nil {
		return err
	}
	err = png.Encode(out, img)
	if err != nil {
		return err
	}
	return nil
}

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

	err = convertToPNG(body, p.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"err": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": fmt.Sprintf("receive your image by %s id :-)", p.ID),
	})

}

func (r *router) PostUserPNG(context echo.Context) error {
	return nil

}

func (r *router) DeleteUserPNG(context echo.Context) error {
	return nil

}
