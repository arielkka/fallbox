package service

type Storage interface {
	GetUser(login, password string) (string, error)
	CreateUser(login, password, id string) error
}
