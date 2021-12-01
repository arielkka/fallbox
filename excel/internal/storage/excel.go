package storage

import (
	"github.com/jmoiron/sqlx"
)

type Excel struct {
	db *sqlx.DB
}

func NewExcelStorage(db *sqlx.DB) *Excel {
	return &Excel{db: db}
}

func (p *Excel) GetTxt(userID string, excelID int) ([]byte, error) {
	panic("implement me")

}

func (p *Excel) AddTxt(userID string, excel []byte) (int, error) {
	panic("implement me")

}

func (p *Excel) DeleteTxt(userID string, excelID int) error {

	panic("implement me")
}
