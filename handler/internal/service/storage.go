package service

import (
	"github.com/arielkka/fallbox/handler/internal/storage"
	"github.com/jmoiron/sqlx"
)

type IUser interface {
	GetUser(login, password string) (string, error)
	CreateUser(login, password, id string) error
}

type Storage struct {
	IUser
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		storage.NewUser(db),
	}
}
