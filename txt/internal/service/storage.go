package service

import (
	"github.com/arielkka/fallbox/txt/internal/storage"
	"github.com/jmoiron/sqlx"
)

type ITxt interface {
	GetExcel(userID string, imgID int) ([]byte, error)
	AddExcel(userID string, img []byte) (int, error)
	DeleteExcel(userID string, imgID int) error
}

type Storage struct {
	ITxt
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		storage.NewTxtStorage(db),
	}
}
