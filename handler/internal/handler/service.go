package handler

import "github.com/arielkka/fallbox/handler/internal/models"

type Service interface {
	GetUser(login, password string) (string,error)
	CreateUser(login, password string) (string, error)

	GetUserPNG(userID, pngID string) ([]byte, error)
	AddUserPNG(userID string, png *models.PNG) (string, error)
	DeleteUserPNG(userID, pngID string) error

	GetUserJPG(userID, jpgID string) ([]byte, error)
	AddUserJPG(userID string, jpg *models.JPG) (string,error)
	DeleteUserJPG(userID, jpgID string) error
}
