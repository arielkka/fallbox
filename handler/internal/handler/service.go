package handler

type Service interface {
	GetUser(login, password string) error

	GetAllUserPNG(userID string) ([]byte, error)
	GetUserPNG(userID, pngID string) ([]byte, error)
	AddUserPNG(userID string) (string, error)
	DeleteUserPNG(userID, pngID string) error

	GetAllUserJPG(userID string) ([]byte, error)
	GetUserJPG(userID, pngID string) ([]byte, error)
	AddUserJPG(userID string) (string, error)
	DeleteUserJPG(userID, pngID string) error
}
