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

func getImage(path string) (image.Image, error) {
	reader, _ :=os.Open(path)
	defer reader.Close()

	return nil,nil
}

func convertToJPG(body []byte, id string) error{
	img, _, err := image.Decode(bytes.NewReader(body))
	if err != nil {
		return err
	}
	out, err := os.Create(fmt.Sprintf("./%s.jpg", id))
	if err != nil {
		return err
	}
	err = png.Encode(out, img)
	if err != nil {
		return err
	}
	return nil
}



func (r *router) GetUserJPG(c echo.Context) error {
	j := new(models.JPG)

	err := json.NewDecoder(c.Request().Body).Decode(j)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	userID := new(models.UserID)

	err = json.NewDecoder(c.Request().Body).Decode(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	body, err := r.service.GetUserJPG(userID.ID, j.ID)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	err = convertToJPG(body, j.ID)
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
	return nil

}

func (r *router) DeleteUserJPG(c echo.Context) error {
	return nil

}
