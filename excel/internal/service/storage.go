package service

import (
	"github.com/arielkka/fallbox/excel/internal/storage"
	"github.com/jmoiron/sqlx"
)

type IExcel interface {
	GetExcel(userID string, excelID int) ([]byte, error)
	AddExcel(userID string, excel []byte) (int, error)
	DeleteExcel(userID string, excelID int) error
}

type Storage struct {
	IExcel
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		storage.NewExcelStorage(db),
	}
}
