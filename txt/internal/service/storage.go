package service

import (
	"github.com/arielkka/fallbox/txt/internal/storage"
	"github.com/jmoiron/sqlx"
)

type ITxt interface {
	GetTxt(userID string, imgID int) ([]byte, error)
	AddTxt(userID string, img []byte) (int, error)
	DeleteTxt(userID string, imgID int) error
}

type Storage struct {
	ITxt
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		storage.NewTxtStorage(db),
	}
}
