package handler

type Service interface {
	GetUser(login, password string) (string, error)
	CreateUser(login, password string) error

	GetUserTxt(userID string, txtID int) error
	AddUserTxt(userID string, path string) (int, error)
	DeleteUserTxt(userID string, txtID int) error

	GetUserExcel(userID string, excelID int) error
	AddUserExcel(userID string, path string) (int, error)
	DeleteUserExcel(userID string, jpgID int) error
}
